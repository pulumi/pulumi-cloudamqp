// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.CloudAmqp
{
    public static class GetPluginsCommunity
    {
        /// <summary>
        /// Use this data source to retrieve information about available community plugins for the CloudAMQP instance.
        /// 
        /// ## Example Usage
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using CloudAmqp = Pulumi.CloudAmqp;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var communitPlugins = CloudAmqp.GetPluginsCommunity.Invoke(new()
        ///     {
        ///         InstanceId = instance.Id,
        ///     });
        /// 
        /// });
        /// ```
        /// 
        /// ## Attributes reference
        /// 
        /// All attributes reference are computed
        /// 
        /// * `id`      - The identifier for this resource.
        /// * `plugins` - An array of community plugins. Each `plugins` block consists of the fields documented below.
        /// * `sleep` - (Optional) Configurable sleep time (seconds) for retries when requesting information
        /// about community plugins. Default set to 10 seconds. *Available from v1.29.0*
        /// * `timeout` - (Optional) - Configurable timeout time (seconds) for retries when requesting
        /// information about community plugins. Default set to 1800 seconds. *Available from v1.29.0*
        /// 
        /// ___
        /// 
        /// The `plugins` block consists of
        /// 
        /// * `name`        - The type of the recipient.
        /// * `require`     - Min. required Rabbit MQ version to be used.
        /// * `description` - Description of what the plugin does.
        /// 
        /// ## Dependency
        /// 
        /// This data source depends on CloudAMQP instance identifier, `cloudamqp_instance.instance.id`.
        /// </summary>
        public static Task<GetPluginsCommunityResult> InvokeAsync(GetPluginsCommunityArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetPluginsCommunityResult>("cloudamqp:index/getPluginsCommunity:getPluginsCommunity", args ?? new GetPluginsCommunityArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to retrieve information about available community plugins for the CloudAMQP instance.
        /// 
        /// ## Example Usage
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using CloudAmqp = Pulumi.CloudAmqp;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var communitPlugins = CloudAmqp.GetPluginsCommunity.Invoke(new()
        ///     {
        ///         InstanceId = instance.Id,
        ///     });
        /// 
        /// });
        /// ```
        /// 
        /// ## Attributes reference
        /// 
        /// All attributes reference are computed
        /// 
        /// * `id`      - The identifier for this resource.
        /// * `plugins` - An array of community plugins. Each `plugins` block consists of the fields documented below.
        /// * `sleep` - (Optional) Configurable sleep time (seconds) for retries when requesting information
        /// about community plugins. Default set to 10 seconds. *Available from v1.29.0*
        /// * `timeout` - (Optional) - Configurable timeout time (seconds) for retries when requesting
        /// information about community plugins. Default set to 1800 seconds. *Available from v1.29.0*
        /// 
        /// ___
        /// 
        /// The `plugins` block consists of
        /// 
        /// * `name`        - The type of the recipient.
        /// * `require`     - Min. required Rabbit MQ version to be used.
        /// * `description` - Description of what the plugin does.
        /// 
        /// ## Dependency
        /// 
        /// This data source depends on CloudAMQP instance identifier, `cloudamqp_instance.instance.id`.
        /// </summary>
        public static Output<GetPluginsCommunityResult> Invoke(GetPluginsCommunityInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetPluginsCommunityResult>("cloudamqp:index/getPluginsCommunity:getPluginsCommunity", args ?? new GetPluginsCommunityInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetPluginsCommunityArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The CloudAMQP instance identifier.
        /// </summary>
        [Input("instanceId", required: true)]
        public int InstanceId { get; set; }

        [Input("sleep")]
        public int? Sleep { get; set; }

        [Input("timeout")]
        public int? Timeout { get; set; }

        public GetPluginsCommunityArgs()
        {
        }
        public static new GetPluginsCommunityArgs Empty => new GetPluginsCommunityArgs();
    }

    public sealed class GetPluginsCommunityInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The CloudAMQP instance identifier.
        /// </summary>
        [Input("instanceId", required: true)]
        public Input<int> InstanceId { get; set; } = null!;

        [Input("sleep")]
        public Input<int>? Sleep { get; set; }

        [Input("timeout")]
        public Input<int>? Timeout { get; set; }

        public GetPluginsCommunityInvokeArgs()
        {
        }
        public static new GetPluginsCommunityInvokeArgs Empty => new GetPluginsCommunityInvokeArgs();
    }


    [OutputType]
    public sealed class GetPluginsCommunityResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly int InstanceId;
        public readonly ImmutableArray<Outputs.GetPluginsCommunityPluginResult> Plugins;
        public readonly int? Sleep;
        public readonly int? Timeout;

        [OutputConstructor]
        private GetPluginsCommunityResult(
            string id,

            int instanceId,

            ImmutableArray<Outputs.GetPluginsCommunityPluginResult> plugins,

            int? sleep,

            int? timeout)
        {
            Id = id;
            InstanceId = instanceId;
            Plugins = plugins;
            Sleep = sleep;
            Timeout = timeout;
        }
    }
}
