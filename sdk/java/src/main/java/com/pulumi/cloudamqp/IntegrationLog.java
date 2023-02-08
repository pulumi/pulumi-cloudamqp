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
 * This resource allows you to create and manage third party log integrations for a CloudAMQP instance. Once configured, the logs produced will be forward to corresponding integration.
 * 
 * Only available for dedicated subscription plans.
 * 
 * ## Argument Reference (cloudwatchlog)
 * 
 * Cloudwatch argument reference and example. Create an IAM user with programmatic access and the following permissions:
 * 
 * * CreateLogGroup
 * * CreateLogStream
 * * DescribeLogGroups
 * * DescribeLogStreams
 * * PutLogEvents
 * 
 * ## Integration service reference
 * 
 * Valid names for third party log integration.
 * 
 * | Name       | Description |
 * |------------|---------------------------------------------------------------|
 * | cloudwatchlog | Create a IAM with programmatic access. |
 * | logentries | Create a Logentries token at https://logentries.com/app#/add-log/manual  |
 * | loggly     | Create a Loggly token at https://your-company}.loggly.com/tokens |
 * | papertrail | Create a Papertrail endpoint https://papertrailapp.com/systems/setup |
 * | splunk     | Create a HTTP Event Collector token at `https://&lt;your-splunk&gt;.cloud.splunk.com/en-US/manager/search/http-eventcollector` |
 * | datadog       | Create a Datadog API key at app.datadoghq.com |
 * | stackdriver   | Create a service account and add &#39;monitor metrics writer&#39; role from your Google Cloud Account |
 * | scalyr        | Create a Log write token at https://app.scalyr.com/keys |
 * 
 * ## Integration Type reference
 * 
 * Valid arguments for third party log integrations.
 * 
 * Required arguments for all integrations: name
 * 
 * | Name | Type | Required arguments |
 * | ---- | ---- | ---- |
 * | CloudWatch | cloudwatchlog | access_key_id, secret_access_key, region |
 * | Log Entries | logentries | token |
 * | Loggly | loggly | token |
 * | Papertrail | papertrail | url |
 * | Splunk | splunk | token, host_port, sourcetype |
 * | Data Dog | datadog | region, api_keys, tags |
 * | Stackdriver | stackdriver | credentials |
 * | Scalyr | scalyr | token, host |
 * 
 * ***Note:*** Stackdriver (v1.20.2 or earlier versions) required arguments  : project_id, private_key, client_email
 * 
 * ## Dependency
 * 
 * This resource depends on CloudAMQP instance identifier, `cloudamqp_instance.instance.id`.
 * 
 * ## Import
 * 
 * `cloudamqp_integration_log`can be imported using the resource identifier together with CloudAMQP instance identifier. The name and identifier are CSV separated, see example below.
 * 
 * ```sh
 *  $ pulumi import cloudamqp:index/integrationLog:IntegrationLog &lt;resource_name&gt; &lt;id&gt;,&lt;instance_id&gt;`
 * ```
 * 
 */
@ResourceType(type="cloudamqp:index/integrationLog:IntegrationLog")
public class IntegrationLog extends com.pulumi.resources.CustomResource {
    /**
     * AWS access key identifier.
     * 
     */
    @Export(name="accessKeyId", type=String.class, parameters={})
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
    @Export(name="apiKey", type=String.class, parameters={})
    private Output</* @Nullable */ String> apiKey;

    /**
     * @return The API key.
     * 
     */
    public Output<Optional<String>> apiKey() {
        return Codegen.optional(this.apiKey);
    }
    /**
     * The client email registered for the integration service.
     * 
     */
    @Export(name="clientEmail", type=String.class, parameters={})
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
    @Export(name="credentials", type=String.class, parameters={})
    private Output</* @Nullable */ String> credentials;

    /**
     * @return Google Service Account private key credentials.
     * 
     */
    public Output<Optional<String>> credentials() {
        return Codegen.optional(this.credentials);
    }
    /**
     * The host for Scalyr integration. (app.scalyr.com, app.eu.scalyr.com)
     * 
     */
    @Export(name="host", type=String.class, parameters={})
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
    @Export(name="hostPort", type=String.class, parameters={})
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
    @Export(name="instanceId", type=Integer.class, parameters={})
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
     * 
     */
    @Export(name="name", type=String.class, parameters={})
    private Output<String> name;

    /**
     * @return The name of the third party log integration. See
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * The private access key.
     * 
     */
    @Export(name="privateKey", type=String.class, parameters={})
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
    @Export(name="privateKeyId", type=String.class, parameters={})
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
    @Export(name="projectId", type=String.class, parameters={})
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
    @Export(name="region", type=String.class, parameters={})
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
    @Export(name="secretAccessKey", type=String.class, parameters={})
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
    @Export(name="sourcetype", type=String.class, parameters={})
    private Output</* @Nullable */ String> sourcetype;

    /**
     * @return Assign source type to the data exported, eg. generic_single_line. (Splunk)
     * 
     */
    public Output<Optional<String>> sourcetype() {
        return Codegen.optional(this.sourcetype);
    }
    /**
     * Tag the integration, e.g. env=prod, region=europe.
     * 
     */
    @Export(name="tags", type=String.class, parameters={})
    private Output</* @Nullable */ String> tags;

    /**
     * @return Tag the integration, e.g. env=prod, region=europe.
     * 
     */
    public Output<Optional<String>> tags() {
        return Codegen.optional(this.tags);
    }
    /**
     * Token used for authentication.
     * 
     */
    @Export(name="token", type=String.class, parameters={})
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
    @Export(name="url", type=String.class, parameters={})
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