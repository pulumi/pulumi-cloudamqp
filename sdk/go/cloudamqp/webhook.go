// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package cloudamqp

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This resource allows you to enable or disable webhooks for a specific vhost and queue.
//
// Only available for dedicated subscription plans.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-cloudamqp/sdk/v3/go/cloudamqp"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := cloudamqp.NewWebhook(ctx, "webhookQueue", &cloudamqp.WebhookArgs{
// 			InstanceId:    pulumi.Any(cloudamqp_instance.Instance.Id),
// 			Vhost:         pulumi.String("myvhost"),
// 			Queue:         pulumi.String("webhook-queue"),
// 			WebhookUri:    pulumi.String("https://example.com/webhook?key=secret"),
// 			RetryInterval: pulumi.Int(5),
// 			Concurrency:   pulumi.Int(5),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ## Dependency
//
// This resource depends on CloudAMQP instance identifier, `cloudamqp_instance.instance.id`.
//
// ## Import
//
// `cloudamqp_webhook` can be imported using the resource identifier together with CloudAMQP instance identifier. The identifiers are CSV separated, see example below.
//
// ```sh
//  $ pulumi import cloudamqp:index/webhook:Webhook webhook_queue <webhook_id>,<instance_id>`
// ```
type Webhook struct {
	pulumi.CustomResourceState

	// Max simultaneous requests to the endpoint.
	Concurrency pulumi.IntOutput `pulumi:"concurrency"`
	// The CloudAMQP instance ID.
	InstanceId pulumi.IntOutput `pulumi:"instanceId"`
	// A (durable) queue on your RabbitMQ instance.
	Queue pulumi.StringOutput `pulumi:"queue"`
	// How often we retry if your endpoint fails (in seconds).
	RetryInterval pulumi.IntOutput `pulumi:"retryInterval"`
	// The vhost the queue resides in.
	Vhost pulumi.StringOutput `pulumi:"vhost"`
	// A POST request will be made for each message in the queue to this endpoint.
	WebhookUri pulumi.StringOutput `pulumi:"webhookUri"`
}

