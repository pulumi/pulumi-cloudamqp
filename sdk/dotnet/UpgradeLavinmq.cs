// *** WARNING: this file was generated by pulumi-language-dotnet. ***
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
    [CloudAmqpResourceType("cloudamqp:index/upgradeLavinmq:UpgradeLavinmq")]
    public partial class UpgradeLavinmq : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The CloudAMQP instance identifier
        /// </summary>
        [Output("instanceId")]
        public Output<int> InstanceId { get; private set; } = null!;

        /// <summary>
        /// The new version to upgrade to
        /// </summary>
        [Output("newVersion")]
        public Output<string?> NewVersion { get; private set; } = null!;


        /// <summary>
        /// Create a UpgradeLavinmq resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public UpgradeLavinmq(string name, UpgradeLavinmqArgs args, CustomResourceOptions? options = null)
            : base("cloudamqp:index/upgradeLavinmq:UpgradeLavinmq", name, args ?? new UpgradeLavinmqArgs(), MakeResourceOptions(options, ""))
        {
        }

        private UpgradeLavinmq(string name, Input<string> id, UpgradeLavinmqState? state = null, CustomResourceOptions? options = null)
            : base("cloudamqp:index/upgradeLavinmq:UpgradeLavinmq", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing UpgradeLavinmq resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static UpgradeLavinmq Get(string name, Input<string> id, UpgradeLavinmqState? state = null, CustomResourceOptions? options = null)
        {
            return new UpgradeLavinmq(name, id, state, options);
        }
    }

    public sealed class UpgradeLavinmqArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The CloudAMQP instance identifier
        /// </summary>
        [Input("instanceId", required: true)]
        public Input<int> InstanceId { get; set; } = null!;

        /// <summary>
        /// The new version to upgrade to
        /// </summary>
        [Input("newVersion")]
        public Input<string>? NewVersion { get; set; }

        public UpgradeLavinmqArgs()
        {
        }
        public static new UpgradeLavinmqArgs Empty => new UpgradeLavinmqArgs();
    }

    public sealed class UpgradeLavinmqState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The CloudAMQP instance identifier
        /// </summary>
        [Input("instanceId")]
        public Input<int>? InstanceId { get; set; }

        /// <summary>
        /// The new version to upgrade to
        /// </summary>
        [Input("newVersion")]
        public Input<string>? NewVersion { get; set; }

        public UpgradeLavinmqState()
        {
        }
        public static new UpgradeLavinmqState Empty => new UpgradeLavinmqState();
    }
}
