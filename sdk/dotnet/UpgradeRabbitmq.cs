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
    /// This resource allows you to upgrade RabbitMQ version. Depending on initial versions of RabbitMQ and Erlang of the CloudAMQP instance, multiple runs may be needed to get to the latest or wanted version. Reason for this is certain supported RabbitMQ version will also automatically upgrade Erlang version.
    /// 
    /// There is three different ways to trigger the version upgrade
    /// 
    /// &gt; - Specify RabbitMQ version to upgrade to
    /// &gt; - Upgrade to latest RabbitMQ version
    /// &gt; - Old behaviour to upgrade to latest RabbitMQ version
    /// 
    /// See, below example usage for the difference.
    /// 
    /// Only available for dedicated subscription plans running ***RabbitMQ***.
    /// 
    /// ## Example Usage
    /// 
    /// &lt;details&gt;
    ///   &lt;summary&gt;
    ///     &lt;b&gt;
    ///       &lt;i&gt;Specify version upgrade, from v1.31.0&lt;/i&gt;
    ///     &lt;/b&gt;
    ///   &lt;/summary&gt;
    /// 
    /// Specify the version to upgrade to. List available upgradable versions, use [CloudAMQP API](https://docs.cloudamqp.com/cloudamqp_api.html#get-available-versions).
    /// After the upgrade finished, there can still be newer versions available.
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using CloudAmqp = Pulumi.CloudAmqp;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var instance = new CloudAmqp.Instance("instance", new()
    ///     {
    ///         Name = "rabbitmq-version-upgrade-test",
    ///         Plan = "bunny-1",
    ///         Region = "amazon-web-services::us-west-1",
    ///     });
    /// 
    ///     var upgrade = new CloudAmqp.UpgradeRabbitmq("upgrade", new()
    ///     {
    ///         InstanceId = instance.Id,
    ///         NewVersion = "3.13.2",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// &lt;/details&gt;
    /// 
    /// &lt;details&gt;
    ///   &lt;summary&gt;
    ///     &lt;b&gt;
    ///       &lt;i&gt;Upgrade to latest possible version, from v1.31.0&lt;/i&gt;
    ///     &lt;/b&gt;
    ///   &lt;/summary&gt;
    /// 
    /// This will upgrade RabbitMQ to the latest possible version detected by the data source `cloudamqp.getUpgradableVersions`.
    /// Multiple runs can be needed to upgrade the version even further.
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using CloudAmqp = Pulumi.CloudAmqp;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var instance = new CloudAmqp.Instance("instance", new()
    ///     {
    ///         Name = "rabbitmq-version-upgrade-test",
    ///         Plan = "bunny-1",
    ///         Region = "amazon-web-services::us-west-1",
    ///     });
    /// 
    ///     var upgradableVersions = CloudAmqp.GetUpgradableVersions.Invoke(new()
    ///     {
    ///         InstanceId = instance.Id,
    ///     });
    /// 
    ///     var upgrade = new CloudAmqp.UpgradeRabbitmq("upgrade", new()
    ///     {
    ///         InstanceId = instance.Id,
    ///         CurrentVersion = instance.RmqVersion,
    ///         NewVersion = upgradableVersions.Apply(getUpgradableVersionsResult =&gt; getUpgradableVersionsResult.NewRabbitmqVersion),
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// &lt;/details&gt;
    /// 
    /// &lt;details&gt;
    ///   &lt;summary&gt;
    ///     &lt;b&gt;
    ///       &lt;i&gt;Upgrade to latest possible version, before v1.31.0&lt;/i&gt;
    ///     &lt;/b&gt;
    ///   &lt;/summary&gt;
    /// 
    /// Old behaviour of the upgrading the RabbitMQ version. No longer recommended.
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using CloudAmqp = Pulumi.CloudAmqp;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     // Retrieve latest possible upgradable versions for RabbitMQ and Erlang
    ///     var versions = CloudAmqp.GetUpgradableVersions.Invoke(new()
    ///     {
    ///         InstanceId = instance.Id,
    ///     });
    /// 
    ///     // Invoke automatically upgrade to latest possible upgradable versions for RabbitMQ and Erlang
    ///     var upgrade = new CloudAmqp.UpgradeRabbitmq("upgrade", new()
    ///     {
    ///         InstanceId = instance.Id,
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using CloudAmqp = Pulumi.CloudAmqp;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     // Retrieve latest possible upgradable versions for RabbitMQ and Erlang
    ///     var versions = CloudAmqp.GetUpgradableVersions.Invoke(new()
    ///     {
    ///         InstanceId = instance.Id,
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// If newer version is still available to be upgradable in the data source, re-run again.
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using CloudAmqp = Pulumi.CloudAmqp;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     // Retrieve latest possible upgradable versions for RabbitMQ and Erlang
    ///     var versions = CloudAmqp.GetUpgradableVersions.Invoke(new()
    ///     {
    ///         InstanceId = instance.Id,
    ///     });
    /// 
    ///     // Invoke automatically upgrade to latest possible upgradable versions for RabbitMQ and Erlang
    ///     var upgrade = new CloudAmqp.UpgradeRabbitmq("upgrade", new()
    ///     {
    ///         InstanceId = instance.Id,
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// &lt;/details&gt;
    /// 
    /// ## Important Upgrade Information
    /// 
    /// &gt; - All single node upgrades will require some downtime since RabbitMQ needs a restart.
    /// &gt; - From RabbitMQ version 3.9, rolling upgrades between minor versions (e.g. 3.9 to 3.10), in a multi-node cluster are possible without downtime. This means that one node is upgraded at a time while the other nodes are still running. For versions older than 3.9, patch version upgrades (e.g. 3.8.x to 3.8.y) are possible without downtime in a multi-node cluster, but minor version upgrades will require downtime.
    /// &gt; - Auto delete queues (queues that are marked AD) will be deleted during the update.
    /// &gt; - Any custom plugins support has installed on your behalf will be disabled and you need to contact support@cloudamqp.com and ask to have them re-installed.
    /// &gt; - TLS 1.0 and 1.1 will not be supported after the update.
    /// 
    /// ## Multiple runs
    /// 
    /// Depending on initial versions of RabbitMQ and Erlang of the CloudAMQP instance, multiple runs may be needed to get to the latest or wanted version.
    /// 
    /// Example steps needed when starting at RabbitMQ version 3.12.2
    /// 
    /// |  Version         | Supported upgrading versions              | Min version to upgrade Erlang |
    /// |------------------|-------------------------------------------|-------------------------------|
    /// | 3.12.2           | 3.12.4, 3.12.6, 3.12.10, 3.12.12, 3.12.13 | 3.12.13                       |
    /// | 3.12.13          | 3.13.2                                    | 3.13.2                        |
    /// | 3.13.2           | -                                         | -                             |
    /// 
    /// ## Import
    /// 
    /// Not possible to import this resource.
    /// </summary>
    [CloudAmqpResourceType("cloudamqp:index/upgradeRabbitmq:UpgradeRabbitmq")]
    public partial class UpgradeRabbitmq : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Helper argument to change upgrade behaviour to latest possible version
        /// </summary>
        [Output("currentVersion")]
        public Output<string?> CurrentVersion { get; private set; } = null!;

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
        /// Helper argument to change upgrade behaviour to latest possible version
        /// </summary>
        [Input("currentVersion")]
        public Input<string>? CurrentVersion { get; set; }

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

        public UpgradeRabbitmqArgs()
        {
        }
        public static new UpgradeRabbitmqArgs Empty => new UpgradeRabbitmqArgs();
    }

    public sealed class UpgradeRabbitmqState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Helper argument to change upgrade behaviour to latest possible version
        /// </summary>
        [Input("currentVersion")]
        public Input<string>? CurrentVersion { get; set; }

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

        public UpgradeRabbitmqState()
        {
        }
        public static new UpgradeRabbitmqState Empty => new UpgradeRabbitmqState();
    }
}
