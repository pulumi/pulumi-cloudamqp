// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cloudamqp

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This resource allows you to create and manage, an [AWS EventBridge](https://aws.amazon.com/eventbridge/) for a CloudAMQP instance. Once created, continue to map the EventBridge in the [AWS Eventbridge console](https://console.aws.amazon.com/events/home).
//
// >  Our consumer needs to have exclusive usage to the configured queue and the maximum body size allowed on msgs by AWS is 256kb. The message body has to be valid JSON for AWS Eventbridge to accept it. If messages are too large or are not valid JSON, they will be rejected (tip: setup a dead-letter queue to catch them).
//
// Not possible to update this resource. Any changes made to the argument will destroy and recreate the resource. Hence why all arguments use ForceNew.
//
// Only available for dedicated subscription plans.
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
//			instance, err := cloudamqp.NewInstance(ctx, "instance", &cloudamqp.InstanceArgs{
//				Plan:       pulumi.String("squirrel-1"),
//				Region:     pulumi.String("amazon-web-services::us-west-1"),
//				RmqVersion: pulumi.String("3.11.5"),
//				Tags: pulumi.StringArray{
//					pulumi.String("aws"),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			_, err = cloudamqp.NewIntegrationAwsEventbridge(ctx, "awsEventbridge", &cloudamqp.IntegrationAwsEventbridgeArgs{
//				InstanceId:   instance.ID(),
//				Vhost:        instance.Vhost,
//				Queue:        pulumi.String("<QUEUE-NAME>"),
//				AwsAccountId: pulumi.String("<AWS-ACCOUNT-ID>"),
//				AwsRegion:    pulumi.String("us-west-1"),
//				WithHeaders:  pulumi.Bool(true),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// ## Argument references
//
// The following arguments are supported:
//
// * `awsAccountId` - (ForceNew/Required) The 12 digit AWS Account ID where you want the events to be sent to.
// * `awsRegion`- (ForceNew/Required) The AWS region where you the events to be sent to. (e.g. us-west-1, us-west-2, ..., etc.)
// * `vhost`- (ForceNew/Required) The VHost the queue resides in.
// * `queue` - (ForceNew/Required) A (durable) queue on your RabbitMQ instance.
// * `withHeaders` - (ForceNew/Required) Include message headers in the event data. `({ "headers": { }, "body": { "your": "message" } })`
//
// ## Dependency
//
// This resource depends on CloudAMQP instance identifier, `cloudamqp_instance.instance.id`.
//
// ## Import
//
// `cloudamqp_integration_aws_eventbridge` can be imported using CloudAMQP internal identifier of the AWS EventBridge together (CSV separated) with the instance identifier. To retrieve the AWS EventBridge identifier, use [CloudAMQP API](https://docs.cloudamqp.com/cloudamqp_api.html#list-eventbridges)
//
// ```sh
//
//	$ pulumi import cloudamqp:index/integrationAwsEventbridge:IntegrationAwsEventbridge aws_eventbridge <id>,<instance_id>`
//
// ```
type IntegrationAwsEventbridge struct {
	pulumi.CustomResourceState

	// The 12 digit AWS Account ID where you want the events to be sent to.
	AwsAccountId pulumi.StringOutput `pulumi:"awsAccountId"`
	// The AWS region where you the events to be sent to. (e.g. us-west-1, us-west-2, ..., etc.)
	AwsRegion pulumi.StringOutput `pulumi:"awsRegion"`
	// Instance identifier
	InstanceId pulumi.IntOutput `pulumi:"instanceId"`
	// A (durable) queue on your RabbitMQ instance.
	Queue pulumi.StringOutput `pulumi:"queue"`
	// Always set to null, unless there is an error starting the EventBridge.
	Status pulumi.StringOutput `pulumi:"status"`
	// The VHost the queue resides in.
	Vhost pulumi.StringOutput `pulumi:"vhost"`
	// Include message headers in the event data.
	WithHeaders pulumi.BoolOutput `pulumi:"withHeaders"`
}

// NewIntegrationAwsEventbridge registers a new resource with the given unique name, arguments, and options.
func NewIntegrationAwsEventbridge(ctx *pulumi.Context,
	name string, args *IntegrationAwsEventbridgeArgs, opts ...pulumi.ResourceOption) (*IntegrationAwsEventbridge, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AwsAccountId == nil {
		return nil, errors.New("invalid value for required argument 'AwsAccountId'")
	}
	if args.AwsRegion == nil {
		return nil, errors.New("invalid value for required argument 'AwsRegion'")
	}
	if args.InstanceId == nil {
		return nil, errors.New("invalid value for required argument 'InstanceId'")
	}
	if args.Queue == nil {
		return nil, errors.New("invalid value for required argument 'Queue'")
	}
	if args.Vhost == nil {
		return nil, errors.New("invalid value for required argument 'Vhost'")
	}
	if args.WithHeaders == nil {
		return nil, errors.New("invalid value for required argument 'WithHeaders'")
	}
	var resource IntegrationAwsEventbridge
	err := ctx.RegisterResource("cloudamqp:index/integrationAwsEventbridge:IntegrationAwsEventbridge", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIntegrationAwsEventbridge gets an existing IntegrationAwsEventbridge resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIntegrationAwsEventbridge(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IntegrationAwsEventbridgeState, opts ...pulumi.ResourceOption) (*IntegrationAwsEventbridge, error) {
	var resource IntegrationAwsEventbridge
	err := ctx.ReadResource("cloudamqp:index/integrationAwsEventbridge:IntegrationAwsEventbridge", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering IntegrationAwsEventbridge resources.
type integrationAwsEventbridgeState struct {
	// The 12 digit AWS Account ID where you want the events to be sent to.
	AwsAccountId *string `pulumi:"awsAccountId"`
	// The AWS region where you the events to be sent to. (e.g. us-west-1, us-west-2, ..., etc.)
	AwsRegion *string `pulumi:"awsRegion"`
	// Instance identifier
	InstanceId *int `pulumi:"instanceId"`
	// A (durable) queue on your RabbitMQ instance.
	Queue *string `pulumi:"queue"`
	// Always set to null, unless there is an error starting the EventBridge.
	Status *string `pulumi:"status"`
	// The VHost the queue resides in.
	Vhost *string `pulumi:"vhost"`
	// Include message headers in the event data.
	WithHeaders *bool `pulumi:"withHeaders"`
}

type IntegrationAwsEventbridgeState struct {
	// The 12 digit AWS Account ID where you want the events to be sent to.
	AwsAccountId pulumi.StringPtrInput
	// The AWS region where you the events to be sent to. (e.g. us-west-1, us-west-2, ..., etc.)
	AwsRegion pulumi.StringPtrInput
	// Instance identifier
	InstanceId pulumi.IntPtrInput
	// A (durable) queue on your RabbitMQ instance.
	Queue pulumi.StringPtrInput
	// Always set to null, unless there is an error starting the EventBridge.
	Status pulumi.StringPtrInput
	// The VHost the queue resides in.
	Vhost pulumi.StringPtrInput
	// Include message headers in the event data.
	WithHeaders pulumi.BoolPtrInput
}

func (IntegrationAwsEventbridgeState) ElementType() reflect.Type {
	return reflect.TypeOf((*integrationAwsEventbridgeState)(nil)).Elem()
}

type integrationAwsEventbridgeArgs struct {
	// The 12 digit AWS Account ID where you want the events to be sent to.
	AwsAccountId string `pulumi:"awsAccountId"`
	// The AWS region where you the events to be sent to. (e.g. us-west-1, us-west-2, ..., etc.)
	AwsRegion string `pulumi:"awsRegion"`
	// Instance identifier
	InstanceId int `pulumi:"instanceId"`
	// A (durable) queue on your RabbitMQ instance.
	Queue string `pulumi:"queue"`
	// The VHost the queue resides in.
	Vhost string `pulumi:"vhost"`
	// Include message headers in the event data.
	WithHeaders bool `pulumi:"withHeaders"`
}

// The set of arguments for constructing a IntegrationAwsEventbridge resource.
type IntegrationAwsEventbridgeArgs struct {
	// The 12 digit AWS Account ID where you want the events to be sent to.
	AwsAccountId pulumi.StringInput
	// The AWS region where you the events to be sent to. (e.g. us-west-1, us-west-2, ..., etc.)
	AwsRegion pulumi.StringInput
	// Instance identifier
	InstanceId pulumi.IntInput
	// A (durable) queue on your RabbitMQ instance.
	Queue pulumi.StringInput
	// The VHost the queue resides in.
	Vhost pulumi.StringInput
	// Include message headers in the event data.
	WithHeaders pulumi.BoolInput
}

func (IntegrationAwsEventbridgeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*integrationAwsEventbridgeArgs)(nil)).Elem()
}

type IntegrationAwsEventbridgeInput interface {
	pulumi.Input

	ToIntegrationAwsEventbridgeOutput() IntegrationAwsEventbridgeOutput
	ToIntegrationAwsEventbridgeOutputWithContext(ctx context.Context) IntegrationAwsEventbridgeOutput
}

func (*IntegrationAwsEventbridge) ElementType() reflect.Type {
	return reflect.TypeOf((**IntegrationAwsEventbridge)(nil)).Elem()
}

func (i *IntegrationAwsEventbridge) ToIntegrationAwsEventbridgeOutput() IntegrationAwsEventbridgeOutput {
	return i.ToIntegrationAwsEventbridgeOutputWithContext(context.Background())
}

func (i *IntegrationAwsEventbridge) ToIntegrationAwsEventbridgeOutputWithContext(ctx context.Context) IntegrationAwsEventbridgeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IntegrationAwsEventbridgeOutput)
}

// IntegrationAwsEventbridgeArrayInput is an input type that accepts IntegrationAwsEventbridgeArray and IntegrationAwsEventbridgeArrayOutput values.
// You can construct a concrete instance of `IntegrationAwsEventbridgeArrayInput` via:
//
//	IntegrationAwsEventbridgeArray{ IntegrationAwsEventbridgeArgs{...} }
type IntegrationAwsEventbridgeArrayInput interface {
	pulumi.Input

	ToIntegrationAwsEventbridgeArrayOutput() IntegrationAwsEventbridgeArrayOutput
	ToIntegrationAwsEventbridgeArrayOutputWithContext(context.Context) IntegrationAwsEventbridgeArrayOutput
}

type IntegrationAwsEventbridgeArray []IntegrationAwsEventbridgeInput

func (IntegrationAwsEventbridgeArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*IntegrationAwsEventbridge)(nil)).Elem()
}

