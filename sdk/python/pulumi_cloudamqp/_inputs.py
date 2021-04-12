# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities, _tables

__all__ = [
    'SecurityFirewallRuleArgs',
    'GetNodesNodeArgs',
    'GetPluginsCommunityPluginArgs',
    'GetPluginsPluginArgs',
]

@pulumi.input_type
class SecurityFirewallRuleArgs:
    def __init__(__self__, *,
                 ip: pulumi.Input[str],
                 description: Optional[pulumi.Input[str]] = None,
                 ports: Optional[pulumi.Input[Sequence[pulumi.Input[int]]]] = None,
                 services: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        """
        :param pulumi.Input[str] ip: Source ip and netmask for the rule. (e.g. 10.56.72.0/24)
        :param pulumi.Input[str] description: Description name of the rule. e.g. Default.
        :param pulumi.Input[Sequence[pulumi.Input[int]]] ports: Custom ports to be opened
        :param pulumi.Input[Sequence[pulumi.Input[str]]] services: Pre-defined service ports
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
        Source ip and netmask for the rule. (e.g. 10.56.72.0/24)
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
        Pre-defined service ports
        """
        return pulumi.get(self, "services")

    @services.setter
    def services(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "services", value)


@pulumi.input_type
class GetNodesNodeArgs:
    def __init__(__self__, *,
                 erlang_version: str,
                 hipe: bool,
                 hostname: str,
                 name: str,
                 rabbitmq_version: str,
                 running: bool):
        pulumi.set(__self__, "erlang_version", erlang_version)
        pulumi.set(__self__, "hipe", hipe)
        pulumi.set(__self__, "hostname", hostname)
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "rabbitmq_version", rabbitmq_version)
        pulumi.set(__self__, "running", running)

    @property
    @pulumi.getter(name="erlangVersion")
    def erlang_version(self) -> str:
        return pulumi.get(self, "erlang_version")

    @erlang_version.setter
    def erlang_version(self, value: str):
        pulumi.set(self, "erlang_version", value)

    @property
    @pulumi.getter
    def hipe(self) -> bool:
        return pulumi.get(self, "hipe")

    @hipe.setter
    def hipe(self, value: bool):
        pulumi.set(self, "hipe", value)

    @property
    @pulumi.getter
    def hostname(self) -> str:
        return pulumi.get(self, "hostname")

    @hostname.setter
    def hostname(self, value: str):
        pulumi.set(self, "hostname", value)

    @property
    @pulumi.getter
    def name(self) -> str:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: str):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="rabbitmqVersion")
    def rabbitmq_version(self) -> str:
        return pulumi.get(self, "rabbitmq_version")

    @rabbitmq_version.setter
    def rabbitmq_version(self, value: str):
        pulumi.set(self, "rabbitmq_version", value)

    @property
    @pulumi.getter
    def running(self) -> bool:
        return pulumi.get(self, "running")

    @running.setter
    def running(self, value: bool):
        pulumi.set(self, "running", value)


@pulumi.input_type
class GetPluginsCommunityPluginArgs:
    def __init__(__self__, *,
                 description: Optional[str] = None,
                 name: Optional[str] = None,
                 require: Optional[str] = None):
        if description is not None:
            pulumi.set(__self__, "description", description)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if require is not None:
            pulumi.set(__self__, "require", require)

    @property
    @pulumi.getter
    def description(self) -> Optional[str]:
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[str]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[str]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def require(self) -> Optional[str]:
        return pulumi.get(self, "require")

    @require.setter
    def require(self, value: Optional[str]):
        pulumi.set(self, "require", value)


@pulumi.input_type
class GetPluginsPluginArgs:
    def __init__(__self__, *,
                 description: Optional[str] = None,
                 enabled: Optional[bool] = None,
                 name: Optional[str] = None,
                 version: Optional[str] = None):
        if description is not None:
            pulumi.set(__self__, "description", description)
        if enabled is not None:
            pulumi.set(__self__, "enabled", enabled)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if version is not None:
            pulumi.set(__self__, "version", version)

    @property
    @pulumi.getter
    def description(self) -> Optional[str]:
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[str]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def enabled(self) -> Optional[bool]:
        return pulumi.get(self, "enabled")

    @enabled.setter
    def enabled(self, value: Optional[bool]):
        pulumi.set(self, "enabled", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[str]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def version(self) -> Optional[str]:
        return pulumi.get(self, "version")

    @version.setter
    def version(self, value: Optional[str]):
        pulumi.set(self, "version", value)


