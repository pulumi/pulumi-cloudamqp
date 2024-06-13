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

// This resouce creates a VPC peering configuration for the CloudAMQP instance. The configuration will connect to another VPC network hosted on Google Cloud Platform (GCP). See the [GCP documentation](https://cloud.google.com/vpc/docs/using-vpc-peering) for more information on how to create the VPC peering configuration.
//
// > **Note:** Creating a VPC peering will automatically add firewall rules for the peered subnet.
//
// <details>
//
//	<summary>
//	   <i>Default VPC peering firewall rule</i>
//	 </summary>
//
// ## Example Usage
//
// <details>
//
//	<summary>
//	  <b>
//	    <i>VPC peering before v1.16.0</i>
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
//			// CloudAMQP instance
//			instance, err := cloudamqp.NewInstance(ctx, "instance", &cloudamqp.InstanceArgs{
//				Name:   pulumi.String("terraform-vpc-peering"),
//				Plan:   pulumi.String("bunny-1"),
//				Region: pulumi.String("google-compute-engine::europe-north1"),
//				Tags: pulumi.StringArray{
//					pulumi.String("terraform"),
//				},
//				VpcSubnet: pulumi.String("10.40.72.0/24"),
//			})
//			if err != nil {
//				return err
//			}
//			// VPC information
//			_ = instance.ID().ApplyT(func(id string) (cloudamqp.GetVpcGcpInfoResult, error) {
//				return cloudamqp.GetVpcGcpInfoOutput(ctx, cloudamqp.GetVpcGcpInfoOutputArgs{
//					InstanceId: id,
//				}, nil), nil
//			}).(cloudamqp.GetVpcGcpInfoResultOutput)
//			// VPC peering configuration
//			_, err = cloudamqp.NewVpcGcpPeering(ctx, "vpc_peering_request", &cloudamqp.VpcGcpPeeringArgs{
//				InstanceId:     instance.ID(),
//				PeerNetworkUri: pulumi.String("https://www.googleapis.com/compute/v1/projects/<PROJECT-NAME>/global/networks/<NETWORK-NAME>"),
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
//	    <i>VPC peering from v1.16.0 (Managed VPC)</i>
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
//			// Managed VPC resource
//			vpc, err := cloudamqp.NewVpc(ctx, "vpc", &cloudamqp.VpcArgs{
//				Name:   pulumi.String("<VPC name>"),
//				Region: pulumi.String("google-compute-engine::europe-north1"),
//				Subnet: pulumi.String("10.56.72.0/24"),
//				Tags:   pulumi.StringArray{},
//			})
//			if err != nil {
//				return err
//			}
//			// CloudAMQP instance
//			_, err = cloudamqp.NewInstance(ctx, "instance", &cloudamqp.InstanceArgs{
//				Name:   pulumi.String("terraform-vpc-peering"),
//				Plan:   pulumi.String("bunny-1"),
//				Region: pulumi.String("google-compute-engine::europe-north1"),
//				Tags: pulumi.StringArray{
//					pulumi.String("terraform"),
//				},
//				VpcId: vpc.ID(),
//			})
//			if err != nil {
//				return err
//			}
//			// VPC information
//			_, err = cloudamqp.GetVpcGcpInfo(ctx, &cloudamqp.GetVpcGcpInfoArgs{
//				VpcId: pulumi.StringRef(vpc.Info),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			// VPC peering configuration
//			_, err = cloudamqp.NewVpcGcpPeering(ctx, "vpc_peering_request", &cloudamqp.VpcGcpPeeringArgs{
//				VpcId:          vpc.ID(),
//				PeerNetworkUri: pulumi.String("https://www.googleapis.com/compute/v1/projects/<PROJECT-NAME>/global/networks/<NETWORK-NAME>"),
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
//	    <i>VPC peering from v1.28.0, waitOnPeeringStatus </i>
//	  </b>
//	</summary>
//
// Default peering request, no need to set `waitOnPeeringStatus`. It's default set to false and will not wait on peering status.
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
//			_, err := cloudamqp.NewVpcGcpPeering(ctx, "vpc_peering_request", &cloudamqp.VpcGcpPeeringArgs{
//				VpcId:          pulumi.Any(vpc.Id),
//				PeerNetworkUri: pulumi.String("https://www.googleapis.com/compute/v1/projects/<PROJECT-NAME>/global/networks/<NETWORK-NAME>"),
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
// Peering request and waiting for peering status.
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
//			_, err := cloudamqp.NewVpcGcpPeering(ctx, "vpc_peering_request", &cloudamqp.VpcGcpPeeringArgs{
//				VpcId:               pulumi.Any(vpc.Id),
//				WaitOnPeeringStatus: pulumi.Bool(true),
//				PeerNetworkUri:      pulumi.String("https://www.googleapis.com/compute/v1/projects/<PROJECT-NAME>/global/networks/<NETWORK-NAME>"),
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
// ### With Additional Firewall Rules
//
// <details>
//
//	<summary>
//	  <b>
//	    <i>VPC peering before v1.16.0</i>
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
//			// VPC peering configuration
//			vpcPeeringRequest, err := cloudamqp.NewVpcGcpPeering(ctx, "vpc_peering_request", &cloudamqp.VpcGcpPeeringArgs{
//				InstanceId:     pulumi.Any(instance.Id),
//				PeerNetworkUri: pulumi.Any(peerNetworkUri),
//			})
//			if err != nil {
//				return err
//			}
//			// Firewall rules
//			_, err = cloudamqp.NewSecurityFirewall(ctx, "firewall_settings", &cloudamqp.SecurityFirewallArgs{
//				InstanceId: pulumi.Any(instance.Id),
//				Rules: cloudamqp.SecurityFirewallRuleArray{
//					&cloudamqp.SecurityFirewallRuleArgs{
//						Ip: pulumi.Any(peerSubnet),
//						Ports: pulumi.IntArray{
//							pulumi.Int(15672),
//						},
//						Services: pulumi.StringArray{
//							pulumi.String("AMQP"),
//							pulumi.String("AMQPS"),
//							pulumi.String("STREAM"),
//							pulumi.String("STREAM_SSL"),
//						},
//						Description: pulumi.String("VPC peering for <NETWORK>"),
//					},
//					&cloudamqp.SecurityFirewallRuleArgs{
//						Ip: pulumi.String("192.168.0.0/24"),
//						Ports: pulumi.IntArray{
//							pulumi.Int(4567),
//							pulumi.Int(4568),
//						},
//						Services: pulumi.StringArray{
//							pulumi.String("AMQP"),
//							pulumi.String("AMQPS"),
//							pulumi.String("HTTPS"),
//						},
//					},
//				},
//			}, pulumi.DependsOn([]pulumi.Resource{
//				vpcPeeringRequest,
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
//
// <details>
//
//	<summary>
//	  <b>
//	    <i>VPC peering from v1.16.0 (Managed VPC)</i>
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
//			// VPC peering configuration
//			vpcPeeringRequest, err := cloudamqp.NewVpcGcpPeering(ctx, "vpc_peering_request", &cloudamqp.VpcGcpPeeringArgs{
//				VpcId:          pulumi.Any(vpc.Id),
//				PeerNetworkUri: pulumi.Any(peerNetworkUri),
//			})
//			if err != nil {
//				return err
//			}
//			// Firewall rules
//			_, err = cloudamqp.NewSecurityFirewall(ctx, "firewall_settings", &cloudamqp.SecurityFirewallArgs{
//				InstanceId: pulumi.Any(instance.Id),
//				Rules: cloudamqp.SecurityFirewallRuleArray{
//					&cloudamqp.SecurityFirewallRuleArgs{
//						Ip: pulumi.Any(peerSubnet),
//						Ports: pulumi.IntArray{
//							pulumi.Int(15672),
//						},
//						Services: pulumi.StringArray{
//							pulumi.String("AMQP"),
//							pulumi.String("AMQPS"),
//							pulumi.String("STREAM"),
//							pulumi.String("STREAM_SSL"),
//						},
//						Description: pulumi.String("VPC peering for <NETWORK>"),
//					},
//					&cloudamqp.SecurityFirewallRuleArgs{
//						Ip:    pulumi.String("0.0.0.0/0"),
//						Ports: pulumi.IntArray{},
//						Services: pulumi.StringArray{
//							pulumi.String("HTTPS"),
//						},
//						Description: pulumi.String("MGMT interface"),
//					},
//				},
//			}, pulumi.DependsOn([]pulumi.Resource{
//				vpcPeeringRequest,
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
//
// ## Depedency
//
// *Before v1.16.0*
// This resource depends on CloudAMQP instance identifier, `cloudamqp_instance.instance.id`.
//
// *From v1.16.0*
// This resource depends on CloudAMQP managed VPC identifier, `cloudamqp_vpc.vpc.id` or instance identifier, `cloudamqp_instance.instance.id`.
//
// ## Create VPC Peering with additional firewall rules
//
// To create a VPC peering configuration with additional firewall rules, it's required to chain the SecurityFirewall
// resource to avoid parallel conflicting resource calls. This is done by adding dependency from the firewall resource to the VPC peering resource.
//
// Furthermore, since all firewall rules are overwritten, the otherwise automatically added rules for the VPC peering also needs to be added.
//
// See example below.
//
// ## Import
//
// Not possible to import this resource.
type VpcGcpPeering struct {
	pulumi.CustomResourceState

	// VPC peering auto created routes
	AutoCreateRoutes pulumi.BoolOutput `pulumi:"autoCreateRoutes"`
	// The CloudAMQP instance identifier. *Deprecated from v1.16.0*
	InstanceId pulumi.IntPtrOutput `pulumi:"instanceId"`
	// Network uri of the VPC network to which you will peer with.
	PeerNetworkUri pulumi.StringOutput `pulumi:"peerNetworkUri"`
	// Configurable sleep time (seconds) between retries when requesting or reading
	// peering. Default set to 10 seconds. *Available from v1.29.0*
	Sleep pulumi.IntPtrOutput `pulumi:"sleep"`
	// VPC peering state
	State pulumi.StringOutput `pulumi:"state"`
	// VPC peering state details
	StateDetails pulumi.StringOutput `pulumi:"stateDetails"`
	// Configurable timeout time (seconds) before retries times out. Default set
	// to 1800 seconds. *Available from v1.29.0*
	Timeout pulumi.IntPtrOutput `pulumi:"timeout"`
	// The managed VPC identifier. *Available from v1.16.0*
	VpcId pulumi.StringPtrOutput `pulumi:"vpcId"`
	// Makes the resource wait until the peering is connected.
	// Default set to false. *Available from v1.28.0*
	WaitOnPeeringStatus pulumi.BoolPtrOutput `pulumi:"waitOnPeeringStatus"`
}

