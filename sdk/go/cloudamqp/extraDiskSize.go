// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cloudamqp

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This resource allows you to expand the disk with additional storage capacity. There is no downtime when expanding the disk.
//
// Only available for dedicated subscription plans hosted at Amazon Web Services (AWS) at this time.
//
// > **WARNING:** Due to restrictions from cloud providers, it's only possible to resize the disk every 8 hours.
//
// Pricing is available at [cloudamqp.com](https://www.cloudamqp.com/).
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
//				Region:     pulumi.String("amazon-web-services::us-west-2"),
//				RmqVersion: pulumi.String("3.10.1"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = cloudamqp.NewExtraDiskSize(ctx, "resizeDisk", &cloudamqp.ExtraDiskSizeArgs{
//				InstanceId:    instance.ID(),
//				ExtraDiskSize: pulumi.Int(25),
//			})
//			if err != nil {
//				return err
//			}
//			_ = instance.ID().ApplyT(func(id string) (cloudamqp.GetNodesResult, error) {
//				return cloudamqp.GetNodesOutput(ctx, cloudamqp.GetNodesOutputArgs{
//					InstanceId: id,
//				}, nil), nil
//			}).(cloudamqp.GetNodesResultOutput)
//			return nil
//		})
//	}
//
// ```
//
// ## Import
//
// Not possible to import this resource.
type ExtraDiskSize struct {
	pulumi.CustomResourceState

	// Extra disk size in GB. Supported values: 25, 50, 100, 250, 500, 1000, 2000
	ExtraDiskSize pulumi.IntOutput `pulumi:"extraDiskSize"`
	// The CloudAMQP instance ID.
	InstanceId pulumi.IntOutput `pulumi:"instanceId"`
}

// NewExtraDiskSize registers a new resource with the given unique name, arguments, and options.
func NewExtraDiskSize(ctx *pulumi.Context,
	name string, args *ExtraDiskSizeArgs, opts ...pulumi.ResourceOption) (*ExtraDiskSize, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ExtraDiskSize == nil {
		return nil, errors.New("invalid value for required argument 'ExtraDiskSize'")
	}
	if args.InstanceId == nil {
		return nil, errors.New("invalid value for required argument 'InstanceId'")
	}
	var resource ExtraDiskSize
	err := ctx.RegisterResource("cloudamqp:index/extraDiskSize:ExtraDiskSize", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetExtraDiskSize gets an existing ExtraDiskSize resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetExtraDiskSize(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ExtraDiskSizeState, opts ...pulumi.ResourceOption) (*ExtraDiskSize, error) {
	var resource ExtraDiskSize
	err := ctx.ReadResource("cloudamqp:index/extraDiskSize:ExtraDiskSize", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ExtraDiskSize resources.
type extraDiskSizeState struct {
	// Extra disk size in GB. Supported values: 25, 50, 100, 250, 500, 1000, 2000
	ExtraDiskSize *int `pulumi:"extraDiskSize"`
	// The CloudAMQP instance ID.
	InstanceId *int `pulumi:"instanceId"`
}

type ExtraDiskSizeState struct {
	// Extra disk size in GB. Supported values: 25, 50, 100, 250, 500, 1000, 2000
	ExtraDiskSize pulumi.IntPtrInput
	// The CloudAMQP instance ID.
	InstanceId pulumi.IntPtrInput
}

func (ExtraDiskSizeState) ElementType() reflect.Type {
	return reflect.TypeOf((*extraDiskSizeState)(nil)).Elem()
}

type extraDiskSizeArgs struct {
	// Extra disk size in GB. Supported values: 25, 50, 100, 250, 500, 1000, 2000
	ExtraDiskSize int `pulumi:"extraDiskSize"`
	// The CloudAMQP instance ID.
	InstanceId int `pulumi:"instanceId"`
}

// The set of arguments for constructing a ExtraDiskSize resource.
type ExtraDiskSizeArgs struct {
	// Extra disk size in GB. Supported values: 25, 50, 100, 250, 500, 1000, 2000
	ExtraDiskSize pulumi.IntInput
	// The CloudAMQP instance ID.
	InstanceId pulumi.IntInput
}

func (ExtraDiskSizeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*extraDiskSizeArgs)(nil)).Elem()
}

type ExtraDiskSizeInput interface {
	pulumi.Input

	ToExtraDiskSizeOutput() ExtraDiskSizeOutput
	ToExtraDiskSizeOutputWithContext(ctx context.Context) ExtraDiskSizeOutput
}

func (*ExtraDiskSize) ElementType() reflect.Type {
	return reflect.TypeOf((**ExtraDiskSize)(nil)).Elem()
}

func (i *ExtraDiskSize) ToExtraDiskSizeOutput() ExtraDiskSizeOutput {
	return i.ToExtraDiskSizeOutputWithContext(context.Background())
}

func (i *ExtraDiskSize) ToExtraDiskSizeOutputWithContext(ctx context.Context) ExtraDiskSizeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExtraDiskSizeOutput)
}

