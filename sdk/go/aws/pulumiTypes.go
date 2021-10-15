// *** WARNING: this file was generated by Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package aws

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The list of Configurations in which to create subnets. You can specify availability
// zone with a private or a public subnet cidr block. You can also associated a private
// subnet with a dedicated network ACL.
type AvailabilityZone struct {
	// The availability zone name to deploy the subnet into
	AvailabilityZone string `pulumi:"availabilityZone"`
	// CIDR block for a private subnet located in the specified availability zone
	PrivateSubnetACidr *string `pulumi:"privateSubnetACidr"`
	// CIDR block for the associated private subnet (2) with a dedicated network ACL located in the specified availability zone. This subnet
	// will only be created if `CreateAdditionalPrivateSubnets` is `true`.
	PrivateSubnetBCidr *string `pulumi:"privateSubnetBCidr"`
	// CIDR block for the public subnet located in the specified availability zone
	PublicSubnetCidr *string `pulumi:"publicSubnetCidr"`
}

// AvailabilityZoneInput is an input type that accepts AvailabilityZoneArgs and AvailabilityZoneOutput values.
// You can construct a concrete instance of `AvailabilityZoneInput` via:
//
//          AvailabilityZoneArgs{...}
type AvailabilityZoneInput interface {
	pulumi.Input

	ToAvailabilityZoneOutput() AvailabilityZoneOutput
	ToAvailabilityZoneOutputWithContext(context.Context) AvailabilityZoneOutput
}

// The list of Configurations in which to create subnets. You can specify availability
// zone with a private or a public subnet cidr block. You can also associated a private
// subnet with a dedicated network ACL.
type AvailabilityZoneArgs struct {
	// The availability zone name to deploy the subnet into
	AvailabilityZone string `pulumi:"availabilityZone"`
	// CIDR block for a private subnet located in the specified availability zone
	PrivateSubnetACidr *string `pulumi:"privateSubnetACidr"`
	// CIDR block for the associated private subnet (2) with a dedicated network ACL located in the specified availability zone. This subnet
	// will only be created if `CreateAdditionalPrivateSubnets` is `true`.
	PrivateSubnetBCidr *string `pulumi:"privateSubnetBCidr"`
	// CIDR block for the public subnet located in the specified availability zone
	PublicSubnetCidr *string `pulumi:"publicSubnetCidr"`
}

func (AvailabilityZoneArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AvailabilityZone)(nil)).Elem()
}

func (i AvailabilityZoneArgs) ToAvailabilityZoneOutput() AvailabilityZoneOutput {
	return i.ToAvailabilityZoneOutputWithContext(context.Background())
}

func (i AvailabilityZoneArgs) ToAvailabilityZoneOutputWithContext(ctx context.Context) AvailabilityZoneOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AvailabilityZoneOutput)
}

// AvailabilityZoneArrayInput is an input type that accepts AvailabilityZoneArray and AvailabilityZoneArrayOutput values.
// You can construct a concrete instance of `AvailabilityZoneArrayInput` via:
//
//          AvailabilityZoneArray{ AvailabilityZoneArgs{...} }
type AvailabilityZoneArrayInput interface {
	pulumi.Input

	ToAvailabilityZoneArrayOutput() AvailabilityZoneArrayOutput
	ToAvailabilityZoneArrayOutputWithContext(context.Context) AvailabilityZoneArrayOutput
}

type AvailabilityZoneArray []AvailabilityZoneInput

func (AvailabilityZoneArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AvailabilityZone)(nil)).Elem()
}

func (i AvailabilityZoneArray) ToAvailabilityZoneArrayOutput() AvailabilityZoneArrayOutput {
	return i.ToAvailabilityZoneArrayOutputWithContext(context.Background())
}

func (i AvailabilityZoneArray) ToAvailabilityZoneArrayOutputWithContext(ctx context.Context) AvailabilityZoneArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AvailabilityZoneArrayOutput)
}

// The list of Configurations in which to create subnets. You can specify availability
// zone with a private or a public subnet cidr block. You can also associated a private
// subnet with a dedicated network ACL.
type AvailabilityZoneOutput struct{ *pulumi.OutputState }

func (AvailabilityZoneOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AvailabilityZone)(nil)).Elem()
}

func (o AvailabilityZoneOutput) ToAvailabilityZoneOutput() AvailabilityZoneOutput {
	return o
}

func (o AvailabilityZoneOutput) ToAvailabilityZoneOutputWithContext(ctx context.Context) AvailabilityZoneOutput {
	return o
}

// The availability zone name to deploy the subnet into
func (o AvailabilityZoneOutput) AvailabilityZone() pulumi.StringOutput {
	return o.ApplyT(func(v AvailabilityZone) string { return v.AvailabilityZone }).(pulumi.StringOutput)
}

// CIDR block for a private subnet located in the specified availability zone
func (o AvailabilityZoneOutput) PrivateSubnetACidr() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AvailabilityZone) *string { return v.PrivateSubnetACidr }).(pulumi.StringPtrOutput)
}

// CIDR block for the associated private subnet (2) with a dedicated network ACL located in the specified availability zone. This subnet
// will only be created if `CreateAdditionalPrivateSubnets` is `true`.
func (o AvailabilityZoneOutput) PrivateSubnetBCidr() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AvailabilityZone) *string { return v.PrivateSubnetBCidr }).(pulumi.StringPtrOutput)
}

// CIDR block for the public subnet located in the specified availability zone
func (o AvailabilityZoneOutput) PublicSubnetCidr() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AvailabilityZone) *string { return v.PublicSubnetCidr }).(pulumi.StringPtrOutput)
}

type AvailabilityZoneArrayOutput struct{ *pulumi.OutputState }

func (AvailabilityZoneArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AvailabilityZone)(nil)).Elem()
}

func (o AvailabilityZoneArrayOutput) ToAvailabilityZoneArrayOutput() AvailabilityZoneArrayOutput {
	return o
}

func (o AvailabilityZoneArrayOutput) ToAvailabilityZoneArrayOutputWithContext(ctx context.Context) AvailabilityZoneArrayOutput {
	return o
}

func (o AvailabilityZoneArrayOutput) Index(i pulumi.IntInput) AvailabilityZoneOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) AvailabilityZone {
		return vs[0].([]AvailabilityZone)[vs[1].(int)]
	}).(AvailabilityZoneOutput)
}

func init() {
	pulumi.RegisterOutputType(AvailabilityZoneOutput{})
	pulumi.RegisterOutputType(AvailabilityZoneArrayOutput{})
}