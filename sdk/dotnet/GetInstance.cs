// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.CloudAmqp
{
    public static class GetInstance
    {
        /// <summary>
        /// Use this data source to retrieve information about an already created CloudAMQP instance. In order to retrieve the correct information, the CoudAMQP instance identifier is needed.
        /// </summary>
        public static Task<GetInstanceResult> InvokeAsync(GetInstanceArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetInstanceResult>("cloudamqp:index/getInstance:getInstance", args ?? new GetInstanceArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to retrieve information about an already created CloudAMQP instance. In order to retrieve the correct information, the CoudAMQP instance identifier is needed.
        /// </summary>
        public static Output<GetInstanceResult> Invoke(GetInstanceInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetInstanceResult>("cloudamqp:index/getInstance:getInstance", args ?? new GetInstanceInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetInstanceArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The CloudAMQP instance identifier.
        /// </summary>
        [Input("instanceId", required: true)]
        public int InstanceId { get; set; }

        public GetInstanceArgs()
        {
        }
        public static new GetInstanceArgs Empty => new GetInstanceArgs();
    }

    public sealed class GetInstanceInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The CloudAMQP instance identifier.
        /// </summary>
        [Input("instanceId", required: true)]
        public Input<int> InstanceId { get; set; } = null!;

        public GetInstanceInvokeArgs()
        {
        }
        public static new GetInstanceInvokeArgs Empty => new GetInstanceInvokeArgs();
    }


    [OutputType]
    public sealed class GetInstanceResult
    {
        public readonly string Apikey;
        public readonly string Backend;
        public readonly bool Dedicated;
        public readonly string Host;
        public readonly string HostInternal;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly int InstanceId;
        public readonly string Name;
        public readonly bool NoDefaultAlarms;
        public readonly int Nodes;
        public readonly string Plan;
        public readonly bool Ready;
        public readonly string Region;
        public readonly string RmqVersion;
        public readonly ImmutableArray<string> Tags;
        public readonly string Url;
        public readonly string Vhost;
        public readonly int VpcId;
        public readonly string VpcSubnet;

        [OutputConstructor]
        private GetInstanceResult(
            string apikey,

            string backend,

            bool dedicated,

            string host,

            string hostInternal,

            string id,

            int instanceId,

            string name,

            bool noDefaultAlarms,

            int nodes,

            string plan,

            bool ready,

            string region,

            string rmqVersion,

            ImmutableArray<string> tags,

            string url,

            string vhost,

            int vpcId,

            string vpcSubnet)
        {
            Apikey = apikey;
            Backend = backend;
            Dedicated = dedicated;
            Host = host;
            HostInternal = hostInternal;
            Id = id;
            InstanceId = instanceId;
            Name = name;
            NoDefaultAlarms = noDefaultAlarms;
            Nodes = nodes;
            Plan = plan;
            Ready = ready;
            Region = region;
            RmqVersion = rmqVersion;
            Tags = tags;
            Url = url;
            Vhost = vhost;
            VpcId = vpcId;
            VpcSubnet = vpcSubnet;
        }
    }
}
