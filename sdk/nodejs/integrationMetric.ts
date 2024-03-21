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
    public readonly clientEmail!: pulumi.Output<string>;
    /**
     * Base64Encoded credentials. (Stackdriver)
     */
    public readonly credentials!: pulumi.Output<string | undefined>;
    /**
     * The email address registred for the integration service. (Librato)
     */
    public readonly email!: pulumi.Output<string | undefined>;
    /**
     * External identifier that match the role you created. (Cloudwatch)
     */
    public readonly iamExternalId!: pulumi.Output<string | undefined>;
    /**
     * The ARN of the role to be assumed when publishing metrics. (Cloudwatch)
     */
    public readonly iamRole!: pulumi.Output<string | undefined>;
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
    public readonly privateKey!: pulumi.Output<string>;
    /**
     * Private key identifier. (Stackdriver)
     */
    public readonly privateKeyId!: pulumi.Output<string>;
    /**
     * Project ID. (Stackdriver)
     */
    public readonly projectId!: pulumi.Output<string>;
    /**
     * (optional) allowlist using regular expression
     */
    public readonly queueAllowlist!: pulumi.Output<string | undefined>;
    /**
     * **Deprecated**
     *
     * @deprecated use queueAllowlist instead
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
     * @deprecated use vhostAllowlist instead
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
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as IntegrationMetricState | undefined;
            resourceInputs["accessKeyId"] = state ? state.accessKeyId : undefined;
            resourceInputs["apiKey"] = state ? state.apiKey : undefined;
            resourceInputs["clientEmail"] = state ? state.clientEmail : undefined;
            resourceInputs["credentials"] = state ? state.credentials : undefined;
            resourceInputs["email"] = state ? state.email : undefined;
            resourceInputs["iamExternalId"] = state ? state.iamExternalId : undefined;
            resourceInputs["iamRole"] = state ? state.iamRole : undefined;
            resourceInputs["instanceId"] = state ? state.instanceId : undefined;
            resourceInputs["licenseKey"] = state ? state.licenseKey : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["privateKey"] = state ? state.privateKey : undefined;
            resourceInputs["privateKeyId"] = state ? state.privateKeyId : undefined;
            resourceInputs["projectId"] = state ? state.projectId : undefined;
            resourceInputs["queueAllowlist"] = state ? state.queueAllowlist : undefined;
            resourceInputs["queueWhitelist"] = state ? state.queueWhitelist : undefined;
            resourceInputs["region"] = state ? state.region : undefined;
            resourceInputs["secretAccessKey"] = state ? state.secretAccessKey : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["vhostAllowlist"] = state ? state.vhostAllowlist : undefined;
            resourceInputs["vhostWhitelist"] = state ? state.vhostWhitelist : undefined;
        } else {
            const args = argsOrState as IntegrationMetricArgs | undefined;
            if ((!args || args.instanceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'instanceId'");
            }
            resourceInputs["accessKeyId"] = args ? args.accessKeyId : undefined;
            resourceInputs["apiKey"] = args ? args.apiKey : undefined;
            resourceInputs["clientEmail"] = args ? args.clientEmail : undefined;
            resourceInputs["credentials"] = args?.credentials ? pulumi.secret(args.credentials) : undefined;
            resourceInputs["email"] = args ? args.email : undefined;
            resourceInputs["iamExternalId"] = args ? args.iamExternalId : undefined;
            resourceInputs["iamRole"] = args ? args.iamRole : undefined;
            resourceInputs["instanceId"] = args ? args.instanceId : undefined;
            resourceInputs["licenseKey"] = args ? args.licenseKey : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["privateKey"] = args?.privateKey ? pulumi.secret(args.privateKey) : undefined;
            resourceInputs["privateKeyId"] = args?.privateKeyId ? pulumi.secret(args.privateKeyId) : undefined;
            resourceInputs["projectId"] = args ? args.projectId : undefined;
            resourceInputs["queueAllowlist"] = args ? args.queueAllowlist : undefined;
            resourceInputs["queueWhitelist"] = args ? args.queueWhitelist : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
            resourceInputs["secretAccessKey"] = args?.secretAccessKey ? pulumi.secret(args.secretAccessKey) : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["vhostAllowlist"] = args ? args.vhostAllowlist : undefined;
            resourceInputs["vhostWhitelist"] = args ? args.vhostWhitelist : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["credentials", "privateKey", "privateKeyId", "secretAccessKey"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(IntegrationMetric.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering IntegrationMetric resources.
 */
