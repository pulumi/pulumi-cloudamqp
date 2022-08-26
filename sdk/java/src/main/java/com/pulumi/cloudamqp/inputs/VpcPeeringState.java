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


public final class VpcPeeringState extends com.pulumi.resources.ResourceArgs {

    public static final VpcPeeringState Empty = new VpcPeeringState();

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
     * Peering identifier created by AW peering request.
     * 
     */
    @Import(name="peeringId")
    private @Nullable Output<String> peeringId;

    /**
     * @return Peering identifier created by AW peering request.
     * 
     */
    public Optional<Output<String>> peeringId() {
        return Optional.ofNullable(this.peeringId);
    }

    /**
     * Configurable sleep time (seconds) between retries for accepting or removing peering. Default set to 60 seconds.
     * 
     */
    @Import(name="sleep")
    private @Nullable Output<Integer> sleep;

    /**
     * @return Configurable sleep time (seconds) between retries for accepting or removing peering. Default set to 60 seconds.
     * 
     */
    public Optional<Output<Integer>> sleep() {
        return Optional.ofNullable(this.sleep);
    }

    /**
     * VPC peering status
     * 
     */
    @Import(name="status")
    private @Nullable Output<String> status;

    /**
     * @return VPC peering status
     * 
     */
    public Optional<Output<String>> status() {
        return Optional.ofNullable(this.status);
    }

    /**
     * - Configurable timeout time (seconds) for accepting or removing peering. Default set to 3600 seconds.
     * 
     */
    @Import(name="timeout")
    private @Nullable Output<Integer> timeout;

    /**
     * @return - Configurable timeout time (seconds) for accepting or removing peering. Default set to 3600 seconds.
     * 
     */
    public Optional<Output<Integer>> timeout() {
        return Optional.ofNullable(this.timeout);
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

    private VpcPeeringState() {}

    private VpcPeeringState(VpcPeeringState $) {
        this.instanceId = $.instanceId;
        this.peeringId = $.peeringId;
        this.sleep = $.sleep;
        this.status = $.status;
        this.timeout = $.timeout;
        this.vpcId = $.vpcId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(VpcPeeringState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private VpcPeeringState $;

        public Builder() {
            $ = new VpcPeeringState();
        }

        public Builder(VpcPeeringState defaults) {
            $ = new VpcPeeringState(Objects.requireNonNull(defaults));
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
         * @param peeringId Peering identifier created by AW peering request.
         * 
         * @return builder
         * 
         */
        public Builder peeringId(@Nullable Output<String> peeringId) {
            $.peeringId = peeringId;
            return this;
        }

        /**
         * @param peeringId Peering identifier created by AW peering request.
         * 
         * @return builder
         * 
         */
        public Builder peeringId(String peeringId) {
            return peeringId(Output.of(peeringId));
        }

        /**
         * @param sleep Configurable sleep time (seconds) between retries for accepting or removing peering. Default set to 60 seconds.
         * 
         * @return builder
         * 
         */
        public Builder sleep(@Nullable Output<Integer> sleep) {
            $.sleep = sleep;
            return this;
        }

        /**
         * @param sleep Configurable sleep time (seconds) between retries for accepting or removing peering. Default set to 60 seconds.
         * 
         * @return builder
         * 
         */
        public Builder sleep(Integer sleep) {
            return sleep(Output.of(sleep));
        }

        /**
         * @param status VPC peering status
         * 
         * @return builder
         * 
         */
        public Builder status(@Nullable Output<String> status) {
            $.status = status;
            return this;
        }

        /**
         * @param status VPC peering status
         * 
         * @return builder
         * 
         */
        public Builder status(String status) {
            return status(Output.of(status));
        }

        /**
         * @param timeout - Configurable timeout time (seconds) for accepting or removing peering. Default set to 3600 seconds.
         * 
         * @return builder
         * 
         */
        public Builder timeout(@Nullable Output<Integer> timeout) {
            $.timeout = timeout;
            return this;
        }

        /**
         * @param timeout - Configurable timeout time (seconds) for accepting or removing peering. Default set to 3600 seconds.
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

        public VpcPeeringState build() {
            return $;
        }
    }

}