// NewWebhook registers a new resource with the given unique name, arguments, and options.
func NewWebhook(ctx *pulumi.Context,
	name string, args *WebhookArgs, opts ...pulumi.ResourceOption) (*Webhook, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Concurrency == nil {
		return nil, errors.New("invalid value for required argument 'Concurrency'")
	}
	if args.InstanceId == nil {
		return nil, errors.New("invalid value for required argument 'InstanceId'")
	}
	if args.Queue == nil {
		return nil, errors.New("invalid value for required argument 'Queue'")
	}
	if args.RetryInterval == nil {
		return nil, errors.New("invalid value for required argument 'RetryInterval'")
	}
	if args.Vhost == nil {
		return nil, errors.New("invalid value for required argument 'Vhost'")
	}
	if args.WebhookUri == nil {
		return nil, errors.New("invalid value for required argument 'WebhookUri'")
	}
	var resource Webhook
	err := ctx.RegisterResource("cloudamqp:index/webhook:Webhook", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWebhook gets an existing Webhook resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWebhook(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WebhookState, opts ...pulumi.ResourceOption) (*Webhook, error) {
	var resource Webhook
	err := ctx.ReadResource("cloudamqp:index/webhook:Webhook", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Webhook resources.
type webhookState struct {
	// Max simultaneous requests to the endpoint.
	Concurrency *int `pulumi:"concurrency"`
	// The CloudAMQP instance ID.
	InstanceId *int `pulumi:"instanceId"`
	// A (durable) queue on your RabbitMQ instance.
	Queue *string `pulumi:"queue"`
	// How often we retry if your endpoint fails (in seconds).
	RetryInterval *int `pulumi:"retryInterval"`
	// The vhost the queue resides in.
	Vhost *string `pulumi:"vhost"`
	// A POST request will be made for each message in the queue to this endpoint.
	WebhookUri *string `pulumi:"webhookUri"`
}

type WebhookState struct {
	// Max simultaneous requests to the endpoint.
	Concurrency pulumi.IntPtrInput
	// The CloudAMQP instance ID.
	InstanceId pulumi.IntPtrInput
	// A (durable) queue on your RabbitMQ instance.
	Queue pulumi.StringPtrInput
	// How often we retry if your endpoint fails (in seconds).
	RetryInterval pulumi.IntPtrInput
	// The vhost the queue resides in.
	Vhost pulumi.StringPtrInput
	// A POST request will be made for each message in the queue to this endpoint.
	WebhookUri pulumi.StringPtrInput
}

func (WebhookState) ElementType() reflect.Type {
	return reflect.TypeOf((*webhookState)(nil)).Elem()
}

type webhookArgs struct {
	// Max simultaneous requests to the endpoint.
	Concurrency int `pulumi:"concurrency"`
	// The CloudAMQP instance ID.
	InstanceId int `pulumi:"instanceId"`
	// A (durable) queue on your RabbitMQ instance.
	Queue string `pulumi:"queue"`
	// How often we retry if your endpoint fails (in seconds).
	RetryInterval int `pulumi:"retryInterval"`
	// The vhost the queue resides in.
	Vhost string `pulumi:"vhost"`
	// A POST request will be made for each message in the queue to this endpoint.
	WebhookUri string `pulumi:"webhookUri"`
}

// The set of arguments for constructing a Webhook resource.
type WebhookArgs struct {
	// Max simultaneous requests to the endpoint.
	Concurrency pulumi.IntInput
	// The CloudAMQP instance ID.
	InstanceId pulumi.IntInput
	// A (durable) queue on your RabbitMQ instance.
	Queue pulumi.StringInput
	// How often we retry if your endpoint fails (in seconds).
	RetryInterval pulumi.IntInput
	// The vhost the queue resides in.
	Vhost pulumi.StringInput
	// A POST request will be made for each message in the queue to this endpoint.
	WebhookUri pulumi.StringInput
}

func (WebhookArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*webhookArgs)(nil)).Elem()
}

type WebhookInput interface {
	pulumi.Input

	ToWebhookOutput() WebhookOutput
	ToWebhookOutputWithContext(ctx context.Context) WebhookOutput
}

func (*Webhook) ElementType() reflect.Type {
	return reflect.TypeOf((*Webhook)(nil))
}

func (i *Webhook) ToWebhookOutput() WebhookOutput {
	return i.ToWebhookOutputWithContext(context.Background())
}

func (i *Webhook) ToWebhookOutputWithContext(ctx context.Context) WebhookOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebhookOutput)
}

func (i *Webhook) ToWebhookPtrOutput() WebhookPtrOutput {
	return i.ToWebhookPtrOutputWithContext(context.Background())
}

func (i *Webhook) ToWebhookPtrOutputWithContext(ctx context.Context) WebhookPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebhookPtrOutput)
}

type WebhookPtrInput interface {
	pulumi.Input

	ToWebhookPtrOutput() WebhookPtrOutput
	ToWebhookPtrOutputWithContext(ctx context.Context) WebhookPtrOutput
}

type webhookPtrType WebhookArgs

func (*webhookPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Webhook)(nil))
}

func (i *webhookPtrType) ToWebhookPtrOutput() WebhookPtrOutput {
	return i.ToWebhookPtrOutputWithContext(context.Background())
}

func (i *webhookPtrType) ToWebhookPtrOutputWithContext(ctx context.Context) WebhookPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebhookPtrOutput)
}

// WebhookArrayInput is an input type that accepts WebhookArray and WebhookArrayOutput values.
// You can construct a concrete instance of `WebhookArrayInput` via:
//
//          WebhookArray{ WebhookArgs{...} }
type WebhookArrayInput interface {
	pulumi.Input

	ToWebhookArrayOutput() WebhookArrayOutput
	ToWebhookArrayOutputWithContext(context.Context) WebhookArrayOutput
}

