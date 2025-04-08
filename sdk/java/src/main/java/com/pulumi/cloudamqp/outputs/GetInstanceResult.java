// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.cloudamqp.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;

@CustomType
public final class GetInstanceResult {
    /**
     * @return (Sensitive) The API key to secondary API handing alarms, integration etc.
     * 
     */
    private String apikey;
    /**
     * @return Information if the CloudAMQP instance runs either RabbitMQ or LavinMQ.
     * 
     */
    private String backend;
    /**
     * @return Information if the CloudAMQP instance is shared or dedicated.
     * 
     */
    private Boolean dedicated;
    /**
     * @return The external hostname for the CloudAMQP instance.
     * 
     */
    private String host;
    /**
     * @return The internal hostname for the CloudAMQP instance.
     * 
     */
    private String hostInternal;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    private Integer instanceId;
    /**
     * @return The name of the CloudAMQP instance.
     * 
     */
    private String name;
    private Boolean noDefaultAlarms;
    /**
     * @return Number of nodes in the cluster of the CloudAMQP instance.
     * 
     */
    private Integer nodes;
    /**
     * @return The subscription plan for the CloudAMQP instance.
     * 
     */
    private String plan;
    private Boolean ready;
    /**
     * @return The cloud platform and region that host the CloudAMQP instance,
     * `{platform}::{region}`.
     * 
     */
    private String region;
    /**
     * @return The version of installed Rabbit MQ.
     * 
     */
    private String rmqVersion;
    /**
     * @return Tags the CloudAMQP instance with categories.
     * 
     */
    private List<String> tags;
    /**
     * @return (Sensitive) The AMQP URL (uses the internal hostname if the instance was created
     * with VPC), used by clients to connect for pub/sub.
     * 
     */
    private String url;
    /**
     * @return The virtual host configured in Rabbit MQ.
     * 
     */
    private String vhost;
    /**
     * @return ID of the VPC configured for the CloudAMQP instance.
     * 
     */
    private Integer vpcId;
    /**
     * @return Dedicated VPC subnet configured for the CloudAMQP instance.
     * 
     */
    private String vpcSubnet;

