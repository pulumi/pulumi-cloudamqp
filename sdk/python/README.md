[![Build Status](https://travis-ci.com/pulumi/pulumi-cloudamqp.svg?token=eHg7Zp5zdDDJfTjY8ejq&branch=master)](https://travis-ci.com/pulumi/pulumi-cloudamqp)

# CloudAMQP Resource Provider

The CloudAMQP resource provider for Pulumi lets you manage CloudAMQP resources in your cloud programs. To use
this package, please [install the Pulumi CLI first](https://pulumi.io/).

## Installing

This package is available in many languages in the standard packaging formats.

### Node.js (Java/TypeScript)

To use from JavaScript or TypeScript in Node.js, install using either `npm`:

    $ npm install @pulumi/cloudamqp

or `yarn`:

    $ yarn add @pulumi/cloudamqp

### Python

To use from Python, install using `pip`:

    $ pip install pulumi_cloudamqp

### Go

To use from Go, use `go get` to grab the latest version of the library

    $ go get github.com/pulumi/pulumi-cloudamqp/sdk/go/...

## Configuration

The following configuration points are available:

- `cloudamqp:apikey` - (Required) Key used to authentication to the CloudAMQP Customer API. May be set via the `CLOUDAMQP_APIKEY` environment variable.
- `cloudamqp:baseurl` - (Optional) Base URL to CloudAMQP Customer website. Default is `https://customer.cloudamqp.com`.

## Reference

For detailed reference documentation, please visit [the API docs](https://pulumi.io/reference/pkg/nodejs/@pulumi/cloudamqp/index.html).
