# coding=utf-8
# *** WARNING: this file was generated by Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities
from ._inputs import *

__all__ = ['VpcArgs', 'Vpc']

@pulumi.input_type
class VpcArgs:
    def __init__(__self__, *,
                 availability_zone_config: Sequence['AvailabilityZoneArgs'],
                 cidr_block: str,
                 create_additional_private_subnets: Optional[bool] = None,
                 create_flow_logs: Optional[bool] = None,
                 create_nat_gateways: Optional[bool] = None,
                 create_private_subnets: Optional[bool] = None,
                 create_public_subnets: Optional[bool] = None,
                 enable_dns_hostnames: Optional[bool] = None,
                 enable_dns_support: Optional[bool] = None,
                 flow_logs_log_format: Optional[str] = None,
                 flow_logs_max_aggregation_interval: Optional[float] = None,
                 flow_logs_retention_period_in_days: Optional[float] = None,
                 flow_logs_traffic_type: Optional[str] = None,
                 instance_tenancy: Optional[str] = None):
        """
        The set of arguments for constructing a Vpc resource.
        :param Sequence['AvailabilityZoneArgs'] availability_zone_config: The list of Configurations in which to create subnets. You can specify availability
               zone with a private or a public subnet cidr block. You can also associated a private
               subnet with a dedicated network ACL.
        :param str cidr_block: CIDR block for the VPC
        :param bool create_additional_private_subnets: Set to `true` to create a network ACL protected subnet in each Availability Zone. If `false`, the CIDR parameters for those subnets will be ignored.
               If `true`, it also requires that the 'Create private subnets' parameter is also `true` to have any effect.
               Default is `true`
        :param bool create_flow_logs: Enable Flow Logs to capture IP traffic for the VPC. Defaults to `true`
        :param bool create_nat_gateways: Set to `false` when creating only private subnets. If `true`, both CreatePublicSubnets and CreatePrivateSubnets must also be `true`.
               Default is `true`
        :param bool create_private_subnets: Set to `false` to create only public subnets. If `false`, the CIDR parameters for ALL private subnets will be ignored.
               Default is `true`.
        :param bool create_public_subnets: Set to `false` to create only private subnets. If `false`, CreatePrivateSubnets must be `true` and the CIDR parameters for ALL public subnets will be
               ignored. Default is `true`
        :param bool enable_dns_hostnames: A boolean flag to enable/disable DNS hostnames in the VPC. Defaults `false`.
        :param bool enable_dns_support: A boolean flag to enable/disable DNS support in the VPC. Defaults true.
        :param str flow_logs_log_format: The fields to include in the flow log record, in the order in which they should appear. Specify the fields using the ${field-id} format,
               separated by spaces. Default is
               `${version} ${account-id} ${interface-id} ${srcaddr} ${dstaddr} ${srcport} ${dstport} ${protocol} ${packets} ${bytes} ${start} ${end} ${action} ${log-status}`
        :param float flow_logs_max_aggregation_interval: The maximum interval of time during which a flow of packets is captured and aggregated into a flow log record. You can specify 60 seconds
               (1 minute) or 600 seconds (10 minutes). Default is `600`
        :param float flow_logs_retention_period_in_days: Number of days to retain the VPC Flow Logs in CloudWatch. Defaults to `14`.
        :param str flow_logs_traffic_type: The type of traffic to log. You can log traffic that the resource accepts or rejects, or all traffic.
               Default is `REJECT`.
        :param str instance_tenancy: The allowed tenancy of instances launched into the VPC. Defaults to `default`.
        """
        pulumi.set(__self__, "availability_zone_config", availability_zone_config)
        pulumi.set(__self__, "cidr_block", cidr_block)
        if create_additional_private_subnets is not None:
            pulumi.set(__self__, "create_additional_private_subnets", create_additional_private_subnets)
        if create_flow_logs is not None:
            pulumi.set(__self__, "create_flow_logs", create_flow_logs)
        if create_nat_gateways is not None:
            pulumi.set(__self__, "create_nat_gateways", create_nat_gateways)
        if create_private_subnets is not None:
            pulumi.set(__self__, "create_private_subnets", create_private_subnets)
        if create_public_subnets is not None:
            pulumi.set(__self__, "create_public_subnets", create_public_subnets)
        if enable_dns_hostnames is not None:
            pulumi.set(__self__, "enable_dns_hostnames", enable_dns_hostnames)
        if enable_dns_support is not None:
            pulumi.set(__self__, "enable_dns_support", enable_dns_support)
        if flow_logs_log_format is not None:
            pulumi.set(__self__, "flow_logs_log_format", flow_logs_log_format)
        if flow_logs_max_aggregation_interval is not None:
            pulumi.set(__self__, "flow_logs_max_aggregation_interval", flow_logs_max_aggregation_interval)
        if flow_logs_retention_period_in_days is not None:
            pulumi.set(__self__, "flow_logs_retention_period_in_days", flow_logs_retention_period_in_days)
        if flow_logs_traffic_type is not None:
            pulumi.set(__self__, "flow_logs_traffic_type", flow_logs_traffic_type)
        if instance_tenancy is not None:
            pulumi.set(__self__, "instance_tenancy", instance_tenancy)

    @property
    @pulumi.getter(name="availabilityZoneConfig")
    def availability_zone_config(self) -> Sequence['AvailabilityZoneArgs']:
        """
        The list of Configurations in which to create subnets. You can specify availability
        zone with a private or a public subnet cidr block. You can also associated a private
        subnet with a dedicated network ACL.
        """
        return pulumi.get(self, "availability_zone_config")

    @availability_zone_config.setter
    def availability_zone_config(self, value: Sequence['AvailabilityZoneArgs']):
        pulumi.set(self, "availability_zone_config", value)

    @property
    @pulumi.getter(name="cidrBlock")
    def cidr_block(self) -> str:
        """
        CIDR block for the VPC
        """
        return pulumi.get(self, "cidr_block")

    @cidr_block.setter
    def cidr_block(self, value: str):
        pulumi.set(self, "cidr_block", value)

    @property
    @pulumi.getter(name="createAdditionalPrivateSubnets")
    def create_additional_private_subnets(self) -> Optional[bool]:
        """
        Set to `true` to create a network ACL protected subnet in each Availability Zone. If `false`, the CIDR parameters for those subnets will be ignored.
        If `true`, it also requires that the 'Create private subnets' parameter is also `true` to have any effect.
        Default is `true`
        """
        return pulumi.get(self, "create_additional_private_subnets")

    @create_additional_private_subnets.setter
    def create_additional_private_subnets(self, value: Optional[bool]):
        pulumi.set(self, "create_additional_private_subnets", value)

    @property
    @pulumi.getter(name="createFlowLogs")
    def create_flow_logs(self) -> Optional[bool]:
        """
        Enable Flow Logs to capture IP traffic for the VPC. Defaults to `true`
        """
        return pulumi.get(self, "create_flow_logs")

    @create_flow_logs.setter
    def create_flow_logs(self, value: Optional[bool]):
        pulumi.set(self, "create_flow_logs", value)

    @property
    @pulumi.getter(name="createNatGateways")
    def create_nat_gateways(self) -> Optional[bool]:
        """
        Set to `false` when creating only private subnets. If `true`, both CreatePublicSubnets and CreatePrivateSubnets must also be `true`.
        Default is `true`
        """
        return pulumi.get(self, "create_nat_gateways")

    @create_nat_gateways.setter
    def create_nat_gateways(self, value: Optional[bool]):
        pulumi.set(self, "create_nat_gateways", value)

    @property
    @pulumi.getter(name="createPrivateSubnets")
    def create_private_subnets(self) -> Optional[bool]:
        """
        Set to `false` to create only public subnets. If `false`, the CIDR parameters for ALL private subnets will be ignored.
        Default is `true`.
        """
        return pulumi.get(self, "create_private_subnets")

    @create_private_subnets.setter
    def create_private_subnets(self, value: Optional[bool]):
        pulumi.set(self, "create_private_subnets", value)

    @property
    @pulumi.getter(name="createPublicSubnets")
    def create_public_subnets(self) -> Optional[bool]:
        """
        Set to `false` to create only private subnets. If `false`, CreatePrivateSubnets must be `true` and the CIDR parameters for ALL public subnets will be
        ignored. Default is `true`
        """
        return pulumi.get(self, "create_public_subnets")

    @create_public_subnets.setter
    def create_public_subnets(self, value: Optional[bool]):
        pulumi.set(self, "create_public_subnets", value)

    @property
    @pulumi.getter(name="enableDnsHostnames")
    def enable_dns_hostnames(self) -> Optional[bool]:
        """
        A boolean flag to enable/disable DNS hostnames in the VPC. Defaults `false`.
        """
        return pulumi.get(self, "enable_dns_hostnames")

    @enable_dns_hostnames.setter
    def enable_dns_hostnames(self, value: Optional[bool]):
        pulumi.set(self, "enable_dns_hostnames", value)

    @property
    @pulumi.getter(name="enableDnsSupport")
    def enable_dns_support(self) -> Optional[bool]:
        """
        A boolean flag to enable/disable DNS support in the VPC. Defaults true.
        """
        return pulumi.get(self, "enable_dns_support")

    @enable_dns_support.setter
    def enable_dns_support(self, value: Optional[bool]):
        pulumi.set(self, "enable_dns_support", value)

    @property
    @pulumi.getter(name="flowLogsLogFormat")
    def flow_logs_log_format(self) -> Optional[str]:
        """
        The fields to include in the flow log record, in the order in which they should appear. Specify the fields using the ${field-id} format,
        separated by spaces. Default is
        `${version} ${account-id} ${interface-id} ${srcaddr} ${dstaddr} ${srcport} ${dstport} ${protocol} ${packets} ${bytes} ${start} ${end} ${action} ${log-status}`
        """
        return pulumi.get(self, "flow_logs_log_format")

    @flow_logs_log_format.setter
    def flow_logs_log_format(self, value: Optional[str]):
        pulumi.set(self, "flow_logs_log_format", value)

    @property
    @pulumi.getter(name="flowLogsMaxAggregationInterval")
    def flow_logs_max_aggregation_interval(self) -> Optional[float]:
        """
        The maximum interval of time during which a flow of packets is captured and aggregated into a flow log record. You can specify 60 seconds
        (1 minute) or 600 seconds (10 minutes). Default is `600`
        """
        return pulumi.get(self, "flow_logs_max_aggregation_interval")

    @flow_logs_max_aggregation_interval.setter
    def flow_logs_max_aggregation_interval(self, value: Optional[float]):
        pulumi.set(self, "flow_logs_max_aggregation_interval", value)

    @property
    @pulumi.getter(name="flowLogsRetentionPeriodInDays")
    def flow_logs_retention_period_in_days(self) -> Optional[float]:
        """
        Number of days to retain the VPC Flow Logs in CloudWatch. Defaults to `14`.
        """
        return pulumi.get(self, "flow_logs_retention_period_in_days")

    @flow_logs_retention_period_in_days.setter
    def flow_logs_retention_period_in_days(self, value: Optional[float]):
        pulumi.set(self, "flow_logs_retention_period_in_days", value)

    @property
    @pulumi.getter(name="flowLogsTrafficType")
    def flow_logs_traffic_type(self) -> Optional[str]:
        """
        The type of traffic to log. You can log traffic that the resource accepts or rejects, or all traffic.
        Default is `REJECT`.
        """
        return pulumi.get(self, "flow_logs_traffic_type")

    @flow_logs_traffic_type.setter
    def flow_logs_traffic_type(self, value: Optional[str]):
        pulumi.set(self, "flow_logs_traffic_type", value)

    @property
    @pulumi.getter(name="instanceTenancy")
    def instance_tenancy(self) -> Optional[str]:
        """
        The allowed tenancy of instances launched into the VPC. Defaults to `default`.
        """
        return pulumi.get(self, "instance_tenancy")

    @instance_tenancy.setter
    def instance_tenancy(self, value: Optional[str]):
        pulumi.set(self, "instance_tenancy", value)


