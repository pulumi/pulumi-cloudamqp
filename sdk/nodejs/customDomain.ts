// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * This resource allows you to configure and manage your custom domain for the CloudAMQP instance.
 *
 * Adding a custom domain to your instance will generate a TLS certificate from [Let's Encrypt], for the given hostname, and install it on all servers in your cluster. The certificate will be automatically renewed going forward.
 *
 * > **WARNING:** Please note that when creating, changing or deleting the custom domain, the listeners on your servers will be restarted in order to apply the changes. This will close your current connections.
 *
 * Your custom domain name needs to point to your CloudAMQP hostname, preferably using a CNAME DNS record.
 *
 * Only available for dedicated subscription plans.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as cloudamqp from "@pulumi/cloudamqp";
 *
 * const settings = new cloudamqp.CustomDomain("settings", {
 *     instanceId: cloudamqp_instance.instance.id,
 *     hostname: "myname.mydomain",
 * });
 * ```
 * ## Depedency
 *
 * This resource depends on CloudAMQP instance identifier, `cloudamqp_instance.instance.id`.
 *
 * ## Import
 *
 * `cloudamqp_custom_domain` can be imported using CloudAMQP instance identifier.
 *
 * ```sh
 * $ pulumi import cloudamqp:index/customDomain:CustomDomain settings <instance_id>`
 * ```
 *
 * [Let's Encrypt]: https://letsencrypt.org/
 */
export class CustomDomain extends pulumi.CustomResource {
    /**
     * Get an existing CustomDomain resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: CustomDomainState, opts?: pulumi.CustomResourceOptions): CustomDomain {
        return new CustomDomain(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'cloudamqp:index/customDomain:CustomDomain';

    /**
     * Returns true if the given object is an instance of CustomDomain.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is CustomDomain {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === CustomDomain.__pulumiType;
    }

    /**
     * Your custom domain name.
     */
    public readonly hostname!: pulumi.Output<string>;
    /**
     * The CloudAMQP instance ID.
     */
    public readonly instanceId!: pulumi.Output<number>;

    /**
     * Create a CustomDomain resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: CustomDomainArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: CustomDomainArgs | CustomDomainState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as CustomDomainState | undefined;
            resourceInputs["hostname"] = state ? state.hostname : undefined;
            resourceInputs["instanceId"] = state ? state.instanceId : undefined;
        } else {
            const args = argsOrState as CustomDomainArgs | undefined;
            if ((!args || args.hostname === undefined) && !opts.urn) {
                throw new Error("Missing required property 'hostname'");
            }
            if ((!args || args.instanceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'instanceId'");
            }
            resourceInputs["hostname"] = args ? args.hostname : undefined;
            resourceInputs["instanceId"] = args ? args.instanceId : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(CustomDomain.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering CustomDomain resources.
 */
export interface CustomDomainState {
    /**
     * Your custom domain name.
     */
    hostname?: pulumi.Input<string>;
    /**
     * The CloudAMQP instance ID.
     */
    instanceId?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a CustomDomain resource.
 */
export interface CustomDomainArgs {
    /**
     * Your custom domain name.
     */
    hostname: pulumi.Input<string>;
    /**
     * The CloudAMQP instance ID.
     */
    instanceId: pulumi.Input<number>;
}
