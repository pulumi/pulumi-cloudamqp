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

// Enable PrivateLink for a CloudAMQP instance hosted in Azure. If no existing VPC available when
// enable PrivateLink, a new VPC will be created with subnet `10.52.72.0/24`.
//
// > **Note:** Enabling PrivateLink will automatically add firewall rules for the peered subnet.
//
// <details>
//
//	<summary>
//	   <i>Default PrivateLink firewall rule</i>
//	 </summary>
//
// ## Example Usage
//
// <details>
//
//	<summary>
//	  <b>
//	    <i>CloudAMQP instance without existing VPC</i>
//	  </b>
//	</summary>
//
// <!--Start PulumiCodeChooser -->
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
//			instance, err := cloudamqp.NewInstance(ctx, "instance", &cloudamqp.InstanceArgs{
//				Name:   pulumi.String("Instance 01"),
//				Plan:   pulumi.String("bunny-1"),
//				Region: pulumi.String("azure-arm::westus"),
//				Tags:   pulumi.StringArray{},
//			})
//			if err != nil {
//				return err
//			}
//			_, err = cloudamqp.NewPrivatelinkAzure(ctx, "privatelink", &cloudamqp.PrivatelinkAzureArgs{
//				InstanceId: instance.ID(),
//				ApprovedSubscriptions: pulumi.StringArray{
//					pulumi.String("XXXXXXXX-XXXX-XXXX-XXXX-XXXXXXXXXXXX"),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// <!--End PulumiCodeChooser -->
// </details>
//
// <details>
//
//	<summary>
//	  <b>
//	    <i>CloudAMQP instance in an existing VPC</i>
//	  </b>
//	</summary>
//
// <!--Start PulumiCodeChooser -->
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
//			vpc, err := cloudamqp.NewVpc(ctx, "vpc", &cloudamqp.VpcArgs{
//				Name:   pulumi.String("Standalone VPC"),
//				Region: pulumi.String("azure-arm::westus"),
//				Subnet: pulumi.String("10.56.72.0/24"),
//				Tags:   pulumi.StringArray{},
//			})
//			if err != nil {
//				return err
//			}
//			instance, err := cloudamqp.NewInstance(ctx, "instance", &cloudamqp.InstanceArgs{
//				Name:              pulumi.String("Instance 01"),
//				Plan:              pulumi.String("bunny-1"),
//				Region:            pulumi.String("azure-arm::westus"),
//				Tags:              pulumi.StringArray{},
//				VpcId:             vpc.ID(),
//				KeepAssociatedVpc: pulumi.Bool(true),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = cloudamqp.NewPrivatelinkAzure(ctx, "privatelink", &cloudamqp.PrivatelinkAzureArgs{
//				InstanceId: instance.ID(),
//				ApprovedSubscriptions: pulumi.StringArray{
//					pulumi.String("XXXXXXXX-XXXX-XXXX-XXXX-XXXXXXXXXXXX"),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// <!--End PulumiCodeChooser -->
//
// </details>
//
// ### With Additional Firewall Rules
//
// <details>
//
//	<summary>
//	  <b>
//	    <i>CloudAMQP instance in an existing VPC with managed firewall rules</i>
//	  </b>
//	</summary>
//
// <!--Start PulumiCodeChooser -->
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
//			vpc, err := cloudamqp.NewVpc(ctx, "vpc", &cloudamqp.VpcArgs{
//				Name:   pulumi.String("Standalone VPC"),
//				Region: pulumi.String("azure-arm::westus"),
//				Subnet: pulumi.String("10.56.72.0/24"),
//				Tags:   pulumi.StringArray{},
//			})
//			if err != nil {
//				return err
//			}
//			instance, err := cloudamqp.NewInstance(ctx, "instance", &cloudamqp.InstanceArgs{
//				Name:              pulumi.String("Instance 01"),
//				Plan:              pulumi.String("bunny-1"),
//				Region:            pulumi.String("azure-arm::westus"),
//				Tags:              pulumi.StringArray{},
//				VpcId:             vpc.ID(),
//				KeepAssociatedVpc: pulumi.Bool(true),
//			})
//			if err != nil {
//				return err
//			}
//			privatelink, err := cloudamqp.NewPrivatelinkAzure(ctx, "privatelink", &cloudamqp.PrivatelinkAzureArgs{
//				InstanceId: instance.ID(),
//				ApprovedSubscriptions: pulumi.StringArray{
//					pulumi.String("XXXXXXXX-XXXX-XXXX-XXXX-XXXXXXXXXXXX"),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			_, err = cloudamqp.NewSecurityFirewall(ctx, "firewall_settings", &cloudamqp.SecurityFirewallArgs{
//				InstanceId: instance.ID(),
//				Rules: cloudamqp.SecurityFirewallRuleArray{
//					&cloudamqp.SecurityFirewallRuleArgs{
//						Description: pulumi.String("Custom PrivateLink setup"),
//						Ip:          vpc.Subnet,
//						Ports:       pulumi.IntArray{},
//						Services: pulumi.StringArray{
//							pulumi.String("AMQP"),
//							pulumi.String("AMQPS"),
//							pulumi.String("HTTPS"),
//							pulumi.String("STREAM"),
//							pulumi.String("STREAM_SSL"),
//						},
//					},
//					&cloudamqp.SecurityFirewallRuleArgs{
//						Description: pulumi.String("MGMT interface"),
//						Ip:          pulumi.String("0.0.0.0/0"),
//						Ports:       pulumi.IntArray{},
//						Services: pulumi.StringArray{
//							pulumi.String("HTTPS"),
//						},
//					},
//				},
//			}, pulumi.DependsOn([]pulumi.Resource{
//				privatelink,
//			}))
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// <!--End PulumiCodeChooser -->
//
// </details>
//
// ## Depedency
//
// This resource depends on CloudAMQP instance identifier, `cloudamqp_instance.instance.id`.
//
// ## Create PrivateLink with additional firewall rules
//
// To create a PrivateLink configuration with additional firewall rules, it's required to chain the SecurityFirewall
// resource to avoid parallel conflicting resource calls. You can do this by making the firewall
// resource depend on the PrivateLink resource, `cloudamqp_privatelink_azure.privatelink`.
//
// Furthermore, since all firewall rules are overwritten, the otherwise automatically added rules for
// the PrivateLink also needs to be added.
//
// ## Import
//
// `cloudamqp_privatelink_aws` can be imported using CloudAMQP internal identifier.
//
// ```sh
// $ pulumi import cloudamqp:index/privatelinkAzure:PrivatelinkAzure privatelink <id>`
// ```
//
// The resource uses the same identifier as the CloudAMQP instance. To retrieve the identifier for an instance, either use [CloudAMQP customer API](https://docs.cloudamqp.com/#list-instances) or use the data source [`cloudamqp_account`](./data-sources/account.md).
type PrivatelinkAzure struct {
	pulumi.CustomResourceState

	// Approved subscriptions to access the endpoint service.
	// See format below.
	ApprovedSubscriptions pulumi.StringArrayOutput `pulumi:"approvedSubscriptions"`
	// The CloudAMQP instance identifier.
	InstanceId pulumi.IntOutput `pulumi:"instanceId"`
	// Name of the server having the PrivateLink enabled.
	ServerName pulumi.StringOutput `pulumi:"serverName"`
	// Service name (alias) of the PrivateLink, needed when creating the endpoint.
	ServiceName pulumi.StringOutput `pulumi:"serviceName"`
	// Configurable sleep time (seconds) when enable PrivateLink.
	// Default set to 10 seconds. *Available from v1.29.0*
	Sleep pulumi.IntPtrOutput `pulumi:"sleep"`
	// PrivateLink status [enable, pending, disable]
	Status pulumi.StringOutput `pulumi:"status"`
	// Configurable timeout time (seconds) when enable PrivateLink.
	// Default set to 1800 seconds. *Available from v1.29.0*
	//
	// Approved subscriptions format (GUID): <br>
	// `XXXXXXXX-XXXX-XXXX-XXXX-XXXXXXXXXXXX`
	Timeout pulumi.IntPtrOutput `pulumi:"timeout"`
}

