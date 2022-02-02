// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Use this data source to retrieve information about an already created CloudAMQP instance. In order to retrieve the correct information, the CoudAMQP instance identifier is needed.
 *
 * ## Attributes reference
 *
 * All attributes reference are computed
 *
 * * `id`          - The identifier for this resource.
 * * `name`        - The name of the CloudAMQP instance.
 * * `plan`        - The subscription plan for the CloudAMQP instance.
 * * `region`      - The cloud platform and region that host the CloudAMQP instance, `{platform}::{region}`.
 * * `vpcId`      - ID of the VPC configured for the CloudAMQP instance.
 * * `vpcSubnet`  - Dedicated VPC subnet configured for the CloudAMQP instance.
 * * `nodes`       - Number of nodes in the cluster of the CloudAMQP instance.
 * * `rmqVersion` - The version of installed Rabbit MQ.
 * * `url`         - (Sensitive) The AMQP URL (uses the internal hostname if the instance was created with VPC), used by clients to connect for pub/sub.
 * * `apikey`      - (Sensitive) The API key to secondary API handing alarms, integration etc.
 * * `tags`        - Tags the CloudAMQP instance with categories.
 * * `host`        - The external hostname for the CloudAMQP instance.
 * * `hostInternal` - The internal hostname for the CloudAMQP instance.
 * * `vhost`       - The virtual host configured in Rabbit MQ.
 */
export function getInstance(args: GetInstanceArgs, opts?: pulumi.InvokeOptions): Promise<GetInstanceResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("cloudamqp:index/getInstance:getInstance", {
        "instanceId": args.instanceId,
    }, opts);
}

/**
 * A collection of arguments for invoking getInstance.
 */
export interface GetInstanceArgs {
    /**
     * The CloudAMQP instance identifier.
     */
    instanceId: number;
}

/**
 * A collection of values returned by getInstance.
 */
export interface GetInstanceResult {
    readonly apikey: string;
    readonly dedicated: boolean;
    readonly host: string;
    readonly hostInternal: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly instanceId: number;
    readonly name: string;
    readonly noDefaultAlarms: boolean;
    readonly nodes: number;
    readonly plan: string;
    readonly ready: boolean;
    readonly region: string;
    readonly rmqVersion: string;
    readonly tags: string[];
    readonly url: string;
    readonly vhost: string;
    readonly vpcId: number;
    readonly vpcSubnet: string;
}

export function getInstanceOutput(args: GetInstanceOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetInstanceResult> {
    return pulumi.output(args).apply(a => getInstance(a, opts))
}

/**
 * A collection of arguments for invoking getInstance.
 */
export interface GetInstanceOutputArgs {
    /**
     * The CloudAMQP instance identifier.
     */
    instanceId: pulumi.Input<number>;
}
