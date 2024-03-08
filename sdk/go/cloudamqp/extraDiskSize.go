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

// This resource allows you to resize the disk with additional storage capacity.
//
// ***Pre v1.25.0***: Only available for Amazon Web Services (AWS) and it done without downtime
//
// ***Post v1.25.0***: Now also available for Google Compute Engine (GCE) and Azure.
//
// Introducing a new optional argument called `allowDowntime`.  Leaving it out or set it to false will proceed to try and resize the disk without downtime, available for *AWS* and *GCE*.
// While *Azure* only support swapping the disk, and this argument needs to be set to *true*.
//
// `allowDowntime` also makes it possible to circumvent the time rate limit or shrinking the disk.
//
// | Cloud Platform        | allow_downtime=false | allow_downtime=true           |
// |-----------------------|----------------------|-------------------------------|
// | amazon-web-services   | Expand current disk* | Try to expand, otherwise swap |
// | google-compute-engine | Expand current disk* | Try to expand, otherwise swap |
// | azure-arm             | Not supported        | Swap disk to new size         |
//
// *Preferable method to use.
//
// > **WARNING:** Due to restrictions from cloud providers, it's only possible to resize the disk every 8 hours. Unless the `allow_downtime=true` is set, then the disk will be swapped for a new.
//
// Pricing is available at [cloudamqp.com](https://www.cloudamqp.com/) and only available for dedicated subscription plans.
//
// ## Example Usage
//
// <details>
//
//	<summary>
//	  <b>
//	    <i>AWS extra disk size (pre v1.25.0)</i>
//	  </b>
//	</summary>
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
//			// Instance
//			instance, err := cloudamqp.NewInstance(ctx, "instance", &cloudamqp.InstanceArgs{
//				Plan:   pulumi.String("bunny-1"),
//				Region: pulumi.String("amazon-web-services::us-west-2"),
//			})
//			if err != nil {
//				return err
//			}
//			// Resize disk with 25 extra GB
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
// <!--End PulumiCodeChooser -->
//
// </details>
//
// <details>
//
//	<summary>
//	  <b>
//	    <i>AWS extra disk size without downtime</i>
//	  </b>
//	</summary>
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
//			// Instance
//			instance, err := cloudamqp.NewInstance(ctx, "instance", &cloudamqp.InstanceArgs{
//				Plan:   pulumi.String("bunny-1"),
//				Region: pulumi.String("amazon-web-services::us-west-2"),
//			})
//			if err != nil {
//				return err
//			}
//			// Resize disk with 25 extra GB, without downtime
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
// <!--End PulumiCodeChooser -->
//
// </details>
//
// <details>
//
//	<summary>
//	  <b>
//	    <i>GCE extra disk size without downtime</i>
//	  </b>
//	</summary>
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
//			// Instance
//			instance, err := cloudamqp.NewInstance(ctx, "instance", &cloudamqp.InstanceArgs{
//				Plan:   pulumi.String("bunny-1"),
//				Region: pulumi.String("google-compute-engine::us-central1"),
//			})
//			if err != nil {
//				return err
//			}
//			// Resize disk with 25 extra GB, without downtime
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
// <!--End PulumiCodeChooser -->
//
// </details>
//
// <details>
//
//	<summary>
//	  <b>
//	    <i>Azure extra disk size with downtime</i>
//	  </b>
//	</summary>
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
//			// Instance
//			instance, err := cloudamqp.NewInstance(ctx, "instance", &cloudamqp.InstanceArgs{
//				Plan:   pulumi.String("bunny-1"),
//				Region: pulumi.String("azure-arm::centralus"),
//			})
//			if err != nil {
//				return err
//			}
//			// Resize disk with 25 extra GB, with downtime
//			_, err = cloudamqp.NewExtraDiskSize(ctx, "resizeDisk", &cloudamqp.ExtraDiskSizeArgs{
//				InstanceId:    instance.ID(),
//				ExtraDiskSize: pulumi.Int(25),
//				AllowDowntime: pulumi.Bool(true),
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
// <!--End PulumiCodeChooser -->
//
// </details>
//
// ## Attributes reference
//
// # All attributes reference are computed
//
// * `id`    - The identifier for this resource.
// * `nodes` - An array of node information. Each `nodes` block consists of the fields documented below.
//
// ***
//
// # The `nodes` block consist of
//
// * `name`                  - Name of the node.
// * `diskSize`             - Subscription plan disk size
// * `additionalDiskSize`  - Additional added disk size
//
// ***Note:*** *Total disk size = diskSize + additional_disk_size*
//
// ## Dependency
//
// This data source depends on CloudAMQP instance identifier, `cloudamqp_instance.instance.id`.
//
// ## Import
//
// Not possible to import this resource.
type ExtraDiskSize struct {
	pulumi.CustomResourceState

	// When resizing the disk, allow cluster downtime if necessary. Default set to false. Required when hosting in *Azure*.
	AllowDowntime pulumi.BoolPtrOutput `pulumi:"allowDowntime"`
	// Extra disk size in GB. Supported values: 0, 25, 50, 100, 250, 500, 1000, 2000
	ExtraDiskSize pulumi.IntOutput `pulumi:"extraDiskSize"`
	// The CloudAMQP instance ID.
	InstanceId pulumi.IntOutput             `pulumi:"instanceId"`
	Nodes      ExtraDiskSizeNodeArrayOutput `pulumi:"nodes"`
	// Configurable sleep time in seconds between retries for resizing the disk. Default set to 30 seconds.
	Sleep pulumi.IntPtrOutput `pulumi:"sleep"`
	// Configurable timeout time in seconds for resizing the disk. Default set to 1800 seconds.
	//
	// ***Note:*** `allowDowntime`, `sleep`, `timeout` only available from v1.25.0.
	Timeout pulumi.IntPtrOutput `pulumi:"timeout"`
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
	opts = internal.PkgResourceDefaultOpts(opts)
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
	// When resizing the disk, allow cluster downtime if necessary. Default set to false. Required when hosting in *Azure*.
	AllowDowntime *bool `pulumi:"allowDowntime"`
	// Extra disk size in GB. Supported values: 0, 25, 50, 100, 250, 500, 1000, 2000
	ExtraDiskSize *int `pulumi:"extraDiskSize"`
	// The CloudAMQP instance ID.
	InstanceId *int                `pulumi:"instanceId"`
	Nodes      []ExtraDiskSizeNode `pulumi:"nodes"`
	// Configurable sleep time in seconds between retries for resizing the disk. Default set to 30 seconds.
	Sleep *int `pulumi:"sleep"`
	// Configurable timeout time in seconds for resizing the disk. Default set to 1800 seconds.
	//
	// ***Note:*** `allowDowntime`, `sleep`, `timeout` only available from v1.25.0.
	Timeout *int `pulumi:"timeout"`
}

