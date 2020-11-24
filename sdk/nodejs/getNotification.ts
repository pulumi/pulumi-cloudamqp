// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * Use this data source to retrieve information about default or created recipients. The recipient will receive notifications assigned to an alarm that has triggered. To retrieve the recipient either use `recipientId` or `name`.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as cloudamqp from "@pulumi/cloudamqp";
 *
 * const defaultRecipient = cloudamqp.getNotification({
 *     instanceId: cloudamqp_instance.instance.id,
 *     name: "default",
 * });
 * ```
 * ## Argument reference
 *
 * * `instanceId`   - (Required) The CloudAMQP instance identifier.
 * * `recipientId`  - (Optional) The recipient identifier.
 * * `name`          - (Optional) The name set for the recipient.
 *
 * ## Attribute reference
 *
 * * `type`  - (Computed) The type of the recipient.
 * * `value` - (Computed) The notification endpoint, where to send the notification.
 *
 * ## Dependency
 *
 * This data source depends on CloudAMQP instance identifier, `cloudamqp_instance.instance.id`.
 */
export function getNotification(args: GetNotificationArgs, opts?: pulumi.InvokeOptions): Promise<GetNotificationResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("cloudamqp:index/getNotification:getNotification", {
        "instanceId": args.instanceId,
        "name": args.name,
        "recipientId": args.recipientId,
    }, opts);
}

/**
 * A collection of arguments for invoking getNotification.
 */
export interface GetNotificationArgs {
    readonly instanceId: number;
    readonly name?: string;
    readonly recipientId?: number;
}

/**
 * A collection of values returned by getNotification.
 */
export interface GetNotificationResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly instanceId: number;
    readonly name?: string;
    readonly recipientId?: number;
    readonly type: string;
    readonly value: string;
}
