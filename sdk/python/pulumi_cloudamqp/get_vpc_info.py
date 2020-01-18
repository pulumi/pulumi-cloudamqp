# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from . import utilities, tables

class GetVpcInfoResult:
    """
    A collection of values returned by getVpcInfo.
    """
    def __init__(__self__, instance_id=None, name=None, owner_id=None, security_group_id=None, vpc_subnet=None, id=None):
        if instance_id and not isinstance(instance_id, float):
            raise TypeError("Expected argument 'instance_id' to be a float")
        __self__.instance_id = instance_id
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        __self__.name = name
        if owner_id and not isinstance(owner_id, str):
            raise TypeError("Expected argument 'owner_id' to be a str")
        __self__.owner_id = owner_id
        if security_group_id and not isinstance(security_group_id, str):
            raise TypeError("Expected argument 'security_group_id' to be a str")
        __self__.security_group_id = security_group_id
        if vpc_subnet and not isinstance(vpc_subnet, str):
            raise TypeError("Expected argument 'vpc_subnet' to be a str")
        __self__.vpc_subnet = vpc_subnet
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        __self__.id = id
        """
        id is the provider-assigned unique ID for this managed resource.
        """
class AwaitableGetVpcInfoResult(GetVpcInfoResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetVpcInfoResult(
            instance_id=self.instance_id,
            name=self.name,
            owner_id=self.owner_id,
            security_group_id=self.security_group_id,
            vpc_subnet=self.vpc_subnet,
            id=self.id)

def get_vpc_info(instance_id=None,opts=None):
    """
    Use this data source to access information about an existing resource.
    
    """
    __args__ = dict()

    __args__['instanceId'] = instance_id
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = utilities.get_version()
    __ret__ = pulumi.runtime.invoke('cloudamqp:index/getVpcInfo:getVpcInfo', __args__, opts=opts).value

    return AwaitableGetVpcInfoResult(
        instance_id=__ret__.get('instanceId'),
        name=__ret__.get('name'),
        owner_id=__ret__.get('ownerId'),
        security_group_id=__ret__.get('securityGroupId'),
        vpc_subnet=__ret__.get('vpcSubnet'),
        id=__ret__.get('id'))
