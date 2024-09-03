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

// This resource is a generic way to handle PrivateLink (AWS and Azure) and Private Service Connect (GCP).
// Communication between resources can be done just as they were living inside a VPC. CloudAMQP creates an Endpoint
// Service to connect the VPC and creating a new network interface to handle the communicate.
//
// If no existing VPC available when enable VPC connect, a new VPC will be created with subnet `10.52.72.0/24`.
//
// More information can be found at: [CloudAMQP VPC Connect](https://www.cloudamqp.com/docs/cloudamqp-vpc-connect.html)
//
// > **Note:** Enabling VPC Connect will automatically add a firewall rule.
//
// <details>
//
//	<summary>
//	   <b>
//	     <i>Default PrivateLink firewall rule [AWS, Azure]</i>
//	   </b>
//	 </summary>
//
// ## Example Usage
//
// <details>
//
//	<summary>
//	  <b>
//	    <i>Enable VPC Connect (PrivateLink) in AWS</i>
//	  </b>
//	</summary>
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
//			vpc, err := cloudamqp.NewVpc(ctx, "vpc", &cloudamqp.VpcArgs{
//				Name:   pulumi.String("Standalone VPC"),
//				Region: pulumi.String("amazon-web-services::us-west-1"),
//				Subnet: pulumi.String("10.56.72.0/24"),
//				Tags:   pulumi.StringArray{},
//			})
//			if err != nil {
//				return err
//			}
//			instance, err := cloudamqp.NewInstance(ctx, "instance", &cloudamqp.InstanceArgs{
//				Name:              pulumi.String("Instance 01"),
//				Plan:              pulumi.String("bunny-1"),
//				Region:            pulumi.String("amazon-web-services::us-west-1"),
//				Tags:              pulumi.StringArray{},
//				VpcId:             vpc.ID(),
//				KeepAssociatedVpc: pulumi.Bool(true),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = cloudamqp.NewVpcConnect(ctx, "vpc_connect", &cloudamqp.VpcConnectArgs{
//				InstanceId: instance.ID(),
//				Region:     instance.Region,
//				AllowedPrincipals: pulumi.StringArray{
//					pulumi.String("arn:aws:iam::aws-account-id:user/user-name"),
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
//
// </details>
//
// <details>
//
//	<summary>
//	  <b>
//	    <i>Enable VPC Connect (PrivateLink) in Azure</i>
//	  </b>
//	</summary>
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
//			_, err = cloudamqp.NewVpcConnect(ctx, "vpc_connect", &cloudamqp.VpcConnectArgs{
//				InstanceId: instance.ID(),
//				Region:     instance.Region,
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
//
// </details>
//
// <details>
//
//	<summary>
//	  <b>
//	    <i>Enable VPC Connect (Private Service Connect) in GCP</i>
//	  </b>
//	</summary>
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
//			vpc, err := cloudamqp.NewVpc(ctx, "vpc", &cloudamqp.VpcArgs{
//				Name:   pulumi.String("Standalone VPC"),
//				Region: pulumi.String("google-compute-engine::us-west1"),
//				Subnet: pulumi.String("10.56.72.0/24"),
//				Tags:   pulumi.StringArray{},
//			})
//			if err != nil {
//				return err
//			}
//			instance, err := cloudamqp.NewInstance(ctx, "instance", &cloudamqp.InstanceArgs{
//				Name:              pulumi.String("Instance 01"),
//				Plan:              pulumi.String("bunny-1"),
//				Region:            pulumi.String("google-compute-engine::us-west1"),
//				Tags:              pulumi.StringArray{},
//				VpcId:             vpc.ID(),
//				KeepAssociatedVpc: pulumi.Bool(true),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = cloudamqp.NewVpcConnect(ctx, "vpc_connect", &cloudamqp.VpcConnectArgs{
//				InstanceId: instance.ID(),
//				Region:     instance.Region,
//				AllowedProjects: pulumi.StringArray{
//					pulumi.String("some-project-123456"),
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
//
// </details>
//
// ### with additional firewall rules
//
// <details>
//
//	<summary>
//	  <b>
//	    <i>CloudAMQP instance in an existing VPC with managed firewall rules</i>
//	  </b>
//	</summary>
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
//			vpc, err := cloudamqp.NewVpc(ctx, "vpc", &cloudamqp.VpcArgs{
//				Name:   pulumi.String("Standalone VPC"),
//				Region: pulumi.String("amazon-web-services::us-west-1"),
//				Subnet: pulumi.String("10.56.72.0/24"),
//				Tags:   pulumi.StringArray{},
//			})
//			if err != nil {
//				return err
//			}
//			instance, err := cloudamqp.NewInstance(ctx, "instance", &cloudamqp.InstanceArgs{
//				Name:              pulumi.String("Instance 01"),
//				Plan:              pulumi.String("bunny-1"),
//				Region:            pulumi.String("amazon-web-services::us-west-1"),
//				Tags:              pulumi.StringArray{},
//				VpcId:             vpc.ID(),
//				KeepAssociatedVpc: pulumi.Bool(true),
//			})
//			if err != nil {
//				return err
//			}
//			vpcConnect, err := cloudamqp.NewVpcConnect(ctx, "vpc_connect", &cloudamqp.VpcConnectArgs{
//				InstanceId: instance.ID(),
//				AllowedPrincipals: pulumi.StringArray{
//					pulumi.String("arn:aws:iam::aws-account-id:user/user-name"),
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
//				vpcConnect,
//			}))
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// </details>
type VpcConnect struct {
	pulumi.CustomResourceState

	// Covering availability zones used when creating an endpoint from other VPC. [AWS]
	ActiveZones pulumi.StringArrayOutput `pulumi:"activeZones"`
	// List of allowed prinicpals used by AWS, see below table.
	AllowedPrincipals pulumi.StringArrayOutput `pulumi:"allowedPrincipals"`
	// List of allowed projects used by GCP, see below table.
	AllowedProjects pulumi.StringArrayOutput `pulumi:"allowedProjects"`
	// List of approved subscriptions used by Azure, see below table.
	ApprovedSubscriptions pulumi.StringArrayOutput `pulumi:"approvedSubscriptions"`
	// The CloudAMQP instance identifier.
	InstanceId pulumi.IntOutput `pulumi:"instanceId"`
	// The region where the CloudAMQP instance is hosted.
	Region pulumi.StringOutput `pulumi:"region"`
	// Service name (alias for Azure) of the PrivateLink.
	ServiceName pulumi.StringOutput `pulumi:"serviceName"`
	// Configurable sleep time (seconds) when enable Private Service Connect.
	// Default set to 10 seconds.
	Sleep pulumi.IntPtrOutput `pulumi:"sleep"`
	// Status of the Private Service Connect [enabled, pending, disabled]
	Status pulumi.StringOutput `pulumi:"status"`
	// Configurable timeout time (seconds) when enable Private Service Connect.
	// Default set to 1800 seconds.
	Timeout pulumi.IntPtrOutput `pulumi:"timeout"`
}

