// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cloudamqp

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-cloudamqp/sdk/v3/go/cloudamqp/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This resource allows you to create and manage third party log integrations for a CloudAMQP instance.
// Once configured, the logs produced will be forward to corresponding integration.
//
// Only available for dedicated subscription plans.
//
// ## Example Usage
//
// <details>
//
//	<summary>
//	  <b>
//	    <i>Azure monitor log integration</i>
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
//			_, err := cloudamqp.NewIntegrationLog(ctx, "azure_monitor", &cloudamqp.IntegrationLogArgs{
//				InstanceId:        pulumi.Any(instance.Id),
//				Name:              pulumi.String("azure_monitor"),
//				TenantId:          pulumi.Any(azmTentantId),
//				ApplicationId:     pulumi.Any(azmApplicationId),
//				ApplicationSecret: pulumi.Any(azmApplicationSecret),
//				DceUri:            pulumi.Any(azmDceUri),
//				Table:             pulumi.Any(azmTable),
//				DcrId:             pulumi.Any(azmDcrId),
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
//	    <i>Cloudwatch log integration</i>
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
//			_, err := cloudamqp.NewIntegrationLog(ctx, "cloudwatch", &cloudamqp.IntegrationLogArgs{
//				InstanceId:      pulumi.Any(instance.Id),
//				Name:            pulumi.String("cloudwatchlog"),
//				AccessKeyId:     pulumi.Any(awsAccessKeyId),
//				SecretAccessKey: pulumi.Any(awsSecretAccessKey),
//				Region:          pulumi.Any(awsRegion),
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
//	    <i>Coralogix log integration</i>
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
//			_, err := cloudamqp.NewIntegrationLog(ctx, "coralogix", &cloudamqp.IntegrationLogArgs{
//				InstanceId:  pulumi.Any(instance.Id),
//				Name:        pulumi.String("coralogix"),
//				PrivateKey:  pulumi.Any(coralogixSendDataKey),
//				Endpoint:    pulumi.Any(coralogixEndpoint),
//				Application: pulumi.Any(coralogixApplication),
//				Subsystem:   pulumi.Any(instance.Host),
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
//	    <i>Datadog log integration</i>
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
//			_, err := cloudamqp.NewIntegrationLog(ctx, "datadog", &cloudamqp.IntegrationLogArgs{
//				InstanceId: pulumi.Any(instance.Id),
//				Name:       pulumi.String("datadog"),
//				Region:     pulumi.Any(datadogRegion),
//				ApiKey:     pulumi.Any(datadogApiKey),
//				Tags:       pulumi.String("env=prod,region=us1,version=v1.0"),
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
//	    <i>Logentries log integration</i>
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
//			_, err := cloudamqp.NewIntegrationLog(ctx, "logentries", &cloudamqp.IntegrationLogArgs{
//				InstanceId: pulumi.Any(instance.Id),
//				Name:       pulumi.String("logentries"),
//				Token:      pulumi.Any(logentriesToken),
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
//	    <i>Loggly log integration</i>
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
//			_, err := cloudamqp.NewIntegrationLog(ctx, "loggly", &cloudamqp.IntegrationLogArgs{
//				InstanceId: pulumi.Any(instance.Id),
//				Name:       pulumi.String("loggly"),
//				Token:      pulumi.Any(logglyToken),
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
//	    <i>Papertrail log integration</i>
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
//			_, err := cloudamqp.NewIntegrationLog(ctx, "papertrail", &cloudamqp.IntegrationLogArgs{
//				InstanceId: pulumi.Any(instance.Id),
//				Name:       pulumi.String("papertrail"),
//				Url:        pulumi.Any(papertrailUrl),
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
//	    <i>Scalyr log integration</i>
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
//			_, err := cloudamqp.NewIntegrationLog(ctx, "scalyr", &cloudamqp.IntegrationLogArgs{
//				InstanceId: pulumi.Any(instance.Id),
//				Name:       pulumi.String("scalyr"),
//				Token:      pulumi.Any(scalyrToken),
//				Host:       pulumi.Any(scalyrHost),
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
//	    <i>Splunk log integration</i>
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
//			_, err := cloudamqp.NewIntegrationLog(ctx, "splunk", &cloudamqp.IntegrationLogArgs{
//				InstanceId: pulumi.Any(instance.Id),
//				Name:       pulumi.String("splunk"),
//				Token:      pulumi.Any(splunkToken),
//				HostPort:   pulumi.Any(splunkHostPort),
//				SourceType: pulumi.String("generic_single_line"),
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
//	    <i>Stackdriver log integration (v1.20.2 or older versions)</i>
//	  </b>
//	</summary>
//
// Use variable file populated with project_id, privateKey and clientEmail
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
//			_, err := cloudamqp.NewIntegrationLog(ctx, "stackdriver", &cloudamqp.IntegrationLogArgs{
//				InstanceId:  pulumi.Any(instance.Id),
//				Name:        pulumi.String("stackdriver"),
//				ProjectId:   pulumi.Any(stackdriverProjectId),
//				PrivateKey:  pulumi.Any(stackdriverPrivateKey),
//				ClientEmail: pulumi.Any(stackdriverClientEmail),
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
// or by using googleServiceAccountKey resource from Google provider
//
// ## Import
//
// `cloudamqp_integration_log`can be imported using the resource identifier together with CloudAMQP
//
// instance identifier. The identifiers are CSV separated, see example below. To retrieve the resource,
//
// use [CloudAMQP API list integration].
//
// From Terraform v1.5.0, the `import` block can be used to import this resource:
//
// hcl
//
// import {
//
//	to = cloudamqp_integration_log.this
//
//	id = format("<id>,%s", cloudamqp_instance.instance.id)
//
// }
//
// ```sh
// $ pulumi import cloudamqp:index/integrationLog:IntegrationLog this <id>,<instance_id>`
// ```
//
// [integration type reference]: #integration-type-reference
//
// [CloudAMQP API list integration]: https://docs.cloudamqp.com/cloudamqp_api.html#list-log-integrations
//
// [CloudAMQP API add integration]: https://docs.cloudamqp.com/cloudamqp_api.html#add-log-integration
// [Datadog documentation]: https://docs.datadoghq.com/getting_started/tagging/#define-tags
type IntegrationLog struct {
	pulumi.CustomResourceState

	// AWS access key identifier.
	AccessKeyId pulumi.StringPtrOutput `pulumi:"accessKeyId"`
	// The API key.
	ApiKey pulumi.StringPtrOutput `pulumi:"apiKey"`
	// The application name for Coralogix.
	Application pulumi.StringPtrOutput `pulumi:"application"`
	// The application identifier for Azure monitor.
	ApplicationId pulumi.StringPtrOutput `pulumi:"applicationId"`
	// The application secret for Azure monitor.
	ApplicationSecret pulumi.StringPtrOutput `pulumi:"applicationSecret"`
	// The client email registered for the integration service.
	ClientEmail pulumi.StringOutput `pulumi:"clientEmail"`
	// Google Service Account private key credentials.
	Credentials pulumi.StringPtrOutput `pulumi:"credentials"`
	// The data collection endpoint for Azure monitor.
	DceUri pulumi.StringPtrOutput `pulumi:"dceUri"`
	// ID of data collection rule that your DCE is linked to for Azure
	// Monitor.
	//
	// This is the full list of all arguments. Only a subset of arguments are used based on which type of
	// integration used. See [integration type reference] table below for more information.
	DcrId pulumi.StringPtrOutput `pulumi:"dcrId"`
	// The syslog destination to send the logs to for Coralogix.
	Endpoint pulumi.StringPtrOutput `pulumi:"endpoint"`
	// The host for Scalyr integration. (app.scalyr.com,
	// app.eu.scalyr.com)
	Host pulumi.StringPtrOutput `pulumi:"host"`
	// Destination to send the logs.
	HostPort pulumi.StringPtrOutput `pulumi:"hostPort"`
	// Instance identifier used to make proxy calls
	InstanceId pulumi.IntOutput `pulumi:"instanceId"`
	// The name of the third party log integration. See
	// [integration type reference]
	Name pulumi.StringOutput `pulumi:"name"`
	// The private access key.
	PrivateKey pulumi.StringOutput `pulumi:"privateKey"`
	// Private key identifier. (Stackdriver)
	PrivateKeyId pulumi.StringOutput `pulumi:"privateKeyId"`
	// The project identifier.
	ProjectId pulumi.StringOutput `pulumi:"projectId"`
	// Region hosting the integration service.
	Region pulumi.StringPtrOutput `pulumi:"region"`
	// AWS secret access key.
	SecretAccessKey pulumi.StringPtrOutput `pulumi:"secretAccessKey"`
	// Assign source type to the data exported, eg. generic_single_line.
	// (Splunk)
	Sourcetype pulumi.StringPtrOutput `pulumi:"sourcetype"`
	// The subsystem name for Coralogix.
	Subsystem pulumi.StringPtrOutput `pulumi:"subsystem"`
	// The table name for Azure monitor.
	Table pulumi.StringPtrOutput `pulumi:"table"`
	// Tags. e.g. `env=prod,region=europe`.
	//
	// ***Note:*** If tags are used with Datadog. The value part (prod, europe, ...) must start with a
	// letter, read more about tags format in the [Datadog documentation].
	Tags pulumi.StringPtrOutput `pulumi:"tags"`
	// The tenant identifier for Azure monitor.
	TenantId pulumi.StringPtrOutput `pulumi:"tenantId"`
	// Token used for authentication.
	Token pulumi.StringPtrOutput `pulumi:"token"`
	// Endpoint to log integration.
	Url pulumi.StringPtrOutput `pulumi:"url"`
}

