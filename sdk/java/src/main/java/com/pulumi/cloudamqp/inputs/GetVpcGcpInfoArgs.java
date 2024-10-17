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
     * ***Deprecated: Changed from required to optional in v1.16.0 will be removed in next major version (v2.0)***
     * 
     */
    @Import(name="instanceId")
    private @Nullable Output<Integer> instanceId;

    /**
     * @return The CloudAMQP instance identifier.
     * 
     * ***Deprecated: Changed from required to optional in v1.16.0 will be removed in next major version (v2.0)***
     * 
     */
    public Optional<Output<Integer>> instanceId() {
        return Optional.ofNullable(this.instanceId);
    }

    /**
     * Configurable sleep time (seconds) between retries when reading peering. Default set to 10 seconds.
     * 
     */
    @Import(name="sleep")
    private @Nullable Output<Integer> sleep;

    /**
     * @return Configurable sleep time (seconds) between retries when reading peering. Default set to 10 seconds.
     * 
     */
    public Optional<Output<Integer>> sleep() {
        return Optional.ofNullable(this.sleep);
    }

    /**
     * Configurable timeout time (seconds) before retries times out. Default set to 1800 seconds.
     * 
     */
    @Import(name="timeout")
    private @Nullable Output<Integer> timeout;

    /**
     * @return Configurable timeout time (seconds) before retries times out. Default set to 1800 seconds.
     * 
     */
    public Optional<Output<Integer>> timeout() {
        return Optional.ofNullable(this.timeout);
    }

    /**
     * The managed VPC identifier.
     * 
     * ***Note: Added as optional in version v1.16.0 and will be required in next major version (v2.0)***
     * 
     */
    @Import(name="vpcId")
    private @Nullable Output<String> vpcId;

    /**
     * @return The managed VPC identifier.
     * 
     * ***Note: Added as optional in version v1.16.0 and will be required in next major version (v2.0)***
     * 
     */
    public Optional<Output<String>> vpcId() {
        return Optional.ofNullable(this.vpcId);
    }

    private GetVpcGcpInfoArgs() {}

    private GetVpcGcpInfoArgs(GetVpcGcpInfoArgs $) {
        this.instanceId = $.instanceId;
        this.sleep = $.sleep;
        this.timeout = $.timeout;
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
         * ***Deprecated: Changed from required to optional in v1.16.0 will be removed in next major version (v2.0)***
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
         * ***Deprecated: Changed from required to optional in v1.16.0 will be removed in next major version (v2.0)***
         * 
         * @return builder
         * 
         */
        public Builder instanceId(Integer instanceId) {
            return instanceId(Output.of(instanceId));
        }

        /**
         * @param sleep Configurable sleep time (seconds) between retries when reading peering. Default set to 10 seconds.
         * 
         * @return builder
         * 
         */
        public Builder sleep(@Nullable Output<Integer> sleep) {
            $.sleep = sleep;
            return this;
        }

        /**
         * @param sleep Configurable sleep time (seconds) between retries when reading peering. Default set to 10 seconds.
         * 
         * @return builder
         * 
         */
        public Builder sleep(Integer sleep) {
            return sleep(Output.of(sleep));
        }

        /**
         * @param timeout Configurable timeout time (seconds) before retries times out. Default set to 1800 seconds.
         * 
         * @return builder
         * 
         */
        public Builder timeout(@Nullable Output<Integer> timeout) {
            $.timeout = timeout;
            return this;
        }

        /**
         * @param timeout Configurable timeout time (seconds) before retries times out. Default set to 1800 seconds.
         * 
         * @return builder
         * 
         */
        public Builder timeout(Integer timeout) {
            return timeout(Output.of(timeout));
        }

        /**
         * @param vpcId The managed VPC identifier.
         * 
         * ***Note: Added as optional in version v1.16.0 and will be required in next major version (v2.0)***
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
         * ***Note: Added as optional in version v1.16.0 and will be required in next major version (v2.0)***
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