// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cloudamqp

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-cloudamqp/sdk/v3/go/cloudamqp/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to retrieve information about the credentials of the configured user in Rabbit MQ. Information is extracted from `cloudamqp_instance.instance.url`.
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
//			_, err := cloudamqp.GetCredentials(ctx, &cloudamqp.GetCredentialsArgs{
//				InstanceId: instance.Id,
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
// ## Attributes reference
//
// All attributes reference are computed.
//
// * `id`          - The identifier for this data source.
// * `username`    - (Sensitive) The username for the configured user in Rabbit MQ.
// * `password`    - (Sensitive) The password used by the `username`.
//
// ## Dependency
//
// This data source depends on CloudAMQP instance identifier, `cloudamqp_instance.instance.id`.
func GetCredentials(ctx *pulumi.Context, args *GetCredentialsArgs, opts ...pulumi.InvokeOption) (*GetCredentialsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetCredentialsResult
	err := ctx.Invoke("cloudamqp:index/getCredentials:getCredentials", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getCredentials.
type GetCredentialsArgs struct {
	// The CloudAMQP instance identifier.
	InstanceId int `pulumi:"instanceId"`
}

// A collection of values returned by getCredentials.
type GetCredentialsResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id         string `pulumi:"id"`
	InstanceId int    `pulumi:"instanceId"`
	Password   string `pulumi:"password"`
	Username   string `pulumi:"username"`
}

func GetCredentialsOutput(ctx *pulumi.Context, args GetCredentialsOutputArgs, opts ...pulumi.InvokeOption) GetCredentialsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetCredentialsResultOutput, error) {
			args := v.(GetCredentialsArgs)
			opts = internal.PkgInvokeDefaultOpts(opts)
			var rv GetCredentialsResult
			secret, err := ctx.InvokePackageRaw("cloudamqp:index/getCredentials:getCredentials", args, &rv, "", opts...)
			if err != nil {
				return GetCredentialsResultOutput{}, err
			}

			output := pulumi.ToOutput(rv).(GetCredentialsResultOutput)
			if secret {
				return pulumi.ToSecret(output).(GetCredentialsResultOutput), nil
			}
			return output, nil
		}).(GetCredentialsResultOutput)
}

// A collection of arguments for invoking getCredentials.
type GetCredentialsOutputArgs struct {
	// The CloudAMQP instance identifier.
	InstanceId pulumi.IntInput `pulumi:"instanceId"`
}

func (GetCredentialsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetCredentialsArgs)(nil)).Elem()
}

// A collection of values returned by getCredentials.
type GetCredentialsResultOutput struct{ *pulumi.OutputState }

func (GetCredentialsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetCredentialsResult)(nil)).Elem()
}

func (o GetCredentialsResultOutput) ToGetCredentialsResultOutput() GetCredentialsResultOutput {
	return o
}

func (o GetCredentialsResultOutput) ToGetCredentialsResultOutputWithContext(ctx context.Context) GetCredentialsResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o GetCredentialsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetCredentialsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetCredentialsResultOutput) InstanceId() pulumi.IntOutput {
	return o.ApplyT(func(v GetCredentialsResult) int { return v.InstanceId }).(pulumi.IntOutput)
}

func (o GetCredentialsResultOutput) Password() pulumi.StringOutput {
	return o.ApplyT(func(v GetCredentialsResult) string { return v.Password }).(pulumi.StringOutput)
}

func (o GetCredentialsResultOutput) Username() pulumi.StringOutput {
	return o.ApplyT(func(v GetCredentialsResult) string { return v.Username }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetCredentialsResultOutput{})
}
