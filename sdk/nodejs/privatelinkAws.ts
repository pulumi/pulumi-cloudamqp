// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Enable PrivateLink for a CloudAMQP instance hosted in AWS. If no existing VPC available when enable
 * PrivateLink, a new VPC will be created with subnet `10.52.72.0/24`.
 *
 * > **Note:** Enabling PrivateLink will automatically add firewall rules for the peered subnet.
 * <details>
 *  <summary>
 *     <i>Default PrivateLink firewall rule</i>
 *   </summary>
 *
 * ## Example Usage
 *
 * <details>
 *   <summary>
 *     <b>
 *       <i>CloudAMQP instance without existing VPC</i>
 *     </b>
 *   </summary>
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as cloudamqp from "@pulumi/cloudamqp";
 *
 * const instance = new cloudamqp.Instance("instance", {
 *     name: "Instance 01",
 *     plan: "bunny-1",
 *     region: "amazon-web-services::us-west-1",
 *     tags: [],
 * });
 * const privatelink = new cloudamqp.PrivatelinkAws("privatelink", {
 *     instanceId: instance.id,
 *     allowedPrincipals: ["arn:aws:iam::aws-account-id:user/user-name"],
 * });
 * ```
 *
 * </details>
 *
 * <details>
 *   <summary>
 *     <b>
 *       <i>CloudAMQP instance in an existing VPC</i>
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
 * const privatelink = new cloudamqp.PrivatelinkAws("privatelink", {
 *     instanceId: instance.id,
 *     allowedPrincipals: ["arn:aws:iam::aws-account-id:user/user-name"],
 * });
 * ```
 *
 * </details>
 *
 * ### With Additional Firewall Rules
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
 * const privatelink = new cloudamqp.PrivatelinkAws("privatelink", {
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
 *     dependsOn: [privatelink],
 * });
 * ```
 *
 * </details>
 *
 * ## Depedency
 *
 * This resource depends on CloudAMQP instance identifier, `cloudamqp_instance.instance.id`.
 *
 * ## Create PrivateLink with additional firewall rules
 *
 * To create a PrivateLink configuration with additional firewall rules, it's required to chain the cloudamqp.SecurityFirewall
 * resource to avoid parallel conflicting resource calls. You can do this by making the firewall
 * resource depend on the PrivateLink resource, `cloudamqp_privatelink_aws.privatelink`.
 *
 * Furthermore, since all firewall rules are overwritten, the otherwise automatically added rules for
 * the PrivateLink also needs to be added.
 *
 * ## Import
 *
 * `cloudamqp_privatelink_aws` can be imported using CloudAMQP internal identifier.
 *
 * ```sh
 * $ pulumi import cloudamqp:index/privatelinkAws:PrivatelinkAws privatelink <id>`
 * ```
 *
 * The resource uses the same identifier as the CloudAMQP instance. To retrieve the identifier for an instance, either use [CloudAMQP customer API](https://docs.cloudamqp.com/#list-instances) or use the data source [`cloudamqp_account`](./data-sources/account.md).
 */
export class PrivatelinkAws extends pulumi.CustomResource {
    /**
     * Get an existing PrivatelinkAws resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: PrivatelinkAwsState, opts?: pulumi.CustomResourceOptions): PrivatelinkAws {
        return new PrivatelinkAws(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'cloudamqp:index/privatelinkAws:PrivatelinkAws';

    /**
     * Returns true if the given object is an instance of PrivatelinkAws.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is PrivatelinkAws {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === PrivatelinkAws.__pulumiType;
    }

    /**
     * Covering availability zones used when creating an Endpoint from other VPC.
     */
    public /*out*/ readonly activeZones!: pulumi.Output<string[]>;
    /**
     * Allowed principals to access the endpoint service.
     */
    public readonly allowedPrincipals!: pulumi.Output<string[]>;
    /**
     * The CloudAMQP instance identifier.
     */
    public readonly instanceId!: pulumi.Output<number>;
    /**
     * Service name of the PrivateLink used when creating the endpoint from other VPC.
     */
    public /*out*/ readonly serviceName!: pulumi.Output<string>;
    /**
     * Configurable sleep time (seconds) when enable PrivateLink.
     * Default set to 10 seconds. *Available from v1.29.0*
     */
    public readonly sleep!: pulumi.Output<number | undefined>;
    /**
     * PrivateLink status [enable, pending, disable]
     */
    public /*out*/ readonly status!: pulumi.Output<string>;
    /**
     * Configurable timeout time (seconds) when enable PrivateLink.
     * Default set to 1800 seconds. *Available from v1.29.0*
     *
     * Allowed principals format: <br>
     * `arn:aws:iam::aws-account-id:root` <br>
     * `arn:aws:iam::aws-account-id:user/user-name` <br>
     * `arn:aws:iam::aws-account-id:role/role-name`
     */
    public readonly timeout!: pulumi.Output<number | undefined>;

    /**
     * Create a PrivatelinkAws resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: PrivatelinkAwsArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: PrivatelinkAwsArgs | PrivatelinkAwsState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as PrivatelinkAwsState | undefined;
            resourceInputs["activeZones"] = state ? state.activeZones : undefined;
            resourceInputs["allowedPrincipals"] = state ? state.allowedPrincipals : undefined;
            resourceInputs["instanceId"] = state ? state.instanceId : undefined;
            resourceInputs["serviceName"] = state ? state.serviceName : undefined;
            resourceInputs["sleep"] = state ? state.sleep : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["timeout"] = state ? state.timeout : undefined;
        } else {
            const args = argsOrState as PrivatelinkAwsArgs | undefined;
            if ((!args || args.allowedPrincipals === undefined) && !opts.urn) {
                throw new Error("Missing required property 'allowedPrincipals'");
            }
            if ((!args || args.instanceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'instanceId'");
            }
            resourceInputs["allowedPrincipals"] = args ? args.allowedPrincipals : undefined;
            resourceInputs["instanceId"] = args ? args.instanceId : undefined;
            resourceInputs["sleep"] = args ? args.sleep : undefined;
            resourceInputs["timeout"] = args ? args.timeout : undefined;
            resourceInputs["activeZones"] = undefined /*out*/;
            resourceInputs["serviceName"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(PrivatelinkAws.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering PrivatelinkAws resources.
 */
export interface PrivatelinkAwsState {
    /**
     * Covering availability zones used when creating an Endpoint from other VPC.
     */
    activeZones?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Allowed principals to access the endpoint service.
     */
    allowedPrincipals?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The CloudAMQP instance identifier.
     */
    instanceId?: pulumi.Input<number>;
    /**
     * Service name of the PrivateLink used when creating the endpoint from other VPC.
     */
    serviceName?: pulumi.Input<string>;
    /**
     * Configurable sleep time (seconds) when enable PrivateLink.
     * Default set to 10 seconds. *Available from v1.29.0*
     */
    sleep?: pulumi.Input<number>;
    /**
     * PrivateLink status [enable, pending, disable]
     */
    status?: pulumi.Input<string>;
    /**
     * Configurable timeout time (seconds) when enable PrivateLink.
     * Default set to 1800 seconds. *Available from v1.29.0*
     *
     * Allowed principals format: <br>
     * `arn:aws:iam::aws-account-id:root` <br>
     * `arn:aws:iam::aws-account-id:user/user-name` <br>
     * `arn:aws:iam::aws-account-id:role/role-name`
     */
    timeout?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a PrivatelinkAws resource.
 */
export interface PrivatelinkAwsArgs {
    /**
     * Allowed principals to access the endpoint service.
     */
    allowedPrincipals: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The CloudAMQP instance identifier.
     */
    instanceId: pulumi.Input<number>;
    /**
     * Configurable sleep time (seconds) when enable PrivateLink.
     * Default set to 10 seconds. *Available from v1.29.0*
     */
    sleep?: pulumi.Input<number>;
    /**
     * Configurable timeout time (seconds) when enable PrivateLink.
     * Default set to 1800 seconds. *Available from v1.29.0*
     *
     * Allowed principals format: <br>
     * `arn:aws:iam::aws-account-id:root` <br>
     * `arn:aws:iam::aws-account-id:user/user-name` <br>
     * `arn:aws:iam::aws-account-id:role/role-name`
     */
    timeout?: pulumi.Input<number>;
}
