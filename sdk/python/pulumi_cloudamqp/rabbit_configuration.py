# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['RabbitConfigurationArgs', 'RabbitConfiguration']

@pulumi.input_type
class RabbitConfigurationArgs:
    def __init__(__self__, *,
                 instance_id: pulumi.Input[int],
                 channel_max: Optional[pulumi.Input[int]] = None,
                 connection_max: Optional[pulumi.Input[int]] = None,
                 consumer_timeout: Optional[pulumi.Input[int]] = None,
                 heartbeat: Optional[pulumi.Input[int]] = None,
                 log_exchange_level: Optional[pulumi.Input[str]] = None,
                 max_message_size: Optional[pulumi.Input[int]] = None,
                 queue_index_embed_msgs_below: Optional[pulumi.Input[int]] = None,
                 vm_memory_high_watermark: Optional[pulumi.Input[float]] = None):
        """
        The set of arguments for constructing a RabbitConfiguration resource.
        :param pulumi.Input[int] instance_id: The CloudAMQP instance ID.
        :param pulumi.Input[int] channel_max: Set the maximum permissible number of channels per connection.
        :param pulumi.Input[int] connection_max: Set the maximum permissible number of connection.
        :param pulumi.Input[int] consumer_timeout: A consumer that has recevied a message and does not acknowledge that message within the timeout in milliseconds
        :param pulumi.Input[int] heartbeat: Set the server AMQP 0-9-1 heartbeat timeout in seconds.
        :param pulumi.Input[str] log_exchange_level: Log level for the logger used for log integrations and the CloudAMQP Console log view.
        :param pulumi.Input[int] max_message_size: The largest allowed message payload size in bytes.
        :param pulumi.Input[int] queue_index_embed_msgs_below: Size in bytes below which to embed messages in the queue index.
        :param pulumi.Input[float] vm_memory_high_watermark: When the server will enter memory based flow-control as relative to the maximum available memory.
        """
        pulumi.set(__self__, "instance_id", instance_id)
        if channel_max is not None:
            pulumi.set(__self__, "channel_max", channel_max)
        if connection_max is not None:
            pulumi.set(__self__, "connection_max", connection_max)
        if consumer_timeout is not None:
            pulumi.set(__self__, "consumer_timeout", consumer_timeout)
        if heartbeat is not None:
            pulumi.set(__self__, "heartbeat", heartbeat)
        if log_exchange_level is not None:
            pulumi.set(__self__, "log_exchange_level", log_exchange_level)
        if max_message_size is not None:
            pulumi.set(__self__, "max_message_size", max_message_size)
        if queue_index_embed_msgs_below is not None:
            pulumi.set(__self__, "queue_index_embed_msgs_below", queue_index_embed_msgs_below)
        if vm_memory_high_watermark is not None:
            pulumi.set(__self__, "vm_memory_high_watermark", vm_memory_high_watermark)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> pulumi.Input[int]:
        """
        The CloudAMQP instance ID.
        """
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: pulumi.Input[int]):
        pulumi.set(self, "instance_id", value)

    @property
    @pulumi.getter(name="channelMax")
    def channel_max(self) -> Optional[pulumi.Input[int]]:
        """
        Set the maximum permissible number of channels per connection.
        """
        return pulumi.get(self, "channel_max")

    @channel_max.setter
    def channel_max(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "channel_max", value)

    @property
    @pulumi.getter(name="connectionMax")
    def connection_max(self) -> Optional[pulumi.Input[int]]:
        """
        Set the maximum permissible number of connection.
        """
        return pulumi.get(self, "connection_max")

    @connection_max.setter
    def connection_max(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "connection_max", value)

    @property
    @pulumi.getter(name="consumerTimeout")
    def consumer_timeout(self) -> Optional[pulumi.Input[int]]:
        """
        A consumer that has recevied a message and does not acknowledge that message within the timeout in milliseconds
        """
        return pulumi.get(self, "consumer_timeout")

    @consumer_timeout.setter
    def consumer_timeout(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "consumer_timeout", value)

    @property
    @pulumi.getter
    def heartbeat(self) -> Optional[pulumi.Input[int]]:
        """
        Set the server AMQP 0-9-1 heartbeat timeout in seconds.
        """
        return pulumi.get(self, "heartbeat")

    @heartbeat.setter
    def heartbeat(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "heartbeat", value)

    @property
    @pulumi.getter(name="logExchangeLevel")
    def log_exchange_level(self) -> Optional[pulumi.Input[str]]:
        """
        Log level for the logger used for log integrations and the CloudAMQP Console log view.
        """
        return pulumi.get(self, "log_exchange_level")

    @log_exchange_level.setter
    def log_exchange_level(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "log_exchange_level", value)

    @property
    @pulumi.getter(name="maxMessageSize")
    def max_message_size(self) -> Optional[pulumi.Input[int]]:
        """
        The largest allowed message payload size in bytes.
        """
        return pulumi.get(self, "max_message_size")

    @max_message_size.setter
    def max_message_size(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "max_message_size", value)

    @property
    @pulumi.getter(name="queueIndexEmbedMsgsBelow")
    def queue_index_embed_msgs_below(self) -> Optional[pulumi.Input[int]]:
        """
        Size in bytes below which to embed messages in the queue index.
        """
        return pulumi.get(self, "queue_index_embed_msgs_below")

    @queue_index_embed_msgs_below.setter
    def queue_index_embed_msgs_below(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "queue_index_embed_msgs_below", value)

    @property
    @pulumi.getter(name="vmMemoryHighWatermark")
    def vm_memory_high_watermark(self) -> Optional[pulumi.Input[float]]:
        """
        When the server will enter memory based flow-control as relative to the maximum available memory.
        """
        return pulumi.get(self, "vm_memory_high_watermark")

    @vm_memory_high_watermark.setter
    def vm_memory_high_watermark(self, value: Optional[pulumi.Input[float]]):
        pulumi.set(self, "vm_memory_high_watermark", value)


