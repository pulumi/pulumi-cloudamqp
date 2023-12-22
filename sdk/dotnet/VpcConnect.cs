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
    /// This resource is a generic way to handle PrivateLink (AWS and Azure) and Private Service Connect (GCP).
    /// Communication between resources can be done just as they were living inside a VPC. CloudAMQP creates an Endpoint
    /// Service to connect the VPC and creating a new network interface to handle the communicate.
    /// 
    /// If no existing VPC available when enable VPC connect, a new VPC will be created with subnet `10.52.72.0/24`.
    /// 
    /// More information can be found at: [CloudAMQP VPC Connect](https://www.cloudamqp.com/docs/cloudamqp-vpc-connect.html)
    /// 
    /// &gt; **Note:** Enabling VPC Connect will automatically add a firewall rule.
    /// 
    /// &lt;details&gt;
    ///  &lt;summary&gt;
    ///     &lt;b&gt;
    ///       &lt;i&gt;Default PrivateLink firewall rule [AWS, Azure]&lt;/i&gt;
    ///     &lt;/b&gt;
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
    /// &lt;details&gt;
    ///  &lt;summary&gt;
    ///     &lt;b&gt;
    ///       &lt;i&gt;Default Private Service Connect firewall rule [GCP]&lt;/i&gt;
    ///     &lt;/b&gt;
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
    /// Only available for dedicated subscription plans.
    /// 
    /// ## Example Usage
    /// 
    /// &lt;details&gt;
    ///   &lt;summary&gt;
    ///     &lt;b&gt;
    ///       &lt;i&gt;Enable VPC Connect (PrivateLink) in AWS&lt;/i&gt;
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
    ///         Region = "amazon-web-services::us-west-1",
    ///         Subnet = "10.56.72.0/24",
    ///         Tags = new[] {},
    ///     });
    /// 
    ///     var instance = new CloudAmqp.Instance("instance", new()
    ///     {
    ///         Plan = "bunny-1",
    ///         Region = "amazon-web-services::us-west-1",
    ///         Tags = new[] {},
    ///         VpcId = vpc.Id,
    ///         KeepAssociatedVpc = true,
    ///     });
    /// 
    ///     var vpcConnect = new CloudAmqp.VpcConnect("vpcConnect", new()
    ///     {
    ///         InstanceId = instance.Id,
    ///         Region = instance.Region,
    ///         AllowedPrincipals = new[]
    ///         {
    ///             "arn:aws:iam::aws-account-id:user/user-name",
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
    ///       &lt;i&gt;Enable VPC Connect (PrivateLink) in Azure&lt;/i&gt;
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
    ///     var vpcConnect = new CloudAmqp.VpcConnect("vpcConnect", new()
    ///     {
    ///         InstanceId = instance.Id,
    ///         Region = instance.Region,
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
    /// 
    /// &lt;details&gt;
    ///   &lt;summary&gt;
    ///     &lt;b&gt;
    ///       &lt;i&gt;Enable VPC Connect (Private Service Connect) in GCP&lt;/i&gt;
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
    ///         Region = "google-compute-engine::us-west1",
    ///         Subnet = "10.56.72.0/24",
    ///         Tags = new[] {},
    ///     });
    /// 
    ///     var instance = new CloudAmqp.Instance("instance", new()
    ///     {
    ///         Plan = "bunny-1",
    ///         Region = "google-compute-engine::us-west1",
    ///         Tags = new[] {},
    ///         VpcId = vpc.Id,
    ///         KeepAssociatedVpc = true,
    ///     });
    /// 
    ///     var vpcConnect = new CloudAmqp.VpcConnect("vpcConnect", new()
    ///     {
    ///         InstanceId = instance.Id,
    ///         Region = instance.Region,
    ///         AllowedProjects = new[]
    ///         {
    ///             "some-project-123456",
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
    ///         Region = "amazon-web-services::us-west-1",
    ///         Subnet = "10.56.72.0/24",
    ///         Tags = new[] {},
    ///     });
    /// 
    ///     var instance = new CloudAmqp.Instance("instance", new()
    ///     {
    ///         Plan = "bunny-1",
    ///         Region = "amazon-web-services::us-west-1",
    ///         Tags = new[] {},
    ///         VpcId = vpc.Id,
    ///         KeepAssociatedVpc = true,
    ///     });
    /// 
    ///     var vpcConnect = new CloudAmqp.VpcConnect("vpcConnect", new()
    ///     {
    ///         InstanceId = instance.Id,
    ///         AllowedPrincipals = new[]
    ///         {
    ///             "arn:aws:iam::aws-account-id:user/user-name",
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
    ///             vpcConnect,
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
    /// Since `region` also is required, suggest to reuse the argument from CloudAMQP instance,
    /// `cloudamqp_instance.instance.region`.
    /// 
    /// ## Create VPC Connect with additional firewall rules
    /// 
    /// To create a PrivateLink/Private Service Connect configuration with additional firewall rules, it's required to chain the cloudamqp.SecurityFirewall
    /// resource to avoid parallel conflicting resource calls. You can do this by making the firewall
    /// resource depend on the VPC Connect resource, `cloudamqp_vpc_connect.vpc_connect`.
    /// 
    /// Furthermore, since all firewall rules are overwritten, the otherwise automatically added rules for
    /// the VPC Connect also needs to be added.
    /// 
    /// ## Import
    /// 
    /// `cloudamqp_vpc_connect` can be imported using CloudAMQP internal identifier.
    /// 
    /// ```sh
    ///  $ pulumi import cloudamqp:index/vpcConnect:VpcConnect vpc_connect &lt;id&gt;`
    /// ```
    /// 
    ///  The resource uses the same identifier as the CloudAMQP instance. To retrieve the identifier for an instance, either use [CloudAMQP customer API](https://docs.cloudamqp.com/#list-instances) or use the data source [`cloudamqp_account`](./data-sources/account.md).
    /// </summary>
    [CloudAmqpResourceType("cloudamqp:index/vpcConnect:VpcConnect")]
    public partial class VpcConnect : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Covering availability zones used when creating an endpoint from other VPC. (AWS)
        /// </summary>
        [Output("activeZones")]
        public Output<ImmutableArray<string>> ActiveZones { get; private set; } = null!;

        /// <summary>
        /// List of allowed prinicpals used by AWS, see below table.
        /// </summary>
        [Output("allowedPrincipals")]
        public Output<ImmutableArray<string>> AllowedPrincipals { get; private set; } = null!;

        /// <summary>
        /// List of allowed projects used by GCP, see below table.
        /// </summary>
        [Output("allowedProjects")]
        public Output<ImmutableArray<string>> AllowedProjects { get; private set; } = null!;

        /// <summary>
        /// List of approved subscriptions used by Azure, see below table.
        /// </summary>
        [Output("approvedSubscriptions")]
        public Output<ImmutableArray<string>> ApprovedSubscriptions { get; private set; } = null!;

        /// <summary>
        /// The CloudAMQP instance identifier.
        /// </summary>
        [Output("instanceId")]
        public Output<int> InstanceId { get; private set; } = null!;

        /// <summary>
        /// The region where the CloudAMQP instance is hosted.
        /// </summary>
        [Output("region")]
        public Output<string> Region { get; private set; } = null!;

        /// <summary>
        /// Service name (alias for Azure) of the PrivateLink.
        /// </summary>
        [Output("serviceName")]
        public Output<string> ServiceName { get; private set; } = null!;

        /// <summary>
        /// Configurable sleep time (seconds) when enable Private Service Connect.
        /// Default set to 10 seconds.
        /// </summary>
        [Output("sleep")]
        public Output<int?> Sleep { get; private set; } = null!;

        /// <summary>
        /// Private Service Connect status [enable, pending, disable]
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// Configurable timeout time (seconds) when enable Private Service Connect.
        /// Default set to 1800 seconds.
        /// 
        /// ___
        /// 
        /// The `allowed_principals`, `approved_subscriptions` or `allowed_projects` data depends on the provider platform:
        /// 
        /// | Platform | Description         | Format                                                                                                                             |
        /// |----------|---------------------|------------------------------------------------------------------------------------------------------------------------------------|
        /// | AWS      | IAM ARN principals  | arn:aws:iam::aws-account-id:root&lt;br /&gt; arn:aws:iam::aws-account-id:user/user-name&lt;br /&gt; arn:aws:iam::aws-account-id:role/role-name |
        /// | Azure    | Subscription (GUID) | XXXXXXXX-XXXX-XXXX-XXXX-XXXXXXXXXXXX                                                                                               |
        /// | GCP      | Project IDs*        | 6 to 30 lowercase letters, digits, or hyphens                                                                                      |
        /// 
        /// *https://cloud.google.com/resource-manager/reference/rest/v1/projects
        /// </summary>
        [Output("timeout")]
        public Output<int?> Timeout { get; private set; } = null!;


        /// <summary>
        /// Create a VpcConnect resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public VpcConnect(string name, VpcConnectArgs args, CustomResourceOptions? options = null)
            : base("cloudamqp:index/vpcConnect:VpcConnect", name, args ?? new VpcConnectArgs(), MakeResourceOptions(options, ""))
        {
        }

        private VpcConnect(string name, Input<string> id, VpcConnectState? state = null, CustomResourceOptions? options = null)
            : base("cloudamqp:index/vpcConnect:VpcConnect", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing VpcConnect resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static VpcConnect Get(string name, Input<string> id, VpcConnectState? state = null, CustomResourceOptions? options = null)
        {
            return new VpcConnect(name, id, state, options);
        }
    }

    public sealed class VpcConnectArgs : global::Pulumi.ResourceArgs
    {
        [Input("allowedPrincipals")]
        private InputList<string>? _allowedPrincipals;

        /// <summary>
        /// List of allowed prinicpals used by AWS, see below table.
        /// </summary>
        public InputList<string> AllowedPrincipals
        {
            get => _allowedPrincipals ?? (_allowedPrincipals = new InputList<string>());
            set => _allowedPrincipals = value;
        }

        [Input("allowedProjects")]
        private InputList<string>? _allowedProjects;

        /// <summary>
        /// List of allowed projects used by GCP, see below table.
        /// </summary>
        public InputList<string> AllowedProjects
        {
            get => _allowedProjects ?? (_allowedProjects = new InputList<string>());
            set => _allowedProjects = value;
        }

        [Input("approvedSubscriptions")]
        private InputList<string>? _approvedSubscriptions;

        /// <summary>
        /// List of approved subscriptions used by Azure, see below table.
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
        /// The region where the CloudAMQP instance is hosted.
        /// </summary>
        [Input("region", required: true)]
        public Input<string> Region { get; set; } = null!;

        /// <summary>
        /// Configurable sleep time (seconds) when enable Private Service Connect.
        /// Default set to 10 seconds.
        /// </summary>
        [Input("sleep")]
        public Input<int>? Sleep { get; set; }

        /// <summary>
        /// Configurable timeout time (seconds) when enable Private Service Connect.
        /// Default set to 1800 seconds.
        /// 
        /// ___
        /// 
        /// The `allowed_principals`, `approved_subscriptions` or `allowed_projects` data depends on the provider platform:
        /// 
        /// | Platform | Description         | Format                                                                                                                             |
        /// |----------|---------------------|------------------------------------------------------------------------------------------------------------------------------------|
        /// | AWS      | IAM ARN principals  | arn:aws:iam::aws-account-id:root&lt;br /&gt; arn:aws:iam::aws-account-id:user/user-name&lt;br /&gt; arn:aws:iam::aws-account-id:role/role-name |
        /// | Azure    | Subscription (GUID) | XXXXXXXX-XXXX-XXXX-XXXX-XXXXXXXXXXXX                                                                                               |
        /// | GCP      | Project IDs*        | 6 to 30 lowercase letters, digits, or hyphens                                                                                      |
        /// 
        /// *https://cloud.google.com/resource-manager/reference/rest/v1/projects
        /// </summary>
        [Input("timeout")]
        public Input<int>? Timeout { get; set; }

        public VpcConnectArgs()
        {
        }
        public static new VpcConnectArgs Empty => new VpcConnectArgs();
    }

    public sealed class VpcConnectState : global::Pulumi.ResourceArgs
    {
        [Input("activeZones")]
        private InputList<string>? _activeZones;

        /// <summary>
        /// Covering availability zones used when creating an endpoint from other VPC. (AWS)
        /// </summary>
        public InputList<string> ActiveZones
        {
            get => _activeZones ?? (_activeZones = new InputList<string>());
            set => _activeZones = value;
        }

        [Input("allowedPrincipals")]
        private InputList<string>? _allowedPrincipals;

        /// <summary>
        /// List of allowed prinicpals used by AWS, see below table.
        /// </summary>
        public InputList<string> AllowedPrincipals
        {
            get => _allowedPrincipals ?? (_allowedPrincipals = new InputList<string>());
            set => _allowedPrincipals = value;
        }

        [Input("allowedProjects")]
        private InputList<string>? _allowedProjects;

        /// <summary>
        /// List of allowed projects used by GCP, see below table.
        /// </summary>
        public InputList<string> AllowedProjects
        {
            get => _allowedProjects ?? (_allowedProjects = new InputList<string>());
            set => _allowedProjects = value;
        }

        [Input("approvedSubscriptions")]
        private InputList<string>? _approvedSubscriptions;

        /// <summary>
        /// List of approved subscriptions used by Azure, see below table.
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
        /// The region where the CloudAMQP instance is hosted.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// Service name (alias for Azure) of the PrivateLink.
        /// </summary>
        [Input("serviceName")]
        public Input<string>? ServiceName { get; set; }

        /// <summary>
        /// Configurable sleep time (seconds) when enable Private Service Connect.
        /// Default set to 10 seconds.
        /// </summary>
        [Input("sleep")]
        public Input<int>? Sleep { get; set; }

        /// <summary>
        /// Private Service Connect status [enable, pending, disable]
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// Configurable timeout time (seconds) when enable Private Service Connect.
        /// Default set to 1800 seconds.
        /// 
        /// ___
        /// 
        /// The `allowed_principals`, `approved_subscriptions` or `allowed_projects` data depends on the provider platform:
        /// 
        /// | Platform | Description         | Format                                                                                                                             |
        /// |----------|---------------------|------------------------------------------------------------------------------------------------------------------------------------|
        /// | AWS      | IAM ARN principals  | arn:aws:iam::aws-account-id:root&lt;br /&gt; arn:aws:iam::aws-account-id:user/user-name&lt;br /&gt; arn:aws:iam::aws-account-id:role/role-name |
        /// | Azure    | Subscription (GUID) | XXXXXXXX-XXXX-XXXX-XXXX-XXXXXXXXXXXX                                                                                               |
        /// | GCP      | Project IDs*        | 6 to 30 lowercase letters, digits, or hyphens                                                                                      |
        /// 
        /// *https://cloud.google.com/resource-manager/reference/rest/v1/projects
        /// </summary>
        [Input("timeout")]
        public Input<int>? Timeout { get; set; }

        public VpcConnectState()
        {
        }
        public static new VpcConnectState Empty => new VpcConnectState();
    }
}
