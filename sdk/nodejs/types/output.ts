// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";

export interface ExtraDiskSizeNode {
    additionalDiskSize: number;
    diskSize: number;
    name: string;
}

export interface GetAccountInstance {
    id: number;
    name: string;
    plan: string;
    region: string;
    tags?: string[];
}

export interface GetAccountVpcsVpc {
    id: number;
    name: string;
    region: string;
    subnet: string;
    tags?: string[];
    vpcName: string;
}

export interface GetNodesNode {
    additionalDiskSize: number;
    configured: boolean;
    diskSize: number;
    erlangVersion: string;
    hipe: boolean;
    hostname: string;
    name: string;
    rabbitmqVersion: string;
    running: boolean;
}

export interface GetPluginsCommunityPlugin {
    description: string;
    name: string;
    require: string;
}

export interface GetPluginsPlugin {
    description: string;
    enabled: boolean;
    name: string;
    version: string;
}

export interface SecurityFirewallRule {
    /**
     * Description name of the rule. e.g. Default.
     *
     * Pre-defined services for RabbitMQ:
     *
     * | Service name | Port  |
     * |--------------|-------|
     * | AMQP         | 5672  |
     * | AMQPS        | 5671  |
     * | HTTPS        | 443   |
     * | MQTT         | 1883  |
     * | MQTTS        | 8883  |
     * | STOMP        | 61613 |
     * | STOMPS       | 61614 |
     * | STREAM       | 5552  |
     * | STREAM_SSL   | 5551  |
     *
     * Pre-defined services for LavinMQ:
     *
     * | Service name | Port  |
     * |--------------|-------|
     * | AMQP         | 5672  |
     * | AMQPS        | 5671  |
     * | HTTPS        | 443   |
     */
    description?: string;
    /**
     * CIDR address: IP address with CIDR notation (e.g. 10.56.72.0/24)
     */
    ip: string;
    /**
     * Custom ports to be opened
     */
    ports?: number[];
    /**
     * Pre-defined service ports, see table below
     */
    services?: string[];
}