// NewVpcConnect registers a new resource with the given unique name, arguments, and options.
func NewVpcConnect(ctx *pulumi.Context,
	name string, args *VpcConnectArgs, opts ...pulumi.ResourceOption) (*VpcConnect, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.InstanceId == nil {
		return nil, errors.New("invalid value for required argument 'InstanceId'")
	}
	if args.Region == nil {
		return nil, errors.New("invalid value for required argument 'Region'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource VpcConnect
	err := ctx.RegisterResource("cloudamqp:index/vpcConnect:VpcConnect", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVpcConnect gets an existing VpcConnect resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVpcConnect(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VpcConnectState, opts ...pulumi.ResourceOption) (*VpcConnect, error) {
	var resource VpcConnect
	err := ctx.ReadResource("cloudamqp:index/vpcConnect:VpcConnect", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VpcConnect resources.
type vpcConnectState struct {
	// Covering availability zones used when creating an endpoint from other VPC. [AWS]
	ActiveZones []string `pulumi:"activeZones"`
	// List of allowed prinicpals used by AWS, see below table.
	AllowedPrincipals []string `pulumi:"allowedPrincipals"`
	// List of allowed projects used by GCP, see below table.
	AllowedProjects []string `pulumi:"allowedProjects"`
	// List of approved subscriptions used by Azure, see below table.
	ApprovedSubscriptions []string `pulumi:"approvedSubscriptions"`
	// The CloudAMQP instance identifier.
	InstanceId *int `pulumi:"instanceId"`
	// The region where the CloudAMQP instance is hosted.
	Region *string `pulumi:"region"`
	// Service name (alias for Azure) of the PrivateLink.
	ServiceName *string `pulumi:"serviceName"`
	// Configurable sleep time (seconds) when enable Private Service Connect.
	// Default set to 10 seconds.
	Sleep *int `pulumi:"sleep"`
	// Status of the Private Service Connect [enabled, pending, disabled]
	Status *string `pulumi:"status"`
	// Configurable timeout time (seconds) when enable Private Service Connect.
	// Default set to 1800 seconds.
	Timeout *int `pulumi:"timeout"`
}

type VpcConnectState struct {
	// Covering availability zones used when creating an endpoint from other VPC. [AWS]
	ActiveZones pulumi.StringArrayInput
	// List of allowed prinicpals used by AWS, see below table.
	AllowedPrincipals pulumi.StringArrayInput
	// List of allowed projects used by GCP, see below table.
	AllowedProjects pulumi.StringArrayInput
	// List of approved subscriptions used by Azure, see below table.
	ApprovedSubscriptions pulumi.StringArrayInput
	// The CloudAMQP instance identifier.
	InstanceId pulumi.IntPtrInput
	// The region where the CloudAMQP instance is hosted.
	Region pulumi.StringPtrInput
	// Service name (alias for Azure) of the PrivateLink.
	ServiceName pulumi.StringPtrInput
	// Configurable sleep time (seconds) when enable Private Service Connect.
	// Default set to 10 seconds.
	Sleep pulumi.IntPtrInput
	// Status of the Private Service Connect [enabled, pending, disabled]
	Status pulumi.StringPtrInput
	// Configurable timeout time (seconds) when enable Private Service Connect.
	// Default set to 1800 seconds.
	Timeout pulumi.IntPtrInput
}

func (VpcConnectState) ElementType() reflect.Type {
	return reflect.TypeOf((*vpcConnectState)(nil)).Elem()
}

type vpcConnectArgs struct {
	// List of allowed prinicpals used by AWS, see below table.
	AllowedPrincipals []string `pulumi:"allowedPrincipals"`
	// List of allowed projects used by GCP, see below table.
	AllowedProjects []string `pulumi:"allowedProjects"`
	// List of approved subscriptions used by Azure, see below table.
	ApprovedSubscriptions []string `pulumi:"approvedSubscriptions"`
	// The CloudAMQP instance identifier.
	InstanceId int `pulumi:"instanceId"`
	// The region where the CloudAMQP instance is hosted.
	Region string `pulumi:"region"`
	// Configurable sleep time (seconds) when enable Private Service Connect.
	// Default set to 10 seconds.
	Sleep *int `pulumi:"sleep"`
	// Configurable timeout time (seconds) when enable Private Service Connect.
	// Default set to 1800 seconds.
	Timeout *int `pulumi:"timeout"`
}

// The set of arguments for constructing a VpcConnect resource.
type VpcConnectArgs struct {
	// List of allowed prinicpals used by AWS, see below table.
	AllowedPrincipals pulumi.StringArrayInput
	// List of allowed projects used by GCP, see below table.
	AllowedProjects pulumi.StringArrayInput
	// List of approved subscriptions used by Azure, see below table.
	ApprovedSubscriptions pulumi.StringArrayInput
	// The CloudAMQP instance identifier.
	InstanceId pulumi.IntInput
	// The region where the CloudAMQP instance is hosted.
	Region pulumi.StringInput
	// Configurable sleep time (seconds) when enable Private Service Connect.
	// Default set to 10 seconds.
	Sleep pulumi.IntPtrInput
	// Configurable timeout time (seconds) when enable Private Service Connect.
	// Default set to 1800 seconds.
	Timeout pulumi.IntPtrInput
}

func (VpcConnectArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*vpcConnectArgs)(nil)).Elem()
}

type VpcConnectInput interface {
	pulumi.Input

	ToVpcConnectOutput() VpcConnectOutput
	ToVpcConnectOutputWithContext(ctx context.Context) VpcConnectOutput
}

func (*VpcConnect) ElementType() reflect.Type {
	return reflect.TypeOf((**VpcConnect)(nil)).Elem()
}

func (i *VpcConnect) ToVpcConnectOutput() VpcConnectOutput {
	return i.ToVpcConnectOutputWithContext(context.Background())
}

func (i *VpcConnect) ToVpcConnectOutputWithContext(ctx context.Context) VpcConnectOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpcConnectOutput)
}

// VpcConnectArrayInput is an input type that accepts VpcConnectArray and VpcConnectArrayOutput values.
// You can construct a concrete instance of `VpcConnectArrayInput` via:
//
//	VpcConnectArray{ VpcConnectArgs{...} }
type VpcConnectArrayInput interface {
	pulumi.Input

	ToVpcConnectArrayOutput() VpcConnectArrayOutput
	ToVpcConnectArrayOutputWithContext(context.Context) VpcConnectArrayOutput
}

type VpcConnectArray []VpcConnectInput

func (VpcConnectArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VpcConnect)(nil)).Elem()
}

