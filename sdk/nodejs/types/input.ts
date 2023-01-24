// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";

export interface SecurityFirewallRule {
    /**
     * Description name of the rule. e.g. Default.
     */
    description?: pulumi.Input<string>;
    /**
     * CIDR address: IP address with CIDR notation (e.g. 10.56.72.0/24)
     */
    ip: pulumi.Input<string>;
    /**
     * Custom ports to be opened
     */
    ports?: pulumi.Input<pulumi.Input<number>[]>;
    /**
     * Pre-defined service ports, see table below
     */
    services?: pulumi.Input<pulumi.Input<string>[]>;
}