// NewIntegrationLog registers a new resource with the given unique name, arguments, and options.
func NewIntegrationLog(ctx *pulumi.Context,
	name string, args *IntegrationLogArgs, opts ...pulumi.ResourceOption) (*IntegrationLog, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.InstanceId == nil {
		return nil, errors.New("invalid value for required argument 'InstanceId'")
	}
	if args.AccessKeyId != nil {
		args.AccessKeyId = pulumi.ToSecret(args.AccessKeyId).(pulumi.StringPtrInput)
	}
	if args.ApiKey != nil {
		args.ApiKey = pulumi.ToSecret(args.ApiKey).(pulumi.StringPtrInput)
	}
	if args.ApplicationSecret != nil {
		args.ApplicationSecret = pulumi.ToSecret(args.ApplicationSecret).(pulumi.StringPtrInput)
	}
	if args.Credentials != nil {
		args.Credentials = pulumi.ToSecret(args.Credentials).(pulumi.StringPtrInput)
	}
	if args.PrivateKey != nil {
		args.PrivateKey = pulumi.ToSecret(args.PrivateKey).(pulumi.StringPtrInput)
	}
	if args.PrivateKeyId != nil {
		args.PrivateKeyId = pulumi.ToSecret(args.PrivateKeyId).(pulumi.StringPtrInput)
	}
	if args.SecretAccessKey != nil {
		args.SecretAccessKey = pulumi.ToSecret(args.SecretAccessKey).(pulumi.StringPtrInput)
	}
	if args.Token != nil {
		args.Token = pulumi.ToSecret(args.Token).(pulumi.StringPtrInput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"accessKeyId",
		"apiKey",
		"applicationSecret",
		"credentials",
		"privateKey",
		"privateKeyId",
		"secretAccessKey",
		"token",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
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
	// The application name for Coralogix.
	Application *string `pulumi:"application"`
	// The application identifier for Azure monitor.
	ApplicationId *string `pulumi:"applicationId"`
	// The application secret for Azure monitor.
	ApplicationSecret *string `pulumi:"applicationSecret"`
	// The client email registered for the integration service.
	ClientEmail *string `pulumi:"clientEmail"`
	// Google Service Account private key credentials.
	Credentials *string `pulumi:"credentials"`
	// The data collection endpoint for Azure monitor.
	DceUri *string `pulumi:"dceUri"`
	// ID of data collection rule that your DCE is linked to for Azure
	// Monitor.
	//
	// This is the full list of all arguments. Only a subset of arguments are used based on which type of
	// integration used. See [integration type reference] table below for more information.
	DcrId *string `pulumi:"dcrId"`
	// The syslog destination to send the logs to for Coralogix.
	Endpoint *string `pulumi:"endpoint"`
	// The host for Scalyr integration. (app.scalyr.com,
	// app.eu.scalyr.com)
	Host *string `pulumi:"host"`
	// Destination to send the logs.
	HostPort *string `pulumi:"hostPort"`
	// Instance identifier used to make proxy calls
	InstanceId *int `pulumi:"instanceId"`
	// The name of the third party log integration. See
	// [integration type reference]
	Name *string `pulumi:"name"`
	// The private access key.
	PrivateKey *string `pulumi:"privateKey"`
	// Private key identifier. (Stackdriver)
	PrivateKeyId *string `pulumi:"privateKeyId"`
	// The project identifier.
	ProjectId *string `pulumi:"projectId"`
	// Region hosting the integration service.
	Region *string `pulumi:"region"`
	// AWS secret access key.
	SecretAccessKey *string `pulumi:"secretAccessKey"`
	// Assign source type to the data exported, eg. generic_single_line.
	// (Splunk)
	Sourcetype *string `pulumi:"sourcetype"`
	// The subsystem name for Coralogix.
	Subsystem *string `pulumi:"subsystem"`
	// The table name for Azure monitor.
	Table *string `pulumi:"table"`
	// Tags. e.g. `env=prod,region=europe`.
	//
	// ***Note:*** If tags are used with Datadog. The value part (prod, europe, ...) must start with a
	// letter, read more about tags format in the [Datadog documentation].
	Tags *string `pulumi:"tags"`
	// The tenant identifier for Azure monitor.
	TenantId *string `pulumi:"tenantId"`
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
	// The application name for Coralogix.
	Application pulumi.StringPtrInput
	// The application identifier for Azure monitor.
	ApplicationId pulumi.StringPtrInput
	// The application secret for Azure monitor.
	ApplicationSecret pulumi.StringPtrInput
	// The client email registered for the integration service.
	ClientEmail pulumi.StringPtrInput
	// Google Service Account private key credentials.
	Credentials pulumi.StringPtrInput
	// The data collection endpoint for Azure monitor.
	DceUri pulumi.StringPtrInput
	// ID of data collection rule that your DCE is linked to for Azure
	// Monitor.
	//
	// This is the full list of all arguments. Only a subset of arguments are used based on which type of
	// integration used. See [integration type reference] table below for more information.
	DcrId pulumi.StringPtrInput
	// The syslog destination to send the logs to for Coralogix.
	Endpoint pulumi.StringPtrInput
	// The host for Scalyr integration. (app.scalyr.com,
	// app.eu.scalyr.com)
	Host pulumi.StringPtrInput
	// Destination to send the logs.
	HostPort pulumi.StringPtrInput
	// Instance identifier used to make proxy calls
	InstanceId pulumi.IntPtrInput
	// The name of the third party log integration. See
	// [integration type reference]
	Name pulumi.StringPtrInput
	// The private access key.
	PrivateKey pulumi.StringPtrInput
	// Private key identifier. (Stackdriver)
	PrivateKeyId pulumi.StringPtrInput
	// The project identifier.
	ProjectId pulumi.StringPtrInput
	// Region hosting the integration service.
	Region pulumi.StringPtrInput
	// AWS secret access key.
	SecretAccessKey pulumi.StringPtrInput
	// Assign source type to the data exported, eg. generic_single_line.
	// (Splunk)
	Sourcetype pulumi.StringPtrInput
	// The subsystem name for Coralogix.
	Subsystem pulumi.StringPtrInput
	// The table name for Azure monitor.
	Table pulumi.StringPtrInput
	// Tags. e.g. `env=prod,region=europe`.
	//
	// ***Note:*** If tags are used with Datadog. The value part (prod, europe, ...) must start with a
	// letter, read more about tags format in the [Datadog documentation].
	Tags pulumi.StringPtrInput
	// The tenant identifier for Azure monitor.
	TenantId pulumi.StringPtrInput
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
	// The application name for Coralogix.
	Application *string `pulumi:"application"`
	// The application identifier for Azure monitor.
	ApplicationId *string `pulumi:"applicationId"`
	// The application secret for Azure monitor.
	ApplicationSecret *string `pulumi:"applicationSecret"`
	// The client email registered for the integration service.
	ClientEmail *string `pulumi:"clientEmail"`
	// Google Service Account private key credentials.
	Credentials *string `pulumi:"credentials"`
	// The data collection endpoint for Azure monitor.
	DceUri *string `pulumi:"dceUri"`
	// ID of data collection rule that your DCE is linked to for Azure
	// Monitor.
	//
	// This is the full list of all arguments. Only a subset of arguments are used based on which type of
	// integration used. See [integration type reference] table below for more information.
	DcrId *string `pulumi:"dcrId"`
	// The syslog destination to send the logs to for Coralogix.
	Endpoint *string `pulumi:"endpoint"`
	// The host for Scalyr integration. (app.scalyr.com,
	// app.eu.scalyr.com)
	Host *string `pulumi:"host"`
	// Destination to send the logs.
	HostPort *string `pulumi:"hostPort"`
	// Instance identifier used to make proxy calls
	InstanceId int `pulumi:"instanceId"`
	// The name of the third party log integration. See
	// [integration type reference]
	Name *string `pulumi:"name"`
	// The private access key.
	PrivateKey *string `pulumi:"privateKey"`
	// Private key identifier. (Stackdriver)
	PrivateKeyId *string `pulumi:"privateKeyId"`
	// The project identifier.
	ProjectId *string `pulumi:"projectId"`
	// Region hosting the integration service.
	Region *string `pulumi:"region"`
	// AWS secret access key.
	SecretAccessKey *string `pulumi:"secretAccessKey"`
	// Assign source type to the data exported, eg. generic_single_line.
	// (Splunk)
	Sourcetype *string `pulumi:"sourcetype"`
	// The subsystem name for Coralogix.
	Subsystem *string `pulumi:"subsystem"`
	// The table name for Azure monitor.
	Table *string `pulumi:"table"`
	// Tags. e.g. `env=prod,region=europe`.
	//
	// ***Note:*** If tags are used with Datadog. The value part (prod, europe, ...) must start with a
	// letter, read more about tags format in the [Datadog documentation].
	Tags *string `pulumi:"tags"`
	// The tenant identifier for Azure monitor.
	TenantId *string `pulumi:"tenantId"`
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
	// The application name for Coralogix.
	Application pulumi.StringPtrInput
	// The application identifier for Azure monitor.
	ApplicationId pulumi.StringPtrInput
	// The application secret for Azure monitor.
	ApplicationSecret pulumi.StringPtrInput
	// The client email registered for the integration service.
	ClientEmail pulumi.StringPtrInput
	// Google Service Account private key credentials.
	Credentials pulumi.StringPtrInput
	// The data collection endpoint for Azure monitor.
	DceUri pulumi.StringPtrInput
	// ID of data collection rule that your DCE is linked to for Azure
	// Monitor.
	//
	// This is the full list of all arguments. Only a subset of arguments are used based on which type of
	// integration used. See [integration type reference] table below for more information.
	DcrId pulumi.StringPtrInput
	// The syslog destination to send the logs to for Coralogix.
	Endpoint pulumi.StringPtrInput
	// The host for Scalyr integration. (app.scalyr.com,
	// app.eu.scalyr.com)
	Host pulumi.StringPtrInput
	// Destination to send the logs.
	HostPort pulumi.StringPtrInput
	// Instance identifier used to make proxy calls
	InstanceId pulumi.IntInput
	// The name of the third party log integration. See
	// [integration type reference]
	Name pulumi.StringPtrInput
	// The private access key.
	PrivateKey pulumi.StringPtrInput
	// Private key identifier. (Stackdriver)
	PrivateKeyId pulumi.StringPtrInput
	// The project identifier.
	ProjectId pulumi.StringPtrInput
	// Region hosting the integration service.
	Region pulumi.StringPtrInput
	// AWS secret access key.
	SecretAccessKey pulumi.StringPtrInput
	// Assign source type to the data exported, eg. generic_single_line.
	// (Splunk)
	Sourcetype pulumi.StringPtrInput
	// The subsystem name for Coralogix.
	Subsystem pulumi.StringPtrInput
	// The table name for Azure monitor.
	Table pulumi.StringPtrInput
	// Tags. e.g. `env=prod,region=europe`.
	//
	// ***Note:*** If tags are used with Datadog. The value part (prod, europe, ...) must start with a
	// letter, read more about tags format in the [Datadog documentation].
	Tags pulumi.StringPtrInput
	// The tenant identifier for Azure monitor.
	TenantId pulumi.StringPtrInput
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

func (*IntegrationLog) ElementType() reflect.Type {
	return reflect.TypeOf((**IntegrationLog)(nil)).Elem()
}

func (i *IntegrationLog) ToIntegrationLogOutput() IntegrationLogOutput {
	return i.ToIntegrationLogOutputWithContext(context.Background())
}

func (i *IntegrationLog) ToIntegrationLogOutputWithContext(ctx context.Context) IntegrationLogOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IntegrationLogOutput)
}

// IntegrationLogArrayInput is an input type that accepts IntegrationLogArray and IntegrationLogArrayOutput values.
// You can construct a concrete instance of `IntegrationLogArrayInput` via:
//
//	IntegrationLogArray{ IntegrationLogArgs{...} }
type IntegrationLogArrayInput interface {
	pulumi.Input

	ToIntegrationLogArrayOutput() IntegrationLogArrayOutput
	ToIntegrationLogArrayOutputWithContext(context.Context) IntegrationLogArrayOutput
}

type IntegrationLogArray []IntegrationLogInput

func (IntegrationLogArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*IntegrationLog)(nil)).Elem()
}

