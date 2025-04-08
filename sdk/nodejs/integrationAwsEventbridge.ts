// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * This resource allows you to create and manage, an [AWS EventBridge] for a CloudAMQP instance. Once
 * created, continue to map the EventBridge in the [AWS Eventbridge console].
 *
 * >  Our consumer needs to have exclusive usage to the configured queue and the maximum body size
 * allowed on msgs by AWS is 256kb. The message body has to be valid JSON for AWS Eventbridge to accept
 * it. If messages are too large or are not valid JSON, they will be rejected (tip: setup a dead-letter
 * queue to catch them).
 *
 * Not possible to update this resource. Any changes made to the argument will destroy and recreate the
 * resource. Hence why all arguments use ForceNew.
 *
 * Only available for dedicated subscription plans.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as cloudamqp from "@pulumi/cloudamqp";
 *
 * const instance = new cloudamqp.Instance("instance", {
 *     name: "Test instance",
 *     plan: "penguin-1",
 *     region: "amazon-web-services::us-west-1",
 *     rmqVersion: "3.11.5",
 *     tags: ["aws"],
 * });
 * const _this = new cloudamqp.IntegrationAwsEventbridge("this", {
 *     instanceId: instance.id,
 *     vhost: instance.vhost,
 *     queue: "<QUEUE-NAME>",
 *     awsAccountId: "<AWS-ACCOUNT-ID>",
 *     awsRegion: "us-west-1",
 *     withHeaders: true,
 * });
 * ```
 *
 * ## Argument References
 *
 * The following arguments are supported:
 *
 * * `awsAccountId`  - (ForceNew/Required) The 12 digit AWS Account ID where you want the events to
 *                       be sent to.
 * * `awsRegion`      - (ForceNew/Required) The AWS region where you the events to be sent to.
 *                       (e.g. us-west-1, us-west-2, ..., etc.)
 * * `vhost`           - (ForceNew/Required) The VHost the queue resides in.
 * * `queue`           - (ForceNew/Required) A (durable) queue on your RabbitMQ instance.
 * * `withHeaders`    - (ForceNew/Required) Include message headers in the event data.
 *                       `({ "headers": { }, "body": { "your": "message" } })`
 *
 * ## Dependency
 *
 * This resource depends on CloudAMQP instance identifier, `cloudamqp_instance.instance.id`.
 *
 * ## Import
 *
 * `cloudamqp_integration_aws_eventbridge` can be imported using the resource identifier together with
 *
 * CloudAMQP instance identifier (CSV separated). To retrieve the resource identifier, use
 *
 * [CloudAMQP API list eventbridges].
 *
 * From Terraform v1.5.0, the `import` block can be used to import this resource:
 *
 * hcl
 *
 * import {
 *
 *   to = cloudamqp_integration_aws_eventbridge.this
 *
 *   id = format("<id>,%s", cloudamqp_instance.instance.id)
 *
 * }
 *
 * Or with Terraform CLI:
 *
 * ```sh
 * $ pulumi import cloudamqp:index/integrationAwsEventbridge:IntegrationAwsEventbridge this <id>,<instance_id>`
 * ```
 *
 * [AWS EventBridge]: https://aws.amazon.com/eventbridge
 *
 * [AWS Eventbridge console]: https://console.aws.amazon.com/events/home
 *
 * [CloudAMQP API list eventbridges]: https://docs.cloudamqp.com/cloudamqp_api.html#list-eventbridges
 */
export class IntegrationAwsEventbridge extends pulumi.CustomResource {
    /**
     * Get an existing IntegrationAwsEventbridge resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: IntegrationAwsEventbridgeState, opts?: pulumi.CustomResourceOptions): IntegrationAwsEventbridge {
        return new IntegrationAwsEventbridge(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'cloudamqp:index/integrationAwsEventbridge:IntegrationAwsEventbridge';

    /**
     * Returns true if the given object is an instance of IntegrationAwsEventbridge.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is IntegrationAwsEventbridge {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === IntegrationAwsEventbridge.__pulumiType;
    }

    /**
     * The 12 digit AWS Account ID where you want the events to be sent to.
     */
    public readonly awsAccountId!: pulumi.Output<string>;
    /**
     * The AWS region where you the events to be sent to. (e.g. us-west-1, us-west-2, ..., etc.)
     */
    public readonly awsRegion!: pulumi.Output<string>;
    /**
     * Instance identifier
     */
    public readonly instanceId!: pulumi.Output<number>;
    /**
     * A (durable) queue on your RabbitMQ instance.
     */
    public readonly queue!: pulumi.Output<string>;
    /**
     * Always set to null, unless there is an error starting the EventBridge.
     */
    public /*out*/ readonly status!: pulumi.Output<string>;
    /**
     * The VHost the queue resides in.
     */
    public readonly vhost!: pulumi.Output<string>;
    /**
     * Include message headers in the event data.
     */
    public readonly withHeaders!: pulumi.Output<boolean>;

