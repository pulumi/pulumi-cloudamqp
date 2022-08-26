// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cloudamqp

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This resource allows you update RabbitMQ config.
//
// Only available for dedicated subscription plans.
//
// ## Argument threshold values
//
// | Argument | Type | Default | Min | Max | Unit | Affect | Note |
// |---|---|---|---|---|---|---|---|
// | heartbeat | int | 120 | 0 | - |  | Only effects new connections |  |
// | connectionMax | int | -1 | 1 | - |  | RabbitMQ restart required | -1 in the provider corresponds to INFINITY in the RabbitMQ config |
// | channelMax | int | 128 | 0 | - |  | Only effects new connections |  |
// | consumerTimeout | int | 7200000 | 10000 | 86400000 | milliseconds | Only effects new channels | -1 in the provider corresponds to false (disable) in the RabbitMQ config |
// | vmMemoryHighWatermark | float | 0.81 | 0.4 | 0.9 |  | Applied immediately |  |
// | queueIndexEmbedMsgsBelow | int | 4096 | 1 | 10485760 | bytes | Applied immediately for new queues, requires restart for existing queues |  |
// | maxMessageSize | int | 134217728 | 1 | 536870912 | bytes | Only effects new channels |  |
// | logExchangeLevel | string | error | - | - |  | RabbitMQ restart required | debug, info, warning, error, critical |
//
// ## Dependency
//
// This resource depends on CloudAMQP instance identifier, `cloudamqp_instance.instance.id`.
//
// ## Import
//
// `cloudamqp_rabbitmq_configuration` can be imported using the CloudAMQP instance identifier.
//
// ```sh
//
//	$ pulumi import cloudamqp:index/rabbitConfiguration:RabbitConfiguration config <instance_id>`
//
// ```
type RabbitConfiguration struct {
	pulumi.CustomResourceState

	// Set the maximum permissible number of channels per connection.
	ChannelMax pulumi.IntPtrOutput `pulumi:"channelMax"`
	// Set the maximum permissible number of connection.
	ConnectionMax pulumi.IntPtrOutput `pulumi:"connectionMax"`
	// A consumer that has recevied a message and does not acknowledge that message within the timeout in milliseconds
	ConsumerTimeout pulumi.IntPtrOutput `pulumi:"consumerTimeout"`
	// Set the server AMQP 0-9-1 heartbeat timeout in seconds.
	Heartbeat pulumi.IntPtrOutput `pulumi:"heartbeat"`
	// The CloudAMQP instance ID.
	InstanceId pulumi.IntOutput `pulumi:"instanceId"`
	// Log level for the logger used for log integrations and the CloudAMQP Console log view.
	LogExchangeLevel pulumi.StringPtrOutput `pulumi:"logExchangeLevel"`
	// The largest allowed message payload size in bytes.
	MaxMessageSize pulumi.IntPtrOutput `pulumi:"maxMessageSize"`
	// Size in bytes below which to embed messages in the queue index.
	QueueIndexEmbedMsgsBelow pulumi.IntPtrOutput `pulumi:"queueIndexEmbedMsgsBelow"`
	// When the server will enter memory based flow-control as relative to the maximum available memory.
	VmMemoryHighWatermark pulumi.Float64PtrOutput `pulumi:"vmMemoryHighWatermark"`
}

