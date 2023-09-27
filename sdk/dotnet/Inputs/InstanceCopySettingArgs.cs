// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.CloudAmqp.Inputs
{

    public sealed class InstanceCopySettingArgs : global::Pulumi.ResourceArgs
    {
        [Input("settings", required: true)]
        private InputList<string>? _settings;

        /// <summary>
        /// Array of one or more settings to be copied. Allowed values: [alarms, config, definitions, firewall, logs, metrics, plugins]
        /// 
        /// See more below, copy settings
        /// </summary>
        public InputList<string> Settings
        {
            get => _settings ?? (_settings = new InputList<string>());
            set => _settings = value;
        }

        /// <summary>
        /// Instance identifier of the CloudAMQP instance to copy the settings from.
        /// </summary>
        [Input("subscriptionId", required: true)]
        public Input<string> SubscriptionId { get; set; } = null!;

        public InstanceCopySettingArgs()
        {
        }
        public static new InstanceCopySettingArgs Empty => new InstanceCopySettingArgs();
    }
}
