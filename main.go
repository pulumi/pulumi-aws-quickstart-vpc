package main

import (
	"fmt"

	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/ec2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AvailabilityZoneConfig struct {
	publicSubnetName   string
	availabilityZone   string
	publicSubnetCidr   string
	privateSubnetAName string
	privateSubnetACidr string
	privateSubnetBName string
	privateSubnetBCidr string
}

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {

		/************************************
		 * Namespace parameters
		 ************************************/
		vpcName := "Pulumi VPC"
		namespace := "pulumi-vpc"

		/************************************
		 * Network parameters
		 ************************************/

		// default || dedicated
		vpcInstanceTenancy := "default"

		createPublicSubnets := true
		createPrivateSubnets := true
		createAdditionalPrivateSubnets := false
		createNatGateways := false

		// VPC block
		vpcCidrBlock := "10.0.0.0/16"

		availabilityZoneConfigs := [4]AvailabilityZoneConfig{
			{
				publicSubnetName:   "Pulumi Public Subnet 1",
				availabilityZone:   "us-east-1a",
				publicSubnetCidr:   "10.0.5.0/24",
				privateSubnetAName: "Pulumi Private Subnet 1A",
				privateSubnetACidr: "10.0.7.0/24",
				privateSubnetBName: "Pulumi Private Subnet 1B",
				privateSubnetBCidr: "10.0.224.0/21",
			},
			{
				publicSubnetName:   "Pulumi Public Subnet 2",
				availabilityZone:   "us-east-1b",
				publicSubnetCidr:   "10.0.8.0/24",
				privateSubnetAName: "Pulumi Private Subnet 2A",
				privateSubnetACidr: "10.0.9.0/24",
				privateSubnetBName: "Pulumi Private Subnet 2B",
				privateSubnetBCidr: "10.0.232.0/21",
			},
			{
				publicSubnetName:   "Pulumi Public Subnet 3",
				availabilityZone:   "us-east-1c",
				publicSubnetCidr:   "10.0.10.0/24",
				privateSubnetAName: "Pulumi Private Subnet 3A",
				privateSubnetACidr: "10.0.11.0/24",
				privateSubnetBName: "Pulumi Private Subnet 3B",
				privateSubnetBCidr: "10.0.240.0/21",
			},
			{
				publicSubnetName:   "Pulumi Public Subnet 4",
				availabilityZone:   "us-east-1d",
				publicSubnetCidr:   "10.0.12.0/24",
				privateSubnetAName: "Pulumi Private Subnet 4A",
				privateSubnetACidr: "10.0.13.0/24",
				privateSubnetBName: "Pulumi Private Subnet 4B",
				privateSubnetBCidr: "10.0.216.0/21",
			},
		}

		// END PARAMETERS

		// @fixme- the quickstart has a dhcp options, is that necessary?
		// https://github.com/aws-quickstart/quickstart-aws-vpc/blob/ffc7af4e59a09dbf52199a3ecf70f3805abeff48/templates/aws-vpc.template.yaml#L457

		vpc, vpcErr := ec2.NewVpc(ctx, "vpc", &ec2.VpcArgs{
			CidrBlock:          pulumi.String(vpcCidrBlock),
			InstanceTenancy:    pulumi.String(vpcInstanceTenancy),
			EnableDnsHostnames: pulumi.Bool(true),
			EnableDnsSupport:   pulumi.Bool(true),
			Tags: pulumi.StringMap{
				"Name": pulumi.String(vpcName),
			},
		})

		if vpcErr != nil {
			return vpcErr
		}

		// @fixme - only create the internet gateway if public subnets are allowed.
		internetGatewayName := "Pulumi Internet Gateway"
		internetGateway, internetGatewayErr := ec2.NewInternetGateway(ctx, resourceName(namespace, "internet-gateway"), &ec2.InternetGatewayArgs{
			VpcId: vpc.ID(),
			Tags: pulumi.StringMap{
				"Name": pulumi.String(internetGatewayName),
			},
		})

		if internetGatewayErr != nil {
			return internetGatewayErr
		}

		for idx, s := range availabilityZoneConfigs {
			createPublicPrivateSubnets(ctx, fmt.Sprintf("%s%s%d", namespace, "-az", idx), vpc, internetGateway, s, createPublicSubnets, createPrivateSubnets, createAdditionalPrivateSubnets, createNatGateways)
		}

		return nil
	})
}

func resourceId(namespace string, resourceName string) string {
	return namespace + "-" + resourceName
}

func resourceName(namespace string, resourceName string) string {
	return namespace + " | " + resourceName
}

