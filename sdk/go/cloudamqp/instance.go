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
// `cloudamqp_instance`can be imported using CloudAMQP internal identifier.
//
// ```sh
//
//	$ pulumi import cloudamqp:index/instance:Instance instance <id>`
//
// ```
//
//	To retrieve the identifier for a VPC, either use [CloudAMQP customer API](https://docs.cloudamqp.com/#list-instances). Or use the data source [`cloudamqp_account`](https://registry.terraform.io/providers/cloudamqp/cloudamqp/latest/docs/data-sources/account) to list all available instances for an account.
type Instance struct {
	pulumi.CustomResourceState

	// API key needed to communicate to CloudAMQP's second API. The second API is used to manage alarms, integration and more, full description [CloudAMQP API](https://docs.cloudamqp.com/cloudamqp_api.html).
	Apikey pulumi.StringOutput `pulumi:"apikey"`
	// Is the instance hosted on a dedicated server
	Dedicated pulumi.BoolOutput `pulumi:"dedicated"`
	// The external hostname for the CloudAMQP instance.
	Host pulumi.StringOutput `pulumi:"host"`
	// The internal hostname for the CloudAMQP instance.
	HostInternal pulumi.StringOutput `pulumi:"hostInternal"`
	// Keep associated VPC when deleting instance, default set to false.
	KeepAssociatedVpc pulumi.BoolPtrOutput `pulumi:"keepAssociatedVpc"`
	// Name of the CloudAMQP instance.
	Name pulumi.StringOutput `pulumi:"name"`
	// Set to true to discard creating default alarms when the instance is created. Can be left out, will then use default value = false.
	NoDefaultAlarms pulumi.BoolOutput `pulumi:"noDefaultAlarms"`
	// Number of nodes, 1, 3 or 5 depending on plan used.
	Nodes pulumi.IntOutput `pulumi:"nodes"`
	// The subscription plan. See available plans
	Plan pulumi.StringOutput `pulumi:"plan"`
	// Flag describing if the resource is ready
	Ready pulumi.BoolOutput `pulumi:"ready"`
	// The region to host the instance in. See Instance regions
	Region pulumi.StringOutput `pulumi:"region"`
	// The Rabbit MQ version. Can be left out, will then be set to default value used by CloudAMQP API.
	RmqVersion pulumi.StringOutput `pulumi:"rmqVersion"`
	// One or more tags for the CloudAMQP instance, makes it possible to categories multiple instances in console view. Default there is no tags assigned.
	Tags pulumi.StringArrayOutput `pulumi:"tags"`
	// The AMQP URL (uses the internal hostname if the instance was created with VPC). Has the format: `amqps://{username}:{password}@{hostname}/{vhost}`
	Url pulumi.StringOutput `pulumi:"url"`
	// The virtual host used by Rabbit MQ.
	Vhost pulumi.StringOutput `pulumi:"vhost"`
	// The VPC ID. Use this to create your instance in an existing VPC. See available example.
	VpcId pulumi.IntOutput `pulumi:"vpcId"`
	// Creates a dedicated VPC subnet, shouldn't overlap with other VPC subnet, default subnet used 10.56.72.0/24.
	VpcSubnet pulumi.StringOutput `pulumi:"vpcSubnet"`
}

