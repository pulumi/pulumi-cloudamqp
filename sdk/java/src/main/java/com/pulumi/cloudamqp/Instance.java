// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.cloudamqp;

import com.pulumi.cloudamqp.InstanceArgs;
import com.pulumi.cloudamqp.Utilities;
import com.pulumi.cloudamqp.inputs.InstanceState;
import com.pulumi.cloudamqp.outputs.InstanceCopySetting;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * This resource allows you to create and manage a CloudAMQP instance running either [**RabbitMQ**](https://www.rabbitmq.com/) or [**LavinMQ**](https://lavinmq.com/) and can be deployed to multiple cloud platforms provider and regions, see instance regions for more information.
 * 
 * Once the instance is created it will be assigned a unique identifier. All other resources and data sources created for this instance needs to reference this unique instance identifier.
 * 
 * Pricing is available at [cloudamqp.com](https://www.cloudamqp.com/plans.html).
 * 
 * ## Example Usage
 * 
 * &lt;details&gt;
 *   &lt;summary&gt;
 *     &lt;b&gt;
 *       &lt;i&gt;Basic example of shared and dedicated instances&lt;/i&gt;
 *     &lt;/b&gt;
 *   &lt;/summary&gt;
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.cloudamqp.Instance;
 * import com.pulumi.cloudamqp.InstanceArgs;
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
 *         // Minimum free lemur instance running RabbitMQ
 *         var lemurInstance = new Instance("lemurInstance", InstanceArgs.builder()
 *             .name("cloudamqp-free-instance")
 *             .plan("lemur")
 *             .region("amazon-web-services::us-west-1")
 *             .tags("rabbitmq")
 *             .build());
 * 
 *         // Minimum free lemming instance running LavinMQ
 *         var lemmingInstance = new Instance("lemmingInstance", InstanceArgs.builder()
 *             .name("cloudamqp-free-instance")
 *             .plan("lemming")
 *             .region("amazon-web-services::us-west-1")
 *             .tags("lavinmq")
 *             .build());
 * 
 *         // New dedicated bunny instance running RabbitMQ
 *         var instance = new Instance("instance", InstanceArgs.builder()
 *             .name("terraform-cloudamqp-instance")
 *             .plan("bunny-1")
 *             .region("amazon-web-services::us-west-1")
 *             .tags("terraform")
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * &lt;/details&gt;
 * 
 * &lt;details&gt;
 *   &lt;summary&gt;
 *     &lt;b&gt;
 *       &lt;i&gt;Dedicated instance using attribute vpc_subnet to create VPC, before v1.16.0&lt;/i&gt;
 *     &lt;/b&gt;
 *   &lt;/summary&gt;
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.cloudamqp.Instance;
 * import com.pulumi.cloudamqp.InstanceArgs;
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
 *         var instance = new Instance("instance", InstanceArgs.builder()
 *             .name("terraform-cloudamqp-instance")
 *             .plan("bunny-1")
 *             .region("amazon-web-services::us-west-1")
 *             .tags("terraform")
 *             .vpcSubnet("10.56.72.0/24")
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * &lt;/details&gt;
 * 
 * &lt;details&gt;
 *   &lt;summary&gt;
 *     &lt;b&gt;
 *       &lt;i&gt;Dedicated instance using attribute vpc_subnet to create VPC and then import managed VPC, from v1.16.0 (Managed VPC)&lt;/i&gt;
 *     &lt;/b&gt;
 *   &lt;/summary&gt;
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.cloudamqp.Instance;
 * import com.pulumi.cloudamqp.InstanceArgs;
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
 *         // Dedicated instance that also creates VPC
 *         var instance01 = new Instance("instance01", InstanceArgs.builder()
 *             .name("terraform-cloudamqp-instance-01")
 *             .plan("bunny-1")
 *             .region("amazon-web-services::us-west-1")
 *             .tags("terraform")
 *             .vpcSubnet("10.56.72.0/24")
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * Once the instance and the VPC are created, the VPC can be imported as managed VPC and added to the configuration file.
 * Set attribute `vpc_id` to the managed VPC identifier. To keep the managed VPC when deleting the instance, set attribute `keep_associated_vpc` to true.
 * For more information see guide Managed VPC.
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.cloudamqp.Vpc;
 * import com.pulumi.cloudamqp.VpcArgs;
 * import com.pulumi.cloudamqp.Instance;
 * import com.pulumi.cloudamqp.InstanceArgs;
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
 *         // Imported managed VPC
 *         var vpc = new Vpc("vpc", VpcArgs.builder()
 *             .name("<vpc-name>")
 *             .region("amazon-web-services::us-east-1")
 *             .subnet("10.56.72.0/24")
 *             .tags()
 *             .build());
 * 
 *         // Add vpc_id and keep_associated_vpc attributes
 *         var instance01 = new Instance("instance01", InstanceArgs.builder()
 *             .name("terraform-cloudamqp-instance-01")
 *             .plan("bunny-1")
 *             .region("amazon-web-services::us-west-1")
 *             .tags("terraform")
 *             .vpcId(vpc.id())
 *             .keepAssociatedVpc(true)
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * &lt;/details&gt;
 * 
 * &lt;details&gt;
 *   &lt;summary&gt;
 *     &lt;b&gt;
 *       &lt;i&gt;Dedicated instances and managed VPC, from v1.16.0 (Managed VPC)&lt;/i&gt;
 *     &lt;/b&gt;
 *   &lt;/summary&gt;
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.cloudamqp.Vpc;
 * import com.pulumi.cloudamqp.VpcArgs;
 * import com.pulumi.cloudamqp.Instance;
 * import com.pulumi.cloudamqp.InstanceArgs;
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
 *         // Managed VPC
 *         var vpc = new Vpc("vpc", VpcArgs.builder()
 *             .name("<vpc-name>")
 *             .region("amazon-web-services::us-east-1")
 *             .subnet("10.56.72.0/24")
 *             .tags()
 *             .build());
 * 
 *         // First instance added to managed VPC
 *         var instance01 = new Instance("instance01", InstanceArgs.builder()
 *             .name("terraform-cloudamqp-instance-01")
 *             .plan("bunny-1")
 *             .region("amazon-web-services::us-west-1")
 *             .tags("terraform")
 *             .vpcId(vpc.id())
 *             .keepAssociatedVpc(true)
 *             .build());
 * 
 *         // Second instance added to managed VPC
 *         var instance02 = new Instance("instance02", InstanceArgs.builder()
 *             .name("terraform-cloudamqp-instance-02")
 *             .plan("bunny-1")
 *             .region("amazon-web-services::us-west-1")
 *             .tags("terraform")
 *             .vpcId(vpc.id())
 *             .keepAssociatedVpc(true)
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * Set attribute `keep_associated_vpc` to true, will keep managed VPC when deleting the instances.
 * 
 * &lt;/details&gt;
 * 
 * ## Copy settings to a new dedicated instance
 * 
 * With copy settings it&#39;s possible to create a new dedicated instance with settings such as alarms, config, etc. from another dedicated instance. This can be done by adding the `copy_settings` block to this resource and populate `subscription_id` with a CloudAMQP instance identifier from another already existing instance.
 * 
 * Then add the settings to be copied over to the new dedicated instance. Settings that can be copied [alarms, config, definitions, firewall, logs, metrics, plugins]
 * 
 * &gt; `rmq_version` argument is required when doing this action. Must match the RabbitMQ version of the dedicated instance to be copied from.
 * 
 * &lt;details&gt;
 *   &lt;summary&gt;
 *     &lt;b&gt;
 *       &lt;i&gt;Copy settings from a dedicated instance to a new dedicated instance&lt;/i&gt;
 *     &lt;/b&gt;
 *   &lt;/summary&gt;
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.cloudamqp.Instance;
 * import com.pulumi.cloudamqp.InstanceArgs;
 * import com.pulumi.cloudamqp.inputs.InstanceCopySettingArgs;
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
 *         var instance02 = new Instance("instance02", InstanceArgs.builder()
 *             .name("terraform-cloudamqp-instance-02")
 *             .plan("squirrel-1")
 *             .region("amazon-web-services::us-west-1")
 *             .rmqVersion("3.12.2")
 *             .tags("terraform")
 *             .copySettings(InstanceCopySettingArgs.builder()
 *                 .subscriptionId(instanceId)
 *                 .settings(                
 *                     "alarms",
 *                     "config",
 *                     "definitions",
 *                     "firewall",
 *                     "logs",
 *                     "metrics",
 *                     "plugins")
 *                 .build())
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * &lt;/details&gt;
 * 
 * ## Import
 * 
 * `cloudamqp_instance`can be imported using CloudAMQP internal identifier.
 * 
 * ```sh
 * $ pulumi import cloudamqp:index/instance:Instance instance &lt;id&gt;`
 * ```
 * 
 * To retrieve the identifier for an instance, either use [CloudAMQP customer API](https://docs.cloudamqp.com/#list-instances) or use the data source [`cloudamqp_account`](./data-sources/account.md) to list all available instances for an account.
 * 
 */
@ResourceType(type="cloudamqp:index/instance:Instance")
public class Instance extends com.pulumi.resources.CustomResource {
    /**
     * API key needed to communicate to CloudAMQP&#39;s second API. The second API is used to manage alarms, integration and more, full description [CloudAMQP API](https://docs.cloudamqp.com/cloudamqp_api.html).
     * 
     */
    @Export(name="apikey", refs={String.class}, tree="[0]")
    private Output<String> apikey;

    /**
     * @return API key needed to communicate to CloudAMQP&#39;s second API. The second API is used to manage alarms, integration and more, full description [CloudAMQP API](https://docs.cloudamqp.com/cloudamqp_api.html).
     * 
     */
    public Output<String> apikey() {
        return this.apikey;
    }
    /**
     * Information if the CloudAMQP instance runs either RabbitMQ or LavinMQ.
     * 
     */
    @Export(name="backend", refs={String.class}, tree="[0]")
    private Output<String> backend;

    /**
     * @return Information if the CloudAMQP instance runs either RabbitMQ or LavinMQ.
     * 
     */
    public Output<String> backend() {
        return this.backend;
    }
    /**
     * Copy settings from one CloudAMQP instance to a new. Consists of the block documented below.
     * 
     */
    @Export(name="copySettings", refs={List.class,InstanceCopySetting.class}, tree="[0,1]")
    private Output</* @Nullable */ List<InstanceCopySetting>> copySettings;

    /**
     * @return Copy settings from one CloudAMQP instance to a new. Consists of the block documented below.
     * 
     */
    public Output<Optional<List<InstanceCopySetting>>> copySettings() {
        return Codegen.optional(this.copySettings);
    }
    /**
     * Information if the CloudAMQP instance is shared or dedicated.
     * 
     */
    @Export(name="dedicated", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> dedicated;

    /**
     * @return Information if the CloudAMQP instance is shared or dedicated.
     * 
     */
    public Output<Boolean> dedicated() {
        return this.dedicated;
    }
    /**
     * The external hostname for the CloudAMQP instance.
     * 
     */
    @Export(name="host", refs={String.class}, tree="[0]")
    private Output<String> host;

    /**
     * @return The external hostname for the CloudAMQP instance.
     * 
     */
    public Output<String> host() {
        return this.host;
    }
    /**
     * The internal hostname for the CloudAMQP instance.
     * 
     */
    @Export(name="hostInternal", refs={String.class}, tree="[0]")
    private Output<String> hostInternal;

    /**
     * @return The internal hostname for the CloudAMQP instance.
     * 
     */
    public Output<String> hostInternal() {
        return this.hostInternal;
    }
    /**
     * Keep associated VPC when deleting instance, default set to false.
     * 
     */
    @Export(name="keepAssociatedVpc", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> keepAssociatedVpc;

    /**
     * @return Keep associated VPC when deleting instance, default set to false.
     * 
     */
    public Output<Optional<Boolean>> keepAssociatedVpc() {
        return Codegen.optional(this.keepAssociatedVpc);
    }
    /**
     * Name of the CloudAMQP instance.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return Name of the CloudAMQP instance.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * Set to true to discard creating default alarms when the instance is created. Can be left out, will then use default value = false.
     * 
     */
    @Export(name="noDefaultAlarms", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> noDefaultAlarms;

    /**
     * @return Set to true to discard creating default alarms when the instance is created. Can be left out, will then use default value = false.
     * 
     */
    public Output<Boolean> noDefaultAlarms() {
        return this.noDefaultAlarms;
    }
    /**
     * Number of nodes, 1, 3 or 5 depending on plan used. Only needed for legacy plans, will otherwise be computed.
     * 
     * ***Deprecated: Legacy subscriptions plan can still change this to scale up or down the instance. New subscriptions plans use the plan to determine number of nodes. In order to change number of nodes the `plan` needs to be updated.***
     * 
     */
    @Export(name="nodes", refs={Integer.class}, tree="[0]")
    private Output<Integer> nodes;

    /**
     * @return Number of nodes, 1, 3 or 5 depending on plan used. Only needed for legacy plans, will otherwise be computed.
     * 
     * ***Deprecated: Legacy subscriptions plan can still change this to scale up or down the instance. New subscriptions plans use the plan to determine number of nodes. In order to change number of nodes the `plan` needs to be updated.***
     * 
     */
    public Output<Integer> nodes() {
        return this.nodes;
    }
    /**
     * The subscription plan. See available plans
     * 
     */
    @Export(name="plan", refs={String.class}, tree="[0]")
    private Output<String> plan;

    /**
     * @return The subscription plan. See available plans
     * 
     */
    public Output<String> plan() {
        return this.plan;
    }
    /**
     * Flag describing if the resource is ready
     * 
     */
    @Export(name="ready", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> ready;

    /**
     * @return Flag describing if the resource is ready
     * 
     */
    public Output<Boolean> ready() {
        return this.ready;
    }
    /**
     * The region to host the instance in. See instance regions
     * 
     * ***Note: Changing region will force the instance to be destroyed and a new created in the new region. All data will be lost and a new name assigned.***
     * 
     */
    @Export(name="region", refs={String.class}, tree="[0]")
    private Output<String> region;

    /**
     * @return The region to host the instance in. See instance regions
     * 
     * ***Note: Changing region will force the instance to be destroyed and a new created in the new region. All data will be lost and a new name assigned.***
     * 
     */
    public Output<String> region() {
        return this.region;
    }
    /**
     * The Rabbit MQ version. Can be left out, will then be set to default value used by CloudAMQP API.
     * 
     * ***Note: There is not yet any support in the provider to change the RMQ version. Once it&#39;s set in the initial creation, it will remain.***
     * 
     */
    @Export(name="rmqVersion", refs={String.class}, tree="[0]")
    private Output<String> rmqVersion;

    /**
     * @return The Rabbit MQ version. Can be left out, will then be set to default value used by CloudAMQP API.
     * 
     * ***Note: There is not yet any support in the provider to change the RMQ version. Once it&#39;s set in the initial creation, it will remain.***
     * 
     */
    public Output<String> rmqVersion() {
        return this.rmqVersion;
    }
    /**
     * One or more tags for the CloudAMQP instance, makes it possible to categories multiple instances in console view. Default there is no tags assigned.
     * 
     */
    @Export(name="tags", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> tags;

    /**
     * @return One or more tags for the CloudAMQP instance, makes it possible to categories multiple instances in console view. Default there is no tags assigned.
     * 
     */
    public Output<Optional<List<String>>> tags() {
        return Codegen.optional(this.tags);
    }
    /**
     * The AMQP URL (uses the internal hostname if the instance was created with VPC). Has the format: `amqps://{username}:{password}{@literal @}{hostname}/{vhost}`
     * 
     */
    @Export(name="url", refs={String.class}, tree="[0]")
    private Output<String> url;

    /**
     * @return The AMQP URL (uses the internal hostname if the instance was created with VPC). Has the format: `amqps://{username}:{password}{@literal @}{hostname}/{vhost}`
     * 
     */
    public Output<String> url() {
        return this.url;
    }
    /**
     * The virtual host used by Rabbit MQ.
     * 
     */
    @Export(name="vhost", refs={String.class}, tree="[0]")
    private Output<String> vhost;

    /**
     * @return The virtual host used by Rabbit MQ.
     * 
     */
    public Output<String> vhost() {
        return this.vhost;
    }
    /**
     * The VPC ID. Use this to create your instance in an existing VPC. See available example.
     * 
     */
    @Export(name="vpcId", refs={Integer.class}, tree="[0]")
    private Output<Integer> vpcId;

    /**
     * @return The VPC ID. Use this to create your instance in an existing VPC. See available example.
     * 
     */
    public Output<Integer> vpcId() {
        return this.vpcId;
    }
    /**
     * Creates a dedicated VPC subnet, shouldn&#39;t overlap with other VPC subnet, default subnet used 10.56.72.0/24.
     * 
     * ***Deprecated: Will be removed in next major version (v2.0)***
     * 
     * ***Note: extra fee will be charged when using VPC, see [CloudAMQP](https://cloudamqp.com) for more information.***
     * 
     */
    @Export(name="vpcSubnet", refs={String.class}, tree="[0]")
    private Output<String> vpcSubnet;

    /**
     * @return Creates a dedicated VPC subnet, shouldn&#39;t overlap with other VPC subnet, default subnet used 10.56.72.0/24.
     * 
     * ***Deprecated: Will be removed in next major version (v2.0)***
     * 
     * ***Note: extra fee will be charged when using VPC, see [CloudAMQP](https://cloudamqp.com) for more information.***
     * 
     */
    public Output<String> vpcSubnet() {
        return this.vpcSubnet;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Instance(java.lang.String name) {
        this(name, InstanceArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Instance(java.lang.String name, InstanceArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Instance(java.lang.String name, InstanceArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("cloudamqp:index/instance:Instance", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private Instance(java.lang.String name, Output<java.lang.String> id, @Nullable InstanceState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("cloudamqp:index/instance:Instance", name, state, makeResourceOptions(options, id), false);
    }

    private static InstanceArgs makeArgs(InstanceArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? InstanceArgs.Empty : args;
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<java.lang.String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .additionalSecretOutputs(List.of(
                "apikey",
                "url"
            ))
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
    public static Instance get(java.lang.String name, Output<java.lang.String> id, @Nullable InstanceState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Instance(name, id, state, options);
    }
}
