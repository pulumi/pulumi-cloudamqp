module github.com/pulumi/pulumi-cloudamqp

go 1.13

require (
	cloud.google.com/go/logging v1.0.0 // indirect
	github.com/cloudamqp/terraform-provider-cloudamqp v1.3.0
	github.com/pkg/errors v0.8.1
	github.com/pulumi/pulumi v1.6.1
	github.com/pulumi/pulumi-terraform v0.18.4-0.20191202134852-87cfb4dc8ae1
)

replace (
	github.com/Azure/go-autorest => github.com/Azure/go-autorest v12.4.3+incompatible
	github.com/hashicorp/vault => github.com/hashicorp/vault v1.2.0
)
