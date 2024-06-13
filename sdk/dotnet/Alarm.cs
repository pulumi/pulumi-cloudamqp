// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.CloudAmqp
{
    /// <summary>
    /// ## Import
    /// 
    /// `cloudamqp_alarm` can be imported using CloudAMQP internal identifier of the alarm together (CSV separated) with the instance identifier. To retrieve the alarm identifier, use [CloudAMQP API](https://docs.cloudamqp.com/cloudamqp_api.html#list-alarms)
    /// 
    /// ```sh
    /// $ pulumi import cloudamqp:index/alarm:Alarm alarm &lt;id&gt;,&lt;instance_id&gt;`
    /// ```
    /// </summary>
    [CloudAmqpResourceType("cloudamqp:index/alarm:Alarm")]
    public partial class Alarm : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Enable or disable the alarm to trigger.
        /// </summary>
        [Output("enabled")]
        public Output<bool> Enabled { get; private set; } = null!;

        /// <summary>
        /// The CloudAMQP instance ID.
        /// </summary>
        [Output("instanceId")]
        public Output<int> InstanceId { get; private set; } = null!;

        /// <summary>
        /// Message type `(total, unacked, ready)` used by queue alarm type.
        /// 
        /// Specific argument for `disk` alarm
        /// </summary>
        [Output("messageType")]
        public Output<string?> MessageType { get; private set; } = null!;

        /// <summary>
        /// Regex for which queue to check.
        /// </summary>
        [Output("queueRegex")]
        public Output<string?> QueueRegex { get; private set; } = null!;

        /// <summary>
        /// Identifier for recipient to be notified. Leave empty to notify all recipients.
        /// </summary>
        [Output("recipients")]
        public Output<ImmutableArray<int>> Recipients { get; private set; } = null!;

        /// <summary>
        /// The reminder interval (in seconds) to resend the alarm if not resolved. Set to 0 for no reminders. The Default is 0.
        /// </summary>
        [Output("reminderInterval")]
        public Output<int?> ReminderInterval { get; private set; } = null!;

        /// <summary>
        /// The time interval (in seconds) the `value_threshold` should be active before triggering an alarm.
        /// </summary>
        [Output("timeThreshold")]
        public Output<int?> TimeThreshold { get; private set; } = null!;

        /// <summary>
        /// The alarm type, see valid options below.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;

        /// <summary>
        /// Disk value threshold calculation, `fixed, percentage` of disk space remaining.
        /// 
        /// Based on alarm type, different arguments are flagged as required or optional.
        /// </summary>
        [Output("valueCalculation")]
        public Output<string?> ValueCalculation { get; private set; } = null!;

        /// <summary>
        /// The value to trigger the alarm for.
        /// </summary>
        [Output("valueThreshold")]
        public Output<int?> ValueThreshold { get; private set; } = null!;

        /// <summary>
        /// Regex for which vhost to check
        /// </summary>
        [Output("vhostRegex")]
        public Output<string?> VhostRegex { get; private set; } = null!;


        /// <summary>
        /// Create a Alarm resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Alarm(string name, AlarmArgs args, CustomResourceOptions? options = null)
            : base("cloudamqp:index/alarm:Alarm", name, args ?? new AlarmArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Alarm(string name, Input<string> id, AlarmState? state = null, CustomResourceOptions? options = null)
            : base("cloudamqp:index/alarm:Alarm", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Alarm resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Alarm Get(string name, Input<string> id, AlarmState? state = null, CustomResourceOptions? options = null)
        {
            return new Alarm(name, id, state, options);
        }
    }

    public sealed class AlarmArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Enable or disable the alarm to trigger.
        /// </summary>
        [Input("enabled", required: true)]
        public Input<bool> Enabled { get; set; } = null!;

        /// <summary>
        /// The CloudAMQP instance ID.
        /// </summary>
        [Input("instanceId", required: true)]
        public Input<int> InstanceId { get; set; } = null!;

        /// <summary>
        /// Message type `(total, unacked, ready)` used by queue alarm type.
        /// 
        /// Specific argument for `disk` alarm
        /// </summary>
        [Input("messageType")]
        public Input<string>? MessageType { get; set; }

        /// <summary>
        /// Regex for which queue to check.
        /// </summary>
        [Input("queueRegex")]
        public Input<string>? QueueRegex { get; set; }

        [Input("recipients", required: true)]
        private InputList<int>? _recipients;

        /// <summary>
        /// Identifier for recipient to be notified. Leave empty to notify all recipients.
        /// </summary>
        public InputList<int> Recipients
        {
            get => _recipients ?? (_recipients = new InputList<int>());
            set => _recipients = value;
        }

        /// <summary>
        /// The reminder interval (in seconds) to resend the alarm if not resolved. Set to 0 for no reminders. The Default is 0.
        /// </summary>
        [Input("reminderInterval")]
        public Input<int>? ReminderInterval { get; set; }

        /// <summary>
        /// The time interval (in seconds) the `value_threshold` should be active before triggering an alarm.
        /// </summary>
        [Input("timeThreshold")]
        public Input<int>? TimeThreshold { get; set; }

        /// <summary>
        /// The alarm type, see valid options below.
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        /// <summary>
        /// Disk value threshold calculation, `fixed, percentage` of disk space remaining.
        /// 
        /// Based on alarm type, different arguments are flagged as required or optional.
        /// </summary>
        [Input("valueCalculation")]
        public Input<string>? ValueCalculation { get; set; }

        /// <summary>
        /// The value to trigger the alarm for.
        /// </summary>
        [Input("valueThreshold")]
        public Input<int>? ValueThreshold { get; set; }

        /// <summary>
        /// Regex for which vhost to check
        /// </summary>
        [Input("vhostRegex")]
        public Input<string>? VhostRegex { get; set; }

        public AlarmArgs()
        {
        }
        public static new AlarmArgs Empty => new AlarmArgs();
    }

    public sealed class AlarmState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Enable or disable the alarm to trigger.
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        /// <summary>
        /// The CloudAMQP instance ID.
        /// </summary>
        [Input("instanceId")]
        public Input<int>? InstanceId { get; set; }

        /// <summary>
        /// Message type `(total, unacked, ready)` used by queue alarm type.
        /// 
        /// Specific argument for `disk` alarm
        /// </summary>
        [Input("messageType")]
        public Input<string>? MessageType { get; set; }

        /// <summary>
        /// Regex for which queue to check.
        /// </summary>
        [Input("queueRegex")]
        public Input<string>? QueueRegex { get; set; }

        [Input("recipients")]
        private InputList<int>? _recipients;

        /// <summary>
        /// Identifier for recipient to be notified. Leave empty to notify all recipients.
        /// </summary>
        public InputList<int> Recipients
        {
            get => _recipients ?? (_recipients = new InputList<int>());
            set => _recipients = value;
        }

        /// <summary>
        /// The reminder interval (in seconds) to resend the alarm if not resolved. Set to 0 for no reminders. The Default is 0.
        /// </summary>
        [Input("reminderInterval")]
        public Input<int>? ReminderInterval { get; set; }

        /// <summary>
        /// The time interval (in seconds) the `value_threshold` should be active before triggering an alarm.
        /// </summary>
        [Input("timeThreshold")]
        public Input<int>? TimeThreshold { get; set; }

        /// <summary>
        /// The alarm type, see valid options below.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        /// <summary>
        /// Disk value threshold calculation, `fixed, percentage` of disk space remaining.
        /// 
        /// Based on alarm type, different arguments are flagged as required or optional.
        /// </summary>
        [Input("valueCalculation")]
        public Input<string>? ValueCalculation { get; set; }

        /// <summary>
        /// The value to trigger the alarm for.
        /// </summary>
        [Input("valueThreshold")]
        public Input<int>? ValueThreshold { get; set; }

        /// <summary>
        /// Regex for which vhost to check
        /// </summary>
        [Input("vhostRegex")]
        public Input<string>? VhostRegex { get; set; }

        public AlarmState()
        {
        }
        public static new AlarmState Empty => new AlarmState();
    }
}
