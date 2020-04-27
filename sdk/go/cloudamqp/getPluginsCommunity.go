// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package cloudamqp

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

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
	InstanceId int                         `pulumi:"instanceId"`
	Plugins    []GetPluginsCommunityPlugin `pulumi:"plugins"`
}

// A collection of values returned by getPluginsCommunity.
type GetPluginsCommunityResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id         string                      `pulumi:"id"`
	InstanceId int                         `pulumi:"instanceId"`
	Plugins    []GetPluginsCommunityPlugin `pulumi:"plugins"`
}
