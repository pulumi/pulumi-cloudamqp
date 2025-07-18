// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.cloudamqp;

import com.pulumi.cloudamqp.RabbitConfigurationArgs;
import com.pulumi.cloudamqp.Utilities;
import com.pulumi.cloudamqp.inputs.RabbitConfigurationState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Double;
import java.lang.Integer;
import java.lang.String;
import javax.annotation.Nullable;

/**
 * ## Import
 * 
 * `cloudamqp_rabbitmq_configuration` can be imported using the CloudAMQP instance identifier.  To
 * 
 * retrieve the identifier, use [CloudAMQP API list intances].
 * 
 * From Terraform v1.5.0, the `import` block can be used to import this resource:
 * 
 * hcl
 * 
 * import {
 * 
 *   to = cloudamqp_rabbitmq_configuration.config
 * 
 *   id = cloudamqp_instance.instance.id
 * 
 * }
 * 
 * Or use Terraform CLI:
 * 
 * ```sh
 * $ pulumi import cloudamqp:index/rabbitConfiguration:RabbitConfiguration config &lt;instance_id&gt;`
 * ```
 * 
 */
@ResourceType(type="cloudamqp:index/rabbitConfiguration:RabbitConfiguration")
public class RabbitConfiguration extends com.pulumi.resources.CustomResource {
    /**
     * Set the maximum permissible number of
     * channels per connection.
     * 
     */
    @Export(name="channelMax", refs={Integer.class}, tree="[0]")
    private Output<Integer> channelMax;

    /**
     * @return Set the maximum permissible number of
     * channels per connection.
     * 
     */
    public Output<Integer> channelMax() {
        return this.channelMax;
    }
    /**
     * Set how the cluster should handle network
     * partition.
     * 
     */
    @Export(name="clusterPartitionHandling", refs={String.class}, tree="[0]")
    private Output<String> clusterPartitionHandling;

    /**
     * @return Set how the cluster should handle network
     * partition.
     * 
     */
    public Output<String> clusterPartitionHandling() {
        return this.clusterPartitionHandling;
    }
    /**
     * Set the maximum permissible number of
     * connection.
     * 
     */
    @Export(name="connectionMax", refs={Integer.class}, tree="[0]")
    private Output<Integer> connectionMax;

    /**
     * @return Set the maximum permissible number of
     * connection.
     * 
     */
    public Output<Integer> connectionMax() {
        return this.connectionMax;
    }
    /**
     * A consumer that has recevied a message and
     * does not acknowledge that message within the timeout in
     * milliseconds
     * 
     */
    @Export(name="consumerTimeout", refs={Integer.class}, tree="[0]")
    private Output<Integer> consumerTimeout;

    /**
     * @return A consumer that has recevied a message and
     * does not acknowledge that message within the timeout in
     * milliseconds
     * 
     */
    public Output<Integer> consumerTimeout() {
        return this.consumerTimeout;
    }
    /**
     * Set the server AMQP 0-9-1 heartbeat timeout
     * in seconds.
     * 
     */
    @Export(name="heartbeat", refs={Integer.class}, tree="[0]")
    private Output<Integer> heartbeat;

    /**
     * @return Set the server AMQP 0-9-1 heartbeat timeout
     * in seconds.
     * 
     */
    public Output<Integer> heartbeat() {
        return this.heartbeat;
    }
    /**
     * The CloudAMQP instance ID.
     * 
     */
    @Export(name="instanceId", refs={Integer.class}, tree="[0]")
    private Output<Integer> instanceId;

    /**
     * @return The CloudAMQP instance ID.
     * 
     */
    public Output<Integer> instanceId() {
        return this.instanceId;
    }
    /**
     * Log level for the logger used for log
     * integrations and the CloudAMQP Console log view.
     * 
     */
    @Export(name="logExchangeLevel", refs={String.class}, tree="[0]")
    private Output<String> logExchangeLevel;

    /**
     * @return Log level for the logger used for log
     * integrations and the CloudAMQP Console log view.
     * 
     */
    public Output<String> logExchangeLevel() {
        return this.logExchangeLevel;
    }
    /**
     * The largest allowed message payload size in
     * bytes.
     * 
     */
    @Export(name="maxMessageSize", refs={Integer.class}, tree="[0]")
    private Output<Integer> maxMessageSize;

    /**
     * @return The largest allowed message payload size in
     * bytes.
     * 
     */
    public Output<Integer> maxMessageSize() {
        return this.maxMessageSize;
    }
    /**
     * Size in bytes below which to embed messages
     * in the queue index. 0 will turn off payload embedding in the
     * queue index.
     * 
     */
    @Export(name="queueIndexEmbedMsgsBelow", refs={Integer.class}, tree="[0]")
    private Output<Integer> queueIndexEmbedMsgsBelow;

    /**
     * @return Size in bytes below which to embed messages
     * in the queue index. 0 will turn off payload embedding in the
     * queue index.
     * 
     */
    public Output<Integer> queueIndexEmbedMsgsBelow() {
        return this.queueIndexEmbedMsgsBelow;
    }
    /**
     * Configurable sleep time in seconds between retries
     * for RabbitMQ configuration. Default set to 60 seconds.
     * 
     */
    @Export(name="sleep", refs={Integer.class}, tree="[0]")
    private Output<Integer> sleep;

    /**
     * @return Configurable sleep time in seconds between retries
     * for RabbitMQ configuration. Default set to 60 seconds.
     * 
     */
    public Output<Integer> sleep() {
        return this.sleep;
    }
    /**
     * Configurable timeout time in seconds for RabbitMQ
     * configuration. Default set to 3600 seconds.
     * 
     */
    @Export(name="timeout", refs={Integer.class}, tree="[0]")
    private Output<Integer> timeout;

    /**
     * @return Configurable timeout time in seconds for RabbitMQ
     * configuration. Default set to 3600 seconds.
     * 
     */
    public Output<Integer> timeout() {
        return this.timeout;
    }
    /**
     * When the server will enter memory based
     * flow-control as relative to the maximum available memory.
     * 
     */
    @Export(name="vmMemoryHighWatermark", refs={Double.class}, tree="[0]")
    private Output<Double> vmMemoryHighWatermark;

    /**
     * @return When the server will enter memory based
     * flow-control as relative to the maximum available memory.
     * 
     */
    public Output<Double> vmMemoryHighWatermark() {
        return this.vmMemoryHighWatermark;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public RabbitConfiguration(java.lang.String name) {
        this(name, RabbitConfigurationArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public RabbitConfiguration(java.lang.String name, RabbitConfigurationArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public RabbitConfiguration(java.lang.String name, RabbitConfigurationArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("cloudamqp:index/rabbitConfiguration:RabbitConfiguration", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private RabbitConfiguration(java.lang.String name, Output<java.lang.String> id, @Nullable RabbitConfigurationState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("cloudamqp:index/rabbitConfiguration:RabbitConfiguration", name, state, makeResourceOptions(options, id), false);
    }

    private static RabbitConfigurationArgs makeArgs(RabbitConfigurationArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? RabbitConfigurationArgs.Empty : args;
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<java.lang.String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .build();
        return com.pulumi.resources.CustomResourceOptions.merge(defaultOptions, options, id);
    }

    /**
     * Get an existing Host resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state
     * @param options Optional settings to control the behavior of the CustomResource.
     */
    public static RabbitConfiguration get(java.lang.String name, Output<java.lang.String> id, @Nullable RabbitConfigurationState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new RabbitConfiguration(name, id, state, options);
    }
}
