# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from . import _utilities, _tables

__all__ = [
    'GetCredentialsResult',
    'AwaitableGetCredentialsResult',
    'get_credentials',
]

@pulumi.output_type
class GetCredentialsResult:
    """
    A collection of values returned by getCredentials.
    """
    def __init__(__self__, id=None, instance_id=None, password=None, username=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if instance_id and not isinstance(instance_id, int):
            raise TypeError("Expected argument 'instance_id' to be a int")
        pulumi.set(__self__, "instance_id", instance_id)
        if password and not isinstance(password, str):
            raise TypeError("Expected argument 'password' to be a str")
        pulumi.set(__self__, "password", password)
        if username and not isinstance(username, str):
            raise TypeError("Expected argument 'username' to be a str")
        pulumi.set(__self__, "username", username)

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
    def password(self) -> Optional[str]:
        return pulumi.get(self, "password")

    @property
    @pulumi.getter
    def username(self) -> Optional[str]:
        return pulumi.get(self, "username")


class AwaitableGetCredentialsResult(GetCredentialsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetCredentialsResult(
            id=self.id,
            instance_id=self.instance_id,
            password=self.password,
            username=self.username)


def get_credentials(instance_id: Optional[int] = None,
                    password: Optional[str] = None,
                    username: Optional[str] = None,
                    opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetCredentialsResult:
    """
    Use this data source to retrieve information about the credentials of the configured user in Rabbit MQ. Information is extracted from `cloudamqp_instance.instance.url`.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_cloudamqp as cloudamqp

    credentials = cloudamqp.get_credentials(instance_id=cloudamqp_instance["instance"]["id"])
    ```
    ## Argument reference

    * `instance_id` - (Required) The CloudAMQP instance identifier.

    ## Attribute reference

    * `username`    - (Computed/Sensitive) The username for the configured user in Rabbit MQ.
    * `password`    - (Computed/Sensitive) The password used by the `username`.

    ## Dependency

    This data source depends on CloudAMQP instance identifier, `cloudamqp_instance.instance.id`.
    """
    __args__ = dict()
    __args__['instanceId'] = instance_id
    __args__['password'] = password
    __args__['username'] = username
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('cloudamqp:index/getCredentials:getCredentials', __args__, opts=opts, typ=GetCredentialsResult).value

    return AwaitableGetCredentialsResult(
        id=__ret__.id,
        instance_id=__ret__.instance_id,
        password=__ret__.password,
        username=__ret__.username)
