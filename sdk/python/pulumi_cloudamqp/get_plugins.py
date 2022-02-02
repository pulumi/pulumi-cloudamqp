# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities
from . import outputs

__all__ = [
    'GetPluginsResult',
    'AwaitableGetPluginsResult',
    'get_plugins',
    'get_plugins_output',
]

@pulumi.output_type
class GetPluginsResult:
    """
    A collection of values returned by getPlugins.
    """
    def __init__(__self__, id=None, instance_id=None, plugins=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if instance_id and not isinstance(instance_id, int):
            raise TypeError("Expected argument 'instance_id' to be a int")
        pulumi.set(__self__, "instance_id", instance_id)
        if plugins and not isinstance(plugins, list):
            raise TypeError("Expected argument 'plugins' to be a list")
        pulumi.set(__self__, "plugins", plugins)

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
    def plugins(self) -> Sequence['outputs.GetPluginsPluginResult']:
        return pulumi.get(self, "plugins")


class AwaitableGetPluginsResult(GetPluginsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetPluginsResult(
            id=self.id,
            instance_id=self.instance_id,
            plugins=self.plugins)


def get_plugins(instance_id: Optional[int] = None,
                opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetPluginsResult:
    """
    Use this data source to retrieve information about installed and available plugins for the CloudAMQP instance.

    ⚠️  From our go API wrapper [v1.4.0](https://github.com/84codes/go-api/releases/tag/v1.4.0) there is support for multiple retries when requesting information about plugins. This was introduced to avoid `ReadPlugin error 400: Timeout talking to backend`.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_cloudamqp as cloudamqp

    plugins = cloudamqp.get_plugins(instance_id=cloudamqp_instance["instance"]["id"])
    ```
    ## Attributes reference

    All attributes reference are computed

    * `id`      - The identifier for this resource.
    * `plugins` - An array of plugins. Each `plugins` block consists of the fields documented below.

    ***

    The `plugins` block consist of

    * `name`        - The type of the recipient.
    * `version`     - Rabbit MQ version that the plugins are shipped with.
    * `description` - Description of what the plugin does.
    * `enabled`     - Enable or disable information for the plugin.

    ## Dependency

    This data source depends on CloudAMQP instance identifier, `cloudamqp_instance.instance.id`.


    :param int instance_id: The CloudAMQP instance identifier.
    """
    __args__ = dict()
    __args__['instanceId'] = instance_id
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('cloudamqp:index/getPlugins:getPlugins', __args__, opts=opts, typ=GetPluginsResult).value

    return AwaitableGetPluginsResult(
        id=__ret__.id,
        instance_id=__ret__.instance_id,
        plugins=__ret__.plugins)


@_utilities.lift_output_func(get_plugins)
def get_plugins_output(instance_id: Optional[pulumi.Input[int]] = None,
                       opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetPluginsResult]:
    """
    Use this data source to retrieve information about installed and available plugins for the CloudAMQP instance.

    ⚠️  From our go API wrapper [v1.4.0](https://github.com/84codes/go-api/releases/tag/v1.4.0) there is support for multiple retries when requesting information about plugins. This was introduced to avoid `ReadPlugin error 400: Timeout talking to backend`.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_cloudamqp as cloudamqp

    plugins = cloudamqp.get_plugins(instance_id=cloudamqp_instance["instance"]["id"])
    ```
    ## Attributes reference

    All attributes reference are computed

    * `id`      - The identifier for this resource.
    * `plugins` - An array of plugins. Each `plugins` block consists of the fields documented below.

    ***

    The `plugins` block consist of

    * `name`        - The type of the recipient.
    * `version`     - Rabbit MQ version that the plugins are shipped with.
    * `description` - Description of what the plugin does.
    * `enabled`     - Enable or disable information for the plugin.

    ## Dependency

    This data source depends on CloudAMQP instance identifier, `cloudamqp_instance.instance.id`.


    :param int instance_id: The CloudAMQP instance identifier.
    """
    ...
