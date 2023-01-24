// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.cloudamqp;

import com.pulumi.cloudamqp.PluginCommunityArgs;
import com.pulumi.cloudamqp.Utilities;
import com.pulumi.cloudamqp.inputs.PluginCommunityState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import javax.annotation.Nullable;

/**
 * ## Example Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.cloudamqp.PluginCommunity;
 * import com.pulumi.cloudamqp.PluginCommunityArgs;
 * import java.util.List;
 * import java.util.ArrayList;
 * import java.util.Map;
 * import java.io.File;
 * import java.nio.file.Files;
 * import java.nio.file.Paths;
 * 
 * public class App {
 *     public static void main(String[] args) {
 *         Pulumi.run(App::stack);
 *     }
 * 
 *     public static void stack(Context ctx) {
 *         var rabbitmqDelayedMessageExchange = new PluginCommunity(&#34;rabbitmqDelayedMessageExchange&#34;, PluginCommunityArgs.builder()        
 *             .instanceId(cloudamqp_instance.instance_01().id())
 *             .enabled(true)
 *             .build());
 * 
 *     }
 * }
 * ```
 * ## Depedency
 * 
 * This resource depends on CloudAMQP instance identifier, `cloudamqp_instance.instance.id`.
 * 
 * ## Import
 * 
 * `cloudamqp_plugin` can be imported using the name argument of the resource together with CloudAMQP instance identifier. The name and identifier are CSV separated, see example below.
 * 
 * ```sh
 *  $ pulumi import cloudamqp:index/pluginCommunity:PluginCommunity &lt;resource_name&gt; &lt;plugin_name&gt;,&lt;instance_id&gt;`
 * ```
 * 
 */
@ResourceType(type="cloudamqp:index/pluginCommunity:PluginCommunity")
public class PluginCommunity extends com.pulumi.resources.CustomResource {
    /**
     * The description of the plugin.
     * 
     */
    @Export(name="description", type=String.class, parameters={})
    private Output<String> description;

    /**
     * @return The description of the plugin.
     * 
     */
    public Output<String> description() {
        return this.description;
    }
    /**
     * Enable or disable the plugins.
     * 
     */
    @Export(name="enabled", type=Boolean.class, parameters={})
    private Output<Boolean> enabled;

    /**
     * @return Enable or disable the plugins.
     * 
     */
    public Output<Boolean> enabled() {
        return this.enabled;
    }
    /**
     * The CloudAMQP instance ID.
     * 
     */
    @Export(name="instanceId", type=Integer.class, parameters={})
    private Output<Integer> instanceId;

    /**
     * @return The CloudAMQP instance ID.
     * 
     */
    public Output<Integer> instanceId() {
        return this.instanceId;
    }
    /**
     * The name of the Rabbit MQ community plugin.
     * 
     */
    @Export(name="name", type=String.class, parameters={})
    private Output<String> name;

    /**
     * @return The name of the Rabbit MQ community plugin.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * Required version of RabbitMQ.
     * 
     */
    @Export(name="require", type=String.class, parameters={})
    private Output<String> require;

    /**
     * @return Required version of RabbitMQ.
     * 
     */
    public Output<String> require() {
        return this.require;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public PluginCommunity(String name) {
        this(name, PluginCommunityArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public PluginCommunity(String name, PluginCommunityArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public PluginCommunity(String name, PluginCommunityArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("cloudamqp:index/pluginCommunity:PluginCommunity", name, args == null ? PluginCommunityArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private PluginCommunity(String name, Output<String> id, @Nullable PluginCommunityState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("cloudamqp:index/pluginCommunity:PluginCommunity", name, state, makeResourceOptions(options, id));
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
    public static PluginCommunity get(String name, Output<String> id, @Nullable PluginCommunityState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new PluginCommunity(name, id, state, options);
    }
}
