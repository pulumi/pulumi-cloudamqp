// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.CloudAmqp
{
    /// <summary>
    /// The provider type for the cloudamqp package. By default, resources use package-wide configuration
    /// settings, however an explicit `Provider` instance may be created and passed during resource
    /// construction to achieve fine-grained programmatic control over provider settings. See the
    /// [documentation](https://www.pulumi.com/docs/reference/programming-model/#providers) for more information.
    /// </summary>
    [CloudAmqpResourceType("pulumi:providers:cloudamqp")]
    public partial class Provider : global::Pulumi.ProviderResource
    {
        /// <summary>
        /// Key used to authentication to the CloudAMQP Customer API
        /// </summary>
        [Output("apikey")]
        public Output<string?> Apikey { get; private set; } = null!;

        /// <summary>
        /// Base URL to CloudAMQP Customer website
        /// </summary>
        [Output("baseurl")]
        public Output<string?> Baseurl { get; private set; } = null!;


        /// <summary>
        /// Create a Provider resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Provider(string name, ProviderArgs? args = null, CustomResourceOptions? options = null)
            : base("cloudamqp", name, args ?? new ProviderArgs(), MakeResourceOptions(options, ""))
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
        /// This function returns a Terraform config object with terraform-namecased keys,to be used with the Terraform Module Provider.
        /// </summary>
        public global::Pulumi.Output<ProviderTerraformConfigResult> TerraformConfig()
            => global::Pulumi.Deployment.Instance.Call<ProviderTerraformConfigResult>("pulumi:providers:cloudamqp/terraformConfig", CallArgs.Empty, this);
    }

    public sealed class ProviderArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Key used to authentication to the CloudAMQP Customer API
        /// </summary>
        [Input("apikey")]
        public Input<string>? Apikey { get; set; }

        /// <summary>
        /// Base URL to CloudAMQP Customer website
        /// </summary>
        [Input("baseurl")]
        public Input<string>? Baseurl { get; set; }

        [Input("enableFasterInstanceDestroy", json: true)]
        public Input<bool>? EnableFasterInstanceDestroy { get; set; }

        public ProviderArgs()
        {
        }
        public static new ProviderArgs Empty => new ProviderArgs();
    }

    /// <summary>
    /// The results of the <see cref="Provider.TerraformConfig"/> method.
    /// </summary>
    [OutputType]
    public sealed class ProviderTerraformConfigResult
    {
        public readonly ImmutableDictionary<string, object> Result;

        [OutputConstructor]
        private ProviderTerraformConfigResult(ImmutableDictionary<string, object> result)
        {
            Result = result;
        }
    }
}
