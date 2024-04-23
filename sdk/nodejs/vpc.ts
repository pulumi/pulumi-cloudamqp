// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * This resource allows you to manage standalone VPC.
 *
 * New Cloudamqp instances can be added to the managed VPC. Set the instance *vpc_id* attribute to the managed vpc identifier , see example below, when creating the instance.
 *
 * Only available for dedicated subscription plans.
 *
 * Pricing is available at [cloudamqp.com](https://www.cloudamqp.com/plans.html).
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as cloudamqp from "@pulumi/cloudamqp";
 *
 * // Managed VPC resource
 * const vpc = new cloudamqp.Vpc("vpc", {
 *     name: "<VPC name>",
 *     region: "amazon-web-services::us-east-1",
 *     subnet: "10.56.72.0/24",
 *     tags: [],
 * });
 * //  New instance, need to be created with a vpc
 * const instance = new cloudamqp.Instance("instance", {
 *     name: "<Instance name>",
 *     plan: "bunny-1",
 *     region: "amazon-web-services::us-east-1",
 *     nodes: 1,
 *     tags: [],
 *     rmqVersion: "3.9.13",
 *     vpcId: vpcCloudamqVpc.id,
 *     keepAssociatedVpc: true,
 * });
 * // Additional VPC information
 * const vpcInfo = cloudamqp.getVpcInfoOutput({
 *     vpcId: vpc.id,
 * });
 * ```
 *
 * ## Import
 *
 * `cloudamqp_vpc` can be imported using the CloudAMQP VPC identifier.
 *
 * ```sh
 * $ pulumi import cloudamqp:index/vpc:Vpc <resource_name> <vpc_id>`
 * ```
 *
 * To retrieve the identifier for a VPC, either use [CloudAMQP customer API](https://docs.cloudamqp.com/#list-vpcs).
 *
 * Or use the data source `cloudamqp_account_vpcs` to list all available standalone VPCs for an account.
 */
export class Vpc extends pulumi.CustomResource {
    /**
     * Get an existing Vpc resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: VpcState, opts?: pulumi.CustomResourceOptions): Vpc {
        return new Vpc(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'cloudamqp:index/vpc:Vpc';

    /**
     * Returns true if the given object is an instance of Vpc.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Vpc {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Vpc.__pulumiType;
    }

    /**
     * The name of the VPC.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The hosted region for the managed standalone VPC
     */
    public readonly region!: pulumi.Output<string>;
    /**
     * The VPC subnet
     */
    public readonly subnet!: pulumi.Output<string>;
    /**
     * Tag the VPC with optional tags
     */
    public readonly tags!: pulumi.Output<string[] | undefined>;
    /**
     * VPC name given when hosted at the cloud provider
     */
    public /*out*/ readonly vpcName!: pulumi.Output<string>;

    /**
     * Create a Vpc resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: VpcArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: VpcArgs | VpcState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as VpcState | undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["region"] = state ? state.region : undefined;
            resourceInputs["subnet"] = state ? state.subnet : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["vpcName"] = state ? state.vpcName : undefined;
        } else {
            const args = argsOrState as VpcArgs | undefined;
            if ((!args || args.region === undefined) && !opts.urn) {
                throw new Error("Missing required property 'region'");
            }
            if ((!args || args.subnet === undefined) && !opts.urn) {
                throw new Error("Missing required property 'subnet'");
            }
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
            resourceInputs["subnet"] = args ? args.subnet : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["vpcName"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Vpc.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Vpc resources.
 */
export interface VpcState {
    /**
     * The name of the VPC.
     */
    name?: pulumi.Input<string>;
    /**
     * The hosted region for the managed standalone VPC
     */
    region?: pulumi.Input<string>;
    /**
     * The VPC subnet
     */
    subnet?: pulumi.Input<string>;
    /**
     * Tag the VPC with optional tags
     */
    tags?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * VPC name given when hosted at the cloud provider
     */
    vpcName?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Vpc resource.
 */
export interface VpcArgs {
    /**
     * The name of the VPC.
     */
    name?: pulumi.Input<string>;
    /**
     * The hosted region for the managed standalone VPC
     */
    region: pulumi.Input<string>;
    /**
     * The VPC subnet
     */
    subnet: pulumi.Input<string>;
    /**
     * Tag the VPC with optional tags
     */
    tags?: pulumi.Input<pulumi.Input<string>[]>;
}
