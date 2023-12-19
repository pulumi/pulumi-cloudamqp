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
    public sealed class GetPluginsCommunityPluginResult
    {
        public readonly string Description;
        public readonly string Name;
        public readonly string Require;
        public readonly int? Sleep;
        public readonly int? Timeout;

        [OutputConstructor]
        private GetPluginsCommunityPluginResult(
            string description,

            string name,

            string require,

            int? sleep,

            int? timeout)
        {
            Description = description;
            Name = name;
            Require = require;
            Sleep = sleep;
            Timeout = timeout;
        }
    }
}
