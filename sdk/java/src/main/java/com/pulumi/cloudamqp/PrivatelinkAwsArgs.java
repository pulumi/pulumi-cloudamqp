// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.cloudamqp;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class PrivatelinkAwsArgs extends com.pulumi.resources.ResourceArgs {

    public static final PrivatelinkAwsArgs Empty = new PrivatelinkAwsArgs();

    /**
     * Allowed principals to access the endpoint service.
     * 
     */
    @Import(name="allowedPrincipals", required=true)
    private Output<List<String>> allowedPrincipals;

    /**
     * @return Allowed principals to access the endpoint service.
     * 
     */
    public Output<List<String>> allowedPrincipals() {
        return this.allowedPrincipals;
    }

    /**
     * The CloudAMQP instance identifier.
     * 
     */
    @Import(name="instanceId", required=true)
    private Output<Integer> instanceId;

    /**
     * @return The CloudAMQP instance identifier.
     * 
     */
    public Output<Integer> instanceId() {
        return this.instanceId;
    }

    /**
     * Configurable sleep time (seconds) when enable PrivateLink. Default set to 60 seconds.
     * 
     */
    @Import(name="sleep")
    private @Nullable Output<Integer> sleep;

    /**
     * @return Configurable sleep time (seconds) when enable PrivateLink. Default set to 60 seconds.
     * 
     */
    public Optional<Output<Integer>> sleep() {
        return Optional.ofNullable(this.sleep);
    }

    /**
     * Configurable timeout time (seconds) when enable PrivateLink. Default set to 3600 seconds.
     * 
     */
    @Import(name="timeout")
    private @Nullable Output<Integer> timeout;

    /**
     * @return Configurable timeout time (seconds) when enable PrivateLink. Default set to 3600 seconds.
     * 
     */
    public Optional<Output<Integer>> timeout() {
        return Optional.ofNullable(this.timeout);
    }

    private PrivatelinkAwsArgs() {}

    private PrivatelinkAwsArgs(PrivatelinkAwsArgs $) {
        this.allowedPrincipals = $.allowedPrincipals;
        this.instanceId = $.instanceId;
        this.sleep = $.sleep;
        this.timeout = $.timeout;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(PrivatelinkAwsArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private PrivatelinkAwsArgs $;

        public Builder() {
            $ = new PrivatelinkAwsArgs();
        }

        public Builder(PrivatelinkAwsArgs defaults) {
            $ = new PrivatelinkAwsArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param allowedPrincipals Allowed principals to access the endpoint service.
         * 
         * @return builder
         * 
         */
        public Builder allowedPrincipals(Output<List<String>> allowedPrincipals) {
            $.allowedPrincipals = allowedPrincipals;
            return this;
        }

        /**
         * @param allowedPrincipals Allowed principals to access the endpoint service.
         * 
         * @return builder
         * 
         */
        public Builder allowedPrincipals(List<String> allowedPrincipals) {
            return allowedPrincipals(Output.of(allowedPrincipals));
        }

        /**
         * @param allowedPrincipals Allowed principals to access the endpoint service.
         * 
         * @return builder
         * 
         */
        public Builder allowedPrincipals(String... allowedPrincipals) {
            return allowedPrincipals(List.of(allowedPrincipals));
        }

        /**
         * @param instanceId The CloudAMQP instance identifier.
         * 
         * @return builder
         * 
         */
        public Builder instanceId(Output<Integer> instanceId) {
            $.instanceId = instanceId;
            return this;
        }

        /**
         * @param instanceId The CloudAMQP instance identifier.
         * 
         * @return builder
         * 
         */
        public Builder instanceId(Integer instanceId) {
            return instanceId(Output.of(instanceId));
        }

        /**
         * @param sleep Configurable sleep time (seconds) when enable PrivateLink. Default set to 60 seconds.
         * 
         * @return builder
         * 
         */
        public Builder sleep(@Nullable Output<Integer> sleep) {
            $.sleep = sleep;
            return this;
        }

        /**
         * @param sleep Configurable sleep time (seconds) when enable PrivateLink. Default set to 60 seconds.
         * 
         * @return builder
         * 
         */
        public Builder sleep(Integer sleep) {
            return sleep(Output.of(sleep));
        }

        /**
         * @param timeout Configurable timeout time (seconds) when enable PrivateLink. Default set to 3600 seconds.
         * 
         * @return builder
         * 
         */
        public Builder timeout(@Nullable Output<Integer> timeout) {
            $.timeout = timeout;
            return this;
        }

        /**
         * @param timeout Configurable timeout time (seconds) when enable PrivateLink. Default set to 3600 seconds.
         * 
         * @return builder
         * 
         */
        public Builder timeout(Integer timeout) {
            return timeout(Output.of(timeout));
        }

        public PrivatelinkAwsArgs build() {
            $.allowedPrincipals = Objects.requireNonNull($.allowedPrincipals, "expected parameter 'allowedPrincipals' to be non-null");
            $.instanceId = Objects.requireNonNull($.instanceId, "expected parameter 'instanceId' to be non-null");
            return $;
        }
    }

}
