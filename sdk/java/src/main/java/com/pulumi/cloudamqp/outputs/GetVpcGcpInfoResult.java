// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.cloudamqp.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
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
    /**
     * @return The name of the VPC.
     * 
     */
    private String name;
    /**
     * @return VPC network uri.
     * 
     */
    private String network;
    private @Nullable Integer sleep;
    private @Nullable Integer timeout;
    private @Nullable String vpcId;
    /**
     * @return Dedicated VPC subnet.
     * 
     */
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
    /**
     * @return The name of the VPC.
     * 
     */
    public String name() {
        return this.name;
    }
    /**
     * @return VPC network uri.
     * 
     */
    public String network() {
        return this.network;
    }
    public Optional<Integer> sleep() {
        return Optional.ofNullable(this.sleep);
    }
    public Optional<Integer> timeout() {
        return Optional.ofNullable(this.timeout);
    }
    public Optional<String> vpcId() {
        return Optional.ofNullable(this.vpcId);
    }
    /**
     * @return Dedicated VPC subnet.
     * 
     */
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
        private @Nullable Integer sleep;
        private @Nullable Integer timeout;
        private @Nullable String vpcId;
        private String vpcSubnet;
        public Builder() {}
        public Builder(GetVpcGcpInfoResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.id = defaults.id;
    	      this.instanceId = defaults.instanceId;
    	      this.name = defaults.name;
    	      this.network = defaults.network;
    	      this.sleep = defaults.sleep;
    	      this.timeout = defaults.timeout;
    	      this.vpcId = defaults.vpcId;
    	      this.vpcSubnet = defaults.vpcSubnet;
        }

        @CustomType.Setter
        public Builder id(String id) {
            if (id == null) {
              throw new MissingRequiredPropertyException("GetVpcGcpInfoResult", "id");
            }
            this.id = id;
            return this;
        }
        @CustomType.Setter
        public Builder instanceId(@Nullable Integer instanceId) {

            this.instanceId = instanceId;
            return this;
        }
        @CustomType.Setter
        public Builder name(String name) {
            if (name == null) {
              throw new MissingRequiredPropertyException("GetVpcGcpInfoResult", "name");
            }
            this.name = name;
            return this;
        }
        @CustomType.Setter
        public Builder network(String network) {
            if (network == null) {
              throw new MissingRequiredPropertyException("GetVpcGcpInfoResult", "network");
            }
            this.network = network;
            return this;
        }
        @CustomType.Setter
        public Builder sleep(@Nullable Integer sleep) {

            this.sleep = sleep;
            return this;
        }
        @CustomType.Setter
        public Builder timeout(@Nullable Integer timeout) {

            this.timeout = timeout;
            return this;
        }
        @CustomType.Setter
        public Builder vpcId(@Nullable String vpcId) {

            this.vpcId = vpcId;
            return this;
        }
        @CustomType.Setter
        public Builder vpcSubnet(String vpcSubnet) {
            if (vpcSubnet == null) {
              throw new MissingRequiredPropertyException("GetVpcGcpInfoResult", "vpcSubnet");
            }
            this.vpcSubnet = vpcSubnet;
            return this;
        }
        public GetVpcGcpInfoResult build() {
            final var _resultValue = new GetVpcGcpInfoResult();
            _resultValue.id = id;
            _resultValue.instanceId = instanceId;
            _resultValue.name = name;
            _resultValue.network = network;
            _resultValue.sleep = sleep;
            _resultValue.timeout = timeout;
            _resultValue.vpcId = vpcId;
            _resultValue.vpcSubnet = vpcSubnet;
            return _resultValue;
        }
    }
}
