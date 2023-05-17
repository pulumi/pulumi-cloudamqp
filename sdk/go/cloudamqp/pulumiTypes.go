// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cloudamqp

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ExtraDiskSizeNode struct {
	AdditionalDiskSize *int    `pulumi:"additionalDiskSize"`
	DiskSize           *int    `pulumi:"diskSize"`
	Name               *string `pulumi:"name"`
}

// ExtraDiskSizeNodeInput is an input type that accepts ExtraDiskSizeNodeArgs and ExtraDiskSizeNodeOutput values.
// You can construct a concrete instance of `ExtraDiskSizeNodeInput` via:
//
//	ExtraDiskSizeNodeArgs{...}
type ExtraDiskSizeNodeInput interface {
	pulumi.Input

	ToExtraDiskSizeNodeOutput() ExtraDiskSizeNodeOutput
	ToExtraDiskSizeNodeOutputWithContext(context.Context) ExtraDiskSizeNodeOutput
}

type ExtraDiskSizeNodeArgs struct {
	AdditionalDiskSize pulumi.IntPtrInput    `pulumi:"additionalDiskSize"`
	DiskSize           pulumi.IntPtrInput    `pulumi:"diskSize"`
	Name               pulumi.StringPtrInput `pulumi:"name"`
}

func (ExtraDiskSizeNodeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ExtraDiskSizeNode)(nil)).Elem()
}

func (i ExtraDiskSizeNodeArgs) ToExtraDiskSizeNodeOutput() ExtraDiskSizeNodeOutput {
	return i.ToExtraDiskSizeNodeOutputWithContext(context.Background())
}

func (i ExtraDiskSizeNodeArgs) ToExtraDiskSizeNodeOutputWithContext(ctx context.Context) ExtraDiskSizeNodeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExtraDiskSizeNodeOutput)
}

// ExtraDiskSizeNodeArrayInput is an input type that accepts ExtraDiskSizeNodeArray and ExtraDiskSizeNodeArrayOutput values.
// You can construct a concrete instance of `ExtraDiskSizeNodeArrayInput` via:
//
//	ExtraDiskSizeNodeArray{ ExtraDiskSizeNodeArgs{...} }
type ExtraDiskSizeNodeArrayInput interface {
	pulumi.Input

	ToExtraDiskSizeNodeArrayOutput() ExtraDiskSizeNodeArrayOutput
	ToExtraDiskSizeNodeArrayOutputWithContext(context.Context) ExtraDiskSizeNodeArrayOutput
}

type ExtraDiskSizeNodeArray []ExtraDiskSizeNodeInput

func (ExtraDiskSizeNodeArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ExtraDiskSizeNode)(nil)).Elem()
}

func (i ExtraDiskSizeNodeArray) ToExtraDiskSizeNodeArrayOutput() ExtraDiskSizeNodeArrayOutput {
	return i.ToExtraDiskSizeNodeArrayOutputWithContext(context.Background())
}

func (i ExtraDiskSizeNodeArray) ToExtraDiskSizeNodeArrayOutputWithContext(ctx context.Context) ExtraDiskSizeNodeArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExtraDiskSizeNodeArrayOutput)
}

type ExtraDiskSizeNodeOutput struct{ *pulumi.OutputState }

func (ExtraDiskSizeNodeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ExtraDiskSizeNode)(nil)).Elem()
}

func (o ExtraDiskSizeNodeOutput) ToExtraDiskSizeNodeOutput() ExtraDiskSizeNodeOutput {
	return o
}

func (o ExtraDiskSizeNodeOutput) ToExtraDiskSizeNodeOutputWithContext(ctx context.Context) ExtraDiskSizeNodeOutput {
	return o
}

func (o ExtraDiskSizeNodeOutput) AdditionalDiskSize() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ExtraDiskSizeNode) *int { return v.AdditionalDiskSize }).(pulumi.IntPtrOutput)
}

func (o ExtraDiskSizeNodeOutput) DiskSize() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ExtraDiskSizeNode) *int { return v.DiskSize }).(pulumi.IntPtrOutput)
}

func (o ExtraDiskSizeNodeOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ExtraDiskSizeNode) *string { return v.Name }).(pulumi.StringPtrOutput)
}

type ExtraDiskSizeNodeArrayOutput struct{ *pulumi.OutputState }

func (ExtraDiskSizeNodeArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ExtraDiskSizeNode)(nil)).Elem()
}

func (o ExtraDiskSizeNodeArrayOutput) ToExtraDiskSizeNodeArrayOutput() ExtraDiskSizeNodeArrayOutput {
	return o
}

func (o ExtraDiskSizeNodeArrayOutput) ToExtraDiskSizeNodeArrayOutputWithContext(ctx context.Context) ExtraDiskSizeNodeArrayOutput {
	return o
}

func (o ExtraDiskSizeNodeArrayOutput) Index(i pulumi.IntInput) ExtraDiskSizeNodeOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ExtraDiskSizeNode {
		return vs[0].([]ExtraDiskSizeNode)[vs[1].(int)]
	}).(ExtraDiskSizeNodeOutput)
}

