// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.cloudamqp.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Integer;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetPluginsArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetPluginsArgs Empty = new GetPluginsArgs();

    /**
     * The CloudAMQP instance identifier.
     * 
     */
    @Import(name="instanceId", required=true)
    private Output<Integer> instanceId;

    /**
     * @return The CloudAMQP instance identifier.
     * 
     */
    public Output<Integer> instanceId() {
        return this.instanceId;
    }

    /**
     * Configurable sleep time (seconds) for retries when requesting
     * information about plugins. Default set to 10 seconds.
     * 
     */
    @Import(name="sleep")
    private @Nullable Output<Integer> sleep;

    /**
     * @return Configurable sleep time (seconds) for retries when requesting
     * information about plugins. Default set to 10 seconds.
     * 
     */
    public Optional<Output<Integer>> sleep() {
        return Optional.ofNullable(this.sleep);
    }

    /**
     * Configurable timeout time (seconds) for retries when requesting
     * information about plugins. Default set to 1800 seconds.
     * 
     */
    @Import(name="timeout")
    private @Nullable Output<Integer> timeout;

    /**
     * @return Configurable timeout time (seconds) for retries when requesting
     * information about plugins. Default set to 1800 seconds.
     * 
     */
    public Optional<Output<Integer>> timeout() {
        return Optional.ofNullable(this.timeout);
    }

    private GetPluginsArgs() {}

    private GetPluginsArgs(GetPluginsArgs $) {
        this.instanceId = $.instanceId;
        this.sleep = $.sleep;
        this.timeout = $.timeout;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetPluginsArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetPluginsArgs $;

        public Builder() {
            $ = new GetPluginsArgs();
        }

        public Builder(GetPluginsArgs defaults) {
            $ = new GetPluginsArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param instanceId The CloudAMQP instance identifier.
         * 
         * @return builder
         * 
         */
        public Builder instanceId(Output<Integer> instanceId) {
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
         * @param sleep Configurable sleep time (seconds) for retries when requesting
         * information about plugins. Default set to 10 seconds.
         * 
         * @return builder
         * 
         */
        public Builder sleep(@Nullable Output<Integer> sleep) {
            $.sleep = sleep;
            return this;
        }

        /**
         * @param sleep Configurable sleep time (seconds) for retries when requesting
         * information about plugins. Default set to 10 seconds.
         * 
         * @return builder
         * 
         */
        public Builder sleep(Integer sleep) {
            return sleep(Output.of(sleep));
        }

        /**
         * @param timeout Configurable timeout time (seconds) for retries when requesting
         * information about plugins. Default set to 1800 seconds.
         * 
         * @return builder
         * 
         */
        public Builder timeout(@Nullable Output<Integer> timeout) {
            $.timeout = timeout;
            return this;
        }

        /**
         * @param timeout Configurable timeout time (seconds) for retries when requesting
         * information about plugins. Default set to 1800 seconds.
         * 
         * @return builder
         * 
         */
        public Builder timeout(Integer timeout) {
            return timeout(Output.of(timeout));
        }

        public GetPluginsArgs build() {
            if ($.instanceId == null) {
                throw new MissingRequiredPropertyException("GetPluginsArgs", "instanceId");
            }
            return $;
        }
    }

}
