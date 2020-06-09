module github.com/pulumi/pulumi-cloudamqp/provider/v2

go 1.14

replace (
	github.com/Azure/go-autorest => github.com/Azure/go-autorest v12.4.3+incompatible
	github.com/hashicorp/vault => github.com/hashicorp/vault v1.2.0
)

require (
	github.com/cloudamqp/terraform-provider-cloudamqp v1.6.0
	github.com/pulumi/pulumi-terraform-bridge/v2 v2.4.1-0.20200608011815-6feeb51f2d39
	github.com/pulumi/pulumi/sdk/v2 v2.3.1-0.20200607162109-9754465b04db
	github.com/pulumi/tf2pulumi v0.8.1-0.20200528170746-c1234defe2b5 // indirect
)
