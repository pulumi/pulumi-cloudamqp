// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.cloudamqp.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.List;
import java.util.Objects;


public final class InstanceCopySettingArgs extends com.pulumi.resources.ResourceArgs {

    public static final InstanceCopySettingArgs Empty = new InstanceCopySettingArgs();

    /**
     * Array of one or more settings to be copied. Allowed values:
     * [alarms, config, definitions, firewall, logs, metrics, plugins]
     * 
     * See more below, [copy settings].
     * 
     */
    @Import(name="settings", required=true)
    private Output<List<String>> settings;

    /**
     * @return Array of one or more settings to be copied. Allowed values:
     * [alarms, config, definitions, firewall, logs, metrics, plugins]
     * 
     * See more below, [copy settings].
     * 
     */
    public Output<List<String>> settings() {
        return this.settings;
    }

    /**
     * Instance identifier of the CloudAMQP instance to copy the settings
     * from.
     * 
     */
    @Import(name="subscriptionId", required=true)
    private Output<String> subscriptionId;

    /**
     * @return Instance identifier of the CloudAMQP instance to copy the settings
     * from.
     * 
     */
    public Output<String> subscriptionId() {
        return this.subscriptionId;
    }

    private InstanceCopySettingArgs() {}

    private InstanceCopySettingArgs(InstanceCopySettingArgs $) {
        this.settings = $.settings;
        this.subscriptionId = $.subscriptionId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(InstanceCopySettingArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private InstanceCopySettingArgs $;

        public Builder() {
            $ = new InstanceCopySettingArgs();
        }

        public Builder(InstanceCopySettingArgs defaults) {
            $ = new InstanceCopySettingArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param settings Array of one or more settings to be copied. Allowed values:
         * [alarms, config, definitions, firewall, logs, metrics, plugins]
         * 
         * See more below, [copy settings].
         * 
         * @return builder
         * 
         */
        public Builder settings(Output<List<String>> settings) {
            $.settings = settings;
            return this;
        }

        /**
         * @param settings Array of one or more settings to be copied. Allowed values:
         * [alarms, config, definitions, firewall, logs, metrics, plugins]
         * 
         * See more below, [copy settings].
         * 
         * @return builder
         * 
         */
        public Builder settings(List<String> settings) {
            return settings(Output.of(settings));
        }

        /**
         * @param settings Array of one or more settings to be copied. Allowed values:
         * [alarms, config, definitions, firewall, logs, metrics, plugins]
         * 
         * See more below, [copy settings].
         * 
         * @return builder
         * 
         */
        public Builder settings(String... settings) {
            return settings(List.of(settings));
        }

        /**
         * @param subscriptionId Instance identifier of the CloudAMQP instance to copy the settings
         * from.
         * 
         * @return builder
         * 
         */
        public Builder subscriptionId(Output<String> subscriptionId) {
            $.subscriptionId = subscriptionId;
            return this;
        }

        /**
         * @param subscriptionId Instance identifier of the CloudAMQP instance to copy the settings
         * from.
         * 
         * @return builder
         * 
         */
        public Builder subscriptionId(String subscriptionId) {
            return subscriptionId(Output.of(subscriptionId));
        }

        public InstanceCopySettingArgs build() {
            if ($.settings == null) {
                throw new MissingRequiredPropertyException("InstanceCopySettingArgs", "settings");
            }
            if ($.subscriptionId == null) {
                throw new MissingRequiredPropertyException("InstanceCopySettingArgs", "subscriptionId");
            }
            return $;
        }
    }

}
