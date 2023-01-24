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
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using Pulumi;
    /// using CloudAmqp = Pulumi.CloudAmqp;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var rabbitmqDelayedMessageExchange = new CloudAmqp.PluginCommunity("rabbitmqDelayedMessageExchange", new()
    ///     {
    ///         InstanceId = cloudamqp_instance.Instance_01.Id,
    ///         Enabled = true,
    ///     });
    /// 
    /// });
    /// ```
    /// ## Depedency
    /// 
    /// This resource depends on CloudAMQP instance identifier, `cloudamqp_instance.instance.id`.
    /// 
    /// ## Import
    /// 
    /// `cloudamqp_plugin` can be imported using the name argument of the resource together with CloudAMQP instance identifier. The name and identifier are CSV separated, see example below.
    /// 
    /// ```sh
    ///  $ pulumi import cloudamqp:index/pluginCommunity:PluginCommunity &lt;resource_name&gt; &lt;plugin_name&gt;,&lt;instance_id&gt;`
    /// ```
    /// </summary>
    [CloudAmqpResourceType("cloudamqp:index/pluginCommunity:PluginCommunity")]
    public partial class PluginCommunity : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The description of the plugin.
        /// </summary>
        [Output("description")]
        public Output<string> Description { get; private set; } = null!;

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
        /// The name of the Rabbit MQ community plugin.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Required version of RabbitMQ.
        /// </summary>
        [Output("require")]
        public Output<string> Require { get; private set; } = null!;


        /// <summary>
        /// Create a PluginCommunity resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public PluginCommunity(string name, PluginCommunityArgs args, CustomResourceOptions? options = null)
            : base("cloudamqp:index/pluginCommunity:PluginCommunity", name, args ?? new PluginCommunityArgs(), MakeResourceOptions(options, ""))
        {
        }

        private PluginCommunity(string name, Input<string> id, PluginCommunityState? state = null, CustomResourceOptions? options = null)
            : base("cloudamqp:index/pluginCommunity:PluginCommunity", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing PluginCommunity resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static PluginCommunity Get(string name, Input<string> id, PluginCommunityState? state = null, CustomResourceOptions? options = null)
        {
            return new PluginCommunity(name, id, state, options);
        }
    }

    public sealed class PluginCommunityArgs : global::Pulumi.ResourceArgs
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
        /// The name of the Rabbit MQ community plugin.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public PluginCommunityArgs()
        {
        }
        public static new PluginCommunityArgs Empty => new PluginCommunityArgs();
    }

    public sealed class PluginCommunityState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The description of the plugin.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

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
        /// The name of the Rabbit MQ community plugin.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Required version of RabbitMQ.
        /// </summary>
        [Input("require")]
        public Input<string>? Require { get; set; }

        public PluginCommunityState()
        {
        }
        public static new PluginCommunityState Empty => new PluginCommunityState();
    }
}
