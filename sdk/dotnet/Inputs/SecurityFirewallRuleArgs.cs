// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.CloudAmqp.Inputs
{

    public sealed class SecurityFirewallRuleArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Description name of the rule. e.g. Default.
        /// 
        /// Pre-defined services for RabbitMQ:
        /// &lt;table&gt;
        /// &lt;thead&gt;
        /// &lt;tr&gt;
        /// &lt;th&gt;Service name&lt;/th&gt;
        /// &lt;th&gt;Port&lt;/th&gt;
        /// &lt;/tr&gt;
        /// &lt;/thead&gt;
        /// &lt;tbody&gt;
        /// &lt;tr&gt;
        /// &lt;td&gt;AMQP&lt;/td&gt;
        /// &lt;td&gt;5672&lt;/td&gt;
        /// &lt;/tr&gt;
        /// &lt;tr&gt;
        /// &lt;td&gt;AMQPS&lt;/td&gt;
        /// &lt;td&gt;5671&lt;/td&gt;
        /// &lt;/tr&gt;
        /// &lt;tr&gt;
        /// &lt;td&gt;HTTPS&lt;/td&gt;
        /// &lt;td&gt;443&lt;/td&gt;
        /// &lt;/tr&gt;
        /// &lt;tr&gt;
        /// &lt;td&gt;MQTT&lt;/td&gt;
        /// &lt;td&gt;1883&lt;/td&gt;
        /// &lt;/tr&gt;
        /// &lt;tr&gt;
        /// &lt;td&gt;MQTTS&lt;/td&gt;
        /// &lt;td&gt;8883&lt;/td&gt;
        /// &lt;/tr&gt;
        /// &lt;tr&gt;
        /// &lt;td&gt;STOMP&lt;/td&gt;
        /// &lt;td&gt;61613&lt;/td&gt;
        /// &lt;/tr&gt;
        /// &lt;tr&gt;
        /// &lt;td&gt;STOMPS&lt;/td&gt;
        /// &lt;td&gt;61614&lt;/td&gt;
        /// &lt;/tr&gt;
        /// &lt;tr&gt;
        /// &lt;td&gt;STREAM&lt;/td&gt;
        /// &lt;td&gt;5552&lt;/td&gt;
        /// &lt;/tr&gt;
        /// &lt;tr&gt;
        /// &lt;td&gt;STREAM_SSL&lt;/td&gt;
        /// &lt;td&gt;5551&lt;/td&gt;
        /// &lt;/tr&gt;
        /// &lt;/tbody&gt;
        /// &lt;/table&gt;
        /// 
        /// Pre-defined services for LavinMQ:
        /// &lt;table&gt;
        /// &lt;thead&gt;
        /// &lt;tr&gt;
        /// &lt;th&gt;Service name&lt;/th&gt;
        /// &lt;th&gt;Port&lt;/th&gt;
        /// &lt;/tr&gt;
        /// &lt;/thead&gt;
        /// &lt;tbody&gt;
        /// &lt;tr&gt;
        /// &lt;td&gt;AMQP&lt;/td&gt;
        /// &lt;td&gt;5672&lt;/td&gt;
        /// &lt;/tr&gt;
        /// &lt;tr&gt;
        /// &lt;td&gt;AMQPS&lt;/td&gt;
        /// &lt;td&gt;5671&lt;/td&gt;
        /// &lt;/tr&gt;
        /// &lt;tr&gt;
        /// &lt;td&gt;HTTPS&lt;/td&gt;
        /// &lt;td&gt;443&lt;/td&gt;
        /// &lt;/tr&gt;
        /// &lt;/tbody&gt;
        /// &lt;/table&gt;
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

        public SecurityFirewallRuleArgs()
        {
        }
        public static new SecurityFirewallRuleArgs Empty => new SecurityFirewallRuleArgs();
    }
}
