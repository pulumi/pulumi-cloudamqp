// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * This resource allows you to create and manage third party log integrations for a CloudAMQP instance. Once configured, the logs produced will be forward to corresponding integration.
 *
 * Only available for dedicated subscription plans.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as cloudamqp from "@pulumi/cloudamqp";
 *
 * const cloudwatch = new cloudamqp.IntegrationLog("cloudwatch", {
 *     instanceId: cloudamqp_instance.instance.id,
 *     accessKeyId: _var.aws_access_key_id,
 *     secretAccessKey: _var.aws_secret_access_key,
 *     region: _var.aws_region,
 * });
 * const logentries = new cloudamqp.IntegrationLog("logentries", {
 *     instanceId: cloudamqp_instance.instance.id,
 *     token: _var.logentries_token,
 * });
 * const loggly = new cloudamqp.IntegrationLog("loggly", {
 *     instanceId: cloudamqp_instance.instance.id,
 *     token: _var.loggly_token,
 * });
 * const papertrail = new cloudamqp.IntegrationLog("papertrail", {
 *     instanceId: cloudamqp_instance.instance.id,
 *     url: _var.papertrail_url,
 * });
 * const splunk = new cloudamqp.IntegrationLog("splunk", {
 *     instanceId: cloudamqp_instance.instance.id,
 *     token: _var.splunk_token,
 *     hostPort: _var.splunk_host_port,
 * });
 * const datadog = new cloudamqp.IntegrationLog("datadog", {
 *     instanceId: cloudamqp_instance.instance.id,
 *     region: _var.datadog_region,
 *     apiKey: _var.datadog_api_key,
 *     tags: _var.datadog_tags,
 * });
 * const stackdriver = new cloudamqp.IntegrationLog("stackdriver", {
 *     instanceId: cloudamqp_instance.instance.id,
 *     projectId: _var.stackdriver_project_id,
 *     privateKey: _var.stackdriver_private_key,
 *     clientEmail: _var.stackdriver_client_email,
 * });
 * ```
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
 * | loggly     | Create a Loggly token at https://{your-company}.loggly.com/tokens |
 * | papertrail | Create a Papertrail endpoint https://papertrailapp.com/systems/setup |
 * | splunk     | Create a HTTP Event Collector token at https://.cloud.splunk.com/en-US/manager/search/http-eventcollector |
 * | datadog       | Create a Datadog API key at app.datadoghq.com |
 * | stackdriver   | Create a service account and add 'monitor metrics writer' role, then download credentials. |
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
 * | Splunk | splunk | token, hostPort |
 * | Data Dog | datadog | region, api_keys, tags |
 * | Stackdriver | stackdriver | project_id, private_key, clientEmail |
 *
 * ## Dependency
 *
 * This resource depends on CloudAMQP instance identifier, `cloudamqp_instance.instance.id`.
 *
 * ## Import
 *
 * `cloudamqp_integration_log`can be imported using the name argument of the resource together with CloudAMQP instance identifier. The name and identifier are CSV separated, see example below.
 *
 * ```sh
 *  $ pulumi import cloudamqp:index/integrationLog:IntegrationLog <resource_name> <name>,<instance_id>`
 * ```
 */
export class IntegrationLog extends pulumi.CustomResource {
    /**
     * Get an existing IntegrationLog resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: IntegrationLogState, opts?: pulumi.CustomResourceOptions): IntegrationLog {
        return new IntegrationLog(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'cloudamqp:index/integrationLog:IntegrationLog';

    /**
     * Returns true if the given object is an instance of IntegrationLog.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is IntegrationLog {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === IntegrationLog.__pulumiType;
    }

    /**
     * AWS access key identifier.
     */
    public readonly accessKeyId!: pulumi.Output<string | undefined>;
    /**
     * The API key.
     */
    public readonly apiKey!: pulumi.Output<string | undefined>;
    /**
     * The client email registered for the integration service.
     */
    public readonly clientEmail!: pulumi.Output<string | undefined>;
    /**
     * Destination to send the logs.
     */
    public readonly hostPort!: pulumi.Output<string | undefined>;
    /**
     * Instance identifier used to make proxy calls
     */
    public readonly instanceId!: pulumi.Output<number>;
    /**
     * The name of the third party log integration. See
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The private access key.
     */
    public readonly privateKey!: pulumi.Output<string | undefined>;
    /**
     * The project identifier.
     */
    public readonly projectId!: pulumi.Output<string | undefined>;
    /**
     * Region hosting the integration service.
     */
    public readonly region!: pulumi.Output<string | undefined>;
    /**
     * AWS secret access key.
     */
    public readonly secretAccessKey!: pulumi.Output<string | undefined>;
    /**
     * Tag the integration, e.g. env=prod, region=europe.
     */
    public readonly tags!: pulumi.Output<string | undefined>;
    /**
     * Token used for authentication.
     */
    public readonly token!: pulumi.Output<string | undefined>;
    /**
     * Endpoint to log integration.
     */
    public readonly url!: pulumi.Output<string | undefined>;