// @fixme - this function includes a "DependsOn block" for the internet gateway
func createPublicSubnet(ctx *pulumi.Context, namespace string, subnetCidr string, subnetName string, availabilityZone string, vpc *ec2.Vpc, internetGateway *ec2.InternetGateway, createNatGateways bool) (*ec2.NatGateway, error) {
	natGatewayName := resourceName(namespace, "Pulumi NAT Gateway")
	publicRouteTableName := resourceName(namespace, "Pulumi Public Route Table")

	entireInternetCidr := "0.0.0.0/0"

	publicSubnet, publicSubnetErr := ec2.NewSubnet(ctx, resourceId(namespace, "public-subnet"), &ec2.SubnetArgs{
		VpcId:            vpc.ID(),
		CidrBlock:        pulumi.String(subnetCidr),
		AvailabilityZone: pulumi.String(availabilityZone),
		Tags: pulumi.StringMap{
			"Name": pulumi.String(subnetName),
		},
	})

	if publicSubnetErr != nil {
		return nil, publicSubnetErr
	}

	publicRouteTable, publicRouteTableErr := ec2.NewRouteTable(ctx, resourceId(namespace, "public-route-table"), &ec2.RouteTableArgs{
		VpcId: vpc.ID(),
		Tags: pulumi.StringMap{
			"Name": pulumi.String(publicRouteTableName),
		},
	})

	if publicRouteTableErr != nil {
		return nil, publicRouteTableErr
	}

	_, publicRouteErr := ec2.NewRoute(ctx, resourceId(namespace, "public-route"), &ec2.RouteArgs{
		RouteTableId:         publicRouteTable.ID(),
		DestinationCidrBlock: pulumi.String(entireInternetCidr),
		GatewayId:            internetGateway.ID(),
	})

	if publicRouteErr != nil {
		return nil, publicRouteErr
	}

	_, publicRouteTableAssociationErr := ec2.NewRouteTableAssociation(ctx, resourceId(namespace, "public-route-table-association"), &ec2.RouteTableAssociationArgs{
		SubnetId:     publicSubnet.ID(),
		RouteTableId: publicRouteTable.ID(),
	})

	if publicRouteTableAssociationErr != nil {
		return nil, publicRouteErr
	}

	// @fixme - this could be moved to a separate function "createNatGateway(subnetId)"
	if createNatGateways {
		// NAT Gateway + Elastic IP
		elasticIp, elasticIpErr := ec2.NewEip(ctx, resourceId(namespace, "elastic-ip"), &ec2.EipArgs{
			Vpc: pulumi.Bool(true),
		})

		if elasticIpErr != nil {
			return nil, elasticIpErr
		}

		// @fixme - return natGateway for the private subnet
		natGateway, natGatewayErr := ec2.NewNatGateway(ctx, resourceId(namespace, "nat-gateway"), &ec2.NatGatewayArgs{
			AllocationId: elasticIp.ID(),
			SubnetId:     publicSubnet.ID(),
			Tags: pulumi.StringMap{
				"Name": pulumi.String(natGatewayName),
			},
			// Explicitly declare that this NAT Gateway depends on the Internet Gateway's deployment to be completed
		}, pulumi.DependsOn([]pulumi.Resource{
			internetGateway,
		}))

		if natGatewayErr != nil {
			return nil, elasticIpErr
		}

		return natGateway, nil
	}

	return nil, nil
}

