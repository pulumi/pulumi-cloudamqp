# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = [
    'SecurityFirewallRule',
    'GetAccountInstanceResult',
    'GetNodesNodeResult',
    'GetPluginsCommunityPluginResult',
    'GetPluginsPluginResult',
]

@pulumi.output_type
class SecurityFirewallRule(dict):
    def __init__(__self__, *,
                 ip: str,
                 description: Optional[str] = None,
                 ports: Optional[Sequence[int]] = None,
                 services: Optional[Sequence[str]] = None):
        """
        :param str ip: Source ip and netmask for the rule. (e.g. 10.56.72.0/24)
        :param str description: Description name of the rule. e.g. Default.
        :param Sequence[int] ports: Custom ports to be opened
        :param Sequence[str] services: Pre-defined service ports, see table below
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
    def ip(self) -> str:
        """
        Source ip and netmask for the rule. (e.g. 10.56.72.0/24)
        """
        return pulumi.get(self, "ip")

    @property
    @pulumi.getter
    def description(self) -> Optional[str]:
        """
        Description name of the rule. e.g. Default.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def ports(self) -> Optional[Sequence[int]]:
        """
        Custom ports to be opened
        """
        return pulumi.get(self, "ports")

    @property
    @pulumi.getter
    def services(self) -> Optional[Sequence[str]]:
        """
        Pre-defined service ports, see table below
        """
        return pulumi.get(self, "services")


@pulumi.output_type
class GetAccountInstanceResult(dict):
    def __init__(__self__, *,
                 id: int,
                 name: str,
                 plan: str,
                 region: str,
                 tags: Optional[Sequence[str]] = None):
        pulumi.set(__self__, "id", id)
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "plan", plan)
        pulumi.set(__self__, "region", region)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter
    def id(self) -> int:
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def name(self) -> str:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def plan(self) -> str:
        return pulumi.get(self, "plan")

    @property
    @pulumi.getter
    def region(self) -> str:
        return pulumi.get(self, "region")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Sequence[str]]:
        return pulumi.get(self, "tags")


@pulumi.output_type
class GetNodesNodeResult(dict):
    def __init__(__self__, *,
                 configured: bool,
                 erlang_version: str,
                 hipe: bool,
                 hostname: str,
                 name: str,
                 rabbitmq_version: str,
                 running: bool):
        pulumi.set(__self__, "configured", configured)
        pulumi.set(__self__, "erlang_version", erlang_version)
        pulumi.set(__self__, "hipe", hipe)
        pulumi.set(__self__, "hostname", hostname)
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "rabbitmq_version", rabbitmq_version)
        pulumi.set(__self__, "running", running)

    @property
    @pulumi.getter
    def configured(self) -> bool:
        return pulumi.get(self, "configured")

    @property
    @pulumi.getter(name="erlangVersion")
    def erlang_version(self) -> str:
        return pulumi.get(self, "erlang_version")

    @property
    @pulumi.getter
    def hipe(self) -> bool:
        return pulumi.get(self, "hipe")

    @property
    @pulumi.getter
    def hostname(self) -> str:
        return pulumi.get(self, "hostname")

    @property
    @pulumi.getter
    def name(self) -> str:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="rabbitmqVersion")
    def rabbitmq_version(self) -> str:
        return pulumi.get(self, "rabbitmq_version")

    @property
    @pulumi.getter
    def running(self) -> bool:
        return pulumi.get(self, "running")


@pulumi.output_type
class GetPluginsCommunityPluginResult(dict):
    def __init__(__self__, *,
                 description: str,
                 name: str,
                 require: str):
        pulumi.set(__self__, "description", description)
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "require", require)

    @property
    @pulumi.getter
    def description(self) -> str:
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def name(self) -> str:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def require(self) -> str:
        return pulumi.get(self, "require")


@pulumi.output_type
class GetPluginsPluginResult(dict):
    def __init__(__self__, *,
                 description: str,
                 enabled: bool,
                 name: str,
                 version: str):
        pulumi.set(__self__, "description", description)
        pulumi.set(__self__, "enabled", enabled)
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "version", version)

    @property
    @pulumi.getter
    def description(self) -> str:
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def enabled(self) -> bool:
        return pulumi.get(self, "enabled")

    @property
    @pulumi.getter
    def name(self) -> str:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def version(self) -> str:
        return pulumi.get(self, "version")


