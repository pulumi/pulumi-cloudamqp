// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.cloudamqp;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;


public final class AccountActionArgs extends com.pulumi.resources.ResourceArgs {

    public static final AccountActionArgs Empty = new AccountActionArgs();

    /**
     * The action to be invoked. Allowed actions `rotate-password`, `rotate-apikey`.
     * 
     */
    @Import(name="action", required=true)
    private Output<String> action;

    /**
     * @return The action to be invoked. Allowed actions `rotate-password`, `rotate-apikey`.
     * 
     */
    public Output<String> action() {
        return this.action;
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

    private AccountActionArgs() {}

    private AccountActionArgs(AccountActionArgs $) {
        this.action = $.action;
        this.instanceId = $.instanceId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(AccountActionArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private AccountActionArgs $;

        public Builder() {
            $ = new AccountActionArgs();
        }

        public Builder(AccountActionArgs defaults) {
            $ = new AccountActionArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param action The action to be invoked. Allowed actions `rotate-password`, `rotate-apikey`.
         * 
         * @return builder
         * 
         */
        public Builder action(Output<String> action) {
            $.action = action;
            return this;
        }

        /**
         * @param action The action to be invoked. Allowed actions `rotate-password`, `rotate-apikey`.
         * 
         * @return builder
         * 
         */
        public Builder action(String action) {
            return action(Output.of(action));
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

        public AccountActionArgs build() {
            if ($.action == null) {
                throw new MissingRequiredPropertyException("AccountActionArgs", "action");
            }
            if ($.instanceId == null) {
                throw new MissingRequiredPropertyException("AccountActionArgs", "instanceId");
            }
            return $;
        }
    }

}