// NewInstance registers a new resource with the given unique name, arguments, and options.
func NewInstance(ctx *pulumi.Context,
	name string, args *InstanceArgs, opts ...pulumi.ResourceOption) (*Instance, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Plan == nil {
		return nil, errors.New("invalid value for required argument 'Plan'")
	}
	if args.Region == nil {
		return nil, errors.New("invalid value for required argument 'Region'")
	}
	var resource Instance
	err := ctx.RegisterResource("cloudamqp:index/instance:Instance", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetInstance gets an existing Instance resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetInstance(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *InstanceState, opts ...pulumi.ResourceOption) (*Instance, error) {
	var resource Instance
	err := ctx.ReadResource("cloudamqp:index/instance:Instance", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Instance resources.
type instanceState struct {
	// API key needed to communicate to CloudAMQP's second API. The second API is used to manage alarms, integration and more, full description [CloudAMQP API](https://docs.cloudamqp.com/cloudamqp_api.html).
	Apikey *string `pulumi:"apikey"`
	// Is the instance hosted on a dedicated server
	Dedicated *bool `pulumi:"dedicated"`
	// The external hostname for the CloudAMQP instance.
	Host *string `pulumi:"host"`
	// The internal hostname for the CloudAMQP instance.
	HostInternal *string `pulumi:"hostInternal"`
	// Keep associated VPC when deleting instance, default set to false.
	KeepAssociatedVpc *bool `pulumi:"keepAssociatedVpc"`
	// Name of the CloudAMQP instance.
	Name *string `pulumi:"name"`
	// Set to true to discard creating default alarms when the instance is created. Can be left out, will then use default value = false.
	NoDefaultAlarms *bool `pulumi:"noDefaultAlarms"`
	// Number of nodes, 1, 3 or 5 depending on plan used.
	Nodes *int `pulumi:"nodes"`
	// The subscription plan. See available plans
	Plan *string `pulumi:"plan"`
	// Flag describing if the resource is ready
	Ready *bool `pulumi:"ready"`
	// The region to host the instance in. See Instance regions
	Region *string `pulumi:"region"`
	// The Rabbit MQ version. Can be left out, will then be set to default value used by CloudAMQP API.
	RmqVersion *string `pulumi:"rmqVersion"`
	// One or more tags for the CloudAMQP instance, makes it possible to categories multiple instances in console view. Default there is no tags assigned.
	Tags []string `pulumi:"tags"`
	// The AMQP URL (uses the internal hostname if the instance was created with VPC). Has the format: `amqps://{username}:{password}@{hostname}/{vhost}`
	Url *string `pulumi:"url"`
	// The virtual host used by Rabbit MQ.
	Vhost *string `pulumi:"vhost"`
	// The VPC ID. Use this to create your instance in an existing VPC. See available example.
	VpcId *int `pulumi:"vpcId"`
	// Creates a dedicated VPC subnet, shouldn't overlap with other VPC subnet, default subnet used 10.56.72.0/24.
	VpcSubnet *string `pulumi:"vpcSubnet"`
}

type InstanceState struct {
	// API key needed to communicate to CloudAMQP's second API. The second API is used to manage alarms, integration and more, full description [CloudAMQP API](https://docs.cloudamqp.com/cloudamqp_api.html).
	Apikey pulumi.StringPtrInput
	// Is the instance hosted on a dedicated server
	Dedicated pulumi.BoolPtrInput
	// The external hostname for the CloudAMQP instance.
	Host pulumi.StringPtrInput
	// The internal hostname for the CloudAMQP instance.
	HostInternal pulumi.StringPtrInput
	// Keep associated VPC when deleting instance, default set to false.
	KeepAssociatedVpc pulumi.BoolPtrInput
	// Name of the CloudAMQP instance.
	Name pulumi.StringPtrInput
	// Set to true to discard creating default alarms when the instance is created. Can be left out, will then use default value = false.
	NoDefaultAlarms pulumi.BoolPtrInput
	// Number of nodes, 1, 3 or 5 depending on plan used.
	Nodes pulumi.IntPtrInput
	// The subscription plan. See available plans
	Plan pulumi.StringPtrInput
	// Flag describing if the resource is ready
	Ready pulumi.BoolPtrInput
	// The region to host the instance in. See Instance regions
	Region pulumi.StringPtrInput
	// The Rabbit MQ version. Can be left out, will then be set to default value used by CloudAMQP API.
	RmqVersion pulumi.StringPtrInput
	// One or more tags for the CloudAMQP instance, makes it possible to categories multiple instances in console view. Default there is no tags assigned.
	Tags pulumi.StringArrayInput
	// The AMQP URL (uses the internal hostname if the instance was created with VPC). Has the format: `amqps://{username}:{password}@{hostname}/{vhost}`
	Url pulumi.StringPtrInput
	// The virtual host used by Rabbit MQ.
	Vhost pulumi.StringPtrInput
	// The VPC ID. Use this to create your instance in an existing VPC. See available example.
	VpcId pulumi.IntPtrInput
	// Creates a dedicated VPC subnet, shouldn't overlap with other VPC subnet, default subnet used 10.56.72.0/24.
	VpcSubnet pulumi.StringPtrInput
}

func (InstanceState) ElementType() reflect.Type {
	return reflect.TypeOf((*instanceState)(nil)).Elem()
}

type instanceArgs struct {
	// Keep associated VPC when deleting instance, default set to false.
	KeepAssociatedVpc *bool `pulumi:"keepAssociatedVpc"`
	// Name of the CloudAMQP instance.
	Name *string `pulumi:"name"`
	// Set to true to discard creating default alarms when the instance is created. Can be left out, will then use default value = false.
	NoDefaultAlarms *bool `pulumi:"noDefaultAlarms"`
	// Number of nodes, 1, 3 or 5 depending on plan used.
	Nodes *int `pulumi:"nodes"`
	// The subscription plan. See available plans
	Plan string `pulumi:"plan"`
	// The region to host the instance in. See Instance regions
	Region string `pulumi:"region"`
	// The Rabbit MQ version. Can be left out, will then be set to default value used by CloudAMQP API.
	RmqVersion *string `pulumi:"rmqVersion"`
	// One or more tags for the CloudAMQP instance, makes it possible to categories multiple instances in console view. Default there is no tags assigned.
	Tags []string `pulumi:"tags"`
	// The VPC ID. Use this to create your instance in an existing VPC. See available example.
	VpcId *int `pulumi:"vpcId"`
	// Creates a dedicated VPC subnet, shouldn't overlap with other VPC subnet, default subnet used 10.56.72.0/24.
	VpcSubnet *string `pulumi:"vpcSubnet"`
}

// The set of arguments for constructing a Instance resource.
type InstanceArgs struct {
	// Keep associated VPC when deleting instance, default set to false.
	KeepAssociatedVpc pulumi.BoolPtrInput
	// Name of the CloudAMQP instance.
	Name pulumi.StringPtrInput
	// Set to true to discard creating default alarms when the instance is created. Can be left out, will then use default value = false.
	NoDefaultAlarms pulumi.BoolPtrInput
	// Number of nodes, 1, 3 or 5 depending on plan used.
	Nodes pulumi.IntPtrInput
	// The subscription plan. See available plans
	Plan pulumi.StringInput
	// The region to host the instance in. See Instance regions
	Region pulumi.StringInput
	// The Rabbit MQ version. Can be left out, will then be set to default value used by CloudAMQP API.
	RmqVersion pulumi.StringPtrInput
	// One or more tags for the CloudAMQP instance, makes it possible to categories multiple instances in console view. Default there is no tags assigned.
	Tags pulumi.StringArrayInput
	// The VPC ID. Use this to create your instance in an existing VPC. See available example.
	VpcId pulumi.IntPtrInput
	// Creates a dedicated VPC subnet, shouldn't overlap with other VPC subnet, default subnet used 10.56.72.0/24.
	VpcSubnet pulumi.StringPtrInput
}

func (InstanceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*instanceArgs)(nil)).Elem()
}

type InstanceInput interface {
	pulumi.Input

	ToInstanceOutput() InstanceOutput
	ToInstanceOutputWithContext(ctx context.Context) InstanceOutput
}

func (*Instance) ElementType() reflect.Type {
	return reflect.TypeOf((**Instance)(nil)).Elem()
}

func (i *Instance) ToInstanceOutput() InstanceOutput {
	return i.ToInstanceOutputWithContext(context.Background())
}

func (i *Instance) ToInstanceOutputWithContext(ctx context.Context) InstanceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceOutput)
}

