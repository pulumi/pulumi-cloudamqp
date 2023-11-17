// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cloudamqp

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-cloudamqp/sdk/v3/go/cloudamqp/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## Import
//
// `cloudamqp_security_firewall` can be imported using CloudAMQP instance identifier.
//
// ```sh
//
//	$ pulumi import cloudamqp:index/securityFirewall:SecurityFirewall firewall <instance_id>`
//
// ```
type SecurityFirewall struct {
	pulumi.CustomResourceState

	// The CloudAMQP instance ID.
	InstanceId pulumi.IntOutput `pulumi:"instanceId"`
	// An array of rules, minimum of 1 needs to be configured. Each `rules` block consists of the field documented below.
	Rules SecurityFirewallRuleArrayOutput `pulumi:"rules"`
	// Configurable sleep time in seconds between retries for firewall configuration. Default set to 30 seconds.
	Sleep pulumi.IntPtrOutput `pulumi:"sleep"`
	// Configurable timeout time in seconds for firewall configuration. Default set to 1800 seconds.
	//
	// The `rules` block consists of:
	Timeout pulumi.IntPtrOutput `pulumi:"timeout"`
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
	opts = internal.PkgResourceDefaultOpts(opts)
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
	// Configurable sleep time in seconds between retries for firewall configuration. Default set to 30 seconds.
	Sleep *int `pulumi:"sleep"`
	// Configurable timeout time in seconds for firewall configuration. Default set to 1800 seconds.
	//
	// The `rules` block consists of:
	Timeout *int `pulumi:"timeout"`
}

type SecurityFirewallState struct {
	// The CloudAMQP instance ID.
	InstanceId pulumi.IntPtrInput
	// An array of rules, minimum of 1 needs to be configured. Each `rules` block consists of the field documented below.
	Rules SecurityFirewallRuleArrayInput
	// Configurable sleep time in seconds between retries for firewall configuration. Default set to 30 seconds.
	Sleep pulumi.IntPtrInput
	// Configurable timeout time in seconds for firewall configuration. Default set to 1800 seconds.
	//
	// The `rules` block consists of:
	Timeout pulumi.IntPtrInput
}

func (SecurityFirewallState) ElementType() reflect.Type {
	return reflect.TypeOf((*securityFirewallState)(nil)).Elem()
}

type securityFirewallArgs struct {
	// The CloudAMQP instance ID.
	InstanceId int `pulumi:"instanceId"`
	// An array of rules, minimum of 1 needs to be configured. Each `rules` block consists of the field documented below.
	Rules []SecurityFirewallRule `pulumi:"rules"`
	// Configurable sleep time in seconds between retries for firewall configuration. Default set to 30 seconds.
	Sleep *int `pulumi:"sleep"`
	// Configurable timeout time in seconds for firewall configuration. Default set to 1800 seconds.
	//
	// The `rules` block consists of:
	Timeout *int `pulumi:"timeout"`
}

// The set of arguments for constructing a SecurityFirewall resource.
type SecurityFirewallArgs struct {
	// The CloudAMQP instance ID.
	InstanceId pulumi.IntInput
	// An array of rules, minimum of 1 needs to be configured. Each `rules` block consists of the field documented below.
	Rules SecurityFirewallRuleArrayInput
	// Configurable sleep time in seconds between retries for firewall configuration. Default set to 30 seconds.
	Sleep pulumi.IntPtrInput
	// Configurable timeout time in seconds for firewall configuration. Default set to 1800 seconds.
	//
	// The `rules` block consists of:
	Timeout pulumi.IntPtrInput
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
	return reflect.TypeOf((**SecurityFirewall)(nil)).Elem()
}

func (i *SecurityFirewall) ToSecurityFirewallOutput() SecurityFirewallOutput {
	return i.ToSecurityFirewallOutputWithContext(context.Background())
}

func (i *SecurityFirewall) ToSecurityFirewallOutputWithContext(ctx context.Context) SecurityFirewallOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecurityFirewallOutput)
}

// SecurityFirewallArrayInput is an input type that accepts SecurityFirewallArray and SecurityFirewallArrayOutput values.
// You can construct a concrete instance of `SecurityFirewallArrayInput` via:
//
//	SecurityFirewallArray{ SecurityFirewallArgs{...} }
type SecurityFirewallArrayInput interface {
	pulumi.Input

	ToSecurityFirewallArrayOutput() SecurityFirewallArrayOutput
	ToSecurityFirewallArrayOutputWithContext(context.Context) SecurityFirewallArrayOutput
}

type SecurityFirewallArray []SecurityFirewallInput

