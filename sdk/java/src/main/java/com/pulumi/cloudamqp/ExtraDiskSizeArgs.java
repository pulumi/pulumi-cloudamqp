// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.cloudamqp;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Integer;
import java.util.Objects;


public final class ExtraDiskSizeArgs extends com.pulumi.resources.ResourceArgs {

    public static final ExtraDiskSizeArgs Empty = new ExtraDiskSizeArgs();

    /**
     * Extra disk size in GB. Supported values: 25, 50, 100, 250, 500, 1000, 2000
     * 
     */
    @Import(name="extraDiskSize", required=true)
    private Output<Integer> extraDiskSize;

    /**
     * @return Extra disk size in GB. Supported values: 25, 50, 100, 250, 500, 1000, 2000
     * 
     */
    public Output<Integer> extraDiskSize() {
        return this.extraDiskSize;
    }

    /**
     * The CloudAMQP instance ID.
     * 
     */
    @Import(name="instanceId", required=true)
    private Output<Integer> instanceId;

    /**
     * @return The CloudAMQP instance ID.
     * 
     */
    public Output<Integer> instanceId() {
        return this.instanceId;
    }

    private ExtraDiskSizeArgs() {}

    private ExtraDiskSizeArgs(ExtraDiskSizeArgs $) {
        this.extraDiskSize = $.extraDiskSize;
        this.instanceId = $.instanceId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ExtraDiskSizeArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ExtraDiskSizeArgs $;

        public Builder() {
            $ = new ExtraDiskSizeArgs();
        }

        public Builder(ExtraDiskSizeArgs defaults) {
            $ = new ExtraDiskSizeArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param extraDiskSize Extra disk size in GB. Supported values: 25, 50, 100, 250, 500, 1000, 2000
         * 
         * @return builder
         * 
         */
        public Builder extraDiskSize(Output<Integer> extraDiskSize) {
            $.extraDiskSize = extraDiskSize;
            return this;
        }

        /**
         * @param extraDiskSize Extra disk size in GB. Supported values: 25, 50, 100, 250, 500, 1000, 2000
         * 
         * @return builder
         * 
         */
        public Builder extraDiskSize(Integer extraDiskSize) {
            return extraDiskSize(Output.of(extraDiskSize));
        }

        /**
         * @param instanceId The CloudAMQP instance ID.
         * 
         * @return builder
         * 
         */
        public Builder instanceId(Output<Integer> instanceId) {
            $.instanceId = instanceId;
            return this;
        }

        /**
         * @param instanceId The CloudAMQP instance ID.
         * 
         * @return builder
         * 
         */
        public Builder instanceId(Integer instanceId) {
            return instanceId(Output.of(instanceId));
        }

        public ExtraDiskSizeArgs build() {
            $.extraDiskSize = Objects.requireNonNull($.extraDiskSize, "expected parameter 'extraDiskSize' to be non-null");
            $.instanceId = Objects.requireNonNull($.instanceId, "expected parameter 'instanceId' to be non-null");
            return $;
        }
    }

}
