// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

<<<<<<< HEAD
/**
 * This resource allows you to create and manage, forwarding metrics to third party integrations for a CloudAMQP instance. Once configured, the metrics produced will be forward to corresponding integration.
 *
 * Only available for dedicated subscription plans.
 *
 * ## Argument references
 *
 * The following arguments are supported:
 *
 * * `name`              - (Required) The name of the third party log integration. See `Integration service reference`
 * * `region`            - (Optional) Region hosting the integration service.
 * * `accessKeyId`     - (Optional) AWS access key identifier.
 * * `secretAccessKey` - (Optional) AWS secret access key.
 * * `apiKey`           - (Optional) The API key for the integration service.
 * * `email`             - (Optional) The email address registred for the integration service.
 * * `projectId`        - (Optional) The project identifier.
 * * `privateKey`       - (Optional) The private access key.
 * * `clientEmail`      - (Optional) The client email registered for the integration service.
 * * `tags`              - (Optional) Tags. e.g. env=prod, region=europe.
 * * `queueWhitelist`   - (Optional) Whitelist queues using regular expression. Leave empty to include all queues.
 * * `vhostWhitelist`   - (Optional) Whitelist vhost using regular expression. Leave empty to include all vhosts.
 *
 * This is the full list of all arguments. Only a subset of arguments are used based on which type of integration used. See Integration type reference below for more information.
 *
 * ## Integration service references
 *
 * Valid names for third party log integration.
 *
 * | Name          | Description |
 * |---------------|---------------------------------------------------------------|
 * | cloudwatch    | Create an IAM with programmatic access. |
 * | cloudwatchV2 | Create an IAM with programmatic access. |
 * | datadog       | Create a Datadog API key at app.datadoghq.com |
 * | datadogV2    | Create a Datadog API key at app.datadoghq.com
 * | librato       | Create a new API token (with record only permissions) here: https://metrics.librato.com/tokens |
 * | newrelic      | Deprecated! |
 * | newrelicV2   | Find or register an Insert API key for your account: Go to insights.newrelic.com > Manage data > API keys. |
 * | stackdriver   | Create a service account and add 'monitor metrics writer' role, then download credentials. |
 *
 * ## Integration type reference
 *
 * Valid arguments for third party log integrations.
 *
 * Required arguments for all integrations: *name*<br>
 * Optional arguments for all integrations: *tags*, *queue_whitelist*, *vhost_whitelist*
 *
 * | Name | Type | Required arguments |
 * | ---- | ---- | ---- |
 * | Cloudwatch             | cloudwatch     | region, access_key_id, secretAccessKey |
 * | Cloudwatch v2          | cloudwatchV2  | region, access_key_id, secretAccessKey |
 * | Datadog                | datadog        | api_key, region |
 * | Datadog v2             | datadogV2     | api_key, region |
 * | Librato                | librato        | email, apiKey |
 * | New relic (deprecated) | newrelic       | - |
 * | New relic v2           | newrelicV2    | api_key, region |
 * | Stackdriver            | stackdriver    | project_id, private_key, clientEmail |
 *
 * ## Dependency
 *
 * This resource depends on CloudAMQP instance identifier, `cloudamqp_instance.instance.id`.
 *
 * ## Import
 *
 * `cloudamqp_integration_metric`can be imported using the name argument of the resource together with CloudAMQP instance identifier. The name and identifier are CSV separated, see example below.
 *
 * ```sh
 *  $ pulumi import cloudamqp:index/integrationMetric:IntegrationMetric <resource_name> <name>,<instance_id>`
 * ```
 */
=======
>>>>>>> 5ed2468... Upgrae to v1.8.3 of the CloudAMQP Terraform Provider
export class IntegrationMetric extends pulumi.CustomResource {
    /**
     * Get an existing IntegrationMetric resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: IntegrationMetricState, opts?: pulumi.CustomResourceOptions): IntegrationMetric {
        return new IntegrationMetric(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'cloudamqp:index/integrationMetric:IntegrationMetric';

    /**
     * Returns true if the given object is an instance of IntegrationMetric.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is IntegrationMetric {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === IntegrationMetric.__pulumiType;
    }

    /**
     * AWS access key identifier. (Cloudwatch)
     */
    public readonly accessKeyId!: pulumi.Output<string | undefined>;
    /**
     * The API key for the integration service. (Librato)
     */
    public readonly apiKey!: pulumi.Output<string | undefined>;
    /**
     * The client email. (Stackdriver)
     */
    public readonly clientEmail!: pulumi.Output<string | undefined>;
    /**
     * The email address registred for the integration service. (Librato)
     */
    public readonly email!: pulumi.Output<string | undefined>;
    /**
     * Instance identifier
     */
    public readonly instanceId!: pulumi.Output<number>;
    /**
     * The license key registred for the integration service. (New Relic)
     */
    public readonly licenseKey!: pulumi.Output<string | undefined>;
    /**
     * The name of metrics integration
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The private key. (Stackdriver)
     */
    public readonly privateKey!: pulumi.Output<string | undefined>;
    /**
     * Project ID. (Stackdriver)
     */
    public readonly projectId!: pulumi.Output<string | undefined>;
    /**
     * (optional) whitelist using regular expression
     */
    public readonly queueWhitelist!: pulumi.Output<string | undefined>;
    /**
     * AWS region for Cloudwatch and [US/EU] for Data dog/New relic. (Cloudwatch, Data Dog, New Relic)
     */
    public readonly region!: pulumi.Output<string | undefined>;
    /**
     * AWS secret key. (Cloudwatch)
     */
    public readonly secretAccessKey!: pulumi.Output<string | undefined>;
    /**
     * (optional) tags. E.g. env=prod,region=europe
     */
    public readonly tags!: pulumi.Output<string | undefined>;
    /**
     * (optional) whitelist using regular expression
     */
    public readonly vhostWhitelist!: pulumi.Output<string | undefined>;

