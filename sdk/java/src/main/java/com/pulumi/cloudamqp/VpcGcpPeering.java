// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.cloudamqp;

import com.pulumi.cloudamqp.Utilities;
import com.pulumi.cloudamqp.VpcGcpPeeringArgs;
import com.pulumi.cloudamqp.inputs.VpcGcpPeeringState;
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
 * This resouce creates a VPC peering configuration for the CloudAMQP instance. The configuration will connect to another VPC network hosted on Google Cloud Platform (GCP). See the [GCP documentation](https://cloud.google.com/vpc/docs/using-vpc-peering) for more information on how to create the VPC peering configuration.
 * 
 * &gt; **Note:** Creating a VPC peering will automatically add firewall rules for the peered subnet.
 * 
 * &lt;details&gt;
 *  &lt;summary&gt;
 *     &lt;i&gt;Default VPC peering firewall rule&lt;/i&gt;
 *   &lt;/summary&gt;
 * 
 * ## Example Usage
 * 
 * &lt;details&gt;
 *   &lt;summary&gt;
 *     &lt;b&gt;
 *       &lt;i&gt;VPC peering before v1.16.0&lt;/i&gt;
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
 * import com.pulumi.cloudamqp.CloudamqpFunctions;
 * import com.pulumi.cloudamqp.inputs.GetVpcGcpInfoArgs;
 * import com.pulumi.cloudamqp.VpcGcpPeering;
 * import com.pulumi.cloudamqp.VpcGcpPeeringArgs;
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
 *         // CloudAMQP instance
 *         var instance = new Instance("instance", InstanceArgs.builder()
 *             .name("terraform-vpc-peering")
 *             .plan("bunny-1")
 *             .region("google-compute-engine::europe-north1")
 *             .tags("terraform")
 *             .vpcSubnet("10.40.72.0/24")
 *             .build());
 * 
 *         // VPC information
 *         final var vpcInfo = CloudamqpFunctions.getVpcGcpInfo(GetVpcGcpInfoArgs.builder()
 *             .instanceId(instance.id())
 *             .build());
 * 
 *         // VPC peering configuration
 *         var vpcPeeringRequest = new VpcGcpPeering("vpcPeeringRequest", VpcGcpPeeringArgs.builder()
 *             .instanceId(instance.id())
 *             .peerNetworkUri("https://www.googleapis.com/compute/v1/projects/<PROJECT-NAME>/global/networks/<NETWORK-NAME>")
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
 *       &lt;i&gt;VPC peering from v1.16.0 (Managed VPC)&lt;/i&gt;
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
 * import com.pulumi.cloudamqp.CloudamqpFunctions;
 * import com.pulumi.cloudamqp.inputs.GetVpcGcpInfoArgs;
 * import com.pulumi.cloudamqp.VpcGcpPeering;
 * import com.pulumi.cloudamqp.VpcGcpPeeringArgs;
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
 *             .region("google-compute-engine::europe-north1")
 *             .subnet("10.56.72.0/24")
 *             .tags()
 *             .build());
 * 
 *         // CloudAMQP instance
 *         var instance = new Instance("instance", InstanceArgs.builder()
 *             .name("terraform-vpc-peering")
 *             .plan("bunny-1")
 *             .region("google-compute-engine::europe-north1")
 *             .tags("terraform")
 *             .vpcId(vpc.id())
 *             .build());
 * 
 *         // VPC information
 *         final var vpcInfo = CloudamqpFunctions.getVpcGcpInfo(GetVpcGcpInfoArgs.builder()
 *             .vpcId(vpc.info())
 *             .build());
 * 
 *         // VPC peering configuration
 *         var vpcPeeringRequest = new VpcGcpPeering("vpcPeeringRequest", VpcGcpPeeringArgs.builder()
 *             .vpcId(vpc.id())
 *             .peerNetworkUri("https://www.googleapis.com/compute/v1/projects/<PROJECT-NAME>/global/networks/<NETWORK-NAME>")
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
 *       &lt;i&gt;VPC peering from v1.28.0, wait_on_peering_status &lt;/i&gt;
 *     &lt;/b&gt;
 *   &lt;/summary&gt;
 * 
 * Default peering request, no need to set `wait_on_peering_status`. It&#39;s default set to false and will not wait on peering status.
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.cloudamqp.VpcGcpPeering;
 * import com.pulumi.cloudamqp.VpcGcpPeeringArgs;
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
 *         var vpcPeeringRequest = new VpcGcpPeering("vpcPeeringRequest", VpcGcpPeeringArgs.builder()
 *             .vpcId(vpc.id())
 *             .peerNetworkUri("https://www.googleapis.com/compute/v1/projects/<PROJECT-NAME>/global/networks/<NETWORK-NAME>")
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * Peering request and waiting for peering status.
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.cloudamqp.VpcGcpPeering;
 * import com.pulumi.cloudamqp.VpcGcpPeeringArgs;
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
 *         var vpcPeeringRequest = new VpcGcpPeering("vpcPeeringRequest", VpcGcpPeeringArgs.builder()
 *             .vpcId(vpc.id())
 *             .waitOnPeeringStatus(true)
 *             .peerNetworkUri("https://www.googleapis.com/compute/v1/projects/<PROJECT-NAME>/global/networks/<NETWORK-NAME>")
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
 * ### With Additional Firewall Rules
 * 
 * &lt;details&gt;
 *   &lt;summary&gt;
 *     &lt;b&gt;
 *       &lt;i&gt;VPC peering before v1.16.0&lt;/i&gt;
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
 * import com.pulumi.cloudamqp.VpcGcpPeering;
 * import com.pulumi.cloudamqp.VpcGcpPeeringArgs;
 * import com.pulumi.cloudamqp.SecurityFirewall;
 * import com.pulumi.cloudamqp.SecurityFirewallArgs;
 * import com.pulumi.cloudamqp.inputs.SecurityFirewallRuleArgs;
 * import com.pulumi.resources.CustomResourceOptions;
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
 *         // VPC peering configuration
 *         var vpcPeeringRequest = new VpcGcpPeering("vpcPeeringRequest", VpcGcpPeeringArgs.builder()
 *             .instanceId(instance.id())
 *             .peerNetworkUri(peerNetworkUri)
 *             .build());
 * 
 *         // Firewall rules
 *         var firewallSettings = new SecurityFirewall("firewallSettings", SecurityFirewallArgs.builder()
 *             .instanceId(instance.id())
 *             .rules(            
 *                 SecurityFirewallRuleArgs.builder()
 *                     .ip(peerSubnet)
 *                     .ports(15672)
 *                     .services(                    
 *                         "AMQP",
 *                         "AMQPS",
 *                         "STREAM",
 *                         "STREAM_SSL")
 *                     .description("VPC peering for <NETWORK>")
 *                     .build(),
 *                 SecurityFirewallRuleArgs.builder()
 *                     .ip("192.168.0.0/24")
 *                     .ports(                    
 *                         4567,
 *                         4568)
 *                     .services(                    
 *                         "AMQP",
 *                         "AMQPS",
 *                         "HTTPS")
 *                     .build())
 *             .build(), CustomResourceOptions.builder()
 *                 .dependsOn(vpcPeeringRequest)
 *                 .build());
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
 *       &lt;i&gt;VPC peering from v1.16.0 (Managed VPC)&lt;/i&gt;
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
 * import com.pulumi.cloudamqp.VpcGcpPeering;
 * import com.pulumi.cloudamqp.VpcGcpPeeringArgs;
 * import com.pulumi.cloudamqp.SecurityFirewall;
 * import com.pulumi.cloudamqp.SecurityFirewallArgs;
 * import com.pulumi.cloudamqp.inputs.SecurityFirewallRuleArgs;
 * import com.pulumi.resources.CustomResourceOptions;
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
 *         // VPC peering configuration
 *         var vpcPeeringRequest = new VpcGcpPeering("vpcPeeringRequest", VpcGcpPeeringArgs.builder()
 *             .vpcId(vpc.id())
 *             .peerNetworkUri(peerNetworkUri)
 *             .build());
 * 
 *         // Firewall rules
 *         var firewallSettings = new SecurityFirewall("firewallSettings", SecurityFirewallArgs.builder()
 *             .instanceId(instance.id())
 *             .rules(            
 *                 SecurityFirewallRuleArgs.builder()
 *                     .ip(peerSubnet)
 *                     .ports(15672)
 *                     .services(                    
 *                         "AMQP",
 *                         "AMQPS",
 *                         "STREAM",
 *                         "STREAM_SSL")
 *                     .description("VPC peering for <NETWORK>")
 *                     .build(),
 *                 SecurityFirewallRuleArgs.builder()
 *                     .ip("0.0.0.0/0")
 *                     .ports()
 *                     .services("HTTPS")
 *                     .description("MGMT interface")
 *                     .build())
 *             .build(), CustomResourceOptions.builder()
 *                 .dependsOn(vpcPeeringRequest)
 *                 .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * &lt;/details&gt;
 * 
 * ## Depedency
 * 
 * *Before v1.16.0*
 * This resource depends on CloudAMQP instance identifier, `cloudamqp_instance.instance.id`.
 * 
 * *From v1.16.0*
 * This resource depends on CloudAMQP managed VPC identifier, `cloudamqp_vpc.vpc.id` or instance identifier, `cloudamqp_instance.instance.id`.
 * 
 * ## Create VPC Peering with additional firewall rules
 * 
 * To create a VPC peering configuration with additional firewall rules, it&#39;s required to chain the cloudamqp.SecurityFirewall
 * resource to avoid parallel conflicting resource calls. This is done by adding dependency from the firewall resource to the VPC peering resource.
 * 
 * Furthermore, since all firewall rules are overwritten, the otherwise automatically added rules for the VPC peering also needs to be added.
 * 
 * See example below.
 * 
 * ## Import
 * 
 * Not possible to import this resource.
 * 
 */
@ResourceType(type="cloudamqp:index/vpcGcpPeering:VpcGcpPeering")
public class VpcGcpPeering extends com.pulumi.resources.CustomResource {
    /**
     * VPC peering auto created routes
     * 
     */
    @Export(name="autoCreateRoutes", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> autoCreateRoutes;

    /**
     * @return VPC peering auto created routes
     * 
     */
    public Output<Boolean> autoCreateRoutes() {
        return this.autoCreateRoutes;
    }
    /**
     * The CloudAMQP instance identifier. *Deprecated from v1.16.0*
     * 
     */
    @Export(name="instanceId", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> instanceId;

    /**
     * @return The CloudAMQP instance identifier. *Deprecated from v1.16.0*
     * 
     */
    public Output<Optional<Integer>> instanceId() {
        return Codegen.optional(this.instanceId);
    }
    /**
     * Network uri of the VPC network to which you will peer with.
     * 
     */
    @Export(name="peerNetworkUri", refs={String.class}, tree="[0]")
    private Output<String> peerNetworkUri;

    /**
     * @return Network uri of the VPC network to which you will peer with.
     * 
     */
    public Output<String> peerNetworkUri() {
        return this.peerNetworkUri;
    }
    /**
     * Configurable sleep time (seconds) between retries when requesting or reading
     * peering. Default set to 10 seconds. *Available from v1.29.0*
     * 
     */
    @Export(name="sleep", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> sleep;

    /**
     * @return Configurable sleep time (seconds) between retries when requesting or reading
     * peering. Default set to 10 seconds. *Available from v1.29.0*
     * 
     */
    public Output<Optional<Integer>> sleep() {
        return Codegen.optional(this.sleep);
    }
    /**
     * VPC peering state
     * 
     */
    @Export(name="state", refs={String.class}, tree="[0]")
    private Output<String> state;

    /**
     * @return VPC peering state
     * 
     */
    public Output<String> state() {
        return this.state;
    }
    /**
     * VPC peering state details
     * 
     */
    @Export(name="stateDetails", refs={String.class}, tree="[0]")
    private Output<String> stateDetails;

    /**
     * @return VPC peering state details
     * 
     */
    public Output<String> stateDetails() {
        return this.stateDetails;
    }
    /**
     * Configurable timeout time (seconds) before retries times out. Default set
     * to 1800 seconds. *Available from v1.29.0*
     * 
     */
    @Export(name="timeout", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> timeout;

    /**
     * @return Configurable timeout time (seconds) before retries times out. Default set
     * to 1800 seconds. *Available from v1.29.0*
     * 
     */
    public Output<Optional<Integer>> timeout() {
        return Codegen.optional(this.timeout);
    }
    /**
     * The managed VPC identifier. *Available from v1.16.0*
     * 
     */
    @Export(name="vpcId", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> vpcId;

    /**
     * @return The managed VPC identifier. *Available from v1.16.0*
     * 
     */
    public Output<Optional<String>> vpcId() {
        return Codegen.optional(this.vpcId);
    }
    /**
     * Makes the resource wait until the peering is connected.
     * Default set to false. *Available from v1.28.0*
     * 
     */
    @Export(name="waitOnPeeringStatus", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> waitOnPeeringStatus;

    /**
     * @return Makes the resource wait until the peering is connected.
     * Default set to false. *Available from v1.28.0*
     * 
     */
    public Output<Optional<Boolean>> waitOnPeeringStatus() {
        return Codegen.optional(this.waitOnPeeringStatus);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public VpcGcpPeering(String name) {
        this(name, VpcGcpPeeringArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public VpcGcpPeering(String name, VpcGcpPeeringArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public VpcGcpPeering(String name, VpcGcpPeeringArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("cloudamqp:index/vpcGcpPeering:VpcGcpPeering", name, args == null ? VpcGcpPeeringArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private VpcGcpPeering(String name, Output<String> id, @Nullable VpcGcpPeeringState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("cloudamqp:index/vpcGcpPeering:VpcGcpPeering", name, state, makeResourceOptions(options, id));
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
    public static VpcGcpPeering get(String name, Output<String> id, @Nullable VpcGcpPeeringState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new VpcGcpPeering(name, id, state, options);
    }
}
