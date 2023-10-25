# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Callable, Mapping, Optional, Sequence, Union, overload
from . import _utilities
from . import outputs

__all__ = [
    'GetAccountResult',
    'AwaitableGetAccountResult',
    'get_account',
    'get_account_output',
]

@pulumi.output_type
class GetAccountResult:
    """
    A collection of values returned by getAccount.
    """
    def __init__(__self__, id=None, instances=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if instances and not isinstance(instances, list):
            raise TypeError("Expected argument 'instances' to be a list")
        pulumi.set(__self__, "instances", instances)

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def instances(self) -> Sequence['outputs.GetAccountInstanceResult']:
        return pulumi.get(self, "instances")


class AwaitableGetAccountResult(GetAccountResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetAccountResult(
            id=self.id,
            instances=self.instances)


def get_account(opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetAccountResult:
    """
    Use this data source to retrieve basic information about all instances available for an account. Uses the included apikey in provider configuration, to determine which account to read from.

    ## Attributes reference

    All attributes reference are computed

    * `id`          - The identifier for this data source. Set to `na` since there is no unique identifier.
    * `instances`   - An array of instances. Each `instances` block consists of the fields documented below.

    ***

    The `instances` block consist of

    * `id`      - The instance identifier.
    * `name`    - The name of the instance.
    * `plan`    - The subscription plan used for the instance.
    * `region`  - The region were the instanece is located in.
    * `tags`    - Optional tags set for the instance.

    ## Dependency

    This data source depends on apikey set in the provider configuration.
    """
    __args__ = dict()
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('cloudamqp:index/getAccount:getAccount', __args__, opts=opts, typ=GetAccountResult).value

    return AwaitableGetAccountResult(
        id=pulumi.get(__ret__, 'id'),
        instances=pulumi.get(__ret__, 'instances'))


@_utilities.lift_output_func(get_account)
def get_account_output(opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetAccountResult]:
    """
    Use this data source to retrieve basic information about all instances available for an account. Uses the included apikey in provider configuration, to determine which account to read from.

    ## Attributes reference

    All attributes reference are computed

    * `id`          - The identifier for this data source. Set to `na` since there is no unique identifier.
    * `instances`   - An array of instances. Each `instances` block consists of the fields documented below.

    ***

    The `instances` block consist of

    * `id`      - The instance identifier.
    * `name`    - The name of the instance.
    * `plan`    - The subscription plan used for the instance.
    * `region`  - The region were the instanece is located in.
    * `tags`    - Optional tags set for the instance.

    ## Dependency

    This data source depends on apikey set in the provider configuration.
    """
    ...
