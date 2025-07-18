// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package config

import (
	"github.com/pulumi/pulumi-cloudamqp/sdk/v3/go/cloudamqp/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
)

var _ = internal.GetEnvOrDefault

// Key used to authentication to the CloudAMQP Customer API
func GetApikey(ctx *pulumi.Context) string {
	return config.Get(ctx, "cloudamqp:apikey")
}

// Base URL to CloudAMQP Customer website
func GetBaseurl(ctx *pulumi.Context) string {
	return config.Get(ctx, "cloudamqp:baseurl")
}
func GetEnableFasterInstanceDestroy(ctx *pulumi.Context) bool {
	return config.GetBool(ctx, "cloudamqp:enableFasterInstanceDestroy")
}
