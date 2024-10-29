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


public final class GetPluginsCommunityArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetPluginsCommunityArgs Empty = new GetPluginsCommunityArgs();

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

    @Import(name="sleep")
    private @Nullable Output<Integer> sleep;

    public Optional<Output<Integer>> sleep() {
        return Optional.ofNullable(this.sleep);
    }

    @Import(name="timeout")
    private @Nullable Output<Integer> timeout;

    public Optional<Output<Integer>> timeout() {
        return Optional.ofNullable(this.timeout);
    }

    private GetPluginsCommunityArgs() {}

    private GetPluginsCommunityArgs(GetPluginsCommunityArgs $) {
        this.instanceId = $.instanceId;
        this.sleep = $.sleep;
        this.timeout = $.timeout;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetPluginsCommunityArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetPluginsCommunityArgs $;

        public Builder() {
            $ = new GetPluginsCommunityArgs();
        }

        public Builder(GetPluginsCommunityArgs defaults) {
            $ = new GetPluginsCommunityArgs(Objects.requireNonNull(defaults));
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

        public Builder sleep(@Nullable Output<Integer> sleep) {
            $.sleep = sleep;
            return this;
        }

        public Builder sleep(Integer sleep) {
            return sleep(Output.of(sleep));
        }

        public Builder timeout(@Nullable Output<Integer> timeout) {
            $.timeout = timeout;
            return this;
        }

        public Builder timeout(Integer timeout) {
            return timeout(Output.of(timeout));
        }

        public GetPluginsCommunityArgs build() {
            if ($.instanceId == null) {
                throw new MissingRequiredPropertyException("GetPluginsCommunityArgs", "instanceId");
            }
            return $;
        }
    }

}
