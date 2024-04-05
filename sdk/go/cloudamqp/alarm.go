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

// This resource allows you to create and manage alarms to trigger based on a set of conditions. Once triggerd a notification will be sent to the assigned recipients. When creating a new instance, there will also be a set of default alarms (cpu, memory and disk) created. All default alarms uses the default recipient for notifications.
//
// By setting `noDefaultAlarms` to *true* in `Instance`. This will create the instance without default alarms and avoid the need to import them to get full control.
//
// Available for all subscription plans, but `lemur`and `tiger`are limited to fewer alarm types. The limited types supported can be seen in the table below in Alarm Type Reference.
//
// ## Example Usage
//
// <details>
//
//	<summary>
//	  <b>
//	    <i>Basic example of CPU and memory alarm</i>
//	  </b>
//	</summary>
//
// <!--Start PulumiCodeChooser -->
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
//			// New recipient
//			recipient01, err := cloudamqp.NewNotification(ctx, "recipient01", &cloudamqp.NotificationArgs{
//				InstanceId: pulumi.Any(cloudamqp_instance.Instance.Id),
//				Type:       pulumi.String("email"),
//				Value:      pulumi.String("alarm@example.com"),
//			})
//			if err != nil {
//				return err
//			}
//			// New cpu alarm
//			_, err = cloudamqp.NewAlarm(ctx, "cpuAlarm", &cloudamqp.AlarmArgs{
//				InstanceId:       pulumi.Any(cloudamqp_instance.Instance.Id),
//				Type:             pulumi.String("cpu"),
//				Enabled:          pulumi.Bool(true),
//				ReminderInterval: pulumi.Int(600),
//				ValueThreshold:   pulumi.Int(95),
//				TimeThreshold:    pulumi.Int(600),
//				Recipients: pulumi.IntArray{
//					recipient01.ID(),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			// New memory alarm
//			_, err = cloudamqp.NewAlarm(ctx, "memoryAlarm", &cloudamqp.AlarmArgs{
//				InstanceId:       pulumi.Any(cloudamqp_instance.Instance.Id),
//				Type:             pulumi.String("memory"),
//				Enabled:          pulumi.Bool(true),
//				ReminderInterval: pulumi.Int(600),
//				ValueThreshold:   pulumi.Int(95),
//				TimeThreshold:    pulumi.Int(600),
//				Recipients: pulumi.IntArray{
//					recipient01.ID(),
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
// <!--End PulumiCodeChooser -->
//
// </details>
//
// <details>
//
//	<summary>
//	  <b>
//	    <i>Manage notice alarm, available from v1.29.5</i>
//	  </b>
//	</summary>
//
// Only one notice alarm can exists and cannot be created, instead the alarm resource will be updated.
//
// <!--Start PulumiCodeChooser -->
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
//			// New recipient
//			recipient01, err := cloudamqp.NewNotification(ctx, "recipient01", &cloudamqp.NotificationArgs{
//				InstanceId: pulumi.Any(cloudamqp_instance.Instance.Id),
//				Type:       pulumi.String("email"),
//				Value:      pulumi.String("alarm@example.com"),
//			})
//			if err != nil {
//				return err
//			}
//			// Update existing notice alarm
//			_, err = cloudamqp.NewAlarm(ctx, "notice", &cloudamqp.AlarmArgs{
//				InstanceId: pulumi.Any(cloudamqp_instance.Instance.Id),
//				Type:       pulumi.String("notice"),
//				Enabled:    pulumi.Bool(true),
//				Recipients: pulumi.IntArray{
//					recipient01.ID(),
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
// <!--End PulumiCodeChooser -->
//
// </details>
//
// ## Alarm Type reference
//
// Supported alarm types: `cpu, memory, disk, queue, connection, flow, consumer, netsplit, server_unreachable, notice`
//
// Required arguments for all alarms: `instance_id, type, enabled`<br>
// Optional argument for all alarms: `tags, queue_regex, vhostRegex`
//
// | Name | Type | Shared | Dedicated | Required arguments |
// | ---- | ---- | ---- | ---- | ---- |
// | CPU | cpu | - | &#10004; | time_threshold, valueThreshold |
// | Memory | memory | - | &#10004;  | time_threshold, valueThreshold |
// | Disk space | disk | - | &#10004;  | time_threshold, valueThreshold |
// | Queue | queue | &#10004;  | &#10004; | time_threshold, value_threshold, queue_regex, vhost_regex, messageType |
// | Connection | connection | &#10004; | &#10004; | time_threshold, valueThreshold |
// | Connection flow | flow | &#10004; | &#10004; | time_threshold, valueThreshold |
// | Consumer | consumer | &#10004; | &#10004; | time_threshold, value_threshold, queue, vhost |
// | Netsplit | netsplit | - | &#10004; | timeThreshold |
// | Server unreachable | serverUnreachable  | - | &#10004;  | timeThreshold |
// | Notice | notice | &#10004; | &#10004; | |
//
// > Notice alarm is manadatory! Only one can exists and cannot be deleted. Setting `noDefaultAlarm` to true, will still create this alarm. See updated changes to notice alarm below.
//
// ## Dependency
//
// This resource depends on CloudAMQP instance identifier, `cloudamqp_instance.instance.id`.
//
// ## Notice alarm
//
// There is a limitation for notice alarm in the API backend. This alarm is mandatory, multiple
// alarms cannot exists or be deleted.
//
// From provider version v1.29.5
// it's possible to manage the notice alarm and no longer needs to be imported. Just create the
// alarm resource as usually and it will be updated with given recipients. If the alarm is deleted
// it will only be removed from the state file, but will still be enabled in the backend.
//
// ## Import
//
// `cloudamqp_alarm` can be imported using CloudAMQP internal identifier of the alarm together (CSV separated) with the instance identifier. To retrieve the alarm identifier, use [CloudAMQP API](https://docs.cloudamqp.com/cloudamqp_api.html#list-alarms)
//
// ```sh
// $ pulumi import cloudamqp:index/alarm:Alarm alarm <id>,<instance_id>`
// ```
type Alarm struct {
	pulumi.CustomResourceState

	// Enable or disable the alarm to trigger.
	Enabled pulumi.BoolOutput `pulumi:"enabled"`
	// The CloudAMQP instance ID.
	InstanceId pulumi.IntOutput `pulumi:"instanceId"`
	// Message type `(total, unacked, ready)` used by queue alarm type.
	//
	// Specific argument for `disk` alarm
	MessageType pulumi.StringPtrOutput `pulumi:"messageType"`
	// Regex for which queue to check.
	QueueRegex pulumi.StringPtrOutput `pulumi:"queueRegex"`
	// Identifier for recipient to be notified. Leave empty to notify all recipients.
	Recipients pulumi.IntArrayOutput `pulumi:"recipients"`
	// The reminder interval (in seconds) to resend the alarm if not resolved. Set to 0 for no reminders. The Default is 0.
	ReminderInterval pulumi.IntPtrOutput `pulumi:"reminderInterval"`
	// The time interval (in seconds) the `valueThreshold` should be active before triggering an alarm.
	TimeThreshold pulumi.IntPtrOutput `pulumi:"timeThreshold"`
	// The alarm type, see valid options below.
	Type pulumi.StringOutput `pulumi:"type"`
	// Disk value threshold calculation, `fixed, percentage` of disk space remaining.
	//
	// Based on alarm type, different arguments are flagged as required or optional.
	ValueCalculation pulumi.StringPtrOutput `pulumi:"valueCalculation"`
	// The value to trigger the alarm for.
	ValueThreshold pulumi.IntPtrOutput `pulumi:"valueThreshold"`
	// Regex for which vhost to check
	VhostRegex pulumi.StringPtrOutput `pulumi:"vhostRegex"`
}

