// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.cloudamqp;

import com.pulumi.cloudamqp.IntegrationLogArgs;
import com.pulumi.cloudamqp.Utilities;
import com.pulumi.cloudamqp.inputs.IntegrationLogState;
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
 * This resource allows you to create and manage third party log integrations for a CloudAMQP instance.
 * Once configured, the logs produced will be forward to corresponding integration.
 * 
 * Only available for dedicated subscription plans.
 * 
 * ## Example Usage
 * 
 * &lt;details&gt;
 *   &lt;summary&gt;
 *     &lt;b&gt;
 *       &lt;i&gt;Azure monitor log integration&lt;/i&gt;
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
 * import com.pulumi.cloudamqp.IntegrationLog;
 * import com.pulumi.cloudamqp.IntegrationLogArgs;
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
 *         var azureMonitor = new IntegrationLog("azureMonitor", IntegrationLogArgs.builder()
 *             .instanceId(instance.id())
 *             .name("azure_monitor")
 *             .tenantId(azmTentantId)
 *             .applicationId(azmApplicationId)
 *             .applicationSecret(azmApplicationSecret)
 *             .dceUri(azmDceUri)
 *             .table(azmTable)
 *             .dcrId(azmDcrId)
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
 *       &lt;i&gt;Cloudwatch log integration&lt;/i&gt;
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
 * import com.pulumi.cloudamqp.IntegrationLog;
 * import com.pulumi.cloudamqp.IntegrationLogArgs;
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
 *         var cloudwatch = new IntegrationLog("cloudwatch", IntegrationLogArgs.builder()
 *             .instanceId(instance.id())
 *             .name("cloudwatchlog")
 *             .accessKeyId(awsAccessKeyId)
 *             .secretAccessKey(awsSecretAccessKey)
 *             .region(awsRegion)
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
 *       &lt;i&gt;Coralogix log integration&lt;/i&gt;
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
 * import com.pulumi.cloudamqp.IntegrationLog;
 * import com.pulumi.cloudamqp.IntegrationLogArgs;
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
 *         var coralogix = new IntegrationLog("coralogix", IntegrationLogArgs.builder()
 *             .instanceId(instance.id())
 *             .name("coralogix")
 *             .privateKey(coralogixSendDataKey)
 *             .endpoint(coralogixEndpoint)
 *             .application(coralogixApplication)
 *             .subsystem(instance.host())
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
 *       &lt;i&gt;Datadog log integration&lt;/i&gt;
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
 * import com.pulumi.cloudamqp.IntegrationLog;
 * import com.pulumi.cloudamqp.IntegrationLogArgs;
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
 *         var datadog = new IntegrationLog("datadog", IntegrationLogArgs.builder()
 *             .instanceId(instance.id())
 *             .name("datadog")
 *             .region(datadogRegion)
 *             .apiKey(datadogApiKey)
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
 *       &lt;i&gt;Logentries log integration&lt;/i&gt;
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
 * import com.pulumi.cloudamqp.IntegrationLog;
 * import com.pulumi.cloudamqp.IntegrationLogArgs;
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
 *         var logentries = new IntegrationLog("logentries", IntegrationLogArgs.builder()
 *             .instanceId(instance.id())
 *             .name("logentries")
 *             .token(logentriesToken)
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
 *       &lt;i&gt;Loggly log integration&lt;/i&gt;
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
 * import com.pulumi.cloudamqp.IntegrationLog;
 * import com.pulumi.cloudamqp.IntegrationLogArgs;
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
 *         var loggly = new IntegrationLog("loggly", IntegrationLogArgs.builder()
 *             .instanceId(instance.id())
 *             .name("loggly")
 *             .token(logglyToken)
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
 *       &lt;i&gt;Papertrail log integration&lt;/i&gt;
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
 * import com.pulumi.cloudamqp.IntegrationLog;
 * import com.pulumi.cloudamqp.IntegrationLogArgs;
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
 *         var papertrail = new IntegrationLog("papertrail", IntegrationLogArgs.builder()
 *             .instanceId(instance.id())
 *             .name("papertrail")
 *             .url(papertrailUrl)
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
 *       &lt;i&gt;Scalyr log integration&lt;/i&gt;
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
 * import com.pulumi.cloudamqp.IntegrationLog;
 * import com.pulumi.cloudamqp.IntegrationLogArgs;
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
 *         var scalyr = new IntegrationLog("scalyr", IntegrationLogArgs.builder()
 *             .instanceId(instance.id())
 *             .name("scalyr")
 *             .token(scalyrToken)
 *             .host(scalyrHost)
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * &lt;details&gt;
 *   &lt;summary&gt;
 *     &lt;b&gt;
 *       &lt;i&gt;Splunk log integration&lt;/i&gt;
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
 * import com.pulumi.cloudamqp.IntegrationLog;
 * import com.pulumi.cloudamqp.IntegrationLogArgs;
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
 *         var splunk = new IntegrationLog("splunk", IntegrationLogArgs.builder()
 *             .instanceId(instance.id())
 *             .name("splunk")
 *             .token(splunkToken)
 *             .hostPort(splunkHostPort)
 *             .sourceType("generic_single_line")
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
 * &lt;/details&gt;
 * 
 * &lt;details&gt;
 *   &lt;summary&gt;
 *     &lt;b&gt;
 *       &lt;i&gt;Stackdriver log integration (v1.20.2 or older versions)&lt;/i&gt;
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
 * import com.pulumi.cloudamqp.IntegrationLog;
 * import com.pulumi.cloudamqp.IntegrationLogArgs;
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
 *         var stackdriver = new IntegrationLog("stackdriver", IntegrationLogArgs.builder()
 *             .instanceId(instance.id())
 *             .name("stackdriver")
 *             .projectId(stackdriverProjectId)
 *             .privateKey(stackdriverPrivateKey)
 *             .clientEmail(stackdriverClientEmail)
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
 * `cloudamqp_integration_log`can be imported using the resource identifier together with CloudAMQP instance identifier. The name and identifier are CSV separated, see example below.
 * 
 * ```sh
 * $ pulumi import cloudamqp:index/integrationLog:IntegrationLog &lt;resource_name&gt; &lt;id&gt;,&lt;instance_id&gt;`
 * ```
 * 
 */
@ResourceType(type="cloudamqp:index/integrationLog:IntegrationLog")
public class IntegrationLog extends com.pulumi.resources.CustomResource {
    /**
     * AWS access key identifier.
     * 
     */
    @Export(name="accessKeyId", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> accessKeyId;

    /**
     * @return AWS access key identifier.
     * 
     */
    public Output<Optional<String>> accessKeyId() {
        return Codegen.optional(this.accessKeyId);
    }
    /**
     * The API key.
     * 
     */
    @Export(name="apiKey", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> apiKey;

    /**
     * @return The API key.
     * 
     */
    public Output<Optional<String>> apiKey() {
        return Codegen.optional(this.apiKey);
    }
    /**
     * The application name for Coralogix.
     * 
     */
    @Export(name="application", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> application;

    /**
     * @return The application name for Coralogix.
     * 
     */
    public Output<Optional<String>> application() {
        return Codegen.optional(this.application);
    }
    /**
     * The application identifier for Azure monitor.
     * 
     */
    @Export(name="applicationId", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> applicationId;

    /**
     * @return The application identifier for Azure monitor.
     * 
     */
    public Output<Optional<String>> applicationId() {
        return Codegen.optional(this.applicationId);
    }
    /**
     * The application secret for Azure monitor.
     * 
     */
    @Export(name="applicationSecret", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> applicationSecret;

    /**
     * @return The application secret for Azure monitor.
     * 
     */
    public Output<Optional<String>> applicationSecret() {
        return Codegen.optional(this.applicationSecret);
    }
    /**
     * The client email registered for the integration service.
     * 
     */
    @Export(name="clientEmail", refs={String.class}, tree="[0]")
    private Output<String> clientEmail;

    /**
     * @return The client email registered for the integration service.
     * 
     */
    public Output<String> clientEmail() {
        return this.clientEmail;
    }
    /**
     * Google Service Account private key credentials.
     * 
     */
    @Export(name="credentials", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> credentials;

    /**
     * @return Google Service Account private key credentials.
     * 
     */
    public Output<Optional<String>> credentials() {
        return Codegen.optional(this.credentials);
    }
    /**
     * The data collection endpoint for Azure monitor.
     * 
     */
    @Export(name="dceUri", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> dceUri;

    /**
     * @return The data collection endpoint for Azure monitor.
     * 
     */
    public Output<Optional<String>> dceUri() {
        return Codegen.optional(this.dceUri);
    }
    /**
     * ID of data collection rule that your DCE is linked to for Azure Monitor.
     * 
     * This is the full list of all arguments. Only a subset of arguments are used based on which type of integration used. See Integration Type reference table below for more information.
     * 
     */
    @Export(name="dcrId", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> dcrId;

    /**
     * @return ID of data collection rule that your DCE is linked to for Azure Monitor.
     * 
     * This is the full list of all arguments. Only a subset of arguments are used based on which type of integration used. See Integration Type reference table below for more information.
     * 
     */
    public Output<Optional<String>> dcrId() {
        return Codegen.optional(this.dcrId);
    }
    /**
     * The syslog destination to send the logs to for Coralogix.
     * 
     */
    @Export(name="endpoint", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> endpoint;

    /**
     * @return The syslog destination to send the logs to for Coralogix.
     * 
     */
    public Output<Optional<String>> endpoint() {
        return Codegen.optional(this.endpoint);
    }
    /**
     * The host for Scalyr integration. (app.scalyr.com, app.eu.scalyr.com)
     * 
     */
    @Export(name="host", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> host;

    /**
     * @return The host for Scalyr integration. (app.scalyr.com, app.eu.scalyr.com)
     * 
     */
    public Output<Optional<String>> host() {
        return Codegen.optional(this.host);
    }
    /**
     * Destination to send the logs.
     * 
     */
    @Export(name="hostPort", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> hostPort;

    /**
     * @return Destination to send the logs.
     * 
     */
    public Output<Optional<String>> hostPort() {
        return Codegen.optional(this.hostPort);
    }
    /**
     * Instance identifier used to make proxy calls
     * 
     */
    @Export(name="instanceId", refs={Integer.class}, tree="[0]")
    private Output<Integer> instanceId;

    /**
     * @return Instance identifier used to make proxy calls
     * 
     */
    public Output<Integer> instanceId() {
        return this.instanceId;
    }
    /**
     * The name of the third party log integration. See
     * Integration type reference
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return The name of the third party log integration. See
     * Integration type reference
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * The private access key.
     * 
     */
    @Export(name="privateKey", refs={String.class}, tree="[0]")
    private Output<String> privateKey;

    /**
     * @return The private access key.
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
     * The project identifier.
     * 
     */
    @Export(name="projectId", refs={String.class}, tree="[0]")
    private Output<String> projectId;

    /**
     * @return The project identifier.
     * 
     */
    public Output<String> projectId() {
        return this.projectId;
    }
    /**
     * Region hosting the integration service.
     * 
     */
    @Export(name="region", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> region;

    /**
     * @return Region hosting the integration service.
     * 
     */
    public Output<Optional<String>> region() {
        return Codegen.optional(this.region);
    }
    /**
     * AWS secret access key.
     * 
     */
    @Export(name="secretAccessKey", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> secretAccessKey;

    /**
     * @return AWS secret access key.
     * 
     */
    public Output<Optional<String>> secretAccessKey() {
        return Codegen.optional(this.secretAccessKey);
    }
    /**
     * Assign source type to the data exported, eg. generic_single_line. (Splunk)
     * 
     */
    @Export(name="sourcetype", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> sourcetype;

    /**
     * @return Assign source type to the data exported, eg. generic_single_line. (Splunk)
     * 
     */
    public Output<Optional<String>> sourcetype() {
        return Codegen.optional(this.sourcetype);
    }
    /**
     * The subsystem name for Coralogix.
     * 
     */
    @Export(name="subsystem", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> subsystem;

    /**
     * @return The subsystem name for Coralogix.
     * 
     */
    public Output<Optional<String>> subsystem() {
        return Codegen.optional(this.subsystem);
    }
    /**
     * The table name for Azure monitor.
     * 
     */
    @Export(name="table", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> table;

    /**
     * @return The table name for Azure monitor.
     * 
     */
    public Output<Optional<String>> table() {
        return Codegen.optional(this.table);
    }
    /**
     * Tags. e.g. `env=prod,region=europe`.
     * 
     * ***Note: If tags are used with Datadog. The value part (prod, europe, ...) must start with a letter, read more about tags format in the [Datadog documentation](https://docs.datadoghq.com/getting_started/tagging/#define-tags)***
     * 
     */
    @Export(name="tags", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> tags;

    /**
     * @return Tags. e.g. `env=prod,region=europe`.
     * 
     * ***Note: If tags are used with Datadog. The value part (prod, europe, ...) must start with a letter, read more about tags format in the [Datadog documentation](https://docs.datadoghq.com/getting_started/tagging/#define-tags)***
     * 
     */
    public Output<Optional<String>> tags() {
        return Codegen.optional(this.tags);
    }
    /**
     * The tenant identifier for Azure monitor.
     * 
     */
    @Export(name="tenantId", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> tenantId;

    /**
     * @return The tenant identifier for Azure monitor.
     * 
     */
    public Output<Optional<String>> tenantId() {
        return Codegen.optional(this.tenantId);
    }
    /**
     * Token used for authentication.
     * 
     */
    @Export(name="token", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> token;

    /**
     * @return Token used for authentication.
     * 
     */
    public Output<Optional<String>> token() {
        return Codegen.optional(this.token);
    }
    /**
     * Endpoint to log integration.
     * 
     */
    @Export(name="url", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> url;

    /**
     * @return Endpoint to log integration.
     * 
     */
    public Output<Optional<String>> url() {
        return Codegen.optional(this.url);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public IntegrationLog(String name) {
        this(name, IntegrationLogArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public IntegrationLog(String name, IntegrationLogArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public IntegrationLog(String name, IntegrationLogArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("cloudamqp:index/integrationLog:IntegrationLog", name, args == null ? IntegrationLogArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private IntegrationLog(String name, Output<String> id, @Nullable IntegrationLogState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("cloudamqp:index/integrationLog:IntegrationLog", name, state, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .additionalSecretOutputs(List.of(
                "accessKeyId",
                "apiKey",
                "applicationSecret",
                "credentials",
                "privateKey",
                "privateKeyId",
                "secretAccessKey",
                "token"
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
    public static IntegrationLog get(String name, Output<String> id, @Nullable IntegrationLogState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new IntegrationLog(name, id, state, options);
    }
}
