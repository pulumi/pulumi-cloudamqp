// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package cloudamqp

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This resource allows you to create and manage alarms to trigger based on a set of conditions. Once triggerd a notification will be sent to the assigned recipients. When creating a new instance, there will also be a set of default alarms (cpu, memory and disk) created. All default alarms uses the default recipient for notifications.
//
// By setting `noDefaultAlarms` to *true* in `Instance`. This will create the instance without default alarms and avoid the need to import them to get full control.
//
// Available for all subscription plans, but `lemur`and `tiger`are limited to fewer alarm types. The limited types supported can be seen in the table below in Alarm Type Reference.
//
// ## Alarm Type reference
//
// Valid options for notification type.
//
// Required arguments for all alarms: *instance_id*, *type* and *enabled*
// Optional argument for all alarms: *tags*, *queue_regex*, *vhost_regex*
//
// | Name | Type | Shared | Dedicated | Required arguments |
// | ---- | ---- | ---- | ---- | ---- |
// | CPU | cpu | - | &#10004; | time_threshold, valueThreshold |
// | Memory | memory | - | &#10004;  | time_threshold, valueThreshold |
// | Disk space | disk | - | &#10004;  | time_threshold, valueThreshold |
// | Queue | queue | &#10004;  | &#10004;  | time_threshold, value_threshold, queue_regex, vhost_regex, messageType |
// | Connection | connection | &#10004; | &#10004; | time_threshold, valueThreshold |
// | Consumer | consumer | &#10004; | &#10004; | time_threshold, value_threshold, queue, vhost |
// | Netsplit | netsplit | - | &#10004; | timeThreshold |
// | Server unreachable | serverUnreachable  | - | &#10004;  | timeThreshold |
// | Notice | notice | &#10004; | &#10004; | |
//
// ## Dependency
//
// This resource depends on CloudAMQP instance identifier, `cloudamqp_instance.instance.id`.
//
// ## Import
//
// `cloudamqp_alarm` can be imported using CloudAMQP internal identifier of the alarm together (CSV separated) with the instance identifier. To retrieve the alarm identifier, use [CloudAMQP API](https://docs.cloudamqp.com/cloudamqp_api.html#list-alarms)
//
// ```sh
//  $ pulumi import cloudamqp:index/alarm:Alarm alarm <id>,<instance_id>`
// ```
type Alarm struct {
	pulumi.CustomResourceState

	// Enable or disable the alarm to trigger.
	Enabled pulumi.BoolOutput `pulumi:"enabled"`
	// The CloudAMQP instance ID.
	InstanceId pulumi.IntOutput `pulumi:"instanceId"`
	// Message type `(total, unacked, ready)` used by queue alarm type.
	MessageType pulumi.StringPtrOutput `pulumi:"messageType"`
	// Regex for which queue to check.
	QueueRegex pulumi.StringPtrOutput `pulumi:"queueRegex"`
	// Identifier for recipient to be notified. Leave empty to notify all recipients.
	Recipients pulumi.IntArrayOutput `pulumi:"recipients"`
	// The time interval (in seconds) the `valueThreshold` should be active before triggering an alarm.
	TimeThreshold pulumi.IntPtrOutput `pulumi:"timeThreshold"`
	// The alarm type, see valid options below.
	Type pulumi.StringOutput `pulumi:"type"`
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
	MessageType *string `pulumi:"messageType"`
	// Regex for which queue to check.
	QueueRegex *string `pulumi:"queueRegex"`
	// Identifier for recipient to be notified. Leave empty to notify all recipients.
	Recipients []int `pulumi:"recipients"`
	// The time interval (in seconds) the `valueThreshold` should be active before triggering an alarm.
	TimeThreshold *int `pulumi:"timeThreshold"`
	// The alarm type, see valid options below.
	Type *string `pulumi:"type"`
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
	MessageType pulumi.StringPtrInput
	// Regex for which queue to check.
	QueueRegex pulumi.StringPtrInput
	// Identifier for recipient to be notified. Leave empty to notify all recipients.
	Recipients pulumi.IntArrayInput
	// The time interval (in seconds) the `valueThreshold` should be active before triggering an alarm.
	TimeThreshold pulumi.IntPtrInput
	// The alarm type, see valid options below.
	Type pulumi.StringPtrInput
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
	MessageType *string `pulumi:"messageType"`
	// Regex for which queue to check.
	QueueRegex *string `pulumi:"queueRegex"`
	// Identifier for recipient to be notified. Leave empty to notify all recipients.
	Recipients []int `pulumi:"recipients"`
	// The time interval (in seconds) the `valueThreshold` should be active before triggering an alarm.
	TimeThreshold *int `pulumi:"timeThreshold"`
	// The alarm type, see valid options below.
	Type string `pulumi:"type"`
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
	MessageType pulumi.StringPtrInput
	// Regex for which queue to check.
	QueueRegex pulumi.StringPtrInput
	// Identifier for recipient to be notified. Leave empty to notify all recipients.
	Recipients pulumi.IntArrayInput
	// The time interval (in seconds) the `valueThreshold` should be active before triggering an alarm.
	TimeThreshold pulumi.IntPtrInput
	// The alarm type, see valid options below.
	Type pulumi.StringInput
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
	return reflect.TypeOf((*Alarm)(nil))
}

