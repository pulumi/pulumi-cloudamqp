// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.cloudamqp.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class GetAlarmResult {
    private @Nullable Integer alarmId;
    private Boolean enabled;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    private Integer instanceId;
    private String messageType;
    private String queueRegex;
    private List<Integer> recipients;
    private Integer reminderInterval;
    private Integer timeThreshold;
    private @Nullable String type;
    private @Nullable String valueCalculation;
    private Integer valueThreshold;
    private String vhostRegex;

    private GetAlarmResult() {}
    public Optional<Integer> alarmId() {
        return Optional.ofNullable(this.alarmId);
    }
    public Boolean enabled() {
        return this.enabled;
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
    public String messageType() {
        return this.messageType;
    }
    public String queueRegex() {
        return this.queueRegex;
    }
    public List<Integer> recipients() {
        return this.recipients;
    }
    public Integer reminderInterval() {
        return this.reminderInterval;
    }
    public Integer timeThreshold() {
        return this.timeThreshold;
    }
    public Optional<String> type() {
        return Optional.ofNullable(this.type);
    }
    public Optional<String> valueCalculation() {
        return Optional.ofNullable(this.valueCalculation);
    }
    public Integer valueThreshold() {
        return this.valueThreshold;
    }
    public String vhostRegex() {
        return this.vhostRegex;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetAlarmResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable Integer alarmId;
        private Boolean enabled;
        private String id;
        private Integer instanceId;
        private String messageType;
        private String queueRegex;
        private List<Integer> recipients;
        private Integer reminderInterval;
        private Integer timeThreshold;
        private @Nullable String type;
        private @Nullable String valueCalculation;
        private Integer valueThreshold;
        private String vhostRegex;
        public Builder() {}
        public Builder(GetAlarmResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.alarmId = defaults.alarmId;
    	      this.enabled = defaults.enabled;
    	      this.id = defaults.id;
    	      this.instanceId = defaults.instanceId;
    	      this.messageType = defaults.messageType;
    	      this.queueRegex = defaults.queueRegex;
    	      this.recipients = defaults.recipients;
    	      this.reminderInterval = defaults.reminderInterval;
    	      this.timeThreshold = defaults.timeThreshold;
    	      this.type = defaults.type;
    	      this.valueCalculation = defaults.valueCalculation;
    	      this.valueThreshold = defaults.valueThreshold;
    	      this.vhostRegex = defaults.vhostRegex;
        }

        @CustomType.Setter
        public Builder alarmId(@Nullable Integer alarmId) {
            this.alarmId = alarmId;
            return this;
        }
        @CustomType.Setter
        public Builder enabled(Boolean enabled) {
            this.enabled = Objects.requireNonNull(enabled);
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        @CustomType.Setter
        public Builder instanceId(Integer instanceId) {
            this.instanceId = Objects.requireNonNull(instanceId);
            return this;
        }
        @CustomType.Setter
        public Builder messageType(String messageType) {
            this.messageType = Objects.requireNonNull(messageType);
            return this;
        }
        @CustomType.Setter
        public Builder queueRegex(String queueRegex) {
            this.queueRegex = Objects.requireNonNull(queueRegex);
            return this;
        }
        @CustomType.Setter
        public Builder recipients(List<Integer> recipients) {
            this.recipients = Objects.requireNonNull(recipients);
            return this;
        }
        public Builder recipients(Integer... recipients) {
            return recipients(List.of(recipients));
        }
        @CustomType.Setter
        public Builder reminderInterval(Integer reminderInterval) {
            this.reminderInterval = Objects.requireNonNull(reminderInterval);
            return this;
        }
        @CustomType.Setter
        public Builder timeThreshold(Integer timeThreshold) {
            this.timeThreshold = Objects.requireNonNull(timeThreshold);
            return this;
        }
        @CustomType.Setter
        public Builder type(@Nullable String type) {
            this.type = type;
            return this;
        }
        @CustomType.Setter
        public Builder valueCalculation(@Nullable String valueCalculation) {
            this.valueCalculation = valueCalculation;
            return this;
        }
        @CustomType.Setter
        public Builder valueThreshold(Integer valueThreshold) {
            this.valueThreshold = Objects.requireNonNull(valueThreshold);
            return this;
        }
        @CustomType.Setter
        public Builder vhostRegex(String vhostRegex) {
            this.vhostRegex = Objects.requireNonNull(vhostRegex);
            return this;
        }
        public GetAlarmResult build() {
            final var _resultValue = new GetAlarmResult();
            _resultValue.alarmId = alarmId;
            _resultValue.enabled = enabled;
            _resultValue.id = id;
            _resultValue.instanceId = instanceId;
            _resultValue.messageType = messageType;
            _resultValue.queueRegex = queueRegex;
            _resultValue.recipients = recipients;
            _resultValue.reminderInterval = reminderInterval;
            _resultValue.timeThreshold = timeThreshold;
            _resultValue.type = type;
            _resultValue.valueCalculation = valueCalculation;
            _resultValue.valueThreshold = valueThreshold;
            _resultValue.vhostRegex = vhostRegex;
            return _resultValue;
        }
    }
}
