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
    public sealed class InstanceCopySetting
    {
        /// <summary>
        /// Array of one or more settings to be copied. Allowed values:
        /// [alarms, config, definitions, firewall, logs, metrics, plugins]
        /// 
        /// See more below, [copy settings].
        /// </summary>
        public readonly ImmutableArray<string> Settings;
        /// <summary>
        /// Instance identifier of the CloudAMQP instance to copy the settings
        /// from.
        /// </summary>
        public readonly string SubscriptionId;

        [OutputConstructor]
        private InstanceCopySetting(
            ImmutableArray<string> settings,

            string subscriptionId)
        {
            Settings = settings;
            SubscriptionId = subscriptionId;
        }
    }
}
