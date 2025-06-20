# coding=utf-8
# *** WARNING: this file was generated by pulumi-language-python. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import builtins
import copy
import warnings
import sys
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
if sys.version_info >= (3, 11):
    from typing import NotRequired, TypedDict, TypeAlias
else:
    from typing_extensions import NotRequired, TypedDict, TypeAlias
from . import _utilities

__all__ = [
    'GetAlarmResult',
    'AwaitableGetAlarmResult',
    'get_alarm',
    'get_alarm_output',
]

@pulumi.output_type
class GetAlarmResult:
    """
    A collection of values returned by getAlarm.
    """
    def __init__(__self__, alarm_id=None, enabled=None, id=None, instance_id=None, message_type=None, queue_regex=None, recipients=None, reminder_interval=None, time_threshold=None, type=None, value_calculation=None, value_threshold=None, vhost_regex=None):
        if alarm_id and not isinstance(alarm_id, int):
            raise TypeError("Expected argument 'alarm_id' to be a int")
        pulumi.set(__self__, "alarm_id", alarm_id)
        if enabled and not isinstance(enabled, bool):
            raise TypeError("Expected argument 'enabled' to be a bool")
        pulumi.set(__self__, "enabled", enabled)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if instance_id and not isinstance(instance_id, int):
            raise TypeError("Expected argument 'instance_id' to be a int")
        pulumi.set(__self__, "instance_id", instance_id)
        if message_type and not isinstance(message_type, str):
            raise TypeError("Expected argument 'message_type' to be a str")
        pulumi.set(__self__, "message_type", message_type)
        if queue_regex and not isinstance(queue_regex, str):
            raise TypeError("Expected argument 'queue_regex' to be a str")
        pulumi.set(__self__, "queue_regex", queue_regex)
        if recipients and not isinstance(recipients, list):
            raise TypeError("Expected argument 'recipients' to be a list")
        pulumi.set(__self__, "recipients", recipients)
        if reminder_interval and not isinstance(reminder_interval, int):
            raise TypeError("Expected argument 'reminder_interval' to be a int")
        pulumi.set(__self__, "reminder_interval", reminder_interval)
        if time_threshold and not isinstance(time_threshold, int):
            raise TypeError("Expected argument 'time_threshold' to be a int")
        pulumi.set(__self__, "time_threshold", time_threshold)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)
        if value_calculation and not isinstance(value_calculation, str):
            raise TypeError("Expected argument 'value_calculation' to be a str")
        pulumi.set(__self__, "value_calculation", value_calculation)
        if value_threshold and not isinstance(value_threshold, int):
            raise TypeError("Expected argument 'value_threshold' to be a int")
        pulumi.set(__self__, "value_threshold", value_threshold)
        if vhost_regex and not isinstance(vhost_regex, str):
            raise TypeError("Expected argument 'vhost_regex' to be a str")
        pulumi.set(__self__, "vhost_regex", vhost_regex)

    @property
    @pulumi.getter(name="alarmId")
    def alarm_id(self) -> Optional[builtins.int]:
        return pulumi.get(self, "alarm_id")

    @property
    @pulumi.getter
    def enabled(self) -> builtins.bool:
        """
        Enable/disable status of the alarm.
        """
        return pulumi.get(self, "enabled")

    @property
    @pulumi.getter
    def id(self) -> builtins.str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> builtins.int:
        return pulumi.get(self, "instance_id")

    @property
    @pulumi.getter(name="messageType")
    def message_type(self) -> builtins.str:
        """
        Message type `(total, unacked, ready)` used by queue alarm type.
        """
        return pulumi.get(self, "message_type")

    @property
    @pulumi.getter(name="queueRegex")
    def queue_regex(self) -> builtins.str:
        """
        Regular expression for which queue to check.
        """
        return pulumi.get(self, "queue_regex")

    @property
    @pulumi.getter
    def recipients(self) -> Sequence[builtins.int]:
        """
        Identifier for recipient to be notified.
        """
        return pulumi.get(self, "recipients")

    @property
    @pulumi.getter(name="reminderInterval")
    def reminder_interval(self) -> builtins.int:
        """
        The reminder interval (in seconds) to resend the alarm if not resolved.
        Set to 0 for no reminders.
        """
        return pulumi.get(self, "reminder_interval")

    @property
    @pulumi.getter(name="timeThreshold")
    def time_threshold(self) -> builtins.int:
        """
        The time interval (in seconds) the `value_threshold` should be active
        before trigger an alarm.
        """
        return pulumi.get(self, "time_threshold")

    @property
    @pulumi.getter
    def type(self) -> Optional[builtins.str]:
        return pulumi.get(self, "type")

    @property
    @pulumi.getter(name="valueCalculation")
    def value_calculation(self) -> Optional[builtins.str]:
        """
        Disk value threshold calculation, `(fixed, percentage)` of disk space
        remaining.
        """
        return pulumi.get(self, "value_calculation")

    @property
    @pulumi.getter(name="valueThreshold")
    def value_threshold(self) -> builtins.int:
        """
        The value threshold that triggers the alarm.
        """
        return pulumi.get(self, "value_threshold")

    @property
    @pulumi.getter(name="vhostRegex")
    def vhost_regex(self) -> builtins.str:
        """
        Regular expression for which vhost to check
        """
        return pulumi.get(self, "vhost_regex")


