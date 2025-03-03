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
    'GetVpcGcpInfoResult',
    'AwaitableGetVpcGcpInfoResult',
    'get_vpc_gcp_info',
    'get_vpc_gcp_info_output',
]

@pulumi.output_type
class GetVpcGcpInfoResult:
    """
    A collection of values returned by getVpcGcpInfo.
    """
    def __init__(__self__, id=None, instance_id=None, name=None, network=None, sleep=None, timeout=None, vpc_id=None, vpc_subnet=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if instance_id and not isinstance(instance_id, int):
            raise TypeError("Expected argument 'instance_id' to be a int")
        pulumi.set(__self__, "instance_id", instance_id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if network and not isinstance(network, str):
            raise TypeError("Expected argument 'network' to be a str")
        pulumi.set(__self__, "network", network)
        if sleep and not isinstance(sleep, int):
            raise TypeError("Expected argument 'sleep' to be a int")
        pulumi.set(__self__, "sleep", sleep)
        if timeout and not isinstance(timeout, int):
            raise TypeError("Expected argument 'timeout' to be a int")
        pulumi.set(__self__, "timeout", timeout)
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
    @pulumi.getter
    def network(self) -> str:
        return pulumi.get(self, "network")

    @property
    @pulumi.getter
    def sleep(self) -> Optional[int]:
        return pulumi.get(self, "sleep")

    @property
    @pulumi.getter
    def timeout(self) -> Optional[int]:
        return pulumi.get(self, "timeout")

    @property
    @pulumi.getter(name="vpcId")
    def vpc_id(self) -> Optional[str]:
        return pulumi.get(self, "vpc_id")

    @property
    @pulumi.getter(name="vpcSubnet")
    def vpc_subnet(self) -> str:
        return pulumi.get(self, "vpc_subnet")


class AwaitableGetVpcGcpInfoResult(GetVpcGcpInfoResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetVpcGcpInfoResult(
            id=self.id,
            instance_id=self.instance_id,
            name=self.name,
            network=self.network,
            sleep=self.sleep,
            timeout=self.timeout,
            vpc_id=self.vpc_id,
            vpc_subnet=self.vpc_subnet)


def get_vpc_gcp_info(instance_id: Optional[int] = None,
                     sleep: Optional[int] = None,
                     timeout: Optional[int] = None,
                     vpc_id: Optional[str] = None,
                     opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetVpcGcpInfoResult:
    """
    Use this data source to retrieve information about VPC for a CloudAMQP instance.

    > **Note:** Only available for CloudAMQP instances/VPCs hosted in Google Cloud Platform (GCP).

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

    vpc_info = cloudamqp.get_vpc_gcp_info(instance_id=instance["id"])
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

    vpc_info = cloudamqp.get_vpc_gcp_info(vpc_id=vpc["id"])
    ```

    </details>

    ## Attributes reference

    All attributes reference are computed

    * `id`                  - The identifier for this resource.
    * `name`                - The name of the VPC.
    * `vpc_subnet`          - Dedicated VPC subnet.
    * `network`             - VPC network uri.

    ## Dependency

    *Pre v1.16.0*
    This resource depends on CloudAMQP instance identifier, `cloudamqp_instance.instance.id`.

    *Post v1.16.0*
    This resource depends on CloudAMQP managed VPC identifier, `cloudamqp_vpc.vpc.id` or instance identifier, `cloudamqp_instance.instance.id`.


    :param int instance_id: The CloudAMQP instance identifier.
           
           ***Deprecated: Changed from required to optional in v1.16.0 will be removed in next major version (v2.0)***
    :param int sleep: Configurable sleep time (seconds) between retries when reading peering. Default set to 10 seconds.
    :param int timeout: Configurable timeout time (seconds) before retries times out. Default set to 1800 seconds.
    :param str vpc_id: The managed VPC identifier.
           
           ***Note: Added as optional in version v1.16.0 and will be required in next major version (v2.0)***
    """
    __args__ = dict()
    __args__['instanceId'] = instance_id
    __args__['sleep'] = sleep
    __args__['timeout'] = timeout
    __args__['vpcId'] = vpc_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('cloudamqp:index/getVpcGcpInfo:getVpcGcpInfo', __args__, opts=opts, typ=GetVpcGcpInfoResult).value

    return AwaitableGetVpcGcpInfoResult(
        id=pulumi.get(__ret__, 'id'),
        instance_id=pulumi.get(__ret__, 'instance_id'),
        name=pulumi.get(__ret__, 'name'),
        network=pulumi.get(__ret__, 'network'),
        sleep=pulumi.get(__ret__, 'sleep'),
        timeout=pulumi.get(__ret__, 'timeout'),
        vpc_id=pulumi.get(__ret__, 'vpc_id'),
        vpc_subnet=pulumi.get(__ret__, 'vpc_subnet'))
def get_vpc_gcp_info_output(instance_id: Optional[pulumi.Input[Optional[int]]] = None,
                            sleep: Optional[pulumi.Input[Optional[int]]] = None,
                            timeout: Optional[pulumi.Input[Optional[int]]] = None,
                            vpc_id: Optional[pulumi.Input[Optional[str]]] = None,
                            opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetVpcGcpInfoResult]:
    """
    Use this data source to retrieve information about VPC for a CloudAMQP instance.

    > **Note:** Only available for CloudAMQP instances/VPCs hosted in Google Cloud Platform (GCP).

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

    vpc_info = cloudamqp.get_vpc_gcp_info(instance_id=instance["id"])
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

    vpc_info = cloudamqp.get_vpc_gcp_info(vpc_id=vpc["id"])
    ```

    </details>

    ## Attributes reference

    All attributes reference are computed

    * `id`                  - The identifier for this resource.
    * `name`                - The name of the VPC.
    * `vpc_subnet`          - Dedicated VPC subnet.
    * `network`             - VPC network uri.

    ## Dependency

    *Pre v1.16.0*
    This resource depends on CloudAMQP instance identifier, `cloudamqp_instance.instance.id`.

    *Post v1.16.0*
    This resource depends on CloudAMQP managed VPC identifier, `cloudamqp_vpc.vpc.id` or instance identifier, `cloudamqp_instance.instance.id`.


    :param int instance_id: The CloudAMQP instance identifier.
           
           ***Deprecated: Changed from required to optional in v1.16.0 will be removed in next major version (v2.0)***
    :param int sleep: Configurable sleep time (seconds) between retries when reading peering. Default set to 10 seconds.
    :param int timeout: Configurable timeout time (seconds) before retries times out. Default set to 1800 seconds.
    :param str vpc_id: The managed VPC identifier.
           
           ***Note: Added as optional in version v1.16.0 and will be required in next major version (v2.0)***
    """
    __args__ = dict()
    __args__['instanceId'] = instance_id
    __args__['sleep'] = sleep
    __args__['timeout'] = timeout
    __args__['vpcId'] = vpc_id
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('cloudamqp:index/getVpcGcpInfo:getVpcGcpInfo', __args__, opts=opts, typ=GetVpcGcpInfoResult)
    return __ret__.apply(lambda __response__: GetVpcGcpInfoResult(
        id=pulumi.get(__response__, 'id'),
        instance_id=pulumi.get(__response__, 'instance_id'),
        name=pulumi.get(__response__, 'name'),
        network=pulumi.get(__response__, 'network'),
        sleep=pulumi.get(__response__, 'sleep'),
        timeout=pulumi.get(__response__, 'timeout'),
        vpc_id=pulumi.get(__response__, 'vpc_id'),
        vpc_subnet=pulumi.get(__response__, 'vpc_subnet')))
