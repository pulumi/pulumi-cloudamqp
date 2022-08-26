// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.CloudAmqp
{
    /// <summary>
    /// ## Import
    /// 
    /// Not possible to import this resource.
    /// </summary>
    [CloudAmqpResourceType("cloudamqp:index/vpcPeering:VpcPeering")]
    public partial class VpcPeering : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The CloudAMQP instance identifier.
        /// </summary>
        [Output("instanceId")]
        public Output<int?> InstanceId { get; private set; } = null!;

        /// <summary>
        /// Peering identifier created by AW peering request.
        /// </summary>
        [Output("peeringId")]
        public Output<string> PeeringId { get; private set; } = null!;

        /// <summary>
        /// Configurable sleep time (seconds) between retries for accepting or removing peering. Default set to 60 seconds.
        /// </summary>
        [Output("sleep")]
        public Output<int?> Sleep { get; private set; } = null!;

        /// <summary>
        /// VPC peering status
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// - Configurable timeout time (seconds) for accepting or removing peering. Default set to 3600 seconds.
        /// </summary>
        [Output("timeout")]
        public Output<int?> Timeout { get; private set; } = null!;

        /// <summary>
        /// The managed VPC identifier.
        /// </summary>
        [Output("vpcId")]
        public Output<string?> VpcId { get; private set; } = null!;


        /// <summary>
        /// Create a VpcPeering resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public VpcPeering(string name, VpcPeeringArgs args, CustomResourceOptions? options = null)
            : base("cloudamqp:index/vpcPeering:VpcPeering", name, args ?? new VpcPeeringArgs(), MakeResourceOptions(options, ""))
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

    public sealed class VpcPeeringArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The CloudAMQP instance identifier.
        /// </summary>
        [Input("instanceId")]
        public Input<int>? InstanceId { get; set; }

        /// <summary>
        /// Peering identifier created by AW peering request.
        /// </summary>
        [Input("peeringId", required: true)]
        public Input<string> PeeringId { get; set; } = null!;

        /// <summary>
        /// Configurable sleep time (seconds) between retries for accepting or removing peering. Default set to 60 seconds.
        /// </summary>
        [Input("sleep")]
        public Input<int>? Sleep { get; set; }

        /// <summary>
        /// - Configurable timeout time (seconds) for accepting or removing peering. Default set to 3600 seconds.
        /// </summary>
        [Input("timeout")]
        public Input<int>? Timeout { get; set; }

        /// <summary>
        /// The managed VPC identifier.
        /// </summary>
        [Input("vpcId")]
        public Input<string>? VpcId { get; set; }

        public VpcPeeringArgs()
        {
        }
        public static new VpcPeeringArgs Empty => new VpcPeeringArgs();
    }

    public sealed class VpcPeeringState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The CloudAMQP instance identifier.
        /// </summary>
        [Input("instanceId")]
        public Input<int>? InstanceId { get; set; }

        /// <summary>
        /// Peering identifier created by AW peering request.
        /// </summary>
        [Input("peeringId")]
        public Input<string>? PeeringId { get; set; }

        /// <summary>
        /// Configurable sleep time (seconds) between retries for accepting or removing peering. Default set to 60 seconds.
        /// </summary>
        [Input("sleep")]
        public Input<int>? Sleep { get; set; }

        /// <summary>
        /// VPC peering status
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// - Configurable timeout time (seconds) for accepting or removing peering. Default set to 3600 seconds.
        /// </summary>
        [Input("timeout")]
        public Input<int>? Timeout { get; set; }

        /// <summary>
        /// The managed VPC identifier.
        /// </summary>
        [Input("vpcId")]
        public Input<string>? VpcId { get; set; }

        public VpcPeeringState()
        {
        }
        public static new VpcPeeringState Empty => new VpcPeeringState();
    }
}