func (i IntegrationLogArray) ToIntegrationLogArrayOutput() IntegrationLogArrayOutput {
	return i.ToIntegrationLogArrayOutputWithContext(context.Background())
}

func (i IntegrationLogArray) ToIntegrationLogArrayOutputWithContext(ctx context.Context) IntegrationLogArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IntegrationLogArrayOutput)
}

// IntegrationLogMapInput is an input type that accepts IntegrationLogMap and IntegrationLogMapOutput values.
// You can construct a concrete instance of `IntegrationLogMapInput` via:
//
//	IntegrationLogMap{ "key": IntegrationLogArgs{...} }
type IntegrationLogMapInput interface {
	pulumi.Input

	ToIntegrationLogMapOutput() IntegrationLogMapOutput
	ToIntegrationLogMapOutputWithContext(context.Context) IntegrationLogMapOutput
}

type IntegrationLogMap map[string]IntegrationLogInput

func (IntegrationLogMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*IntegrationLog)(nil)).Elem()
}

func (i IntegrationLogMap) ToIntegrationLogMapOutput() IntegrationLogMapOutput {
	return i.ToIntegrationLogMapOutputWithContext(context.Background())
}

func (i IntegrationLogMap) ToIntegrationLogMapOutputWithContext(ctx context.Context) IntegrationLogMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IntegrationLogMapOutput)
}

