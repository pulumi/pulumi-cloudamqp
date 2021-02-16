// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package cloudamqp

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// This resource allows you to configure and manage firewall rules for the CloudAMQP instance. Beware that all rules need to be present, since all older configurations will be overwritten.
//
// Only available for dedicated subscription plans.
//
// ## Depedency
//
// This resource depends on CloudAMQP instance identifier, `cloudamqp_instance.instance.id`.
//
// ## Import
//
// `cloudamqp_security_firewall` can be imported using CloudAMQP instance identifier.
//
// ```sh
//  $ pulumi import cloudamqp:index/securityFirewall:SecurityFirewall firewall <instance_id>`
// ```
type SecurityFirewall struct {
	pulumi.CustomResourceState

	// The CloudAMQP instance ID.
	InstanceId pulumi.IntOutput `pulumi:"instanceId"`
	// An array of rules, minimum of 1 needs to be configured. Each `rules` block consists of the field documented below.
	Rules SecurityFirewallRuleArrayOutput `pulumi:"rules"`
}

// NewSecurityFirewall registers a new resource with the given unique name, arguments, and options.
func NewSecurityFirewall(ctx *pulumi.Context,
	name string, args *SecurityFirewallArgs, opts ...pulumi.ResourceOption) (*SecurityFirewall, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.InstanceId == nil {
		return nil, errors.New("invalid value for required argument 'InstanceId'")
	}
	if args.Rules == nil {
		return nil, errors.New("invalid value for required argument 'Rules'")
	}
	var resource SecurityFirewall
	err := ctx.RegisterResource("cloudamqp:index/securityFirewall:SecurityFirewall", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSecurityFirewall gets an existing SecurityFirewall resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSecurityFirewall(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SecurityFirewallState, opts ...pulumi.ResourceOption) (*SecurityFirewall, error) {
	var resource SecurityFirewall
	err := ctx.ReadResource("cloudamqp:index/securityFirewall:SecurityFirewall", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SecurityFirewall resources.
type securityFirewallState struct {
	// The CloudAMQP instance ID.
	InstanceId *int `pulumi:"instanceId"`
	// An array of rules, minimum of 1 needs to be configured. Each `rules` block consists of the field documented below.
	Rules []SecurityFirewallRule `pulumi:"rules"`
}

type SecurityFirewallState struct {
	// The CloudAMQP instance ID.
	InstanceId pulumi.IntPtrInput
	// An array of rules, minimum of 1 needs to be configured. Each `rules` block consists of the field documented below.
	Rules SecurityFirewallRuleArrayInput
}

func (SecurityFirewallState) ElementType() reflect.Type {
	return reflect.TypeOf((*securityFirewallState)(nil)).Elem()
}

type securityFirewallArgs struct {
	// The CloudAMQP instance ID.
	InstanceId int `pulumi:"instanceId"`
	// An array of rules, minimum of 1 needs to be configured. Each `rules` block consists of the field documented below.
	Rules []SecurityFirewallRule `pulumi:"rules"`
}

// The set of arguments for constructing a SecurityFirewall resource.
type SecurityFirewallArgs struct {
	// The CloudAMQP instance ID.
	InstanceId pulumi.IntInput
	// An array of rules, minimum of 1 needs to be configured. Each `rules` block consists of the field documented below.
	Rules SecurityFirewallRuleArrayInput
}

func (SecurityFirewallArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*securityFirewallArgs)(nil)).Elem()
}

type SecurityFirewallInput interface {
	pulumi.Input

	ToSecurityFirewallOutput() SecurityFirewallOutput
	ToSecurityFirewallOutputWithContext(ctx context.Context) SecurityFirewallOutput
}

func (*SecurityFirewall) ElementType() reflect.Type {
	return reflect.TypeOf((*SecurityFirewall)(nil))
}

func (i *SecurityFirewall) ToSecurityFirewallOutput() SecurityFirewallOutput {
	return i.ToSecurityFirewallOutputWithContext(context.Background())
}

func (i *SecurityFirewall) ToSecurityFirewallOutputWithContext(ctx context.Context) SecurityFirewallOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecurityFirewallOutput)
}

func (i *SecurityFirewall) ToSecurityFirewallPtrOutput() SecurityFirewallPtrOutput {
	return i.ToSecurityFirewallPtrOutputWithContext(context.Background())
}

func (i *SecurityFirewall) ToSecurityFirewallPtrOutputWithContext(ctx context.Context) SecurityFirewallPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecurityFirewallPtrOutput)
}

type SecurityFirewallPtrInput interface {
	pulumi.Input

	ToSecurityFirewallPtrOutput() SecurityFirewallPtrOutput
	ToSecurityFirewallPtrOutputWithContext(ctx context.Context) SecurityFirewallPtrOutput
}

type securityFirewallPtrType SecurityFirewallArgs

func (*securityFirewallPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SecurityFirewall)(nil))
}

func (i *securityFirewallPtrType) ToSecurityFirewallPtrOutput() SecurityFirewallPtrOutput {
	return i.ToSecurityFirewallPtrOutputWithContext(context.Background())
}

func (i *securityFirewallPtrType) ToSecurityFirewallPtrOutputWithContext(ctx context.Context) SecurityFirewallPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecurityFirewallPtrOutput)
}

