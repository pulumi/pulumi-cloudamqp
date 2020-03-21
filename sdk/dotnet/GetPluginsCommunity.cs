// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.CloudAmqp
{
    public static partial class Invokes
    {
        public static Task<GetPluginsCommunityResult> GetPluginsCommunity(GetPluginsCommunityArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetPluginsCommunityResult>("cloudamqp:index/getPluginsCommunity:getPluginsCommunity", args ?? InvokeArgs.Empty, options.WithVersion());
    }

    public sealed class GetPluginsCommunityArgs : Pulumi.InvokeArgs
    {
        [Input("instanceId", required: true)]
        public int InstanceId { get; set; }

        [Input("plugins")]
        private List<Inputs.GetPluginsCommunityPluginsArgs>? _plugins;
        public List<Inputs.GetPluginsCommunityPluginsArgs> Plugins
        {
            get => _plugins ?? (_plugins = new List<Inputs.GetPluginsCommunityPluginsArgs>());
            set => _plugins = value;
        }

        public GetPluginsCommunityArgs()
        {
        }
    }

    [OutputType]
    public sealed class GetPluginsCommunityResult
    {
        public readonly int InstanceId;
        public readonly ImmutableArray<Outputs.GetPluginsCommunityPluginsResult> Plugins;
        /// <summary>
        /// id is the provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;

        [OutputConstructor]
        private GetPluginsCommunityResult(
            int instanceId,
            ImmutableArray<Outputs.GetPluginsCommunityPluginsResult> plugins,
            string id)
        {
            InstanceId = instanceId;
            Plugins = plugins;
            Id = id;
        }
    }

    namespace Inputs
    {

    public sealed class GetPluginsCommunityPluginsArgs : Pulumi.InvokeArgs
    {
        [Input("description")]
        public string? Description { get; set; }

        [Input("name")]
        public string? Name { get; set; }

        [Input("require")]
        public string? Require { get; set; }

        public GetPluginsCommunityPluginsArgs()
        {
        }
    }
    }

    namespace Outputs
    {

    [OutputType]
    public sealed class GetPluginsCommunityPluginsResult
    {
        public readonly string? Description;
        public readonly string? Name;
        public readonly string? Require;

        [OutputConstructor]
        private GetPluginsCommunityPluginsResult(
            string? description,
            string? name,
            string? require)
        {
            Description = description;
            Name = name;
            Require = require;
        }
    }
    }
}
