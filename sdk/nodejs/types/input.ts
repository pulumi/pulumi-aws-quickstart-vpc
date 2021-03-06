// *** WARNING: this file was generated by Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";

/**
 * The list of Configurations in which to create subnets. You can specify availability
 * zone with a private or a public subnet cidr block. You can also associated a private
 * subnet with a dedicated network ACL.
 */
export interface AvailabilityZoneArgs {
    /**
     * The availability zone name to deploy the subnet into
     */
    availabilityZone: string;
    /**
     * CIDR block for a private subnet located in the specified availability zone
     */
    privateSubnetACidr?: string;
    /**
     * CIDR block for the associated private subnet (2) with a dedicated network ACL located in the specified availability zone. This subnet
     * will only be created if `CreateAdditionalPrivateSubnets` is `true`.
     */
    privateSubnetBCidr?: string;
    /**
     * CIDR block for the public subnet located in the specified availability zone
     */
    publicSubnetCidr?: string;
}