type IntegrationLogOutput struct{ *pulumi.OutputState }

func (IntegrationLogOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IntegrationLog)(nil)).Elem()
}

func (o IntegrationLogOutput) ToIntegrationLogOutput() IntegrationLogOutput {
	return o
}

func (o IntegrationLogOutput) ToIntegrationLogOutputWithContext(ctx context.Context) IntegrationLogOutput {
	return o
}

// AWS access key identifier.
func (o IntegrationLogOutput) AccessKeyId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IntegrationLog) pulumi.StringPtrOutput { return v.AccessKeyId }).(pulumi.StringPtrOutput)
}

// The API key.
func (o IntegrationLogOutput) ApiKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IntegrationLog) pulumi.StringPtrOutput { return v.ApiKey }).(pulumi.StringPtrOutput)
}

// The application name for Coralogix.
func (o IntegrationLogOutput) Application() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IntegrationLog) pulumi.StringPtrOutput { return v.Application }).(pulumi.StringPtrOutput)
}

// The application identifier for Azure monitor.
func (o IntegrationLogOutput) ApplicationId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IntegrationLog) pulumi.StringPtrOutput { return v.ApplicationId }).(pulumi.StringPtrOutput)
}

// The application secret for Azure monitor.
func (o IntegrationLogOutput) ApplicationSecret() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IntegrationLog) pulumi.StringPtrOutput { return v.ApplicationSecret }).(pulumi.StringPtrOutput)
}

