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
        /// 
        /// ## Argument reference
        /// 
        /// * `instance_id` - (Required) The CloudAMQP instance identifier.
        /// 
        /// ## Attribute reference
        /// 
        /// * `name`        - (Computed) The name of the CloudAMQP instance.
        /// * `plan`        - (Computed) The subscription plan for the CloudAMQP instance.
        /// * `region`      - (Computed) The cloud platform and region that host the CloudAMQP instance, `{platform}::{region}`.
        /// * `vpc_subnet`  - (Computed) Dedicated VPC subnet configured for the CloudAMQP instance.
        /// * `nodes`       - (Computed) Number of nodes in the cluster of the CloudAMQP instance.
        /// * `rmq_version` - (Computed) The version of installed Rabbit MQ.
        /// * `url`         - (Computed/Sensitive) The AMQP url, used by clients to connect for pub/sub.
        /// * `apikey`      - (Computed/Sensitive) The API key to secondary API handing alarms, integration etc.
        /// * `tags`        - (Computed) Tags the CloudAMQP instance with categories.
        /// * `host`        - (Computed) The hostname for the CloudAMQP instance.
        /// * `vhost`       - (Computed) The virtual host configured in Rabbit MQ.
        /// </summary>
        public static Task<GetInstanceResult> InvokeAsync(GetInstanceArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetInstanceResult>("cloudamqp:index/getInstance:getInstance", args ?? new GetInstanceArgs(), options.WithVersion());
    }


    public sealed class GetInstanceArgs : Pulumi.InvokeArgs
    {
        [Input("instanceId", required: true)]
        public int InstanceId { get; set; }

        [Input("vpcSubnet")]
        public string? VpcSubnet { get; set; }

        public GetInstanceArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetInstanceResult
    {
        public readonly string Apikey;
        public readonly bool Dedicated;
        public readonly string Host;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly int InstanceId;
        public readonly string Name;
        public readonly int Nodes;
        public readonly string Plan;
        public readonly string Region;
        public readonly string RmqVersion;
        public readonly ImmutableArray<string> Tags;
        public readonly string Url;
        public readonly string Vhost;
        public readonly string? VpcSubnet;

        [OutputConstructor]
        private GetInstanceResult(
            string apikey,

            bool dedicated,

            string host,

            string id,

            int instanceId,

            string name,

            int nodes,

            string plan,

            string region,

            string rmqVersion,

            ImmutableArray<string> tags,

            string url,

            string vhost,

            string? vpcSubnet)
        {
            Apikey = apikey;
            Dedicated = dedicated;
            Host = host;
            Id = id;
            InstanceId = instanceId;
            Name = name;
            Nodes = nodes;
            Plan = plan;
            Region = region;
            RmqVersion = rmqVersion;
            Tags = tags;
            Url = url;
            Vhost = vhost;
            VpcSubnet = vpcSubnet;
        }
    }
}
