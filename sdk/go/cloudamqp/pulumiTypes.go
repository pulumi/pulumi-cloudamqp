// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package cloudamqp

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type SecurityFirewallRule struct {
	// Description name of the rule. e.g. Default.
	Description *string `pulumi:"description"`
	// Source ip and netmask for the rule. (e.g. 10.56.72.0/24)
	Ip string `pulumi:"ip"`
	// Custom ports to be opened
	Ports []int `pulumi:"ports"`
	// Pre-defined service ports
	Services []string `pulumi:"services"`
}

// SecurityFirewallRuleInput is an input type that accepts SecurityFirewallRuleArgs and SecurityFirewallRuleOutput values.
// You can construct a concrete instance of `SecurityFirewallRuleInput` via:
//
//          SecurityFirewallRuleArgs{...}
type SecurityFirewallRuleInput interface {
	pulumi.Input

	ToSecurityFirewallRuleOutput() SecurityFirewallRuleOutput
	ToSecurityFirewallRuleOutputWithContext(context.Context) SecurityFirewallRuleOutput
}

type SecurityFirewallRuleArgs struct {
	// Description name of the rule. e.g. Default.
	Description pulumi.StringPtrInput `pulumi:"description"`
	// Source ip and netmask for the rule. (e.g. 10.56.72.0/24)
	Ip pulumi.StringInput `pulumi:"ip"`
	// Custom ports to be opened
	Ports pulumi.IntArrayInput `pulumi:"ports"`
	// Pre-defined service ports
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
//          SecurityFirewallRuleArray{ SecurityFirewallRuleArgs{...} }
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
func (o SecurityFirewallRuleOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SecurityFirewallRule) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// Source ip and netmask for the rule. (e.g. 10.56.72.0/24)
func (o SecurityFirewallRuleOutput) Ip() pulumi.StringOutput {
	return o.ApplyT(func(v SecurityFirewallRule) string { return v.Ip }).(pulumi.StringOutput)
}

// Custom ports to be opened
func (o SecurityFirewallRuleOutput) Ports() pulumi.IntArrayOutput {
	return o.ApplyT(func(v SecurityFirewallRule) []int { return v.Ports }).(pulumi.IntArrayOutput)
}

// Pre-defined service ports
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

type GetNodesNode struct {
	ErlangVersion   string `pulumi:"erlangVersion"`
	Hipe            bool   `pulumi:"hipe"`
	Hostname        string `pulumi:"hostname"`
	Name            string `pulumi:"name"`
	RabbitmqVersion string `pulumi:"rabbitmqVersion"`
	Running         bool   `pulumi:"running"`
}

// GetNodesNodeInput is an input type that accepts GetNodesNodeArgs and GetNodesNodeOutput values.
// You can construct a concrete instance of `GetNodesNodeInput` via:
//
//          GetNodesNodeArgs{...}
type GetNodesNodeInput interface {
	pulumi.Input

	ToGetNodesNodeOutput() GetNodesNodeOutput
	ToGetNodesNodeOutputWithContext(context.Context) GetNodesNodeOutput
}

type GetNodesNodeArgs struct {
	ErlangVersion   pulumi.StringInput `pulumi:"erlangVersion"`
	Hipe            pulumi.BoolInput   `pulumi:"hipe"`
	Hostname        pulumi.StringInput `pulumi:"hostname"`
	Name            pulumi.StringInput `pulumi:"name"`
	RabbitmqVersion pulumi.StringInput `pulumi:"rabbitmqVersion"`
	Running         pulumi.BoolInput   `pulumi:"running"`
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
//          GetNodesNodeArray{ GetNodesNodeArgs{...} }
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
	Description *string `pulumi:"description"`
	Name        *string `pulumi:"name"`
	Require     *string `pulumi:"require"`
}

// GetPluginsCommunityPluginInput is an input type that accepts GetPluginsCommunityPluginArgs and GetPluginsCommunityPluginOutput values.
// You can construct a concrete instance of `GetPluginsCommunityPluginInput` via:
//
//          GetPluginsCommunityPluginArgs{...}
type GetPluginsCommunityPluginInput interface {
	pulumi.Input

	ToGetPluginsCommunityPluginOutput() GetPluginsCommunityPluginOutput
	ToGetPluginsCommunityPluginOutputWithContext(context.Context) GetPluginsCommunityPluginOutput
}

type GetPluginsCommunityPluginArgs struct {
	Description pulumi.StringPtrInput `pulumi:"description"`
	Name        pulumi.StringPtrInput `pulumi:"name"`
	Require     pulumi.StringPtrInput `pulumi:"require"`
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
//          GetPluginsCommunityPluginArray{ GetPluginsCommunityPluginArgs{...} }
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

func (o GetPluginsCommunityPluginOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetPluginsCommunityPlugin) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o GetPluginsCommunityPluginOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetPluginsCommunityPlugin) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o GetPluginsCommunityPluginOutput) Require() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetPluginsCommunityPlugin) *string { return v.Require }).(pulumi.StringPtrOutput)
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
	Description *string `pulumi:"description"`
	Enabled     *bool   `pulumi:"enabled"`
	Name        *string `pulumi:"name"`
	Version     *string `pulumi:"version"`
}

// GetPluginsPluginInput is an input type that accepts GetPluginsPluginArgs and GetPluginsPluginOutput values.
// You can construct a concrete instance of `GetPluginsPluginInput` via:
//
//          GetPluginsPluginArgs{...}
type GetPluginsPluginInput interface {
	pulumi.Input

	ToGetPluginsPluginOutput() GetPluginsPluginOutput
	ToGetPluginsPluginOutputWithContext(context.Context) GetPluginsPluginOutput
}

type GetPluginsPluginArgs struct {
	Description pulumi.StringPtrInput `pulumi:"description"`
	Enabled     pulumi.BoolPtrInput   `pulumi:"enabled"`
	Name        pulumi.StringPtrInput `pulumi:"name"`
	Version     pulumi.StringPtrInput `pulumi:"version"`
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
//          GetPluginsPluginArray{ GetPluginsPluginArgs{...} }
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

func (o GetPluginsPluginOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetPluginsPlugin) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o GetPluginsPluginOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetPluginsPlugin) *bool { return v.Enabled }).(pulumi.BoolPtrOutput)
}

func (o GetPluginsPluginOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetPluginsPlugin) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o GetPluginsPluginOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetPluginsPlugin) *string { return v.Version }).(pulumi.StringPtrOutput)
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
	pulumi.RegisterOutputType(SecurityFirewallRuleOutput{})
	pulumi.RegisterOutputType(SecurityFirewallRuleArrayOutput{})
	pulumi.RegisterOutputType(GetNodesNodeOutput{})
	pulumi.RegisterOutputType(GetNodesNodeArrayOutput{})
	pulumi.RegisterOutputType(GetPluginsCommunityPluginOutput{})
	pulumi.RegisterOutputType(GetPluginsCommunityPluginArrayOutput{})
	pulumi.RegisterOutputType(GetPluginsPluginOutput{})
	pulumi.RegisterOutputType(GetPluginsPluginArrayOutput{})
}
