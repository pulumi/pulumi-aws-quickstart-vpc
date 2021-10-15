// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package provider

import (
	"encoding/json"
	"fmt"

	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/cloudwatch"
	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/ec2"
	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/iam"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AvailabilityZone struct {
	AvailabilityZone   string `pulumi:"availabilityZone"`
	PublicSubnetCidr   string `pulumi:"publicSubnetCidr"`
	PrivateSubnetACidr string `pulumi:"privateSubnetACidr"`
	PrivateSubnetBCidr string `pulumi:"privateSubnetBCidr"`
}

type VpcArgs struct {
	CidrBlock                      string             `pulumi:"cidrBlock"`
	InstanceTenancy                string             `pulumi:"instanceTenancy"`
	EnableDnsHostnames             *bool              `pulumi:"enableDnsHostnames"`
	EnableDnsSupport               *bool              `pulumi:"enableDnsSupport"`
	CreateFlowLogs                 *bool              `pulumi:"createFlowLogs"`
	FlowLogsRetentionPeriodInDays  int                `pulumi:"flowLogsRetentionPeriodInDays"`
	FlowLogsMaxAggregationInterval int                `pulumi:"flowLogsMaxAggregationInterval"`
	FlowLogsLogFormat              string             `pulumi:"flowLogsLogFormat"`
	FlowLogsTrafficType            string             `pulumi:"flowLogsTrafficType"`
	AvailabilityZoneConfig         []AvailabilityZone `pulumi:"availabilityZoneConfig"`
	CreatePrivateSubnets           *bool              `pulumi:"createPrivateSubnets"`
	CreatePublicSubnets            *bool              `pulumi:"createPublicSubnets"`
	CreateNatGateways              *bool              `pulumi:"createNatGateways"`
	CreateAdditionalPrivateSubnets *bool              `pulumi:"createAdditionalPrivateSubnets"`
}

type Vpc struct {
	pulumi.ResourceState

	VpcID            pulumi.StringOutput      `pulumi:"vpcID"`
	PrivateSubnetIDs pulumi.StringArrayOutput `pulumi:"privateSubnetIDs"`
	PublicSubnetIDs  pulumi.StringArrayOutput `pulumi:"publicSubnetIDs"`
	NatGatewayIPs    pulumi.StringArrayOutput `pulumi:"natGatewayIPs"`
}