type ExtraDiskSizeState struct {
	// When resizing the disk, allow cluster downtime if necessary. Default set to false. Required when hosting in *Azure*.
	AllowDowntime pulumi.BoolPtrInput
	// Extra disk size in GB. Supported values: 0, 25, 50, 100, 250, 500, 1000, 2000
	ExtraDiskSize pulumi.IntPtrInput
	// The CloudAMQP instance ID.
	InstanceId pulumi.IntPtrInput
	Nodes      ExtraDiskSizeNodeArrayInput
	// Configurable sleep time in seconds between retries for resizing the disk. Default set to 30 seconds.
	Sleep pulumi.IntPtrInput
	// Configurable timeout time in seconds for resizing the disk. Default set to 1800 seconds.
	//
	// ***Note:*** `allowDowntime`, `sleep`, `timeout` only available from v1.25.0.
	Timeout pulumi.IntPtrInput
}

func (ExtraDiskSizeState) ElementType() reflect.Type {
	return reflect.TypeOf((*extraDiskSizeState)(nil)).Elem()
}

type extraDiskSizeArgs struct {
	// When resizing the disk, allow cluster downtime if necessary. Default set to false. Required when hosting in *Azure*.
	AllowDowntime *bool `pulumi:"allowDowntime"`
	// Extra disk size in GB. Supported values: 0, 25, 50, 100, 250, 500, 1000, 2000
	ExtraDiskSize int `pulumi:"extraDiskSize"`
	// The CloudAMQP instance ID.
	InstanceId int `pulumi:"instanceId"`
	// Configurable sleep time in seconds between retries for resizing the disk. Default set to 30 seconds.
	Sleep *int `pulumi:"sleep"`
	// Configurable timeout time in seconds for resizing the disk. Default set to 1800 seconds.
	//
	// ***Note:*** `allowDowntime`, `sleep`, `timeout` only available from v1.25.0.
	Timeout *int `pulumi:"timeout"`
}

