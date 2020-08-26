# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from . import _utilities, _tables
from . import outputs
from ._inputs import *

__all__ = [
    'GetNodesResult',
    'AwaitableGetNodesResult',
    'get_nodes',
]

@pulumi.output_type
class GetNodesResult:
    """
    A collection of values returned by getNodes.
    """
    def __init__(__self__, id=None, instance_id=None, nodes=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if instance_id and not isinstance(instance_id, float):
            raise TypeError("Expected argument 'instance_id' to be a float")
        pulumi.set(__self__, "instance_id", instance_id)
        if nodes and not isinstance(nodes, list):
            raise TypeError("Expected argument 'nodes' to be a list")
        pulumi.set(__self__, "nodes", nodes)

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> float:
        return pulumi.get(self, "instance_id")

    @property
    @pulumi.getter
    def nodes(self) -> List['outputs.GetNodesNodeResult']:
        return pulumi.get(self, "nodes")


class AwaitableGetNodesResult(GetNodesResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetNodesResult(
            id=self.id,
            instance_id=self.instance_id,
            nodes=self.nodes)


def get_nodes(instance_id: Optional[float] = None,
              nodes: Optional[List[pulumi.InputType['GetNodesNodeArgs']]] = None,
              opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetNodesResult:
    """
    Use this data source to retrieve information about the node(s) created by CloudAMQP instance.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_cloudamqp as cloudamqp

    nodes = cloudamqp.get_nodes(instance_id=cloudamqp_instance["instance"]["id"])
    ```
    ## Argument reference

    * `instance_id` - (Required) The CloudAMQP instance identifier.

    ## Attribute reference

    * `nodes` - (Computed) An array of node information. Each `nodes` block consists of the fields documented below.

    ***

    The `nodes` block consist of

    * `hostname`          - (Computed) Hostname assigned to the node.
    * `name`              - (Computed) Name of the node.
    * `running`           - (Computed) Is the node running?
    * `rabbitmq_version`  - (Computed) Currently configured Rabbit MQ version on the node.
    * `erlang_version`    - (Computed) Currently used Erlanbg version on the node.
    * `hipe`              - (Computed) Enable or disable High-performance Erlang.

    ## Dependency

    This data source depends on CloudAMQP instance identifier, `cloudamqp_instance.instance.id`.
    """
    __args__ = dict()
    __args__['instanceId'] = instance_id
    __args__['nodes'] = nodes
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('cloudamqp:index/getNodes:getNodes', __args__, opts=opts, typ=GetNodesResult).value

    return AwaitableGetNodesResult(
        id=__ret__.id,
        instance_id=__ret__.instance_id,
        nodes=__ret__.nodes)