func createPrivateSubnet(ctx *pulumi.Context, namespace string, subnetCidr string, subnetName string, availabilityZone string, vpc *ec2.Vpc, natGateway *ec2.NatGateway, withNetworkAcl bool) error {
	entireInternetCidr := "0.0.0.0/0"

	privateSubnet, privateSubnetErr := ec2.NewSubnet(ctx, resourceId(namespace, "private-subnet"), &ec2.SubnetArgs{
		VpcId:            vpc.ID(),
		CidrBlock:        pulumi.String(subnetCidr),
		AvailabilityZone: pulumi.String(availabilityZone),
		Tags: pulumi.StringMap{
			"Name": pulumi.String(subnetName),
		},
	})

	if privateSubnetErr != nil {
		return privateSubnetErr
	}

	if nil != natGateway {
		privateRouteTableName := resourceName(namespace, "Pulumi Private Route Table")
		privateRouteTable, privateRouteTableErr := ec2.NewRouteTable(ctx, resourceId(namespace, "private-route-table"), &ec2.RouteTableArgs{
			VpcId: vpc.ID(),
			Tags: pulumi.StringMap{
				"Name": pulumi.String(privateRouteTableName),
			},
		})

		if privateRouteTableErr != nil {
			return privateRouteTableErr
		}

		_, privateRouteErr := ec2.NewRoute(ctx, resourceId(namespace, "private-route"), &ec2.RouteArgs{
			RouteTableId:         privateRouteTable.ID(),
			DestinationCidrBlock: pulumi.String(entireInternetCidr),
			NatGatewayId:         natGateway.ID(),
		})

		if privateRouteErr != nil {
			return privateRouteErr
		}

		_, privateRouteTableAssociationErr := ec2.NewRouteTableAssociation(ctx, resourceId(namespace, "private-route-table-association"), &ec2.RouteTableAssociationArgs{
			SubnetId:     privateSubnet.ID(),
			RouteTableId: privateRouteTable.ID(),
		})

		if privateRouteTableAssociationErr != nil {
			return privateRouteTableAssociationErr
		}
	}

	/***********************************************************
	 * Begin Private Network ACL
	 ***********************************************************/
	if withNetworkAcl {
		privateNetworkAclName := resourceName(namespace, "Pulumi Private Network ACL")
		_, privateNetworkAclErr := ec2.NewNetworkAcl(ctx, resourceId(namespace, "private-network-acl"), &ec2.NetworkAclArgs{
			VpcId:     vpc.ID(),
			SubnetIds: pulumi.StringArray{privateSubnet.ID()},
			Egress: ec2.NetworkAclEgressArray{
				ec2.NetworkAclEgressArgs{
					Action:    pulumi.String("allow"),
					CidrBlock: pulumi.String(entireInternetCidr),
					RuleNo:    pulumi.Int(100),

					// Note: Protocol "all" ignores the FromPort and ToPort fields. Using "0" as placeholder.
					Protocol: pulumi.String("all"),
					FromPort: pulumi.Int(0),
					ToPort:   pulumi.Int(0),
				},
			},
			Ingress: ec2.NetworkAclIngressArray{
				ec2.NetworkAclIngressArgs{
					Action:    pulumi.String("allow"),
					CidrBlock: pulumi.String(entireInternetCidr),
					RuleNo:    pulumi.Int(100),

					// Note: Protocol "all" ignores the FromPort and ToPort fields. Using "0" as placeholder.
					Protocol: pulumi.String("all"),
					FromPort: pulumi.Int(0),
					ToPort:   pulumi.Int(0),
				},
			},
			Tags: pulumi.StringMap{
				"Name":    pulumi.String(privateNetworkAclName),
				"Network": pulumi.String("NACL Protected"),
			},
		})

		if privateNetworkAclErr != nil {
			return privateNetworkAclErr
		}
	}

	return nil
}

// func createPublicPrivateSubnets(ctx *pulumi.Context, namespace string, vpc *ec2.Vpc, availabilityZone string, publicSubnetCidr string, publicSubnetName string, internetGateway *ec2.InternetGateway, privateSubnetCidr string, privateSubnetName string, privateSubnetBCidr string, privateSubnetBName string, withPublicSubnet bool, withPrivateSubnet bool, withAdditionalPrivateSubnet bool, createNatGateways bool) error {
func createPublicPrivateSubnets(ctx *pulumi.Context, namespace string, vpc *ec2.Vpc, internetGateway *ec2.InternetGateway, config AvailabilityZoneConfig, withPublicSubnet bool, withPrivateSubnet bool, withAdditionalPrivateSubnet bool, createNatGateways bool) error {

	var natGateway *ec2.NatGateway = nil

	if withPublicSubnet {
		var createPublicSubnetErr error = nil
		natGateway, createPublicSubnetErr = createPublicSubnet(ctx, resourceId(namespace, "public"), config.publicSubnetCidr, resourceName(namespace, config.publicSubnetName), config.availabilityZone, vpc, internetGateway, createNatGateways)

		if createPublicSubnetErr != nil {
			return createPublicSubnetErr
		}
	}

	if withPrivateSubnet {
		createPrivateSubnetErr := createPrivateSubnet(ctx, resourceId(namespace, "privateA"), config.privateSubnetACidr, resourceName(namespace, config.privateSubnetAName), config.availabilityZone, vpc, natGateway, false)

		if createPrivateSubnetErr != nil {
			return createPrivateSubnetErr
		}

	}

	if withAdditionalPrivateSubnet {
		createAdditionalPrivateSubnetErr := createPrivateSubnet(ctx, resourceId(namespace, "privateB"), config.privateSubnetBCidr, resourceName(namespace, config.privateSubnetBName), config.availabilityZone, vpc, natGateway, true)

		if createAdditionalPrivateSubnetErr != nil {
			return createAdditionalPrivateSubnetErr
		}
	}

	return nil
}
