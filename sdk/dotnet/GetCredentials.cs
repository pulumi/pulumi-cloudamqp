// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.CloudAmqp
{
    public static class GetCredentials
    {
        /// <summary>
        /// Use this data source to retrieve information about the credentials of the configured user in Rabbit MQ. Information is extracted from `cloudamqp_instance.instance.url`.
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
        ///         var credentials = Output.Create(CloudAmqp.GetCredentials.InvokeAsync(new CloudAmqp.GetCredentialsArgs
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
        /// All attributes reference are computed.
        /// 
        /// * `id`          - The identifier for this data source.
        /// * `username`    - (Sensitive) The username for the configured user in Rabbit MQ.
        /// * `password`    - (Sensitive) The password used by the `username`.
        /// 
        /// ## Dependency
        /// 
        /// This data source depends on CloudAMQP instance identifier, `cloudamqp_instance.instance.id`.
        /// </summary>
        public static Task<GetCredentialsResult> InvokeAsync(GetCredentialsArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetCredentialsResult>("cloudamqp:index/getCredentials:getCredentials", args ?? new GetCredentialsArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to retrieve information about the credentials of the configured user in Rabbit MQ. Information is extracted from `cloudamqp_instance.instance.url`.
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
        ///         var credentials = Output.Create(CloudAmqp.GetCredentials.InvokeAsync(new CloudAmqp.GetCredentialsArgs
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
        /// All attributes reference are computed.
        /// 
        /// * `id`          - The identifier for this data source.
        /// * `username`    - (Sensitive) The username for the configured user in Rabbit MQ.
        /// * `password`    - (Sensitive) The password used by the `username`.
        /// 
        /// ## Dependency
        /// 
        /// This data source depends on CloudAMQP instance identifier, `cloudamqp_instance.instance.id`.
        /// </summary>
        public static Output<GetCredentialsResult> Invoke(GetCredentialsInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetCredentialsResult>("cloudamqp:index/getCredentials:getCredentials", args ?? new GetCredentialsInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetCredentialsArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The CloudAMQP instance identifier.
        /// </summary>
        [Input("instanceId", required: true)]
        public int InstanceId { get; set; }

        public GetCredentialsArgs()
        {
        }
    }

    public sealed class GetCredentialsInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The CloudAMQP instance identifier.
        /// </summary>
        [Input("instanceId", required: true)]
        public Input<int> InstanceId { get; set; } = null!;

        public GetCredentialsInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetCredentialsResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly int InstanceId;
        public readonly string Password;
        public readonly string Username;

        [OutputConstructor]
        private GetCredentialsResult(
            string id,

            int instanceId,

            string password,

            string username)
        {
            Id = id;
            InstanceId = instanceId;
            Password = password;
            Username = username;
        }
    }
}
