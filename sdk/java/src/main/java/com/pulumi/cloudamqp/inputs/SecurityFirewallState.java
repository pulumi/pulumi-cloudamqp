// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.cloudamqp.inputs;

import com.pulumi.cloudamqp.inputs.SecurityFirewallRuleArgs;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Integer;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class SecurityFirewallState extends com.pulumi.resources.ResourceArgs {

    public static final SecurityFirewallState Empty = new SecurityFirewallState();

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
     * An array of rules, minimum of 1 needs to be configured. Each `rules` block consists of the field documented below.
     * 
     */
    @Import(name="rules")
    private @Nullable Output<List<SecurityFirewallRuleArgs>> rules;

    /**
     * @return An array of rules, minimum of 1 needs to be configured. Each `rules` block consists of the field documented below.
     * 
     */
    public Optional<Output<List<SecurityFirewallRuleArgs>>> rules() {
        return Optional.ofNullable(this.rules);
    }

    /**
     * Configurable sleep time in seconds between retries for firewall configuration. Default set to 30 seconds.
     * 
     */
    @Import(name="sleep")
    private @Nullable Output<Integer> sleep;

    /**
     * @return Configurable sleep time in seconds between retries for firewall configuration. Default set to 30 seconds.
     * 
     */
    public Optional<Output<Integer>> sleep() {
        return Optional.ofNullable(this.sleep);
    }

    /**
     * Configurable timeout time in seconds for firewall configuration. Default set to 1800 seconds.
     * 
     * ***
     * 
     * The `rules` block consists of:
     * 
     */
    @Import(name="timeout")
    private @Nullable Output<Integer> timeout;

    /**
     * @return Configurable timeout time in seconds for firewall configuration. Default set to 1800 seconds.
     * 
     * ***
     * 
     * The `rules` block consists of:
     * 
     */
    public Optional<Output<Integer>> timeout() {
        return Optional.ofNullable(this.timeout);
    }

    private SecurityFirewallState() {}

    private SecurityFirewallState(SecurityFirewallState $) {
        this.instanceId = $.instanceId;
        this.rules = $.rules;
        this.sleep = $.sleep;
        this.timeout = $.timeout;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(SecurityFirewallState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private SecurityFirewallState $;

        public Builder() {
            $ = new SecurityFirewallState();
        }

        public Builder(SecurityFirewallState defaults) {
            $ = new SecurityFirewallState(Objects.requireNonNull(defaults));
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
         * @param rules An array of rules, minimum of 1 needs to be configured. Each `rules` block consists of the field documented below.
         * 
         * @return builder
         * 
         */
        public Builder rules(@Nullable Output<List<SecurityFirewallRuleArgs>> rules) {
            $.rules = rules;
            return this;
        }

        /**
         * @param rules An array of rules, minimum of 1 needs to be configured. Each `rules` block consists of the field documented below.
         * 
         * @return builder
         * 
         */
        public Builder rules(List<SecurityFirewallRuleArgs> rules) {
            return rules(Output.of(rules));
        }

        /**
         * @param rules An array of rules, minimum of 1 needs to be configured. Each `rules` block consists of the field documented below.
         * 
         * @return builder
         * 
         */
        public Builder rules(SecurityFirewallRuleArgs... rules) {
            return rules(List.of(rules));
        }

        /**
         * @param sleep Configurable sleep time in seconds between retries for firewall configuration. Default set to 30 seconds.
         * 
         * @return builder
         * 
         */
        public Builder sleep(@Nullable Output<Integer> sleep) {
            $.sleep = sleep;
            return this;
        }

        /**
         * @param sleep Configurable sleep time in seconds between retries for firewall configuration. Default set to 30 seconds.
         * 
         * @return builder
         * 
         */
        public Builder sleep(Integer sleep) {
            return sleep(Output.of(sleep));
        }

        /**
         * @param timeout Configurable timeout time in seconds for firewall configuration. Default set to 1800 seconds.
         * 
         * ***
         * 
         * The `rules` block consists of:
         * 
         * @return builder
         * 
         */
        public Builder timeout(@Nullable Output<Integer> timeout) {
            $.timeout = timeout;
            return this;
        }

        /**
         * @param timeout Configurable timeout time in seconds for firewall configuration. Default set to 1800 seconds.
         * 
         * ***
         * 
         * The `rules` block consists of:
         * 
         * @return builder
         * 
         */
        public Builder timeout(Integer timeout) {
            return timeout(Output.of(timeout));
        }

        public SecurityFirewallState build() {
            return $;
        }
    }

}
