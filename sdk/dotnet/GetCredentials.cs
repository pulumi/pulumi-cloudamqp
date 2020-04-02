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
        [Obsolete("Use GetCredentials.InvokeAsync() instead")]
        public static Task<GetCredentialsResult> GetCredentials(GetCredentialsArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetCredentialsResult>("cloudamqp:index/getCredentials:getCredentials", args ?? InvokeArgs.Empty, options.WithVersion());
    }
    public static class GetCredentials
    {
        public static Task<GetCredentialsResult> InvokeAsync(GetCredentialsArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetCredentialsResult>("cloudamqp:index/getCredentials:getCredentials", args ?? InvokeArgs.Empty, options.WithVersion());
    }

    public sealed class GetCredentialsArgs : Pulumi.InvokeArgs
    {
        [Input("instanceId", required: true)]
        public int InstanceId { get; set; }

        [Input("password")]
        public string? Password { get; set; }

        [Input("username")]
        public string? Username { get; set; }

        public GetCredentialsArgs()
        {
        }
    }

    [OutputType]
    public sealed class GetCredentialsResult
    {
        public readonly int InstanceId;
        public readonly string? Password;
        public readonly string? Username;
        /// <summary>
        /// id is the provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;

        [OutputConstructor]
        private GetCredentialsResult(
            int instanceId,
            string? password,
            string? username,
            string id)
        {
            InstanceId = instanceId;
            Password = password;
            Username = username;
            Id = id;
        }
    }
}
