// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.cloudamqp.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import javax.annotation.Nullable;

@CustomType
public final class GetAccountInstance {
    /**
     * @return The instance identifier.
     * 
     */
    private Integer id;
    /**
     * @return The name of the instance.
     * 
     */
    private String name;
    /**
     * @return The subscription plan used for the instance.
     * 
     */
    private String plan;
    /**
     * @return The region were the instanece is located in.
     * 
     */
    private String region;
    /**
     * @return Optional tags set for the instance.
     * 
     */
    private @Nullable List<String> tags;

    private GetAccountInstance() {}
    /**
     * @return The instance identifier.
     * 
     */
    public Integer id() {
        return this.id;
    }
    /**
     * @return The name of the instance.
     * 
     */
    public String name() {
        return this.name;
    }
    /**
     * @return The subscription plan used for the instance.
     * 
     */
    public String plan() {
        return this.plan;
    }
    /**
     * @return The region were the instanece is located in.
     * 
     */
    public String region() {
        return this.region;
    }
    /**
     * @return Optional tags set for the instance.
     * 
     */
    public List<String> tags() {
        return this.tags == null ? List.of() : this.tags;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetAccountInstance defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private Integer id;
        private String name;
        private String plan;
        private String region;
        private @Nullable List<String> tags;
        public Builder() {}
        public Builder(GetAccountInstance defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.id = defaults.id;
    	      this.name = defaults.name;
    	      this.plan = defaults.plan;
    	      this.region = defaults.region;
    	      this.tags = defaults.tags;
        }

        @CustomType.Setter
        public Builder id(Integer id) {
            if (id == null) {
              throw new MissingRequiredPropertyException("GetAccountInstance", "id");
            }
            this.id = id;
            return this;
        }
        @CustomType.Setter
        public Builder name(String name) {
            if (name == null) {
              throw new MissingRequiredPropertyException("GetAccountInstance", "name");
            }
            this.name = name;
            return this;
        }
        @CustomType.Setter
        public Builder plan(String plan) {
            if (plan == null) {
              throw new MissingRequiredPropertyException("GetAccountInstance", "plan");
            }
            this.plan = plan;
            return this;
        }
        @CustomType.Setter
        public Builder region(String region) {
            if (region == null) {
              throw new MissingRequiredPropertyException("GetAccountInstance", "region");
            }
            this.region = region;
            return this;
        }
        @CustomType.Setter
        public Builder tags(@Nullable List<String> tags) {

            this.tags = tags;
            return this;
        }
        public Builder tags(String... tags) {
            return tags(List.of(tags));
        }
        public GetAccountInstance build() {
            final var _resultValue = new GetAccountInstance();
            _resultValue.id = id;
            _resultValue.name = name;
            _resultValue.plan = plan;
            _resultValue.region = region;
            _resultValue.tags = tags;
            return _resultValue;
        }
    }
}
