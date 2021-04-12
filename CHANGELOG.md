CHANGELOG
=========

## HEAD (Unreleased)
_(none)_

---

## 2.9.0 (2021-04-12)
* Upgrade to pulumi-terraform-bridge v2.23.0

## 2.8.1 (2021-03-23)
* Upgrade to pulumi-terraform-bridge v2.22.1  
  **Please Note:** This includes a bug fix to the refresh operation regarding arrays

## 2.8.0 (2021-03-16)
* Upgrade to pulumi-terraform-bridge v2.21.0
* Release macOS arm64 binary

## 2.7.2 (2021-02-16)
* Upgrade to pulumi-terraform-bridge v2.19.0  
  **Please Note:** This includes a bug fix that stops mutating resources options in the nodejs provider
* Avoid storing config from the environment into the state

## 2.7.1 (2021-02-08)
* Upgrade to v1.9.1 of the CloudAMQP Terraform Provider

## 2.7.0 (2021-02-03)
* Upgrade to v1.9.0 of the CloudAMQP Terraform Provider

## 2.6.0 (2021-01-29)
* Upgrading pulumi-terraform-bridge to v2.18.1

## 2.5.4 (2021-01-13)
* Upgrade to pulumi-terraform-bridge v2.17.0
* Upgrade to Pulumi v2.17.0

## 2.5.3 (2020-12-10)
* Upgrade to v1.8.6 of the CloudAMQP Terraform Provider

## 2.5.2 (2020-12-04)
* Upgrade to v1.8.4 of the CloudAMQP Terraform Provider

## 2.5.1 (2020-11-24)
* Upgrade to pulumi-terraform-bridge v2.13.2  
  * This adds support for import specific examples in documentation

## 2.5.0 (2020-10-26)
* Upgrade to Pulumi v2.12.0 and pulumi-terraform-bridge v2.11.0
* Improving the accuracy of previews leading to a more accurate understanding of what will actually change rather than assuming all output properties will change.  
  ** PLEASE NOTE:**  
  This new preview functionality can be disabled by setting `PULUMI_DISABLE_PROVIDER_PREVIEW` to `1` or `false`.

## 2.4.0 (2020-09-04)
* Upgrade to v1.8.1 of the CloudAMQP Terraform Provider

## 2.3.0 (2020-08-31)
* Upgrade to v1.7.3 of the CloudAMQP Terraform Provider
* Upgrade to pulumi-terraform-bridge v2.7.3
* Upgrade to Pulumi v2.9.0, which adds type annotations and input/output classes to Python

## 2.2.3 (2020-06-12)
* Switch to GitHub actions for build

## 2.2.2 (2020-05-28)
* Upgrade to Pulumi v2.3.0
* Upgrade to pulumi-terraform-bridge v2.4.0

## 2.2.1 (2020-05-12)
* Upgrade to pulumi-terraform-bridge v2.3.1

## 2.2.0 (2020-04-29)
* Upgrade to v1.6.0 of the CloudAMQP Terraform Provider

## 2.1.0 (2020-04-28)
* Upgrade to pulumi-terraform-bridge v2.1.0

## 2.0.0 (2020-04-18)
* Upgrade to Pulumi v2.0.0
* Upgrade to pulumi-terraform-bridge v2.0.0

## 1.6.0 (2020-04-03)
* Upgrade to Pulumi v1.13.1
* Upgrade to pulumi-terraform-bridge v1.8.4
* Refactor layout to support Go modules

## 1.5.0 (2020-03-23)
* Upgrade to Pulumi v1.12.1
* Upgrade to pulumi-terraform-bridge v1.8.2

## 1.4.0 (2020-03-10)
* Upgrade to v1.5.0 of the CloudAMQP Terraform Provider

## 1.3.1 (2020-03-06)
* Upgrade to v1.4.1 of the CloudAMQP Terraform Provider

## 1.3.0 (2020-03-04)
* Upgrade to v1.4.0 of the CloudAMQP Terraform Provider

## 1.2.0 (2020-01-19)
* Upgrade to v1.3.0 of the CloudAMQP Terraform Provider

## 1.1.0 (2019-12-17)
* Namespace names in .NET SDK are adjusted to PascalCase
([#2](https://github.com/pulumi/pulumi-cloudamqp/pull/2)).
* Upgrade to v1.2.2 of the CloudAMQP Terraform Provider

## 1.0.0 (2019-12-06)
* Initial release of the provider