func (i VpcConnectArray) ToVpcConnectArrayOutput() VpcConnectArrayOutput {
	return i.ToVpcConnectArrayOutputWithContext(context.Background())
}

func (i VpcConnectArray) ToVpcConnectArrayOutputWithContext(ctx context.Context) VpcConnectArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpcConnectArrayOutput)
}

// VpcConnectMapInput is an input type that accepts VpcConnectMap and VpcConnectMapOutput values.
// You can construct a concrete instance of `VpcConnectMapInput` via:
//
//	VpcConnectMap{ "key": VpcConnectArgs{...} }
type VpcConnectMapInput interface {
	pulumi.Input

	ToVpcConnectMapOutput() VpcConnectMapOutput
	ToVpcConnectMapOutputWithContext(context.Context) VpcConnectMapOutput
}

type VpcConnectMap map[string]VpcConnectInput

func (VpcConnectMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VpcConnect)(nil)).Elem()
}

func (i VpcConnectMap) ToVpcConnectMapOutput() VpcConnectMapOutput {
	return i.ToVpcConnectMapOutputWithContext(context.Background())
}

func (i VpcConnectMap) ToVpcConnectMapOutputWithContext(ctx context.Context) VpcConnectMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpcConnectMapOutput)
}

type VpcConnectOutput struct{ *pulumi.OutputState }

