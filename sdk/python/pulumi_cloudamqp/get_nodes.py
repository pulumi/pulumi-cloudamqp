# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from . import utilities, tables

class GetNodesResult:
    """
    A collection of values returned by getNodes.
    """
    def __init__(__self__, id=None, instance_id=None, nodes=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        __self__.id = id
        """
        The provider-assigned unique ID for this managed resource.
        """
        if instance_id and not isinstance(instance_id, float):
            raise TypeError("Expected argument 'instance_id' to be a float")
        __self__.instance_id = instance_id
        if nodes and not isinstance(nodes, list):
            raise TypeError("Expected argument 'nodes' to be a list")
        __self__.nodes = nodes
class AwaitableGetNodesResult(GetNodesResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetNodesResult(
            id=self.id,
            instance_id=self.instance_id,
            nodes=self.nodes)

def get_nodes(instance_id=None,nodes=None,opts=None):
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



    The **nodes** object supports the following:

      * `erlangVersion` (`str`)
      * `hipe` (`bool`)
      * `hostname` (`str`)
      * `name` (`str`)
      * `rabbitmqVersion` (`str`)
      * `running` (`bool`)
    """
    __args__ = dict()


    __args__['instanceId'] = instance_id
    __args__['nodes'] = nodes
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = utilities.get_version()
    __ret__ = pulumi.runtime.invoke('cloudamqp:index/getNodes:getNodes', __args__, opts=opts).value

    return AwaitableGetNodesResult(
        id=__ret__.get('id'),
        instance_id=__ret__.get('instanceId'),
        nodes=__ret__.get('nodes'))