// NewRabbitConfiguration registers a new resource with the given unique name, arguments, and options.
func NewRabbitConfiguration(ctx *pulumi.Context,
	name string, args *RabbitConfigurationArgs, opts ...pulumi.ResourceOption) (*RabbitConfiguration, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.InstanceId == nil {
		return nil, errors.New("invalid value for required argument 'InstanceId'")
	}
	var resource RabbitConfiguration
	err := ctx.RegisterResource("cloudamqp:index/rabbitConfiguration:RabbitConfiguration", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRabbitConfiguration gets an existing RabbitConfiguration resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRabbitConfiguration(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RabbitConfigurationState, opts ...pulumi.ResourceOption) (*RabbitConfiguration, error) {
	var resource RabbitConfiguration
	err := ctx.ReadResource("cloudamqp:index/rabbitConfiguration:RabbitConfiguration", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RabbitConfiguration resources.
type rabbitConfigurationState struct {
	// Set the maximum permissible number of channels per connection.
	ChannelMax *int `pulumi:"channelMax"`
	// Set the maximum permissible number of connection.
	ConnectionMax *int `pulumi:"connectionMax"`
	// A consumer that has recevied a message and does not acknowledge that message within the timeout in milliseconds
	ConsumerTimeout *int `pulumi:"consumerTimeout"`
	// Set the server AMQP 0-9-1 heartbeat timeout in seconds.
	Heartbeat *int `pulumi:"heartbeat"`
	// The CloudAMQP instance ID.
	InstanceId *int `pulumi:"instanceId"`
	// Log level for the logger used for log integrations and the CloudAMQP Console log view.
	LogExchangeLevel *string `pulumi:"logExchangeLevel"`
	// The largest allowed message payload size in bytes.
	MaxMessageSize *int `pulumi:"maxMessageSize"`
	// Size in bytes below which to embed messages in the queue index.
	QueueIndexEmbedMsgsBelow *int `pulumi:"queueIndexEmbedMsgsBelow"`
	// When the server will enter memory based flow-control as relative to the maximum available memory.
	VmMemoryHighWatermark *float64 `pulumi:"vmMemoryHighWatermark"`
}

type RabbitConfigurationState struct {
	// Set the maximum permissible number of channels per connection.
	ChannelMax pulumi.IntPtrInput
	// Set the maximum permissible number of connection.
	ConnectionMax pulumi.IntPtrInput
	// A consumer that has recevied a message and does not acknowledge that message within the timeout in milliseconds
	ConsumerTimeout pulumi.IntPtrInput
	// Set the server AMQP 0-9-1 heartbeat timeout in seconds.
	Heartbeat pulumi.IntPtrInput
	// The CloudAMQP instance ID.
	InstanceId pulumi.IntPtrInput
	// Log level for the logger used for log integrations and the CloudAMQP Console log view.
	LogExchangeLevel pulumi.StringPtrInput
	// The largest allowed message payload size in bytes.
	MaxMessageSize pulumi.IntPtrInput
	// Size in bytes below which to embed messages in the queue index.
	QueueIndexEmbedMsgsBelow pulumi.IntPtrInput
	// When the server will enter memory based flow-control as relative to the maximum available memory.
	VmMemoryHighWatermark pulumi.Float64PtrInput
}

func (RabbitConfigurationState) ElementType() reflect.Type {
	return reflect.TypeOf((*rabbitConfigurationState)(nil)).Elem()
}

type rabbitConfigurationArgs struct {
	// Set the maximum permissible number of channels per connection.
	ChannelMax *int `pulumi:"channelMax"`
	// Set the maximum permissible number of connection.
	ConnectionMax *int `pulumi:"connectionMax"`
	// A consumer that has recevied a message and does not acknowledge that message within the timeout in milliseconds
	ConsumerTimeout *int `pulumi:"consumerTimeout"`
	// Set the server AMQP 0-9-1 heartbeat timeout in seconds.
	Heartbeat *int `pulumi:"heartbeat"`
	// The CloudAMQP instance ID.
	InstanceId int `pulumi:"instanceId"`
	// Log level for the logger used for log integrations and the CloudAMQP Console log view.
	LogExchangeLevel *string `pulumi:"logExchangeLevel"`
	// The largest allowed message payload size in bytes.
	MaxMessageSize *int `pulumi:"maxMessageSize"`
	// Size in bytes below which to embed messages in the queue index.
	QueueIndexEmbedMsgsBelow *int `pulumi:"queueIndexEmbedMsgsBelow"`
	// When the server will enter memory based flow-control as relative to the maximum available memory.
	VmMemoryHighWatermark *float64 `pulumi:"vmMemoryHighWatermark"`
}

// The set of arguments for constructing a RabbitConfiguration resource.
type RabbitConfigurationArgs struct {
	// Set the maximum permissible number of channels per connection.
	ChannelMax pulumi.IntPtrInput
	// Set the maximum permissible number of connection.
	ConnectionMax pulumi.IntPtrInput
	// A consumer that has recevied a message and does not acknowledge that message within the timeout in milliseconds
	ConsumerTimeout pulumi.IntPtrInput
	// Set the server AMQP 0-9-1 heartbeat timeout in seconds.
	Heartbeat pulumi.IntPtrInput
	// The CloudAMQP instance ID.
	InstanceId pulumi.IntInput
	// Log level for the logger used for log integrations and the CloudAMQP Console log view.
	LogExchangeLevel pulumi.StringPtrInput
	// The largest allowed message payload size in bytes.
	MaxMessageSize pulumi.IntPtrInput
	// Size in bytes below which to embed messages in the queue index.
	QueueIndexEmbedMsgsBelow pulumi.IntPtrInput
	// When the server will enter memory based flow-control as relative to the maximum available memory.
	VmMemoryHighWatermark pulumi.Float64PtrInput
}

func (RabbitConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*rabbitConfigurationArgs)(nil)).Elem()
}

type RabbitConfigurationInput interface {
	pulumi.Input

	ToRabbitConfigurationOutput() RabbitConfigurationOutput
	ToRabbitConfigurationOutputWithContext(ctx context.Context) RabbitConfigurationOutput
}

func (*RabbitConfiguration) ElementType() reflect.Type {
	return reflect.TypeOf((**RabbitConfiguration)(nil)).Elem()
}

func (i *RabbitConfiguration) ToRabbitConfigurationOutput() RabbitConfigurationOutput {
	return i.ToRabbitConfigurationOutputWithContext(context.Background())
}

func (i *RabbitConfiguration) ToRabbitConfigurationOutputWithContext(ctx context.Context) RabbitConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RabbitConfigurationOutput)
}