// NewAlarm registers a new resource with the given unique name, arguments, and options.
func NewAlarm(ctx *pulumi.Context,
	name string, args *AlarmArgs, opts ...pulumi.ResourceOption) (*Alarm, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Enabled == nil {
		return nil, errors.New("invalid value for required argument 'Enabled'")
	}
	if args.InstanceId == nil {
		return nil, errors.New("invalid value for required argument 'InstanceId'")
	}
	if args.Recipients == nil {
		return nil, errors.New("invalid value for required argument 'Recipients'")
	}
	if args.Type == nil {
		return nil, errors.New("invalid value for required argument 'Type'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Alarm
	err := ctx.RegisterResource("cloudamqp:index/alarm:Alarm", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAlarm gets an existing Alarm resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAlarm(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AlarmState, opts ...pulumi.ResourceOption) (*Alarm, error) {
	var resource Alarm
	err := ctx.ReadResource("cloudamqp:index/alarm:Alarm", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Alarm resources.
type alarmState struct {
	// Enable or disable the alarm to trigger.
	Enabled *bool `pulumi:"enabled"`
	// The CloudAMQP instance ID.
	InstanceId *int `pulumi:"instanceId"`
	// Message type `(total, unacked, ready)` used by queue alarm type.
	//
	// Specific argument for `disk` alarm
	MessageType *string `pulumi:"messageType"`
	// Regex for which queue to check.
	QueueRegex *string `pulumi:"queueRegex"`
	// Identifier for recipient to be notified. Leave empty to notify all recipients.
	Recipients []int `pulumi:"recipients"`
	// The reminder interval (in seconds) to resend the alarm if not resolved. Set to 0 for no reminders. The Default is 0.
	ReminderInterval *int `pulumi:"reminderInterval"`
	// The time interval (in seconds) the `valueThreshold` should be active before triggering an alarm.
	TimeThreshold *int `pulumi:"timeThreshold"`
	// The alarm type, see valid options below.
	Type *string `pulumi:"type"`
	// Disk value threshold calculation, `fixed, percentage` of disk space remaining.
	//
	// Based on alarm type, different arguments are flagged as required or optional.
	ValueCalculation *string `pulumi:"valueCalculation"`
	// The value to trigger the alarm for.
	ValueThreshold *int `pulumi:"valueThreshold"`
	// Regex for which vhost to check
	VhostRegex *string `pulumi:"vhostRegex"`
}

type AlarmState struct {
	// Enable or disable the alarm to trigger.
	Enabled pulumi.BoolPtrInput
	// The CloudAMQP instance ID.
	InstanceId pulumi.IntPtrInput
	// Message type `(total, unacked, ready)` used by queue alarm type.
	//
	// Specific argument for `disk` alarm
	MessageType pulumi.StringPtrInput
	// Regex for which queue to check.
	QueueRegex pulumi.StringPtrInput
	// Identifier for recipient to be notified. Leave empty to notify all recipients.
	Recipients pulumi.IntArrayInput
	// The reminder interval (in seconds) to resend the alarm if not resolved. Set to 0 for no reminders. The Default is 0.
	ReminderInterval pulumi.IntPtrInput
	// The time interval (in seconds) the `valueThreshold` should be active before triggering an alarm.
	TimeThreshold pulumi.IntPtrInput
	// The alarm type, see valid options below.
	Type pulumi.StringPtrInput
	// Disk value threshold calculation, `fixed, percentage` of disk space remaining.
	//
	// Based on alarm type, different arguments are flagged as required or optional.
	ValueCalculation pulumi.StringPtrInput
	// The value to trigger the alarm for.
	ValueThreshold pulumi.IntPtrInput
	// Regex for which vhost to check
	VhostRegex pulumi.StringPtrInput
}

func (AlarmState) ElementType() reflect.Type {
	return reflect.TypeOf((*alarmState)(nil)).Elem()
}

type alarmArgs struct {
	// Enable or disable the alarm to trigger.
	Enabled bool `pulumi:"enabled"`
	// The CloudAMQP instance ID.
	InstanceId int `pulumi:"instanceId"`
	// Message type `(total, unacked, ready)` used by queue alarm type.
	//
	// Specific argument for `disk` alarm
	MessageType *string `pulumi:"messageType"`
	// Regex for which queue to check.
	QueueRegex *string `pulumi:"queueRegex"`
	// Identifier for recipient to be notified. Leave empty to notify all recipients.
	Recipients []int `pulumi:"recipients"`
	// The reminder interval (in seconds) to resend the alarm if not resolved. Set to 0 for no reminders. The Default is 0.
	ReminderInterval *int `pulumi:"reminderInterval"`
	// The time interval (in seconds) the `valueThreshold` should be active before triggering an alarm.
	TimeThreshold *int `pulumi:"timeThreshold"`
	// The alarm type, see valid options below.
	Type string `pulumi:"type"`
	// Disk value threshold calculation, `fixed, percentage` of disk space remaining.
	//
	// Based on alarm type, different arguments are flagged as required or optional.
	ValueCalculation *string `pulumi:"valueCalculation"`
	// The value to trigger the alarm for.
	ValueThreshold *int `pulumi:"valueThreshold"`
	// Regex for which vhost to check
	VhostRegex *string `pulumi:"vhostRegex"`
}

// The set of arguments for constructing a Alarm resource.
type AlarmArgs struct {
	// Enable or disable the alarm to trigger.
	Enabled pulumi.BoolInput
	// The CloudAMQP instance ID.
	InstanceId pulumi.IntInput
	// Message type `(total, unacked, ready)` used by queue alarm type.
	//
	// Specific argument for `disk` alarm
	MessageType pulumi.StringPtrInput
	// Regex for which queue to check.
	QueueRegex pulumi.StringPtrInput
	// Identifier for recipient to be notified. Leave empty to notify all recipients.
	Recipients pulumi.IntArrayInput
	// The reminder interval (in seconds) to resend the alarm if not resolved. Set to 0 for no reminders. The Default is 0.
	ReminderInterval pulumi.IntPtrInput
	// The time interval (in seconds) the `valueThreshold` should be active before triggering an alarm.
	TimeThreshold pulumi.IntPtrInput
	// The alarm type, see valid options below.
	Type pulumi.StringInput
	// Disk value threshold calculation, `fixed, percentage` of disk space remaining.
	//
	// Based on alarm type, different arguments are flagged as required or optional.
	ValueCalculation pulumi.StringPtrInput
	// The value to trigger the alarm for.
	ValueThreshold pulumi.IntPtrInput
	// Regex for which vhost to check
	VhostRegex pulumi.StringPtrInput
}

func (AlarmArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*alarmArgs)(nil)).Elem()
}

type AlarmInput interface {
	pulumi.Input

	ToAlarmOutput() AlarmOutput
	ToAlarmOutputWithContext(ctx context.Context) AlarmOutput
}

func (*Alarm) ElementType() reflect.Type {
	return reflect.TypeOf((**Alarm)(nil)).Elem()
}

func (i *Alarm) ToAlarmOutput() AlarmOutput {
	return i.ToAlarmOutputWithContext(context.Background())
}

func (i *Alarm) ToAlarmOutputWithContext(ctx context.Context) AlarmOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AlarmOutput)
}

// AlarmArrayInput is an input type that accepts AlarmArray and AlarmArrayOutput values.
// You can construct a concrete instance of `AlarmArrayInput` via:
//
//	AlarmArray{ AlarmArgs{...} }
type AlarmArrayInput interface {
	pulumi.Input

	ToAlarmArrayOutput() AlarmArrayOutput
	ToAlarmArrayOutputWithContext(context.Context) AlarmArrayOutput
}

type AlarmArray []AlarmInput

func (AlarmArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Alarm)(nil)).Elem()
}

