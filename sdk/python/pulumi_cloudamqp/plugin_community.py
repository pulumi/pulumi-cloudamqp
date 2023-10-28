# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['PluginCommunityArgs', 'PluginCommunity']

@pulumi.input_type
class PluginCommunityArgs:
    def __init__(__self__, *,
                 enabled: pulumi.Input[bool],
                 instance_id: pulumi.Input[int],
                 name: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a PluginCommunity resource.
        :param pulumi.Input[bool] enabled: Enable or disable the plugins.
        :param pulumi.Input[int] instance_id: The CloudAMQP instance ID.
        :param pulumi.Input[str] name: The name of the Rabbit MQ community plugin.
        """
        pulumi.set(__self__, "enabled", enabled)
        pulumi.set(__self__, "instance_id", instance_id)
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def enabled(self) -> pulumi.Input[bool]:
        """
        Enable or disable the plugins.
        """
        return pulumi.get(self, "enabled")

    @enabled.setter
    def enabled(self, value: pulumi.Input[bool]):
        pulumi.set(self, "enabled", value)

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
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the Rabbit MQ community plugin.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


@pulumi.input_type
class _PluginCommunityState:
    def __init__(__self__, *,
                 description: Optional[pulumi.Input[str]] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 instance_id: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 require: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering PluginCommunity resources.
        :param pulumi.Input[str] description: The description of the plugin.
        :param pulumi.Input[bool] enabled: Enable or disable the plugins.
        :param pulumi.Input[int] instance_id: The CloudAMQP instance ID.
        :param pulumi.Input[str] name: The name of the Rabbit MQ community plugin.
        :param pulumi.Input[str] require: Required version of RabbitMQ.
        """
        if description is not None:
            pulumi.set(__self__, "description", description)
        if enabled is not None:
            pulumi.set(__self__, "enabled", enabled)
        if instance_id is not None:
            pulumi.set(__self__, "instance_id", instance_id)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if require is not None:
            pulumi.set(__self__, "require", require)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The description of the plugin.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        Enable or disable the plugins.
        """
        return pulumi.get(self, "enabled")

    @enabled.setter
    def enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enabled", value)

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
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the Rabbit MQ community plugin.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def require(self) -> Optional[pulumi.Input[str]]:
        """
        Required version of RabbitMQ.
        """
        return pulumi.get(self, "require")

    @require.setter
    def require(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "require", value)


class PluginCommunity(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 instance_id: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        ## Import

        `cloudamqp_plugin` can be imported using the name argument of the resource together with CloudAMQP instance identifier. The name and identifier are CSV separated, see example below.

        ```sh
         $ pulumi import cloudamqp:index/pluginCommunity:PluginCommunity <resource_name> <plugin_name>,<instance_id>`
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] enabled: Enable or disable the plugins.
        :param pulumi.Input[int] instance_id: The CloudAMQP instance ID.
        :param pulumi.Input[str] name: The name of the Rabbit MQ community plugin.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: PluginCommunityArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        ## Import

        `cloudamqp_plugin` can be imported using the name argument of the resource together with CloudAMQP instance identifier. The name and identifier are CSV separated, see example below.

        ```sh
         $ pulumi import cloudamqp:index/pluginCommunity:PluginCommunity <resource_name> <plugin_name>,<instance_id>`
        ```

        :param str resource_name: The name of the resource.
        :param PluginCommunityArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(PluginCommunityArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 instance_id: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = PluginCommunityArgs.__new__(PluginCommunityArgs)

            if enabled is None and not opts.urn:
                raise TypeError("Missing required property 'enabled'")
            __props__.__dict__["enabled"] = enabled
            if instance_id is None and not opts.urn:
                raise TypeError("Missing required property 'instance_id'")
            __props__.__dict__["instance_id"] = instance_id
            __props__.__dict__["name"] = name
            __props__.__dict__["description"] = None
            __props__.__dict__["require"] = None
        super(PluginCommunity, __self__).__init__(
            'cloudamqp:index/pluginCommunity:PluginCommunity',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            description: Optional[pulumi.Input[str]] = None,
            enabled: Optional[pulumi.Input[bool]] = None,
            instance_id: Optional[pulumi.Input[int]] = None,
            name: Optional[pulumi.Input[str]] = None,
            require: Optional[pulumi.Input[str]] = None) -> 'PluginCommunity':
        """
        Get an existing PluginCommunity resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: The description of the plugin.
        :param pulumi.Input[bool] enabled: Enable or disable the plugins.
        :param pulumi.Input[int] instance_id: The CloudAMQP instance ID.
        :param pulumi.Input[str] name: The name of the Rabbit MQ community plugin.
        :param pulumi.Input[str] require: Required version of RabbitMQ.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _PluginCommunityState.__new__(_PluginCommunityState)

        __props__.__dict__["description"] = description
        __props__.__dict__["enabled"] = enabled
        __props__.__dict__["instance_id"] = instance_id
        __props__.__dict__["name"] = name
        __props__.__dict__["require"] = require
        return PluginCommunity(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[str]:
        """
        The description of the plugin.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def enabled(self) -> pulumi.Output[bool]:
        """
        Enable or disable the plugins.
        """
        return pulumi.get(self, "enabled")

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> pulumi.Output[int]:
        """
        The CloudAMQP instance ID.
        """
        return pulumi.get(self, "instance_id")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the Rabbit MQ community plugin.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def require(self) -> pulumi.Output[str]:
        """
        Required version of RabbitMQ.
        """
        return pulumi.get(self, "require")

