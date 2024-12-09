// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Use this data source to retrieve information about an already created CloudAMQP instance. In order to retrieve the correct information, the CoudAMQP instance identifier is needed.
 */
export function getInstance(args: GetInstanceArgs, opts?: pulumi.InvokeOptions): Promise<GetInstanceResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
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
    readonly backend: string;
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
/**
 * Use this data source to retrieve information about an already created CloudAMQP instance. In order to retrieve the correct information, the CoudAMQP instance identifier is needed.
 */
export function getInstanceOutput(args: GetInstanceOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetInstanceResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("cloudamqp:index/getInstance:getInstance", {
        "instanceId": args.instanceId,
    }, opts);
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
