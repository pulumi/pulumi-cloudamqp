// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.cloudamqp.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class AccountActionsState extends com.pulumi.resources.ResourceArgs {

    public static final AccountActionsState Empty = new AccountActionsState();

    /**
     * The action to perform on the node
     * 
     */
    @Import(name="action")
    private @Nullable Output<String> action;

    /**
     * @return The action to perform on the node
     * 
     */
    public Optional<Output<String>> action() {
        return Optional.ofNullable(this.action);
    }

    /**
     * Instance identifier
     * 
     */
    @Import(name="instanceId")
    private @Nullable Output<Integer> instanceId;

    /**
     * @return Instance identifier
     * 
     */
    public Optional<Output<Integer>> instanceId() {
        return Optional.ofNullable(this.instanceId);
    }

    private AccountActionsState() {}

    private AccountActionsState(AccountActionsState $) {
        this.action = $.action;
        this.instanceId = $.instanceId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(AccountActionsState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private AccountActionsState $;

        public Builder() {
            $ = new AccountActionsState();
        }

        public Builder(AccountActionsState defaults) {
            $ = new AccountActionsState(Objects.requireNonNull(defaults));
        }

        /**
         * @param action The action to perform on the node
         * 
         * @return builder
         * 
         */
        public Builder action(@Nullable Output<String> action) {
            $.action = action;
            return this;
        }

        /**
         * @param action The action to perform on the node
         * 
         * @return builder
         * 
         */
        public Builder action(String action) {
            return action(Output.of(action));
        }

        /**
         * @param instanceId Instance identifier
         * 
         * @return builder
         * 
         */
        public Builder instanceId(@Nullable Output<Integer> instanceId) {
            $.instanceId = instanceId;
            return this;
        }

        /**
         * @param instanceId Instance identifier
         * 
         * @return builder
         * 
         */
        public Builder instanceId(Integer instanceId) {
            return instanceId(Output.of(instanceId));
        }

        public AccountActionsState build() {
            return $;
        }
    }

}
