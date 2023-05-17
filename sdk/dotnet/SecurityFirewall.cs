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
    /// ## Import
    /// 
    /// `cloudamqp_security_firewall` can be imported using CloudAMQP instance identifier.
    /// 
    /// ```sh
    ///  $ pulumi import cloudamqp:index/securityFirewall:SecurityFirewall firewall &lt;instance_id&gt;`
    /// ```
    /// </summary>
    [CloudAmqpResourceType("cloudamqp:index/securityFirewall:SecurityFirewall")]
    public partial class SecurityFirewall : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The CloudAMQP instance ID.
        /// </summary>
        [Output("instanceId")]
        public Output<int> InstanceId { get; private set; } = null!;

        /// <summary>
        /// An array of rules, minimum of 1 needs to be configured. Each `rules` block consists of the field documented below.
        /// </summary>
        [Output("rules")]
        public Output<ImmutableArray<Outputs.SecurityFirewallRule>> Rules { get; private set; } = null!;

        /// <summary>
        /// Configurable sleep time in seconds between retries for firewall configuration. Default set to 30 seconds.
        /// </summary>
        [Output("sleep")]
        public Output<int?> Sleep { get; private set; } = null!;

        /// <summary>
        /// Configurable timeout time in seconds for firewall configuration. Default set to 1800 seconds.
        /// 
        /// ___
        /// 
        /// The `rules` block consists of:
        /// </summary>
        [Output("timeout")]
        public Output<int?> Timeout { get; private set; } = null!;


        /// <summary>
        /// Create a SecurityFirewall resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public SecurityFirewall(string name, SecurityFirewallArgs args, CustomResourceOptions? options = null)
            : base("cloudamqp:index/securityFirewall:SecurityFirewall", name, args ?? new SecurityFirewallArgs(), MakeResourceOptions(options, ""))
        {
        }

        private SecurityFirewall(string name, Input<string> id, SecurityFirewallState? state = null, CustomResourceOptions? options = null)
            : base("cloudamqp:index/securityFirewall:SecurityFirewall", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing SecurityFirewall resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static SecurityFirewall Get(string name, Input<string> id, SecurityFirewallState? state = null, CustomResourceOptions? options = null)
        {
            return new SecurityFirewall(name, id, state, options);
        }
    }

    public sealed class SecurityFirewallArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The CloudAMQP instance ID.
        /// </summary>
        [Input("instanceId", required: true)]
        public Input<int> InstanceId { get; set; } = null!;

        [Input("rules", required: true)]
        private InputList<Inputs.SecurityFirewallRuleArgs>? _rules;

        /// <summary>
        /// An array of rules, minimum of 1 needs to be configured. Each `rules` block consists of the field documented below.
        /// </summary>
        public InputList<Inputs.SecurityFirewallRuleArgs> Rules
        {
            get => _rules ?? (_rules = new InputList<Inputs.SecurityFirewallRuleArgs>());
            set => _rules = value;
        }

        /// <summary>
        /// Configurable sleep time in seconds between retries for firewall configuration. Default set to 30 seconds.
        /// </summary>
        [Input("sleep")]
        public Input<int>? Sleep { get; set; }

        /// <summary>
        /// Configurable timeout time in seconds for firewall configuration. Default set to 1800 seconds.
        /// 
        /// ___
        /// 
        /// The `rules` block consists of:
        /// </summary>
        [Input("timeout")]
        public Input<int>? Timeout { get; set; }

        public SecurityFirewallArgs()
        {
        }
        public static new SecurityFirewallArgs Empty => new SecurityFirewallArgs();
    }

    public sealed class SecurityFirewallState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The CloudAMQP instance ID.
        /// </summary>
        [Input("instanceId")]
        public Input<int>? InstanceId { get; set; }

        [Input("rules")]
        private InputList<Inputs.SecurityFirewallRuleGetArgs>? _rules;

        /// <summary>
        /// An array of rules, minimum of 1 needs to be configured. Each `rules` block consists of the field documented below.
        /// </summary>
        public InputList<Inputs.SecurityFirewallRuleGetArgs> Rules
        {
            get => _rules ?? (_rules = new InputList<Inputs.SecurityFirewallRuleGetArgs>());
            set => _rules = value;
        }

        /// <summary>
        /// Configurable sleep time in seconds between retries for firewall configuration. Default set to 30 seconds.
        /// </summary>
        [Input("sleep")]
        public Input<int>? Sleep { get; set; }

        /// <summary>
        /// Configurable timeout time in seconds for firewall configuration. Default set to 1800 seconds.
        /// 
        /// ___
        /// 
        /// The `rules` block consists of:
        /// </summary>
        [Input("timeout")]
        public Input<int>? Timeout { get; set; }

        public SecurityFirewallState()
        {
        }
        public static new SecurityFirewallState Empty => new SecurityFirewallState();
    }
}