class Vpc(pulumi.ComponentResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 availability_zone_config: Optional[Sequence[pulumi.InputType['AvailabilityZoneArgs']]] = None,
                 cidr_block: Optional[str] = None,
                 create_additional_private_subnets: Optional[bool] = None,
                 create_flow_logs: Optional[bool] = None,
                 create_nat_gateways: Optional[bool] = None,
                 create_private_subnets: Optional[bool] = None,
                 create_public_subnets: Optional[bool] = None,
                 enable_dns_hostnames: Optional[bool] = None,
                 enable_dns_support: Optional[bool] = None,
                 flow_logs_log_format: Optional[str] = None,
                 flow_logs_max_aggregation_interval: Optional[float] = None,
                 flow_logs_retention_period_in_days: Optional[float] = None,
                 flow_logs_traffic_type: Optional[str] = None,
                 instance_tenancy: Optional[str] = None,
                 __props__=None):
        """
        Create a Vpc resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param Sequence[pulumi.InputType['AvailabilityZoneArgs']] availability_zone_config: The list of Configurations in which to create subnets. You can specify availability
               zone with a private or a public subnet cidr block. You can also associated a private
               subnet with a dedicated network ACL.
        :param str cidr_block: CIDR block for the VPC
        :param bool create_additional_private_subnets: Set to `true` to create a network ACL protected subnet in each Availability Zone. If `false`, the CIDR parameters for those subnets will be ignored.
               If `true`, it also requires that the 'Create private subnets' parameter is also `true` to have any effect.
               Default is `true`
        :param bool create_flow_logs: Enable Flow Logs to capture IP traffic for the VPC. Defaults to `true`
        :param bool create_nat_gateways: Set to `false` when creating only private subnets. If `true`, both CreatePublicSubnets and CreatePrivateSubnets must also be `true`.
               Default is `true`
        :param bool create_private_subnets: Set to `false` to create only public subnets. If `false`, the CIDR parameters for ALL private subnets will be ignored.
               Default is `true`.
        :param bool create_public_subnets: Set to `false` to create only private subnets. If `false`, CreatePrivateSubnets must be `true` and the CIDR parameters for ALL public subnets will be
               ignored. Default is `true`
        :param bool enable_dns_hostnames: A boolean flag to enable/disable DNS hostnames in the VPC. Defaults `false`.
        :param bool enable_dns_support: A boolean flag to enable/disable DNS support in the VPC. Defaults true.
        :param str flow_logs_log_format: The fields to include in the flow log record, in the order in which they should appear. Specify the fields using the ${field-id} format,
               separated by spaces. Default is
               `${version} ${account-id} ${interface-id} ${srcaddr} ${dstaddr} ${srcport} ${dstport} ${protocol} ${packets} ${bytes} ${start} ${end} ${action} ${log-status}`
        :param float flow_logs_max_aggregation_interval: The maximum interval of time during which a flow of packets is captured and aggregated into a flow log record. You can specify 60 seconds
               (1 minute) or 600 seconds (10 minutes). Default is `600`
        :param float flow_logs_retention_period_in_days: Number of days to retain the VPC Flow Logs in CloudWatch. Defaults to `14`.
        :param str flow_logs_traffic_type: The type of traffic to log. You can log traffic that the resource accepts or rejects, or all traffic.
               Default is `REJECT`.
        :param str instance_tenancy: The allowed tenancy of instances launched into the VPC. Defaults to `default`.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: VpcArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a Vpc resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param VpcArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(VpcArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 availability_zone_config: Optional[Sequence[pulumi.InputType['AvailabilityZoneArgs']]] = None,
                 cidr_block: Optional[str] = None,
                 create_additional_private_subnets: Optional[bool] = None,
                 create_flow_logs: Optional[bool] = None,
                 create_nat_gateways: Optional[bool] = None,
                 create_private_subnets: Optional[bool] = None,
                 create_public_subnets: Optional[bool] = None,
                 enable_dns_hostnames: Optional[bool] = None,
                 enable_dns_support: Optional[bool] = None,
                 flow_logs_log_format: Optional[str] = None,
                 flow_logs_max_aggregation_interval: Optional[float] = None,
                 flow_logs_retention_period_in_days: Optional[float] = None,
                 flow_logs_traffic_type: Optional[str] = None,
                 instance_tenancy: Optional[str] = None,
                 __props__=None):
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is not None:
            raise ValueError('ComponentResource classes do not support opts.id')
        else:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = VpcArgs.__new__(VpcArgs)

            if availability_zone_config is None and not opts.urn:
                raise TypeError("Missing required property 'availability_zone_config'")
            __props__.__dict__["availability_zone_config"] = availability_zone_config
            if cidr_block is None and not opts.urn:
                raise TypeError("Missing required property 'cidr_block'")
            __props__.__dict__["cidr_block"] = cidr_block
            __props__.__dict__["create_additional_private_subnets"] = create_additional_private_subnets
            __props__.__dict__["create_flow_logs"] = create_flow_logs
            __props__.__dict__["create_nat_gateways"] = create_nat_gateways
            __props__.__dict__["create_private_subnets"] = create_private_subnets
            __props__.__dict__["create_public_subnets"] = create_public_subnets
            __props__.__dict__["enable_dns_hostnames"] = enable_dns_hostnames
            __props__.__dict__["enable_dns_support"] = enable_dns_support
            __props__.__dict__["flow_logs_log_format"] = flow_logs_log_format
            __props__.__dict__["flow_logs_max_aggregation_interval"] = flow_logs_max_aggregation_interval
            __props__.__dict__["flow_logs_retention_period_in_days"] = flow_logs_retention_period_in_days
            __props__.__dict__["flow_logs_traffic_type"] = flow_logs_traffic_type
            __props__.__dict__["instance_tenancy"] = instance_tenancy
            __props__.__dict__["nat_gateway_ips"] = None
            __props__.__dict__["private_subnet_ids"] = None
            __props__.__dict__["public_subnet_ids"] = None
            __props__.__dict__["vpc_id"] = None
        super(Vpc, __self__).__init__(
            'aws-quickstart-vpc:index:Vpc',
            resource_name,
            __props__,
            opts,
            remote=True)

    @property
    @pulumi.getter(name="natGatewayIPs")
    def nat_gateway_ips(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        The IPs of the EIPs associated with the Nat Gateways
        """
        return pulumi.get(self, "nat_gateway_ips")

    @property
    @pulumi.getter(name="privateSubnetIDs")
    def private_subnet_ids(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        The IDs of the Private Subnets Created
        """
        return pulumi.get(self, "private_subnet_ids")

    @property
    @pulumi.getter(name="publicSubnetIDs")
    def public_subnet_ids(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        The IDs of the Public Subnets Created
        """
        return pulumi.get(self, "public_subnet_ids")

    @property
    @pulumi.getter(name="vpcID")
    def vpc_id(self) -> pulumi.Output[str]:
        """
        The VPC ID
        """
        return pulumi.get(self, "vpc_id")

