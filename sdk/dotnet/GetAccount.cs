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
        /// Use this data source to retrieve basic information about all instances available for an account. Uses the included apikey in provider configuration, to determine which account to read from.
        /// 
        /// ## Attributes reference
        /// 
        /// All attributes reference are computed
        /// 
        /// * `id`      - The identifier for this resource. Set to `na` since there is no unique identifier.
        /// * `name`    - The name of the instance.
        /// * `plan`    - The subscription plan used for the instance.
        /// * `region`  - The region were the instanece is located in.
        /// * `tags`    - The tags set for the instance.
        /// 
        /// ## Dependency
        /// 
        /// This data source depends on apikey set in the provider configuration.
        /// </summary>
        public static Task<GetAccountResult> InvokeAsync(InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetAccountResult>("cloudamqp:index/getAccount:getAccount", InvokeArgs.Empty, options.WithDefaults());
    }


    [OutputType]
    public sealed class GetAccountResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
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
