// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package cloudamqp

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

type PluginCommunity struct {
	pulumi.CustomResourceState

	// If the plugin is enabled
	Enabled pulumi.BoolOutput `pulumi:"enabled"`
	// Instance identifier
	InstanceId pulumi.IntOutput `pulumi:"instanceId"`
	// The name of the plugin
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
	// If the plugin is enabled
	Enabled *bool `pulumi:"enabled"`
	// Instance identifier
	InstanceId *int `pulumi:"instanceId"`
	// The name of the plugin
	Name *string `pulumi:"name"`
}

type PluginCommunityState struct {
	// If the plugin is enabled
	Enabled pulumi.BoolPtrInput
	// Instance identifier
	InstanceId pulumi.IntPtrInput
	// The name of the plugin
	Name pulumi.StringPtrInput
}

func (PluginCommunityState) ElementType() reflect.Type {
	return reflect.TypeOf((*pluginCommunityState)(nil)).Elem()
}

type pluginCommunityArgs struct {
	// If the plugin is enabled
	Enabled bool `pulumi:"enabled"`
	// Instance identifier
	InstanceId int `pulumi:"instanceId"`
	// The name of the plugin
	Name *string `pulumi:"name"`
}

// The set of arguments for constructing a PluginCommunity resource.
type PluginCommunityArgs struct {
	// If the plugin is enabled
	Enabled pulumi.BoolInput
	// Instance identifier
	InstanceId pulumi.IntInput
	// The name of the plugin
	Name pulumi.StringPtrInput
}

func (PluginCommunityArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*pluginCommunityArgs)(nil)).Elem()
}