// RabbitConfigurationArrayInput is an input type that accepts RabbitConfigurationArray and RabbitConfigurationArrayOutput values.
// You can construct a concrete instance of `RabbitConfigurationArrayInput` via:
//
//	RabbitConfigurationArray{ RabbitConfigurationArgs{...} }
type RabbitConfigurationArrayInput interface {
	pulumi.Input

	ToRabbitConfigurationArrayOutput() RabbitConfigurationArrayOutput
	ToRabbitConfigurationArrayOutputWithContext(context.Context) RabbitConfigurationArrayOutput
}

type RabbitConfigurationArray []RabbitConfigurationInput

func (RabbitConfigurationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RabbitConfiguration)(nil)).Elem()
}

func (i RabbitConfigurationArray) ToRabbitConfigurationArrayOutput() RabbitConfigurationArrayOutput {
	return i.ToRabbitConfigurationArrayOutputWithContext(context.Background())
}

func (i RabbitConfigurationArray) ToRabbitConfigurationArrayOutputWithContext(ctx context.Context) RabbitConfigurationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RabbitConfigurationArrayOutput)
}

// RabbitConfigurationMapInput is an input type that accepts RabbitConfigurationMap and RabbitConfigurationMapOutput values.
// You can construct a concrete instance of `RabbitConfigurationMapInput` via:
//
//	RabbitConfigurationMap{ "key": RabbitConfigurationArgs{...} }
type RabbitConfigurationMapInput interface {
	pulumi.Input

	ToRabbitConfigurationMapOutput() RabbitConfigurationMapOutput
	ToRabbitConfigurationMapOutputWithContext(context.Context) RabbitConfigurationMapOutput
}

type RabbitConfigurationMap map[string]RabbitConfigurationInput

func (RabbitConfigurationMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RabbitConfiguration)(nil)).Elem()
}

func (i RabbitConfigurationMap) ToRabbitConfigurationMapOutput() RabbitConfigurationMapOutput {
	return i.ToRabbitConfigurationMapOutputWithContext(context.Background())
}

func (i RabbitConfigurationMap) ToRabbitConfigurationMapOutputWithContext(ctx context.Context) RabbitConfigurationMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RabbitConfigurationMapOutput)
}

type RabbitConfigurationOutput struct{ *pulumi.OutputState }

func (RabbitConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RabbitConfiguration)(nil)).Elem()
}

func (o RabbitConfigurationOutput) ToRabbitConfigurationOutput() RabbitConfigurationOutput {
	return o
}

func (o RabbitConfigurationOutput) ToRabbitConfigurationOutputWithContext(ctx context.Context) RabbitConfigurationOutput {
	return o
}

// Set the maximum permissible number of channels per connection.
func (o RabbitConfigurationOutput) ChannelMax() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *RabbitConfiguration) pulumi.IntPtrOutput { return v.ChannelMax }).(pulumi.IntPtrOutput)
}

