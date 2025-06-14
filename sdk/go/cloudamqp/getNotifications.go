// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cloudamqp

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-cloudamqp/sdk/v3/go/cloudamqp/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to retrieve information about all notification recipients. Each recipient will
// receive notifications assigned to an alarm that has triggered.
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
//			_, err := cloudamqp.GetNotifications(ctx, &cloudamqp.GetNotificationsArgs{
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
func GetNotifications(ctx *pulumi.Context, args *GetNotificationsArgs, opts ...pulumi.InvokeOption) (*GetNotificationsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetNotificationsResult
	err := ctx.Invoke("cloudamqp:index/getNotifications:getNotifications", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getNotifications.
type GetNotificationsArgs struct {
	// The CloudAMQP instance identifier.
	InstanceId int `pulumi:"instanceId"`
}

// A collection of values returned by getNotifications.
type GetNotificationsResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id         string `pulumi:"id"`
	InstanceId int    `pulumi:"instanceId"`
	// List of alarms (see below for nested schema)
	Recipients []GetNotificationsRecipient `pulumi:"recipients"`
}

func GetNotificationsOutput(ctx *pulumi.Context, args GetNotificationsOutputArgs, opts ...pulumi.InvokeOption) GetNotificationsResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (GetNotificationsResultOutput, error) {
			args := v.(GetNotificationsArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("cloudamqp:index/getNotifications:getNotifications", args, GetNotificationsResultOutput{}, options).(GetNotificationsResultOutput), nil
		}).(GetNotificationsResultOutput)
}

// A collection of arguments for invoking getNotifications.
type GetNotificationsOutputArgs struct {
	// The CloudAMQP instance identifier.
	InstanceId pulumi.IntInput `pulumi:"instanceId"`
}

func (GetNotificationsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetNotificationsArgs)(nil)).Elem()
}

// A collection of values returned by getNotifications.
type GetNotificationsResultOutput struct{ *pulumi.OutputState }

func (GetNotificationsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetNotificationsResult)(nil)).Elem()
}

func (o GetNotificationsResultOutput) ToGetNotificationsResultOutput() GetNotificationsResultOutput {
	return o
}

func (o GetNotificationsResultOutput) ToGetNotificationsResultOutputWithContext(ctx context.Context) GetNotificationsResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o GetNotificationsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetNotificationsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetNotificationsResultOutput) InstanceId() pulumi.IntOutput {
	return o.ApplyT(func(v GetNotificationsResult) int { return v.InstanceId }).(pulumi.IntOutput)
}

// List of alarms (see below for nested schema)
func (o GetNotificationsResultOutput) Recipients() GetNotificationsRecipientArrayOutput {
	return o.ApplyT(func(v GetNotificationsResult) []GetNotificationsRecipient { return v.Recipients }).(GetNotificationsRecipientArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetNotificationsResultOutput{})
}