// NewPrivatelinkAzure registers a new resource with the given unique name, arguments, and options.
func NewPrivatelinkAzure(ctx *pulumi.Context,
	name string, args *PrivatelinkAzureArgs, opts ...pulumi.ResourceOption) (*PrivatelinkAzure, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ApprovedSubscriptions == nil {
		return nil, errors.New("invalid value for required argument 'ApprovedSubscriptions'")
	}
	if args.InstanceId == nil {
		return nil, errors.New("invalid value for required argument 'InstanceId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource PrivatelinkAzure
	err := ctx.RegisterResource("cloudamqp:index/privatelinkAzure:PrivatelinkAzure", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPrivatelinkAzure gets an existing PrivatelinkAzure resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPrivatelinkAzure(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PrivatelinkAzureState, opts ...pulumi.ResourceOption) (*PrivatelinkAzure, error) {
	var resource PrivatelinkAzure
	err := ctx.ReadResource("cloudamqp:index/privatelinkAzure:PrivatelinkAzure", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PrivatelinkAzure resources.
type privatelinkAzureState struct {
	// Approved subscriptions to access the endpoint service.
	// See format below.
	ApprovedSubscriptions []string `pulumi:"approvedSubscriptions"`
	// The CloudAMQP instance identifier.
	InstanceId *int `pulumi:"instanceId"`
	// Name of the server having the PrivateLink enabled.
	ServerName *string `pulumi:"serverName"`
	// Service name (alias) of the PrivateLink, needed when creating the endpoint.
	ServiceName *string `pulumi:"serviceName"`
	// Configurable sleep time (seconds) when enable PrivateLink.
	// Default set to 10 seconds. *Available from v1.29.0*
	Sleep *int `pulumi:"sleep"`
	// PrivateLink status [enable, pending, disable]
	Status *string `pulumi:"status"`
	// Configurable timeout time (seconds) when enable PrivateLink.
	// Default set to 1800 seconds. *Available from v1.29.0*
	//
	// Approved subscriptions format (GUID): <br>
	// `XXXXXXXX-XXXX-XXXX-XXXX-XXXXXXXXXXXX`
	Timeout *int `pulumi:"timeout"`
}

type PrivatelinkAzureState struct {
	// Approved subscriptions to access the endpoint service.
	// See format below.
	ApprovedSubscriptions pulumi.StringArrayInput
	// The CloudAMQP instance identifier.
	InstanceId pulumi.IntPtrInput
	// Name of the server having the PrivateLink enabled.
	ServerName pulumi.StringPtrInput
	// Service name (alias) of the PrivateLink, needed when creating the endpoint.
	ServiceName pulumi.StringPtrInput
	// Configurable sleep time (seconds) when enable PrivateLink.
	// Default set to 10 seconds. *Available from v1.29.0*
	Sleep pulumi.IntPtrInput
	// PrivateLink status [enable, pending, disable]
	Status pulumi.StringPtrInput
	// Configurable timeout time (seconds) when enable PrivateLink.
	// Default set to 1800 seconds. *Available from v1.29.0*
	//
	// Approved subscriptions format (GUID): <br>
	// `XXXXXXXX-XXXX-XXXX-XXXX-XXXXXXXXXXXX`
	Timeout pulumi.IntPtrInput
}

func (PrivatelinkAzureState) ElementType() reflect.Type {
	return reflect.TypeOf((*privatelinkAzureState)(nil)).Elem()
}

type privatelinkAzureArgs struct {
	// Approved subscriptions to access the endpoint service.
	// See format below.
	ApprovedSubscriptions []string `pulumi:"approvedSubscriptions"`
	// The CloudAMQP instance identifier.
	InstanceId int `pulumi:"instanceId"`
	// Configurable sleep time (seconds) when enable PrivateLink.
	// Default set to 10 seconds. *Available from v1.29.0*
	Sleep *int `pulumi:"sleep"`
	// Configurable timeout time (seconds) when enable PrivateLink.
	// Default set to 1800 seconds. *Available from v1.29.0*
	//
	// Approved subscriptions format (GUID): <br>
	// `XXXXXXXX-XXXX-XXXX-XXXX-XXXXXXXXXXXX`
	Timeout *int `pulumi:"timeout"`
}

// The set of arguments for constructing a PrivatelinkAzure resource.
type PrivatelinkAzureArgs struct {
	// Approved subscriptions to access the endpoint service.
	// See format below.
	ApprovedSubscriptions pulumi.StringArrayInput
	// The CloudAMQP instance identifier.
	InstanceId pulumi.IntInput
	// Configurable sleep time (seconds) when enable PrivateLink.
	// Default set to 10 seconds. *Available from v1.29.0*
	Sleep pulumi.IntPtrInput
	// Configurable timeout time (seconds) when enable PrivateLink.
	// Default set to 1800 seconds. *Available from v1.29.0*
	//
	// Approved subscriptions format (GUID): <br>
	// `XXXXXXXX-XXXX-XXXX-XXXX-XXXXXXXXXXXX`
	Timeout pulumi.IntPtrInput
}

func (PrivatelinkAzureArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*privatelinkAzureArgs)(nil)).Elem()
}

type PrivatelinkAzureInput interface {
	pulumi.Input

	ToPrivatelinkAzureOutput() PrivatelinkAzureOutput
	ToPrivatelinkAzureOutputWithContext(ctx context.Context) PrivatelinkAzureOutput
}

func (*PrivatelinkAzure) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivatelinkAzure)(nil)).Elem()
}

func (i *PrivatelinkAzure) ToPrivatelinkAzureOutput() PrivatelinkAzureOutput {
	return i.ToPrivatelinkAzureOutputWithContext(context.Background())
}

func (i *PrivatelinkAzure) ToPrivatelinkAzureOutputWithContext(ctx context.Context) PrivatelinkAzureOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivatelinkAzureOutput)
}

