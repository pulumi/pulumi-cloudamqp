// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.cloudamqp.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetVpcGcpInfoArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetVpcGcpInfoArgs Empty = new GetVpcGcpInfoArgs();

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
     * The managed VPC identifier.
     * 
     */
    @Import(name="vpcId")
    private @Nullable Output<String> vpcId;

    /**
     * @return The managed VPC identifier.
     * 
     */
    public Optional<Output<String>> vpcId() {
        return Optional.ofNullable(this.vpcId);
    }

    private GetVpcGcpInfoArgs() {}

    private GetVpcGcpInfoArgs(GetVpcGcpInfoArgs $) {
        this.instanceId = $.instanceId;
        this.vpcId = $.vpcId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetVpcGcpInfoArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetVpcGcpInfoArgs $;

        public Builder() {
            $ = new GetVpcGcpInfoArgs();
        }

        public Builder(GetVpcGcpInfoArgs defaults) {
            $ = new GetVpcGcpInfoArgs(Objects.requireNonNull(defaults));
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
         * @param vpcId The managed VPC identifier.
         * 
         * @return builder
         * 
         */
        public Builder vpcId(@Nullable Output<String> vpcId) {
            $.vpcId = vpcId;
            return this;
        }

        /**
         * @param vpcId The managed VPC identifier.
         * 
         * @return builder
         * 
         */
        public Builder vpcId(String vpcId) {
            return vpcId(Output.of(vpcId));
        }

        public GetVpcGcpInfoArgs build() {
            return $;
        }
    }

}
