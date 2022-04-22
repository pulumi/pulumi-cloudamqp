// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cloudamqp

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type IntegrationMetric struct {
	pulumi.CustomResourceState

	// AWS access key identifier. (Cloudwatch)
	AccessKeyId pulumi.StringPtrOutput `pulumi:"accessKeyId"`
	// The API key for the integration service. (Librato)
	ApiKey pulumi.StringPtrOutput `pulumi:"apiKey"`
	// The client email. (Stackdriver)
	ClientEmail pulumi.StringPtrOutput `pulumi:"clientEmail"`
	// The email address registred for the integration service. (Librato)
	Email pulumi.StringPtrOutput `pulumi:"email"`
	// Instance identifier
	InstanceId pulumi.IntOutput `pulumi:"instanceId"`
	// The license key registred for the integration service. (New Relic)
	LicenseKey pulumi.StringPtrOutput `pulumi:"licenseKey"`
	// The name of metrics integration
	Name pulumi.StringOutput `pulumi:"name"`
	// The private key. (Stackdriver)
	PrivateKey pulumi.StringPtrOutput `pulumi:"privateKey"`
	// Project ID. (Stackdriver)
	ProjectId pulumi.StringPtrOutput `pulumi:"projectId"`
	// (optional) allowlist using regular expression
	QueueAllowlist pulumi.StringPtrOutput `pulumi:"queueAllowlist"`
	// **Deprecated**
	//
	// Deprecated: use queue_allowlist instead
	QueueWhitelist pulumi.StringPtrOutput `pulumi:"queueWhitelist"`
	// AWS region for Cloudwatch and [US/EU] for Data dog/New relic. (Cloudwatch, Data Dog, New Relic)
	Region pulumi.StringPtrOutput `pulumi:"region"`
	// AWS secret key. (Cloudwatch)
	SecretAccessKey pulumi.StringPtrOutput `pulumi:"secretAccessKey"`
	// (optional) tags. E.g. env=prod,region=europe
	Tags pulumi.StringPtrOutput `pulumi:"tags"`
	// (optional) allowlist using regular expression
	VhostAllowlist pulumi.StringPtrOutput `pulumi:"vhostAllowlist"`
	// **Deprecated**
	//
	// Deprecated: use vhost_allowlist instead
	VhostWhitelist pulumi.StringPtrOutput `pulumi:"vhostWhitelist"`
}