type SecurityFirewallRule struct {
	// Description name of the rule. e.g. Default.
	//
	// Pre-defined services for RabbitMQ:
	//
	// | Service name | Port  |
	// |--------------|-------|
	// | AMQP         | 5672  |
	// | AMQPS        | 5671  |
	// | HTTPS        | 443   |
	// | MQTT         | 1883  |
	// | MQTTS        | 8883  |
	// | STOMP        | 61613 |
	// | STOMPS       | 61614 |
	// | STREAM       | 5552  |
	// | STREAM_SSL   | 5551  |
	//
	// Pre-defined services for LavinMQ:
	//
	// | Service name | Port  |
	// |--------------|-------|
	// | AMQP         | 5672  |
	// | AMQPS        | 5671  |
	// | HTTPS        | 443   |
	Description *string `pulumi:"description"`
	// CIDR address: IP address with CIDR notation (e.g. 10.56.72.0/24)
	Ip string `pulumi:"ip"`
	// Custom ports to be opened
	Ports []int `pulumi:"ports"`
	// Pre-defined service ports, see table below
	Services []string `pulumi:"services"`
}

// SecurityFirewallRuleInput is an input type that accepts SecurityFirewallRuleArgs and SecurityFirewallRuleOutput values.
// You can construct a concrete instance of `SecurityFirewallRuleInput` via:
//
//	SecurityFirewallRuleArgs{...}
type SecurityFirewallRuleInput interface {
	pulumi.Input

	ToSecurityFirewallRuleOutput() SecurityFirewallRuleOutput
	ToSecurityFirewallRuleOutputWithContext(context.Context) SecurityFirewallRuleOutput
}

type SecurityFirewallRuleArgs struct {
	// Description name of the rule. e.g. Default.
	//
	// Pre-defined services for RabbitMQ:
	//
	// | Service name | Port  |
	// |--------------|-------|
	// | AMQP         | 5672  |
	// | AMQPS        | 5671  |
	// | HTTPS        | 443   |
	// | MQTT         | 1883  |
	// | MQTTS        | 8883  |
	// | STOMP        | 61613 |
	// | STOMPS       | 61614 |
	// | STREAM       | 5552  |
	// | STREAM_SSL   | 5551  |
	//
	// Pre-defined services for LavinMQ:
	//
	// | Service name | Port  |
	// |--------------|-------|
	// | AMQP         | 5672  |
	// | AMQPS        | 5671  |
	// | HTTPS        | 443   |
	Description pulumi.StringPtrInput `pulumi:"description"`
	// CIDR address: IP address with CIDR notation (e.g. 10.56.72.0/24)
	Ip pulumi.StringInput `pulumi:"ip"`
	// Custom ports to be opened
	Ports pulumi.IntArrayInput `pulumi:"ports"`
	// Pre-defined service ports, see table below
	Services pulumi.StringArrayInput `pulumi:"services"`
}

func (SecurityFirewallRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SecurityFirewallRule)(nil)).Elem()
}

func (i SecurityFirewallRuleArgs) ToSecurityFirewallRuleOutput() SecurityFirewallRuleOutput {
	return i.ToSecurityFirewallRuleOutputWithContext(context.Background())
}

func (i SecurityFirewallRuleArgs) ToSecurityFirewallRuleOutputWithContext(ctx context.Context) SecurityFirewallRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecurityFirewallRuleOutput)
}

// SecurityFirewallRuleArrayInput is an input type that accepts SecurityFirewallRuleArray and SecurityFirewallRuleArrayOutput values.
// You can construct a concrete instance of `SecurityFirewallRuleArrayInput` via:
//
//	SecurityFirewallRuleArray{ SecurityFirewallRuleArgs{...} }
type SecurityFirewallRuleArrayInput interface {
	pulumi.Input

	ToSecurityFirewallRuleArrayOutput() SecurityFirewallRuleArrayOutput
	ToSecurityFirewallRuleArrayOutputWithContext(context.Context) SecurityFirewallRuleArrayOutput
}

type SecurityFirewallRuleArray []SecurityFirewallRuleInput

func (SecurityFirewallRuleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SecurityFirewallRule)(nil)).Elem()
}

func (i SecurityFirewallRuleArray) ToSecurityFirewallRuleArrayOutput() SecurityFirewallRuleArrayOutput {
	return i.ToSecurityFirewallRuleArrayOutputWithContext(context.Background())
}

func (i SecurityFirewallRuleArray) ToSecurityFirewallRuleArrayOutputWithContext(ctx context.Context) SecurityFirewallRuleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecurityFirewallRuleArrayOutput)
}

type SecurityFirewallRuleOutput struct{ *pulumi.OutputState }

func (SecurityFirewallRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SecurityFirewallRule)(nil)).Elem()
}

func (o SecurityFirewallRuleOutput) ToSecurityFirewallRuleOutput() SecurityFirewallRuleOutput {
	return o
}

func (o SecurityFirewallRuleOutput) ToSecurityFirewallRuleOutputWithContext(ctx context.Context) SecurityFirewallRuleOutput {
	return o
}

