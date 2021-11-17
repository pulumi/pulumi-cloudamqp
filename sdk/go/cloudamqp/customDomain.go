// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package cloudamqp

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This resource allows you to configure and manage your custom domain for the CloudAMQP instance.
//
// Adding a custom domain to your instance will generate a TLS certificate from [Let's Encrypt], for the given hostname, and install it on all servers in your cluster. The certificate will be automatically renewed going forward.
//
// ⚠️ Please note that when creating, changing or deleting the custom domain, the listeners on your servers will be restarted in order to apply the changes. This will close your current connections.
//
// Your custom domain name needs to point to your CloudAMQP hostname, preferably using a CNAME DNS record.
//
// Only available for dedicated subscription plans.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-cloudamqp/sdk/v3/go/cloudamqp"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := cloudamqp.NewCustomDomain(ctx, "settings", &cloudamqp.CustomDomainArgs{
// 			InstanceId: pulumi.Any(cloudamqp_instance.Instance.Id),
// 			Hostname:   pulumi.String("myname.mydomain"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ## Depedency
//
// This resource depends on CloudAMQP instance identifier, `cloudamqp_instance.instance.id`.
//
// ## Import
//
// `cloudamqp_custom_domain` can be imported using CloudAMQP instance identifier.
//
// ```sh
//  $ pulumi import cloudamqp:index/customDomain:CustomDomain settings <instance_id>`
// ```
//
//  [Let's Encrypt]https://letsencrypt.org/
type CustomDomain struct {
	pulumi.CustomResourceState

	// Your custom domain name.
	Hostname pulumi.StringOutput `pulumi:"hostname"`
	// The CloudAMQP instance ID.
	InstanceId pulumi.IntOutput `pulumi:"instanceId"`
}

// NewCustomDomain registers a new resource with the given unique name, arguments, and options.
func NewCustomDomain(ctx *pulumi.Context,
	name string, args *CustomDomainArgs, opts ...pulumi.ResourceOption) (*CustomDomain, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Hostname == nil {
		return nil, errors.New("invalid value for required argument 'Hostname'")
	}
	if args.InstanceId == nil {
		return nil, errors.New("invalid value for required argument 'InstanceId'")
	}
	var resource CustomDomain
	err := ctx.RegisterResource("cloudamqp:index/customDomain:CustomDomain", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCustomDomain gets an existing CustomDomain resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCustomDomain(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CustomDomainState, opts ...pulumi.ResourceOption) (*CustomDomain, error) {
	var resource CustomDomain
	err := ctx.ReadResource("cloudamqp:index/customDomain:CustomDomain", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CustomDomain resources.
type customDomainState struct {
	// Your custom domain name.
	Hostname *string `pulumi:"hostname"`
	// The CloudAMQP instance ID.
	InstanceId *int `pulumi:"instanceId"`
}

type CustomDomainState struct {
	// Your custom domain name.
	Hostname pulumi.StringPtrInput
	// The CloudAMQP instance ID.
	InstanceId pulumi.IntPtrInput
}

func (CustomDomainState) ElementType() reflect.Type {
	return reflect.TypeOf((*customDomainState)(nil)).Elem()
}

type customDomainArgs struct {
	// Your custom domain name.
	Hostname string `pulumi:"hostname"`
	// The CloudAMQP instance ID.
	InstanceId int `pulumi:"instanceId"`
}

// The set of arguments for constructing a CustomDomain resource.
type CustomDomainArgs struct {
	// Your custom domain name.
	Hostname pulumi.StringInput
	// The CloudAMQP instance ID.
	InstanceId pulumi.IntInput
}

func (CustomDomainArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*customDomainArgs)(nil)).Elem()
}

type CustomDomainInput interface {
	pulumi.Input

	ToCustomDomainOutput() CustomDomainOutput
	ToCustomDomainOutputWithContext(ctx context.Context) CustomDomainOutput
}

func (*CustomDomain) ElementType() reflect.Type {
	return reflect.TypeOf((*CustomDomain)(nil))
}

func (i *CustomDomain) ToCustomDomainOutput() CustomDomainOutput {
	return i.ToCustomDomainOutputWithContext(context.Background())
}

func (i *CustomDomain) ToCustomDomainOutputWithContext(ctx context.Context) CustomDomainOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomDomainOutput)
}

func (i *CustomDomain) ToCustomDomainPtrOutput() CustomDomainPtrOutput {
	return i.ToCustomDomainPtrOutputWithContext(context.Background())
}

func (i *CustomDomain) ToCustomDomainPtrOutputWithContext(ctx context.Context) CustomDomainPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomDomainPtrOutput)
}

type CustomDomainPtrInput interface {
	pulumi.Input

	ToCustomDomainPtrOutput() CustomDomainPtrOutput
	ToCustomDomainPtrOutputWithContext(ctx context.Context) CustomDomainPtrOutput
}

type customDomainPtrType CustomDomainArgs

func (*customDomainPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CustomDomain)(nil))
}

func (i *customDomainPtrType) ToCustomDomainPtrOutput() CustomDomainPtrOutput {
	return i.ToCustomDomainPtrOutputWithContext(context.Background())
}

func (i *customDomainPtrType) ToCustomDomainPtrOutputWithContext(ctx context.Context) CustomDomainPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomDomainPtrOutput)
}

