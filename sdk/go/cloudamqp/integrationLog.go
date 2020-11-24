// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package cloudamqp

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// This resource allows you to create and manage third party log integrations for a CloudAMQP instance. Once configured, the logs produced will be forward to corresponding integration.
//
// Only available for dedicated subscription plans.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-cloudamqp/sdk/v2/go/cloudamqp"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := cloudamqp.NewIntegrationLog(ctx, "cloudwatch", &cloudamqp.IntegrationLogArgs{
// 			InstanceId:      pulumi.Any(cloudamqp_instance.Instance.Id),
// 			AccessKeyId:     pulumi.Any(_var.Aws_access_key_id),
// 			SecretAccessKey: pulumi.Any(_var.Aws_secret_access_key),
// 			Region:          pulumi.Any(_var.Aws_region),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = cloudamqp.NewIntegrationLog(ctx, "logentries", &cloudamqp.IntegrationLogArgs{
// 			InstanceId: pulumi.Any(cloudamqp_instance.Instance.Id),
// 			Token:      pulumi.Any(_var.Logentries_token),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = cloudamqp.NewIntegrationLog(ctx, "loggly", &cloudamqp.IntegrationLogArgs{
// 			InstanceId: pulumi.Any(cloudamqp_instance.Instance.Id),
// 			Token:      pulumi.Any(_var.Loggly_token),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = cloudamqp.NewIntegrationLog(ctx, "papertrail", &cloudamqp.IntegrationLogArgs{
// 			InstanceId: pulumi.Any(cloudamqp_instance.Instance.Id),
// 			Url:        pulumi.Any(_var.Papertrail_url),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = cloudamqp.NewIntegrationLog(ctx, "splunk", &cloudamqp.IntegrationLogArgs{
// 			InstanceId: pulumi.Any(cloudamqp_instance.Instance.Id),
// 			Token:      pulumi.Any(_var.Splunk_token),
// 			HostPort:   pulumi.Any(_var.Splunk_host_port),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = cloudamqp.NewIntegrationLog(ctx, "datadog", &cloudamqp.IntegrationLogArgs{
// 			InstanceId: pulumi.Any(cloudamqp_instance.Instance.Id),
// 			Region:     pulumi.Any(_var.Datadog_region),
// 			ApiKey:     pulumi.Any(_var.Datadog_api_key),
// 			Tags:       pulumi.Any(_var.Datadog_tags),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = cloudamqp.NewIntegrationLog(ctx, "stackdriver", &cloudamqp.IntegrationLogArgs{
// 			InstanceId:  pulumi.Any(cloudamqp_instance.Instance.Id),
// 			ProjectId:   pulumi.Any(_var.Stackdriver_project_id),
// 			PrivateKey:  pulumi.Any(_var.Stackdriver_private_key),
// 			ClientEmail: pulumi.Any(_var.Stackdriver_client_email),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ## Argument Reference (cloudwatchlog)
//
// Cloudwatch argument reference and example. Create an IAM user with programmatic access and the following permissions:
//
// * CreateLogGroup
// * CreateLogStream
// * DescribeLogGroups
// * DescribeLogStreams
// * PutLogEvents
//
// ## Integration service reference
//
// Valid names for third party log integration.
//
// | Name       | Description |
// |------------|---------------------------------------------------------------|
// | cloudwatchlog | Create a IAM with programmatic access. |
// | logentries | Create a Logentries token at https://logentries.com/app#/add-log/manual  |
// | loggly     | Create a Loggly token at https://{your-company}.loggly.com/tokens |
// | papertrail | Create a Papertrail endpoint https://papertrailapp.com/systems/setup |
// | splunk     | Create a HTTP Event Collector token at https://.cloud.splunk.com/en-US/manager/search/http-eventcollector |
// | datadog       | Create a Datadog API key at app.datadoghq.com |
// | stackdriver   | Create a service account and add 'monitor metrics writer' role, then download credentials. |
//
// ## Integration Type reference
//
// Valid arguments for third party log integrations.
//
// Required arguments for all integrations: name
//
// | Name | Type | Required arguments |
// | ---- | ---- | ---- |
// | CloudWatch | cloudwatchlog | access_key_id, secret_access_key, region |
// | Log Entries | logentries | token |
// | Loggly | loggly | token |
// | Papertrail | papertrail | url |
// | Splunk | splunk | token, hostPort |
// | Data Dog | datadog | region, api_keys, tags |
// | Stackdriver | stackdriver | project_id, private_key, clientEmail |
//
// ## Dependency
//
// This resource depends on CloudAMQP instance identifier, `cloudamqp_instance.instance.id`.
//
// ## Import
//
// `cloudamqp_integration_log`can be imported using the name argument of the resource together with CloudAMQP instance identifier. The name and identifier are CSV separated, see example below.
//
// ```sh
//  $ pulumi import cloudamqp:index/integrationLog:IntegrationLog <resource_name> <name>,<instance_id>`
// ```
type IntegrationLog struct {
	pulumi.CustomResourceState

	// AWS access key identifier.
	AccessKeyId pulumi.StringPtrOutput `pulumi:"accessKeyId"`
	// The API key.
	ApiKey pulumi.StringPtrOutput `pulumi:"apiKey"`
	// The client email registered for the integration service.
	ClientEmail pulumi.StringPtrOutput `pulumi:"clientEmail"`
	// Destination to send the logs.
	HostPort pulumi.StringPtrOutput `pulumi:"hostPort"`
	// Instance identifier used to make proxy calls
	InstanceId pulumi.IntOutput `pulumi:"instanceId"`
	// The name of the third party log integration. See
	Name pulumi.StringOutput `pulumi:"name"`
	// The private access key.
	PrivateKey pulumi.StringPtrOutput `pulumi:"privateKey"`
	// The project identifier.
	ProjectId pulumi.StringPtrOutput `pulumi:"projectId"`
	// Region hosting the integration service.
	Region pulumi.StringPtrOutput `pulumi:"region"`
	// AWS secret access key.
	SecretAccessKey pulumi.StringPtrOutput `pulumi:"secretAccessKey"`
	// Tag the integration, e.g. env=prod, region=europe.
	Tags pulumi.StringPtrOutput `pulumi:"tags"`
	// Token used for authentication.
	Token pulumi.StringPtrOutput `pulumi:"token"`
	// Endpoint to log integration.
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
	// AWS access key identifier.
	AccessKeyId *string `pulumi:"accessKeyId"`
	// The API key.
	ApiKey *string `pulumi:"apiKey"`
	// The client email registered for the integration service.
	ClientEmail *string `pulumi:"clientEmail"`
	// Destination to send the logs.
	HostPort *string `pulumi:"hostPort"`
	// Instance identifier used to make proxy calls
	InstanceId *int `pulumi:"instanceId"`
	// The name of the third party log integration. See
	Name *string `pulumi:"name"`
	// The private access key.
	PrivateKey *string `pulumi:"privateKey"`
	// The project identifier.
	ProjectId *string `pulumi:"projectId"`
	// Region hosting the integration service.
	Region *string `pulumi:"region"`
	// AWS secret access key.
	SecretAccessKey *string `pulumi:"secretAccessKey"`
	// Tag the integration, e.g. env=prod, region=europe.
	Tags *string `pulumi:"tags"`
	// Token used for authentication.
	Token *string `pulumi:"token"`
	// Endpoint to log integration.
	Url *string `pulumi:"url"`
}

