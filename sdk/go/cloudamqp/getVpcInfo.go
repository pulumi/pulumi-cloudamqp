// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cloudamqp

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-cloudamqp/sdk/v3/go/cloudamqp/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to retrieve information about VPC for a CloudAMQP instance.
//
// Only available for CloudAMQP instances hosted in AWS.
//
// ## Example Usage
//
// <details>
//
//	<summary>
//	  <b>
//	    <i>AWS VPC peering pre v1.16.0</i>
//	  </b>
//	</summary>
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-cloudamqp/sdk/v3/go/cloudamqp"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := cloudamqp.GetVpcInfo(ctx, &cloudamqp.GetVpcInfoArgs{
//				InstanceId: pulumi.IntRef(instance.Id),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// </details>
//
// <details>
//
//	<summary>
//	  <b>
//	    <i>AWS VPC peering post v1.16.0 (Managed VPC)</i>
//	  </b>
//	</summary>
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-cloudamqp/sdk/v3/go/cloudamqp"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := cloudamqp.GetVpcInfo(ctx, &cloudamqp.GetVpcInfoArgs{
//				VpcId: pulumi.StringRef(vpc.Id),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// </details>
//
// ## Attributes reference
//
// # All attributes reference are computed
//
// * `id`                  - The identifier for this resource.
// * `name`                - The name of the CloudAMQP instance.
// * `vpcSubnet`          - Dedicated VPC subnet.
// * `ownerId`            - AWS account identifier.
// * `securityGroupId`   - AWS security group identifier.
//
// ## Dependency
//
// *Pre v1.16.0*
// This resource depends on CloudAMQP instance identifier, `cloudamqp_instance.instance.id`.
//
// *Post v1.16.0*
// This resource depends on CloudAMQP managed VPC identifier, `cloudamqp_vpc.vpc.id` or instance identifier, `cloudamqp_instance.instance.id`.
func GetVpcInfo(ctx *pulumi.Context, args *GetVpcInfoArgs, opts ...pulumi.InvokeOption) (*GetVpcInfoResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetVpcInfoResult
	err := ctx.Invoke("cloudamqp:index/getVpcInfo:getVpcInfo", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getVpcInfo.
type GetVpcInfoArgs struct {
	// The CloudAMQP instance identifier.
	//
	// ***Deprecated: Changed from required to optional in v1.16.0 will be removed in next major version (v2.0)***
	InstanceId *int `pulumi:"instanceId"`
	// The managed VPC identifier.
	//
	// ***Note: Added as optional in version v1.16.0 and will be required in next major version (v2.0)***
	VpcId *string `pulumi:"vpcId"`
}

// A collection of values returned by getVpcInfo.
type GetVpcInfoResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id              string  `pulumi:"id"`
	InstanceId      *int    `pulumi:"instanceId"`
	Name            string  `pulumi:"name"`
	OwnerId         string  `pulumi:"ownerId"`
	SecurityGroupId string  `pulumi:"securityGroupId"`
	VpcId           *string `pulumi:"vpcId"`
	VpcSubnet       string  `pulumi:"vpcSubnet"`
}

func GetVpcInfoOutput(ctx *pulumi.Context, args GetVpcInfoOutputArgs, opts ...pulumi.InvokeOption) GetVpcInfoResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetVpcInfoResult, error) {
			args := v.(GetVpcInfoArgs)
			r, err := GetVpcInfo(ctx, &args, opts...)
			var s GetVpcInfoResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetVpcInfoResultOutput)
}

// A collection of arguments for invoking getVpcInfo.
type GetVpcInfoOutputArgs struct {
	// The CloudAMQP instance identifier.
	//
	// ***Deprecated: Changed from required to optional in v1.16.0 will be removed in next major version (v2.0)***
	InstanceId pulumi.IntPtrInput `pulumi:"instanceId"`
	// The managed VPC identifier.
	//
	// ***Note: Added as optional in version v1.16.0 and will be required in next major version (v2.0)***
	VpcId pulumi.StringPtrInput `pulumi:"vpcId"`
}

func (GetVpcInfoOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetVpcInfoArgs)(nil)).Elem()
}

// A collection of values returned by getVpcInfo.
type GetVpcInfoResultOutput struct{ *pulumi.OutputState }

func (GetVpcInfoResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetVpcInfoResult)(nil)).Elem()
}

func (o GetVpcInfoResultOutput) ToGetVpcInfoResultOutput() GetVpcInfoResultOutput {
	return o
}

func (o GetVpcInfoResultOutput) ToGetVpcInfoResultOutputWithContext(ctx context.Context) GetVpcInfoResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o GetVpcInfoResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetVpcInfoResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetVpcInfoResultOutput) InstanceId() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GetVpcInfoResult) *int { return v.InstanceId }).(pulumi.IntPtrOutput)
}

func (o GetVpcInfoResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetVpcInfoResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o GetVpcInfoResultOutput) OwnerId() pulumi.StringOutput {
	return o.ApplyT(func(v GetVpcInfoResult) string { return v.OwnerId }).(pulumi.StringOutput)
}

func (o GetVpcInfoResultOutput) SecurityGroupId() pulumi.StringOutput {
	return o.ApplyT(func(v GetVpcInfoResult) string { return v.SecurityGroupId }).(pulumi.StringOutput)
}

func (o GetVpcInfoResultOutput) VpcId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetVpcInfoResult) *string { return v.VpcId }).(pulumi.StringPtrOutput)
}

func (o GetVpcInfoResultOutput) VpcSubnet() pulumi.StringOutput {
	return o.ApplyT(func(v GetVpcInfoResult) string { return v.VpcSubnet }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetVpcInfoResultOutput{})
}
