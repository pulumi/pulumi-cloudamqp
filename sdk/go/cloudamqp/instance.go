// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package cloudamqp

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// This resource allows you to create and manage a CloudAMQP instance running Rabbit MQ and deploy to multiple cloud platforms provider and over multiple regions, see Instance regions for more information.
//
// Once the instance is created it will be assigned a unique identifier. All other resource and data sources created for this instance needs to reference the instance identifier.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-cloudamqp/sdk/v2/go/cloudamqp/"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := cloudamqp.NewInstance(ctx, "lemurInstance", &cloudamqp.InstanceArgs{
// 			Plan:   pulumi.String("lemur"),
// 			Region: pulumi.String("amazon-web-services::us-west-1"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = cloudamqp.NewInstance(ctx, "instance", &cloudamqp.InstanceArgs{
// 			NoDefaultAlarms: pulumi.Bool(true),
// 			Nodes:           pulumi.Int(1),
// 			Plan:            pulumi.String("bunny"),
// 			Region:          pulumi.String("amazon-web-services::us-west-1"),
// 			RmqVersion:      pulumi.String("3.8.3"),
// 			Tags: pulumi.StringArray{
// 				pulumi.String("terraform"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## Import
//
// `cloudamqp_instance`can be imported using CloudAMQP internal identifier. To retrieve the identifier for an instance, use [CloudAMQP customer API](https://docs.cloudamqp.com/#list-instances).
//
// ```sh
//  $ pulumi import cloudamqp:index/instance:Instance instance <instance_id>`
// ```
type Instance struct {
	pulumi.CustomResourceState

	// (Computed) API key needed to communicate to CloudAMQP's second API. The second API is used to manage alarms, integration and more, full description [CloudAMQP API](https://docs.cloudamqp.com/cloudamqp_api.html).
	Apikey pulumi.StringOutput `pulumi:"apikey"`
	// Is the instance hosted on a dedicated server
	Dedicated pulumi.BoolOutput `pulumi:"dedicated"`
	// (Computed) The host name for the CloudAMQP instance.
	Host pulumi.StringOutput `pulumi:"host"`
	// Name of the CloudAMQP instance.
	Name pulumi.StringOutput `pulumi:"name"`
	// Set to true to discard creating default alarms when the instance is created. Can be left out, will then use default value = false.
	NoDefaultAlarms pulumi.BoolOutput `pulumi:"noDefaultAlarms"`
	// Number of nodes, 1 to 3, in the CloudAMQP instance, default set to 1. The plan chosen must support the number of nodes.
	Nodes pulumi.IntPtrOutput `pulumi:"nodes"`
	// The subscription plan. See available plans
	Plan pulumi.StringOutput `pulumi:"plan"`
	// Flag describing if the resource is ready
	Ready pulumi.BoolOutput `pulumi:"ready"`
	// The region to host the instance in. See Instance regions
	Region pulumi.StringOutput `pulumi:"region"`
	// The Rabbit MQ version. Can be left out, will then be set to default value used by CloudAMQP API. **Note: There is not yet any support in the provider to change the RMQ version. Once it's set in the initial creation, it will remain.**
	RmqVersion pulumi.StringOutput `pulumi:"rmqVersion"`
	// One or more tags for the CloudAMQP instance, makes it possible to categories multiple instances in console view. Default there is no tags assigned.
	Tags pulumi.StringArrayOutput `pulumi:"tags"`
	// (Computed) AMQP server endpoint. `amqps://{username}:{password}@{hostname}/{vhost}`
	Url pulumi.StringOutput `pulumi:"url"`
	// (Computed) The virtual host used by Rabbit MQ.
	Vhost pulumi.StringOutput `pulumi:"vhost"`
	// Creates a dedicated VPC subnet, shouldn't overlap with other VPC subnet, default subnet used 10.56.72.0/24. **NOTE: extra fee will be charged when using VPC, see [CloudAMQP](https://cloudamqp.com) for more information.**
	VpcSubnet pulumi.StringPtrOutput `pulumi:"vpcSubnet"`
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
	// (Computed) API key needed to communicate to CloudAMQP's second API. The second API is used to manage alarms, integration and more, full description [CloudAMQP API](https://docs.cloudamqp.com/cloudamqp_api.html).
	Apikey *string `pulumi:"apikey"`
	// Is the instance hosted on a dedicated server
	Dedicated *bool `pulumi:"dedicated"`
	// (Computed) The host name for the CloudAMQP instance.
	Host *string `pulumi:"host"`
	// Name of the CloudAMQP instance.
	Name *string `pulumi:"name"`
	// Set to true to discard creating default alarms when the instance is created. Can be left out, will then use default value = false.
	NoDefaultAlarms *bool `pulumi:"noDefaultAlarms"`
	// Number of nodes, 1 to 3, in the CloudAMQP instance, default set to 1. The plan chosen must support the number of nodes.
	Nodes *int `pulumi:"nodes"`
	// The subscription plan. See available plans
	Plan *string `pulumi:"plan"`
	// Flag describing if the resource is ready
	Ready *bool `pulumi:"ready"`
	// The region to host the instance in. See Instance regions
	Region *string `pulumi:"region"`
	// The Rabbit MQ version. Can be left out, will then be set to default value used by CloudAMQP API. **Note: There is not yet any support in the provider to change the RMQ version. Once it's set in the initial creation, it will remain.**
	RmqVersion *string `pulumi:"rmqVersion"`
	// One or more tags for the CloudAMQP instance, makes it possible to categories multiple instances in console view. Default there is no tags assigned.
	Tags []string `pulumi:"tags"`
	// (Computed) AMQP server endpoint. `amqps://{username}:{password}@{hostname}/{vhost}`
	Url *string `pulumi:"url"`
	// (Computed) The virtual host used by Rabbit MQ.
	Vhost *string `pulumi:"vhost"`
	// Creates a dedicated VPC subnet, shouldn't overlap with other VPC subnet, default subnet used 10.56.72.0/24. **NOTE: extra fee will be charged when using VPC, see [CloudAMQP](https://cloudamqp.com) for more information.**
	VpcSubnet *string `pulumi:"vpcSubnet"`
}

