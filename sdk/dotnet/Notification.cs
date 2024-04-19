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
    /// This resource allows you to create and manage recipients to receive alarm notifications. There will
    /// always be a default recipient created upon instance creation. This recipient will use team email and
    /// receive notifications from default alarms.
    /// 
    /// Available for all subscription plans.
    /// 
    /// ## Example Usage
    /// 
    /// &lt;details&gt;
    ///   &lt;summary&gt;
    ///     &lt;b&gt;Email recipient&lt;/b&gt;
    ///   &lt;/summary&gt;
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using CloudAmqp = Pulumi.CloudAmqp;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var emailRecipient = new CloudAmqp.Notification("email_recipient", new()
    ///     {
    ///         InstanceId = instance.Id,
    ///         Type = "email",
    ///         Value = "alarm@example.com",
    ///         Name = "alarm",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// &lt;/details&gt;
    /// 
    /// &lt;details&gt;
    ///   &lt;summary&gt;
    ///     &lt;b&gt;OpsGenie recipient with optional responders&lt;/b&gt;
    ///   &lt;/summary&gt;
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using CloudAmqp = Pulumi.CloudAmqp;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var opsgenieRecipient = new CloudAmqp.Notification("opsgenie_recipient", new()
    ///     {
    ///         InstanceId = instance.Id,
    ///         Type = "opsgenie",
    ///         Value = "&lt;api-key&gt;",
    ///         Name = "OpsGenie",
    ///         Responders = new[]
    ///         {
    ///             new CloudAmqp.Inputs.NotificationResponderArgs
    ///             {
    ///                 Type = "team",
    ///                 Id = "&lt;team-uuid&gt;",
    ///             },
    ///             new CloudAmqp.Inputs.NotificationResponderArgs
    ///             {
    ///                 Type = "user",
    ///                 Username = "&lt;username&gt;",
    ///             },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// &lt;/details&gt;
    /// 
    /// &lt;details&gt;
    ///   &lt;summary&gt;
    ///     &lt;b&gt;Pagerduty recipient with optional dedup key&lt;/b&gt;
    ///   &lt;/summary&gt;
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using CloudAmqp = Pulumi.CloudAmqp;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var pagerdutyRecipient = new CloudAmqp.Notification("pagerduty_recipient", new()
    ///     {
    ///         InstanceId = instance.Id,
    ///         Type = "pagerduty",
    ///         Value = "&lt;integration-key&gt;",
    ///         Name = "PagerDuty",
    ///         Options = 
    ///         {
    ///             { "dedupkey", "DEDUPKEY" },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// &lt;/details&gt;
    /// 
    /// &lt;details&gt;
    ///   &lt;summary&gt;
    ///     &lt;b&gt;Signl4 recipient&lt;/b&gt;
    ///   &lt;/summary&gt;
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using CloudAmqp = Pulumi.CloudAmqp;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var signl4Recipient = new CloudAmqp.Notification("signl4_recipient", new()
    ///     {
    ///         InstanceId = instance.Id,
    ///         Type = "signl4",
    ///         Value = "&lt;team-secret&gt;",
    ///         Name = "Signl4",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// &lt;/details&gt;
    /// 
    /// &lt;details&gt;
    ///   &lt;summary&gt;
    ///     &lt;b&gt;Teams recipient&lt;/b&gt;
    ///   &lt;/summary&gt;
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using CloudAmqp = Pulumi.CloudAmqp;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var teamsRecipient = new CloudAmqp.Notification("teams_recipient", new()
    ///     {
    ///         InstanceId = instance.Id,
    ///         Type = "teams",
    ///         Value = "&lt;teams-webhook-url&gt;",
    ///         Name = "Teams",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// &lt;/details&gt;
    /// 
    /// &lt;details&gt;
    ///   &lt;summary&gt;
    ///     &lt;b&gt;Victorops recipient with optional routing key (rk)&lt;/b&gt;
    ///   &lt;/summary&gt;
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using CloudAmqp = Pulumi.CloudAmqp;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var victoropsRecipient = new CloudAmqp.Notification("victorops_recipient", new()
    ///     {
    ///         InstanceId = instance.Id,
    ///         Type = "victorops",
    ///         Value = "&lt;integration-key&gt;",
    ///         Name = "Victorops",
    ///         Options = 
    ///         {
    ///             { "rk", "ROUTINGKEY" },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// &lt;/details&gt;
    /// 
    /// &lt;details&gt;
    ///   &lt;summary&gt;
    ///     &lt;b&gt;Webhook recipient&lt;/b&gt;
    ///   &lt;/summary&gt;
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using CloudAmqp = Pulumi.CloudAmqp;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var webhookRecipient = new CloudAmqp.Notification("webhook_recipient", new()
    ///     {
    ///         InstanceId = instance.Id,
    ///         Type = "webhook",
    ///         Value = "&lt;webhook-url&gt;",
    ///         Name = "Webhook",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// &lt;/details&gt;
    /// 
    /// ## Notification Type reference
    /// 
    /// Valid options for notification type.
    /// 
    /// * email
    /// * opsgenie
    /// * opsgenie-eu
    /// * pagerduty
    /// * signl4
    /// * slack
    /// * teams
    /// * victorops
    /// * webhook
    /// 
    /// ## Options parameter
    /// 
    /// | Type      | Options  | Description | Note |
    /// |---|---|---|---|
    /// | Victorops | rk       | Routing key to route alarm notification | - |
    /// | PagerDuty | dedupkey | Default the dedup key for PagerDuty is generated depending on what alarm has triggered, but here you can set what `dedup` key to use so even if the same alarm is triggered for different resources you only get one notification. Leave blank to use the generated dedup key. | If multiple alarms are triggered using this recipient, since they all share `dedup` key only the first alarm will be shown in PagerDuty |
    /// 
    /// ## Dependency
    /// 
    /// This resource depends on CloudAMQP instance identifier, `cloudamqp_instance.instance.id`.
    /// 
    /// ## Import
    /// 
    /// `cloudamqp_notification` can be imported using CloudAMQP internal identifier of a recipient together
    /// 
    /// (CSV separated) with the instance identifier. To retrieve the identifier of a recipient, use
    /// 
    /// [CloudAMQP API](https://docs.cloudamqp.com/cloudamqp_api.html#list-notification-recipients)
    /// 
    /// ```sh
    /// $ pulumi import cloudamqp:index/notification:Notification recipient &lt;id&gt;,&lt;instance_id&gt;`
    /// ```
    /// </summary>
    [CloudAmqpResourceType("cloudamqp:index/notification:Notification")]
    public partial class Notification : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The CloudAMQP instance ID.
        /// </summary>
        [Output("instanceId")]
        public Output<int> InstanceId { get; private set; } = null!;

        /// <summary>
        /// Name of the responder
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Options argument (e.g. `rk` used for VictorOps routing key).
        /// </summary>
        [Output("options")]
        public Output<ImmutableDictionary<string, string>?> Options { get; private set; } = null!;

        /// <summary>
        /// An array of reponders (only for OpsGenie). Each `responders` block
        /// consists of the field documented below.
        /// 
        /// ___
        /// 
        /// The `responders` block consists of:
        /// </summary>
        [Output("responders")]
        public Output<ImmutableArray<Outputs.NotificationResponder>> Responders { get; private set; } = null!;

        /// <summary>
        /// Type of responder. [`team`, `user`, `escalation`, `schedule`]
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;

        /// <summary>
        /// Integration/API key or endpoint to send the notification.
        /// </summary>
        [Output("value")]
        public Output<string> Value { get; private set; } = null!;


        /// <summary>
        /// Create a Notification resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Notification(string name, NotificationArgs args, CustomResourceOptions? options = null)
            : base("cloudamqp:index/notification:Notification", name, args ?? new NotificationArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Notification(string name, Input<string> id, NotificationState? state = null, CustomResourceOptions? options = null)
            : base("cloudamqp:index/notification:Notification", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Notification resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Notification Get(string name, Input<string> id, NotificationState? state = null, CustomResourceOptions? options = null)
        {
            return new Notification(name, id, state, options);
        }
    }

    public sealed class NotificationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The CloudAMQP instance ID.
        /// </summary>
        [Input("instanceId", required: true)]
        public Input<int> InstanceId { get; set; } = null!;

        /// <summary>
        /// Name of the responder
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("options")]
        private InputMap<string>? _options;

        /// <summary>
        /// Options argument (e.g. `rk` used for VictorOps routing key).
        /// </summary>
        public InputMap<string> Options
        {
            get => _options ?? (_options = new InputMap<string>());
            set => _options = value;
        }

        [Input("responders")]
        private InputList<Inputs.NotificationResponderArgs>? _responders;

        /// <summary>
        /// An array of reponders (only for OpsGenie). Each `responders` block
        /// consists of the field documented below.
        /// 
        /// ___
        /// 
        /// The `responders` block consists of:
        /// </summary>
        public InputList<Inputs.NotificationResponderArgs> Responders
        {
            get => _responders ?? (_responders = new InputList<Inputs.NotificationResponderArgs>());
            set => _responders = value;
        }

        /// <summary>
        /// Type of responder. [`team`, `user`, `escalation`, `schedule`]
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        /// <summary>
        /// Integration/API key or endpoint to send the notification.
        /// </summary>
        [Input("value", required: true)]
        public Input<string> Value { get; set; } = null!;

        public NotificationArgs()
        {
        }
        public static new NotificationArgs Empty => new NotificationArgs();
    }

    public sealed class NotificationState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The CloudAMQP instance ID.
        /// </summary>
        [Input("instanceId")]
        public Input<int>? InstanceId { get; set; }

        /// <summary>
        /// Name of the responder
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("options")]
        private InputMap<string>? _options;

        /// <summary>
        /// Options argument (e.g. `rk` used for VictorOps routing key).
        /// </summary>
        public InputMap<string> Options
        {
            get => _options ?? (_options = new InputMap<string>());
            set => _options = value;
        }

        [Input("responders")]
        private InputList<Inputs.NotificationResponderGetArgs>? _responders;

        /// <summary>
        /// An array of reponders (only for OpsGenie). Each `responders` block
        /// consists of the field documented below.
        /// 
        /// ___
        /// 
        /// The `responders` block consists of:
        /// </summary>
        public InputList<Inputs.NotificationResponderGetArgs> Responders
        {
            get => _responders ?? (_responders = new InputList<Inputs.NotificationResponderGetArgs>());
            set => _responders = value;
        }

        /// <summary>
        /// Type of responder. [`team`, `user`, `escalation`, `schedule`]
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        /// <summary>
        /// Integration/API key or endpoint to send the notification.
        /// </summary>
        [Input("value")]
        public Input<string>? Value { get; set; }

        public NotificationState()
        {
        }
        public static new NotificationState Empty => new NotificationState();
    }
}