func (SecurityFirewallArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SecurityFirewall)(nil)).Elem()
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
//	SecurityFirewallMap{ "key": SecurityFirewallArgs{...} }
type SecurityFirewallMapInput interface {
	pulumi.Input

	ToSecurityFirewallMapOutput() SecurityFirewallMapOutput
	ToSecurityFirewallMapOutputWithContext(context.Context) SecurityFirewallMapOutput
}

type SecurityFirewallMap map[string]SecurityFirewallInput

func (SecurityFirewallMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SecurityFirewall)(nil)).Elem()
}

func (i SecurityFirewallMap) ToSecurityFirewallMapOutput() SecurityFirewallMapOutput {
	return i.ToSecurityFirewallMapOutputWithContext(context.Background())
}

func (i SecurityFirewallMap) ToSecurityFirewallMapOutputWithContext(ctx context.Context) SecurityFirewallMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecurityFirewallMapOutput)
}

type SecurityFirewallOutput struct{ *pulumi.OutputState }

func (SecurityFirewallOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SecurityFirewall)(nil)).Elem()
}

func (o SecurityFirewallOutput) ToSecurityFirewallOutput() SecurityFirewallOutput {
	return o
}

func (o SecurityFirewallOutput) ToSecurityFirewallOutputWithContext(ctx context.Context) SecurityFirewallOutput {
	return o
}

// The CloudAMQP instance ID.
func (o SecurityFirewallOutput) InstanceId() pulumi.IntOutput {
	return o.ApplyT(func(v *SecurityFirewall) pulumi.IntOutput { return v.InstanceId }).(pulumi.IntOutput)
}

// An array of rules, minimum of 1 needs to be configured. Each `rules` block consists of the field documented below.
func (o SecurityFirewallOutput) Rules() SecurityFirewallRuleArrayOutput {
	return o.ApplyT(func(v *SecurityFirewall) SecurityFirewallRuleArrayOutput { return v.Rules }).(SecurityFirewallRuleArrayOutput)
}

// Configurable sleep time in seconds between retries for firewall configuration. Default set to 30 seconds.
func (o SecurityFirewallOutput) Sleep() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SecurityFirewall) pulumi.IntPtrOutput { return v.Sleep }).(pulumi.IntPtrOutput)
}

// Configurable timeout time in seconds for firewall configuration. Default set to 1800 seconds.
//
// The `rules` block consists of:
func (o SecurityFirewallOutput) Timeout() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SecurityFirewall) pulumi.IntPtrOutput { return v.Timeout }).(pulumi.IntPtrOutput)
}

type SecurityFirewallArrayOutput struct{ *pulumi.OutputState }

func (SecurityFirewallArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SecurityFirewall)(nil)).Elem()
}

func (o SecurityFirewallArrayOutput) ToSecurityFirewallArrayOutput() SecurityFirewallArrayOutput {
	return o
}

func (o SecurityFirewallArrayOutput) ToSecurityFirewallArrayOutputWithContext(ctx context.Context) SecurityFirewallArrayOutput {
	return o
}

func (o SecurityFirewallArrayOutput) Index(i pulumi.IntInput) SecurityFirewallOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SecurityFirewall {
		return vs[0].([]*SecurityFirewall)[vs[1].(int)]
	}).(SecurityFirewallOutput)
}

type SecurityFirewallMapOutput struct{ *pulumi.OutputState }

func (SecurityFirewallMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SecurityFirewall)(nil)).Elem()
}

func (o SecurityFirewallMapOutput) ToSecurityFirewallMapOutput() SecurityFirewallMapOutput {
	return o
}

func (o SecurityFirewallMapOutput) ToSecurityFirewallMapOutputWithContext(ctx context.Context) SecurityFirewallMapOutput {
	return o
}

func (o SecurityFirewallMapOutput) MapIndex(k pulumi.StringInput) SecurityFirewallOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SecurityFirewall {
		return vs[0].(map[string]*SecurityFirewall)[vs[1].(string)]
	}).(SecurityFirewallOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SecurityFirewallInput)(nil)).Elem(), &SecurityFirewall{})
	pulumi.RegisterInputType(reflect.TypeOf((*SecurityFirewallArrayInput)(nil)).Elem(), SecurityFirewallArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SecurityFirewallMapInput)(nil)).Elem(), SecurityFirewallMap{})
	pulumi.RegisterOutputType(SecurityFirewallOutput{})
	pulumi.RegisterOutputType(SecurityFirewallArrayOutput{})
	pulumi.RegisterOutputType(SecurityFirewallMapOutput{})
}
