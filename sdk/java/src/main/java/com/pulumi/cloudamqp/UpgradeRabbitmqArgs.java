// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.cloudamqp;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class UpgradeRabbitmqArgs extends com.pulumi.resources.ResourceArgs {

    public static final UpgradeRabbitmqArgs Empty = new UpgradeRabbitmqArgs();

    /**
     * Helper argument to change upgrade behaviour to latest possible
     * version
     * 
     */
    @Import(name="currentVersion")
    private @Nullable Output<String> currentVersion;

    /**
     * @return Helper argument to change upgrade behaviour to latest possible
     * version
     * 
     */
    public Optional<Output<String>> currentVersion() {
        return Optional.ofNullable(this.currentVersion);
    }

    /**
     * The CloudAMQP instance identifier
     * 
     */
    @Import(name="instanceId", required=true)
    private Output<Integer> instanceId;

    /**
     * @return The CloudAMQP instance identifier
     * 
     */
    public Output<Integer> instanceId() {
        return this.instanceId;
    }

    /**
     * The new version to upgrade to
     * 
     */
    @Import(name="newVersion")
    private @Nullable Output<String> newVersion;

    /**
     * @return The new version to upgrade to
     * 
     */
    public Optional<Output<String>> newVersion() {
        return Optional.ofNullable(this.newVersion);
    }

    private UpgradeRabbitmqArgs() {}

    private UpgradeRabbitmqArgs(UpgradeRabbitmqArgs $) {
        this.currentVersion = $.currentVersion;
        this.instanceId = $.instanceId;
        this.newVersion = $.newVersion;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(UpgradeRabbitmqArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private UpgradeRabbitmqArgs $;

        public Builder() {
            $ = new UpgradeRabbitmqArgs();
        }

        public Builder(UpgradeRabbitmqArgs defaults) {
            $ = new UpgradeRabbitmqArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param currentVersion Helper argument to change upgrade behaviour to latest possible
         * version
         * 
         * @return builder
         * 
         */
        public Builder currentVersion(@Nullable Output<String> currentVersion) {
            $.currentVersion = currentVersion;
            return this;
        }

        /**
         * @param currentVersion Helper argument to change upgrade behaviour to latest possible
         * version
         * 
         * @return builder
         * 
         */
        public Builder currentVersion(String currentVersion) {
            return currentVersion(Output.of(currentVersion));
        }

        /**
         * @param instanceId The CloudAMQP instance identifier
         * 
         * @return builder
         * 
         */
        public Builder instanceId(Output<Integer> instanceId) {
            $.instanceId = instanceId;
            return this;
        }

        /**
         * @param instanceId The CloudAMQP instance identifier
         * 
         * @return builder
         * 
         */
        public Builder instanceId(Integer instanceId) {
            return instanceId(Output.of(instanceId));
        }

        /**
         * @param newVersion The new version to upgrade to
         * 
         * @return builder
         * 
         */
        public Builder newVersion(@Nullable Output<String> newVersion) {
            $.newVersion = newVersion;
            return this;
        }

        /**
         * @param newVersion The new version to upgrade to
         * 
         * @return builder
         * 
         */
        public Builder newVersion(String newVersion) {
            return newVersion(Output.of(newVersion));
        }

        public UpgradeRabbitmqArgs build() {
            if ($.instanceId == null) {
                throw new MissingRequiredPropertyException("UpgradeRabbitmqArgs", "instanceId");
            }
            return $;
        }
    }

}
