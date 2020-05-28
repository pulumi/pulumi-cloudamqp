// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package cloudamqp

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

type IntegrationMetric struct {
	pulumi.CustomResourceState

	// AWS access key identifier. (Cloudwatch)
	AccessKeyId pulumi.StringPtrOutput `pulumi:"accessKeyId"`
	// The API key for the integration service. (Librato)
	ApiKey pulumi.StringPtrOutput `pulumi:"apiKey"`
	// The email address registred for the integration service. (Librato)
	Email pulumi.StringPtrOutput `pulumi:"email"`
	// Instance identifier
	InstanceId pulumi.IntOutput `pulumi:"instanceId"`
	// The license key registred for the integration service. (New Relic)
	LicenseKey pulumi.StringPtrOutput `pulumi:"licenseKey"`
	// The name of metrics integration
	Name pulumi.StringOutput `pulumi:"name"`
	// (optional) whitelist using regular expression
	QueueWhitelist pulumi.StringPtrOutput `pulumi:"queueWhitelist"`
	// AWS region for Cloudwatch and [US/EU] for Data dog/New relic. (Cloudwatch, Data Dog, New Relic)
	Region pulumi.StringPtrOutput `pulumi:"region"`
	// AWS secret key. (Cloudwatch)
	SecretAccessKey pulumi.StringPtrOutput `pulumi:"secretAccessKey"`
	// (optional) tags. E.g. env=prod,region=europe
	Tags pulumi.StringPtrOutput `pulumi:"tags"`
	// (optional) whitelist using regular expression
	VhostWhitelist pulumi.StringPtrOutput `pulumi:"vhostWhitelist"`
}

// NewIntegrationMetric registers a new resource with the given unique name, arguments, and options.
func NewIntegrationMetric(ctx *pulumi.Context,
	name string, args *IntegrationMetricArgs, opts ...pulumi.ResourceOption) (*IntegrationMetric, error) {
	if args == nil || args.InstanceId == nil {
		return nil, errors.New("missing required argument 'InstanceId'")
	}
	if args == nil {
		args = &IntegrationMetricArgs{}
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
	// The email address registred for the integration service. (Librato)
	Email *string `pulumi:"email"`
	// Instance identifier
	InstanceId *int `pulumi:"instanceId"`
	// The license key registred for the integration service. (New Relic)
	LicenseKey *string `pulumi:"licenseKey"`
	// The name of metrics integration
	Name *string `pulumi:"name"`
	// (optional) whitelist using regular expression
	QueueWhitelist *string `pulumi:"queueWhitelist"`
	// AWS region for Cloudwatch and [US/EU] for Data dog/New relic. (Cloudwatch, Data Dog, New Relic)
	Region *string `pulumi:"region"`
	// AWS secret key. (Cloudwatch)
	SecretAccessKey *string `pulumi:"secretAccessKey"`
	// (optional) tags. E.g. env=prod,region=europe
	Tags *string `pulumi:"tags"`
	// (optional) whitelist using regular expression
	VhostWhitelist *string `pulumi:"vhostWhitelist"`
}

type IntegrationMetricState struct {
	// AWS access key identifier. (Cloudwatch)
	AccessKeyId pulumi.StringPtrInput
	// The API key for the integration service. (Librato)
	ApiKey pulumi.StringPtrInput
	// The email address registred for the integration service. (Librato)
	Email pulumi.StringPtrInput
	// Instance identifier
	InstanceId pulumi.IntPtrInput
	// The license key registred for the integration service. (New Relic)
	LicenseKey pulumi.StringPtrInput
	// The name of metrics integration
	Name pulumi.StringPtrInput
	// (optional) whitelist using regular expression
	QueueWhitelist pulumi.StringPtrInput
	// AWS region for Cloudwatch and [US/EU] for Data dog/New relic. (Cloudwatch, Data Dog, New Relic)
	Region pulumi.StringPtrInput
	// AWS secret key. (Cloudwatch)
	SecretAccessKey pulumi.StringPtrInput
	// (optional) tags. E.g. env=prod,region=europe
	Tags pulumi.StringPtrInput
	// (optional) whitelist using regular expression
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
	// The email address registred for the integration service. (Librato)
	Email *string `pulumi:"email"`
	// Instance identifier
	InstanceId int `pulumi:"instanceId"`
	// The license key registred for the integration service. (New Relic)
	LicenseKey *string `pulumi:"licenseKey"`
	// The name of metrics integration
	Name *string `pulumi:"name"`
	// (optional) whitelist using regular expression
	QueueWhitelist *string `pulumi:"queueWhitelist"`
	// AWS region for Cloudwatch and [US/EU] for Data dog/New relic. (Cloudwatch, Data Dog, New Relic)
	Region *string `pulumi:"region"`
	// AWS secret key. (Cloudwatch)
	SecretAccessKey *string `pulumi:"secretAccessKey"`
	// (optional) tags. E.g. env=prod,region=europe
	Tags *string `pulumi:"tags"`
	// (optional) whitelist using regular expression
	VhostWhitelist *string `pulumi:"vhostWhitelist"`
}

// The set of arguments for constructing a IntegrationMetric resource.
type IntegrationMetricArgs struct {
	// AWS access key identifier. (Cloudwatch)
	AccessKeyId pulumi.StringPtrInput
	// The API key for the integration service. (Librato)
	ApiKey pulumi.StringPtrInput
	// The email address registred for the integration service. (Librato)
	Email pulumi.StringPtrInput
	// Instance identifier
	InstanceId pulumi.IntInput
	// The license key registred for the integration service. (New Relic)
	LicenseKey pulumi.StringPtrInput
	// The name of metrics integration
	Name pulumi.StringPtrInput
	// (optional) whitelist using regular expression
	QueueWhitelist pulumi.StringPtrInput
	// AWS region for Cloudwatch and [US/EU] for Data dog/New relic. (Cloudwatch, Data Dog, New Relic)
	Region pulumi.StringPtrInput
	// AWS secret key. (Cloudwatch)
	SecretAccessKey pulumi.StringPtrInput
	// (optional) tags. E.g. env=prod,region=europe
	Tags pulumi.StringPtrInput
	// (optional) whitelist using regular expression
	VhostWhitelist pulumi.StringPtrInput
}

func (IntegrationMetricArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*integrationMetricArgs)(nil)).Elem()
}