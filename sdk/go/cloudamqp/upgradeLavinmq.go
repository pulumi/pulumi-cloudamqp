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

// This resource allows you to upgrade LavinMQ version.
//
// See below example usage.
//
// Only available for dedicated subscription plans running ***LavinMQ***.
//
// ## Example Usage
//
// <details>
//
//	<summary>
//	  <b>
//	    <i>Upgrade LavinMQ, specify which version to upgrade to, from v1.32.0</i>
//	  </b>
//	</summary>
//
// Specify the version to upgrade to. List available upgradable versions, use [CloudAMQP API](https://docs.cloudamqp.com/cloudamqp_api.html#get-available-versions).
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
//				Name:   pulumi.String("lavinmq-version-upgrade-test"),
//				Plan:   pulumi.String("lynx-1"),
//				Region: pulumi.String("amazon-web-services::us-west-1"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = cloudamqp.NewUpgradeLavinmq(ctx, "upgrade", &cloudamqp.UpgradeLavinmqArgs{
//				InstanceId: instance.ID(),
//				NewVersion: pulumi.String("1.3.1"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// </details>
//
// ## Important Upgrade Information
//
// > - All single node upgrades will require some downtime since LavinMQ needs a restart.
// > - Auto delete queues (queues that are marked AD) will be deleted during the update.
//
// ## Import
//
// Not possible to import this resource.
type UpgradeLavinmq struct {
	pulumi.CustomResourceState

	// The CloudAMQP instance identifier
	InstanceId pulumi.IntOutput `pulumi:"instanceId"`
	// The new version to upgrade to
	NewVersion pulumi.StringPtrOutput `pulumi:"newVersion"`
}

// NewUpgradeLavinmq registers a new resource with the given unique name, arguments, and options.
func NewUpgradeLavinmq(ctx *pulumi.Context,
	name string, args *UpgradeLavinmqArgs, opts ...pulumi.ResourceOption) (*UpgradeLavinmq, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.InstanceId == nil {
		return nil, errors.New("invalid value for required argument 'InstanceId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource UpgradeLavinmq
	err := ctx.RegisterResource("cloudamqp:index/upgradeLavinmq:UpgradeLavinmq", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetUpgradeLavinmq gets an existing UpgradeLavinmq resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetUpgradeLavinmq(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *UpgradeLavinmqState, opts ...pulumi.ResourceOption) (*UpgradeLavinmq, error) {
	var resource UpgradeLavinmq
	err := ctx.ReadResource("cloudamqp:index/upgradeLavinmq:UpgradeLavinmq", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering UpgradeLavinmq resources.
type upgradeLavinmqState struct {
	// The CloudAMQP instance identifier
	InstanceId *int `pulumi:"instanceId"`
	// The new version to upgrade to
	NewVersion *string `pulumi:"newVersion"`
}

type UpgradeLavinmqState struct {
	// The CloudAMQP instance identifier
	InstanceId pulumi.IntPtrInput
	// The new version to upgrade to
	NewVersion pulumi.StringPtrInput
}

func (UpgradeLavinmqState) ElementType() reflect.Type {
	return reflect.TypeOf((*upgradeLavinmqState)(nil)).Elem()
}

type upgradeLavinmqArgs struct {
	// The CloudAMQP instance identifier
	InstanceId int `pulumi:"instanceId"`
	// The new version to upgrade to
	NewVersion *string `pulumi:"newVersion"`
}

// The set of arguments for constructing a UpgradeLavinmq resource.
type UpgradeLavinmqArgs struct {
	// The CloudAMQP instance identifier
	InstanceId pulumi.IntInput
	// The new version to upgrade to
	NewVersion pulumi.StringPtrInput
}

func (UpgradeLavinmqArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*upgradeLavinmqArgs)(nil)).Elem()
}

type UpgradeLavinmqInput interface {
	pulumi.Input

	ToUpgradeLavinmqOutput() UpgradeLavinmqOutput
	ToUpgradeLavinmqOutputWithContext(ctx context.Context) UpgradeLavinmqOutput
}

func (*UpgradeLavinmq) ElementType() reflect.Type {
	return reflect.TypeOf((**UpgradeLavinmq)(nil)).Elem()
}

func (i *UpgradeLavinmq) ToUpgradeLavinmqOutput() UpgradeLavinmqOutput {
	return i.ToUpgradeLavinmqOutputWithContext(context.Background())
}

func (i *UpgradeLavinmq) ToUpgradeLavinmqOutputWithContext(ctx context.Context) UpgradeLavinmqOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UpgradeLavinmqOutput)
}

// UpgradeLavinmqArrayInput is an input type that accepts UpgradeLavinmqArray and UpgradeLavinmqArrayOutput values.
// You can construct a concrete instance of `UpgradeLavinmqArrayInput` via:
//
//	UpgradeLavinmqArray{ UpgradeLavinmqArgs{...} }
type UpgradeLavinmqArrayInput interface {
	pulumi.Input

	ToUpgradeLavinmqArrayOutput() UpgradeLavinmqArrayOutput
	ToUpgradeLavinmqArrayOutputWithContext(context.Context) UpgradeLavinmqArrayOutput
}

type UpgradeLavinmqArray []UpgradeLavinmqInput

func (UpgradeLavinmqArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*UpgradeLavinmq)(nil)).Elem()
}

