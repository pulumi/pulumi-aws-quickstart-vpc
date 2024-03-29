// Code generated by Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package aws

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Vpc struct {
	pulumi.ResourceState

	// The IPs of the EIPs associated with the Nat Gateways
	NatGatewayIPs pulumi.StringArrayOutput `pulumi:"natGatewayIPs"`
	// The IDs of the Private Subnets Created
	PrivateSubnetIDs pulumi.StringArrayOutput `pulumi:"privateSubnetIDs"`
	// The IDs of the Public Subnets Created
	PublicSubnetIDs pulumi.StringArrayOutput `pulumi:"publicSubnetIDs"`
	// The VPC ID
	VpcID pulumi.StringOutput `pulumi:"vpcID"`
}

// NewVpc registers a new resource with the given unique name, arguments, and options.
func NewVpc(ctx *pulumi.Context,
	name string, args *VpcArgs, opts ...pulumi.ResourceOption) (*Vpc, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AvailabilityZoneConfig == nil {
		return nil, errors.New("invalid value for required argument 'AvailabilityZoneConfig'")
	}
	var resource Vpc
	err := ctx.RegisterRemoteComponentResource("aws-quickstart-vpc:index:Vpc", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

type vpcArgs struct {
	// The list of Configurations in which to create subnets. You can specify availability
	// zone with a private or a public subnet cidr block. You can also associated a private
	// subnet with a dedicated network ACL.
	AvailabilityZoneConfig []AvailabilityZone `pulumi:"availabilityZoneConfig"`
	// CIDR block for the VPC
	CidrBlock string `pulumi:"cidrBlock"`
	// Set to `true` to create a network ACL protected subnet in each Availability Zone. If `false`, the CIDR parameters for those subnets will be ignored.
	// If `true`, it also requires that the 'Create private subnets' parameter is also `true` to have any effect.
	// Default is `true`
	CreateAdditionalPrivateSubnets *bool `pulumi:"createAdditionalPrivateSubnets"`
	// Enable Flow Logs to capture IP traffic for the VPC. Defaults to `true`
	CreateFlowLogs *bool `pulumi:"createFlowLogs"`
	// Set to `false` when creating only private subnets. If `true`, both CreatePublicSubnets and CreatePrivateSubnets must also be `true`.
	// Default is `true`
	CreateNatGateways *bool `pulumi:"createNatGateways"`
	// Set to `false` to create only public subnets. If `false`, the CIDR parameters for ALL private subnets will be ignored.
	// Default is `true`.
	CreatePrivateSubnets *bool `pulumi:"createPrivateSubnets"`
	// Set to `false` to create only private subnets. If `false`, CreatePrivateSubnets must be `true` and the CIDR parameters for ALL public subnets will be
	// ignored. Default is `true`
	CreatePublicSubnets *bool `pulumi:"createPublicSubnets"`
	// A boolean flag to enable/disable DNS hostnames in the VPC. Defaults `false`.
	EnableDnsHostnames *bool `pulumi:"enableDnsHostnames"`
	// A boolean flag to enable/disable DNS support in the VPC. Defaults true.
	EnableDnsSupport *bool `pulumi:"enableDnsSupport"`
	// The fields to include in the flow log record, in the order in which they should appear. Specify the fields using the ${field-id} format,
	// separated by spaces. Default is
	// `${version} ${account-id} ${interface-id} ${srcaddr} ${dstaddr} ${srcport} ${dstport} ${protocol} ${packets} ${bytes} ${start} ${end} ${action} ${log-status}`
	FlowLogsLogFormat *string `pulumi:"flowLogsLogFormat"`
	// The maximum interval of time during which a flow of packets is captured and aggregated into a flow log record. You can specify 60 seconds
	// (1 minute) or 600 seconds (10 minutes). Default is `600`
	FlowLogsMaxAggregationInterval *float64 `pulumi:"flowLogsMaxAggregationInterval"`
	// Number of days to retain the VPC Flow Logs in CloudWatch. Defaults to `14`.
	FlowLogsRetentionPeriodInDays *float64 `pulumi:"flowLogsRetentionPeriodInDays"`
	// The type of traffic to log. You can log traffic that the resource accepts or rejects, or all traffic.
	// Default is `REJECT`.
	FlowLogsTrafficType *string `pulumi:"flowLogsTrafficType"`
	// The allowed tenancy of instances launched into the VPC. Defaults to `default`.
	InstanceTenancy *string `pulumi:"instanceTenancy"`
}

// The set of arguments for constructing a Vpc resource.
type VpcArgs struct {
	// The list of Configurations in which to create subnets. You can specify availability
	// zone with a private or a public subnet cidr block. You can also associated a private
	// subnet with a dedicated network ACL.
	AvailabilityZoneConfig []AvailabilityZoneArgs
	// CIDR block for the VPC
	CidrBlock string
	// Set to `true` to create a network ACL protected subnet in each Availability Zone. If `false`, the CIDR parameters for those subnets will be ignored.
	// If `true`, it also requires that the 'Create private subnets' parameter is also `true` to have any effect.
	// Default is `true`
	CreateAdditionalPrivateSubnets *bool
	// Enable Flow Logs to capture IP traffic for the VPC. Defaults to `true`
	CreateFlowLogs *bool
	// Set to `false` when creating only private subnets. If `true`, both CreatePublicSubnets and CreatePrivateSubnets must also be `true`.
	// Default is `true`
	CreateNatGateways *bool
	// Set to `false` to create only public subnets. If `false`, the CIDR parameters for ALL private subnets will be ignored.
	// Default is `true`.
	CreatePrivateSubnets *bool
	// Set to `false` to create only private subnets. If `false`, CreatePrivateSubnets must be `true` and the CIDR parameters for ALL public subnets will be
	// ignored. Default is `true`
	CreatePublicSubnets *bool
	// A boolean flag to enable/disable DNS hostnames in the VPC. Defaults `false`.
	EnableDnsHostnames *bool
	// A boolean flag to enable/disable DNS support in the VPC. Defaults true.
	EnableDnsSupport *bool
	// The fields to include in the flow log record, in the order in which they should appear. Specify the fields using the ${field-id} format,
	// separated by spaces. Default is
	// `${version} ${account-id} ${interface-id} ${srcaddr} ${dstaddr} ${srcport} ${dstport} ${protocol} ${packets} ${bytes} ${start} ${end} ${action} ${log-status}`
	FlowLogsLogFormat *string
	// The maximum interval of time during which a flow of packets is captured and aggregated into a flow log record. You can specify 60 seconds
	// (1 minute) or 600 seconds (10 minutes). Default is `600`
	FlowLogsMaxAggregationInterval *float64
	// Number of days to retain the VPC Flow Logs in CloudWatch. Defaults to `14`.
	FlowLogsRetentionPeriodInDays *float64
	// The type of traffic to log. You can log traffic that the resource accepts or rejects, or all traffic.
	// Default is `REJECT`.
	FlowLogsTrafficType *string
	// The allowed tenancy of instances launched into the VPC. Defaults to `default`.
	InstanceTenancy *string
}

func (VpcArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*vpcArgs)(nil)).Elem()
}

