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
    /// This resource allows you to create and manage alarms to trigger based on a set of conditions. Once triggerd a notification will be sent to the assigned recipients. When creating a new instance, there will also be a set of default alarms (cpu, memory and disk) created. All default alarms uses the default recipient for notifications.
    /// 
    /// By setting `no_default_alarms` to *true* in `cloudamqp.Instance`. This will create the instance without default alarms and avoid the need to import them to get full control.
    /// 
    /// Available for all subscription plans, but `lemur`and `tiger`are limited to fewer alarm types. The limited types supported can be seen in the table below in Alarm Type Reference.
    /// 
    /// ## Alarm Type reference
    /// 
    /// Valid options for notification type.
    /// 
    /// Required arguments for all alarms: *instance_id*, *type* and *enabled*&lt;br&gt;
    /// Optional argument for all alarms: *tags*, *queue_regex*, *vhost_regex*
    /// 
    /// | Name | Type | Shared | Dedicated | Required arguments |
    /// | ---- | ---- | ---- | ---- | ---- | ---- |
    /// | CPU | cpu | - | &amp;#10004; | time_threshold, value_threshold |
    /// | Memory | memory | - | &amp;#10004;  | time_threshold, value_threshold |
    /// | Disk space | disk | - | &amp;#10004;  | time_threshold, value_threshold |
    /// | Queue | queue | &amp;#10004;  | &amp;#10004;  | time_threshold, value_threshold, queue_regex, vhost_regex, message_type |
    /// | Connection | connection | &amp;#10004; | &amp;#10004; | time_threshold, value_threshold |
    /// | Consumer | consumer | &amp;#10004; | &amp;#10004; | time_threshold, value_threshold, queue, vhost |
    /// | Netsplit | netsplit | - | &amp;#10004; | time_threshold |
    /// | Server unreachable | server_unreachable  | - | &amp;#10004;  | time_threshold |
    /// | Notice | notice | &amp;#10004; | &amp;#10004; | |
    /// 
    /// ## Dependency
    /// 
    /// This resource depends on CloudAMQP instance identifier, `cloudamqp_instance.instance.id`.
    /// 
    /// ## Import
    /// 
    /// `cloudamqp_alarm` can be imported using CloudAMQP internal identifier of the alarm together (CSV separated) with the instance identifier. To retrieve the alarm identifier, use [CloudAMQP API](https://docs.cloudamqp.com/cloudamqp_api.html#list-alarms)
    /// 
    /// ```sh
    ///  $ pulumi import cloudamqp:index/alarm:Alarm alarm &lt;alarm_id&gt;,&lt;instance_id&gt;`
    /// ```
    /// </summary>
    [CloudAmqpResourceType("cloudamqp:index/alarm:Alarm")]
    public partial class Alarm : Pulumi.CustomResource
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

    public sealed class AlarmArgs : Pulumi.ResourceArgs
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
    }

    public sealed class AlarmState : Pulumi.ResourceArgs
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
    }
}