// Description name of the rule. e.g. Default.
//
// Pre-defined services for RabbitMQ:
//
// | Service name | Port  |
// |--------------|-------|
// | AMQP         | 5672  |
// | AMQPS        | 5671  |
// | HTTPS        | 443   |
// | MQTT         | 1883  |
// | MQTTS        | 8883  |
// | STOMP        | 61613 |
// | STOMPS       | 61614 |
// | STREAM       | 5552  |
// | STREAM_SSL   | 5551  |
//
// Pre-defined services for LavinMQ:
//
// | Service name | Port  |
// |--------------|-------|
// | AMQP         | 5672  |
// | AMQPS        | 5671  |
// | HTTPS        | 443   |
func (o SecurityFirewallRuleOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SecurityFirewallRule) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// CIDR address: IP address with CIDR notation (e.g. 10.56.72.0/24)
func (o SecurityFirewallRuleOutput) Ip() pulumi.StringOutput {
	return o.ApplyT(func(v SecurityFirewallRule) string { return v.Ip }).(pulumi.StringOutput)
}

// Custom ports to be opened
func (o SecurityFirewallRuleOutput) Ports() pulumi.IntArrayOutput {
	return o.ApplyT(func(v SecurityFirewallRule) []int { return v.Ports }).(pulumi.IntArrayOutput)
}

// Pre-defined service ports, see table below
func (o SecurityFirewallRuleOutput) Services() pulumi.StringArrayOutput {
	return o.ApplyT(func(v SecurityFirewallRule) []string { return v.Services }).(pulumi.StringArrayOutput)
}

type SecurityFirewallRuleArrayOutput struct{ *pulumi.OutputState }

func (SecurityFirewallRuleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SecurityFirewallRule)(nil)).Elem()
}

func (o SecurityFirewallRuleArrayOutput) ToSecurityFirewallRuleArrayOutput() SecurityFirewallRuleArrayOutput {
	return o
}

func (o SecurityFirewallRuleArrayOutput) ToSecurityFirewallRuleArrayOutputWithContext(ctx context.Context) SecurityFirewallRuleArrayOutput {
	return o
}

func (o SecurityFirewallRuleArrayOutput) Index(i pulumi.IntInput) SecurityFirewallRuleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SecurityFirewallRule {
		return vs[0].([]SecurityFirewallRule)[vs[1].(int)]
	}).(SecurityFirewallRuleOutput)
}

type GetAccountInstance struct {
	Id     int      `pulumi:"id"`
	Name   string   `pulumi:"name"`
	Plan   string   `pulumi:"plan"`
	Region string   `pulumi:"region"`
	Tags   []string `pulumi:"tags"`
}

// GetAccountInstanceInput is an input type that accepts GetAccountInstanceArgs and GetAccountInstanceOutput values.
// You can construct a concrete instance of `GetAccountInstanceInput` via:
//
//	GetAccountInstanceArgs{...}
type GetAccountInstanceInput interface {
	pulumi.Input

	ToGetAccountInstanceOutput() GetAccountInstanceOutput
	ToGetAccountInstanceOutputWithContext(context.Context) GetAccountInstanceOutput
}

type GetAccountInstanceArgs struct {
	Id     pulumi.IntInput         `pulumi:"id"`
	Name   pulumi.StringInput      `pulumi:"name"`
	Plan   pulumi.StringInput      `pulumi:"plan"`
	Region pulumi.StringInput      `pulumi:"region"`
	Tags   pulumi.StringArrayInput `pulumi:"tags"`
}

func (GetAccountInstanceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetAccountInstance)(nil)).Elem()
}

func (i GetAccountInstanceArgs) ToGetAccountInstanceOutput() GetAccountInstanceOutput {
	return i.ToGetAccountInstanceOutputWithContext(context.Background())
}

func (i GetAccountInstanceArgs) ToGetAccountInstanceOutputWithContext(ctx context.Context) GetAccountInstanceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetAccountInstanceOutput)
}

// GetAccountInstanceArrayInput is an input type that accepts GetAccountInstanceArray and GetAccountInstanceArrayOutput values.
// You can construct a concrete instance of `GetAccountInstanceArrayInput` via:
//
//	GetAccountInstanceArray{ GetAccountInstanceArgs{...} }
type GetAccountInstanceArrayInput interface {
	pulumi.Input

	ToGetAccountInstanceArrayOutput() GetAccountInstanceArrayOutput
	ToGetAccountInstanceArrayOutputWithContext(context.Context) GetAccountInstanceArrayOutput
}

type GetAccountInstanceArray []GetAccountInstanceInput

func (GetAccountInstanceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetAccountInstance)(nil)).Elem()
}

func (i GetAccountInstanceArray) ToGetAccountInstanceArrayOutput() GetAccountInstanceArrayOutput {
	return i.ToGetAccountInstanceArrayOutputWithContext(context.Background())
}

func (i GetAccountInstanceArray) ToGetAccountInstanceArrayOutputWithContext(ctx context.Context) GetAccountInstanceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetAccountInstanceArrayOutput)
}

type GetAccountInstanceOutput struct{ *pulumi.OutputState }

