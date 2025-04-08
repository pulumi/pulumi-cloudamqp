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
        /// Use this data source to retrieve information about the credentials of the configured user in
        /// RabbitMQ. Information is extracted from `cloudamqp_instance.instance.url`.
        /// 
        /// ## Example Usage
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using CloudAmqp = Pulumi.CloudAmqp;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var credentials = CloudAmqp.GetCredentials.Invoke(new()
        ///     {
        ///         InstanceId = instance.Id,
        ///     });
        /// 
        /// });
        /// ```
        /// 
        /// ## Dependency
        /// 
        /// This data source depends on CloudAMQP instance identifier, `cloudamqp_instance.instance.id`.
        /// </summary>
        public static Task<GetCredentialsResult> InvokeAsync(GetCredentialsArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetCredentialsResult>("cloudamqp:index/getCredentials:getCredentials", args ?? new GetCredentialsArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to retrieve information about the credentials of the configured user in
        /// RabbitMQ. Information is extracted from `cloudamqp_instance.instance.url`.
        /// 
        /// ## Example Usage
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using CloudAmqp = Pulumi.CloudAmqp;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var credentials = CloudAmqp.GetCredentials.Invoke(new()
        ///     {
        ///         InstanceId = instance.Id,
        ///     });
        /// 
        /// });
        /// ```
        /// 
        /// ## Dependency
        /// 
        /// This data source depends on CloudAMQP instance identifier, `cloudamqp_instance.instance.id`.
        /// </summary>
        public static Output<GetCredentialsResult> Invoke(GetCredentialsInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetCredentialsResult>("cloudamqp:index/getCredentials:getCredentials", args ?? new GetCredentialsInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to retrieve information about the credentials of the configured user in
        /// RabbitMQ. Information is extracted from `cloudamqp_instance.instance.url`.
        /// 
        /// ## Example Usage
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using CloudAmqp = Pulumi.CloudAmqp;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var credentials = CloudAmqp.GetCredentials.Invoke(new()
        ///     {
        ///         InstanceId = instance.Id,
        ///     });
        /// 
        /// });
        /// ```
        /// 
        /// ## Dependency
        /// 
        /// This data source depends on CloudAMQP instance identifier, `cloudamqp_instance.instance.id`.
        /// </summary>
        public static Output<GetCredentialsResult> Invoke(GetCredentialsInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetCredentialsResult>("cloudamqp:index/getCredentials:getCredentials", args ?? new GetCredentialsInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetCredentialsArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The CloudAMQP instance identifier.
        /// </summary>
        [Input("instanceId", required: true)]
        public int InstanceId { get; set; }

        public GetCredentialsArgs()
        {
        }
        public static new GetCredentialsArgs Empty => new GetCredentialsArgs();
    }

    public sealed class GetCredentialsInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The CloudAMQP instance identifier.
        /// </summary>
        [Input("instanceId", required: true)]
        public Input<int> InstanceId { get; set; } = null!;

        public GetCredentialsInvokeArgs()
        {
        }
        public static new GetCredentialsInvokeArgs Empty => new GetCredentialsInvokeArgs();
    }


    [OutputType]
    public sealed class GetCredentialsResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly int InstanceId;
        /// <summary>
        /// (Sensitive) The password used by the `username`.
        /// </summary>
        public readonly string Password;
        /// <summary>
        /// (Sensitive) The username for the configured user in Rabbit MQ.
        /// </summary>
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
