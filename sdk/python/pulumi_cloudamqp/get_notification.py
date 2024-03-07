# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = [
    'GetNotificationResult',
    'AwaitableGetNotificationResult',
    'get_notification',
    'get_notification_output',
]

@pulumi.output_type
class GetNotificationResult:
    """
    A collection of values returned by getNotification.
    """
    def __init__(__self__, id=None, instance_id=None, name=None, options=None, recipient_id=None, type=None, value=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if instance_id and not isinstance(instance_id, int):
            raise TypeError("Expected argument 'instance_id' to be a int")
        pulumi.set(__self__, "instance_id", instance_id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if options and not isinstance(options, dict):
            raise TypeError("Expected argument 'options' to be a dict")
        pulumi.set(__self__, "options", options)
        if recipient_id and not isinstance(recipient_id, int):
            raise TypeError("Expected argument 'recipient_id' to be a int")
        pulumi.set(__self__, "recipient_id", recipient_id)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)
        if value and not isinstance(value, str):
            raise TypeError("Expected argument 'value' to be a str")
        pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> int:
        return pulumi.get(self, "instance_id")

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def options(self) -> Optional[Mapping[str, str]]:
        return pulumi.get(self, "options")

    @property
    @pulumi.getter(name="recipientId")
    def recipient_id(self) -> Optional[int]:
        return pulumi.get(self, "recipient_id")

    @property
    @pulumi.getter
    def type(self) -> str:
        return pulumi.get(self, "type")

    @property
    @pulumi.getter
    def value(self) -> str:
        return pulumi.get(self, "value")


class AwaitableGetNotificationResult(GetNotificationResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetNotificationResult(
            id=self.id,
            instance_id=self.instance_id,
            name=self.name,
            options=self.options,
            recipient_id=self.recipient_id,
            type=self.type,
            value=self.value)


def get_notification(instance_id: Optional[int] = None,
                     name: Optional[str] = None,
                     options: Optional[Mapping[str, str]] = None,
                     recipient_id: Optional[int] = None,
                     opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetNotificationResult:
    """
    Use this data source to retrieve information about default or created recipients. The recipient will receive notifications assigned to an alarm that has triggered. To retrieve the recipient either use `recipient_id` or `name`.

    ## Example Usage

    <!--Start PulumiCodeChooser -->
    ```python
    import pulumi
    import pulumi_cloudamqp as cloudamqp

    default_recipient = cloudamqp.get_notification(instance_id=cloudamqp_instance["instance"]["id"],
        name="default")
    ```
    <!--End PulumiCodeChooser -->

    ## Attributes reference

    All attributes reference are computed

    * `id`    - The identifier for this resource.
    * `type`  - The type of the recipient.
    * `value` - The notification endpoint, where to send the notification.
    * `options`- Options argument (e.g. `rk` used for VictorOps routing key).

    ## Dependency

    This data source depends on CloudAMQP instance identifier, `cloudamqp_instance.instance.id`.


    :param int instance_id: The CloudAMQP instance identifier.
    :param str name: The name set for the recipient.
    :param int recipient_id: The recipient identifier.
    """
    __args__ = dict()
    __args__['instanceId'] = instance_id
    __args__['name'] = name
    __args__['options'] = options
    __args__['recipientId'] = recipient_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('cloudamqp:index/getNotification:getNotification', __args__, opts=opts, typ=GetNotificationResult).value

    return AwaitableGetNotificationResult(
        id=pulumi.get(__ret__, 'id'),
        instance_id=pulumi.get(__ret__, 'instance_id'),
        name=pulumi.get(__ret__, 'name'),
        options=pulumi.get(__ret__, 'options'),
        recipient_id=pulumi.get(__ret__, 'recipient_id'),
        type=pulumi.get(__ret__, 'type'),
        value=pulumi.get(__ret__, 'value'))


@_utilities.lift_output_func(get_notification)
def get_notification_output(instance_id: Optional[pulumi.Input[int]] = None,
                            name: Optional[pulumi.Input[Optional[str]]] = None,
                            options: Optional[pulumi.Input[Optional[Mapping[str, str]]]] = None,
                            recipient_id: Optional[pulumi.Input[Optional[int]]] = None,
                            opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetNotificationResult]:
    """
    Use this data source to retrieve information about default or created recipients. The recipient will receive notifications assigned to an alarm that has triggered. To retrieve the recipient either use `recipient_id` or `name`.

    ## Example Usage

    <!--Start PulumiCodeChooser -->
    ```python
    import pulumi
    import pulumi_cloudamqp as cloudamqp

    default_recipient = cloudamqp.get_notification(instance_id=cloudamqp_instance["instance"]["id"],
        name="default")
    ```
    <!--End PulumiCodeChooser -->

    ## Attributes reference

    All attributes reference are computed

    * `id`    - The identifier for this resource.
    * `type`  - The type of the recipient.
    * `value` - The notification endpoint, where to send the notification.
    * `options`- Options argument (e.g. `rk` used for VictorOps routing key).

    ## Dependency

    This data source depends on CloudAMQP instance identifier, `cloudamqp_instance.instance.id`.


    :param int instance_id: The CloudAMQP instance identifier.
    :param str name: The name set for the recipient.
    :param int recipient_id: The recipient identifier.
    """
    ...
