// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * This resource allows you to create and manage recipients to receive alarm notifications. There will always be a default recipient created upon instance creation. This recipient will use team email and receive notifications from default alarms.
 *
 * Available for all subscription plans.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as cloudamqp from "@pulumi/cloudamqp";
 *
 * // New recipient to receieve notifications
 * const recipient01 = new cloudamqp.Notification("recipient01", {
 *     instanceId: cloudamqp_instance.instance.id,
 *     type: "email",
 *     value: "alarm@example.com",
 * });
 * ```
 * ## Notification Type reference
 *
 * Valid options for notification type.
 *
 * * email
 * * webhook
 * * pagerduty
 * * victorops
 * * opsgenie
 * * opsgenie-eu
 * * slack
 *
 * ## Dependency
 *
 * This resource depends on CloudAMQP instance identifier, `cloudamqp_instance.instance.id`.
 *
 * ## Import
 *
 * `cloudamqp_notification` can be imported using CloudAMQP internal identifier of a recipient together (CSV separated) with the instance identifier. To retrieve the identifier of a recipient, use [CloudAMQP API](https://docs.cloudamqp.com/cloudamqp_api.html#list-notification-recipients)
 *
 * ```sh
 *  $ pulumi import cloudamqp:index/notification:Notification recipient <recpient_id>,<indstance_id>`
 * ```
 */
export class Notification extends pulumi.CustomResource {
    /**
     * Get an existing Notification resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: NotificationState, opts?: pulumi.CustomResourceOptions): Notification {
        return new Notification(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'cloudamqp:index/notification:Notification';

    /**
     * Returns true if the given object is an instance of Notification.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Notification {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Notification.__pulumiType;
    }

    /**
     * The CloudAMQP instance ID.
     */
    public readonly instanceId!: pulumi.Output<number>;
    /**
     * Display name of the recipient.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Type of the notification. See valid options below.
     */
    public readonly type!: pulumi.Output<string>;
    /**
     * Endpoint to send the notification.
     */
    public readonly value!: pulumi.Output<string>;

    /**
     * Create a Notification resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: NotificationArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: NotificationArgs | NotificationState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as NotificationState | undefined;
            inputs["instanceId"] = state ? state.instanceId : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["type"] = state ? state.type : undefined;
            inputs["value"] = state ? state.value : undefined;
        } else {
            const args = argsOrState as NotificationArgs | undefined;
            if (!args || args.instanceId === undefined) {
                throw new Error("Missing required property 'instanceId'");
            }
            if (!args || args.type === undefined) {
                throw new Error("Missing required property 'type'");
            }
            if (!args || args.value === undefined) {
                throw new Error("Missing required property 'value'");
            }
            inputs["instanceId"] = args ? args.instanceId : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["type"] = args ? args.type : undefined;
            inputs["value"] = args ? args.value : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(Notification.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Notification resources.
 */
export interface NotificationState {
    /**
     * The CloudAMQP instance ID.
     */
    readonly instanceId?: pulumi.Input<number>;
    /**
     * Display name of the recipient.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Type of the notification. See valid options below.
     */
    readonly type?: pulumi.Input<string>;
    /**
     * Endpoint to send the notification.
     */
    readonly value?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Notification resource.
 */
export interface NotificationArgs {
    /**
     * The CloudAMQP instance ID.
     */
    readonly instanceId: pulumi.Input<number>;
    /**
     * Display name of the recipient.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Type of the notification. See valid options below.
     */
    readonly type: pulumi.Input<string>;
    /**
     * Endpoint to send the notification.
     */
    readonly value: pulumi.Input<string>;
}