// InstanceArrayInput is an input type that accepts InstanceArray and InstanceArrayOutput values.
// You can construct a concrete instance of `InstanceArrayInput` via:
//
//	InstanceArray{ InstanceArgs{...} }
type InstanceArrayInput interface {
	pulumi.Input

	ToInstanceArrayOutput() InstanceArrayOutput
	ToInstanceArrayOutputWithContext(context.Context) InstanceArrayOutput
}

type InstanceArray []InstanceInput

func (InstanceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Instance)(nil)).Elem()
}

func (i InstanceArray) ToInstanceArrayOutput() InstanceArrayOutput {
	return i.ToInstanceArrayOutputWithContext(context.Background())
}

func (i InstanceArray) ToInstanceArrayOutputWithContext(ctx context.Context) InstanceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceArrayOutput)
}

// InstanceMapInput is an input type that accepts InstanceMap and InstanceMapOutput values.
// You can construct a concrete instance of `InstanceMapInput` via:
//
//	InstanceMap{ "key": InstanceArgs{...} }
type InstanceMapInput interface {
	pulumi.Input

	ToInstanceMapOutput() InstanceMapOutput
	ToInstanceMapOutputWithContext(context.Context) InstanceMapOutput
}

type InstanceMap map[string]InstanceInput

func (InstanceMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Instance)(nil)).Elem()
}

func (i InstanceMap) ToInstanceMapOutput() InstanceMapOutput {
	return i.ToInstanceMapOutputWithContext(context.Background())
}

func (i InstanceMap) ToInstanceMapOutputWithContext(ctx context.Context) InstanceMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceMapOutput)
}

type InstanceOutput struct{ *pulumi.OutputState }

func (InstanceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Instance)(nil)).Elem()
}

func (o InstanceOutput) ToInstanceOutput() InstanceOutput {
	return o
}

func (o InstanceOutput) ToInstanceOutputWithContext(ctx context.Context) InstanceOutput {
	return o
}

// API key needed to communicate to CloudAMQP's second API. The second API is used to manage alarms, integration and more, full description [CloudAMQP API](https://docs.cloudamqp.com/cloudamqp_api.html).
func (o InstanceOutput) Apikey() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.Apikey }).(pulumi.StringOutput)
}

