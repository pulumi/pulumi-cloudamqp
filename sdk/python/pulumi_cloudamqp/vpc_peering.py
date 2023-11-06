# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Callable, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['VpcPeeringArgs', 'VpcPeering']

@pulumi.input_type
class VpcPeeringArgs:
    def __init__(__self__, *,
                 peering_id: pulumi.Input[str],
                 instance_id: Optional[pulumi.Input[int]] = None,
                 sleep: Optional[pulumi.Input[int]] = None,
                 timeout: Optional[pulumi.Input[int]] = None,
                 vpc_id: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a VpcPeering resource.
        :param pulumi.Input[str] peering_id: Peering identifier created by AW peering request.
        :param pulumi.Input[int] instance_id: The CloudAMQP instance identifier.
               
               ***Deprecated: Changed from required to optional in v1.16.0, will be removed in next major version (v2.0)***
        :param pulumi.Input[int] sleep: Configurable sleep time (seconds) between retries for accepting or removing peering. Default set to 60 seconds.
        :param pulumi.Input[int] timeout: Configurable timeout time (seconds) for accepting or removing peering. Default set to 3600 seconds.
        :param pulumi.Input[str] vpc_id: The managed VPC identifier.
               
               ***Note: Introduced as optional in version v1.16.0, will be required in next major version (v2.0)***
        """
        VpcPeeringArgs._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            peering_id=peering_id,
            instance_id=instance_id,
            sleep=sleep,
            timeout=timeout,
            vpc_id=vpc_id,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             peering_id: Optional[pulumi.Input[str]] = None,
             instance_id: Optional[pulumi.Input[int]] = None,
             sleep: Optional[pulumi.Input[int]] = None,
             timeout: Optional[pulumi.Input[int]] = None,
             vpc_id: Optional[pulumi.Input[str]] = None,
             opts: Optional[pulumi.ResourceOptions]=None,
             **kwargs):
        if peering_id is None and 'peeringId' in kwargs:
            peering_id = kwargs['peeringId']
        if peering_id is None:
            raise TypeError("Missing 'peering_id' argument")
        if instance_id is None and 'instanceId' in kwargs:
            instance_id = kwargs['instanceId']
        if vpc_id is None and 'vpcId' in kwargs:
            vpc_id = kwargs['vpcId']

        _setter("peering_id", peering_id)
        if instance_id is not None:
            _setter("instance_id", instance_id)
        if sleep is not None:
            _setter("sleep", sleep)
        if timeout is not None:
            _setter("timeout", timeout)
        if vpc_id is not None:
            _setter("vpc_id", vpc_id)

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

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> Optional[pulumi.Input[int]]:
        """
        The CloudAMQP instance identifier.

        ***Deprecated: Changed from required to optional in v1.16.0, will be removed in next major version (v2.0)***
        """
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "instance_id", value)

    @property
    @pulumi.getter
    def sleep(self) -> Optional[pulumi.Input[int]]:
        """
        Configurable sleep time (seconds) between retries for accepting or removing peering. Default set to 60 seconds.
        """
        return pulumi.get(self, "sleep")

    @sleep.setter
    def sleep(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "sleep", value)

    @property
    @pulumi.getter
    def timeout(self) -> Optional[pulumi.Input[int]]:
        """
        Configurable timeout time (seconds) for accepting or removing peering. Default set to 3600 seconds.
        """
        return pulumi.get(self, "timeout")

    @timeout.setter
    def timeout(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "timeout", value)

    @property
    @pulumi.getter(name="vpcId")
    def vpc_id(self) -> Optional[pulumi.Input[str]]:
        """
        The managed VPC identifier.

        ***Note: Introduced as optional in version v1.16.0, will be required in next major version (v2.0)***
        """
        return pulumi.get(self, "vpc_id")

    @vpc_id.setter
    def vpc_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vpc_id", value)


@pulumi.input_type
class _VpcPeeringState:
    def __init__(__self__, *,
                 instance_id: Optional[pulumi.Input[int]] = None,
                 peering_id: Optional[pulumi.Input[str]] = None,
                 sleep: Optional[pulumi.Input[int]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 timeout: Optional[pulumi.Input[int]] = None,
                 vpc_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering VpcPeering resources.
        :param pulumi.Input[int] instance_id: The CloudAMQP instance identifier.
               
               ***Deprecated: Changed from required to optional in v1.16.0, will be removed in next major version (v2.0)***
        :param pulumi.Input[str] peering_id: Peering identifier created by AW peering request.
        :param pulumi.Input[int] sleep: Configurable sleep time (seconds) between retries for accepting or removing peering. Default set to 60 seconds.
        :param pulumi.Input[str] status: VPC peering status
        :param pulumi.Input[int] timeout: Configurable timeout time (seconds) for accepting or removing peering. Default set to 3600 seconds.
        :param pulumi.Input[str] vpc_id: The managed VPC identifier.
               
               ***Note: Introduced as optional in version v1.16.0, will be required in next major version (v2.0)***
        """
        _VpcPeeringState._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            instance_id=instance_id,
            peering_id=peering_id,
            sleep=sleep,
            status=status,
            timeout=timeout,
            vpc_id=vpc_id,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             instance_id: Optional[pulumi.Input[int]] = None,
             peering_id: Optional[pulumi.Input[str]] = None,
             sleep: Optional[pulumi.Input[int]] = None,
             status: Optional[pulumi.Input[str]] = None,
             timeout: Optional[pulumi.Input[int]] = None,
             vpc_id: Optional[pulumi.Input[str]] = None,
             opts: Optional[pulumi.ResourceOptions]=None,
             **kwargs):
        if instance_id is None and 'instanceId' in kwargs:
            instance_id = kwargs['instanceId']
        if peering_id is None and 'peeringId' in kwargs:
            peering_id = kwargs['peeringId']
        if vpc_id is None and 'vpcId' in kwargs:
            vpc_id = kwargs['vpcId']

        if instance_id is not None:
            _setter("instance_id", instance_id)
        if peering_id is not None:
            _setter("peering_id", peering_id)
        if sleep is not None:
            _setter("sleep", sleep)
        if status is not None:
            _setter("status", status)
        if timeout is not None:
            _setter("timeout", timeout)
        if vpc_id is not None:
            _setter("vpc_id", vpc_id)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> Optional[pulumi.Input[int]]:
        """
        The CloudAMQP instance identifier.

        ***Deprecated: Changed from required to optional in v1.16.0, will be removed in next major version (v2.0)***
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
    def sleep(self) -> Optional[pulumi.Input[int]]:
        """
        Configurable sleep time (seconds) between retries for accepting or removing peering. Default set to 60 seconds.
        """
        return pulumi.get(self, "sleep")

    @sleep.setter
    def sleep(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "sleep", value)

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

    @property
    @pulumi.getter
    def timeout(self) -> Optional[pulumi.Input[int]]:
        """
        Configurable timeout time (seconds) for accepting or removing peering. Default set to 3600 seconds.
        """
        return pulumi.get(self, "timeout")

    @timeout.setter
    def timeout(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "timeout", value)

    @property
    @pulumi.getter(name="vpcId")
    def vpc_id(self) -> Optional[pulumi.Input[str]]:
        """
        The managed VPC identifier.

        ***Note: Introduced as optional in version v1.16.0, will be required in next major version (v2.0)***
        """
        return pulumi.get(self, "vpc_id")

    @vpc_id.setter
    def vpc_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vpc_id", value)


class VpcPeering(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 instance_id: Optional[pulumi.Input[int]] = None,
                 peering_id: Optional[pulumi.Input[str]] = None,
                 sleep: Optional[pulumi.Input[int]] = None,
                 timeout: Optional[pulumi.Input[int]] = None,
                 vpc_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        ## Import

        Not possible to import this resource.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] instance_id: The CloudAMQP instance identifier.
               
               ***Deprecated: Changed from required to optional in v1.16.0, will be removed in next major version (v2.0)***
        :param pulumi.Input[str] peering_id: Peering identifier created by AW peering request.
        :param pulumi.Input[int] sleep: Configurable sleep time (seconds) between retries for accepting or removing peering. Default set to 60 seconds.
        :param pulumi.Input[int] timeout: Configurable timeout time (seconds) for accepting or removing peering. Default set to 3600 seconds.
        :param pulumi.Input[str] vpc_id: The managed VPC identifier.
               
               ***Note: Introduced as optional in version v1.16.0, will be required in next major version (v2.0)***
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: VpcPeeringArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        ## Import

        Not possible to import this resource.

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
            kwargs = kwargs or {}
            def _setter(key, value):
                kwargs[key] = value
            VpcPeeringArgs._configure(_setter, **kwargs)
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 instance_id: Optional[pulumi.Input[int]] = None,
                 peering_id: Optional[pulumi.Input[str]] = None,
                 sleep: Optional[pulumi.Input[int]] = None,
                 timeout: Optional[pulumi.Input[int]] = None,
                 vpc_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = VpcPeeringArgs.__new__(VpcPeeringArgs)

            __props__.__dict__["instance_id"] = instance_id
            if peering_id is None and not opts.urn:
                raise TypeError("Missing required property 'peering_id'")
            __props__.__dict__["peering_id"] = peering_id
            __props__.__dict__["sleep"] = sleep
            __props__.__dict__["timeout"] = timeout
            __props__.__dict__["vpc_id"] = vpc_id
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
            sleep: Optional[pulumi.Input[int]] = None,
            status: Optional[pulumi.Input[str]] = None,
            timeout: Optional[pulumi.Input[int]] = None,
            vpc_id: Optional[pulumi.Input[str]] = None) -> 'VpcPeering':
        """
        Get an existing VpcPeering resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] instance_id: The CloudAMQP instance identifier.
               
               ***Deprecated: Changed from required to optional in v1.16.0, will be removed in next major version (v2.0)***
        :param pulumi.Input[str] peering_id: Peering identifier created by AW peering request.
        :param pulumi.Input[int] sleep: Configurable sleep time (seconds) between retries for accepting or removing peering. Default set to 60 seconds.
        :param pulumi.Input[str] status: VPC peering status
        :param pulumi.Input[int] timeout: Configurable timeout time (seconds) for accepting or removing peering. Default set to 3600 seconds.
        :param pulumi.Input[str] vpc_id: The managed VPC identifier.
               
               ***Note: Introduced as optional in version v1.16.0, will be required in next major version (v2.0)***
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _VpcPeeringState.__new__(_VpcPeeringState)

        __props__.__dict__["instance_id"] = instance_id
        __props__.__dict__["peering_id"] = peering_id
        __props__.__dict__["sleep"] = sleep
        __props__.__dict__["status"] = status
        __props__.__dict__["timeout"] = timeout
        __props__.__dict__["vpc_id"] = vpc_id
        return VpcPeering(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> pulumi.Output[Optional[int]]:
        """
        The CloudAMQP instance identifier.

        ***Deprecated: Changed from required to optional in v1.16.0, will be removed in next major version (v2.0)***
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
    def sleep(self) -> pulumi.Output[Optional[int]]:
        """
        Configurable sleep time (seconds) between retries for accepting or removing peering. Default set to 60 seconds.
        """
        return pulumi.get(self, "sleep")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[str]:
        """
        VPC peering status
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter
    def timeout(self) -> pulumi.Output[Optional[int]]:
        """
        Configurable timeout time (seconds) for accepting or removing peering. Default set to 3600 seconds.
        """
        return pulumi.get(self, "timeout")

    @property
    @pulumi.getter(name="vpcId")
    def vpc_id(self) -> pulumi.Output[Optional[str]]:
        """
        The managed VPC identifier.

        ***Note: Introduced as optional in version v1.16.0, will be required in next major version (v2.0)***
        """
        return pulumi.get(self, "vpc_id")

