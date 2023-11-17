// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.cloudamqp;

import com.pulumi.cloudamqp.inputs.SecurityFirewallRuleArgs;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Integer;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class SecurityFirewallArgs extends com.pulumi.resources.ResourceArgs {

    public static final SecurityFirewallArgs Empty = new SecurityFirewallArgs();

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
     * An array of rules, minimum of 1 needs to be configured. Each `rules` block consists of the field documented below.
     * 
     */
    @Import(name="rules", required=true)
    private Output<List<SecurityFirewallRuleArgs>> rules;

    /**
     * @return An array of rules, minimum of 1 needs to be configured. Each `rules` block consists of the field documented below.
     * 
     */
    public Output<List<SecurityFirewallRuleArgs>> rules() {
        return this.rules;
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
     * The `rules` block consists of:
     * 
     */
    @Import(name="timeout")
    private @Nullable Output<Integer> timeout;

    /**
     * @return Configurable timeout time in seconds for firewall configuration. Default set to 1800 seconds.
     * 
     * The `rules` block consists of:
     * 
     */
    public Optional<Output<Integer>> timeout() {
        return Optional.ofNullable(this.timeout);
    }

    private SecurityFirewallArgs() {}

    private SecurityFirewallArgs(SecurityFirewallArgs $) {
        this.instanceId = $.instanceId;
        this.rules = $.rules;
        this.sleep = $.sleep;
        this.timeout = $.timeout;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(SecurityFirewallArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private SecurityFirewallArgs $;

        public Builder() {
            $ = new SecurityFirewallArgs();
        }

        public Builder(SecurityFirewallArgs defaults) {
            $ = new SecurityFirewallArgs(Objects.requireNonNull(defaults));
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
         * @param rules An array of rules, minimum of 1 needs to be configured. Each `rules` block consists of the field documented below.
         * 
         * @return builder
         * 
         */
        public Builder rules(Output<List<SecurityFirewallRuleArgs>> rules) {
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
         * The `rules` block consists of:
         * 
         * @return builder
         * 
         */
        public Builder timeout(Integer timeout) {
            return timeout(Output.of(timeout));
        }

        public SecurityFirewallArgs build() {
            $.instanceId = Objects.requireNonNull($.instanceId, "expected parameter 'instanceId' to be non-null");
            $.rules = Objects.requireNonNull($.rules, "expected parameter 'rules' to be non-null");
            return $;
        }
    }

}
