package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/ec2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		vpcName := "Pulumi VPC"

		namespace := "pulumi-vpc"

		// @fixme - not implemented yet. The main challenge is the false choice.
		// createPublicSubnets := true

		createPrivateSubnets := true
		createAdditionalPrivateSubnets := true

		// @fixme - not implemented yet. The main challenge is actually the "false" choice.
		// createNatGateways := true

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

		// @fixme - only create the internet gateway if public subnets are allowed.
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

		// Add a public subnet
		publicSubnet1Name := "Pulumi Public Subnet 1"
		publicSubnet1Cidr := "10.0.5.0/24" // Add a public subnet
		privateSubnet1AName := "Pulumi Private Subnet 1A"
		privateSubnet1ACidr := "10.0.7.0/24"
		privateSubnet1BName := "Pulumi Private Subnet 1B"
		privateSubnet1BCidr := "10.0.14.0/24"

		publicSubnet2Name := "Pulumi Public Subnet 2"
		publicSubnet2Cidr := "10.0.8.0/24"
		privateSubnet2AName := "Pulumi Private Subnet 2A"
		privateSubnet2ACidr := "10.0.9.0/24"
		privateSubnet2BName := "Pulumi Private Subnet 1A"
		privateSubnet2BCidr := "10.0.15.0/24"

		publicSubnet3Name := "Pulumi Public Subnet 3"
		publicSubnet3Cidr := "10.0.10.0/24"
		privateSubnet3AName := "Pulumi Private Subnet 3A"
		privateSubnet3ACidr := "10.0.11.0/24"
		privateSubnet3BName := "Pulumi Private Subnet 3B"
		privateSubnet3BCidr := "10.0.16.0/24"

		publicSubnet4Name := "Pulumi Public Subnet 4"
		publicSubnet4Cidr := "10.0.12.0/24"
		privateSubnet4AName := "Pulumi Private Subnet 4A"
		privateSubnet4ACidr := "10.0.13.0/24"
		privateSubnet4BName := "Pulumi Private Subnet 4B"
		privateSubnet4BCidr := "10.0.17.0/24"

		// Item 1
		createPublicPrivateSubnets(ctx, namespace+"1", vpc, publicSubnet1Cidr, publicSubnet1Name, internetGateway, privateSubnet1ACidr, privateSubnet1AName, privateSubnet1BCidr, privateSubnet1BName, createPrivateSubnets, createAdditionalPrivateSubnets)

		// Item 2
		createPublicPrivateSubnets(ctx, namespace+"2", vpc, publicSubnet2Cidr, publicSubnet2Name, internetGateway, privateSubnet2ACidr, privateSubnet2AName, privateSubnet2BCidr, privateSubnet2BName, createPrivateSubnets, createAdditionalPrivateSubnets)

		// // Item 1
		createPublicPrivateSubnets(ctx, namespace+"3", vpc, publicSubnet3Cidr, publicSubnet3Name, internetGateway, privateSubnet3ACidr, privateSubnet3AName, privateSubnet3BCidr, privateSubnet3BName, createPrivateSubnets, createAdditionalPrivateSubnets)

		// // Item 2
		createPublicPrivateSubnets(ctx, namespace+"4", vpc, publicSubnet4Cidr, publicSubnet4Name, internetGateway, privateSubnet4ACidr, privateSubnet4AName, privateSubnet4BCidr, privateSubnet4BName, createPrivateSubnets, createAdditionalPrivateSubnets)

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
func createPublicSubnet(ctx *pulumi.Context, namespace string, subnetCidr string, subnetName string, vpc *ec2.Vpc, internetGateway *ec2.InternetGateway) (*ec2.NatGateway, error) {
	natGatewayName := resourceName(namespace, "Pulumi NAT Gateway")
	publicRouteTableName := resourceName(namespace, "Pulumi Public Route Table")

	entireInternetCidr := "0.0.0.0/0"

	publicSubnet, publicSubnetErr := ec2.NewSubnet(ctx, resourceId(namespace, "public-subnet"), &ec2.SubnetArgs{
		VpcId:     vpc.ID(),
		CidrBlock: pulumi.String(subnetCidr),
		Tags: pulumi.StringMap{
			"Name": pulumi.String(subnetName),
		},
	})

	if publicSubnetErr != nil {
		return nil, publicSubnetErr
	}

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

	return natGateway, nil
}

func createPrivateSubnet(ctx *pulumi.Context, namespace string, subnetCidr string, subnetName string, vpc *ec2.Vpc, natGateway *ec2.NatGateway) error {
	entireInternetCidr := "0.0.0.0/0"

	privateSubnet, privateSubnetErr := ec2.NewSubnet(ctx, resourceId(namespace, "private-subnet"), &ec2.SubnetArgs{
		VpcId:     vpc.ID(),
		CidrBlock: pulumi.String(subnetCidr),
		Tags: pulumi.StringMap{
			"Name": pulumi.String(subnetName),
		},
	})

	if privateSubnetErr != nil {
		return privateSubnetErr
	}

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

	/***********************************************************
	 * Begin Private Network ACL
	 ***********************************************************/
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

	return nil
}

func createPublicPrivateSubnets(ctx *pulumi.Context, namespace string, vpc *ec2.Vpc, publicSubnetCidr string, publicSubnetName string, internetGateway *ec2.InternetGateway, privateSubnetCidr string, privateSubnetName string, privateSubnetBCidr string, privateSubnetBName string, withPrivateSubnet bool, withAdditionalPrivateSubnet bool) error {
	natGateway, createPublicSubnetErr := createPublicSubnet(ctx, resourceId(namespace, "public"), publicSubnetCidr, publicSubnetName, vpc, internetGateway)

	if createPublicSubnetErr != nil {
		return createPublicSubnetErr
	}

	if withPrivateSubnet {
		createPrivateSubnetErr := createPrivateSubnet(ctx, resourceId(namespace, "privateA"), privateSubnetCidr, resourceName(namespace, privateSubnetName), vpc, natGateway)

		if createPrivateSubnetErr != nil {
			return createPrivateSubnetErr
		}

	}

	if withAdditionalPrivateSubnet {
		createAdditionalPrivateSubnetErr := createPrivateSubnet(ctx, resourceId(namespace, "privateB"), privateSubnetBCidr, resourceName(namespace, privateSubnetBName), vpc, natGateway)

		if createAdditionalPrivateSubnetErr != nil {
			return createAdditionalPrivateSubnetErr
		}
	}

	return nil
}
