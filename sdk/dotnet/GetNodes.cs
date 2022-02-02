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
        /// ## Attributes reference
        /// 
        /// All attributes reference are computed
        /// 
        /// * `id`    - The identifier for this resource.
        /// * `nodes` - An array of node information. Each `nodes` block consists of the fields documented below.
        /// 
        /// ___
        /// 
        /// The `nodes` block consist of
        /// 
        /// * `hostname`          - External hostname assigned to the node.
        /// * `name`              - Name of the node.
        /// * `running`           - Is the node running?
        /// * `rabbitmq_version`  - Currently configured Rabbit MQ version on the node.
        /// * `erlang_version`    - Currently used Erlanbg version on the node.
        /// * `hipe`              - Enable or disable High-performance Erlang.
        /// * `configured`        - Is the node configured?
        /// 
        /// ## Dependency
        /// 
        /// This data source depends on CloudAMQP instance identifier, `cloudamqp_instance.instance.id`.
        /// </summary>
        public static Task<GetNodesResult> InvokeAsync(GetNodesArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetNodesResult>("cloudamqp:index/getNodes:getNodes", args ?? new GetNodesArgs(), options.WithDefaults());

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
        /// ## Attributes reference
        /// 
        /// All attributes reference are computed
        /// 
        /// * `id`    - The identifier for this resource.
        /// * `nodes` - An array of node information. Each `nodes` block consists of the fields documented below.
        /// 
        /// ___
        /// 
        /// The `nodes` block consist of
        /// 
        /// * `hostname`          - External hostname assigned to the node.
        /// * `name`              - Name of the node.
        /// * `running`           - Is the node running?
        /// * `rabbitmq_version`  - Currently configured Rabbit MQ version on the node.
        /// * `erlang_version`    - Currently used Erlanbg version on the node.
        /// * `hipe`              - Enable or disable High-performance Erlang.
        /// * `configured`        - Is the node configured?
        /// 
        /// ## Dependency
        /// 
        /// This data source depends on CloudAMQP instance identifier, `cloudamqp_instance.instance.id`.
        /// </summary>
        public static Output<GetNodesResult> Invoke(GetNodesInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetNodesResult>("cloudamqp:index/getNodes:getNodes", args ?? new GetNodesInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetNodesArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The CloudAMQP instance identifier.
        /// </summary>
        [Input("instanceId", required: true)]
        public int InstanceId { get; set; }

        public GetNodesArgs()
        {
        }
    }

    public sealed class GetNodesInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The CloudAMQP instance identifier.
        /// </summary>
        [Input("instanceId", required: true)]
        public Input<int> InstanceId { get; set; } = null!;

        public GetNodesInvokeArgs()
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
