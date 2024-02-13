// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * This resource allows you to create and manage third party log integrations for a CloudAMQP instance.
 * Once configured, the logs produced will be forward to corresponding integration.
 *
 * Only available for dedicated subscription plans.
 *
 * ## Integration Type reference
 *
 * Valid arguments for third party log integrations. See more information at [docs.cloudamqp.com](https://docs.cloudamqp.com/cloudamqp_api.html#add-log-integration)
 *
 * Required arguments for all integrations: name
 *
 * | Integration | name | Required arguments |
 * | ---- | ---- | ---- |
 * | Azure monitor | azureMonitor | tenant_id, application_id, application_secret, dce_uri, table, dcrId |
 * | CloudWatch | cloudwatchlog | access_key_id, secret_access_key, region |
 * | Coralogix | coralogix | private_key, endpoint, application, subsystem |
 * | Data Dog | datadog | region, api_keys, tags |
 * | Log Entries | logentries | token |
 * | Loggly | loggly | token |
 * | Papertrail | papertrail | url |
 * | Scalyr | scalyr | token, host |
 * | Splunk | splunk | token, host_port, sourcetype |
 * | Stackdriver | stackdriver | credentials |
 *
 * ***Note:*** Stackdriver (v1.20.2 or earlier versions) required arguments  : project_id, private_key, clientEmail
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
 * $ pulumi import cloudamqp:index/integrationLog:IntegrationLog <resource_name> <id>,<instance_id>`
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
     * The application name for Coralogix.
     */
    public readonly application!: pulumi.Output<string | undefined>;
    /**
     * The application identifier for Azure monitor.
     */
    public readonly applicationId!: pulumi.Output<string | undefined>;
    /**
     * The application secret for Azure monitor.
     */
    public readonly applicationSecret!: pulumi.Output<string | undefined>;
    /**
     * The client email registered for the integration service.
     */
    public readonly clientEmail!: pulumi.Output<string>;
    /**
     * Google Service Account private key credentials.
     */
    public readonly credentials!: pulumi.Output<string | undefined>;
    /**
     * The data collection endpoint for Azure monitor.
     */
    public readonly dceUri!: pulumi.Output<string | undefined>;
    /**
     * ID of data collection rule that your DCE is linked to for Azure Monitor.
     *
     * This is the full list of all arguments. Only a subset of arguments are used based on which type of integration used. See Integration Type reference table below for more information.
     */
    public readonly dcrId!: pulumi.Output<string | undefined>;
    /**
     * The syslog destination to send the logs to for Coralogix.
     */
    public readonly endpoint!: pulumi.Output<string | undefined>;
    /**
     * The host for Scalyr integration. (app.scalyr.com, app.eu.scalyr.com)
     */
    public readonly host!: pulumi.Output<string | undefined>;
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
     * Integration type reference
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The private access key.
     */
    public readonly privateKey!: pulumi.Output<string>;
    /**
     * Private key identifier. (Stackdriver)
     */
    public readonly privateKeyId!: pulumi.Output<string>;
    /**
     * The project identifier.
     */
    public readonly projectId!: pulumi.Output<string>;
    /**
     * Region hosting the integration service.
     */
    public readonly region!: pulumi.Output<string | undefined>;
    /**
     * AWS secret access key.
     */
    public readonly secretAccessKey!: pulumi.Output<string | undefined>;
    /**
     * Assign source type to the data exported, eg. generic_single_line. (Splunk)
     */
    public readonly sourcetype!: pulumi.Output<string | undefined>;
    /**
     * The subsystem name for Coralogix.
     */
    public readonly subsystem!: pulumi.Output<string | undefined>;
    /**
     * The table name for Azure monitor.
     */
    public readonly table!: pulumi.Output<string | undefined>;
    /**
     * Tag the integration, e.g. env=prod,region=europe.
     */
    public readonly tags!: pulumi.Output<string | undefined>;
    /**
     * The tenant identifier for Azure monitor.
     */
    public readonly tenantId!: pulumi.Output<string | undefined>;
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
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as IntegrationLogState | undefined;
            resourceInputs["accessKeyId"] = state ? state.accessKeyId : undefined;
            resourceInputs["apiKey"] = state ? state.apiKey : undefined;
            resourceInputs["application"] = state ? state.application : undefined;
            resourceInputs["applicationId"] = state ? state.applicationId : undefined;
            resourceInputs["applicationSecret"] = state ? state.applicationSecret : undefined;
            resourceInputs["clientEmail"] = state ? state.clientEmail : undefined;
            resourceInputs["credentials"] = state ? state.credentials : undefined;
            resourceInputs["dceUri"] = state ? state.dceUri : undefined;
            resourceInputs["dcrId"] = state ? state.dcrId : undefined;
            resourceInputs["endpoint"] = state ? state.endpoint : undefined;
            resourceInputs["host"] = state ? state.host : undefined;
            resourceInputs["hostPort"] = state ? state.hostPort : undefined;
            resourceInputs["instanceId"] = state ? state.instanceId : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["privateKey"] = state ? state.privateKey : undefined;
            resourceInputs["privateKeyId"] = state ? state.privateKeyId : undefined;
            resourceInputs["projectId"] = state ? state.projectId : undefined;
            resourceInputs["region"] = state ? state.region : undefined;
            resourceInputs["secretAccessKey"] = state ? state.secretAccessKey : undefined;
            resourceInputs["sourcetype"] = state ? state.sourcetype : undefined;
            resourceInputs["subsystem"] = state ? state.subsystem : undefined;
            resourceInputs["table"] = state ? state.table : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["tenantId"] = state ? state.tenantId : undefined;
            resourceInputs["token"] = state ? state.token : undefined;
            resourceInputs["url"] = state ? state.url : undefined;
        } else {
            const args = argsOrState as IntegrationLogArgs | undefined;
            if ((!args || args.instanceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'instanceId'");
            }
            resourceInputs["accessKeyId"] = args?.accessKeyId ? pulumi.secret(args.accessKeyId) : undefined;
            resourceInputs["apiKey"] = args?.apiKey ? pulumi.secret(args.apiKey) : undefined;
            resourceInputs["application"] = args ? args.application : undefined;
            resourceInputs["applicationId"] = args ? args.applicationId : undefined;
            resourceInputs["applicationSecret"] = args?.applicationSecret ? pulumi.secret(args.applicationSecret) : undefined;
            resourceInputs["clientEmail"] = args ? args.clientEmail : undefined;
            resourceInputs["credentials"] = args?.credentials ? pulumi.secret(args.credentials) : undefined;
            resourceInputs["dceUri"] = args ? args.dceUri : undefined;
            resourceInputs["dcrId"] = args ? args.dcrId : undefined;
            resourceInputs["endpoint"] = args ? args.endpoint : undefined;
            resourceInputs["host"] = args ? args.host : undefined;
            resourceInputs["hostPort"] = args ? args.hostPort : undefined;
            resourceInputs["instanceId"] = args ? args.instanceId : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["privateKey"] = args?.privateKey ? pulumi.secret(args.privateKey) : undefined;
            resourceInputs["privateKeyId"] = args?.privateKeyId ? pulumi.secret(args.privateKeyId) : undefined;
            resourceInputs["projectId"] = args ? args.projectId : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
            resourceInputs["secretAccessKey"] = args?.secretAccessKey ? pulumi.secret(args.secretAccessKey) : undefined;
            resourceInputs["sourcetype"] = args ? args.sourcetype : undefined;
            resourceInputs["subsystem"] = args ? args.subsystem : undefined;
            resourceInputs["table"] = args ? args.table : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["tenantId"] = args ? args.tenantId : undefined;
            resourceInputs["token"] = args?.token ? pulumi.secret(args.token) : undefined;
            resourceInputs["url"] = args ? args.url : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["accessKeyId", "apiKey", "applicationSecret", "credentials", "privateKey", "privateKeyId", "secretAccessKey", "token"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(IntegrationLog.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering IntegrationLog resources.
 */
export interface IntegrationLogState {
    /**
     * AWS access key identifier.
     */
    accessKeyId?: pulumi.Input<string>;
    /**
     * The API key.
     */
    apiKey?: pulumi.Input<string>;
    /**
     * The application name for Coralogix.
     */
    application?: pulumi.Input<string>;
    /**
     * The application identifier for Azure monitor.
     */
    applicationId?: pulumi.Input<string>;
    /**
     * The application secret for Azure monitor.
     */
    applicationSecret?: pulumi.Input<string>;
    /**
     * The client email registered for the integration service.
     */
    clientEmail?: pulumi.Input<string>;
    /**
     * Google Service Account private key credentials.
     */
    credentials?: pulumi.Input<string>;
    /**
     * The data collection endpoint for Azure monitor.
     */
    dceUri?: pulumi.Input<string>;
    /**
     * ID of data collection rule that your DCE is linked to for Azure Monitor.
     *
     * This is the full list of all arguments. Only a subset of arguments are used based on which type of integration used. See Integration Type reference table below for more information.
     */
    dcrId?: pulumi.Input<string>;
    /**
     * The syslog destination to send the logs to for Coralogix.
     */
    endpoint?: pulumi.Input<string>;
    /**
     * The host for Scalyr integration. (app.scalyr.com, app.eu.scalyr.com)
     */
    host?: pulumi.Input<string>;
    /**
     * Destination to send the logs.
     */
    hostPort?: pulumi.Input<string>;
    /**
     * Instance identifier used to make proxy calls
     */
    instanceId?: pulumi.Input<number>;
    /**
     * The name of the third party log integration. See
     * Integration type reference
     */
    name?: pulumi.Input<string>;
    /**
     * The private access key.
     */
    privateKey?: pulumi.Input<string>;
    /**
     * Private key identifier. (Stackdriver)
     */
    privateKeyId?: pulumi.Input<string>;
    /**
     * The project identifier.
     */
    projectId?: pulumi.Input<string>;
    /**
     * Region hosting the integration service.
     */
    region?: pulumi.Input<string>;
    /**
     * AWS secret access key.
     */
    secretAccessKey?: pulumi.Input<string>;
    /**
     * Assign source type to the data exported, eg. generic_single_line. (Splunk)
     */
    sourcetype?: pulumi.Input<string>;
    /**
     * The subsystem name for Coralogix.
     */
    subsystem?: pulumi.Input<string>;
    /**
     * The table name for Azure monitor.
     */
    table?: pulumi.Input<string>;
    /**
     * Tag the integration, e.g. env=prod,region=europe.
     */
    tags?: pulumi.Input<string>;
    /**
     * The tenant identifier for Azure monitor.
     */
    tenantId?: pulumi.Input<string>;
    /**
     * Token used for authentication.
     */
    token?: pulumi.Input<string>;
    /**
     * Endpoint to log integration.
     */
    url?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a IntegrationLog resource.
 */
export interface IntegrationLogArgs {
    /**
     * AWS access key identifier.
     */
    accessKeyId?: pulumi.Input<string>;
    /**
     * The API key.
     */
    apiKey?: pulumi.Input<string>;
    /**
     * The application name for Coralogix.
     */
    application?: pulumi.Input<string>;
    /**
     * The application identifier for Azure monitor.
     */
    applicationId?: pulumi.Input<string>;
    /**
     * The application secret for Azure monitor.
     */
    applicationSecret?: pulumi.Input<string>;
    /**
     * The client email registered for the integration service.
     */
    clientEmail?: pulumi.Input<string>;
    /**
     * Google Service Account private key credentials.
     */
    credentials?: pulumi.Input<string>;
    /**
     * The data collection endpoint for Azure monitor.
     */
    dceUri?: pulumi.Input<string>;
    /**
     * ID of data collection rule that your DCE is linked to for Azure Monitor.
     *
     * This is the full list of all arguments. Only a subset of arguments are used based on which type of integration used. See Integration Type reference table below for more information.
     */
    dcrId?: pulumi.Input<string>;
    /**
     * The syslog destination to send the logs to for Coralogix.
     */
    endpoint?: pulumi.Input<string>;
    /**
     * The host for Scalyr integration. (app.scalyr.com, app.eu.scalyr.com)
     */
    host?: pulumi.Input<string>;
    /**
     * Destination to send the logs.
     */
    hostPort?: pulumi.Input<string>;
    /**
     * Instance identifier used to make proxy calls
     */
    instanceId: pulumi.Input<number>;
    /**
     * The name of the third party log integration. See
     * Integration type reference
     */
    name?: pulumi.Input<string>;
    /**
     * The private access key.
     */
    privateKey?: pulumi.Input<string>;
    /**
     * Private key identifier. (Stackdriver)
     */
    privateKeyId?: pulumi.Input<string>;
    /**
     * The project identifier.
     */
    projectId?: pulumi.Input<string>;
    /**
     * Region hosting the integration service.
     */
    region?: pulumi.Input<string>;
    /**
     * AWS secret access key.
     */
    secretAccessKey?: pulumi.Input<string>;
    /**
     * Assign source type to the data exported, eg. generic_single_line. (Splunk)
     */
    sourcetype?: pulumi.Input<string>;
    /**
     * The subsystem name for Coralogix.
     */
    subsystem?: pulumi.Input<string>;
    /**
     * The table name for Azure monitor.
     */
    table?: pulumi.Input<string>;
    /**
     * Tag the integration, e.g. env=prod,region=europe.
     */
    tags?: pulumi.Input<string>;
    /**
     * The tenant identifier for Azure monitor.
     */
    tenantId?: pulumi.Input<string>;
    /**
     * Token used for authentication.
     */
    token?: pulumi.Input<string>;
    /**
     * Endpoint to log integration.
     */
    url?: pulumi.Input<string>;
}
