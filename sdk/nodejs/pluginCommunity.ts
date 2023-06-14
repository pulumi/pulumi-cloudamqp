// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * ## Import
 *
 * `cloudamqp_plugin` can be imported using the name argument of the resource together with CloudAMQP instance identifier. The name and identifier are CSV separated, see example below.
 *
 * ```sh
 *  $ pulumi import cloudamqp:index/pluginCommunity:PluginCommunity <resource_name> <plugin_name>,<instance_id>`
 * ```
 */
export class PluginCommunity extends pulumi.CustomResource {
    /**
     * Get an existing PluginCommunity resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: PluginCommunityState, opts?: pulumi.CustomResourceOptions): PluginCommunity {
        return new PluginCommunity(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'cloudamqp:index/pluginCommunity:PluginCommunity';

    /**
     * Returns true if the given object is an instance of PluginCommunity.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is PluginCommunity {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === PluginCommunity.__pulumiType;
    }

    /**
     * The description of the plugin.
     */
    public /*out*/ readonly description!: pulumi.Output<string>;
    /**
     * Enable or disable the plugins.
     */
    public readonly enabled!: pulumi.Output<boolean>;
    /**
     * The CloudAMQP instance ID.
     */
    public readonly instanceId!: pulumi.Output<number>;
    /**
     * The name of the Rabbit MQ community plugin.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Required version of RabbitMQ.
     */
    public /*out*/ readonly require!: pulumi.Output<string>;

    /**
     * Create a PluginCommunity resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: PluginCommunityArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: PluginCommunityArgs | PluginCommunityState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as PluginCommunityState | undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["enabled"] = state ? state.enabled : undefined;
            resourceInputs["instanceId"] = state ? state.instanceId : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["require"] = state ? state.require : undefined;
        } else {
            const args = argsOrState as PluginCommunityArgs | undefined;
            if ((!args || args.enabled === undefined) && !opts.urn) {
                throw new Error("Missing required property 'enabled'");
            }
            if ((!args || args.instanceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'instanceId'");
            }
            resourceInputs["enabled"] = args ? args.enabled : undefined;
            resourceInputs["instanceId"] = args ? args.instanceId : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["description"] = undefined /*out*/;
            resourceInputs["require"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(PluginCommunity.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering PluginCommunity resources.
 */
export interface PluginCommunityState {
    /**
     * The description of the plugin.
     */
    description?: pulumi.Input<string>;
    /**
     * Enable or disable the plugins.
     */
    enabled?: pulumi.Input<boolean>;
    /**
     * The CloudAMQP instance ID.
     */
    instanceId?: pulumi.Input<number>;
    /**
     * The name of the Rabbit MQ community plugin.
     */
    name?: pulumi.Input<string>;
    /**
     * Required version of RabbitMQ.
     */
    require?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a PluginCommunity resource.
 */
export interface PluginCommunityArgs {
    /**
     * Enable or disable the plugins.
     */
    enabled: pulumi.Input<boolean>;
    /**
     * The CloudAMQP instance ID.
     */
    instanceId: pulumi.Input<number>;
    /**
     * The name of the Rabbit MQ community plugin.
     */
    name?: pulumi.Input<string>;
}