@pulumi.input_type
class _RabbitConfigurationState:
    def __init__(__self__, *,
                 channel_max: Optional[pulumi.Input[int]] = None,
                 connection_max: Optional[pulumi.Input[int]] = None,
                 consumer_timeout: Optional[pulumi.Input[int]] = None,
                 heartbeat: Optional[pulumi.Input[int]] = None,
                 instance_id: Optional[pulumi.Input[int]] = None,
                 log_exchange_level: Optional[pulumi.Input[str]] = None,
                 max_message_size: Optional[pulumi.Input[int]] = None,
                 queue_index_embed_msgs_below: Optional[pulumi.Input[int]] = None,
                 vm_memory_high_watermark: Optional[pulumi.Input[float]] = None):
        """
        Input properties used for looking up and filtering RabbitConfiguration resources.
        :param pulumi.Input[int] channel_max: Set the maximum permissible number of channels per connection.
        :param pulumi.Input[int] connection_max: Set the maximum permissible number of connection.
        :param pulumi.Input[int] consumer_timeout: A consumer that has recevied a message and does not acknowledge that message within the timeout in milliseconds
        :param pulumi.Input[int] heartbeat: Set the server AMQP 0-9-1 heartbeat timeout in seconds.
        :param pulumi.Input[int] instance_id: The CloudAMQP instance ID.
        :param pulumi.Input[str] log_exchange_level: Log level for the logger used for log integrations and the CloudAMQP Console log view.
        :param pulumi.Input[int] max_message_size: The largest allowed message payload size in bytes.
        :param pulumi.Input[int] queue_index_embed_msgs_below: Size in bytes below which to embed messages in the queue index.
        :param pulumi.Input[float] vm_memory_high_watermark: When the server will enter memory based flow-control as relative to the maximum available memory.
        """
        if channel_max is not None:
            pulumi.set(__self__, "channel_max", channel_max)
        if connection_max is not None:
            pulumi.set(__self__, "connection_max", connection_max)
        if consumer_timeout is not None:
            pulumi.set(__self__, "consumer_timeout", consumer_timeout)
        if heartbeat is not None:
            pulumi.set(__self__, "heartbeat", heartbeat)
        if instance_id is not None:
            pulumi.set(__self__, "instance_id", instance_id)
        if log_exchange_level is not None:
            pulumi.set(__self__, "log_exchange_level", log_exchange_level)
        if max_message_size is not None:
            pulumi.set(__self__, "max_message_size", max_message_size)
        if queue_index_embed_msgs_below is not None:
            pulumi.set(__self__, "queue_index_embed_msgs_below", queue_index_embed_msgs_below)
        if vm_memory_high_watermark is not None:
            pulumi.set(__self__, "vm_memory_high_watermark", vm_memory_high_watermark)

    @property
    @pulumi.getter(name="channelMax")
    def channel_max(self) -> Optional[pulumi.Input[int]]:
        """
        Set the maximum permissible number of channels per connection.
        """
        return pulumi.get(self, "channel_max")

    @channel_max.setter
    def channel_max(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "channel_max", value)

    @property
    @pulumi.getter(name="connectionMax")
    def connection_max(self) -> Optional[pulumi.Input[int]]:
        """
        Set the maximum permissible number of connection.
        """
        return pulumi.get(self, "connection_max")

    @connection_max.setter
    def connection_max(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "connection_max", value)

    @property
    @pulumi.getter(name="consumerTimeout")
    def consumer_timeout(self) -> Optional[pulumi.Input[int]]:
        """
        A consumer that has recevied a message and does not acknowledge that message within the timeout in milliseconds
        """
        return pulumi.get(self, "consumer_timeout")

    @consumer_timeout.setter
    def consumer_timeout(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "consumer_timeout", value)

    @property
    @pulumi.getter
    def heartbeat(self) -> Optional[pulumi.Input[int]]:
        """
        Set the server AMQP 0-9-1 heartbeat timeout in seconds.
        """
        return pulumi.get(self, "heartbeat")

    @heartbeat.setter
    def heartbeat(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "heartbeat", value)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> Optional[pulumi.Input[int]]:
        """
        The CloudAMQP instance ID.
        """
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "instance_id", value)

    @property
    @pulumi.getter(name="logExchangeLevel")
    def log_exchange_level(self) -> Optional[pulumi.Input[str]]:
        """
        Log level for the logger used for log integrations and the CloudAMQP Console log view.
        """
        return pulumi.get(self, "log_exchange_level")

    @log_exchange_level.setter
    def log_exchange_level(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "log_exchange_level", value)

    @property
    @pulumi.getter(name="maxMessageSize")
    def max_message_size(self) -> Optional[pulumi.Input[int]]:
        """
        The largest allowed message payload size in bytes.
        """
        return pulumi.get(self, "max_message_size")

    @max_message_size.setter
    def max_message_size(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "max_message_size", value)

    @property
    @pulumi.getter(name="queueIndexEmbedMsgsBelow")
    def queue_index_embed_msgs_below(self) -> Optional[pulumi.Input[int]]:
        """
        Size in bytes below which to embed messages in the queue index.
        """
        return pulumi.get(self, "queue_index_embed_msgs_below")

    @queue_index_embed_msgs_below.setter
    def queue_index_embed_msgs_below(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "queue_index_embed_msgs_below", value)

    @property
    @pulumi.getter(name="vmMemoryHighWatermark")
    def vm_memory_high_watermark(self) -> Optional[pulumi.Input[float]]:
        """
        When the server will enter memory based flow-control as relative to the maximum available memory.
        """
        return pulumi.get(self, "vm_memory_high_watermark")

    @vm_memory_high_watermark.setter
    def vm_memory_high_watermark(self, value: Optional[pulumi.Input[float]]):
        pulumi.set(self, "vm_memory_high_watermark", value)


