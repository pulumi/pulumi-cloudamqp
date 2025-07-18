// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * ## Import
 *
 * Not possible to import this resource.
 */
export class UpgradeRabbitmq extends pulumi.CustomResource {
    /**
     * Get an existing UpgradeRabbitmq resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: UpgradeRabbitmqState, opts?: pulumi.CustomResourceOptions): UpgradeRabbitmq {
        return new UpgradeRabbitmq(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'cloudamqp:index/upgradeRabbitmq:UpgradeRabbitmq';

    /**
     * Returns true if the given object is an instance of UpgradeRabbitmq.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is UpgradeRabbitmq {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === UpgradeRabbitmq.__pulumiType;
    }

    /**
     * Helper argument to change upgrade behaviour to latest possible
     * version
     */
    public readonly currentVersion!: pulumi.Output<string | undefined>;
    /**
     * The CloudAMQP instance identifier
     */
    public readonly instanceId!: pulumi.Output<number>;
    /**
     * The new version to upgrade to
     */
    public readonly newVersion!: pulumi.Output<string | undefined>;

    /**
     * Create a UpgradeRabbitmq resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: UpgradeRabbitmqArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: UpgradeRabbitmqArgs | UpgradeRabbitmqState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as UpgradeRabbitmqState | undefined;
            resourceInputs["currentVersion"] = state ? state.currentVersion : undefined;
            resourceInputs["instanceId"] = state ? state.instanceId : undefined;
            resourceInputs["newVersion"] = state ? state.newVersion : undefined;
        } else {
            const args = argsOrState as UpgradeRabbitmqArgs | undefined;
            if ((!args || args.instanceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'instanceId'");
            }
            resourceInputs["currentVersion"] = args ? args.currentVersion : undefined;
            resourceInputs["instanceId"] = args ? args.instanceId : undefined;
            resourceInputs["newVersion"] = args ? args.newVersion : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(UpgradeRabbitmq.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering UpgradeRabbitmq resources.
 */
export interface UpgradeRabbitmqState {
    /**
     * Helper argument to change upgrade behaviour to latest possible
     * version
     */
    currentVersion?: pulumi.Input<string>;
    /**
     * The CloudAMQP instance identifier
     */
    instanceId?: pulumi.Input<number>;
    /**
     * The new version to upgrade to
     */
    newVersion?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a UpgradeRabbitmq resource.
 */
export interface UpgradeRabbitmqArgs {
    /**
     * Helper argument to change upgrade behaviour to latest possible
     * version
     */
    currentVersion?: pulumi.Input<string>;
    /**
     * The CloudAMQP instance identifier
     */
    instanceId: pulumi.Input<number>;
    /**
     * The new version to upgrade to
     */
    newVersion?: pulumi.Input<string>;
}
