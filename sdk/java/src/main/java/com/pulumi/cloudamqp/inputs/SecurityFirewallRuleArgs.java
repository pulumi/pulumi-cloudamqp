// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.cloudamqp.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class SecurityFirewallRuleArgs extends com.pulumi.resources.ResourceArgs {

    public static final SecurityFirewallRuleArgs Empty = new SecurityFirewallRuleArgs();

    /**
     * Description name of the rule. e.g. Default.
     * 
     * Pre-defined services for RabbitMQ:
     * 
     * | Service | Port  |
     * |---------|-------|
     * | AMQP    |  5672 |
     * | AMQPS   |  5671 |
     * | HTTPS   |   443 |
     * | MQTT    |  1883 |
     * | MQTTS   |  8883 |
     * | STOMP   | 61613 |
     * | STOMPS  | 61614 |
     * | STREAM  |  5552 |
     * | STREAM_ |  5551 |
     * 
     */
    @Import(name="description")
    private @Nullable Output<String> description;

    /**
     * @return Description name of the rule. e.g. Default.
     * 
     * Pre-defined services for RabbitMQ:
     * 
     * | Service | Port  |
     * |---------|-------|
     * | AMQP    |  5672 |
     * | AMQPS   |  5671 |
     * | HTTPS   |   443 |
     * | MQTT    |  1883 |
     * | MQTTS   |  8883 |
     * | STOMP   | 61613 |
     * | STOMPS  | 61614 |
     * | STREAM  |  5552 |
     * | STREAM_ |  5551 |
     * 
     */
    public Optional<Output<String>> description() {
        return Optional.ofNullable(this.description);
    }

    /**
     * CIDR address: IP address with CIDR notation (e.g. 10.56.72.0/24)
     * 
     */
    @Import(name="ip", required=true)
    private Output<String> ip;

    /**
     * @return CIDR address: IP address with CIDR notation (e.g. 10.56.72.0/24)
     * 
     */
    public Output<String> ip() {
        return this.ip;
    }

    /**
     * Custom ports to be opened
     * 
     */
    @Import(name="ports")
    private @Nullable Output<List<Integer>> ports;

    /**
     * @return Custom ports to be opened
     * 
     */
    public Optional<Output<List<Integer>>> ports() {
        return Optional.ofNullable(this.ports);
    }

    /**
     * Pre-defined service ports, see table below
     * 
     */
    @Import(name="services")
    private @Nullable Output<List<String>> services;

    /**
     * @return Pre-defined service ports, see table below
     * 
     */
    public Optional<Output<List<String>>> services() {
        return Optional.ofNullable(this.services);
    }

    private SecurityFirewallRuleArgs() {}

    private SecurityFirewallRuleArgs(SecurityFirewallRuleArgs $) {
        this.description = $.description;
        this.ip = $.ip;
        this.ports = $.ports;
        this.services = $.services;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(SecurityFirewallRuleArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private SecurityFirewallRuleArgs $;

        public Builder() {
            $ = new SecurityFirewallRuleArgs();
        }

        public Builder(SecurityFirewallRuleArgs defaults) {
            $ = new SecurityFirewallRuleArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param description Description name of the rule. e.g. Default.
         * 
         * Pre-defined services for RabbitMQ:
         * 
         * | Service | Port  |
         * |---------|-------|
         * | AMQP    |  5672 |
         * | AMQPS   |  5671 |
         * | HTTPS   |   443 |
         * | MQTT    |  1883 |
         * | MQTTS   |  8883 |
         * | STOMP   | 61613 |
         * | STOMPS  | 61614 |
         * | STREAM  |  5552 |
         * | STREAM_ |  5551 |
         * 
         * @return builder
         * 
         */
        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        /**
         * @param description Description name of the rule. e.g. Default.
         * 
         * Pre-defined services for RabbitMQ:
         * 
         * | Service | Port  |
         * |---------|-------|
         * | AMQP    |  5672 |
         * | AMQPS   |  5671 |
         * | HTTPS   |   443 |
         * | MQTT    |  1883 |
         * | MQTTS   |  8883 |
         * | STOMP   | 61613 |
         * | STOMPS  | 61614 |
         * | STREAM  |  5552 |
         * | STREAM_ |  5551 |
         * 
         * @return builder
         * 
         */
        public Builder description(String description) {
            return description(Output.of(description));
        }

        /**
         * @param ip CIDR address: IP address with CIDR notation (e.g. 10.56.72.0/24)
         * 
         * @return builder
         * 
         */
        public Builder ip(Output<String> ip) {
            $.ip = ip;
            return this;
        }

        /**
         * @param ip CIDR address: IP address with CIDR notation (e.g. 10.56.72.0/24)
         * 
         * @return builder
         * 
         */
        public Builder ip(String ip) {
            return ip(Output.of(ip));
        }

        /**
         * @param ports Custom ports to be opened
         * 
         * @return builder
         * 
         */
        public Builder ports(@Nullable Output<List<Integer>> ports) {
            $.ports = ports;
            return this;
        }

        /**
         * @param ports Custom ports to be opened
         * 
         * @return builder
         * 
         */
        public Builder ports(List<Integer> ports) {
            return ports(Output.of(ports));
        }

        /**
         * @param ports Custom ports to be opened
         * 
         * @return builder
         * 
         */
        public Builder ports(Integer... ports) {
            return ports(List.of(ports));
        }

        /**
         * @param services Pre-defined service ports, see table below
         * 
         * @return builder
         * 
         */
        public Builder services(@Nullable Output<List<String>> services) {
            $.services = services;
            return this;
        }

        /**
         * @param services Pre-defined service ports, see table below
         * 
         * @return builder
         * 
         */
        public Builder services(List<String> services) {
            return services(Output.of(services));
        }

        /**
         * @param services Pre-defined service ports, see table below
         * 
         * @return builder
         * 
         */
        public Builder services(String... services) {
            return services(List.of(services));
        }

        public SecurityFirewallRuleArgs build() {
            if ($.ip == null) {
                throw new MissingRequiredPropertyException("SecurityFirewallRuleArgs", "ip");
            }
            return $;
        }
    }

}