// Set the maximum permissible number of connection.
func (o RabbitConfigurationOutput) ConnectionMax() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *RabbitConfiguration) pulumi.IntPtrOutput { return v.ConnectionMax }).(pulumi.IntPtrOutput)
}

// A consumer that has recevied a message and does not acknowledge that message within the timeout in milliseconds
func (o RabbitConfigurationOutput) ConsumerTimeout() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *RabbitConfiguration) pulumi.IntPtrOutput { return v.ConsumerTimeout }).(pulumi.IntPtrOutput)
}

// Set the server AMQP 0-9-1 heartbeat timeout in seconds.
func (o RabbitConfigurationOutput) Heartbeat() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *RabbitConfiguration) pulumi.IntPtrOutput { return v.Heartbeat }).(pulumi.IntPtrOutput)
}

// The CloudAMQP instance ID.
func (o RabbitConfigurationOutput) InstanceId() pulumi.IntOutput {
	return o.ApplyT(func(v *RabbitConfiguration) pulumi.IntOutput { return v.InstanceId }).(pulumi.IntOutput)
}

// Log level for the logger used for log integrations and the CloudAMQP Console log view.
func (o RabbitConfigurationOutput) LogExchangeLevel() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RabbitConfiguration) pulumi.StringPtrOutput { return v.LogExchangeLevel }).(pulumi.StringPtrOutput)
}

// The largest allowed message payload size in bytes.
func (o RabbitConfigurationOutput) MaxMessageSize() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *RabbitConfiguration) pulumi.IntPtrOutput { return v.MaxMessageSize }).(pulumi.IntPtrOutput)
}

// Size in bytes below which to embed messages in the queue index.
func (o RabbitConfigurationOutput) QueueIndexEmbedMsgsBelow() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *RabbitConfiguration) pulumi.IntPtrOutput { return v.QueueIndexEmbedMsgsBelow }).(pulumi.IntPtrOutput)
}

// When the server will enter memory based flow-control as relative to the maximum available memory.
func (o RabbitConfigurationOutput) VmMemoryHighWatermark() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *RabbitConfiguration) pulumi.Float64PtrOutput { return v.VmMemoryHighWatermark }).(pulumi.Float64PtrOutput)
}

type RabbitConfigurationArrayOutput struct{ *pulumi.OutputState }

func (RabbitConfigurationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RabbitConfiguration)(nil)).Elem()
}

func (o RabbitConfigurationArrayOutput) ToRabbitConfigurationArrayOutput() RabbitConfigurationArrayOutput {
	return o
}

func (o RabbitConfigurationArrayOutput) ToRabbitConfigurationArrayOutputWithContext(ctx context.Context) RabbitConfigurationArrayOutput {
	return o
}

func (o RabbitConfigurationArrayOutput) Index(i pulumi.IntInput) RabbitConfigurationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *RabbitConfiguration {
		return vs[0].([]*RabbitConfiguration)[vs[1].(int)]
	}).(RabbitConfigurationOutput)
}

type RabbitConfigurationMapOutput struct{ *pulumi.OutputState }

func (RabbitConfigurationMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RabbitConfiguration)(nil)).Elem()
}

func (o RabbitConfigurationMapOutput) ToRabbitConfigurationMapOutput() RabbitConfigurationMapOutput {
	return o
}

func (o RabbitConfigurationMapOutput) ToRabbitConfigurationMapOutputWithContext(ctx context.Context) RabbitConfigurationMapOutput {
	return o
}

func (o RabbitConfigurationMapOutput) MapIndex(k pulumi.StringInput) RabbitConfigurationOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *RabbitConfiguration {
		return vs[0].(map[string]*RabbitConfiguration)[vs[1].(string)]
	}).(RabbitConfigurationOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RabbitConfigurationInput)(nil)).Elem(), &RabbitConfiguration{})
	pulumi.RegisterInputType(reflect.TypeOf((*RabbitConfigurationArrayInput)(nil)).Elem(), RabbitConfigurationArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*RabbitConfigurationMapInput)(nil)).Elem(), RabbitConfigurationMap{})
	pulumi.RegisterOutputType(RabbitConfigurationOutput{})
	pulumi.RegisterOutputType(RabbitConfigurationArrayOutput{})
	pulumi.RegisterOutputType(RabbitConfigurationMapOutput{})
}