// NewIntegrationMetric registers a new resource with the given unique name, arguments, and options.
func NewIntegrationMetric(ctx *pulumi.Context,
	name string, args *IntegrationMetricArgs, opts ...pulumi.ResourceOption) (*IntegrationMetric, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.InstanceId == nil {
		return nil, errors.New("invalid value for required argument 'InstanceId'")
	}
	var resource IntegrationMetric
	err := ctx.RegisterResource("cloudamqp:index/integrationMetric:IntegrationMetric", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIntegrationMetric gets an existing IntegrationMetric resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIntegrationMetric(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IntegrationMetricState, opts ...pulumi.ResourceOption) (*IntegrationMetric, error) {
	var resource IntegrationMetric
	err := ctx.ReadResource("cloudamqp:index/integrationMetric:IntegrationMetric", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering IntegrationMetric resources.
type integrationMetricState struct {
	// AWS access key identifier. (Cloudwatch)
	AccessKeyId *string `pulumi:"accessKeyId"`
	// The API key for the integration service. (Librato)
	ApiKey *string `pulumi:"apiKey"`
	// The client email. (Stackdriver)
	ClientEmail *string `pulumi:"clientEmail"`
	// The email address registred for the integration service. (Librato)
	Email *string `pulumi:"email"`
	// Instance identifier
	InstanceId *int `pulumi:"instanceId"`
	// The license key registred for the integration service. (New Relic)
	LicenseKey *string `pulumi:"licenseKey"`
	// The name of metrics integration
	Name *string `pulumi:"name"`
	// The private key. (Stackdriver)
	PrivateKey *string `pulumi:"privateKey"`
	// Project ID. (Stackdriver)
	ProjectId *string `pulumi:"projectId"`
	// (optional) allowlist using regular expression
	QueueAllowlist *string `pulumi:"queueAllowlist"`
	// **Deprecated**
	//
	// Deprecated: use queue_allowlist instead
	QueueWhitelist *string `pulumi:"queueWhitelist"`
	// AWS region for Cloudwatch and [US/EU] for Data dog/New relic. (Cloudwatch, Data Dog, New Relic)
	Region *string `pulumi:"region"`
	// AWS secret key. (Cloudwatch)
	SecretAccessKey *string `pulumi:"secretAccessKey"`
	// (optional) tags. E.g. env=prod,region=europe
	Tags *string `pulumi:"tags"`
	// (optional) allowlist using regular expression
	VhostAllowlist *string `pulumi:"vhostAllowlist"`
	// **Deprecated**
	//
	// Deprecated: use vhost_allowlist instead
	VhostWhitelist *string `pulumi:"vhostWhitelist"`
}

type IntegrationMetricState struct {
	// AWS access key identifier. (Cloudwatch)
	AccessKeyId pulumi.StringPtrInput
	// The API key for the integration service. (Librato)
	ApiKey pulumi.StringPtrInput
	// The client email. (Stackdriver)
	ClientEmail pulumi.StringPtrInput
	// The email address registred for the integration service. (Librato)
	Email pulumi.StringPtrInput
	// Instance identifier
	InstanceId pulumi.IntPtrInput
	// The license key registred for the integration service. (New Relic)
	LicenseKey pulumi.StringPtrInput
	// The name of metrics integration
	Name pulumi.StringPtrInput
	// The private key. (Stackdriver)
	PrivateKey pulumi.StringPtrInput
	// Project ID. (Stackdriver)
	ProjectId pulumi.StringPtrInput
	// (optional) allowlist using regular expression
	QueueAllowlist pulumi.StringPtrInput
	// **Deprecated**
	//
	// Deprecated: use queue_allowlist instead
	QueueWhitelist pulumi.StringPtrInput
	// AWS region for Cloudwatch and [US/EU] for Data dog/New relic. (Cloudwatch, Data Dog, New Relic)
	Region pulumi.StringPtrInput
	// AWS secret key. (Cloudwatch)
	SecretAccessKey pulumi.StringPtrInput
	// (optional) tags. E.g. env=prod,region=europe
	Tags pulumi.StringPtrInput
	// (optional) allowlist using regular expression
	VhostAllowlist pulumi.StringPtrInput
	// **Deprecated**
	//
	// Deprecated: use vhost_allowlist instead
	VhostWhitelist pulumi.StringPtrInput
}

func (IntegrationMetricState) ElementType() reflect.Type {
	return reflect.TypeOf((*integrationMetricState)(nil)).Elem()
}

type integrationMetricArgs struct {
	// AWS access key identifier. (Cloudwatch)
	AccessKeyId *string `pulumi:"accessKeyId"`
	// The API key for the integration service. (Librato)
	ApiKey *string `pulumi:"apiKey"`
	// The client email. (Stackdriver)
	ClientEmail *string `pulumi:"clientEmail"`
	// The email address registred for the integration service. (Librato)
	Email *string `pulumi:"email"`
	// Instance identifier
	InstanceId int `pulumi:"instanceId"`
	// The license key registred for the integration service. (New Relic)
	LicenseKey *string `pulumi:"licenseKey"`
	// The name of metrics integration
	Name *string `pulumi:"name"`
	// The private key. (Stackdriver)
	PrivateKey *string `pulumi:"privateKey"`
	// Project ID. (Stackdriver)
	ProjectId *string `pulumi:"projectId"`
	// (optional) allowlist using regular expression
	QueueAllowlist *string `pulumi:"queueAllowlist"`
	// **Deprecated**
	//
	// Deprecated: use queue_allowlist instead
	QueueWhitelist *string `pulumi:"queueWhitelist"`
	// AWS region for Cloudwatch and [US/EU] for Data dog/New relic. (Cloudwatch, Data Dog, New Relic)
	Region *string `pulumi:"region"`
	// AWS secret key. (Cloudwatch)
	SecretAccessKey *string `pulumi:"secretAccessKey"`
	// (optional) tags. E.g. env=prod,region=europe
	Tags *string `pulumi:"tags"`
	// (optional) allowlist using regular expression
	VhostAllowlist *string `pulumi:"vhostAllowlist"`
	// **Deprecated**
	//
	// Deprecated: use vhost_allowlist instead
	VhostWhitelist *string `pulumi:"vhostWhitelist"`
}

// The set of arguments for constructing a IntegrationMetric resource.
type IntegrationMetricArgs struct {
	// AWS access key identifier. (Cloudwatch)
	AccessKeyId pulumi.StringPtrInput
	// The API key for the integration service. (Librato)
	ApiKey pulumi.StringPtrInput
	// The client email. (Stackdriver)
	ClientEmail pulumi.StringPtrInput
	// The email address registred for the integration service. (Librato)
	Email pulumi.StringPtrInput
	// Instance identifier
	InstanceId pulumi.IntInput
	// The license key registred for the integration service. (New Relic)
	LicenseKey pulumi.StringPtrInput
	// The name of metrics integration
	Name pulumi.StringPtrInput
	// The private key. (Stackdriver)
	PrivateKey pulumi.StringPtrInput
	// Project ID. (Stackdriver)
	ProjectId pulumi.StringPtrInput
	// (optional) allowlist using regular expression
	QueueAllowlist pulumi.StringPtrInput
	// **Deprecated**
	//
	// Deprecated: use queue_allowlist instead
	QueueWhitelist pulumi.StringPtrInput
	// AWS region for Cloudwatch and [US/EU] for Data dog/New relic. (Cloudwatch, Data Dog, New Relic)
	Region pulumi.StringPtrInput
	// AWS secret key. (Cloudwatch)
	SecretAccessKey pulumi.StringPtrInput
	// (optional) tags. E.g. env=prod,region=europe
	Tags pulumi.StringPtrInput
	// (optional) allowlist using regular expression
	VhostAllowlist pulumi.StringPtrInput
	// **Deprecated**
	//
	// Deprecated: use vhost_allowlist instead
	VhostWhitelist pulumi.StringPtrInput
}

func (IntegrationMetricArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*integrationMetricArgs)(nil)).Elem()
}

type IntegrationMetricInput interface {
	pulumi.Input

	ToIntegrationMetricOutput() IntegrationMetricOutput
	ToIntegrationMetricOutputWithContext(ctx context.Context) IntegrationMetricOutput
}

func (*IntegrationMetric) ElementType() reflect.Type {
	return reflect.TypeOf((**IntegrationMetric)(nil)).Elem()
}

func (i *IntegrationMetric) ToIntegrationMetricOutput() IntegrationMetricOutput {
	return i.ToIntegrationMetricOutputWithContext(context.Background())
}

func (i *IntegrationMetric) ToIntegrationMetricOutputWithContext(ctx context.Context) IntegrationMetricOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IntegrationMetricOutput)
}

