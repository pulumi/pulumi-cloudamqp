// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.cloudamqp;

import com.pulumi.cloudamqp.PrivatelinkAwsArgs;
import com.pulumi.cloudamqp.Utilities;
import com.pulumi.cloudamqp.inputs.PrivatelinkAwsState;
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
 * Enable PrivateLink for a CloudAMQP instance hosted in AWS. If no existing VPC available when enable PrivateLink, a new VPC will be created with subnet `10.52.72.0/24`.
 * 
 * More information about [CloudAMQP Privatelink](https://www.cloudamqp.com/docs/cloudamqp-privatelink.html#aws-privatelink).
 * 
 * Only available for dedicated subscription plans.
 * 
 * Pricing is available at [cloudamqp.com](https://www.cloudamqp.com/plans.html).
 * 
 * ## Example Usage
 * 
 * CloudAMQP instance without existing VPC
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.cloudamqp.Instance;
 * import com.pulumi.cloudamqp.InstanceArgs;
 * import com.pulumi.cloudamqp.PrivatelinkAws;
 * import com.pulumi.cloudamqp.PrivatelinkAwsArgs;
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
 *         var instance = new Instance(&#34;instance&#34;, InstanceArgs.builder()        
 *             .plan(&#34;squirrel-1&#34;)
 *             .region(&#34;amazon-web-services::us-west-1&#34;)
 *             .tags(&#34;test&#34;)
 *             .rmqVersion(&#34;3.10.8&#34;)
 *             .build());
 * 
 *         var privatelink = new PrivatelinkAws(&#34;privatelink&#34;, PrivatelinkAwsArgs.builder()        
 *             .instanceId(instance.id())
 *             .allowedPrincipals(&#34;arn:aws:iam::aws-account-id:user/user-name&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * CloudAMQP instance already in an existing VPC.
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.cloudamqp.Vpc;
 * import com.pulumi.cloudamqp.VpcArgs;
 * import com.pulumi.cloudamqp.Instance;
 * import com.pulumi.cloudamqp.InstanceArgs;
 * import com.pulumi.cloudamqp.PrivatelinkAws;
 * import com.pulumi.cloudamqp.PrivatelinkAwsArgs;
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
 *         var vpc = new Vpc(&#34;vpc&#34;, VpcArgs.builder()        
 *             .region(&#34;amazon-web-services::us-west-1&#34;)
 *             .subnet(&#34;10.56.72.0/24&#34;)
 *             .tags(&#34;test&#34;)
 *             .build());
 * 
 *         var instance = new Instance(&#34;instance&#34;, InstanceArgs.builder()        
 *             .plan(&#34;squirrel-1&#34;)
 *             .region(&#34;amazon-web-services::us-west-1&#34;)
 *             .tags(&#34;test&#34;)
 *             .rmqVersion(&#34;3.10.8&#34;)
 *             .vpcId(vpc.id())
 *             .keepAssociatedVpc(true)
 *             .build());
 * 
 *         var privatelink = new PrivatelinkAws(&#34;privatelink&#34;, PrivatelinkAwsArgs.builder()        
 *             .instanceId(instance.id())
 *             .allowedPrincipals(&#34;arn:aws:iam::aws-account-id:user/user-name&#34;)
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
 * `cloudamqp_privatelink_aws` can be imported using CloudAMQP internal identifier.
 * 
 * ```sh
 *  $ pulumi import cloudamqp:index/privatelinkAws:PrivatelinkAws privatelink &lt;id&gt;`
 * ```
 * 
 */
@ResourceType(type="cloudamqp:index/privatelinkAws:PrivatelinkAws")
public class PrivatelinkAws extends com.pulumi.resources.CustomResource {
    /**
     * Covering availability zones used when creating an Endpoint from other VPC.
     * 
     */
    @Export(name="activeZones", type=List.class, parameters={String.class})
    private Output<List<String>> activeZones;

    /**
     * @return Covering availability zones used when creating an Endpoint from other VPC.
     * 
     */
    public Output<List<String>> activeZones() {
        return this.activeZones;
    }
    /**
     * Allowed principals to access the endpoint service.
     * 
     */
    @Export(name="allowedPrincipals", type=List.class, parameters={String.class})
    private Output<List<String>> allowedPrincipals;

    /**
     * @return Allowed principals to access the endpoint service.
     * 
     */
    public Output<List<String>> allowedPrincipals() {
        return this.allowedPrincipals;
    }
    /**
     * The CloudAMQP instance identifier.
     * 
     */
    @Export(name="instanceId", type=Integer.class, parameters={})
    private Output<Integer> instanceId;

    /**
     * @return The CloudAMQP instance identifier.
     * 
     */
    public Output<Integer> instanceId() {
        return this.instanceId;
    }
    /**
     * Service name of the PrivateLink used when creating the endpoint from other VPC.
     * 
     */
    @Export(name="serviceName", type=String.class, parameters={})
    private Output<String> serviceName;

    /**
     * @return Service name of the PrivateLink used when creating the endpoint from other VPC.
     * 
     */
    public Output<String> serviceName() {
        return this.serviceName;
    }
    /**
     * Configurable sleep time (seconds) when enable PrivateLink. Default set to 60 seconds.
     * 
     */
    @Export(name="sleep", type=Integer.class, parameters={})
    private Output</* @Nullable */ Integer> sleep;

    /**
     * @return Configurable sleep time (seconds) when enable PrivateLink. Default set to 60 seconds.
     * 
     */
    public Output<Optional<Integer>> sleep() {
        return Codegen.optional(this.sleep);
    }
    /**
     * PrivateLink status [enable, pending, disable]
     * 
     */
    @Export(name="status", type=String.class, parameters={})
    private Output<String> status;

    /**
     * @return PrivateLink status [enable, pending, disable]
     * 
     */
    public Output<String> status() {
        return this.status;
    }
    /**
     * Configurable timeout time (seconds) when enable PrivateLink. Default set to 3600 seconds.
     * 
     */
    @Export(name="timeout", type=Integer.class, parameters={})
    private Output</* @Nullable */ Integer> timeout;

    /**
     * @return Configurable timeout time (seconds) when enable PrivateLink. Default set to 3600 seconds.
     * 
     */
    public Output<Optional<Integer>> timeout() {
        return Codegen.optional(this.timeout);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public PrivatelinkAws(String name) {
        this(name, PrivatelinkAwsArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public PrivatelinkAws(String name, PrivatelinkAwsArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public PrivatelinkAws(String name, PrivatelinkAwsArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("cloudamqp:index/privatelinkAws:PrivatelinkAws", name, args == null ? PrivatelinkAwsArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private PrivatelinkAws(String name, Output<String> id, @Nullable PrivatelinkAwsState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("cloudamqp:index/privatelinkAws:PrivatelinkAws", name, state, makeResourceOptions(options, id));
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
    public static PrivatelinkAws get(String name, Output<String> id, @Nullable PrivatelinkAwsState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new PrivatelinkAws(name, id, state, options);
    }
}