// NewVpcGcpPeering registers a new resource with the given unique name, arguments, and options.
func NewVpcGcpPeering(ctx *pulumi.Context,
	name string, args *VpcGcpPeeringArgs, opts ...pulumi.ResourceOption) (*VpcGcpPeering, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.PeerNetworkUri == nil {
		return nil, errors.New("invalid value for required argument 'PeerNetworkUri'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource VpcGcpPeering
	err := ctx.RegisterResource("cloudamqp:index/vpcGcpPeering:VpcGcpPeering", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVpcGcpPeering gets an existing VpcGcpPeering resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVpcGcpPeering(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VpcGcpPeeringState, opts ...pulumi.ResourceOption) (*VpcGcpPeering, error) {
	var resource VpcGcpPeering
	err := ctx.ReadResource("cloudamqp:index/vpcGcpPeering:VpcGcpPeering", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VpcGcpPeering resources.
type vpcGcpPeeringState struct {
	// VPC peering auto created routes
	AutoCreateRoutes *bool `pulumi:"autoCreateRoutes"`
	// The CloudAMQP instance identifier. *Deprecated from v1.16.0*
	InstanceId *int `pulumi:"instanceId"`
	// Network uri of the VPC network to which you will peer with.
	PeerNetworkUri *string `pulumi:"peerNetworkUri"`
	// Configurable sleep time (seconds) between retries when requesting or reading
	// peering. Default set to 10 seconds. *Available from v1.29.0*
	Sleep *int `pulumi:"sleep"`
	// VPC peering state
	State *string `pulumi:"state"`
	// VPC peering state details
	StateDetails *string `pulumi:"stateDetails"`
	// Configurable timeout time (seconds) before retries times out. Default set
	// to 1800 seconds. *Available from v1.29.0*
	Timeout *int `pulumi:"timeout"`
	// The managed VPC identifier. *Available from v1.16.0*
	VpcId *string `pulumi:"vpcId"`
	// Makes the resource wait until the peering is connected.
	// Default set to false. *Available from v1.28.0*
	WaitOnPeeringStatus *bool `pulumi:"waitOnPeeringStatus"`
}

type VpcGcpPeeringState struct {
	// VPC peering auto created routes
	AutoCreateRoutes pulumi.BoolPtrInput
	// The CloudAMQP instance identifier. *Deprecated from v1.16.0*
	InstanceId pulumi.IntPtrInput
	// Network uri of the VPC network to which you will peer with.
	PeerNetworkUri pulumi.StringPtrInput
	// Configurable sleep time (seconds) between retries when requesting or reading
	// peering. Default set to 10 seconds. *Available from v1.29.0*
	Sleep pulumi.IntPtrInput
	// VPC peering state
	State pulumi.StringPtrInput
	// VPC peering state details
	StateDetails pulumi.StringPtrInput
	// Configurable timeout time (seconds) before retries times out. Default set
	// to 1800 seconds. *Available from v1.29.0*
	Timeout pulumi.IntPtrInput
	// The managed VPC identifier. *Available from v1.16.0*
	VpcId pulumi.StringPtrInput
	// Makes the resource wait until the peering is connected.
	// Default set to false. *Available from v1.28.0*
	WaitOnPeeringStatus pulumi.BoolPtrInput
}

func (VpcGcpPeeringState) ElementType() reflect.Type {
	return reflect.TypeOf((*vpcGcpPeeringState)(nil)).Elem()
}

type vpcGcpPeeringArgs struct {
	// The CloudAMQP instance identifier. *Deprecated from v1.16.0*
	InstanceId *int `pulumi:"instanceId"`
	// Network uri of the VPC network to which you will peer with.
	PeerNetworkUri string `pulumi:"peerNetworkUri"`
	// Configurable sleep time (seconds) between retries when requesting or reading
	// peering. Default set to 10 seconds. *Available from v1.29.0*
	Sleep *int `pulumi:"sleep"`
	// Configurable timeout time (seconds) before retries times out. Default set
	// to 1800 seconds. *Available from v1.29.0*
	Timeout *int `pulumi:"timeout"`
	// The managed VPC identifier. *Available from v1.16.0*
	VpcId *string `pulumi:"vpcId"`
	// Makes the resource wait until the peering is connected.
	// Default set to false. *Available from v1.28.0*
	WaitOnPeeringStatus *bool `pulumi:"waitOnPeeringStatus"`
}

// The set of arguments for constructing a VpcGcpPeering resource.
type VpcGcpPeeringArgs struct {
	// The CloudAMQP instance identifier. *Deprecated from v1.16.0*
	InstanceId pulumi.IntPtrInput
	// Network uri of the VPC network to which you will peer with.
	PeerNetworkUri pulumi.StringInput
	// Configurable sleep time (seconds) between retries when requesting or reading
	// peering. Default set to 10 seconds. *Available from v1.29.0*
	Sleep pulumi.IntPtrInput
	// Configurable timeout time (seconds) before retries times out. Default set
	// to 1800 seconds. *Available from v1.29.0*
	Timeout pulumi.IntPtrInput
	// The managed VPC identifier. *Available from v1.16.0*
	VpcId pulumi.StringPtrInput
	// Makes the resource wait until the peering is connected.
	// Default set to false. *Available from v1.28.0*
	WaitOnPeeringStatus pulumi.BoolPtrInput
}

func (VpcGcpPeeringArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*vpcGcpPeeringArgs)(nil)).Elem()
}

type VpcGcpPeeringInput interface {
	pulumi.Input

	ToVpcGcpPeeringOutput() VpcGcpPeeringOutput
	ToVpcGcpPeeringOutputWithContext(ctx context.Context) VpcGcpPeeringOutput
}

func (*VpcGcpPeering) ElementType() reflect.Type {
	return reflect.TypeOf((**VpcGcpPeering)(nil)).Elem()
}

func (i *VpcGcpPeering) ToVpcGcpPeeringOutput() VpcGcpPeeringOutput {
	return i.ToVpcGcpPeeringOutputWithContext(context.Background())
}

func (i *VpcGcpPeering) ToVpcGcpPeeringOutputWithContext(ctx context.Context) VpcGcpPeeringOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpcGcpPeeringOutput)
}

// VpcGcpPeeringArrayInput is an input type that accepts VpcGcpPeeringArray and VpcGcpPeeringArrayOutput values.
// You can construct a concrete instance of `VpcGcpPeeringArrayInput` via:
//
//	VpcGcpPeeringArray{ VpcGcpPeeringArgs{...} }
type VpcGcpPeeringArrayInput interface {
	pulumi.Input

	ToVpcGcpPeeringArrayOutput() VpcGcpPeeringArrayOutput
	ToVpcGcpPeeringArrayOutputWithContext(context.Context) VpcGcpPeeringArrayOutput
}

type VpcGcpPeeringArray []VpcGcpPeeringInput

func (VpcGcpPeeringArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VpcGcpPeering)(nil)).Elem()
}

