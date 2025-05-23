// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.cloudamqp;

import com.pulumi.cloudamqp.IntegrationAwsEventbridgeArgs;
import com.pulumi.cloudamqp.Utilities;
import com.pulumi.cloudamqp.inputs.IntegrationAwsEventbridgeState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import javax.annotation.Nullable;

/**
 * This resource allows you to create and manage, an [AWS EventBridge] for a CloudAMQP instance. Once
 * created, continue to map the EventBridge in the [AWS Eventbridge console].
 * 
 * &gt;  Our consumer needs to have exclusive usage to the configured queue and the maximum body size
 * allowed on msgs by AWS is 256kb. The message body has to be valid JSON for AWS Eventbridge to accept
 * it. If messages are too large or are not valid JSON, they will be rejected (tip: setup a dead-letter
 * queue to catch them).
 * 
 * Not possible to update this resource. Any changes made to the argument will destroy and recreate the
 * resource. Hence why all arguments use ForceNew.
 * 
 * Only available for dedicated subscription plans.
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
 * import com.pulumi.cloudamqp.Instance;
 * import com.pulumi.cloudamqp.InstanceArgs;
 * import com.pulumi.cloudamqp.IntegrationAwsEventbridge;
 * import com.pulumi.cloudamqp.IntegrationAwsEventbridgeArgs;
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
 *             .name("Test instance")
 *             .plan("penguin-1")
 *             .region("amazon-web-services::us-west-1")
 *             .rmqVersion("3.11.5")
 *             .tags("aws")
 *             .build());
 * 
 *         var this_ = new IntegrationAwsEventbridge("this", IntegrationAwsEventbridgeArgs.builder()
 *             .instanceId(instance.id())
 *             .vhost(instance.vhost())
 *             .queue("<QUEUE-NAME>")
 *             .awsAccountId("<AWS-ACCOUNT-ID>")
 *             .awsRegion("us-west-1")
 *             .withHeaders(true)
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Argument References
 * 
 * The following arguments are supported:
 * 
 * * `aws_account_id`  - (ForceNew/Required) The 12 digit AWS Account ID where you want the events to
 *                       be sent to.
 * * `aws_region`      - (ForceNew/Required) The AWS region where you the events to be sent to.
 *                       (e.g. us-west-1, us-west-2, ..., etc.)
 * * `vhost`           - (ForceNew/Required) The VHost the queue resides in.
 * * `queue`           - (ForceNew/Required) A (durable) queue on your RabbitMQ instance.
 * * `with_headers`    - (ForceNew/Required) Include message headers in the event data.
 *                       `({ &#34;headers&#34;: { }, &#34;body&#34;: { &#34;your&#34;: &#34;message&#34; } })`
 * 
 * ## Dependency
 * 
 * This resource depends on CloudAMQP instance identifier, `cloudamqp_instance.instance.id`.
 * 
 * ## Import
 * 
 * `cloudamqp_integration_aws_eventbridge` can be imported using the resource identifier together with
 * 
 * CloudAMQP instance identifier (CSV separated). To retrieve the resource identifier, use
 * 
 * [CloudAMQP API list eventbridges].
 * 
 * From Terraform v1.5.0, the `import` block can be used to import this resource:
 * 
 * hcl
 * 
 * import {
 * 
 *   to = cloudamqp_integration_aws_eventbridge.this
 * 
 *   id = format(&#34;&lt;id&gt;,%s&#34;, cloudamqp_instance.instance.id)
 * 
 * }
 * 
 * Or with Terraform CLI:
 * 
 * ```sh
 * $ pulumi import cloudamqp:index/integrationAwsEventbridge:IntegrationAwsEventbridge this &lt;id&gt;,&lt;instance_id&gt;`
 * ```
 * 
 * [AWS EventBridge]: https://aws.amazon.com/eventbridge
 * 
 * [AWS Eventbridge console]: https://console.aws.amazon.com/events/home
 * 
 * [CloudAMQP API list eventbridges]: https://docs.cloudamqp.com/cloudamqp_api.html#list-eventbridges
 * 
 */
@ResourceType(type="cloudamqp:index/integrationAwsEventbridge:IntegrationAwsEventbridge")
public class IntegrationAwsEventbridge extends com.pulumi.resources.CustomResource {
    /**
     * The 12 digit AWS Account ID where you want the events to be sent to.
     * 
     */
    @Export(name="awsAccountId", refs={String.class}, tree="[0]")
    private Output<String> awsAccountId;

    /**
     * @return The 12 digit AWS Account ID where you want the events to be sent to.
     * 
     */
    public Output<String> awsAccountId() {
        return this.awsAccountId;
    }
    /**
     * The AWS region where you the events to be sent to. (e.g. us-west-1, us-west-2, ..., etc.)
     * 
     */
    @Export(name="awsRegion", refs={String.class}, tree="[0]")
    private Output<String> awsRegion;

    /**
     * @return The AWS region where you the events to be sent to. (e.g. us-west-1, us-west-2, ..., etc.)
     * 
     */
    public Output<String> awsRegion() {
        return this.awsRegion;
    }
    /**
     * Instance identifier
     * 
     */
    @Export(name="instanceId", refs={Integer.class}, tree="[0]")
    private Output<Integer> instanceId;

    /**
     * @return Instance identifier
     * 
     */
    public Output<Integer> instanceId() {
        return this.instanceId;
    }
    /**
     * A (durable) queue on your RabbitMQ instance.
     * 
     */
    @Export(name="queue", refs={String.class}, tree="[0]")
    private Output<String> queue;

    /**
     * @return A (durable) queue on your RabbitMQ instance.
     * 
     */
    public Output<String> queue() {
        return this.queue;
    }
    /**
     * Always set to null, unless there is an error starting the EventBridge.
     * 
     */
    @Export(name="status", refs={String.class}, tree="[0]")
    private Output<String> status;

    /**
     * @return Always set to null, unless there is an error starting the EventBridge.
     * 
     */
    public Output<String> status() {
        return this.status;
    }
    /**
     * The VHost the queue resides in.
     * 
     */
    @Export(name="vhost", refs={String.class}, tree="[0]")
    private Output<String> vhost;

    /**
     * @return The VHost the queue resides in.
     * 
     */
    public Output<String> vhost() {
        return this.vhost;
    }
    /**
     * Include message headers in the event data.
     * 
     */
    @Export(name="withHeaders", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> withHeaders;

    /**
     * @return Include message headers in the event data.
     * 
     */
    public Output<Boolean> withHeaders() {
        return this.withHeaders;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public IntegrationAwsEventbridge(java.lang.String name) {
        this(name, IntegrationAwsEventbridgeArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public IntegrationAwsEventbridge(java.lang.String name, IntegrationAwsEventbridgeArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public IntegrationAwsEventbridge(java.lang.String name, IntegrationAwsEventbridgeArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("cloudamqp:index/integrationAwsEventbridge:IntegrationAwsEventbridge", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private IntegrationAwsEventbridge(java.lang.String name, Output<java.lang.String> id, @Nullable IntegrationAwsEventbridgeState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("cloudamqp:index/integrationAwsEventbridge:IntegrationAwsEventbridge", name, state, makeResourceOptions(options, id), false);
    }

    private static IntegrationAwsEventbridgeArgs makeArgs(IntegrationAwsEventbridgeArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? IntegrationAwsEventbridgeArgs.Empty : args;
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
    public static IntegrationAwsEventbridge get(java.lang.String name, Output<java.lang.String> id, @Nullable IntegrationAwsEventbridgeState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new IntegrationAwsEventbridge(name, id, state, options);
    }
}
