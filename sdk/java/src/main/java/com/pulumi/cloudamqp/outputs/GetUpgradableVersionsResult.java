// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.cloudamqp.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetUpgradableVersionsResult {
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    private Integer instanceId;
    private String newErlangVersion;
    private String newRabbitmqVersion;

    private GetUpgradableVersionsResult() {}
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    public String id() {
        return this.id;
    }
    public Integer instanceId() {
        return this.instanceId;
    }
    public String newErlangVersion() {
        return this.newErlangVersion;
    }
    public String newRabbitmqVersion() {
        return this.newRabbitmqVersion;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetUpgradableVersionsResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String id;
        private Integer instanceId;
        private String newErlangVersion;
        private String newRabbitmqVersion;
        public Builder() {}
        public Builder(GetUpgradableVersionsResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.id = defaults.id;
    	      this.instanceId = defaults.instanceId;
    	      this.newErlangVersion = defaults.newErlangVersion;
    	      this.newRabbitmqVersion = defaults.newRabbitmqVersion;
        }

        @CustomType.Setter
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        @CustomType.Setter
        public Builder instanceId(Integer instanceId) {
            this.instanceId = Objects.requireNonNull(instanceId);
            return this;
        }
        @CustomType.Setter
        public Builder newErlangVersion(String newErlangVersion) {
            this.newErlangVersion = Objects.requireNonNull(newErlangVersion);
            return this;
        }
        @CustomType.Setter
        public Builder newRabbitmqVersion(String newRabbitmqVersion) {
            this.newRabbitmqVersion = Objects.requireNonNull(newRabbitmqVersion);
            return this;
        }
        public GetUpgradableVersionsResult build() {
            final var o = new GetUpgradableVersionsResult();
            o.id = id;
            o.instanceId = instanceId;
            o.newErlangVersion = newErlangVersion;
            o.newRabbitmqVersion = newRabbitmqVersion;
            return o;
        }
    }
}