func (i VpcGcpPeeringArray) ToVpcGcpPeeringArrayOutput() VpcGcpPeeringArrayOutput {
	return i.ToVpcGcpPeeringArrayOutputWithContext(context.Background())
}

func (i VpcGcpPeeringArray) ToVpcGcpPeeringArrayOutputWithContext(ctx context.Context) VpcGcpPeeringArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpcGcpPeeringArrayOutput)
}

// VpcGcpPeeringMapInput is an input type that accepts VpcGcpPeeringMap and VpcGcpPeeringMapOutput values.
// You can construct a concrete instance of `VpcGcpPeeringMapInput` via:
//
//	VpcGcpPeeringMap{ "key": VpcGcpPeeringArgs{...} }
type VpcGcpPeeringMapInput interface {
	pulumi.Input

	ToVpcGcpPeeringMapOutput() VpcGcpPeeringMapOutput
	ToVpcGcpPeeringMapOutputWithContext(context.Context) VpcGcpPeeringMapOutput
}

type VpcGcpPeeringMap map[string]VpcGcpPeeringInput

func (VpcGcpPeeringMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VpcGcpPeering)(nil)).Elem()
}

func (i VpcGcpPeeringMap) ToVpcGcpPeeringMapOutput() VpcGcpPeeringMapOutput {
	return i.ToVpcGcpPeeringMapOutputWithContext(context.Background())
}

