# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Callable, Mapping, Optional, Sequence, Union, overload
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
        ExtraDiskSizeNodeArgs._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            additional_disk_size=additional_disk_size,
            disk_size=disk_size,
            name=name,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             additional_disk_size: Optional[pulumi.Input[int]] = None,
             disk_size: Optional[pulumi.Input[int]] = None,
             name: Optional[pulumi.Input[str]] = None,
             opts: Optional[pulumi.ResourceOptions]=None,
             **kwargs):
        if additional_disk_size is None and 'additionalDiskSize' in kwargs:
            additional_disk_size = kwargs['additionalDiskSize']
        if disk_size is None and 'diskSize' in kwargs:
            disk_size = kwargs['diskSize']

        if additional_disk_size is not None:
            _setter("additional_disk_size", additional_disk_size)
        if disk_size is not None:
            _setter("disk_size", disk_size)
        if name is not None:
            _setter("name", name)

    @property
    @pulumi.getter(name="additionalDiskSize")
    def additional_disk_size(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "additional_disk_size")

    @additional_disk_size.setter
    def additional_disk_size(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "additional_disk_size", value)

    @property
    @pulumi.getter(name="diskSize")
    def disk_size(self) -> Optional[pulumi.Input[int]]:
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
        InstanceCopySettingArgs._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            settings=settings,
            subscription_id=subscription_id,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             settings: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
             subscription_id: Optional[pulumi.Input[str]] = None,
             opts: Optional[pulumi.ResourceOptions]=None,
             **kwargs):
        if settings is None:
            raise TypeError("Missing 'settings' argument")
        if subscription_id is None and 'subscriptionId' in kwargs:
            subscription_id = kwargs['subscriptionId']
        if subscription_id is None:
            raise TypeError("Missing 'subscription_id' argument")

        _setter("settings", settings)
        _setter("subscription_id", subscription_id)

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
        :param pulumi.Input[str] description: Description name of the rule. e.g. Default.
               
               Pre-defined services for RabbitMQ:
               
               | Service name | Port  |
               |--------------|-------|
               | AMQP         | 5672  |
               | AMQPS        | 5671  |
               | HTTPS        | 443   |
               | MQTT         | 1883  |
               | MQTTS        | 8883  |
               | STOMP        | 61613 |
               | STOMPS       | 61614 |
               | STREAM       | 5552  |
               | STREAM_SSL   | 5551  |
               
               Pre-defined services for LavinMQ:
               
               | Service name | Port  |
               |--------------|-------|
               | AMQP         | 5672  |
               | AMQPS        | 5671  |
               | HTTPS        | 443   |
        :param pulumi.Input[Sequence[pulumi.Input[int]]] ports: Custom ports to be opened
        :param pulumi.Input[Sequence[pulumi.Input[str]]] services: Pre-defined service ports, see table below
        """
        SecurityFirewallRuleArgs._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            ip=ip,
            description=description,
            ports=ports,
            services=services,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             ip: Optional[pulumi.Input[str]] = None,
             description: Optional[pulumi.Input[str]] = None,
             ports: Optional[pulumi.Input[Sequence[pulumi.Input[int]]]] = None,
             services: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
             opts: Optional[pulumi.ResourceOptions]=None,
             **kwargs):
        if ip is None:
            raise TypeError("Missing 'ip' argument")

        _setter("ip", ip)
        if description is not None:
            _setter("description", description)
        if ports is not None:
            _setter("ports", ports)
        if services is not None:
            _setter("services", services)

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
        Description name of the rule. e.g. Default.

        Pre-defined services for RabbitMQ:

        | Service name | Port  |
        |--------------|-------|
        | AMQP         | 5672  |
        | AMQPS        | 5671  |
        | HTTPS        | 443   |
        | MQTT         | 1883  |
        | MQTTS        | 8883  |
        | STOMP        | 61613 |
        | STOMPS       | 61614 |
        | STREAM       | 5552  |
        | STREAM_SSL   | 5551  |

        Pre-defined services for LavinMQ:

        | Service name | Port  |
        |--------------|-------|
        | AMQP         | 5672  |
        | AMQPS        | 5671  |
        | HTTPS        | 443   |
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def ports(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[int]]]]:
        """
        Custom ports to be opened
        """
        return pulumi.get(self, "ports")

    @ports.setter
    def ports(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[int]]]]):
        pulumi.set(self, "ports", value)

    @property
    @pulumi.getter
    def services(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Pre-defined service ports, see table below
        """
        return pulumi.get(self, "services")

    @services.setter
    def services(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "services", value)


