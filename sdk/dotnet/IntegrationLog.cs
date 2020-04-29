// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.CloudAmqp
{
    public partial class IntegrationLog : Pulumi.CustomResource
    {
        /// <summary>
        /// AWS access key identifier. (Cloudwatch)
        /// </summary>
        [Output("accessKeyId")]
        public Output<string?> AccessKeyId { get; private set; } = null!;

        /// <summary>
        /// Destination to send the logs. (Splunk)
        /// </summary>
        [Output("hostPort")]
        public Output<string?> HostPort { get; private set; } = null!;

        /// <summary>
        /// Instance identifier used to make proxy calls
        /// </summary>
        [Output("instanceId")]
        public Output<int> InstanceId { get; private set; } = null!;

        /// <summary>
        /// The name of log integration
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The region hosting integration service. (Cloudwatch)
        /// </summary>
        [Output("region")]
        public Output<string?> Region { get; private set; } = null!;

        /// <summary>
        /// AWS secret access key. (Cloudwatch)
        /// </summary>
        [Output("secretAccessKey")]
        public Output<string?> SecretAccessKey { get; private set; } = null!;

        /// <summary>
        /// The token used for authentication. (Loggly, Logentries, Splunk)
        /// </summary>
        [Output("token")]
        public Output<string?> Token { get; private set; } = null!;

        /// <summary>
        /// The URL to push the logs to. (Papertrail)
        /// </summary>
        [Output("url")]
        public Output<string?> Url { get; private set; } = null!;


        /// <summary>
        /// Create a IntegrationLog resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public IntegrationLog(string name, IntegrationLogArgs args, CustomResourceOptions? options = null)
            : base("cloudamqp:index/integrationLog:IntegrationLog", name, args ?? new IntegrationLogArgs(), MakeResourceOptions(options, ""))
        {
        }

        private IntegrationLog(string name, Input<string> id, IntegrationLogState? state = null, CustomResourceOptions? options = null)
            : base("cloudamqp:index/integrationLog:IntegrationLog", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing IntegrationLog resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static IntegrationLog Get(string name, Input<string> id, IntegrationLogState? state = null, CustomResourceOptions? options = null)
        {
            return new IntegrationLog(name, id, state, options);
        }
    }

    public sealed class IntegrationLogArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// AWS access key identifier. (Cloudwatch)
        /// </summary>
        [Input("accessKeyId")]
        public Input<string>? AccessKeyId { get; set; }

        /// <summary>
        /// Destination to send the logs. (Splunk)
        /// </summary>
        [Input("hostPort")]
        public Input<string>? HostPort { get; set; }

        /// <summary>
        /// Instance identifier used to make proxy calls
        /// </summary>
        [Input("instanceId", required: true)]
        public Input<int> InstanceId { get; set; } = null!;

        /// <summary>
        /// The name of log integration
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The region hosting integration service. (Cloudwatch)
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// AWS secret access key. (Cloudwatch)
        /// </summary>
        [Input("secretAccessKey")]
        public Input<string>? SecretAccessKey { get; set; }

        /// <summary>
        /// The token used for authentication. (Loggly, Logentries, Splunk)
        /// </summary>
        [Input("token")]
        public Input<string>? Token { get; set; }

        /// <summary>
        /// The URL to push the logs to. (Papertrail)
        /// </summary>
        [Input("url")]
        public Input<string>? Url { get; set; }

        public IntegrationLogArgs()
        {
        }
    }

    public sealed class IntegrationLogState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// AWS access key identifier. (Cloudwatch)
        /// </summary>
        [Input("accessKeyId")]
        public Input<string>? AccessKeyId { get; set; }

        /// <summary>
        /// Destination to send the logs. (Splunk)
        /// </summary>
        [Input("hostPort")]
        public Input<string>? HostPort { get; set; }

        /// <summary>
        /// Instance identifier used to make proxy calls
        /// </summary>
        [Input("instanceId")]
        public Input<int>? InstanceId { get; set; }

        /// <summary>
        /// The name of log integration
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The region hosting integration service. (Cloudwatch)
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// AWS secret access key. (Cloudwatch)
        /// </summary>
        [Input("secretAccessKey")]
        public Input<string>? SecretAccessKey { get; set; }

        /// <summary>
        /// The token used for authentication. (Loggly, Logentries, Splunk)
        /// </summary>
        [Input("token")]
        public Input<string>? Token { get; set; }

        /// <summary>
        /// The URL to push the logs to. (Papertrail)
        /// </summary>
        [Input("url")]
        public Input<string>? Url { get; set; }

        public IntegrationLogState()
        {
        }
    }
}
