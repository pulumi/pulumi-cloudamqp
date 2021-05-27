// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * This resource allows you to create and manage a CloudAMQP instance running Rabbit MQ and deploy to multiple cloud platforms provider and over multiple regions, see Instance regions for more information.
 *
 * Once the instance is created it will be assigned a unique identifier. All other resource and data sources created for this instance needs to reference the instance identifier.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as cloudamqp from "@pulumi/cloudamqp";
 *
 * // Minimum free lemur instance
 * const lemurInstance = new cloudamqp.Instance("lemur_instance", {
 *     plan: "lemur",
 *     region: "amazon-web-services::us-west-1",
 * });
 * // New dedicated bunny instance
 * const instance = new cloudamqp.Instance("instance", {
 *     noDefaultAlarms: true,
 *     nodes: 1,
 *     plan: "bunny-1",
 *     region: "amazon-web-services::us-west-1",
 *     rmqVersion: "3.8.3",
 *     tags: ["terraform"],
 * });
 * ```
 *
 * ## Import
 *
 * `cloudamqp_instance`can be imported using CloudAMQP internal identifier. To retrieve the identifier for an instance, use [CloudAMQP customer API](https://docs.cloudamqp.com/#list-instances).
 *
 * ```sh
 *  $ pulumi import cloudamqp:index/instance:Instance instance <instance_id>`
 * ```
 */
export class Instance extends pulumi.CustomResource {
    /**
     * Get an existing Instance resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: InstanceState, opts?: pulumi.CustomResourceOptions): Instance {
        return new Instance(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'cloudamqp:index/instance:Instance';

    /**
     * Returns true if the given object is an instance of Instance.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Instance {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Instance.__pulumiType;
    }

    /**
     * (Computed) API key needed to communicate to CloudAMQP's second API. The second API is used to manage alarms, integration and more, full description [CloudAMQP API](https://docs.cloudamqp.com/cloudamqp_api.html).
     */
    public /*out*/ readonly apikey!: pulumi.Output<string>;
    /**
     * Is the instance hosted on a dedicated server
     */
    public /*out*/ readonly dedicated!: pulumi.Output<boolean>;
    /**
     * (Computed) The host name for the CloudAMQP instance.
     */
    public /*out*/ readonly host!: pulumi.Output<string>;
    /**
     * Name of the CloudAMQP instance.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Set to true to discard creating default alarms when the instance is created. Can be left out, will then use default value = false.
     */
    public readonly noDefaultAlarms!: pulumi.Output<boolean>;
    /**
     * Number of nodes, 1, 3 or 5. **Note: Changed from optional to computed. In order to change number of nodes, the subscription plan needs to be updated.**
     */
    public readonly nodes!: pulumi.Output<number | undefined>;
    /**
     * The subscription plan. See available plans
     */
    public readonly plan!: pulumi.Output<string>;
    /**
     * Flag describing if the resource is ready
     */
    public /*out*/ readonly ready!: pulumi.Output<boolean>;
    /**
     * The region to host the instance in. See Instance regions
     */
    public readonly region!: pulumi.Output<string>;
    /**
     * The Rabbit MQ version. Can be left out, will then be set to default value used by CloudAMQP API. **Note: There is not yet any support in the provider to change the RMQ version. Once it's set in the initial creation, it will remain.**
     */
    public readonly rmqVersion!: pulumi.Output<string>;
    /**
     * One or more tags for the CloudAMQP instance, makes it possible to categories multiple instances in console view. Default there is no tags assigned.
     */
    public readonly tags!: pulumi.Output<string[] | undefined>;
    /**
     * (Computed) AMQP server endpoint. `amqps://{username}:{password}@{hostname}/{vhost}`
     */
    public /*out*/ readonly url!: pulumi.Output<string>;
    /**
     * (Computed) The virtual host used by Rabbit MQ.
     */
    public /*out*/ readonly vhost!: pulumi.Output<string>;
    /**
     * Creates a dedicated VPC subnet, shouldn't overlap with other VPC subnet, default subnet used 10.56.72.0/24. **NOTE: extra fee will be charged when using VPC, see [CloudAMQP](https://cloudamqp.com) for more information.**
     */
    public readonly vpcSubnet!: pulumi.Output<string | undefined>;

