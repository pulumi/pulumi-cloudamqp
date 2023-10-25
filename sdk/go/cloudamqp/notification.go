// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cloudamqp

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-cloudamqp/sdk/v3/go/cloudamqp/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// This resource allows you to create and manage recipients to receive alarm notifications. There will always be a default recipient created upon instance creation. This recipient will use team email and receive notifications from default alarms.
//
// Available for all subscription plans.
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
//			_, err := cloudamqp.NewNotification(ctx, "emailRecipient", &cloudamqp.NotificationArgs{
//				InstanceId: pulumi.Any(cloudamqp_instance.Instance.Id),
//				Type:       pulumi.String("email"),
//				Value:      pulumi.String("alarm@example.com"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = cloudamqp.NewNotification(ctx, "victoropsRecipient", &cloudamqp.NotificationArgs{
//				InstanceId: pulumi.Any(cloudamqp_instance.Instance.Id),
//				Type:       pulumi.String("victorops"),
//				Value:      pulumi.String("<UUID>"),
//				Options: pulumi.StringMap{
//					"rk": pulumi.String("ROUTINGKEY"),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			_, err = cloudamqp.NewNotification(ctx, "pagerdutyRecipient", &cloudamqp.NotificationArgs{
//				InstanceId: pulumi.Any(cloudamqp_instance.Instance.Id),
//				Type:       pulumi.String("pagerduty"),
//				Value:      pulumi.String("<integration-key>"),
//				Options: pulumi.StringMap{
//					"dedupkey": pulumi.String("DEDUPKEY"),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// ## Notification Type reference
//
// Valid options for notification type.
//
// * email
// * webhook
// * pagerduty
// * victorops
// * opsgenie
// * opsgenie-eu
// * slack
// * teams
//
// ## Options parameter
//
// | Type      | Options  | Description                                                                                                                                                                                                                                                                      | Note                                                                                                                                    |
// |-----------|----------|----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|-----------------------------------------------------------------------------------------------------------------------------------------|
// | Victorops | rk       | Routing key to route alarm notification                                                                                                                                                                                                                                          | -                                                                                                                                        |
// | PagerDuty | dedupkey | Default the dedup key for PagerDuty is generated depending on what alarm has triggered, but here you can set what `dedup` key to use so even if the same alarm is triggered for different resources you only get one notification. Leave blank to use the generated dedup key. | If multiple alarms are triggered using this recipient, since they all share `dedup` key only the first alarm will be shown in PagerDuty |
//
// ## Dependency
//
// This resource depends on CloudAMQP instance identifier, `cloudamqp_instance.instance.id`.
//
// ## Import
//
// `cloudamqp_notification` can be imported using CloudAMQP internal identifier of a recipient together (CSV separated) with the instance identifier. To retrieve the identifier of a recipient, use [CloudAMQP API](https://docs.cloudamqp.com/cloudamqp_api.html#list-notification-recipients)
//
// ```sh
//
//	$ pulumi import cloudamqp:index/notification:Notification recipient <id>,<instance_id>`
//
// ```
type Notification struct {
	pulumi.CustomResourceState

	// The CloudAMQP instance ID.
	InstanceId pulumi.IntOutput `pulumi:"instanceId"`
	// Display name of the recipient.
	Name pulumi.StringOutput `pulumi:"name"`
	// Options argument (e.g. `rk` used for VictorOps routing key).
	Options pulumi.StringMapOutput `pulumi:"options"`
	// Type of the notification. See valid options below.
	Type pulumi.StringOutput `pulumi:"type"`
	// Integration/API key or endpoint to send the notification.
	Value pulumi.StringOutput `pulumi:"value"`
}