type VpcInput interface {
	pulumi.Input

	ToVpcOutput() VpcOutput
	ToVpcOutputWithContext(ctx context.Context) VpcOutput
}

func (*Vpc) ElementType() reflect.Type {
	return reflect.TypeOf((**Vpc)(nil)).Elem()
}

func (i *Vpc) ToVpcOutput() VpcOutput {
	return i.ToVpcOutputWithContext(context.Background())
}

func (i *Vpc) ToVpcOutputWithContext(ctx context.Context) VpcOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpcOutput)
}

// VpcArrayInput is an input type that accepts VpcArray and VpcArrayOutput values.
// You can construct a concrete instance of `VpcArrayInput` via:
//
//	VpcArray{ VpcArgs{...} }
type VpcArrayInput interface {
	pulumi.Input

	ToVpcArrayOutput() VpcArrayOutput
	ToVpcArrayOutputWithContext(context.Context) VpcArrayOutput
}

type VpcArray []VpcInput

func (VpcArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Vpc)(nil)).Elem()
}

func (i VpcArray) ToVpcArrayOutput() VpcArrayOutput {
	return i.ToVpcArrayOutputWithContext(context.Background())
}

func (i VpcArray) ToVpcArrayOutputWithContext(ctx context.Context) VpcArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpcArrayOutput)
}

