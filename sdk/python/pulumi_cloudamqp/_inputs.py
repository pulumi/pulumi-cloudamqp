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
    'ExtraDiskSizeNodeArgs',
    'InstanceCopySettingArgs',
    'SecurityFirewallRuleArgs',
]

@pulumi.input_type
class ExtraDiskSizeNodeArgs:
    def __init__(__self__, *,
                 additional_disk_size: Optional[pulumi.Input[int]] = None,
                 disk_size: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[int] additional_disk_size: Additional added disk size
        :param pulumi.Input[int] disk_size: Subscription plan disk size
        """
        if additional_disk_size is not None:
            pulumi.set(__self__, "additional_disk_size", additional_disk_size)
        if disk_size is not None:
            pulumi.set(__self__, "disk_size", disk_size)
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter(name="additionalDiskSize")
    def additional_disk_size(self) -> Optional[pulumi.Input[int]]:
        """
        Additional added disk size
        """
        return pulumi.get(self, "additional_disk_size")

    @additional_disk_size.setter
    def additional_disk_size(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "additional_disk_size", value)

    @property
    @pulumi.getter(name="diskSize")
    def disk_size(self) -> Optional[pulumi.Input[int]]:
        """
        Subscription plan disk size
        """
        return pulumi.get(self, "disk_size")

    @disk_size.setter
    def disk_size(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "disk_size", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


@pulumi.input_type
class InstanceCopySettingArgs:
    def __init__(__self__, *,
                 settings: pulumi.Input[Sequence[pulumi.Input[str]]],
                 subscription_id: pulumi.Input[str]):
        """
        :param pulumi.Input[Sequence[pulumi.Input[str]]] settings: Array of one or more settings to be copied. Allowed values: [alarms, config, definitions, firewall, logs, metrics, plugins]
               
               See more below, copy settings
        :param pulumi.Input[str] subscription_id: Instance identifier of the CloudAMQP instance to copy the settings from.
        """
        pulumi.set(__self__, "settings", settings)
        pulumi.set(__self__, "subscription_id", subscription_id)

    @property
    @pulumi.getter
    def settings(self) -> pulumi.Input[Sequence[pulumi.Input[str]]]:
        """
        Array of one or more settings to be copied. Allowed values: [alarms, config, definitions, firewall, logs, metrics, plugins]

        See more below, copy settings
        """
        return pulumi.get(self, "settings")

    @settings.setter
    def settings(self, value: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(self, "settings", value)

    @property
    @pulumi.getter(name="subscriptionId")
    def subscription_id(self) -> pulumi.Input[str]:
        """
        Instance identifier of the CloudAMQP instance to copy the settings from.
        """
        return pulumi.get(self, "subscription_id")

    @subscription_id.setter
    def subscription_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "subscription_id", value)


@pulumi.input_type
class SecurityFirewallRuleArgs:
    def __init__(__self__, *,
                 ip: pulumi.Input[str],
                 description: Optional[pulumi.Input[str]] = None,
                 ports: Optional[pulumi.Input[Sequence[pulumi.Input[int]]]] = None,
                 services: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        """
        :param pulumi.Input[str] ip: CIDR address: IP address with CIDR notation (e.g. 10.56.72.0/24)
        :param pulumi.Input[str] description: Naming descripton e.g. 'Default'
        :param pulumi.Input[Sequence[pulumi.Input[int]]] ports: Custom ports between 0 - 65554
        :param pulumi.Input[Sequence[pulumi.Input[str]]] services: Pre-defined services 'AMQP', 'AMQPS', 'HTTPS', 'MQTT', 'MQTTS', 'STOMP', 'STOMPS', 'STREAM', 'STREAM_SSL'
        """
        pulumi.set(__self__, "ip", ip)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if ports is not None:
            pulumi.set(__self__, "ports", ports)
        if services is not None:
            pulumi.set(__self__, "services", services)

    @property
    @pulumi.getter
    def ip(self) -> pulumi.Input[str]:
        """
        CIDR address: IP address with CIDR notation (e.g. 10.56.72.0/24)
        """
        return pulumi.get(self, "ip")

    @ip.setter
    def ip(self, value: pulumi.Input[str]):
        pulumi.set(self, "ip", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Naming descripton e.g. 'Default'
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def ports(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[int]]]]:
        """
        Custom ports between 0 - 65554
        """
        return pulumi.get(self, "ports")

    @ports.setter
    def ports(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[int]]]]):
        pulumi.set(self, "ports", value)

    @property
    @pulumi.getter
    def services(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Pre-defined services 'AMQP', 'AMQPS', 'HTTPS', 'MQTT', 'MQTTS', 'STOMP', 'STOMPS', 'STREAM', 'STREAM_SSL'
        """
        return pulumi.get(self, "services")

    @services.setter
    def services(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "services", value)


