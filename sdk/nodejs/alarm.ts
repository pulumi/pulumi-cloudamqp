// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export class Alarm extends pulumi.CustomResource {
    /**
     * Get an existing Alarm resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
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
     * Instance identifier
     */
    public readonly instanceId!: pulumi.Output<number>;
    /**
     * Identifiers for recipients to be notified. Leave empty to notifiy all recipients.
     */
    public readonly notificationIds!: pulumi.Output<number[] | undefined>;
    /**
     * Regex for which queues to check
     */
    public readonly queueRegex!: pulumi.Output<string | undefined>;
    /**
     * For how long (in seconds) the value_threshold should be active before trigger alarm
     */
    public readonly timeThreshold!: pulumi.Output<number | undefined>;
    /**
     * Type of the alarm, valid options are: cpu, memory, disk_usage, queue_length, connection_count, consumers_count,
     * net_split
     */
    public readonly type!: pulumi.Output<string>;
    /**
     * What value to trigger the alarm for
     */
    public readonly valueThreshold!: pulumi.Output<number | undefined>;
    /**
     * Regex for which vhost the queues are in
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
            inputs["instanceId"] = state ? state.instanceId : undefined;
            inputs["notificationIds"] = state ? state.notificationIds : undefined;
            inputs["queueRegex"] = state ? state.queueRegex : undefined;
            inputs["timeThreshold"] = state ? state.timeThreshold : undefined;
            inputs["type"] = state ? state.type : undefined;
            inputs["valueThreshold"] = state ? state.valueThreshold : undefined;
            inputs["vhostRegex"] = state ? state.vhostRegex : undefined;
        } else {
            const args = argsOrState as AlarmArgs | undefined;
            if (!args || args.instanceId === undefined) {
                throw new Error("Missing required property 'instanceId'");
            }
            if (!args || args.type === undefined) {
                throw new Error("Missing required property 'type'");
            }
            inputs["instanceId"] = args ? args.instanceId : undefined;
            inputs["notificationIds"] = args ? args.notificationIds : undefined;
            inputs["queueRegex"] = args ? args.queueRegex : undefined;
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
     * Instance identifier
     */
    readonly instanceId?: pulumi.Input<number>;
    /**
     * Identifiers for recipients to be notified. Leave empty to notifiy all recipients.
     */
    readonly notificationIds?: pulumi.Input<pulumi.Input<number>[]>;
    /**
     * Regex for which queues to check
     */
    readonly queueRegex?: pulumi.Input<string>;
    /**
     * For how long (in seconds) the value_threshold should be active before trigger alarm
     */
    readonly timeThreshold?: pulumi.Input<number>;
    /**
     * Type of the alarm, valid options are: cpu, memory, disk_usage, queue_length, connection_count, consumers_count,
     * net_split
     */
    readonly type?: pulumi.Input<string>;
    /**
     * What value to trigger the alarm for
     */
    readonly valueThreshold?: pulumi.Input<number>;
    /**
     * Regex for which vhost the queues are in
     */
    readonly vhostRegex?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Alarm resource.
 */
export interface AlarmArgs {
    /**
     * Instance identifier
     */
    readonly instanceId: pulumi.Input<number>;
    /**
     * Identifiers for recipients to be notified. Leave empty to notifiy all recipients.
     */
    readonly notificationIds?: pulumi.Input<pulumi.Input<number>[]>;
    /**
     * Regex for which queues to check
     */
    readonly queueRegex?: pulumi.Input<string>;
    /**
     * For how long (in seconds) the value_threshold should be active before trigger alarm
     */
    readonly timeThreshold?: pulumi.Input<number>;
    /**
     * Type of the alarm, valid options are: cpu, memory, disk_usage, queue_length, connection_count, consumers_count,
     * net_split
     */
    readonly type: pulumi.Input<string>;
    /**
     * What value to trigger the alarm for
     */
    readonly valueThreshold?: pulumi.Input<number>;
    /**
     * Regex for which vhost the queues are in
     */
    readonly vhostRegex?: pulumi.Input<string>;
}