    /**
     * Create a IntegrationAwsEventbridge resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: IntegrationAwsEventbridgeArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: IntegrationAwsEventbridgeArgs | IntegrationAwsEventbridgeState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as IntegrationAwsEventbridgeState | undefined;
            resourceInputs["awsAccountId"] = state ? state.awsAccountId : undefined;
            resourceInputs["awsRegion"] = state ? state.awsRegion : undefined;
            resourceInputs["instanceId"] = state ? state.instanceId : undefined;
            resourceInputs["queue"] = state ? state.queue : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["vhost"] = state ? state.vhost : undefined;
            resourceInputs["withHeaders"] = state ? state.withHeaders : undefined;
        } else {
            const args = argsOrState as IntegrationAwsEventbridgeArgs | undefined;
            if ((!args || args.awsAccountId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'awsAccountId'");
            }
            if ((!args || args.awsRegion === undefined) && !opts.urn) {
                throw new Error("Missing required property 'awsRegion'");
            }
            if ((!args || args.instanceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'instanceId'");
            }
            if ((!args || args.queue === undefined) && !opts.urn) {
                throw new Error("Missing required property 'queue'");
            }
            if ((!args || args.vhost === undefined) && !opts.urn) {
                throw new Error("Missing required property 'vhost'");
            }
            if ((!args || args.withHeaders === undefined) && !opts.urn) {
                throw new Error("Missing required property 'withHeaders'");
            }
            resourceInputs["awsAccountId"] = args ? args.awsAccountId : undefined;
            resourceInputs["awsRegion"] = args ? args.awsRegion : undefined;
            resourceInputs["instanceId"] = args ? args.instanceId : undefined;
            resourceInputs["queue"] = args ? args.queue : undefined;
            resourceInputs["vhost"] = args ? args.vhost : undefined;
            resourceInputs["withHeaders"] = args ? args.withHeaders : undefined;
            resourceInputs["status"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(IntegrationAwsEventbridge.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering IntegrationAwsEventbridge resources.
 */
export interface IntegrationAwsEventbridgeState {
    /**
     * The 12 digit AWS Account ID where you want the events to be sent to.
     */
    awsAccountId?: pulumi.Input<string>;
    /**
     * The AWS region where you the events to be sent to. (e.g. us-west-1, us-west-2, ..., etc.)
     */
    awsRegion?: pulumi.Input<string>;
    /**
     * Instance identifier
     */
    instanceId?: pulumi.Input<number>;
    /**
     * A (durable) queue on your RabbitMQ instance.
     */
    queue?: pulumi.Input<string>;
    /**
     * Always set to null, unless there is an error starting the EventBridge.
     */
    status?: pulumi.Input<string>;
    /**
     * The VHost the queue resides in.
     */
    vhost?: pulumi.Input<string>;
    /**
     * Include message headers in the event data.
     */
    withHeaders?: pulumi.Input<boolean>;
}

/**
 * The set of arguments for constructing a IntegrationAwsEventbridge resource.
 */
export interface IntegrationAwsEventbridgeArgs {
    /**
     * The 12 digit AWS Account ID where you want the events to be sent to.
     */
    awsAccountId: pulumi.Input<string>;
    /**
     * The AWS region where you the events to be sent to. (e.g. us-west-1, us-west-2, ..., etc.)
     */
    awsRegion: pulumi.Input<string>;
    /**
     * Instance identifier
     */
    instanceId: pulumi.Input<number>;
    /**
     * A (durable) queue on your RabbitMQ instance.
     */
    queue: pulumi.Input<string>;
    /**
     * The VHost the queue resides in.
     */
    vhost: pulumi.Input<string>;
    /**
     * Include message headers in the event data.
     */
    withHeaders: pulumi.Input<boolean>;
}