func (GetAccountInstanceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetAccountInstance)(nil)).Elem()
}

func (o GetAccountInstanceOutput) ToGetAccountInstanceOutput() GetAccountInstanceOutput {
	return o
}

func (o GetAccountInstanceOutput) ToGetAccountInstanceOutputWithContext(ctx context.Context) GetAccountInstanceOutput {
	return o
}

func (o GetAccountInstanceOutput) Id() pulumi.IntOutput {
	return o.ApplyT(func(v GetAccountInstance) int { return v.Id }).(pulumi.IntOutput)
}

func (o GetAccountInstanceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetAccountInstance) string { return v.Name }).(pulumi.StringOutput)
}

func (o GetAccountInstanceOutput) Plan() pulumi.StringOutput {
	return o.ApplyT(func(v GetAccountInstance) string { return v.Plan }).(pulumi.StringOutput)
}

func (o GetAccountInstanceOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v GetAccountInstance) string { return v.Region }).(pulumi.StringOutput)
}

func (o GetAccountInstanceOutput) Tags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetAccountInstance) []string { return v.Tags }).(pulumi.StringArrayOutput)
}

type GetAccountInstanceArrayOutput struct{ *pulumi.OutputState }

func (GetAccountInstanceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetAccountInstance)(nil)).Elem()
}

func (o GetAccountInstanceArrayOutput) ToGetAccountInstanceArrayOutput() GetAccountInstanceArrayOutput {
	return o
}

func (o GetAccountInstanceArrayOutput) ToGetAccountInstanceArrayOutputWithContext(ctx context.Context) GetAccountInstanceArrayOutput {
	return o
}

func (o GetAccountInstanceArrayOutput) Index(i pulumi.IntInput) GetAccountInstanceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) GetAccountInstance {
		return vs[0].([]GetAccountInstance)[vs[1].(int)]
	}).(GetAccountInstanceOutput)
}

type GetAccountVpcsVpc struct {
	Id      int      `pulumi:"id"`
	Name    string   `pulumi:"name"`
	Region  string   `pulumi:"region"`
	Subnet  string   `pulumi:"subnet"`
	Tags    []string `pulumi:"tags"`
	VpcName string   `pulumi:"vpcName"`
}

// GetAccountVpcsVpcInput is an input type that accepts GetAccountVpcsVpcArgs and GetAccountVpcsVpcOutput values.
// You can construct a concrete instance of `GetAccountVpcsVpcInput` via:
//
//	GetAccountVpcsVpcArgs{...}
type GetAccountVpcsVpcInput interface {
	pulumi.Input

	ToGetAccountVpcsVpcOutput() GetAccountVpcsVpcOutput
	ToGetAccountVpcsVpcOutputWithContext(context.Context) GetAccountVpcsVpcOutput
}

type GetAccountVpcsVpcArgs struct {
	Id      pulumi.IntInput         `pulumi:"id"`
	Name    pulumi.StringInput      `pulumi:"name"`
	Region  pulumi.StringInput      `pulumi:"region"`
	Subnet  pulumi.StringInput      `pulumi:"subnet"`
	Tags    pulumi.StringArrayInput `pulumi:"tags"`
	VpcName pulumi.StringInput      `pulumi:"vpcName"`
}

func (GetAccountVpcsVpcArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetAccountVpcsVpc)(nil)).Elem()
}

func (i GetAccountVpcsVpcArgs) ToGetAccountVpcsVpcOutput() GetAccountVpcsVpcOutput {
	return i.ToGetAccountVpcsVpcOutputWithContext(context.Background())
}

func (i GetAccountVpcsVpcArgs) ToGetAccountVpcsVpcOutputWithContext(ctx context.Context) GetAccountVpcsVpcOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetAccountVpcsVpcOutput)
}

// GetAccountVpcsVpcArrayInput is an input type that accepts GetAccountVpcsVpcArray and GetAccountVpcsVpcArrayOutput values.
// You can construct a concrete instance of `GetAccountVpcsVpcArrayInput` via:
//
//	GetAccountVpcsVpcArray{ GetAccountVpcsVpcArgs{...} }
type GetAccountVpcsVpcArrayInput interface {
	pulumi.Input

	ToGetAccountVpcsVpcArrayOutput() GetAccountVpcsVpcArrayOutput
	ToGetAccountVpcsVpcArrayOutputWithContext(context.Context) GetAccountVpcsVpcArrayOutput
}

type GetAccountVpcsVpcArray []GetAccountVpcsVpcInput

func (GetAccountVpcsVpcArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetAccountVpcsVpc)(nil)).Elem()
}

func (i GetAccountVpcsVpcArray) ToGetAccountVpcsVpcArrayOutput() GetAccountVpcsVpcArrayOutput {
	return i.ToGetAccountVpcsVpcArrayOutputWithContext(context.Background())
}

func (i GetAccountVpcsVpcArray) ToGetAccountVpcsVpcArrayOutputWithContext(ctx context.Context) GetAccountVpcsVpcArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetAccountVpcsVpcArrayOutput)
}

type GetAccountVpcsVpcOutput struct{ *pulumi.OutputState }

func (GetAccountVpcsVpcOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetAccountVpcsVpc)(nil)).Elem()
}

func (o GetAccountVpcsVpcOutput) ToGetAccountVpcsVpcOutput() GetAccountVpcsVpcOutput {
	return o
}

func (o GetAccountVpcsVpcOutput) ToGetAccountVpcsVpcOutputWithContext(ctx context.Context) GetAccountVpcsVpcOutput {
	return o
}

func (o GetAccountVpcsVpcOutput) Id() pulumi.IntOutput {
	return o.ApplyT(func(v GetAccountVpcsVpc) int { return v.Id }).(pulumi.IntOutput)
}

func (o GetAccountVpcsVpcOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetAccountVpcsVpc) string { return v.Name }).(pulumi.StringOutput)
}

func (o GetAccountVpcsVpcOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v GetAccountVpcsVpc) string { return v.Region }).(pulumi.StringOutput)
}

func (o GetAccountVpcsVpcOutput) Subnet() pulumi.StringOutput {
	return o.ApplyT(func(v GetAccountVpcsVpc) string { return v.Subnet }).(pulumi.StringOutput)
}

func (o GetAccountVpcsVpcOutput) Tags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetAccountVpcsVpc) []string { return v.Tags }).(pulumi.StringArrayOutput)
}

func (o GetAccountVpcsVpcOutput) VpcName() pulumi.StringOutput {
	return o.ApplyT(func(v GetAccountVpcsVpc) string { return v.VpcName }).(pulumi.StringOutput)
}

type GetAccountVpcsVpcArrayOutput struct{ *pulumi.OutputState }

func (GetAccountVpcsVpcArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetAccountVpcsVpc)(nil)).Elem()
}

func (o GetAccountVpcsVpcArrayOutput) ToGetAccountVpcsVpcArrayOutput() GetAccountVpcsVpcArrayOutput {
	return o
}

func (o GetAccountVpcsVpcArrayOutput) ToGetAccountVpcsVpcArrayOutputWithContext(ctx context.Context) GetAccountVpcsVpcArrayOutput {
	return o
}

func (o GetAccountVpcsVpcArrayOutput) Index(i pulumi.IntInput) GetAccountVpcsVpcOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) GetAccountVpcsVpc {
		return vs[0].([]GetAccountVpcsVpc)[vs[1].(int)]
	}).(GetAccountVpcsVpcOutput)
}

type GetNodesNode struct {
	AdditionalDiskSize int    `pulumi:"additionalDiskSize"`
	Configured         bool   `pulumi:"configured"`
	DiskSize           int    `pulumi:"diskSize"`
	ErlangVersion      string `pulumi:"erlangVersion"`
	Hipe               bool   `pulumi:"hipe"`
	Hostname           string `pulumi:"hostname"`
	Name               string `pulumi:"name"`
	RabbitmqVersion    string `pulumi:"rabbitmqVersion"`
	Running            bool   `pulumi:"running"`
}

// GetNodesNodeInput is an input type that accepts GetNodesNodeArgs and GetNodesNodeOutput values.
// You can construct a concrete instance of `GetNodesNodeInput` via:
//
//	GetNodesNodeArgs{...}
type GetNodesNodeInput interface {
	pulumi.Input

	ToGetNodesNodeOutput() GetNodesNodeOutput
	ToGetNodesNodeOutputWithContext(context.Context) GetNodesNodeOutput
}

type GetNodesNodeArgs struct {
	AdditionalDiskSize pulumi.IntInput    `pulumi:"additionalDiskSize"`
	Configured         pulumi.BoolInput   `pulumi:"configured"`
	DiskSize           pulumi.IntInput    `pulumi:"diskSize"`
	ErlangVersion      pulumi.StringInput `pulumi:"erlangVersion"`
	Hipe               pulumi.BoolInput   `pulumi:"hipe"`
	Hostname           pulumi.StringInput `pulumi:"hostname"`
	Name               pulumi.StringInput `pulumi:"name"`
	RabbitmqVersion    pulumi.StringInput `pulumi:"rabbitmqVersion"`
	Running            pulumi.BoolInput   `pulumi:"running"`
}

func (GetNodesNodeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetNodesNode)(nil)).Elem()
}

func (i GetNodesNodeArgs) ToGetNodesNodeOutput() GetNodesNodeOutput {
	return i.ToGetNodesNodeOutputWithContext(context.Background())
}

func (i GetNodesNodeArgs) ToGetNodesNodeOutputWithContext(ctx context.Context) GetNodesNodeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetNodesNodeOutput)
}

// GetNodesNodeArrayInput is an input type that accepts GetNodesNodeArray and GetNodesNodeArrayOutput values.
// You can construct a concrete instance of `GetNodesNodeArrayInput` via:
//
//	GetNodesNodeArray{ GetNodesNodeArgs{...} }
type GetNodesNodeArrayInput interface {
	pulumi.Input

	ToGetNodesNodeArrayOutput() GetNodesNodeArrayOutput
	ToGetNodesNodeArrayOutputWithContext(context.Context) GetNodesNodeArrayOutput
}