// IntegrationMetricArrayInput is an input type that accepts IntegrationMetricArray and IntegrationMetricArrayOutput values.
// You can construct a concrete instance of `IntegrationMetricArrayInput` via:
//
//          IntegrationMetricArray{ IntegrationMetricArgs{...} }
type IntegrationMetricArrayInput interface {
	pulumi.Input

	ToIntegrationMetricArrayOutput() IntegrationMetricArrayOutput
	ToIntegrationMetricArrayOutputWithContext(context.Context) IntegrationMetricArrayOutput
}

type IntegrationMetricArray []IntegrationMetricInput

func (IntegrationMetricArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*IntegrationMetric)(nil)).Elem()
}

func (i IntegrationMetricArray) ToIntegrationMetricArrayOutput() IntegrationMetricArrayOutput {
	return i.ToIntegrationMetricArrayOutputWithContext(context.Background())
}

func (i IntegrationMetricArray) ToIntegrationMetricArrayOutputWithContext(ctx context.Context) IntegrationMetricArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IntegrationMetricArrayOutput)
}

// IntegrationMetricMapInput is an input type that accepts IntegrationMetricMap and IntegrationMetricMapOutput values.
// You can construct a concrete instance of `IntegrationMetricMapInput` via:
//
//          IntegrationMetricMap{ "key": IntegrationMetricArgs{...} }
type IntegrationMetricMapInput interface {
	pulumi.Input

	ToIntegrationMetricMapOutput() IntegrationMetricMapOutput
	ToIntegrationMetricMapOutputWithContext(context.Context) IntegrationMetricMapOutput
}