func (VpcConnectOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VpcConnect)(nil)).Elem()
}

func (o VpcConnectOutput) ToVpcConnectOutput() VpcConnectOutput {
	return o
}

func (o VpcConnectOutput) ToVpcConnectOutputWithContext(ctx context.Context) VpcConnectOutput {
	return o
}

// Covering availability zones used when creating an endpoint from other VPC. [AWS]
func (o VpcConnectOutput) ActiveZones() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *VpcConnect) pulumi.StringArrayOutput { return v.ActiveZones }).(pulumi.StringArrayOutput)
}

// List of allowed prinicpals used by AWS, see below table.
func (o VpcConnectOutput) AllowedPrincipals() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *VpcConnect) pulumi.StringArrayOutput { return v.AllowedPrincipals }).(pulumi.StringArrayOutput)
}

// List of allowed projects used by GCP, see below table.
func (o VpcConnectOutput) AllowedProjects() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *VpcConnect) pulumi.StringArrayOutput { return v.AllowedProjects }).(pulumi.StringArrayOutput)
}

// List of approved subscriptions used by Azure, see below table.
func (o VpcConnectOutput) ApprovedSubscriptions() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *VpcConnect) pulumi.StringArrayOutput { return v.ApprovedSubscriptions }).(pulumi.StringArrayOutput)
}

