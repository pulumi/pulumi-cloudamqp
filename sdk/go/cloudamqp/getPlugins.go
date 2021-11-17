// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package cloudamqp

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to retrieve information about installed and available plugins for the CloudAMQP instance.
//
// ⚠️  From our go API wrapper [v1.4.0](https://github.com/84codes/go-api/releases/tag/v1.4.0) there is support for multiple retries when requesting information about plugins. This was introduced to avoid `ReadPlugin error 400: Timeout talking to backend`.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-cloudamqp/sdk/v3/go/cloudamqp"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := cloudamqp.GetPlugins(ctx, &GetPluginsArgs{
// 			InstanceId: cloudamqp_instance.Instance.Id,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ## Argument reference
//
// * `instanceId` - (Required) The CloudAMQP instance identifier.
//
// ## Attributes reference
//
// All attributes reference are computed
//
// * `id`      - The identifier for this resource.
// * `plugins` - An array of plugins. Each `plugins` block consists of the fields documented below.
//
// ***
//
// The `plugins` block consist of
//
// * `name`        - The type of the recipient.
// * `version`     - Rabbit MQ version that the plugins are shipped with.
// * `description` - Description of what the plugin does.
// * `enabled`     - Enable or disable information for the plugin.
//
// ## Dependency
//
// This data source depends on CloudAMQP instance identifier, `cloudamqp_instance.instance.id`.
func GetPlugins(ctx *pulumi.Context, args *GetPluginsArgs, opts ...pulumi.InvokeOption) (*GetPluginsResult, error) {
	var rv GetPluginsResult
	err := ctx.Invoke("cloudamqp:index/getPlugins:getPlugins", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getPlugins.
type GetPluginsArgs struct {
	InstanceId int `pulumi:"instanceId"`
}

// A collection of values returned by getPlugins.
type GetPluginsResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id         string             `pulumi:"id"`
	InstanceId int                `pulumi:"instanceId"`
	Plugins    []GetPluginsPlugin `pulumi:"plugins"`
}

func GetPluginsOutput(ctx *pulumi.Context, args GetPluginsOutputArgs, opts ...pulumi.InvokeOption) GetPluginsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetPluginsResult, error) {
			args := v.(GetPluginsArgs)
			r, err := GetPlugins(ctx, &args, opts...)
			return *r, err
		}).(GetPluginsResultOutput)
}

// A collection of arguments for invoking getPlugins.
type GetPluginsOutputArgs struct {
	InstanceId pulumi.IntInput `pulumi:"instanceId"`
}

func (GetPluginsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetPluginsArgs)(nil)).Elem()
}

// A collection of values returned by getPlugins.
type GetPluginsResultOutput struct{ *pulumi.OutputState }

func (GetPluginsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetPluginsResult)(nil)).Elem()
}

func (o GetPluginsResultOutput) ToGetPluginsResultOutput() GetPluginsResultOutput {
	return o
}

func (o GetPluginsResultOutput) ToGetPluginsResultOutputWithContext(ctx context.Context) GetPluginsResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o GetPluginsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetPluginsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetPluginsResultOutput) InstanceId() pulumi.IntOutput {
	return o.ApplyT(func(v GetPluginsResult) int { return v.InstanceId }).(pulumi.IntOutput)
}

func (o GetPluginsResultOutput) Plugins() GetPluginsPluginArrayOutput {
	return o.ApplyT(func(v GetPluginsResult) []GetPluginsPlugin { return v.Plugins }).(GetPluginsPluginArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetPluginsResultOutput{})
}
