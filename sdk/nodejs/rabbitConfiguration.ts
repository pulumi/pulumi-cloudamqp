// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * This resource allows you update RabbitMQ config.
 *
 * Only available for dedicated subscription plans running ***RabbitMQ***.
 *
 * ## Example Usage
 *
 * <details>
 *   <summary>
 *     <b>
 *       <i>RabbitMQ configuration with default values</i>
 *     </b>
 *   </summary>
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as cloudamqp from "@pulumi/cloudamqp";
 *
 * const rabbitmqConfig = new cloudamqp.RabbitConfiguration("rabbitmq_config", {
 *     instanceId: instance.id,
 *     channelMax: 0,
 *     connectionMax: -1,
 *     consumerTimeout: 7200000,
 *     heartbeat: 120,
 *     logExchangeLevel: "error",
 *     maxMessageSize: 134217728,
 *     queueIndexEmbedMsgsBelow: 4096,
 *     vmMemoryHighWatermark: 0.81,
 *     clusterPartitionHandling: "autoheal",
 * });
 * ```
 *
 * </details>
 *
 * <details>
 *   <summary>
 *     <b>
 *       <i>Change log level and combine `cloudamqp.NodeActions` for RabbitMQ restart</i>
 *     </b>
 *   </summary>
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as cloudamqp from "@pulumi/cloudamqp";
 *
 * const rabbitmqConfig = new cloudamqp.RabbitConfiguration("rabbitmq_config", {
 *     instanceId: instance.id,
 *     channelMax: 0,
 *     connectionMax: -1,
 *     consumerTimeout: 7200000,
 *     heartbeat: 120,
 *     logExchangeLevel: "info",
 *     maxMessageSize: 134217728,
 *     queueIndexEmbedMsgsBelow: 4096,
 *     vmMemoryHighWatermark: 0.81,
 *     clusterPartitionHandling: "autoheal",
 * });
 * const listNodes = cloudamqp.getNodes({
 *     instanceId: instance.id,
 * });
 * const nodeAction = new cloudamqp.NodeActions("node_action", {
 *     instanceId: instance.id,
 *     nodeName: listNodes.then(listNodes => listNodes.nodes?.[0]?.name),
 *     action: "restart",
 * }, {
 *     dependsOn: [rabbitmqConfig],
 * });
 * ```
 *
 * </details>
 *
 * <details>
 *   <summary>
 *     <b>
 *       <i>Only change log level for exchange. All other values will be read from the RabbitMQ configuration.</i>
 *     </b>
 *   </summary>
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as cloudamqp from "@pulumi/cloudamqp";
 *
 * const rabbitConfig = new cloudamqp.RabbitConfiguration("rabbit_config", {
 *     instanceId: instance.id,
 *     logExchangeLevel: "info",
 * });
 * ```
 *
 * </details>
 *
 * ## Argument threshold values
 *
 * |  Argument   |  Type  |  Default  |  Min  |    Max    |     Unit     |                              Affect                               |                               Note                                |
 * |-------------|--------|-----------|-------|-----------|--------------|-------------------------------------------------------------------|-------------------------------------------------------------------|
 * | heartbeat   | int    |       120 |     0 | -         |              | Only effects new                                                  |                                                                   |
 * | connection_ | int    |        -1 |     1 | -         |              | RabbitMQ restart                                                  | -1 in the provider corresponds to INFINITY in the RabbitMQ        |
 * | channel_    | int    |       128 |     0 | -         |              | Only effects new                                                  |                                                                   |
 * | consumer_   | int    |   7200000 | 10000 |  86400000 | milliseconds | Only effects new                                                  | -1 in the provider corresponds to false (disable) in the RabbitMQ |
 * | vm_         | float  |      0.81 |   0.4 |       0.9 |              | Applied                                                           |                                                                   |
 * | queue_      | int    |      4096 |     0 |  10485760 | bytes        | Applied immediately for new queues, requires restart for existing |                                                                   |
 * | max_        | int    | 134217728 |     1 | 536870912 | bytes        | Only effects new                                                  |                                                                   |
 * | log_        | string | error     | -     | -         |              | RabbitMQ restart                                                  | debug, info, warning, error,                                      |
 * | cluster_    | string | see       | -     | -         |              | Applied                                                           | autoheal, pause_                                                  |
 */
export class RabbitConfiguration extends pulumi.CustomResource {
    /**
     * Get an existing RabbitConfiguration resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: RabbitConfigurationState, opts?: pulumi.CustomResourceOptions): RabbitConfiguration {
        return new RabbitConfiguration(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'cloudamqp:index/rabbitConfiguration:RabbitConfiguration';

    /**
     * Returns true if the given object is an instance of RabbitConfiguration.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is RabbitConfiguration {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === RabbitConfiguration.__pulumiType;
    }

    /**
     * Set the maximum permissible number of channels per connection.
     */
    public readonly channelMax!: pulumi.Output<number>;
    /**
     * Set how the cluster should handle network partition.
     */
    public readonly clusterPartitionHandling!: pulumi.Output<string>;
    /**
     * Set the maximum permissible number of connection.
     */
    public readonly connectionMax!: pulumi.Output<number>;
    /**
     * A consumer that has recevied a message and does not acknowledge that message within the timeout in milliseconds
     */
    public readonly consumerTimeout!: pulumi.Output<number>;
    /**
     * Set the server AMQP 0-9-1 heartbeat timeout in seconds.
     */
    public readonly heartbeat!: pulumi.Output<number>;
    /**
     * The CloudAMQP instance ID.
     */
    public readonly instanceId!: pulumi.Output<number>;
    /**
     * Log level for the logger used for log integrations and the CloudAMQP Console log view.
     *
     * *Note: Requires a restart of RabbitMQ to be applied.*
     */
    public readonly logExchangeLevel!: pulumi.Output<string>;
    /**
     * The largest allowed message payload size in bytes.
     */
    public readonly maxMessageSize!: pulumi.Output<number>;
    /**
     * Size in bytes below which to embed messages in the queue index. 0 will turn off payload embedding in the queue index.
     */
    public readonly queueIndexEmbedMsgsBelow!: pulumi.Output<number>;
    /**
     * Configurable sleep time in seconds between retries for RabbitMQ configuration. Default set to 60 seconds.
     */
    public readonly sleep!: pulumi.Output<number | undefined>;
    /**
     * Configurable timeout time in seconds for RabbitMQ configuration. Default set to 3600 seconds.
     */
    public readonly timeout!: pulumi.Output<number | undefined>;
    /**
     * When the server will enter memory based flow-control as relative to the maximum available memory.
     */
    public readonly vmMemoryHighWatermark!: pulumi.Output<number>;

    /**
     * Create a RabbitConfiguration resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: RabbitConfigurationArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: RabbitConfigurationArgs | RabbitConfigurationState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as RabbitConfigurationState | undefined;
            resourceInputs["channelMax"] = state ? state.channelMax : undefined;
            resourceInputs["clusterPartitionHandling"] = state ? state.clusterPartitionHandling : undefined;
            resourceInputs["connectionMax"] = state ? state.connectionMax : undefined;
            resourceInputs["consumerTimeout"] = state ? state.consumerTimeout : undefined;
            resourceInputs["heartbeat"] = state ? state.heartbeat : undefined;
            resourceInputs["instanceId"] = state ? state.instanceId : undefined;
            resourceInputs["logExchangeLevel"] = state ? state.logExchangeLevel : undefined;
            resourceInputs["maxMessageSize"] = state ? state.maxMessageSize : undefined;
            resourceInputs["queueIndexEmbedMsgsBelow"] = state ? state.queueIndexEmbedMsgsBelow : undefined;
            resourceInputs["sleep"] = state ? state.sleep : undefined;
            resourceInputs["timeout"] = state ? state.timeout : undefined;
            resourceInputs["vmMemoryHighWatermark"] = state ? state.vmMemoryHighWatermark : undefined;
        } else {
            const args = argsOrState as RabbitConfigurationArgs | undefined;
            if ((!args || args.instanceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'instanceId'");
            }
            resourceInputs["channelMax"] = args ? args.channelMax : undefined;
            resourceInputs["clusterPartitionHandling"] = args ? args.clusterPartitionHandling : undefined;
            resourceInputs["connectionMax"] = args ? args.connectionMax : undefined;
            resourceInputs["consumerTimeout"] = args ? args.consumerTimeout : undefined;
            resourceInputs["heartbeat"] = args ? args.heartbeat : undefined;
            resourceInputs["instanceId"] = args ? args.instanceId : undefined;
            resourceInputs["logExchangeLevel"] = args ? args.logExchangeLevel : undefined;
            resourceInputs["maxMessageSize"] = args ? args.maxMessageSize : undefined;
            resourceInputs["queueIndexEmbedMsgsBelow"] = args ? args.queueIndexEmbedMsgsBelow : undefined;
            resourceInputs["sleep"] = args ? args.sleep : undefined;
            resourceInputs["timeout"] = args ? args.timeout : undefined;
            resourceInputs["vmMemoryHighWatermark"] = args ? args.vmMemoryHighWatermark : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(RabbitConfiguration.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering RabbitConfiguration resources.
 */
export interface RabbitConfigurationState {
    /**
     * Set the maximum permissible number of channels per connection.
     */
    channelMax?: pulumi.Input<number>;
    /**
     * Set how the cluster should handle network partition.
     */
    clusterPartitionHandling?: pulumi.Input<string>;
    /**
     * Set the maximum permissible number of connection.
     */
    connectionMax?: pulumi.Input<number>;
    /**
     * A consumer that has recevied a message and does not acknowledge that message within the timeout in milliseconds
     */
    consumerTimeout?: pulumi.Input<number>;
    /**
     * Set the server AMQP 0-9-1 heartbeat timeout in seconds.
     */
    heartbeat?: pulumi.Input<number>;
    /**
     * The CloudAMQP instance ID.
     */
    instanceId?: pulumi.Input<number>;
    /**
     * Log level for the logger used for log integrations and the CloudAMQP Console log view.
     *
     * *Note: Requires a restart of RabbitMQ to be applied.*
     */
    logExchangeLevel?: pulumi.Input<string>;
    /**
     * The largest allowed message payload size in bytes.
     */
    maxMessageSize?: pulumi.Input<number>;
    /**
     * Size in bytes below which to embed messages in the queue index. 0 will turn off payload embedding in the queue index.
     */
    queueIndexEmbedMsgsBelow?: pulumi.Input<number>;
    /**
     * Configurable sleep time in seconds between retries for RabbitMQ configuration. Default set to 60 seconds.
     */
    sleep?: pulumi.Input<number>;
    /**
     * Configurable timeout time in seconds for RabbitMQ configuration. Default set to 3600 seconds.
     */
    timeout?: pulumi.Input<number>;
    /**
     * When the server will enter memory based flow-control as relative to the maximum available memory.
     */
    vmMemoryHighWatermark?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a RabbitConfiguration resource.
 */
export interface RabbitConfigurationArgs {
    /**
     * Set the maximum permissible number of channels per connection.
     */
    channelMax?: pulumi.Input<number>;
    /**
     * Set how the cluster should handle network partition.
     */
    clusterPartitionHandling?: pulumi.Input<string>;
    /**
     * Set the maximum permissible number of connection.
     */
    connectionMax?: pulumi.Input<number>;
    /**
     * A consumer that has recevied a message and does not acknowledge that message within the timeout in milliseconds
     */
    consumerTimeout?: pulumi.Input<number>;
    /**
     * Set the server AMQP 0-9-1 heartbeat timeout in seconds.
     */
    heartbeat?: pulumi.Input<number>;
    /**
     * The CloudAMQP instance ID.
     */
    instanceId: pulumi.Input<number>;
    /**
     * Log level for the logger used for log integrations and the CloudAMQP Console log view.
     *
     * *Note: Requires a restart of RabbitMQ to be applied.*
     */
    logExchangeLevel?: pulumi.Input<string>;
    /**
     * The largest allowed message payload size in bytes.
     */
    maxMessageSize?: pulumi.Input<number>;
    /**
     * Size in bytes below which to embed messages in the queue index. 0 will turn off payload embedding in the queue index.
     */
    queueIndexEmbedMsgsBelow?: pulumi.Input<number>;
    /**
     * Configurable sleep time in seconds between retries for RabbitMQ configuration. Default set to 60 seconds.
     */
    sleep?: pulumi.Input<number>;
    /**
     * Configurable timeout time in seconds for RabbitMQ configuration. Default set to 3600 seconds.
     */
    timeout?: pulumi.Input<number>;
    /**
     * When the server will enter memory based flow-control as relative to the maximum available memory.
     */
    vmMemoryHighWatermark?: pulumi.Input<number>;
}
