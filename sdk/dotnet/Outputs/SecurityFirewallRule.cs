// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.CloudAmqp.Outputs
{

    [OutputType]
    public sealed class SecurityFirewallRule
    {
        /// <summary>
        /// Description name of the rule. e.g. Default.
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// Source ip and netmask for the rule. (e.g. 10.56.72.0/24)
        /// </summary>
        public readonly string Ip;
        /// <summary>
        /// Custom ports to be opened
        /// </summary>
        public readonly ImmutableArray<int> Ports;
        /// <summary>
        /// Pre-defined service ports, see table below
        /// </summary>
        public readonly ImmutableArray<string> Services;

        [OutputConstructor]
        private SecurityFirewallRule(
            string? description,

            string ip,

            ImmutableArray<int> ports,

            ImmutableArray<string> services)
        {
            Description = description;
            Ip = ip;
            Ports = ports;
            Services = services;
        }
    }
}
