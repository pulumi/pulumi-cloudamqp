// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.cloudamqp;

import com.pulumi.cloudamqp.PrivatelinkAzureArgs;
import com.pulumi.cloudamqp.Utilities;
import com.pulumi.cloudamqp.inputs.PrivatelinkAzureState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Enable PrivateLink for a CloudAMQP instance hosted in Azure. If no existing VPC available when
 * enable PrivateLink, a new VPC will be created with subnet `10.52.72.0/24`.
 * 
 * &gt; **Note:** Enabling PrivateLink will automatically add firewall rules for the peered subnet.
 * 
 * &lt;details&gt;
 *  &lt;summary&gt;
 *     &lt;i&gt;Default PrivateLink firewall rule&lt;/i&gt;
 *   &lt;/summary&gt;
 * 
 * ## Example Usage
 * 
 * &lt;details&gt;
 *   &lt;summary&gt;
 *     &lt;b&gt;
 *       &lt;i&gt;CloudAMQP instance without existing VPC&lt;/i&gt;
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
 * import com.pulumi.cloudamqp.PrivatelinkAzure;
 * import com.pulumi.cloudamqp.PrivatelinkAzureArgs;
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
 *             .name("Instance 01")
 *             .plan("bunny-1")
 *             .region("azure-arm::westus")
 *             .tags()
 *             .build());
 * 
 *         var privatelink = new PrivatelinkAzure("privatelink", PrivatelinkAzureArgs.builder()        
 *             .instanceId(instance.id())
 *             .approvedSubscriptions("XXXXXXXX-XXXX-XXXX-XXXX-XXXXXXXXXXXX")
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * &lt;/details&gt;
 * 
 * &lt;details&gt;
 *   &lt;summary&gt;
 *     &lt;b&gt;
 *       &lt;i&gt;CloudAMQP instance in an existing VPC&lt;/i&gt;
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
 * import com.pulumi.cloudamqp.PrivatelinkAzure;
 * import com.pulumi.cloudamqp.PrivatelinkAzureArgs;
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
 *         var vpc = new Vpc("vpc", VpcArgs.builder()        
 *             .name("Standalone VPC")
 *             .region("azure-arm::westus")
 *             .subnet("10.56.72.0/24")
 *             .tags()
 *             .build());
 * 
 *         var instance = new Instance("instance", InstanceArgs.builder()        
 *             .name("Instance 01")
 *             .plan("bunny-1")
 *             .region("azure-arm::westus")
 *             .tags()
 *             .vpcId(vpc.id())
 *             .keepAssociatedVpc(true)
 *             .build());
 * 
 *         var privatelink = new PrivatelinkAzure("privatelink", PrivatelinkAzureArgs.builder()        
 *             .instanceId(instance.id())
 *             .approvedSubscriptions("XXXXXXXX-XXXX-XXXX-XXXX-XXXXXXXXXXXX")
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
 *       &lt;i&gt;CloudAMQP instance in an existing VPC with managed firewall rules&lt;/i&gt;
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
 * import com.pulumi.cloudamqp.PrivatelinkAzure;
 * import com.pulumi.cloudamqp.PrivatelinkAzureArgs;
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
 *         var vpc = new Vpc("vpc", VpcArgs.builder()        
 *             .name("Standalone VPC")
 *             .region("azure-arm::westus")
 *             .subnet("10.56.72.0/24")
 *             .tags()
 *             .build());
 * 
 *         var instance = new Instance("instance", InstanceArgs.builder()        
 *             .name("Instance 01")
 *             .plan("bunny-1")
 *             .region("azure-arm::westus")
 *             .tags()
 *             .vpcId(vpc.id())
 *             .keepAssociatedVpc(true)
 *             .build());
 * 
 *         var privatelink = new PrivatelinkAzure("privatelink", PrivatelinkAzureArgs.builder()        
 *             .instanceId(instance.id())
 *             .approvedSubscriptions("XXXXXXXX-XXXX-XXXX-XXXX-XXXXXXXXXXXX")
 *             .build());
 * 
 *         var firewallSettings = new SecurityFirewall("firewallSettings", SecurityFirewallArgs.builder()        
 *             .instanceId(instance.id())
 *             .rules(            
 *                 SecurityFirewallRuleArgs.builder()
 *                     .description("Custom PrivateLink setup")
 *                     .ip(vpc.subnet())
 *                     .ports()
 *                     .services(                    
 *                         "AMQP",
 *                         "AMQPS",
 *                         "HTTPS",
 *                         "STREAM",
 *                         "STREAM_SSL")
 *                     .build(),
 *                 SecurityFirewallRuleArgs.builder()
 *                     .description("MGMT interface")
 *                     .ip("0.0.0.0/0")
 *                     .ports()
 *                     .services("HTTPS")
 *                     .build())
 *             .build(), CustomResourceOptions.builder()
 *                 .dependsOn(privatelink)
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
 * This resource depends on CloudAMQP instance identifier, `cloudamqp_instance.instance.id`.
 * 
 * ## Create PrivateLink with additional firewall rules
 * 
 * To create a PrivateLink configuration with additional firewall rules, it&#39;s required to chain the cloudamqp.SecurityFirewall
 * resource to avoid parallel conflicting resource calls. You can do this by making the firewall
 * resource depend on the PrivateLink resource, `cloudamqp_privatelink_azure.privatelink`.
 * 
 * Furthermore, since all firewall rules are overwritten, the otherwise automatically added rules for
 * the PrivateLink also needs to be added.
 * 
 * ## Import
 * 
 * `cloudamqp_privatelink_aws` can be imported using CloudAMQP internal identifier.
 * 
 * ```sh
 * $ pulumi import cloudamqp:index/privatelinkAzure:PrivatelinkAzure privatelink &lt;id&gt;`
 * ```
 * 
 * The resource uses the same identifier as the CloudAMQP instance. To retrieve the identifier for an instance, either use [CloudAMQP customer API](https://docs.cloudamqp.com/#list-instances) or use the data source [`cloudamqp_account`](./data-sources/account.md).
 * 
 */
@ResourceType(type="cloudamqp:index/privatelinkAzure:PrivatelinkAzure")
public class PrivatelinkAzure extends com.pulumi.resources.CustomResource {
    /**
     * Approved subscriptions to access the endpoint service.
     * See format below.
     * 
     */
    @Export(name="approvedSubscriptions", refs={List.class,String.class}, tree="[0,1]")
    private Output<List<String>> approvedSubscriptions;

    /**
     * @return Approved subscriptions to access the endpoint service.
     * See format below.
     * 
     */
    public Output<List<String>> approvedSubscriptions() {
        return this.approvedSubscriptions;
    }
    /**
     * The CloudAMQP instance identifier.
     * 
     */
    @Export(name="instanceId", refs={Integer.class}, tree="[0]")
    private Output<Integer> instanceId;

    /**
     * @return The CloudAMQP instance identifier.
     * 
     */
    public Output<Integer> instanceId() {
        return this.instanceId;
    }
    /**
     * Name of the server having the PrivateLink enabled.
     * 
     */
    @Export(name="serverName", refs={String.class}, tree="[0]")
    private Output<String> serverName;

    /**
     * @return Name of the server having the PrivateLink enabled.
     * 
     */
    public Output<String> serverName() {
        return this.serverName;
    }
    /**
     * Service name (alias) of the PrivateLink, needed when creating the endpoint.
     * 
     */
    @Export(name="serviceName", refs={String.class}, tree="[0]")
    private Output<String> serviceName;

    /**
     * @return Service name (alias) of the PrivateLink, needed when creating the endpoint.
     * 
     */
    public Output<String> serviceName() {
        return this.serviceName;
    }
    /**
     * Configurable sleep time (seconds) when enable PrivateLink.
     * Default set to 10 seconds. *Available from v1.29.0*
     * 
     */
    @Export(name="sleep", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> sleep;

    /**
     * @return Configurable sleep time (seconds) when enable PrivateLink.
     * Default set to 10 seconds. *Available from v1.29.0*
     * 
     */
    public Output<Optional<Integer>> sleep() {
        return Codegen.optional(this.sleep);
    }
    /**
     * PrivateLink status [enable, pending, disable]
     * 
     */
    @Export(name="status", refs={String.class}, tree="[0]")
    private Output<String> status;

    /**
     * @return PrivateLink status [enable, pending, disable]
     * 
     */
    public Output<String> status() {
        return this.status;
    }
    /**
     * Configurable timeout time (seconds) when enable PrivateLink.
     * Default set to 1800 seconds. *Available from v1.29.0*
     * 
     * Approved subscriptions format (GUID): &lt;br&gt;
     * `XXXXXXXX-XXXX-XXXX-XXXX-XXXXXXXXXXXX`
     * 
     */
    @Export(name="timeout", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> timeout;

    /**
     * @return Configurable timeout time (seconds) when enable PrivateLink.
     * Default set to 1800 seconds. *Available from v1.29.0*
     * 
     * Approved subscriptions format (GUID): &lt;br&gt;
     * `XXXXXXXX-XXXX-XXXX-XXXX-XXXXXXXXXXXX`
     * 
     */
    public Output<Optional<Integer>> timeout() {
        return Codegen.optional(this.timeout);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public PrivatelinkAzure(String name) {
        this(name, PrivatelinkAzureArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public PrivatelinkAzure(String name, PrivatelinkAzureArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public PrivatelinkAzure(String name, PrivatelinkAzureArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("cloudamqp:index/privatelinkAzure:PrivatelinkAzure", name, args == null ? PrivatelinkAzureArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private PrivatelinkAzure(String name, Output<String> id, @Nullable PrivatelinkAzureState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("cloudamqp:index/privatelinkAzure:PrivatelinkAzure", name, state, makeResourceOptions(options, id));
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
    public static PrivatelinkAzure get(String name, Output<String> id, @Nullable PrivatelinkAzureState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new PrivatelinkAzure(name, id, state, options);
    }
}
