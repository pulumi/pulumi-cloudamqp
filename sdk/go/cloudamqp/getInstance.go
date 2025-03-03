// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cloudamqp

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-cloudamqp/sdk/v3/go/cloudamqp/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to retrieve information about an already created CloudAMQP instance. In order to retrieve the correct information, the CoudAMQP instance identifier is needed.
func LookupInstance(ctx *pulumi.Context, args *LookupInstanceArgs, opts ...pulumi.InvokeOption) (*LookupInstanceResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupInstanceResult
	err := ctx.Invoke("cloudamqp:index/getInstance:getInstance", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getInstance.
type LookupInstanceArgs struct {
	// The CloudAMQP instance identifier.
	InstanceId int `pulumi:"instanceId"`
}

// A collection of values returned by getInstance.
type LookupInstanceResult struct {
	Apikey       string `pulumi:"apikey"`
	Backend      string `pulumi:"backend"`
	Dedicated    bool   `pulumi:"dedicated"`
	Host         string `pulumi:"host"`
	HostInternal string `pulumi:"hostInternal"`
	// The provider-assigned unique ID for this managed resource.
	Id              string   `pulumi:"id"`
	InstanceId      int      `pulumi:"instanceId"`
	Name            string   `pulumi:"name"`
	NoDefaultAlarms bool     `pulumi:"noDefaultAlarms"`
	Nodes           int      `pulumi:"nodes"`
	Plan            string   `pulumi:"plan"`
	Ready           bool     `pulumi:"ready"`
	Region          string   `pulumi:"region"`
	RmqVersion      string   `pulumi:"rmqVersion"`
	Tags            []string `pulumi:"tags"`
	Url             string   `pulumi:"url"`
	Vhost           string   `pulumi:"vhost"`
	VpcId           int      `pulumi:"vpcId"`
	VpcSubnet       string   `pulumi:"vpcSubnet"`
}

func LookupInstanceOutput(ctx *pulumi.Context, args LookupInstanceOutputArgs, opts ...pulumi.InvokeOption) LookupInstanceResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupInstanceResultOutput, error) {
			args := v.(LookupInstanceArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("cloudamqp:index/getInstance:getInstance", args, LookupInstanceResultOutput{}, options).(LookupInstanceResultOutput), nil
		}).(LookupInstanceResultOutput)
}

// A collection of arguments for invoking getInstance.
type LookupInstanceOutputArgs struct {
	// The CloudAMQP instance identifier.
	InstanceId pulumi.IntInput `pulumi:"instanceId"`
}

func (LookupInstanceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupInstanceArgs)(nil)).Elem()
}

// A collection of values returned by getInstance.
type LookupInstanceResultOutput struct{ *pulumi.OutputState }

func (LookupInstanceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupInstanceResult)(nil)).Elem()
}

func (o LookupInstanceResultOutput) ToLookupInstanceResultOutput() LookupInstanceResultOutput {
	return o
}

func (o LookupInstanceResultOutput) ToLookupInstanceResultOutputWithContext(ctx context.Context) LookupInstanceResultOutput {
	return o
}

func (o LookupInstanceResultOutput) Apikey() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceResult) string { return v.Apikey }).(pulumi.StringOutput)
}

func (o LookupInstanceResultOutput) Backend() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceResult) string { return v.Backend }).(pulumi.StringOutput)
}

func (o LookupInstanceResultOutput) Dedicated() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupInstanceResult) bool { return v.Dedicated }).(pulumi.BoolOutput)
}

func (o LookupInstanceResultOutput) Host() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceResult) string { return v.Host }).(pulumi.StringOutput)
}

func (o LookupInstanceResultOutput) HostInternal() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceResult) string { return v.HostInternal }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupInstanceResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupInstanceResultOutput) InstanceId() pulumi.IntOutput {
	return o.ApplyT(func(v LookupInstanceResult) int { return v.InstanceId }).(pulumi.IntOutput)
}

func (o LookupInstanceResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupInstanceResultOutput) NoDefaultAlarms() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupInstanceResult) bool { return v.NoDefaultAlarms }).(pulumi.BoolOutput)
}

func (o LookupInstanceResultOutput) Nodes() pulumi.IntOutput {
	return o.ApplyT(func(v LookupInstanceResult) int { return v.Nodes }).(pulumi.IntOutput)
}

func (o LookupInstanceResultOutput) Plan() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceResult) string { return v.Plan }).(pulumi.StringOutput)
}

func (o LookupInstanceResultOutput) Ready() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupInstanceResult) bool { return v.Ready }).(pulumi.BoolOutput)
}

func (o LookupInstanceResultOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceResult) string { return v.Region }).(pulumi.StringOutput)
}

func (o LookupInstanceResultOutput) RmqVersion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceResult) string { return v.RmqVersion }).(pulumi.StringOutput)
}

func (o LookupInstanceResultOutput) Tags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupInstanceResult) []string { return v.Tags }).(pulumi.StringArrayOutput)
}

func (o LookupInstanceResultOutput) Url() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceResult) string { return v.Url }).(pulumi.StringOutput)
}

func (o LookupInstanceResultOutput) Vhost() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceResult) string { return v.Vhost }).(pulumi.StringOutput)
}

func (o LookupInstanceResultOutput) VpcId() pulumi.IntOutput {
	return o.ApplyT(func(v LookupInstanceResult) int { return v.VpcId }).(pulumi.IntOutput)
}

func (o LookupInstanceResultOutput) VpcSubnet() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceResult) string { return v.VpcSubnet }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupInstanceResultOutput{})
}