// The CloudAMQP instance identifier.
func (o VpcConnectOutput) InstanceId() pulumi.IntOutput {
	return o.ApplyT(func(v *VpcConnect) pulumi.IntOutput { return v.InstanceId }).(pulumi.IntOutput)
}

// The region where the CloudAMQP instance is hosted.
func (o VpcConnectOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *VpcConnect) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

// Service name (alias for Azure) of the PrivateLink.
func (o VpcConnectOutput) ServiceName() pulumi.StringOutput {
	return o.ApplyT(func(v *VpcConnect) pulumi.StringOutput { return v.ServiceName }).(pulumi.StringOutput)
}

// Configurable sleep time (seconds) when enable Private Service Connect.
// Default set to 10 seconds.
func (o VpcConnectOutput) Sleep() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *VpcConnect) pulumi.IntPtrOutput { return v.Sleep }).(pulumi.IntPtrOutput)
}

// Status of the Private Service Connect [enabled, pending, disabled]
func (o VpcConnectOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *VpcConnect) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// Configurable timeout time (seconds) when enable Private Service Connect.
// Default set to 1800 seconds.
func (o VpcConnectOutput) Timeout() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *VpcConnect) pulumi.IntPtrOutput { return v.Timeout }).(pulumi.IntPtrOutput)
}

type VpcConnectArrayOutput struct{ *pulumi.OutputState }

func (VpcConnectArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VpcConnect)(nil)).Elem()
}

func (o VpcConnectArrayOutput) ToVpcConnectArrayOutput() VpcConnectArrayOutput {
	return o
}

func (o VpcConnectArrayOutput) ToVpcConnectArrayOutputWithContext(ctx context.Context) VpcConnectArrayOutput {
	return o
}

func (o VpcConnectArrayOutput) Index(i pulumi.IntInput) VpcConnectOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *VpcConnect {
		return vs[0].([]*VpcConnect)[vs[1].(int)]
	}).(VpcConnectOutput)
}

type VpcConnectMapOutput struct{ *pulumi.OutputState }

func (VpcConnectMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VpcConnect)(nil)).Elem()
}

func (o VpcConnectMapOutput) ToVpcConnectMapOutput() VpcConnectMapOutput {
	return o
}

func (o VpcConnectMapOutput) ToVpcConnectMapOutputWithContext(ctx context.Context) VpcConnectMapOutput {
	return o
}

func (o VpcConnectMapOutput) MapIndex(k pulumi.StringInput) VpcConnectOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *VpcConnect {
		return vs[0].(map[string]*VpcConnect)[vs[1].(string)]
	}).(VpcConnectOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*VpcConnectInput)(nil)).Elem(), &VpcConnect{})
	pulumi.RegisterInputType(reflect.TypeOf((*VpcConnectArrayInput)(nil)).Elem(), VpcConnectArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*VpcConnectMapInput)(nil)).Elem(), VpcConnectMap{})
	pulumi.RegisterOutputType(VpcConnectOutput{})
	pulumi.RegisterOutputType(VpcConnectArrayOutput{})
	pulumi.RegisterOutputType(VpcConnectMapOutput{})
}
