// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cloudamqp

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-cloudamqp/sdk/v3/go/cloudamqp/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This resource allows you to create and manage recipients to receive alarm notifications. There will
// always be a default recipient created upon instance creation. This recipient will use team email and
// receive notifications from default alarms.
//
// Available for all subscription plans.
//
// ## Example Usage
//
// <details>
//
//	<summary>
//	  <b>Email recipient</b>
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
//			_, err := cloudamqp.NewNotification(ctx, "email_recipient", &cloudamqp.NotificationArgs{
//				InstanceId: pulumi.Any(instance.Id),
//				Type:       pulumi.String("email"),
//				Value:      pulumi.String("alarm@example.com"),
//				Name:       pulumi.String("alarm"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// </details>
//
// <details>
//
//	<summary>
//	  <b>OpsGenie recipient with optional responders</b>
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
//			_, err := cloudamqp.NewNotification(ctx, "opsgenie_recipient", &cloudamqp.NotificationArgs{
//				InstanceId: pulumi.Any(instance.Id),
//				Type:       pulumi.String("opsgenie"),
//				Value:      pulumi.String("<api-key>"),
//				Name:       pulumi.String("OpsGenie"),
//				Responders: cloudamqp.NotificationResponderArray{
//					&cloudamqp.NotificationResponderArgs{
//						Type: pulumi.String("team"),
//						Id:   pulumi.String("<team-uuid>"),
//					},
//					&cloudamqp.NotificationResponderArgs{
//						Type:     pulumi.String("user"),
//						Username: pulumi.String("<username>"),
//					},
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
//
// </details>
//
// <details>
//
//	<summary>
//	  <b>Pagerduty recipient with optional dedup key</b>
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
//			_, err := cloudamqp.NewNotification(ctx, "pagerduty_recipient", &cloudamqp.NotificationArgs{
//				InstanceId: pulumi.Any(instance.Id),
//				Type:       pulumi.String("pagerduty"),
//				Value:      pulumi.String("<integration-key>"),
//				Name:       pulumi.String("PagerDuty"),
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
//
// </details>
//
// <details>
//
//	<summary>
//	  <b>Signl4 recipient</b>
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
//			_, err := cloudamqp.NewNotification(ctx, "signl4_recipient", &cloudamqp.NotificationArgs{
//				InstanceId: pulumi.Any(instance.Id),
//				Type:       pulumi.String("signl4"),
//				Value:      pulumi.String("<team-secret>"),
//				Name:       pulumi.String("Signl4"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// </details>
//
// <details>
//
//	<summary>
//	  <b>Teams recipient</b>
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
//			_, err := cloudamqp.NewNotification(ctx, "teams_recipient", &cloudamqp.NotificationArgs{
//				InstanceId: pulumi.Any(instance.Id),
//				Type:       pulumi.String("teams"),
//				Value:      pulumi.String("<teams-webhook-url>"),
//				Name:       pulumi.String("Teams"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// </details>
//
// <details>
//
//	<summary>
//	  <b>Victorops recipient with optional routing key (rk)</b>
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
//			_, err := cloudamqp.NewNotification(ctx, "victorops_recipient", &cloudamqp.NotificationArgs{
//				InstanceId: pulumi.Any(instance.Id),
//				Type:       pulumi.String("victorops"),
//				Value:      pulumi.String("<integration-key>"),
//				Name:       pulumi.String("Victorops"),
//				Options: pulumi.StringMap{
//					"rk": pulumi.String("ROUTINGKEY"),
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
//
// </details>
//
// <details>
//
//	<summary>
//	  <b>Slack recipient</b>
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
//			_, err := cloudamqp.NewNotification(ctx, "slack_recipient", &cloudamqp.NotificationArgs{
//				InstanceId: pulumi.Any(instance.Id),
//				Type:       pulumi.String("slack"),
//				Value:      pulumi.String("<slack-webhook-url>"),
//				Name:       pulumi.String("Slack webhook recipient"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// </details>
//
// <details>
//
//	<summary>
//	  <b>Webhook recipient</b>
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
//			_, err := cloudamqp.NewNotification(ctx, "webhook_recipient", &cloudamqp.NotificationArgs{
//				InstanceId: pulumi.Any(instance.Id),
//				Type:       pulumi.String("webhook"),
//				Value:      pulumi.String("<webhook-url>"),
//				Name:       pulumi.String("Webhook"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// </details>
//
// ## Notification Type reference
//
// Valid options for notification type.
//
// * email
// * opsgenie
// * opsgenie-eu
// * pagerduty
// * signl4
// * slack
// * teams
// * victorops
// * webhook
//
// ## Options parameter
//
// <table>
// <thead>
// <tr>
// <th>Type</th>
// <th>Options</th>
// <th>Description</th>
// <th>Note</th>
// </tr>
// </thead>
// <tbody>
// <tr>
// <td>Victorops</td>
// <td>rk</td>
// <td>Routing key to route alarm notification</td>
// <td>-</td>
// </tr>
// <tr>
// <td>PagerDuty</td>
// <td>dedupkey</td>
// <td>Default the dedup key for PagerDuty is generated depending on what alarm has triggered, but here you can set what `dedup` key to use so even if the same alarm is triggered for different resources you only get one notification. Leave blank to use the generated dedup key.</td>
// <td>If multiple alarms are triggered using this recipient, since they all share `dedup` key only the first alarm will be shown in PagerDuty</td>
// </tr>
// </tbody>
// </table>
//
// ## Dependency
//
// This resource depends on CloudAMQP instance identifier, `cloudamqp_instance.instance.id`.
//
// ## Import
//
// `cloudamqp_notification` can be imported using CloudAMQP internal identifier of a recipient together
//
// (CSV separated) with the instance identifier. To retrieve the identifier of a recipient, use
//
// [CloudAMQP API](https://docs.cloudamqp.com/cloudamqp_api.html#list-recipients).
//
// ```sh
// $ pulumi import cloudamqp:index/notification:Notification recipient <id>,<instance_id>`
// ```
type Notification struct {
	pulumi.CustomResourceState

	// The CloudAMQP instance ID.
	InstanceId pulumi.IntOutput `pulumi:"instanceId"`
	// Name of the responder
	Name pulumi.StringOutput `pulumi:"name"`
	// Options argument (e.g. `rk` used for VictorOps routing key).
	Options pulumi.StringMapOutput `pulumi:"options"`
	// An array of reponders (only for OpsGenie). Each `responders` block
	// consists of the field documented below.
	Responders NotificationResponderArrayOutput `pulumi:"responders"`
	// Type of responder. [`team`, `user`, `escalation`, `schedule`]
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
	// Name of the responder
	Name *string `pulumi:"name"`
	// Options argument (e.g. `rk` used for VictorOps routing key).
	Options map[string]string `pulumi:"options"`
	// An array of reponders (only for OpsGenie). Each `responders` block
	// consists of the field documented below.
	Responders []NotificationResponder `pulumi:"responders"`
	// Type of responder. [`team`, `user`, `escalation`, `schedule`]
	Type *string `pulumi:"type"`
	// Integration/API key or endpoint to send the notification.
	Value *string `pulumi:"value"`
}

