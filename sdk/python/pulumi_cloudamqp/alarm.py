# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from . import utilities, tables


class Alarm(pulumi.CustomResource):
    enabled: pulumi.Output[bool]
    """
    Enable or disable an alarm
    """
    instance_id: pulumi.Output[float]
    """
    Instance identifier
    """
    message_type: pulumi.Output[str]
    """
    Message types (total, unacked, ready) of the queue to trigger the alarm
    """
    queue_regex: pulumi.Output[str]
    """
    Regex for which queues to check
    """
    recipients: pulumi.Output[list]
    """
    Identifiers for recipients to be notified.
    """
    time_threshold: pulumi.Output[float]
    """
    For how long (in seconds) the value_threshold should be active before trigger alarm
    """
    type: pulumi.Output[str]
    """
    Type of the alarm, valid options are: cpu, memory, disk_usage, queue_length, connection_count, consumers_count,
    net_split
    """
    value_threshold: pulumi.Output[float]
    """
    What value to trigger the alarm for
    """
    vhost_regex: pulumi.Output[str]
    """
    Regex for which vhost the queues are in
    """
    def __init__(__self__, resource_name, opts=None, enabled=None, instance_id=None, message_type=None, queue_regex=None, recipients=None, time_threshold=None, type=None, value_threshold=None, vhost_regex=None, __props__=None, __name__=None, __opts__=None):
        """
        Create a Alarm resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] enabled: Enable or disable an alarm
        :param pulumi.Input[float] instance_id: Instance identifier
        :param pulumi.Input[str] message_type: Message types (total, unacked, ready) of the queue to trigger the alarm
        :param pulumi.Input[str] queue_regex: Regex for which queues to check
        :param pulumi.Input[list] recipients: Identifiers for recipients to be notified.
        :param pulumi.Input[float] time_threshold: For how long (in seconds) the value_threshold should be active before trigger alarm
        :param pulumi.Input[str] type: Type of the alarm, valid options are: cpu, memory, disk_usage, queue_length, connection_count, consumers_count,
               net_split
        :param pulumi.Input[float] value_threshold: What value to trigger the alarm for
        :param pulumi.Input[str] vhost_regex: Regex for which vhost the queues are in
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
            opts.version = utilities.get_version()
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
            __props__['message_type'] = message_type
            __props__['queue_regex'] = queue_regex
            if recipients is None:
                raise TypeError("Missing required property 'recipients'")
            __props__['recipients'] = recipients
            __props__['time_threshold'] = time_threshold
            if type is None:
                raise TypeError("Missing required property 'type'")
            __props__['type'] = type
            __props__['value_threshold'] = value_threshold
            __props__['vhost_regex'] = vhost_regex
        super(Alarm, __self__).__init__(
            'cloudamqp:index/alarm:Alarm',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, enabled=None, instance_id=None, message_type=None, queue_regex=None, recipients=None, time_threshold=None, type=None, value_threshold=None, vhost_regex=None):
        """
        Get an existing Alarm resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] enabled: Enable or disable an alarm
        :param pulumi.Input[float] instance_id: Instance identifier
        :param pulumi.Input[str] message_type: Message types (total, unacked, ready) of the queue to trigger the alarm
        :param pulumi.Input[str] queue_regex: Regex for which queues to check
        :param pulumi.Input[list] recipients: Identifiers for recipients to be notified.
        :param pulumi.Input[float] time_threshold: For how long (in seconds) the value_threshold should be active before trigger alarm
        :param pulumi.Input[str] type: Type of the alarm, valid options are: cpu, memory, disk_usage, queue_length, connection_count, consumers_count,
               net_split
        :param pulumi.Input[float] value_threshold: What value to trigger the alarm for
        :param pulumi.Input[str] vhost_regex: Regex for which vhost the queues are in
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["enabled"] = enabled
        __props__["instance_id"] = instance_id
        __props__["message_type"] = message_type
        __props__["queue_regex"] = queue_regex
        __props__["recipients"] = recipients
        __props__["time_threshold"] = time_threshold
        __props__["type"] = type
        __props__["value_threshold"] = value_threshold
        __props__["vhost_regex"] = vhost_regex
        return Alarm(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
