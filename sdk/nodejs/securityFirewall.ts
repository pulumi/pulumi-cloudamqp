// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * ## Import
 *
 * `cloudamqp_security_firewall` can be imported using CloudAMQP instance identifier.
 *
 * ```sh
 *  $ pulumi import cloudamqp:index/securityFirewall:SecurityFirewall firewall <instance_id>`
 * ```
 */
export class SecurityFirewall extends pulumi.CustomResource {
    /**
     * Get an existing SecurityFirewall resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SecurityFirewallState, opts?: pulumi.CustomResourceOptions): SecurityFirewall {
        return new SecurityFirewall(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'cloudamqp:index/securityFirewall:SecurityFirewall';

    /**
     * Returns true if the given object is an instance of SecurityFirewall.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SecurityFirewall {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SecurityFirewall.__pulumiType;
    }

    /**
     * The CloudAMQP instance ID.
     */
    public readonly instanceId!: pulumi.Output<number>;
    /**
     * An array of rules, minimum of 1 needs to be configured. Each `rules` block consists of the field documented below.
     */
    public readonly rules!: pulumi.Output<outputs.SecurityFirewallRule[]>;

    /**
     * Create a SecurityFirewall resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SecurityFirewallArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SecurityFirewallArgs | SecurityFirewallState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SecurityFirewallState | undefined;
            resourceInputs["instanceId"] = state ? state.instanceId : undefined;
            resourceInputs["rules"] = state ? state.rules : undefined;
        } else {
            const args = argsOrState as SecurityFirewallArgs | undefined;
            if ((!args || args.instanceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'instanceId'");
            }
            if ((!args || args.rules === undefined) && !opts.urn) {
                throw new Error("Missing required property 'rules'");
            }
            resourceInputs["instanceId"] = args ? args.instanceId : undefined;
            resourceInputs["rules"] = args ? args.rules : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(SecurityFirewall.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SecurityFirewall resources.
 */
export interface SecurityFirewallState {
    /**
     * The CloudAMQP instance ID.
     */
    instanceId?: pulumi.Input<number>;
    /**
     * An array of rules, minimum of 1 needs to be configured. Each `rules` block consists of the field documented below.
     */
    rules?: pulumi.Input<pulumi.Input<inputs.SecurityFirewallRule>[]>;
}

/**
 * The set of arguments for constructing a SecurityFirewall resource.
 */
export interface SecurityFirewallArgs {
    /**
     * The CloudAMQP instance ID.
     */
    instanceId: pulumi.Input<number>;
    /**
     * An array of rules, minimum of 1 needs to be configured. Each `rules` block consists of the field documented below.
     */
    rules: pulumi.Input<pulumi.Input<inputs.SecurityFirewallRule>[]>;
}
