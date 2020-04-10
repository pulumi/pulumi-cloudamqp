module github.com/pulumi/pulumi-cloudamqp/provider/v2

go 1.13

replace (
	github.com/Azure/go-autorest => github.com/Azure/go-autorest v12.4.3+incompatible
	github.com/hashicorp/vault => github.com/hashicorp/vault v1.2.0
)

require (
	github.com/cloudamqp/terraform-provider-cloudamqp v1.5.0
	github.com/pulumi/pulumi-terraform-bridge/v2 v2.0.0-beta.1
	github.com/pulumi/pulumi/sdk/v2 v2.0.0-beta.3
)