type WebhookArray []WebhookInput

func (WebhookArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*Webhook)(nil))
}

func (i WebhookArray) ToWebhookArrayOutput() WebhookArrayOutput {
	return i.ToWebhookArrayOutputWithContext(context.Background())
}

func (i WebhookArray) ToWebhookArrayOutputWithContext(ctx context.Context) WebhookArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebhookArrayOutput)
}

// WebhookMapInput is an input type that accepts WebhookMap and WebhookMapOutput values.
// You can construct a concrete instance of `WebhookMapInput` via:
//
//          WebhookMap{ "key": WebhookArgs{...} }
type WebhookMapInput interface {
	pulumi.Input

	ToWebhookMapOutput() WebhookMapOutput
	ToWebhookMapOutputWithContext(context.Context) WebhookMapOutput
}

type WebhookMap map[string]WebhookInput

func (WebhookMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*Webhook)(nil))
}

func (i WebhookMap) ToWebhookMapOutput() WebhookMapOutput {
	return i.ToWebhookMapOutputWithContext(context.Background())
}

func (i WebhookMap) ToWebhookMapOutputWithContext(ctx context.Context) WebhookMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebhookMapOutput)
}

type WebhookOutput struct {
	*pulumi.OutputState
}

func (WebhookOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Webhook)(nil))
}

func (o WebhookOutput) ToWebhookOutput() WebhookOutput {
	return o
}

func (o WebhookOutput) ToWebhookOutputWithContext(ctx context.Context) WebhookOutput {
	return o
}

func (o WebhookOutput) ToWebhookPtrOutput() WebhookPtrOutput {
	return o.ToWebhookPtrOutputWithContext(context.Background())
}

func (o WebhookOutput) ToWebhookPtrOutputWithContext(ctx context.Context) WebhookPtrOutput {
	return o.ApplyT(func(v Webhook) *Webhook {
		return &v
	}).(WebhookPtrOutput)
}

type WebhookPtrOutput struct {
	*pulumi.OutputState
}

func (WebhookPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Webhook)(nil))
}

func (o WebhookPtrOutput) ToWebhookPtrOutput() WebhookPtrOutput {
	return o
}

func (o WebhookPtrOutput) ToWebhookPtrOutputWithContext(ctx context.Context) WebhookPtrOutput {
	return o
}

type WebhookArrayOutput struct{ *pulumi.OutputState }

func (WebhookArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Webhook)(nil))
}

func (o WebhookArrayOutput) ToWebhookArrayOutput() WebhookArrayOutput {
	return o
}

func (o WebhookArrayOutput) ToWebhookArrayOutputWithContext(ctx context.Context) WebhookArrayOutput {
	return o
}

func (o WebhookArrayOutput) Index(i pulumi.IntInput) WebhookOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Webhook {
		return vs[0].([]Webhook)[vs[1].(int)]
	}).(WebhookOutput)
}

type WebhookMapOutput struct{ *pulumi.OutputState }

func (WebhookMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]Webhook)(nil))
}

func (o WebhookMapOutput) ToWebhookMapOutput() WebhookMapOutput {
	return o
}

func (o WebhookMapOutput) ToWebhookMapOutputWithContext(ctx context.Context) WebhookMapOutput {
	return o
}

func (o WebhookMapOutput) MapIndex(k pulumi.StringInput) WebhookOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) Webhook {
		return vs[0].(map[string]Webhook)[vs[1].(string)]
	}).(WebhookOutput)
}

func init() {
	pulumi.RegisterOutputType(WebhookOutput{})
	pulumi.RegisterOutputType(WebhookPtrOutput{})
	pulumi.RegisterOutputType(WebhookArrayOutput{})
	pulumi.RegisterOutputType(WebhookMapOutput{})
}
