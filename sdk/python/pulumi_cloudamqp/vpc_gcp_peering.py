# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['VpcGcpPeeringArgs', 'VpcGcpPeering']

@pulumi.input_type
class VpcGcpPeeringArgs:
    def __init__(__self__, *,
                 peer_network_uri: pulumi.Input[str],
                 instance_id: Optional[pulumi.Input[int]] = None,
                 vpc_id: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a VpcGcpPeering resource.
        :param pulumi.Input[str] peer_network_uri: Network uri of the VPC network to which you will peer with.
        :param pulumi.Input[int] instance_id: The CloudAMQP instance identifier.
        :param pulumi.Input[str] vpc_id: The managed VPC identifier.
        """
        pulumi.set(__self__, "peer_network_uri", peer_network_uri)
        if instance_id is not None:
            pulumi.set(__self__, "instance_id", instance_id)
        if vpc_id is not None:
            pulumi.set(__self__, "vpc_id", vpc_id)

    @property
    @pulumi.getter(name="peerNetworkUri")
    def peer_network_uri(self) -> pulumi.Input[str]:
        """
        Network uri of the VPC network to which you will peer with.
        """
        return pulumi.get(self, "peer_network_uri")

    @peer_network_uri.setter
    def peer_network_uri(self, value: pulumi.Input[str]):
        pulumi.set(self, "peer_network_uri", value)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> Optional[pulumi.Input[int]]:
        """
        The CloudAMQP instance identifier.
        """
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "instance_id", value)

    @property
    @pulumi.getter(name="vpcId")
    def vpc_id(self) -> Optional[pulumi.Input[str]]:
        """
        The managed VPC identifier.
        """
        return pulumi.get(self, "vpc_id")

    @vpc_id.setter
    def vpc_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vpc_id", value)


@pulumi.input_type
class _VpcGcpPeeringState:
    def __init__(__self__, *,
                 auto_create_routes: Optional[pulumi.Input[bool]] = None,
                 instance_id: Optional[pulumi.Input[int]] = None,
                 peer_network_uri: Optional[pulumi.Input[str]] = None,
                 state: Optional[pulumi.Input[str]] = None,
                 state_details: Optional[pulumi.Input[str]] = None,
                 vpc_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering VpcGcpPeering resources.
        :param pulumi.Input[bool] auto_create_routes: VPC peering auto created routes
        :param pulumi.Input[int] instance_id: The CloudAMQP instance identifier.
        :param pulumi.Input[str] peer_network_uri: Network uri of the VPC network to which you will peer with.
        :param pulumi.Input[str] state: VPC peering state
        :param pulumi.Input[str] state_details: VPC peering state details
        :param pulumi.Input[str] vpc_id: The managed VPC identifier.
        """
        if auto_create_routes is not None:
            pulumi.set(__self__, "auto_create_routes", auto_create_routes)
        if instance_id is not None:
            pulumi.set(__self__, "instance_id", instance_id)
        if peer_network_uri is not None:
            pulumi.set(__self__, "peer_network_uri", peer_network_uri)
        if state is not None:
            pulumi.set(__self__, "state", state)
        if state_details is not None:
            pulumi.set(__self__, "state_details", state_details)
        if vpc_id is not None:
            pulumi.set(__self__, "vpc_id", vpc_id)

    @property
    @pulumi.getter(name="autoCreateRoutes")
    def auto_create_routes(self) -> Optional[pulumi.Input[bool]]:
        """
        VPC peering auto created routes
        """
        return pulumi.get(self, "auto_create_routes")

    @auto_create_routes.setter
    def auto_create_routes(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "auto_create_routes", value)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> Optional[pulumi.Input[int]]:
        """
        The CloudAMQP instance identifier.
        """
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "instance_id", value)

    @property
    @pulumi.getter(name="peerNetworkUri")
    def peer_network_uri(self) -> Optional[pulumi.Input[str]]:
        """
        Network uri of the VPC network to which you will peer with.
        """
        return pulumi.get(self, "peer_network_uri")

    @peer_network_uri.setter
    def peer_network_uri(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "peer_network_uri", value)

    @property
    @pulumi.getter
    def state(self) -> Optional[pulumi.Input[str]]:
        """
        VPC peering state
        """
        return pulumi.get(self, "state")

    @state.setter
    def state(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "state", value)

    @property
    @pulumi.getter(name="stateDetails")
    def state_details(self) -> Optional[pulumi.Input[str]]:
        """
        VPC peering state details
        """
        return pulumi.get(self, "state_details")

    @state_details.setter
    def state_details(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "state_details", value)

    @property
    @pulumi.getter(name="vpcId")
    def vpc_id(self) -> Optional[pulumi.Input[str]]:
        """
        The managed VPC identifier.
        """
        return pulumi.get(self, "vpc_id")

    @vpc_id.setter
    def vpc_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vpc_id", value)


