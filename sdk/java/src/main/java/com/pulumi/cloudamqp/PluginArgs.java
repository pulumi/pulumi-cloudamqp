// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.cloudamqp;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class PluginArgs extends com.pulumi.resources.ResourceArgs {

    public static final PluginArgs Empty = new PluginArgs();

    /**
     * Enable or disable the plugins.
     * 
     */
    @Import(name="enabled", required=true)
    private Output<Boolean> enabled;

    /**
     * @return Enable or disable the plugins.
     * 
     */
    public Output<Boolean> enabled() {
        return this.enabled;
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
     * The name of the Rabbit MQ plugin.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return The name of the Rabbit MQ plugin.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * Configurable sleep time (seconds) for retries when requesting
     * information about plugins. Default set to 10 seconds.
     * 
     * ***Note:*** Available from [v1.29.0]
     * 
     */
    @Import(name="sleep")
    private @Nullable Output<Integer> sleep;

    /**
     * @return Configurable sleep time (seconds) for retries when requesting
     * information about plugins. Default set to 10 seconds.
     * 
     * ***Note:*** Available from [v1.29.0]
     * 
     */
    public Optional<Output<Integer>> sleep() {
        return Optional.ofNullable(this.sleep);
    }

    /**
     * Configurable timeout time (seconds) for retries when requesting
     * information about plugins. Default set to 1800 seconds.
     * 
     * ***Note:*** Available from [v1.29.0]
     * 
     */
    @Import(name="timeout")
    private @Nullable Output<Integer> timeout;

    /**
     * @return Configurable timeout time (seconds) for retries when requesting
     * information about plugins. Default set to 1800 seconds.
     * 
     * ***Note:*** Available from [v1.29.0]
     * 
     */
    public Optional<Output<Integer>> timeout() {
        return Optional.ofNullable(this.timeout);
    }

    private PluginArgs() {}

    private PluginArgs(PluginArgs $) {
        this.enabled = $.enabled;
        this.instanceId = $.instanceId;
        this.name = $.name;
        this.sleep = $.sleep;
        this.timeout = $.timeout;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(PluginArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private PluginArgs $;

        public Builder() {
            $ = new PluginArgs();
        }

        public Builder(PluginArgs defaults) {
            $ = new PluginArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param enabled Enable or disable the plugins.
         * 
         * @return builder
         * 
         */
        public Builder enabled(Output<Boolean> enabled) {
            $.enabled = enabled;
            return this;
        }

        /**
         * @param enabled Enable or disable the plugins.
         * 
         * @return builder
         * 
         */
        public Builder enabled(Boolean enabled) {
            return enabled(Output.of(enabled));
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
         * @param name The name of the Rabbit MQ plugin.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name The name of the Rabbit MQ plugin.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param sleep Configurable sleep time (seconds) for retries when requesting
         * information about plugins. Default set to 10 seconds.
         * 
         * ***Note:*** Available from [v1.29.0]
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
         * ***Note:*** Available from [v1.29.0]
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
         * ***Note:*** Available from [v1.29.0]
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
         * ***Note:*** Available from [v1.29.0]
         * 
         * @return builder
         * 
         */
        public Builder timeout(Integer timeout) {
            return timeout(Output.of(timeout));
        }

        public PluginArgs build() {
            if ($.enabled == null) {
                throw new MissingRequiredPropertyException("PluginArgs", "enabled");
            }
            if ($.instanceId == null) {
                throw new MissingRequiredPropertyException("PluginArgs", "instanceId");
            }
            return $;
        }
    }

}
