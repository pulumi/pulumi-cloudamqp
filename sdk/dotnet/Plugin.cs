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
    /// `cloudamqp_plugin` can be imported using the name argument of the resource together with CloudAMQP instance identifier. The name and identifier are CSV separated, see example below.
    /// 
    /// ```sh
    /// $ pulumi import cloudamqp:index/plugin:Plugin rabbitmq_management rabbitmq_management,&lt;instance_id&gt;`
    /// ```
    /// </summary>
    [CloudAmqpResourceType("cloudamqp:index/plugin:Plugin")]
    public partial class Plugin : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The description of the plugin.
        /// </summary>
        [Output("description")]
        public Output<string> Description { get; private set; } = null!;

        /// <summary>
        /// If the plugin is enabled
        /// </summary>
        [Output("enabled")]
        public Output<bool> Enabled { get; private set; } = null!;

        /// <summary>
        /// Instance identifier
        /// </summary>
        [Output("instanceId")]
        public Output<int> InstanceId { get; private set; } = null!;

        /// <summary>
        /// The name of the plugin
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Configurable sleep time in seconds between retries for plugins
        /// </summary>
        [Output("sleep")]
        public Output<int?> Sleep { get; private set; } = null!;

        /// <summary>
        /// Configurable timeout time in seconds for plugins
        /// </summary>
        [Output("timeout")]
        public Output<int?> Timeout { get; private set; } = null!;

        /// <summary>
        /// The version of the plugin.
        /// </summary>
        [Output("version")]
        public Output<string> Version { get; private set; } = null!;


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

    public sealed class PluginArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// If the plugin is enabled
        /// </summary>
        [Input("enabled", required: true)]
        public Input<bool> Enabled { get; set; } = null!;

        /// <summary>
        /// Instance identifier
        /// </summary>
        [Input("instanceId", required: true)]
        public Input<int> InstanceId { get; set; } = null!;

        /// <summary>
        /// The name of the plugin
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Configurable sleep time in seconds between retries for plugins
        /// </summary>
        [Input("sleep")]
        public Input<int>? Sleep { get; set; }

        /// <summary>
        /// Configurable timeout time in seconds for plugins
        /// </summary>
        [Input("timeout")]
        public Input<int>? Timeout { get; set; }

        public PluginArgs()
        {
        }
        public static new PluginArgs Empty => new PluginArgs();
    }

    public sealed class PluginState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The description of the plugin.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// If the plugin is enabled
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        /// <summary>
        /// Instance identifier
        /// </summary>
        [Input("instanceId")]
        public Input<int>? InstanceId { get; set; }

        /// <summary>
        /// The name of the plugin
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Configurable sleep time in seconds between retries for plugins
        /// </summary>
        [Input("sleep")]
        public Input<int>? Sleep { get; set; }

        /// <summary>
        /// Configurable timeout time in seconds for plugins
        /// </summary>
        [Input("timeout")]
        public Input<int>? Timeout { get; set; }

        /// <summary>
        /// The version of the plugin.
        /// </summary>
        [Input("version")]
        public Input<string>? Version { get; set; }

        public PluginState()
        {
        }
        public static new PluginState Empty => new PluginState();
    }
}
