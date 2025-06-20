// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cloudamqp

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-cloudamqp/sdk/v3/go/cloudamqp/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to retrieve information about possible upgradable versions for RabbitMQ and
// Erlang.
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
//			_, err := cloudamqp.GetUpgradableVersions(ctx, &cloudamqp.GetUpgradableVersionsArgs{
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
func GetUpgradableVersions(ctx *pulumi.Context, args *GetUpgradableVersionsArgs, opts ...pulumi.InvokeOption) (*GetUpgradableVersionsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetUpgradableVersionsResult
	err := ctx.Invoke("cloudamqp:index/getUpgradableVersions:getUpgradableVersions", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getUpgradableVersions.
type GetUpgradableVersionsArgs struct {
	// The CloudAMQP instance identifier.
	InstanceId int `pulumi:"instanceId"`
}

// A collection of values returned by getUpgradableVersions.
type GetUpgradableVersionsResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id         string `pulumi:"id"`
	InstanceId int    `pulumi:"instanceId"`
	// Possible upgradable version for Erlang.
	NewErlangVersion string `pulumi:"newErlangVersion"`
	// Possible upgradable version for RabbitMQ.
	NewRabbitmqVersion string `pulumi:"newRabbitmqVersion"`
}

func GetUpgradableVersionsOutput(ctx *pulumi.Context, args GetUpgradableVersionsOutputArgs, opts ...pulumi.InvokeOption) GetUpgradableVersionsResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (GetUpgradableVersionsResultOutput, error) {
			args := v.(GetUpgradableVersionsArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("cloudamqp:index/getUpgradableVersions:getUpgradableVersions", args, GetUpgradableVersionsResultOutput{}, options).(GetUpgradableVersionsResultOutput), nil
		}).(GetUpgradableVersionsResultOutput)
}

// A collection of arguments for invoking getUpgradableVersions.
type GetUpgradableVersionsOutputArgs struct {
	// The CloudAMQP instance identifier.
	InstanceId pulumi.IntInput `pulumi:"instanceId"`
}

func (GetUpgradableVersionsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetUpgradableVersionsArgs)(nil)).Elem()
}

// A collection of values returned by getUpgradableVersions.
type GetUpgradableVersionsResultOutput struct{ *pulumi.OutputState }

func (GetUpgradableVersionsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetUpgradableVersionsResult)(nil)).Elem()
}

func (o GetUpgradableVersionsResultOutput) ToGetUpgradableVersionsResultOutput() GetUpgradableVersionsResultOutput {
	return o
}

func (o GetUpgradableVersionsResultOutput) ToGetUpgradableVersionsResultOutputWithContext(ctx context.Context) GetUpgradableVersionsResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o GetUpgradableVersionsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetUpgradableVersionsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetUpgradableVersionsResultOutput) InstanceId() pulumi.IntOutput {
	return o.ApplyT(func(v GetUpgradableVersionsResult) int { return v.InstanceId }).(pulumi.IntOutput)
}

// Possible upgradable version for Erlang.
func (o GetUpgradableVersionsResultOutput) NewErlangVersion() pulumi.StringOutput {
	return o.ApplyT(func(v GetUpgradableVersionsResult) string { return v.NewErlangVersion }).(pulumi.StringOutput)
}

// Possible upgradable version for RabbitMQ.
func (o GetUpgradableVersionsResultOutput) NewRabbitmqVersion() pulumi.StringOutput {
	return o.ApplyT(func(v GetUpgradableVersionsResult) string { return v.NewRabbitmqVersion }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetUpgradableVersionsResultOutput{})
}
