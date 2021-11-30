// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * Use this data source to retrieve information about the node(s) created by CloudAMQP instance.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as cloudamqp from "@pulumi/cloudamqp";
 *
 * const nodes = cloudamqp.getNodes({
 *     instanceId: cloudamqp_instance.instance.id,
 * });
 * ```
 * ## Argument reference
 *
 * * `instanceId` - (Required) The CloudAMQP instance identifier.
 *
 * ## Attributes reference
 *
 * All attributes reference are computed
 *
 * * `id`    - The identifier for this resource.
 * * `nodes` - An array of node information. Each `nodes` block consists of the fields documented below.
 *
 * ***
 *
 * The `nodes` block consist of
 *
 * * `hostname`          - External hostname assigned to the node.
 * * `name`              - Name of the node.
 * * `running`           - Is the node running?
 * * `rabbitmqVersion`  - Currently configured Rabbit MQ version on the node.
 * * `erlangVersion`    - Currently used Erlanbg version on the node.
 * * `hipe`              - Enable or disable High-performance Erlang.
 * * `configured`        - Is the node configured?
 *
 * ## Dependency
 *
 * This data source depends on CloudAMQP instance identifier, `cloudamqp_instance.instance.id`.
 */
export function getNodes(args: GetNodesArgs, opts?: pulumi.InvokeOptions): Promise<GetNodesResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("cloudamqp:index/getNodes:getNodes", {
        "instanceId": args.instanceId,
    }, opts);
}

/**
 * A collection of arguments for invoking getNodes.
 */
export interface GetNodesArgs {
    instanceId: number;
}

/**
 * A collection of values returned by getNodes.
 */
export interface GetNodesResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly instanceId: number;
    readonly nodes: outputs.GetNodesNode[];
}

export function getNodesOutput(args: GetNodesOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetNodesResult> {
    return pulumi.output(args).apply(a => getNodes(a, opts))
}

/**
 * A collection of arguments for invoking getNodes.
 */
export interface GetNodesOutputArgs {
    instanceId: pulumi.Input<number>;
}
