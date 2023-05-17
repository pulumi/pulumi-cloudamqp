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
 * Only available for dedicated subscription plans.
 * 
 * Pricing is available at [cloudamqp.com](https://www.cloudamqp.com/plans.html).
 * 
 * ## Example Usage
 * ### With Additional Firewall Rules
 * 
 * &lt;details&gt;
 *   &lt;summary&gt;
 *     &lt;b&gt;
 *       &lt;i&gt;VPC peering pre v1.16.0&lt;/i&gt;
 *     &lt;/b&gt;
 *   &lt;/summary&gt;
 * ```java
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
 *         var vpcPeeringRequest = new VpcGcpPeering(&#34;vpcPeeringRequest&#34;, VpcGcpPeeringArgs.builder()        
 *             .instanceId(cloudamqp_instance.instance().id())
 *             .peerNetworkUri(var_.peer_network_uri())
 *             .build());
 * 
 *         var firewallSettings = new SecurityFirewall(&#34;firewallSettings&#34;, SecurityFirewallArgs.builder()        
 *             .instanceId(cloudamqp_instance.instance().id())
 *             .rules(            
 *                 SecurityFirewallRuleArgs.builder()
 *                     .ip(var_.peer_subnet())
 *                     .ports(15672)
 *                     .services(                    
 *                         &#34;AMQP&#34;,
 *                         &#34;AMQPS&#34;,
 *                         &#34;STREAM&#34;,
 *                         &#34;STREAM_SSL&#34;)
 *                     .description(&#34;VPC peering for &lt;NETWORK&gt;&#34;)
 *                     .build(),
 *                 SecurityFirewallRuleArgs.builder()
 *                     .ip(&#34;192.168.0.0/24&#34;)
 *                     .ports(                    
 *                         4567,
 *                         4568)
 *                     .services(                    
 *                         &#34;AMQP&#34;,
 *                         &#34;AMQPS&#34;,
 *                         &#34;HTTPS&#34;)
 *                     .build())
 *             .build(), CustomResourceOptions.builder()
 *                 .dependsOn(vpcPeeringRequest)
 *                 .build());
 * 
 *     }
 * }
 * ```
 * &lt;/details&gt;
 * 
 * &lt;details&gt;
 *   &lt;summary&gt;
 *     &lt;b&gt;
 *       &lt;i&gt;VPC peering post v1.16.0 (Managed VPC)&lt;/i&gt;
 *     &lt;/b&gt;
 *   &lt;/summary&gt;
 * ```java
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
 *         var vpcPeeringRequest = new VpcGcpPeering(&#34;vpcPeeringRequest&#34;, VpcGcpPeeringArgs.builder()        
 *             .vpcId(cloudamqp_vpc.vpc().id())
 *             .peerNetworkUri(var_.peer_network_uri())
 *             .build());
 * 
 *         var firewallSettings = new SecurityFirewall(&#34;firewallSettings&#34;, SecurityFirewallArgs.builder()        
 *             .instanceId(cloudamqp_instance.instance().id())
 *             .rules(            
 *                 SecurityFirewallRuleArgs.builder()
 *                     .ip(var_.peer_subnet())
 *                     .ports(15672)
 *                     .services(                    
 *                         &#34;AMQP&#34;,
 *                         &#34;AMQPS&#34;,
 *                         &#34;STREAM&#34;,
 *                         &#34;STREAM_SSL&#34;)
 *                     .description(&#34;VPC peering for &lt;NETWORK&gt;&#34;)
 *                     .build(),
 *                 SecurityFirewallRuleArgs.builder()
 *                     .ip(&#34;192.168.0.0/24&#34;)
 *                     .ports(                    
 *                         4567,
 *                         4568)
 *                     .services(                    
 *                         &#34;AMQP&#34;,
 *                         &#34;AMQPS&#34;,
 *                         &#34;HTTPS&#34;)
 *                     .build())
 *             .build(), CustomResourceOptions.builder()
 *                 .dependsOn(vpcPeeringRequest)
 *                 .build());
 * 
 *     }
 * }
 * ```
 * &lt;/details&gt;
 * ## Depedency
 * 
 * *Pre v1.16.0*
 * This resource depends on CloudAMQP instance identifier, `cloudamqp_instance.instance.id`.
 * 
 * *Post v1.16.0*
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
    @Export(name="autoCreateRoutes", type=Boolean.class, parameters={})
    private Output<Boolean> autoCreateRoutes;

    /**
     * @return VPC peering auto created routes
     * 
     */
    public Output<Boolean> autoCreateRoutes() {
        return this.autoCreateRoutes;
    }
    /**
     * The CloudAMQP instance identifier.
     * 
     * ***Depreacted: Changed from required to optional in v1.16.0, will be removed in next major version (v2.0)***
     * 
     */
    @Export(name="instanceId", type=Integer.class, parameters={})
    private Output</* @Nullable */ Integer> instanceId;

    /**
     * @return The CloudAMQP instance identifier.
     * 
     * ***Depreacted: Changed from required to optional in v1.16.0, will be removed in next major version (v2.0)***
     * 
     */
    public Output<Optional<Integer>> instanceId() {
        return Codegen.optional(this.instanceId);
    }
    /**
     * Network uri of the VPC network to which you will peer with.
     * 
     */
    @Export(name="peerNetworkUri", type=String.class, parameters={})
    private Output<String> peerNetworkUri;

    /**
     * @return Network uri of the VPC network to which you will peer with.
     * 
     */
    public Output<String> peerNetworkUri() {
        return this.peerNetworkUri;
    }
    /**
     * VPC peering state
     * 
     */
    @Export(name="state", type=String.class, parameters={})
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
    @Export(name="stateDetails", type=String.class, parameters={})
    private Output<String> stateDetails;

    /**
     * @return VPC peering state details
     * 
     */
    public Output<String> stateDetails() {
        return this.stateDetails;
    }
    /**
     * The managed VPC identifier.
     * 
     * ***Note: Added as optional in version v1.16.0, will be required in next major version (v2.0)***
     * 
     */
    @Export(name="vpcId", type=String.class, parameters={})
    private Output</* @Nullable */ String> vpcId;

    /**
     * @return The managed VPC identifier.
     * 
     * ***Note: Added as optional in version v1.16.0, will be required in next major version (v2.0)***
     * 
     */
    public Output<Optional<String>> vpcId() {
        return Codegen.optional(this.vpcId);
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