func (i *Alarm) ToAlarmOutput() AlarmOutput {
	return i.ToAlarmOutputWithContext(context.Background())
}

func (i *Alarm) ToAlarmOutputWithContext(ctx context.Context) AlarmOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AlarmOutput)
}

func (i *Alarm) ToAlarmPtrOutput() AlarmPtrOutput {
	return i.ToAlarmPtrOutputWithContext(context.Background())
}

func (i *Alarm) ToAlarmPtrOutputWithContext(ctx context.Context) AlarmPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AlarmPtrOutput)
}

type AlarmPtrInput interface {
	pulumi.Input

	ToAlarmPtrOutput() AlarmPtrOutput
	ToAlarmPtrOutputWithContext(ctx context.Context) AlarmPtrOutput
}

type alarmPtrType AlarmArgs

func (*alarmPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Alarm)(nil))
}

func (i *alarmPtrType) ToAlarmPtrOutput() AlarmPtrOutput {
	return i.ToAlarmPtrOutputWithContext(context.Background())
}

func (i *alarmPtrType) ToAlarmPtrOutputWithContext(ctx context.Context) AlarmPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AlarmPtrOutput)
}

// AlarmArrayInput is an input type that accepts AlarmArray and AlarmArrayOutput values.
// You can construct a concrete instance of `AlarmArrayInput` via:
//
//          AlarmArray{ AlarmArgs{...} }
type AlarmArrayInput interface {
	pulumi.Input

	ToAlarmArrayOutput() AlarmArrayOutput
	ToAlarmArrayOutputWithContext(context.Context) AlarmArrayOutput
}

type AlarmArray []AlarmInput

func (AlarmArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*Alarm)(nil))
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
//          AlarmMap{ "key": AlarmArgs{...} }
type AlarmMapInput interface {
	pulumi.Input

	ToAlarmMapOutput() AlarmMapOutput
	ToAlarmMapOutputWithContext(context.Context) AlarmMapOutput
}

type AlarmMap map[string]AlarmInput

func (AlarmMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*Alarm)(nil))
}

func (i AlarmMap) ToAlarmMapOutput() AlarmMapOutput {
	return i.ToAlarmMapOutputWithContext(context.Background())
}

func (i AlarmMap) ToAlarmMapOutputWithContext(ctx context.Context) AlarmMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AlarmMapOutput)
}

type AlarmOutput struct {
	*pulumi.OutputState
}

func (AlarmOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Alarm)(nil))
}

func (o AlarmOutput) ToAlarmOutput() AlarmOutput {
	return o
}

func (o AlarmOutput) ToAlarmOutputWithContext(ctx context.Context) AlarmOutput {
	return o
}

func (o AlarmOutput) ToAlarmPtrOutput() AlarmPtrOutput {
	return o.ToAlarmPtrOutputWithContext(context.Background())
}

func (o AlarmOutput) ToAlarmPtrOutputWithContext(ctx context.Context) AlarmPtrOutput {
	return o.ApplyT(func(v Alarm) *Alarm {
		return &v
	}).(AlarmPtrOutput)
}

type AlarmPtrOutput struct {
	*pulumi.OutputState
}

func (AlarmPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Alarm)(nil))
}

func (o AlarmPtrOutput) ToAlarmPtrOutput() AlarmPtrOutput {
	return o
}

func (o AlarmPtrOutput) ToAlarmPtrOutputWithContext(ctx context.Context) AlarmPtrOutput {
	return o
}

type AlarmArrayOutput struct{ *pulumi.OutputState }

func (AlarmArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Alarm)(nil))
}

func (o AlarmArrayOutput) ToAlarmArrayOutput() AlarmArrayOutput {
	return o
}

func (o AlarmArrayOutput) ToAlarmArrayOutputWithContext(ctx context.Context) AlarmArrayOutput {
	return o
}

func (o AlarmArrayOutput) Index(i pulumi.IntInput) AlarmOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Alarm {
		return vs[0].([]Alarm)[vs[1].(int)]
	}).(AlarmOutput)
}

type AlarmMapOutput struct{ *pulumi.OutputState }

func (AlarmMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]Alarm)(nil))
}

func (o AlarmMapOutput) ToAlarmMapOutput() AlarmMapOutput {
	return o
}

func (o AlarmMapOutput) ToAlarmMapOutputWithContext(ctx context.Context) AlarmMapOutput {
	return o
}

func (o AlarmMapOutput) MapIndex(k pulumi.StringInput) AlarmOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) Alarm {
		return vs[0].(map[string]Alarm)[vs[1].(string)]
	}).(AlarmOutput)
}

func init() {
	pulumi.RegisterOutputType(AlarmOutput{})
	pulumi.RegisterOutputType(AlarmPtrOutput{})
	pulumi.RegisterOutputType(AlarmArrayOutput{})
	pulumi.RegisterOutputType(AlarmMapOutput{})
}
