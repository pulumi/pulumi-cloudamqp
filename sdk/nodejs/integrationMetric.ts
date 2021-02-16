// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

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
     * (optional) allowlist using regular expression
     */
    public readonly queueAllowlist!: pulumi.Output<string | undefined>;
    /**
     * **Deprecated**
     *
     * @deprecated use queue_allowlist instead
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
     * (optional) allowlist using regular expression
     */
    public readonly vhostAllowlist!: pulumi.Output<string | undefined>;
    /**
     * **Deprecated**
     *
     * @deprecated use vhost_allowlist instead
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
        opts = opts || {};
        if (opts.id) {
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
            inputs["queueAllowlist"] = state ? state.queueAllowlist : undefined;
            inputs["queueWhitelist"] = state ? state.queueWhitelist : undefined;
            inputs["region"] = state ? state.region : undefined;
            inputs["secretAccessKey"] = state ? state.secretAccessKey : undefined;
            inputs["tags"] = state ? state.tags : undefined;
            inputs["vhostAllowlist"] = state ? state.vhostAllowlist : undefined;
            inputs["vhostWhitelist"] = state ? state.vhostWhitelist : undefined;
        } else {
            const args = argsOrState as IntegrationMetricArgs | undefined;
            if ((!args || args.instanceId === undefined) && !opts.urn) {
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
            inputs["queueAllowlist"] = args ? args.queueAllowlist : undefined;
            inputs["queueWhitelist"] = args ? args.queueWhitelist : undefined;
            inputs["region"] = args ? args.region : undefined;
            inputs["secretAccessKey"] = args ? args.secretAccessKey : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["vhostAllowlist"] = args ? args.vhostAllowlist : undefined;
            inputs["vhostWhitelist"] = args ? args.vhostWhitelist : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
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
     * (optional) allowlist using regular expression
     */
    readonly queueAllowlist?: pulumi.Input<string>;
    /**
     * **Deprecated**
     *
     * @deprecated use queue_allowlist instead
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
     * (optional) allowlist using regular expression
     */
    readonly vhostAllowlist?: pulumi.Input<string>;
    /**
     * **Deprecated**
     *
     * @deprecated use vhost_allowlist instead
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
     * (optional) allowlist using regular expression
     */
    readonly queueAllowlist?: pulumi.Input<string>;
    /**
     * **Deprecated**
     *
     * @deprecated use queue_allowlist instead
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
     * (optional) allowlist using regular expression
     */
    readonly vhostAllowlist?: pulumi.Input<string>;
    /**
     * **Deprecated**
     *
     * @deprecated use vhost_allowlist instead
     */
    readonly vhostWhitelist?: pulumi.Input<string>;
}