func (i IntegrationAwsEventbridgeArray) ToIntegrationAwsEventbridgeArrayOutput() IntegrationAwsEventbridgeArrayOutput {
	return i.ToIntegrationAwsEventbridgeArrayOutputWithContext(context.Background())
}

func (i IntegrationAwsEventbridgeArray) ToIntegrationAwsEventbridgeArrayOutputWithContext(ctx context.Context) IntegrationAwsEventbridgeArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IntegrationAwsEventbridgeArrayOutput)
}

// IntegrationAwsEventbridgeMapInput is an input type that accepts IntegrationAwsEventbridgeMap and IntegrationAwsEventbridgeMapOutput values.
// You can construct a concrete instance of `IntegrationAwsEventbridgeMapInput` via:
//
//	IntegrationAwsEventbridgeMap{ "key": IntegrationAwsEventbridgeArgs{...} }
type IntegrationAwsEventbridgeMapInput interface {
	pulumi.Input

	ToIntegrationAwsEventbridgeMapOutput() IntegrationAwsEventbridgeMapOutput
	ToIntegrationAwsEventbridgeMapOutputWithContext(context.Context) IntegrationAwsEventbridgeMapOutput
}

type IntegrationAwsEventbridgeMap map[string]IntegrationAwsEventbridgeInput

func (IntegrationAwsEventbridgeMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*IntegrationAwsEventbridge)(nil)).Elem()
}

func (i IntegrationAwsEventbridgeMap) ToIntegrationAwsEventbridgeMapOutput() IntegrationAwsEventbridgeMapOutput {
	return i.ToIntegrationAwsEventbridgeMapOutputWithContext(context.Background())
}

func (i IntegrationAwsEventbridgeMap) ToIntegrationAwsEventbridgeMapOutputWithContext(ctx context.Context) IntegrationAwsEventbridgeMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IntegrationAwsEventbridgeMapOutput)
}

type IntegrationAwsEventbridgeOutput struct{ *pulumi.OutputState }

func (IntegrationAwsEventbridgeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IntegrationAwsEventbridge)(nil)).Elem()
}

func (o IntegrationAwsEventbridgeOutput) ToIntegrationAwsEventbridgeOutput() IntegrationAwsEventbridgeOutput {
	return o
}

func (o IntegrationAwsEventbridgeOutput) ToIntegrationAwsEventbridgeOutputWithContext(ctx context.Context) IntegrationAwsEventbridgeOutput {
	return o
}

// The 12 digit AWS Account ID where you want the events to be sent to.
func (o IntegrationAwsEventbridgeOutput) AwsAccountId() pulumi.StringOutput {
	return o.ApplyT(func(v *IntegrationAwsEventbridge) pulumi.StringOutput { return v.AwsAccountId }).(pulumi.StringOutput)
}

