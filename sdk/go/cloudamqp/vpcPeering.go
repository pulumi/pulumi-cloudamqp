// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cloudamqp

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## Import
//
// *Pre v1.16.0* `cloudamqp_vpc_peering` can be imported using the CloudAMQP instance identifier.
//
// ```sh
//  $ pulumi import cloudamqp:index/vpcPeering:VpcPeering <resource_name> <instance_id>`
// ```
//
//  *Post v1.16.0* `cloudamqp_vpc_peering` can be imported using the CloudAMQP managed VPC identifier or instance identifier.
//
// ```sh
//  $ pulumi import cloudamqp:index/vpcPeering:VpcPeering <resource_name> <vpc_id>`
// ```
//
// ```sh
//  $ pulumi import cloudamqp:index/vpcPeering:VpcPeering <resource_name> <instance_id>`
// ```
type VpcPeering struct {
	pulumi.CustomResourceState

	// The CloudAMQP instance identifier.
	InstanceId pulumi.IntPtrOutput `pulumi:"instanceId"`
	// Peering identifier created by AW peering request.
	PeeringId pulumi.StringOutput `pulumi:"peeringId"`
	// Configurable sleep time (seconds) between retries for accepting or removing peering. Default set to 60 seconds.
	Sleep pulumi.IntPtrOutput `pulumi:"sleep"`
	// VPC peering status
	Status pulumi.StringOutput `pulumi:"status"`
	// - Configurable timeout time (seconds) for accepting or removing peering. Default set to 3600 seconds.
	Timeout pulumi.IntPtrOutput `pulumi:"timeout"`
	// The managed VPC identifier.
	VpcId pulumi.StringPtrOutput `pulumi:"vpcId"`
}

// NewVpcPeering registers a new resource with the given unique name, arguments, and options.
func NewVpcPeering(ctx *pulumi.Context,
	name string, args *VpcPeeringArgs, opts ...pulumi.ResourceOption) (*VpcPeering, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.PeeringId == nil {
		return nil, errors.New("invalid value for required argument 'PeeringId'")
	}
	var resource VpcPeering
	err := ctx.RegisterResource("cloudamqp:index/vpcPeering:VpcPeering", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVpcPeering gets an existing VpcPeering resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVpcPeering(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VpcPeeringState, opts ...pulumi.ResourceOption) (*VpcPeering, error) {
	var resource VpcPeering
	err := ctx.ReadResource("cloudamqp:index/vpcPeering:VpcPeering", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VpcPeering resources.
type vpcPeeringState struct {
	// The CloudAMQP instance identifier.
	InstanceId *int `pulumi:"instanceId"`
	// Peering identifier created by AW peering request.
	PeeringId *string `pulumi:"peeringId"`
	// Configurable sleep time (seconds) between retries for accepting or removing peering. Default set to 60 seconds.
	Sleep *int `pulumi:"sleep"`
	// VPC peering status
	Status *string `pulumi:"status"`
	// - Configurable timeout time (seconds) for accepting or removing peering. Default set to 3600 seconds.
	Timeout *int `pulumi:"timeout"`
	// The managed VPC identifier.
	VpcId *string `pulumi:"vpcId"`
}

type VpcPeeringState struct {
	// The CloudAMQP instance identifier.
	InstanceId pulumi.IntPtrInput
	// Peering identifier created by AW peering request.
	PeeringId pulumi.StringPtrInput
	// Configurable sleep time (seconds) between retries for accepting or removing peering. Default set to 60 seconds.
	Sleep pulumi.IntPtrInput
	// VPC peering status
	Status pulumi.StringPtrInput
	// - Configurable timeout time (seconds) for accepting or removing peering. Default set to 3600 seconds.
	Timeout pulumi.IntPtrInput
	// The managed VPC identifier.
	VpcId pulumi.StringPtrInput
}

func (VpcPeeringState) ElementType() reflect.Type {
	return reflect.TypeOf((*vpcPeeringState)(nil)).Elem()
}

type vpcPeeringArgs struct {
	// The CloudAMQP instance identifier.
	InstanceId *int `pulumi:"instanceId"`
	// Peering identifier created by AW peering request.
	PeeringId string `pulumi:"peeringId"`
	// Configurable sleep time (seconds) between retries for accepting or removing peering. Default set to 60 seconds.
	Sleep *int `pulumi:"sleep"`
	// - Configurable timeout time (seconds) for accepting or removing peering. Default set to 3600 seconds.
	Timeout *int `pulumi:"timeout"`
	// The managed VPC identifier.
	VpcId *string `pulumi:"vpcId"`
}

// The set of arguments for constructing a VpcPeering resource.
type VpcPeeringArgs struct {
	// The CloudAMQP instance identifier.
	InstanceId pulumi.IntPtrInput
	// Peering identifier created by AW peering request.
	PeeringId pulumi.StringInput
	// Configurable sleep time (seconds) between retries for accepting or removing peering. Default set to 60 seconds.
	Sleep pulumi.IntPtrInput
	// - Configurable timeout time (seconds) for accepting or removing peering. Default set to 3600 seconds.
	Timeout pulumi.IntPtrInput
	// The managed VPC identifier.
	VpcId pulumi.StringPtrInput
}

func (VpcPeeringArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*vpcPeeringArgs)(nil)).Elem()
}

type VpcPeeringInput interface {
	pulumi.Input

	ToVpcPeeringOutput() VpcPeeringOutput
	ToVpcPeeringOutputWithContext(ctx context.Context) VpcPeeringOutput
}

func (*VpcPeering) ElementType() reflect.Type {
	return reflect.TypeOf((**VpcPeering)(nil)).Elem()
}

func (i *VpcPeering) ToVpcPeeringOutput() VpcPeeringOutput {
	return i.ToVpcPeeringOutputWithContext(context.Background())
}

func (i *VpcPeering) ToVpcPeeringOutputWithContext(ctx context.Context) VpcPeeringOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpcPeeringOutput)
}

