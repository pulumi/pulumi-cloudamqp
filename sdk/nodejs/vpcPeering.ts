// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * ## Import
 *
 * Not possible to import this resource.
 */
export class VpcPeering extends pulumi.CustomResource {
    /**
     * Get an existing VpcPeering resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: VpcPeeringState, opts?: pulumi.CustomResourceOptions): VpcPeering {
        return new VpcPeering(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'cloudamqp:index/vpcPeering:VpcPeering';

    /**
     * Returns true if the given object is an instance of VpcPeering.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is VpcPeering {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === VpcPeering.__pulumiType;
    }

    /**
     * The CloudAMQP instance identifier.
     *
     * ***Deprecated: Changed from required to optional in v1.16.0, will be removed in next major version (v2.0)***
     */
    public readonly instanceId!: pulumi.Output<number | undefined>;
    /**
     * Peering identifier created by AW peering request.
     */
    public readonly peeringId!: pulumi.Output<string>;
    /**
     * Configurable sleep time (seconds) between retries for accepting or removing peering. Default set to 60 seconds.
     */
    public readonly sleep!: pulumi.Output<number | undefined>;
    /**
     * VPC peering status
     */
    public /*out*/ readonly status!: pulumi.Output<string>;
    /**
     * Configurable timeout time (seconds) for accepting or removing peering. Default set to 3600 seconds.
     */
    public readonly timeout!: pulumi.Output<number | undefined>;
    /**
     * The managed VPC identifier.
     *
     * ***Note: Introduced as optional in version v1.16.0, will be required in next major version (v2.0)***
     */
    public readonly vpcId!: pulumi.Output<string | undefined>;

    /**
     * Create a VpcPeering resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: VpcPeeringArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: VpcPeeringArgs | VpcPeeringState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as VpcPeeringState | undefined;
            resourceInputs["instanceId"] = state ? state.instanceId : undefined;
            resourceInputs["peeringId"] = state ? state.peeringId : undefined;
            resourceInputs["sleep"] = state ? state.sleep : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["timeout"] = state ? state.timeout : undefined;
            resourceInputs["vpcId"] = state ? state.vpcId : undefined;
        } else {
            const args = argsOrState as VpcPeeringArgs | undefined;
            if ((!args || args.peeringId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'peeringId'");
            }
            resourceInputs["instanceId"] = args ? args.instanceId : undefined;
            resourceInputs["peeringId"] = args ? args.peeringId : undefined;
            resourceInputs["sleep"] = args ? args.sleep : undefined;
            resourceInputs["timeout"] = args ? args.timeout : undefined;
            resourceInputs["vpcId"] = args ? args.vpcId : undefined;
            resourceInputs["status"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(VpcPeering.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering VpcPeering resources.
 */
export interface VpcPeeringState {
    /**
     * The CloudAMQP instance identifier.
     *
     * ***Deprecated: Changed from required to optional in v1.16.0, will be removed in next major version (v2.0)***
     */
    instanceId?: pulumi.Input<number>;
    /**
     * Peering identifier created by AW peering request.
     */
    peeringId?: pulumi.Input<string>;
    /**
     * Configurable sleep time (seconds) between retries for accepting or removing peering. Default set to 60 seconds.
     */
    sleep?: pulumi.Input<number>;
    /**
     * VPC peering status
     */
    status?: pulumi.Input<string>;
    /**
     * Configurable timeout time (seconds) for accepting or removing peering. Default set to 3600 seconds.
     */
    timeout?: pulumi.Input<number>;
    /**
     * The managed VPC identifier.
     *
     * ***Note: Introduced as optional in version v1.16.0, will be required in next major version (v2.0)***
     */
    vpcId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a VpcPeering resource.
 */
export interface VpcPeeringArgs {
    /**
     * The CloudAMQP instance identifier.
     *
     * ***Deprecated: Changed from required to optional in v1.16.0, will be removed in next major version (v2.0)***
     */
    instanceId?: pulumi.Input<number>;
    /**
     * Peering identifier created by AW peering request.
     */
    peeringId: pulumi.Input<string>;
    /**
     * Configurable sleep time (seconds) between retries for accepting or removing peering. Default set to 60 seconds.
     */
    sleep?: pulumi.Input<number>;
    /**
     * Configurable timeout time (seconds) for accepting or removing peering. Default set to 3600 seconds.
     */
    timeout?: pulumi.Input<number>;
    /**
     * The managed VPC identifier.
     *
     * ***Note: Introduced as optional in version v1.16.0, will be required in next major version (v2.0)***
     */
    vpcId?: pulumi.Input<string>;
}