type InstanceState struct {
	// (Computed) API key needed to communicate to CloudAMQP's second API. The second API is used to manage alarms, integration and more, full description [CloudAMQP API](https://docs.cloudamqp.com/cloudamqp_api.html).
	Apikey pulumi.StringPtrInput
	// Is the instance hosted on a dedicated server
	Dedicated pulumi.BoolPtrInput
	// (Computed) The host name for the CloudAMQP instance.
	Host pulumi.StringPtrInput
	// Name of the CloudAMQP instance.
	Name pulumi.StringPtrInput
	// Set to true to discard creating default alarms when the instance is created. Can be left out, will then use default value = false.
	NoDefaultAlarms pulumi.BoolPtrInput
	// Number of nodes, 1 to 3, in the CloudAMQP instance, default set to 1. The plan chosen must support the number of nodes.
	Nodes pulumi.IntPtrInput
	// The subscription plan. See available plans
	Plan pulumi.StringPtrInput
	// Flag describing if the resource is ready
	Ready pulumi.BoolPtrInput
	// The region to host the instance in. See Instance regions
	Region pulumi.StringPtrInput
	// The Rabbit MQ version. Can be left out, will then be set to default value used by CloudAMQP API. **Note: There is not yet any support in the provider to change the RMQ version. Once it's set in the initial creation, it will remain.**
	RmqVersion pulumi.StringPtrInput
	// One or more tags for the CloudAMQP instance, makes it possible to categories multiple instances in console view. Default there is no tags assigned.
	Tags pulumi.StringArrayInput
	// (Computed) AMQP server endpoint. `amqps://{username}:{password}@{hostname}/{vhost}`
	Url pulumi.StringPtrInput
	// (Computed) The virtual host used by Rabbit MQ.
	Vhost pulumi.StringPtrInput
	// Creates a dedicated VPC subnet, shouldn't overlap with other VPC subnet, default subnet used 10.56.72.0/24. **NOTE: extra fee will be charged when using VPC, see [CloudAMQP](https://cloudamqp.com) for more information.**
	VpcSubnet pulumi.StringPtrInput
}

