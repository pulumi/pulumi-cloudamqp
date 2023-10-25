// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Enable PrivateLink for a CloudAMQP instance hosted in Azure. If no existing VPC available when enable PrivateLink, a new VPC will be created with subnet `10.52.72.0/24`.
 *
 * > **Note:** Enabling PrivateLink will automatically add firewall rules for the peered subnet.
 * <details>
 *  <summary>
 *     <i>Default PrivateLink firewall rule</i>
 *   </summary>
 * </details>
 *
 * Pricing is available at [cloudamqp.com](https://www.cloudamqp.com/plans.html) where you can also find more information about [CloudAMQP PrivateLink](https://www.cloudamqp.com/docs/cloudamqp-privatelink.html#azure-privatelink).
 *
 * Only available for dedicated subscription plans.
 *
 * ## Depedency
 *
 * This resource depends on CloudAMQP instance identifier, `cloudamqp_instance.instance.id`.
 *
 * ## Create PrivateLink with additional firewall rules
 *
 * To create a PrivateLink configuration with additional firewall rules, it's required to chain the cloudamqp.SecurityFirewall
 * resource to avoid parallel conflicting resource calls. You can do this by making the firewall resource depend on the PrivateLink resource, `cloudamqp_privatelink_azure.privatelink`.
 *
 * Furthermore, since all firewall rules are overwritten, the otherwise automatically added rules for the PrivateLink also needs to be added.
 *
 * ## Import
 *
 * `cloudamqp_privatelink_aws` can be imported using CloudAMQP internal identifier.
 *
 * ```sh
 *  $ pulumi import cloudamqp:index/privatelinkAzure:PrivatelinkAzure privatelink <id>`
 * ```
 */
export class PrivatelinkAzure extends pulumi.CustomResource {
    /**
     * Get an existing PrivatelinkAzure resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: PrivatelinkAzureState, opts?: pulumi.CustomResourceOptions): PrivatelinkAzure {
        return new PrivatelinkAzure(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'cloudamqp:index/privatelinkAzure:PrivatelinkAzure';

    /**
     * Returns true if the given object is an instance of PrivatelinkAzure.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is PrivatelinkAzure {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === PrivatelinkAzure.__pulumiType;
    }

    /**
     * Approved subscriptions to access the endpoint service. See format below.
     */
    public readonly approvedSubscriptions!: pulumi.Output<string[]>;
    /**
     * The CloudAMQP instance identifier.
     */
    public readonly instanceId!: pulumi.Output<number>;
    /**
     * Name of the server having the PrivateLink enabled.
     */
    public /*out*/ readonly serverName!: pulumi.Output<string>;
    /**
     * Service name (alias) of the PrivateLink, needed when creating the endpoint.
     */
    public /*out*/ readonly serviceName!: pulumi.Output<string>;
    /**
     * Configurable sleep time (seconds) when enable PrivateLink. Default set to 60 seconds.
     */
    public readonly sleep!: pulumi.Output<number | undefined>;
    /**
     * PrivateLink status [enable, pending, disable]
     */
    public /*out*/ readonly status!: pulumi.Output<string>;
    /**
     * Configurable timeout time (seconds) when enable PrivateLink. Default set to 3600 seconds.
     *
     * Approved subscriptions format: <br>
     * `XXXXXXXX-XXXX-XXXX-XXXX-XXXXXXXXXXXX`
     */
    public readonly timeout!: pulumi.Output<number | undefined>;

    /**
     * Create a PrivatelinkAzure resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: PrivatelinkAzureArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: PrivatelinkAzureArgs | PrivatelinkAzureState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as PrivatelinkAzureState | undefined;
            resourceInputs["approvedSubscriptions"] = state ? state.approvedSubscriptions : undefined;
            resourceInputs["instanceId"] = state ? state.instanceId : undefined;
            resourceInputs["serverName"] = state ? state.serverName : undefined;
            resourceInputs["serviceName"] = state ? state.serviceName : undefined;
            resourceInputs["sleep"] = state ? state.sleep : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["timeout"] = state ? state.timeout : undefined;
        } else {
            const args = argsOrState as PrivatelinkAzureArgs | undefined;
            if ((!args || args.approvedSubscriptions === undefined) && !opts.urn) {
                throw new Error("Missing required property 'approvedSubscriptions'");
            }
            if ((!args || args.instanceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'instanceId'");
            }
            resourceInputs["approvedSubscriptions"] = args ? args.approvedSubscriptions : undefined;
            resourceInputs["instanceId"] = args ? args.instanceId : undefined;
            resourceInputs["sleep"] = args ? args.sleep : undefined;
            resourceInputs["timeout"] = args ? args.timeout : undefined;
            resourceInputs["serverName"] = undefined /*out*/;
            resourceInputs["serviceName"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(PrivatelinkAzure.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering PrivatelinkAzure resources.
 */
export interface PrivatelinkAzureState {
    /**
     * Approved subscriptions to access the endpoint service. See format below.
     */
    approvedSubscriptions?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The CloudAMQP instance identifier.
     */
    instanceId?: pulumi.Input<number>;
    /**
     * Name of the server having the PrivateLink enabled.
     */
    serverName?: pulumi.Input<string>;
    /**
     * Service name (alias) of the PrivateLink, needed when creating the endpoint.
     */
    serviceName?: pulumi.Input<string>;
    /**
     * Configurable sleep time (seconds) when enable PrivateLink. Default set to 60 seconds.
     */
    sleep?: pulumi.Input<number>;
    /**
     * PrivateLink status [enable, pending, disable]
     */
    status?: pulumi.Input<string>;
    /**
     * Configurable timeout time (seconds) when enable PrivateLink. Default set to 3600 seconds.
     *
     * Approved subscriptions format: <br>
     * `XXXXXXXX-XXXX-XXXX-XXXX-XXXXXXXXXXXX`
     */
    timeout?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a PrivatelinkAzure resource.
 */
export interface PrivatelinkAzureArgs {
    /**
     * Approved subscriptions to access the endpoint service. See format below.
     */
    approvedSubscriptions: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The CloudAMQP instance identifier.
     */
    instanceId: pulumi.Input<number>;
    /**
     * Configurable sleep time (seconds) when enable PrivateLink. Default set to 60 seconds.
     */
    sleep?: pulumi.Input<number>;
    /**
     * Configurable timeout time (seconds) when enable PrivateLink. Default set to 3600 seconds.
     *
     * Approved subscriptions format: <br>
     * `XXXXXXXX-XXXX-XXXX-XXXX-XXXXXXXXXXXX`
     */
    timeout?: pulumi.Input<number>;
}