func (i VpcGcpPeeringMap) ToVpcGcpPeeringMapOutputWithContext(ctx context.Context) VpcGcpPeeringMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpcGcpPeeringMapOutput)
}

type VpcGcpPeeringOutput struct{ *pulumi.OutputState }

func (VpcGcpPeeringOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VpcGcpPeering)(nil)).Elem()
}

func (o VpcGcpPeeringOutput) ToVpcGcpPeeringOutput() VpcGcpPeeringOutput {
	return o
}

func (o VpcGcpPeeringOutput) ToVpcGcpPeeringOutputWithContext(ctx context.Context) VpcGcpPeeringOutput {
	return o
}

// VPC peering auto created routes
func (o VpcGcpPeeringOutput) AutoCreateRoutes() pulumi.BoolOutput {
	return o.ApplyT(func(v *VpcGcpPeering) pulumi.BoolOutput { return v.AutoCreateRoutes }).(pulumi.BoolOutput)
}

// The CloudAMQP instance identifier. *Deprecated from v1.16.0*
func (o VpcGcpPeeringOutput) InstanceId() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *VpcGcpPeering) pulumi.IntPtrOutput { return v.InstanceId }).(pulumi.IntPtrOutput)
}

// Network uri of the VPC network to which you will peer with.
func (o VpcGcpPeeringOutput) PeerNetworkUri() pulumi.StringOutput {
	return o.ApplyT(func(v *VpcGcpPeering) pulumi.StringOutput { return v.PeerNetworkUri }).(pulumi.StringOutput)
}

