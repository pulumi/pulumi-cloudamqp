module github.com/pulumi/pulumi-cloudamqp/provider/v3

go 1.16

replace github.com/hashicorp/vault => github.com/hashicorp/vault v1.2.0

require (
	github.com/cloudamqp/terraform-provider-cloudamqp v1.9.1
	github.com/pulumi/pulumi-terraform-bridge/v3 v3.0.0
	github.com/pulumi/pulumi/pkg/v3 v3.0.0
	github.com/pulumi/pulumi/sdk/v3 v3.0.0
)
