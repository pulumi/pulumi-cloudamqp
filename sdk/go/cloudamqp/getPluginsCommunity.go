// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package cloudamqp

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to retrieve information about available community plugins for the CloudAMQP instance.
//
// ⚠️  From our go API wrapper [v1.5.0](https://github.com/84codes/go-api/releases/tag/v1.5.0) there is support for multiple retries when requesting information about community plugins. This was introduced to avoid `ReadPluginCommunity error 400: Timeout talking to backend`.
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
// 		_, err := cloudamqp.GetPluginsCommunity(ctx, &cloudamqp.GetPluginsCommunityArgs{
// 			InstanceId: cloudamqp_instance.Instance.Id,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ## Argument reference
//
// * `instanceId` - (Required) The CloudAMQP instance identifier.
//
// ## Attributes reference
//
// All attributes reference are computed
//
// * `id`      - The identifier for this resource.
// * `plugins` - An array of community plugins. Each `plugins` block consists of the fields documented below.
//
// ***
//
// The `plugins` block consists of
//
// * `name`        - The type of the recipient.
// * `require`     - Min. required Rabbit MQ version to be used.
// * `description` - Description of what the plugin does.
//
// ## Dependency
//
// This data source depends on CloudAMQP instance identifier, `cloudamqp_instance.instance.id`.
func GetPluginsCommunity(ctx *pulumi.Context, args *GetPluginsCommunityArgs, opts ...pulumi.InvokeOption) (*GetPluginsCommunityResult, error) {
	var rv GetPluginsCommunityResult
	err := ctx.Invoke("cloudamqp:index/getPluginsCommunity:getPluginsCommunity", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getPluginsCommunity.
type GetPluginsCommunityArgs struct {
	InstanceId int `pulumi:"instanceId"`
}

// A collection of values returned by getPluginsCommunity.
type GetPluginsCommunityResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id         string                      `pulumi:"id"`
	InstanceId int                         `pulumi:"instanceId"`
	Plugins    []GetPluginsCommunityPlugin `pulumi:"plugins"`
}
