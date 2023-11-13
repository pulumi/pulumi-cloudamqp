// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cloudamqp

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-cloudamqp/sdk/v3/go/cloudamqp/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## Example Usage
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
//			_, err := cloudamqp.GetPluginsCommunity(ctx, &cloudamqp.GetPluginsCommunityArgs{
//				InstanceId: cloudamqp_instance.Instance.Id,
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// ## Attributes reference
//
// # All attributes reference are computed
//
// * `id`      - The identifier for this resource.
// * `plugins` - An array of community plugins. Each `plugins` block consists of the fields documented below.
//
// ***
//
// # The `plugins` block consists of
//
// * `name`        - The type of the recipient.
// * `require`     - Min. required Rabbit MQ version to be used.
// * `description` - Description of what the plugin does.
//
// ## Dependency
//
// This data source depends on CloudAMQP instance identifier, `cloudamqp_instance.instance.id`.
func GetPluginsCommunity(ctx *pulumi.Context, args *GetPluginsCommunityArgs, opts ...pulumi.InvokeOption) (*GetPluginsCommunityResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetPluginsCommunityResult
	err := ctx.Invoke("cloudamqp:index/getPluginsCommunity:getPluginsCommunity", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getPluginsCommunity.
type GetPluginsCommunityArgs struct {
	// The CloudAMQP instance identifier.
	InstanceId int `pulumi:"instanceId"`
}

// A collection of values returned by getPluginsCommunity.
type GetPluginsCommunityResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id         string                      `pulumi:"id"`
	InstanceId int                         `pulumi:"instanceId"`
	Plugins    []GetPluginsCommunityPlugin `pulumi:"plugins"`
}

func GetPluginsCommunityOutput(ctx *pulumi.Context, args GetPluginsCommunityOutputArgs, opts ...pulumi.InvokeOption) GetPluginsCommunityResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetPluginsCommunityResult, error) {
			args := v.(GetPluginsCommunityArgs)
			r, err := GetPluginsCommunity(ctx, &args, opts...)
			var s GetPluginsCommunityResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetPluginsCommunityResultOutput)
}

// A collection of arguments for invoking getPluginsCommunity.
type GetPluginsCommunityOutputArgs struct {
	// The CloudAMQP instance identifier.
	InstanceId pulumi.IntInput `pulumi:"instanceId"`
}

func (GetPluginsCommunityOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetPluginsCommunityArgs)(nil)).Elem()
}

// A collection of values returned by getPluginsCommunity.
type GetPluginsCommunityResultOutput struct{ *pulumi.OutputState }

func (GetPluginsCommunityResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetPluginsCommunityResult)(nil)).Elem()
}

func (o GetPluginsCommunityResultOutput) ToGetPluginsCommunityResultOutput() GetPluginsCommunityResultOutput {
	return o
}

func (o GetPluginsCommunityResultOutput) ToGetPluginsCommunityResultOutputWithContext(ctx context.Context) GetPluginsCommunityResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o GetPluginsCommunityResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetPluginsCommunityResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetPluginsCommunityResultOutput) InstanceId() pulumi.IntOutput {
	return o.ApplyT(func(v GetPluginsCommunityResult) int { return v.InstanceId }).(pulumi.IntOutput)
}

func (o GetPluginsCommunityResultOutput) Plugins() GetPluginsCommunityPluginArrayOutput {
	return o.ApplyT(func(v GetPluginsCommunityResult) []GetPluginsCommunityPlugin { return v.Plugins }).(GetPluginsCommunityPluginArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetPluginsCommunityResultOutput{})
}
