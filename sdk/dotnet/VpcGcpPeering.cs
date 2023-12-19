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
    /// This resouce creates a VPC peering configuration for the CloudAMQP instance. The configuration will connect to another VPC network hosted on Google Cloud Platform (GCP). See the [GCP documentation](https://cloud.google.com/vpc/docs/using-vpc-peering) for more information on how to create the VPC peering configuration.
    /// 
    /// &gt; **Note:** Creating a VPC peering will automatically add firewall rules for the peered subnet.
    /// 
    /// &lt;details&gt;
    ///  &lt;summary&gt;
    ///     &lt;i&gt;Default VPC peering firewall rule&lt;/i&gt;
    ///   &lt;/summary&gt;
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    /// });
    /// ```
    /// 
    /// &lt;/details&gt;
    /// 
    /// Pricing is available at [cloudamqp.com](https://www.cloudamqp.com/plans.html).
    /// 
    /// Only available for dedicated subscription plans.
    /// 
    /// ## Example Usage
    /// ### With Additional Firewall Rules
    /// 
    /// &lt;details&gt;
    ///   &lt;summary&gt;
    ///     &lt;b&gt;
    ///       &lt;i&gt;VPC peering pre v1.16.0&lt;/i&gt;
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
    ///     // VPC peering configuration
    ///     var vpcPeeringRequest = new CloudAmqp.VpcGcpPeering("vpcPeeringRequest", new()
    ///     {
    ///         InstanceId = cloudamqp_instance.Instance.Id,
    ///         PeerNetworkUri = @var.Peer_network_uri,
    ///     });
    /// 
    ///     // Firewall rules
    ///     var firewallSettings = new CloudAmqp.SecurityFirewall("firewallSettings", new()
    ///     {
    ///         InstanceId = cloudamqp_instance.Instance.Id,
    ///         Rules = new[]
    ///         {
    ///             new CloudAmqp.Inputs.SecurityFirewallRuleArgs
    ///             {
    ///                 Ip = @var.Peer_subnet,
    ///                 Ports = new[]
    ///                 {
    ///                     15672,
    ///                 },
    ///                 Services = new[]
    ///                 {
    ///                     "AMQP",
    ///                     "AMQPS",
    ///                     "STREAM",
    ///                     "STREAM_SSL",
    ///                 },
    ///                 Description = "VPC peering for &lt;NETWORK&gt;",
    ///             },
    ///             new CloudAmqp.Inputs.SecurityFirewallRuleArgs
    ///             {
    ///                 Ip = "192.168.0.0/24",
    ///                 Ports = new[]
    ///                 {
    ///                     4567,
    ///                     4568,
    ///                 },
    ///                 Services = new[]
    ///                 {
    ///                     "AMQP",
    ///                     "AMQPS",
    ///                     "HTTPS",
    ///                 },
    ///             },
    ///         },
    ///     }, new CustomResourceOptions
    ///     {
    ///         DependsOn = new[]
    ///         {
    ///             vpcPeeringRequest,
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
    ///       &lt;i&gt;VPC peering post v1.16.0 (Managed VPC)&lt;/i&gt;
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
    ///     // VPC peering configuration
    ///     var vpcPeeringRequest = new CloudAmqp.VpcGcpPeering("vpcPeeringRequest", new()
    ///     {
    ///         VpcId = cloudamqp_vpc.Vpc.Id,
    ///         PeerNetworkUri = @var.Peer_network_uri,
    ///     });
    /// 
    ///     // Firewall rules
    ///     var firewallSettings = new CloudAmqp.SecurityFirewall("firewallSettings", new()
    ///     {
    ///         InstanceId = cloudamqp_instance.Instance.Id,
    ///         Rules = new[]
    ///         {
    ///             new CloudAmqp.Inputs.SecurityFirewallRuleArgs
    ///             {
    ///                 Ip = @var.Peer_subnet,
    ///                 Ports = new[]
    ///                 {
    ///                     15672,
    ///                 },
    ///                 Services = new[]
    ///                 {
    ///                     "AMQP",
    ///                     "AMQPS",
    ///                     "STREAM",
    ///                     "STREAM_SSL",
    ///                 },
    ///                 Description = "VPC peering for &lt;NETWORK&gt;",
    ///             },
    ///             new CloudAmqp.Inputs.SecurityFirewallRuleArgs
    ///             {
    ///                 Ip = "0.0.0.0/0",
    ///                 Ports = new() { },
    ///                 Services = new[]
    ///                 {
    ///                     "HTTPS",
    ///                 },
    ///                 Description = "MGMT interface",
    ///             },
    ///         },
    ///     }, new CustomResourceOptions
    ///     {
    ///         DependsOn = new[]
    ///         {
    ///             vpcPeeringRequest,
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// &lt;/details&gt;
    /// ## Depedency
    /// 
    /// *Pre v1.16.0*
    /// This resource depends on CloudAMQP instance identifier, `cloudamqp_instance.instance.id`.
    /// 
    /// *Post v1.16.0*
    /// This resource depends on CloudAMQP managed VPC identifier, `cloudamqp_vpc.vpc.id` or instance identifier, `cloudamqp_instance.instance.id`.
    /// 
    /// ## Create VPC Peering with additional firewall rules
    /// 
    /// To create a VPC peering configuration with additional firewall rules, it's required to chain the cloudamqp.SecurityFirewall
    /// resource to avoid parallel conflicting resource calls. This is done by adding dependency from the firewall resource to the VPC peering resource.
    /// 
    /// Furthermore, since all firewall rules are overwritten, the otherwise automatically added rules for the VPC peering also needs to be added.
    /// 
    /// See example below.
    /// 
    /// ## Import
    /// 
    /// Not possible to import this resource.
    /// </summary>
    [CloudAmqpResourceType("cloudamqp:index/vpcGcpPeering:VpcGcpPeering")]
    public partial class VpcGcpPeering : global::Pulumi.CustomResource
    {
        /// <summary>
        /// VPC peering auto created routes
        /// </summary>
        [Output("autoCreateRoutes")]
        public Output<bool> AutoCreateRoutes { get; private set; } = null!;

        /// <summary>
        /// The CloudAMQP instance identifier. *Deprecated from v1.16.0*
        /// </summary>
        [Output("instanceId")]
        public Output<int?> InstanceId { get; private set; } = null!;

        /// <summary>
        /// Network uri of the VPC network to which you will peer with.
        /// </summary>
        [Output("peerNetworkUri")]
        public Output<string> PeerNetworkUri { get; private set; } = null!;

        /// <summary>
        /// Configurable sleep time (seconds) between retries when requesting or reading
        /// peering. Default set to 10 seconds. *Available from v1.29.0*
        /// </summary>
        [Output("sleep")]
        public Output<int?> Sleep { get; private set; } = null!;

        /// <summary>
        /// VPC peering state
        /// </summary>
        [Output("state")]
        public Output<string> State { get; private set; } = null!;

        /// <summary>
        /// VPC peering state details
        /// </summary>
        [Output("stateDetails")]
        public Output<string> StateDetails { get; private set; } = null!;

        /// <summary>
        /// Configurable timeout time (seconds) before retries times out. Default set
        /// to 1800 seconds. *Available from v1.29.0*
        /// </summary>
        [Output("timeout")]
        public Output<int?> Timeout { get; private set; } = null!;

        /// <summary>
        /// The managed VPC identifier. *Available from v1.16.0*
        /// </summary>
        [Output("vpcId")]
        public Output<string?> VpcId { get; private set; } = null!;

        /// <summary>
        /// Makes the resource wait until the peering is connected.
        /// Default set to false. *Available from v1.28.0*
        /// </summary>
        [Output("waitOnPeeringStatus")]
        public Output<bool?> WaitOnPeeringStatus { get; private set; } = null!;


        /// <summary>
        /// Create a VpcGcpPeering resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public VpcGcpPeering(string name, VpcGcpPeeringArgs args, CustomResourceOptions? options = null)
            : base("cloudamqp:index/vpcGcpPeering:VpcGcpPeering", name, args ?? new VpcGcpPeeringArgs(), MakeResourceOptions(options, ""))
        {
        }

        private VpcGcpPeering(string name, Input<string> id, VpcGcpPeeringState? state = null, CustomResourceOptions? options = null)
            : base("cloudamqp:index/vpcGcpPeering:VpcGcpPeering", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing VpcGcpPeering resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static VpcGcpPeering Get(string name, Input<string> id, VpcGcpPeeringState? state = null, CustomResourceOptions? options = null)
        {
            return new VpcGcpPeering(name, id, state, options);
        }
    }

    public sealed class VpcGcpPeeringArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The CloudAMQP instance identifier. *Deprecated from v1.16.0*
        /// </summary>
        [Input("instanceId")]
        public Input<int>? InstanceId { get; set; }

        /// <summary>
        /// Network uri of the VPC network to which you will peer with.
        /// </summary>
        [Input("peerNetworkUri", required: true)]
        public Input<string> PeerNetworkUri { get; set; } = null!;

        /// <summary>
        /// Configurable sleep time (seconds) between retries when requesting or reading
        /// peering. Default set to 10 seconds. *Available from v1.29.0*
        /// </summary>
        [Input("sleep")]
        public Input<int>? Sleep { get; set; }

        /// <summary>
        /// Configurable timeout time (seconds) before retries times out. Default set
        /// to 1800 seconds. *Available from v1.29.0*
        /// </summary>
        [Input("timeout")]
        public Input<int>? Timeout { get; set; }

        /// <summary>
        /// The managed VPC identifier. *Available from v1.16.0*
        /// </summary>
        [Input("vpcId")]
        public Input<string>? VpcId { get; set; }

        /// <summary>
        /// Makes the resource wait until the peering is connected.
        /// Default set to false. *Available from v1.28.0*
        /// </summary>
        [Input("waitOnPeeringStatus")]
        public Input<bool>? WaitOnPeeringStatus { get; set; }

        public VpcGcpPeeringArgs()
        {
        }
        public static new VpcGcpPeeringArgs Empty => new VpcGcpPeeringArgs();
    }

    public sealed class VpcGcpPeeringState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// VPC peering auto created routes
        /// </summary>
        [Input("autoCreateRoutes")]
        public Input<bool>? AutoCreateRoutes { get; set; }

        /// <summary>
        /// The CloudAMQP instance identifier. *Deprecated from v1.16.0*
        /// </summary>
        [Input("instanceId")]
        public Input<int>? InstanceId { get; set; }

        /// <summary>
        /// Network uri of the VPC network to which you will peer with.
        /// </summary>
        [Input("peerNetworkUri")]
        public Input<string>? PeerNetworkUri { get; set; }

        /// <summary>
        /// Configurable sleep time (seconds) between retries when requesting or reading
        /// peering. Default set to 10 seconds. *Available from v1.29.0*
        /// </summary>
        [Input("sleep")]
        public Input<int>? Sleep { get; set; }

        /// <summary>
        /// VPC peering state
        /// </summary>
        [Input("state")]
        public Input<string>? State { get; set; }

        /// <summary>
        /// VPC peering state details
        /// </summary>
        [Input("stateDetails")]
        public Input<string>? StateDetails { get; set; }

        /// <summary>
        /// Configurable timeout time (seconds) before retries times out. Default set
        /// to 1800 seconds. *Available from v1.29.0*
        /// </summary>
        [Input("timeout")]
        public Input<int>? Timeout { get; set; }

        /// <summary>
        /// The managed VPC identifier. *Available from v1.16.0*
        /// </summary>
        [Input("vpcId")]
        public Input<string>? VpcId { get; set; }

        /// <summary>
        /// Makes the resource wait until the peering is connected.
        /// Default set to false. *Available from v1.28.0*
        /// </summary>
        [Input("waitOnPeeringStatus")]
        public Input<bool>? WaitOnPeeringStatus { get; set; }

        public VpcGcpPeeringState()
        {
        }
        public static new VpcGcpPeeringState Empty => new VpcGcpPeeringState();
    }
}
