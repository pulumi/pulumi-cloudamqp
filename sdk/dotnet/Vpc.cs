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
    /// This resource allows you to manage standalone VPC.
    /// 
    /// New Cloudamqp instances can be added to the managed VPC. Set the instance *vpc_id* attribute to the managed vpc identifier , see example below, when creating the instance.
    /// 
    /// Only available for dedicated subscription plans.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using CloudAmqp = Pulumi.CloudAmqp;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         // Managed VPC resource
    ///         var vpc = new CloudAmqp.Vpc("vpc", new CloudAmqp.VpcArgs
    ///         {
    ///             Region = "amazon-web-services::us-east-1",
    ///             Subnet = "10.56.72.0/24",
    ///             Tags = {},
    ///         });
    ///         //  New instance, need to be created with a vpc
    ///         var instance = new CloudAmqp.Instance("instance", new CloudAmqp.InstanceArgs
    ///         {
    ///             Plan = "bunny-1",
    ///             Region = "amazon-web-services::us-east-1",
    ///             Nodes = 1,
    ///             Tags = {},
    ///             RmqVersion = "3.9.13",
    ///             VpcId = cloudamq_vpc.Vpc.Id,
    ///         });
    ///         var vpcInfo = CloudAmqp.GetVpcInfo.Invoke(new CloudAmqp.GetVpcInfoInvokeArgs
    ///         {
    ///             VpcId = vpc.Id,
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// `cloudamqp_vpc` can be imported using the CloudAMQP VPC identifier.
    /// 
    /// ```sh
    ///  $ pulumi import cloudamqp:index/vpc:Vpc &lt;resource_name&gt; &lt;vpc_id&gt;`
    /// ```
    /// 
    ///  To retrieve the identifier for a VPC, either use [CloudAMQP customer API](https://docs.cloudamqp.com/#list-vpcs). Or use the data source [`cloudamqp_account_vpcs`](https://registry.terraform.io/providers/cloudamqp/cloudamqp/latest/docs/data-sources/account_vpcs) to list all available standalone VPCs for an account.
    /// </summary>
    [CloudAmqpResourceType("cloudamqp:index/vpc:Vpc")]
    public partial class Vpc : Pulumi.CustomResource
    {
        /// <summary>
        /// The name of the VPC.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The hosted region for the managed standalone VPC
        /// </summary>
        [Output("region")]
        public Output<string> Region { get; private set; } = null!;

        /// <summary>
        /// The VPC subnet
        /// </summary>
        [Output("subnet")]
        public Output<string> Subnet { get; private set; } = null!;

        /// <summary>
        /// Tag the VPC with optional tags
        /// </summary>
        [Output("tags")]
        public Output<ImmutableArray<string>> Tags { get; private set; } = null!;

        /// <summary>
        /// VPC name given when hosted at the cloud provider
        /// </summary>
        [Output("vpcName")]
        public Output<string> VpcName { get; private set; } = null!;


        /// <summary>
        /// Create a Vpc resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Vpc(string name, VpcArgs args, CustomResourceOptions? options = null)
            : base("cloudamqp:index/vpc:Vpc", name, args ?? new VpcArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Vpc(string name, Input<string> id, VpcState? state = null, CustomResourceOptions? options = null)
            : base("cloudamqp:index/vpc:Vpc", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Vpc resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Vpc Get(string name, Input<string> id, VpcState? state = null, CustomResourceOptions? options = null)
        {
            return new Vpc(name, id, state, options);
        }
    }

    public sealed class VpcArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the VPC.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The hosted region for the managed standalone VPC
        /// </summary>
        [Input("region", required: true)]
        public Input<string> Region { get; set; } = null!;

        /// <summary>
        /// The VPC subnet
        /// </summary>
        [Input("subnet", required: true)]
        public Input<string> Subnet { get; set; } = null!;

        [Input("tags")]
        private InputList<string>? _tags;

        /// <summary>
        /// Tag the VPC with optional tags
        /// </summary>
        public InputList<string> Tags
        {
            get => _tags ?? (_tags = new InputList<string>());
            set => _tags = value;
        }

        public VpcArgs()
        {
        }
    }

    public sealed class VpcState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the VPC.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The hosted region for the managed standalone VPC
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// The VPC subnet
        /// </summary>
        [Input("subnet")]
        public Input<string>? Subnet { get; set; }

        [Input("tags")]
        private InputList<string>? _tags;

        /// <summary>
        /// Tag the VPC with optional tags
        /// </summary>
        public InputList<string> Tags
        {
            get => _tags ?? (_tags = new InputList<string>());
            set => _tags = value;
        }

        /// <summary>
        /// VPC name given when hosted at the cloud provider
        /// </summary>
        [Input("vpcName")]
        public Input<string>? VpcName { get; set; }

        public VpcState()
        {
        }
    }
}
