# Pulumi AWS VPC

Easily deploy an AWS VPC with multiple public, private, and/or isolated subnets. Also provisions networking resources like NAT Gateways, Internet Gateways, Route Tables, Security Group, and VPC Flow Logs. This component is based on the best practices recommended by AWS in the [VPC Architecture](https://aws.amazon.com/quickstart/architecture/vpc/)

# Examples

See the `/examples` directory for more

Go:
```go
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
```

Typescript:
```typescript
const myVpc = new quickstartVpc.Vpc("simple-vpc", {
    cidrBlock: "10.0.0.0/16",
    availabilityZoneConfig: [{
        availabilityZone: "us-east-1a",
        publicSubnetCidr: "10.0.128.0/20",
        privateSubnetACidr: "10.0.32.0/19",
    }, {
        availabilityZone: "us-east-1b",
        privateSubnetACidr: "10.0.64.0/19",
    }]
})

```