// The client email registered for the integration service.
func (o IntegrationLogOutput) ClientEmail() pulumi.StringOutput {
	return o.ApplyT(func(v *IntegrationLog) pulumi.StringOutput { return v.ClientEmail }).(pulumi.StringOutput)
}

// Google Service Account private key credentials.
func (o IntegrationLogOutput) Credentials() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IntegrationLog) pulumi.StringPtrOutput { return v.Credentials }).(pulumi.StringPtrOutput)
}

// The data collection endpoint for Azure monitor.
func (o IntegrationLogOutput) DceUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IntegrationLog) pulumi.StringPtrOutput { return v.DceUri }).(pulumi.StringPtrOutput)
}

// ID of data collection rule that your DCE is linked to for Azure
// Monitor.
//
// This is the full list of all arguments. Only a subset of arguments are used based on which type of
// integration used. See [integration type reference] table below for more information.
func (o IntegrationLogOutput) DcrId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IntegrationLog) pulumi.StringPtrOutput { return v.DcrId }).(pulumi.StringPtrOutput)
}

// The syslog destination to send the logs to for Coralogix.
func (o IntegrationLogOutput) Endpoint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IntegrationLog) pulumi.StringPtrOutput { return v.Endpoint }).(pulumi.StringPtrOutput)
}

