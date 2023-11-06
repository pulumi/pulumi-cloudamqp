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

// ## Import
//
// `cloudamqp_plugin` can be imported using the name argument of the resource together with CloudAMQP instance identifier. The name and identifier are CSV separated, see example below.
//
// ```sh
//
//	$ pulumi import cloudamqp:index/pluginCommunity:PluginCommunity <resource_name> <plugin_name>,<instance_id>`
//
// ```
type PluginCommunity struct {
	pulumi.CustomResourceState

	// The description of the plugin.
	Description pulumi.StringOutput `pulumi:"description"`
	// Enable or disable the plugins.
	Enabled pulumi.BoolOutput `pulumi:"enabled"`
	// The CloudAMQP instance ID.
	InstanceId pulumi.IntOutput `pulumi:"instanceId"`
	// The name of the Rabbit MQ community plugin.
	Name pulumi.StringOutput `pulumi:"name"`
	// Required version of RabbitMQ.
	Require pulumi.StringOutput `pulumi:"require"`
}

// NewPluginCommunity registers a new resource with the given unique name, arguments, and options.
func NewPluginCommunity(ctx *pulumi.Context,
	name string, args *PluginCommunityArgs, opts ...pulumi.ResourceOption) (*PluginCommunity, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Enabled == nil {
		return nil, errors.New("invalid value for required argument 'Enabled'")
	}
	if args.InstanceId == nil {
		return nil, errors.New("invalid value for required argument 'InstanceId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource PluginCommunity
	err := ctx.RegisterResource("cloudamqp:index/pluginCommunity:PluginCommunity", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPluginCommunity gets an existing PluginCommunity resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPluginCommunity(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PluginCommunityState, opts ...pulumi.ResourceOption) (*PluginCommunity, error) {
	var resource PluginCommunity
	err := ctx.ReadResource("cloudamqp:index/pluginCommunity:PluginCommunity", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PluginCommunity resources.
type pluginCommunityState struct {
	// The description of the plugin.
	Description *string `pulumi:"description"`
	// Enable or disable the plugins.
	Enabled *bool `pulumi:"enabled"`
	// The CloudAMQP instance ID.
	InstanceId *int `pulumi:"instanceId"`
	// The name of the Rabbit MQ community plugin.
	Name *string `pulumi:"name"`
	// Required version of RabbitMQ.
	Require *string `pulumi:"require"`
}

type PluginCommunityState struct {
	// The description of the plugin.
	Description pulumi.StringPtrInput
	// Enable or disable the plugins.
	Enabled pulumi.BoolPtrInput
	// The CloudAMQP instance ID.
	InstanceId pulumi.IntPtrInput
	// The name of the Rabbit MQ community plugin.
	Name pulumi.StringPtrInput
	// Required version of RabbitMQ.
	Require pulumi.StringPtrInput
}

func (PluginCommunityState) ElementType() reflect.Type {
	return reflect.TypeOf((*pluginCommunityState)(nil)).Elem()
}

type pluginCommunityArgs struct {
	// Enable or disable the plugins.
	Enabled bool `pulumi:"enabled"`
	// The CloudAMQP instance ID.
	InstanceId int `pulumi:"instanceId"`
	// The name of the Rabbit MQ community plugin.
	Name *string `pulumi:"name"`
}

// The set of arguments for constructing a PluginCommunity resource.
type PluginCommunityArgs struct {
	// Enable or disable the plugins.
	Enabled pulumi.BoolInput
	// The CloudAMQP instance ID.
	InstanceId pulumi.IntInput
	// The name of the Rabbit MQ community plugin.
	Name pulumi.StringPtrInput
}

func (PluginCommunityArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*pluginCommunityArgs)(nil)).Elem()
}

type PluginCommunityInput interface {
	pulumi.Input

	ToPluginCommunityOutput() PluginCommunityOutput
	ToPluginCommunityOutputWithContext(ctx context.Context) PluginCommunityOutput
}

func (*PluginCommunity) ElementType() reflect.Type {
	return reflect.TypeOf((**PluginCommunity)(nil)).Elem()
}

func (i *PluginCommunity) ToPluginCommunityOutput() PluginCommunityOutput {
	return i.ToPluginCommunityOutputWithContext(context.Background())
}

func (i *PluginCommunity) ToPluginCommunityOutputWithContext(ctx context.Context) PluginCommunityOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PluginCommunityOutput)
}

// PluginCommunityArrayInput is an input type that accepts PluginCommunityArray and PluginCommunityArrayOutput values.
// You can construct a concrete instance of `PluginCommunityArrayInput` via:
//
//	PluginCommunityArray{ PluginCommunityArgs{...} }
type PluginCommunityArrayInput interface {
	pulumi.Input

	ToPluginCommunityArrayOutput() PluginCommunityArrayOutput
	ToPluginCommunityArrayOutputWithContext(context.Context) PluginCommunityArrayOutput
}

type PluginCommunityArray []PluginCommunityInput

func (PluginCommunityArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*PluginCommunity)(nil)).Elem()
}

