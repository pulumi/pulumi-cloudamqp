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
    /// Enable PrivateLink for a CloudAMQP instance hosted in Azure. If no existing VPC available when
    /// enable PrivateLink, a new VPC will be created with subnet `10.52.72.0/24`.
    /// 
    /// &gt; **Note:** Enabling PrivateLink will automatically add firewall rules for the peered subnet.
    /// 
    /// &lt;details&gt;
    ///  &lt;summary&gt;
    ///     &lt;i&gt;Default PrivateLink firewall rule&lt;/i&gt;
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
    /// Pricing is available at [cloudamqp.com](https://www.cloudamqp.com/plans.html) where you can also
    /// find more information about
    /// [CloudAMQP PrivateLink](https://www.cloudamqp.com/docs/cloudamqp-privatelink.html#azure-privatelink).
    /// 
    /// Only available for dedicated subscription plans.
    /// 
    /// &gt; **Warning:** This resource considered deprecated and will be removed in next major version (v2.0).
    /// Recommended to start using the new resource`cloudamqp.VpcConnect`.
    /// 
    /// ## Example Usage
    /// 
    /// &lt;details&gt;
    ///   &lt;summary&gt;
    ///     &lt;b&gt;
    ///       &lt;i&gt;CloudAMQP instance without existing VPC&lt;/i&gt;
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
    ///         Plan = "bunny-1",
    ///         Region = "azure-arm::westus",
    ///         Tags = new[] {},
    ///     });
    /// 
    ///     var privatelink = new CloudAmqp.PrivatelinkAzure("privatelink", new()
    ///     {
    ///         InstanceId = instance.Id,
    ///         ApprovedSubscriptions = new[]
    ///         {
    ///             "XXXXXXXX-XXXX-XXXX-XXXX-XXXXXXXXXXXX",
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;/details&gt;
    /// 
    /// &lt;details&gt;
    ///   &lt;summary&gt;
    ///     &lt;b&gt;
    ///       &lt;i&gt;CloudAMQP instance in an existing VPC&lt;/i&gt;
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
    ///     var vpc = new CloudAmqp.Vpc("vpc", new()
    ///     {
    ///         Region = "azure-arm::westus",
    ///         Subnet = "10.56.72.0/24",
    ///         Tags = new[] {},
    ///     });
    /// 
    ///     var instance = new CloudAmqp.Instance("instance", new()
    ///     {
    ///         Plan = "bunny-1",
    ///         Region = "azure-arm::westus",
    ///         Tags = new[] {},
    ///         VpcId = vpc.Id,
    ///         KeepAssociatedVpc = true,
    ///     });
    /// 
    ///     var privatelink = new CloudAmqp.PrivatelinkAzure("privatelink", new()
    ///     {
    ///         InstanceId = instance.Id,
    ///         ApprovedSubscriptions = new[]
    ///         {
    ///             "XXXXXXXX-XXXX-XXXX-XXXX-XXXXXXXXXXXX",
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// &lt;/details&gt;
    /// ### With Additional Firewall Rules
    /// 
    /// &lt;details&gt;
    ///   &lt;summary&gt;
    ///     &lt;b&gt;
    ///       &lt;i&gt;CloudAMQP instance in an existing VPC with managed firewall rules&lt;/i&gt;
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
    ///     var vpc = new CloudAmqp.Vpc("vpc", new()
    ///     {
    ///         Region = "azure-arm::westus",
    ///         Subnet = "10.56.72.0/24",
    ///         Tags = new[] {},
    ///     });
    /// 
    ///     var instance = new CloudAmqp.Instance("instance", new()
    ///     {
    ///         Plan = "bunny-1",
    ///         Region = "azure-arm::westus",
    ///         Tags = new[] {},
    ///         VpcId = vpc.Id,
    ///         KeepAssociatedVpc = true,
    ///     });
    /// 
    ///     var privatelink = new CloudAmqp.PrivatelinkAzure("privatelink", new()
    ///     {
    ///         InstanceId = instance.Id,
    ///         ApprovedSubscriptions = new[]
    ///         {
    ///             "XXXXXXXX-XXXX-XXXX-XXXX-XXXXXXXXXXXX",
    ///         },
    ///     });
    /// 
    ///     var firewallSettings = new CloudAmqp.SecurityFirewall("firewallSettings", new()
    ///     {
    ///         InstanceId = instance.Id,
    ///         Rules = new[]
    ///         {
    ///             new CloudAmqp.Inputs.SecurityFirewallRuleArgs
    ///             {
    ///                 Description = "Custom PrivateLink setup",
    ///                 Ip = vpc.Subnet,
    ///                 Ports = new() { },
    ///                 Services = new[]
    ///                 {
    ///                     "AMQP",
    ///                     "AMQPS",
    ///                     "HTTPS",
    ///                     "STREAM",
    ///                     "STREAM_SSL",
    ///                 },
    ///             },
    ///             new CloudAmqp.Inputs.SecurityFirewallRuleArgs
    ///             {
    ///                 Description = "MGMT interface",
    ///                 Ip = "0.0.0.0/0",
    ///                 Ports = new() { },
    ///                 Services = new[]
    ///                 {
    ///                     "HTTPS",
    ///                 },
    ///             },
    ///         },
    ///     }, new CustomResourceOptions
    ///     {
    ///         DependsOn = new[]
    ///         {
    ///             privatelink,
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// &lt;/details&gt;
    /// ## Depedency
    /// 
    /// This resource depends on CloudAMQP instance identifier, `cloudamqp_instance.instance.id`.
    /// 
    /// ## Create PrivateLink with additional firewall rules
    /// 
    /// To create a PrivateLink configuration with additional firewall rules, it's required to chain the cloudamqp.SecurityFirewall
    /// resource to avoid parallel conflicting resource calls. You can do this by making the firewall
    /// resource depend on the PrivateLink resource, `cloudamqp_privatelink_azure.privatelink`.
    /// 
    /// Furthermore, since all firewall rules are overwritten, the otherwise automatically added rules for
    /// the PrivateLink also needs to be added.
    /// 
    /// ## Import
    /// 
    /// `cloudamqp_privatelink_aws` can be imported using CloudAMQP internal identifier.
    /// 
    /// ```sh
    ///  $ pulumi import cloudamqp:index/privatelinkAzure:PrivatelinkAzure privatelink &lt;id&gt;`
    /// ```
    /// 
    ///  The resource uses the same identifier as the CloudAMQP instance. To retrieve the identifier for an instance, either use [CloudAMQP customer API](https://docs.cloudamqp.com/#list-instances) or use the data source [`cloudamqp_account`](./data-sources/account.md).
    /// </summary>
    [CloudAmqpResourceType("cloudamqp:index/privatelinkAzure:PrivatelinkAzure")]
    public partial class PrivatelinkAzure : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Approved subscriptions to access the endpoint service.
        /// See format below.
        /// </summary>
        [Output("approvedSubscriptions")]
        public Output<ImmutableArray<string>> ApprovedSubscriptions { get; private set; } = null!;

        /// <summary>
        /// The CloudAMQP instance identifier.
        /// </summary>
        [Output("instanceId")]
        public Output<int> InstanceId { get; private set; } = null!;

        /// <summary>
        /// Name of the server having the PrivateLink enabled.
        /// </summary>
        [Output("serverName")]
        public Output<string> ServerName { get; private set; } = null!;

        /// <summary>
        /// Service name (alias) of the PrivateLink, needed when creating the endpoint.
        /// </summary>
        [Output("serviceName")]
        public Output<string> ServiceName { get; private set; } = null!;

        /// <summary>
        /// Configurable sleep time (seconds) when enable PrivateLink.
        /// Default set to 10 seconds. *Available from v1.29.0*
        /// </summary>
        [Output("sleep")]
        public Output<int?> Sleep { get; private set; } = null!;

        /// <summary>
        /// PrivateLink status [enable, pending, disable]
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// Configurable timeout time (seconds) when enable PrivateLink.
        /// Default set to 1800 seconds. *Available from v1.29.0*
        /// 
        /// Approved subscriptions format (GUID): &lt;br&gt;
        /// `XXXXXXXX-XXXX-XXXX-XXXX-XXXXXXXXXXXX`
        /// </summary>
        [Output("timeout")]
        public Output<int?> Timeout { get; private set; } = null!;


        /// <summary>
        /// Create a PrivatelinkAzure resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public PrivatelinkAzure(string name, PrivatelinkAzureArgs args, CustomResourceOptions? options = null)
            : base("cloudamqp:index/privatelinkAzure:PrivatelinkAzure", name, args ?? new PrivatelinkAzureArgs(), MakeResourceOptions(options, ""))
        {
        }

        private PrivatelinkAzure(string name, Input<string> id, PrivatelinkAzureState? state = null, CustomResourceOptions? options = null)
            : base("cloudamqp:index/privatelinkAzure:PrivatelinkAzure", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing PrivatelinkAzure resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static PrivatelinkAzure Get(string name, Input<string> id, PrivatelinkAzureState? state = null, CustomResourceOptions? options = null)
        {
            return new PrivatelinkAzure(name, id, state, options);
        }
    }

    public sealed class PrivatelinkAzureArgs : global::Pulumi.ResourceArgs
    {
        [Input("approvedSubscriptions", required: true)]
        private InputList<string>? _approvedSubscriptions;

        /// <summary>
        /// Approved subscriptions to access the endpoint service.
        /// See format below.
        /// </summary>
        public InputList<string> ApprovedSubscriptions
        {
            get => _approvedSubscriptions ?? (_approvedSubscriptions = new InputList<string>());
            set => _approvedSubscriptions = value;
        }

        /// <summary>
        /// The CloudAMQP instance identifier.
        /// </summary>
        [Input("instanceId", required: true)]
        public Input<int> InstanceId { get; set; } = null!;

        /// <summary>
        /// Configurable sleep time (seconds) when enable PrivateLink.
        /// Default set to 10 seconds. *Available from v1.29.0*
        /// </summary>
        [Input("sleep")]
        public Input<int>? Sleep { get; set; }

        /// <summary>
        /// Configurable timeout time (seconds) when enable PrivateLink.
        /// Default set to 1800 seconds. *Available from v1.29.0*
        /// 
        /// Approved subscriptions format (GUID): &lt;br&gt;
        /// `XXXXXXXX-XXXX-XXXX-XXXX-XXXXXXXXXXXX`
        /// </summary>
        [Input("timeout")]
        public Input<int>? Timeout { get; set; }

        public PrivatelinkAzureArgs()
        {
        }
        public static new PrivatelinkAzureArgs Empty => new PrivatelinkAzureArgs();
    }

    public sealed class PrivatelinkAzureState : global::Pulumi.ResourceArgs
    {
        [Input("approvedSubscriptions")]
        private InputList<string>? _approvedSubscriptions;

        /// <summary>
        /// Approved subscriptions to access the endpoint service.
        /// See format below.
        /// </summary>
        public InputList<string> ApprovedSubscriptions
        {
            get => _approvedSubscriptions ?? (_approvedSubscriptions = new InputList<string>());
            set => _approvedSubscriptions = value;
        }

        /// <summary>
        /// The CloudAMQP instance identifier.
        /// </summary>
        [Input("instanceId")]
        public Input<int>? InstanceId { get; set; }

        /// <summary>
        /// Name of the server having the PrivateLink enabled.
        /// </summary>
        [Input("serverName")]
        public Input<string>? ServerName { get; set; }

        /// <summary>
        /// Service name (alias) of the PrivateLink, needed when creating the endpoint.
        /// </summary>
        [Input("serviceName")]
        public Input<string>? ServiceName { get; set; }

        /// <summary>
        /// Configurable sleep time (seconds) when enable PrivateLink.
        /// Default set to 10 seconds. *Available from v1.29.0*
        /// </summary>
        [Input("sleep")]
        public Input<int>? Sleep { get; set; }

        /// <summary>
        /// PrivateLink status [enable, pending, disable]
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// Configurable timeout time (seconds) when enable PrivateLink.
        /// Default set to 1800 seconds. *Available from v1.29.0*
        /// 
        /// Approved subscriptions format (GUID): &lt;br&gt;
        /// `XXXXXXXX-XXXX-XXXX-XXXX-XXXXXXXXXXXX`
        /// </summary>
        [Input("timeout")]
        public Input<int>? Timeout { get; set; }

        public PrivatelinkAzureState()
        {
        }
        public static new PrivatelinkAzureState Empty => new PrivatelinkAzureState();
    }
}