func (i UpgradeLavinmqArray) ToUpgradeLavinmqArrayOutput() UpgradeLavinmqArrayOutput {
	return i.ToUpgradeLavinmqArrayOutputWithContext(context.Background())
}

func (i UpgradeLavinmqArray) ToUpgradeLavinmqArrayOutputWithContext(ctx context.Context) UpgradeLavinmqArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UpgradeLavinmqArrayOutput)
}

// UpgradeLavinmqMapInput is an input type that accepts UpgradeLavinmqMap and UpgradeLavinmqMapOutput values.
// You can construct a concrete instance of `UpgradeLavinmqMapInput` via:
//
//	UpgradeLavinmqMap{ "key": UpgradeLavinmqArgs{...} }
type UpgradeLavinmqMapInput interface {
	pulumi.Input

	ToUpgradeLavinmqMapOutput() UpgradeLavinmqMapOutput
	ToUpgradeLavinmqMapOutputWithContext(context.Context) UpgradeLavinmqMapOutput
}

type UpgradeLavinmqMap map[string]UpgradeLavinmqInput

func (UpgradeLavinmqMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*UpgradeLavinmq)(nil)).Elem()
}

func (i UpgradeLavinmqMap) ToUpgradeLavinmqMapOutput() UpgradeLavinmqMapOutput {
	return i.ToUpgradeLavinmqMapOutputWithContext(context.Background())
}

func (i UpgradeLavinmqMap) ToUpgradeLavinmqMapOutputWithContext(ctx context.Context) UpgradeLavinmqMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UpgradeLavinmqMapOutput)
}

type UpgradeLavinmqOutput struct{ *pulumi.OutputState }

func (UpgradeLavinmqOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**UpgradeLavinmq)(nil)).Elem()
}

func (o UpgradeLavinmqOutput) ToUpgradeLavinmqOutput() UpgradeLavinmqOutput {
	return o
}

func (o UpgradeLavinmqOutput) ToUpgradeLavinmqOutputWithContext(ctx context.Context) UpgradeLavinmqOutput {
	return o
}

// The CloudAMQP instance identifier
func (o UpgradeLavinmqOutput) InstanceId() pulumi.IntOutput {
	return o.ApplyT(func(v *UpgradeLavinmq) pulumi.IntOutput { return v.InstanceId }).(pulumi.IntOutput)
}

// The new version to upgrade to
func (o UpgradeLavinmqOutput) NewVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UpgradeLavinmq) pulumi.StringPtrOutput { return v.NewVersion }).(pulumi.StringPtrOutput)
}

type UpgradeLavinmqArrayOutput struct{ *pulumi.OutputState }

func (UpgradeLavinmqArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*UpgradeLavinmq)(nil)).Elem()
}

func (o UpgradeLavinmqArrayOutput) ToUpgradeLavinmqArrayOutput() UpgradeLavinmqArrayOutput {
	return o
}

func (o UpgradeLavinmqArrayOutput) ToUpgradeLavinmqArrayOutputWithContext(ctx context.Context) UpgradeLavinmqArrayOutput {
	return o
}

func (o UpgradeLavinmqArrayOutput) Index(i pulumi.IntInput) UpgradeLavinmqOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *UpgradeLavinmq {
		return vs[0].([]*UpgradeLavinmq)[vs[1].(int)]
	}).(UpgradeLavinmqOutput)
}

type UpgradeLavinmqMapOutput struct{ *pulumi.OutputState }

func (UpgradeLavinmqMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*UpgradeLavinmq)(nil)).Elem()
}

func (o UpgradeLavinmqMapOutput) ToUpgradeLavinmqMapOutput() UpgradeLavinmqMapOutput {
	return o
}

func (o UpgradeLavinmqMapOutput) ToUpgradeLavinmqMapOutputWithContext(ctx context.Context) UpgradeLavinmqMapOutput {
	return o
}

func (o UpgradeLavinmqMapOutput) MapIndex(k pulumi.StringInput) UpgradeLavinmqOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *UpgradeLavinmq {
		return vs[0].(map[string]*UpgradeLavinmq)[vs[1].(string)]
	}).(UpgradeLavinmqOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*UpgradeLavinmqInput)(nil)).Elem(), &UpgradeLavinmq{})
	pulumi.RegisterInputType(reflect.TypeOf((*UpgradeLavinmqArrayInput)(nil)).Elem(), UpgradeLavinmqArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*UpgradeLavinmqMapInput)(nil)).Elem(), UpgradeLavinmqMap{})
	pulumi.RegisterOutputType(UpgradeLavinmqOutput{})
	pulumi.RegisterOutputType(UpgradeLavinmqArrayOutput{})
	pulumi.RegisterOutputType(UpgradeLavinmqMapOutput{})
}
