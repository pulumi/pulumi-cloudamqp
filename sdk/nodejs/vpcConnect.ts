// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * This resource is a generic way to handle PrivateLink (AWS and Azure) and Private Service Connect (GCP).
 * Communication between resources can be done just as they were living inside a VPC. CloudAMQP creates an Endpoint
 * Service to connect the VPC and creating a new network interface to handle the communicate.
 *
 * If no existing VPC available when enable VPC connect, a new VPC will be created with subnet `10.52.72.0/24`.
 *
 * More information can be found at: [CloudAMQP VPC Connect](https://www.cloudamqp.com/docs/cloudamqp-vpc-connect.html)
 *
 * > **Note:** Enabling VPC Connect will automatically add a firewall rule.
 *
 * <details>
 *  <summary>
 *     <b>
 *       <i>Default PrivateLink firewall rule [AWS, Azure]</i>
 *     </b>
 *   </summary>
 *
 * ## Example Usage
 *
 * <details>
 *   <summary>
 *     <b>
 *       <i>Enable VPC Connect (PrivateLink) in AWS</i>
 *     </b>
 *   </summary>
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as cloudamqp from "@pulumi/cloudamqp";
 *
 * const vpc = new cloudamqp.Vpc("vpc", {
 *     name: "Standalone VPC",
 *     region: "amazon-web-services::us-west-1",
 *     subnet: "10.56.72.0/24",
 *     tags: [],
 * });
 * const instance = new cloudamqp.Instance("instance", {
 *     name: "Instance 01",
 *     plan: "bunny-1",
 *     region: "amazon-web-services::us-west-1",
 *     tags: [],
 *     vpcId: vpc.id,
 *     keepAssociatedVpc: true,
 * });
 * const vpcConnect = new cloudamqp.VpcConnect("vpc_connect", {
 *     instanceId: instance.id,
 *     region: instance.region,
 *     allowedPrincipals: ["arn:aws:iam::aws-account-id:user/user-name"],
 * });
 * ```
 *
 * </details>
 *
 * <details>
 *   <summary>
 *     <b>
 *       <i>Enable VPC Connect (PrivateLink) in Azure</i>
 *     </b>
 *   </summary>
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as cloudamqp from "@pulumi/cloudamqp";
 *
 * const vpc = new cloudamqp.Vpc("vpc", {
 *     name: "Standalone VPC",
 *     region: "azure-arm::westus",
 *     subnet: "10.56.72.0/24",
 *     tags: [],
 * });
 * const instance = new cloudamqp.Instance("instance", {
 *     name: "Instance 01",
 *     plan: "bunny-1",
 *     region: "azure-arm::westus",
 *     tags: [],
 *     vpcId: vpc.id,
 *     keepAssociatedVpc: true,
 * });
 * const vpcConnect = new cloudamqp.VpcConnect("vpc_connect", {
 *     instanceId: instance.id,
 *     region: instance.region,
 *     approvedSubscriptions: ["XXXXXXXX-XXXX-XXXX-XXXX-XXXXXXXXXXXX"],
 * });
 * ```
 *
 * </details>
 *
 * <details>
 *   <summary>
 *     <b>
 *       <i>Enable VPC Connect (Private Service Connect) in GCP</i>
 *     </b>
 *   </summary>
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as cloudamqp from "@pulumi/cloudamqp";
 *
 * const vpc = new cloudamqp.Vpc("vpc", {
 *     name: "Standalone VPC",
 *     region: "google-compute-engine::us-west1",
 *     subnet: "10.56.72.0/24",
 *     tags: [],
 * });
 * const instance = new cloudamqp.Instance("instance", {
 *     name: "Instance 01",
 *     plan: "bunny-1",
 *     region: "google-compute-engine::us-west1",
 *     tags: [],
 *     vpcId: vpc.id,
 *     keepAssociatedVpc: true,
 * });
 * const vpcConnect = new cloudamqp.VpcConnect("vpc_connect", {
 *     instanceId: instance.id,
 *     region: instance.region,
 *     allowedProjects: ["some-project-123456"],
 * });
 * ```
 *
 * </details>
 *
 * ### with additional firewall rules
 *
 * <details>
 *   <summary>
 *     <b>
 *       <i>CloudAMQP instance in an existing VPC with managed firewall rules</i>
 *     </b>
 *   </summary>
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as cloudamqp from "@pulumi/cloudamqp";
 *
 * const vpc = new cloudamqp.Vpc("vpc", {
 *     name: "Standalone VPC",
 *     region: "amazon-web-services::us-west-1",
 *     subnet: "10.56.72.0/24",
 *     tags: [],
 * });
 * const instance = new cloudamqp.Instance("instance", {
 *     name: "Instance 01",
 *     plan: "bunny-1",
 *     region: "amazon-web-services::us-west-1",
 *     tags: [],
 *     vpcId: vpc.id,
 *     keepAssociatedVpc: true,
 * });
 * const vpcConnect = new cloudamqp.VpcConnect("vpc_connect", {
 *     instanceId: instance.id,
 *     allowedPrincipals: ["arn:aws:iam::aws-account-id:user/user-name"],
 * });
 * const firewallSettings = new cloudamqp.SecurityFirewall("firewall_settings", {
 *     instanceId: instance.id,
 *     rules: [
 *         {
 *             description: "Custom PrivateLink setup",
 *             ip: vpc.subnet,
 *             ports: [],
 *             services: [
 *                 "AMQP",
 *                 "AMQPS",
 *                 "HTTPS",
 *                 "STREAM",
 *                 "STREAM_SSL",
 *             ],
 *         },
 *         {
 *             description: "MGMT interface",
 *             ip: "0.0.0.0/0",
 *             ports: [],
 *             services: ["HTTPS"],
 *         },
 *     ],
 * }, {
 *     dependsOn: [vpcConnect],
 * });
 * ```
 *
 * </details>
 */
export class VpcConnect extends pulumi.CustomResource {
    /**
     * Get an existing VpcConnect resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: VpcConnectState, opts?: pulumi.CustomResourceOptions): VpcConnect {
        return new VpcConnect(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'cloudamqp:index/vpcConnect:VpcConnect';

    /**
     * Returns true if the given object is an instance of VpcConnect.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is VpcConnect {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === VpcConnect.__pulumiType;
    }

    /**
     * Covering availability zones used when creating an endpoint from other VPC. [AWS]
     */
    public /*out*/ readonly activeZones!: pulumi.Output<string[]>;
    /**
     * List of allowed prinicpals used by AWS, see below table.
     */
    public readonly allowedPrincipals!: pulumi.Output<string[] | undefined>;
    /**
     * List of allowed projects used by GCP, see below table.
     */
    public readonly allowedProjects!: pulumi.Output<string[] | undefined>;
    /**
     * List of approved subscriptions used by Azure, see below table.
     */
    public readonly approvedSubscriptions!: pulumi.Output<string[] | undefined>;
    /**
     * The CloudAMQP instance identifier.
     */
    public readonly instanceId!: pulumi.Output<number>;
    /**
     * The region where the CloudAMQP instance is hosted.
     */
    public readonly region!: pulumi.Output<string>;
    /**
     * Service name (alias for Azure) of the PrivateLink.
     */
    public /*out*/ readonly serviceName!: pulumi.Output<string>;
    /**
     * Configurable sleep time (seconds) when enable Private Service Connect.
     * Default set to 10 seconds.
     */
    public readonly sleep!: pulumi.Output<number | undefined>;
    /**
     * Status of the Private Service Connect [enabled, pending, disabled]
     */
    public /*out*/ readonly status!: pulumi.Output<string>;
    /**
     * Configurable timeout time (seconds) when enable Private Service Connect.
     * Default set to 1800 seconds.
     */
    public readonly timeout!: pulumi.Output<number | undefined>;

    /**
     * Create a VpcConnect resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: VpcConnectArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: VpcConnectArgs | VpcConnectState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as VpcConnectState | undefined;
            resourceInputs["activeZones"] = state ? state.activeZones : undefined;
            resourceInputs["allowedPrincipals"] = state ? state.allowedPrincipals : undefined;
            resourceInputs["allowedProjects"] = state ? state.allowedProjects : undefined;
            resourceInputs["approvedSubscriptions"] = state ? state.approvedSubscriptions : undefined;
            resourceInputs["instanceId"] = state ? state.instanceId : undefined;
            resourceInputs["region"] = state ? state.region : undefined;
            resourceInputs["serviceName"] = state ? state.serviceName : undefined;
            resourceInputs["sleep"] = state ? state.sleep : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["timeout"] = state ? state.timeout : undefined;
        } else {
            const args = argsOrState as VpcConnectArgs | undefined;
            if ((!args || args.instanceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'instanceId'");
            }
            if ((!args || args.region === undefined) && !opts.urn) {
                throw new Error("Missing required property 'region'");
            }
            resourceInputs["allowedPrincipals"] = args ? args.allowedPrincipals : undefined;
            resourceInputs["allowedProjects"] = args ? args.allowedProjects : undefined;
            resourceInputs["approvedSubscriptions"] = args ? args.approvedSubscriptions : undefined;
            resourceInputs["instanceId"] = args ? args.instanceId : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
            resourceInputs["sleep"] = args ? args.sleep : undefined;
            resourceInputs["timeout"] = args ? args.timeout : undefined;
            resourceInputs["activeZones"] = undefined /*out*/;
            resourceInputs["serviceName"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(VpcConnect.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering VpcConnect resources.
 */
export interface VpcConnectState {
    /**
     * Covering availability zones used when creating an endpoint from other VPC. [AWS]
     */
    activeZones?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * List of allowed prinicpals used by AWS, see below table.
     */
    allowedPrincipals?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * List of allowed projects used by GCP, see below table.
     */
    allowedProjects?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * List of approved subscriptions used by Azure, see below table.
     */
    approvedSubscriptions?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The CloudAMQP instance identifier.
     */
    instanceId?: pulumi.Input<number>;
    /**
     * The region where the CloudAMQP instance is hosted.
     */
    region?: pulumi.Input<string>;
    /**
     * Service name (alias for Azure) of the PrivateLink.
     */
    serviceName?: pulumi.Input<string>;
    /**
     * Configurable sleep time (seconds) when enable Private Service Connect.
     * Default set to 10 seconds.
     */
    sleep?: pulumi.Input<number>;
    /**
     * Status of the Private Service Connect [enabled, pending, disabled]
     */
    status?: pulumi.Input<string>;
    /**
     * Configurable timeout time (seconds) when enable Private Service Connect.
     * Default set to 1800 seconds.
     */
    timeout?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a VpcConnect resource.
 */
export interface VpcConnectArgs {
    /**
     * List of allowed prinicpals used by AWS, see below table.
     */
    allowedPrincipals?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * List of allowed projects used by GCP, see below table.
     */
    allowedProjects?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * List of approved subscriptions used by Azure, see below table.
     */
    approvedSubscriptions?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The CloudAMQP instance identifier.
     */
    instanceId: pulumi.Input<number>;
    /**
     * The region where the CloudAMQP instance is hosted.
     */
    region: pulumi.Input<string>;
    /**
     * Configurable sleep time (seconds) when enable Private Service Connect.
     * Default set to 10 seconds.
     */
    sleep?: pulumi.Input<number>;
    /**
     * Configurable timeout time (seconds) when enable Private Service Connect.
     * Default set to 1800 seconds.
     */
    timeout?: pulumi.Input<number>;
}
