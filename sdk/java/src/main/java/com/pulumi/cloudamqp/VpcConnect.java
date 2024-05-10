// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.cloudamqp;

import com.pulumi.cloudamqp.Utilities;
import com.pulumi.cloudamqp.VpcConnectArgs;
import com.pulumi.cloudamqp.inputs.VpcConnectState;
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
 * This resource is a generic way to handle PrivateLink (AWS and Azure) and Private Service Connect (GCP).
 * Communication between resources can be done just as they were living inside a VPC. CloudAMQP creates an Endpoint
 * Service to connect the VPC and creating a new network interface to handle the communicate.
 * 
 * If no existing VPC available when enable VPC connect, a new VPC will be created with subnet `10.52.72.0/24`.
 * 
 * More information can be found at: [CloudAMQP VPC Connect](https://www.cloudamqp.com/docs/cloudamqp-vpc-connect.html)
 * 
 * &gt; **Note:** Enabling VPC Connect will automatically add a firewall rule.
 * 
 * &lt;details&gt;
 *  &lt;summary&gt;
 *     &lt;b&gt;
 *       &lt;i&gt;Default PrivateLink firewall rule [AWS, Azure]&lt;/i&gt;
 *     &lt;/b&gt;
 *   &lt;/summary&gt;
 * 
 * ## Example Usage
 * 
 * &lt;details&gt;
 *   &lt;summary&gt;
 *     &lt;b&gt;
 *       &lt;i&gt;Enable VPC Connect (PrivateLink) in AWS&lt;/i&gt;
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
 * import com.pulumi.cloudamqp.VpcConnect;
 * import com.pulumi.cloudamqp.VpcConnectArgs;
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
 *             .region("amazon-web-services::us-west-1")
 *             .subnet("10.56.72.0/24")
 *             .tags()
 *             .build());
 * 
 *         var instance = new Instance("instance", InstanceArgs.builder()        
 *             .name("Instance 01")
 *             .plan("bunny-1")
 *             .region("amazon-web-services::us-west-1")
 *             .tags()
 *             .vpcId(vpc.id())
 *             .keepAssociatedVpc(true)
 *             .build());
 * 
 *         var vpcConnect = new VpcConnect("vpcConnect", VpcConnectArgs.builder()        
 *             .instanceId(instance.id())
 *             .region(instance.region())
 *             .allowedPrincipals("arn:aws:iam::aws-account-id:user/user-name")
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
 *       &lt;i&gt;Enable VPC Connect (PrivateLink) in Azure&lt;/i&gt;
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
 * import com.pulumi.cloudamqp.VpcConnect;
 * import com.pulumi.cloudamqp.VpcConnectArgs;
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
 *         var vpcConnect = new VpcConnect("vpcConnect", VpcConnectArgs.builder()        
 *             .instanceId(instance.id())
 *             .region(instance.region())
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
 * &lt;details&gt;
 *   &lt;summary&gt;
 *     &lt;b&gt;
 *       &lt;i&gt;Enable VPC Connect (Private Service Connect) in GCP&lt;/i&gt;
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
 * import com.pulumi.cloudamqp.VpcConnect;
 * import com.pulumi.cloudamqp.VpcConnectArgs;
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
 *             .region("google-compute-engine::us-west1")
 *             .subnet("10.56.72.0/24")
 *             .tags()
 *             .build());
 * 
 *         var instance = new Instance("instance", InstanceArgs.builder()        
 *             .name("Instance 01")
 *             .plan("bunny-1")
 *             .region("google-compute-engine::us-west1")
 *             .tags()
 *             .vpcId(vpc.id())
 *             .keepAssociatedVpc(true)
 *             .build());
 * 
 *         var vpcConnect = new VpcConnect("vpcConnect", VpcConnectArgs.builder()        
 *             .instanceId(instance.id())
 *             .region(instance.region())
 *             .allowedProjects("some-project-123456")
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
 * import com.pulumi.cloudamqp.VpcConnect;
 * import com.pulumi.cloudamqp.VpcConnectArgs;
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
 *             .region("amazon-web-services::us-west-1")
 *             .subnet("10.56.72.0/24")
 *             .tags()
 *             .build());
 * 
 *         var instance = new Instance("instance", InstanceArgs.builder()        
 *             .name("Instance 01")
 *             .plan("bunny-1")
 *             .region("amazon-web-services::us-west-1")
 *             .tags()
 *             .vpcId(vpc.id())
 *             .keepAssociatedVpc(true)
 *             .build());
 * 
 *         var vpcConnect = new VpcConnect("vpcConnect", VpcConnectArgs.builder()        
 *             .instanceId(instance.id())
 *             .allowedPrincipals("arn:aws:iam::aws-account-id:user/user-name")
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
 *                 .dependsOn(vpcConnect)
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
 * Since `region` also is required, suggest to reuse the argument from CloudAMQP instance,
 * `cloudamqp_instance.instance.region`.
 * 
 * ## Create VPC Connect with additional firewall rules
 * 
 * To create a PrivateLink/Private Service Connect configuration with additional firewall rules, it&#39;s required to chain the cloudamqp.SecurityFirewall
 * resource to avoid parallel conflicting resource calls. You can do this by making the firewall
 * resource depend on the VPC Connect resource, `cloudamqp_vpc_connect.vpc_connect`.
 * 
 * Furthermore, since all firewall rules are overwritten, the otherwise automatically added rules for
 * the VPC Connect also needs to be added.
 * 
 * ## Import
 * 
 * `cloudamqp_vpc_connect` can be imported using CloudAMQP internal identifier.
 * 
 * ```sh
 * $ pulumi import cloudamqp:index/vpcConnect:VpcConnect vpc_connect &lt;id&gt;`
 * ```
 * 
 * The resource uses the same identifier as the CloudAMQP instance. To retrieve the identifier for an instance, either use [CloudAMQP customer API](https://docs.cloudamqp.com/#list-instances) or use the data source [`cloudamqp_account`](./data-sources/account.md).
 * 
 */
@ResourceType(type="cloudamqp:index/vpcConnect:VpcConnect")
public class VpcConnect extends com.pulumi.resources.CustomResource {
    /**
     * Covering availability zones used when creating an endpoint from other VPC. (AWS)
     * 
     */
    @Export(name="activeZones", refs={List.class,String.class}, tree="[0,1]")
    private Output<List<String>> activeZones;

    /**
     * @return Covering availability zones used when creating an endpoint from other VPC. (AWS)
     * 
     */
    public Output<List<String>> activeZones() {
        return this.activeZones;
    }
    /**
     * List of allowed prinicpals used by AWS, see below table.
     * 
     */
    @Export(name="allowedPrincipals", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> allowedPrincipals;

    /**
     * @return List of allowed prinicpals used by AWS, see below table.
     * 
     */
    public Output<Optional<List<String>>> allowedPrincipals() {
        return Codegen.optional(this.allowedPrincipals);
    }
    /**
     * List of allowed projects used by GCP, see below table.
     * 
     */
    @Export(name="allowedProjects", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> allowedProjects;

    /**
     * @return List of allowed projects used by GCP, see below table.
     * 
     */
    public Output<Optional<List<String>>> allowedProjects() {
        return Codegen.optional(this.allowedProjects);
    }
    /**
     * List of approved subscriptions used by Azure, see below table.
     * 
     */
    @Export(name="approvedSubscriptions", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> approvedSubscriptions;

    /**
     * @return List of approved subscriptions used by Azure, see below table.
     * 
     */
    public Output<Optional<List<String>>> approvedSubscriptions() {
        return Codegen.optional(this.approvedSubscriptions);
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
     * The region where the CloudAMQP instance is hosted.
     * 
     */
    @Export(name="region", refs={String.class}, tree="[0]")
    private Output<String> region;

    /**
     * @return The region where the CloudAMQP instance is hosted.
     * 
     */
    public Output<String> region() {
        return this.region;
    }
    /**
     * Service name (alias for Azure) of the PrivateLink.
     * 
     */
    @Export(name="serviceName", refs={String.class}, tree="[0]")
    private Output<String> serviceName;

    /**
     * @return Service name (alias for Azure) of the PrivateLink.
     * 
     */
    public Output<String> serviceName() {
        return this.serviceName;
    }
    /**
     * Configurable sleep time (seconds) when enable Private Service Connect.
     * Default set to 10 seconds.
     * 
     */
    @Export(name="sleep", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> sleep;

    /**
     * @return Configurable sleep time (seconds) when enable Private Service Connect.
     * Default set to 10 seconds.
     * 
     */
    public Output<Optional<Integer>> sleep() {
        return Codegen.optional(this.sleep);
    }
    /**
     * Private Service Connect status [enable, pending, disable]
     * 
     */
    @Export(name="status", refs={String.class}, tree="[0]")
    private Output<String> status;

    /**
     * @return Private Service Connect status [enable, pending, disable]
     * 
     */
    public Output<String> status() {
        return this.status;
    }
    /**
     * Configurable timeout time (seconds) when enable Private Service Connect.
     * Default set to 1800 seconds.
     * 
     * ***
     * 
     * The `allowed_principals`, `approved_subscriptions` or `allowed_projects` data depends on the provider platform:
     * 
     * | Platform | Description         | Format                                                                                                                             |
     * |----------|---------------------|------------------------------------------------------------------------------------------------------------------------------------|
     * | AWS      | IAM ARN principals  | arn:aws:iam::aws-account-id:root&lt;br /&gt; arn:aws:iam::aws-account-id:user/user-name&lt;br /&gt; arn:aws:iam::aws-account-id:role/role-name |
     * | Azure    | Subscription (GUID) | XXXXXXXX-XXXX-XXXX-XXXX-XXXXXXXXXXXX                                                                                               |
     * | GCP      | Project IDs*        | 6 to 30 lowercase letters, digits, or hyphens                                                                                      |
     * 
     * *https://cloud.google.com/resource-manager/reference/rest/v1/projects
     * 
     */
    @Export(name="timeout", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> timeout;

    /**
     * @return Configurable timeout time (seconds) when enable Private Service Connect.
     * Default set to 1800 seconds.
     * 
     * ***
     * 
     * The `allowed_principals`, `approved_subscriptions` or `allowed_projects` data depends on the provider platform:
     * 
     * | Platform | Description         | Format                                                                                                                             |
     * |----------|---------------------|------------------------------------------------------------------------------------------------------------------------------------|
     * | AWS      | IAM ARN principals  | arn:aws:iam::aws-account-id:root&lt;br /&gt; arn:aws:iam::aws-account-id:user/user-name&lt;br /&gt; arn:aws:iam::aws-account-id:role/role-name |
     * | Azure    | Subscription (GUID) | XXXXXXXX-XXXX-XXXX-XXXX-XXXXXXXXXXXX                                                                                               |
     * | GCP      | Project IDs*        | 6 to 30 lowercase letters, digits, or hyphens                                                                                      |
     * 
     * *https://cloud.google.com/resource-manager/reference/rest/v1/projects
     * 
     */
    public Output<Optional<Integer>> timeout() {
        return Codegen.optional(this.timeout);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public VpcConnect(String name) {
        this(name, VpcConnectArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public VpcConnect(String name, VpcConnectArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public VpcConnect(String name, VpcConnectArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("cloudamqp:index/vpcConnect:VpcConnect", name, args == null ? VpcConnectArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private VpcConnect(String name, Output<String> id, @Nullable VpcConnectState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("cloudamqp:index/vpcConnect:VpcConnect", name, state, makeResourceOptions(options, id));
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
    public static VpcConnect get(String name, Output<String> id, @Nullable VpcConnectState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new VpcConnect(name, id, state, options);
    }
}
