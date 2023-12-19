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
 *  $ pulumi import cloudamqp:index/plugin:Plugin rabbitmq_management rabbitmq_management,<instance_id>`
 * ```
 */
export class Plugin extends pulumi.CustomResource {
    /**
     * Get an existing Plugin resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: PluginState, opts?: pulumi.CustomResourceOptions): Plugin {
        return new Plugin(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'cloudamqp:index/plugin:Plugin';

    /**
     * Returns true if the given object is an instance of Plugin.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Plugin {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Plugin.__pulumiType;
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
     * The name of the Rabbit MQ plugin.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Configurable sleep time (seconds) for retries when requesting information
     * about plugins. Default set to 10 seconds. *Available from v1.29.0*
     */
    public readonly sleep!: pulumi.Output<number | undefined>;
    /**
     * Configurable timeout time (seconds) for retries when requesting
     * information about plugins. Default set to 1800 seconds. *Available from v1.29.0*
     */
    public readonly timeout!: pulumi.Output<number | undefined>;
    /**
     * The version of the plugin.
     */
    public /*out*/ readonly version!: pulumi.Output<string>;

    /**
     * Create a Plugin resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: PluginArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: PluginArgs | PluginState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as PluginState | undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["enabled"] = state ? state.enabled : undefined;
            resourceInputs["instanceId"] = state ? state.instanceId : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["sleep"] = state ? state.sleep : undefined;
            resourceInputs["timeout"] = state ? state.timeout : undefined;
            resourceInputs["version"] = state ? state.version : undefined;
        } else {
            const args = argsOrState as PluginArgs | undefined;
            if ((!args || args.enabled === undefined) && !opts.urn) {
                throw new Error("Missing required property 'enabled'");
            }
            if ((!args || args.instanceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'instanceId'");
            }
            resourceInputs["enabled"] = args ? args.enabled : undefined;
            resourceInputs["instanceId"] = args ? args.instanceId : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["sleep"] = args ? args.sleep : undefined;
            resourceInputs["timeout"] = args ? args.timeout : undefined;
            resourceInputs["description"] = undefined /*out*/;
            resourceInputs["version"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Plugin.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Plugin resources.
 */
export interface PluginState {
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
     * The name of the Rabbit MQ plugin.
     */
    name?: pulumi.Input<string>;
    /**
     * Configurable sleep time (seconds) for retries when requesting information
     * about plugins. Default set to 10 seconds. *Available from v1.29.0*
     */
    sleep?: pulumi.Input<number>;
    /**
     * Configurable timeout time (seconds) for retries when requesting
     * information about plugins. Default set to 1800 seconds. *Available from v1.29.0*
     */
    timeout?: pulumi.Input<number>;
    /**
     * The version of the plugin.
     */
    version?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Plugin resource.
 */
export interface PluginArgs {
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
    /**
     * Configurable sleep time (seconds) for retries when requesting information
     * about plugins. Default set to 10 seconds. *Available from v1.29.0*
     */
    sleep?: pulumi.Input<number>;
    /**
     * Configurable timeout time (seconds) for retries when requesting
     * information about plugins. Default set to 1800 seconds. *Available from v1.29.0*
     */
    timeout?: pulumi.Input<number>;
}
