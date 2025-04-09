# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
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

__all__ = ['NodeActionsArgs', 'NodeActions']

@pulumi.input_type
class NodeActionsArgs:
    def __init__(__self__, *,
                 action: pulumi.Input[builtins.str],
                 instance_id: pulumi.Input[builtins.int],
                 node_name: pulumi.Input[builtins.str]):
        """
        The set of arguments for constructing a NodeActions resource.
        :param pulumi.Input[builtins.str] action: The action to invoke on the node.
        :param pulumi.Input[builtins.int] instance_id: The CloudAMQP instance ID.
        :param pulumi.Input[builtins.str] node_name: The node name, e.g `green-guinea-pig-01`.
        """
        pulumi.set(__self__, "action", action)
        pulumi.set(__self__, "instance_id", instance_id)
        pulumi.set(__self__, "node_name", node_name)

    @property
    @pulumi.getter
    def action(self) -> pulumi.Input[builtins.str]:
        """
        The action to invoke on the node.
        """
        return pulumi.get(self, "action")

    @action.setter
    def action(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "action", value)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> pulumi.Input[builtins.int]:
        """
        The CloudAMQP instance ID.
        """
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: pulumi.Input[builtins.int]):
        pulumi.set(self, "instance_id", value)

    @property
    @pulumi.getter(name="nodeName")
    def node_name(self) -> pulumi.Input[builtins.str]:
        """
        The node name, e.g `green-guinea-pig-01`.
        """
        return pulumi.get(self, "node_name")

    @node_name.setter
    def node_name(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "node_name", value)


@pulumi.input_type
class _NodeActionsState:
    def __init__(__self__, *,
                 action: Optional[pulumi.Input[builtins.str]] = None,
                 instance_id: Optional[pulumi.Input[builtins.int]] = None,
                 node_name: Optional[pulumi.Input[builtins.str]] = None,
                 running: Optional[pulumi.Input[builtins.bool]] = None):
        """
        Input properties used for looking up and filtering NodeActions resources.
        :param pulumi.Input[builtins.str] action: The action to invoke on the node.
        :param pulumi.Input[builtins.int] instance_id: The CloudAMQP instance ID.
        :param pulumi.Input[builtins.str] node_name: The node name, e.g `green-guinea-pig-01`.
        :param pulumi.Input[builtins.bool] running: If the node is running.
        """
        if action is not None:
            pulumi.set(__self__, "action", action)
        if instance_id is not None:
            pulumi.set(__self__, "instance_id", instance_id)
        if node_name is not None:
            pulumi.set(__self__, "node_name", node_name)
        if running is not None:
            pulumi.set(__self__, "running", running)

    @property
    @pulumi.getter
    def action(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The action to invoke on the node.
        """
        return pulumi.get(self, "action")

    @action.setter
    def action(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "action", value)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> Optional[pulumi.Input[builtins.int]]:
        """
        The CloudAMQP instance ID.
        """
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: Optional[pulumi.Input[builtins.int]]):
        pulumi.set(self, "instance_id", value)

    @property
    @pulumi.getter(name="nodeName")
    def node_name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The node name, e.g `green-guinea-pig-01`.
        """
        return pulumi.get(self, "node_name")

    @node_name.setter
    def node_name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "node_name", value)

    @property
    @pulumi.getter
    def running(self) -> Optional[pulumi.Input[builtins.bool]]:
        """
        If the node is running.
        """
        return pulumi.get(self, "running")

    @running.setter
    def running(self, value: Optional[pulumi.Input[builtins.bool]]):
        pulumi.set(self, "running", value)


