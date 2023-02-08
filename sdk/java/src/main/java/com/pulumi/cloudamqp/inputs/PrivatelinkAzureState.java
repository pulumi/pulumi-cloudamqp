// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.cloudamqp.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class PrivatelinkAzureState extends com.pulumi.resources.ResourceArgs {

    public static final PrivatelinkAzureState Empty = new PrivatelinkAzureState();

    /**
     * Approved subscriptions to access the endpoint service. See format below.
     * 
     */
    @Import(name="approvedSubscriptions")
    private @Nullable Output<List<String>> approvedSubscriptions;

    /**
     * @return Approved subscriptions to access the endpoint service. See format below.
     * 
     */
    public Optional<Output<List<String>>> approvedSubscriptions() {
        return Optional.ofNullable(this.approvedSubscriptions);
    }

    /**
     * The CloudAMQP instance identifier.
     * 
     */
    @Import(name="instanceId")
    private @Nullable Output<Integer> instanceId;

    /**
     * @return The CloudAMQP instance identifier.
     * 
     */
    public Optional<Output<Integer>> instanceId() {
        return Optional.ofNullable(this.instanceId);
    }

    /**
     * Name of the server having the PrivateLink enabled.
     * 
     */
    @Import(name="serverName")
    private @Nullable Output<String> serverName;

    /**
     * @return Name of the server having the PrivateLink enabled.
     * 
     */
    public Optional<Output<String>> serverName() {
        return Optional.ofNullable(this.serverName);
    }

    /**
     * Service name (alias) of the PrivateLink, needed when creating the endpoint.
     * 
     */
    @Import(name="serviceName")
    private @Nullable Output<String> serviceName;

    /**
     * @return Service name (alias) of the PrivateLink, needed when creating the endpoint.
     * 
     */
    public Optional<Output<String>> serviceName() {
        return Optional.ofNullable(this.serviceName);
    }

    /**
     * Configurable sleep time (seconds) when enable PrivateLink. Default set to 60 seconds.
     * 
     */
    @Import(name="sleep")
    private @Nullable Output<Integer> sleep;

    /**
     * @return Configurable sleep time (seconds) when enable PrivateLink. Default set to 60 seconds.
     * 
     */
    public Optional<Output<Integer>> sleep() {
        return Optional.ofNullable(this.sleep);
    }

    /**
     * PrivateLink status [enable, pending, disable]
     * 
     */
    @Import(name="status")
    private @Nullable Output<String> status;

    /**
     * @return PrivateLink status [enable, pending, disable]
     * 
     */
    public Optional<Output<String>> status() {
        return Optional.ofNullable(this.status);
    }

    /**
     * Configurable timeout time (seconds) when enable PrivateLink. Default set to 3600 seconds.
     * 
     */
    @Import(name="timeout")
    private @Nullable Output<Integer> timeout;

    /**
     * @return Configurable timeout time (seconds) when enable PrivateLink. Default set to 3600 seconds.
     * 
     */
    public Optional<Output<Integer>> timeout() {
        return Optional.ofNullable(this.timeout);
    }

    private PrivatelinkAzureState() {}

    private PrivatelinkAzureState(PrivatelinkAzureState $) {
        this.approvedSubscriptions = $.approvedSubscriptions;
        this.instanceId = $.instanceId;
        this.serverName = $.serverName;
        this.serviceName = $.serviceName;
        this.sleep = $.sleep;
        this.status = $.status;
        this.timeout = $.timeout;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(PrivatelinkAzureState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private PrivatelinkAzureState $;

        public Builder() {
            $ = new PrivatelinkAzureState();
        }

        public Builder(PrivatelinkAzureState defaults) {
            $ = new PrivatelinkAzureState(Objects.requireNonNull(defaults));
        }

        /**
         * @param approvedSubscriptions Approved subscriptions to access the endpoint service. See format below.
         * 
         * @return builder
         * 
         */
        public Builder approvedSubscriptions(@Nullable Output<List<String>> approvedSubscriptions) {
            $.approvedSubscriptions = approvedSubscriptions;
            return this;
        }

        /**
         * @param approvedSubscriptions Approved subscriptions to access the endpoint service. See format below.
         * 
         * @return builder
         * 
         */
        public Builder approvedSubscriptions(List<String> approvedSubscriptions) {
            return approvedSubscriptions(Output.of(approvedSubscriptions));
        }

        /**
         * @param approvedSubscriptions Approved subscriptions to access the endpoint service. See format below.
         * 
         * @return builder
         * 
         */
        public Builder approvedSubscriptions(String... approvedSubscriptions) {
            return approvedSubscriptions(List.of(approvedSubscriptions));
        }

        /**
         * @param instanceId The CloudAMQP instance identifier.
         * 
         * @return builder
         * 
         */
        public Builder instanceId(@Nullable Output<Integer> instanceId) {
            $.instanceId = instanceId;
            return this;
        }

        /**
         * @param instanceId The CloudAMQP instance identifier.
         * 
         * @return builder
         * 
         */
        public Builder instanceId(Integer instanceId) {
            return instanceId(Output.of(instanceId));
        }

        /**
         * @param serverName Name of the server having the PrivateLink enabled.
         * 
         * @return builder
         * 
         */
        public Builder serverName(@Nullable Output<String> serverName) {
            $.serverName = serverName;
            return this;
        }

        /**
         * @param serverName Name of the server having the PrivateLink enabled.
         * 
         * @return builder
         * 
         */
        public Builder serverName(String serverName) {
            return serverName(Output.of(serverName));
        }

        /**
         * @param serviceName Service name (alias) of the PrivateLink, needed when creating the endpoint.
         * 
         * @return builder
         * 
         */
        public Builder serviceName(@Nullable Output<String> serviceName) {
            $.serviceName = serviceName;
            return this;
        }

        /**
         * @param serviceName Service name (alias) of the PrivateLink, needed when creating the endpoint.
         * 
         * @return builder
         * 
         */
        public Builder serviceName(String serviceName) {
            return serviceName(Output.of(serviceName));
        }

        /**
         * @param sleep Configurable sleep time (seconds) when enable PrivateLink. Default set to 60 seconds.
         * 
         * @return builder
         * 
         */
        public Builder sleep(@Nullable Output<Integer> sleep) {
            $.sleep = sleep;
            return this;
        }

        /**
         * @param sleep Configurable sleep time (seconds) when enable PrivateLink. Default set to 60 seconds.
         * 
         * @return builder
         * 
         */
        public Builder sleep(Integer sleep) {
            return sleep(Output.of(sleep));
        }

        /**
         * @param status PrivateLink status [enable, pending, disable]
         * 
         * @return builder
         * 
         */
        public Builder status(@Nullable Output<String> status) {
            $.status = status;
            return this;
        }

        /**
         * @param status PrivateLink status [enable, pending, disable]
         * 
         * @return builder
         * 
         */
        public Builder status(String status) {
            return status(Output.of(status));
        }

        /**
         * @param timeout Configurable timeout time (seconds) when enable PrivateLink. Default set to 3600 seconds.
         * 
         * @return builder
         * 
         */
        public Builder timeout(@Nullable Output<Integer> timeout) {
            $.timeout = timeout;
            return this;
        }

        /**
         * @param timeout Configurable timeout time (seconds) when enable PrivateLink. Default set to 3600 seconds.
         * 
         * @return builder
         * 
         */
        public Builder timeout(Integer timeout) {
            return timeout(Output.of(timeout));
        }

        public PrivatelinkAzureState build() {
            return $;
        }
    }

}