func (i AlarmArray) ToAlarmArrayOutput() AlarmArrayOutput {
	return i.ToAlarmArrayOutputWithContext(context.Background())
}

func (i AlarmArray) ToAlarmArrayOutputWithContext(ctx context.Context) AlarmArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AlarmArrayOutput)
}

// AlarmMapInput is an input type that accepts AlarmMap and AlarmMapOutput values.
// You can construct a concrete instance of `AlarmMapInput` via:
//
//	AlarmMap{ "key": AlarmArgs{...} }
type AlarmMapInput interface {
	pulumi.Input

	ToAlarmMapOutput() AlarmMapOutput
	ToAlarmMapOutputWithContext(context.Context) AlarmMapOutput
}

type AlarmMap map[string]AlarmInput

func (AlarmMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Alarm)(nil)).Elem()
}

func (i AlarmMap) ToAlarmMapOutput() AlarmMapOutput {
	return i.ToAlarmMapOutputWithContext(context.Background())
}

func (i AlarmMap) ToAlarmMapOutputWithContext(ctx context.Context) AlarmMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AlarmMapOutput)
}

type AlarmOutput struct{ *pulumi.OutputState }

func (AlarmOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Alarm)(nil)).Elem()
}

func (o AlarmOutput) ToAlarmOutput() AlarmOutput {
	return o
}

