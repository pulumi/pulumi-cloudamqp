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
    /// This resource allows you to create and manage third party log integrations for a CloudAMQP instance. Once configured, the logs produced will be forward to corresponding integration.
    /// 
    /// Only available for dedicated subscription plans.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using CloudAmqp = Pulumi.CloudAmqp;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var cloudwatch = new CloudAmqp.IntegrationLog("cloudwatch", new CloudAmqp.IntegrationLogArgs
    ///         {
    ///             InstanceId = cloudamqp_instance.Instance.Id,
    ///             AccessKeyId = @var.Aws_access_key_id,
    ///             SecretAccessKey = @var.Aws_secret_access_key,
    ///             Region = @var.Aws_region,
    ///         });
    ///         var logentries = new CloudAmqp.IntegrationLog("logentries", new CloudAmqp.IntegrationLogArgs
    ///         {
    ///             InstanceId = cloudamqp_instance.Instance.Id,
    ///             Token = @var.Logentries_token,
    ///         });
    ///         var loggly = new CloudAmqp.IntegrationLog("loggly", new CloudAmqp.IntegrationLogArgs
    ///         {
    ///             InstanceId = cloudamqp_instance.Instance.Id,
    ///             Token = @var.Loggly_token,
    ///         });
    ///         var papertrail = new CloudAmqp.IntegrationLog("papertrail", new CloudAmqp.IntegrationLogArgs
    ///         {
    ///             InstanceId = cloudamqp_instance.Instance.Id,
    ///             Url = @var.Papertrail_url,
    ///         });
    ///         var splunk = new CloudAmqp.IntegrationLog("splunk", new CloudAmqp.IntegrationLogArgs
    ///         {
    ///             InstanceId = cloudamqp_instance.Instance.Id,
    ///             Token = @var.Splunk_token,
    ///             HostPort = @var.Splunk_host_port,
    ///         });
    ///         var datadog = new CloudAmqp.IntegrationLog("datadog", new CloudAmqp.IntegrationLogArgs
    ///         {
    ///             InstanceId = cloudamqp_instance.Instance.Id,
    ///             Region = @var.Datadog_region,
    ///             ApiKey = @var.Datadog_api_key,
    ///             Tags = @var.Datadog_tags,
    ///         });
    ///         var stackdriver = new CloudAmqp.IntegrationLog("stackdriver", new CloudAmqp.IntegrationLogArgs
    ///         {
    ///             InstanceId = cloudamqp_instance.Instance.Id,
    ///             ProjectId = @var.Stackdriver_project_id,
    ///             PrivateKey = @var.Stackdriver_private_key,
    ///             ClientEmail = @var.Stackdriver_client_email,
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// ## Argument Reference (cloudwatchlog)
    /// 
    /// Cloudwatch argument reference and example. Create an IAM user with programmatic access and the following permissions:
    /// 
    /// * CreateLogGroup
    /// * CreateLogStream
    /// * DescribeLogGroups
    /// * DescribeLogStreams
    /// * PutLogEvents
    /// 
    /// ## Integration service reference
    /// 
    /// Valid names for third party log integration.
    /// 
    /// | Name       | Description |
    /// |------------|---------------------------------------------------------------|
    /// | cloudwatchlog | Create a IAM with programmatic access. |
    /// | logentries | Create a Logentries token at https://logentries.com/app#/add-log/manual  |
    /// | loggly     | Create a Loggly token at https://{your-company}.loggly.com/tokens |
    /// | papertrail | Create a Papertrail endpoint https://papertrailapp.com/systems/setup |
    /// | splunk     | Create a HTTP Event Collector token at https://.cloud.splunk.com/en-US/manager/search/http-eventcollector |
    /// | datadog       | Create a Datadog API key at app.datadoghq.com |
    /// | stackdriver   | Create a service account and add 'monitor metrics writer' role, then download credentials. |
    /// 
    /// ## Integration Type reference
    /// 
    /// Valid arguments for third party log integrations.
    /// 
    /// Required arguments for all integrations: name
    /// 
    /// | Name | Type | Required arguments |
    /// | ---- | ---- | ---- |
    /// | CloudWatch | cloudwatchlog | access_key_id, secret_access_key, region |
    /// | Log Entries | logentries | token |
    /// | Loggly | loggly | token |
    /// | Papertrail | papertrail | url |
    /// | Splunk | splunk | token, host_port |
    /// | Data Dog | datadog | region, api_keys, tags |
    /// | Stackdriver | stackdriver | project_id, private_key, client_email |
    /// 
    /// ## Dependency
    /// 
    /// This resource depends on CloudAMQP instance identifier, `cloudamqp_instance.instance.id`.
    /// 
    /// ## Import
    /// 
    /// `cloudamqp_integration_log`can be imported using the name argument of the resource together with CloudAMQP instance identifier. The name and identifier are CSV separated, see example below.
    /// 
    /// ```sh
    ///  $ pulumi import cloudamqp:index/integrationLog:IntegrationLog &lt;resource_name&gt; &lt;name&gt;,&lt;instance_id&gt;`
    /// ```
    /// </summary>
    [CloudAmqpResourceType("cloudamqp:index/integrationLog:IntegrationLog")]
    public partial class IntegrationLog : Pulumi.CustomResource
    {
        /// <summary>
        /// AWS access key identifier.
        /// </summary>
        [Output("accessKeyId")]
        public Output<string?> AccessKeyId { get; private set; } = null!;

        /// <summary>
        /// The API key.
        /// </summary>
        [Output("apiKey")]
        public Output<string?> ApiKey { get; private set; } = null!;

        /// <summary>
        /// The client email registered for the integration service.
        /// </summary>
        [Output("clientEmail")]
        public Output<string?> ClientEmail { get; private set; } = null!;

        /// <summary>
        /// Destination to send the logs.
        /// </summary>
        [Output("hostPort")]
        public Output<string?> HostPort { get; private set; } = null!;

        /// <summary>
        /// Instance identifier used to make proxy calls
        /// </summary>
        [Output("instanceId")]
        public Output<int> InstanceId { get; private set; } = null!;

        /// <summary>
        /// The name of the third party log integration. See
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The private access key.
        /// </summary>
        [Output("privateKey")]
        public Output<string?> PrivateKey { get; private set; } = null!;

        /// <summary>
        /// The project identifier.
        /// </summary>
        [Output("projectId")]
        public Output<string?> ProjectId { get; private set; } = null!;

        /// <summary>
        /// Region hosting the integration service.
        /// </summary>
        [Output("region")]
        public Output<string?> Region { get; private set; } = null!;

        /// <summary>
        /// AWS secret access key.
        /// </summary>
        [Output("secretAccessKey")]
        public Output<string?> SecretAccessKey { get; private set; } = null!;

        /// <summary>
        /// Tag the integration, e.g. env=prod, region=europe.
        /// </summary>
        [Output("tags")]
        public Output<string?> Tags { get; private set; } = null!;

        /// <summary>
        /// Token used for authentication.
        /// </summary>
        [Output("token")]
        public Output<string?> Token { get; private set; } = null!;

        /// <summary>
        /// Endpoint to log integration.
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
        /// AWS access key identifier.
        /// </summary>
        [Input("accessKeyId")]
        public Input<string>? AccessKeyId { get; set; }

        /// <summary>
        /// The API key.
        /// </summary>
        [Input("apiKey")]
        public Input<string>? ApiKey { get; set; }

        /// <summary>
        /// The client email registered for the integration service.
        /// </summary>
        [Input("clientEmail")]
        public Input<string>? ClientEmail { get; set; }

        /// <summary>
        /// Destination to send the logs.
        /// </summary>
        [Input("hostPort")]
        public Input<string>? HostPort { get; set; }

        /// <summary>
        /// Instance identifier used to make proxy calls
        /// </summary>
        [Input("instanceId", required: true)]
        public Input<int> InstanceId { get; set; } = null!;

        /// <summary>
        /// The name of the third party log integration. See
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The private access key.
        /// </summary>
        [Input("privateKey")]
        public Input<string>? PrivateKey { get; set; }

        /// <summary>
        /// The project identifier.
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        /// <summary>
        /// Region hosting the integration service.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// AWS secret access key.
        /// </summary>
        [Input("secretAccessKey")]
        public Input<string>? SecretAccessKey { get; set; }

        /// <summary>
        /// Tag the integration, e.g. env=prod, region=europe.
        /// </summary>
        [Input("tags")]
        public Input<string>? Tags { get; set; }

        /// <summary>
        /// Token used for authentication.
        /// </summary>
        [Input("token")]
        public Input<string>? Token { get; set; }

        /// <summary>
        /// Endpoint to log integration.
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
        /// AWS access key identifier.
        /// </summary>
        [Input("accessKeyId")]
        public Input<string>? AccessKeyId { get; set; }

        /// <summary>
        /// The API key.
        /// </summary>
        [Input("apiKey")]
        public Input<string>? ApiKey { get; set; }

        /// <summary>
        /// The client email registered for the integration service.
        /// </summary>
        [Input("clientEmail")]
        public Input<string>? ClientEmail { get; set; }

        /// <summary>
        /// Destination to send the logs.
        /// </summary>
        [Input("hostPort")]
        public Input<string>? HostPort { get; set; }

        /// <summary>
        /// Instance identifier used to make proxy calls
        /// </summary>
        [Input("instanceId")]
        public Input<int>? InstanceId { get; set; }

        /// <summary>
        /// The name of the third party log integration. See
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The private access key.
        /// </summary>
        [Input("privateKey")]
        public Input<string>? PrivateKey { get; set; }

        /// <summary>
        /// The project identifier.
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        /// <summary>
        /// Region hosting the integration service.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// AWS secret access key.
        /// </summary>
        [Input("secretAccessKey")]
        public Input<string>? SecretAccessKey { get; set; }

        /// <summary>
        /// Tag the integration, e.g. env=prod, region=europe.
        /// </summary>
        [Input("tags")]
        public Input<string>? Tags { get; set; }

        /// <summary>
        /// Token used for authentication.
        /// </summary>
        [Input("token")]
        public Input<string>? Token { get; set; }

        /// <summary>
        /// Endpoint to log integration.
        /// </summary>
        [Input("url")]
        public Input<string>? Url { get; set; }

        public IntegrationLogState()
        {
        }
    }
}