// PrivatelinkAzureArrayInput is an input type that accepts PrivatelinkAzureArray and PrivatelinkAzureArrayOutput values.
// You can construct a concrete instance of `PrivatelinkAzureArrayInput` via:
//
//	PrivatelinkAzureArray{ PrivatelinkAzureArgs{...} }
type PrivatelinkAzureArrayInput interface {
	pulumi.Input

	ToPrivatelinkAzureArrayOutput() PrivatelinkAzureArrayOutput
	ToPrivatelinkAzureArrayOutputWithContext(context.Context) PrivatelinkAzureArrayOutput
}

type PrivatelinkAzureArray []PrivatelinkAzureInput

func (PrivatelinkAzureArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*PrivatelinkAzure)(nil)).Elem()
}

func (i PrivatelinkAzureArray) ToPrivatelinkAzureArrayOutput() PrivatelinkAzureArrayOutput {
	return i.ToPrivatelinkAzureArrayOutputWithContext(context.Background())
}

func (i PrivatelinkAzureArray) ToPrivatelinkAzureArrayOutputWithContext(ctx context.Context) PrivatelinkAzureArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivatelinkAzureArrayOutput)
}

// PrivatelinkAzureMapInput is an input type that accepts PrivatelinkAzureMap and PrivatelinkAzureMapOutput values.
// You can construct a concrete instance of `PrivatelinkAzureMapInput` via:
//
//	PrivatelinkAzureMap{ "key": PrivatelinkAzureArgs{...} }
type PrivatelinkAzureMapInput interface {
	pulumi.Input

	ToPrivatelinkAzureMapOutput() PrivatelinkAzureMapOutput
	ToPrivatelinkAzureMapOutputWithContext(context.Context) PrivatelinkAzureMapOutput
}

