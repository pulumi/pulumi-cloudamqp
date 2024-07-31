// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.cloudamqp;

import com.pulumi.cloudamqp.UpgradeRabbitmqArgs;
import com.pulumi.cloudamqp.Utilities;
import com.pulumi.cloudamqp.inputs.UpgradeRabbitmqState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Integer;
import javax.annotation.Nullable;

/**
 * This resource allows you to automatically upgrade to the latest possible upgradable versions for RabbitMQ and Erlang. Depending on initial versions of RabbitMQ and Erlang of the CloudAMQP instance, multiple runs may be needed to get to the latest versions. After completed upgrade, check data source `cloudamqp.getUpgradableVersions` to see if newer versions is available. Then delete `cloudamqp.UpgradeRabbitmq` and create it again to invoke the upgrade.
 * 
 * &gt; **Important Upgrade Information**
 * &gt; - All single node upgrades will require some downtime since RabbitMQ needs a restart.
 * &gt; - From RabbitMQ version 3.9, rolling upgrades between minor versions (e.g. 3.9 to 3.10), in a multi-node cluster are possible without downtime. This means that one node is upgraded at a time while the other nodes are still running. For versions older than 3.9, patch version upgrades (e.g. 3.8.x to 3.8.y) are possible without downtime in a multi-node cluster, but minor version upgrades will require downtime.
 * &gt; - Auto delete queues (queues that are marked AD) will be deleted during the update.
 * &gt; - Any custom plugins support has installed on your behalf will be disabled and you need to contact support{@literal @}cloudamqp.com and ask to have them re-installed.
 * &gt; - TLS 1.0 and 1.1 will not be supported after the update.
 * 
 * Only available for dedicated subscription plans running ***RabbitMQ***.
 * 
 * ## Example Usage
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.cloudamqp.CloudamqpFunctions;
 * import com.pulumi.cloudamqp.inputs.GetUpgradableVersionsArgs;
 * import com.pulumi.cloudamqp.UpgradeRabbitmq;
 * import com.pulumi.cloudamqp.UpgradeRabbitmqArgs;
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
 *         // Retrieve latest possible upgradable versions for RabbitMQ and Erlang
 *         final var versions = CloudamqpFunctions.getUpgradableVersions(GetUpgradableVersionsArgs.builder()
 *             .instanceId(instance.id())
 *             .build());
 * 
 *         // Invoke automatically upgrade to latest possible upgradable versions for RabbitMQ and Erlang
 *         var upgrade = new UpgradeRabbitmq("upgrade", UpgradeRabbitmqArgs.builder()
 *             .instanceId(instance.id())
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.cloudamqp.CloudamqpFunctions;
 * import com.pulumi.cloudamqp.inputs.GetUpgradableVersionsArgs;
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
 *         // Retrieve latest possible upgradable versions for RabbitMQ and Erlang
 *         final var versions = CloudamqpFunctions.getUpgradableVersions(GetUpgradableVersionsArgs.builder()
 *             .instanceId(instance.id())
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * If newer version is still available to be upgradable in the data source, re-run again.
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.cloudamqp.CloudamqpFunctions;
 * import com.pulumi.cloudamqp.inputs.GetUpgradableVersionsArgs;
 * import com.pulumi.cloudamqp.UpgradeRabbitmq;
 * import com.pulumi.cloudamqp.UpgradeRabbitmqArgs;
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
 *         // Retrieve latest possible upgradable versions for RabbitMQ and Erlang
 *         final var versions = CloudamqpFunctions.getUpgradableVersions(GetUpgradableVersionsArgs.builder()
 *             .instanceId(instance.id())
 *             .build());
 * 
 *         // Invoke automatically upgrade to latest possible upgradable versions for RabbitMQ and Erlang
 *         var upgrade = new UpgradeRabbitmq("upgrade", UpgradeRabbitmqArgs.builder()
 *             .instanceId(instance.id())
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Import
 * 
 * Not possible to import this resource.
 * 
 */
@ResourceType(type="cloudamqp:index/upgradeRabbitmq:UpgradeRabbitmq")
public class UpgradeRabbitmq extends com.pulumi.resources.CustomResource {
    /**
     * The CloudAMQP instance identifier
     * 
     */
    @Export(name="instanceId", refs={Integer.class}, tree="[0]")
    private Output<Integer> instanceId;

    /**
     * @return The CloudAMQP instance identifier
     * 
     */
    public Output<Integer> instanceId() {
        return this.instanceId;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public UpgradeRabbitmq(String name) {
        this(name, UpgradeRabbitmqArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public UpgradeRabbitmq(String name, UpgradeRabbitmqArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public UpgradeRabbitmq(String name, UpgradeRabbitmqArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("cloudamqp:index/upgradeRabbitmq:UpgradeRabbitmq", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()));
    }

    private UpgradeRabbitmq(String name, Output<String> id, @Nullable UpgradeRabbitmqState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("cloudamqp:index/upgradeRabbitmq:UpgradeRabbitmq", name, state, makeResourceOptions(options, id));
    }

    private static UpgradeRabbitmqArgs makeArgs(UpgradeRabbitmqArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? UpgradeRabbitmqArgs.Empty : args;
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
    public static UpgradeRabbitmq get(String name, Output<String> id, @Nullable UpgradeRabbitmqState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new UpgradeRabbitmq(name, id, state, options);
    }
}
