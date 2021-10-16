package main

import (
	quickstartVpc "github.com/pulumi/pulumi-aws-quickstart-vpc/sdk/go/aws"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		publicSubnet1Cidr := "10.0.128.0/20"
		privateSubnet1ACidr := "10.0.32.0/19"

		publicSubnet2Cidr := "10.0.64.0/19"

		// Create an AWS resource (S3 Bucket)
		_, err := quickstartVpc.NewVpc(ctx, "simple-vpc", &quickstartVpc.VpcArgs{
			CidrBlock: "10.0.0.0/16",
			AvailabilityZoneConfig: []quickstartVpc.AvailabilityZoneArgs{
				quickstartVpc.AvailabilityZoneArgs{
					AvailabilityZone:   "us-east-1a",
					PublicSubnetCidr:   &publicSubnet1Cidr,
					PrivateSubnetACidr: &privateSubnet1ACidr,
				},
				quickstartVpc.AvailabilityZoneArgs{
					AvailabilityZone: "us-east-1b",
					PublicSubnetCidr: &publicSubnet2Cidr,
				},
			},
		})

		if err != nil {
			return err
		}

		return nil
	})
}
