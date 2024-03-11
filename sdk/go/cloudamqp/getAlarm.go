// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cloudamqp

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-cloudamqp/sdk/v3/go/cloudamqp/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to retrieve information about default or created alarms. Either use `alarmId` or `type` to retrieve the alarm.
//
// ## Example Usage
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
//			_, err := cloudamqp.LookupAlarm(ctx, &cloudamqp.LookupAlarmArgs{
//				InstanceId: cloudamqp_instance.Instance.Id,
//				Type:       pulumi.StringRef("cpu"),
//			}, nil)
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
// ## Attributes reference
//
// # All attributes reference are computed
//
// * `id`                  - The identifier for this resource.
// * `enabled`             - Enable/disable status of the alarm.
// * `valueThreshold`     - The value threshold that triggers the alarm.
// * `reminderInterval`   - The reminder interval (in seconds) to resend the alarm if not resolved. Set to 0 for no reminders.
// * `timeThreshold`      - The time interval (in seconds) the `valueThreshold` should be active before trigger an alarm.
// * `queueRegex`         - Regular expression for which queue to check.
// * `vhostRegex`         - Regular expression for which vhost to check
// * `recipients`          - Identifier for recipient to be notified.
// * `messageType`        - Message type `(total, unacked, ready)` used by queue alarm type.
//
// # Specific attribute for `disk` alarm
//
// * `valueCalculation`   - Disk value threshold calculation, `(fixed, percentage)` of disk space remaining.
//
// ## Dependency
//
// This data source depends on CloudAMQP instance identifier, `cloudamqp_instance.instance.id`.
//
// ## Alarm types
//
// `cpu, memory, disk, queue, connection, flow, consumer, netsplit, server_unreachable, notice`
func LookupAlarm(ctx *pulumi.Context, args *LookupAlarmArgs, opts ...pulumi.InvokeOption) (*LookupAlarmResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupAlarmResult
	err := ctx.Invoke("cloudamqp:index/getAlarm:getAlarm", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getAlarm.
type LookupAlarmArgs struct {
	// The alarm identifier. Either use this or `type` to give `Alarm` necessary information to retrieve the alarm.
	AlarmId *int `pulumi:"alarmId"`
	// The CloudAMQP instance identifier.
	InstanceId int `pulumi:"instanceId"`
	// The alarm type. Either use this or `alarmId` to give `Alarm` necessary information when retrieve the alarm. Supported alarm types
	Type             *string `pulumi:"type"`
	ValueCalculation *string `pulumi:"valueCalculation"`
}

// A collection of values returned by getAlarm.
type LookupAlarmResult struct {
	AlarmId *int `pulumi:"alarmId"`
	Enabled bool `pulumi:"enabled"`
	// The provider-assigned unique ID for this managed resource.
	Id               string  `pulumi:"id"`
	InstanceId       int     `pulumi:"instanceId"`
	MessageType      string  `pulumi:"messageType"`
	QueueRegex       string  `pulumi:"queueRegex"`
	Recipients       []int   `pulumi:"recipients"`
	ReminderInterval int     `pulumi:"reminderInterval"`
	TimeThreshold    int     `pulumi:"timeThreshold"`
	Type             *string `pulumi:"type"`
	ValueCalculation *string `pulumi:"valueCalculation"`
	ValueThreshold   int     `pulumi:"valueThreshold"`
	VhostRegex       string  `pulumi:"vhostRegex"`
}

func LookupAlarmOutput(ctx *pulumi.Context, args LookupAlarmOutputArgs, opts ...pulumi.InvokeOption) LookupAlarmResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupAlarmResult, error) {
			args := v.(LookupAlarmArgs)
			r, err := LookupAlarm(ctx, &args, opts...)
			var s LookupAlarmResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupAlarmResultOutput)
}

// A collection of arguments for invoking getAlarm.
type LookupAlarmOutputArgs struct {
	// The alarm identifier. Either use this or `type` to give `Alarm` necessary information to retrieve the alarm.
	AlarmId pulumi.IntPtrInput `pulumi:"alarmId"`
	// The CloudAMQP instance identifier.
	InstanceId pulumi.IntInput `pulumi:"instanceId"`
	// The alarm type. Either use this or `alarmId` to give `Alarm` necessary information when retrieve the alarm. Supported alarm types
	Type             pulumi.StringPtrInput `pulumi:"type"`
	ValueCalculation pulumi.StringPtrInput `pulumi:"valueCalculation"`
}

func (LookupAlarmOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAlarmArgs)(nil)).Elem()
}

// A collection of values returned by getAlarm.
type LookupAlarmResultOutput struct{ *pulumi.OutputState }

func (LookupAlarmResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAlarmResult)(nil)).Elem()
}

func (o LookupAlarmResultOutput) ToLookupAlarmResultOutput() LookupAlarmResultOutput {
	return o
}

func (o LookupAlarmResultOutput) ToLookupAlarmResultOutputWithContext(ctx context.Context) LookupAlarmResultOutput {
	return o
}

func (o LookupAlarmResultOutput) AlarmId() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupAlarmResult) *int { return v.AlarmId }).(pulumi.IntPtrOutput)
}

func (o LookupAlarmResultOutput) Enabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupAlarmResult) bool { return v.Enabled }).(pulumi.BoolOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupAlarmResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAlarmResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupAlarmResultOutput) InstanceId() pulumi.IntOutput {
	return o.ApplyT(func(v LookupAlarmResult) int { return v.InstanceId }).(pulumi.IntOutput)
}

func (o LookupAlarmResultOutput) MessageType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAlarmResult) string { return v.MessageType }).(pulumi.StringOutput)
}

func (o LookupAlarmResultOutput) QueueRegex() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAlarmResult) string { return v.QueueRegex }).(pulumi.StringOutput)
}

func (o LookupAlarmResultOutput) Recipients() pulumi.IntArrayOutput {
	return o.ApplyT(func(v LookupAlarmResult) []int { return v.Recipients }).(pulumi.IntArrayOutput)
}

func (o LookupAlarmResultOutput) ReminderInterval() pulumi.IntOutput {
	return o.ApplyT(func(v LookupAlarmResult) int { return v.ReminderInterval }).(pulumi.IntOutput)
}

func (o LookupAlarmResultOutput) TimeThreshold() pulumi.IntOutput {
	return o.ApplyT(func(v LookupAlarmResult) int { return v.TimeThreshold }).(pulumi.IntOutput)
}

func (o LookupAlarmResultOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAlarmResult) *string { return v.Type }).(pulumi.StringPtrOutput)
}

func (o LookupAlarmResultOutput) ValueCalculation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAlarmResult) *string { return v.ValueCalculation }).(pulumi.StringPtrOutput)
}

func (o LookupAlarmResultOutput) ValueThreshold() pulumi.IntOutput {
	return o.ApplyT(func(v LookupAlarmResult) int { return v.ValueThreshold }).(pulumi.IntOutput)
}

func (o LookupAlarmResultOutput) VhostRegex() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAlarmResult) string { return v.VhostRegex }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupAlarmResultOutput{})
}
