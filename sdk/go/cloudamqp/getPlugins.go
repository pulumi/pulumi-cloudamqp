// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package cloudamqp

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

func GetPlugins(ctx *pulumi.Context, args *GetPluginsArgs, opts ...pulumi.InvokeOption) (*GetPluginsResult, error) {
	var rv GetPluginsResult
	err := ctx.Invoke("cloudamqp:index/getPlugins:getPlugins", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getPlugins.
type GetPluginsArgs struct {
	InstanceId int `pulumi:"instanceId"`
	Plugins []GetPluginsPlugin `pulumi:"plugins"`
}


// A collection of values returned by getPlugins.
type GetPluginsResult struct {
	// id is the provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	InstanceId int `pulumi:"instanceId"`
	Plugins []GetPluginsPlugin `pulumi:"plugins"`
}

