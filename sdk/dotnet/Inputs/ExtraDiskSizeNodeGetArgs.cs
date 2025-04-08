// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.CloudAmqp.Inputs
{

    public sealed class ExtraDiskSizeNodeGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Additional added disk size
        /// </summary>
        [Input("additionalDiskSize")]
        public Input<int>? AdditionalDiskSize { get; set; }

        /// <summary>
        /// Subscription plan disk size
        /// </summary>
        [Input("diskSize")]
        public Input<int>? DiskSize { get; set; }

        /// <summary>
        /// Name of the node.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public ExtraDiskSizeNodeGetArgs()
        {
        }
        public static new ExtraDiskSizeNodeGetArgs Empty => new ExtraDiskSizeNodeGetArgs();
    }
}
