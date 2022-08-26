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
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Source ip and netmask for the rule. (e.g. 10.56.72.0/24)
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
