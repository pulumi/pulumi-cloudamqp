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
    /// This resource allows you to create and manage, an [AWS EventBridge](https://aws.amazon.com/eventbridge/) for a CloudAMQP instance. Once created, continue to map the EventBridge in the [AWS Eventbridge console](https://console.aws.amazon.com/events/home).
    /// 
    /// &gt;  Our consumer needs to have exclusive usage to the configured queue and the maximum body size allowed on msgs by AWS is 256kb. The message body has to be valid JSON for AWS Eventbridge to accept it. If messages are too large or are not valid JSON, they will be rejected (tip: setup a dead-letter queue to catch them).
    /// 
    /// Not possible to update this resource. Any changes made to the argument will destroy and recreate the resource. Hence why all arguments use ForceNew.
    /// 
    /// Only available for dedicated subscription plans.
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
    ///     var instance = new CloudAmqp.Instance("instance", new()
    ///     {
    ///         Name = "Test instance",
    ///         Plan = "squirrel-1",
    ///         Region = "amazon-web-services::us-west-1",
    ///         RmqVersion = "3.11.5",
    ///         Tags = new[]
    ///         {
    ///             "aws",
    ///         },
    ///     });
    /// 
    ///     var awsEventbridge = new CloudAmqp.IntegrationAwsEventbridge("aws_eventbridge", new()
    ///     {
    ///         InstanceId = instance.Id,
    ///         Vhost = instance.Vhost,
    ///         Queue = "&lt;QUEUE-NAME&gt;",
    ///         AwsAccountId = "&lt;AWS-ACCOUNT-ID&gt;",
    ///         AwsRegion = "us-west-1",
    ///         WithHeaders = true,
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Argument references
    /// 
    /// The following arguments are supported:
    /// 
    /// * `aws_account_id` - (ForceNew/Required) The 12 digit AWS Account ID where you want the events to be sent to.
    /// * `aws_region`- (ForceNew/Required) The AWS region where you the events to be sent to. (e.g. us-west-1, us-west-2, ..., etc.)
    /// * `vhost`- (ForceNew/Required) The VHost the queue resides in.
    /// * `queue` - (ForceNew/Required) A (durable) queue on your RabbitMQ instance.
    /// * `with_headers` - (ForceNew/Required) Include message headers in the event data. `({ "headers": { }, "body": { "your": "message" } })`
    /// 
    /// ## Dependency
    /// 
    /// This resource depends on CloudAMQP instance identifier, `cloudamqp_instance.instance.id`.
    /// 
    /// ## Import
    /// 
    /// `cloudamqp_integration_aws_eventbridge` can be imported using CloudAMQP internal identifier of the AWS EventBridge together (CSV separated) with the instance identifier. To retrieve the AWS EventBridge identifier, use [CloudAMQP API](https://docs.cloudamqp.com/cloudamqp_api.html#list-eventbridges)
    /// 
    /// ```sh
    /// $ pulumi import cloudamqp:index/integrationAwsEventbridge:IntegrationAwsEventbridge aws_eventbridge &lt;id&gt;,&lt;instance_id&gt;`
    /// ```
    /// </summary>
    [CloudAmqpResourceType("cloudamqp:index/integrationAwsEventbridge:IntegrationAwsEventbridge")]
    public partial class IntegrationAwsEventbridge : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The 12 digit AWS Account ID where you want the events to be sent to.
        /// </summary>
        [Output("awsAccountId")]
        public Output<string> AwsAccountId { get; private set; } = null!;

        /// <summary>
        /// The AWS region where you the events to be sent to. (e.g. us-west-1, us-west-2, ..., etc.)
        /// </summary>
        [Output("awsRegion")]
        public Output<string> AwsRegion { get; private set; } = null!;

        /// <summary>
        /// Instance identifier
        /// </summary>
        [Output("instanceId")]
        public Output<int> InstanceId { get; private set; } = null!;

        /// <summary>
        /// A (durable) queue on your RabbitMQ instance.
        /// </summary>
        [Output("queue")]
        public Output<string> Queue { get; private set; } = null!;

        /// <summary>
        /// Always set to null, unless there is an error starting the EventBridge.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// The VHost the queue resides in.
        /// </summary>
        [Output("vhost")]
        public Output<string> Vhost { get; private set; } = null!;

        /// <summary>
        /// Include message headers in the event data.
        /// </summary>
        [Output("withHeaders")]
        public Output<bool> WithHeaders { get; private set; } = null!;


        /// <summary>
        /// Create a IntegrationAwsEventbridge resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public IntegrationAwsEventbridge(string name, IntegrationAwsEventbridgeArgs args, CustomResourceOptions? options = null)
            : base("cloudamqp:index/integrationAwsEventbridge:IntegrationAwsEventbridge", name, args ?? new IntegrationAwsEventbridgeArgs(), MakeResourceOptions(options, ""))
        {
        }

        private IntegrationAwsEventbridge(string name, Input<string> id, IntegrationAwsEventbridgeState? state = null, CustomResourceOptions? options = null)
            : base("cloudamqp:index/integrationAwsEventbridge:IntegrationAwsEventbridge", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing IntegrationAwsEventbridge resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static IntegrationAwsEventbridge Get(string name, Input<string> id, IntegrationAwsEventbridgeState? state = null, CustomResourceOptions? options = null)
        {
            return new IntegrationAwsEventbridge(name, id, state, options);
        }
    }

    public sealed class IntegrationAwsEventbridgeArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The 12 digit AWS Account ID where you want the events to be sent to.
        /// </summary>
        [Input("awsAccountId", required: true)]
        public Input<string> AwsAccountId { get; set; } = null!;

        /// <summary>
        /// The AWS region where you the events to be sent to. (e.g. us-west-1, us-west-2, ..., etc.)
        /// </summary>
        [Input("awsRegion", required: true)]
        public Input<string> AwsRegion { get; set; } = null!;

        /// <summary>
        /// Instance identifier
        /// </summary>
        [Input("instanceId", required: true)]
        public Input<int> InstanceId { get; set; } = null!;

        /// <summary>
        /// A (durable) queue on your RabbitMQ instance.
        /// </summary>
        [Input("queue", required: true)]
        public Input<string> Queue { get; set; } = null!;

        /// <summary>
        /// The VHost the queue resides in.
        /// </summary>
        [Input("vhost", required: true)]
        public Input<string> Vhost { get; set; } = null!;

        /// <summary>
        /// Include message headers in the event data.
        /// </summary>
        [Input("withHeaders", required: true)]
        public Input<bool> WithHeaders { get; set; } = null!;

        public IntegrationAwsEventbridgeArgs()
        {
        }
        public static new IntegrationAwsEventbridgeArgs Empty => new IntegrationAwsEventbridgeArgs();
    }

    public sealed class IntegrationAwsEventbridgeState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The 12 digit AWS Account ID where you want the events to be sent to.
        /// </summary>
        [Input("awsAccountId")]
        public Input<string>? AwsAccountId { get; set; }

        /// <summary>
        /// The AWS region where you the events to be sent to. (e.g. us-west-1, us-west-2, ..., etc.)
        /// </summary>
        [Input("awsRegion")]
        public Input<string>? AwsRegion { get; set; }

        /// <summary>
        /// Instance identifier
        /// </summary>
        [Input("instanceId")]
        public Input<int>? InstanceId { get; set; }

        /// <summary>
        /// A (durable) queue on your RabbitMQ instance.
        /// </summary>
        [Input("queue")]
        public Input<string>? Queue { get; set; }

        /// <summary>
        /// Always set to null, unless there is an error starting the EventBridge.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// The VHost the queue resides in.
        /// </summary>
        [Input("vhost")]
        public Input<string>? Vhost { get; set; }

        /// <summary>
        /// Include message headers in the event data.
        /// </summary>
        [Input("withHeaders")]
        public Input<bool>? WithHeaders { get; set; }

        public IntegrationAwsEventbridgeState()
        {
        }
        public static new IntegrationAwsEventbridgeState Empty => new IntegrationAwsEventbridgeState();
    }
}
