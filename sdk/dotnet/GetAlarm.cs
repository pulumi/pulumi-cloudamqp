// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.CloudAmqp
{
    public static class GetAlarm
    {
        /// <summary>
        /// Use this data source to retrieve information about default or created alarms. Either use `alarm_id`
        /// or `type` to retrieve the alarm.
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
        ///     var defaultCpuAlarm = CloudAmqp.GetAlarm.Invoke(new()
        ///     {
        ///         InstanceId = instance.Id,
        ///         Type = "cpu",
        ///     });
        /// 
        /// });
        /// ```
        /// 
        /// ## Dependency
        /// 
        /// This data source depends on CloudAMQP instance identifier, `cloudamqp_instance.instance.id`.
        /// 
        /// ## Alarm Types
        /// 
        /// `cpu, memory, disk, queue, connection, flow, consumer, netsplit, server_unreachable, notice`
        /// </summary>
        public static Task<GetAlarmResult> InvokeAsync(GetAlarmArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetAlarmResult>("cloudamqp:index/getAlarm:getAlarm", args ?? new GetAlarmArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to retrieve information about default or created alarms. Either use `alarm_id`
        /// or `type` to retrieve the alarm.
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
        ///     var defaultCpuAlarm = CloudAmqp.GetAlarm.Invoke(new()
        ///     {
        ///         InstanceId = instance.Id,
        ///         Type = "cpu",
        ///     });
        /// 
        /// });
        /// ```
        /// 
        /// ## Dependency
        /// 
        /// This data source depends on CloudAMQP instance identifier, `cloudamqp_instance.instance.id`.
        /// 
        /// ## Alarm Types
        /// 
        /// `cpu, memory, disk, queue, connection, flow, consumer, netsplit, server_unreachable, notice`
        /// </summary>
        public static Output<GetAlarmResult> Invoke(GetAlarmInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetAlarmResult>("cloudamqp:index/getAlarm:getAlarm", args ?? new GetAlarmInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to retrieve information about default or created alarms. Either use `alarm_id`
        /// or `type` to retrieve the alarm.
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
        ///     var defaultCpuAlarm = CloudAmqp.GetAlarm.Invoke(new()
        ///     {
        ///         InstanceId = instance.Id,
        ///         Type = "cpu",
        ///     });
        /// 
        /// });
        /// ```
        /// 
        /// ## Dependency
        /// 
        /// This data source depends on CloudAMQP instance identifier, `cloudamqp_instance.instance.id`.
        /// 
        /// ## Alarm Types
        /// 
        /// `cpu, memory, disk, queue, connection, flow, consumer, netsplit, server_unreachable, notice`
        /// </summary>
        public static Output<GetAlarmResult> Invoke(GetAlarmInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetAlarmResult>("cloudamqp:index/getAlarm:getAlarm", args ?? new GetAlarmInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetAlarmArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The alarm identifier. Either use this or `type` to give
        /// `cloudamqp.Alarm` necessary information to retrieve the alarm.
        /// </summary>
        [Input("alarmId")]
        public int? AlarmId { get; set; }

        /// <summary>
        /// The CloudAMQP instance identifier.
        /// </summary>
        [Input("instanceId", required: true)]
        public int InstanceId { get; set; }

        /// <summary>
        /// The alarm type. Either use this or `alarm_id` to give `cloudamqp.Alarm`
        /// necessary information when retrieve the alarm. Supported
        /// alarm types.
        /// </summary>
        [Input("type")]
        public string? Type { get; set; }

        /// <summary>
        /// Disk value threshold calculation, `(fixed, percentage)` of disk space
        /// remaining.
        /// </summary>
        [Input("valueCalculation")]
        public string? ValueCalculation { get; set; }

        public GetAlarmArgs()
        {
        }
        public static new GetAlarmArgs Empty => new GetAlarmArgs();
    }

    public sealed class GetAlarmInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The alarm identifier. Either use this or `type` to give
        /// `cloudamqp.Alarm` necessary information to retrieve the alarm.
        /// </summary>
        [Input("alarmId")]
        public Input<int>? AlarmId { get; set; }

        /// <summary>
        /// The CloudAMQP instance identifier.
        /// </summary>
        [Input("instanceId", required: true)]
        public Input<int> InstanceId { get; set; } = null!;

        /// <summary>
        /// The alarm type. Either use this or `alarm_id` to give `cloudamqp.Alarm`
        /// necessary information when retrieve the alarm. Supported
        /// alarm types.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        /// <summary>
        /// Disk value threshold calculation, `(fixed, percentage)` of disk space
        /// remaining.
        /// </summary>
        [Input("valueCalculation")]
        public Input<string>? ValueCalculation { get; set; }

        public GetAlarmInvokeArgs()
        {
        }
        public static new GetAlarmInvokeArgs Empty => new GetAlarmInvokeArgs();
    }


    [OutputType]
    public sealed class GetAlarmResult
    {
        public readonly int? AlarmId;
        /// <summary>
        /// Enable/disable status of the alarm.
        /// </summary>
        public readonly bool Enabled;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly int InstanceId;
        /// <summary>
        /// Message type `(total, unacked, ready)` used by queue alarm type.
        /// </summary>
        public readonly string MessageType;
        /// <summary>
        /// Regular expression for which queue to check.
        /// </summary>
        public readonly string QueueRegex;
        /// <summary>
        /// Identifier for recipient to be notified.
        /// </summary>
        public readonly ImmutableArray<int> Recipients;
        /// <summary>
        /// The reminder interval (in seconds) to resend the alarm if not resolved.
        /// Set to 0 for no reminders.
        /// </summary>
        public readonly int ReminderInterval;
        /// <summary>
        /// The time interval (in seconds) the `value_threshold` should be active
        /// before trigger an alarm.
        /// </summary>
        public readonly int TimeThreshold;
        public readonly string? Type;
        /// <summary>
        /// Disk value threshold calculation, `(fixed, percentage)` of disk space
        /// remaining.
        /// </summary>
        public readonly string? ValueCalculation;
        /// <summary>
        /// The value threshold that triggers the alarm.
        /// </summary>
        public readonly int ValueThreshold;
        /// <summary>
        /// Regular expression for which vhost to check
        /// </summary>
        public readonly string VhostRegex;

        [OutputConstructor]
        private GetAlarmResult(
            int? alarmId,

            bool enabled,

            string id,

            int instanceId,

            string messageType,

            string queueRegex,

            ImmutableArray<int> recipients,

            int reminderInterval,

            int timeThreshold,

            string? type,

            string? valueCalculation,

            int valueThreshold,

            string vhostRegex)
        {
            AlarmId = alarmId;
            Enabled = enabled;
            Id = id;
            InstanceId = instanceId;
            MessageType = messageType;
            QueueRegex = queueRegex;
            Recipients = recipients;
            ReminderInterval = reminderInterval;
            TimeThreshold = timeThreshold;
            Type = type;
            ValueCalculation = valueCalculation;
            ValueThreshold = valueThreshold;
            VhostRegex = vhostRegex;
        }
    }
}