// NewNotification registers a new resource with the given unique name, arguments, and options.
func NewNotification(ctx *pulumi.Context,
	name string, args *NotificationArgs, opts ...pulumi.ResourceOption) (*Notification, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.InstanceId == nil {
		return nil, errors.New("invalid value for required argument 'InstanceId'")
	}
	if args.Type == nil {
		return nil, errors.New("invalid value for required argument 'Type'")
	}
	if args.Value == nil {
		return nil, errors.New("invalid value for required argument 'Value'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Notification
	err := ctx.RegisterResource("cloudamqp:index/notification:Notification", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNotification gets an existing Notification resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNotification(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NotificationState, opts ...pulumi.ResourceOption) (*Notification, error) {
	var resource Notification
	err := ctx.ReadResource("cloudamqp:index/notification:Notification", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Notification resources.
type notificationState struct {
	// The CloudAMQP instance ID.
	InstanceId *int `pulumi:"instanceId"`
	// Display name of the recipient.
	Name *string `pulumi:"name"`
	// Options argument (e.g. `rk` used for VictorOps routing key).
	Options map[string]string `pulumi:"options"`
	// Type of the notification. See valid options below.
	Type *string `pulumi:"type"`
	// Integration/API key or endpoint to send the notification.
	Value *string `pulumi:"value"`
}

type NotificationState struct {
	// The CloudAMQP instance ID.
	InstanceId pulumi.IntPtrInput
	// Display name of the recipient.
	Name pulumi.StringPtrInput
	// Options argument (e.g. `rk` used for VictorOps routing key).
	Options pulumi.StringMapInput
	// Type of the notification. See valid options below.
	Type pulumi.StringPtrInput
	// Integration/API key or endpoint to send the notification.
	Value pulumi.StringPtrInput
}

func (NotificationState) ElementType() reflect.Type {
	return reflect.TypeOf((*notificationState)(nil)).Elem()
}

type notificationArgs struct {
	// The CloudAMQP instance ID.
	InstanceId int `pulumi:"instanceId"`
	// Display name of the recipient.
	Name *string `pulumi:"name"`
	// Options argument (e.g. `rk` used for VictorOps routing key).
	Options map[string]string `pulumi:"options"`
	// Type of the notification. See valid options below.
	Type string `pulumi:"type"`
	// Integration/API key or endpoint to send the notification.
	Value string `pulumi:"value"`
}

// The set of arguments for constructing a Notification resource.
type NotificationArgs struct {
	// The CloudAMQP instance ID.
	InstanceId pulumi.IntInput
	// Display name of the recipient.
	Name pulumi.StringPtrInput
	// Options argument (e.g. `rk` used for VictorOps routing key).
	Options pulumi.StringMapInput
	// Type of the notification. See valid options below.
	Type pulumi.StringInput
	// Integration/API key or endpoint to send the notification.
	Value pulumi.StringInput
}

func (NotificationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*notificationArgs)(nil)).Elem()
}

type NotificationInput interface {
	pulumi.Input

	ToNotificationOutput() NotificationOutput
	ToNotificationOutputWithContext(ctx context.Context) NotificationOutput
}

func (*Notification) ElementType() reflect.Type {
	return reflect.TypeOf((**Notification)(nil)).Elem()
}

func (i *Notification) ToNotificationOutput() NotificationOutput {
	return i.ToNotificationOutputWithContext(context.Background())
}

func (i *Notification) ToNotificationOutputWithContext(ctx context.Context) NotificationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NotificationOutput)
}

func (i *Notification) ToOutput(ctx context.Context) pulumix.Output[*Notification] {
	return pulumix.Output[*Notification]{
		OutputState: i.ToNotificationOutputWithContext(ctx).OutputState,
	}
}

// NotificationArrayInput is an input type that accepts NotificationArray and NotificationArrayOutput values.
// You can construct a concrete instance of `NotificationArrayInput` via:
//
//	NotificationArray{ NotificationArgs{...} }
type NotificationArrayInput interface {
	pulumi.Input

	ToNotificationArrayOutput() NotificationArrayOutput
	ToNotificationArrayOutputWithContext(context.Context) NotificationArrayOutput
}

type NotificationArray []NotificationInput

func (NotificationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Notification)(nil)).Elem()
}

func (i NotificationArray) ToNotificationArrayOutput() NotificationArrayOutput {
	return i.ToNotificationArrayOutputWithContext(context.Background())
}

func (i NotificationArray) ToNotificationArrayOutputWithContext(ctx context.Context) NotificationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NotificationArrayOutput)
}

func (i NotificationArray) ToOutput(ctx context.Context) pulumix.Output[[]*Notification] {
	return pulumix.Output[[]*Notification]{
		OutputState: i.ToNotificationArrayOutputWithContext(ctx).OutputState,
	}
}

