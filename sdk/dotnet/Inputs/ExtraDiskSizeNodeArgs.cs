// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.CloudAmqp.Inputs
{

    public sealed class ExtraDiskSizeNodeArgs : global::Pulumi.ResourceArgs
    {
        [Input("additionalDiskSize")]
        public Input<int>? AdditionalDiskSize { get; set; }

        [Input("diskSize")]
        public Input<int>? DiskSize { get; set; }

        [Input("name")]
        public Input<string>? Name { get; set; }

        public ExtraDiskSizeNodeArgs()
        {
        }
        public static new ExtraDiskSizeNodeArgs Empty => new ExtraDiskSizeNodeArgs();
    }
}