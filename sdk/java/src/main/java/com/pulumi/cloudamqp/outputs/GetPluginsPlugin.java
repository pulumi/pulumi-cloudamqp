// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.cloudamqp.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class GetPluginsPlugin {
    private String description;
    private Boolean enabled;
    private String name;
    /**
     * @return Configurable sleep time in seconds between retries for plugins
     * 
     */
    private @Nullable Integer sleep;
    /**
     * @return Configurable timeout time in seconds for plugins
     * 
     */
    private @Nullable Integer timeout;
    private String version;

    private GetPluginsPlugin() {}
    public String description() {
        return this.description;
    }
    public Boolean enabled() {
        return this.enabled;
    }
    public String name() {
        return this.name;
    }
    /**
     * @return Configurable sleep time in seconds between retries for plugins
     * 
     */
    public Optional<Integer> sleep() {
        return Optional.ofNullable(this.sleep);
    }
    /**
     * @return Configurable timeout time in seconds for plugins
     * 
     */
    public Optional<Integer> timeout() {
        return Optional.ofNullable(this.timeout);
    }
    public String version() {
        return this.version;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetPluginsPlugin defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String description;
        private Boolean enabled;
        private String name;
        private @Nullable Integer sleep;
        private @Nullable Integer timeout;
        private String version;
        public Builder() {}
        public Builder(GetPluginsPlugin defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.description = defaults.description;
    	      this.enabled = defaults.enabled;
    	      this.name = defaults.name;
    	      this.sleep = defaults.sleep;
    	      this.timeout = defaults.timeout;
    	      this.version = defaults.version;
        }

        @CustomType.Setter
        public Builder description(String description) {
            if (description == null) {
              throw new MissingRequiredPropertyException("GetPluginsPlugin", "description");
            }
            this.description = description;
            return this;
        }
        @CustomType.Setter
        public Builder enabled(Boolean enabled) {
            if (enabled == null) {
              throw new MissingRequiredPropertyException("GetPluginsPlugin", "enabled");
            }
            this.enabled = enabled;
            return this;
        }
        @CustomType.Setter
        public Builder name(String name) {
            if (name == null) {
              throw new MissingRequiredPropertyException("GetPluginsPlugin", "name");
            }
            this.name = name;
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
        public Builder version(String version) {
            if (version == null) {
              throw new MissingRequiredPropertyException("GetPluginsPlugin", "version");
            }
            this.version = version;
            return this;
        }
        public GetPluginsPlugin build() {
            final var _resultValue = new GetPluginsPlugin();
            _resultValue.description = description;
            _resultValue.enabled = enabled;
            _resultValue.name = name;
            _resultValue.sleep = sleep;
            _resultValue.timeout = timeout;
            _resultValue.version = version;
            return _resultValue;
        }
    }
}
