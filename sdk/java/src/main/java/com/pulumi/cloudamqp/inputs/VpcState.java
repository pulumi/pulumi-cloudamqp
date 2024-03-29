// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.cloudamqp.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class VpcState extends com.pulumi.resources.ResourceArgs {

    public static final VpcState Empty = new VpcState();

    /**
     * The name of the VPC.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return The name of the VPC.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * The hosted region for the managed standalone VPC
     * 
     */
    @Import(name="region")
    private @Nullable Output<String> region;

    /**
     * @return The hosted region for the managed standalone VPC
     * 
     */
    public Optional<Output<String>> region() {
        return Optional.ofNullable(this.region);
    }

    /**
     * The VPC subnet
     * 
     */
    @Import(name="subnet")
    private @Nullable Output<String> subnet;

    /**
     * @return The VPC subnet
     * 
     */
    public Optional<Output<String>> subnet() {
        return Optional.ofNullable(this.subnet);
    }

    /**
     * Tag the VPC with optional tags
     * 
     */
    @Import(name="tags")
    private @Nullable Output<List<String>> tags;

    /**
     * @return Tag the VPC with optional tags
     * 
     */
    public Optional<Output<List<String>>> tags() {
        return Optional.ofNullable(this.tags);
    }

    /**
     * VPC name given when hosted at the cloud provider
     * 
     */
    @Import(name="vpcName")
    private @Nullable Output<String> vpcName;

    /**
     * @return VPC name given when hosted at the cloud provider
     * 
     */
    public Optional<Output<String>> vpcName() {
        return Optional.ofNullable(this.vpcName);
    }

    private VpcState() {}

    private VpcState(VpcState $) {
        this.name = $.name;
        this.region = $.region;
        this.subnet = $.subnet;
        this.tags = $.tags;
        this.vpcName = $.vpcName;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(VpcState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private VpcState $;

        public Builder() {
            $ = new VpcState();
        }

        public Builder(VpcState defaults) {
            $ = new VpcState(Objects.requireNonNull(defaults));
        }

        /**
         * @param name The name of the VPC.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name The name of the VPC.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param region The hosted region for the managed standalone VPC
         * 
         * @return builder
         * 
         */
        public Builder region(@Nullable Output<String> region) {
            $.region = region;
            return this;
        }

        /**
         * @param region The hosted region for the managed standalone VPC
         * 
         * @return builder
         * 
         */
        public Builder region(String region) {
            return region(Output.of(region));
        }

        /**
         * @param subnet The VPC subnet
         * 
         * @return builder
         * 
         */
        public Builder subnet(@Nullable Output<String> subnet) {
            $.subnet = subnet;
            return this;
        }

        /**
         * @param subnet The VPC subnet
         * 
         * @return builder
         * 
         */
        public Builder subnet(String subnet) {
            return subnet(Output.of(subnet));
        }

        /**
         * @param tags Tag the VPC with optional tags
         * 
         * @return builder
         * 
         */
        public Builder tags(@Nullable Output<List<String>> tags) {
            $.tags = tags;
            return this;
        }

        /**
         * @param tags Tag the VPC with optional tags
         * 
         * @return builder
         * 
         */
        public Builder tags(List<String> tags) {
            return tags(Output.of(tags));
        }

        /**
         * @param tags Tag the VPC with optional tags
         * 
         * @return builder
         * 
         */
        public Builder tags(String... tags) {
            return tags(List.of(tags));
        }

        /**
         * @param vpcName VPC name given when hosted at the cloud provider
         * 
         * @return builder
         * 
         */
        public Builder vpcName(@Nullable Output<String> vpcName) {
            $.vpcName = vpcName;
            return this;
        }

        /**
         * @param vpcName VPC name given when hosted at the cloud provider
         * 
         * @return builder
         * 
         */
        public Builder vpcName(String vpcName) {
            return vpcName(Output.of(vpcName));
        }

        public VpcState build() {
            return $;
        }
    }

}