// Configurable sleep time (seconds) between retries when requesting or reading
// peering. Default set to 10 seconds. *Available from v1.29.0*
func (o VpcGcpPeeringOutput) Sleep() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *VpcGcpPeering) pulumi.IntPtrOutput { return v.Sleep }).(pulumi.IntPtrOutput)
}

// VPC peering state
func (o VpcGcpPeeringOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v *VpcGcpPeering) pulumi.StringOutput { return v.State }).(pulumi.StringOutput)
}

// VPC peering state details
func (o VpcGcpPeeringOutput) StateDetails() pulumi.StringOutput {
	return o.ApplyT(func(v *VpcGcpPeering) pulumi.StringOutput { return v.StateDetails }).(pulumi.StringOutput)
}

// Configurable timeout time (seconds) before retries times out. Default set
// to 1800 seconds. *Available from v1.29.0*
func (o VpcGcpPeeringOutput) Timeout() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *VpcGcpPeering) pulumi.IntPtrOutput { return v.Timeout }).(pulumi.IntPtrOutput)
}

// The managed VPC identifier. *Available from v1.16.0*
func (o VpcGcpPeeringOutput) VpcId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VpcGcpPeering) pulumi.StringPtrOutput { return v.VpcId }).(pulumi.StringPtrOutput)
}

