[![Actions Status](https://github.com/pulumi/pulumi-cloudamqp/workflows/master/badge.svg)](https://github.com/pulumi/pulumi-cloudamqp/actions)
[![Slack](http://www.pulumi.com/images/docs/badges/slack.svg)](https://slack.pulumi.com)
[![NPM version](https://badge.fury.io/js/%40pulumi%2Fcloudamqp.svg)](https://www.npmjs.com/package/@pulumi/cloudamqp)
[![Python version](https://badge.fury.io/py/pulumi-cloudamqp.svg)](https://pypi.org/project/pulumi-cloudamqp)
[![NuGet version](https://badge.fury.io/nu/pulumi.cloudamqp.svg)](https://badge.fury.io/nu/pulumi.cloudamqp)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/pulumi/pulumi-cloudamqp/sdk/v2/go)](https://pkg.go.dev/github.com/pulumi/pulumi-cloudamqp/sdk/v2/go)
[![License](https://img.shields.io/npm/l/%40pulumi%2Fpulumi.svg)](https://github.com/pulumi/pulumi-cloudamqp/blob/master/LICENSE)

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

    $ go get github.com/pulumi/pulumi-cloudamqp/sdk/v2/go/...

### .NET

To use from .NET, install using `dotnet add package`:

    $ dotnet add package Pulumi.cloudamqp

## Configuration

The following configuration points are available:

- `cloudamqp:apikey` - (Required) Key used to authentication to the CloudAMQP Customer API. May be set via the `CLOUDAMQP_APIKEY` environment variable.
- `cloudamqp:baseurl` - (Optional) Base URL to CloudAMQP Customer website. Default is `https://customer.cloudamqp.com`.

## Reference

For further information, please visit [the CloudAMQP provider docs](https://www.pulumi.com/docs/intro/cloud-providers/cloudamqp) or for detailed reference documentation, please visit [the API docs](https://www.pulumi.com/docs/reference/pkg/cloudamp).