// SecurityFirewallArrayInput is an input type that accepts SecurityFirewallArray and SecurityFirewallArrayOutput values.
// You can construct a concrete instance of `SecurityFirewallArrayInput` via:
//
//          SecurityFirewallArray{ SecurityFirewallArgs{...} }
type SecurityFirewallArrayInput interface {
	pulumi.Input

	ToSecurityFirewallArrayOutput() SecurityFirewallArrayOutput
	ToSecurityFirewallArrayOutputWithContext(context.Context) SecurityFirewallArrayOutput
}

type SecurityFirewallArray []SecurityFirewallInput

func (SecurityFirewallArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*SecurityFirewall)(nil))
}

func (i SecurityFirewallArray) ToSecurityFirewallArrayOutput() SecurityFirewallArrayOutput {
	return i.ToSecurityFirewallArrayOutputWithContext(context.Background())
}

func (i SecurityFirewallArray) ToSecurityFirewallArrayOutputWithContext(ctx context.Context) SecurityFirewallArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecurityFirewallArrayOutput)
}

// SecurityFirewallMapInput is an input type that accepts SecurityFirewallMap and SecurityFirewallMapOutput values.
// You can construct a concrete instance of `SecurityFirewallMapInput` via:
//
//          SecurityFirewallMap{ "key": SecurityFirewallArgs{...} }
type SecurityFirewallMapInput interface {
	pulumi.Input

	ToSecurityFirewallMapOutput() SecurityFirewallMapOutput
	ToSecurityFirewallMapOutputWithContext(context.Context) SecurityFirewallMapOutput
}

type SecurityFirewallMap map[string]SecurityFirewallInput

func (SecurityFirewallMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*SecurityFirewall)(nil))
}

func (i SecurityFirewallMap) ToSecurityFirewallMapOutput() SecurityFirewallMapOutput {
	return i.ToSecurityFirewallMapOutputWithContext(context.Background())
}

func (i SecurityFirewallMap) ToSecurityFirewallMapOutputWithContext(ctx context.Context) SecurityFirewallMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecurityFirewallMapOutput)
}

type SecurityFirewallOutput struct {
	*pulumi.OutputState
}

func (SecurityFirewallOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SecurityFirewall)(nil))
}

func (o SecurityFirewallOutput) ToSecurityFirewallOutput() SecurityFirewallOutput {
	return o
}

func (o SecurityFirewallOutput) ToSecurityFirewallOutputWithContext(ctx context.Context) SecurityFirewallOutput {
	return o
}

func (o SecurityFirewallOutput) ToSecurityFirewallPtrOutput() SecurityFirewallPtrOutput {
	return o.ToSecurityFirewallPtrOutputWithContext(context.Background())
}

func (o SecurityFirewallOutput) ToSecurityFirewallPtrOutputWithContext(ctx context.Context) SecurityFirewallPtrOutput {
	return o.ApplyT(func(v SecurityFirewall) *SecurityFirewall {
		return &v
	}).(SecurityFirewallPtrOutput)
}

type SecurityFirewallPtrOutput struct {
	*pulumi.OutputState
}

func (SecurityFirewallPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SecurityFirewall)(nil))
}

func (o SecurityFirewallPtrOutput) ToSecurityFirewallPtrOutput() SecurityFirewallPtrOutput {
	return o
}

func (o SecurityFirewallPtrOutput) ToSecurityFirewallPtrOutputWithContext(ctx context.Context) SecurityFirewallPtrOutput {
	return o
}

type SecurityFirewallArrayOutput struct{ *pulumi.OutputState }

func (SecurityFirewallArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SecurityFirewall)(nil))
}

func (o SecurityFirewallArrayOutput) ToSecurityFirewallArrayOutput() SecurityFirewallArrayOutput {
	return o
}

func (o SecurityFirewallArrayOutput) ToSecurityFirewallArrayOutputWithContext(ctx context.Context) SecurityFirewallArrayOutput {
	return o
}

func (o SecurityFirewallArrayOutput) Index(i pulumi.IntInput) SecurityFirewallOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SecurityFirewall {
		return vs[0].([]SecurityFirewall)[vs[1].(int)]
	}).(SecurityFirewallOutput)
}

type SecurityFirewallMapOutput struct{ *pulumi.OutputState }

func (SecurityFirewallMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]SecurityFirewall)(nil))
}

func (o SecurityFirewallMapOutput) ToSecurityFirewallMapOutput() SecurityFirewallMapOutput {
	return o
}

func (o SecurityFirewallMapOutput) ToSecurityFirewallMapOutputWithContext(ctx context.Context) SecurityFirewallMapOutput {
	return o
}

func (o SecurityFirewallMapOutput) MapIndex(k pulumi.StringInput) SecurityFirewallOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) SecurityFirewall {
		return vs[0].(map[string]SecurityFirewall)[vs[1].(string)]
	}).(SecurityFirewallOutput)
}

func init() {
	pulumi.RegisterOutputType(SecurityFirewallOutput{})
	pulumi.RegisterOutputType(SecurityFirewallPtrOutput{})
	pulumi.RegisterOutputType(SecurityFirewallArrayOutput{})
	pulumi.RegisterOutputType(SecurityFirewallMapOutput{})
}
