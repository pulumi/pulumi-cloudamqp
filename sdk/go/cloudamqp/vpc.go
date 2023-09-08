// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cloudamqp

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-cloudamqp/sdk/v3/go/cloudamqp/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// This resource allows you to manage standalone VPC.
//
// New Cloudamqp instances can be added to the managed VPC. Set the instance *vpc_id* attribute to the managed vpc identifier , see example below, when creating the instance.
//
// Only available for dedicated subscription plans.
//
// Pricing is available at [cloudamqp.com](https://www.cloudamqp.com/plans.html).
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
//			vpc, err := cloudamqp.NewVpc(ctx, "vpc", &cloudamqp.VpcArgs{
//				Region: pulumi.String("amazon-web-services::us-east-1"),
//				Subnet: pulumi.String("10.56.72.0/24"),
//				Tags:   pulumi.StringArray{},
//			})
//			if err != nil {
//				return err
//			}
//			_, err = cloudamqp.NewInstance(ctx, "instance", &cloudamqp.InstanceArgs{
//				Plan:              pulumi.String("bunny-1"),
//				Region:            pulumi.String("amazon-web-services::us-east-1"),
//				Nodes:             pulumi.Int(1),
//				Tags:              pulumi.StringArray{},
//				RmqVersion:        pulumi.String("3.9.13"),
//				VpcId:             pulumi.Any(cloudamq_vpc.Vpc.Id),
//				KeepAssociatedVpc: pulumi.Bool(true),
//			})
//			if err != nil {
//				return err
//			}
//			_ = cloudamqp.GetVpcInfoOutput(ctx, cloudamqp.GetVpcInfoOutputArgs{
//				VpcId: vpc.ID(),
//			}, nil)
//			return nil
//		})
//	}
//
// ```
//
// ## Import
//
// `cloudamqp_vpc` can be imported using the CloudAMQP VPC identifier.
//
// ```sh
//
//	$ pulumi import cloudamqp:index/vpc:Vpc <resource_name> <vpc_id>`
//
// ```
//
//	To retrieve the identifier for a VPC, either use [CloudAMQP customer API](https://docs.cloudamqp.com/#list-vpcs). Or use the data source `cloudamqp_account_vpcs` to list all available standalone VPCs for an account.
type Vpc struct {
	pulumi.CustomResourceState

	// The name of the VPC.
	Name pulumi.StringOutput `pulumi:"name"`
	// The hosted region for the managed standalone VPC
	Region pulumi.StringOutput `pulumi:"region"`
	// The VPC subnet
	Subnet pulumi.StringOutput `pulumi:"subnet"`
	// Tag the VPC with optional tags
	Tags pulumi.StringArrayOutput `pulumi:"tags"`
	// VPC name given when hosted at the cloud provider
	VpcName pulumi.StringOutput `pulumi:"vpcName"`
}

// NewVpc registers a new resource with the given unique name, arguments, and options.
func NewVpc(ctx *pulumi.Context,
	name string, args *VpcArgs, opts ...pulumi.ResourceOption) (*Vpc, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Region == nil {
		return nil, errors.New("invalid value for required argument 'Region'")
	}
	if args.Subnet == nil {
		return nil, errors.New("invalid value for required argument 'Subnet'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Vpc
	err := ctx.RegisterResource("cloudamqp:index/vpc:Vpc", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVpc gets an existing Vpc resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVpc(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VpcState, opts ...pulumi.ResourceOption) (*Vpc, error) {
	var resource Vpc
	err := ctx.ReadResource("cloudamqp:index/vpc:Vpc", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Vpc resources.
type vpcState struct {
	// The name of the VPC.
	Name *string `pulumi:"name"`
	// The hosted region for the managed standalone VPC
	Region *string `pulumi:"region"`
	// The VPC subnet
	Subnet *string `pulumi:"subnet"`
	// Tag the VPC with optional tags
	Tags []string `pulumi:"tags"`
	// VPC name given when hosted at the cloud provider
	VpcName *string `pulumi:"vpcName"`
}

type VpcState struct {
	// The name of the VPC.
	Name pulumi.StringPtrInput
	// The hosted region for the managed standalone VPC
	Region pulumi.StringPtrInput
	// The VPC subnet
	Subnet pulumi.StringPtrInput
	// Tag the VPC with optional tags
	Tags pulumi.StringArrayInput
	// VPC name given when hosted at the cloud provider
	VpcName pulumi.StringPtrInput
}

func (VpcState) ElementType() reflect.Type {
	return reflect.TypeOf((*vpcState)(nil)).Elem()
}

type vpcArgs struct {
	// The name of the VPC.
	Name *string `pulumi:"name"`
	// The hosted region for the managed standalone VPC
	Region string `pulumi:"region"`
	// The VPC subnet
	Subnet string `pulumi:"subnet"`
	// Tag the VPC with optional tags
	Tags []string `pulumi:"tags"`
}

// The set of arguments for constructing a Vpc resource.
type VpcArgs struct {
	// The name of the VPC.
	Name pulumi.StringPtrInput
	// The hosted region for the managed standalone VPC
	Region pulumi.StringInput
	// The VPC subnet
	Subnet pulumi.StringInput
	// Tag the VPC with optional tags
	Tags pulumi.StringArrayInput
}

func (VpcArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*vpcArgs)(nil)).Elem()
}

type VpcInput interface {
	pulumi.Input

	ToVpcOutput() VpcOutput
	ToVpcOutputWithContext(ctx context.Context) VpcOutput
}

func (*Vpc) ElementType() reflect.Type {
	return reflect.TypeOf((**Vpc)(nil)).Elem()
}

func (i *Vpc) ToVpcOutput() VpcOutput {
	return i.ToVpcOutputWithContext(context.Background())
}

func (i *Vpc) ToVpcOutputWithContext(ctx context.Context) VpcOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpcOutput)
}

func (i *Vpc) ToOutput(ctx context.Context) pulumix.Output[*Vpc] {
	return pulumix.Output[*Vpc]{
		OutputState: i.ToVpcOutputWithContext(ctx).OutputState,
	}
}

// VpcArrayInput is an input type that accepts VpcArray and VpcArrayOutput values.
// You can construct a concrete instance of `VpcArrayInput` via:
//
//	VpcArray{ VpcArgs{...} }
type VpcArrayInput interface {
	pulumi.Input

	ToVpcArrayOutput() VpcArrayOutput
	ToVpcArrayOutputWithContext(context.Context) VpcArrayOutput
}

type VpcArray []VpcInput

func (VpcArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Vpc)(nil)).Elem()
}

