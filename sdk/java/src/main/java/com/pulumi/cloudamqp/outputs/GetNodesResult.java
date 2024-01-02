// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.cloudamqp.outputs;

import com.pulumi.cloudamqp.outputs.GetNodesNode;
import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;

@CustomType
public final class GetNodesResult {
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    private Integer instanceId;
    private List<GetNodesNode> nodes;

    private GetNodesResult() {}
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
    public List<GetNodesNode> nodes() {
        return this.nodes;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetNodesResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String id;
        private Integer instanceId;
        private List<GetNodesNode> nodes;
        public Builder() {}
        public Builder(GetNodesResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.id = defaults.id;
    	      this.instanceId = defaults.instanceId;
    	      this.nodes = defaults.nodes;
        }

        @CustomType.Setter
        public Builder id(String id) {
            if (id == null) {
              throw new MissingRequiredPropertyException("GetNodesResult", "id");
            }
            this.id = id;
            return this;
        }
        @CustomType.Setter
        public Builder instanceId(Integer instanceId) {
            if (instanceId == null) {
              throw new MissingRequiredPropertyException("GetNodesResult", "instanceId");
            }
            this.instanceId = instanceId;
            return this;
        }
        @CustomType.Setter
        public Builder nodes(List<GetNodesNode> nodes) {
            if (nodes == null) {
              throw new MissingRequiredPropertyException("GetNodesResult", "nodes");
            }
            this.nodes = nodes;
            return this;
        }
        public Builder nodes(GetNodesNode... nodes) {
            return nodes(List.of(nodes));
        }
        public GetNodesResult build() {
            final var _resultValue = new GetNodesResult();
            _resultValue.id = id;
            _resultValue.instanceId = instanceId;
            _resultValue.nodes = nodes;
            return _resultValue;
        }
    }
}
