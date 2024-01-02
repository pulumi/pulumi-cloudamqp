// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.cloudamqp.inputs;

import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Integer;
import java.lang.String;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetNotificationPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetNotificationPlainArgs Empty = new GetNotificationPlainArgs();

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
     * The name set for the recipient.
     * 
     */
    @Import(name="name")
    private @Nullable String name;

    /**
     * @return The name set for the recipient.
     * 
     */
    public Optional<String> name() {
        return Optional.ofNullable(this.name);
    }

    @Import(name="options")
    private @Nullable Map<String,String> options;

    public Optional<Map<String,String>> options() {
        return Optional.ofNullable(this.options);
    }

    /**
     * The recipient identifier.
     * 
     */
    @Import(name="recipientId")
    private @Nullable Integer recipientId;

    /**
     * @return The recipient identifier.
     * 
     */
    public Optional<Integer> recipientId() {
        return Optional.ofNullable(this.recipientId);
    }

    private GetNotificationPlainArgs() {}

    private GetNotificationPlainArgs(GetNotificationPlainArgs $) {
        this.instanceId = $.instanceId;
        this.name = $.name;
        this.options = $.options;
        this.recipientId = $.recipientId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetNotificationPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetNotificationPlainArgs $;

        public Builder() {
            $ = new GetNotificationPlainArgs();
        }

        public Builder(GetNotificationPlainArgs defaults) {
            $ = new GetNotificationPlainArgs(Objects.requireNonNull(defaults));
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
         * @param name The name set for the recipient.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable String name) {
            $.name = name;
            return this;
        }

        public Builder options(@Nullable Map<String,String> options) {
            $.options = options;
            return this;
        }

        /**
         * @param recipientId The recipient identifier.
         * 
         * @return builder
         * 
         */
        public Builder recipientId(@Nullable Integer recipientId) {
            $.recipientId = recipientId;
            return this;
        }

        public GetNotificationPlainArgs build() {
            if ($.instanceId == null) {
                throw new MissingRequiredPropertyException("GetNotificationPlainArgs", "instanceId");
            }
            return $;
        }
    }

}
