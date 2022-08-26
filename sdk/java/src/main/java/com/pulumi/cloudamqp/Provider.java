// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.cloudamqp;

import com.pulumi.cloudamqp.ProviderArgs;
import com.pulumi.cloudamqp.Utilities;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * The provider type for the cloudamqp package. By default, resources use package-wide configuration
 * settings, however an explicit `Provider` instance may be created and passed during resource
 * construction to achieve fine-grained programmatic control over provider settings. See the
 * [documentation](https://www.pulumi.com/docs/reference/programming-model/#providers) for more information.
 * 
 */
@ResourceType(type="pulumi:providers:cloudamqp")
public class Provider extends com.pulumi.resources.ProviderResource {
    /**
     * Key used to authentication to the CloudAMQP Customer API
     * 
     */
    @Export(name="apikey", type=String.class, parameters={})
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
    @Export(name="baseurl", type=String.class, parameters={})
    private Output</* @Nullable */ String> baseurl;

    /**
     * @return Base URL to CloudAMQP Customer website
     * 
     */
    public Output<Optional<String>> baseurl() {
        return Codegen.optional(this.baseurl);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Provider(String name) {
        this(name, ProviderArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Provider(String name, ProviderArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Provider(String name, ProviderArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("cloudamqp", name, args == null ? ProviderArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .build();
        return com.pulumi.resources.CustomResourceOptions.merge(defaultOptions, options, id);
    }

}
