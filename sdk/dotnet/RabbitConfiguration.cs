// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.CloudAmqp
{
    /// <summary>
    /// This resource allows you update RabbitMQ config.
    /// 
    /// Only available for dedicated subscription plans.
    /// 
    /// ## Argument threshold values
    /// 
    /// | Argument | Type | Default | Min | Max | Unit | Affect | Note |
    /// |---|---|---|---|---|---|---|---|
    /// | heartbeat | int | 120 | 0 | - |  | Only effects new connections |  |
    /// | connection_max | int | -1 | 1 | - |  | RabbitMQ restart required | -1 in the provider corresponds to INFINITY in the RabbitMQ config |
    /// | channel_max | int | 128 | 0 | - |  | Only effects new connections |  |
    /// | consumer_timeout | int | 7200000 | 10000 | 86400000 | milliseconds | Only effects new channels | -1 in the provider corresponds to false (disable) in the RabbitMQ config |
    /// | vm_memory_high_watermark | float | 0.81 | 0.4 | 0.9 |  | Applied immediately |  |
    /// | queue_index_embed_msgs_below | int | 4096 | 1 | 10485760 | bytes | Applied immediately for new queues, requires restart for existing queues |  |
    /// | max_message_size | int | 134217728 | 1 | 536870912 | bytes | Only effects new channels |  |
    /// | log_exchange_level | string | error | - | - |  | RabbitMQ restart required | debug, info, warning, error, critical |
    /// 
    /// ## Dependency
    /// 
    /// This resource depends on CloudAMQP instance identifier, `cloudamqp_instance.instance.id`.
    /// 
    /// ## Import
    /// 
    /// `cloudamqp_rabbitmq_configuration` can be imported using the CloudAMQP instance identifier.
    /// 
    /// ```sh
    ///  $ pulumi import cloudamqp:index/rabbitConfiguration:RabbitConfiguration config &lt;instance_id&gt;`
    /// ```
    /// </summary>
    [CloudAmqpResourceType("cloudamqp:index/rabbitConfiguration:RabbitConfiguration")]
    public partial class RabbitConfiguration : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Set the maximum permissible number of channels per connection.
        /// </summary>
        [Output("channelMax")]
        public Output<int> ChannelMax { get; private set; } = null!;

        /// <summary>
        /// Set the maximum permissible number of connection.
        /// </summary>
        [Output("connectionMax")]
        public Output<int> ConnectionMax { get; private set; } = null!;

        /// <summary>
        /// A consumer that has recevied a message and does not acknowledge that message within the timeout in milliseconds
        /// </summary>
        [Output("consumerTimeout")]
        public Output<int> ConsumerTimeout { get; private set; } = null!;

        /// <summary>
        /// Set the server AMQP 0-9-1 heartbeat timeout in seconds.
        /// </summary>
        [Output("heartbeat")]
        public Output<int> Heartbeat { get; private set; } = null!;

        /// <summary>
        /// The CloudAMQP instance ID.
        /// </summary>
        [Output("instanceId")]
        public Output<int> InstanceId { get; private set; } = null!;

        /// <summary>
        /// Log level for the logger used for log integrations and the CloudAMQP Console log view.
        /// </summary>
        [Output("logExchangeLevel")]
        public Output<string> LogExchangeLevel { get; private set; } = null!;

        /// <summary>
        /// The largest allowed message payload size in bytes.
        /// </summary>
        [Output("maxMessageSize")]
        public Output<int> MaxMessageSize { get; private set; } = null!;

        /// <summary>
        /// Size in bytes below which to embed messages in the queue index.
        /// </summary>
        [Output("queueIndexEmbedMsgsBelow")]
        public Output<int> QueueIndexEmbedMsgsBelow { get; private set; } = null!;

        /// <summary>
        /// When the server will enter memory based flow-control as relative to the maximum available memory.
        /// </summary>
        [Output("vmMemoryHighWatermark")]
        public Output<double> VmMemoryHighWatermark { get; private set; } = null!;


        /// <summary>
        /// Create a RabbitConfiguration resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public RabbitConfiguration(string name, RabbitConfigurationArgs args, CustomResourceOptions? options = null)
            : base("cloudamqp:index/rabbitConfiguration:RabbitConfiguration", name, args ?? new RabbitConfigurationArgs(), MakeResourceOptions(options, ""))
        {
        }

        private RabbitConfiguration(string name, Input<string> id, RabbitConfigurationState? state = null, CustomResourceOptions? options = null)
            : base("cloudamqp:index/rabbitConfiguration:RabbitConfiguration", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing RabbitConfiguration resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static RabbitConfiguration Get(string name, Input<string> id, RabbitConfigurationState? state = null, CustomResourceOptions? options = null)
        {
            return new RabbitConfiguration(name, id, state, options);
        }
    }

    public sealed class RabbitConfigurationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Set the maximum permissible number of channels per connection.
        /// </summary>
        [Input("channelMax")]
        public Input<int>? ChannelMax { get; set; }

        /// <summary>
        /// Set the maximum permissible number of connection.
        /// </summary>
        [Input("connectionMax")]
        public Input<int>? ConnectionMax { get; set; }

        /// <summary>
        /// A consumer that has recevied a message and does not acknowledge that message within the timeout in milliseconds
        /// </summary>
        [Input("consumerTimeout")]
        public Input<int>? ConsumerTimeout { get; set; }

        /// <summary>
        /// Set the server AMQP 0-9-1 heartbeat timeout in seconds.
        /// </summary>
        [Input("heartbeat")]
        public Input<int>? Heartbeat { get; set; }

        /// <summary>
        /// The CloudAMQP instance ID.
        /// </summary>
        [Input("instanceId", required: true)]
        public Input<int> InstanceId { get; set; } = null!;

        /// <summary>
        /// Log level for the logger used for log integrations and the CloudAMQP Console log view.
        /// </summary>
        [Input("logExchangeLevel")]
        public Input<string>? LogExchangeLevel { get; set; }

        /// <summary>
        /// The largest allowed message payload size in bytes.
        /// </summary>
        [Input("maxMessageSize")]
        public Input<int>? MaxMessageSize { get; set; }

        /// <summary>
        /// Size in bytes below which to embed messages in the queue index.
        /// </summary>
        [Input("queueIndexEmbedMsgsBelow")]
        public Input<int>? QueueIndexEmbedMsgsBelow { get; set; }

        /// <summary>
        /// When the server will enter memory based flow-control as relative to the maximum available memory.
        /// </summary>
        [Input("vmMemoryHighWatermark")]
        public Input<double>? VmMemoryHighWatermark { get; set; }

        public RabbitConfigurationArgs()
        {
        }
        public static new RabbitConfigurationArgs Empty => new RabbitConfigurationArgs();
    }

    public sealed class RabbitConfigurationState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Set the maximum permissible number of channels per connection.
        /// </summary>
        [Input("channelMax")]
        public Input<int>? ChannelMax { get; set; }

        /// <summary>
        /// Set the maximum permissible number of connection.
        /// </summary>
        [Input("connectionMax")]
        public Input<int>? ConnectionMax { get; set; }

        /// <summary>
        /// A consumer that has recevied a message and does not acknowledge that message within the timeout in milliseconds
        /// </summary>
        [Input("consumerTimeout")]
        public Input<int>? ConsumerTimeout { get; set; }

        /// <summary>
        /// Set the server AMQP 0-9-1 heartbeat timeout in seconds.
        /// </summary>
        [Input("heartbeat")]
        public Input<int>? Heartbeat { get; set; }

        /// <summary>
        /// The CloudAMQP instance ID.
        /// </summary>
        [Input("instanceId")]
        public Input<int>? InstanceId { get; set; }

        /// <summary>
        /// Log level for the logger used for log integrations and the CloudAMQP Console log view.
        /// </summary>
        [Input("logExchangeLevel")]
        public Input<string>? LogExchangeLevel { get; set; }

        /// <summary>
        /// The largest allowed message payload size in bytes.
        /// </summary>
        [Input("maxMessageSize")]
        public Input<int>? MaxMessageSize { get; set; }

        /// <summary>
        /// Size in bytes below which to embed messages in the queue index.
        /// </summary>
        [Input("queueIndexEmbedMsgsBelow")]
        public Input<int>? QueueIndexEmbedMsgsBelow { get; set; }

        /// <summary>
        /// When the server will enter memory based flow-control as relative to the maximum available memory.
        /// </summary>
        [Input("vmMemoryHighWatermark")]
        public Input<double>? VmMemoryHighWatermark { get; set; }

        public RabbitConfigurationState()
        {
        }
        public static new RabbitConfigurationState Empty => new RabbitConfigurationState();
    }
}