// The host for Scalyr integration. (app.scalyr.com,
// app.eu.scalyr.com)
func (o IntegrationLogOutput) Host() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IntegrationLog) pulumi.StringPtrOutput { return v.Host }).(pulumi.StringPtrOutput)
}

// Destination to send the logs.
func (o IntegrationLogOutput) HostPort() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IntegrationLog) pulumi.StringPtrOutput { return v.HostPort }).(pulumi.StringPtrOutput)
}

// Instance identifier used to make proxy calls
func (o IntegrationLogOutput) InstanceId() pulumi.IntOutput {
	return o.ApplyT(func(v *IntegrationLog) pulumi.IntOutput { return v.InstanceId }).(pulumi.IntOutput)
}

// The name of the third party log integration. See
// [integration type reference]
func (o IntegrationLogOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *IntegrationLog) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The private access key.
func (o IntegrationLogOutput) PrivateKey() pulumi.StringOutput {
	return o.ApplyT(func(v *IntegrationLog) pulumi.StringOutput { return v.PrivateKey }).(pulumi.StringOutput)
}

// Private key identifier. (Stackdriver)
func (o IntegrationLogOutput) PrivateKeyId() pulumi.StringOutput {
	return o.ApplyT(func(v *IntegrationLog) pulumi.StringOutput { return v.PrivateKeyId }).(pulumi.StringOutput)
}

