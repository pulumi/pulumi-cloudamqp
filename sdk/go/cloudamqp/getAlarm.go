// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package cloudamqp

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Use this data source to retrieve information about default or created alarms. Either use `alarmId` or `type` to retrieve the alarm.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-cloudamqp/sdk/v2/go/cloudamqp"
// 	"github.com/pulumi/pulumi-cloudamqp/sdk/v2/go/cloudamqp/"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		opt0 := "cpu"
// 		_, err := cloudamqp.LookupAlarm(ctx, &cloudamqp.LookupAlarmArgs{
// 			InstanceId: cloudamqp_instance.Instance.Id,
// 			Type:       &opt0,
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
// * `alarmId`    - (Optional) The alarm identifier. Either use this or `type` to give `Alarm` necessary information to retrieve the alarm.
// * `type`        - (Optional) The alarm type. Either use this or `alarmId` to give `Alarm` necessary information when retrieve the alarm.
//
// ## Attribute reference
//
// * `enabled`         - (Computed) Enable/disable status of the alarm.
// * `valueThreshold` - (Computed) The value threshold that triggers the alarm.
// * `timeThreshold`  - (Computed) The time interval (in seconds) the `valueThreshold` should be active before trigger an alarm.
// * `queueRegex`     - (Computed) Regular expression for which queue to check.
// * `vhostRegex`     - (Computed) Regular expression for which vhost to check
// * `recipients`      - (Computed) Identifier for recipient to be notified.
// * `messageType`    - (Computed) Message type `(total, unacked, ready)` used by queue alarm type.
//
// ## Dependency
//
// This data source depends on CloudAMQP instance identifier, `cloudamqp_instance.instance.id`.
func LookupAlarm(ctx *pulumi.Context, args *LookupAlarmArgs, opts ...pulumi.InvokeOption) (*LookupAlarmResult, error) {
	var rv LookupAlarmResult
	err := ctx.Invoke("cloudamqp:index/getAlarm:getAlarm", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getAlarm.
type LookupAlarmArgs struct {
	AlarmId    *int    `pulumi:"alarmId"`
	InstanceId int     `pulumi:"instanceId"`
	Type       *string `pulumi:"type"`
}

// A collection of values returned by getAlarm.
type LookupAlarmResult struct {
	AlarmId *int `pulumi:"alarmId"`
	Enabled bool `pulumi:"enabled"`
	// The provider-assigned unique ID for this managed resource.
	Id             string  `pulumi:"id"`
	InstanceId     int     `pulumi:"instanceId"`
	MessageType    string  `pulumi:"messageType"`
	QueueRegex     string  `pulumi:"queueRegex"`
	Recipients     []int   `pulumi:"recipients"`
	TimeThreshold  int     `pulumi:"timeThreshold"`
	Type           *string `pulumi:"type"`
	ValueThreshold int     `pulumi:"valueThreshold"`
	VhostRegex     string  `pulumi:"vhostRegex"`
}
