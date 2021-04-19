# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['VpcPeeringArgs', 'VpcPeering']

@pulumi.input_type
class VpcPeeringArgs:
    def __init__(__self__, *,
                 instance_id: pulumi.Input[int],
                 peering_id: pulumi.Input[str]):
        """
        The set of arguments for constructing a VpcPeering resource.
        :param pulumi.Input[int] instance_id: The CloudAMQP instance ID.
        :param pulumi.Input[str] peering_id: Peering identifier created by AW peering request.
        """
        pulumi.set(__self__, "instance_id", instance_id)
        pulumi.set(__self__, "peering_id", peering_id)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> pulumi.Input[int]:
        """
        The CloudAMQP instance ID.
        """
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: pulumi.Input[int]):
        pulumi.set(self, "instance_id", value)

    @property
    @pulumi.getter(name="peeringId")
    def peering_id(self) -> pulumi.Input[str]:
        """
        Peering identifier created by AW peering request.
        """
        return pulumi.get(self, "peering_id")

    @peering_id.setter
    def peering_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "peering_id", value)


@pulumi.input_type
class _VpcPeeringState:
    def __init__(__self__, *,
                 instance_id: Optional[pulumi.Input[int]] = None,
                 peering_id: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering VpcPeering resources.
        :param pulumi.Input[int] instance_id: The CloudAMQP instance ID.
        :param pulumi.Input[str] peering_id: Peering identifier created by AW peering request.
        :param pulumi.Input[str] status: VPC peering status
        """
        if instance_id is not None:
            pulumi.set(__self__, "instance_id", instance_id)
        if peering_id is not None:
            pulumi.set(__self__, "peering_id", peering_id)
        if status is not None:
            pulumi.set(__self__, "status", status)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> Optional[pulumi.Input[int]]:
        """
        The CloudAMQP instance ID.
        """
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "instance_id", value)

    @property
    @pulumi.getter(name="peeringId")
    def peering_id(self) -> Optional[pulumi.Input[str]]:
        """
        Peering identifier created by AW peering request.
        """
        return pulumi.get(self, "peering_id")

    @peering_id.setter
    def peering_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "peering_id", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        """
        VPC peering status
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)


class VpcPeering(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 instance_id: Optional[pulumi.Input[int]] = None,
                 peering_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        ## Import

        `cloudamqp_vpc_peering` can be imported using the CloudAMQP instance identifier.

        ```sh
         $ pulumi import cloudamqp:index/vpcPeering:VpcPeering <resource_name> <instance_id>`
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] instance_id: The CloudAMQP instance ID.
        :param pulumi.Input[str] peering_id: Peering identifier created by AW peering request.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: VpcPeeringArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        ## Import

        `cloudamqp_vpc_peering` can be imported using the CloudAMQP instance identifier.

        ```sh
         $ pulumi import cloudamqp:index/vpcPeering:VpcPeering <resource_name> <instance_id>`
        ```

        :param str resource_name: The name of the resource.
        :param VpcPeeringArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(VpcPeeringArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 instance_id: Optional[pulumi.Input[int]] = None,
                 peering_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = VpcPeeringArgs.__new__(VpcPeeringArgs)

            if instance_id is None and not opts.urn:
                raise TypeError("Missing required property 'instance_id'")
            __props__.__dict__["instance_id"] = instance_id
            if peering_id is None and not opts.urn:
                raise TypeError("Missing required property 'peering_id'")
            __props__.__dict__["peering_id"] = peering_id
            __props__.__dict__["status"] = None
        super(VpcPeering, __self__).__init__(
            'cloudamqp:index/vpcPeering:VpcPeering',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            instance_id: Optional[pulumi.Input[int]] = None,
            peering_id: Optional[pulumi.Input[str]] = None,
            status: Optional[pulumi.Input[str]] = None) -> 'VpcPeering':
        """
        Get an existing VpcPeering resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] instance_id: The CloudAMQP instance ID.
        :param pulumi.Input[str] peering_id: Peering identifier created by AW peering request.
        :param pulumi.Input[str] status: VPC peering status
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _VpcPeeringState.__new__(_VpcPeeringState)

        __props__.__dict__["instance_id"] = instance_id
        __props__.__dict__["peering_id"] = peering_id
        __props__.__dict__["status"] = status
        return VpcPeering(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> pulumi.Output[int]:
        """
        The CloudAMQP instance ID.
        """
        return pulumi.get(self, "instance_id")

    @property
    @pulumi.getter(name="peeringId")
    def peering_id(self) -> pulumi.Output[str]:
        """
        Peering identifier created by AW peering request.
        """
        return pulumi.get(self, "peering_id")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[str]:
        """
        VPC peering status
        """
        return pulumi.get(self, "status")

