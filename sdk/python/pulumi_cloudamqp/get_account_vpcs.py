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
    'GetAccountVpcsResult',
    'AwaitableGetAccountVpcsResult',
    'get_account_vpcs',
    'get_account_vpcs_output',
]

@pulumi.output_type
class GetAccountVpcsResult:
    """
    A collection of values returned by getAccountVpcs.
    """
    def __init__(__self__, id=None, vpcs=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if vpcs and not isinstance(vpcs, list):
            raise TypeError("Expected argument 'vpcs' to be a list")
        pulumi.set(__self__, "vpcs", vpcs)

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def vpcs(self) -> Sequence['outputs.GetAccountVpcsVpcResult']:
        return pulumi.get(self, "vpcs")


class AwaitableGetAccountVpcsResult(GetAccountVpcsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetAccountVpcsResult(
            id=self.id,
            vpcs=self.vpcs)


def get_account_vpcs(opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetAccountVpcsResult:
    """
    Use this data source to retrieve basic information about all standalone VPCs available for an account. Uses the included apikey in provider configuration to determine which account to read from.

    ## Example Usage

    Can be used in other resources/data sources when the VPC identifier is unknown, while other attributes are known. E.g. find correct VPC using the `name` you gave your VPC. Then iterate over VPCs to find the matching one and extract the VPC identifier.

    ```python
    import pulumi
    import pulumi_cloudamqp as cloudamqp

    my_vpc_name = "<your VPC name>"
    vpc_list = cloudamqp.get_account_vpcs()
    pulumi.export("vpcId", [vpc for vpc in vpc_list.vpcs if vpc.name == my_vpc_name][0].id)
    ```
    ## Attributes reference

    All attributes reference are computed

    * `id`      - The identifier for this data source. Set to `na` since there is no unique identifier.
    * `vpcs`    - An array of VPCs. Each `vpcs` block consists of the fields documented below.

    ***

    The `vpcs` block consist of

    * `id`          - The VPC identifier.
    * `name`        - The VPC instance name.
    * `region`      - The region the VPC is hosted in.
    * `subnet`      - The VPC subnet.
    * `tags`        - Optional tags set for the VPC.
    * `vpc_name`    - VPC name given when hosted at the cloud provider.

    ## Dependency

    This data source depends on apikey set in the provider configuration.
    """
    __args__ = dict()
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('cloudamqp:index/getAccountVpcs:getAccountVpcs', __args__, opts=opts, typ=GetAccountVpcsResult).value

    return AwaitableGetAccountVpcsResult(
        id=pulumi.get(__ret__, 'id'),
        vpcs=pulumi.get(__ret__, 'vpcs'))


@_utilities.lift_output_func(get_account_vpcs)
def get_account_vpcs_output(opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetAccountVpcsResult]:
    """
    Use this data source to retrieve basic information about all standalone VPCs available for an account. Uses the included apikey in provider configuration to determine which account to read from.

    ## Example Usage

    Can be used in other resources/data sources when the VPC identifier is unknown, while other attributes are known. E.g. find correct VPC using the `name` you gave your VPC. Then iterate over VPCs to find the matching one and extract the VPC identifier.

    ```python
    import pulumi
    import pulumi_cloudamqp as cloudamqp

    my_vpc_name = "<your VPC name>"
    vpc_list = cloudamqp.get_account_vpcs()
    pulumi.export("vpcId", [vpc for vpc in vpc_list.vpcs if vpc.name == my_vpc_name][0].id)
    ```
    ## Attributes reference

    All attributes reference are computed

    * `id`      - The identifier for this data source. Set to `na` since there is no unique identifier.
    * `vpcs`    - An array of VPCs. Each `vpcs` block consists of the fields documented below.

    ***

    The `vpcs` block consist of

    * `id`          - The VPC identifier.
    * `name`        - The VPC instance name.
    * `region`      - The region the VPC is hosted in.
    * `subnet`      - The VPC subnet.
    * `tags`        - Optional tags set for the VPC.
    * `vpc_name`    - VPC name given when hosted at the cloud provider.

    ## Dependency

    This data source depends on apikey set in the provider configuration.
    """
    ...
