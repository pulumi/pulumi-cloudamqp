// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.CloudAmqp
{
    public static class GetVpcGcpInfo
    {
        /// <summary>
        /// Use this data source to retrieve information about VPC for a CloudAMQP instance hosted in GCP.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// &lt;details&gt;
        ///   &lt;summary&gt;
        ///     &lt;b&gt;
        ///       &lt;i&gt;AWS VPC peering pre v1.16.0&lt;/i&gt;
        ///     &lt;/b&gt;
        ///   &lt;/summary&gt;
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using Pulumi;
        /// using CloudAmqp = Pulumi.CloudAmqp;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var vpcInfo = CloudAmqp.GetVpcGcpInfo.Invoke(new()
        ///     {
        ///         InstanceId = cloudamqp_instance.Instance.Id,
        ///     });
        /// 
        /// });
        /// ```
        /// &lt;/details&gt;
        /// 
        /// &lt;details&gt;
        ///   &lt;summary&gt;
        ///     &lt;b&gt;
        ///       &lt;i&gt;AWS VPC peering post v1.16.0 (Managed VPC)&lt;/i&gt;
        ///     &lt;/b&gt;
        ///   &lt;/summary&gt;
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using Pulumi;
        /// using CloudAmqp = Pulumi.CloudAmqp;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var vpcInfo = CloudAmqp.GetVpcGcpInfo.Invoke(new()
        ///     {
        ///         VpcId = cloudamqp_vpc.Vpc.Id,
        ///     });
        /// 
        /// });
        /// ```
        /// &lt;/details&gt;
        /// {{% /example %}}
        /// {{% /examples %}}
        /// ## Attributes reference
        /// 
        /// All attributes reference are computed
        /// 
        /// * `id`                  - The identifier for this resource.
        /// * `name`                - The name of the VPC.
        /// * `vpc_subnet`          - Dedicated VPC subnet.
        /// * `network`             - VPC network uri.
        /// 
        /// ## Dependency
        /// 
        /// *Pre v1.16.0*
        /// This resource depends on CloudAMQP instance identifier, `cloudamqp_instance.instance.id`.
        /// 
        /// *Post v1.16.0*
        /// This resource depends on CloudAMQP managed VPC identifier, `cloudamqp_vpc.vpc.id` or instance identifier, `cloudamqp_instance.instance.id`.
        /// </summary>
        public static Task<GetVpcGcpInfoResult> InvokeAsync(GetVpcGcpInfoArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetVpcGcpInfoResult>("cloudamqp:index/getVpcGcpInfo:getVpcGcpInfo", args ?? new GetVpcGcpInfoArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to retrieve information about VPC for a CloudAMQP instance hosted in GCP.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// &lt;details&gt;
        ///   &lt;summary&gt;
        ///     &lt;b&gt;
        ///       &lt;i&gt;AWS VPC peering pre v1.16.0&lt;/i&gt;
        ///     &lt;/b&gt;
        ///   &lt;/summary&gt;
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using Pulumi;
        /// using CloudAmqp = Pulumi.CloudAmqp;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var vpcInfo = CloudAmqp.GetVpcGcpInfo.Invoke(new()
        ///     {
        ///         InstanceId = cloudamqp_instance.Instance.Id,
        ///     });
        /// 
        /// });
        /// ```
        /// &lt;/details&gt;
        /// 
        /// &lt;details&gt;
        ///   &lt;summary&gt;
        ///     &lt;b&gt;
        ///       &lt;i&gt;AWS VPC peering post v1.16.0 (Managed VPC)&lt;/i&gt;
        ///     &lt;/b&gt;
        ///   &lt;/summary&gt;
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using Pulumi;
        /// using CloudAmqp = Pulumi.CloudAmqp;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var vpcInfo = CloudAmqp.GetVpcGcpInfo.Invoke(new()
        ///     {
        ///         VpcId = cloudamqp_vpc.Vpc.Id,
        ///     });
        /// 
        /// });
        /// ```
        /// &lt;/details&gt;
        /// {{% /example %}}
        /// {{% /examples %}}
        /// ## Attributes reference
        /// 
        /// All attributes reference are computed
        /// 
        /// * `id`                  - The identifier for this resource.
        /// * `name`                - The name of the VPC.
        /// * `vpc_subnet`          - Dedicated VPC subnet.
        /// * `network`             - VPC network uri.
        /// 
        /// ## Dependency
        /// 
        /// *Pre v1.16.0*
        /// This resource depends on CloudAMQP instance identifier, `cloudamqp_instance.instance.id`.
        /// 
        /// *Post v1.16.0*
        /// This resource depends on CloudAMQP managed VPC identifier, `cloudamqp_vpc.vpc.id` or instance identifier, `cloudamqp_instance.instance.id`.
        /// </summary>
        public static Output<GetVpcGcpInfoResult> Invoke(GetVpcGcpInfoInvokeArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetVpcGcpInfoResult>("cloudamqp:index/getVpcGcpInfo:getVpcGcpInfo", args ?? new GetVpcGcpInfoInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetVpcGcpInfoArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The CloudAMQP instance identifier.
        /// </summary>
        [Input("instanceId")]
        public int? InstanceId { get; set; }

        /// <summary>
        /// The managed VPC identifier.
        /// </summary>
        [Input("vpcId")]
        public string? VpcId { get; set; }

        public GetVpcGcpInfoArgs()
        {
        }
        public static new GetVpcGcpInfoArgs Empty => new GetVpcGcpInfoArgs();
    }

    public sealed class GetVpcGcpInfoInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The CloudAMQP instance identifier.
        /// </summary>
        [Input("instanceId")]
        public Input<int>? InstanceId { get; set; }

        /// <summary>
        /// The managed VPC identifier.
        /// </summary>
        [Input("vpcId")]
        public Input<string>? VpcId { get; set; }

        public GetVpcGcpInfoInvokeArgs()
        {
        }
        public static new GetVpcGcpInfoInvokeArgs Empty => new GetVpcGcpInfoInvokeArgs();
    }


    [OutputType]
    public sealed class GetVpcGcpInfoResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly int? InstanceId;
        public readonly string Name;
        public readonly string Network;
        public readonly string? VpcId;
        public readonly string VpcSubnet;

        [OutputConstructor]
        private GetVpcGcpInfoResult(
            string id,

            int? instanceId,

            string name,

            string network,

            string? vpcId,

            string vpcSubnet)
        {
            Id = id;
            InstanceId = instanceId;
            Name = name;
            Network = network;
            VpcId = vpcId;
            VpcSubnet = vpcSubnet;
        }
    }
}