// Makes the resource wait until the peering is connected.
// Default set to false. *Available from v1.28.0*
func (o VpcGcpPeeringOutput) WaitOnPeeringStatus() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *VpcGcpPeering) pulumi.BoolPtrOutput { return v.WaitOnPeeringStatus }).(pulumi.BoolPtrOutput)
}

type VpcGcpPeeringArrayOutput struct{ *pulumi.OutputState }

func (VpcGcpPeeringArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VpcGcpPeering)(nil)).Elem()
}

func (o VpcGcpPeeringArrayOutput) ToVpcGcpPeeringArrayOutput() VpcGcpPeeringArrayOutput {
	return o
}

func (o VpcGcpPeeringArrayOutput) ToVpcGcpPeeringArrayOutputWithContext(ctx context.Context) VpcGcpPeeringArrayOutput {
	return o
}

func (o VpcGcpPeeringArrayOutput) Index(i pulumi.IntInput) VpcGcpPeeringOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *VpcGcpPeering {
		return vs[0].([]*VpcGcpPeering)[vs[1].(int)]
	}).(VpcGcpPeeringOutput)
}

type VpcGcpPeeringMapOutput struct{ *pulumi.OutputState }

func (VpcGcpPeeringMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VpcGcpPeering)(nil)).Elem()
}

func (o VpcGcpPeeringMapOutput) ToVpcGcpPeeringMapOutput() VpcGcpPeeringMapOutput {
	return o
}

func (o VpcGcpPeeringMapOutput) ToVpcGcpPeeringMapOutputWithContext(ctx context.Context) VpcGcpPeeringMapOutput {
	return o
}

func (o VpcGcpPeeringMapOutput) MapIndex(k pulumi.StringInput) VpcGcpPeeringOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *VpcGcpPeering {
		return vs[0].(map[string]*VpcGcpPeering)[vs[1].(string)]
	}).(VpcGcpPeeringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*VpcGcpPeeringInput)(nil)).Elem(), &VpcGcpPeering{})
	pulumi.RegisterInputType(reflect.TypeOf((*VpcGcpPeeringArrayInput)(nil)).Elem(), VpcGcpPeeringArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*VpcGcpPeeringMapInput)(nil)).Elem(), VpcGcpPeeringMap{})
	pulumi.RegisterOutputType(VpcGcpPeeringOutput{})
	pulumi.RegisterOutputType(VpcGcpPeeringArrayOutput{})
	pulumi.RegisterOutputType(VpcGcpPeeringMapOutput{})
}
