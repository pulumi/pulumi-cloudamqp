// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.cloudamqp;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Double;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class RabbitConfigurationArgs extends com.pulumi.resources.ResourceArgs {

    public static final RabbitConfigurationArgs Empty = new RabbitConfigurationArgs();

    /**
     * Set the maximum permissible number of channels per connection.
     * 
     */
    @Import(name="channelMax")
    private @Nullable Output<Integer> channelMax;

    /**
     * @return Set the maximum permissible number of channels per connection.
     * 
     */
    public Optional<Output<Integer>> channelMax() {
        return Optional.ofNullable(this.channelMax);
    }

    /**
     * Set how the cluster should handle network partition.
     * 
     */
    @Import(name="clusterPartitionHandling")
    private @Nullable Output<String> clusterPartitionHandling;

    /**
     * @return Set how the cluster should handle network partition.
     * 
     */
    public Optional<Output<String>> clusterPartitionHandling() {
        return Optional.ofNullable(this.clusterPartitionHandling);
    }

    /**
     * Set the maximum permissible number of connection.
     * 
     */
    @Import(name="connectionMax")
    private @Nullable Output<Integer> connectionMax;

    /**
     * @return Set the maximum permissible number of connection.
     * 
     */
    public Optional<Output<Integer>> connectionMax() {
        return Optional.ofNullable(this.connectionMax);
    }

    /**
     * A consumer that has recevied a message and does not acknowledge that message within the timeout in milliseconds
     * 
     */
    @Import(name="consumerTimeout")
    private @Nullable Output<Integer> consumerTimeout;

    /**
     * @return A consumer that has recevied a message and does not acknowledge that message within the timeout in milliseconds
     * 
     */
    public Optional<Output<Integer>> consumerTimeout() {
        return Optional.ofNullable(this.consumerTimeout);
    }

    /**
     * Set the server AMQP 0-9-1 heartbeat timeout in seconds.
     * 
     */
    @Import(name="heartbeat")
    private @Nullable Output<Integer> heartbeat;

    /**
     * @return Set the server AMQP 0-9-1 heartbeat timeout in seconds.
     * 
     */
    public Optional<Output<Integer>> heartbeat() {
        return Optional.ofNullable(this.heartbeat);
    }

    /**
     * The CloudAMQP instance ID.
     * 
     */
    @Import(name="instanceId", required=true)
    private Output<Integer> instanceId;

    /**
     * @return The CloudAMQP instance ID.
     * 
     */
    public Output<Integer> instanceId() {
        return this.instanceId;
    }

    /**
     * Log level for the logger used for log integrations and the CloudAMQP Console log view.
     * 
     * *Note: Requires a restart of RabbitMQ to be applied.*
     * 
     */
    @Import(name="logExchangeLevel")
    private @Nullable Output<String> logExchangeLevel;

    /**
     * @return Log level for the logger used for log integrations and the CloudAMQP Console log view.
     * 
     * *Note: Requires a restart of RabbitMQ to be applied.*
     * 
     */
    public Optional<Output<String>> logExchangeLevel() {
        return Optional.ofNullable(this.logExchangeLevel);
    }

    /**
     * The largest allowed message payload size in bytes.
     * 
     */
    @Import(name="maxMessageSize")
    private @Nullable Output<Integer> maxMessageSize;

    /**
     * @return The largest allowed message payload size in bytes.
     * 
     */
    public Optional<Output<Integer>> maxMessageSize() {
        return Optional.ofNullable(this.maxMessageSize);
    }

    /**
     * Size in bytes below which to embed messages in the queue index. 0 will turn off payload embedding in the queue index.
     * 
     */
    @Import(name="queueIndexEmbedMsgsBelow")
    private @Nullable Output<Integer> queueIndexEmbedMsgsBelow;

    /**
     * @return Size in bytes below which to embed messages in the queue index. 0 will turn off payload embedding in the queue index.
     * 
     */
    public Optional<Output<Integer>> queueIndexEmbedMsgsBelow() {
        return Optional.ofNullable(this.queueIndexEmbedMsgsBelow);
    }

    /**
     * Configurable sleep time in seconds between retries for RabbitMQ configuration. Default set to 60 seconds.
     * 
     */
    @Import(name="sleep")
    private @Nullable Output<Integer> sleep;

    /**
     * @return Configurable sleep time in seconds between retries for RabbitMQ configuration. Default set to 60 seconds.
     * 
     */
    public Optional<Output<Integer>> sleep() {
        return Optional.ofNullable(this.sleep);
    }

    /**
     * Configurable timeout time in seconds for RabbitMQ configuration. Default set to 3600 seconds.
     * 
     */
    @Import(name="timeout")
    private @Nullable Output<Integer> timeout;

    /**
     * @return Configurable timeout time in seconds for RabbitMQ configuration. Default set to 3600 seconds.
     * 
     */
    public Optional<Output<Integer>> timeout() {
        return Optional.ofNullable(this.timeout);
    }

    /**
     * When the server will enter memory based flow-control as relative to the maximum available memory.
     * 
     */
    @Import(name="vmMemoryHighWatermark")
    private @Nullable Output<Double> vmMemoryHighWatermark;

    /**
     * @return When the server will enter memory based flow-control as relative to the maximum available memory.
     * 
     */
    public Optional<Output<Double>> vmMemoryHighWatermark() {
        return Optional.ofNullable(this.vmMemoryHighWatermark);
    }

    private RabbitConfigurationArgs() {}

    private RabbitConfigurationArgs(RabbitConfigurationArgs $) {
        this.channelMax = $.channelMax;
        this.clusterPartitionHandling = $.clusterPartitionHandling;
        this.connectionMax = $.connectionMax;
        this.consumerTimeout = $.consumerTimeout;
        this.heartbeat = $.heartbeat;
        this.instanceId = $.instanceId;
        this.logExchangeLevel = $.logExchangeLevel;
        this.maxMessageSize = $.maxMessageSize;
        this.queueIndexEmbedMsgsBelow = $.queueIndexEmbedMsgsBelow;
        this.sleep = $.sleep;
        this.timeout = $.timeout;
        this.vmMemoryHighWatermark = $.vmMemoryHighWatermark;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(RabbitConfigurationArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private RabbitConfigurationArgs $;

        public Builder() {
            $ = new RabbitConfigurationArgs();
        }

        public Builder(RabbitConfigurationArgs defaults) {
            $ = new RabbitConfigurationArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param channelMax Set the maximum permissible number of channels per connection.
         * 
         * @return builder
         * 
         */
        public Builder channelMax(@Nullable Output<Integer> channelMax) {
            $.channelMax = channelMax;
            return this;
        }

        /**
         * @param channelMax Set the maximum permissible number of channels per connection.
         * 
         * @return builder
         * 
         */
        public Builder channelMax(Integer channelMax) {
            return channelMax(Output.of(channelMax));
        }

        /**
         * @param clusterPartitionHandling Set how the cluster should handle network partition.
         * 
         * @return builder
         * 
         */
        public Builder clusterPartitionHandling(@Nullable Output<String> clusterPartitionHandling) {
            $.clusterPartitionHandling = clusterPartitionHandling;
            return this;
        }

        /**
         * @param clusterPartitionHandling Set how the cluster should handle network partition.
         * 
         * @return builder
         * 
         */
        public Builder clusterPartitionHandling(String clusterPartitionHandling) {
            return clusterPartitionHandling(Output.of(clusterPartitionHandling));
        }

        /**
         * @param connectionMax Set the maximum permissible number of connection.
         * 
         * @return builder
         * 
         */
        public Builder connectionMax(@Nullable Output<Integer> connectionMax) {
            $.connectionMax = connectionMax;
            return this;
        }

        /**
         * @param connectionMax Set the maximum permissible number of connection.
         * 
         * @return builder
         * 
         */
        public Builder connectionMax(Integer connectionMax) {
            return connectionMax(Output.of(connectionMax));
        }

        /**
         * @param consumerTimeout A consumer that has recevied a message and does not acknowledge that message within the timeout in milliseconds
         * 
         * @return builder
         * 
         */
        public Builder consumerTimeout(@Nullable Output<Integer> consumerTimeout) {
            $.consumerTimeout = consumerTimeout;
            return this;
        }

        /**
         * @param consumerTimeout A consumer that has recevied a message and does not acknowledge that message within the timeout in milliseconds
         * 
         * @return builder
         * 
         */
        public Builder consumerTimeout(Integer consumerTimeout) {
            return consumerTimeout(Output.of(consumerTimeout));
        }

        /**
         * @param heartbeat Set the server AMQP 0-9-1 heartbeat timeout in seconds.
         * 
         * @return builder
         * 
         */
        public Builder heartbeat(@Nullable Output<Integer> heartbeat) {
            $.heartbeat = heartbeat;
            return this;
        }

        /**
         * @param heartbeat Set the server AMQP 0-9-1 heartbeat timeout in seconds.
         * 
         * @return builder
         * 
         */
        public Builder heartbeat(Integer heartbeat) {
            return heartbeat(Output.of(heartbeat));
        }

        /**
         * @param instanceId The CloudAMQP instance ID.
         * 
         * @return builder
         * 
         */
        public Builder instanceId(Output<Integer> instanceId) {
            $.instanceId = instanceId;
            return this;
        }

        /**
         * @param instanceId The CloudAMQP instance ID.
         * 
         * @return builder
         * 
         */
        public Builder instanceId(Integer instanceId) {
            return instanceId(Output.of(instanceId));
        }

        /**
         * @param logExchangeLevel Log level for the logger used for log integrations and the CloudAMQP Console log view.
         * 
         * *Note: Requires a restart of RabbitMQ to be applied.*
         * 
         * @return builder
         * 
         */
        public Builder logExchangeLevel(@Nullable Output<String> logExchangeLevel) {
            $.logExchangeLevel = logExchangeLevel;
            return this;
        }

        /**
         * @param logExchangeLevel Log level for the logger used for log integrations and the CloudAMQP Console log view.
         * 
         * *Note: Requires a restart of RabbitMQ to be applied.*
         * 
         * @return builder
         * 
         */
        public Builder logExchangeLevel(String logExchangeLevel) {
            return logExchangeLevel(Output.of(logExchangeLevel));
        }

        /**
         * @param maxMessageSize The largest allowed message payload size in bytes.
         * 
         * @return builder
         * 
         */
        public Builder maxMessageSize(@Nullable Output<Integer> maxMessageSize) {
            $.maxMessageSize = maxMessageSize;
            return this;
        }

        /**
         * @param maxMessageSize The largest allowed message payload size in bytes.
         * 
         * @return builder
         * 
         */
        public Builder maxMessageSize(Integer maxMessageSize) {
            return maxMessageSize(Output.of(maxMessageSize));
        }

        /**
         * @param queueIndexEmbedMsgsBelow Size in bytes below which to embed messages in the queue index. 0 will turn off payload embedding in the queue index.
         * 
         * @return builder
         * 
         */
        public Builder queueIndexEmbedMsgsBelow(@Nullable Output<Integer> queueIndexEmbedMsgsBelow) {
            $.queueIndexEmbedMsgsBelow = queueIndexEmbedMsgsBelow;
            return this;
        }

        /**
         * @param queueIndexEmbedMsgsBelow Size in bytes below which to embed messages in the queue index. 0 will turn off payload embedding in the queue index.
         * 
         * @return builder
         * 
         */
        public Builder queueIndexEmbedMsgsBelow(Integer queueIndexEmbedMsgsBelow) {
            return queueIndexEmbedMsgsBelow(Output.of(queueIndexEmbedMsgsBelow));
        }

        /**
         * @param sleep Configurable sleep time in seconds between retries for RabbitMQ configuration. Default set to 60 seconds.
         * 
         * @return builder
         * 
         */
        public Builder sleep(@Nullable Output<Integer> sleep) {
            $.sleep = sleep;
            return this;
        }

        /**
         * @param sleep Configurable sleep time in seconds between retries for RabbitMQ configuration. Default set to 60 seconds.
         * 
         * @return builder
         * 
         */
        public Builder sleep(Integer sleep) {
            return sleep(Output.of(sleep));
        }

        /**
         * @param timeout Configurable timeout time in seconds for RabbitMQ configuration. Default set to 3600 seconds.
         * 
         * @return builder
         * 
         */
        public Builder timeout(@Nullable Output<Integer> timeout) {
            $.timeout = timeout;
            return this;
        }

        /**
         * @param timeout Configurable timeout time in seconds for RabbitMQ configuration. Default set to 3600 seconds.
         * 
         * @return builder
         * 
         */
        public Builder timeout(Integer timeout) {
            return timeout(Output.of(timeout));
        }

        /**
         * @param vmMemoryHighWatermark When the server will enter memory based flow-control as relative to the maximum available memory.
         * 
         * @return builder
         * 
         */
        public Builder vmMemoryHighWatermark(@Nullable Output<Double> vmMemoryHighWatermark) {
            $.vmMemoryHighWatermark = vmMemoryHighWatermark;
            return this;
        }

        /**
         * @param vmMemoryHighWatermark When the server will enter memory based flow-control as relative to the maximum available memory.
         * 
         * @return builder
         * 
         */
        public Builder vmMemoryHighWatermark(Double vmMemoryHighWatermark) {
            return vmMemoryHighWatermark(Output.of(vmMemoryHighWatermark));
        }

        public RabbitConfigurationArgs build() {
            if ($.instanceId == null) {
                throw new MissingRequiredPropertyException("RabbitConfigurationArgs", "instanceId");
            }
            return $;
        }
    }

}