// The set of arguments for constructing a ExtraDiskSize resource.
type ExtraDiskSizeArgs struct {
	// When resizing the disk, allow cluster downtime if necessary. Default set to false. Required when hosting in *Azure*.
	AllowDowntime pulumi.BoolPtrInput
	// Extra disk size in GB. Supported values: 0, 25, 50, 100, 250, 500, 1000, 2000
	ExtraDiskSize pulumi.IntInput
	// The CloudAMQP instance ID.
	InstanceId pulumi.IntInput
	// Configurable sleep time in seconds between retries for resizing the disk. Default set to 30 seconds.
	Sleep pulumi.IntPtrInput
	// Configurable timeout time in seconds for resizing the disk. Default set to 1800 seconds.
	//
	// ***Note:*** `allowDowntime`, `sleep`, `timeout` only available from v1.25.0.
	Timeout pulumi.IntPtrInput
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

// When resizing the disk, allow cluster downtime if necessary. Default set to false. Required when hosting in *Azure*.
func (o ExtraDiskSizeOutput) AllowDowntime() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ExtraDiskSize) pulumi.BoolPtrOutput { return v.AllowDowntime }).(pulumi.BoolPtrOutput)
}

// Extra disk size in GB. Supported values: 0, 25, 50, 100, 250, 500, 1000, 2000
func (o ExtraDiskSizeOutput) ExtraDiskSize() pulumi.IntOutput {
	return o.ApplyT(func(v *ExtraDiskSize) pulumi.IntOutput { return v.ExtraDiskSize }).(pulumi.IntOutput)
}

// The CloudAMQP instance ID.
func (o ExtraDiskSizeOutput) InstanceId() pulumi.IntOutput {
	return o.ApplyT(func(v *ExtraDiskSize) pulumi.IntOutput { return v.InstanceId }).(pulumi.IntOutput)
}

func (o ExtraDiskSizeOutput) Nodes() ExtraDiskSizeNodeArrayOutput {
	return o.ApplyT(func(v *ExtraDiskSize) ExtraDiskSizeNodeArrayOutput { return v.Nodes }).(ExtraDiskSizeNodeArrayOutput)
}

// Configurable sleep time in seconds between retries for resizing the disk. Default set to 30 seconds.
func (o ExtraDiskSizeOutput) Sleep() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ExtraDiskSize) pulumi.IntPtrOutput { return v.Sleep }).(pulumi.IntPtrOutput)
}

// Configurable timeout time in seconds for resizing the disk. Default set to 1800 seconds.
//
// ***Note:*** `allowDowntime`, `sleep`, `timeout` only available from v1.25.0.
func (o ExtraDiskSizeOutput) Timeout() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ExtraDiskSize) pulumi.IntPtrOutput { return v.Timeout }).(pulumi.IntPtrOutput)
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
