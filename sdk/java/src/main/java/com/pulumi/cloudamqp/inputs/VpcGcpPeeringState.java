// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.cloudamqp.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class VpcGcpPeeringState extends com.pulumi.resources.ResourceArgs {

    public static final VpcGcpPeeringState Empty = new VpcGcpPeeringState();

    /**
     * VPC peering auto created routes
     * 
     */
    @Import(name="autoCreateRoutes")
    private @Nullable Output<Boolean> autoCreateRoutes;

    /**
     * @return VPC peering auto created routes
     * 
     */
    public Optional<Output<Boolean>> autoCreateRoutes() {
        return Optional.ofNullable(this.autoCreateRoutes);
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
     * Network uri of the VPC network to which you will peer with.
     * 
     */
    @Import(name="peerNetworkUri")
    private @Nullable Output<String> peerNetworkUri;

    /**
     * @return Network uri of the VPC network to which you will peer with.
     * 
     */
    public Optional<Output<String>> peerNetworkUri() {
        return Optional.ofNullable(this.peerNetworkUri);
    }

    /**
     * VPC peering state
     * 
     */
    @Import(name="state")
    private @Nullable Output<String> state;

    /**
     * @return VPC peering state
     * 
     */
    public Optional<Output<String>> state() {
        return Optional.ofNullable(this.state);
    }

    /**
     * VPC peering state details
     * 
     */
    @Import(name="stateDetails")
    private @Nullable Output<String> stateDetails;

    /**
     * @return VPC peering state details
     * 
     */
    public Optional<Output<String>> stateDetails() {
        return Optional.ofNullable(this.stateDetails);
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

    private VpcGcpPeeringState() {}

    private VpcGcpPeeringState(VpcGcpPeeringState $) {
        this.autoCreateRoutes = $.autoCreateRoutes;
        this.instanceId = $.instanceId;
        this.peerNetworkUri = $.peerNetworkUri;
        this.state = $.state;
        this.stateDetails = $.stateDetails;
        this.vpcId = $.vpcId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(VpcGcpPeeringState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private VpcGcpPeeringState $;

        public Builder() {
            $ = new VpcGcpPeeringState();
        }

        public Builder(VpcGcpPeeringState defaults) {
            $ = new VpcGcpPeeringState(Objects.requireNonNull(defaults));
        }

        /**
         * @param autoCreateRoutes VPC peering auto created routes
         * 
         * @return builder
         * 
         */
        public Builder autoCreateRoutes(@Nullable Output<Boolean> autoCreateRoutes) {
            $.autoCreateRoutes = autoCreateRoutes;
            return this;
        }

        /**
         * @param autoCreateRoutes VPC peering auto created routes
         * 
         * @return builder
         * 
         */
        public Builder autoCreateRoutes(Boolean autoCreateRoutes) {
            return autoCreateRoutes(Output.of(autoCreateRoutes));
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
         * @param peerNetworkUri Network uri of the VPC network to which you will peer with.
         * 
         * @return builder
         * 
         */
        public Builder peerNetworkUri(@Nullable Output<String> peerNetworkUri) {
            $.peerNetworkUri = peerNetworkUri;
            return this;
        }

        /**
         * @param peerNetworkUri Network uri of the VPC network to which you will peer with.
         * 
         * @return builder
         * 
         */
        public Builder peerNetworkUri(String peerNetworkUri) {
            return peerNetworkUri(Output.of(peerNetworkUri));
        }

        /**
         * @param state VPC peering state
         * 
         * @return builder
         * 
         */
        public Builder state(@Nullable Output<String> state) {
            $.state = state;
            return this;
        }

        /**
         * @param state VPC peering state
         * 
         * @return builder
         * 
         */
        public Builder state(String state) {
            return state(Output.of(state));
        }

        /**
         * @param stateDetails VPC peering state details
         * 
         * @return builder
         * 
         */
        public Builder stateDetails(@Nullable Output<String> stateDetails) {
            $.stateDetails = stateDetails;
            return this;
        }

        /**
         * @param stateDetails VPC peering state details
         * 
         * @return builder
         * 
         */
        public Builder stateDetails(String stateDetails) {
            return stateDetails(Output.of(stateDetails));
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

        public VpcGcpPeeringState build() {
            return $;
        }
    }

}