    private GetInstanceResult() {}
    /**
     * @return (Sensitive) The API key to secondary API handing alarms, integration etc.
     * 
     */
    public String apikey() {
        return this.apikey;
    }
    /**
     * @return Information if the CloudAMQP instance runs either RabbitMQ or LavinMQ.
     * 
     */
    public String backend() {
        return this.backend;
    }
    /**
     * @return Information if the CloudAMQP instance is shared or dedicated.
     * 
     */
    public Boolean dedicated() {
        return this.dedicated;
    }
    /**
     * @return The external hostname for the CloudAMQP instance.
     * 
     */
    public String host() {
        return this.host;
    }
    /**
     * @return The internal hostname for the CloudAMQP instance.
     * 
     */
    public String hostInternal() {
        return this.hostInternal;
    }
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
    /**
     * @return The name of the CloudAMQP instance.
     * 
     */
    public String name() {
        return this.name;
    }
    public Boolean noDefaultAlarms() {
        return this.noDefaultAlarms;
    }
    /**
     * @return Number of nodes in the cluster of the CloudAMQP instance.
     * 
     */
    public Integer nodes() {
        return this.nodes;
    }
    /**
     * @return The subscription plan for the CloudAMQP instance.
     * 
     */
    public String plan() {
        return this.plan;
    }
    public Boolean ready() {
        return this.ready;
    }
    /**
     * @return The cloud platform and region that host the CloudAMQP instance,
     * `{platform}::{region}`.
     * 
     */
    public String region() {
        return this.region;
    }
    /**
     * @return The version of installed Rabbit MQ.
     * 
     */
    public String rmqVersion() {
        return this.rmqVersion;
    }
    /**
     * @return Tags the CloudAMQP instance with categories.
     * 
     */
    public List<String> tags() {
        return this.tags;
    }
    /**
     * @return (Sensitive) The AMQP URL (uses the internal hostname if the instance was created
     * with VPC), used by clients to connect for pub/sub.
     * 
     */
    public String url() {
        return this.url;
    }
    /**
     * @return The virtual host configured in Rabbit MQ.
     * 
     */
    public String vhost() {
        return this.vhost;
    }
    /**
     * @return ID of the VPC configured for the CloudAMQP instance.
     * 
     */
    public Integer vpcId() {
        return this.vpcId;
    }
    /**
     * @return Dedicated VPC subnet configured for the CloudAMQP instance.
     * 
     */
    public String vpcSubnet() {
        return this.vpcSubnet;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetInstanceResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String apikey;
        private String backend;
        private Boolean dedicated;
        private String host;
        private String hostInternal;
        private String id;
        private Integer instanceId;
        private String name;
        private Boolean noDefaultAlarms;
        private Integer nodes;
        private String plan;
        private Boolean ready;
        private String region;
        private String rmqVersion;
        private List<String> tags;
        private String url;
        private String vhost;
        private Integer vpcId;
        private String vpcSubnet;
        public Builder() {}
        public Builder(GetInstanceResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.apikey = defaults.apikey;
    	      this.backend = defaults.backend;
    	      this.dedicated = defaults.dedicated;
    	      this.host = defaults.host;
    	      this.hostInternal = defaults.hostInternal;
    	      this.id = defaults.id;
    	      this.instanceId = defaults.instanceId;
    	      this.name = defaults.name;
    	      this.noDefaultAlarms = defaults.noDefaultAlarms;
    	      this.nodes = defaults.nodes;
    	      this.plan = defaults.plan;
    	      this.ready = defaults.ready;
    	      this.region = defaults.region;
    	      this.rmqVersion = defaults.rmqVersion;
    	      this.tags = defaults.tags;
    	      this.url = defaults.url;
    	      this.vhost = defaults.vhost;
    	      this.vpcId = defaults.vpcId;
    	      this.vpcSubnet = defaults.vpcSubnet;
        }

        @CustomType.Setter
        public Builder apikey(String apikey) {
            if (apikey == null) {
              throw new MissingRequiredPropertyException("GetInstanceResult", "apikey");
            }
            this.apikey = apikey;
            return this;
        }
        @CustomType.Setter
        public Builder backend(String backend) {
            if (backend == null) {
              throw new MissingRequiredPropertyException("GetInstanceResult", "backend");
            }
            this.backend = backend;
            return this;
        }
        @CustomType.Setter
        public Builder dedicated(Boolean dedicated) {
            if (dedicated == null) {
              throw new MissingRequiredPropertyException("GetInstanceResult", "dedicated");
            }
            this.dedicated = dedicated;
            return this;
        }
        @CustomType.Setter
        public Builder host(String host) {
            if (host == null) {
              throw new MissingRequiredPropertyException("GetInstanceResult", "host");
            }
            this.host = host;
            return this;
        }
        @CustomType.Setter
        public Builder hostInternal(String hostInternal) {
            if (hostInternal == null) {
              throw new MissingRequiredPropertyException("GetInstanceResult", "hostInternal");
            }
            this.hostInternal = hostInternal;
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            if (id == null) {
              throw new MissingRequiredPropertyException("GetInstanceResult", "id");
            }
            this.id = id;
            return this;
        }
        @CustomType.Setter
        public Builder instanceId(Integer instanceId) {
            if (instanceId == null) {
              throw new MissingRequiredPropertyException("GetInstanceResult", "instanceId");
            }
            this.instanceId = instanceId;
            return this;
        }
        @CustomType.Setter
        public Builder name(String name) {
            if (name == null) {
              throw new MissingRequiredPropertyException("GetInstanceResult", "name");
            }
            this.name = name;
            return this;
        }
        @CustomType.Setter
        public Builder noDefaultAlarms(Boolean noDefaultAlarms) {
            if (noDefaultAlarms == null) {
              throw new MissingRequiredPropertyException("GetInstanceResult", "noDefaultAlarms");
            }
            this.noDefaultAlarms = noDefaultAlarms;
            return this;
        }
        @CustomType.Setter
        public Builder nodes(Integer nodes) {
            if (nodes == null) {
              throw new MissingRequiredPropertyException("GetInstanceResult", "nodes");
            }
            this.nodes = nodes;
            return this;
        }
        @CustomType.Setter
        public Builder plan(String plan) {
            if (plan == null) {
              throw new MissingRequiredPropertyException("GetInstanceResult", "plan");
            }
            this.plan = plan;
            return this;
        }
        @CustomType.Setter
        public Builder ready(Boolean ready) {
            if (ready == null) {
              throw new MissingRequiredPropertyException("GetInstanceResult", "ready");
            }
            this.ready = ready;
            return this;
        }
        @CustomType.Setter
        public Builder region(String region) {
            if (region == null) {
              throw new MissingRequiredPropertyException("GetInstanceResult", "region");
            }
            this.region = region;
            return this;
        }
        @CustomType.Setter
        public Builder rmqVersion(String rmqVersion) {
            if (rmqVersion == null) {
              throw new MissingRequiredPropertyException("GetInstanceResult", "rmqVersion");
            }
            this.rmqVersion = rmqVersion;
            return this;
        }
        @CustomType.Setter
        public Builder tags(List<String> tags) {
            if (tags == null) {
              throw new MissingRequiredPropertyException("GetInstanceResult", "tags");
            }
            this.tags = tags;
            return this;
        }
        public Builder tags(String... tags) {
            return tags(List.of(tags));
        }
        @CustomType.Setter
        public Builder url(String url) {
            if (url == null) {
              throw new MissingRequiredPropertyException("GetInstanceResult", "url");
            }
            this.url = url;
            return this;
        }
        @CustomType.Setter
        public Builder vhost(String vhost) {
            if (vhost == null) {
              throw new MissingRequiredPropertyException("GetInstanceResult", "vhost");
            }
            this.vhost = vhost;
            return this;
        }
        @CustomType.Setter
        public Builder vpcId(Integer vpcId) {
            if (vpcId == null) {
              throw new MissingRequiredPropertyException("GetInstanceResult", "vpcId");
            }
            this.vpcId = vpcId;
            return this;
        }
        @CustomType.Setter
        public Builder vpcSubnet(String vpcSubnet) {
            if (vpcSubnet == null) {
              throw new MissingRequiredPropertyException("GetInstanceResult", "vpcSubnet");
            }
            this.vpcSubnet = vpcSubnet;
            return this;
        }
        public GetInstanceResult build() {
            final var _resultValue = new GetInstanceResult();
            _resultValue.apikey = apikey;
            _resultValue.backend = backend;
            _resultValue.dedicated = dedicated;
            _resultValue.host = host;
            _resultValue.hostInternal = hostInternal;
            _resultValue.id = id;
            _resultValue.instanceId = instanceId;
            _resultValue.name = name;
            _resultValue.noDefaultAlarms = noDefaultAlarms;
            _resultValue.nodes = nodes;
            _resultValue.plan = plan;
            _resultValue.ready = ready;
            _resultValue.region = region;
            _resultValue.rmqVersion = rmqVersion;
            _resultValue.tags = tags;
            _resultValue.url = url;
            _resultValue.vhost = vhost;
            _resultValue.vpcId = vpcId;
            _resultValue.vpcSubnet = vpcSubnet;
            return _resultValue;
        }
    }
}