export interface IntegrationMetricState {
    /**
     * AWS access key identifier. (Cloudwatch)
     */
    accessKeyId?: pulumi.Input<string>;
    /**
     * The API key for the integration service. (Librato)
     */
    apiKey?: pulumi.Input<string>;
    /**
     * The client email. (Stackdriver)
     */
    clientEmail?: pulumi.Input<string>;
    /**
     * Base64Encoded credentials. (Stackdriver)
     */
    credentials?: pulumi.Input<string>;
    /**
     * The email address registred for the integration service. (Librato)
     */
    email?: pulumi.Input<string>;
    /**
     * External identifier that match the role you created. (Cloudwatch)
     */
    iamExternalId?: pulumi.Input<string>;
    /**
     * The ARN of the role to be assumed when publishing metrics. (Cloudwatch)
     */
    iamRole?: pulumi.Input<string>;
    /**
     * Instance identifier
     */
    instanceId?: pulumi.Input<number>;
    /**
     * The license key registred for the integration service. (New Relic)
     */
    licenseKey?: pulumi.Input<string>;
    /**
     * The name of metrics integration
     */
    name?: pulumi.Input<string>;
    /**
     * The private key. (Stackdriver)
     */
    privateKey?: pulumi.Input<string>;
    /**
     * Private key identifier. (Stackdriver)
     */
    privateKeyId?: pulumi.Input<string>;
    /**
     * Project ID. (Stackdriver)
     */
    projectId?: pulumi.Input<string>;
    /**
     * (optional) allowlist using regular expression
     */
    queueAllowlist?: pulumi.Input<string>;
    /**
     * **Deprecated**
     *
     * @deprecated use queueAllowlist instead
     */
    queueWhitelist?: pulumi.Input<string>;
    /**
     * AWS region for Cloudwatch and [US/EU] for Data dog/New relic. (Cloudwatch, Data Dog, New Relic)
     */
    region?: pulumi.Input<string>;
    /**
     * AWS secret key. (Cloudwatch)
     */
    secretAccessKey?: pulumi.Input<string>;
    /**
     * (optional) tags. E.g. env=prod,region=europe
     */
    tags?: pulumi.Input<string>;
    /**
     * (optional) allowlist using regular expression
     */
    vhostAllowlist?: pulumi.Input<string>;
    /**
     * **Deprecated**
     *
     * @deprecated use vhostAllowlist instead
     */
    vhostWhitelist?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a IntegrationMetric resource.
 */
export interface IntegrationMetricArgs {
    /**
     * AWS access key identifier. (Cloudwatch)
     */
    accessKeyId?: pulumi.Input<string>;
    /**
     * The API key for the integration service. (Librato)
     */
    apiKey?: pulumi.Input<string>;
    /**
     * The client email. (Stackdriver)
     */
    clientEmail?: pulumi.Input<string>;
    /**
     * Base64Encoded credentials. (Stackdriver)
     */
    credentials?: pulumi.Input<string>;
    /**
     * The email address registred for the integration service. (Librato)
     */
    email?: pulumi.Input<string>;
    /**
     * External identifier that match the role you created. (Cloudwatch)
     */
    iamExternalId?: pulumi.Input<string>;
    /**
     * The ARN of the role to be assumed when publishing metrics. (Cloudwatch)
     */
    iamRole?: pulumi.Input<string>;
    /**
     * Instance identifier
     */
    instanceId: pulumi.Input<number>;
    /**
     * The license key registred for the integration service. (New Relic)
     */
    licenseKey?: pulumi.Input<string>;
    /**
     * The name of metrics integration
     */
    name?: pulumi.Input<string>;
    /**
     * The private key. (Stackdriver)
     */
    privateKey?: pulumi.Input<string>;
    /**
     * Private key identifier. (Stackdriver)
     */
    privateKeyId?: pulumi.Input<string>;
    /**
     * Project ID. (Stackdriver)
     */
    projectId?: pulumi.Input<string>;
    /**
     * (optional) allowlist using regular expression
     */
    queueAllowlist?: pulumi.Input<string>;
    /**
     * **Deprecated**
     *
     * @deprecated use queueAllowlist instead
     */
    queueWhitelist?: pulumi.Input<string>;
    /**
     * AWS region for Cloudwatch and [US/EU] for Data dog/New relic. (Cloudwatch, Data Dog, New Relic)
     */
    region?: pulumi.Input<string>;
    /**
     * AWS secret key. (Cloudwatch)
     */
    secretAccessKey?: pulumi.Input<string>;
    /**
     * (optional) tags. E.g. env=prod,region=europe
     */
    tags?: pulumi.Input<string>;
    /**
     * (optional) allowlist using regular expression
     */
    vhostAllowlist?: pulumi.Input<string>;
    /**
     * **Deprecated**
     *
     * @deprecated use vhostAllowlist instead
     */
    vhostWhitelist?: pulumi.Input<string>;
}