// NotificationMapInput is an input type that accepts NotificationMap and NotificationMapOutput values.
// You can construct a concrete instance of `NotificationMapInput` via:
//
//	NotificationMap{ "key": NotificationArgs{...} }
type NotificationMapInput interface {
	pulumi.Input

	ToNotificationMapOutput() NotificationMapOutput
	ToNotificationMapOutputWithContext(context.Context) NotificationMapOutput
}

type NotificationMap map[string]NotificationInput

func (NotificationMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Notification)(nil)).Elem()
}

func (i NotificationMap) ToNotificationMapOutput() NotificationMapOutput {
	return i.ToNotificationMapOutputWithContext(context.Background())
}

func (i NotificationMap) ToNotificationMapOutputWithContext(ctx context.Context) NotificationMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NotificationMapOutput)
}

func (i NotificationMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*Notification] {
	return pulumix.Output[map[string]*Notification]{
		OutputState: i.ToNotificationMapOutputWithContext(ctx).OutputState,
	}
}

type NotificationOutput struct{ *pulumi.OutputState }

func (NotificationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Notification)(nil)).Elem()
}

func (o NotificationOutput) ToNotificationOutput() NotificationOutput {
	return o
}

func (o NotificationOutput) ToNotificationOutputWithContext(ctx context.Context) NotificationOutput {
	return o
}

func (o NotificationOutput) ToOutput(ctx context.Context) pulumix.Output[*Notification] {
	return pulumix.Output[*Notification]{
		OutputState: o.OutputState,
	}
}

// The CloudAMQP instance ID.
func (o NotificationOutput) InstanceId() pulumi.IntOutput {
	return o.ApplyT(func(v *Notification) pulumi.IntOutput { return v.InstanceId }).(pulumi.IntOutput)
}

// Display name of the recipient.
func (o NotificationOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Notification) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Options argument (e.g. `rk` used for VictorOps routing key).
func (o NotificationOutput) Options() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Notification) pulumi.StringMapOutput { return v.Options }).(pulumi.StringMapOutput)
}

// Type of the notification. See valid options below.
func (o NotificationOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Notification) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// Integration/API key or endpoint to send the notification.
func (o NotificationOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v *Notification) pulumi.StringOutput { return v.Value }).(pulumi.StringOutput)
}

type NotificationArrayOutput struct{ *pulumi.OutputState }

func (NotificationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Notification)(nil)).Elem()
}

func (o NotificationArrayOutput) ToNotificationArrayOutput() NotificationArrayOutput {
	return o
}

func (o NotificationArrayOutput) ToNotificationArrayOutputWithContext(ctx context.Context) NotificationArrayOutput {
	return o
}

func (o NotificationArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*Notification] {
	return pulumix.Output[[]*Notification]{
		OutputState: o.OutputState,
	}
}

func (o NotificationArrayOutput) Index(i pulumi.IntInput) NotificationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Notification {
		return vs[0].([]*Notification)[vs[1].(int)]
	}).(NotificationOutput)
}

type NotificationMapOutput struct{ *pulumi.OutputState }

func (NotificationMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Notification)(nil)).Elem()
}

func (o NotificationMapOutput) ToNotificationMapOutput() NotificationMapOutput {
	return o
}

func (o NotificationMapOutput) ToNotificationMapOutputWithContext(ctx context.Context) NotificationMapOutput {
	return o
}

func (o NotificationMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*Notification] {
	return pulumix.Output[map[string]*Notification]{
		OutputState: o.OutputState,
	}
}

func (o NotificationMapOutput) MapIndex(k pulumi.StringInput) NotificationOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Notification {
		return vs[0].(map[string]*Notification)[vs[1].(string)]
	}).(NotificationOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*NotificationInput)(nil)).Elem(), &Notification{})
	pulumi.RegisterInputType(reflect.TypeOf((*NotificationArrayInput)(nil)).Elem(), NotificationArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*NotificationMapInput)(nil)).Elem(), NotificationMap{})
	pulumi.RegisterOutputType(NotificationOutput{})
	pulumi.RegisterOutputType(NotificationArrayOutput{})
	pulumi.RegisterOutputType(NotificationMapOutput{})
}
