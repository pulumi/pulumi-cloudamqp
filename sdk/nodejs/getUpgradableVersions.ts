// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Use this data source to retrieve information about possible upgradable versions for RabbitMQ and Erlang.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as cloudamqp from "@pulumi/cloudamqp";
 *
 * const versions = cloudamqp.getUpgradableVersions({
 *     instanceId: instance.id,
 * });
 * ```
 *
 * ## Attributes reference
 *
 * All attributes reference are computed
 *
 * * `newRabbitmqVersion`  - Possible upgradable version for RabbitMQ.
 * * `newErlangVersion`    - Possible upgradable version for Erlang.
 *
 * ## Dependency
 *
 * This data source depends on CloudAMQP instance identifier, `cloudamqp_instance.instance.id`.
 */
export function getUpgradableVersions(args: GetUpgradableVersionsArgs, opts?: pulumi.InvokeOptions): Promise<GetUpgradableVersionsResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("cloudamqp:index/getUpgradableVersions:getUpgradableVersions", {
        "instanceId": args.instanceId,
    }, opts);
}

/**
 * A collection of arguments for invoking getUpgradableVersions.
 */
export interface GetUpgradableVersionsArgs {
    /**
     * The CloudAMQP instance identifier.
     */
    instanceId: number;
}

/**
 * A collection of values returned by getUpgradableVersions.
 */
export interface GetUpgradableVersionsResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly instanceId: number;
    readonly newErlangVersion: string;
    readonly newRabbitmqVersion: string;
}
/**
 * Use this data source to retrieve information about possible upgradable versions for RabbitMQ and Erlang.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as cloudamqp from "@pulumi/cloudamqp";
 *
 * const versions = cloudamqp.getUpgradableVersions({
 *     instanceId: instance.id,
 * });
 * ```
 *
 * ## Attributes reference
 *
 * All attributes reference are computed
 *
 * * `newRabbitmqVersion`  - Possible upgradable version for RabbitMQ.
 * * `newErlangVersion`    - Possible upgradable version for Erlang.
 *
 * ## Dependency
 *
 * This data source depends on CloudAMQP instance identifier, `cloudamqp_instance.instance.id`.
 */
export function getUpgradableVersionsOutput(args: GetUpgradableVersionsOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetUpgradableVersionsResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("cloudamqp:index/getUpgradableVersions:getUpgradableVersions", {
        "instanceId": args.instanceId,
    }, opts);
}

/**
 * A collection of arguments for invoking getUpgradableVersions.
 */
export interface GetUpgradableVersionsOutputArgs {
    /**
     * The CloudAMQP instance identifier.
     */
    instanceId: pulumi.Input<number>;
}