type GetNodesNodeArray []GetNodesNodeInput

func (GetNodesNodeArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetNodesNode)(nil)).Elem()
}

func (i GetNodesNodeArray) ToGetNodesNodeArrayOutput() GetNodesNodeArrayOutput {
	return i.ToGetNodesNodeArrayOutputWithContext(context.Background())
}

func (i GetNodesNodeArray) ToGetNodesNodeArrayOutputWithContext(ctx context.Context) GetNodesNodeArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetNodesNodeArrayOutput)
}

type GetNodesNodeOutput struct{ *pulumi.OutputState }

func (GetNodesNodeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetNodesNode)(nil)).Elem()
}

func (o GetNodesNodeOutput) ToGetNodesNodeOutput() GetNodesNodeOutput {
	return o
}

func (o GetNodesNodeOutput) ToGetNodesNodeOutputWithContext(ctx context.Context) GetNodesNodeOutput {
	return o
}

func (o GetNodesNodeOutput) AdditionalDiskSize() pulumi.IntOutput {
	return o.ApplyT(func(v GetNodesNode) int { return v.AdditionalDiskSize }).(pulumi.IntOutput)
}

func (o GetNodesNodeOutput) Configured() pulumi.BoolOutput {
	return o.ApplyT(func(v GetNodesNode) bool { return v.Configured }).(pulumi.BoolOutput)
}

func (o GetNodesNodeOutput) DiskSize() pulumi.IntOutput {
	return o.ApplyT(func(v GetNodesNode) int { return v.DiskSize }).(pulumi.IntOutput)
}

func (o GetNodesNodeOutput) ErlangVersion() pulumi.StringOutput {
	return o.ApplyT(func(v GetNodesNode) string { return v.ErlangVersion }).(pulumi.StringOutput)
}

func (o GetNodesNodeOutput) Hipe() pulumi.BoolOutput {
	return o.ApplyT(func(v GetNodesNode) bool { return v.Hipe }).(pulumi.BoolOutput)
}

func (o GetNodesNodeOutput) Hostname() pulumi.StringOutput {
	return o.ApplyT(func(v GetNodesNode) string { return v.Hostname }).(pulumi.StringOutput)
}

func (o GetNodesNodeOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetNodesNode) string { return v.Name }).(pulumi.StringOutput)
}

func (o GetNodesNodeOutput) RabbitmqVersion() pulumi.StringOutput {
	return o.ApplyT(func(v GetNodesNode) string { return v.RabbitmqVersion }).(pulumi.StringOutput)
}

func (o GetNodesNodeOutput) Running() pulumi.BoolOutput {
	return o.ApplyT(func(v GetNodesNode) bool { return v.Running }).(pulumi.BoolOutput)
}

type GetNodesNodeArrayOutput struct{ *pulumi.OutputState }

func (GetNodesNodeArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetNodesNode)(nil)).Elem()
}

func (o GetNodesNodeArrayOutput) ToGetNodesNodeArrayOutput() GetNodesNodeArrayOutput {
	return o
}

func (o GetNodesNodeArrayOutput) ToGetNodesNodeArrayOutputWithContext(ctx context.Context) GetNodesNodeArrayOutput {
	return o
}

func (o GetNodesNodeArrayOutput) Index(i pulumi.IntInput) GetNodesNodeOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) GetNodesNode {
		return vs[0].([]GetNodesNode)[vs[1].(int)]
	}).(GetNodesNodeOutput)
}

type GetPluginsCommunityPlugin struct {
	Description string `pulumi:"description"`
	Name        string `pulumi:"name"`
	Require     string `pulumi:"require"`
}

// GetPluginsCommunityPluginInput is an input type that accepts GetPluginsCommunityPluginArgs and GetPluginsCommunityPluginOutput values.
// You can construct a concrete instance of `GetPluginsCommunityPluginInput` via:
//
//	GetPluginsCommunityPluginArgs{...}
type GetPluginsCommunityPluginInput interface {
	pulumi.Input

	ToGetPluginsCommunityPluginOutput() GetPluginsCommunityPluginOutput
	ToGetPluginsCommunityPluginOutputWithContext(context.Context) GetPluginsCommunityPluginOutput
}

type GetPluginsCommunityPluginArgs struct {
	Description pulumi.StringInput `pulumi:"description"`
	Name        pulumi.StringInput `pulumi:"name"`
	Require     pulumi.StringInput `pulumi:"require"`
}

func (GetPluginsCommunityPluginArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetPluginsCommunityPlugin)(nil)).Elem()
}

func (i GetPluginsCommunityPluginArgs) ToGetPluginsCommunityPluginOutput() GetPluginsCommunityPluginOutput {
	return i.ToGetPluginsCommunityPluginOutputWithContext(context.Background())
}

func (i GetPluginsCommunityPluginArgs) ToGetPluginsCommunityPluginOutputWithContext(ctx context.Context) GetPluginsCommunityPluginOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetPluginsCommunityPluginOutput)
}

