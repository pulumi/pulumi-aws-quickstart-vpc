package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/ec2"
	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/s3"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		// Create an AWS resource (S3 Bucket)
		bucket, bucketErr := s3.NewBucket(ctx, "my-bucket", nil)

		if bucketErr != nil {
			return bucketErr
		}

		vpcName := "Pulumi VPC"

		// VPC block
		cidrBlock := "10.0.0.0/16"
		vpc, vpcErr := ec2.NewVpc(ctx, "vpc", &ec2.VpcArgs{
			CidrBlock: pulumi.String(cidrBlock),
			Tags: pulumi.StringMap{
				"Name": pulumi.String(vpcName),
			},
		})

		if vpcErr != nil {
			return vpcErr
		}

		// Add a public subnet
		publicSubnetName := "Public Subnet"
		_, publicSubnetErr := ec2.NewSubnet(ctx, "public-subnet", &ec2.SubnetArgs{
			VpcId:     vpc.ID(),
			CidrBlock: pulumi.String("10.0.1.0/24"),
			Tags: pulumi.StringMap{
				"Name": pulumi.String(publicSubnetName),
			},
		})

		if publicSubnetErr != nil {
			return publicSubnetErr
		}

		// Export the name of the bucket
		ctx.Export("bucketName", bucket.ID())
		return nil
	})
}
