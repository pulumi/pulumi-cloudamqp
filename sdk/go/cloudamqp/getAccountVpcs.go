// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cloudamqp

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-cloudamqp/sdk/v3/go/cloudamqp/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to retrieve basic information about all standalone VPCs available for an account. Uses the included apikey in provider configuration to determine which account to read from.
func GetAccountVpcs(ctx *pulumi.Context, opts ...pulumi.InvokeOption) (*GetAccountVpcsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetAccountVpcsResult
	err := ctx.Invoke("cloudamqp:index/getAccountVpcs:getAccountVpcs", nil, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of values returned by getAccountVpcs.
type GetAccountVpcsResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id   string              `pulumi:"id"`
	Vpcs []GetAccountVpcsVpc `pulumi:"vpcs"`
}

func GetAccountVpcsOutput(ctx *pulumi.Context, opts ...pulumi.InvokeOption) GetAccountVpcsResultOutput {
	return pulumi.ToOutput(0).ApplyT(func(int) (GetAccountVpcsResultOutput, error) {
		options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
		return ctx.InvokeOutput("cloudamqp:index/getAccountVpcs:getAccountVpcs", nil, GetAccountVpcsResultOutput{}, options).(GetAccountVpcsResultOutput), nil
	}).(GetAccountVpcsResultOutput)
}

// A collection of values returned by getAccountVpcs.
type GetAccountVpcsResultOutput struct{ *pulumi.OutputState }

func (GetAccountVpcsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetAccountVpcsResult)(nil)).Elem()
}

func (o GetAccountVpcsResultOutput) ToGetAccountVpcsResultOutput() GetAccountVpcsResultOutput {
	return o
}

func (o GetAccountVpcsResultOutput) ToGetAccountVpcsResultOutputWithContext(ctx context.Context) GetAccountVpcsResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o GetAccountVpcsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetAccountVpcsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetAccountVpcsResultOutput) Vpcs() GetAccountVpcsVpcArrayOutput {
	return o.ApplyT(func(v GetAccountVpcsResult) []GetAccountVpcsVpc { return v.Vpcs }).(GetAccountVpcsVpcArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetAccountVpcsResultOutput{})
}
