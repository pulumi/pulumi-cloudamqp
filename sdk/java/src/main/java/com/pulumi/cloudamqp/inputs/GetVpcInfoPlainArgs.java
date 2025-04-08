// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.cloudamqp.inputs;

import com.pulumi.core.annotations.Import;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetVpcInfoPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetVpcInfoPlainArgs Empty = new GetVpcInfoPlainArgs();

    /**
     * The CloudAMQP instance identifier.
     * 
     * ***Deprecated:*** from [v1.16.0], will be removed in next major version (v2.0)
     * 
     */
    @Import(name="instanceId")
    private @Nullable Integer instanceId;

    /**
     * @return The CloudAMQP instance identifier.
     * 
     * ***Deprecated:*** from [v1.16.0], will be removed in next major version (v2.0)
     * 
     */
    public Optional<Integer> instanceId() {
        return Optional.ofNullable(this.instanceId);
    }

    /**
     * The managed VPC identifier.
     * 
     * ***Note:*** Available from [v1.16.0], will be removed in next major version (v2.0)
     * 
     */
    @Import(name="vpcId")
    private @Nullable String vpcId;

    /**
     * @return The managed VPC identifier.
     * 
     * ***Note:*** Available from [v1.16.0], will be removed in next major version (v2.0)
     * 
     */
    public Optional<String> vpcId() {
        return Optional.ofNullable(this.vpcId);
    }

    private GetVpcInfoPlainArgs() {}

    private GetVpcInfoPlainArgs(GetVpcInfoPlainArgs $) {
        this.instanceId = $.instanceId;
        this.vpcId = $.vpcId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetVpcInfoPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetVpcInfoPlainArgs $;

        public Builder() {
            $ = new GetVpcInfoPlainArgs();
        }

        public Builder(GetVpcInfoPlainArgs defaults) {
            $ = new GetVpcInfoPlainArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param instanceId The CloudAMQP instance identifier.
         * 
         * ***Deprecated:*** from [v1.16.0], will be removed in next major version (v2.0)
         * 
         * @return builder
         * 
         */
        public Builder instanceId(@Nullable Integer instanceId) {
            $.instanceId = instanceId;
            return this;
        }

        /**
         * @param vpcId The managed VPC identifier.
         * 
         * ***Note:*** Available from [v1.16.0], will be removed in next major version (v2.0)
         * 
         * @return builder
         * 
         */
        public Builder vpcId(@Nullable String vpcId) {
            $.vpcId = vpcId;
            return this;
        }

        public GetVpcInfoPlainArgs build() {
            return $;
        }
    }

}
