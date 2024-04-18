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
    /// This resource allows you to configure and manage your custom domain for the CloudAMQP instance.
    /// 
    /// Adding a custom domain to your instance will generate a TLS certificate from [Let's Encrypt], for the given hostname, and install it on all servers in your cluster. The certificate will be automatically renewed going forward.
    /// 
    /// &gt; **WARNING:** Please note that when creating, changing or deleting the custom domain, the listeners on your servers will be restarted in order to apply the changes. This will close your current connections.
    /// 
    /// Your custom domain name needs to point to your CloudAMQP hostname, preferably using a CNAME DNS record.
    /// 
    /// Only available for dedicated subscription plans.
    /// 
    /// ## Example Usage
    /// 
    /// &lt;!--Start PulumiCodeChooser --&gt;
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using CloudAmqp = Pulumi.CloudAmqp;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var settings = new CloudAmqp.CustomDomain("settings", new()
    ///     {
    ///         InstanceId = instance.Id,
    ///         Hostname = "myname.mydomain",
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ## Depedency
    /// 
    /// This resource depends on CloudAMQP instance identifier, `cloudamqp_instance.instance.id`.
    /// 
    /// ## Import
    /// 
    /// `cloudamqp_custom_domain` can be imported using CloudAMQP instance identifier.
    /// 
    /// ```sh
    /// $ pulumi import cloudamqp:index/customDomain:CustomDomain settings &lt;instance_id&gt;`
    /// ```
    /// 
    /// [Let's Encrypt]: https://letsencrypt.org/
    /// </summary>
    [CloudAmqpResourceType("cloudamqp:index/customDomain:CustomDomain")]
    public partial class CustomDomain : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Your custom domain name.
        /// </summary>
        [Output("hostname")]
        public Output<string> Hostname { get; private set; } = null!;

        /// <summary>
        /// The CloudAMQP instance ID.
        /// </summary>
        [Output("instanceId")]
        public Output<int> InstanceId { get; private set; } = null!;


        /// <summary>
        /// Create a CustomDomain resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public CustomDomain(string name, CustomDomainArgs args, CustomResourceOptions? options = null)
            : base("cloudamqp:index/customDomain:CustomDomain", name, args ?? new CustomDomainArgs(), MakeResourceOptions(options, ""))
        {
        }

        private CustomDomain(string name, Input<string> id, CustomDomainState? state = null, CustomResourceOptions? options = null)
            : base("cloudamqp:index/customDomain:CustomDomain", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing CustomDomain resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static CustomDomain Get(string name, Input<string> id, CustomDomainState? state = null, CustomResourceOptions? options = null)
        {
            return new CustomDomain(name, id, state, options);
        }
    }

    public sealed class CustomDomainArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Your custom domain name.
        /// </summary>
        [Input("hostname", required: true)]
        public Input<string> Hostname { get; set; } = null!;

        /// <summary>
        /// The CloudAMQP instance ID.
        /// </summary>
        [Input("instanceId", required: true)]
        public Input<int> InstanceId { get; set; } = null!;

        public CustomDomainArgs()
        {
        }
        public static new CustomDomainArgs Empty => new CustomDomainArgs();
    }

    public sealed class CustomDomainState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Your custom domain name.
        /// </summary>
        [Input("hostname")]
        public Input<string>? Hostname { get; set; }

        /// <summary>
        /// The CloudAMQP instance ID.
        /// </summary>
        [Input("instanceId")]
        public Input<int>? InstanceId { get; set; }

        public CustomDomainState()
        {
        }
        public static new CustomDomainState Empty => new CustomDomainState();
    }
}