func (i PluginCommunityArray) ToPluginCommunityArrayOutput() PluginCommunityArrayOutput {
	return i.ToPluginCommunityArrayOutputWithContext(context.Background())
}

func (i PluginCommunityArray) ToPluginCommunityArrayOutputWithContext(ctx context.Context) PluginCommunityArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PluginCommunityArrayOutput)
}

// PluginCommunityMapInput is an input type that accepts PluginCommunityMap and PluginCommunityMapOutput values.
// You can construct a concrete instance of `PluginCommunityMapInput` via:
//
//	PluginCommunityMap{ "key": PluginCommunityArgs{...} }
type PluginCommunityMapInput interface {
	pulumi.Input

	ToPluginCommunityMapOutput() PluginCommunityMapOutput
	ToPluginCommunityMapOutputWithContext(context.Context) PluginCommunityMapOutput
}

type PluginCommunityMap map[string]PluginCommunityInput

func (PluginCommunityMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*PluginCommunity)(nil)).Elem()
}

func (i PluginCommunityMap) ToPluginCommunityMapOutput() PluginCommunityMapOutput {
	return i.ToPluginCommunityMapOutputWithContext(context.Background())
}

func (i PluginCommunityMap) ToPluginCommunityMapOutputWithContext(ctx context.Context) PluginCommunityMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PluginCommunityMapOutput)
}

type PluginCommunityOutput struct{ *pulumi.OutputState }

func (PluginCommunityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PluginCommunity)(nil)).Elem()
}

func (o PluginCommunityOutput) ToPluginCommunityOutput() PluginCommunityOutput {
	return o
}

func (o PluginCommunityOutput) ToPluginCommunityOutputWithContext(ctx context.Context) PluginCommunityOutput {
	return o
}

// The description of the plugin.
func (o PluginCommunityOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v *PluginCommunity) pulumi.StringOutput { return v.Description }).(pulumi.StringOutput)
}

// Enable or disable the plugins.
func (o PluginCommunityOutput) Enabled() pulumi.BoolOutput {
	return o.ApplyT(func(v *PluginCommunity) pulumi.BoolOutput { return v.Enabled }).(pulumi.BoolOutput)
}

// The CloudAMQP instance ID.
func (o PluginCommunityOutput) InstanceId() pulumi.IntOutput {
	return o.ApplyT(func(v *PluginCommunity) pulumi.IntOutput { return v.InstanceId }).(pulumi.IntOutput)
}

// The name of the Rabbit MQ community plugin.
func (o PluginCommunityOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *PluginCommunity) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Required version of RabbitMQ.
func (o PluginCommunityOutput) Require() pulumi.StringOutput {
	return o.ApplyT(func(v *PluginCommunity) pulumi.StringOutput { return v.Require }).(pulumi.StringOutput)
}

type PluginCommunityArrayOutput struct{ *pulumi.OutputState }

func (PluginCommunityArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*PluginCommunity)(nil)).Elem()
}

func (o PluginCommunityArrayOutput) ToPluginCommunityArrayOutput() PluginCommunityArrayOutput {
	return o
}

func (o PluginCommunityArrayOutput) ToPluginCommunityArrayOutputWithContext(ctx context.Context) PluginCommunityArrayOutput {
	return o
}

func (o PluginCommunityArrayOutput) Index(i pulumi.IntInput) PluginCommunityOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *PluginCommunity {
		return vs[0].([]*PluginCommunity)[vs[1].(int)]
	}).(PluginCommunityOutput)
}

type PluginCommunityMapOutput struct{ *pulumi.OutputState }

func (PluginCommunityMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*PluginCommunity)(nil)).Elem()
}

func (o PluginCommunityMapOutput) ToPluginCommunityMapOutput() PluginCommunityMapOutput {
	return o
}

func (o PluginCommunityMapOutput) ToPluginCommunityMapOutputWithContext(ctx context.Context) PluginCommunityMapOutput {
	return o
}

func (o PluginCommunityMapOutput) MapIndex(k pulumi.StringInput) PluginCommunityOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *PluginCommunity {
		return vs[0].(map[string]*PluginCommunity)[vs[1].(string)]
	}).(PluginCommunityOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*PluginCommunityInput)(nil)).Elem(), &PluginCommunity{})
	pulumi.RegisterInputType(reflect.TypeOf((*PluginCommunityArrayInput)(nil)).Elem(), PluginCommunityArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*PluginCommunityMapInput)(nil)).Elem(), PluginCommunityMap{})
	pulumi.RegisterOutputType(PluginCommunityOutput{})
	pulumi.RegisterOutputType(PluginCommunityArrayOutput{})
	pulumi.RegisterOutputType(PluginCommunityMapOutput{})
}