// GetPluginsCommunityPluginArrayInput is an input type that accepts GetPluginsCommunityPluginArray and GetPluginsCommunityPluginArrayOutput values.
// You can construct a concrete instance of `GetPluginsCommunityPluginArrayInput` via:
//
//	GetPluginsCommunityPluginArray{ GetPluginsCommunityPluginArgs{...} }
type GetPluginsCommunityPluginArrayInput interface {
	pulumi.Input

	ToGetPluginsCommunityPluginArrayOutput() GetPluginsCommunityPluginArrayOutput
	ToGetPluginsCommunityPluginArrayOutputWithContext(context.Context) GetPluginsCommunityPluginArrayOutput
}

type GetPluginsCommunityPluginArray []GetPluginsCommunityPluginInput

func (GetPluginsCommunityPluginArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetPluginsCommunityPlugin)(nil)).Elem()
}

func (i GetPluginsCommunityPluginArray) ToGetPluginsCommunityPluginArrayOutput() GetPluginsCommunityPluginArrayOutput {
	return i.ToGetPluginsCommunityPluginArrayOutputWithContext(context.Background())
}

func (i GetPluginsCommunityPluginArray) ToGetPluginsCommunityPluginArrayOutputWithContext(ctx context.Context) GetPluginsCommunityPluginArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetPluginsCommunityPluginArrayOutput)
}

type GetPluginsCommunityPluginOutput struct{ *pulumi.OutputState }

func (GetPluginsCommunityPluginOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetPluginsCommunityPlugin)(nil)).Elem()
}

func (o GetPluginsCommunityPluginOutput) ToGetPluginsCommunityPluginOutput() GetPluginsCommunityPluginOutput {
	return o
}

func (o GetPluginsCommunityPluginOutput) ToGetPluginsCommunityPluginOutputWithContext(ctx context.Context) GetPluginsCommunityPluginOutput {
	return o
}

func (o GetPluginsCommunityPluginOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v GetPluginsCommunityPlugin) string { return v.Description }).(pulumi.StringOutput)
}

func (o GetPluginsCommunityPluginOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetPluginsCommunityPlugin) string { return v.Name }).(pulumi.StringOutput)
}

func (o GetPluginsCommunityPluginOutput) Require() pulumi.StringOutput {
	return o.ApplyT(func(v GetPluginsCommunityPlugin) string { return v.Require }).(pulumi.StringOutput)
}

type GetPluginsCommunityPluginArrayOutput struct{ *pulumi.OutputState }

func (GetPluginsCommunityPluginArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetPluginsCommunityPlugin)(nil)).Elem()
}

func (o GetPluginsCommunityPluginArrayOutput) ToGetPluginsCommunityPluginArrayOutput() GetPluginsCommunityPluginArrayOutput {
	return o
}

func (o GetPluginsCommunityPluginArrayOutput) ToGetPluginsCommunityPluginArrayOutputWithContext(ctx context.Context) GetPluginsCommunityPluginArrayOutput {
	return o
}

func (o GetPluginsCommunityPluginArrayOutput) Index(i pulumi.IntInput) GetPluginsCommunityPluginOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) GetPluginsCommunityPlugin {
		return vs[0].([]GetPluginsCommunityPlugin)[vs[1].(int)]
	}).(GetPluginsCommunityPluginOutput)
}

type GetPluginsPlugin struct {
	Description string `pulumi:"description"`
	Enabled     bool   `pulumi:"enabled"`
	Name        string `pulumi:"name"`
	Version     string `pulumi:"version"`
}

// GetPluginsPluginInput is an input type that accepts GetPluginsPluginArgs and GetPluginsPluginOutput values.
// You can construct a concrete instance of `GetPluginsPluginInput` via:
//
//	GetPluginsPluginArgs{...}
type GetPluginsPluginInput interface {
	pulumi.Input

	ToGetPluginsPluginOutput() GetPluginsPluginOutput
	ToGetPluginsPluginOutputWithContext(context.Context) GetPluginsPluginOutput
}

type GetPluginsPluginArgs struct {
	Description pulumi.StringInput `pulumi:"description"`
	Enabled     pulumi.BoolInput   `pulumi:"enabled"`
	Name        pulumi.StringInput `pulumi:"name"`
	Version     pulumi.StringInput `pulumi:"version"`
}

func (GetPluginsPluginArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetPluginsPlugin)(nil)).Elem()
}

func (i GetPluginsPluginArgs) ToGetPluginsPluginOutput() GetPluginsPluginOutput {
	return i.ToGetPluginsPluginOutputWithContext(context.Background())
}

func (i GetPluginsPluginArgs) ToGetPluginsPluginOutputWithContext(ctx context.Context) GetPluginsPluginOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetPluginsPluginOutput)
}

// GetPluginsPluginArrayInput is an input type that accepts GetPluginsPluginArray and GetPluginsPluginArrayOutput values.
// You can construct a concrete instance of `GetPluginsPluginArrayInput` via:
//
//	GetPluginsPluginArray{ GetPluginsPluginArgs{...} }
type GetPluginsPluginArrayInput interface {
	pulumi.Input

	ToGetPluginsPluginArrayOutput() GetPluginsPluginArrayOutput
	ToGetPluginsPluginArrayOutputWithContext(context.Context) GetPluginsPluginArrayOutput
}

