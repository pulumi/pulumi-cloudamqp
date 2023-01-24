// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.CloudAmqp
{
    public static class GetNotification
    {
        /// <summary>
        /// Use this data source to retrieve information about default or created recipients. The recipient will receive notifications assigned to an alarm that has triggered. To retrieve the recipient either use `recipient_id` or `name`.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using Pulumi;
        /// using CloudAmqp = Pulumi.CloudAmqp;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var defaultRecipient = CloudAmqp.GetNotification.Invoke(new()
        ///     {
        ///         InstanceId = cloudamqp_instance.Instance.Id,
        ///         Name = "default",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// ## Attributes reference
        /// 
        /// All attributes reference are computed
        /// 
        /// * `id`    - The identifier for this resource.
        /// * `type`  - The type of the recipient.
        /// * `value` - The notification endpoint, where to send the notification.
        /// 
        /// ## Dependency
        /// 
        /// This data source depends on CloudAMQP instance identifier, `cloudamqp_instance.instance.id`.
        /// </summary>
        public static Task<GetNotificationResult> InvokeAsync(GetNotificationArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetNotificationResult>("cloudamqp:index/getNotification:getNotification", args ?? new GetNotificationArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to retrieve information about default or created recipients. The recipient will receive notifications assigned to an alarm that has triggered. To retrieve the recipient either use `recipient_id` or `name`.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using Pulumi;
        /// using CloudAmqp = Pulumi.CloudAmqp;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var defaultRecipient = CloudAmqp.GetNotification.Invoke(new()
        ///     {
        ///         InstanceId = cloudamqp_instance.Instance.Id,
        ///         Name = "default",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// ## Attributes reference
        /// 
        /// All attributes reference are computed
        /// 
        /// * `id`    - The identifier for this resource.
        /// * `type`  - The type of the recipient.
        /// * `value` - The notification endpoint, where to send the notification.
        /// 
        /// ## Dependency
        /// 
        /// This data source depends on CloudAMQP instance identifier, `cloudamqp_instance.instance.id`.
        /// </summary>
        public static Output<GetNotificationResult> Invoke(GetNotificationInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetNotificationResult>("cloudamqp:index/getNotification:getNotification", args ?? new GetNotificationInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetNotificationArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The CloudAMQP instance identifier.
        /// </summary>
        [Input("instanceId", required: true)]
        public int InstanceId { get; set; }

        /// <summary>
        /// The name set for the recipient.
        /// </summary>
        [Input("name")]
        public string? Name { get; set; }

        /// <summary>
        /// The recipient identifier.
        /// </summary>
        [Input("recipientId")]
        public int? RecipientId { get; set; }

        public GetNotificationArgs()
        {
        }
        public static new GetNotificationArgs Empty => new GetNotificationArgs();
    }

    public sealed class GetNotificationInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The CloudAMQP instance identifier.
        /// </summary>
        [Input("instanceId", required: true)]
        public Input<int> InstanceId { get; set; } = null!;

        /// <summary>
        /// The name set for the recipient.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The recipient identifier.
        /// </summary>
        [Input("recipientId")]
        public Input<int>? RecipientId { get; set; }

        public GetNotificationInvokeArgs()
        {
        }
        public static new GetNotificationInvokeArgs Empty => new GetNotificationInvokeArgs();
    }


    [OutputType]
    public sealed class GetNotificationResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly int InstanceId;
        public readonly string? Name;
        public readonly int? RecipientId;
        public readonly string Type;
        public readonly string Value;

        [OutputConstructor]
        private GetNotificationResult(
            string id,

            int instanceId,

            string? name,

            int? recipientId,

            string type,

            string value)
        {
            Id = id;
            InstanceId = instanceId;
            Name = name;
            RecipientId = recipientId;
            Type = type;
            Value = value;
        }
    }
}
