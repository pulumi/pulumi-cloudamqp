// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.cloudamqp;

import com.pulumi.cloudamqp.IntegrationMetricArgs;
import com.pulumi.cloudamqp.Utilities;
import com.pulumi.cloudamqp.inputs.IntegrationMetricState;
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
 * This resource allows you to create and manage, forwarding metrics to third party integrations for a CloudAMQP instance. Once configured, the metrics produced will be forward to corresponding integration.
 * 
 * Only available for dedicated subscription plans.
 * 
 * ## Example Usage
 * 
 * &lt;details&gt;
 *   &lt;summary&gt;
 *     &lt;b&gt;
 *       &lt;i&gt;Cloudwatch v1 and v2 metric integration&lt;/i&gt;
 *     &lt;/b&gt;
 *   &lt;/summary&gt;
 * 
 * ***Access key***
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.cloudamqp.IntegrationMetric;
 * import com.pulumi.cloudamqp.IntegrationMetricArgs;
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
 *         var cloudwatch = new IntegrationMetric("cloudwatch", IntegrationMetricArgs.builder()
 *             .instanceId(instance.id())
 *             .name("cloudwatch")
 *             .accessKeyId(awsAccessKeyId)
 *             .secretAccessKey(varAwsSecretAcccessKey)
 *             .region(awsRegion)
 *             .build());
 * 
 *         var cloudwatchV2 = new IntegrationMetric("cloudwatchV2", IntegrationMetricArgs.builder()
 *             .instanceId(instance.id())
 *             .name("cloudwatch_v2")
 *             .accessKeyId(awsAccessKeyId)
 *             .secretAccessKey(varAwsSecretAcccessKey)
 *             .region(awsRegion)
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ***Assume role***
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.cloudamqp.IntegrationMetric;
 * import com.pulumi.cloudamqp.IntegrationMetricArgs;
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
 *         var cloudwatch = new IntegrationMetric("cloudwatch", IntegrationMetricArgs.builder()
 *             .instanceId(instance.id())
 *             .name("cloudwatch")
 *             .iamRole(awsIamRole)
 *             .iamExternalId(externalId)
 *             .region(awsRegion)
 *             .build());
 * 
 *         var cloudwatchV2 = new IntegrationMetric("cloudwatchV2", IntegrationMetricArgs.builder()
 *             .instanceId(instance.id())
 *             .name("cloudwatch_v2")
 *             .iamRole(awsIamRole)
 *             .iamExternalId(externalId)
 *             .region(awsRegion)
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * * AWS IAM role: arn:aws:iam::ACCOUNT-ID:role/ROLE-NAME
 * * External id: Create own external identifier that match the role created. E.g. &#34;cloudamqp-abc123&#34;.
 * 
 * &lt;/details&gt;
 * 
 * &lt;details&gt;
 *   &lt;summary&gt;
 *     &lt;b&gt;
 *       &lt;i&gt;Datadog v1 and v2 metric integration&lt;/i&gt;
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
 * import com.pulumi.cloudamqp.IntegrationMetric;
 * import com.pulumi.cloudamqp.IntegrationMetricArgs;
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
 *         var datadog = new IntegrationMetric("datadog", IntegrationMetricArgs.builder()
 *             .instanceId(instance.id())
 *             .name("datadog")
 *             .apiKey(datadogApiKey)
 *             .region(datadogRegion)
 *             .tags("env=prod,region=us1,version=v1.0")
 *             .build());
 * 
 *         var datadogV2 = new IntegrationMetric("datadogV2", IntegrationMetricArgs.builder()
 *             .instanceId(instance.id())
 *             .name("datadog_v2")
 *             .apiKey(datadogApiKey)
 *             .region(datadogRegion)
 *             .tags("env=prod,region=us1,version=v1.0")
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
 *       &lt;i&gt;Librato metric integration&lt;/i&gt;
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
 * import com.pulumi.cloudamqp.IntegrationMetric;
 * import com.pulumi.cloudamqp.IntegrationMetricArgs;
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
 *         var librato = new IntegrationMetric("librato", IntegrationMetricArgs.builder()
 *             .instanceId(instance.id())
 *             .name("librato")
 *             .email(libratoEmail)
 *             .apiKey(libratoApiKey)
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
 *       &lt;i&gt;New relic v2 metric integration&lt;/i&gt;
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
 * import com.pulumi.cloudamqp.IntegrationMetric;
 * import com.pulumi.cloudamqp.IntegrationMetricArgs;
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
 *         var newrelic = new IntegrationMetric("newrelic", IntegrationMetricArgs.builder()
 *             .instanceId(instance.id())
 *             .name("newrelic_v2")
 *             .apiKey(newrelicApiKey)
 *             .region(newrelicRegion)
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
 *       &lt;i&gt;Stackdriver metric integration (v1.20.2 or earlier versions)&lt;/i&gt;
 *     &lt;/b&gt;
 *   &lt;/summary&gt;
 * 
 * Use variable file populated with project_id, private_key and client_email
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.cloudamqp.IntegrationMetric;
 * import com.pulumi.cloudamqp.IntegrationMetricArgs;
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
 *         var stackdriver = new IntegrationMetric("stackdriver", IntegrationMetricArgs.builder()
 *             .instanceId(instance.id())
 *             .name("stackdriver")
 *             .projectId(stackdriverProjectId)
 *             .privateKey(stackdriverPrivateKey)
 *             .clientEmail(stackriverEmail)
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * or by using google_service_account_key resource from Google provider
 * 
 * ## Import
 * 
 * `cloudamqp_integration_metric`can be imported using the resource identifier together with CloudAMQP instance identifier. The name and identifier are CSV separated, see example below.
 * 
 * ```sh
 * $ pulumi import cloudamqp:index/integrationMetric:IntegrationMetric &lt;resource_name&gt; &lt;resource_id&gt;,&lt;instance_id&gt;`
 * ```
 * 
 */
@ResourceType(type="cloudamqp:index/integrationMetric:IntegrationMetric")
public class IntegrationMetric extends com.pulumi.resources.CustomResource {
    /**
     * AWS access key identifier. (Cloudwatch)
     * 
     */
    @Export(name="accessKeyId", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> accessKeyId;

    /**
     * @return AWS access key identifier. (Cloudwatch)
     * 
     */
    public Output<Optional<String>> accessKeyId() {
        return Codegen.optional(this.accessKeyId);
    }
    /**
     * The API key for the integration service. (Librato)
     * 
     */
    @Export(name="apiKey", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> apiKey;

    /**
     * @return The API key for the integration service. (Librato)
     * 
     */
    public Output<Optional<String>> apiKey() {
        return Codegen.optional(this.apiKey);
    }
    /**
     * The client email. (Stackdriver)
     * 
     */
    @Export(name="clientEmail", refs={String.class}, tree="[0]")
    private Output<String> clientEmail;

    /**
     * @return The client email. (Stackdriver)
     * 
     */
    public Output<String> clientEmail() {
        return this.clientEmail;
    }
    /**
     * Base64Encoded credentials. (Stackdriver)
     * 
     */
    @Export(name="credentials", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> credentials;

    /**
     * @return Base64Encoded credentials. (Stackdriver)
     * 
     */
    public Output<Optional<String>> credentials() {
        return Codegen.optional(this.credentials);
    }
    /**
     * The email address registred for the integration service. (Librato)
     * 
     */
    @Export(name="email", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> email;

    /**
     * @return The email address registred for the integration service. (Librato)
     * 
     */
    public Output<Optional<String>> email() {
        return Codegen.optional(this.email);
    }
    /**
     * External identifier that match the role you created. (Cloudwatch)
     * 
     */
    @Export(name="iamExternalId", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> iamExternalId;

    /**
     * @return External identifier that match the role you created. (Cloudwatch)
     * 
     */
    public Output<Optional<String>> iamExternalId() {
        return Codegen.optional(this.iamExternalId);
    }
    /**
     * The ARN of the role to be assumed when publishing metrics. (Cloudwatch)
     * 
     */
    @Export(name="iamRole", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> iamRole;

    /**
     * @return The ARN of the role to be assumed when publishing metrics. (Cloudwatch)
     * 
     */
    public Output<Optional<String>> iamRole() {
        return Codegen.optional(this.iamRole);
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
     * The license key registred for the integration service. (New Relic)
     * 
     */
    @Export(name="licenseKey", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> licenseKey;

    /**
     * @return The license key registred for the integration service. (New Relic)
     * 
     */
    public Output<Optional<String>> licenseKey() {
        return Codegen.optional(this.licenseKey);
    }
    /**
     * The name of metrics integration
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return The name of metrics integration
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * The private key. (Stackdriver)
     * 
     */
    @Export(name="privateKey", refs={String.class}, tree="[0]")
    private Output<String> privateKey;

    /**
     * @return The private key. (Stackdriver)
     * 
     */
    public Output<String> privateKey() {
        return this.privateKey;
    }
    /**
     * Private key identifier. (Stackdriver)
     * 
     */
    @Export(name="privateKeyId", refs={String.class}, tree="[0]")
    private Output<String> privateKeyId;

    /**
     * @return Private key identifier. (Stackdriver)
     * 
     */
    public Output<String> privateKeyId() {
        return this.privateKeyId;
    }
    /**
     * Project ID. (Stackdriver)
     * 
     */
    @Export(name="projectId", refs={String.class}, tree="[0]")
    private Output<String> projectId;

    /**
     * @return Project ID. (Stackdriver)
     * 
     */
    public Output<String> projectId() {
        return this.projectId;
    }
    /**
     * (optional) allowlist using regular expression
     * 
     */
    @Export(name="queueAllowlist", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> queueAllowlist;

    /**
     * @return (optional) allowlist using regular expression
     * 
     */
    public Output<Optional<String>> queueAllowlist() {
        return Codegen.optional(this.queueAllowlist);
    }
    /**
     * **Deprecated**
     * 
     * @deprecated
     * use queue_allowlist instead
     * 
     */
    @Deprecated /* use queue_allowlist instead */
    @Export(name="queueWhitelist", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> queueWhitelist;

    /**
     * @return **Deprecated**
     * 
     */
    public Output<Optional<String>> queueWhitelist() {
        return Codegen.optional(this.queueWhitelist);
    }
    /**
     * AWS region for Cloudwatch and [US/EU] for Data dog/New relic. (Cloudwatch, Data Dog, New Relic)
     * 
     */
    @Export(name="region", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> region;

    /**
     * @return AWS region for Cloudwatch and [US/EU] for Data dog/New relic. (Cloudwatch, Data Dog, New Relic)
     * 
     */
    public Output<Optional<String>> region() {
        return Codegen.optional(this.region);
    }
    /**
     * AWS secret key. (Cloudwatch)
     * 
     */
    @Export(name="secretAccessKey", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> secretAccessKey;

    /**
     * @return AWS secret key. (Cloudwatch)
     * 
     */
    public Output<Optional<String>> secretAccessKey() {
        return Codegen.optional(this.secretAccessKey);
    }
    /**
     * (optional) tags. E.g. env=prod,region=europe
     * 
     */
    @Export(name="tags", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> tags;

    /**
     * @return (optional) tags. E.g. env=prod,region=europe
     * 
     */
    public Output<Optional<String>> tags() {
        return Codegen.optional(this.tags);
    }
    /**
     * (optional) allowlist using regular expression
     * 
     */
    @Export(name="vhostAllowlist", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> vhostAllowlist;

    /**
     * @return (optional) allowlist using regular expression
     * 
     */
    public Output<Optional<String>> vhostAllowlist() {
        return Codegen.optional(this.vhostAllowlist);
    }
    /**
     * **Deprecated**
     * 
     * @deprecated
     * use vhost_allowlist instead
     * 
     */
    @Deprecated /* use vhost_allowlist instead */
    @Export(name="vhostWhitelist", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> vhostWhitelist;

    /**
     * @return **Deprecated**
     * 
     */
    public Output<Optional<String>> vhostWhitelist() {
        return Codegen.optional(this.vhostWhitelist);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public IntegrationMetric(java.lang.String name) {
        this(name, IntegrationMetricArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public IntegrationMetric(java.lang.String name, IntegrationMetricArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public IntegrationMetric(java.lang.String name, IntegrationMetricArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("cloudamqp:index/integrationMetric:IntegrationMetric", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private IntegrationMetric(java.lang.String name, Output<java.lang.String> id, @Nullable IntegrationMetricState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("cloudamqp:index/integrationMetric:IntegrationMetric", name, state, makeResourceOptions(options, id), false);
    }

    private static IntegrationMetricArgs makeArgs(IntegrationMetricArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? IntegrationMetricArgs.Empty : args;
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<java.lang.String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .additionalSecretOutputs(List.of(
                "apiKey",
                "credentials",
                "licenseKey",
                "privateKey",
                "privateKeyId",
                "secretAccessKey"
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
    public static IntegrationMetric get(java.lang.String name, Output<java.lang.String> id, @Nullable IntegrationMetricState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new IntegrationMetric(name, id, state, options);
    }
}
