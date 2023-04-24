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
    public sealed class ExtraDiskSizeNode
    {
        public readonly int? AdditionalDiskSize;
        public readonly int? DiskSize;
        public readonly string? Name;

        [OutputConstructor]
        private ExtraDiskSizeNode(
            int? additionalDiskSize,

            int? diskSize,

            string? name)
        {
            AdditionalDiskSize = additionalDiskSize;
            DiskSize = diskSize;
            Name = name;
        }
    }
}