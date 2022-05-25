// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.CloudAmqp
{
    public static class GetUpgradableVersions
    {
        /// <summary>
        /// Use this data source to retrieve information about possible upgradable versions for RabbitMQ and Erlang.
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
        ///         var versions = Output.Create(CloudAmqp.GetUpgradableVersions.InvokeAsync(new CloudAmqp.GetUpgradableVersionsArgs
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
        /// * `new_rabbitmq_version`  - Possible upgradable version for RabbitMQ.
        /// * `new_erlang_version`    - Possible upgradable version for Erlang.
        /// 
        /// ## Dependency
        /// 
        /// This data source depends on CloudAMQP instance identifier, `cloudamqp_instance.instance.id`.
        /// </summary>
        public static Task<GetUpgradableVersionsResult> InvokeAsync(GetUpgradableVersionsArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetUpgradableVersionsResult>("cloudamqp:index/getUpgradableVersions:getUpgradableVersions", args ?? new GetUpgradableVersionsArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to retrieve information about possible upgradable versions for RabbitMQ and Erlang.
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
        ///         var versions = Output.Create(CloudAmqp.GetUpgradableVersions.InvokeAsync(new CloudAmqp.GetUpgradableVersionsArgs
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
        /// * `new_rabbitmq_version`  - Possible upgradable version for RabbitMQ.
        /// * `new_erlang_version`    - Possible upgradable version for Erlang.
        /// 
        /// ## Dependency
        /// 
        /// This data source depends on CloudAMQP instance identifier, `cloudamqp_instance.instance.id`.
        /// </summary>
        public static Output<GetUpgradableVersionsResult> Invoke(GetUpgradableVersionsInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetUpgradableVersionsResult>("cloudamqp:index/getUpgradableVersions:getUpgradableVersions", args ?? new GetUpgradableVersionsInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetUpgradableVersionsArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The CloudAMQP instance identifier.
        /// </summary>
        [Input("instanceId", required: true)]
        public int InstanceId { get; set; }

        public GetUpgradableVersionsArgs()
        {
        }
    }

    public sealed class GetUpgradableVersionsInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The CloudAMQP instance identifier.
        /// </summary>
        [Input("instanceId", required: true)]
        public Input<int> InstanceId { get; set; } = null!;

        public GetUpgradableVersionsInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetUpgradableVersionsResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly int InstanceId;
        public readonly string NewErlangVersion;
        public readonly string NewRabbitmqVersion;

        [OutputConstructor]
        private GetUpgradableVersionsResult(
            string id,

            int instanceId,

            string newErlangVersion,

            string newRabbitmqVersion)
        {
            Id = id;
            InstanceId = instanceId;
            NewErlangVersion = newErlangVersion;
            NewRabbitmqVersion = newRabbitmqVersion;
        }
    }
}