class VpcGcpPeering(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 instance_id: Optional[pulumi.Input[int]] = None,
                 peer_network_uri: Optional[pulumi.Input[str]] = None,
                 vpc_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Create a VpcGcpPeering resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] instance_id: The CloudAMQP instance identifier.
        :param pulumi.Input[str] peer_network_uri: Network uri of the VPC network to which you will peer with.
        :param pulumi.Input[str] vpc_id: The managed VPC identifier.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: VpcGcpPeeringArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a VpcGcpPeering resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param VpcGcpPeeringArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(VpcGcpPeeringArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 instance_id: Optional[pulumi.Input[int]] = None,
                 peer_network_uri: Optional[pulumi.Input[str]] = None,
                 vpc_id: Optional[pulumi.Input[str]] = None,
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
            __props__ = VpcGcpPeeringArgs.__new__(VpcGcpPeeringArgs)

            __props__.__dict__["instance_id"] = instance_id
            if peer_network_uri is None and not opts.urn:
                raise TypeError("Missing required property 'peer_network_uri'")
            __props__.__dict__["peer_network_uri"] = peer_network_uri
            __props__.__dict__["vpc_id"] = vpc_id
            __props__.__dict__["auto_create_routes"] = None
            __props__.__dict__["state"] = None
            __props__.__dict__["state_details"] = None
        super(VpcGcpPeering, __self__).__init__(
            'cloudamqp:index/vpcGcpPeering:VpcGcpPeering',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            auto_create_routes: Optional[pulumi.Input[bool]] = None,
            instance_id: Optional[pulumi.Input[int]] = None,
            peer_network_uri: Optional[pulumi.Input[str]] = None,
            state: Optional[pulumi.Input[str]] = None,
            state_details: Optional[pulumi.Input[str]] = None,
            vpc_id: Optional[pulumi.Input[str]] = None) -> 'VpcGcpPeering':
        """
        Get an existing VpcGcpPeering resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] auto_create_routes: VPC peering auto created routes
        :param pulumi.Input[int] instance_id: The CloudAMQP instance identifier.
        :param pulumi.Input[str] peer_network_uri: Network uri of the VPC network to which you will peer with.
        :param pulumi.Input[str] state: VPC peering state
        :param pulumi.Input[str] state_details: VPC peering state details
        :param pulumi.Input[str] vpc_id: The managed VPC identifier.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _VpcGcpPeeringState.__new__(_VpcGcpPeeringState)

        __props__.__dict__["auto_create_routes"] = auto_create_routes
        __props__.__dict__["instance_id"] = instance_id
        __props__.__dict__["peer_network_uri"] = peer_network_uri
        __props__.__dict__["state"] = state
        __props__.__dict__["state_details"] = state_details
        __props__.__dict__["vpc_id"] = vpc_id
        return VpcGcpPeering(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="autoCreateRoutes")
    def auto_create_routes(self) -> pulumi.Output[bool]:
        """
        VPC peering auto created routes
        """
        return pulumi.get(self, "auto_create_routes")

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> pulumi.Output[Optional[int]]:
        """
        The CloudAMQP instance identifier.
        """
        return pulumi.get(self, "instance_id")

    @property
    @pulumi.getter(name="peerNetworkUri")
    def peer_network_uri(self) -> pulumi.Output[str]:
        """
        Network uri of the VPC network to which you will peer with.
        """
        return pulumi.get(self, "peer_network_uri")

    @property
    @pulumi.getter
    def state(self) -> pulumi.Output[str]:
        """
        VPC peering state
        """
        return pulumi.get(self, "state")

    @property
    @pulumi.getter(name="stateDetails")
    def state_details(self) -> pulumi.Output[str]:
        """
        VPC peering state details
        """
        return pulumi.get(self, "state_details")

    @property
    @pulumi.getter(name="vpcId")
    def vpc_id(self) -> pulumi.Output[Optional[str]]:
        """
        The managed VPC identifier.
        """
        return pulumi.get(self, "vpc_id")

