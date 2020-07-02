// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * This resource allows you to create and manage alarms to trigger based on a set of conditions. Once triggerd a notification will be sent to the assigned recipients. When creating a new instance, there will also be a set of default alarms (cpu, memory and disk) created. All default alarms uses the default recipient for notifications.
 *
 * By setting `noDefaultAlarms` to *true* in `cloudamqp.Instance`. This will create the instance without default alarms and avoid the need to import them to get full control.
 *
 * Available for all subscription plans, but `lemur`and `tiger`are limited to fewer alarm types. The limited types supported can be seen in the table below in Alarm Type Reference.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as cloudamqp from "@pulumi/cloudamqp";
 *
 * // New recipient
 * const recipient01 = new cloudamqp.Notification("recipient01", {
 *     instanceId: cloudamqp_instance.instance.id,
 *     type: "email",
 *     value: "alarm@example.com",
 * });
 * // New cpu alarm
 * const cpuAlarm = new cloudamqp.Alarm("cpuAlarm", {
 *     instanceId: cloudamqp_instance.instance.id,
 *     type: "cpu",
 *     enabled: true,
 *     valueThreshold: 95,
 *     timeThreshold: 600,
 *     recipient: [2],
 * });
 * // New memory alarm
 * const memoryAlarm = new cloudamqp.Alarm("memoryAlarm", {
 *     instanceId: cloudamqp_instance.instance.id,
 *     type: "memory",
 *     enabled: true,
 *     valueThreshold: 95,
 *     timeThreshold: 600,
 *     recipient: [2],
 * });
 * ```
 * ## Alarm Type reference
 *
 * Valid options for notification type.
 *
 * Required arguments for all alarms: *instance_id*, *type* and *enabled*<br>
 * Optional argument for all alarms: *tags*, *queue_regex*, *vhost_regex*
 *
 * | Name | Type | Shared | Dedicated | Required arguments |
 * | ---- | ---- | ---- | ---- | ---- | ---- |
 * | CPU | cpu | - | &#10004; | time_threshold, valueThreshold |
 * | Memory | memory | - | &#10004;  | time_threshold, valueThreshold |
 * | Disk space | disk | - | &#10004;  | time_threshold, valueThreshold |
 * | Queue | queue | &#10004;  | &#10004;  | time_threshold, value_threshold, queue_regex, vhost_regex, messageType |
 * | Connection | connection | &#10004; | &#10004; | time_threshold, valueThreshold |
 * | Consumer | consumer | &#10004; | &#10004; | time_threshold, value_threshold, queue, vhost |
 * | Netsplit | netsplit | - | &#10004; | timeThreshold |
 * | Server unreachable | serverUnreachable  | - | &#10004;  | timeThreshold |
 * | Notice | notice | &#10004; | &#10004; |
 *
 * ## Dependency
 *
 * This resource depends on CloudAMQP instance identifier, `cloudamqp_instance.instance.id`.
 */
export class Alarm extends pulumi.CustomResource {
    /**
     * Get an existing Alarm resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AlarmState, opts?: pulumi.CustomResourceOptions): Alarm {
        return new Alarm(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'cloudamqp:index/alarm:Alarm';

    /**
     * Returns true if the given object is an instance of Alarm.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Alarm {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Alarm.__pulumiType;
    }

    /**
     * Enable or disable the alarm to trigger.
     */
    public readonly enabled!: pulumi.Output<boolean>;
    /**
     * The CloudAMQP instance ID.
     */
    public readonly instanceId!: pulumi.Output<number>;
    /**
     * Message type `(total, unacked, ready)` used by queue alarm type.
     */
    public readonly messageType!: pulumi.Output<string | undefined>;
    /**
     * Regex for which queue to check.
     */
    public readonly queueRegex!: pulumi.Output<string | undefined>;
    /**
     * Identifier for recipient to be notified. Leave empty to notify all recipients.
     */
    public readonly recipients!: pulumi.Output<number[]>;
    /**
     * The time interval (in seconds) the `valueThreshold` should be active before triggering an alarm.
     */
    public readonly timeThreshold!: pulumi.Output<number | undefined>;
    /**
     * The alarm type, see valid options below.
     */
    public readonly type!: pulumi.Output<string>;
    /**
     * The value to trigger the alarm for.
     */
    public readonly valueThreshold!: pulumi.Output<number | undefined>;
    /**
     * Regex for which vhost to check
     */
    public readonly vhostRegex!: pulumi.Output<string | undefined>;

    /**
     * Create a Alarm resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AlarmArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AlarmArgs | AlarmState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as AlarmState | undefined;
            inputs["enabled"] = state ? state.enabled : undefined;
            inputs["instanceId"] = state ? state.instanceId : undefined;
            inputs["messageType"] = state ? state.messageType : undefined;
            inputs["queueRegex"] = state ? state.queueRegex : undefined;
            inputs["recipients"] = state ? state.recipients : undefined;
            inputs["timeThreshold"] = state ? state.timeThreshold : undefined;
            inputs["type"] = state ? state.type : undefined;
            inputs["valueThreshold"] = state ? state.valueThreshold : undefined;
            inputs["vhostRegex"] = state ? state.vhostRegex : undefined;
        } else {
            const args = argsOrState as AlarmArgs | undefined;
            if (!args || args.enabled === undefined) {
                throw new Error("Missing required property 'enabled'");
            }
            if (!args || args.instanceId === undefined) {
                throw new Error("Missing required property 'instanceId'");
            }
            if (!args || args.recipients === undefined) {
                throw new Error("Missing required property 'recipients'");
            }
            if (!args || args.type === undefined) {
                throw new Error("Missing required property 'type'");
            }
            inputs["enabled"] = args ? args.enabled : undefined;
            inputs["instanceId"] = args ? args.instanceId : undefined;
            inputs["messageType"] = args ? args.messageType : undefined;
            inputs["queueRegex"] = args ? args.queueRegex : undefined;
            inputs["recipients"] = args ? args.recipients : undefined;
            inputs["timeThreshold"] = args ? args.timeThreshold : undefined;
            inputs["type"] = args ? args.type : undefined;
            inputs["valueThreshold"] = args ? args.valueThreshold : undefined;
            inputs["vhostRegex"] = args ? args.vhostRegex : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(Alarm.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Alarm resources.
 */
export interface AlarmState {
    /**
     * Enable or disable the alarm to trigger.
     */
    readonly enabled?: pulumi.Input<boolean>;
    /**
     * The CloudAMQP instance ID.
     */
    readonly instanceId?: pulumi.Input<number>;
    /**
     * Message type `(total, unacked, ready)` used by queue alarm type.
     */
    readonly messageType?: pulumi.Input<string>;
    /**
     * Regex for which queue to check.
     */
    readonly queueRegex?: pulumi.Input<string>;
    /**
     * Identifier for recipient to be notified. Leave empty to notify all recipients.
     */
    readonly recipients?: pulumi.Input<pulumi.Input<number>[]>;
    /**
     * The time interval (in seconds) the `valueThreshold` should be active before triggering an alarm.
     */
    readonly timeThreshold?: pulumi.Input<number>;
    /**
     * The alarm type, see valid options below.
     */
    readonly type?: pulumi.Input<string>;
    /**
     * The value to trigger the alarm for.
     */
    readonly valueThreshold?: pulumi.Input<number>;
    /**
     * Regex for which vhost to check
     */
    readonly vhostRegex?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Alarm resource.
 */
export interface AlarmArgs {
    /**
     * Enable or disable the alarm to trigger.
     */
    readonly enabled: pulumi.Input<boolean>;
    /**
     * The CloudAMQP instance ID.
     */
    readonly instanceId: pulumi.Input<number>;
    /**
     * Message type `(total, unacked, ready)` used by queue alarm type.
     */
    readonly messageType?: pulumi.Input<string>;
    /**
     * Regex for which queue to check.
     */
    readonly queueRegex?: pulumi.Input<string>;
    /**
     * Identifier for recipient to be notified. Leave empty to notify all recipients.
     */
    readonly recipients: pulumi.Input<pulumi.Input<number>[]>;
    /**
     * The time interval (in seconds) the `valueThreshold` should be active before triggering an alarm.
     */
    readonly timeThreshold?: pulumi.Input<number>;
    /**
     * The alarm type, see valid options below.
     */
    readonly type: pulumi.Input<string>;
    /**
     * The value to trigger the alarm for.
     */
    readonly valueThreshold?: pulumi.Input<number>;
    /**
     * Regex for which vhost to check
     */
    readonly vhostRegex?: pulumi.Input<string>;
}