// VpcMapInput is an input type that accepts VpcMap and VpcMapOutput values.
// You can construct a concrete instance of `VpcMapInput` via:
//
//	VpcMap{ "key": VpcArgs{...} }
type VpcMapInput interface {
	pulumi.Input

	ToVpcMapOutput() VpcMapOutput
	ToVpcMapOutputWithContext(context.Context) VpcMapOutput
}

type VpcMap map[string]VpcInput

func (VpcMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Vpc)(nil)).Elem()
}

func (i VpcMap) ToVpcMapOutput() VpcMapOutput {
	return i.ToVpcMapOutputWithContext(context.Background())
}

func (i VpcMap) ToVpcMapOutputWithContext(ctx context.Context) VpcMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpcMapOutput)
}

type VpcOutput struct{ *pulumi.OutputState }

func (VpcOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Vpc)(nil)).Elem()
}

func (o VpcOutput) ToVpcOutput() VpcOutput {
	return o
}

func (o VpcOutput) ToVpcOutputWithContext(ctx context.Context) VpcOutput {
	return o
}

// The IPs of the EIPs associated with the Nat Gateways
func (o VpcOutput) NatGatewayIPs() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Vpc) pulumi.StringArrayOutput { return v.NatGatewayIPs }).(pulumi.StringArrayOutput)
}

// The IDs of the Private Subnets Created
func (o VpcOutput) PrivateSubnetIDs() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Vpc) pulumi.StringArrayOutput { return v.PrivateSubnetIDs }).(pulumi.StringArrayOutput)
}

// The IDs of the Public Subnets Created
func (o VpcOutput) PublicSubnetIDs() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Vpc) pulumi.StringArrayOutput { return v.PublicSubnetIDs }).(pulumi.StringArrayOutput)
}

// The VPC ID
func (o VpcOutput) VpcID() pulumi.StringOutput {
	return o.ApplyT(func(v *Vpc) pulumi.StringOutput { return v.VpcID }).(pulumi.StringOutput)
}

type VpcArrayOutput struct{ *pulumi.OutputState }

func (VpcArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Vpc)(nil)).Elem()
}

func (o VpcArrayOutput) ToVpcArrayOutput() VpcArrayOutput {
	return o
}

func (o VpcArrayOutput) ToVpcArrayOutputWithContext(ctx context.Context) VpcArrayOutput {
	return o
}

func (o VpcArrayOutput) Index(i pulumi.IntInput) VpcOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Vpc {
		return vs[0].([]*Vpc)[vs[1].(int)]
	}).(VpcOutput)
}

type VpcMapOutput struct{ *pulumi.OutputState }

func (VpcMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Vpc)(nil)).Elem()
}

func (o VpcMapOutput) ToVpcMapOutput() VpcMapOutput {
	return o
}

func (o VpcMapOutput) ToVpcMapOutputWithContext(ctx context.Context) VpcMapOutput {
	return o
}

func (o VpcMapOutput) MapIndex(k pulumi.StringInput) VpcOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Vpc {
		return vs[0].(map[string]*Vpc)[vs[1].(string)]
	}).(VpcOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*VpcInput)(nil)).Elem(), &Vpc{})
	pulumi.RegisterInputType(reflect.TypeOf((*VpcArrayInput)(nil)).Elem(), VpcArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*VpcMapInput)(nil)).Elem(), VpcMap{})
	pulumi.RegisterOutputType(VpcOutput{})
	pulumi.RegisterOutputType(VpcArrayOutput{})
	pulumi.RegisterOutputType(VpcMapOutput{})
}
