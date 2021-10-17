import * as quickstartVpc from "@pulumi/aws-quickstart-vpc";

const myVpc = new quickstartVpc.Vpc("test", {
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

// Export the name of the bucket
export const bucketName = myVpc.vpcID;
