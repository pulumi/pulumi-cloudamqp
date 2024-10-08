// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.cloudamqp;

import com.pulumi.cloudamqp.ExtraDiskSizeArgs;
import com.pulumi.cloudamqp.Utilities;
import com.pulumi.cloudamqp.inputs.ExtraDiskSizeState;
import com.pulumi.cloudamqp.outputs.ExtraDiskSizeNode;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Boolean;
import java.lang.Integer;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * This resource allows you to resize the disk with additional storage capacity.
 * 
 * ***Before v1.25.0***: Only available for Amazon Web Services (AWS) without downtime.
 * 
 * ***From v1.25.0***: Google Compute Engine (GCE) and Azure available.
 * 
 * Introducing a new optional argument called `allow_downtime`.  Leaving it out or set it to false will proceed to try and resize the disk without downtime, available for *AWS* and *GCE*.
 * While *Azure* only support swapping the disk, and this argument needs to be set to *true*.
 * 
 * `allow_downtime` also makes it possible to circumvent the time rate limit or shrinking the disk.
 * 
 * | Cloud Platform        | allow_downtime=false | allow_downtime=true           |
 * |-----------------------|----------------------|-------------------------------|
 * | amazon-web-services   | Expand current disk* | Try to expand, otherwise swap |
 * | google-compute-engine | Expand current disk* | Try to expand, otherwise swap |
 * | azure-arm             | Not supported        | Swap disk to new size         |
 * 
 * *Preferable method to use.
 * 
 * &gt; **WARNING:** Due to restrictions from cloud providers, it&#39;s only possible to resize the disk every 8 hours. Unless the `allow_downtime=true` is set, then the disk will be swapped for a new.
 * 
 * Pricing is available at [cloudamqp.com](https://www.cloudamqp.com/) and only available for dedicated subscription plans.
 * 
 * ## Example Usage
 * 
 * &lt;details&gt;
 *   &lt;summary&gt;
 *     &lt;b&gt;
 *       &lt;i&gt;AWS extra disk size (before v1.25.0)&lt;/i&gt;
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
 * import com.pulumi.cloudamqp.ExtraDiskSize;
 * import com.pulumi.cloudamqp.ExtraDiskSizeArgs;
 * import com.pulumi.cloudamqp.CloudamqpFunctions;
 * import com.pulumi.cloudamqp.inputs.GetNodesArgs;
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
 *         // Instance
 *         var instance = new Instance("instance", InstanceArgs.builder()
 *             .name("Instance")
 *             .plan("bunny-1")
 *             .region("amazon-web-services::us-west-2")
 *             .build());
 * 
 *         // Resize disk with 25 extra GB
 *         var resizeDisk = new ExtraDiskSize("resizeDisk", ExtraDiskSizeArgs.builder()
 *             .instanceId(instance.id())
 *             .extraDiskSize(25)
 *             .build());
 * 
 *         // Optional, refresh nodes info after disk resize by adding dependency
 *         // to cloudamqp_extra_disk_size.resize_disk resource
 *         final var nodes = CloudamqpFunctions.getNodes(GetNodesArgs.builder()
 *             .instanceId(instance.id())
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
 *       &lt;i&gt;AWS extra disk size without downtime&lt;/i&gt;
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
 * import com.pulumi.cloudamqp.ExtraDiskSize;
 * import com.pulumi.cloudamqp.ExtraDiskSizeArgs;
 * import com.pulumi.cloudamqp.CloudamqpFunctions;
 * import com.pulumi.cloudamqp.inputs.GetNodesArgs;
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
 *         // Instance
 *         var instance = new Instance("instance", InstanceArgs.builder()
 *             .name("Instance")
 *             .plan("bunny-1")
 *             .region("amazon-web-services::us-west-2")
 *             .build());
 * 
 *         // Resize disk with 25 extra GB, without downtime
 *         var resizeDisk = new ExtraDiskSize("resizeDisk", ExtraDiskSizeArgs.builder()
 *             .instanceId(instance.id())
 *             .extraDiskSize(25)
 *             .build());
 * 
 *         // Optional, refresh nodes info after disk resize by adding dependency
 *         // to cloudamqp_extra_disk_size.resize_disk resource
 *         final var nodes = CloudamqpFunctions.getNodes(GetNodesArgs.builder()
 *             .instanceId(instance.id())
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
 *       &lt;i&gt;GCE extra disk size without downtime&lt;/i&gt;
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
 * import com.pulumi.cloudamqp.ExtraDiskSize;
 * import com.pulumi.cloudamqp.ExtraDiskSizeArgs;
 * import com.pulumi.cloudamqp.CloudamqpFunctions;
 * import com.pulumi.cloudamqp.inputs.GetNodesArgs;
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
 *         // Instance
 *         var instance = new Instance("instance", InstanceArgs.builder()
 *             .name("Instance")
 *             .plan("bunny-1")
 *             .region("google-compute-engine::us-central1")
 *             .build());
 * 
 *         // Resize disk with 25 extra GB, without downtime
 *         var resizeDisk = new ExtraDiskSize("resizeDisk", ExtraDiskSizeArgs.builder()
 *             .instanceId(instance.id())
 *             .extraDiskSize(25)
 *             .build());
 * 
 *         // Optional, refresh nodes info after disk resize by adding dependency
 *         // to cloudamqp_extra_disk_size.resize_disk resource
 *         final var nodes = CloudamqpFunctions.getNodes(GetNodesArgs.builder()
 *             .instanceId(instance.id())
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
 *       &lt;i&gt;Azure extra disk size with downtime&lt;/i&gt;
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
 * import com.pulumi.cloudamqp.ExtraDiskSize;
 * import com.pulumi.cloudamqp.ExtraDiskSizeArgs;
 * import com.pulumi.cloudamqp.CloudamqpFunctions;
 * import com.pulumi.cloudamqp.inputs.GetNodesArgs;
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
 *         // Instance
 *         var instance = new Instance("instance", InstanceArgs.builder()
 *             .name("Instance")
 *             .plan("bunny-1")
 *             .region("azure-arm::centralus")
 *             .build());
 * 
 *         // Resize disk with 25 extra GB, with downtime
 *         var resizeDisk = new ExtraDiskSize("resizeDisk", ExtraDiskSizeArgs.builder()
 *             .instanceId(instance.id())
 *             .extraDiskSize(25)
 *             .allowDowntime(true)
 *             .build());
 * 
 *         // Optional, refresh nodes info after disk resize by adding dependency
 *         // to cloudamqp_extra_disk_size.resize_disk resource
 *         final var nodes = CloudamqpFunctions.getNodes(GetNodesArgs.builder()
 *             .instanceId(instance.id())
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
 * ## Attributes reference
 * 
 * All attributes reference are computed
 * 
 * * `id`    - The identifier for this resource.
 * * `nodes` - An array of node information. Each `nodes` block consists of the fields documented below.
 * 
 * ***
 * 
 * The `nodes` block consist of
 * 
 * * `name`                  - Name of the node.
 * * `disk_size`             - Subscription plan disk size
 * * `additional_disk_size`  - Additional added disk size
 * 
 * ***Note:*** *Total disk size = disk_size + additional_disk_size*
 * 
 * ## Dependency
 * 
 * This data source depends on CloudAMQP instance identifier, `cloudamqp_instance.instance.id`.
 * 
 * ## Import
 * 
 * Not possible to import this resource.
 * 
 */
@ResourceType(type="cloudamqp:index/extraDiskSize:ExtraDiskSize")
public class ExtraDiskSize extends com.pulumi.resources.CustomResource {
    /**
     * When resizing the disk, allow cluster downtime if necessary. Default set to false. Required when hosting in *Azure*.
     * 
     */
    @Export(name="allowDowntime", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> allowDowntime;

    /**
     * @return When resizing the disk, allow cluster downtime if necessary. Default set to false. Required when hosting in *Azure*.
     * 
     */
    public Output<Optional<Boolean>> allowDowntime() {
        return Codegen.optional(this.allowDowntime);
    }
    /**
     * Extra disk size in GB. Supported values: 0, 25, 50, 100, 250, 500, 1000, 2000
     * 
     */
    @Export(name="extraDiskSize", refs={Integer.class}, tree="[0]")
    private Output<Integer> extraDiskSize;

    /**
     * @return Extra disk size in GB. Supported values: 0, 25, 50, 100, 250, 500, 1000, 2000
     * 
     */
    public Output<Integer> extraDiskSize() {
        return this.extraDiskSize;
    }
    /**
     * The CloudAMQP instance ID.
     * 
     */
    @Export(name="instanceId", refs={Integer.class}, tree="[0]")
    private Output<Integer> instanceId;

    /**
     * @return The CloudAMQP instance ID.
     * 
     */
    public Output<Integer> instanceId() {
        return this.instanceId;
    }
    @Export(name="nodes", refs={List.class,ExtraDiskSizeNode.class}, tree="[0,1]")
    private Output<List<ExtraDiskSizeNode>> nodes;

    public Output<List<ExtraDiskSizeNode>> nodes() {
        return this.nodes;
    }
    /**
     * Configurable sleep time in seconds between retries for resizing the disk. Default set to 30 seconds.
     * 
     */
    @Export(name="sleep", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> sleep;

    /**
     * @return Configurable sleep time in seconds between retries for resizing the disk. Default set to 30 seconds.
     * 
     */
    public Output<Optional<Integer>> sleep() {
        return Codegen.optional(this.sleep);
    }
    /**
     * Configurable timeout time in seconds for resizing the disk. Default set to 1800 seconds.
     * 
     * ***Note:*** `allow_downtime`, `sleep`, `timeout` only available from v1.25.0.
     * 
     */
    @Export(name="timeout", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> timeout;

    /**
     * @return Configurable timeout time in seconds for resizing the disk. Default set to 1800 seconds.
     * 
     * ***Note:*** `allow_downtime`, `sleep`, `timeout` only available from v1.25.0.
     * 
     */
    public Output<Optional<Integer>> timeout() {
        return Codegen.optional(this.timeout);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public ExtraDiskSize(java.lang.String name) {
        this(name, ExtraDiskSizeArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public ExtraDiskSize(java.lang.String name, ExtraDiskSizeArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public ExtraDiskSize(java.lang.String name, ExtraDiskSizeArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("cloudamqp:index/extraDiskSize:ExtraDiskSize", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private ExtraDiskSize(java.lang.String name, Output<java.lang.String> id, @Nullable ExtraDiskSizeState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("cloudamqp:index/extraDiskSize:ExtraDiskSize", name, state, makeResourceOptions(options, id), false);
    }

    private static ExtraDiskSizeArgs makeArgs(ExtraDiskSizeArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? ExtraDiskSizeArgs.Empty : args;
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<java.lang.String> id) {
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
    public static ExtraDiskSize get(java.lang.String name, Output<java.lang.String> id, @Nullable ExtraDiskSizeState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new ExtraDiskSize(name, id, state, options);
    }
}
