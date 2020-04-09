// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package cloudamqp

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

type VpcPeering struct {
	pulumi.CustomResourceState

	// Instance identifier
	InstanceId pulumi.IntOutput `pulumi:"instanceId"`
	// VPC peering identifier
	PeeringId pulumi.StringOutput `pulumi:"peeringId"`
	// VPC peering status
	Status pulumi.StringOutput `pulumi:"status"`
}

// NewVpcPeering registers a new resource with the given unique name, arguments, and options.
func NewVpcPeering(ctx *pulumi.Context,
	name string, args *VpcPeeringArgs, opts ...pulumi.ResourceOption) (*VpcPeering, error) {
	if args == nil || args.InstanceId == nil {
		return nil, errors.New("missing required argument 'InstanceId'")
	}
	if args == nil || args.PeeringId == nil {
		return nil, errors.New("missing required argument 'PeeringId'")
	}
	if args == nil {
		args = &VpcPeeringArgs{}
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
	// Instance identifier
	InstanceId *int `pulumi:"instanceId"`
	// VPC peering identifier
	PeeringId *string `pulumi:"peeringId"`
	// VPC peering status
	Status *string `pulumi:"status"`
}

type VpcPeeringState struct {
	// Instance identifier
	InstanceId pulumi.IntPtrInput
	// VPC peering identifier
	PeeringId pulumi.StringPtrInput
	// VPC peering status
	Status pulumi.StringPtrInput
}

func (VpcPeeringState) ElementType() reflect.Type {
	return reflect.TypeOf((*vpcPeeringState)(nil)).Elem()
}

type vpcPeeringArgs struct {
	// Instance identifier
	InstanceId int `pulumi:"instanceId"`
	// VPC peering identifier
	PeeringId string `pulumi:"peeringId"`
}

// The set of arguments for constructing a VpcPeering resource.
type VpcPeeringArgs struct {
	// Instance identifier
	InstanceId pulumi.IntInput
	// VPC peering identifier
	PeeringId pulumi.StringInput
}

func (VpcPeeringArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*vpcPeeringArgs)(nil)).Elem()
}
