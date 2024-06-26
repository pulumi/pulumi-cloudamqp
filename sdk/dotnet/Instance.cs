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
    /// This resource allows you to create and manage a CloudAMQP instance running either [**RabbitMQ**](https://www.rabbitmq.com/) or [**LavinMQ**](https://lavinmq.com/) and can be deployed to multiple cloud platforms provider and regions, see instance regions for more information.
    /// 
    /// Once the instance is created it will be assigned a unique identifier. All other resources and data sources created for this instance needs to reference this unique instance identifier.
    /// 
    /// Pricing is available at [cloudamqp.com](https://www.cloudamqp.com/plans.html).
    /// 
    /// ## Example Usage
    /// 
    /// &lt;details&gt;
    ///   &lt;summary&gt;
    ///     &lt;b&gt;
    ///       &lt;i&gt;Basic example of shared and dedicated instances&lt;/i&gt;
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
    ///     // Minimum free lemur instance running RabbitMQ
    ///     var lemurInstance = new CloudAmqp.Instance("lemur_instance", new()
    ///     {
    ///         Name = "cloudamqp-free-instance",
    ///         Plan = "lemur",
    ///         Region = "amazon-web-services::us-west-1",
    ///         Tags = new[]
    ///         {
    ///             "rabbitmq",
    ///         },
    ///     });
    /// 
    ///     // Minimum free lemming instance running LavinMQ
    ///     var lemmingInstance = new CloudAmqp.Instance("lemming_instance", new()
    ///     {
    ///         Name = "cloudamqp-free-instance",
    ///         Plan = "lemming",
    ///         Region = "amazon-web-services::us-west-1",
    ///         Tags = new[]
    ///         {
    ///             "lavinmq",
    ///         },
    ///     });
    /// 
    ///     // New dedicated bunny instance running RabbitMQ
    ///     var instance = new CloudAmqp.Instance("instance", new()
    ///     {
    ///         Name = "terraform-cloudamqp-instance",
    ///         Plan = "bunny-1",
    ///         Region = "amazon-web-services::us-west-1",
    ///         Tags = new[]
    ///         {
    ///             "terraform",
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
    ///     &lt;b&gt;
    ///       &lt;i&gt;Dedicated instance using attribute vpc_subnet to create VPC, before v1.16.0&lt;/i&gt;
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
    ///     var instance = new CloudAmqp.Instance("instance", new()
    ///     {
    ///         Name = "terraform-cloudamqp-instance",
    ///         Plan = "bunny-1",
    ///         Region = "amazon-web-services::us-west-1",
    ///         Tags = new[]
    ///         {
    ///             "terraform",
    ///         },
    ///         VpcSubnet = "10.56.72.0/24",
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
    ///       &lt;i&gt;Dedicated instance using attribute vpc_subnet to create VPC and then import managed VPC, from v1.16.0 (Managed VPC)&lt;/i&gt;
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
    ///     // Dedicated instance that also creates VPC
    ///     var instance01 = new CloudAmqp.Instance("instance_01", new()
    ///     {
    ///         Name = "terraform-cloudamqp-instance-01",
    ///         Plan = "bunny-1",
    ///         Region = "amazon-web-services::us-west-1",
    ///         Tags = new[]
    ///         {
    ///             "terraform",
    ///         },
    ///         VpcSubnet = "10.56.72.0/24",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// Once the instance and the VPC are created, the VPC can be imported as managed VPC and added to the configuration file.
    /// Set attribute `vpc_id` to the managed VPC identifier. To keep the managed VPC when deleting the instance, set attribute `keep_associated_vpc` to true.
    /// For more information see guide Managed VPC.
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using CloudAmqp = Pulumi.CloudAmqp;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     // Imported managed VPC
    ///     var vpc = new CloudAmqp.Vpc("vpc", new()
    ///     {
    ///         Name = "&lt;vpc-name&gt;",
    ///         Region = "amazon-web-services::us-east-1",
    ///         Subnet = "10.56.72.0/24",
    ///         Tags = new[] {},
    ///     });
    /// 
    ///     // Add vpc_id and keep_associated_vpc attributes
    ///     var instance01 = new CloudAmqp.Instance("instance_01", new()
    ///     {
    ///         Name = "terraform-cloudamqp-instance-01",
    ///         Plan = "bunny-1",
    ///         Region = "amazon-web-services::us-west-1",
    ///         Tags = new[]
    ///         {
    ///             "terraform",
    ///         },
    ///         VpcId = vpc.Id,
    ///         KeepAssociatedVpc = true,
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
    ///       &lt;i&gt;Dedicated instances and managed VPC, from v1.16.0 (Managed VPC)&lt;/i&gt;
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
    ///     // Managed VPC
    ///     var vpc = new CloudAmqp.Vpc("vpc", new()
    ///     {
    ///         Name = "&lt;vpc-name&gt;",
    ///         Region = "amazon-web-services::us-east-1",
    ///         Subnet = "10.56.72.0/24",
    ///         Tags = new[] {},
    ///     });
    /// 
    ///     // First instance added to managed VPC
    ///     var instance01 = new CloudAmqp.Instance("instance_01", new()
    ///     {
    ///         Name = "terraform-cloudamqp-instance-01",
    ///         Plan = "bunny-1",
    ///         Region = "amazon-web-services::us-west-1",
    ///         Tags = new[]
    ///         {
    ///             "terraform",
    ///         },
    ///         VpcId = vpc.Id,
    ///         KeepAssociatedVpc = true,
    ///     });
    /// 
    ///     // Second instance added to managed VPC
    ///     var instance02 = new CloudAmqp.Instance("instance_02", new()
    ///     {
    ///         Name = "terraform-cloudamqp-instance-02",
    ///         Plan = "bunny-1",
    ///         Region = "amazon-web-services::us-west-1",
    ///         Tags = new[]
    ///         {
    ///             "terraform",
    ///         },
    ///         VpcId = vpc.Id,
    ///         KeepAssociatedVpc = true,
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// Set attribute `keep_associated_vpc` to true, will keep managed VPC when deleting the instances.
    /// 
    /// &lt;/details&gt;
    /// 
    /// ## Copy settings to a new dedicated instance
    /// 
    /// With copy settings it's possible to create a new dedicated instance with settings such as alarms, config, etc. from another dedicated instance. This can be done by adding the `copy_settings` block to this resource and populate `subscription_id` with a CloudAMQP instance identifier from another already existing instance.
    /// 
    /// Then add the settings to be copied over to the new dedicated instance. Settings that can be copied [alarms, config, definitions, firewall, logs, metrics, plugins]
    /// 
    /// &gt; `rmq_version` argument is required when doing this action. Must match the RabbitMQ version of the dedicated instance to be copied from.
    /// 
    /// &lt;details&gt;
    ///   &lt;summary&gt;
    ///     &lt;b&gt;
    ///       &lt;i&gt;Copy settings from a dedicated instance to a new dedicated instance&lt;/i&gt;
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
    ///     var instance02 = new CloudAmqp.Instance("instance_02", new()
    ///     {
    ///         Name = "terraform-cloudamqp-instance-02",
    ///         Plan = "squirrel-1",
    ///         Region = "amazon-web-services::us-west-1",
    ///         RmqVersion = "3.12.2",
    ///         Tags = new[]
    ///         {
    ///             "terraform",
    ///         },
    ///         CopySettings = new[]
    ///         {
    ///             new CloudAmqp.Inputs.InstanceCopySettingArgs
    ///             {
    ///                 SubscriptionId = instanceId,
    ///                 Settings = new[]
    ///                 {
    ///                     "alarms",
    ///                     "config",
    ///                     "definitions",
    ///                     "firewall",
    ///                     "logs",
    ///                     "metrics",
    ///                     "plugins",
    ///                 },
    ///             },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// &lt;/details&gt;
    /// 
    /// ## Import
    /// 
    /// `cloudamqp_instance`can be imported using CloudAMQP internal identifier.
    /// 
    /// ```sh
    /// $ pulumi import cloudamqp:index/instance:Instance instance &lt;id&gt;`
    /// ```
    /// 
    /// To retrieve the identifier for an instance, either use [CloudAMQP customer API](https://docs.cloudamqp.com/#list-instances) or use the data source [`cloudamqp_account`](./data-sources/account.md) to list all available instances for an account.
    /// </summary>
    [CloudAmqpResourceType("cloudamqp:index/instance:Instance")]
    public partial class Instance : global::Pulumi.CustomResource
    {
        /// <summary>
        /// API key needed to communicate to CloudAMQP's second API. The second API is used to manage alarms, integration and more, full description [CloudAMQP API](https://docs.cloudamqp.com/cloudamqp_api.html).
        /// </summary>
        [Output("apikey")]
        public Output<string> Apikey { get; private set; } = null!;

        /// <summary>
        /// Information if the CloudAMQP instance runs either RabbitMQ or LavinMQ.
        /// </summary>
        [Output("backend")]
        public Output<string> Backend { get; private set; } = null!;

        /// <summary>
        /// Copy settings from one CloudAMQP instance to a new. Consists of the block documented below.
        /// 
        /// ___
        /// 
        /// The `copy_settings` block consists of:
        /// </summary>
        [Output("copySettings")]
        public Output<ImmutableArray<Outputs.InstanceCopySetting>> CopySettings { get; private set; } = null!;

        /// <summary>
        /// Information if the CloudAMQP instance is shared or dedicated.
        /// </summary>
        [Output("dedicated")]
        public Output<bool> Dedicated { get; private set; } = null!;

        /// <summary>
        /// The external hostname for the CloudAMQP instance.
        /// </summary>
        [Output("host")]
        public Output<string> Host { get; private set; } = null!;

        /// <summary>
        /// The internal hostname for the CloudAMQP instance.
        /// </summary>
        [Output("hostInternal")]
        public Output<string> HostInternal { get; private set; } = null!;

        /// <summary>
        /// Keep associated VPC when deleting instance, default set to false.
        /// </summary>
        [Output("keepAssociatedVpc")]
        public Output<bool?> KeepAssociatedVpc { get; private set; } = null!;

        /// <summary>
        /// Name of the CloudAMQP instance.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Set to true to discard creating default alarms when the instance is created. Can be left out, will then use default value = false.
        /// </summary>
        [Output("noDefaultAlarms")]
        public Output<bool> NoDefaultAlarms { get; private set; } = null!;

        /// <summary>
        /// Number of nodes, 1, 3 or 5 depending on plan used. Only needed for legacy plans, will otherwise be computed.
        /// 
        /// ***Deprecated: Legacy subscriptions plan can still change this to scale up or down the instance. New subscriptions plans use the plan to determine number of nodes. In order to change number of nodes the `plan` needs to be updated.***
        /// </summary>
        [Output("nodes")]
        public Output<int> Nodes { get; private set; } = null!;

        /// <summary>
        /// The subscription plan. See available plans
        /// </summary>
        [Output("plan")]
        public Output<string> Plan { get; private set; } = null!;

        /// <summary>
        /// Flag describing if the resource is ready
        /// </summary>
        [Output("ready")]
        public Output<bool> Ready { get; private set; } = null!;

        /// <summary>
        /// The region to host the instance in. See instance regions
        /// 
        /// ***Note: Changing region will force the instance to be destroyed and a new created in the new region. All data will be lost and a new name assigned.***
        /// </summary>
        [Output("region")]
        public Output<string> Region { get; private set; } = null!;

        /// <summary>
        /// The Rabbit MQ version. Can be left out, will then be set to default value used by CloudAMQP API.
        /// 
        /// ***Note: There is not yet any support in the provider to change the RMQ version. Once it's set in the initial creation, it will remain.***
        /// </summary>
        [Output("rmqVersion")]
        public Output<string> RmqVersion { get; private set; } = null!;

        /// <summary>
        /// One or more tags for the CloudAMQP instance, makes it possible to categories multiple instances in console view. Default there is no tags assigned.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableArray<string>> Tags { get; private set; } = null!;

        /// <summary>
        /// The AMQP URL (uses the internal hostname if the instance was created with VPC). Has the format: `amqps://{username}:{password}@{hostname}/{vhost}`
        /// </summary>
        [Output("url")]
        public Output<string> Url { get; private set; } = null!;

        /// <summary>
        /// The virtual host used by Rabbit MQ.
        /// </summary>
        [Output("vhost")]
        public Output<string> Vhost { get; private set; } = null!;

        /// <summary>
        /// The VPC ID. Use this to create your instance in an existing VPC. See available example.
        /// </summary>
        [Output("vpcId")]
        public Output<int> VpcId { get; private set; } = null!;

        /// <summary>
        /// Creates a dedicated VPC subnet, shouldn't overlap with other VPC subnet, default subnet used 10.56.72.0/24.
        /// 
        /// ***Deprecated: Will be removed in next major version (v2.0)***
        /// 
        /// ***Note: extra fee will be charged when using VPC, see [CloudAMQP](https://cloudamqp.com) for more information.***
        /// </summary>
        [Output("vpcSubnet")]
        public Output<string> VpcSubnet { get; private set; } = null!;


        /// <summary>
        /// Create a Instance resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Instance(string name, InstanceArgs args, CustomResourceOptions? options = null)
            : base("cloudamqp:index/instance:Instance", name, args ?? new InstanceArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Instance(string name, Input<string> id, InstanceState? state = null, CustomResourceOptions? options = null)
            : base("cloudamqp:index/instance:Instance", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                AdditionalSecretOutputs =
                {
                    "apikey",
                    "url",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Instance resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Instance Get(string name, Input<string> id, InstanceState? state = null, CustomResourceOptions? options = null)
        {
            return new Instance(name, id, state, options);
        }
    }

    public sealed class InstanceArgs : global::Pulumi.ResourceArgs
    {
        [Input("copySettings")]
        private InputList<Inputs.InstanceCopySettingArgs>? _copySettings;

        /// <summary>
        /// Copy settings from one CloudAMQP instance to a new. Consists of the block documented below.
        /// 
        /// ___
        /// 
        /// The `copy_settings` block consists of:
        /// </summary>
        public InputList<Inputs.InstanceCopySettingArgs> CopySettings
        {
            get => _copySettings ?? (_copySettings = new InputList<Inputs.InstanceCopySettingArgs>());
            set => _copySettings = value;
        }

        /// <summary>
        /// Keep associated VPC when deleting instance, default set to false.
        /// </summary>
        [Input("keepAssociatedVpc")]
        public Input<bool>? KeepAssociatedVpc { get; set; }

        /// <summary>
        /// Name of the CloudAMQP instance.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Set to true to discard creating default alarms when the instance is created. Can be left out, will then use default value = false.
        /// </summary>
        [Input("noDefaultAlarms")]
        public Input<bool>? NoDefaultAlarms { get; set; }

        /// <summary>
        /// Number of nodes, 1, 3 or 5 depending on plan used. Only needed for legacy plans, will otherwise be computed.
        /// 
        /// ***Deprecated: Legacy subscriptions plan can still change this to scale up or down the instance. New subscriptions plans use the plan to determine number of nodes. In order to change number of nodes the `plan` needs to be updated.***
        /// </summary>
        [Input("nodes")]
        public Input<int>? Nodes { get; set; }

        /// <summary>
        /// The subscription plan. See available plans
        /// </summary>
        [Input("plan", required: true)]
        public Input<string> Plan { get; set; } = null!;

        /// <summary>
        /// The region to host the instance in. See instance regions
        /// 
        /// ***Note: Changing region will force the instance to be destroyed and a new created in the new region. All data will be lost and a new name assigned.***
        /// </summary>
        [Input("region", required: true)]
        public Input<string> Region { get; set; } = null!;

        /// <summary>
        /// The Rabbit MQ version. Can be left out, will then be set to default value used by CloudAMQP API.
        /// 
        /// ***Note: There is not yet any support in the provider to change the RMQ version. Once it's set in the initial creation, it will remain.***
        /// </summary>
        [Input("rmqVersion")]
        public Input<string>? RmqVersion { get; set; }

        [Input("tags")]
        private InputList<string>? _tags;

        /// <summary>
        /// One or more tags for the CloudAMQP instance, makes it possible to categories multiple instances in console view. Default there is no tags assigned.
        /// </summary>
        public InputList<string> Tags
        {
            get => _tags ?? (_tags = new InputList<string>());
            set => _tags = value;
        }

        /// <summary>
        /// The VPC ID. Use this to create your instance in an existing VPC. See available example.
        /// </summary>
        [Input("vpcId")]
        public Input<int>? VpcId { get; set; }

        /// <summary>
        /// Creates a dedicated VPC subnet, shouldn't overlap with other VPC subnet, default subnet used 10.56.72.0/24.
        /// 
        /// ***Deprecated: Will be removed in next major version (v2.0)***
        /// 
        /// ***Note: extra fee will be charged when using VPC, see [CloudAMQP](https://cloudamqp.com) for more information.***
        /// </summary>
        [Input("vpcSubnet")]
        public Input<string>? VpcSubnet { get; set; }

        public InstanceArgs()
        {
        }
        public static new InstanceArgs Empty => new InstanceArgs();
    }

    public sealed class InstanceState : global::Pulumi.ResourceArgs
    {
        [Input("apikey")]
        private Input<string>? _apikey;

        /// <summary>
        /// API key needed to communicate to CloudAMQP's second API. The second API is used to manage alarms, integration and more, full description [CloudAMQP API](https://docs.cloudamqp.com/cloudamqp_api.html).
        /// </summary>
        public Input<string>? Apikey
        {
            get => _apikey;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _apikey = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// Information if the CloudAMQP instance runs either RabbitMQ or LavinMQ.
        /// </summary>
        [Input("backend")]
        public Input<string>? Backend { get; set; }

        [Input("copySettings")]
        private InputList<Inputs.InstanceCopySettingGetArgs>? _copySettings;

        /// <summary>
        /// Copy settings from one CloudAMQP instance to a new. Consists of the block documented below.
        /// 
        /// ___
        /// 
        /// The `copy_settings` block consists of:
        /// </summary>
        public InputList<Inputs.InstanceCopySettingGetArgs> CopySettings
        {
            get => _copySettings ?? (_copySettings = new InputList<Inputs.InstanceCopySettingGetArgs>());
            set => _copySettings = value;
        }

        /// <summary>
        /// Information if the CloudAMQP instance is shared or dedicated.
        /// </summary>
        [Input("dedicated")]
        public Input<bool>? Dedicated { get; set; }

        /// <summary>
        /// The external hostname for the CloudAMQP instance.
        /// </summary>
        [Input("host")]
        public Input<string>? Host { get; set; }

        /// <summary>
        /// The internal hostname for the CloudAMQP instance.
        /// </summary>
        [Input("hostInternal")]
        public Input<string>? HostInternal { get; set; }

        /// <summary>
        /// Keep associated VPC when deleting instance, default set to false.
        /// </summary>
        [Input("keepAssociatedVpc")]
        public Input<bool>? KeepAssociatedVpc { get; set; }

        /// <summary>
        /// Name of the CloudAMQP instance.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Set to true to discard creating default alarms when the instance is created. Can be left out, will then use default value = false.
        /// </summary>
        [Input("noDefaultAlarms")]
        public Input<bool>? NoDefaultAlarms { get; set; }

        /// <summary>
        /// Number of nodes, 1, 3 or 5 depending on plan used. Only needed for legacy plans, will otherwise be computed.
        /// 
        /// ***Deprecated: Legacy subscriptions plan can still change this to scale up or down the instance. New subscriptions plans use the plan to determine number of nodes. In order to change number of nodes the `plan` needs to be updated.***
        /// </summary>
        [Input("nodes")]
        public Input<int>? Nodes { get; set; }

        /// <summary>
        /// The subscription plan. See available plans
        /// </summary>
        [Input("plan")]
        public Input<string>? Plan { get; set; }

        /// <summary>
        /// Flag describing if the resource is ready
        /// </summary>
        [Input("ready")]
        public Input<bool>? Ready { get; set; }

        /// <summary>
        /// The region to host the instance in. See instance regions
        /// 
        /// ***Note: Changing region will force the instance to be destroyed and a new created in the new region. All data will be lost and a new name assigned.***
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// The Rabbit MQ version. Can be left out, will then be set to default value used by CloudAMQP API.
        /// 
        /// ***Note: There is not yet any support in the provider to change the RMQ version. Once it's set in the initial creation, it will remain.***
        /// </summary>
        [Input("rmqVersion")]
        public Input<string>? RmqVersion { get; set; }

        [Input("tags")]
        private InputList<string>? _tags;

        /// <summary>
        /// One or more tags for the CloudAMQP instance, makes it possible to categories multiple instances in console view. Default there is no tags assigned.
        /// </summary>
        public InputList<string> Tags
        {
            get => _tags ?? (_tags = new InputList<string>());
            set => _tags = value;
        }

        [Input("url")]
        private Input<string>? _url;

        /// <summary>
        /// The AMQP URL (uses the internal hostname if the instance was created with VPC). Has the format: `amqps://{username}:{password}@{hostname}/{vhost}`
        /// </summary>
        public Input<string>? Url
        {
            get => _url;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _url = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// The virtual host used by Rabbit MQ.
        /// </summary>
        [Input("vhost")]
        public Input<string>? Vhost { get; set; }

        /// <summary>
        /// The VPC ID. Use this to create your instance in an existing VPC. See available example.
        /// </summary>
        [Input("vpcId")]
        public Input<int>? VpcId { get; set; }

        /// <summary>
        /// Creates a dedicated VPC subnet, shouldn't overlap with other VPC subnet, default subnet used 10.56.72.0/24.
        /// 
        /// ***Deprecated: Will be removed in next major version (v2.0)***
        /// 
        /// ***Note: extra fee will be charged when using VPC, see [CloudAMQP](https://cloudamqp.com) for more information.***
        /// </summary>
        [Input("vpcSubnet")]
        public Input<string>? VpcSubnet { get; set; }

        public InstanceState()
        {
        }
        public static new InstanceState Empty => new InstanceState();
    }
}