func NewVpc(ctx *pulumi.Context,
	name string, args *VpcArgs, opts ...pulumi.ResourceOption) (*Vpc, error) {
	if args == nil {
		args = &VpcArgs{}
	}

	component := &Vpc{}
	err := ctx.RegisterComponentResource("aws-quickstart-vpc:index:Vpc", name, component, opts...)
	if err != nil {
		return nil, err
	}

	instanceTenancy := "default"
	if args.InstanceTenancy != "" {
		instanceTenancy = args.InstanceTenancy
	}

	vpcEnableDnsHostnames := true
	if args.EnableDnsHostnames != nil {
		vpcEnableDnsHostnames = *args.EnableDnsHostnames
	}

	vpcEnableDnsSupport := true
	if args.EnableDnsSupport != nil {
		vpcEnableDnsSupport = *args.EnableDnsSupport
	}

	vpc, vpcErr := ec2.NewVpc(ctx, fmt.Sprintf("%s-vpc", name), &ec2.VpcArgs{
		CidrBlock:          pulumi.String(args.CidrBlock),
		InstanceTenancy:    pulumi.String(instanceTenancy),
		EnableDnsHostnames: pulumi.Bool(vpcEnableDnsHostnames),
		EnableDnsSupport:   pulumi.Bool(vpcEnableDnsSupport),
	}, pulumi.Parent(component))
	if vpcErr != nil {
		return nil, vpcErr
	}

	createFlowLogs := true
	if args.CreateFlowLogs != nil {
		createFlowLogs = *args.CreateFlowLogs
	}

	if createFlowLogs {
		flowLogsRetentionPeriodInDays := 14
		if args.FlowLogsRetentionPeriodInDays > 0 {
			flowLogsRetentionPeriodInDays = args.FlowLogsRetentionPeriodInDays
		}
		vpcFlowLogGroup, vpcFlowLogGroupErr := cloudwatch.NewLogGroup(ctx, fmt.Sprintf("%s-flow-logs", name), &cloudwatch.LogGroupArgs{
			RetentionInDays: pulumi.Int(flowLogsRetentionPeriodInDays),
		}, pulumi.Parent(component))
		if vpcFlowLogGroupErr != nil {
			return nil, vpcFlowLogGroupErr
		}

		assumeRolePolicyString, assumeRolePolicyStringErr := json.Marshal(
			map[string]interface{}{
				"Version": "2012-10-17",
				"Statement": []map[string]interface{}{
					{
						"Action": "sts:AssumeRole",
						"Effect": "Allow",
						"Sid":    "",
						"Principal": map[string]interface{}{
							"Service": "vpc-flow-logs.amazonaws.com",
						},
					},
				},
			},
		)
		if assumeRolePolicyStringErr != nil {
			return nil, assumeRolePolicyStringErr
		}

		vpcFlowLogRole, vpcFlowLogRoleErr := iam.NewRole(ctx, fmt.Sprintf("%s-vpc-flow-log-role", name), &iam.RoleArgs{
			AssumeRolePolicy: pulumi.String(assumeRolePolicyString),
		}, pulumi.Parent(component))
		if vpcFlowLogRoleErr != nil {
			return nil, vpcFlowLogRoleErr
		}

		policyStatement, policyStatementErr := json.Marshal(
			map[string]interface{}{
				"Version": "2012-10-17",
				"Statement": []map[string]interface{}{
					{
						"Action": []string{
							"logs:CreateLogStream",
							"logs:PutLogEvents",
							"logs:DescribeLogGroups",
							"logs:DescribeLogStreams",
						},
						"Effect":   "Allow",
						"Resource": "*",
					},
				},
			},
		)
		if policyStatementErr != nil {
			return nil, policyStatementErr
		}

		_, vpcFlowLogRolePolicyErr := iam.NewRolePolicy(ctx, fmt.Sprintf("%s-vpc-flow-log-policy", name), &iam.RolePolicyArgs{
			Role:   vpcFlowLogRole,
			Policy: pulumi.String(policyStatement),
		}, pulumi.Parent(component))
		if vpcFlowLogRolePolicyErr != nil {
			return nil, vpcFlowLogRolePolicyErr
		}

		flowLogsTrafficType := "REJECT"
		if args.FlowLogsTrafficType != "" {
			flowLogsTrafficType = args.FlowLogsTrafficType
		}

		flowLogsLogFormat := "${version} ${account-id} ${interface-id} ${srcaddr} ${dstaddr} ${srcport} ${dstport} ${protocol} ${packets} ${bytes} ${start} ${end} ${action} ${log-status}"
		if args.FlowLogsLogFormat != "" {
			flowLogsLogFormat = args.FlowLogsLogFormat
		}

		flowLogsMaxAggregationInterval := 600
		if args.FlowLogsMaxAggregationInterval > 0 {
			flowLogsMaxAggregationInterval = args.FlowLogsMaxAggregationInterval
		}

		_, vpcFlowLogErr := ec2.NewFlowLog(ctx, fmt.Sprintf("%s-vpc-flow-log", name), &ec2.FlowLogArgs{
			IamRoleArn:             vpcFlowLogRole.Arn,
			LogDestination:         vpcFlowLogGroup.Arn,
			TrafficType:            pulumi.String(flowLogsTrafficType),
			VpcId:                  vpc.ID(),
			LogFormat:              pulumi.String(flowLogsLogFormat),
			MaxAggregationInterval: pulumi.Int(flowLogsMaxAggregationInterval),
		}, pulumi.Parent(component))
		if vpcFlowLogErr != nil {
			return nil, vpcFlowLogErr
		}

	}

	internetGateway, internetGatewayErr := ec2.NewInternetGateway(ctx, fmt.Sprintf("%s-internet-gateway", name), &ec2.InternetGatewayArgs{
		VpcId: vpc.ID(),
	}, pulumi.Parent(component))
	if internetGatewayErr != nil {
		return nil, internetGatewayErr
	}

	createPublicSubnets := true
	if args.CreatePublicSubnets != nil {
		createPublicSubnets = *args.CreatePublicSubnets
	}
	createPrivateSubnets := true
	if args.CreatePrivateSubnets != nil {
		createPrivateSubnets = *args.CreatePrivateSubnets
	}
	createNatGateways := true
	if args.CreateNatGateways != nil {
		createNatGateways = *args.CreateNatGateways
	}
	createAdditionalPrivateSubnets := true
	if args.CreateAdditionalPrivateSubnets != nil {
		createAdditionalPrivateSubnets = *args.CreateAdditionalPrivateSubnets
	}

	var privateSubnetIds []pulumi.StringOutput
	var publicSubnetIds []pulumi.StringOutput
	var natGatewayIPs []pulumi.StringOutput
	for i, az := range args.AvailabilityZoneConfig {
		var natGateway *ec2.NatGateway
		if createPublicSubnets && az.PublicSubnetCidr != "" && az.AvailabilityZone != "" {
			publicSubnet, publicSubnetErr := ec2.NewSubnet(ctx, fmt.Sprintf("%s-public-subnet-%d", name, i), &ec2.SubnetArgs{
				VpcId:            vpc.ID(),
				CidrBlock:        pulumi.String(az.PublicSubnetCidr),
				AvailabilityZone: pulumi.String(az.AvailabilityZone),
			}, pulumi.Parent(component))
			if publicSubnetErr != nil {
				return nil, publicSubnetErr
			}

			publicRouteTable, publicRouteTableErr := ec2.NewRouteTable(ctx, fmt.Sprintf("%s-public-route-table-%d", name, i), &ec2.RouteTableArgs{
				VpcId: vpc.ID(),
			}, pulumi.Parent(component))
			if publicRouteTableErr != nil {
				return nil, publicRouteTableErr
			}

			_, publicRouteErr := ec2.NewRoute(ctx, fmt.Sprintf("%s-public-route-%d", name, i), &ec2.RouteArgs{
				RouteTableId:         publicRouteTable.ID(),
				DestinationCidrBlock: pulumi.String("0.0.0.0/0"),
				GatewayId:            internetGateway.ID(),
			}, pulumi.Parent(component))
			if publicRouteErr != nil {
				return nil, publicRouteErr
			}

			_, publicRouteTableAssociationErr := ec2.NewRouteTableAssociation(ctx, fmt.Sprintf("%s-public-route-table-association-%d", name, i), &ec2.RouteTableAssociationArgs{
				SubnetId:     publicSubnet.ID(),
				RouteTableId: publicRouteTable.ID(),
			}, pulumi.Parent(component))
			if publicRouteTableAssociationErr != nil {
				return nil, publicRouteErr
			}

			if createNatGateways {
				elasticIp, elasticIpErr := ec2.NewEip(ctx, fmt.Sprintf("%s-elastic-ip-%d", name, i), &ec2.EipArgs{
					Vpc: pulumi.Bool(true),
				}, pulumi.Parent(component))
				if elasticIpErr != nil {
					return nil, elasticIpErr
				}

				natGateway, natGatewayErr := ec2.NewNatGateway(ctx, fmt.Sprintf("%s-nat-gateway-%d", name, i), &ec2.NatGatewayArgs{
					AllocationId: elasticIp.ID(),
					SubnetId:     publicSubnet.ID(),
				}, pulumi.Parent(component))
				if natGatewayErr != nil {
					return nil, elasticIpErr
				}

				natGateway = natGateway
				natGatewayIPs = append(natGatewayIPs, elasticIp.PublicIp)
			}

			publicSubnetIds = append(publicSubnetIds, publicSubnet.ID().ToStringOutput())
		}

		if createPrivateSubnets && az.AvailabilityZone != "" && az.PrivateSubnetACidr != "" {
			privateSubnet, privateSubnetErr := ec2.NewSubnet(ctx, fmt.Sprintf("%s-private-subnet-a-%d", name, i), &ec2.SubnetArgs{
				VpcId:            vpc.ID(),
				CidrBlock:        pulumi.String(az.PrivateSubnetACidr),
				AvailabilityZone: pulumi.String(az.AvailabilityZone),
			}, pulumi.Parent(component))
			if privateSubnetErr != nil {
				return nil, privateSubnetErr
			}

			privateRouteTable, privateRouteTableErr := ec2.NewRouteTable(ctx, fmt.Sprintf("%s-private-route-table-%d", name, i), &ec2.RouteTableArgs{
				VpcId: vpc.ID(),
			}, pulumi.Parent(component))
			if privateRouteTableErr != nil {
				return nil, privateRouteTableErr
			}

			_, privateRouteErr := ec2.NewRoute(ctx, fmt.Sprintf("%s-private-route-%d", name, i), &ec2.RouteArgs{
				RouteTableId:         privateRouteTable.ID(),
				DestinationCidrBlock: pulumi.String("0.0.0.0/0"),
				NatGatewayId:         natGateway.ID(),
			}, pulumi.Parent(component))
			if privateRouteErr != nil {
				return nil, privateRouteErr
			}

			_, privateRouteTableAssociationErr := ec2.NewRouteTableAssociation(ctx, fmt.Sprintf("%s-private-route-table-association-%d", name, i), &ec2.RouteTableAssociationArgs{
				SubnetId:     privateSubnet.ID(),
				RouteTableId: privateRouteTable.ID(),
			}, pulumi.Parent(component))
			if privateRouteTableAssociationErr != nil {
				return nil, privateRouteTableAssociationErr
			}

			privateSubnetIds = append(publicSubnetIds, privateSubnet.ID().ToStringOutput())
		}

		if createAdditionalPrivateSubnets && az.AvailabilityZone != "" && az.PrivateSubnetBCidr != "" {
			privateSubnet, privateSubnetErr := ec2.NewSubnet(ctx, fmt.Sprintf("%s-private-subnet-b-%d", name, i), &ec2.SubnetArgs{
				VpcId:            vpc.ID(),
				CidrBlock:        pulumi.String(az.PrivateSubnetACidr),
				AvailabilityZone: pulumi.String(az.AvailabilityZone),
			}, pulumi.Parent(component))
			if privateSubnetErr != nil {
				return nil, privateSubnetErr
			}

			_, privateNetworkAclErr := ec2.NewNetworkAcl(ctx, fmt.Sprintf("%s-private-network-acl-%d", name, i), &ec2.NetworkAclArgs{
				VpcId:     vpc.ID(),
				SubnetIds: pulumi.StringArray{privateSubnet.ID()},
				Egress: ec2.NetworkAclEgressArray{
					ec2.NetworkAclEgressArgs{
						Action:    pulumi.String("allow"),
						CidrBlock: pulumi.String("0.0.0.0/0"),
						RuleNo:    pulumi.Int(100),
						Protocol:  pulumi.String("all"),
						FromPort:  pulumi.Int(0),
						ToPort:    pulumi.Int(0),
					},
				},
				Ingress: ec2.NetworkAclIngressArray{
					ec2.NetworkAclIngressArgs{
						Action:    pulumi.String("allow"),
						CidrBlock: pulumi.String("0.0.0.0/0"),
						RuleNo:    pulumi.Int(100),
						Protocol:  pulumi.String("all"),
						FromPort:  pulumi.Int(0),
						ToPort:    pulumi.Int(0),
					},
				},
			}, pulumi.Parent(component))
			if privateNetworkAclErr != nil {
				return nil, privateNetworkAclErr
			}
			privateSubnetIds = append(publicSubnetIds, privateSubnet.ID().ToStringOutput())
		}
	}

	component.VpcID = vpc.ID().ToStringOutput()
	component.PrivateSubnetIDs = pulumi.ToStringArrayOutput(privateSubnetIds)
	component.PublicSubnetIDs = pulumi.ToStringArrayOutput(publicSubnetIds)
	component.NatGatewayIPs = pulumi.ToStringArrayOutput(natGatewayIPs)

	if err := ctx.RegisterResourceOutputs(component, pulumi.Map{
		"vpcID":            component.VpcID,
		"privateSubnetIDs": component.PrivateSubnetIDs,
		"publicSubnetIDs":  component.PublicSubnetIDs,
		"natGatewayIPs":    component.NatGatewayIPs,
	}); err != nil {
		return nil, err
	}

	return component, nil
}
