// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cloudamqp

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to retrieve information about the node(s) created by CloudAMQP instance.
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
//			_, err := cloudamqp.GetNodes(ctx, &cloudamqp.GetNodesArgs{
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
// * `hostname`              - External hostname assigned to the node.
// * `name`                  - Name of the node.
// * `running`               - Is the node running?
// * `rabbitmqVersion`      - Currently configured Rabbit MQ version on the node.
// * `erlangVersion`        - Currently used Erlanbg version on the node.
// * `hipe`                  - Enable or disable High-performance Erlang.
// * `configured`            - Is the node configured?
// * `diskSize`             - Subscription plan disk size
// * `additionalDiskSize`  - Additional added disk size
//
// ***Note:*** *Total disk size = diskSize + additional_disk_size*
//
// ## Dependency
//
// This data source depends on CloudAMQP instance identifier, `cloudamqp_instance.instance.id`.
func GetNodes(ctx *pulumi.Context, args *GetNodesArgs, opts ...pulumi.InvokeOption) (*GetNodesResult, error) {
	var rv GetNodesResult
	err := ctx.Invoke("cloudamqp:index/getNodes:getNodes", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getNodes.
type GetNodesArgs struct {
	// The CloudAMQP instance identifier.
	InstanceId int `pulumi:"instanceId"`
}

// A collection of values returned by getNodes.
type GetNodesResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id         string         `pulumi:"id"`
	InstanceId int            `pulumi:"instanceId"`
	Nodes      []GetNodesNode `pulumi:"nodes"`
}

func GetNodesOutput(ctx *pulumi.Context, args GetNodesOutputArgs, opts ...pulumi.InvokeOption) GetNodesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetNodesResult, error) {
			args := v.(GetNodesArgs)
			r, err := GetNodes(ctx, &args, opts...)
			var s GetNodesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetNodesResultOutput)
}

// A collection of arguments for invoking getNodes.
type GetNodesOutputArgs struct {
	// The CloudAMQP instance identifier.
	InstanceId pulumi.IntInput `pulumi:"instanceId"`
}

func (GetNodesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetNodesArgs)(nil)).Elem()
}

// A collection of values returned by getNodes.
type GetNodesResultOutput struct{ *pulumi.OutputState }

func (GetNodesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetNodesResult)(nil)).Elem()
}

func (o GetNodesResultOutput) ToGetNodesResultOutput() GetNodesResultOutput {
	return o
}

func (o GetNodesResultOutput) ToGetNodesResultOutputWithContext(ctx context.Context) GetNodesResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o GetNodesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetNodesResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetNodesResultOutput) InstanceId() pulumi.IntOutput {
	return o.ApplyT(func(v GetNodesResult) int { return v.InstanceId }).(pulumi.IntOutput)
}

func (o GetNodesResultOutput) Nodes() GetNodesNodeArrayOutput {
	return o.ApplyT(func(v GetNodesResult) []GetNodesNode { return v.Nodes }).(GetNodesNodeArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetNodesResultOutput{})
}
