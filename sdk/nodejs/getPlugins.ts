// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * Use this data source to retrieve information about installed and available plugins for the CloudAMQP instance.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as cloudamqp from "@pulumi/cloudamqp";
 *
 * const plugins = cloudamqp.getPlugins({
 *     instanceId: instance.id,
 * });
 * ```
 *
 * ## Attributes reference
 *
 * All attributes reference are computed
 *
 * * `id`      - The identifier for this resource.
 * * `plugins` - An array of plugins. Each `plugins` block consists of the fields documented below.
 * * `sleep` - (Optional) Configurable sleep time (seconds) for retries when requesting information
 *   about plugins. Default set to 10 seconds. *Available from v1.29.0*
 * * `timeout` - (Optional) - Configurable timeout time (seconds) for retries when requesting
 *   information about plugins. Default set to 1800 seconds. *Available from v1.29.0*
 *
 * ***
 *
 * The `plugins` block consist of
 *
 * * `name`        - The type of the recipient.
 * * `version`     - Rabbit MQ version that the plugins are shipped with.
 * * `description` - Description of what the plugin does.
 * * `enabled`     - Enable or disable information for the plugin.
 *
 * ## Dependency
 *
 * This data source depends on CloudAMQP instance identifier, `cloudamqp_instance.instance.id`.
 */
export function getPlugins(args: GetPluginsArgs, opts?: pulumi.InvokeOptions): Promise<GetPluginsResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("cloudamqp:index/getPlugins:getPlugins", {
        "instanceId": args.instanceId,
    }, opts);
}

/**
 * A collection of arguments for invoking getPlugins.
 */
export interface GetPluginsArgs {
    /**
     * The CloudAMQP instance identifier.
     */
    instanceId: number;
}

/**
 * A collection of values returned by getPlugins.
 */
export interface GetPluginsResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly instanceId: number;
    readonly plugins: outputs.GetPluginsPlugin[];
}
/**
 * Use this data source to retrieve information about installed and available plugins for the CloudAMQP instance.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as cloudamqp from "@pulumi/cloudamqp";
 *
 * const plugins = cloudamqp.getPlugins({
 *     instanceId: instance.id,
 * });
 * ```
 *
 * ## Attributes reference
 *
 * All attributes reference are computed
 *
 * * `id`      - The identifier for this resource.
 * * `plugins` - An array of plugins. Each `plugins` block consists of the fields documented below.
 * * `sleep` - (Optional) Configurable sleep time (seconds) for retries when requesting information
 *   about plugins. Default set to 10 seconds. *Available from v1.29.0*
 * * `timeout` - (Optional) - Configurable timeout time (seconds) for retries when requesting
 *   information about plugins. Default set to 1800 seconds. *Available from v1.29.0*
 *
 * ***
 *
 * The `plugins` block consist of
 *
 * * `name`        - The type of the recipient.
 * * `version`     - Rabbit MQ version that the plugins are shipped with.
 * * `description` - Description of what the plugin does.
 * * `enabled`     - Enable or disable information for the plugin.
 *
 * ## Dependency
 *
 * This data source depends on CloudAMQP instance identifier, `cloudamqp_instance.instance.id`.
 */
export function getPluginsOutput(args: GetPluginsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetPluginsResult> {
    return pulumi.output(args).apply((a: any) => getPlugins(a, opts))
}

/**
 * A collection of arguments for invoking getPlugins.
 */
export interface GetPluginsOutputArgs {
    /**
     * The CloudAMQP instance identifier.
     */
    instanceId: pulumi.Input<number>;
}