type PrivatelinkAzureMap map[string]PrivatelinkAzureInput

func (PrivatelinkAzureMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*PrivatelinkAzure)(nil)).Elem()
}

func (i PrivatelinkAzureMap) ToPrivatelinkAzureMapOutput() PrivatelinkAzureMapOutput {
	return i.ToPrivatelinkAzureMapOutputWithContext(context.Background())
}

func (i PrivatelinkAzureMap) ToPrivatelinkAzureMapOutputWithContext(ctx context.Context) PrivatelinkAzureMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivatelinkAzureMapOutput)
}

type PrivatelinkAzureOutput struct{ *pulumi.OutputState }

func (PrivatelinkAzureOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivatelinkAzure)(nil)).Elem()
}

func (o PrivatelinkAzureOutput) ToPrivatelinkAzureOutput() PrivatelinkAzureOutput {
	return o
}

func (o PrivatelinkAzureOutput) ToPrivatelinkAzureOutputWithContext(ctx context.Context) PrivatelinkAzureOutput {
	return o
}

// Approved subscriptions to access the endpoint service.
// See format below.
func (o PrivatelinkAzureOutput) ApprovedSubscriptions() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *PrivatelinkAzure) pulumi.StringArrayOutput { return v.ApprovedSubscriptions }).(pulumi.StringArrayOutput)
}

