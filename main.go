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

		// VPC block
		cidrBlock := "10.0.0.0/16"
		vpc, vpcErr := ec2.NewVpc(ctx, "vpc", &ec2.VpcArgs{
			CidrBlock: pulumi.String(cidrBlock),
		})

		if vpcErr != nil {
			return vpcErr
		}

		vpcName := "Pulumi VPC"

		// Add a name tag to the VPC.
		// @fixme, can this be added to the ec2.NewVpc constructor?
		_, vpcNameTagErr := ec2.NewTag(ctx, "vpcNameTag", &ec2.TagArgs{
			ResourceId: vpc.ID(),
			Key:        pulumi.String("Name"),
			Value:      pulumi.String(vpcName),
		})

		if vpcNameTagErr != nil {
			return vpcNameTagErr
		}

		// Export the name of the bucket
		ctx.Export("bucketName", bucket.ID())
		return nil
	})
}