// The project identifier.
func (o IntegrationLogOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v *IntegrationLog) pulumi.StringOutput { return v.ProjectId }).(pulumi.StringOutput)
}

// Region hosting the integration service.
func (o IntegrationLogOutput) Region() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IntegrationLog) pulumi.StringPtrOutput { return v.Region }).(pulumi.StringPtrOutput)
}

// AWS secret access key.
func (o IntegrationLogOutput) SecretAccessKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IntegrationLog) pulumi.StringPtrOutput { return v.SecretAccessKey }).(pulumi.StringPtrOutput)
}

// Assign source type to the data exported, eg. generic_single_line.
// (Splunk)
func (o IntegrationLogOutput) Sourcetype() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IntegrationLog) pulumi.StringPtrOutput { return v.Sourcetype }).(pulumi.StringPtrOutput)
}

// The subsystem name for Coralogix.
func (o IntegrationLogOutput) Subsystem() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IntegrationLog) pulumi.StringPtrOutput { return v.Subsystem }).(pulumi.StringPtrOutput)
}

// The table name for Azure monitor.
func (o IntegrationLogOutput) Table() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IntegrationLog) pulumi.StringPtrOutput { return v.Table }).(pulumi.StringPtrOutput)
}

// Tags. e.g. `env=prod,region=europe`.
//
// ***Note:*** If tags are used with Datadog. The value part (prod, europe, ...) must start with a
// letter, read more about tags format in the [Datadog documentation].
func (o IntegrationLogOutput) Tags() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IntegrationLog) pulumi.StringPtrOutput { return v.Tags }).(pulumi.StringPtrOutput)
}