// VpcPeeringArrayInput is an input type that accepts VpcPeeringArray and VpcPeeringArrayOutput values.
// You can construct a concrete instance of `VpcPeeringArrayInput` via:
//
//          VpcPeeringArray{ VpcPeeringArgs{...} }
type VpcPeeringArrayInput interface {
	pulumi.Input

	ToVpcPeeringArrayOutput() VpcPeeringArrayOutput
	ToVpcPeeringArrayOutputWithContext(context.Context) VpcPeeringArrayOutput
}

type VpcPeeringArray []VpcPeeringInput

func (VpcPeeringArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VpcPeering)(nil)).Elem()
}

func (i VpcPeeringArray) ToVpcPeeringArrayOutput() VpcPeeringArrayOutput {
	return i.ToVpcPeeringArrayOutputWithContext(context.Background())
}

func (i VpcPeeringArray) ToVpcPeeringArrayOutputWithContext(ctx context.Context) VpcPeeringArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpcPeeringArrayOutput)
}

// VpcPeeringMapInput is an input type that accepts VpcPeeringMap and VpcPeeringMapOutput values.
// You can construct a concrete instance of `VpcPeeringMapInput` via:
//
//          VpcPeeringMap{ "key": VpcPeeringArgs{...} }
type VpcPeeringMapInput interface {
	pulumi.Input

	ToVpcPeeringMapOutput() VpcPeeringMapOutput
	ToVpcPeeringMapOutputWithContext(context.Context) VpcPeeringMapOutput
}

type VpcPeeringMap map[string]VpcPeeringInput

func (VpcPeeringMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VpcPeering)(nil)).Elem()
}

func (i VpcPeeringMap) ToVpcPeeringMapOutput() VpcPeeringMapOutput {
	return i.ToVpcPeeringMapOutputWithContext(context.Background())
}

func (i VpcPeeringMap) ToVpcPeeringMapOutputWithContext(ctx context.Context) VpcPeeringMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpcPeeringMapOutput)
}

type VpcPeeringOutput struct{ *pulumi.OutputState }

func (VpcPeeringOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VpcPeering)(nil)).Elem()
}

func (o VpcPeeringOutput) ToVpcPeeringOutput() VpcPeeringOutput {
	return o
}

func (o VpcPeeringOutput) ToVpcPeeringOutputWithContext(ctx context.Context) VpcPeeringOutput {
	return o
}

type VpcPeeringArrayOutput struct{ *pulumi.OutputState }

func (VpcPeeringArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VpcPeering)(nil)).Elem()
}

func (o VpcPeeringArrayOutput) ToVpcPeeringArrayOutput() VpcPeeringArrayOutput {
	return o
}

func (o VpcPeeringArrayOutput) ToVpcPeeringArrayOutputWithContext(ctx context.Context) VpcPeeringArrayOutput {
	return o
}

func (o VpcPeeringArrayOutput) Index(i pulumi.IntInput) VpcPeeringOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *VpcPeering {
		return vs[0].([]*VpcPeering)[vs[1].(int)]
	}).(VpcPeeringOutput)
}

type VpcPeeringMapOutput struct{ *pulumi.OutputState }

func (VpcPeeringMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VpcPeering)(nil)).Elem()
}

func (o VpcPeeringMapOutput) ToVpcPeeringMapOutput() VpcPeeringMapOutput {
	return o
}

func (o VpcPeeringMapOutput) ToVpcPeeringMapOutputWithContext(ctx context.Context) VpcPeeringMapOutput {
	return o
}

func (o VpcPeeringMapOutput) MapIndex(k pulumi.StringInput) VpcPeeringOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *VpcPeering {
		return vs[0].(map[string]*VpcPeering)[vs[1].(string)]
	}).(VpcPeeringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*VpcPeeringInput)(nil)).Elem(), &VpcPeering{})
	pulumi.RegisterInputType(reflect.TypeOf((*VpcPeeringArrayInput)(nil)).Elem(), VpcPeeringArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*VpcPeeringMapInput)(nil)).Elem(), VpcPeeringMap{})
	pulumi.RegisterOutputType(VpcPeeringOutput{})
	pulumi.RegisterOutputType(VpcPeeringArrayOutput{})
	pulumi.RegisterOutputType(VpcPeeringMapOutput{})
}
