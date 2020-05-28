# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from . import utilities, tables

class GetNotificationResult:
    """
    A collection of values returned by getNotification.
    """
    def __init__(__self__, id=None, instance_id=None, name=None, recipient_id=None, type=None, value=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        __self__.id = id
        """
        The provider-assigned unique ID for this managed resource.
        """
        if instance_id and not isinstance(instance_id, float):
            raise TypeError("Expected argument 'instance_id' to be a float")
        __self__.instance_id = instance_id
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        __self__.name = name
        if recipient_id and not isinstance(recipient_id, float):
            raise TypeError("Expected argument 'recipient_id' to be a float")
        __self__.recipient_id = recipient_id
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        __self__.type = type
        if value and not isinstance(value, str):
            raise TypeError("Expected argument 'value' to be a str")
        __self__.value = value
class AwaitableGetNotificationResult(GetNotificationResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetNotificationResult(
            id=self.id,
            instance_id=self.instance_id,
            name=self.name,
            recipient_id=self.recipient_id,
            type=self.type,
            value=self.value)

def get_notification(instance_id=None,name=None,recipient_id=None,opts=None):
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()


    __args__['instanceId'] = instance_id
    __args__['name'] = name
    __args__['recipientId'] = recipient_id
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = utilities.get_version()
    __ret__ = pulumi.runtime.invoke('cloudamqp:index/getNotification:getNotification', __args__, opts=opts).value

    return AwaitableGetNotificationResult(
        id=__ret__.get('id'),
        instance_id=__ret__.get('instanceId'),
        name=__ret__.get('name'),
        recipient_id=__ret__.get('recipientId'),
        type=__ret__.get('type'),
        value=__ret__.get('value'))