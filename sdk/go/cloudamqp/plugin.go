// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package cloudamqp

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// This resource allows you to enable or disable Rabbit MQ plugins.
//
// Only available for dedicated subscription plans.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-cloudamqp/sdk/v2/go/cloudamqp"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := cloudamqp.NewPlugin(ctx, "pluginRabbitmqTop", &cloudamqp.PluginArgs{
// 			InstanceId: pulumi.Any(cloudamqp_instance.Instance.Id),
// 			Enabled:    pulumi.Bool(true),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ## Dependency
//
// This resource depends on CloudAMQP instance identifier, `cloudamqp_instance.instance.id`.
//
// ## Import
//
// `cloudamqp_plugin` can be imported using the name argument of the resource together with CloudAMQP instance identifier. The name and identifier are CSV separated, see example below.
//
// ```sh
//  $ pulumi import cloudamqp:index/plugin:Plugin rabbitmq_management rabbitmq_management,<instance_id>`
// ```
type Plugin struct {
	pulumi.CustomResourceState

	// Enable or disable the plugins.
	Enabled pulumi.BoolOutput `pulumi:"enabled"`
	// The CloudAMQP instance ID.
	InstanceId pulumi.IntOutput `pulumi:"instanceId"`
	// The name of the Rabbit MQ plugin.
	Name pulumi.StringOutput `pulumi:"name"`
}

// NewPlugin registers a new resource with the given unique name, arguments, and options.
func NewPlugin(ctx *pulumi.Context,
	name string, args *PluginArgs, opts ...pulumi.ResourceOption) (*Plugin, error) {
	if args == nil || args.Enabled == nil {
		return nil, errors.New("missing required argument 'Enabled'")
	}
	if args == nil || args.InstanceId == nil {
		return nil, errors.New("missing required argument 'InstanceId'")
	}
	if args == nil {
		args = &PluginArgs{}
	}
	var resource Plugin
	err := ctx.RegisterResource("cloudamqp:index/plugin:Plugin", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPlugin gets an existing Plugin resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPlugin(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PluginState, opts ...pulumi.ResourceOption) (*Plugin, error) {
	var resource Plugin
	err := ctx.ReadResource("cloudamqp:index/plugin:Plugin", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Plugin resources.
type pluginState struct {
	// Enable or disable the plugins.
	Enabled *bool `pulumi:"enabled"`
	// The CloudAMQP instance ID.
	InstanceId *int `pulumi:"instanceId"`
	// The name of the Rabbit MQ plugin.
	Name *string `pulumi:"name"`
}

type PluginState struct {
	// Enable or disable the plugins.
	Enabled pulumi.BoolPtrInput
	// The CloudAMQP instance ID.
	InstanceId pulumi.IntPtrInput
	// The name of the Rabbit MQ plugin.
	Name pulumi.StringPtrInput
}

func (PluginState) ElementType() reflect.Type {
	return reflect.TypeOf((*pluginState)(nil)).Elem()
}

type pluginArgs struct {
	// Enable or disable the plugins.
	Enabled bool `pulumi:"enabled"`
	// The CloudAMQP instance ID.
	InstanceId int `pulumi:"instanceId"`
	// The name of the Rabbit MQ plugin.
	Name *string `pulumi:"name"`
}

// The set of arguments for constructing a Plugin resource.
type PluginArgs struct {
	// Enable or disable the plugins.
	Enabled pulumi.BoolInput
	// The CloudAMQP instance ID.
	InstanceId pulumi.IntInput
	// The name of the Rabbit MQ plugin.
	Name pulumi.StringPtrInput
}

func (PluginArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*pluginArgs)(nil)).Elem()
}

type PluginInput interface {
	pulumi.Input

	ToPluginOutput() PluginOutput
	ToPluginOutputWithContext(ctx context.Context) PluginOutput
}

func (Plugin) ElementType() reflect.Type {
	return reflect.TypeOf((*Plugin)(nil)).Elem()
}

func (i Plugin) ToPluginOutput() PluginOutput {
	return i.ToPluginOutputWithContext(context.Background())
}

func (i Plugin) ToPluginOutputWithContext(ctx context.Context) PluginOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PluginOutput)
}

type PluginOutput struct {
	*pulumi.OutputState
}

func (PluginOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PluginOutput)(nil)).Elem()
}

func (o PluginOutput) ToPluginOutput() PluginOutput {
	return o
}

func (o PluginOutput) ToPluginOutputWithContext(ctx context.Context) PluginOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(PluginOutput{})
}