// The CloudAMQP instance identifier.
func (o PrivatelinkAzureOutput) InstanceId() pulumi.IntOutput {
	return o.ApplyT(func(v *PrivatelinkAzure) pulumi.IntOutput { return v.InstanceId }).(pulumi.IntOutput)
}

// Name of the server having the PrivateLink enabled.
func (o PrivatelinkAzureOutput) ServerName() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivatelinkAzure) pulumi.StringOutput { return v.ServerName }).(pulumi.StringOutput)
}

// Service name (alias) of the PrivateLink, needed when creating the endpoint.
func (o PrivatelinkAzureOutput) ServiceName() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivatelinkAzure) pulumi.StringOutput { return v.ServiceName }).(pulumi.StringOutput)
}

// Configurable sleep time (seconds) when enable PrivateLink.
// Default set to 10 seconds. *Available from v1.29.0*
func (o PrivatelinkAzureOutput) Sleep() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *PrivatelinkAzure) pulumi.IntPtrOutput { return v.Sleep }).(pulumi.IntPtrOutput)
}

// PrivateLink status [enable, pending, disable]
func (o PrivatelinkAzureOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivatelinkAzure) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// Configurable timeout time (seconds) when enable PrivateLink.
// Default set to 1800 seconds. *Available from v1.29.0*
//
// Approved subscriptions format (GUID): <br>
// `XXXXXXXX-XXXX-XXXX-XXXX-XXXXXXXXXXXX`
func (o PrivatelinkAzureOutput) Timeout() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *PrivatelinkAzure) pulumi.IntPtrOutput { return v.Timeout }).(pulumi.IntPtrOutput)
}

type PrivatelinkAzureArrayOutput struct{ *pulumi.OutputState }

func (PrivatelinkAzureArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*PrivatelinkAzure)(nil)).Elem()
}

func (o PrivatelinkAzureArrayOutput) ToPrivatelinkAzureArrayOutput() PrivatelinkAzureArrayOutput {
	return o
}

func (o PrivatelinkAzureArrayOutput) ToPrivatelinkAzureArrayOutputWithContext(ctx context.Context) PrivatelinkAzureArrayOutput {
	return o
}

func (o PrivatelinkAzureArrayOutput) Index(i pulumi.IntInput) PrivatelinkAzureOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *PrivatelinkAzure {
		return vs[0].([]*PrivatelinkAzure)[vs[1].(int)]
	}).(PrivatelinkAzureOutput)
}

type PrivatelinkAzureMapOutput struct{ *pulumi.OutputState }

func (PrivatelinkAzureMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*PrivatelinkAzure)(nil)).Elem()
}

func (o PrivatelinkAzureMapOutput) ToPrivatelinkAzureMapOutput() PrivatelinkAzureMapOutput {
	return o
}

func (o PrivatelinkAzureMapOutput) ToPrivatelinkAzureMapOutputWithContext(ctx context.Context) PrivatelinkAzureMapOutput {
	return o
}

func (o PrivatelinkAzureMapOutput) MapIndex(k pulumi.StringInput) PrivatelinkAzureOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *PrivatelinkAzure {
		return vs[0].(map[string]*PrivatelinkAzure)[vs[1].(string)]
	}).(PrivatelinkAzureOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*PrivatelinkAzureInput)(nil)).Elem(), &PrivatelinkAzure{})
	pulumi.RegisterInputType(reflect.TypeOf((*PrivatelinkAzureArrayInput)(nil)).Elem(), PrivatelinkAzureArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*PrivatelinkAzureMapInput)(nil)).Elem(), PrivatelinkAzureMap{})
	pulumi.RegisterOutputType(PrivatelinkAzureOutput{})
	pulumi.RegisterOutputType(PrivatelinkAzureArrayOutput{})
	pulumi.RegisterOutputType(PrivatelinkAzureMapOutput{})
}