class RabbitConfiguration(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 channel_max: Optional[pulumi.Input[int]] = None,
                 connection_max: Optional[pulumi.Input[int]] = None,
                 consumer_timeout: Optional[pulumi.Input[int]] = None,
                 heartbeat: Optional[pulumi.Input[int]] = None,
                 instance_id: Optional[pulumi.Input[int]] = None,
                 log_exchange_level: Optional[pulumi.Input[str]] = None,
                 max_message_size: Optional[pulumi.Input[int]] = None,
                 queue_index_embed_msgs_below: Optional[pulumi.Input[int]] = None,
                 vm_memory_high_watermark: Optional[pulumi.Input[float]] = None,
                 __props__=None):
        """
        This resource allows you update RabbitMQ config.

        Only available for dedicated subscription plans.

        ## Argument threshold values

        | Argument | Type | Default | Min | Max | Unit | Affect | Note |
        |---|---|---|---|---|---|---|---|
        | heartbeat | int | 120 | 0 | - |  | Only effects new connections |  |
        | connection_max | int | -1 | 1 | - |  | RabbitMQ restart required | -1 in the provider corresponds to INFINITY in the RabbitMQ config |
        | channel_max | int | 128 | 0 | - |  | Only effects new connections |  |
        | consumer_timeout | int | 7200000 | 10000 | 86400000 | milliseconds | Only effects new channels | -1 in the provider corresponds to false (disable) in the RabbitMQ config |
        | vm_memory_high_watermark | float | 0.81 | 0.4 | 0.9 |  | Applied immediately |  |
        | queue_index_embed_msgs_below | int | 4096 | 1 | 10485760 | bytes | Applied immediately for new queues, requires restart for existing queues |  |
        | max_message_size | int | 134217728 | 1 | 536870912 | bytes | Only effects new channels |  |
        | log_exchange_level | string | error | - | - |  | RabbitMQ restart required | debug, info, warning, error, critical |

        ## Dependency

        This resource depends on CloudAMQP instance identifier, `cloudamqp_instance.instance.id`.

        ## Import

        `cloudamqp_rabbitmq_configuration` can be imported using the CloudAMQP instance identifier.

        ```sh
         $ pulumi import cloudamqp:index/rabbitConfiguration:RabbitConfiguration config <instance_id>`
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] channel_max: Set the maximum permissible number of channels per connection.
        :param pulumi.Input[int] connection_max: Set the maximum permissible number of connection.
        :param pulumi.Input[int] consumer_timeout: A consumer that has recevied a message and does not acknowledge that message within the timeout in milliseconds
        :param pulumi.Input[int] heartbeat: Set the server AMQP 0-9-1 heartbeat timeout in seconds.
        :param pulumi.Input[int] instance_id: The CloudAMQP instance ID.
        :param pulumi.Input[str] log_exchange_level: Log level for the logger used for log integrations and the CloudAMQP Console log view.
        :param pulumi.Input[int] max_message_size: The largest allowed message payload size in bytes.
        :param pulumi.Input[int] queue_index_embed_msgs_below: Size in bytes below which to embed messages in the queue index.
        :param pulumi.Input[float] vm_memory_high_watermark: When the server will enter memory based flow-control as relative to the maximum available memory.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: RabbitConfigurationArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        This resource allows you update RabbitMQ config.

        Only available for dedicated subscription plans.

        ## Argument threshold values

        | Argument | Type | Default | Min | Max | Unit | Affect | Note |
        |---|---|---|---|---|---|---|---|
        | heartbeat | int | 120 | 0 | - |  | Only effects new connections |  |
        | connection_max | int | -1 | 1 | - |  | RabbitMQ restart required | -1 in the provider corresponds to INFINITY in the RabbitMQ config |
        | channel_max | int | 128 | 0 | - |  | Only effects new connections |  |
        | consumer_timeout | int | 7200000 | 10000 | 86400000 | milliseconds | Only effects new channels | -1 in the provider corresponds to false (disable) in the RabbitMQ config |
        | vm_memory_high_watermark | float | 0.81 | 0.4 | 0.9 |  | Applied immediately |  |
        | queue_index_embed_msgs_below | int | 4096 | 1 | 10485760 | bytes | Applied immediately for new queues, requires restart for existing queues |  |
        | max_message_size | int | 134217728 | 1 | 536870912 | bytes | Only effects new channels |  |
        | log_exchange_level | string | error | - | - |  | RabbitMQ restart required | debug, info, warning, error, critical |

        ## Dependency

        This resource depends on CloudAMQP instance identifier, `cloudamqp_instance.instance.id`.

        ## Import

        `cloudamqp_rabbitmq_configuration` can be imported using the CloudAMQP instance identifier.

        ```sh
         $ pulumi import cloudamqp:index/rabbitConfiguration:RabbitConfiguration config <instance_id>`
        ```

        :param str resource_name: The name of the resource.
        :param RabbitConfigurationArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(RabbitConfigurationArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 channel_max: Optional[pulumi.Input[int]] = None,
                 connection_max: Optional[pulumi.Input[int]] = None,
                 consumer_timeout: Optional[pulumi.Input[int]] = None,
                 heartbeat: Optional[pulumi.Input[int]] = None,
                 instance_id: Optional[pulumi.Input[int]] = None,
                 log_exchange_level: Optional[pulumi.Input[str]] = None,
                 max_message_size: Optional[pulumi.Input[int]] = None,
                 queue_index_embed_msgs_below: Optional[pulumi.Input[int]] = None,
                 vm_memory_high_watermark: Optional[pulumi.Input[float]] = None,
                 __props__=None):
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = RabbitConfigurationArgs.__new__(RabbitConfigurationArgs)

            __props__.__dict__["channel_max"] = channel_max
            __props__.__dict__["connection_max"] = connection_max
            __props__.__dict__["consumer_timeout"] = consumer_timeout
            __props__.__dict__["heartbeat"] = heartbeat
            if instance_id is None and not opts.urn:
                raise TypeError("Missing required property 'instance_id'")
            __props__.__dict__["instance_id"] = instance_id
            __props__.__dict__["log_exchange_level"] = log_exchange_level
            __props__.__dict__["max_message_size"] = max_message_size
            __props__.__dict__["queue_index_embed_msgs_below"] = queue_index_embed_msgs_below
            __props__.__dict__["vm_memory_high_watermark"] = vm_memory_high_watermark
        super(RabbitConfiguration, __self__).__init__(
            'cloudamqp:index/rabbitConfiguration:RabbitConfiguration',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            channel_max: Optional[pulumi.Input[int]] = None,
            connection_max: Optional[pulumi.Input[int]] = None,
            consumer_timeout: Optional[pulumi.Input[int]] = None,
            heartbeat: Optional[pulumi.Input[int]] = None,
            instance_id: Optional[pulumi.Input[int]] = None,
            log_exchange_level: Optional[pulumi.Input[str]] = None,
            max_message_size: Optional[pulumi.Input[int]] = None,
            queue_index_embed_msgs_below: Optional[pulumi.Input[int]] = None,
            vm_memory_high_watermark: Optional[pulumi.Input[float]] = None) -> 'RabbitConfiguration':
        """
        Get an existing RabbitConfiguration resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] channel_max: Set the maximum permissible number of channels per connection.
        :param pulumi.Input[int] connection_max: Set the maximum permissible number of connection.
        :param pulumi.Input[int] consumer_timeout: A consumer that has recevied a message and does not acknowledge that message within the timeout in milliseconds
        :param pulumi.Input[int] heartbeat: Set the server AMQP 0-9-1 heartbeat timeout in seconds.
        :param pulumi.Input[int] instance_id: The CloudAMQP instance ID.
        :param pulumi.Input[str] log_exchange_level: Log level for the logger used for log integrations and the CloudAMQP Console log view.
        :param pulumi.Input[int] max_message_size: The largest allowed message payload size in bytes.
        :param pulumi.Input[int] queue_index_embed_msgs_below: Size in bytes below which to embed messages in the queue index.
        :param pulumi.Input[float] vm_memory_high_watermark: When the server will enter memory based flow-control as relative to the maximum available memory.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _RabbitConfigurationState.__new__(_RabbitConfigurationState)

        __props__.__dict__["channel_max"] = channel_max
        __props__.__dict__["connection_max"] = connection_max
        __props__.__dict__["consumer_timeout"] = consumer_timeout
        __props__.__dict__["heartbeat"] = heartbeat
        __props__.__dict__["instance_id"] = instance_id
        __props__.__dict__["log_exchange_level"] = log_exchange_level
        __props__.__dict__["max_message_size"] = max_message_size
        __props__.__dict__["queue_index_embed_msgs_below"] = queue_index_embed_msgs_below
        __props__.__dict__["vm_memory_high_watermark"] = vm_memory_high_watermark
        return RabbitConfiguration(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="channelMax")
    def channel_max(self) -> pulumi.Output[Optional[int]]:
        """
        Set the maximum permissible number of channels per connection.
        """
        return pulumi.get(self, "channel_max")

    @property
    @pulumi.getter(name="connectionMax")
    def connection_max(self) -> pulumi.Output[Optional[int]]:
        """
        Set the maximum permissible number of connection.
        """
        return pulumi.get(self, "connection_max")

    @property
    @pulumi.getter(name="consumerTimeout")
    def consumer_timeout(self) -> pulumi.Output[Optional[int]]:
        """
        A consumer that has recevied a message and does not acknowledge that message within the timeout in milliseconds
        """
        return pulumi.get(self, "consumer_timeout")

    @property
    @pulumi.getter
    def heartbeat(self) -> pulumi.Output[Optional[int]]:
        """
        Set the server AMQP 0-9-1 heartbeat timeout in seconds.
        """
        return pulumi.get(self, "heartbeat")

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> pulumi.Output[int]:
        """
        The CloudAMQP instance ID.
        """
        return pulumi.get(self, "instance_id")

    @property
    @pulumi.getter(name="logExchangeLevel")
    def log_exchange_level(self) -> pulumi.Output[Optional[str]]:
        """
        Log level for the logger used for log integrations and the CloudAMQP Console log view.
        """
        return pulumi.get(self, "log_exchange_level")

    @property
    @pulumi.getter(name="maxMessageSize")
    def max_message_size(self) -> pulumi.Output[Optional[int]]:
        """
        The largest allowed message payload size in bytes.
        """
        return pulumi.get(self, "max_message_size")

    @property
    @pulumi.getter(name="queueIndexEmbedMsgsBelow")
    def queue_index_embed_msgs_below(self) -> pulumi.Output[Optional[int]]:
        """
        Size in bytes below which to embed messages in the queue index.
        """
        return pulumi.get(self, "queue_index_embed_msgs_below")

    @property
    @pulumi.getter(name="vmMemoryHighWatermark")
    def vm_memory_high_watermark(self) -> pulumi.Output[Optional[float]]:
        """
        When the server will enter memory based flow-control as relative to the maximum available memory.
        """
        return pulumi.get(self, "vm_memory_high_watermark")

