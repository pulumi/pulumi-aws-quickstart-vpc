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
		publicSubnetName := "Pulumi Public Subnet"
		publicSubnetCidr := "10.0.5.0/24"
		publicSubnet, publicSubnetErr := ec2.NewSubnet(ctx, "public-subnet", &ec2.SubnetArgs{
			VpcId:     vpc.ID(),
			CidrBlock: pulumi.String(publicSubnetCidr),
			Tags: pulumi.StringMap{
				"Name": pulumi.String(publicSubnetName),
			},
		})

		if publicSubnetErr != nil {
			return publicSubnetErr
		}

		// Add a private subnet
		privateSubnetName := "Pulumi Private Subnet"
		privateSubnetCidr := "10.0.6.0/24"
		_, privateSubnetErr := ec2.NewSubnet(ctx, "private-subnet", &ec2.SubnetArgs{
			VpcId:     vpc.ID(),
			CidrBlock: pulumi.String(privateSubnetCidr),
			Tags: pulumi.StringMap{
				"Name": pulumi.String(privateSubnetName),
			},
		})

		if privateSubnetErr != nil {
			return privateSubnetErr
		}

		// Internet Gateway
		internetGatewayName := "Pulumi Internet Gateway"
		internetGateway, internetGatewayErr := ec2.NewInternetGateway(ctx, "internet-gateway", &ec2.InternetGatewayArgs{
			VpcId: vpc.ID(),
			Tags: pulumi.StringMap{
				"Name": pulumi.String(internetGatewayName),
			},
		})

		if internetGatewayErr != nil {
			return internetGatewayErr
		}

		// NAT Gateway + Elastic IP
		elasticIp, elasticIpErr := ec2.NewEip(ctx, "lb", &ec2.EipArgs{
			Vpc: pulumi.Bool(true),
		})

		if elasticIpErr != nil {
			return elasticIpErr
		}

		natGatewayName := "Pulumi NAT Gateway"
		_, natGatewayErr := ec2.NewNatGateway(ctx, "nat-gateway", &ec2.NatGatewayArgs{
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
			return elasticIpErr
		}

		publicRouteTableName := "Pulumi Public Route Table"
		publicRouteTable, publicRouteTableErr := ec2.NewRouteTable(ctx, "public-route-table", &ec2.RouteTableArgs{
			VpcId: vpc.ID(),
			Tags: pulumi.StringMap{
				"Name": pulumi.String(publicRouteTableName),
			},
		})

		entireInternetCidr := "0.0.0.0/0"
		ec2.NewRoute(ctx, "temp", &ec2.RouteArgs{
			RouteTableId:         publicRouteTable.ID(),
			DestinationCidrBlock: pulumi.String(entireInternetCidr),
			GatewayId:            internetGateway.ID(),
		})

		if publicRouteTableErr != nil {
			return publicRouteTableErr
		}

		// Export the name of the bucket
		ctx.Export("bucketName", bucket.ID())
		return nil
	})
}