    /**
     * Create a IntegrationMetric resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: IntegrationMetricArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: IntegrationMetricArgs | IntegrationMetricState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as IntegrationMetricState | undefined;
            inputs["accessKeyId"] = state ? state.accessKeyId : undefined;
            inputs["apiKey"] = state ? state.apiKey : undefined;
            inputs["clientEmail"] = state ? state.clientEmail : undefined;
            inputs["email"] = state ? state.email : undefined;
            inputs["instanceId"] = state ? state.instanceId : undefined;
            inputs["licenseKey"] = state ? state.licenseKey : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["privateKey"] = state ? state.privateKey : undefined;
            inputs["projectId"] = state ? state.projectId : undefined;
            inputs["queueWhitelist"] = state ? state.queueWhitelist : undefined;
            inputs["region"] = state ? state.region : undefined;
            inputs["secretAccessKey"] = state ? state.secretAccessKey : undefined;
            inputs["tags"] = state ? state.tags : undefined;
            inputs["vhostWhitelist"] = state ? state.vhostWhitelist : undefined;
        } else {
            const args = argsOrState as IntegrationMetricArgs | undefined;
            if (!args || args.instanceId === undefined) {
                throw new Error("Missing required property 'instanceId'");
            }
            inputs["accessKeyId"] = args ? args.accessKeyId : undefined;
            inputs["apiKey"] = args ? args.apiKey : undefined;
            inputs["clientEmail"] = args ? args.clientEmail : undefined;
            inputs["email"] = args ? args.email : undefined;
            inputs["instanceId"] = args ? args.instanceId : undefined;
            inputs["licenseKey"] = args ? args.licenseKey : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["privateKey"] = args ? args.privateKey : undefined;
            inputs["projectId"] = args ? args.projectId : undefined;
            inputs["queueWhitelist"] = args ? args.queueWhitelist : undefined;
            inputs["region"] = args ? args.region : undefined;
            inputs["secretAccessKey"] = args ? args.secretAccessKey : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["vhostWhitelist"] = args ? args.vhostWhitelist : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(IntegrationMetric.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering IntegrationMetric resources.
 */
export interface IntegrationMetricState {
    /**
     * AWS access key identifier. (Cloudwatch)
     */
    readonly accessKeyId?: pulumi.Input<string>;
    /**
     * The API key for the integration service. (Librato)
     */
    readonly apiKey?: pulumi.Input<string>;
    /**
     * The client email. (Stackdriver)
     */
    readonly clientEmail?: pulumi.Input<string>;
    /**
     * The email address registred for the integration service. (Librato)
     */
    readonly email?: pulumi.Input<string>;
    /**
     * Instance identifier
     */
    readonly instanceId?: pulumi.Input<number>;
    /**
     * The license key registred for the integration service. (New Relic)
     */
    readonly licenseKey?: pulumi.Input<string>;
    /**
     * The name of metrics integration
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The private key. (Stackdriver)
     */
    readonly privateKey?: pulumi.Input<string>;
    /**
     * Project ID. (Stackdriver)
     */
    readonly projectId?: pulumi.Input<string>;
    /**
     * (optional) whitelist using regular expression
     */
    readonly queueWhitelist?: pulumi.Input<string>;
    /**
     * AWS region for Cloudwatch and [US/EU] for Data dog/New relic. (Cloudwatch, Data Dog, New Relic)
     */
    readonly region?: pulumi.Input<string>;
    /**
     * AWS secret key. (Cloudwatch)
     */
    readonly secretAccessKey?: pulumi.Input<string>;
    /**
     * (optional) tags. E.g. env=prod,region=europe
     */
    readonly tags?: pulumi.Input<string>;
    /**
     * (optional) whitelist using regular expression
     */
    readonly vhostWhitelist?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a IntegrationMetric resource.
 */
export interface IntegrationMetricArgs {
    /**
     * AWS access key identifier. (Cloudwatch)
     */
    readonly accessKeyId?: pulumi.Input<string>;
    /**
     * The API key for the integration service. (Librato)
     */
    readonly apiKey?: pulumi.Input<string>;
    /**
     * The client email. (Stackdriver)
     */
    readonly clientEmail?: pulumi.Input<string>;
    /**
     * The email address registred for the integration service. (Librato)
     */
    readonly email?: pulumi.Input<string>;
    /**
     * Instance identifier
     */
    readonly instanceId: pulumi.Input<number>;
    /**
     * The license key registred for the integration service. (New Relic)
     */
    readonly licenseKey?: pulumi.Input<string>;
    /**
     * The name of metrics integration
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The private key. (Stackdriver)
     */
    readonly privateKey?: pulumi.Input<string>;
    /**
     * Project ID. (Stackdriver)
     */
    readonly projectId?: pulumi.Input<string>;
    /**
     * (optional) whitelist using regular expression
     */
    readonly queueWhitelist?: pulumi.Input<string>;
    /**
     * AWS region for Cloudwatch and [US/EU] for Data dog/New relic. (Cloudwatch, Data Dog, New Relic)
     */
    readonly region?: pulumi.Input<string>;
    /**
     * AWS secret key. (Cloudwatch)
     */
    readonly secretAccessKey?: pulumi.Input<string>;
    /**
     * (optional) tags. E.g. env=prod,region=europe
     */
    readonly tags?: pulumi.Input<string>;
    /**
     * (optional) whitelist using regular expression
     */
    readonly vhostWhitelist?: pulumi.Input<string>;
}
