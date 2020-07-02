// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";

export interface GetNodesNode {
    erlangVersion: string;
    hipe: boolean;
    hostname: string;
    name: string;
    rabbitmqVersion: string;
    running: boolean;
}

export interface GetPluginsCommunityPlugin {
    description?: string;
    name?: string;
    require?: string;
}

export interface GetPluginsPlugin {
    description?: string;
    enabled?: boolean;
    name?: string;
    version?: string;
}

export interface SecurityFirewallRule {
    /**
     * Description name of the rule. e.g. Default.
     */
    description?: string;
    /**
     * Source ip and netmask for the rule. (e.g. 10.56.72.0/24)
     */
    ip: string;
    /**
     * Custom ports to be opened
     */
    ports?: number[];
    /**
     * Pre-defined service ports
     */
    services?: string[];
}
