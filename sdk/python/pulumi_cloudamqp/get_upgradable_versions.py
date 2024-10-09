# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

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
    'GetUpgradableVersionsResult',
    'AwaitableGetUpgradableVersionsResult',
    'get_upgradable_versions',
    'get_upgradable_versions_output',
]

@pulumi.output_type
class GetUpgradableVersionsResult:
    """
    A collection of values returned by getUpgradableVersions.
    """
    def __init__(__self__, id=None, instance_id=None, new_erlang_version=None, new_rabbitmq_version=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if instance_id and not isinstance(instance_id, int):
            raise TypeError("Expected argument 'instance_id' to be a int")
        pulumi.set(__self__, "instance_id", instance_id)
        if new_erlang_version and not isinstance(new_erlang_version, str):
            raise TypeError("Expected argument 'new_erlang_version' to be a str")
        pulumi.set(__self__, "new_erlang_version", new_erlang_version)
        if new_rabbitmq_version and not isinstance(new_rabbitmq_version, str):
            raise TypeError("Expected argument 'new_rabbitmq_version' to be a str")
        pulumi.set(__self__, "new_rabbitmq_version", new_rabbitmq_version)

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
    @pulumi.getter(name="newErlangVersion")
    def new_erlang_version(self) -> str:
        return pulumi.get(self, "new_erlang_version")

    @property
    @pulumi.getter(name="newRabbitmqVersion")
    def new_rabbitmq_version(self) -> str:
        return pulumi.get(self, "new_rabbitmq_version")


class AwaitableGetUpgradableVersionsResult(GetUpgradableVersionsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetUpgradableVersionsResult(
            id=self.id,
            instance_id=self.instance_id,
            new_erlang_version=self.new_erlang_version,
            new_rabbitmq_version=self.new_rabbitmq_version)


def get_upgradable_versions(instance_id: Optional[int] = None,
                            opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetUpgradableVersionsResult:
    """
    Use this data source to retrieve information about possible upgradable versions for RabbitMQ and Erlang.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_cloudamqp as cloudamqp

    versions = cloudamqp.get_upgradable_versions(instance_id=instance["id"])
    ```

    ## Attributes reference

    All attributes reference are computed

    * `new_rabbitmq_version`  - Possible upgradable version for RabbitMQ.
    * `new_erlang_version`    - Possible upgradable version for Erlang.

    ## Dependency

    This data source depends on CloudAMQP instance identifier, `cloudamqp_instance.instance.id`.


    :param int instance_id: The CloudAMQP instance identifier.
    """
    __args__ = dict()
    __args__['instanceId'] = instance_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('cloudamqp:index/getUpgradableVersions:getUpgradableVersions', __args__, opts=opts, typ=GetUpgradableVersionsResult).value

    return AwaitableGetUpgradableVersionsResult(
        id=pulumi.get(__ret__, 'id'),
        instance_id=pulumi.get(__ret__, 'instance_id'),
        new_erlang_version=pulumi.get(__ret__, 'new_erlang_version'),
        new_rabbitmq_version=pulumi.get(__ret__, 'new_rabbitmq_version'))
def get_upgradable_versions_output(instance_id: Optional[pulumi.Input[int]] = None,
                                   opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetUpgradableVersionsResult]:
    """
    Use this data source to retrieve information about possible upgradable versions for RabbitMQ and Erlang.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_cloudamqp as cloudamqp

    versions = cloudamqp.get_upgradable_versions(instance_id=instance["id"])
    ```

    ## Attributes reference

    All attributes reference are computed

    * `new_rabbitmq_version`  - Possible upgradable version for RabbitMQ.
    * `new_erlang_version`    - Possible upgradable version for Erlang.

    ## Dependency

    This data source depends on CloudAMQP instance identifier, `cloudamqp_instance.instance.id`.


    :param int instance_id: The CloudAMQP instance identifier.
    """
    __args__ = dict()
    __args__['instanceId'] = instance_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('cloudamqp:index/getUpgradableVersions:getUpgradableVersions', __args__, opts=opts, typ=GetUpgradableVersionsResult)
    return __ret__.apply(lambda __response__: GetUpgradableVersionsResult(
        id=pulumi.get(__response__, 'id'),
        instance_id=pulumi.get(__response__, 'instance_id'),
        new_erlang_version=pulumi.get(__response__, 'new_erlang_version'),
        new_rabbitmq_version=pulumi.get(__response__, 'new_rabbitmq_version')))
