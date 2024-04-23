# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['AlarmArgs', 'Alarm']

@pulumi.input_type
class AlarmArgs:
    def __init__(__self__, *,
                 enabled: pulumi.Input[bool],
                 instance_id: pulumi.Input[int],
                 recipients: pulumi.Input[Sequence[pulumi.Input[int]]],
                 type: pulumi.Input[str],
                 message_type: Optional[pulumi.Input[str]] = None,
                 queue_regex: Optional[pulumi.Input[str]] = None,
                 reminder_interval: Optional[pulumi.Input[int]] = None,
                 time_threshold: Optional[pulumi.Input[int]] = None,
                 value_calculation: Optional[pulumi.Input[str]] = None,
                 value_threshold: Optional[pulumi.Input[int]] = None,
                 vhost_regex: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Alarm resource.
        :param pulumi.Input[bool] enabled: Enable or disable the alarm to trigger.
        :param pulumi.Input[int] instance_id: The CloudAMQP instance ID.
        :param pulumi.Input[Sequence[pulumi.Input[int]]] recipients: Identifier for recipient to be notified. Leave empty to notify all recipients.
        :param pulumi.Input[str] type: The alarm type, see valid options below.
        :param pulumi.Input[str] message_type: Message type `(total, unacked, ready)` used by queue alarm type.
               
               Specific argument for `disk` alarm
        :param pulumi.Input[str] queue_regex: Regex for which queue to check.
        :param pulumi.Input[int] reminder_interval: The reminder interval (in seconds) to resend the alarm if not resolved. Set to 0 for no reminders. The Default is 0.
        :param pulumi.Input[int] time_threshold: The time interval (in seconds) the `value_threshold` should be active before triggering an alarm.
        :param pulumi.Input[str] value_calculation: Disk value threshold calculation, `fixed, percentage` of disk space remaining.
               
               Based on alarm type, different arguments are flagged as required or optional.
        :param pulumi.Input[int] value_threshold: The value to trigger the alarm for.
        :param pulumi.Input[str] vhost_regex: Regex for which vhost to check
        """
        pulumi.set(__self__, "enabled", enabled)
        pulumi.set(__self__, "instance_id", instance_id)
        pulumi.set(__self__, "recipients", recipients)
        pulumi.set(__self__, "type", type)
        if message_type is not None:
            pulumi.set(__self__, "message_type", message_type)
        if queue_regex is not None:
            pulumi.set(__self__, "queue_regex", queue_regex)
        if reminder_interval is not None:
            pulumi.set(__self__, "reminder_interval", reminder_interval)
        if time_threshold is not None:
            pulumi.set(__self__, "time_threshold", time_threshold)
        if value_calculation is not None:
            pulumi.set(__self__, "value_calculation", value_calculation)
        if value_threshold is not None:
            pulumi.set(__self__, "value_threshold", value_threshold)
        if vhost_regex is not None:
            pulumi.set(__self__, "vhost_regex", vhost_regex)

    @property
    @pulumi.getter
    def enabled(self) -> pulumi.Input[bool]:
        """
        Enable or disable the alarm to trigger.
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
    def recipients(self) -> pulumi.Input[Sequence[pulumi.Input[int]]]:
        """
        Identifier for recipient to be notified. Leave empty to notify all recipients.
        """
        return pulumi.get(self, "recipients")

    @recipients.setter
    def recipients(self, value: pulumi.Input[Sequence[pulumi.Input[int]]]):
        pulumi.set(self, "recipients", value)

    @property
    @pulumi.getter
    def type(self) -> pulumi.Input[str]:
        """
        The alarm type, see valid options below.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: pulumi.Input[str]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter(name="messageType")
    def message_type(self) -> Optional[pulumi.Input[str]]:
        """
        Message type `(total, unacked, ready)` used by queue alarm type.

        Specific argument for `disk` alarm
        """
        return pulumi.get(self, "message_type")

    @message_type.setter
    def message_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "message_type", value)

    @property
    @pulumi.getter(name="queueRegex")
    def queue_regex(self) -> Optional[pulumi.Input[str]]:
        """
        Regex for which queue to check.
        """
        return pulumi.get(self, "queue_regex")

    @queue_regex.setter
    def queue_regex(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "queue_regex", value)

    @property
    @pulumi.getter(name="reminderInterval")
    def reminder_interval(self) -> Optional[pulumi.Input[int]]:
        """
        The reminder interval (in seconds) to resend the alarm if not resolved. Set to 0 for no reminders. The Default is 0.
        """
        return pulumi.get(self, "reminder_interval")

    @reminder_interval.setter
    def reminder_interval(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "reminder_interval", value)

    @property
    @pulumi.getter(name="timeThreshold")
    def time_threshold(self) -> Optional[pulumi.Input[int]]:
        """
        The time interval (in seconds) the `value_threshold` should be active before triggering an alarm.
        """
        return pulumi.get(self, "time_threshold")

    @time_threshold.setter
    def time_threshold(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "time_threshold", value)

    @property
    @pulumi.getter(name="valueCalculation")
    def value_calculation(self) -> Optional[pulumi.Input[str]]:
        """
        Disk value threshold calculation, `fixed, percentage` of disk space remaining.

        Based on alarm type, different arguments are flagged as required or optional.
        """
        return pulumi.get(self, "value_calculation")

    @value_calculation.setter
    def value_calculation(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "value_calculation", value)

    @property
    @pulumi.getter(name="valueThreshold")
    def value_threshold(self) -> Optional[pulumi.Input[int]]:
        """
        The value to trigger the alarm for.
        """
        return pulumi.get(self, "value_threshold")

    @value_threshold.setter
    def value_threshold(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "value_threshold", value)

    @property
    @pulumi.getter(name="vhostRegex")
    def vhost_regex(self) -> Optional[pulumi.Input[str]]:
        """
        Regex for which vhost to check
        """
        return pulumi.get(self, "vhost_regex")

    @vhost_regex.setter
    def vhost_regex(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vhost_regex", value)


@pulumi.input_type
class _AlarmState:
    def __init__(__self__, *,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 instance_id: Optional[pulumi.Input[int]] = None,
                 message_type: Optional[pulumi.Input[str]] = None,
                 queue_regex: Optional[pulumi.Input[str]] = None,
                 recipients: Optional[pulumi.Input[Sequence[pulumi.Input[int]]]] = None,
                 reminder_interval: Optional[pulumi.Input[int]] = None,
                 time_threshold: Optional[pulumi.Input[int]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 value_calculation: Optional[pulumi.Input[str]] = None,
                 value_threshold: Optional[pulumi.Input[int]] = None,
                 vhost_regex: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Alarm resources.
        :param pulumi.Input[bool] enabled: Enable or disable the alarm to trigger.
        :param pulumi.Input[int] instance_id: The CloudAMQP instance ID.
        :param pulumi.Input[str] message_type: Message type `(total, unacked, ready)` used by queue alarm type.
               
               Specific argument for `disk` alarm
        :param pulumi.Input[str] queue_regex: Regex for which queue to check.
        :param pulumi.Input[Sequence[pulumi.Input[int]]] recipients: Identifier for recipient to be notified. Leave empty to notify all recipients.
        :param pulumi.Input[int] reminder_interval: The reminder interval (in seconds) to resend the alarm if not resolved. Set to 0 for no reminders. The Default is 0.
        :param pulumi.Input[int] time_threshold: The time interval (in seconds) the `value_threshold` should be active before triggering an alarm.
        :param pulumi.Input[str] type: The alarm type, see valid options below.
        :param pulumi.Input[str] value_calculation: Disk value threshold calculation, `fixed, percentage` of disk space remaining.
               
               Based on alarm type, different arguments are flagged as required or optional.
        :param pulumi.Input[int] value_threshold: The value to trigger the alarm for.
        :param pulumi.Input[str] vhost_regex: Regex for which vhost to check
        """
        if enabled is not None:
            pulumi.set(__self__, "enabled", enabled)
        if instance_id is not None:
            pulumi.set(__self__, "instance_id", instance_id)
        if message_type is not None:
            pulumi.set(__self__, "message_type", message_type)
        if queue_regex is not None:
            pulumi.set(__self__, "queue_regex", queue_regex)
        if recipients is not None:
            pulumi.set(__self__, "recipients", recipients)
        if reminder_interval is not None:
            pulumi.set(__self__, "reminder_interval", reminder_interval)
        if time_threshold is not None:
            pulumi.set(__self__, "time_threshold", time_threshold)
        if type is not None:
            pulumi.set(__self__, "type", type)
        if value_calculation is not None:
            pulumi.set(__self__, "value_calculation", value_calculation)
        if value_threshold is not None:
            pulumi.set(__self__, "value_threshold", value_threshold)
        if vhost_regex is not None:
            pulumi.set(__self__, "vhost_regex", vhost_regex)

    @property
    @pulumi.getter
    def enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        Enable or disable the alarm to trigger.
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
    @pulumi.getter(name="messageType")
    def message_type(self) -> Optional[pulumi.Input[str]]:
        """
        Message type `(total, unacked, ready)` used by queue alarm type.

        Specific argument for `disk` alarm
        """
        return pulumi.get(self, "message_type")

    @message_type.setter
    def message_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "message_type", value)

    @property
    @pulumi.getter(name="queueRegex")
    def queue_regex(self) -> Optional[pulumi.Input[str]]:
        """
        Regex for which queue to check.
        """
        return pulumi.get(self, "queue_regex")

    @queue_regex.setter
    def queue_regex(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "queue_regex", value)

    @property
    @pulumi.getter
    def recipients(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[int]]]]:
        """
        Identifier for recipient to be notified. Leave empty to notify all recipients.
        """
        return pulumi.get(self, "recipients")

    @recipients.setter
    def recipients(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[int]]]]):
        pulumi.set(self, "recipients", value)

    @property
    @pulumi.getter(name="reminderInterval")
    def reminder_interval(self) -> Optional[pulumi.Input[int]]:
        """
        The reminder interval (in seconds) to resend the alarm if not resolved. Set to 0 for no reminders. The Default is 0.
        """
        return pulumi.get(self, "reminder_interval")

    @reminder_interval.setter
    def reminder_interval(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "reminder_interval", value)

    @property
    @pulumi.getter(name="timeThreshold")
    def time_threshold(self) -> Optional[pulumi.Input[int]]:
        """
        The time interval (in seconds) the `value_threshold` should be active before triggering an alarm.
        """
        return pulumi.get(self, "time_threshold")

    @time_threshold.setter
    def time_threshold(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "time_threshold", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[str]]:
        """
        The alarm type, see valid options below.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter(name="valueCalculation")
    def value_calculation(self) -> Optional[pulumi.Input[str]]:
        """
        Disk value threshold calculation, `fixed, percentage` of disk space remaining.

        Based on alarm type, different arguments are flagged as required or optional.
        """
        return pulumi.get(self, "value_calculation")

    @value_calculation.setter
    def value_calculation(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "value_calculation", value)

    @property
    @pulumi.getter(name="valueThreshold")
    def value_threshold(self) -> Optional[pulumi.Input[int]]:
        """
        The value to trigger the alarm for.
        """
        return pulumi.get(self, "value_threshold")

    @value_threshold.setter
    def value_threshold(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "value_threshold", value)

    @property
    @pulumi.getter(name="vhostRegex")
    def vhost_regex(self) -> Optional[pulumi.Input[str]]:
        """
        Regex for which vhost to check
        """
        return pulumi.get(self, "vhost_regex")

    @vhost_regex.setter
    def vhost_regex(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vhost_regex", value)


class Alarm(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 instance_id: Optional[pulumi.Input[int]] = None,
                 message_type: Optional[pulumi.Input[str]] = None,
                 queue_regex: Optional[pulumi.Input[str]] = None,
                 recipients: Optional[pulumi.Input[Sequence[pulumi.Input[int]]]] = None,
                 reminder_interval: Optional[pulumi.Input[int]] = None,
                 time_threshold: Optional[pulumi.Input[int]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 value_calculation: Optional[pulumi.Input[str]] = None,
                 value_threshold: Optional[pulumi.Input[int]] = None,
                 vhost_regex: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        This resource allows you to create and manage alarms to trigger based on a set of conditions. Once triggerd a notification will be sent to the assigned recipients. When creating a new instance, there will also be a set of default alarms (cpu, memory and disk) created. All default alarms uses the default recipient for notifications.

        By setting `no_default_alarms` to *true* in `Instance`. This will create the instance without default alarms and avoid the need to import them to get full control.

        Available for all subscription plans, but `lemur`and `tiger`are limited to fewer alarm types. The limited types supported can be seen in the table below in Alarm Type Reference.

        ## Example Usage

        <details>
          <summary>
            <b>
              <i>Basic example of CPU and memory alarm</i>
            </b>
          </summary>

        ```python
        import pulumi
        import pulumi_cloudamqp as cloudamqp

        # New recipient
        recipient01 = cloudamqp.Notification("recipient_01",
            instance_id=instance["id"],
            type="email",
            value="alarm@example.com",
            name="alarm")
        # New cpu alarm
        cpu_alarm = cloudamqp.Alarm("cpu_alarm",
            instance_id=instance["id"],
            type="cpu",
            enabled=True,
            reminder_interval=600,
            value_threshold=95,
            time_threshold=600,
            recipients=[recipient01.id])
        # New memory alarm
        memory_alarm = cloudamqp.Alarm("memory_alarm",
            instance_id=instance["id"],
            type="memory",
            enabled=True,
            reminder_interval=600,
            value_threshold=95,
            time_threshold=600,
            recipients=[recipient01.id])
        ```

        </details>

        <details>
          <summary>
            <b>
              <i>Manage notice alarm, available from v1.29.5</i>
            </b>
          </summary>

        Only one notice alarm can exists and cannot be created, instead the alarm resource will be updated.

        ```python
        import pulumi
        import pulumi_cloudamqp as cloudamqp

        # New recipient
        recipient01 = cloudamqp.Notification("recipient_01",
            instance_id=instance["id"],
            type="email",
            value="alarm@example.com",
            name="alarm")
        # Update existing notice alarm
        notice = cloudamqp.Alarm("notice",
            instance_id=instance["id"],
            type="notice",
            enabled=True,
            recipients=[recipient01.id])
        ```

        </details>

        ## Alarm Type reference

        Supported alarm types: `cpu, memory, disk, queue, connection, flow, consumer, netsplit, server_unreachable, notice`

        Required arguments for all alarms: `instance_id, type, enabled`<br>
        Optional argument for all alarms: `tags, queue_regex, vhost_regex`

        | Name | Type | Shared | Dedicated | Required arguments |
        | ---- | ---- | ---- | ---- | ---- |
        | CPU | cpu | - | &#10004; | time_threshold, value_threshold |
        | Memory | memory | - | &#10004;  | time_threshold, value_threshold |
        | Disk space | disk | - | &#10004;  | time_threshold, value_threshold |
        | Queue | queue | &#10004;  | &#10004; | time_threshold, value_threshold, queue_regex, vhost_regex, message_type |
        | Connection | connection | &#10004; | &#10004; | time_threshold, value_threshold |
        | Connection flow | flow | &#10004; | &#10004; | time_threshold, value_threshold |
        | Consumer | consumer | &#10004; | &#10004; | time_threshold, value_threshold, queue, vhost |
        | Netsplit | netsplit | - | &#10004; | time_threshold |
        | Server unreachable | server_unreachable  | - | &#10004;  | time_threshold |
        | Notice | notice | &#10004; | &#10004; | |

        > Notice alarm is manadatory! Only one can exists and cannot be deleted. Setting `no_default_alarm` to true, will still create this alarm. See updated changes to notice alarm below.

        ## Dependency

        This resource depends on CloudAMQP instance identifier, `cloudamqp_instance.instance.id`.

        ## Notice alarm

        There is a limitation for notice alarm in the API backend. This alarm is mandatory, multiple
        alarms cannot exists or be deleted.

        From provider version v1.29.5
        it's possible to manage the notice alarm and no longer needs to be imported. Just create the
        alarm resource as usually and it will be updated with given recipients. If the alarm is deleted
        it will only be removed from the state file, but will still be enabled in the backend.

        ## Import

        `cloudamqp_alarm` can be imported using CloudAMQP internal identifier of the alarm together (CSV separated) with the instance identifier. To retrieve the alarm identifier, use [CloudAMQP API](https://docs.cloudamqp.com/cloudamqp_api.html#list-alarms)

        ```sh
        $ pulumi import cloudamqp:index/alarm:Alarm alarm <id>,<instance_id>`
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] enabled: Enable or disable the alarm to trigger.
        :param pulumi.Input[int] instance_id: The CloudAMQP instance ID.
        :param pulumi.Input[str] message_type: Message type `(total, unacked, ready)` used by queue alarm type.
               
               Specific argument for `disk` alarm
        :param pulumi.Input[str] queue_regex: Regex for which queue to check.
        :param pulumi.Input[Sequence[pulumi.Input[int]]] recipients: Identifier for recipient to be notified. Leave empty to notify all recipients.
        :param pulumi.Input[int] reminder_interval: The reminder interval (in seconds) to resend the alarm if not resolved. Set to 0 for no reminders. The Default is 0.
        :param pulumi.Input[int] time_threshold: The time interval (in seconds) the `value_threshold` should be active before triggering an alarm.
        :param pulumi.Input[str] type: The alarm type, see valid options below.
        :param pulumi.Input[str] value_calculation: Disk value threshold calculation, `fixed, percentage` of disk space remaining.
               
               Based on alarm type, different arguments are flagged as required or optional.
        :param pulumi.Input[int] value_threshold: The value to trigger the alarm for.
        :param pulumi.Input[str] vhost_regex: Regex for which vhost to check
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: AlarmArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        This resource allows you to create and manage alarms to trigger based on a set of conditions. Once triggerd a notification will be sent to the assigned recipients. When creating a new instance, there will also be a set of default alarms (cpu, memory and disk) created. All default alarms uses the default recipient for notifications.

        By setting `no_default_alarms` to *true* in `Instance`. This will create the instance without default alarms and avoid the need to import them to get full control.

        Available for all subscription plans, but `lemur`and `tiger`are limited to fewer alarm types. The limited types supported can be seen in the table below in Alarm Type Reference.

        ## Example Usage

        <details>
          <summary>
            <b>
              <i>Basic example of CPU and memory alarm</i>
            </b>
          </summary>

        ```python
        import pulumi
        import pulumi_cloudamqp as cloudamqp

        # New recipient
        recipient01 = cloudamqp.Notification("recipient_01",
            instance_id=instance["id"],
            type="email",
            value="alarm@example.com",
            name="alarm")
        # New cpu alarm
        cpu_alarm = cloudamqp.Alarm("cpu_alarm",
            instance_id=instance["id"],
            type="cpu",
            enabled=True,
            reminder_interval=600,
            value_threshold=95,
            time_threshold=600,
            recipients=[recipient01.id])
        # New memory alarm
        memory_alarm = cloudamqp.Alarm("memory_alarm",
            instance_id=instance["id"],
            type="memory",
            enabled=True,
            reminder_interval=600,
            value_threshold=95,
            time_threshold=600,
            recipients=[recipient01.id])
        ```

        </details>

        <details>
          <summary>
            <b>
              <i>Manage notice alarm, available from v1.29.5</i>
            </b>
          </summary>

        Only one notice alarm can exists and cannot be created, instead the alarm resource will be updated.

        ```python
        import pulumi
        import pulumi_cloudamqp as cloudamqp

        # New recipient
        recipient01 = cloudamqp.Notification("recipient_01",
            instance_id=instance["id"],
            type="email",
            value="alarm@example.com",
            name="alarm")
        # Update existing notice alarm
        notice = cloudamqp.Alarm("notice",
            instance_id=instance["id"],
            type="notice",
            enabled=True,
            recipients=[recipient01.id])
        ```

        </details>

        ## Alarm Type reference

        Supported alarm types: `cpu, memory, disk, queue, connection, flow, consumer, netsplit, server_unreachable, notice`

        Required arguments for all alarms: `instance_id, type, enabled`<br>
        Optional argument for all alarms: `tags, queue_regex, vhost_regex`

        | Name | Type | Shared | Dedicated | Required arguments |
        | ---- | ---- | ---- | ---- | ---- |
        | CPU | cpu | - | &#10004; | time_threshold, value_threshold |
        | Memory | memory | - | &#10004;  | time_threshold, value_threshold |
        | Disk space | disk | - | &#10004;  | time_threshold, value_threshold |
        | Queue | queue | &#10004;  | &#10004; | time_threshold, value_threshold, queue_regex, vhost_regex, message_type |
        | Connection | connection | &#10004; | &#10004; | time_threshold, value_threshold |
        | Connection flow | flow | &#10004; | &#10004; | time_threshold, value_threshold |
        | Consumer | consumer | &#10004; | &#10004; | time_threshold, value_threshold, queue, vhost |
        | Netsplit | netsplit | - | &#10004; | time_threshold |
        | Server unreachable | server_unreachable  | - | &#10004;  | time_threshold |
        | Notice | notice | &#10004; | &#10004; | |

        > Notice alarm is manadatory! Only one can exists and cannot be deleted. Setting `no_default_alarm` to true, will still create this alarm. See updated changes to notice alarm below.

        ## Dependency

        This resource depends on CloudAMQP instance identifier, `cloudamqp_instance.instance.id`.

        ## Notice alarm

        There is a limitation for notice alarm in the API backend. This alarm is mandatory, multiple
        alarms cannot exists or be deleted.

        From provider version v1.29.5
        it's possible to manage the notice alarm and no longer needs to be imported. Just create the
        alarm resource as usually and it will be updated with given recipients. If the alarm is deleted
        it will only be removed from the state file, but will still be enabled in the backend.

        ## Import

        `cloudamqp_alarm` can be imported using CloudAMQP internal identifier of the alarm together (CSV separated) with the instance identifier. To retrieve the alarm identifier, use [CloudAMQP API](https://docs.cloudamqp.com/cloudamqp_api.html#list-alarms)

        ```sh
        $ pulumi import cloudamqp:index/alarm:Alarm alarm <id>,<instance_id>`
        ```

        :param str resource_name: The name of the resource.
        :param AlarmArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(AlarmArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 instance_id: Optional[pulumi.Input[int]] = None,
                 message_type: Optional[pulumi.Input[str]] = None,
                 queue_regex: Optional[pulumi.Input[str]] = None,
                 recipients: Optional[pulumi.Input[Sequence[pulumi.Input[int]]]] = None,
                 reminder_interval: Optional[pulumi.Input[int]] = None,
                 time_threshold: Optional[pulumi.Input[int]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 value_calculation: Optional[pulumi.Input[str]] = None,
                 value_threshold: Optional[pulumi.Input[int]] = None,
                 vhost_regex: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = AlarmArgs.__new__(AlarmArgs)

            if enabled is None and not opts.urn:
                raise TypeError("Missing required property 'enabled'")
            __props__.__dict__["enabled"] = enabled
            if instance_id is None and not opts.urn:
                raise TypeError("Missing required property 'instance_id'")
            __props__.__dict__["instance_id"] = instance_id
            __props__.__dict__["message_type"] = message_type
            __props__.__dict__["queue_regex"] = queue_regex
            if recipients is None and not opts.urn:
                raise TypeError("Missing required property 'recipients'")
            __props__.__dict__["recipients"] = recipients
            __props__.__dict__["reminder_interval"] = reminder_interval
            __props__.__dict__["time_threshold"] = time_threshold
            if type is None and not opts.urn:
                raise TypeError("Missing required property 'type'")
            __props__.__dict__["type"] = type
            __props__.__dict__["value_calculation"] = value_calculation
            __props__.__dict__["value_threshold"] = value_threshold
            __props__.__dict__["vhost_regex"] = vhost_regex
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
            instance_id: Optional[pulumi.Input[int]] = None,
            message_type: Optional[pulumi.Input[str]] = None,
            queue_regex: Optional[pulumi.Input[str]] = None,
            recipients: Optional[pulumi.Input[Sequence[pulumi.Input[int]]]] = None,
            reminder_interval: Optional[pulumi.Input[int]] = None,
            time_threshold: Optional[pulumi.Input[int]] = None,
            type: Optional[pulumi.Input[str]] = None,
            value_calculation: Optional[pulumi.Input[str]] = None,
            value_threshold: Optional[pulumi.Input[int]] = None,
            vhost_regex: Optional[pulumi.Input[str]] = None) -> 'Alarm':
        """
        Get an existing Alarm resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] enabled: Enable or disable the alarm to trigger.
        :param pulumi.Input[int] instance_id: The CloudAMQP instance ID.
        :param pulumi.Input[str] message_type: Message type `(total, unacked, ready)` used by queue alarm type.
               
               Specific argument for `disk` alarm
        :param pulumi.Input[str] queue_regex: Regex for which queue to check.
        :param pulumi.Input[Sequence[pulumi.Input[int]]] recipients: Identifier for recipient to be notified. Leave empty to notify all recipients.
        :param pulumi.Input[int] reminder_interval: The reminder interval (in seconds) to resend the alarm if not resolved. Set to 0 for no reminders. The Default is 0.
        :param pulumi.Input[int] time_threshold: The time interval (in seconds) the `value_threshold` should be active before triggering an alarm.
        :param pulumi.Input[str] type: The alarm type, see valid options below.
        :param pulumi.Input[str] value_calculation: Disk value threshold calculation, `fixed, percentage` of disk space remaining.
               
               Based on alarm type, different arguments are flagged as required or optional.
        :param pulumi.Input[int] value_threshold: The value to trigger the alarm for.
        :param pulumi.Input[str] vhost_regex: Regex for which vhost to check
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _AlarmState.__new__(_AlarmState)

        __props__.__dict__["enabled"] = enabled
        __props__.__dict__["instance_id"] = instance_id
        __props__.__dict__["message_type"] = message_type
        __props__.__dict__["queue_regex"] = queue_regex
        __props__.__dict__["recipients"] = recipients
        __props__.__dict__["reminder_interval"] = reminder_interval
        __props__.__dict__["time_threshold"] = time_threshold
        __props__.__dict__["type"] = type
        __props__.__dict__["value_calculation"] = value_calculation
        __props__.__dict__["value_threshold"] = value_threshold
        __props__.__dict__["vhost_regex"] = vhost_regex
        return Alarm(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def enabled(self) -> pulumi.Output[bool]:
        """
        Enable or disable the alarm to trigger.
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
    @pulumi.getter(name="messageType")
    def message_type(self) -> pulumi.Output[Optional[str]]:
        """
        Message type `(total, unacked, ready)` used by queue alarm type.

        Specific argument for `disk` alarm
        """
        return pulumi.get(self, "message_type")

    @property
    @pulumi.getter(name="queueRegex")
    def queue_regex(self) -> pulumi.Output[Optional[str]]:
        """
        Regex for which queue to check.
        """
        return pulumi.get(self, "queue_regex")

    @property
    @pulumi.getter
    def recipients(self) -> pulumi.Output[Sequence[int]]:
        """
        Identifier for recipient to be notified. Leave empty to notify all recipients.
        """
        return pulumi.get(self, "recipients")

    @property
    @pulumi.getter(name="reminderInterval")
    def reminder_interval(self) -> pulumi.Output[Optional[int]]:
        """
        The reminder interval (in seconds) to resend the alarm if not resolved. Set to 0 for no reminders. The Default is 0.
        """
        return pulumi.get(self, "reminder_interval")

    @property
    @pulumi.getter(name="timeThreshold")
    def time_threshold(self) -> pulumi.Output[Optional[int]]:
        """
        The time interval (in seconds) the `value_threshold` should be active before triggering an alarm.
        """
        return pulumi.get(self, "time_threshold")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[str]:
        """
        The alarm type, see valid options below.
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter(name="valueCalculation")
    def value_calculation(self) -> pulumi.Output[Optional[str]]:
        """
        Disk value threshold calculation, `fixed, percentage` of disk space remaining.

        Based on alarm type, different arguments are flagged as required or optional.
        """
        return pulumi.get(self, "value_calculation")

    @property
    @pulumi.getter(name="valueThreshold")
    def value_threshold(self) -> pulumi.Output[Optional[int]]:
        """
        The value to trigger the alarm for.
        """
        return pulumi.get(self, "value_threshold")

    @property
    @pulumi.getter(name="vhostRegex")
    def vhost_regex(self) -> pulumi.Output[Optional[str]]:
        """
        Regex for which vhost to check
        """
        return pulumi.get(self, "vhost_regex")

