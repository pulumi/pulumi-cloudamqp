// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * This resource allows you to install or uninstall community plugins. Once installed the plugin will be available in `cloudamqp.Plugin`.
 *
 * Only available for dedicated subscription plans.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as cloudamqp from "@pulumi/cloudamqp";
 *
 * const rabbitmqDelayedMessageExchange = new cloudamqp.PluginCommunity("rabbitmqDelayedMessageExchange", {
 *     instanceId: cloudamqp_instance.instance_01.id,
 *     enabled: true,
 * });
 * ```
 * ## Depedency
 *
 * This resource depends on CloudAMQP instance identifier, `cloudamqp_instance.instance.id`.
 *
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
     * Enable or disable the plugins.
     */
    public readonly enabled!: pulumi.Output<boolean>;
    /**
     * The CloudAMQP instance ID.
     */
    public readonly instanceId!: pulumi.Output<number>;
    /**
     * The name of the Rabbit MQ plugin.
     */
    public readonly name!: pulumi.Output<string>;

    /**
     * Create a PluginCommunity resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: PluginCommunityArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: PluginCommunityArgs | PluginCommunityState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as PluginCommunityState | undefined;
            inputs["enabled"] = state ? state.enabled : undefined;
            inputs["instanceId"] = state ? state.instanceId : undefined;
            inputs["name"] = state ? state.name : undefined;
        } else {
            const args = argsOrState as PluginCommunityArgs | undefined;
            if ((!args || args.enabled === undefined) && !opts.urn) {
                throw new Error("Missing required property 'enabled'");
            }
            if ((!args || args.instanceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'instanceId'");
            }
            inputs["enabled"] = args ? args.enabled : undefined;
            inputs["instanceId"] = args ? args.instanceId : undefined;
            inputs["name"] = args ? args.name : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(PluginCommunity.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering PluginCommunity resources.
 */
export interface PluginCommunityState {
    /**
     * Enable or disable the plugins.
     */
    enabled?: pulumi.Input<boolean>;
    /**
     * The CloudAMQP instance ID.
     */
    instanceId?: pulumi.Input<number>;
    /**
     * The name of the Rabbit MQ plugin.
     */
    name?: pulumi.Input<string>;
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
     * The name of the Rabbit MQ plugin.
     */
    name?: pulumi.Input<string>;
}
