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
    /// This resource allows you to create and manage a CloudAMQP instance running Rabbit MQ and deploy to multiple cloud platforms provider and over multiple regions, see Instance regions for more information.
    /// 
    /// Once the instance is created it will be assigned a unique identifier. All other resource and data sources created for this instance needs to reference the instance identifier.
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
    ///         // Minimum free lemur instance
    ///         var lemurInstance = new CloudAmqp.Instance("lemurInstance", new CloudAmqp.InstanceArgs
    ///         {
    ///             Plan = "lemur",
    ///             Region = "amazon-web-services::us-west-1",
    ///         });
    ///         // New dedicated bunny instance
    ///         var instance = new CloudAmqp.Instance("instance", new CloudAmqp.InstanceArgs
    ///         {
    ///             NoDefaultAlarms = true,
    ///             Nodes = 1,
    ///             Plan = "bunny",
    ///             Region = "amazon-web-services::us-west-1",
    ///             RmqVersion = "3.8.3",
    ///             Tags = 
    ///             {
    ///                 "terraform",
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// `cloudamqp_instance`can be imported using CloudAMQP internal identifier. To retrieve the identifier for an instance, use [CloudAMQP customer API](https://docs.cloudamqp.com/#list-instances).
    /// 
    /// ```sh
    ///  $ pulumi import cloudamqp:index/instance:Instance instance &lt;instance_id&gt;`
    /// ```
    /// </summary>
    [CloudAmqpResourceType("cloudamqp:index/instance:Instance")]
    public partial class Instance : Pulumi.CustomResource
    {
        /// <summary>
        /// (Computed) API key needed to communicate to CloudAMQP's second API. The second API is used to manage alarms, integration and more, full description [CloudAMQP API](https://docs.cloudamqp.com/cloudamqp_api.html).
        /// </summary>
        [Output("apikey")]
        public Output<string> Apikey { get; private set; } = null!;

        /// <summary>
        /// Is the instance hosted on a dedicated server
        /// </summary>
        [Output("dedicated")]
        public Output<bool> Dedicated { get; private set; } = null!;

        /// <summary>
        /// (Computed) The host name for the CloudAMQP instance.
        /// </summary>
        [Output("host")]
        public Output<string> Host { get; private set; } = null!;

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
        /// Number of nodes, 1 to 3, in the CloudAMQP instance, default set to 1. The plan chosen must support the number of nodes.
        /// </summary>
        [Output("nodes")]
        public Output<int?> Nodes { get; private set; } = null!;

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
        /// The region to host the instance in. See Instance regions
        /// </summary>
        [Output("region")]
        public Output<string> Region { get; private set; } = null!;

        /// <summary>
        /// The Rabbit MQ version. Can be left out, will then be set to default value used by CloudAMQP API. **Note: There is not yet any support in the provider to change the RMQ version. Once it's set in the initial creation, it will remain.**
        /// </summary>
        [Output("rmqVersion")]
        public Output<string> RmqVersion { get; private set; } = null!;

        /// <summary>
        /// One or more tags for the CloudAMQP instance, makes it possible to categories multiple instances in console view. Default there is no tags assigned.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableArray<string>> Tags { get; private set; } = null!;

        /// <summary>
        /// (Computed) AMQP server endpoint. `amqps://{username}:{password}@{hostname}/{vhost}`
        /// </summary>
        [Output("url")]
        public Output<string> Url { get; private set; } = null!;

        /// <summary>
        /// (Computed) The virtual host used by Rabbit MQ.
        /// </summary>
        [Output("vhost")]
        public Output<string> Vhost { get; private set; } = null!;

        /// <summary>
        /// Creates a dedicated VPC subnet, shouldn't overlap with other VPC subnet, default subnet used 10.56.72.0/24. **NOTE: extra fee will be charged when using VPC, see [CloudAMQP](https://cloudamqp.com) for more information.**
        /// </summary>
        [Output("vpcSubnet")]
        public Output<string?> VpcSubnet { get; private set; } = null!;


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

    public sealed class InstanceArgs : Pulumi.ResourceArgs
    {
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
        /// Number of nodes, 1 to 3, in the CloudAMQP instance, default set to 1. The plan chosen must support the number of nodes.
        /// </summary>
        [Input("nodes")]
        public Input<int>? Nodes { get; set; }

        /// <summary>
        /// The subscription plan. See available plans
        /// </summary>
        [Input("plan", required: true)]
        public Input<string> Plan { get; set; } = null!;

        /// <summary>
        /// The region to host the instance in. See Instance regions
        /// </summary>
        [Input("region", required: true)]
        public Input<string> Region { get; set; } = null!;

        /// <summary>
        /// The Rabbit MQ version. Can be left out, will then be set to default value used by CloudAMQP API. **Note: There is not yet any support in the provider to change the RMQ version. Once it's set in the initial creation, it will remain.**
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
        /// Creates a dedicated VPC subnet, shouldn't overlap with other VPC subnet, default subnet used 10.56.72.0/24. **NOTE: extra fee will be charged when using VPC, see [CloudAMQP](https://cloudamqp.com) for more information.**
        /// </summary>
        [Input("vpcSubnet")]
        public Input<string>? VpcSubnet { get; set; }

        public InstanceArgs()
        {
        }
    }

    public sealed class InstanceState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// (Computed) API key needed to communicate to CloudAMQP's second API. The second API is used to manage alarms, integration and more, full description [CloudAMQP API](https://docs.cloudamqp.com/cloudamqp_api.html).
        /// </summary>
        [Input("apikey")]
        public Input<string>? Apikey { get; set; }

        /// <summary>
        /// Is the instance hosted on a dedicated server
        /// </summary>
        [Input("dedicated")]
        public Input<bool>? Dedicated { get; set; }

        /// <summary>
        /// (Computed) The host name for the CloudAMQP instance.
        /// </summary>
        [Input("host")]
        public Input<string>? Host { get; set; }

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
        /// Number of nodes, 1 to 3, in the CloudAMQP instance, default set to 1. The plan chosen must support the number of nodes.
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
        /// The region to host the instance in. See Instance regions
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// The Rabbit MQ version. Can be left out, will then be set to default value used by CloudAMQP API. **Note: There is not yet any support in the provider to change the RMQ version. Once it's set in the initial creation, it will remain.**
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
        /// (Computed) AMQP server endpoint. `amqps://{username}:{password}@{hostname}/{vhost}`
        /// </summary>
        [Input("url")]
        public Input<string>? Url { get; set; }

        /// <summary>
        /// (Computed) The virtual host used by Rabbit MQ.
        /// </summary>
        [Input("vhost")]
        public Input<string>? Vhost { get; set; }

        /// <summary>
        /// Creates a dedicated VPC subnet, shouldn't overlap with other VPC subnet, default subnet used 10.56.72.0/24. **NOTE: extra fee will be charged when using VPC, see [CloudAMQP](https://cloudamqp.com) for more information.**
        /// </summary>
        [Input("vpcSubnet")]
        public Input<string>? VpcSubnet { get; set; }

        public InstanceState()
        {
        }
    }
}