// Is the instance hosted on a dedicated server
func (o InstanceOutput) Dedicated() pulumi.BoolOutput {
	return o.ApplyT(func(v *Instance) pulumi.BoolOutput { return v.Dedicated }).(pulumi.BoolOutput)
}

// The external hostname for the CloudAMQP instance.
func (o InstanceOutput) Host() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.Host }).(pulumi.StringOutput)
}

// The internal hostname for the CloudAMQP instance.
func (o InstanceOutput) HostInternal() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.HostInternal }).(pulumi.StringOutput)
}

// Keep associated VPC when deleting instance, default set to false.
func (o InstanceOutput) KeepAssociatedVpc() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Instance) pulumi.BoolPtrOutput { return v.KeepAssociatedVpc }).(pulumi.BoolPtrOutput)
}

// Name of the CloudAMQP instance.
func (o InstanceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Set to true to discard creating default alarms when the instance is created. Can be left out, will then use default value = false.
func (o InstanceOutput) NoDefaultAlarms() pulumi.BoolOutput {
	return o.ApplyT(func(v *Instance) pulumi.BoolOutput { return v.NoDefaultAlarms }).(pulumi.BoolOutput)
}

// Number of nodes, 1, 3 or 5 depending on plan used.
func (o InstanceOutput) Nodes() pulumi.IntOutput {
	return o.ApplyT(func(v *Instance) pulumi.IntOutput { return v.Nodes }).(pulumi.IntOutput)
}

// The subscription plan. See available plans
func (o InstanceOutput) Plan() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.Plan }).(pulumi.StringOutput)
}

// Flag describing if the resource is ready
func (o InstanceOutput) Ready() pulumi.BoolOutput {
	return o.ApplyT(func(v *Instance) pulumi.BoolOutput { return v.Ready }).(pulumi.BoolOutput)
}

// The region to host the instance in. See Instance regions
func (o InstanceOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

// The Rabbit MQ version. Can be left out, will then be set to default value used by CloudAMQP API.
func (o InstanceOutput) RmqVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.RmqVersion }).(pulumi.StringOutput)
}

// One or more tags for the CloudAMQP instance, makes it possible to categories multiple instances in console view. Default there is no tags assigned.
func (o InstanceOutput) Tags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringArrayOutput { return v.Tags }).(pulumi.StringArrayOutput)
}

// The AMQP URL (uses the internal hostname if the instance was created with VPC). Has the format: `amqps://{username}:{password}@{hostname}/{vhost}`
func (o InstanceOutput) Url() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.Url }).(pulumi.StringOutput)
}

// The virtual host used by Rabbit MQ.
func (o InstanceOutput) Vhost() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.Vhost }).(pulumi.StringOutput)
}

// The VPC ID. Use this to create your instance in an existing VPC. See available example.
func (o InstanceOutput) VpcId() pulumi.IntOutput {
	return o.ApplyT(func(v *Instance) pulumi.IntOutput { return v.VpcId }).(pulumi.IntOutput)
}

// Creates a dedicated VPC subnet, shouldn't overlap with other VPC subnet, default subnet used 10.56.72.0/24.
func (o InstanceOutput) VpcSubnet() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.VpcSubnet }).(pulumi.StringOutput)
}

type InstanceArrayOutput struct{ *pulumi.OutputState }

func (InstanceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Instance)(nil)).Elem()
}

func (o InstanceArrayOutput) ToInstanceArrayOutput() InstanceArrayOutput {
	return o
}

func (o InstanceArrayOutput) ToInstanceArrayOutputWithContext(ctx context.Context) InstanceArrayOutput {
	return o
}

func (o InstanceArrayOutput) Index(i pulumi.IntInput) InstanceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Instance {
		return vs[0].([]*Instance)[vs[1].(int)]
	}).(InstanceOutput)
}

type InstanceMapOutput struct{ *pulumi.OutputState }

func (InstanceMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Instance)(nil)).Elem()
}

func (o InstanceMapOutput) ToInstanceMapOutput() InstanceMapOutput {
	return o
}

func (o InstanceMapOutput) ToInstanceMapOutputWithContext(ctx context.Context) InstanceMapOutput {
	return o
}

func (o InstanceMapOutput) MapIndex(k pulumi.StringInput) InstanceOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Instance {
		return vs[0].(map[string]*Instance)[vs[1].(string)]
	}).(InstanceOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*InstanceInput)(nil)).Elem(), &Instance{})
	pulumi.RegisterInputType(reflect.TypeOf((*InstanceArrayInput)(nil)).Elem(), InstanceArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*InstanceMapInput)(nil)).Elem(), InstanceMap{})
	pulumi.RegisterOutputType(InstanceOutput{})
	pulumi.RegisterOutputType(InstanceArrayOutput{})
	pulumi.RegisterOutputType(InstanceMapOutput{})
}
