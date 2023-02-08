// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cloudamqp

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This resource allows you to automatically upgrade to latest possible upgradable versions for RabbitMQ and Erlang. Depending on initial versions of RabbitMQ and Erlang of the CloudAMQP instance, multiple runs may be needed to get to latest versions. After completed upgrade, check data source `getUpgradableVersions` to see if newer versions is available. Then delete `UpgradeRabbitmq` and create it again to invoke the upgrade.
//
// > **Important Upgrade Information**
// > - All nodes in a cluster must run the same major and minor version of RabbitMQ. The entire cluster will be offline while upgrading major or minor versions.
// > - Auto delete queues (queues that are marked AD) will be deleted during the update.
// > - Any custom plugins support has installed on your behalf will be disabled and you need to contact support@cloudamqp.com and ask to have them re-installed.
// > - TLS 1.0 and 1.1 will not be supported after the update.
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
//			_, err := cloudamqp.GetUpgradableVersions(ctx, &cloudamqp.GetUpgradableVersionsArgs{
//				InstanceId: cloudamqp_instance.Instance.Id,
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = cloudamqp.NewUpgradeRabbitmq(ctx, "upgrade", &cloudamqp.UpgradeRabbitmqArgs{
//				InstanceId: pulumi.Any(cloudamqp_instance.Instance.Id),
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
//			_, err := cloudamqp.GetUpgradableVersions(ctx, &cloudamqp.GetUpgradableVersionsArgs{
//				InstanceId: cloudamqp_instance.Instance.Id,
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// If newer version is still available to be upgradable in the data source, re-run again.
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
//			_, err := cloudamqp.GetUpgradableVersions(ctx, &cloudamqp.GetUpgradableVersionsArgs{
//				InstanceId: cloudamqp_instance.Instance.Id,
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = cloudamqp.NewUpgradeRabbitmq(ctx, "upgrade", &cloudamqp.UpgradeRabbitmqArgs{
//				InstanceId: pulumi.Any(cloudamqp_instance.Instance.Id),
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
// ## Import
//
// Not possible to import this resource.
type UpgradeRabbitmq struct {
	pulumi.CustomResourceState

	// The CloudAMQP instance identifier
	InstanceId pulumi.IntOutput `pulumi:"instanceId"`
}

// NewUpgradeRabbitmq registers a new resource with the given unique name, arguments, and options.
func NewUpgradeRabbitmq(ctx *pulumi.Context,
	name string, args *UpgradeRabbitmqArgs, opts ...pulumi.ResourceOption) (*UpgradeRabbitmq, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.InstanceId == nil {
		return nil, errors.New("invalid value for required argument 'InstanceId'")
	}
	var resource UpgradeRabbitmq
	err := ctx.RegisterResource("cloudamqp:index/upgradeRabbitmq:UpgradeRabbitmq", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetUpgradeRabbitmq gets an existing UpgradeRabbitmq resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetUpgradeRabbitmq(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *UpgradeRabbitmqState, opts ...pulumi.ResourceOption) (*UpgradeRabbitmq, error) {
	var resource UpgradeRabbitmq
	err := ctx.ReadResource("cloudamqp:index/upgradeRabbitmq:UpgradeRabbitmq", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering UpgradeRabbitmq resources.
type upgradeRabbitmqState struct {
	// The CloudAMQP instance identifier
	InstanceId *int `pulumi:"instanceId"`
}

type UpgradeRabbitmqState struct {
	// The CloudAMQP instance identifier
	InstanceId pulumi.IntPtrInput
}

func (UpgradeRabbitmqState) ElementType() reflect.Type {
	return reflect.TypeOf((*upgradeRabbitmqState)(nil)).Elem()
}

type upgradeRabbitmqArgs struct {
	// The CloudAMQP instance identifier
	InstanceId int `pulumi:"instanceId"`
}

// The set of arguments for constructing a UpgradeRabbitmq resource.
type UpgradeRabbitmqArgs struct {
	// The CloudAMQP instance identifier
	InstanceId pulumi.IntInput
}

func (UpgradeRabbitmqArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*upgradeRabbitmqArgs)(nil)).Elem()
}

type UpgradeRabbitmqInput interface {
	pulumi.Input

	ToUpgradeRabbitmqOutput() UpgradeRabbitmqOutput
	ToUpgradeRabbitmqOutputWithContext(ctx context.Context) UpgradeRabbitmqOutput
}

func (*UpgradeRabbitmq) ElementType() reflect.Type {
	return reflect.TypeOf((**UpgradeRabbitmq)(nil)).Elem()
}

func (i *UpgradeRabbitmq) ToUpgradeRabbitmqOutput() UpgradeRabbitmqOutput {
	return i.ToUpgradeRabbitmqOutputWithContext(context.Background())
}

func (i *UpgradeRabbitmq) ToUpgradeRabbitmqOutputWithContext(ctx context.Context) UpgradeRabbitmqOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UpgradeRabbitmqOutput)
}

// UpgradeRabbitmqArrayInput is an input type that accepts UpgradeRabbitmqArray and UpgradeRabbitmqArrayOutput values.
// You can construct a concrete instance of `UpgradeRabbitmqArrayInput` via:
//
//	UpgradeRabbitmqArray{ UpgradeRabbitmqArgs{...} }
type UpgradeRabbitmqArrayInput interface {
	pulumi.Input

	ToUpgradeRabbitmqArrayOutput() UpgradeRabbitmqArrayOutput
	ToUpgradeRabbitmqArrayOutputWithContext(context.Context) UpgradeRabbitmqArrayOutput
}

type UpgradeRabbitmqArray []UpgradeRabbitmqInput

func (UpgradeRabbitmqArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*UpgradeRabbitmq)(nil)).Elem()
}

func (i UpgradeRabbitmqArray) ToUpgradeRabbitmqArrayOutput() UpgradeRabbitmqArrayOutput {
	return i.ToUpgradeRabbitmqArrayOutputWithContext(context.Background())
}

func (i UpgradeRabbitmqArray) ToUpgradeRabbitmqArrayOutputWithContext(ctx context.Context) UpgradeRabbitmqArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UpgradeRabbitmqArrayOutput)
}

