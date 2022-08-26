// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.CloudAmqp
{
    public static class GetVpcInfo
    {
        /// <summary>
        /// Use this data source to retrieve information about VPC for a CloudAMQP instance.
        /// 
        /// Only available for CloudAMQP instances hosted in AWS.
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
        ///     var vpcInfo = CloudAmqp.GetVpcInfo.Invoke(new()
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
        ///     var vpcInfo = CloudAmqp.GetVpcInfo.Invoke(new()
        ///     {
        ///         VpcId = cloudamqp_vpc.Vpc.Id,
        ///     });
        /// 
        /// });
        /// ```
        /// &lt;/details&gt;
        /// 
        /// {{% /example %}}
        /// {{% /examples %}}
        /// ## Attributes reference
        /// 
        /// All attributes reference are computed
        /// 
        /// * `id`                  - The identifier for this resource.
        /// * `name`                - The name of the CloudAMQP instance.
        /// * `vpc_subnet`          - Dedicated VPC subnet.
        /// * `owner_id`            - AWS account identifier.
        /// * `security_group_id`   - AWS security group identifier.
        /// 
        /// ## Dependency
        /// 
        /// *Pre v1.16.0*
        /// This resource depends on CloudAMQP instance identifier, `cloudamqp_instance.instance.id`.
        /// 
        /// *Post v1.16.0*
        /// This resource depends on CloudAMQP managed VPC identifier, `cloudamqp_vpc.vpc.id` or instance identifier, `cloudamqp_instance.instance.id`.
        /// </summary>
        public static Task<GetVpcInfoResult> InvokeAsync(GetVpcInfoArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetVpcInfoResult>("cloudamqp:index/getVpcInfo:getVpcInfo", args ?? new GetVpcInfoArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to retrieve information about VPC for a CloudAMQP instance.
        /// 
        /// Only available for CloudAMQP instances hosted in AWS.
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
        ///     var vpcInfo = CloudAmqp.GetVpcInfo.Invoke(new()
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
        ///     var vpcInfo = CloudAmqp.GetVpcInfo.Invoke(new()
        ///     {
        ///         VpcId = cloudamqp_vpc.Vpc.Id,
        ///     });
        /// 
        /// });
        /// ```
        /// &lt;/details&gt;
        /// 
        /// {{% /example %}}
        /// {{% /examples %}}
        /// ## Attributes reference
        /// 
        /// All attributes reference are computed
        /// 
        /// * `id`                  - The identifier for this resource.
        /// * `name`                - The name of the CloudAMQP instance.
        /// * `vpc_subnet`          - Dedicated VPC subnet.
        /// * `owner_id`            - AWS account identifier.
        /// * `security_group_id`   - AWS security group identifier.
        /// 
        /// ## Dependency
        /// 
        /// *Pre v1.16.0*
        /// This resource depends on CloudAMQP instance identifier, `cloudamqp_instance.instance.id`.
        /// 
        /// *Post v1.16.0*
        /// This resource depends on CloudAMQP managed VPC identifier, `cloudamqp_vpc.vpc.id` or instance identifier, `cloudamqp_instance.instance.id`.
        /// </summary>
        public static Output<GetVpcInfoResult> Invoke(GetVpcInfoInvokeArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetVpcInfoResult>("cloudamqp:index/getVpcInfo:getVpcInfo", args ?? new GetVpcInfoInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetVpcInfoArgs : global::Pulumi.InvokeArgs
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

        public GetVpcInfoArgs()
        {
        }
        public static new GetVpcInfoArgs Empty => new GetVpcInfoArgs();
    }

    public sealed class GetVpcInfoInvokeArgs : global::Pulumi.InvokeArgs
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

        public GetVpcInfoInvokeArgs()
        {
        }
        public static new GetVpcInfoInvokeArgs Empty => new GetVpcInfoInvokeArgs();
    }


    [OutputType]
    public sealed class GetVpcInfoResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly int? InstanceId;
        public readonly string Name;
        public readonly string OwnerId;
        public readonly string SecurityGroupId;
        public readonly string? VpcId;
        public readonly string VpcSubnet;

        [OutputConstructor]
        private GetVpcInfoResult(
            string id,

            int? instanceId,

            string name,

            string ownerId,

            string securityGroupId,

            string? vpcId,

            string vpcSubnet)
        {
            Id = id;
            InstanceId = instanceId;
            Name = name;
            OwnerId = ownerId;
            SecurityGroupId = securityGroupId;
            VpcId = vpcId;
            VpcSubnet = vpcSubnet;
        }
    }
}