// ExtraDiskSizeArrayInput is an input type that accepts ExtraDiskSizeArray and ExtraDiskSizeArrayOutput values.
// You can construct a concrete instance of `ExtraDiskSizeArrayInput` via:
//
//	ExtraDiskSizeArray{ ExtraDiskSizeArgs{...} }
type ExtraDiskSizeArrayInput interface {
	pulumi.Input

	ToExtraDiskSizeArrayOutput() ExtraDiskSizeArrayOutput
	ToExtraDiskSizeArrayOutputWithContext(context.Context) ExtraDiskSizeArrayOutput
}

type ExtraDiskSizeArray []ExtraDiskSizeInput

func (ExtraDiskSizeArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ExtraDiskSize)(nil)).Elem()
}

func (i ExtraDiskSizeArray) ToExtraDiskSizeArrayOutput() ExtraDiskSizeArrayOutput {
	return i.ToExtraDiskSizeArrayOutputWithContext(context.Background())
}

func (i ExtraDiskSizeArray) ToExtraDiskSizeArrayOutputWithContext(ctx context.Context) ExtraDiskSizeArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExtraDiskSizeArrayOutput)
}

// ExtraDiskSizeMapInput is an input type that accepts ExtraDiskSizeMap and ExtraDiskSizeMapOutput values.
// You can construct a concrete instance of `ExtraDiskSizeMapInput` via:
//
//	ExtraDiskSizeMap{ "key": ExtraDiskSizeArgs{...} }
type ExtraDiskSizeMapInput interface {
	pulumi.Input

	ToExtraDiskSizeMapOutput() ExtraDiskSizeMapOutput
	ToExtraDiskSizeMapOutputWithContext(context.Context) ExtraDiskSizeMapOutput
}

type ExtraDiskSizeMap map[string]ExtraDiskSizeInput

func (ExtraDiskSizeMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ExtraDiskSize)(nil)).Elem()
}

func (i ExtraDiskSizeMap) ToExtraDiskSizeMapOutput() ExtraDiskSizeMapOutput {
	return i.ToExtraDiskSizeMapOutputWithContext(context.Background())
}

func (i ExtraDiskSizeMap) ToExtraDiskSizeMapOutputWithContext(ctx context.Context) ExtraDiskSizeMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExtraDiskSizeMapOutput)
}

type ExtraDiskSizeOutput struct{ *pulumi.OutputState }

func (ExtraDiskSizeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ExtraDiskSize)(nil)).Elem()
}

func (o ExtraDiskSizeOutput) ToExtraDiskSizeOutput() ExtraDiskSizeOutput {
	return o
}

func (o ExtraDiskSizeOutput) ToExtraDiskSizeOutputWithContext(ctx context.Context) ExtraDiskSizeOutput {
	return o
}

// Extra disk size in GB. Supported values: 25, 50, 100, 250, 500, 1000, 2000
func (o ExtraDiskSizeOutput) ExtraDiskSize() pulumi.IntOutput {
	return o.ApplyT(func(v *ExtraDiskSize) pulumi.IntOutput { return v.ExtraDiskSize }).(pulumi.IntOutput)
}

// The CloudAMQP instance ID.
func (o ExtraDiskSizeOutput) InstanceId() pulumi.IntOutput {
	return o.ApplyT(func(v *ExtraDiskSize) pulumi.IntOutput { return v.InstanceId }).(pulumi.IntOutput)
}

type ExtraDiskSizeArrayOutput struct{ *pulumi.OutputState }

func (ExtraDiskSizeArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ExtraDiskSize)(nil)).Elem()
}

func (o ExtraDiskSizeArrayOutput) ToExtraDiskSizeArrayOutput() ExtraDiskSizeArrayOutput {
	return o
}

func (o ExtraDiskSizeArrayOutput) ToExtraDiskSizeArrayOutputWithContext(ctx context.Context) ExtraDiskSizeArrayOutput {
	return o
}

func (o ExtraDiskSizeArrayOutput) Index(i pulumi.IntInput) ExtraDiskSizeOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ExtraDiskSize {
		return vs[0].([]*ExtraDiskSize)[vs[1].(int)]
	}).(ExtraDiskSizeOutput)
}

type ExtraDiskSizeMapOutput struct{ *pulumi.OutputState }

func (ExtraDiskSizeMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ExtraDiskSize)(nil)).Elem()
}

func (o ExtraDiskSizeMapOutput) ToExtraDiskSizeMapOutput() ExtraDiskSizeMapOutput {
	return o
}

func (o ExtraDiskSizeMapOutput) ToExtraDiskSizeMapOutputWithContext(ctx context.Context) ExtraDiskSizeMapOutput {
	return o
}

func (o ExtraDiskSizeMapOutput) MapIndex(k pulumi.StringInput) ExtraDiskSizeOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ExtraDiskSize {
		return vs[0].(map[string]*ExtraDiskSize)[vs[1].(string)]
	}).(ExtraDiskSizeOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ExtraDiskSizeInput)(nil)).Elem(), &ExtraDiskSize{})
	pulumi.RegisterInputType(reflect.TypeOf((*ExtraDiskSizeArrayInput)(nil)).Elem(), ExtraDiskSizeArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ExtraDiskSizeMapInput)(nil)).Elem(), ExtraDiskSizeMap{})
	pulumi.RegisterOutputType(ExtraDiskSizeOutput{})
	pulumi.RegisterOutputType(ExtraDiskSizeArrayOutput{})
	pulumi.RegisterOutputType(ExtraDiskSizeMapOutput{})
}
