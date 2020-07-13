// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package cloudamqp

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// This resource allows you to install or uninstall community plugins. Once installed the plugin will be available in `Plugin`.
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
// 		_, err := cloudamqp.NewPluginCommunity(ctx, "rabbitmqDelayedMessageExchange", &cloudamqp.PluginCommunityArgs{
// 			InstanceId: pulumi.Any(cloudamqp_instance.Instance_01.Id),
// 			Enabled:    pulumi.Bool(true),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ## Depedency
//
// This resource depends on CloudAMQP instance identifier, `cloudamqp_instance.instance.id`.
type PluginCommunity struct {
	pulumi.CustomResourceState

	// Enable or disable the plugins.
	Enabled pulumi.BoolOutput `pulumi:"enabled"`
	// The CloudAMQP instance ID.
	InstanceId pulumi.IntOutput `pulumi:"instanceId"`
	// The name of the Rabbit MQ plugin.
	Name pulumi.StringOutput `pulumi:"name"`
}

// NewPluginCommunity registers a new resource with the given unique name, arguments, and options.
func NewPluginCommunity(ctx *pulumi.Context,
	name string, args *PluginCommunityArgs, opts ...pulumi.ResourceOption) (*PluginCommunity, error) {
	if args == nil || args.Enabled == nil {
		return nil, errors.New("missing required argument 'Enabled'")
	}
	if args == nil || args.InstanceId == nil {
		return nil, errors.New("missing required argument 'InstanceId'")
	}
	if args == nil {
		args = &PluginCommunityArgs{}
	}
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
	// Enable or disable the plugins.
	Enabled *bool `pulumi:"enabled"`
	// The CloudAMQP instance ID.
	InstanceId *int `pulumi:"instanceId"`
	// The name of the Rabbit MQ plugin.
	Name *string `pulumi:"name"`
}

type PluginCommunityState struct {
	// Enable or disable the plugins.
	Enabled pulumi.BoolPtrInput
	// The CloudAMQP instance ID.
	InstanceId pulumi.IntPtrInput
	// The name of the Rabbit MQ plugin.
	Name pulumi.StringPtrInput
}

func (PluginCommunityState) ElementType() reflect.Type {
	return reflect.TypeOf((*pluginCommunityState)(nil)).Elem()
}

type pluginCommunityArgs struct {
	// Enable or disable the plugins.
	Enabled bool `pulumi:"enabled"`
	// The CloudAMQP instance ID.
	InstanceId int `pulumi:"instanceId"`
	// The name of the Rabbit MQ plugin.
	Name *string `pulumi:"name"`
}

// The set of arguments for constructing a PluginCommunity resource.
type PluginCommunityArgs struct {
	// Enable or disable the plugins.
	Enabled pulumi.BoolInput
	// The CloudAMQP instance ID.
	InstanceId pulumi.IntInput
	// The name of the Rabbit MQ plugin.
	Name pulumi.StringPtrInput
}

func (PluginCommunityArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*pluginCommunityArgs)(nil)).Elem()
}
