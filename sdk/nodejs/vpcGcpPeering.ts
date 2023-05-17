// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * This resouce creates a VPC peering configuration for the CloudAMQP instance. The configuration will connect to another VPC network hosted on Google Cloud Platform (GCP). See the [GCP documentation](https://cloud.google.com/vpc/docs/using-vpc-peering) for more information on how to create the VPC peering configuration.
 *
 * Only available for dedicated subscription plans.
 *
 * Pricing is available at [cloudamqp.com](https://www.cloudamqp.com/plans.html).
 *
 * ## Example Usage
 * ### With Additional Firewall Rules
 *
 * <details>
 *   <summary>
 *     <b>
 *       <i>VPC peering pre v1.16.0</i>
 *     </b>
 *   </summary>
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as cloudamqp from "@pulumi/cloudamqp";
 *
 * // VPC peering configuration
 * const vpcPeeringRequest = new cloudamqp.VpcGcpPeering("vpcPeeringRequest", {
 *     instanceId: cloudamqp_instance.instance.id,
 *     peerNetworkUri: _var.peer_network_uri,
 * });
 * // Firewall rules
 * const firewallSettings = new cloudamqp.SecurityFirewall("firewallSettings", {
 *     instanceId: cloudamqp_instance.instance.id,
 *     rules: [
 *         {
 *             ip: _var.peer_subnet,
 *             ports: [15672],
 *             services: [
 *                 "AMQP",
 *                 "AMQPS",
 *                 "STREAM",
 *                 "STREAM_SSL",
 *             ],
 *             description: "VPC peering for <NETWORK>",
 *         },
 *         {
 *             ip: "192.168.0.0/24",
 *             ports: [
 *                 4567,
 *                 4568,
 *             ],
 *             services: [
 *                 "AMQP",
 *                 "AMQPS",
 *                 "HTTPS",
 *             ],
 *         },
 *     ],
 * }, {
 *     dependsOn: [vpcPeeringRequest],
 * });
 * ```
 * </details>
 *
 * <details>
 *   <summary>
 *     <b>
 *       <i>VPC peering post v1.16.0 (Managed VPC)</i>
 *     </b>
 *   </summary>
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as cloudamqp from "@pulumi/cloudamqp";
 *
 * // VPC peering configuration
 * const vpcPeeringRequest = new cloudamqp.VpcGcpPeering("vpcPeeringRequest", {
 *     vpcId: cloudamqp_vpc.vpc.id,
 *     peerNetworkUri: _var.peer_network_uri,
 * });
 * // Firewall rules
 * const firewallSettings = new cloudamqp.SecurityFirewall("firewallSettings", {
 *     instanceId: cloudamqp_instance.instance.id,
 *     rules: [
 *         {
 *             ip: _var.peer_subnet,
 *             ports: [15672],
 *             services: [
 *                 "AMQP",
 *                 "AMQPS",
 *                 "STREAM",
 *                 "STREAM_SSL",
 *             ],
 *             description: "VPC peering for <NETWORK>",
 *         },
 *         {
 *             ip: "192.168.0.0/24",
 *             ports: [
 *                 4567,
 *                 4568,
 *             ],
 *             services: [
 *                 "AMQP",
 *                 "AMQPS",
 *                 "HTTPS",
 *             ],
 *         },
 *     ],
 * }, {
 *     dependsOn: [vpcPeeringRequest],
 * });
 * ```
 * </details>
 * ## Depedency
 *
 * *Pre v1.16.0*
 * This resource depends on CloudAMQP instance identifier, `cloudamqp_instance.instance.id`.
 *
 * *Post v1.16.0*
 * This resource depends on CloudAMQP managed VPC identifier, `cloudamqp_vpc.vpc.id` or instance identifier, `cloudamqp_instance.instance.id`.
 *
 * ## Create VPC Peering with additional firewall rules
 *
 * To create a VPC peering configuration with additional firewall rules, it's required to chain the cloudamqp.SecurityFirewall
 * resource to avoid parallel conflicting resource calls. This is done by adding dependency from the firewall resource to the VPC peering resource.
 *
 * Furthermore, since all firewall rules are overwritten, the otherwise automatically added rules for the VPC peering also needs to be added.
 *
 * See example below.
 *
 * ## Import
 *
 * Not possible to import this resource.
 */
export class VpcGcpPeering extends pulumi.CustomResource {
    /**
     * Get an existing VpcGcpPeering resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: VpcGcpPeeringState, opts?: pulumi.CustomResourceOptions): VpcGcpPeering {
        return new VpcGcpPeering(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'cloudamqp:index/vpcGcpPeering:VpcGcpPeering';

    /**
     * Returns true if the given object is an instance of VpcGcpPeering.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is VpcGcpPeering {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === VpcGcpPeering.__pulumiType;
    }

    /**
     * VPC peering auto created routes
     */
    public /*out*/ readonly autoCreateRoutes!: pulumi.Output<boolean>;
    /**
     * The CloudAMQP instance identifier.
     *
     * ***Depreacted: Changed from required to optional in v1.16.0, will be removed in next major version (v2.0)***
     */
    public readonly instanceId!: pulumi.Output<number | undefined>;
    /**
     * Network uri of the VPC network to which you will peer with.
     */
    public readonly peerNetworkUri!: pulumi.Output<string>;
    /**
     * VPC peering state
     */
    public /*out*/ readonly state!: pulumi.Output<string>;
    /**
     * VPC peering state details
     */
    public /*out*/ readonly stateDetails!: pulumi.Output<string>;
    /**
     * The managed VPC identifier.
     *
     * ***Note: Added as optional in version v1.16.0, will be required in next major version (v2.0)***
     */
    public readonly vpcId!: pulumi.Output<string | undefined>;

    /**
     * Create a VpcGcpPeering resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: VpcGcpPeeringArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: VpcGcpPeeringArgs | VpcGcpPeeringState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as VpcGcpPeeringState | undefined;
            resourceInputs["autoCreateRoutes"] = state ? state.autoCreateRoutes : undefined;
            resourceInputs["instanceId"] = state ? state.instanceId : undefined;
            resourceInputs["peerNetworkUri"] = state ? state.peerNetworkUri : undefined;
            resourceInputs["state"] = state ? state.state : undefined;
            resourceInputs["stateDetails"] = state ? state.stateDetails : undefined;
            resourceInputs["vpcId"] = state ? state.vpcId : undefined;
        } else {
            const args = argsOrState as VpcGcpPeeringArgs | undefined;
            if ((!args || args.peerNetworkUri === undefined) && !opts.urn) {
                throw new Error("Missing required property 'peerNetworkUri'");
            }
            resourceInputs["instanceId"] = args ? args.instanceId : undefined;
            resourceInputs["peerNetworkUri"] = args ? args.peerNetworkUri : undefined;
            resourceInputs["vpcId"] = args ? args.vpcId : undefined;
            resourceInputs["autoCreateRoutes"] = undefined /*out*/;
            resourceInputs["state"] = undefined /*out*/;
            resourceInputs["stateDetails"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(VpcGcpPeering.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering VpcGcpPeering resources.
 */
export interface VpcGcpPeeringState {
    /**
     * VPC peering auto created routes
     */
    autoCreateRoutes?: pulumi.Input<boolean>;
    /**
     * The CloudAMQP instance identifier.
     *
     * ***Depreacted: Changed from required to optional in v1.16.0, will be removed in next major version (v2.0)***
     */
    instanceId?: pulumi.Input<number>;
    /**
     * Network uri of the VPC network to which you will peer with.
     */
    peerNetworkUri?: pulumi.Input<string>;
    /**
     * VPC peering state
     */
    state?: pulumi.Input<string>;
    /**
     * VPC peering state details
     */
    stateDetails?: pulumi.Input<string>;
    /**
     * The managed VPC identifier.
     *
     * ***Note: Added as optional in version v1.16.0, will be required in next major version (v2.0)***
     */
    vpcId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a VpcGcpPeering resource.
 */
export interface VpcGcpPeeringArgs {
    /**
     * The CloudAMQP instance identifier.
     *
     * ***Depreacted: Changed from required to optional in v1.16.0, will be removed in next major version (v2.0)***
     */
    instanceId?: pulumi.Input<number>;
    /**
     * Network uri of the VPC network to which you will peer with.
     */
    peerNetworkUri: pulumi.Input<string>;
    /**
     * The managed VPC identifier.
     *
     * ***Note: Added as optional in version v1.16.0, will be required in next major version (v2.0)***
     */
    vpcId?: pulumi.Input<string>;
}
