// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.cloudamqp.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class SecurityFirewallRule {
    /**
     * @return Description name of the rule. e.g. Default.
     * 
     */
    private @Nullable String description;
    /**
     * @return CIDR address: IP address with CIDR notation (e.g. 10.56.72.0/24)
     * 
     */
    private String ip;
    /**
     * @return Custom ports to be opened
     * 
     */
    private @Nullable List<Integer> ports;
    /**
     * @return Pre-defined service ports, see table below
     * 
     */
    private @Nullable List<String> services;

    private SecurityFirewallRule() {}
    /**
     * @return Description name of the rule. e.g. Default.
     * 
     */
    public Optional<String> description() {
        return Optional.ofNullable(this.description);
    }
    /**
     * @return CIDR address: IP address with CIDR notation (e.g. 10.56.72.0/24)
     * 
     */
    public String ip() {
        return this.ip;
    }
    /**
     * @return Custom ports to be opened
     * 
     */
    public List<Integer> ports() {
        return this.ports == null ? List.of() : this.ports;
    }
    /**
     * @return Pre-defined service ports, see table below
     * 
     */
    public List<String> services() {
        return this.services == null ? List.of() : this.services;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(SecurityFirewallRule defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable String description;
        private String ip;
        private @Nullable List<Integer> ports;
        private @Nullable List<String> services;
        public Builder() {}
        public Builder(SecurityFirewallRule defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.description = defaults.description;
    	      this.ip = defaults.ip;
    	      this.ports = defaults.ports;
    	      this.services = defaults.services;
        }

        @CustomType.Setter
        public Builder description(@Nullable String description) {
            this.description = description;
            return this;
        }
        @CustomType.Setter
        public Builder ip(String ip) {
            this.ip = Objects.requireNonNull(ip);
            return this;
        }
        @CustomType.Setter
        public Builder ports(@Nullable List<Integer> ports) {
            this.ports = ports;
            return this;
        }
        public Builder ports(Integer... ports) {
            return ports(List.of(ports));
        }
        @CustomType.Setter
        public Builder services(@Nullable List<String> services) {
            this.services = services;
            return this;
        }
        public Builder services(String... services) {
            return services(List.of(services));
        }
        public SecurityFirewallRule build() {
            final var o = new SecurityFirewallRule();
            o.description = description;
            o.ip = ip;
            o.ports = ports;
            o.services = services;
            return o;
        }
    }
}