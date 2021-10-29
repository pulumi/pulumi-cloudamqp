// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package cloudamqp

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to retrieve information about an already created CloudAMQP instance. In order to retrieve the correct information, the CoudAMQP instance identifier is needed.
//
// ## Argument reference
//
// * `instanceId` - (Required) The CloudAMQP instance identifier.
//
// ## Attributes reference
//
// All attributes reference are computed
//
// * `id`          - The identifier for this resource.
// * `name`        - The name of the CloudAMQP instance.
// * `plan`        - The subscription plan for the CloudAMQP instance.
// * `region`      - The cloud platform and region that host the CloudAMQP instance, `{platform}::{region}`.
// * `vpcId`      - ID of the VPC configured for the CloudAMQP instance.
// * `vpcSubnet`  - Dedicated VPC subnet configured for the CloudAMQP instance.
// * `nodes`       - Number of nodes in the cluster of the CloudAMQP instance.
// * `rmqVersion` - The version of installed Rabbit MQ.
// * `url`         - (Sensitive) The AMQP url, used by clients to connect for pub/sub.
// * `apikey`      - (Sensitive) The API key to secondary API handing alarms, integration etc.
// * `tags`        - Tags the CloudAMQP instance with categories.
// * `host`        - The hostname for the CloudAMQP instance.
// * `vhost`       - The virtual host configured in Rabbit MQ.
func LookupInstance(ctx *pulumi.Context, args *LookupInstanceArgs, opts ...pulumi.InvokeOption) (*LookupInstanceResult, error) {
	var rv LookupInstanceResult
	err := ctx.Invoke("cloudamqp:index/getInstance:getInstance", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getInstance.
type LookupInstanceArgs struct {
	InstanceId int `pulumi:"instanceId"`
}

// A collection of values returned by getInstance.
type LookupInstanceResult struct {
	Apikey    string `pulumi:"apikey"`
	Dedicated bool   `pulumi:"dedicated"`
	Host      string `pulumi:"host"`
	// The provider-assigned unique ID for this managed resource.
	Id              string   `pulumi:"id"`
	InstanceId      int      `pulumi:"instanceId"`
	Name            string   `pulumi:"name"`
	NoDefaultAlarms bool     `pulumi:"noDefaultAlarms"`
	Nodes           int      `pulumi:"nodes"`
	Plan            string   `pulumi:"plan"`
	Ready           bool     `pulumi:"ready"`
	Region          string   `pulumi:"region"`
	RmqVersion      string   `pulumi:"rmqVersion"`
	Tags            []string `pulumi:"tags"`
	Url             string   `pulumi:"url"`
	Vhost           string   `pulumi:"vhost"`
	VpcId           int      `pulumi:"vpcId"`
	VpcSubnet       string   `pulumi:"vpcSubnet"`
}
