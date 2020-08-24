# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from . import _utilities, _tables

__all__ = ['Alarm']


class Alarm(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 instance_id: Optional[pulumi.Input[float]] = None,
                 message_type: Optional[pulumi.Input[str]] = None,
                 queue_regex: Optional[pulumi.Input[str]] = None,
                 recipients: Optional[pulumi.Input[List[pulumi.Input[float]]]] = None,
                 time_threshold: Optional[pulumi.Input[float]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 value_threshold: Optional[pulumi.Input[float]] = None,
                 vhost_regex: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        This resource allows you to create and manage alarms to trigger based on a set of conditions. Once triggerd a notification will be sent to the assigned recipients. When creating a new instance, there will also be a set of default alarms (cpu, memory and disk) created. All default alarms uses the default recipient for notifications.

        By setting `no_default_alarms` to *true* in `Instance`. This will create the instance without default alarms and avoid the need to import them to get full control.

        Available for all subscription plans, but `lemur`and `tiger`are limited to fewer alarm types. The limited types supported can be seen in the table below in Alarm Type Reference.

        ## Alarm Type reference

        Valid options for notification type.

        Required arguments for all alarms: *instance_id*, *type* and *enabled*<br>
        Optional argument for all alarms: *tags*, *queue_regex*, *vhost_regex*

        | Name | Type | Shared | Dedicated | Required arguments |
        | ---- | ---- | ---- | ---- | ---- | ---- |
        | CPU | cpu | - | &#10004; | time_threshold, value_threshold |
        | Memory | memory | - | &#10004;  | time_threshold, value_threshold |
        | Disk space | disk | - | &#10004;  | time_threshold, value_threshold |
        | Queue | queue | &#10004;  | &#10004;  | time_threshold, value_threshold, queue_regex, vhost_regex, message_type |
        | Connection | connection | &#10004; | &#10004; | time_threshold, value_threshold |
        | Consumer | consumer | &#10004; | &#10004; | time_threshold, value_threshold, queue, vhost |
        | Netsplit | netsplit | - | &#10004; | time_threshold |
        | Server unreachable | server_unreachable  | - | &#10004;  | time_threshold |
        | Notice | notice | &#10004; | &#10004; |

        ## Dependency

        This resource depends on CloudAMQP instance identifier, `cloudamqp_instance.instance.id`.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] enabled: Enable or disable the alarm to trigger.
        :param pulumi.Input[float] instance_id: The CloudAMQP instance ID.
        :param pulumi.Input[str] message_type: Message type `(total, unacked, ready)` used by queue alarm type.
        :param pulumi.Input[str] queue_regex: Regex for which queue to check.
        :param pulumi.Input[List[pulumi.Input[float]]] recipients: Identifier for recipient to be notified. Leave empty to notify all recipients.
        :param pulumi.Input[float] time_threshold: The time interval (in seconds) the `value_threshold` should be active before triggering an alarm.
        :param pulumi.Input[str] type: The alarm type, see valid options below.
        :param pulumi.Input[float] value_threshold: The value to trigger the alarm for.
        :param pulumi.Input[str] vhost_regex: Regex for which vhost to check
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
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            enabled: Optional[pulumi.Input[bool]] = None,
            instance_id: Optional[pulumi.Input[float]] = None,
            message_type: Optional[pulumi.Input[str]] = None,
            queue_regex: Optional[pulumi.Input[str]] = None,
            recipients: Optional[pulumi.Input[List[pulumi.Input[float]]]] = None,
            time_threshold: Optional[pulumi.Input[float]] = None,
            type: Optional[pulumi.Input[str]] = None,
            value_threshold: Optional[pulumi.Input[float]] = None,
            vhost_regex: Optional[pulumi.Input[str]] = None) -> 'Alarm':
        """
        Get an existing Alarm resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] enabled: Enable or disable the alarm to trigger.
        :param pulumi.Input[float] instance_id: The CloudAMQP instance ID.
        :param pulumi.Input[str] message_type: Message type `(total, unacked, ready)` used by queue alarm type.
        :param pulumi.Input[str] queue_regex: Regex for which queue to check.
        :param pulumi.Input[List[pulumi.Input[float]]] recipients: Identifier for recipient to be notified. Leave empty to notify all recipients.
        :param pulumi.Input[float] time_threshold: The time interval (in seconds) the `value_threshold` should be active before triggering an alarm.
        :param pulumi.Input[str] type: The alarm type, see valid options below.
        :param pulumi.Input[float] value_threshold: The value to trigger the alarm for.
        :param pulumi.Input[str] vhost_regex: Regex for which vhost to check
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

    @property
    @pulumi.getter
    def enabled(self) -> bool:
        """
        Enable or disable the alarm to trigger.
        """
        return pulumi.get(self, "enabled")

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> float:
        """
        The CloudAMQP instance ID.
        """
        return pulumi.get(self, "instance_id")

    @property
    @pulumi.getter(name="messageType")
    def message_type(self) -> Optional[str]:
        """
        Message type `(total, unacked, ready)` used by queue alarm type.
        """
        return pulumi.get(self, "message_type")

    @property
    @pulumi.getter(name="queueRegex")
    def queue_regex(self) -> Optional[str]:
        """
        Regex for which queue to check.
        """
        return pulumi.get(self, "queue_regex")

    @property
    @pulumi.getter
    def recipients(self) -> List[float]:
        """
        Identifier for recipient to be notified. Leave empty to notify all recipients.
        """
        return pulumi.get(self, "recipients")

    @property
    @pulumi.getter(name="timeThreshold")
    def time_threshold(self) -> Optional[float]:
        """
        The time interval (in seconds) the `value_threshold` should be active before triggering an alarm.
        """
        return pulumi.get(self, "time_threshold")

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        The alarm type, see valid options below.
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter(name="valueThreshold")
    def value_threshold(self) -> Optional[float]:
        """
        The value to trigger the alarm for.
        """
        return pulumi.get(self, "value_threshold")

    @property
    @pulumi.getter(name="vhostRegex")
    def vhost_regex(self) -> Optional[str]:
        """
        Regex for which vhost to check
        """
        return pulumi.get(self, "vhost_regex")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