// The AWS region where you the events to be sent to. (e.g. us-west-1, us-west-2, ..., etc.)
func (o IntegrationAwsEventbridgeOutput) AwsRegion() pulumi.StringOutput {
	return o.ApplyT(func(v *IntegrationAwsEventbridge) pulumi.StringOutput { return v.AwsRegion }).(pulumi.StringOutput)
}

// Instance identifier
func (o IntegrationAwsEventbridgeOutput) InstanceId() pulumi.IntOutput {
	return o.ApplyT(func(v *IntegrationAwsEventbridge) pulumi.IntOutput { return v.InstanceId }).(pulumi.IntOutput)
}

// A (durable) queue on your RabbitMQ instance.
func (o IntegrationAwsEventbridgeOutput) Queue() pulumi.StringOutput {
	return o.ApplyT(func(v *IntegrationAwsEventbridge) pulumi.StringOutput { return v.Queue }).(pulumi.StringOutput)
}

// Always set to null, unless there is an error starting the EventBridge.
func (o IntegrationAwsEventbridgeOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *IntegrationAwsEventbridge) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// The VHost the queue resides in.
func (o IntegrationAwsEventbridgeOutput) Vhost() pulumi.StringOutput {
	return o.ApplyT(func(v *IntegrationAwsEventbridge) pulumi.StringOutput { return v.Vhost }).(pulumi.StringOutput)
}

// Include message headers in the event data.
func (o IntegrationAwsEventbridgeOutput) WithHeaders() pulumi.BoolOutput {
	return o.ApplyT(func(v *IntegrationAwsEventbridge) pulumi.BoolOutput { return v.WithHeaders }).(pulumi.BoolOutput)
}

type IntegrationAwsEventbridgeArrayOutput struct{ *pulumi.OutputState }

func (IntegrationAwsEventbridgeArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*IntegrationAwsEventbridge)(nil)).Elem()
}

func (o IntegrationAwsEventbridgeArrayOutput) ToIntegrationAwsEventbridgeArrayOutput() IntegrationAwsEventbridgeArrayOutput {
	return o
}

func (o IntegrationAwsEventbridgeArrayOutput) ToIntegrationAwsEventbridgeArrayOutputWithContext(ctx context.Context) IntegrationAwsEventbridgeArrayOutput {
	return o
}

func (o IntegrationAwsEventbridgeArrayOutput) Index(i pulumi.IntInput) IntegrationAwsEventbridgeOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *IntegrationAwsEventbridge {
		return vs[0].([]*IntegrationAwsEventbridge)[vs[1].(int)]
	}).(IntegrationAwsEventbridgeOutput)
}

type IntegrationAwsEventbridgeMapOutput struct{ *pulumi.OutputState }

func (IntegrationAwsEventbridgeMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*IntegrationAwsEventbridge)(nil)).Elem()
}

func (o IntegrationAwsEventbridgeMapOutput) ToIntegrationAwsEventbridgeMapOutput() IntegrationAwsEventbridgeMapOutput {
	return o
}

func (o IntegrationAwsEventbridgeMapOutput) ToIntegrationAwsEventbridgeMapOutputWithContext(ctx context.Context) IntegrationAwsEventbridgeMapOutput {
	return o
}

func (o IntegrationAwsEventbridgeMapOutput) MapIndex(k pulumi.StringInput) IntegrationAwsEventbridgeOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *IntegrationAwsEventbridge {
		return vs[0].(map[string]*IntegrationAwsEventbridge)[vs[1].(string)]
	}).(IntegrationAwsEventbridgeOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*IntegrationAwsEventbridgeInput)(nil)).Elem(), &IntegrationAwsEventbridge{})
	pulumi.RegisterInputType(reflect.TypeOf((*IntegrationAwsEventbridgeArrayInput)(nil)).Elem(), IntegrationAwsEventbridgeArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*IntegrationAwsEventbridgeMapInput)(nil)).Elem(), IntegrationAwsEventbridgeMap{})
	pulumi.RegisterOutputType(IntegrationAwsEventbridgeOutput{})
	pulumi.RegisterOutputType(IntegrationAwsEventbridgeArrayOutput{})
	pulumi.RegisterOutputType(IntegrationAwsEventbridgeMapOutput{})
}