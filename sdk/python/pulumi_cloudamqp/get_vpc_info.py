# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = [
    'GetVpcInfoResult',
    'AwaitableGetVpcInfoResult',
    'get_vpc_info',
    'get_vpc_info_output',
]

@pulumi.output_type
class GetVpcInfoResult:
    """
    A collection of values returned by getVpcInfo.
    """
    def __init__(__self__, id=None, instance_id=None, name=None, owner_id=None, security_group_id=None, vpc_id=None, vpc_subnet=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if instance_id and not isinstance(instance_id, int):
            raise TypeError("Expected argument 'instance_id' to be a int")
        pulumi.set(__self__, "instance_id", instance_id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if owner_id and not isinstance(owner_id, str):
            raise TypeError("Expected argument 'owner_id' to be a str")
        pulumi.set(__self__, "owner_id", owner_id)
        if security_group_id and not isinstance(security_group_id, str):
            raise TypeError("Expected argument 'security_group_id' to be a str")
        pulumi.set(__self__, "security_group_id", security_group_id)
        if vpc_id and not isinstance(vpc_id, str):
            raise TypeError("Expected argument 'vpc_id' to be a str")
        pulumi.set(__self__, "vpc_id", vpc_id)
        if vpc_subnet and not isinstance(vpc_subnet, str):
            raise TypeError("Expected argument 'vpc_subnet' to be a str")
        pulumi.set(__self__, "vpc_subnet", vpc_subnet)

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> Optional[int]:
        return pulumi.get(self, "instance_id")

    @property
    @pulumi.getter
    def name(self) -> str:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="ownerId")
    def owner_id(self) -> str:
        return pulumi.get(self, "owner_id")

    @property
    @pulumi.getter(name="securityGroupId")
    def security_group_id(self) -> str:
        return pulumi.get(self, "security_group_id")

    @property
    @pulumi.getter(name="vpcId")
    def vpc_id(self) -> Optional[str]:
        return pulumi.get(self, "vpc_id")

    @property
    @pulumi.getter(name="vpcSubnet")
    def vpc_subnet(self) -> str:
        return pulumi.get(self, "vpc_subnet")


class AwaitableGetVpcInfoResult(GetVpcInfoResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetVpcInfoResult(
            id=self.id,
            instance_id=self.instance_id,
            name=self.name,
            owner_id=self.owner_id,
            security_group_id=self.security_group_id,
            vpc_id=self.vpc_id,
            vpc_subnet=self.vpc_subnet)


def get_vpc_info(instance_id: Optional[int] = None,
                 vpc_id: Optional[str] = None,
                 opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetVpcInfoResult:
    """
    Use this data source to retrieve information about VPC for a CloudAMQP instance.

    Only available for CloudAMQP instances hosted in AWS.

    ## Example Usage

    <details>
      <summary>
        <b>
          <i>AWS VPC peering pre v1.16.0</i>
        </b>
      </summary>

    ```python
    import pulumi
    import pulumi_cloudamqp as cloudamqp

    vpc_info = cloudamqp.get_vpc_info(instance_id=instance["id"])
    ```
    </details>

    <details>
      <summary>
        <b>
          <i>AWS VPC peering post v1.16.0 (Managed VPC)</i>
        </b>
      </summary>

    ```python
    import pulumi
    import pulumi_cloudamqp as cloudamqp

    vpc_info = cloudamqp.get_vpc_info(vpc_id=vpc["id"])
    ```
    </details>

    ## Attributes reference

    All attributes reference are computed

    * `id`                  - The identifier for this resource.
    * `name`                - The name of the CloudAMQP instance.
    * `vpc_subnet`          - Dedicated VPC subnet.
    * `owner_id`            - AWS account identifier.
    * `security_group_id`   - AWS security group identifier.

    ## Dependency

    *Pre v1.16.0*
    This resource depends on CloudAMQP instance identifier, `cloudamqp_instance.instance.id`.

    *Post v1.16.0*
    This resource depends on CloudAMQP managed VPC identifier, `cloudamqp_vpc.vpc.id` or instance identifier, `cloudamqp_instance.instance.id`.


    :param int instance_id: The CloudAMQP instance identifier.
           
           ***Deprecated: Changed from required to optional in v1.16.0 will be removed in next major version (v2.0)***
    :param str vpc_id: The managed VPC identifier.
           
           ***Note: Added as optional in version v1.16.0 and will be required in next major version (v2.0)***
    """
    __args__ = dict()
    __args__['instanceId'] = instance_id
    __args__['vpcId'] = vpc_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('cloudamqp:index/getVpcInfo:getVpcInfo', __args__, opts=opts, typ=GetVpcInfoResult).value

    return AwaitableGetVpcInfoResult(
        id=pulumi.get(__ret__, 'id'),
        instance_id=pulumi.get(__ret__, 'instance_id'),
        name=pulumi.get(__ret__, 'name'),
        owner_id=pulumi.get(__ret__, 'owner_id'),
        security_group_id=pulumi.get(__ret__, 'security_group_id'),
        vpc_id=pulumi.get(__ret__, 'vpc_id'),
        vpc_subnet=pulumi.get(__ret__, 'vpc_subnet'))


@_utilities.lift_output_func(get_vpc_info)
def get_vpc_info_output(instance_id: Optional[pulumi.Input[Optional[int]]] = None,
                        vpc_id: Optional[pulumi.Input[Optional[str]]] = None,
                        opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetVpcInfoResult]:
    """
    Use this data source to retrieve information about VPC for a CloudAMQP instance.

    Only available for CloudAMQP instances hosted in AWS.

    ## Example Usage

    <details>
      <summary>
        <b>
          <i>AWS VPC peering pre v1.16.0</i>
        </b>
      </summary>

    ```python
    import pulumi
    import pulumi_cloudamqp as cloudamqp

    vpc_info = cloudamqp.get_vpc_info(instance_id=instance["id"])
    ```
    </details>

    <details>
      <summary>
        <b>
          <i>AWS VPC peering post v1.16.0 (Managed VPC)</i>
        </b>
      </summary>

    ```python
    import pulumi
    import pulumi_cloudamqp as cloudamqp

    vpc_info = cloudamqp.get_vpc_info(vpc_id=vpc["id"])
    ```
    </details>

    ## Attributes reference

    All attributes reference are computed

    * `id`                  - The identifier for this resource.
    * `name`                - The name of the CloudAMQP instance.
    * `vpc_subnet`          - Dedicated VPC subnet.
    * `owner_id`            - AWS account identifier.
    * `security_group_id`   - AWS security group identifier.

    ## Dependency

    *Pre v1.16.0*
    This resource depends on CloudAMQP instance identifier, `cloudamqp_instance.instance.id`.

    *Post v1.16.0*
    This resource depends on CloudAMQP managed VPC identifier, `cloudamqp_vpc.vpc.id` or instance identifier, `cloudamqp_instance.instance.id`.


    :param int instance_id: The CloudAMQP instance identifier.
           
           ***Deprecated: Changed from required to optional in v1.16.0 will be removed in next major version (v2.0)***
    :param str vpc_id: The managed VPC identifier.
           
           ***Note: Added as optional in version v1.16.0 and will be required in next major version (v2.0)***
    """
    ...
