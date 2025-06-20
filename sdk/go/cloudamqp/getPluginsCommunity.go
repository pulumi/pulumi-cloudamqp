// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cloudamqp

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-cloudamqp/sdk/v3/go/cloudamqp/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to retrieve information about available community plugins for the CloudAMQP
// instance.
//
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
//				InstanceId: instance.Id,
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
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
	// Configurable sleep time (seconds) for retries when requesting
	// information about community plugins. Default set to 10 seconds.
	Sleep *int `pulumi:"sleep"`
	// Configurable timeout time (seconds) for retries when requesting
	// information about community plugins. Default set to 1800 seconds.
	Timeout *int `pulumi:"timeout"`
}

// A collection of values returned by getPluginsCommunity.
type GetPluginsCommunityResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id         string `pulumi:"id"`
	InstanceId int    `pulumi:"instanceId"`
	// An array of community plugins. Each `plugins` block consists of the fields documented
	// below.
	Plugins []GetPluginsCommunityPlugin `pulumi:"plugins"`
	Sleep   *int                        `pulumi:"sleep"`
	Timeout *int                        `pulumi:"timeout"`
}

func GetPluginsCommunityOutput(ctx *pulumi.Context, args GetPluginsCommunityOutputArgs, opts ...pulumi.InvokeOption) GetPluginsCommunityResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (GetPluginsCommunityResultOutput, error) {
			args := v.(GetPluginsCommunityArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("cloudamqp:index/getPluginsCommunity:getPluginsCommunity", args, GetPluginsCommunityResultOutput{}, options).(GetPluginsCommunityResultOutput), nil
		}).(GetPluginsCommunityResultOutput)
}

// A collection of arguments for invoking getPluginsCommunity.
type GetPluginsCommunityOutputArgs struct {
	// The CloudAMQP instance identifier.
	InstanceId pulumi.IntInput `pulumi:"instanceId"`
	// Configurable sleep time (seconds) for retries when requesting
	// information about community plugins. Default set to 10 seconds.
	Sleep pulumi.IntPtrInput `pulumi:"sleep"`
	// Configurable timeout time (seconds) for retries when requesting
	// information about community plugins. Default set to 1800 seconds.
	Timeout pulumi.IntPtrInput `pulumi:"timeout"`
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

// An array of community plugins. Each `plugins` block consists of the fields documented
// below.
func (o GetPluginsCommunityResultOutput) Plugins() GetPluginsCommunityPluginArrayOutput {
	return o.ApplyT(func(v GetPluginsCommunityResult) []GetPluginsCommunityPlugin { return v.Plugins }).(GetPluginsCommunityPluginArrayOutput)
}

func (o GetPluginsCommunityResultOutput) Sleep() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GetPluginsCommunityResult) *int { return v.Sleep }).(pulumi.IntPtrOutput)
}

func (o GetPluginsCommunityResultOutput) Timeout() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GetPluginsCommunityResult) *int { return v.Timeout }).(pulumi.IntPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetPluginsCommunityResultOutput{})
}
