// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.CloudAmqp.Inputs
{

    public sealed class SecurityFirewallRuleGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Description name of the rule. e.g. Default.
        /// 
        /// Pre-defined services for RabbitMQ:
        /// 
        /// | Service | Port  |
        /// |---------|-------|
        /// | AMQP    |  5672 |
        /// | AMQPS   |  5671 |
        /// | HTTPS   |   443 |
        /// | MQTT    |  1883 |
        /// | MQTTS   |  8883 |
        /// | STOMP   | 61613 |
        /// | STOMPS  | 61614 |
        /// | STREAM  |  5552 |
        /// | STREAM_ |  5551 |
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// CIDR address: IP address with CIDR notation (e.g. 10.56.72.0/24)
        /// </summary>
        [Input("ip", required: true)]
        public Input<string> Ip { get; set; } = null!;

        [Input("ports")]
        private InputList<int>? _ports;

        /// <summary>
        /// Custom ports to be opened
        /// </summary>
        public InputList<int> Ports
        {
            get => _ports ?? (_ports = new InputList<int>());
            set => _ports = value;
        }

        [Input("services")]
        private InputList<string>? _services;

        /// <summary>
        /// Pre-defined service ports, see table below
        /// </summary>
        public InputList<string> Services
        {
            get => _services ?? (_services = new InputList<string>());
            set => _services = value;
        }

        public SecurityFirewallRuleGetArgs()
        {
        }
        public static new SecurityFirewallRuleGetArgs Empty => new SecurityFirewallRuleGetArgs();
    }
}
