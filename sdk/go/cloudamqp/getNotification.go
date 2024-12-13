// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cloudamqp

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-cloudamqp/sdk/v3/go/cloudamqp/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to retrieve information about default or created recipients. The recipient will receive notifications assigned to an alarm that has triggered. To retrieve the recipient either use `recipientId` or `name`.
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
//			_, err := cloudamqp.LookupNotification(ctx, &cloudamqp.LookupNotificationArgs{
//				InstanceId: instance.Id,
//				Name:       pulumi.StringRef("default"),
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
// ## Attributes reference
//
// # All attributes reference are computed
//
// * `id`    - The identifier for this resource.
// * `type`  - The type of the recipient.
// * `value` - The notification endpoint, where to send the notification.
// * `options`- Options argument (e.g. `rk` used for VictorOps routing key).
//
// ## Dependency
//
// This data source depends on CloudAMQP instance identifier, `cloudamqp_instance.instance.id`.
func LookupNotification(ctx *pulumi.Context, args *LookupNotificationArgs, opts ...pulumi.InvokeOption) (*LookupNotificationResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupNotificationResult
	err := ctx.Invoke("cloudamqp:index/getNotification:getNotification", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getNotification.
type LookupNotificationArgs struct {
	// The CloudAMQP instance identifier.
	InstanceId int `pulumi:"instanceId"`
	// The name set for the recipient.
	Name    *string           `pulumi:"name"`
	Options map[string]string `pulumi:"options"`
	// The recipient identifier.
	RecipientId *int `pulumi:"recipientId"`
}

// A collection of values returned by getNotification.
type LookupNotificationResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id          string            `pulumi:"id"`
	InstanceId  int               `pulumi:"instanceId"`
	Name        *string           `pulumi:"name"`
	Options     map[string]string `pulumi:"options"`
	RecipientId *int              `pulumi:"recipientId"`
	Type        string            `pulumi:"type"`
	Value       string            `pulumi:"value"`
}

func LookupNotificationOutput(ctx *pulumi.Context, args LookupNotificationOutputArgs, opts ...pulumi.InvokeOption) LookupNotificationResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupNotificationResultOutput, error) {
			args := v.(LookupNotificationArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("cloudamqp:index/getNotification:getNotification", args, LookupNotificationResultOutput{}, options).(LookupNotificationResultOutput), nil
		}).(LookupNotificationResultOutput)
}

// A collection of arguments for invoking getNotification.
type LookupNotificationOutputArgs struct {
	// The CloudAMQP instance identifier.
	InstanceId pulumi.IntInput `pulumi:"instanceId"`
	// The name set for the recipient.
	Name    pulumi.StringPtrInput `pulumi:"name"`
	Options pulumi.StringMapInput `pulumi:"options"`
	// The recipient identifier.
	RecipientId pulumi.IntPtrInput `pulumi:"recipientId"`
}

func (LookupNotificationOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNotificationArgs)(nil)).Elem()
}

// A collection of values returned by getNotification.
type LookupNotificationResultOutput struct{ *pulumi.OutputState }

func (LookupNotificationResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNotificationResult)(nil)).Elem()
}

func (o LookupNotificationResultOutput) ToLookupNotificationResultOutput() LookupNotificationResultOutput {
	return o
}

func (o LookupNotificationResultOutput) ToLookupNotificationResultOutputWithContext(ctx context.Context) LookupNotificationResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o LookupNotificationResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNotificationResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupNotificationResultOutput) InstanceId() pulumi.IntOutput {
	return o.ApplyT(func(v LookupNotificationResult) int { return v.InstanceId }).(pulumi.IntOutput)
}

func (o LookupNotificationResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNotificationResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o LookupNotificationResultOutput) Options() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupNotificationResult) map[string]string { return v.Options }).(pulumi.StringMapOutput)
}

func (o LookupNotificationResultOutput) RecipientId() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupNotificationResult) *int { return v.RecipientId }).(pulumi.IntPtrOutput)
}

func (o LookupNotificationResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNotificationResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupNotificationResultOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNotificationResult) string { return v.Value }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupNotificationResultOutput{})
}
