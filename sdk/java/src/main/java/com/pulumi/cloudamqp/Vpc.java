// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.cloudamqp;

import com.pulumi.cloudamqp.Utilities;
import com.pulumi.cloudamqp.VpcArgs;
import com.pulumi.cloudamqp.inputs.VpcState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * This resource allows you to manage standalone VPC.
 * 
 * New Cloudamqp instances can be added to the managed VPC. Set the instance *vpc_id* attribute to the
 * managed vpc identifier, see example below, when creating the instance.
 * 
 * Only available for dedicated subscription plans.
 * 
 * Pricing is available at [CloudAMQP plans].
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
 * import com.pulumi.cloudamqp.Vpc;
 * import com.pulumi.cloudamqp.VpcArgs;
 * import com.pulumi.cloudamqp.Instance;
 * import com.pulumi.cloudamqp.InstanceArgs;
 * import com.pulumi.cloudamqp.CloudamqpFunctions;
 * import com.pulumi.cloudamqp.inputs.GetVpcInfoArgs;
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
 *         // Managed VPC resource
 *         var vpc = new Vpc("vpc", VpcArgs.builder()
 *             .name("<VPC name>")
 *             .region("amazon-web-services::us-east-1")
 *             .subnet("10.56.72.0/24")
 *             .tags()
 *             .build());
 * 
 *         //  New instance, need to be created with a vpc
 *         var instance = new Instance("instance", InstanceArgs.builder()
 *             .name("<Instance name>")
 *             .plan("penguin-1")
 *             .region("amazon-web-services::us-east-1")
 *             .nodes(1)
 *             .tags()
 *             .vpcId(vpcCloudamqVpc.id())
 *             .keepAssociatedVpc(true)
 *             .build());
 * 
 *         // Additional VPC information
 *         final var vpcInfo = CloudamqpFunctions.getVpcInfo(GetVpcInfoArgs.builder()
 *             .vpcId(vpc.id())
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
 * `cloudamqp_vpc` can be imported using the CloudAMQP VPC identifier. To retrieve the identifier of an
 * 
 * instance, use [CloudAMQP API list vpcs].
 * 
 * From Terraform v1.5.0, the `import` block can be used to import this resource:
 * 
 * hcl
 * 
 * import {
 * 
 *   to = cloudamqp_vpc.vpc
 * 
 *   id = &lt;vpc_id&gt;
 * 
 * }
 * 
 * Or with Terraform CLI:
 * 
 * ```sh
 * $ pulumi import cloudamqp:index/vpc:Vpc vpc &lt;vpc_id&gt;`
 * ```
 * 
 * [CloudAMQP API list vpcs]: https://docs.cloudamqp.com/#list-vpcs
 * 
 * [CloudAMQP plans]: https://www.cloudamqp.com/plans.html
 * 
 */
@ResourceType(type="cloudamqp:index/vpc:Vpc")
public class Vpc extends com.pulumi.resources.CustomResource {
    /**
     * The name of the VPC.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return The name of the VPC.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * The hosted region for the managed standalone VPC
     * 
     */
    @Export(name="region", refs={String.class}, tree="[0]")
    private Output<String> region;

    /**
     * @return The hosted region for the managed standalone VPC
     * 
     */
    public Output<String> region() {
        return this.region;
    }
    /**
     * The VPC subnet
     * 
     */
    @Export(name="subnet", refs={String.class}, tree="[0]")
    private Output<String> subnet;

    /**
     * @return The VPC subnet
     * 
     */
    public Output<String> subnet() {
        return this.subnet;
    }
    /**
     * Tag the VPC with optional tags
     * 
     */
    @Export(name="tags", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> tags;

    /**
     * @return Tag the VPC with optional tags
     * 
     */
    public Output<Optional<List<String>>> tags() {
        return Codegen.optional(this.tags);
    }
    /**
     * VPC name given when hosted at the cloud provider
     * 
     */
    @Export(name="vpcName", refs={String.class}, tree="[0]")
    private Output<String> vpcName;

    /**
     * @return VPC name given when hosted at the cloud provider
     * 
     */
    public Output<String> vpcName() {
        return this.vpcName;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Vpc(java.lang.String name) {
        this(name, VpcArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Vpc(java.lang.String name, VpcArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Vpc(java.lang.String name, VpcArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("cloudamqp:index/vpc:Vpc", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private Vpc(java.lang.String name, Output<java.lang.String> id, @Nullable VpcState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("cloudamqp:index/vpc:Vpc", name, state, makeResourceOptions(options, id), false);
    }

    private static VpcArgs makeArgs(VpcArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? VpcArgs.Empty : args;
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
    public static Vpc get(java.lang.String name, Output<java.lang.String> id, @Nullable VpcState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Vpc(name, id, state, options);
    }
}
