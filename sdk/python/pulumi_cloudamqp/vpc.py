# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['VpcArgs', 'Vpc']

@pulumi.input_type
class VpcArgs:
    def __init__(__self__, *,
                 region: pulumi.Input[str],
                 subnet: pulumi.Input[str],
                 name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        """
        The set of arguments for constructing a Vpc resource.
        :param pulumi.Input[str] region: The hosted region for the managed standalone VPC
        :param pulumi.Input[str] subnet: The VPC subnet
        :param pulumi.Input[str] name: The name of the VPC.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] tags: Tag the VPC with optional tags
        """
        pulumi.set(__self__, "region", region)
        pulumi.set(__self__, "subnet", subnet)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter
    def region(self) -> pulumi.Input[str]:
        """
        The hosted region for the managed standalone VPC
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: pulumi.Input[str]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter
    def subnet(self) -> pulumi.Input[str]:
        """
        The VPC subnet
        """
        return pulumi.get(self, "subnet")

    @subnet.setter
    def subnet(self, value: pulumi.Input[str]):
        pulumi.set(self, "subnet", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the VPC.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Tag the VPC with optional tags
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)


@pulumi.input_type
class _VpcState:
    def __init__(__self__, *,
                 name: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 subnet: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 vpc_name: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Vpc resources.
        :param pulumi.Input[str] name: The name of the VPC.
        :param pulumi.Input[str] region: The hosted region for the managed standalone VPC
        :param pulumi.Input[str] subnet: The VPC subnet
        :param pulumi.Input[Sequence[pulumi.Input[str]]] tags: Tag the VPC with optional tags
        :param pulumi.Input[str] vpc_name: VPC name given when hosted at the cloud provider
        """
        if name is not None:
            pulumi.set(__self__, "name", name)
        if region is not None:
            pulumi.set(__self__, "region", region)
        if subnet is not None:
            pulumi.set(__self__, "subnet", subnet)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if vpc_name is not None:
            pulumi.set(__self__, "vpc_name", vpc_name)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the VPC.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        """
        The hosted region for the managed standalone VPC
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter
    def subnet(self) -> Optional[pulumi.Input[str]]:
        """
        The VPC subnet
        """
        return pulumi.get(self, "subnet")

    @subnet.setter
    def subnet(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "subnet", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Tag the VPC with optional tags
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)

    @property
    @pulumi.getter(name="vpcName")
    def vpc_name(self) -> Optional[pulumi.Input[str]]:
        """
        VPC name given when hosted at the cloud provider
        """
        return pulumi.get(self, "vpc_name")

    @vpc_name.setter
    def vpc_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vpc_name", value)


class Vpc(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 subnet: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 __props__=None):
        """
        This resource allows you to manage standalone VPC.

        New Cloudamqp instances can be added to the managed VPC. Set the instance *vpc_id* attribute to the managed vpc identifier , see example below, when creating the instance.

        Only available for dedicated subscription plans.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_cloudamqp as cloudamqp

        # Managed VPC resource
        vpc = cloudamqp.Vpc("vpc",
            region="amazon-web-services::us-east-1",
            subnet="10.56.72.0/24",
            tags=[])
        #  New instance, need to be created with a vpc
        instance = cloudamqp.Instance("instance",
            plan="bunny-1",
            region="amazon-web-services::us-east-1",
            nodes=1,
            tags=[],
            rmq_version="3.9.13",
            vpc_id=cloudamq_vpc["vpc"]["id"])
        vpc_info = cloudamqp.get_vpc_info_output(vpc_id=vpc.id)
        ```

        ## Import

        `cloudamqp_vpc` can be imported using the CloudAMQP VPC identifier.

        ```sh
         $ pulumi import cloudamqp:index/vpc:Vpc <resource_name> <vpc_id>`
        ```

         To retrieve the identifier for a VPC, either use [CloudAMQP customer API](https://docs.cloudamqp.com/#list-vpcs). Or use the data source [`cloudamqp_account_vpcs`](https://registry.terraform.io/providers/cloudamqp/cloudamqp/latest/docs/data-sources/account_vpcs) to list all available standalone VPCs for an account.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] name: The name of the VPC.
        :param pulumi.Input[str] region: The hosted region for the managed standalone VPC
        :param pulumi.Input[str] subnet: The VPC subnet
        :param pulumi.Input[Sequence[pulumi.Input[str]]] tags: Tag the VPC with optional tags
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: VpcArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        This resource allows you to manage standalone VPC.

        New Cloudamqp instances can be added to the managed VPC. Set the instance *vpc_id* attribute to the managed vpc identifier , see example below, when creating the instance.

        Only available for dedicated subscription plans.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_cloudamqp as cloudamqp

        # Managed VPC resource
        vpc = cloudamqp.Vpc("vpc",
            region="amazon-web-services::us-east-1",
            subnet="10.56.72.0/24",
            tags=[])
        #  New instance, need to be created with a vpc
        instance = cloudamqp.Instance("instance",
            plan="bunny-1",
            region="amazon-web-services::us-east-1",
            nodes=1,
            tags=[],
            rmq_version="3.9.13",
            vpc_id=cloudamq_vpc["vpc"]["id"])
        vpc_info = cloudamqp.get_vpc_info_output(vpc_id=vpc.id)
        ```

        ## Import

        `cloudamqp_vpc` can be imported using the CloudAMQP VPC identifier.

        ```sh
         $ pulumi import cloudamqp:index/vpc:Vpc <resource_name> <vpc_id>`
        ```

         To retrieve the identifier for a VPC, either use [CloudAMQP customer API](https://docs.cloudamqp.com/#list-vpcs). Or use the data source [`cloudamqp_account_vpcs`](https://registry.terraform.io/providers/cloudamqp/cloudamqp/latest/docs/data-sources/account_vpcs) to list all available standalone VPCs for an account.

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
                 name: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 subnet: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = VpcArgs.__new__(VpcArgs)

            __props__.__dict__["name"] = name
            if region is None and not opts.urn:
                raise TypeError("Missing required property 'region'")
            __props__.__dict__["region"] = region
            if subnet is None and not opts.urn:
                raise TypeError("Missing required property 'subnet'")
            __props__.__dict__["subnet"] = subnet
            __props__.__dict__["tags"] = tags
            __props__.__dict__["vpc_name"] = None
        super(Vpc, __self__).__init__(
            'cloudamqp:index/vpc:Vpc',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            name: Optional[pulumi.Input[str]] = None,
            region: Optional[pulumi.Input[str]] = None,
            subnet: Optional[pulumi.Input[str]] = None,
            tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            vpc_name: Optional[pulumi.Input[str]] = None) -> 'Vpc':
        """
        Get an existing Vpc resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] name: The name of the VPC.
        :param pulumi.Input[str] region: The hosted region for the managed standalone VPC
        :param pulumi.Input[str] subnet: The VPC subnet
        :param pulumi.Input[Sequence[pulumi.Input[str]]] tags: Tag the VPC with optional tags
        :param pulumi.Input[str] vpc_name: VPC name given when hosted at the cloud provider
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _VpcState.__new__(_VpcState)

        __props__.__dict__["name"] = name
        __props__.__dict__["region"] = region
        __props__.__dict__["subnet"] = subnet
        __props__.__dict__["tags"] = tags
        __props__.__dict__["vpc_name"] = vpc_name
        return Vpc(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the VPC.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def region(self) -> pulumi.Output[str]:
        """
        The hosted region for the managed standalone VPC
        """
        return pulumi.get(self, "region")

    @property
    @pulumi.getter
    def subnet(self) -> pulumi.Output[str]:
        """
        The VPC subnet
        """
        return pulumi.get(self, "subnet")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        Tag the VPC with optional tags
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="vpcName")
    def vpc_name(self) -> pulumi.Output[str]:
        """
        VPC name given when hosted at the cloud provider
        """
        return pulumi.get(self, "vpc_name")

