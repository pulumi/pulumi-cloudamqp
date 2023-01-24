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
 * This resource allows you to automatically upgrade to latest possible upgradable versions for RabbitMQ and Erlang. Depending on initial versions of RabbitMQ and Erlang of the CloudAMQP instance, multiple runs may be needed to get to latest versions. After completed upgrade, check data source `cloudamqp.getUpgradableVersions` to see if newer versions is available. Then delete `cloudamqp.UpgradeRabbitmq` and create it again to invoke the upgrade.
 * 
 * &gt; **Important Upgrade Information**
 * &gt; - All nodes in a cluster must run the same major and minor version of RabbitMQ. The entire cluster will be offline while upgrading major or minor versions.
 * &gt; - Auto delete queues (queues that are marked AD) will be deleted during the update.
 * &gt; - Any custom plugins support has installed on your behalf will be disabled and you need to contact support@cloudamqp.com and ask to have them re-installed.
 * &gt; - TLS 1.0 and 1.1 will not be supported after the update.
 * 
 * Only available for dedicated subscription plans.
 * 
 * ## Example Usage
 * ```java
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
 *         final var versions = CloudamqpFunctions.getUpgradableVersions(GetUpgradableVersionsArgs.builder()
 *             .instanceId(cloudamqp_instance.instance().id())
 *             .build());
 * 
 *         var upgrade = new UpgradeRabbitmq(&#34;upgrade&#34;, UpgradeRabbitmqArgs.builder()        
 *             .instanceId(cloudamqp_instance.instance().id())
 *             .build());
 * 
 *     }
 * }
 * ```
 * ```java
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
 *         final var versions = CloudamqpFunctions.getUpgradableVersions(GetUpgradableVersionsArgs.builder()
 *             .instanceId(cloudamqp_instance.instance().id())
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * If newer version is still available to be upgradable in the data source, re-run again.
 * ```java
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
 *         final var versions = CloudamqpFunctions.getUpgradableVersions(GetUpgradableVersionsArgs.builder()
 *             .instanceId(cloudamqp_instance.instance().id())
 *             .build());
 * 
 *         var upgrade = new UpgradeRabbitmq(&#34;upgrade&#34;, UpgradeRabbitmqArgs.builder()        
 *             .instanceId(cloudamqp_instance.instance().id())
 *             .build());
 * 
 *     }
 * }
 * ```
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
    @Export(name="instanceId", type=Integer.class, parameters={})
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
        super("cloudamqp:index/upgradeRabbitmq:UpgradeRabbitmq", name, args == null ? UpgradeRabbitmqArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private UpgradeRabbitmq(String name, Output<String> id, @Nullable UpgradeRabbitmqState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("cloudamqp:index/upgradeRabbitmq:UpgradeRabbitmq", name, state, makeResourceOptions(options, id));
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