type IntegrationLogState struct {
	// AWS access key identifier.
	AccessKeyId pulumi.StringPtrInput
	// The API key.
	ApiKey pulumi.StringPtrInput
	// The client email registered for the integration service.
	ClientEmail pulumi.StringPtrInput
	// Destination to send the logs.
	HostPort pulumi.StringPtrInput
	// Instance identifier used to make proxy calls
	InstanceId pulumi.IntPtrInput
	// The name of the third party log integration. See
	Name pulumi.StringPtrInput
	// The private access key.
	PrivateKey pulumi.StringPtrInput
	// The project identifier.
	ProjectId pulumi.StringPtrInput
	// Region hosting the integration service.
	Region pulumi.StringPtrInput
	// AWS secret access key.
	SecretAccessKey pulumi.StringPtrInput
	// Tag the integration, e.g. env=prod, region=europe.
	Tags pulumi.StringPtrInput
	// Token used for authentication.
	Token pulumi.StringPtrInput
	// Endpoint to log integration.
	Url pulumi.StringPtrInput
}

func (IntegrationLogState) ElementType() reflect.Type {
	return reflect.TypeOf((*integrationLogState)(nil)).Elem()
}

type integrationLogArgs struct {
	// AWS access key identifier.
	AccessKeyId *string `pulumi:"accessKeyId"`
	// The API key.
	ApiKey *string `pulumi:"apiKey"`
	// The client email registered for the integration service.
	ClientEmail *string `pulumi:"clientEmail"`
	// Destination to send the logs.
	HostPort *string `pulumi:"hostPort"`
	// Instance identifier used to make proxy calls
	InstanceId int `pulumi:"instanceId"`
	// The name of the third party log integration. See
	Name *string `pulumi:"name"`
	// The private access key.
	PrivateKey *string `pulumi:"privateKey"`
	// The project identifier.
	ProjectId *string `pulumi:"projectId"`
	// Region hosting the integration service.
	Region *string `pulumi:"region"`
	// AWS secret access key.
	SecretAccessKey *string `pulumi:"secretAccessKey"`
	// Tag the integration, e.g. env=prod, region=europe.
	Tags *string `pulumi:"tags"`
	// Token used for authentication.
	Token *string `pulumi:"token"`
	// Endpoint to log integration.
	Url *string `pulumi:"url"`
}

