// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.cloudamqp.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class GetVpcGcpInfoResult {
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    private @Nullable Integer instanceId;
    private String name;
    private String network;
    private @Nullable String vpcId;
    private String vpcSubnet;

    private GetVpcGcpInfoResult() {}
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    public String id() {
        return this.id;
    }
    public Optional<Integer> instanceId() {
        return Optional.ofNullable(this.instanceId);
    }
    public String name() {
        return this.name;
    }
    public String network() {
        return this.network;
    }
    public Optional<String> vpcId() {
        return Optional.ofNullable(this.vpcId);
    }
    public String vpcSubnet() {
        return this.vpcSubnet;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetVpcGcpInfoResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String id;
        private @Nullable Integer instanceId;
        private String name;
        private String network;
        private @Nullable String vpcId;
        private String vpcSubnet;
        public Builder() {}
        public Builder(GetVpcGcpInfoResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.id = defaults.id;
    	      this.instanceId = defaults.instanceId;
    	      this.name = defaults.name;
    	      this.network = defaults.network;
    	      this.vpcId = defaults.vpcId;
    	      this.vpcSubnet = defaults.vpcSubnet;
        }

        @CustomType.Setter
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        @CustomType.Setter
        public Builder instanceId(@Nullable Integer instanceId) {
            this.instanceId = instanceId;
            return this;
        }
        @CustomType.Setter
        public Builder name(String name) {
            this.name = Objects.requireNonNull(name);
            return this;
        }
        @CustomType.Setter
        public Builder network(String network) {
            this.network = Objects.requireNonNull(network);
            return this;
        }
        @CustomType.Setter
        public Builder vpcId(@Nullable String vpcId) {
            this.vpcId = vpcId;
            return this;
        }
        @CustomType.Setter
        public Builder vpcSubnet(String vpcSubnet) {
            this.vpcSubnet = Objects.requireNonNull(vpcSubnet);
            return this;
        }
        public GetVpcGcpInfoResult build() {
            final var o = new GetVpcGcpInfoResult();
            o.id = id;
            o.instanceId = instanceId;
            o.name = name;
            o.network = network;
            o.vpcId = vpcId;
            o.vpcSubnet = vpcSubnet;
            return o;
        }
    }
}