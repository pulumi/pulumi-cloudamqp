// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.cloudamqp;

import com.pulumi.cloudamqp.inputs.NotificationResponderArgs;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class NotificationArgs extends com.pulumi.resources.ResourceArgs {

    public static final NotificationArgs Empty = new NotificationArgs();

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

    /**
     * Name of the responder
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return Name of the responder
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * Options argument (e.g. `rk` used for VictorOps routing key).
     * 
     */
    @Import(name="options")
    private @Nullable Output<Map<String,String>> options;

    /**
     * @return Options argument (e.g. `rk` used for VictorOps routing key).
     * 
     */
    public Optional<Output<Map<String,String>>> options() {
        return Optional.ofNullable(this.options);
    }

    /**
     * An array of reponders (only for OpsGenie). Each `responders` block
     * consists of the field documented below.
     * 
     * ***
     * 
     * The `responders` block consists of:
     * 
     */
    @Import(name="responders")
    private @Nullable Output<List<NotificationResponderArgs>> responders;

    /**
     * @return An array of reponders (only for OpsGenie). Each `responders` block
     * consists of the field documented below.
     * 
     * ***
     * 
     * The `responders` block consists of:
     * 
     */
    public Optional<Output<List<NotificationResponderArgs>>> responders() {
        return Optional.ofNullable(this.responders);
    }

    /**
     * Type of responder. [`team`, `user`, `escalation`, `schedule`]
     * 
     */
    @Import(name="type", required=true)
    private Output<String> type;

    /**
     * @return Type of responder. [`team`, `user`, `escalation`, `schedule`]
     * 
     */
    public Output<String> type() {
        return this.type;
    }

    /**
     * Integration/API key or endpoint to send the notification.
     * 
     */
    @Import(name="value", required=true)
    private Output<String> value;

    /**
     * @return Integration/API key or endpoint to send the notification.
     * 
     */
    public Output<String> value() {
        return this.value;
    }

    private NotificationArgs() {}

    private NotificationArgs(NotificationArgs $) {
        this.instanceId = $.instanceId;
        this.name = $.name;
        this.options = $.options;
        this.responders = $.responders;
        this.type = $.type;
        this.value = $.value;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(NotificationArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private NotificationArgs $;

        public Builder() {
            $ = new NotificationArgs();
        }

        public Builder(NotificationArgs defaults) {
            $ = new NotificationArgs(Objects.requireNonNull(defaults));
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

        /**
         * @param name Name of the responder
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name Name of the responder
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param options Options argument (e.g. `rk` used for VictorOps routing key).
         * 
         * @return builder
         * 
         */
        public Builder options(@Nullable Output<Map<String,String>> options) {
            $.options = options;
            return this;
        }

        /**
         * @param options Options argument (e.g. `rk` used for VictorOps routing key).
         * 
         * @return builder
         * 
         */
        public Builder options(Map<String,String> options) {
            return options(Output.of(options));
        }

        /**
         * @param responders An array of reponders (only for OpsGenie). Each `responders` block
         * consists of the field documented below.
         * 
         * ***
         * 
         * The `responders` block consists of:
         * 
         * @return builder
         * 
         */
        public Builder responders(@Nullable Output<List<NotificationResponderArgs>> responders) {
            $.responders = responders;
            return this;
        }

        /**
         * @param responders An array of reponders (only for OpsGenie). Each `responders` block
         * consists of the field documented below.
         * 
         * ***
         * 
         * The `responders` block consists of:
         * 
         * @return builder
         * 
         */
        public Builder responders(List<NotificationResponderArgs> responders) {
            return responders(Output.of(responders));
        }

        /**
         * @param responders An array of reponders (only for OpsGenie). Each `responders` block
         * consists of the field documented below.
         * 
         * ***
         * 
         * The `responders` block consists of:
         * 
         * @return builder
         * 
         */
        public Builder responders(NotificationResponderArgs... responders) {
            return responders(List.of(responders));
        }

        /**
         * @param type Type of responder. [`team`, `user`, `escalation`, `schedule`]
         * 
         * @return builder
         * 
         */
        public Builder type(Output<String> type) {
            $.type = type;
            return this;
        }

        /**
         * @param type Type of responder. [`team`, `user`, `escalation`, `schedule`]
         * 
         * @return builder
         * 
         */
        public Builder type(String type) {
            return type(Output.of(type));
        }

        /**
         * @param value Integration/API key or endpoint to send the notification.
         * 
         * @return builder
         * 
         */
        public Builder value(Output<String> value) {
            $.value = value;
            return this;
        }

        /**
         * @param value Integration/API key or endpoint to send the notification.
         * 
         * @return builder
         * 
         */
        public Builder value(String value) {
            return value(Output.of(value));
        }

        public NotificationArgs build() {
            if ($.instanceId == null) {
                throw new MissingRequiredPropertyException("NotificationArgs", "instanceId");
            }
            if ($.type == null) {
                throw new MissingRequiredPropertyException("NotificationArgs", "type");
            }
            if ($.value == null) {
                throw new MissingRequiredPropertyException("NotificationArgs", "value");
            }
            return $;
        }
    }

}