class NodeActions(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 action: Optional[pulumi.Input[builtins.str]] = None,
                 instance_id: Optional[pulumi.Input[builtins.int]] = None,
                 node_name: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        """
        This resource allows you to invoke actions on a specific node.

        Only available for dedicated subscription plans.

        ## Example Usage

        <details>
          <summary>
            <b>
              <i>Already know the node identifier (e.g. from state file)</i>
            </b>
          </summary>

        ```python
        import pulumi
        import pulumi_cloudamqp as cloudamqp

        # New recipient to receieve notifications
        node_action = cloudamqp.NodeActions("node_action",
            instance_id=instance["id"],
            node_name="<node name>",
            action="restart")
        ```

        </details>

        <details>
          <summary>
            <b>
              <i>Multi node RabbitMQ restart</i>
            </b>
          </summary>

        Using data source `get_nodes` to restart RabbitMQ on all nodes.

        > **Note:** RabbitMQ restart on multiple nodes need to be chained, let one node restart at the time.

        ```python
        import pulumi
        import pulumi_cloudamqp as cloudamqp

        list_nodes = cloudamqp.get_nodes(instance_id=instance["id"])
        restart01 = cloudamqp.NodeActions("restart_01",
            instance_id=instance["id"],
            action="restart",
            node_name=list_nodes.nodes[0].name)
        restart02 = cloudamqp.NodeActions("restart_02",
            instance_id=instance["id"],
            action="restart",
            node_name=list_nodes.nodes[1].name,
            opts = pulumi.ResourceOptions(depends_on=[restart01]))
        restart03 = cloudamqp.NodeActions("restart_03",
            instance_id=instance["id"],
            action="restart",
            node_name=list_nodes.nodes[2].name,
            opts = pulumi.ResourceOptions(depends_on=[
                    restart01,
                    restart02,
                ]))
        ```

        </details>

        <details>
          <summary>
            <b>
              <i>Combine log level configuration change with multi node RabbitMQ restart</i>
            </b>
          </summary>

        ```python
        import pulumi
        import pulumi_cloudamqp as cloudamqp

        list_nodes = cloudamqp.get_nodes(instance_id=instance["id"])
        rabbitmq_config = cloudamqp.RabbitConfiguration("rabbitmq_config",
            instance_id=instance["id"],
            log_exchange_level="info")
        restart01 = cloudamqp.NodeActions("restart_01",
            instance_id=instance["id"],
            action="restart",
            node_name=list_nodes.nodes[0].name,
            opts = pulumi.ResourceOptions(depends_on=[rabbitmq_config]))
        restart02 = cloudamqp.NodeActions("restart_02",
            instance_id=instance["id"],
            action="restart",
            node_name=list_nodes.nodes[1].name,
            opts = pulumi.ResourceOptions(depends_on=[
                    rabbitmq_config,
                    restart01,
                ]))
        restart03 = cloudamqp.NodeActions("restart_03",
            instance_id=instance["id"],
            action="restart",
            node_name=list_nodes.nodes[2].name,
            opts = pulumi.ResourceOptions(depends_on=[
                    rabbitmq_config,
                    restart01,
                    restart02,
                ]))
        ```

        </details>

        ## Action reference

        Valid actions for ***LavinMQ***.

        | Action       | Info                               |
        |--------------|------------------------------------|
        | start        | Start LavinMQ                      |
        | stop         | Stop LavinMQ                       |
        | restart      | Restart LavinMQ                    |
        | reboot       | Reboot the node                    |

        Valid actions for ***RabbitMQ***.

        | Action       | Info                               |
        |--------------|------------------------------------|
        | start        | Start RabbitMQ                     |
        | stop         | Stop RabbitMQ                      |
        | restart      | Restart RabbitMQ                   |
        | reboot       | Reboot the node                    |
        | mgmt.restart | Restart the RabbitMQ mgmt interace |

        ## Dependency

        This resource depends on CloudAMQP instance identifier, `cloudamqp_instance.instance.id` and node
        name.

        ## Import

        This resource cannot be imported.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.str] action: The action to invoke on the node.
        :param pulumi.Input[builtins.int] instance_id: The CloudAMQP instance ID.
        :param pulumi.Input[builtins.str] node_name: The node name, e.g `green-guinea-pig-01`.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: NodeActionsArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        This resource allows you to invoke actions on a specific node.

        Only available for dedicated subscription plans.

        ## Example Usage

        <details>
          <summary>
            <b>
              <i>Already know the node identifier (e.g. from state file)</i>
            </b>
          </summary>

        ```python
        import pulumi
        import pulumi_cloudamqp as cloudamqp

        # New recipient to receieve notifications
        node_action = cloudamqp.NodeActions("node_action",
            instance_id=instance["id"],
            node_name="<node name>",
            action="restart")
        ```

        </details>

        <details>
          <summary>
            <b>
              <i>Multi node RabbitMQ restart</i>
            </b>
          </summary>

        Using data source `get_nodes` to restart RabbitMQ on all nodes.

        > **Note:** RabbitMQ restart on multiple nodes need to be chained, let one node restart at the time.

        ```python
        import pulumi
        import pulumi_cloudamqp as cloudamqp

        list_nodes = cloudamqp.get_nodes(instance_id=instance["id"])
        restart01 = cloudamqp.NodeActions("restart_01",
            instance_id=instance["id"],
            action="restart",
            node_name=list_nodes.nodes[0].name)
        restart02 = cloudamqp.NodeActions("restart_02",
            instance_id=instance["id"],
            action="restart",
            node_name=list_nodes.nodes[1].name,
            opts = pulumi.ResourceOptions(depends_on=[restart01]))
        restart03 = cloudamqp.NodeActions("restart_03",
            instance_id=instance["id"],
            action="restart",
            node_name=list_nodes.nodes[2].name,
            opts = pulumi.ResourceOptions(depends_on=[
                    restart01,
                    restart02,
                ]))
        ```

        </details>

        <details>
          <summary>
            <b>
              <i>Combine log level configuration change with multi node RabbitMQ restart</i>
            </b>
          </summary>

        ```python
        import pulumi
        import pulumi_cloudamqp as cloudamqp

        list_nodes = cloudamqp.get_nodes(instance_id=instance["id"])
        rabbitmq_config = cloudamqp.RabbitConfiguration("rabbitmq_config",
            instance_id=instance["id"],
            log_exchange_level="info")
        restart01 = cloudamqp.NodeActions("restart_01",
            instance_id=instance["id"],
            action="restart",
            node_name=list_nodes.nodes[0].name,
            opts = pulumi.ResourceOptions(depends_on=[rabbitmq_config]))
        restart02 = cloudamqp.NodeActions("restart_02",
            instance_id=instance["id"],
            action="restart",
            node_name=list_nodes.nodes[1].name,
            opts = pulumi.ResourceOptions(depends_on=[
                    rabbitmq_config,
                    restart01,
                ]))
        restart03 = cloudamqp.NodeActions("restart_03",
            instance_id=instance["id"],
            action="restart",
            node_name=list_nodes.nodes[2].name,
            opts = pulumi.ResourceOptions(depends_on=[
                    rabbitmq_config,
                    restart01,
                    restart02,
                ]))
        ```

        </details>

        ## Action reference

        Valid actions for ***LavinMQ***.

        | Action       | Info                               |
        |--------------|------------------------------------|
        | start        | Start LavinMQ                      |
        | stop         | Stop LavinMQ                       |
        | restart      | Restart LavinMQ                    |
        | reboot       | Reboot the node                    |

        Valid actions for ***RabbitMQ***.

        | Action       | Info                               |
        |--------------|------------------------------------|
        | start        | Start RabbitMQ                     |
        | stop         | Stop RabbitMQ                      |
        | restart      | Restart RabbitMQ                   |
        | reboot       | Reboot the node                    |
        | mgmt.restart | Restart the RabbitMQ mgmt interace |

        ## Dependency

        This resource depends on CloudAMQP instance identifier, `cloudamqp_instance.instance.id` and node
        name.

        ## Import

        This resource cannot be imported.

        :param str resource_name: The name of the resource.
        :param NodeActionsArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(NodeActionsArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 action: Optional[pulumi.Input[builtins.str]] = None,
                 instance_id: Optional[pulumi.Input[builtins.int]] = None,
                 node_name: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = NodeActionsArgs.__new__(NodeActionsArgs)

            if action is None and not opts.urn:
                raise TypeError("Missing required property 'action'")
            __props__.__dict__["action"] = action
            if instance_id is None and not opts.urn:
                raise TypeError("Missing required property 'instance_id'")
            __props__.__dict__["instance_id"] = instance_id
            if node_name is None and not opts.urn:
                raise TypeError("Missing required property 'node_name'")
            __props__.__dict__["node_name"] = node_name
            __props__.__dict__["running"] = None
        super(NodeActions, __self__).__init__(
            'cloudamqp:index/nodeActions:NodeActions',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            action: Optional[pulumi.Input[builtins.str]] = None,
            instance_id: Optional[pulumi.Input[builtins.int]] = None,
            node_name: Optional[pulumi.Input[builtins.str]] = None,
            running: Optional[pulumi.Input[builtins.bool]] = None) -> 'NodeActions':
        """
        Get an existing NodeActions resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.str] action: The action to invoke on the node.
        :param pulumi.Input[builtins.int] instance_id: The CloudAMQP instance ID.
        :param pulumi.Input[builtins.str] node_name: The node name, e.g `green-guinea-pig-01`.
        :param pulumi.Input[builtins.bool] running: If the node is running.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _NodeActionsState.__new__(_NodeActionsState)

        __props__.__dict__["action"] = action
        __props__.__dict__["instance_id"] = instance_id
        __props__.__dict__["node_name"] = node_name
        __props__.__dict__["running"] = running
        return NodeActions(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def action(self) -> pulumi.Output[builtins.str]:
        """
        The action to invoke on the node.
        """
        return pulumi.get(self, "action")

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> pulumi.Output[builtins.int]:
        """
        The CloudAMQP instance ID.
        """
        return pulumi.get(self, "instance_id")

    @property
    @pulumi.getter(name="nodeName")
    def node_name(self) -> pulumi.Output[builtins.str]:
        """
        The node name, e.g `green-guinea-pig-01`.
        """
        return pulumi.get(self, "node_name")

    @property
    @pulumi.getter
    def running(self) -> pulumi.Output[builtins.bool]:
        """
        If the node is running.
        """
        return pulumi.get(self, "running")

