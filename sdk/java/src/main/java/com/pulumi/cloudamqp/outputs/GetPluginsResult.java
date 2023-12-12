// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.cloudamqp.outputs;

import com.pulumi.cloudamqp.outputs.GetPluginsPlugin;
import com.pulumi.core.annotations.CustomType;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;

@CustomType
public final class GetPluginsResult {
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    private Integer instanceId;
    private List<GetPluginsPlugin> plugins;

    private GetPluginsResult() {}
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
    public List<GetPluginsPlugin> plugins() {
        return this.plugins;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetPluginsResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String id;
        private Integer instanceId;
        private List<GetPluginsPlugin> plugins;
        public Builder() {}
        public Builder(GetPluginsResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.id = defaults.id;
    	      this.instanceId = defaults.instanceId;
    	      this.plugins = defaults.plugins;
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
        public Builder plugins(List<GetPluginsPlugin> plugins) {
            this.plugins = Objects.requireNonNull(plugins);
            return this;
        }
        public Builder plugins(GetPluginsPlugin... plugins) {
            return plugins(List.of(plugins));
        }
        public GetPluginsResult build() {
            final var _resultValue = new GetPluginsResult();
            _resultValue.id = id;
            _resultValue.instanceId = instanceId;
            _resultValue.plugins = plugins;
            return _resultValue;
        }
    }
}
