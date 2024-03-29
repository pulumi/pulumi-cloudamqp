// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.cloudamqp.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Integer;
import java.util.Objects;


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

    private GetPluginsCommunityArgs() {}

    private GetPluginsCommunityArgs(GetPluginsCommunityArgs $) {
        this.instanceId = $.instanceId;
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

        public GetPluginsCommunityArgs build() {
            if ($.instanceId == null) {
                throw new MissingRequiredPropertyException("GetPluginsCommunityArgs", "instanceId");
            }
            return $;
        }
    }

}
