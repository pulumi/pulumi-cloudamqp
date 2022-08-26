// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.cloudamqp.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class AlarmState extends com.pulumi.resources.ResourceArgs {

    public static final AlarmState Empty = new AlarmState();

    /**
     * Enable or disable the alarm to trigger.
     * 
     */
    @Import(name="enabled")
    private @Nullable Output<Boolean> enabled;

    /**
     * @return Enable or disable the alarm to trigger.
     * 
     */
    public Optional<Output<Boolean>> enabled() {
        return Optional.ofNullable(this.enabled);
    }

    /**
     * The CloudAMQP instance ID.
     * 
     */
    @Import(name="instanceId")
    private @Nullable Output<Integer> instanceId;

    /**
     * @return The CloudAMQP instance ID.
     * 
     */
    public Optional<Output<Integer>> instanceId() {
        return Optional.ofNullable(this.instanceId);
    }

    /**
     * Message type `(total, unacked, ready)` used by queue alarm type.
     * 
     */
    @Import(name="messageType")
    private @Nullable Output<String> messageType;

    /**
     * @return Message type `(total, unacked, ready)` used by queue alarm type.
     * 
     */
    public Optional<Output<String>> messageType() {
        return Optional.ofNullable(this.messageType);
    }

    /**
     * Regex for which queue to check.
     * 
     */
    @Import(name="queueRegex")
    private @Nullable Output<String> queueRegex;

    /**
     * @return Regex for which queue to check.
     * 
     */
    public Optional<Output<String>> queueRegex() {
        return Optional.ofNullable(this.queueRegex);
    }

    /**
     * Identifier for recipient to be notified. Leave empty to notify all recipients.
     * 
     */
    @Import(name="recipients")
    private @Nullable Output<List<Integer>> recipients;

    /**
     * @return Identifier for recipient to be notified. Leave empty to notify all recipients.
     * 
     */
    public Optional<Output<List<Integer>>> recipients() {
        return Optional.ofNullable(this.recipients);
    }

    /**
     * The reminder interval (in seconds) to resend the alarm if not resolved. Set to 0 for no reminders. The Default is 0.
     * 
     */
    @Import(name="reminderInterval")
    private @Nullable Output<Integer> reminderInterval;

    /**
     * @return The reminder interval (in seconds) to resend the alarm if not resolved. Set to 0 for no reminders. The Default is 0.
     * 
     */
    public Optional<Output<Integer>> reminderInterval() {
        return Optional.ofNullable(this.reminderInterval);
    }

    /**
     * The time interval (in seconds) the `value_threshold` should be active before triggering an alarm.
     * 
     */
    @Import(name="timeThreshold")
    private @Nullable Output<Integer> timeThreshold;

    /**
     * @return The time interval (in seconds) the `value_threshold` should be active before triggering an alarm.
     * 
     */
    public Optional<Output<Integer>> timeThreshold() {
        return Optional.ofNullable(this.timeThreshold);
    }

    /**
     * The alarm type, see valid options below.
     * 
     */
    @Import(name="type")
    private @Nullable Output<String> type;

    /**
     * @return The alarm type, see valid options below.
     * 
     */
    public Optional<Output<String>> type() {
        return Optional.ofNullable(this.type);
    }

    /**
     * Disk value threshold calculation, `fixed, percentage` of disk space remaining.
     * 
     */
    @Import(name="valueCalculation")
    private @Nullable Output<String> valueCalculation;

    /**
     * @return Disk value threshold calculation, `fixed, percentage` of disk space remaining.
     * 
     */
    public Optional<Output<String>> valueCalculation() {
        return Optional.ofNullable(this.valueCalculation);
    }

    /**
     * The value to trigger the alarm for.
     * 
     */
    @Import(name="valueThreshold")
    private @Nullable Output<Integer> valueThreshold;

    /**
     * @return The value to trigger the alarm for.
     * 
     */
    public Optional<Output<Integer>> valueThreshold() {
        return Optional.ofNullable(this.valueThreshold);
    }

    /**
     * Regex for which vhost to check
     * 
     */
    @Import(name="vhostRegex")
    private @Nullable Output<String> vhostRegex;

    /**
     * @return Regex for which vhost to check
     * 
     */
    public Optional<Output<String>> vhostRegex() {
        return Optional.ofNullable(this.vhostRegex);
    }

    private AlarmState() {}

    private AlarmState(AlarmState $) {
        this.enabled = $.enabled;
        this.instanceId = $.instanceId;
        this.messageType = $.messageType;
        this.queueRegex = $.queueRegex;
        this.recipients = $.recipients;
        this.reminderInterval = $.reminderInterval;
        this.timeThreshold = $.timeThreshold;
        this.type = $.type;
        this.valueCalculation = $.valueCalculation;
        this.valueThreshold = $.valueThreshold;
        this.vhostRegex = $.vhostRegex;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(AlarmState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private AlarmState $;

        public Builder() {
            $ = new AlarmState();
        }

        public Builder(AlarmState defaults) {
            $ = new AlarmState(Objects.requireNonNull(defaults));
        }

        /**
         * @param enabled Enable or disable the alarm to trigger.
         * 
         * @return builder
         * 
         */
        public Builder enabled(@Nullable Output<Boolean> enabled) {
            $.enabled = enabled;
            return this;
        }

        /**
         * @param enabled Enable or disable the alarm to trigger.
         * 
         * @return builder
         * 
         */
        public Builder enabled(Boolean enabled) {
            return enabled(Output.of(enabled));
        }

        /**
         * @param instanceId The CloudAMQP instance ID.
         * 
         * @return builder
         * 
         */
        public Builder instanceId(@Nullable Output<Integer> instanceId) {
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
         * @param messageType Message type `(total, unacked, ready)` used by queue alarm type.
         * 
         * @return builder
         * 
         */
        public Builder messageType(@Nullable Output<String> messageType) {
            $.messageType = messageType;
            return this;
        }

        /**
         * @param messageType Message type `(total, unacked, ready)` used by queue alarm type.
         * 
         * @return builder
         * 
         */
        public Builder messageType(String messageType) {
            return messageType(Output.of(messageType));
        }

        /**
         * @param queueRegex Regex for which queue to check.
         * 
         * @return builder
         * 
         */
        public Builder queueRegex(@Nullable Output<String> queueRegex) {
            $.queueRegex = queueRegex;
            return this;
        }

        /**
         * @param queueRegex Regex for which queue to check.
         * 
         * @return builder
         * 
         */
        public Builder queueRegex(String queueRegex) {
            return queueRegex(Output.of(queueRegex));
        }

        /**
         * @param recipients Identifier for recipient to be notified. Leave empty to notify all recipients.
         * 
         * @return builder
         * 
         */
        public Builder recipients(@Nullable Output<List<Integer>> recipients) {
            $.recipients = recipients;
            return this;
        }

        /**
         * @param recipients Identifier for recipient to be notified. Leave empty to notify all recipients.
         * 
         * @return builder
         * 
         */
        public Builder recipients(List<Integer> recipients) {
            return recipients(Output.of(recipients));
        }

        /**
         * @param recipients Identifier for recipient to be notified. Leave empty to notify all recipients.
         * 
         * @return builder
         * 
         */
        public Builder recipients(Integer... recipients) {
            return recipients(List.of(recipients));
        }

        /**
         * @param reminderInterval The reminder interval (in seconds) to resend the alarm if not resolved. Set to 0 for no reminders. The Default is 0.
         * 
         * @return builder
         * 
         */
        public Builder reminderInterval(@Nullable Output<Integer> reminderInterval) {
            $.reminderInterval = reminderInterval;
            return this;
        }

        /**
         * @param reminderInterval The reminder interval (in seconds) to resend the alarm if not resolved. Set to 0 for no reminders. The Default is 0.
         * 
         * @return builder
         * 
         */
        public Builder reminderInterval(Integer reminderInterval) {
            return reminderInterval(Output.of(reminderInterval));
        }

        /**
         * @param timeThreshold The time interval (in seconds) the `value_threshold` should be active before triggering an alarm.
         * 
         * @return builder
         * 
         */
        public Builder timeThreshold(@Nullable Output<Integer> timeThreshold) {
            $.timeThreshold = timeThreshold;
            return this;
        }

        /**
         * @param timeThreshold The time interval (in seconds) the `value_threshold` should be active before triggering an alarm.
         * 
         * @return builder
         * 
         */
        public Builder timeThreshold(Integer timeThreshold) {
            return timeThreshold(Output.of(timeThreshold));
        }

        /**
         * @param type The alarm type, see valid options below.
         * 
         * @return builder
         * 
         */
        public Builder type(@Nullable Output<String> type) {
            $.type = type;
            return this;
        }

        /**
         * @param type The alarm type, see valid options below.
         * 
         * @return builder
         * 
         */
        public Builder type(String type) {
            return type(Output.of(type));
        }

        /**
         * @param valueCalculation Disk value threshold calculation, `fixed, percentage` of disk space remaining.
         * 
         * @return builder
         * 
         */
        public Builder valueCalculation(@Nullable Output<String> valueCalculation) {
            $.valueCalculation = valueCalculation;
            return this;
        }

        /**
         * @param valueCalculation Disk value threshold calculation, `fixed, percentage` of disk space remaining.
         * 
         * @return builder
         * 
         */
        public Builder valueCalculation(String valueCalculation) {
            return valueCalculation(Output.of(valueCalculation));
        }

        /**
         * @param valueThreshold The value to trigger the alarm for.
         * 
         * @return builder
         * 
         */
        public Builder valueThreshold(@Nullable Output<Integer> valueThreshold) {
            $.valueThreshold = valueThreshold;
            return this;
        }

        /**
         * @param valueThreshold The value to trigger the alarm for.
         * 
         * @return builder
         * 
         */
        public Builder valueThreshold(Integer valueThreshold) {
            return valueThreshold(Output.of(valueThreshold));
        }

        /**
         * @param vhostRegex Regex for which vhost to check
         * 
         * @return builder
         * 
         */
        public Builder vhostRegex(@Nullable Output<String> vhostRegex) {
            $.vhostRegex = vhostRegex;
            return this;
        }

        /**
         * @param vhostRegex Regex for which vhost to check
         * 
         * @return builder
         * 
         */
        public Builder vhostRegex(String vhostRegex) {
            return vhostRegex(Output.of(vhostRegex));
        }

        public AlarmState build() {
            return $;
        }
    }

}
