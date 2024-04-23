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
 * <details>
 *   <summary>
 *     <b>
 *       <i>Basic example of CPU and memory alarm</i>
 *     </b>
 *   </summary>
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as cloudamqp from "@pulumi/cloudamqp";
 *
 * // New recipient
 * const recipient01 = new cloudamqp.Notification("recipient_01", {
 *     instanceId: instance.id,
 *     type: "email",
 *     value: "alarm@example.com",
 *     name: "alarm",
 * });
 * // New cpu alarm
 * const cpuAlarm = new cloudamqp.Alarm("cpu_alarm", {
 *     instanceId: instance.id,
 *     type: "cpu",
 *     enabled: true,
 *     reminderInterval: 600,
 *     valueThreshold: 95,
 *     timeThreshold: 600,
 *     recipients: [recipient01.id],
 * });
 * // New memory alarm
 * const memoryAlarm = new cloudamqp.Alarm("memory_alarm", {
 *     instanceId: instance.id,
 *     type: "memory",
 *     enabled: true,
 *     reminderInterval: 600,
 *     valueThreshold: 95,
 *     timeThreshold: 600,
 *     recipients: [recipient01.id],
 * });
 * ```
 *
 * </details>
 *
 * <details>
 *   <summary>
 *     <b>
 *       <i>Manage notice alarm, available from v1.29.5</i>
 *     </b>
 *   </summary>
 *
 * Only one notice alarm can exists and cannot be created, instead the alarm resource will be updated.
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as cloudamqp from "@pulumi/cloudamqp";
 *
 * // New recipient
 * const recipient01 = new cloudamqp.Notification("recipient_01", {
 *     instanceId: instance.id,
 *     type: "email",
 *     value: "alarm@example.com",
 *     name: "alarm",
 * });
 * // Update existing notice alarm
 * const notice = new cloudamqp.Alarm("notice", {
 *     instanceId: instance.id,
 *     type: "notice",
 *     enabled: true,
 *     recipients: [recipient01.id],
 * });
 * ```
 *
 * </details>
 *
 * ## Alarm Type reference
 *
 * Supported alarm types: `cpu, memory, disk, queue, connection, flow, consumer, netsplit, server_unreachable, notice`
 *
 * Required arguments for all alarms: `instance_id, type, enabled`<br>
 * Optional argument for all alarms: `tags, queue_regex, vhostRegex`
 *
 * | Name | Type | Shared | Dedicated | Required arguments |
 * | ---- | ---- | ---- | ---- | ---- |
 * | CPU | cpu | - | &#10004; | time_threshold, valueThreshold |
 * | Memory | memory | - | &#10004;  | time_threshold, valueThreshold |
 * | Disk space | disk | - | &#10004;  | time_threshold, valueThreshold |
 * | Queue | queue | &#10004;  | &#10004; | time_threshold, value_threshold, queue_regex, vhost_regex, messageType |
 * | Connection | connection | &#10004; | &#10004; | time_threshold, valueThreshold |
 * | Connection flow | flow | &#10004; | &#10004; | time_threshold, valueThreshold |
 * | Consumer | consumer | &#10004; | &#10004; | time_threshold, value_threshold, queue, vhost |
 * | Netsplit | netsplit | - | &#10004; | timeThreshold |
 * | Server unreachable | serverUnreachable  | - | &#10004;  | timeThreshold |
 * | Notice | notice | &#10004; | &#10004; | |
 *
 * > Notice alarm is manadatory! Only one can exists and cannot be deleted. Setting `noDefaultAlarm` to true, will still create this alarm. See updated changes to notice alarm below.
 *
 * ## Dependency
 *
 * This resource depends on CloudAMQP instance identifier, `cloudamqp_instance.instance.id`.
 *
 * ## Notice alarm
 *
 * There is a limitation for notice alarm in the API backend. This alarm is mandatory, multiple
 * alarms cannot exists or be deleted.
 *
 * From provider version v1.29.5
 * it's possible to manage the notice alarm and no longer needs to be imported. Just create the
 * alarm resource as usually and it will be updated with given recipients. If the alarm is deleted
 * it will only be removed from the state file, but will still be enabled in the backend.
 *
 * ## Import
 *
 * `cloudamqp_alarm` can be imported using CloudAMQP internal identifier of the alarm together (CSV separated) with the instance identifier. To retrieve the alarm identifier, use [CloudAMQP API](https://docs.cloudamqp.com/cloudamqp_api.html#list-alarms)
 *
 * ```sh
 * $ pulumi import cloudamqp:index/alarm:Alarm alarm <id>,<instance_id>`
 * ```
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
     *
     * Specific argument for `disk` alarm
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
     * The reminder interval (in seconds) to resend the alarm if not resolved. Set to 0 for no reminders. The Default is 0.
     */
    public readonly reminderInterval!: pulumi.Output<number | undefined>;
    /**
     * The time interval (in seconds) the `valueThreshold` should be active before triggering an alarm.
     */
    public readonly timeThreshold!: pulumi.Output<number | undefined>;
    /**
     * The alarm type, see valid options below.
     */
    public readonly type!: pulumi.Output<string>;
    /**
     * Disk value threshold calculation, `fixed, percentage` of disk space remaining.
     *
     * Based on alarm type, different arguments are flagged as required or optional.
     */
    public readonly valueCalculation!: pulumi.Output<string | undefined>;
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
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as AlarmState | undefined;
            resourceInputs["enabled"] = state ? state.enabled : undefined;
            resourceInputs["instanceId"] = state ? state.instanceId : undefined;
            resourceInputs["messageType"] = state ? state.messageType : undefined;
            resourceInputs["queueRegex"] = state ? state.queueRegex : undefined;
            resourceInputs["recipients"] = state ? state.recipients : undefined;
            resourceInputs["reminderInterval"] = state ? state.reminderInterval : undefined;
            resourceInputs["timeThreshold"] = state ? state.timeThreshold : undefined;
            resourceInputs["type"] = state ? state.type : undefined;
            resourceInputs["valueCalculation"] = state ? state.valueCalculation : undefined;
            resourceInputs["valueThreshold"] = state ? state.valueThreshold : undefined;
            resourceInputs["vhostRegex"] = state ? state.vhostRegex : undefined;
        } else {
            const args = argsOrState as AlarmArgs | undefined;
            if ((!args || args.enabled === undefined) && !opts.urn) {
                throw new Error("Missing required property 'enabled'");
            }
            if ((!args || args.instanceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'instanceId'");
            }
            if ((!args || args.recipients === undefined) && !opts.urn) {
                throw new Error("Missing required property 'recipients'");
            }
            if ((!args || args.type === undefined) && !opts.urn) {
                throw new Error("Missing required property 'type'");
            }
            resourceInputs["enabled"] = args ? args.enabled : undefined;
            resourceInputs["instanceId"] = args ? args.instanceId : undefined;
            resourceInputs["messageType"] = args ? args.messageType : undefined;
            resourceInputs["queueRegex"] = args ? args.queueRegex : undefined;
            resourceInputs["recipients"] = args ? args.recipients : undefined;
            resourceInputs["reminderInterval"] = args ? args.reminderInterval : undefined;
            resourceInputs["timeThreshold"] = args ? args.timeThreshold : undefined;
            resourceInputs["type"] = args ? args.type : undefined;
            resourceInputs["valueCalculation"] = args ? args.valueCalculation : undefined;
            resourceInputs["valueThreshold"] = args ? args.valueThreshold : undefined;
            resourceInputs["vhostRegex"] = args ? args.vhostRegex : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Alarm.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Alarm resources.
 */
export interface AlarmState {
    /**
     * Enable or disable the alarm to trigger.
     */
    enabled?: pulumi.Input<boolean>;
    /**
     * The CloudAMQP instance ID.
     */
    instanceId?: pulumi.Input<number>;
    /**
     * Message type `(total, unacked, ready)` used by queue alarm type.
     *
     * Specific argument for `disk` alarm
     */
    messageType?: pulumi.Input<string>;
    /**
     * Regex for which queue to check.
     */
    queueRegex?: pulumi.Input<string>;
    /**
     * Identifier for recipient to be notified. Leave empty to notify all recipients.
     */
    recipients?: pulumi.Input<pulumi.Input<number>[]>;
    /**
     * The reminder interval (in seconds) to resend the alarm if not resolved. Set to 0 for no reminders. The Default is 0.
     */
    reminderInterval?: pulumi.Input<number>;
    /**
     * The time interval (in seconds) the `valueThreshold` should be active before triggering an alarm.
     */
    timeThreshold?: pulumi.Input<number>;
    /**
     * The alarm type, see valid options below.
     */
    type?: pulumi.Input<string>;
    /**
     * Disk value threshold calculation, `fixed, percentage` of disk space remaining.
     *
     * Based on alarm type, different arguments are flagged as required or optional.
     */
    valueCalculation?: pulumi.Input<string>;
    /**
     * The value to trigger the alarm for.
     */
    valueThreshold?: pulumi.Input<number>;
    /**
     * Regex for which vhost to check
     */
    vhostRegex?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Alarm resource.
 */
export interface AlarmArgs {
    /**
     * Enable or disable the alarm to trigger.
     */
    enabled: pulumi.Input<boolean>;
    /**
     * The CloudAMQP instance ID.
     */
    instanceId: pulumi.Input<number>;
    /**
     * Message type `(total, unacked, ready)` used by queue alarm type.
     *
     * Specific argument for `disk` alarm
     */
    messageType?: pulumi.Input<string>;
    /**
     * Regex for which queue to check.
     */
    queueRegex?: pulumi.Input<string>;
    /**
     * Identifier for recipient to be notified. Leave empty to notify all recipients.
     */
    recipients: pulumi.Input<pulumi.Input<number>[]>;
    /**
     * The reminder interval (in seconds) to resend the alarm if not resolved. Set to 0 for no reminders. The Default is 0.
     */
    reminderInterval?: pulumi.Input<number>;
    /**
     * The time interval (in seconds) the `valueThreshold` should be active before triggering an alarm.
     */
    timeThreshold?: pulumi.Input<number>;
    /**
     * The alarm type, see valid options below.
     */
    type: pulumi.Input<string>;
    /**
     * Disk value threshold calculation, `fixed, percentage` of disk space remaining.
     *
     * Based on alarm type, different arguments are flagged as required or optional.
     */
    valueCalculation?: pulumi.Input<string>;
    /**
     * The value to trigger the alarm for.
     */
    valueThreshold?: pulumi.Input<number>;
    /**
     * Regex for which vhost to check
     */
    vhostRegex?: pulumi.Input<string>;
}
