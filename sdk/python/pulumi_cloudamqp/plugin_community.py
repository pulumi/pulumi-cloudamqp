# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from . import _utilities, _tables

__all__ = ['PluginCommunity']


class PluginCommunity(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 instance_id: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        This resource allows you to install or uninstall community plugins. Once installed the plugin will be available in `Plugin`.

        Only available for dedicated subscription plans.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_cloudamqp as cloudamqp

        rabbitmq_delayed_message_exchange = cloudamqp.PluginCommunity("rabbitmqDelayedMessageExchange",
            instance_id=cloudamqp_instance["instance_01"]["id"],
            enabled=True)
        ```
        ## Depedency

        This resource depends on CloudAMQP instance identifier, `cloudamqp_instance.instance.id`.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] enabled: Enable or disable the plugins.
        :param pulumi.Input[int] instance_id: The CloudAMQP instance ID.
        :param pulumi.Input[str] name: The name of the Rabbit MQ plugin.
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            if enabled is None:
                raise TypeError("Missing required property 'enabled'")
            __props__['enabled'] = enabled
            if instance_id is None:
                raise TypeError("Missing required property 'instance_id'")
            __props__['instance_id'] = instance_id
            __props__['name'] = name
        super(PluginCommunity, __self__).__init__(
            'cloudamqp:index/pluginCommunity:PluginCommunity',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            enabled: Optional[pulumi.Input[bool]] = None,
            instance_id: Optional[pulumi.Input[int]] = None,
            name: Optional[pulumi.Input[str]] = None) -> 'PluginCommunity':
        """
        Get an existing PluginCommunity resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] enabled: Enable or disable the plugins.
        :param pulumi.Input[int] instance_id: The CloudAMQP instance ID.
        :param pulumi.Input[str] name: The name of the Rabbit MQ plugin.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["enabled"] = enabled
        __props__["instance_id"] = instance_id
        __props__["name"] = name
        return PluginCommunity(resource_name, opts=opts, __props__=__props__)

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
        The name of the Rabbit MQ plugin.
        """
        return pulumi.get(self, "name")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

