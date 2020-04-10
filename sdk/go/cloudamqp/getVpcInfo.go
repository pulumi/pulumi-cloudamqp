// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package cloudamqp

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func GetVpcInfo(ctx *pulumi.Context, args *GetVpcInfoArgs, opts ...pulumi.InvokeOption) (*GetVpcInfoResult, error) {
	var rv GetVpcInfoResult
	err := ctx.Invoke("cloudamqp:index/getVpcInfo:getVpcInfo", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getVpcInfo.
type GetVpcInfoArgs struct {
	InstanceId int `pulumi:"instanceId"`
}

// A collection of values returned by getVpcInfo.
type GetVpcInfoResult struct {
	// id is the provider-assigned unique ID for this managed resource.
	Id              string `pulumi:"id"`
	InstanceId      int    `pulumi:"instanceId"`
	Name            string `pulumi:"name"`
	OwnerId         string `pulumi:"ownerId"`
	SecurityGroupId string `pulumi:"securityGroupId"`
	VpcSubnet       string `pulumi:"vpcSubnet"`
}
