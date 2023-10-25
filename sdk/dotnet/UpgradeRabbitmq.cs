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
    /// This resource allows you to automatically upgrade to the latest possible upgradable versions for RabbitMQ and Erlang. Depending on initial versions of RabbitMQ and Erlang of the CloudAMQP instance, multiple runs may be needed to get to the latest versions. After completed upgrade, check data source `cloudamqp.getUpgradableVersions` to see if newer versions is available. Then delete `cloudamqp.UpgradeRabbitmq` and create it again to invoke the upgrade.
    /// 
    /// &gt; **Important Upgrade Information**
    /// &gt; - All single node upgrades will require some downtime since RabbitMQ needs a restart.
    /// &gt; - From RabbitMQ version 3.9, rolling upgrades between minor versions (e.g. 3.9 to 3.10), in a multi-node cluster are possible without downtime. This means that one node is upgraded at a time while the other nodes are still running. For versions older than 3.9, patch version upgrades (e.g. 3.8.x to 3.8.y) are possible without downtime in a multi-node cluster, but minor version upgrades will require downtime.
    /// &gt; - Auto delete queues (queues that are marked AD) will be deleted during the update.
    /// &gt; - Any custom plugins support has installed on your behalf will be disabled and you need to contact support@cloudamqp.com and ask to have them re-installed.
    /// &gt; - TLS 1.0 and 1.1 will not be supported after the update.
    /// 
    /// Only available for dedicated subscription plans running ***RabbitMQ***.
    /// 
    /// ## Import
    /// 
    /// Not possible to import this resource.
    /// </summary>
    [CloudAmqpResourceType("cloudamqp:index/upgradeRabbitmq:UpgradeRabbitmq")]
    public partial class UpgradeRabbitmq : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The CloudAMQP instance identifier
        /// </summary>
        [Output("instanceId")]
        public Output<int> InstanceId { get; private set; } = null!;


        /// <summary>
        /// Create a UpgradeRabbitmq resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public UpgradeRabbitmq(string name, UpgradeRabbitmqArgs args, CustomResourceOptions? options = null)
            : base("cloudamqp:index/upgradeRabbitmq:UpgradeRabbitmq", name, args ?? new UpgradeRabbitmqArgs(), MakeResourceOptions(options, ""))
        {
        }

        private UpgradeRabbitmq(string name, Input<string> id, UpgradeRabbitmqState? state = null, CustomResourceOptions? options = null)
            : base("cloudamqp:index/upgradeRabbitmq:UpgradeRabbitmq", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing UpgradeRabbitmq resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static UpgradeRabbitmq Get(string name, Input<string> id, UpgradeRabbitmqState? state = null, CustomResourceOptions? options = null)
        {
            return new UpgradeRabbitmq(name, id, state, options);
        }
    }

    public sealed class UpgradeRabbitmqArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The CloudAMQP instance identifier
        /// </summary>
        [Input("instanceId", required: true)]
        public Input<int> InstanceId { get; set; } = null!;

        public UpgradeRabbitmqArgs()
        {
        }
        public static new UpgradeRabbitmqArgs Empty => new UpgradeRabbitmqArgs();
    }

    public sealed class UpgradeRabbitmqState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The CloudAMQP instance identifier
        /// </summary>
        [Input("instanceId")]
        public Input<int>? InstanceId { get; set; }

        public UpgradeRabbitmqState()
        {
        }
        public static new UpgradeRabbitmqState Empty => new UpgradeRabbitmqState();
    }
}