    /**
     * Create a IntegrationLog resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: IntegrationLogArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: IntegrationLogArgs | IntegrationLogState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as IntegrationLogState | undefined;
            inputs["accessKeyId"] = state ? state.accessKeyId : undefined;
            inputs["apiKey"] = state ? state.apiKey : undefined;
            inputs["clientEmail"] = state ? state.clientEmail : undefined;
            inputs["hostPort"] = state ? state.hostPort : undefined;
            inputs["instanceId"] = state ? state.instanceId : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["privateKey"] = state ? state.privateKey : undefined;
            inputs["projectId"] = state ? state.projectId : undefined;
            inputs["region"] = state ? state.region : undefined;
            inputs["secretAccessKey"] = state ? state.secretAccessKey : undefined;
            inputs["tags"] = state ? state.tags : undefined;
            inputs["token"] = state ? state.token : undefined;
            inputs["url"] = state ? state.url : undefined;
        } else {
            const args = argsOrState as IntegrationLogArgs | undefined;
            if ((!args || args.instanceId === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'instanceId'");
            }
            inputs["accessKeyId"] = args ? args.accessKeyId : undefined;
            inputs["apiKey"] = args ? args.apiKey : undefined;
            inputs["clientEmail"] = args ? args.clientEmail : undefined;
            inputs["hostPort"] = args ? args.hostPort : undefined;
            inputs["instanceId"] = args ? args.instanceId : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["privateKey"] = args ? args.privateKey : undefined;
            inputs["projectId"] = args ? args.projectId : undefined;
            inputs["region"] = args ? args.region : undefined;
            inputs["secretAccessKey"] = args ? args.secretAccessKey : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["token"] = args ? args.token : undefined;
            inputs["url"] = args ? args.url : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(IntegrationLog.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering IntegrationLog resources.
 */
export interface IntegrationLogState {
    /**
     * AWS access key identifier.
     */
    readonly accessKeyId?: pulumi.Input<string>;
    /**
     * The API key.
     */
    readonly apiKey?: pulumi.Input<string>;
    /**
     * The client email registered for the integration service.
     */
    readonly clientEmail?: pulumi.Input<string>;
    /**
     * Destination to send the logs.
     */
    readonly hostPort?: pulumi.Input<string>;
    /**
     * Instance identifier used to make proxy calls
     */
    readonly instanceId?: pulumi.Input<number>;
    /**
     * The name of the third party log integration. See
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The private access key.
     */
    readonly privateKey?: pulumi.Input<string>;
    /**
     * The project identifier.
     */
    readonly projectId?: pulumi.Input<string>;
    /**
     * Region hosting the integration service.
     */
    readonly region?: pulumi.Input<string>;
    /**
     * AWS secret access key.
     */
    readonly secretAccessKey?: pulumi.Input<string>;
    /**
     * Tag the integration, e.g. env=prod, region=europe.
     */
    readonly tags?: pulumi.Input<string>;
    /**
     * Token used for authentication.
     */
    readonly token?: pulumi.Input<string>;
    /**
     * Endpoint to log integration.
     */
    readonly url?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a IntegrationLog resource.
 */
export interface IntegrationLogArgs {
    /**
     * AWS access key identifier.
     */
    readonly accessKeyId?: pulumi.Input<string>;
    /**
     * The API key.
     */
    readonly apiKey?: pulumi.Input<string>;
    /**
     * The client email registered for the integration service.
     */
    readonly clientEmail?: pulumi.Input<string>;
    /**
     * Destination to send the logs.
     */
    readonly hostPort?: pulumi.Input<string>;
    /**
     * Instance identifier used to make proxy calls
     */
    readonly instanceId: pulumi.Input<number>;
    /**
     * The name of the third party log integration. See
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The private access key.
     */
    readonly privateKey?: pulumi.Input<string>;
    /**
     * The project identifier.
     */
    readonly projectId?: pulumi.Input<string>;
    /**
     * Region hosting the integration service.
     */
    readonly region?: pulumi.Input<string>;
    /**
     * AWS secret access key.
     */
    readonly secretAccessKey?: pulumi.Input<string>;
    /**
     * Tag the integration, e.g. env=prod, region=europe.
     */
    readonly tags?: pulumi.Input<string>;
    /**
     * Token used for authentication.
     */
    readonly token?: pulumi.Input<string>;
    /**
     * Endpoint to log integration.
     */
    readonly url?: pulumi.Input<string>;
}
