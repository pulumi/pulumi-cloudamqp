// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.cloudamqp;

import com.pulumi.cloudamqp.PluginArgs;
import com.pulumi.cloudamqp.Utilities;
import com.pulumi.cloudamqp.inputs.PluginState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * ## Import
 * 
 * `cloudamqp_plugin` can be imported using the name argument of the resource together with CloudAMQP instance identifier. The name and identifier are CSV separated, see example below.
 * 
 * ```sh
 *  $ pulumi import cloudamqp:index/plugin:Plugin rabbitmq_management rabbitmq_management,&lt;instance_id&gt;`
 * ```
 * 
 */
@ResourceType(type="cloudamqp:index/plugin:Plugin")
public class Plugin extends com.pulumi.resources.CustomResource {
    /**
     * The description of the plugin.
     * 
     */
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output<String> description;

    /**
     * @return The description of the plugin.
     * 
     */
    public Output<String> description() {
        return this.description;
    }
    /**
     * If the plugin is enabled
     * 
     */
    @Export(name="enabled", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> enabled;

    /**
     * @return If the plugin is enabled
     * 
     */
    public Output<Boolean> enabled() {
        return this.enabled;
    }
    /**
     * Instance identifier
     * 
     */
    @Export(name="instanceId", refs={Integer.class}, tree="[0]")
    private Output<Integer> instanceId;

    /**
     * @return Instance identifier
     * 
     */
    public Output<Integer> instanceId() {
        return this.instanceId;
    }
    /**
     * The name of the plugin
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return The name of the plugin
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * Configurable sleep time in seconds between retries for plugins
     * 
     */
    @Export(name="sleep", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> sleep;

    /**
     * @return Configurable sleep time in seconds between retries for plugins
     * 
     */
    public Output<Optional<Integer>> sleep() {
        return Codegen.optional(this.sleep);
    }
    /**
     * Configurable timeout time in seconds for plugins
     * 
     */
    @Export(name="timeout", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> timeout;

    /**
     * @return Configurable timeout time in seconds for plugins
     * 
     */
    public Output<Optional<Integer>> timeout() {
        return Codegen.optional(this.timeout);
    }
    /**
     * The version of the plugin.
     * 
     */
    @Export(name="version", refs={String.class}, tree="[0]")
    private Output<String> version;

    /**
     * @return The version of the plugin.
     * 
     */
    public Output<String> version() {
        return this.version;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Plugin(String name) {
        this(name, PluginArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Plugin(String name, PluginArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Plugin(String name, PluginArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("cloudamqp:index/plugin:Plugin", name, args == null ? PluginArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Plugin(String name, Output<String> id, @Nullable PluginState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("cloudamqp:index/plugin:Plugin", name, state, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .build();
        return com.pulumi.resources.CustomResourceOptions.merge(defaultOptions, options, id);
    }

    /**
     * Get an existing Host resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state
     * @param options Optional settings to control the behavior of the CustomResource.
     */
    public static Plugin get(String name, Output<String> id, @Nullable PluginState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Plugin(name, id, state, options);
    }
}
