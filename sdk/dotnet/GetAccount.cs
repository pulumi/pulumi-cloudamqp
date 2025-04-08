// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.CloudAmqp
{
    public static class GetAccount
    {
        /// <summary>
        /// Use this data source to retrieve basic information about all instances available for an account.
        /// Uses the included apikey in provider configuration, to determine which account to read from.
        /// </summary>
        public static Task<GetAccountResult> InvokeAsync(InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetAccountResult>("cloudamqp:index/getAccount:getAccount", InvokeArgs.Empty, options.WithDefaults());

        /// <summary>
        /// Use this data source to retrieve basic information about all instances available for an account.
        /// Uses the included apikey in provider configuration, to determine which account to read from.
        /// </summary>
        public static Output<GetAccountResult> Invoke(InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetAccountResult>("cloudamqp:index/getAccount:getAccount", InvokeArgs.Empty, options.WithDefaults());

        /// <summary>
        /// Use this data source to retrieve basic information about all instances available for an account.
        /// Uses the included apikey in provider configuration, to determine which account to read from.
        /// </summary>
        public static Output<GetAccountResult> Invoke(InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetAccountResult>("cloudamqp:index/getAccount:getAccount", InvokeArgs.Empty, options.WithDefaults());
    }


    [OutputType]
    public sealed class GetAccountResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// An array of instances. Each `instances` block consists of the fields documented
        /// below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetAccountInstanceResult> Instances;

        [OutputConstructor]
        private GetAccountResult(
            string id,

            ImmutableArray<Outputs.GetAccountInstanceResult> instances)
        {
            Id = id;
            Instances = instances;
        }
    }
}
