module github.com/pulumi/pulumi-cloudamqp

go 1.13

replace (
	github.com/Azure/go-autorest => github.com/Azure/go-autorest v12.4.3+incompatible
	github.com/hashicorp/vault => github.com/hashicorp/vault v1.2.0
)

require (
	github.com/cloudamqp/terraform-provider-cloudamqp v1.5.0
	github.com/pkg/errors v0.9.1
	github.com/pulumi/pulumi v1.12.2-0.20200313044354-8111d33438b9
	github.com/pulumi/pulumi-terraform-bridge v1.8.2
)
