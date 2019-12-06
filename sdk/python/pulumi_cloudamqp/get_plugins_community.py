# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from . import utilities, tables

class GetPluginsCommunityResult:
    """
    A collection of values returned by getPluginsCommunity.
    """
    def __init__(__self__, instance_id=None, plugins=None, id=None):
        if instance_id and not isinstance(instance_id, float):
            raise TypeError("Expected argument 'instance_id' to be a float")
        __self__.instance_id = instance_id
        if plugins and not isinstance(plugins, list):
            raise TypeError("Expected argument 'plugins' to be a list")
        __self__.plugins = plugins
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        __self__.id = id
        """
        id is the provider-assigned unique ID for this managed resource.
        """
class AwaitableGetPluginsCommunityResult(GetPluginsCommunityResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetPluginsCommunityResult(
            instance_id=self.instance_id,
            plugins=self.plugins,
            id=self.id)

def get_plugins_community(instance_id=None,plugins=None,opts=None):
    """
    Use this data source to access information about an existing resource.
    
    
    The **plugins** object supports the following:
    
      * `description` (`str`)
      * `name` (`str`)
      * `require` (`str`)
    """
    __args__ = dict()

    __args__['instanceId'] = instance_id
    __args__['plugins'] = plugins
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = utilities.get_version()
    __ret__ = pulumi.runtime.invoke('cloudamqp:index/getPluginsCommunity:getPluginsCommunity', __args__, opts=opts).value

    return AwaitableGetPluginsCommunityResult(
        instance_id=__ret__.get('instanceId'),
        plugins=__ret__.get('plugins'),
        id=__ret__.get('id'))
