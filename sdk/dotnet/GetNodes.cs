// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.CloudAmqp
{
    public static class GetNodes
    {
        /// <summary>
        /// Use this data source to retrieve information about the node(s) created by CloudAMQP instance.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using CloudAmqp = Pulumi.CloudAmqp;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var nodes = Output.Create(CloudAmqp.GetNodes.InvokeAsync(new CloudAmqp.GetNodesArgs
        ///         {
        ///             InstanceId = cloudamqp_instance.Instance.Id,
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// ## Argument reference
        /// 
        /// * `instance_id` - (Required) The CloudAMQP instance identifier.
        /// 
        /// ## Attribute reference
        /// 
        /// * `nodes` - (Computed) An array of node information. Each `nodes` block consists of the fields documented below.
        /// 
        /// ___
        /// 
        /// The `nodes` block consist of
        /// 
        /// * `hostname`          - (Computed) Hostname assigned to the node.
        /// * `name`              - (Computed) Name of the node.
        /// * `running`           - (Computed) Is the node running?
        /// * `rabbitmq_version`  - (Computed) Currently configured Rabbit MQ version on the node.
        /// * `erlang_version`    - (Computed) Currently used Erlanbg version on the node.
        /// * `hipe`              - (Computed) Enable or disable High-performance Erlang.
        /// 
        /// ## Dependency
        /// 
        /// This data source depends on CloudAMQP instance identifier, `cloudamqp_instance.instance.id`.
        /// </summary>
        public static Task<GetNodesResult> InvokeAsync(GetNodesArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetNodesResult>("cloudamqp:index/getNodes:getNodes", args ?? new GetNodesArgs(), options.WithVersion());
    }


    public sealed class GetNodesArgs : Pulumi.InvokeArgs
    {
        [Input("instanceId", required: true)]
        public int InstanceId { get; set; }

        [Input("nodes")]
        private List<Inputs.GetNodesNodeArgs>? _nodes;
        public List<Inputs.GetNodesNodeArgs> Nodes
        {
            get => _nodes ?? (_nodes = new List<Inputs.GetNodesNodeArgs>());
            set => _nodes = value;
        }

        public GetNodesArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetNodesResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly int InstanceId;
        public readonly ImmutableArray<Outputs.GetNodesNodeResult> Nodes;

        [OutputConstructor]
        private GetNodesResult(
            string id,

            int instanceId,

            ImmutableArray<Outputs.GetNodesNodeResult> nodes)
        {
            Id = id;
            InstanceId = instanceId;
            Nodes = nodes;
        }
    }
}
