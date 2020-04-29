// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package cloudamqp

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

type IntegrationLog struct {
	pulumi.CustomResourceState

	// AWS access key identifier. (Cloudwatch)
	AccessKeyId pulumi.StringPtrOutput `pulumi:"accessKeyId"`
	// Destination to send the logs. (Splunk)
	HostPort pulumi.StringPtrOutput `pulumi:"hostPort"`
	// Instance identifier used to make proxy calls
	InstanceId pulumi.IntOutput `pulumi:"instanceId"`
	// The name of log integration
	Name pulumi.StringOutput `pulumi:"name"`
	// The region hosting integration service. (Cloudwatch)
	Region pulumi.StringPtrOutput `pulumi:"region"`
	// AWS secret access key. (Cloudwatch)
	SecretAccessKey pulumi.StringPtrOutput `pulumi:"secretAccessKey"`
	// The token used for authentication. (Loggly, Logentries, Splunk)
	Token pulumi.StringPtrOutput `pulumi:"token"`
	// The URL to push the logs to. (Papertrail)
	Url pulumi.StringPtrOutput `pulumi:"url"`
}

// NewIntegrationLog registers a new resource with the given unique name, arguments, and options.
func NewIntegrationLog(ctx *pulumi.Context,
	name string, args *IntegrationLogArgs, opts ...pulumi.ResourceOption) (*IntegrationLog, error) {
	if args == nil || args.InstanceId == nil {
		return nil, errors.New("missing required argument 'InstanceId'")
	}
	if args == nil {
		args = &IntegrationLogArgs{}
	}
	var resource IntegrationLog
	err := ctx.RegisterResource("cloudamqp:index/integrationLog:IntegrationLog", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIntegrationLog gets an existing IntegrationLog resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIntegrationLog(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IntegrationLogState, opts ...pulumi.ResourceOption) (*IntegrationLog, error) {
	var resource IntegrationLog
	err := ctx.ReadResource("cloudamqp:index/integrationLog:IntegrationLog", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering IntegrationLog resources.
type integrationLogState struct {
	// AWS access key identifier. (Cloudwatch)
	AccessKeyId *string `pulumi:"accessKeyId"`
	// Destination to send the logs. (Splunk)
	HostPort *string `pulumi:"hostPort"`
	// Instance identifier used to make proxy calls
	InstanceId *int `pulumi:"instanceId"`
	// The name of log integration
	Name *string `pulumi:"name"`
	// The region hosting integration service. (Cloudwatch)
	Region *string `pulumi:"region"`
	// AWS secret access key. (Cloudwatch)
	SecretAccessKey *string `pulumi:"secretAccessKey"`
	// The token used for authentication. (Loggly, Logentries, Splunk)
	Token *string `pulumi:"token"`
	// The URL to push the logs to. (Papertrail)
	Url *string `pulumi:"url"`
}

type IntegrationLogState struct {
	// AWS access key identifier. (Cloudwatch)
	AccessKeyId pulumi.StringPtrInput
	// Destination to send the logs. (Splunk)
	HostPort pulumi.StringPtrInput
	// Instance identifier used to make proxy calls
	InstanceId pulumi.IntPtrInput
	// The name of log integration
	Name pulumi.StringPtrInput
	// The region hosting integration service. (Cloudwatch)
	Region pulumi.StringPtrInput
	// AWS secret access key. (Cloudwatch)
	SecretAccessKey pulumi.StringPtrInput
	// The token used for authentication. (Loggly, Logentries, Splunk)
	Token pulumi.StringPtrInput
	// The URL to push the logs to. (Papertrail)
	Url pulumi.StringPtrInput
}

func (IntegrationLogState) ElementType() reflect.Type {
	return reflect.TypeOf((*integrationLogState)(nil)).Elem()
}

type integrationLogArgs struct {
	// AWS access key identifier. (Cloudwatch)
	AccessKeyId *string `pulumi:"accessKeyId"`
	// Destination to send the logs. (Splunk)
	HostPort *string `pulumi:"hostPort"`
	// Instance identifier used to make proxy calls
	InstanceId int `pulumi:"instanceId"`
	// The name of log integration
	Name *string `pulumi:"name"`
	// The region hosting integration service. (Cloudwatch)
	Region *string `pulumi:"region"`
	// AWS secret access key. (Cloudwatch)
	SecretAccessKey *string `pulumi:"secretAccessKey"`
	// The token used for authentication. (Loggly, Logentries, Splunk)
	Token *string `pulumi:"token"`
	// The URL to push the logs to. (Papertrail)
	Url *string `pulumi:"url"`
}

// The set of arguments for constructing a IntegrationLog resource.
type IntegrationLogArgs struct {
	// AWS access key identifier. (Cloudwatch)
	AccessKeyId pulumi.StringPtrInput
	// Destination to send the logs. (Splunk)
	HostPort pulumi.StringPtrInput
	// Instance identifier used to make proxy calls
	InstanceId pulumi.IntInput
	// The name of log integration
	Name pulumi.StringPtrInput
	// The region hosting integration service. (Cloudwatch)
	Region pulumi.StringPtrInput
	// AWS secret access key. (Cloudwatch)
	SecretAccessKey pulumi.StringPtrInput
	// The token used for authentication. (Loggly, Logentries, Splunk)
	Token pulumi.StringPtrInput
	// The URL to push the logs to. (Papertrail)
	Url pulumi.StringPtrInput
}

func (IntegrationLogArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*integrationLogArgs)(nil)).Elem()
}
