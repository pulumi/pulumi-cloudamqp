// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.cloudamqp.outputs;

import com.pulumi.cloudamqp.outputs.GetAlarmsAlarm;
import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class GetAlarmsResult {
    /**
     * @return List of alarms (see below for nested schema)
     * 
     */
    private List<GetAlarmsAlarm> alarms;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    private Integer instanceId;
    /**
     * @return The type of the alarm.
     * 
     */
    private @Nullable String type;

    private GetAlarmsResult() {}
    /**
     * @return List of alarms (see below for nested schema)
     * 
     */
    public List<GetAlarmsAlarm> alarms() {
        return this.alarms;
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
     * @return The type of the alarm.
     * 
     */
    public Optional<String> type() {
        return Optional.ofNullable(this.type);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetAlarmsResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private List<GetAlarmsAlarm> alarms;
        private String id;
        private Integer instanceId;
        private @Nullable String type;
        public Builder() {}
        public Builder(GetAlarmsResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.alarms = defaults.alarms;
    	      this.id = defaults.id;
    	      this.instanceId = defaults.instanceId;
    	      this.type = defaults.type;
        }

        @CustomType.Setter
        public Builder alarms(List<GetAlarmsAlarm> alarms) {
            if (alarms == null) {
              throw new MissingRequiredPropertyException("GetAlarmsResult", "alarms");
            }
            this.alarms = alarms;
            return this;
        }
        public Builder alarms(GetAlarmsAlarm... alarms) {
            return alarms(List.of(alarms));
        }
        @CustomType.Setter
        public Builder id(String id) {
            if (id == null) {
              throw new MissingRequiredPropertyException("GetAlarmsResult", "id");
            }
            this.id = id;
            return this;
        }
        @CustomType.Setter
        public Builder instanceId(Integer instanceId) {
            if (instanceId == null) {
              throw new MissingRequiredPropertyException("GetAlarmsResult", "instanceId");
            }
            this.instanceId = instanceId;
            return this;
        }
        @CustomType.Setter
        public Builder type(@Nullable String type) {

            this.type = type;
            return this;
        }
        public GetAlarmsResult build() {
            final var _resultValue = new GetAlarmsResult();
            _resultValue.alarms = alarms;
            _resultValue.id = id;
            _resultValue.instanceId = instanceId;
            _resultValue.type = type;
            return _resultValue;
        }
    }
}