func (o AlarmOutput) ToAlarmOutputWithContext(ctx context.Context) AlarmOutput {
	return o
}

// Enable or disable the alarm to trigger.
func (o AlarmOutput) Enabled() pulumi.BoolOutput {
	return o.ApplyT(func(v *Alarm) pulumi.BoolOutput { return v.Enabled }).(pulumi.BoolOutput)
}

// The CloudAMQP instance ID.
func (o AlarmOutput) InstanceId() pulumi.IntOutput {
	return o.ApplyT(func(v *Alarm) pulumi.IntOutput { return v.InstanceId }).(pulumi.IntOutput)
}

// Message type `(total, unacked, ready)` used by queue alarm type.
//
// Specific argument for `disk` alarm
func (o AlarmOutput) MessageType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Alarm) pulumi.StringPtrOutput { return v.MessageType }).(pulumi.StringPtrOutput)
}

// Regex for which queue to check.
func (o AlarmOutput) QueueRegex() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Alarm) pulumi.StringPtrOutput { return v.QueueRegex }).(pulumi.StringPtrOutput)
}

// Identifier for recipient to be notified. Leave empty to notify all recipients.
func (o AlarmOutput) Recipients() pulumi.IntArrayOutput {
	return o.ApplyT(func(v *Alarm) pulumi.IntArrayOutput { return v.Recipients }).(pulumi.IntArrayOutput)
}

