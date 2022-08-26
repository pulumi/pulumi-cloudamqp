// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.cloudamqp;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class InstanceArgs extends com.pulumi.resources.ResourceArgs {

    public static final InstanceArgs Empty = new InstanceArgs();

    /**
     * Keep associated VPC when deleting instance, default set to false.
     * 
     */
    @Import(name="keepAssociatedVpc")
    private @Nullable Output<Boolean> keepAssociatedVpc;

    /**
     * @return Keep associated VPC when deleting instance, default set to false.
     * 
     */
    public Optional<Output<Boolean>> keepAssociatedVpc() {
        return Optional.ofNullable(this.keepAssociatedVpc);
    }

    /**
     * Name of the CloudAMQP instance.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return Name of the CloudAMQP instance.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * Set to true to discard creating default alarms when the instance is created. Can be left out, will then use default value = false.
     * 
     */
    @Import(name="noDefaultAlarms")
    private @Nullable Output<Boolean> noDefaultAlarms;

    /**
     * @return Set to true to discard creating default alarms when the instance is created. Can be left out, will then use default value = false.
     * 
     */
    public Optional<Output<Boolean>> noDefaultAlarms() {
        return Optional.ofNullable(this.noDefaultAlarms);
    }

    /**
     * Number of nodes, 1, 3 or 5 depending on plan used.
     * 
     */
    @Import(name="nodes")
    private @Nullable Output<Integer> nodes;

    /**
     * @return Number of nodes, 1, 3 or 5 depending on plan used.
     * 
     */
    public Optional<Output<Integer>> nodes() {
        return Optional.ofNullable(this.nodes);
    }

    /**
     * The subscription plan. See available plans
     * 
     */
    @Import(name="plan", required=true)
    private Output<String> plan;

    /**
     * @return The subscription plan. See available plans
     * 
     */
    public Output<String> plan() {
        return this.plan;
    }

    /**
     * The region to host the instance in. See Instance regions
     * 
     */
    @Import(name="region", required=true)
    private Output<String> region;

    /**
     * @return The region to host the instance in. See Instance regions
     * 
     */
    public Output<String> region() {
        return this.region;
    }

    /**
     * The Rabbit MQ version. Can be left out, will then be set to default value used by CloudAMQP API.
     * 
     */
    @Import(name="rmqVersion")
    private @Nullable Output<String> rmqVersion;

    /**
     * @return The Rabbit MQ version. Can be left out, will then be set to default value used by CloudAMQP API.
     * 
     */
    public Optional<Output<String>> rmqVersion() {
        return Optional.ofNullable(this.rmqVersion);
    }

    /**
     * One or more tags for the CloudAMQP instance, makes it possible to categories multiple instances in console view. Default there is no tags assigned.
     * 
     */
    @Import(name="tags")
    private @Nullable Output<List<String>> tags;

    /**
     * @return One or more tags for the CloudAMQP instance, makes it possible to categories multiple instances in console view. Default there is no tags assigned.
     * 
     */
    public Optional<Output<List<String>>> tags() {
        return Optional.ofNullable(this.tags);
    }

    /**
     * The VPC ID. Use this to create your instance in an existing VPC. See available example.
     * 
     */
    @Import(name="vpcId")
    private @Nullable Output<Integer> vpcId;

    /**
     * @return The VPC ID. Use this to create your instance in an existing VPC. See available example.
     * 
     */
    public Optional<Output<Integer>> vpcId() {
        return Optional.ofNullable(this.vpcId);
    }

    /**
     * Creates a dedicated VPC subnet, shouldn&#39;t overlap with other VPC subnet, default subnet used 10.56.72.0/24.
     * 
     */
    @Import(name="vpcSubnet")
    private @Nullable Output<String> vpcSubnet;

    /**
     * @return Creates a dedicated VPC subnet, shouldn&#39;t overlap with other VPC subnet, default subnet used 10.56.72.0/24.
     * 
     */
    public Optional<Output<String>> vpcSubnet() {
        return Optional.ofNullable(this.vpcSubnet);
    }

    private InstanceArgs() {}

    private InstanceArgs(InstanceArgs $) {
        this.keepAssociatedVpc = $.keepAssociatedVpc;
        this.name = $.name;
        this.noDefaultAlarms = $.noDefaultAlarms;
        this.nodes = $.nodes;
        this.plan = $.plan;
        this.region = $.region;
        this.rmqVersion = $.rmqVersion;
        this.tags = $.tags;
        this.vpcId = $.vpcId;
        this.vpcSubnet = $.vpcSubnet;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(InstanceArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private InstanceArgs $;

        public Builder() {
            $ = new InstanceArgs();
        }

        public Builder(InstanceArgs defaults) {
            $ = new InstanceArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param keepAssociatedVpc Keep associated VPC when deleting instance, default set to false.
         * 
         * @return builder
         * 
         */
        public Builder keepAssociatedVpc(@Nullable Output<Boolean> keepAssociatedVpc) {
            $.keepAssociatedVpc = keepAssociatedVpc;
            return this;
        }

        /**
         * @param keepAssociatedVpc Keep associated VPC when deleting instance, default set to false.
         * 
         * @return builder
         * 
         */
        public Builder keepAssociatedVpc(Boolean keepAssociatedVpc) {
            return keepAssociatedVpc(Output.of(keepAssociatedVpc));
        }

        /**
         * @param name Name of the CloudAMQP instance.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name Name of the CloudAMQP instance.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param noDefaultAlarms Set to true to discard creating default alarms when the instance is created. Can be left out, will then use default value = false.
         * 
         * @return builder
         * 
         */
        public Builder noDefaultAlarms(@Nullable Output<Boolean> noDefaultAlarms) {
            $.noDefaultAlarms = noDefaultAlarms;
            return this;
        }

        /**
         * @param noDefaultAlarms Set to true to discard creating default alarms when the instance is created. Can be left out, will then use default value = false.
         * 
         * @return builder
         * 
         */
        public Builder noDefaultAlarms(Boolean noDefaultAlarms) {
            return noDefaultAlarms(Output.of(noDefaultAlarms));
        }

        /**
         * @param nodes Number of nodes, 1, 3 or 5 depending on plan used.
         * 
         * @return builder
         * 
         */
        public Builder nodes(@Nullable Output<Integer> nodes) {
            $.nodes = nodes;
            return this;
        }

        /**
         * @param nodes Number of nodes, 1, 3 or 5 depending on plan used.
         * 
         * @return builder
         * 
         */
        public Builder nodes(Integer nodes) {
            return nodes(Output.of(nodes));
        }

        /**
         * @param plan The subscription plan. See available plans
         * 
         * @return builder
         * 
         */
        public Builder plan(Output<String> plan) {
            $.plan = plan;
            return this;
        }

        /**
         * @param plan The subscription plan. See available plans
         * 
         * @return builder
         * 
         */
        public Builder plan(String plan) {
            return plan(Output.of(plan));
        }

        /**
         * @param region The region to host the instance in. See Instance regions
         * 
         * @return builder
         * 
         */
        public Builder region(Output<String> region) {
            $.region = region;
            return this;
        }

        /**
         * @param region The region to host the instance in. See Instance regions
         * 
         * @return builder
         * 
         */
        public Builder region(String region) {
            return region(Output.of(region));
        }

        /**
         * @param rmqVersion The Rabbit MQ version. Can be left out, will then be set to default value used by CloudAMQP API.
         * 
         * @return builder
         * 
         */
        public Builder rmqVersion(@Nullable Output<String> rmqVersion) {
            $.rmqVersion = rmqVersion;
            return this;
        }

        /**
         * @param rmqVersion The Rabbit MQ version. Can be left out, will then be set to default value used by CloudAMQP API.
         * 
         * @return builder
         * 
         */
        public Builder rmqVersion(String rmqVersion) {
            return rmqVersion(Output.of(rmqVersion));
        }

        /**
         * @param tags One or more tags for the CloudAMQP instance, makes it possible to categories multiple instances in console view. Default there is no tags assigned.
         * 
         * @return builder
         * 
         */
        public Builder tags(@Nullable Output<List<String>> tags) {
            $.tags = tags;
            return this;
        }

        /**
         * @param tags One or more tags for the CloudAMQP instance, makes it possible to categories multiple instances in console view. Default there is no tags assigned.
         * 
         * @return builder
         * 
         */
        public Builder tags(List<String> tags) {
            return tags(Output.of(tags));
        }

        /**
         * @param tags One or more tags for the CloudAMQP instance, makes it possible to categories multiple instances in console view. Default there is no tags assigned.
         * 
         * @return builder
         * 
         */
        public Builder tags(String... tags) {
            return tags(List.of(tags));
        }

        /**
         * @param vpcId The VPC ID. Use this to create your instance in an existing VPC. See available example.
         * 
         * @return builder
         * 
         */
        public Builder vpcId(@Nullable Output<Integer> vpcId) {
            $.vpcId = vpcId;
            return this;
        }

        /**
         * @param vpcId The VPC ID. Use this to create your instance in an existing VPC. See available example.
         * 
         * @return builder
         * 
         */
        public Builder vpcId(Integer vpcId) {
            return vpcId(Output.of(vpcId));
        }

        /**
         * @param vpcSubnet Creates a dedicated VPC subnet, shouldn&#39;t overlap with other VPC subnet, default subnet used 10.56.72.0/24.
         * 
         * @return builder
         * 
         */
        public Builder vpcSubnet(@Nullable Output<String> vpcSubnet) {
            $.vpcSubnet = vpcSubnet;
            return this;
        }

        /**
         * @param vpcSubnet Creates a dedicated VPC subnet, shouldn&#39;t overlap with other VPC subnet, default subnet used 10.56.72.0/24.
         * 
         * @return builder
         * 
         */
        public Builder vpcSubnet(String vpcSubnet) {
            return vpcSubnet(Output.of(vpcSubnet));
        }

        public InstanceArgs build() {
            $.plan = Objects.requireNonNull($.plan, "expected parameter 'plan' to be non-null");
            $.region = Objects.requireNonNull($.region, "expected parameter 'region' to be non-null");
            return $;
        }
    }

}
