// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

export function getAlarm(args: GetAlarmArgs, opts?: pulumi.InvokeOptions): Promise<GetAlarmResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("cloudamqp:index/getAlarm:getAlarm", {
        "alarmId": args.alarmId,
        "instanceId": args.instanceId,
        "type": args.type,
    }, opts);
}

/**
 * A collection of arguments for invoking getAlarm.
 */
export interface GetAlarmArgs {
    readonly alarmId?: number;
    readonly instanceId: number;
    readonly type?: string;
}

/**
 * A collection of values returned by getAlarm.
 */
export interface GetAlarmResult {
    readonly alarmId?: number;
    readonly enabled: boolean;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly instanceId: number;
    readonly messageType: string;
    readonly queueRegex: string;
    readonly recipients: number[];
    readonly timeThreshold: number;
    readonly type?: string;
    readonly valueThreshold: number;
    readonly vhostRegex: string;
}
