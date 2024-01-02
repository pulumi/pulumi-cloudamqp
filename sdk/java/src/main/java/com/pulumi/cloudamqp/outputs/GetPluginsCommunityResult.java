// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.cloudamqp.outputs;

import com.pulumi.cloudamqp.outputs.GetPluginsCommunityPlugin;
import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;

@CustomType
public final class GetPluginsCommunityResult {
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    private Integer instanceId;
    private List<GetPluginsCommunityPlugin> plugins;

    private GetPluginsCommunityResult() {}
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
    public List<GetPluginsCommunityPlugin> plugins() {
        return this.plugins;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetPluginsCommunityResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String id;
        private Integer instanceId;
        private List<GetPluginsCommunityPlugin> plugins;
        public Builder() {}
        public Builder(GetPluginsCommunityResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.id = defaults.id;
    	      this.instanceId = defaults.instanceId;
    	      this.plugins = defaults.plugins;
        }

        @CustomType.Setter
        public Builder id(String id) {
            if (id == null) {
              throw new MissingRequiredPropertyException("GetPluginsCommunityResult", "id");
            }
            this.id = id;
            return this;
        }
        @CustomType.Setter
        public Builder instanceId(Integer instanceId) {
            if (instanceId == null) {
              throw new MissingRequiredPropertyException("GetPluginsCommunityResult", "instanceId");
            }
            this.instanceId = instanceId;
            return this;
        }
        @CustomType.Setter
        public Builder plugins(List<GetPluginsCommunityPlugin> plugins) {
            if (plugins == null) {
              throw new MissingRequiredPropertyException("GetPluginsCommunityResult", "plugins");
            }
            this.plugins = plugins;
            return this;
        }
        public Builder plugins(GetPluginsCommunityPlugin... plugins) {
            return plugins(List.of(plugins));
        }
        public GetPluginsCommunityResult build() {
            final var _resultValue = new GetPluginsCommunityResult();
            _resultValue.id = id;
            _resultValue.instanceId = instanceId;
            _resultValue.plugins = plugins;
            return _resultValue;
        }
    }
}