func (InstanceState) ElementType() reflect.Type {
	return reflect.TypeOf((*instanceState)(nil)).Elem()
}

type instanceArgs struct {
	// Name of the CloudAMQP instance.
	Name *string `pulumi:"name"`
	// Set to true to discard creating default alarms when the instance is created. Can be left out, will then use default value = false.
	NoDefaultAlarms *bool `pulumi:"noDefaultAlarms"`
	// Number of nodes, 1 to 3, in the CloudAMQP instance, default set to 1. The plan chosen must support the number of nodes.
	Nodes *int `pulumi:"nodes"`
	// The subscription plan. See available plans
	Plan string `pulumi:"plan"`
	// The region to host the instance in. See Instance regions
	Region string `pulumi:"region"`
	// The Rabbit MQ version. Can be left out, will then be set to default value used by CloudAMQP API. **Note: There is not yet any support in the provider to change the RMQ version. Once it's set in the initial creation, it will remain.**
	RmqVersion *string `pulumi:"rmqVersion"`
	// One or more tags for the CloudAMQP instance, makes it possible to categories multiple instances in console view. Default there is no tags assigned.
	Tags []string `pulumi:"tags"`
	// Creates a dedicated VPC subnet, shouldn't overlap with other VPC subnet, default subnet used 10.56.72.0/24. **NOTE: extra fee will be charged when using VPC, see [CloudAMQP](https://cloudamqp.com) for more information.**
	VpcSubnet *string `pulumi:"vpcSubnet"`
}

// The set of arguments for constructing a Instance resource.
type InstanceArgs struct {
	// Name of the CloudAMQP instance.
	Name pulumi.StringPtrInput
	// Set to true to discard creating default alarms when the instance is created. Can be left out, will then use default value = false.
	NoDefaultAlarms pulumi.BoolPtrInput
	// Number of nodes, 1 to 3, in the CloudAMQP instance, default set to 1. The plan chosen must support the number of nodes.
	Nodes pulumi.IntPtrInput
	// The subscription plan. See available plans
	Plan pulumi.StringInput
	// The region to host the instance in. See Instance regions
	Region pulumi.StringInput
	// The Rabbit MQ version. Can be left out, will then be set to default value used by CloudAMQP API. **Note: There is not yet any support in the provider to change the RMQ version. Once it's set in the initial creation, it will remain.**
	RmqVersion pulumi.StringPtrInput
	// One or more tags for the CloudAMQP instance, makes it possible to categories multiple instances in console view. Default there is no tags assigned.
	Tags pulumi.StringArrayInput
	// Creates a dedicated VPC subnet, shouldn't overlap with other VPC subnet, default subnet used 10.56.72.0/24. **NOTE: extra fee will be charged when using VPC, see [CloudAMQP](https://cloudamqp.com) for more information.**
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
	return reflect.TypeOf((*Instance)(nil))
}

func (i *Instance) ToInstanceOutput() InstanceOutput {
	return i.ToInstanceOutputWithContext(context.Background())
}

func (i *Instance) ToInstanceOutputWithContext(ctx context.Context) InstanceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceOutput)
}

type InstanceOutput struct {
	*pulumi.OutputState
}

func (InstanceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Instance)(nil))
}

func (o InstanceOutput) ToInstanceOutput() InstanceOutput {
	return o
}

func (o InstanceOutput) ToInstanceOutputWithContext(ctx context.Context) InstanceOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(InstanceOutput{})
}
