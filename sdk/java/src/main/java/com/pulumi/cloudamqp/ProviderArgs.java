// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.cloudamqp;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ProviderArgs extends com.pulumi.resources.ResourceArgs {

    public static final ProviderArgs Empty = new ProviderArgs();

    /**
     * Key used to authentication to the CloudAMQP Customer API
     * 
     */
    @Import(name="apikey", required=true)
    private Output<String> apikey;

    /**
     * @return Key used to authentication to the CloudAMQP Customer API
     * 
     */
    public Output<String> apikey() {
        return this.apikey;
    }

    /**
     * Base URL to CloudAMQP Customer website
     * 
     */
    @Import(name="baseurl")
    private @Nullable Output<String> baseurl;

    /**
     * @return Base URL to CloudAMQP Customer website
     * 
     */
    public Optional<Output<String>> baseurl() {
        return Optional.ofNullable(this.baseurl);
    }

    /**
     * Skips destroying backend resources on &#39;terraform destroy&#39;
     * 
     */
    @Import(name="enableFasterInstanceDestroy", json=true)
    private @Nullable Output<Boolean> enableFasterInstanceDestroy;

    /**
     * @return Skips destroying backend resources on &#39;terraform destroy&#39;
     * 
     */
    public Optional<Output<Boolean>> enableFasterInstanceDestroy() {
        return Optional.ofNullable(this.enableFasterInstanceDestroy);
    }

    private ProviderArgs() {}

    private ProviderArgs(ProviderArgs $) {
        this.apikey = $.apikey;
        this.baseurl = $.baseurl;
        this.enableFasterInstanceDestroy = $.enableFasterInstanceDestroy;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ProviderArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ProviderArgs $;

        public Builder() {
            $ = new ProviderArgs();
        }

        public Builder(ProviderArgs defaults) {
            $ = new ProviderArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param apikey Key used to authentication to the CloudAMQP Customer API
         * 
         * @return builder
         * 
         */
        public Builder apikey(Output<String> apikey) {
            $.apikey = apikey;
            return this;
        }

        /**
         * @param apikey Key used to authentication to the CloudAMQP Customer API
         * 
         * @return builder
         * 
         */
        public Builder apikey(String apikey) {
            return apikey(Output.of(apikey));
        }

        /**
         * @param baseurl Base URL to CloudAMQP Customer website
         * 
         * @return builder
         * 
         */
        public Builder baseurl(@Nullable Output<String> baseurl) {
            $.baseurl = baseurl;
            return this;
        }

        /**
         * @param baseurl Base URL to CloudAMQP Customer website
         * 
         * @return builder
         * 
         */
        public Builder baseurl(String baseurl) {
            return baseurl(Output.of(baseurl));
        }

        /**
         * @param enableFasterInstanceDestroy Skips destroying backend resources on &#39;terraform destroy&#39;
         * 
         * @return builder
         * 
         */
        public Builder enableFasterInstanceDestroy(@Nullable Output<Boolean> enableFasterInstanceDestroy) {
            $.enableFasterInstanceDestroy = enableFasterInstanceDestroy;
            return this;
        }

        /**
         * @param enableFasterInstanceDestroy Skips destroying backend resources on &#39;terraform destroy&#39;
         * 
         * @return builder
         * 
         */
        public Builder enableFasterInstanceDestroy(Boolean enableFasterInstanceDestroy) {
            return enableFasterInstanceDestroy(Output.of(enableFasterInstanceDestroy));
        }

        public ProviderArgs build() {
            $.apikey = Objects.requireNonNull($.apikey, "expected parameter 'apikey' to be non-null");
            return $;
        }
    }

}
