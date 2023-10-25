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
        /// * `hostname`              - External hostname assigned to the node.
        /// * `name`                  - Name of the node.
        /// * `running`               - Is the node running?
        /// * `rabbitmq_version`      - Currently configured Rabbit MQ version on the node.
        /// * `erlang_version`        - Currently used Erlang version on the node.
        /// * `hipe`                  - Enable or disable High-performance Erlang.
        /// * `configured`            - Is the node configured?
        /// * `disk_size`             - Subscription plan disk size
        /// * `additional_disk_size`  - Additional added disk size
        /// 
        /// ***Note:*** *Total disk size = disk_size + additional_disk_size*
        /// 
        /// ## Dependency
        /// 
        /// This data source depends on CloudAMQP instance identifier, `cloudamqp_instance.instance.id`.
        /// </summary>
        public static Task<GetNodesResult> InvokeAsync(GetNodesArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetNodesResult>("cloudamqp:index/getNodes:getNodes", args ?? new GetNodesArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to retrieve information about the node(s) created by CloudAMQP instance.
        /// 
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
        /// * `hostname`              - External hostname assigned to the node.
        /// * `name`                  - Name of the node.
        /// * `running`               - Is the node running?
        /// * `rabbitmq_version`      - Currently configured Rabbit MQ version on the node.
        /// * `erlang_version`        - Currently used Erlang version on the node.
        /// * `hipe`                  - Enable or disable High-performance Erlang.
        /// * `configured`            - Is the node configured?
        /// * `disk_size`             - Subscription plan disk size
        /// * `additional_disk_size`  - Additional added disk size
        /// 
        /// ***Note:*** *Total disk size = disk_size + additional_disk_size*
        /// 
        /// ## Dependency
        /// 
        /// This data source depends on CloudAMQP instance identifier, `cloudamqp_instance.instance.id`.
        /// </summary>
        public static Output<GetNodesResult> Invoke(GetNodesInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetNodesResult>("cloudamqp:index/getNodes:getNodes", args ?? new GetNodesInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetNodesArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The CloudAMQP instance identifier.
        /// </summary>
        [Input("instanceId", required: true)]
        public int InstanceId { get; set; }

        public GetNodesArgs()
        {
        }
        public static new GetNodesArgs Empty => new GetNodesArgs();
    }

    public sealed class GetNodesInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The CloudAMQP instance identifier.
        /// </summary>
        [Input("instanceId", required: true)]
        public Input<int> InstanceId { get; set; } = null!;

        public GetNodesInvokeArgs()
        {
        }
        public static new GetNodesInvokeArgs Empty => new GetNodesInvokeArgs();
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
