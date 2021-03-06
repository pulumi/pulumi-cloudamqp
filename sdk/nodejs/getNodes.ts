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
 * ## Attribute reference
 *
 * * `nodes` - (Computed) An array of node information. Each `nodes` block consists of the fields documented below.
 *
 * ***
 *
 * The `nodes` block consist of
 *
 * * `hostname`          - (Computed) Hostname assigned to the node.
 * * `name`              - (Computed) Name of the node.
 * * `running`           - (Computed) Is the node running?
 * * `rabbitmqVersion`  - (Computed) Currently configured Rabbit MQ version on the node.
 * * `erlangVersion`    - (Computed) Currently used Erlanbg version on the node.
 * * `hipe`              - (Computed) Enable or disable High-performance Erlang.
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
        "nodes": args.nodes,
    }, opts);
}

/**
 * A collection of arguments for invoking getNodes.
 */
export interface GetNodesArgs {
    readonly instanceId: number;
    readonly nodes?: inputs.GetNodesNode[];
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
