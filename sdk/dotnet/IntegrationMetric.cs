// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.CloudAmqp
{
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
                    "credentials",
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

        /// <summary>
        /// The API key for the integration service. (Librato)
        /// </summary>
        [Input("apiKey")]
        public Input<string>? ApiKey { get; set; }

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
        /// Instance identifier
        /// </summary>
        [Input("instanceId", required: true)]
        public Input<int> InstanceId { get; set; } = null!;

        /// <summary>
        /// The license key registred for the integration service. (New Relic)
        /// </summary>
        [Input("licenseKey")]
        public Input<string>? LicenseKey { get; set; }

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

        /// <summary>
        /// The API key for the integration service. (Librato)
        /// </summary>
        [Input("apiKey")]
        public Input<string>? ApiKey { get; set; }

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
        /// Instance identifier
        /// </summary>
        [Input("instanceId")]
        public Input<int>? InstanceId { get; set; }

        /// <summary>
        /// The license key registred for the integration service. (New Relic)
        /// </summary>
        [Input("licenseKey")]
        public Input<string>? LicenseKey { get; set; }

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