    /**
     * Create a Instance resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: InstanceArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: InstanceArgs | InstanceState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as InstanceState | undefined;
            inputs["apikey"] = state ? state.apikey : undefined;
            inputs["dedicated"] = state ? state.dedicated : undefined;
            inputs["host"] = state ? state.host : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["noDefaultAlarms"] = state ? state.noDefaultAlarms : undefined;
            inputs["nodes"] = state ? state.nodes : undefined;
            inputs["plan"] = state ? state.plan : undefined;
            inputs["ready"] = state ? state.ready : undefined;
            inputs["region"] = state ? state.region : undefined;
            inputs["rmqVersion"] = state ? state.rmqVersion : undefined;
            inputs["tags"] = state ? state.tags : undefined;
            inputs["url"] = state ? state.url : undefined;
            inputs["vhost"] = state ? state.vhost : undefined;
            inputs["vpcSubnet"] = state ? state.vpcSubnet : undefined;
        } else {
            const args = argsOrState as InstanceArgs | undefined;
            if ((!args || args.plan === undefined) && !opts.urn) {
                throw new Error("Missing required property 'plan'");
            }
            if ((!args || args.region === undefined) && !opts.urn) {
                throw new Error("Missing required property 'region'");
            }
            inputs["name"] = args ? args.name : undefined;
            inputs["noDefaultAlarms"] = args ? args.noDefaultAlarms : undefined;
            inputs["nodes"] = args ? args.nodes : undefined;
            inputs["plan"] = args ? args.plan : undefined;
            inputs["region"] = args ? args.region : undefined;
            inputs["rmqVersion"] = args ? args.rmqVersion : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["vpcSubnet"] = args ? args.vpcSubnet : undefined;
            inputs["apikey"] = undefined /*out*/;
            inputs["dedicated"] = undefined /*out*/;
            inputs["host"] = undefined /*out*/;
            inputs["ready"] = undefined /*out*/;
            inputs["url"] = undefined /*out*/;
            inputs["vhost"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(Instance.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Instance resources.
 */
export interface InstanceState {
    /**
     * (Computed) API key needed to communicate to CloudAMQP's second API. The second API is used to manage alarms, integration and more, full description [CloudAMQP API](https://docs.cloudamqp.com/cloudamqp_api.html).
     */
    apikey?: pulumi.Input<string>;
    /**
     * Is the instance hosted on a dedicated server
     */
    dedicated?: pulumi.Input<boolean>;
    /**
     * (Computed) The host name for the CloudAMQP instance.
     */
    host?: pulumi.Input<string>;
    /**
     * Name of the CloudAMQP instance.
     */
    name?: pulumi.Input<string>;
    /**
     * Set to true to discard creating default alarms when the instance is created. Can be left out, will then use default value = false.
     */
    noDefaultAlarms?: pulumi.Input<boolean>;
    /**
     * Number of nodes, 1, 3 or 5. **Note: Changed from optional to computed. In order to change number of nodes, the subscription plan needs to be updated.**
     */
    nodes?: pulumi.Input<number>;
    /**
     * The subscription plan. See available plans
     */
    plan?: pulumi.Input<string>;
    /**
     * Flag describing if the resource is ready
     */
    ready?: pulumi.Input<boolean>;
    /**
     * The region to host the instance in. See Instance regions
     */
    region?: pulumi.Input<string>;
    /**
     * The Rabbit MQ version. Can be left out, will then be set to default value used by CloudAMQP API. **Note: There is not yet any support in the provider to change the RMQ version. Once it's set in the initial creation, it will remain.**
     */
    rmqVersion?: pulumi.Input<string>;
    /**
     * One or more tags for the CloudAMQP instance, makes it possible to categories multiple instances in console view. Default there is no tags assigned.
     */
    tags?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * (Computed) AMQP server endpoint. `amqps://{username}:{password}@{hostname}/{vhost}`
     */
    url?: pulumi.Input<string>;
    /**
     * (Computed) The virtual host used by Rabbit MQ.
     */
    vhost?: pulumi.Input<string>;
    /**
     * Creates a dedicated VPC subnet, shouldn't overlap with other VPC subnet, default subnet used 10.56.72.0/24. **NOTE: extra fee will be charged when using VPC, see [CloudAMQP](https://cloudamqp.com) for more information.**
     */
    vpcSubnet?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Instance resource.
 */
export interface InstanceArgs {
    /**
     * Name of the CloudAMQP instance.
     */
    name?: pulumi.Input<string>;
    /**
     * Set to true to discard creating default alarms when the instance is created. Can be left out, will then use default value = false.
     */
    noDefaultAlarms?: pulumi.Input<boolean>;
    /**
     * Number of nodes, 1, 3 or 5. **Note: Changed from optional to computed. In order to change number of nodes, the subscription plan needs to be updated.**
     */
    nodes?: pulumi.Input<number>;
    /**
     * The subscription plan. See available plans
     */
    plan: pulumi.Input<string>;
    /**
     * The region to host the instance in. See Instance regions
     */
    region: pulumi.Input<string>;
    /**
     * The Rabbit MQ version. Can be left out, will then be set to default value used by CloudAMQP API. **Note: There is not yet any support in the provider to change the RMQ version. Once it's set in the initial creation, it will remain.**
     */
    rmqVersion?: pulumi.Input<string>;
    /**
     * One or more tags for the CloudAMQP instance, makes it possible to categories multiple instances in console view. Default there is no tags assigned.
     */
    tags?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Creates a dedicated VPC subnet, shouldn't overlap with other VPC subnet, default subnet used 10.56.72.0/24. **NOTE: extra fee will be charged when using VPC, see [CloudAMQP](https://cloudamqp.com) for more information.**
     */
    vpcSubnet?: pulumi.Input<string>;
}