type GetPluginsPluginArray []GetPluginsPluginInput

func (GetPluginsPluginArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetPluginsPlugin)(nil)).Elem()
}

func (i GetPluginsPluginArray) ToGetPluginsPluginArrayOutput() GetPluginsPluginArrayOutput {
	return i.ToGetPluginsPluginArrayOutputWithContext(context.Background())
}

func (i GetPluginsPluginArray) ToGetPluginsPluginArrayOutputWithContext(ctx context.Context) GetPluginsPluginArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetPluginsPluginArrayOutput)
}

type GetPluginsPluginOutput struct{ *pulumi.OutputState }

func (GetPluginsPluginOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetPluginsPlugin)(nil)).Elem()
}

func (o GetPluginsPluginOutput) ToGetPluginsPluginOutput() GetPluginsPluginOutput {
	return o
}

func (o GetPluginsPluginOutput) ToGetPluginsPluginOutputWithContext(ctx context.Context) GetPluginsPluginOutput {
	return o
}

func (o GetPluginsPluginOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v GetPluginsPlugin) string { return v.Description }).(pulumi.StringOutput)
}

func (o GetPluginsPluginOutput) Enabled() pulumi.BoolOutput {
	return o.ApplyT(func(v GetPluginsPlugin) bool { return v.Enabled }).(pulumi.BoolOutput)
}

func (o GetPluginsPluginOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetPluginsPlugin) string { return v.Name }).(pulumi.StringOutput)
}

func (o GetPluginsPluginOutput) Version() pulumi.StringOutput {
	return o.ApplyT(func(v GetPluginsPlugin) string { return v.Version }).(pulumi.StringOutput)
}

type GetPluginsPluginArrayOutput struct{ *pulumi.OutputState }

func (GetPluginsPluginArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetPluginsPlugin)(nil)).Elem()
}

func (o GetPluginsPluginArrayOutput) ToGetPluginsPluginArrayOutput() GetPluginsPluginArrayOutput {
	return o
}

func (o GetPluginsPluginArrayOutput) ToGetPluginsPluginArrayOutputWithContext(ctx context.Context) GetPluginsPluginArrayOutput {
	return o
}

func (o GetPluginsPluginArrayOutput) Index(i pulumi.IntInput) GetPluginsPluginOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) GetPluginsPlugin {
		return vs[0].([]GetPluginsPlugin)[vs[1].(int)]
	}).(GetPluginsPluginOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ExtraDiskSizeNodeInput)(nil)).Elem(), ExtraDiskSizeNodeArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ExtraDiskSizeNodeArrayInput)(nil)).Elem(), ExtraDiskSizeNodeArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SecurityFirewallRuleInput)(nil)).Elem(), SecurityFirewallRuleArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SecurityFirewallRuleArrayInput)(nil)).Elem(), SecurityFirewallRuleArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*GetAccountInstanceInput)(nil)).Elem(), GetAccountInstanceArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*GetAccountInstanceArrayInput)(nil)).Elem(), GetAccountInstanceArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*GetAccountVpcsVpcInput)(nil)).Elem(), GetAccountVpcsVpcArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*GetAccountVpcsVpcArrayInput)(nil)).Elem(), GetAccountVpcsVpcArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*GetNodesNodeInput)(nil)).Elem(), GetNodesNodeArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*GetNodesNodeArrayInput)(nil)).Elem(), GetNodesNodeArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*GetPluginsCommunityPluginInput)(nil)).Elem(), GetPluginsCommunityPluginArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*GetPluginsCommunityPluginArrayInput)(nil)).Elem(), GetPluginsCommunityPluginArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*GetPluginsPluginInput)(nil)).Elem(), GetPluginsPluginArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*GetPluginsPluginArrayInput)(nil)).Elem(), GetPluginsPluginArray{})
	pulumi.RegisterOutputType(ExtraDiskSizeNodeOutput{})
	pulumi.RegisterOutputType(ExtraDiskSizeNodeArrayOutput{})
	pulumi.RegisterOutputType(SecurityFirewallRuleOutput{})
	pulumi.RegisterOutputType(SecurityFirewallRuleArrayOutput{})
	pulumi.RegisterOutputType(GetAccountInstanceOutput{})
	pulumi.RegisterOutputType(GetAccountInstanceArrayOutput{})
	pulumi.RegisterOutputType(GetAccountVpcsVpcOutput{})
	pulumi.RegisterOutputType(GetAccountVpcsVpcArrayOutput{})
	pulumi.RegisterOutputType(GetNodesNodeOutput{})
	pulumi.RegisterOutputType(GetNodesNodeArrayOutput{})
	pulumi.RegisterOutputType(GetPluginsCommunityPluginOutput{})
	pulumi.RegisterOutputType(GetPluginsCommunityPluginArrayOutput{})
	pulumi.RegisterOutputType(GetPluginsPluginOutput{})
	pulumi.RegisterOutputType(GetPluginsPluginArrayOutput{})
}
