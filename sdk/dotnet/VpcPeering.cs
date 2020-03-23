// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.CloudAmqp
{
    public partial class VpcPeering : Pulumi.CustomResource
    {
        /// <summary>
        /// Instance identifier
        /// </summary>
        [Output("instanceId")]
        public Output<int> InstanceId { get; private set; } = null!;

        /// <summary>
        /// VPC peering identifier
        /// </summary>
        [Output("peeringId")]
        public Output<string> PeeringId { get; private set; } = null!;

        /// <summary>
        /// VPC peering status
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;


        /// <summary>
        /// Create a VpcPeering resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public VpcPeering(string name, VpcPeeringArgs args, CustomResourceOptions? options = null)
            : base("cloudamqp:index/vpcPeering:VpcPeering", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private VpcPeering(string name, Input<string> id, VpcPeeringState? state = null, CustomResourceOptions? options = null)
            : base("cloudamqp:index/vpcPeering:VpcPeering", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing VpcPeering resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static VpcPeering Get(string name, Input<string> id, VpcPeeringState? state = null, CustomResourceOptions? options = null)
        {
            return new VpcPeering(name, id, state, options);
        }
    }

    public sealed class VpcPeeringArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Instance identifier
        /// </summary>
        [Input("instanceId", required: true)]
        public Input<int> InstanceId { get; set; } = null!;

        /// <summary>
        /// VPC peering identifier
        /// </summary>
        [Input("peeringId", required: true)]
        public Input<string> PeeringId { get; set; } = null!;

        public VpcPeeringArgs()
        {
        }
    }

    public sealed class VpcPeeringState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Instance identifier
        /// </summary>
        [Input("instanceId")]
        public Input<int>? InstanceId { get; set; }

        /// <summary>
        /// VPC peering identifier
        /// </summary>
        [Input("peeringId")]
        public Input<string>? PeeringId { get; set; }

        /// <summary>
        /// VPC peering status
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        public VpcPeeringState()
        {
        }
    }
}
