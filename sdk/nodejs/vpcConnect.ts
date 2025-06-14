// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * ## Import
 *
 * `cloudamqp_vpc_connect` can be imported using CloudAMQP instance identifier. To
 *
 * retrieve the identifier, use [CloudAMQP API list intances].
 *
 * From Terraform v1.5.0, the `import` block can be used to import this resource:
 *
 * hcl
 *
 * import {
 *
 *   to = cloudamqp_vpc_connect.this
 *
 *   id = cloudamqp_instance.instance.id
 *
 * }
 *
 * Or use Terraform CLI:
 *
 * ```sh
 * $ pulumi import cloudamqp:index/vpcConnect:VpcConnect vpc_connect <id>`
 * ```
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
     * Covering availability zones used when creating an endpoint from other VPC. (AWS)
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
     * List of approved subscriptions used by Azure, see below
     * table.
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
     * Service name (alias for Azure, see example above) of the PrivateLink.
     */
    public /*out*/ readonly serviceName!: pulumi.Output<string>;
    /**
     * Configurable sleep time (seconds) when enable Private
     * Service Connect. Default set to 10 seconds.
     */
    public readonly sleep!: pulumi.Output<number | undefined>;
    /**
     * Private Service Connect status [enable, pending, disable]
     */
    public /*out*/ readonly status!: pulumi.Output<string>;
    /**
     * Configurable timeout time (seconds) when enable Private
     * Service Connect. Default set to 1800 seconds.
     *
     * ___
     *
     * The `allowedPrincipals`, `approvedSubscriptions` or `allowedProjects` data depends on the
     * provider platform:
     *
     * | Platform | Description | Format |
     * |---|---|---|
     * | AWS | IAM ARN principals | arn:aws:iam::aws-account-id:root<br>arn:aws:iam::aws-account-id:user/user-name<br> arn:aws:iam::aws-account-id:role/role-name |
     * | Azure | Subscription (GUID) | XXXXXXXX-XXXX-XXXX-XXXX-XXXXXXXXXXXX |
     * | GCP | Project IDs [Google docs] | 6 to 30 lowercase letters, digits, or hyphens |
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
     * Covering availability zones used when creating an endpoint from other VPC. (AWS)
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
     * List of approved subscriptions used by Azure, see below
     * table.
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
     * Service name (alias for Azure, see example above) of the PrivateLink.
     */
    serviceName?: pulumi.Input<string>;
    /**
     * Configurable sleep time (seconds) when enable Private
     * Service Connect. Default set to 10 seconds.
     */
    sleep?: pulumi.Input<number>;
    /**
     * Private Service Connect status [enable, pending, disable]
     */
    status?: pulumi.Input<string>;
    /**
     * Configurable timeout time (seconds) when enable Private
     * Service Connect. Default set to 1800 seconds.
     *
     * ___
     *
     * The `allowedPrincipals`, `approvedSubscriptions` or `allowedProjects` data depends on the
     * provider platform:
     *
     * | Platform | Description | Format |
     * |---|---|---|
     * | AWS | IAM ARN principals | arn:aws:iam::aws-account-id:root<br>arn:aws:iam::aws-account-id:user/user-name<br> arn:aws:iam::aws-account-id:role/role-name |
     * | Azure | Subscription (GUID) | XXXXXXXX-XXXX-XXXX-XXXX-XXXXXXXXXXXX |
     * | GCP | Project IDs [Google docs] | 6 to 30 lowercase letters, digits, or hyphens |
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
     * List of approved subscriptions used by Azure, see below
     * table.
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
     * Configurable sleep time (seconds) when enable Private
     * Service Connect. Default set to 10 seconds.
     */
    sleep?: pulumi.Input<number>;
    /**
     * Configurable timeout time (seconds) when enable Private
     * Service Connect. Default set to 1800 seconds.
     *
     * ___
     *
     * The `allowedPrincipals`, `approvedSubscriptions` or `allowedProjects` data depends on the
     * provider platform:
     *
     * | Platform | Description | Format |
     * |---|---|---|
     * | AWS | IAM ARN principals | arn:aws:iam::aws-account-id:root<br>arn:aws:iam::aws-account-id:user/user-name<br> arn:aws:iam::aws-account-id:role/role-name |
     * | Azure | Subscription (GUID) | XXXXXXXX-XXXX-XXXX-XXXX-XXXXXXXXXXXX |
     * | GCP | Project IDs [Google docs] | 6 to 30 lowercase letters, digits, or hyphens |
     */
    timeout?: pulumi.Input<number>;
}
