// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.cloudamqp.inputs;

import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetAlarmPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetAlarmPlainArgs Empty = new GetAlarmPlainArgs();

    /**
     * The alarm identifier. Either use this or `type` to give `cloudamqp.Alarm` necessary information to retrieve the alarm.
     * 
     */
    @Import(name="alarmId")
    private @Nullable Integer alarmId;

    /**
     * @return The alarm identifier. Either use this or `type` to give `cloudamqp.Alarm` necessary information to retrieve the alarm.
     * 
     */
    public Optional<Integer> alarmId() {
        return Optional.ofNullable(this.alarmId);
    }

    /**
     * The CloudAMQP instance identifier.
     * 
     */
    @Import(name="instanceId", required=true)
    private Integer instanceId;

    /**
     * @return The CloudAMQP instance identifier.
     * 
     */
    public Integer instanceId() {
        return this.instanceId;
    }

    /**
     * The alarm type. Either use this or `alarm_id` to give `cloudamqp.Alarm` necessary information when retrieve the alarm. Supported alarm types
     * 
     */
    @Import(name="type")
    private @Nullable String type;

    /**
     * @return The alarm type. Either use this or `alarm_id` to give `cloudamqp.Alarm` necessary information when retrieve the alarm. Supported alarm types
     * 
     */
    public Optional<String> type() {
        return Optional.ofNullable(this.type);
    }

    @Import(name="valueCalculation")
    private @Nullable String valueCalculation;

    public Optional<String> valueCalculation() {
        return Optional.ofNullable(this.valueCalculation);
    }

    private GetAlarmPlainArgs() {}

    private GetAlarmPlainArgs(GetAlarmPlainArgs $) {
        this.alarmId = $.alarmId;
        this.instanceId = $.instanceId;
        this.type = $.type;
        this.valueCalculation = $.valueCalculation;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetAlarmPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetAlarmPlainArgs $;

        public Builder() {
            $ = new GetAlarmPlainArgs();
        }

        public Builder(GetAlarmPlainArgs defaults) {
            $ = new GetAlarmPlainArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param alarmId The alarm identifier. Either use this or `type` to give `cloudamqp.Alarm` necessary information to retrieve the alarm.
         * 
         * @return builder
         * 
         */
        public Builder alarmId(@Nullable Integer alarmId) {
            $.alarmId = alarmId;
            return this;
        }

        /**
         * @param instanceId The CloudAMQP instance identifier.
         * 
         * @return builder
         * 
         */
        public Builder instanceId(Integer instanceId) {
            $.instanceId = instanceId;
            return this;
        }

        /**
         * @param type The alarm type. Either use this or `alarm_id` to give `cloudamqp.Alarm` necessary information when retrieve the alarm. Supported alarm types
         * 
         * @return builder
         * 
         */
        public Builder type(@Nullable String type) {
            $.type = type;
            return this;
        }

        public Builder valueCalculation(@Nullable String valueCalculation) {
            $.valueCalculation = valueCalculation;
            return this;
        }

        public GetAlarmPlainArgs build() {
            if ($.instanceId == null) {
                throw new MissingRequiredPropertyException("GetAlarmPlainArgs", "instanceId");
            }
            return $;
        }
    }

}
