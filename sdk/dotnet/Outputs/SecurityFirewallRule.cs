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
        public readonly string Ip;
        public readonly ImmutableArray<int> Ports;
        public readonly ImmutableArray<string> Services;

        [OutputConstructor]
        private SecurityFirewallRule(
            string ip,

            ImmutableArray<int> ports,

            ImmutableArray<string> services)
        {
            Ip = ip;
            Ports = ports;
            Services = services;
        }
    }
}