// CustomDomainArrayInput is an input type that accepts CustomDomainArray and CustomDomainArrayOutput values.
// You can construct a concrete instance of `CustomDomainArrayInput` via:
//
//          CustomDomainArray{ CustomDomainArgs{...} }
type CustomDomainArrayInput interface {
	pulumi.Input

	ToCustomDomainArrayOutput() CustomDomainArrayOutput
	ToCustomDomainArrayOutputWithContext(context.Context) CustomDomainArrayOutput
}

type CustomDomainArray []CustomDomainInput

func (CustomDomainArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CustomDomain)(nil)).Elem()
}

func (i CustomDomainArray) ToCustomDomainArrayOutput() CustomDomainArrayOutput {
	return i.ToCustomDomainArrayOutputWithContext(context.Background())
}

func (i CustomDomainArray) ToCustomDomainArrayOutputWithContext(ctx context.Context) CustomDomainArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomDomainArrayOutput)
}

// CustomDomainMapInput is an input type that accepts CustomDomainMap and CustomDomainMapOutput values.
// You can construct a concrete instance of `CustomDomainMapInput` via:
//
//          CustomDomainMap{ "key": CustomDomainArgs{...} }
type CustomDomainMapInput interface {
	pulumi.Input

	ToCustomDomainMapOutput() CustomDomainMapOutput
	ToCustomDomainMapOutputWithContext(context.Context) CustomDomainMapOutput
}

type CustomDomainMap map[string]CustomDomainInput

func (CustomDomainMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CustomDomain)(nil)).Elem()
}

func (i CustomDomainMap) ToCustomDomainMapOutput() CustomDomainMapOutput {
	return i.ToCustomDomainMapOutputWithContext(context.Background())
}

func (i CustomDomainMap) ToCustomDomainMapOutputWithContext(ctx context.Context) CustomDomainMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomDomainMapOutput)
}

type CustomDomainOutput struct{ *pulumi.OutputState }

func (CustomDomainOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CustomDomain)(nil))
}

func (o CustomDomainOutput) ToCustomDomainOutput() CustomDomainOutput {
	return o
}

func (o CustomDomainOutput) ToCustomDomainOutputWithContext(ctx context.Context) CustomDomainOutput {
	return o
}

func (o CustomDomainOutput) ToCustomDomainPtrOutput() CustomDomainPtrOutput {
	return o.ToCustomDomainPtrOutputWithContext(context.Background())
}

func (o CustomDomainOutput) ToCustomDomainPtrOutputWithContext(ctx context.Context) CustomDomainPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CustomDomain) *CustomDomain {
		return &v
	}).(CustomDomainPtrOutput)
}

type CustomDomainPtrOutput struct{ *pulumi.OutputState }

func (CustomDomainPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CustomDomain)(nil))
}

func (o CustomDomainPtrOutput) ToCustomDomainPtrOutput() CustomDomainPtrOutput {
	return o
}

func (o CustomDomainPtrOutput) ToCustomDomainPtrOutputWithContext(ctx context.Context) CustomDomainPtrOutput {
	return o
}

func (o CustomDomainPtrOutput) Elem() CustomDomainOutput {
	return o.ApplyT(func(v *CustomDomain) CustomDomain {
		if v != nil {
			return *v
		}
		var ret CustomDomain
		return ret
	}).(CustomDomainOutput)
}

type CustomDomainArrayOutput struct{ *pulumi.OutputState }

func (CustomDomainArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]CustomDomain)(nil))
}

func (o CustomDomainArrayOutput) ToCustomDomainArrayOutput() CustomDomainArrayOutput {
	return o
}

func (o CustomDomainArrayOutput) ToCustomDomainArrayOutputWithContext(ctx context.Context) CustomDomainArrayOutput {
	return o
}

func (o CustomDomainArrayOutput) Index(i pulumi.IntInput) CustomDomainOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) CustomDomain {
		return vs[0].([]CustomDomain)[vs[1].(int)]
	}).(CustomDomainOutput)
}

type CustomDomainMapOutput struct{ *pulumi.OutputState }

func (CustomDomainMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]CustomDomain)(nil))
}

func (o CustomDomainMapOutput) ToCustomDomainMapOutput() CustomDomainMapOutput {
	return o
}

func (o CustomDomainMapOutput) ToCustomDomainMapOutputWithContext(ctx context.Context) CustomDomainMapOutput {
	return o
}

func (o CustomDomainMapOutput) MapIndex(k pulumi.StringInput) CustomDomainOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) CustomDomain {
		return vs[0].(map[string]CustomDomain)[vs[1].(string)]
	}).(CustomDomainOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*CustomDomainInput)(nil)).Elem(), &CustomDomain{})
	pulumi.RegisterInputType(reflect.TypeOf((*CustomDomainPtrInput)(nil)).Elem(), &CustomDomain{})
	pulumi.RegisterInputType(reflect.TypeOf((*CustomDomainArrayInput)(nil)).Elem(), CustomDomainArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*CustomDomainMapInput)(nil)).Elem(), CustomDomainMap{})
	pulumi.RegisterOutputType(CustomDomainOutput{})
	pulumi.RegisterOutputType(CustomDomainPtrOutput{})
	pulumi.RegisterOutputType(CustomDomainArrayOutput{})
	pulumi.RegisterOutputType(CustomDomainMapOutput{})
}