type IntegrationMetricMap map[string]IntegrationMetricInput

func (IntegrationMetricMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*IntegrationMetric)(nil)).Elem()
}

func (i IntegrationMetricMap) ToIntegrationMetricMapOutput() IntegrationMetricMapOutput {
	return i.ToIntegrationMetricMapOutputWithContext(context.Background())
}

func (i IntegrationMetricMap) ToIntegrationMetricMapOutputWithContext(ctx context.Context) IntegrationMetricMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IntegrationMetricMapOutput)
}

type IntegrationMetricOutput struct{ *pulumi.OutputState }

func (IntegrationMetricOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IntegrationMetric)(nil)).Elem()
}

func (o IntegrationMetricOutput) ToIntegrationMetricOutput() IntegrationMetricOutput {
	return o
}

func (o IntegrationMetricOutput) ToIntegrationMetricOutputWithContext(ctx context.Context) IntegrationMetricOutput {
	return o
}

type IntegrationMetricArrayOutput struct{ *pulumi.OutputState }

func (IntegrationMetricArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*IntegrationMetric)(nil)).Elem()
}

func (o IntegrationMetricArrayOutput) ToIntegrationMetricArrayOutput() IntegrationMetricArrayOutput {
	return o
}

func (o IntegrationMetricArrayOutput) ToIntegrationMetricArrayOutputWithContext(ctx context.Context) IntegrationMetricArrayOutput {
	return o
}

func (o IntegrationMetricArrayOutput) Index(i pulumi.IntInput) IntegrationMetricOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *IntegrationMetric {
		return vs[0].([]*IntegrationMetric)[vs[1].(int)]
	}).(IntegrationMetricOutput)
}

type IntegrationMetricMapOutput struct{ *pulumi.OutputState }

func (IntegrationMetricMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*IntegrationMetric)(nil)).Elem()
}

func (o IntegrationMetricMapOutput) ToIntegrationMetricMapOutput() IntegrationMetricMapOutput {
	return o
}

func (o IntegrationMetricMapOutput) ToIntegrationMetricMapOutputWithContext(ctx context.Context) IntegrationMetricMapOutput {
	return o
}

func (o IntegrationMetricMapOutput) MapIndex(k pulumi.StringInput) IntegrationMetricOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *IntegrationMetric {
		return vs[0].(map[string]*IntegrationMetric)[vs[1].(string)]
	}).(IntegrationMetricOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*IntegrationMetricInput)(nil)).Elem(), &IntegrationMetric{})
	pulumi.RegisterInputType(reflect.TypeOf((*IntegrationMetricArrayInput)(nil)).Elem(), IntegrationMetricArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*IntegrationMetricMapInput)(nil)).Elem(), IntegrationMetricMap{})
	pulumi.RegisterOutputType(IntegrationMetricOutput{})
	pulumi.RegisterOutputType(IntegrationMetricArrayOutput{})
	pulumi.RegisterOutputType(IntegrationMetricMapOutput{})
}
