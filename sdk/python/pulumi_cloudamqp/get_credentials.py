# coding=utf-8
# *** WARNING: this file was generated by pulumi-language-python. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import builtins
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
    'GetCredentialsResult',
    'AwaitableGetCredentialsResult',
    'get_credentials',
    'get_credentials_output',
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
    def id(self) -> builtins.str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> builtins.int:
        return pulumi.get(self, "instance_id")

    @property
    @pulumi.getter
    def password(self) -> builtins.str:
        """
        (Sensitive) The password used by the `username`.
        """
        return pulumi.get(self, "password")

    @property
    @pulumi.getter
    def username(self) -> builtins.str:
        """
        (Sensitive) The username for the configured user in Rabbit MQ.
        """
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


def get_credentials(instance_id: Optional[builtins.int] = None,
                    opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetCredentialsResult:
    """
    Use this data source to retrieve information about the credentials of the configured user in
    RabbitMQ. Information is extracted from `cloudamqp_instance.instance.url`.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_cloudamqp as cloudamqp

    credentials = cloudamqp.get_credentials(instance_id=instance["id"])
    ```

    ## Dependency

    This data source depends on CloudAMQP instance identifier, `cloudamqp_instance.instance.id`.


    :param builtins.int instance_id: The CloudAMQP instance identifier.
    """
    __args__ = dict()
    __args__['instanceId'] = instance_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('cloudamqp:index/getCredentials:getCredentials', __args__, opts=opts, typ=GetCredentialsResult).value

    return AwaitableGetCredentialsResult(
        id=pulumi.get(__ret__, 'id'),
        instance_id=pulumi.get(__ret__, 'instance_id'),
        password=pulumi.get(__ret__, 'password'),
        username=pulumi.get(__ret__, 'username'))
def get_credentials_output(instance_id: Optional[pulumi.Input[builtins.int]] = None,
                           opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetCredentialsResult]:
    """
    Use this data source to retrieve information about the credentials of the configured user in
    RabbitMQ. Information is extracted from `cloudamqp_instance.instance.url`.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_cloudamqp as cloudamqp

    credentials = cloudamqp.get_credentials(instance_id=instance["id"])
    ```

    ## Dependency

    This data source depends on CloudAMQP instance identifier, `cloudamqp_instance.instance.id`.


    :param builtins.int instance_id: The CloudAMQP instance identifier.
    """
    __args__ = dict()
    __args__['instanceId'] = instance_id
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('cloudamqp:index/getCredentials:getCredentials', __args__, opts=opts, typ=GetCredentialsResult)
    return __ret__.apply(lambda __response__: GetCredentialsResult(
        id=pulumi.get(__response__, 'id'),
        instance_id=pulumi.get(__response__, 'instance_id'),
        password=pulumi.get(__response__, 'password'),
        username=pulumi.get(__response__, 'username')))