class AwaitableGetAlarmResult(GetAlarmResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetAlarmResult(
            alarm_id=self.alarm_id,
            enabled=self.enabled,
            id=self.id,
            instance_id=self.instance_id,
            message_type=self.message_type,
            queue_regex=self.queue_regex,
            recipients=self.recipients,
            reminder_interval=self.reminder_interval,
            time_threshold=self.time_threshold,
            type=self.type,
            value_calculation=self.value_calculation,
            value_threshold=self.value_threshold,
            vhost_regex=self.vhost_regex)


def get_alarm(alarm_id: Optional[builtins.int] = None,
              instance_id: Optional[builtins.int] = None,
              type: Optional[builtins.str] = None,
              value_calculation: Optional[builtins.str] = None,
              opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetAlarmResult:
    """
    Use this data source to retrieve information about default or created alarms. Either use `alarm_id`
    or `type` to retrieve the alarm.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_cloudamqp as cloudamqp

    default_cpu_alarm = cloudamqp.get_alarm(instance_id=instance["id"],
        type="cpu")
    ```

    ## Dependency

    This data source depends on CloudAMQP instance identifier, `cloudamqp_instance.instance.id`.

    ## Alarm Types

    `cpu, memory, disk, queue, connection, flow, consumer, netsplit, server_unreachable, notice`


    :param builtins.int alarm_id: The alarm identifier. Either use this or `type` to give
           `Alarm` necessary information to retrieve the alarm.
    :param builtins.int instance_id: The CloudAMQP instance identifier.
    :param builtins.str type: The alarm type. Either use this or `alarm_id` to give `Alarm`
           necessary information when retrieve the alarm. Supported
           alarm types.
    :param builtins.str value_calculation: Disk value threshold calculation, `(fixed, percentage)` of disk space
           remaining.
    """
    __args__ = dict()
    __args__['alarmId'] = alarm_id
    __args__['instanceId'] = instance_id
    __args__['type'] = type
    __args__['valueCalculation'] = value_calculation
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('cloudamqp:index/getAlarm:getAlarm', __args__, opts=opts, typ=GetAlarmResult).value

    return AwaitableGetAlarmResult(
        alarm_id=pulumi.get(__ret__, 'alarm_id'),
        enabled=pulumi.get(__ret__, 'enabled'),
        id=pulumi.get(__ret__, 'id'),
        instance_id=pulumi.get(__ret__, 'instance_id'),
        message_type=pulumi.get(__ret__, 'message_type'),
        queue_regex=pulumi.get(__ret__, 'queue_regex'),
        recipients=pulumi.get(__ret__, 'recipients'),
        reminder_interval=pulumi.get(__ret__, 'reminder_interval'),
        time_threshold=pulumi.get(__ret__, 'time_threshold'),
        type=pulumi.get(__ret__, 'type'),
        value_calculation=pulumi.get(__ret__, 'value_calculation'),
        value_threshold=pulumi.get(__ret__, 'value_threshold'),
        vhost_regex=pulumi.get(__ret__, 'vhost_regex'))
def get_alarm_output(alarm_id: Optional[pulumi.Input[Optional[builtins.int]]] = None,
                     instance_id: Optional[pulumi.Input[builtins.int]] = None,
                     type: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                     value_calculation: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                     opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetAlarmResult]:
    """
    Use this data source to retrieve information about default or created alarms. Either use `alarm_id`
    or `type` to retrieve the alarm.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_cloudamqp as cloudamqp

    default_cpu_alarm = cloudamqp.get_alarm(instance_id=instance["id"],
        type="cpu")
    ```

    ## Dependency

    This data source depends on CloudAMQP instance identifier, `cloudamqp_instance.instance.id`.

    ## Alarm Types

    `cpu, memory, disk, queue, connection, flow, consumer, netsplit, server_unreachable, notice`


    :param builtins.int alarm_id: The alarm identifier. Either use this or `type` to give
           `Alarm` necessary information to retrieve the alarm.
    :param builtins.int instance_id: The CloudAMQP instance identifier.
    :param builtins.str type: The alarm type. Either use this or `alarm_id` to give `Alarm`
           necessary information when retrieve the alarm. Supported
           alarm types.
    :param builtins.str value_calculation: Disk value threshold calculation, `(fixed, percentage)` of disk space
           remaining.
    """
    __args__ = dict()
    __args__['alarmId'] = alarm_id
    __args__['instanceId'] = instance_id
    __args__['type'] = type
    __args__['valueCalculation'] = value_calculation
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('cloudamqp:index/getAlarm:getAlarm', __args__, opts=opts, typ=GetAlarmResult)
    return __ret__.apply(lambda __response__: GetAlarmResult(
        alarm_id=pulumi.get(__response__, 'alarm_id'),
        enabled=pulumi.get(__response__, 'enabled'),
        id=pulumi.get(__response__, 'id'),
        instance_id=pulumi.get(__response__, 'instance_id'),
        message_type=pulumi.get(__response__, 'message_type'),
        queue_regex=pulumi.get(__response__, 'queue_regex'),
        recipients=pulumi.get(__response__, 'recipients'),
        reminder_interval=pulumi.get(__response__, 'reminder_interval'),
        time_threshold=pulumi.get(__response__, 'time_threshold'),
        type=pulumi.get(__response__, 'type'),
        value_calculation=pulumi.get(__response__, 'value_calculation'),
        value_threshold=pulumi.get(__response__, 'value_threshold'),
        vhost_regex=pulumi.get(__response__, 'vhost_regex')))
