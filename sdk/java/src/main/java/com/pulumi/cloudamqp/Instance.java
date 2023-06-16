// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.cloudamqp;

import com.pulumi.cloudamqp.InstanceArgs;
import com.pulumi.cloudamqp.Utilities;
import com.pulumi.cloudamqp.inputs.InstanceState;
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
 * This resource allows you to create and manage a CloudAMQP instance running either [**RabbitMQ**](https://www.rabbitmq.com/) or [**LavinMQ**](https://lavinmq.com/) and can be deployed to multiple cloud platforms provider and regions, see Instance regions for more information.
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
 * ```java
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
 *         var lemurInstance = new Instance(&#34;lemurInstance&#34;, InstanceArgs.builder()        
 *             .plan(&#34;lemur&#34;)
 *             .region(&#34;amazon-web-services::us-west-1&#34;)
 *             .tags(&#34;rabbitmq&#34;)
 *             .build());
 * 
 *         var lemmingInstance = new Instance(&#34;lemmingInstance&#34;, InstanceArgs.builder()        
 *             .plan(&#34;lemming&#34;)
 *             .region(&#34;amazon-web-services::us-west-1&#34;)
 *             .tags(&#34;lavinmq&#34;)
 *             .build());
 * 
 *         var instance = new Instance(&#34;instance&#34;, InstanceArgs.builder()        
 *             .plan(&#34;bunny-1&#34;)
 *             .region(&#34;amazon-web-services::us-west-1&#34;)
 *             .tags(&#34;terraform&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * &lt;/details&gt;
 * 
 * &lt;details&gt;
 *   &lt;summary&gt;
 *     &lt;b&gt;
 *       &lt;i&gt;Dedicated instance using attribute vpc_subnet to create VPC, pre v1.16.0&lt;/i&gt;
 *     &lt;/b&gt;
 *   &lt;/summary&gt;
 * ```java
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
 *         var instance = new Instance(&#34;instance&#34;, InstanceArgs.builder()        
 *             .plan(&#34;bunny-1&#34;)
 *             .region(&#34;amazon-web-services::us-west-1&#34;)
 *             .tags(&#34;terraform&#34;)
 *             .vpcSubnet(&#34;10.56.72.0/24&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * &lt;/details&gt;
 * 
 * &lt;details&gt;
 *   &lt;summary&gt;
 *     &lt;b&gt;
 *       &lt;i&gt;Dedicated instance using attribute vpc_subnet to create VPC and then import managed VPC, post v1.16.0 (Managed VPC)&lt;/i&gt;
 *     &lt;/b&gt;
 *   &lt;/summary&gt;
 * ```java
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
 *         var instance01 = new Instance(&#34;instance01&#34;, InstanceArgs.builder()        
 *             .plan(&#34;bunny-1&#34;)
 *             .region(&#34;amazon-web-services::us-west-1&#34;)
 *             .tags(&#34;terraform&#34;)
 *             .vpcSubnet(&#34;10.56.72.0/24&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * Once the instance and the VPC are created, the VPC can be imported as managed VPC and added to the configuration file.
 * Set attribute `vpc_id` to the managed VPC identifier. To keep the managed VPC when deleting the instance, set attribute `keep_associated_vpc` to true.
 * For more information see guide Managed VPC.
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
 *             .region(&#34;amazon-web-services::us-east-1&#34;)
 *             .subnet(&#34;10.56.72.0/24&#34;)
 *             .tags()
 *             .build());
 * 
 *         var instance01 = new Instance(&#34;instance01&#34;, InstanceArgs.builder()        
 *             .plan(&#34;bunny-1&#34;)
 *             .region(&#34;amazon-web-services::us-west-1&#34;)
 *             .tags(&#34;terraform&#34;)
 *             .vpcId(vpc.id())
 *             .keepAssociatedVpc(true)
 *             .build());
 * 
 *     }
 * }
 * ```
 * &lt;/details&gt;
 * 
 * &lt;details&gt;
 *   &lt;summary&gt;
 *     &lt;b&gt;
 *       &lt;i&gt;Dedicated instances and managed VPC, post v1.16.0 (Managed VPC)&lt;/i&gt;
 *     &lt;/b&gt;
 *   &lt;/summary&gt;
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
 *             .region(&#34;amazon-web-services::us-east-1&#34;)
 *             .subnet(&#34;10.56.72.0/24&#34;)
 *             .tags()
 *             .build());
 * 
 *         var instance01 = new Instance(&#34;instance01&#34;, InstanceArgs.builder()        
 *             .plan(&#34;bunny-1&#34;)
 *             .region(&#34;amazon-web-services::us-west-1&#34;)
 *             .tags(&#34;terraform&#34;)
 *             .vpcId(vpc.id())
 *             .keepAssociatedVpc(true)
 *             .build());
 * 
 *         var instance02 = new Instance(&#34;instance02&#34;, InstanceArgs.builder()        
 *             .plan(&#34;bunny-1&#34;)
 *             .region(&#34;amazon-web-services::us-west-1&#34;)
 *             .tags(&#34;terraform&#34;)
 *             .vpcId(vpc.id())
 *             .keepAssociatedVpc(true)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * Set attribute `keep_associated_vpc` to true, will keep managed VPC when deleting the instances.
 * &lt;/details&gt;
 * 
 * ## Import
 * 
 * `cloudamqp_instance`can be imported using CloudAMQP internal identifier.
 * 
 * ```sh
 *  $ pulumi import cloudamqp:index/instance:Instance instance &lt;id&gt;`
 * ```
 * 
 *  To retrieve the identifier for a VPC, either use [CloudAMQP customer API](https://docs.cloudamqp.com/#list-instances). Or use the data source `cloudamqp_account` to list all available instances for an account.
 * 
 */
@ResourceType(type="cloudamqp:index/instance:Instance")
public class Instance extends com.pulumi.resources.CustomResource {
    /**
     * API key needed to communicate to CloudAMQP&#39;s second API. The second API is used to manage alarms, integration and more, full description [CloudAMQP API](https://docs.cloudamqp.com/cloudamqp_api.html).
     * 
     */
    @Export(name="apikey", type=String.class, parameters={})
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
    @Export(name="backend", type=String.class, parameters={})
    private Output<String> backend;

    /**
     * @return Information if the CloudAMQP instance runs either RabbitMQ or LavinMQ.
     * 
     */
    public Output<String> backend() {
        return this.backend;
    }
    /**
     * Information if the CloudAMQP instance is shared or dedicated.
     * 
     */
    @Export(name="dedicated", type=Boolean.class, parameters={})
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
    @Export(name="host", type=String.class, parameters={})
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
    @Export(name="hostInternal", type=String.class, parameters={})
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
    @Export(name="keepAssociatedVpc", type=Boolean.class, parameters={})
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
    @Export(name="name", type=String.class, parameters={})
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
    @Export(name="noDefaultAlarms", type=Boolean.class, parameters={})
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
    @Export(name="nodes", type=Integer.class, parameters={})
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
    @Export(name="plan", type=String.class, parameters={})
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
    @Export(name="ready", type=Boolean.class, parameters={})
    private Output<Boolean> ready;

    /**
     * @return Flag describing if the resource is ready
     * 
     */
    public Output<Boolean> ready() {
        return this.ready;
    }
    /**
     * The region to host the instance in. See Instance regions
     * 
     * ***Note: Changing region will force the instance to be destroyed and a new created in the new region. All data will be lost and a new name assigned.***
     * 
     */
    @Export(name="region", type=String.class, parameters={})
    private Output<String> region;

    /**
     * @return The region to host the instance in. See Instance regions
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
    @Export(name="rmqVersion", type=String.class, parameters={})
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
    @Export(name="tags", type=List.class, parameters={String.class})
    private Output</* @Nullable */ List<String>> tags;

    /**
     * @return One or more tags for the CloudAMQP instance, makes it possible to categories multiple instances in console view. Default there is no tags assigned.
     * 
     */
    public Output<Optional<List<String>>> tags() {
        return Codegen.optional(this.tags);
    }
    /**
     * The AMQP URL (uses the internal hostname if the instance was created with VPC). Has the format: `amqps://{username}:{password}@{hostname}/{vhost}`
     * 
     */
    @Export(name="url", type=String.class, parameters={})
    private Output<String> url;

    /**
     * @return The AMQP URL (uses the internal hostname if the instance was created with VPC). Has the format: `amqps://{username}:{password}@{hostname}/{vhost}`
     * 
     */
    public Output<String> url() {
        return this.url;
    }
    /**
     * The virtual host used by Rabbit MQ.
     * 
     */
    @Export(name="vhost", type=String.class, parameters={})
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
    @Export(name="vpcId", type=Integer.class, parameters={})
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
    @Export(name="vpcSubnet", type=String.class, parameters={})
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
    public Instance(String name) {
        this(name, InstanceArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Instance(String name, InstanceArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Instance(String name, InstanceArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("cloudamqp:index/instance:Instance", name, args == null ? InstanceArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Instance(String name, Output<String> id, @Nullable InstanceState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("cloudamqp:index/instance:Instance", name, state, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
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
    public static Instance get(String name, Output<String> id, @Nullable InstanceState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Instance(name, id, state, options);
    }
}