func (i VpcArray) ToVpcArrayOutput() VpcArrayOutput {
	return i.ToVpcArrayOutputWithContext(context.Background())
}

func (i VpcArray) ToVpcArrayOutputWithContext(ctx context.Context) VpcArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpcArrayOutput)
}

func (i VpcArray) ToOutput(ctx context.Context) pulumix.Output[[]*Vpc] {
	return pulumix.Output[[]*Vpc]{
		OutputState: i.ToVpcArrayOutputWithContext(ctx).OutputState,
	}
}

// VpcMapInput is an input type that accepts VpcMap and VpcMapOutput values.
// You can construct a concrete instance of `VpcMapInput` via:
//
//	VpcMap{ "key": VpcArgs{...} }
type VpcMapInput interface {
	pulumi.Input

	ToVpcMapOutput() VpcMapOutput
	ToVpcMapOutputWithContext(context.Context) VpcMapOutput
}

type VpcMap map[string]VpcInput

func (VpcMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Vpc)(nil)).Elem()
}

func (i VpcMap) ToVpcMapOutput() VpcMapOutput {
	return i.ToVpcMapOutputWithContext(context.Background())
}

func (i VpcMap) ToVpcMapOutputWithContext(ctx context.Context) VpcMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpcMapOutput)
}

func (i VpcMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*Vpc] {
	return pulumix.Output[map[string]*Vpc]{
		OutputState: i.ToVpcMapOutputWithContext(ctx).OutputState,
	}
}

type VpcOutput struct{ *pulumi.OutputState }

func (VpcOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Vpc)(nil)).Elem()
}

func (o VpcOutput) ToVpcOutput() VpcOutput {
	return o
}

func (o VpcOutput) ToVpcOutputWithContext(ctx context.Context) VpcOutput {
	return o
}

func (o VpcOutput) ToOutput(ctx context.Context) pulumix.Output[*Vpc] {
	return pulumix.Output[*Vpc]{
		OutputState: o.OutputState,
	}
}

// The name of the VPC.
func (o VpcOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Vpc) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The hosted region for the managed standalone VPC
func (o VpcOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *Vpc) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

// The VPC subnet
func (o VpcOutput) Subnet() pulumi.StringOutput {
	return o.ApplyT(func(v *Vpc) pulumi.StringOutput { return v.Subnet }).(pulumi.StringOutput)
}

// Tag the VPC with optional tags
func (o VpcOutput) Tags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Vpc) pulumi.StringArrayOutput { return v.Tags }).(pulumi.StringArrayOutput)
}

// VPC name given when hosted at the cloud provider
func (o VpcOutput) VpcName() pulumi.StringOutput {
	return o.ApplyT(func(v *Vpc) pulumi.StringOutput { return v.VpcName }).(pulumi.StringOutput)
}

type VpcArrayOutput struct{ *pulumi.OutputState }

func (VpcArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Vpc)(nil)).Elem()
}

func (o VpcArrayOutput) ToVpcArrayOutput() VpcArrayOutput {
	return o
}

func (o VpcArrayOutput) ToVpcArrayOutputWithContext(ctx context.Context) VpcArrayOutput {
	return o
}

func (o VpcArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*Vpc] {
	return pulumix.Output[[]*Vpc]{
		OutputState: o.OutputState,
	}
}

func (o VpcArrayOutput) Index(i pulumi.IntInput) VpcOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Vpc {
		return vs[0].([]*Vpc)[vs[1].(int)]
	}).(VpcOutput)
}

type VpcMapOutput struct{ *pulumi.OutputState }

func (VpcMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Vpc)(nil)).Elem()
}

func (o VpcMapOutput) ToVpcMapOutput() VpcMapOutput {
	return o
}

func (o VpcMapOutput) ToVpcMapOutputWithContext(ctx context.Context) VpcMapOutput {
	return o
}

func (o VpcMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*Vpc] {
	return pulumix.Output[map[string]*Vpc]{
		OutputState: o.OutputState,
	}
}

func (o VpcMapOutput) MapIndex(k pulumi.StringInput) VpcOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Vpc {
		return vs[0].(map[string]*Vpc)[vs[1].(string)]
	}).(VpcOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*VpcInput)(nil)).Elem(), &Vpc{})
	pulumi.RegisterInputType(reflect.TypeOf((*VpcArrayInput)(nil)).Elem(), VpcArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*VpcMapInput)(nil)).Elem(), VpcMap{})
	pulumi.RegisterOutputType(VpcOutput{})
	pulumi.RegisterOutputType(VpcArrayOutput{})
	pulumi.RegisterOutputType(VpcMapOutput{})
}