// The set of arguments for constructing a IntegrationLog resource.
type IntegrationLogArgs struct {
	// AWS access key identifier.
	AccessKeyId pulumi.StringPtrInput
	// The API key.
	ApiKey pulumi.StringPtrInput
	// The client email registered for the integration service.
	ClientEmail pulumi.StringPtrInput
	// Destination to send the logs.
	HostPort pulumi.StringPtrInput
	// Instance identifier used to make proxy calls
	InstanceId pulumi.IntInput
	// The name of the third party log integration. See
	Name pulumi.StringPtrInput
	// The private access key.
	PrivateKey pulumi.StringPtrInput
	// The project identifier.
	ProjectId pulumi.StringPtrInput
	// Region hosting the integration service.
	Region pulumi.StringPtrInput
	// AWS secret access key.
	SecretAccessKey pulumi.StringPtrInput
	// Tag the integration, e.g. env=prod, region=europe.
	Tags pulumi.StringPtrInput
	// Token used for authentication.
	Token pulumi.StringPtrInput
	// Endpoint to log integration.
	Url pulumi.StringPtrInput
}

func (IntegrationLogArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*integrationLogArgs)(nil)).Elem()
}

type IntegrationLogInput interface {
	pulumi.Input

	ToIntegrationLogOutput() IntegrationLogOutput
	ToIntegrationLogOutputWithContext(ctx context.Context) IntegrationLogOutput
}

func (IntegrationLog) ElementType() reflect.Type {
	return reflect.TypeOf((*IntegrationLog)(nil)).Elem()
}

func (i IntegrationLog) ToIntegrationLogOutput() IntegrationLogOutput {
	return i.ToIntegrationLogOutputWithContext(context.Background())
}

func (i IntegrationLog) ToIntegrationLogOutputWithContext(ctx context.Context) IntegrationLogOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IntegrationLogOutput)
}

type IntegrationLogOutput struct {
	*pulumi.OutputState
}

func (IntegrationLogOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IntegrationLogOutput)(nil)).Elem()
}

func (o IntegrationLogOutput) ToIntegrationLogOutput() IntegrationLogOutput {
	return o
}

func (o IntegrationLogOutput) ToIntegrationLogOutputWithContext(ctx context.Context) IntegrationLogOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(IntegrationLogOutput{})
}
