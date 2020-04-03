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
        [Obsolete("Use GetVpcInfo.InvokeAsync() instead")]
        public static Task<GetVpcInfoResult> GetVpcInfo(GetVpcInfoArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetVpcInfoResult>("cloudamqp:index/getVpcInfo:getVpcInfo", args ?? InvokeArgs.Empty, options.WithVersion());
    }
    public static class GetVpcInfo
    {
        public static Task<GetVpcInfoResult> InvokeAsync(GetVpcInfoArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetVpcInfoResult>("cloudamqp:index/getVpcInfo:getVpcInfo", args ?? InvokeArgs.Empty, options.WithVersion());
    }

    public sealed class GetVpcInfoArgs : Pulumi.InvokeArgs
    {
        [Input("instanceId", required: true)]
        public int InstanceId { get; set; }

        public GetVpcInfoArgs()
        {
        }
    }

    [OutputType]
    public sealed class GetVpcInfoResult
    {
        public readonly int InstanceId;
        public readonly string Name;
        public readonly string OwnerId;
        public readonly string SecurityGroupId;
        public readonly string VpcSubnet;
        /// <summary>
        /// id is the provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;

        [OutputConstructor]
        private GetVpcInfoResult(
            int instanceId,
            string name,
            string ownerId,
            string securityGroupId,
            string vpcSubnet,
            string id)
        {
            InstanceId = instanceId;
            Name = name;
            OwnerId = ownerId;
            SecurityGroupId = securityGroupId;
            VpcSubnet = vpcSubnet;
            Id = id;
        }
    }
}
