// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Use this data source to retrieve information about VPC for a CloudAMQP instance.
 *
 * Only available for CloudAMQP instances hosted in AWS.
 *
 * ## Example Usage
 *
 * <details>
 *   <summary>
 *     <b>
 *       <i>AWS VPC peering pre v1.16.0</i>
 *     </b>
 *   </summary>
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as cloudamqp from "@pulumi/cloudamqp";
 *
 * const vpcInfo = cloudamqp.getVpcInfo({
 *     instanceId: instance.id,
 * });
 * ```
 * </details>
 *
 * <details>
 *   <summary>
 *     <b>
 *       <i>AWS VPC peering post v1.16.0 (Managed VPC)</i>
 *     </b>
 *   </summary>
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as cloudamqp from "@pulumi/cloudamqp";
 *
 * const vpcInfo = cloudamqp.getVpcInfo({
 *     vpcId: vpc.id,
 * });
 * ```
 * </details>
 *
 * ## Attributes reference
 *
 * All attributes reference are computed
 *
 * * `id`                  - The identifier for this resource.
 * * `name`                - The name of the CloudAMQP instance.
 * * `vpcSubnet`          - Dedicated VPC subnet.
 * * `ownerId`            - AWS account identifier.
 * * `securityGroupId`   - AWS security group identifier.
 *
 * ## Dependency
 *
 * *Pre v1.16.0*
 * This resource depends on CloudAMQP instance identifier, `cloudamqp_instance.instance.id`.
 *
 * *Post v1.16.0*
 * This resource depends on CloudAMQP managed VPC identifier, `cloudamqp_vpc.vpc.id` or instance identifier, `cloudamqp_instance.instance.id`.
 */
export function getVpcInfo(args?: GetVpcInfoArgs, opts?: pulumi.InvokeOptions): Promise<GetVpcInfoResult> {
    args = args || {};
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("cloudamqp:index/getVpcInfo:getVpcInfo", {
        "instanceId": args.instanceId,
        "vpcId": args.vpcId,
    }, opts);
}

/**
 * A collection of arguments for invoking getVpcInfo.
 */
export interface GetVpcInfoArgs {
    /**
     * The CloudAMQP instance identifier.
     *
     * ***Deprecated: Changed from required to optional in v1.16.0 will be removed in next major version (v2.0)***
     */
    instanceId?: number;
    /**
     * The managed VPC identifier.
     *
     * ***Note: Added as optional in version v1.16.0 and will be required in next major version (v2.0)***
     */
    vpcId?: string;
}

/**
 * A collection of values returned by getVpcInfo.
 */
export interface GetVpcInfoResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly instanceId?: number;
    readonly name: string;
    readonly ownerId: string;
    readonly securityGroupId: string;
    readonly vpcId?: string;
    readonly vpcSubnet: string;
}
/**
 * Use this data source to retrieve information about VPC for a CloudAMQP instance.
 *
 * Only available for CloudAMQP instances hosted in AWS.
 *
 * ## Example Usage
 *
 * <details>
 *   <summary>
 *     <b>
 *       <i>AWS VPC peering pre v1.16.0</i>
 *     </b>
 *   </summary>
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as cloudamqp from "@pulumi/cloudamqp";
 *
 * const vpcInfo = cloudamqp.getVpcInfo({
 *     instanceId: instance.id,
 * });
 * ```
 * </details>
 *
 * <details>
 *   <summary>
 *     <b>
 *       <i>AWS VPC peering post v1.16.0 (Managed VPC)</i>
 *     </b>
 *   </summary>
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as cloudamqp from "@pulumi/cloudamqp";
 *
 * const vpcInfo = cloudamqp.getVpcInfo({
 *     vpcId: vpc.id,
 * });
 * ```
 * </details>
 *
 * ## Attributes reference
 *
 * All attributes reference are computed
 *
 * * `id`                  - The identifier for this resource.
 * * `name`                - The name of the CloudAMQP instance.
 * * `vpcSubnet`          - Dedicated VPC subnet.
 * * `ownerId`            - AWS account identifier.
 * * `securityGroupId`   - AWS security group identifier.
 *
 * ## Dependency
 *
 * *Pre v1.16.0*
 * This resource depends on CloudAMQP instance identifier, `cloudamqp_instance.instance.id`.
 *
 * *Post v1.16.0*
 * This resource depends on CloudAMQP managed VPC identifier, `cloudamqp_vpc.vpc.id` or instance identifier, `cloudamqp_instance.instance.id`.
 */
export function getVpcInfoOutput(args?: GetVpcInfoOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetVpcInfoResult> {
    args = args || {};
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("cloudamqp:index/getVpcInfo:getVpcInfo", {
        "instanceId": args.instanceId,
        "vpcId": args.vpcId,
    }, opts);
}

/**
 * A collection of arguments for invoking getVpcInfo.
 */
export interface GetVpcInfoOutputArgs {
    /**
     * The CloudAMQP instance identifier.
     *
     * ***Deprecated: Changed from required to optional in v1.16.0 will be removed in next major version (v2.0)***
     */
    instanceId?: pulumi.Input<number>;
    /**
     * The managed VPC identifier.
     *
     * ***Note: Added as optional in version v1.16.0 and will be required in next major version (v2.0)***
     */
    vpcId?: pulumi.Input<string>;
}
