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
    /// This resource allows you to create and manage, forwarding metrics to third party integrations for a
    /// CloudAMQP instance. Once configured, the metrics produced will be forward to corresponding
    /// integration.
    /// 
    /// Only available for dedicated subscription plans.
    /// 
    /// ## Example Usage
    /// 
    /// &lt;details&gt;
    ///   &lt;summary&gt;
    ///     &lt;b&gt;
    ///       &lt;i&gt;Cloudwatch v1 and v2 metric integration&lt;/i&gt;
    ///     &lt;/b&gt;
    ///   &lt;/summary&gt;
    /// 
    /// ***Access key***
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using CloudAmqp = Pulumi.CloudAmqp;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var cloudwatch = new CloudAmqp.IntegrationMetric("cloudwatch", new()
    ///     {
    ///         InstanceId = instance.Id,
    ///         Name = "cloudwatch",
    ///         AccessKeyId = awsAccessKeyId,
    ///         SecretAccessKey = varAwsSecretAcccessKey,
    ///         Region = awsRegion,
    ///     });
    /// 
    ///     var cloudwatchV2 = new CloudAmqp.IntegrationMetric("cloudwatch_v2", new()
    ///     {
    ///         InstanceId = instance.Id,
    ///         Name = "cloudwatch_v2",
    ///         AccessKeyId = awsAccessKeyId,
    ///         SecretAccessKey = varAwsSecretAcccessKey,
    ///         Region = awsRegion,
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ***Assume role***
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using CloudAmqp = Pulumi.CloudAmqp;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var cloudwatch = new CloudAmqp.IntegrationMetric("cloudwatch", new()
    ///     {
    ///         InstanceId = instance.Id,
    ///         Name = "cloudwatch",
    ///         IamRole = awsIamRole,
    ///         IamExternalId = externalId,
    ///         Region = awsRegion,
    ///     });
    /// 
    ///     var cloudwatchV2 = new CloudAmqp.IntegrationMetric("cloudwatch_v2", new()
    ///     {
    ///         InstanceId = instance.Id,
    ///         Name = "cloudwatch_v2",
    ///         IamRole = awsIamRole,
    ///         IamExternalId = externalId,
    ///         Region = awsRegion,
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// * AWS IAM role: arn:aws:iam::ACCOUNT-ID:role/ROLE-NAME
    /// * External id: Create own external identifier that match the role created. E.g. "cloudamqp-abc123".
    /// 
    /// &lt;/details&gt;
    /// 
    /// &lt;details&gt;
    ///   &lt;summary&gt;
    ///     &lt;b&gt;
    ///       &lt;i&gt;Datadog v1 and v2 metric integration&lt;/i&gt;
    ///     &lt;/b&gt;
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
    ///     var datadog = new CloudAmqp.IntegrationMetric("datadog", new()
    ///     {
    ///         InstanceId = instance.Id,
    ///         Name = "datadog",
    ///         ApiKey = datadogApiKey,
    ///         Region = datadogRegion,
    ///         Tags = "env=prod,region=us1,version=v1.0",
    ///     });
    /// 
    ///     var datadogV2 = new CloudAmqp.IntegrationMetric("datadog_v2", new()
    ///     {
    ///         InstanceId = instance.Id,
    ///         Name = "datadog_v2",
    ///         ApiKey = datadogApiKey,
    ///         Region = datadogRegion,
    ///         Tags = "env=prod,region=us1,version=v1.0",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// &lt;/details&gt;
    /// 
    /// &lt;details&gt;
    ///   &lt;summary&gt;
    ///     &lt;b&gt;
    ///       &lt;i&gt;Librato metric integration&lt;/i&gt;
    ///     &lt;/b&gt;
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
    ///     var librato = new CloudAmqp.IntegrationMetric("librato", new()
    ///     {
    ///         InstanceId = instance.Id,
    ///         Name = "librato",
    ///         Email = libratoEmail,
    ///         ApiKey = libratoApiKey,
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// &lt;/details&gt;
    /// 
    /// &lt;details&gt;
    ///   &lt;summary&gt;
    ///     &lt;b&gt;
    ///       &lt;i&gt;New relic v2 metric integration&lt;/i&gt;
    ///     &lt;/b&gt;
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
    ///     var newrelic = new CloudAmqp.IntegrationMetric("newrelic", new()
    ///     {
    ///         InstanceId = instance.Id,
    ///         Name = "newrelic_v2",
    ///         ApiKey = newrelicApiKey,
    ///         Region = newrelicRegion,
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// &lt;/details&gt;
    /// 
    /// &lt;details&gt;
    ///   &lt;summary&gt;
    ///     &lt;b&gt;
    ///       &lt;i&gt;Stackdriver metric integration (v1.20.2 or earlier versions)&lt;/i&gt;
    ///     &lt;/b&gt;
    ///   &lt;/summary&gt;
    /// 
    /// Use variable file populated with project_id, private_key and client_email
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using CloudAmqp = Pulumi.CloudAmqp;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var stackdriver = new CloudAmqp.IntegrationMetric("stackdriver", new()
    ///     {
    ///         InstanceId = instance.Id,
    ///         Name = "stackdriver",
    ///         ProjectId = stackdriverProjectId,
    ///         PrivateKey = stackdriverPrivateKey,
    ///         ClientEmail = stackriverEmail,
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// or by using google_service_account_key resource from Google provider
    /// 
    /// ## Import
    /// 
    /// `cloudamqp_integration_metric`can be imported using the resource identifier together with CloudAMQP
    /// 
    /// instance identifier (CSV separated). To retrieve the resource identifier, use
    /// 
    /// [CloudAMQP API list integrations].
    /// 
    /// From Terraform v1.5.0, the `import` block can be used to import this resource:
    /// 
    /// hcl
    /// 
    /// import {
    /// 
    ///   to = cloudamqp_alarm.alarm
    /// 
    ///   id = format("&lt;id&gt;,%s", cloudamqp_instance.instance.id)
    /// 
    /// }
    /// 
    /// Or use Terraform CLI:
    /// 
    /// ```sh
    /// $ pulumi import cloudamqp:index/integrationMetric:IntegrationMetric &lt;resource_name&gt; &lt;resource_id&gt;,&lt;instance_id&gt;`
    /// ```
    /// 
    /// [CloudAMQP API add integrations]: https://docs.cloudamqp.com/cloudamqp_api.html#add-metrics-integration
    /// 
    /// [CloudAMQP API list integrations]: https://docs.cloudamqp.com/cloudamqp_api.html#list-metrics-integrations
    /// 
    /// [Datadog documentation]: https://docs.datadoghq.com/getting_started/tagging/#define-tags
    /// 
    /// [integration type reference]: #integration-type-reference
    /// </summary>
    [CloudAmqpResourceType("cloudamqp:index/integrationMetric:IntegrationMetric")]
    public partial class IntegrationMetric : global::Pulumi.CustomResource
    {
        /// <summary>
        /// AWS access key identifier. (Cloudwatch)
        /// </summary>
        [Output("accessKeyId")]
        public Output<string?> AccessKeyId { get; private set; } = null!;

        /// <summary>
        /// The API key for the integration service. (Librato)
        /// </summary>
        [Output("apiKey")]
        public Output<string?> ApiKey { get; private set; } = null!;

        /// <summary>
        /// The client email. (Stackdriver)
        /// </summary>
        [Output("clientEmail")]
        public Output<string> ClientEmail { get; private set; } = null!;

        /// <summary>
        /// Base64Encoded credentials. (Stackdriver)
        /// </summary>
        [Output("credentials")]
        public Output<string?> Credentials { get; private set; } = null!;

        /// <summary>
        /// The email address registred for the integration service. (Librato)
        /// </summary>
        [Output("email")]
        public Output<string?> Email { get; private set; } = null!;

        /// <summary>
        /// External identifier that match the role you created. (Cloudwatch)
        /// </summary>
        [Output("iamExternalId")]
        public Output<string?> IamExternalId { get; private set; } = null!;

        /// <summary>
        /// The ARN of the role to be assumed when publishing metrics. (Cloudwatch)
        /// </summary>
        [Output("iamRole")]
        public Output<string?> IamRole { get; private set; } = null!;

        /// <summary>
        /// (optional) Include Auto-Delete queues
        /// </summary>
        [Output("includeAdQueues")]
        public Output<bool?> IncludeAdQueues { get; private set; } = null!;

        /// <summary>
        /// Instance identifier
        /// </summary>
        [Output("instanceId")]
        public Output<int> InstanceId { get; private set; } = null!;

        /// <summary>
        /// The license key registred for the integration service. (New Relic)
        /// </summary>
        [Output("licenseKey")]
        public Output<string?> LicenseKey { get; private set; } = null!;

        /// <summary>
        /// The name of metrics integration
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The private key. (Stackdriver)
        /// </summary>
        [Output("privateKey")]
        public Output<string> PrivateKey { get; private set; } = null!;

        /// <summary>
        /// Private key identifier. (Stackdriver)
        /// </summary>
        [Output("privateKeyId")]
        public Output<string> PrivateKeyId { get; private set; } = null!;

        /// <summary>
        /// Project ID. (Stackdriver)
        /// </summary>
        [Output("projectId")]
        public Output<string> ProjectId { get; private set; } = null!;

        /// <summary>
        /// (optional) allowlist using regular expression
        /// </summary>
        [Output("queueAllowlist")]
        public Output<string?> QueueAllowlist { get; private set; } = null!;

        /// <summary>
        /// **Deprecated**
        /// </summary>
        [Output("queueWhitelist")]
        public Output<string?> QueueWhitelist { get; private set; } = null!;

        /// <summary>
        /// AWS region for Cloudwatch and [US/EU] for Data dog/New relic. (Cloudwatch, Data Dog, New Relic)
        /// </summary>
        [Output("region")]
        public Output<string?> Region { get; private set; } = null!;

        /// <summary>
        /// AWS secret key. (Cloudwatch)
        /// </summary>
        [Output("secretAccessKey")]
        public Output<string?> SecretAccessKey { get; private set; } = null!;

        /// <summary>
        /// (optional) tags. E.g. env=prod,region=europe
        /// </summary>
        [Output("tags")]
        public Output<string?> Tags { get; private set; } = null!;

        /// <summary>
        /// (optional) allowlist using regular expression
        /// </summary>
        [Output("vhostAllowlist")]
        public Output<string?> VhostAllowlist { get; private set; } = null!;

        /// <summary>
        /// **Deprecated**
        /// </summary>
        [Output("vhostWhitelist")]
        public Output<string?> VhostWhitelist { get; private set; } = null!;


        /// <summary>
        /// Create a IntegrationMetric resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public IntegrationMetric(string name, IntegrationMetricArgs args, CustomResourceOptions? options = null)
            : base("cloudamqp:index/integrationMetric:IntegrationMetric", name, args ?? new IntegrationMetricArgs(), MakeResourceOptions(options, ""))
        {
        }

        private IntegrationMetric(string name, Input<string> id, IntegrationMetricState? state = null, CustomResourceOptions? options = null)
            : base("cloudamqp:index/integrationMetric:IntegrationMetric", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                AdditionalSecretOutputs =
                {
                    "apiKey",
                    "credentials",
                    "licenseKey",
                    "privateKey",
                    "privateKeyId",
                    "secretAccessKey",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing IntegrationMetric resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static IntegrationMetric Get(string name, Input<string> id, IntegrationMetricState? state = null, CustomResourceOptions? options = null)
        {
            return new IntegrationMetric(name, id, state, options);
        }
    }

    public sealed class IntegrationMetricArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// AWS access key identifier. (Cloudwatch)
        /// </summary>
        [Input("accessKeyId")]
        public Input<string>? AccessKeyId { get; set; }

        [Input("apiKey")]
        private Input<string>? _apiKey;

        /// <summary>
        /// The API key for the integration service. (Librato)
        /// </summary>
        public Input<string>? ApiKey
        {
            get => _apiKey;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _apiKey = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// The client email. (Stackdriver)
        /// </summary>
        [Input("clientEmail")]
        public Input<string>? ClientEmail { get; set; }

        [Input("credentials")]
        private Input<string>? _credentials;

        /// <summary>
        /// Base64Encoded credentials. (Stackdriver)
        /// </summary>
        public Input<string>? Credentials
        {
            get => _credentials;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _credentials = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// The email address registred for the integration service. (Librato)
        /// </summary>
        [Input("email")]
        public Input<string>? Email { get; set; }

        /// <summary>
        /// External identifier that match the role you created. (Cloudwatch)
        /// </summary>
        [Input("iamExternalId")]
        public Input<string>? IamExternalId { get; set; }

        /// <summary>
        /// The ARN of the role to be assumed when publishing metrics. (Cloudwatch)
        /// </summary>
        [Input("iamRole")]
        public Input<string>? IamRole { get; set; }

        /// <summary>
        /// (optional) Include Auto-Delete queues
        /// </summary>
        [Input("includeAdQueues")]
        public Input<bool>? IncludeAdQueues { get; set; }

        /// <summary>
        /// Instance identifier
        /// </summary>
        [Input("instanceId", required: true)]
        public Input<int> InstanceId { get; set; } = null!;

        [Input("licenseKey")]
        private Input<string>? _licenseKey;

        /// <summary>
        /// The license key registred for the integration service. (New Relic)
        /// </summary>
        public Input<string>? LicenseKey
        {
            get => _licenseKey;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _licenseKey = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// The name of metrics integration
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("privateKey")]
        private Input<string>? _privateKey;

        /// <summary>
        /// The private key. (Stackdriver)
        /// </summary>
        public Input<string>? PrivateKey
        {
            get => _privateKey;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _privateKey = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        [Input("privateKeyId")]
        private Input<string>? _privateKeyId;

        /// <summary>
        /// Private key identifier. (Stackdriver)
        /// </summary>
        public Input<string>? PrivateKeyId
        {
            get => _privateKeyId;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _privateKeyId = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// Project ID. (Stackdriver)
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        /// <summary>
        /// (optional) allowlist using regular expression
        /// </summary>
        [Input("queueAllowlist")]
        public Input<string>? QueueAllowlist { get; set; }

        /// <summary>
        /// **Deprecated**
        /// </summary>
        [Input("queueWhitelist")]
        public Input<string>? QueueWhitelist { get; set; }

        /// <summary>
        /// AWS region for Cloudwatch and [US/EU] for Data dog/New relic. (Cloudwatch, Data Dog, New Relic)
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        [Input("secretAccessKey")]
        private Input<string>? _secretAccessKey;

        /// <summary>
        /// AWS secret key. (Cloudwatch)
        /// </summary>
        public Input<string>? SecretAccessKey
        {
            get => _secretAccessKey;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _secretAccessKey = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// (optional) tags. E.g. env=prod,region=europe
        /// </summary>
        [Input("tags")]
        public Input<string>? Tags { get; set; }

        /// <summary>
        /// (optional) allowlist using regular expression
        /// </summary>
        [Input("vhostAllowlist")]
        public Input<string>? VhostAllowlist { get; set; }

        /// <summary>
        /// **Deprecated**
        /// </summary>
        [Input("vhostWhitelist")]
        public Input<string>? VhostWhitelist { get; set; }

        public IntegrationMetricArgs()
        {
        }
        public static new IntegrationMetricArgs Empty => new IntegrationMetricArgs();
    }

    public sealed class IntegrationMetricState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// AWS access key identifier. (Cloudwatch)
        /// </summary>
        [Input("accessKeyId")]
        public Input<string>? AccessKeyId { get; set; }

        [Input("apiKey")]
        private Input<string>? _apiKey;

        /// <summary>
        /// The API key for the integration service. (Librato)
        /// </summary>
        public Input<string>? ApiKey
        {
            get => _apiKey;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _apiKey = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// The client email. (Stackdriver)
        /// </summary>
        [Input("clientEmail")]
        public Input<string>? ClientEmail { get; set; }

        [Input("credentials")]
        private Input<string>? _credentials;

        /// <summary>
        /// Base64Encoded credentials. (Stackdriver)
        /// </summary>
        public Input<string>? Credentials
        {
            get => _credentials;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _credentials = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// The email address registred for the integration service. (Librato)
        /// </summary>
        [Input("email")]
        public Input<string>? Email { get; set; }

        /// <summary>
        /// External identifier that match the role you created. (Cloudwatch)
        /// </summary>
        [Input("iamExternalId")]
        public Input<string>? IamExternalId { get; set; }

        /// <summary>
        /// The ARN of the role to be assumed when publishing metrics. (Cloudwatch)
        /// </summary>
        [Input("iamRole")]
        public Input<string>? IamRole { get; set; }

        /// <summary>
        /// (optional) Include Auto-Delete queues
        /// </summary>
        [Input("includeAdQueues")]
        public Input<bool>? IncludeAdQueues { get; set; }

        /// <summary>
        /// Instance identifier
        /// </summary>
        [Input("instanceId")]
        public Input<int>? InstanceId { get; set; }

        [Input("licenseKey")]
        private Input<string>? _licenseKey;

        /// <summary>
        /// The license key registred for the integration service. (New Relic)
        /// </summary>
        public Input<string>? LicenseKey
        {
            get => _licenseKey;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _licenseKey = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// The name of metrics integration
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("privateKey")]
        private Input<string>? _privateKey;

        /// <summary>
        /// The private key. (Stackdriver)
        /// </summary>
        public Input<string>? PrivateKey
        {
            get => _privateKey;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _privateKey = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        [Input("privateKeyId")]
        private Input<string>? _privateKeyId;

        /// <summary>
        /// Private key identifier. (Stackdriver)
        /// </summary>
        public Input<string>? PrivateKeyId
        {
            get => _privateKeyId;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _privateKeyId = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// Project ID. (Stackdriver)
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        /// <summary>
        /// (optional) allowlist using regular expression
        /// </summary>
        [Input("queueAllowlist")]
        public Input<string>? QueueAllowlist { get; set; }

        /// <summary>
        /// **Deprecated**
        /// </summary>
        [Input("queueWhitelist")]
        public Input<string>? QueueWhitelist { get; set; }

        /// <summary>
        /// AWS region for Cloudwatch and [US/EU] for Data dog/New relic. (Cloudwatch, Data Dog, New Relic)
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        [Input("secretAccessKey")]
        private Input<string>? _secretAccessKey;

        /// <summary>
        /// AWS secret key. (Cloudwatch)
        /// </summary>
        public Input<string>? SecretAccessKey
        {
            get => _secretAccessKey;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _secretAccessKey = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// (optional) tags. E.g. env=prod,region=europe
        /// </summary>
        [Input("tags")]
        public Input<string>? Tags { get; set; }

        /// <summary>
        /// (optional) allowlist using regular expression
        /// </summary>
        [Input("vhostAllowlist")]
        public Input<string>? VhostAllowlist { get; set; }

        /// <summary>
        /// **Deprecated**
        /// </summary>
        [Input("vhostWhitelist")]
        public Input<string>? VhostWhitelist { get; set; }

        public IntegrationMetricState()
        {
        }
        public static new IntegrationMetricState Empty => new IntegrationMetricState();
    }
}