type NotificationState struct {
	// The CloudAMQP instance ID.
	InstanceId pulumi.IntPtrInput
	// Name of the responder
	Name pulumi.StringPtrInput
	// Options argument (e.g. `rk` used for VictorOps routing key).
	Options pulumi.StringMapInput
	// An array of reponders (only for OpsGenie). Each `responders` block
	// consists of the field documented below.
	Responders NotificationResponderArrayInput
	// Type of responder. [`team`, `user`, `escalation`, `schedule`]
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
	// Name of the responder
	Name *string `pulumi:"name"`
	// Options argument (e.g. `rk` used for VictorOps routing key).
	Options map[string]string `pulumi:"options"`
	// An array of reponders (only for OpsGenie). Each `responders` block
	// consists of the field documented below.
	Responders []NotificationResponder `pulumi:"responders"`
	// Type of responder. [`team`, `user`, `escalation`, `schedule`]
	Type string `pulumi:"type"`
	// Integration/API key or endpoint to send the notification.
	Value string `pulumi:"value"`
}

// The set of arguments for constructing a Notification resource.
type NotificationArgs struct {
	// The CloudAMQP instance ID.
	InstanceId pulumi.IntInput
	// Name of the responder
	Name pulumi.StringPtrInput
	// Options argument (e.g. `rk` used for VictorOps routing key).
	Options pulumi.StringMapInput
	// An array of reponders (only for OpsGenie). Each `responders` block
	// consists of the field documented below.
	Responders NotificationResponderArrayInput
	// Type of responder. [`team`, `user`, `escalation`, `schedule`]
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

// The CloudAMQP instance ID.
func (o NotificationOutput) InstanceId() pulumi.IntOutput {
	return o.ApplyT(func(v *Notification) pulumi.IntOutput { return v.InstanceId }).(pulumi.IntOutput)
}

// Name of the responder
func (o NotificationOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Notification) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Options argument (e.g. `rk` used for VictorOps routing key).
func (o NotificationOutput) Options() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Notification) pulumi.StringMapOutput { return v.Options }).(pulumi.StringMapOutput)
}

// An array of reponders (only for OpsGenie). Each `responders` block
// consists of the field documented below.
func (o NotificationOutput) Responders() NotificationResponderArrayOutput {
	return o.ApplyT(func(v *Notification) NotificationResponderArrayOutput { return v.Responders }).(NotificationResponderArrayOutput)
}

// Type of responder. [`team`, `user`, `escalation`, `schedule`]
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