// The reminder interval (in seconds) to resend the alarm if not resolved. Set to 0 for no reminders. The Default is 0.
func (o AlarmOutput) ReminderInterval() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Alarm) pulumi.IntPtrOutput { return v.ReminderInterval }).(pulumi.IntPtrOutput)
}

// The time interval (in seconds) the `valueThreshold` should be active before triggering an alarm.
func (o AlarmOutput) TimeThreshold() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Alarm) pulumi.IntPtrOutput { return v.TimeThreshold }).(pulumi.IntPtrOutput)
}

// The alarm type, see valid options below.
func (o AlarmOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Alarm) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// Disk value threshold calculation, `fixed, percentage` of disk space remaining.
//
// Based on alarm type, different arguments are flagged as required or optional.
func (o AlarmOutput) ValueCalculation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Alarm) pulumi.StringPtrOutput { return v.ValueCalculation }).(pulumi.StringPtrOutput)
}

// The value to trigger the alarm for.
func (o AlarmOutput) ValueThreshold() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Alarm) pulumi.IntPtrOutput { return v.ValueThreshold }).(pulumi.IntPtrOutput)
}

// Regex for which vhost to check
func (o AlarmOutput) VhostRegex() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Alarm) pulumi.StringPtrOutput { return v.VhostRegex }).(pulumi.StringPtrOutput)
}

type AlarmArrayOutput struct{ *pulumi.OutputState }

func (AlarmArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Alarm)(nil)).Elem()
}

func (o AlarmArrayOutput) ToAlarmArrayOutput() AlarmArrayOutput {
	return o
}

func (o AlarmArrayOutput) ToAlarmArrayOutputWithContext(ctx context.Context) AlarmArrayOutput {
	return o
}

func (o AlarmArrayOutput) Index(i pulumi.IntInput) AlarmOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Alarm {
		return vs[0].([]*Alarm)[vs[1].(int)]
	}).(AlarmOutput)
}

type AlarmMapOutput struct{ *pulumi.OutputState }

func (AlarmMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Alarm)(nil)).Elem()
}

func (o AlarmMapOutput) ToAlarmMapOutput() AlarmMapOutput {
	return o
}

func (o AlarmMapOutput) ToAlarmMapOutputWithContext(ctx context.Context) AlarmMapOutput {
	return o
}

func (o AlarmMapOutput) MapIndex(k pulumi.StringInput) AlarmOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Alarm {
		return vs[0].(map[string]*Alarm)[vs[1].(string)]
	}).(AlarmOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AlarmInput)(nil)).Elem(), &Alarm{})
	pulumi.RegisterInputType(reflect.TypeOf((*AlarmArrayInput)(nil)).Elem(), AlarmArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AlarmMapInput)(nil)).Elem(), AlarmMap{})
	pulumi.RegisterOutputType(AlarmOutput{})
	pulumi.RegisterOutputType(AlarmArrayOutput{})
	pulumi.RegisterOutputType(AlarmMapOutput{})
}