// The tenant identifier for Azure monitor.
func (o IntegrationLogOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IntegrationLog) pulumi.StringPtrOutput { return v.TenantId }).(pulumi.StringPtrOutput)
}

// Token used for authentication.
func (o IntegrationLogOutput) Token() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IntegrationLog) pulumi.StringPtrOutput { return v.Token }).(pulumi.StringPtrOutput)
}

// Endpoint to log integration.
func (o IntegrationLogOutput) Url() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IntegrationLog) pulumi.StringPtrOutput { return v.Url }).(pulumi.StringPtrOutput)
}

type IntegrationLogArrayOutput struct{ *pulumi.OutputState }

func (IntegrationLogArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*IntegrationLog)(nil)).Elem()
}

func (o IntegrationLogArrayOutput) ToIntegrationLogArrayOutput() IntegrationLogArrayOutput {
	return o
}

func (o IntegrationLogArrayOutput) ToIntegrationLogArrayOutputWithContext(ctx context.Context) IntegrationLogArrayOutput {
	return o
}

func (o IntegrationLogArrayOutput) Index(i pulumi.IntInput) IntegrationLogOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *IntegrationLog {
		return vs[0].([]*IntegrationLog)[vs[1].(int)]
	}).(IntegrationLogOutput)
}

type IntegrationLogMapOutput struct{ *pulumi.OutputState }

func (IntegrationLogMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*IntegrationLog)(nil)).Elem()
}

func (o IntegrationLogMapOutput) ToIntegrationLogMapOutput() IntegrationLogMapOutput {
	return o
}

func (o IntegrationLogMapOutput) ToIntegrationLogMapOutputWithContext(ctx context.Context) IntegrationLogMapOutput {
	return o
}

func (o IntegrationLogMapOutput) MapIndex(k pulumi.StringInput) IntegrationLogOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *IntegrationLog {
		return vs[0].(map[string]*IntegrationLog)[vs[1].(string)]
	}).(IntegrationLogOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*IntegrationLogInput)(nil)).Elem(), &IntegrationLog{})
	pulumi.RegisterInputType(reflect.TypeOf((*IntegrationLogArrayInput)(nil)).Elem(), IntegrationLogArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*IntegrationLogMapInput)(nil)).Elem(), IntegrationLogMap{})
	pulumi.RegisterOutputType(IntegrationLogOutput{})
	pulumi.RegisterOutputType(IntegrationLogArrayOutput{})
	pulumi.RegisterOutputType(IntegrationLogMapOutput{})
}