// UpgradeRabbitmqMapInput is an input type that accepts UpgradeRabbitmqMap and UpgradeRabbitmqMapOutput values.
// You can construct a concrete instance of `UpgradeRabbitmqMapInput` via:
//
//	UpgradeRabbitmqMap{ "key": UpgradeRabbitmqArgs{...} }
type UpgradeRabbitmqMapInput interface {
	pulumi.Input

	ToUpgradeRabbitmqMapOutput() UpgradeRabbitmqMapOutput
	ToUpgradeRabbitmqMapOutputWithContext(context.Context) UpgradeRabbitmqMapOutput
}

type UpgradeRabbitmqMap map[string]UpgradeRabbitmqInput

func (UpgradeRabbitmqMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*UpgradeRabbitmq)(nil)).Elem()
}

func (i UpgradeRabbitmqMap) ToUpgradeRabbitmqMapOutput() UpgradeRabbitmqMapOutput {
	return i.ToUpgradeRabbitmqMapOutputWithContext(context.Background())
}

func (i UpgradeRabbitmqMap) ToUpgradeRabbitmqMapOutputWithContext(ctx context.Context) UpgradeRabbitmqMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UpgradeRabbitmqMapOutput)
}

type UpgradeRabbitmqOutput struct{ *pulumi.OutputState }

func (UpgradeRabbitmqOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**UpgradeRabbitmq)(nil)).Elem()
}

func (o UpgradeRabbitmqOutput) ToUpgradeRabbitmqOutput() UpgradeRabbitmqOutput {
	return o
}

func (o UpgradeRabbitmqOutput) ToUpgradeRabbitmqOutputWithContext(ctx context.Context) UpgradeRabbitmqOutput {
	return o
}

// The CloudAMQP instance identifier
func (o UpgradeRabbitmqOutput) InstanceId() pulumi.IntOutput {
	return o.ApplyT(func(v *UpgradeRabbitmq) pulumi.IntOutput { return v.InstanceId }).(pulumi.IntOutput)
}

type UpgradeRabbitmqArrayOutput struct{ *pulumi.OutputState }

func (UpgradeRabbitmqArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*UpgradeRabbitmq)(nil)).Elem()
}

func (o UpgradeRabbitmqArrayOutput) ToUpgradeRabbitmqArrayOutput() UpgradeRabbitmqArrayOutput {
	return o
}

func (o UpgradeRabbitmqArrayOutput) ToUpgradeRabbitmqArrayOutputWithContext(ctx context.Context) UpgradeRabbitmqArrayOutput {
	return o
}

func (o UpgradeRabbitmqArrayOutput) Index(i pulumi.IntInput) UpgradeRabbitmqOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *UpgradeRabbitmq {
		return vs[0].([]*UpgradeRabbitmq)[vs[1].(int)]
	}).(UpgradeRabbitmqOutput)
}

type UpgradeRabbitmqMapOutput struct{ *pulumi.OutputState }

func (UpgradeRabbitmqMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*UpgradeRabbitmq)(nil)).Elem()
}

func (o UpgradeRabbitmqMapOutput) ToUpgradeRabbitmqMapOutput() UpgradeRabbitmqMapOutput {
	return o
}

func (o UpgradeRabbitmqMapOutput) ToUpgradeRabbitmqMapOutputWithContext(ctx context.Context) UpgradeRabbitmqMapOutput {
	return o
}

func (o UpgradeRabbitmqMapOutput) MapIndex(k pulumi.StringInput) UpgradeRabbitmqOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *UpgradeRabbitmq {
		return vs[0].(map[string]*UpgradeRabbitmq)[vs[1].(string)]
	}).(UpgradeRabbitmqOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*UpgradeRabbitmqInput)(nil)).Elem(), &UpgradeRabbitmq{})
	pulumi.RegisterInputType(reflect.TypeOf((*UpgradeRabbitmqArrayInput)(nil)).Elem(), UpgradeRabbitmqArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*UpgradeRabbitmqMapInput)(nil)).Elem(), UpgradeRabbitmqMap{})
	pulumi.RegisterOutputType(UpgradeRabbitmqOutput{})
	pulumi.RegisterOutputType(UpgradeRabbitmqArrayOutput{})
	pulumi.RegisterOutputType(UpgradeRabbitmqMapOutput{})
}