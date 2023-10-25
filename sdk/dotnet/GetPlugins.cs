// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.CloudAmqp
{
    public static class GetPlugins
    {
        /// <summary>
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using CloudAmqp = Pulumi.CloudAmqp;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var plugins = CloudAmqp.GetPlugins.Invoke(new()
        ///     {
        ///         InstanceId = cloudamqp_instance.Instance.Id,
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// ## Attributes reference
        /// 
        /// All attributes reference are computed
        /// 
        /// * `id`      - The identifier for this resource.
        /// * `plugins` - An array of plugins. Each `plugins` block consists of the fields documented below.
        /// 
        /// ___
        /// 
        /// The `plugins` block consist of
        /// 
        /// * `name`        - The type of the recipient.
        /// * `version`     - Rabbit MQ version that the plugins are shipped with.
        /// * `description` - Description of what the plugin does.
        /// * `enabled`     - Enable or disable information for the plugin.
        /// 
        /// ## Dependency
        /// 
        /// This data source depends on CloudAMQP instance identifier, `cloudamqp_instance.instance.id`.
        /// </summary>
        public static Task<GetPluginsResult> InvokeAsync(GetPluginsArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetPluginsResult>("cloudamqp:index/getPlugins:getPlugins", args ?? new GetPluginsArgs(), options.WithDefaults());

        /// <summary>
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using CloudAmqp = Pulumi.CloudAmqp;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var plugins = CloudAmqp.GetPlugins.Invoke(new()
        ///     {
        ///         InstanceId = cloudamqp_instance.Instance.Id,
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// ## Attributes reference
        /// 
        /// All attributes reference are computed
        /// 
        /// * `id`      - The identifier for this resource.
        /// * `plugins` - An array of plugins. Each `plugins` block consists of the fields documented below.
        /// 
        /// ___
        /// 
        /// The `plugins` block consist of
        /// 
        /// * `name`        - The type of the recipient.
        /// * `version`     - Rabbit MQ version that the plugins are shipped with.
        /// * `description` - Description of what the plugin does.
        /// * `enabled`     - Enable or disable information for the plugin.
        /// 
        /// ## Dependency
        /// 
        /// This data source depends on CloudAMQP instance identifier, `cloudamqp_instance.instance.id`.
        /// </summary>
        public static Output<GetPluginsResult> Invoke(GetPluginsInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetPluginsResult>("cloudamqp:index/getPlugins:getPlugins", args ?? new GetPluginsInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetPluginsArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The CloudAMQP instance identifier.
        /// </summary>
        [Input("instanceId", required: true)]
        public int InstanceId { get; set; }

        public GetPluginsArgs()
        {
        }
        public static new GetPluginsArgs Empty => new GetPluginsArgs();
    }

    public sealed class GetPluginsInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The CloudAMQP instance identifier.
        /// </summary>
        [Input("instanceId", required: true)]
        public Input<int> InstanceId { get; set; } = null!;

        public GetPluginsInvokeArgs()
        {
        }
        public static new GetPluginsInvokeArgs Empty => new GetPluginsInvokeArgs();
    }


    [OutputType]
    public sealed class GetPluginsResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly int InstanceId;
        public readonly ImmutableArray<Outputs.GetPluginsPluginResult> Plugins;

        [OutputConstructor]
        private GetPluginsResult(
            string id,

            int instanceId,

            ImmutableArray<Outputs.GetPluginsPluginResult> plugins)
        {
            Id = id;
            InstanceId = instanceId;
            Plugins = plugins;
        }
    }
}
