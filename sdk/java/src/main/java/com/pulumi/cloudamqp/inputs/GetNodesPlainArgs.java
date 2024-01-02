// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.cloudamqp.inputs;

import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Integer;
import java.util.Objects;


public final class GetNodesPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetNodesPlainArgs Empty = new GetNodesPlainArgs();

    /**
     * The CloudAMQP instance identifier.
     * 
     */
    @Import(name="instanceId", required=true)
    private Integer instanceId;

    /**
     * @return The CloudAMQP instance identifier.
     * 
     */
    public Integer instanceId() {
        return this.instanceId;
    }

    private GetNodesPlainArgs() {}

    private GetNodesPlainArgs(GetNodesPlainArgs $) {
        this.instanceId = $.instanceId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetNodesPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetNodesPlainArgs $;

        public Builder() {
            $ = new GetNodesPlainArgs();
        }

        public Builder(GetNodesPlainArgs defaults) {
            $ = new GetNodesPlainArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param instanceId The CloudAMQP instance identifier.
         * 
         * @return builder
         * 
         */
        public Builder instanceId(Integer instanceId) {
            $.instanceId = instanceId;
            return this;
        }

        public GetNodesPlainArgs build() {
            if ($.instanceId == null) {
                throw new MissingRequiredPropertyException("GetNodesPlainArgs", "instanceId");
            }
            return $;
        }
    }

}
