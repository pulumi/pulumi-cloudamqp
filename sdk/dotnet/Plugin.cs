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
    /// This resource allows you to enable or disable Rabbit MQ plugins.
    /// 
    /// Only available for dedicated subscription plans.
    /// 
    /// ⚠️  From our go API wrapper [v1.4.0](https://github.com/84codes/go-api/releases/tag/v1.4.0) there is support for multiple retries when requesting information about plugins. This was introduced to avoid `ReadPlugin error 400: Timeout talking to backend`.
    /// 
    /// **Enable multiple plugins:** Rabbit MQ can only change one plugin at a time. It will fail if multiple plugins resources are used, unless by creating dependencies with `depend_on` between the resources. Once one plugin has been enabled, the other will continue. See example below.
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
    ///         var rabbitmqTop = new CloudAmqp.Plugin("rabbitmqTop", new CloudAmqp.PluginArgs
    ///         {
    ///             InstanceId = cloudamqp_instance.Instance.Id,
    ///             Enabled = true,
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// **Enable multiple plugins**
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using CloudAmqp = Pulumi.CloudAmqp;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var rabbitmqTop = new CloudAmqp.Plugin("rabbitmqTop", new CloudAmqp.PluginArgs
    ///         {
    ///             InstanceId = cloudamqp_instance.Instance.Id,
    ///             Enabled = true,
    ///         });
    ///         var rabbitmqAmqp10 = new CloudAmqp.Plugin("rabbitmqAmqp10", new CloudAmqp.PluginArgs
    ///         {
    ///             InstanceId = cloudamqp_instance.Instance.Id,
    ///             Enabled = true,
    ///         }, new CustomResourceOptions
    ///         {
    ///             DependsOn = 
    ///             {
    ///                 rabbitmqTop,
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// ## Dependency
    /// 
    /// This resource depends on CloudAMQP instance identifier, `cloudamqp_instance.instance.id`.
    /// 
    /// If multiple plugins should be enable, create dependencies between the plugin resources. See example above.
    /// 
    /// ## Import
    /// 
    /// `cloudamqp_plugin` can be imported using the name argument of the resource together with CloudAMQP instance identifier. The name and identifier are CSV separated, see example below.
    /// 
    /// ```sh
    ///  $ pulumi import cloudamqp:index/plugin:Plugin rabbitmq_management rabbitmq_management,&lt;instance_id&gt;`
    /// ```
    /// </summary>
    [CloudAmqpResourceType("cloudamqp:index/plugin:Plugin")]
    public partial class Plugin : Pulumi.CustomResource
    {
        /// <summary>
        /// Enable or disable the plugins.
        /// </summary>
        [Output("enabled")]
        public Output<bool> Enabled { get; private set; } = null!;

        /// <summary>
        /// The CloudAMQP instance ID.
        /// </summary>
        [Output("instanceId")]
        public Output<int> InstanceId { get; private set; } = null!;

        /// <summary>
        /// The name of the Rabbit MQ plugin.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;


        /// <summary>
        /// Create a Plugin resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Plugin(string name, PluginArgs args, CustomResourceOptions? options = null)
            : base("cloudamqp:index/plugin:Plugin", name, args ?? new PluginArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Plugin(string name, Input<string> id, PluginState? state = null, CustomResourceOptions? options = null)
            : base("cloudamqp:index/plugin:Plugin", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Plugin resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Plugin Get(string name, Input<string> id, PluginState? state = null, CustomResourceOptions? options = null)
        {
            return new Plugin(name, id, state, options);
        }
    }

    public sealed class PluginArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Enable or disable the plugins.
        /// </summary>
        [Input("enabled", required: true)]
        public Input<bool> Enabled { get; set; } = null!;

        /// <summary>
        /// The CloudAMQP instance ID.
        /// </summary>
        [Input("instanceId", required: true)]
        public Input<int> InstanceId { get; set; } = null!;

        /// <summary>
        /// The name of the Rabbit MQ plugin.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public PluginArgs()
        {
        }
    }

    public sealed class PluginState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Enable or disable the plugins.
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        /// <summary>
        /// The CloudAMQP instance ID.
        /// </summary>
        [Input("instanceId")]
        public Input<int>? InstanceId { get; set; }

        /// <summary>
        /// The name of the Rabbit MQ plugin.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public PluginState()
        {
        }
    }
}
