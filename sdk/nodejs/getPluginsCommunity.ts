// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as cloudamqp from "@pulumi/cloudamqp";
 *
 * const communitPlugins = cloudamqp.getPluginsCommunity({
 *     instanceId: cloudamqp_instance.instance.id,
 * });
 * ```
 * ## Attributes reference
 *
 * All attributes reference are computed
 *
 * * `id`      - The identifier for this resource.
 * * `plugins` - An array of community plugins. Each `plugins` block consists of the fields documented below.
 *
 * ***
 *
 * The `plugins` block consists of
 *
 * * `name`        - The type of the recipient.
 * * `require`     - Min. required Rabbit MQ version to be used.
 * * `description` - Description of what the plugin does.
 *
 * ## Dependency
 *
 * This data source depends on CloudAMQP instance identifier, `cloudamqp_instance.instance.id`.
 */
export function getPluginsCommunity(args: GetPluginsCommunityArgs, opts?: pulumi.InvokeOptions): Promise<GetPluginsCommunityResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("cloudamqp:index/getPluginsCommunity:getPluginsCommunity", {
        "instanceId": args.instanceId,
    }, opts);
}

/**
 * A collection of arguments for invoking getPluginsCommunity.
 */
export interface GetPluginsCommunityArgs {
    /**
     * The CloudAMQP instance identifier.
     */
    instanceId: number;
}

/**
 * A collection of values returned by getPluginsCommunity.
 */
export interface GetPluginsCommunityResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly instanceId: number;
    readonly plugins: outputs.GetPluginsCommunityPlugin[];
}
/**
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as cloudamqp from "@pulumi/cloudamqp";
 *
 * const communitPlugins = cloudamqp.getPluginsCommunity({
 *     instanceId: cloudamqp_instance.instance.id,
 * });
 * ```
 * ## Attributes reference
 *
 * All attributes reference are computed
 *
 * * `id`      - The identifier for this resource.
 * * `plugins` - An array of community plugins. Each `plugins` block consists of the fields documented below.
 *
 * ***
 *
 * The `plugins` block consists of
 *
 * * `name`        - The type of the recipient.
 * * `require`     - Min. required Rabbit MQ version to be used.
 * * `description` - Description of what the plugin does.
 *
 * ## Dependency
 *
 * This data source depends on CloudAMQP instance identifier, `cloudamqp_instance.instance.id`.
 */
export function getPluginsCommunityOutput(args: GetPluginsCommunityOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetPluginsCommunityResult> {
    return pulumi.output(args).apply((a: any) => getPluginsCommunity(a, opts))
}

/**
 * A collection of arguments for invoking getPluginsCommunity.
 */
export interface GetPluginsCommunityOutputArgs {
    /**
     * The CloudAMQP instance identifier.
     */
    instanceId: pulumi.Input<number>;
}
