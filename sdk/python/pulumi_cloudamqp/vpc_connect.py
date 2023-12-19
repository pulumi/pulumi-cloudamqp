# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['VpcConnectArgs', 'VpcConnect']

@pulumi.input_type
class VpcConnectArgs:
    def __init__(__self__, *,
                 instance_id: pulumi.Input[int],
                 region: pulumi.Input[str],
                 allowed_principals: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 allowed_projects: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 approved_subscriptions: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 sleep: Optional[pulumi.Input[int]] = None,
                 timeout: Optional[pulumi.Input[int]] = None):
        """
        The set of arguments for constructing a VpcConnect resource.
        :param pulumi.Input[int] instance_id: The CloudAMQP instance identifier.
        :param pulumi.Input[str] region: The region where the CloudAMQP instance is hosted.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] allowed_principals: List of allowed prinicpals used by AWS, see below table.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] allowed_projects: List of allowed projects used by GCP, see below table.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] approved_subscriptions: List of approved subscriptions used by Azure, see below table.
        :param pulumi.Input[int] sleep: Configurable sleep time (seconds) when enable Private Service Connect.
               Default set to 10 seconds.
        :param pulumi.Input[int] timeout: Configurable timeout time (seconds) when enable Private Service Connect.
               Default set to 1800 seconds.
               
               ___
               
               The `allowed_principals`, `approved_subscriptions` or `allowed_projects` data depends on the provider platform:
               
               | Platform | Description         | Format                                                                                                                             |
               |----------|---------------------|------------------------------------------------------------------------------------------------------------------------------------|
               | AWS      | IAM ARN principals  | arn:aws:iam::aws-account-id:root<br /> arn:aws:iam::aws-account-id:user/user-name<br /> arn:aws:iam::aws-account-id:role/role-name |
               | Azure    | Subscription (GUID) | XXXXXXXX-XXXX-XXXX-XXXX-XXXXXXXXXXXX                                                                                               |
               | GCP      | Project IDs*        | 6 to 30 lowercase letters, digits, or hyphens                                                                                      |
               
               *https://cloud.google.com/resource-manager/reference/rest/v1/projects
        """
        pulumi.set(__self__, "instance_id", instance_id)
        pulumi.set(__self__, "region", region)
        if allowed_principals is not None:
            pulumi.set(__self__, "allowed_principals", allowed_principals)
        if allowed_projects is not None:
            pulumi.set(__self__, "allowed_projects", allowed_projects)
        if approved_subscriptions is not None:
            pulumi.set(__self__, "approved_subscriptions", approved_subscriptions)
        if sleep is not None:
            pulumi.set(__self__, "sleep", sleep)
        if timeout is not None:
            pulumi.set(__self__, "timeout", timeout)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> pulumi.Input[int]:
        """
        The CloudAMQP instance identifier.
        """
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: pulumi.Input[int]):
        pulumi.set(self, "instance_id", value)

    @property
    @pulumi.getter
    def region(self) -> pulumi.Input[str]:
        """
        The region where the CloudAMQP instance is hosted.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: pulumi.Input[str]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter(name="allowedPrincipals")
    def allowed_principals(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        List of allowed prinicpals used by AWS, see below table.
        """
        return pulumi.get(self, "allowed_principals")

    @allowed_principals.setter
    def allowed_principals(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "allowed_principals", value)

    @property
    @pulumi.getter(name="allowedProjects")
    def allowed_projects(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        List of allowed projects used by GCP, see below table.
        """
        return pulumi.get(self, "allowed_projects")

    @allowed_projects.setter
    def allowed_projects(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "allowed_projects", value)

    @property
    @pulumi.getter(name="approvedSubscriptions")
    def approved_subscriptions(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        List of approved subscriptions used by Azure, see below table.
        """
        return pulumi.get(self, "approved_subscriptions")

    @approved_subscriptions.setter
    def approved_subscriptions(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "approved_subscriptions", value)

    @property
    @pulumi.getter
    def sleep(self) -> Optional[pulumi.Input[int]]:
        """
        Configurable sleep time (seconds) when enable Private Service Connect.
        Default set to 10 seconds.
        """
        return pulumi.get(self, "sleep")

    @sleep.setter
    def sleep(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "sleep", value)

    @property
    @pulumi.getter
    def timeout(self) -> Optional[pulumi.Input[int]]:
        """
        Configurable timeout time (seconds) when enable Private Service Connect.
        Default set to 1800 seconds.

        ___

        The `allowed_principals`, `approved_subscriptions` or `allowed_projects` data depends on the provider platform:

        | Platform | Description         | Format                                                                                                                             |
        |----------|---------------------|------------------------------------------------------------------------------------------------------------------------------------|
        | AWS      | IAM ARN principals  | arn:aws:iam::aws-account-id:root<br /> arn:aws:iam::aws-account-id:user/user-name<br /> arn:aws:iam::aws-account-id:role/role-name |
        | Azure    | Subscription (GUID) | XXXXXXXX-XXXX-XXXX-XXXX-XXXXXXXXXXXX                                                                                               |
        | GCP      | Project IDs*        | 6 to 30 lowercase letters, digits, or hyphens                                                                                      |

        *https://cloud.google.com/resource-manager/reference/rest/v1/projects
        """
        return pulumi.get(self, "timeout")

    @timeout.setter
    def timeout(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "timeout", value)


@pulumi.input_type
class _VpcConnectState:
    def __init__(__self__, *,
                 active_zones: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 allowed_principals: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 allowed_projects: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 approved_subscriptions: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 instance_id: Optional[pulumi.Input[int]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 service_name: Optional[pulumi.Input[str]] = None,
                 sleep: Optional[pulumi.Input[int]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 timeout: Optional[pulumi.Input[int]] = None):
        """
        Input properties used for looking up and filtering VpcConnect resources.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] active_zones: Covering availability zones used when creating an endpoint from other VPC. (AWS)
        :param pulumi.Input[Sequence[pulumi.Input[str]]] allowed_principals: List of allowed prinicpals used by AWS, see below table.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] allowed_projects: List of allowed projects used by GCP, see below table.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] approved_subscriptions: List of approved subscriptions used by Azure, see below table.
        :param pulumi.Input[int] instance_id: The CloudAMQP instance identifier.
        :param pulumi.Input[str] region: The region where the CloudAMQP instance is hosted.
        :param pulumi.Input[str] service_name: Service name (alias for Azure) of the PrivateLink.
        :param pulumi.Input[int] sleep: Configurable sleep time (seconds) when enable Private Service Connect.
               Default set to 10 seconds.
        :param pulumi.Input[str] status: Private Service Connect status [enable, pending, disable]
        :param pulumi.Input[int] timeout: Configurable timeout time (seconds) when enable Private Service Connect.
               Default set to 1800 seconds.
               
               ___
               
               The `allowed_principals`, `approved_subscriptions` or `allowed_projects` data depends on the provider platform:
               
               | Platform | Description         | Format                                                                                                                             |
               |----------|---------------------|------------------------------------------------------------------------------------------------------------------------------------|
               | AWS      | IAM ARN principals  | arn:aws:iam::aws-account-id:root<br /> arn:aws:iam::aws-account-id:user/user-name<br /> arn:aws:iam::aws-account-id:role/role-name |
               | Azure    | Subscription (GUID) | XXXXXXXX-XXXX-XXXX-XXXX-XXXXXXXXXXXX                                                                                               |
               | GCP      | Project IDs*        | 6 to 30 lowercase letters, digits, or hyphens                                                                                      |
               
               *https://cloud.google.com/resource-manager/reference/rest/v1/projects
        """
        if active_zones is not None:
            pulumi.set(__self__, "active_zones", active_zones)
        if allowed_principals is not None:
            pulumi.set(__self__, "allowed_principals", allowed_principals)
        if allowed_projects is not None:
            pulumi.set(__self__, "allowed_projects", allowed_projects)
        if approved_subscriptions is not None:
            pulumi.set(__self__, "approved_subscriptions", approved_subscriptions)
        if instance_id is not None:
            pulumi.set(__self__, "instance_id", instance_id)
        if region is not None:
            pulumi.set(__self__, "region", region)
        if service_name is not None:
            pulumi.set(__self__, "service_name", service_name)
        if sleep is not None:
            pulumi.set(__self__, "sleep", sleep)
        if status is not None:
            pulumi.set(__self__, "status", status)
        if timeout is not None:
            pulumi.set(__self__, "timeout", timeout)

    @property
    @pulumi.getter(name="activeZones")
    def active_zones(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Covering availability zones used when creating an endpoint from other VPC. (AWS)
        """
        return pulumi.get(self, "active_zones")

    @active_zones.setter
    def active_zones(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "active_zones", value)

    @property
    @pulumi.getter(name="allowedPrincipals")
    def allowed_principals(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        List of allowed prinicpals used by AWS, see below table.
        """
        return pulumi.get(self, "allowed_principals")

    @allowed_principals.setter
    def allowed_principals(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "allowed_principals", value)

    @property
    @pulumi.getter(name="allowedProjects")
    def allowed_projects(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        List of allowed projects used by GCP, see below table.
        """
        return pulumi.get(self, "allowed_projects")

    @allowed_projects.setter
    def allowed_projects(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "allowed_projects", value)

    @property
    @pulumi.getter(name="approvedSubscriptions")
    def approved_subscriptions(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        List of approved subscriptions used by Azure, see below table.
        """
        return pulumi.get(self, "approved_subscriptions")

    @approved_subscriptions.setter
    def approved_subscriptions(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "approved_subscriptions", value)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> Optional[pulumi.Input[int]]:
        """
        The CloudAMQP instance identifier.
        """
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "instance_id", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        """
        The region where the CloudAMQP instance is hosted.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter(name="serviceName")
    def service_name(self) -> Optional[pulumi.Input[str]]:
        """
        Service name (alias for Azure) of the PrivateLink.
        """
        return pulumi.get(self, "service_name")

    @service_name.setter
    def service_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "service_name", value)

    @property
    @pulumi.getter
    def sleep(self) -> Optional[pulumi.Input[int]]:
        """
        Configurable sleep time (seconds) when enable Private Service Connect.
        Default set to 10 seconds.
        """
        return pulumi.get(self, "sleep")

    @sleep.setter
    def sleep(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "sleep", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        """
        Private Service Connect status [enable, pending, disable]
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)

    @property
    @pulumi.getter
    def timeout(self) -> Optional[pulumi.Input[int]]:
        """
        Configurable timeout time (seconds) when enable Private Service Connect.
        Default set to 1800 seconds.

        ___

        The `allowed_principals`, `approved_subscriptions` or `allowed_projects` data depends on the provider platform:

        | Platform | Description         | Format                                                                                                                             |
        |----------|---------------------|------------------------------------------------------------------------------------------------------------------------------------|
        | AWS      | IAM ARN principals  | arn:aws:iam::aws-account-id:root<br /> arn:aws:iam::aws-account-id:user/user-name<br /> arn:aws:iam::aws-account-id:role/role-name |
        | Azure    | Subscription (GUID) | XXXXXXXX-XXXX-XXXX-XXXX-XXXXXXXXXXXX                                                                                               |
        | GCP      | Project IDs*        | 6 to 30 lowercase letters, digits, or hyphens                                                                                      |

        *https://cloud.google.com/resource-manager/reference/rest/v1/projects
        """
        return pulumi.get(self, "timeout")

    @timeout.setter
    def timeout(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "timeout", value)


class VpcConnect(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 allowed_principals: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 allowed_projects: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 approved_subscriptions: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 instance_id: Optional[pulumi.Input[int]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 sleep: Optional[pulumi.Input[int]] = None,
                 timeout: Optional[pulumi.Input[int]] = None,
                 __props__=None):
        """
        This resource is a generic way to handle PrivateLink (AWS and Azure) and Private Service Connect (GCP).
        Communication between resources can be done just as they were living inside a VPC. CloudAMQP creates an Endpoint
        Service to connect the VPC and creating a new network interface to handle the communicate.

        If no existing VPC available when enable VPC connect, a new VPC will be created with subnet `10.52.72.0/24`.

        More information can be found at: [CloudAMQP VPC Connect](https://www.cloudamqp.com/docs/cloudamqp-vpc-connect.html)

        > **Note:** Enabling VPC Connect will automatically add a firewall rule.

        <details>
         <summary>
            <b>
              <i>Default PrivateLink firewall rule [AWS, Azure]</i>
            </b>
          </summary>

        ```python
        import pulumi
        ```

        </details>

        <details>
         <summary>
            <b>
              <i>Default Private Service Connect firewall rule [GCP]</i>
            </b>
          </summary>

        ```python
        import pulumi
        ```

        </details>

        Only available for dedicated subscription plans.

        ## Example Usage

        <details>
          <summary>
            <b>
              <i>Enable VPC Connect (PrivateLink) in AWS</i>
            </b>
          </summary>

        ```python
        import pulumi
        import pulumi_cloudamqp as cloudamqp

        vpc = cloudamqp.Vpc("vpc",
            region="amazon-web-services::us-west-1",
            subnet="10.56.72.0/24",
            tags=[])
        instance = cloudamqp.Instance("instance",
            plan="bunny-1",
            region="amazon-web-services::us-west-1",
            tags=[],
            vpc_id=vpc.id,
            keep_associated_vpc=True)
        vpc_connect = cloudamqp.VpcConnect("vpcConnect",
            instance_id=instance.id,
            region=instance.region,
            allowed_principals=["arn:aws:iam::aws-account-id:user/user-name"])
        ```

        </details>

        <details>
          <summary>
            <b>
              <i>Enable VPC Connect (PrivateLink) in Azure</i>
            </b>
          </summary>

        ```python
        import pulumi
        import pulumi_cloudamqp as cloudamqp

        vpc = cloudamqp.Vpc("vpc",
            region="azure-arm::westus",
            subnet="10.56.72.0/24",
            tags=[])
        instance = cloudamqp.Instance("instance",
            plan="bunny-1",
            region="azure-arm::westus",
            tags=[],
            vpc_id=vpc.id,
            keep_associated_vpc=True)
        vpc_connect = cloudamqp.VpcConnect("vpcConnect",
            instance_id=instance.id,
            region=instance.region,
            approved_subscriptions=["XXXXXXXX-XXXX-XXXX-XXXX-XXXXXXXXXXXX"])
        ```

        </details>

        <details>
          <summary>
            <b>
              <i>Enable VPC Connect (Private Service Connect) in GCP</i>
            </b>
          </summary>

        ```python
        import pulumi
        import pulumi_cloudamqp as cloudamqp

        vpc = cloudamqp.Vpc("vpc",
            region="google-compute-engine::us-west1",
            subnet="10.56.72.0/24",
            tags=[])
        instance = cloudamqp.Instance("instance",
            plan="bunny-1",
            region="google-compute-engine::us-west1",
            tags=[],
            vpc_id=vpc.id,
            keep_associated_vpc=True)
        vpc_connect = cloudamqp.VpcConnect("vpcConnect",
            instance_id=instance.id,
            region=instance.region,
            allowed_projects=["some-project-123456"])
        ```

        </details>
        ## Depedency

        This resource depends on CloudAMQP instance identifier, `cloudamqp_instance.instance.id`.

        Since `region` also is required, suggest to reuse the argument from CloudAMQP instance,
        `cloudamqp_instance.instance.region`.

        ## Import

        `cloudamqp_vpc_connect` can be imported using CloudAMQP internal identifier.

        ```sh
         $ pulumi import cloudamqp:index/vpcConnect:VpcConnect vpc_connect <id>`
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] allowed_principals: List of allowed prinicpals used by AWS, see below table.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] allowed_projects: List of allowed projects used by GCP, see below table.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] approved_subscriptions: List of approved subscriptions used by Azure, see below table.
        :param pulumi.Input[int] instance_id: The CloudAMQP instance identifier.
        :param pulumi.Input[str] region: The region where the CloudAMQP instance is hosted.
        :param pulumi.Input[int] sleep: Configurable sleep time (seconds) when enable Private Service Connect.
               Default set to 10 seconds.
        :param pulumi.Input[int] timeout: Configurable timeout time (seconds) when enable Private Service Connect.
               Default set to 1800 seconds.
               
               ___
               
               The `allowed_principals`, `approved_subscriptions` or `allowed_projects` data depends on the provider platform:
               
               | Platform | Description         | Format                                                                                                                             |
               |----------|---------------------|------------------------------------------------------------------------------------------------------------------------------------|
               | AWS      | IAM ARN principals  | arn:aws:iam::aws-account-id:root<br /> arn:aws:iam::aws-account-id:user/user-name<br /> arn:aws:iam::aws-account-id:role/role-name |
               | Azure    | Subscription (GUID) | XXXXXXXX-XXXX-XXXX-XXXX-XXXXXXXXXXXX                                                                                               |
               | GCP      | Project IDs*        | 6 to 30 lowercase letters, digits, or hyphens                                                                                      |
               
               *https://cloud.google.com/resource-manager/reference/rest/v1/projects
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: VpcConnectArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        This resource is a generic way to handle PrivateLink (AWS and Azure) and Private Service Connect (GCP).
        Communication between resources can be done just as they were living inside a VPC. CloudAMQP creates an Endpoint
        Service to connect the VPC and creating a new network interface to handle the communicate.

        If no existing VPC available when enable VPC connect, a new VPC will be created with subnet `10.52.72.0/24`.

        More information can be found at: [CloudAMQP VPC Connect](https://www.cloudamqp.com/docs/cloudamqp-vpc-connect.html)

        > **Note:** Enabling VPC Connect will automatically add a firewall rule.

        <details>
         <summary>
            <b>
              <i>Default PrivateLink firewall rule [AWS, Azure]</i>
            </b>
          </summary>

        ```python
        import pulumi
        ```

        </details>

        <details>
         <summary>
            <b>
              <i>Default Private Service Connect firewall rule [GCP]</i>
            </b>
          </summary>

        ```python
        import pulumi
        ```

        </details>

        Only available for dedicated subscription plans.

        ## Example Usage

        <details>
          <summary>
            <b>
              <i>Enable VPC Connect (PrivateLink) in AWS</i>
            </b>
          </summary>

        ```python
        import pulumi
        import pulumi_cloudamqp as cloudamqp

        vpc = cloudamqp.Vpc("vpc",
            region="amazon-web-services::us-west-1",
            subnet="10.56.72.0/24",
            tags=[])
        instance = cloudamqp.Instance("instance",
            plan="bunny-1",
            region="amazon-web-services::us-west-1",
            tags=[],
            vpc_id=vpc.id,
            keep_associated_vpc=True)
        vpc_connect = cloudamqp.VpcConnect("vpcConnect",
            instance_id=instance.id,
            region=instance.region,
            allowed_principals=["arn:aws:iam::aws-account-id:user/user-name"])
        ```

        </details>

        <details>
          <summary>
            <b>
              <i>Enable VPC Connect (PrivateLink) in Azure</i>
            </b>
          </summary>

        ```python
        import pulumi
        import pulumi_cloudamqp as cloudamqp

        vpc = cloudamqp.Vpc("vpc",
            region="azure-arm::westus",
            subnet="10.56.72.0/24",
            tags=[])
        instance = cloudamqp.Instance("instance",
            plan="bunny-1",
            region="azure-arm::westus",
            tags=[],
            vpc_id=vpc.id,
            keep_associated_vpc=True)
        vpc_connect = cloudamqp.VpcConnect("vpcConnect",
            instance_id=instance.id,
            region=instance.region,
            approved_subscriptions=["XXXXXXXX-XXXX-XXXX-XXXX-XXXXXXXXXXXX"])
        ```

        </details>

        <details>
          <summary>
            <b>
              <i>Enable VPC Connect (Private Service Connect) in GCP</i>
            </b>
          </summary>

        ```python
        import pulumi
        import pulumi_cloudamqp as cloudamqp

        vpc = cloudamqp.Vpc("vpc",
            region="google-compute-engine::us-west1",
            subnet="10.56.72.0/24",
            tags=[])
        instance = cloudamqp.Instance("instance",
            plan="bunny-1",
            region="google-compute-engine::us-west1",
            tags=[],
            vpc_id=vpc.id,
            keep_associated_vpc=True)
        vpc_connect = cloudamqp.VpcConnect("vpcConnect",
            instance_id=instance.id,
            region=instance.region,
            allowed_projects=["some-project-123456"])
        ```

        </details>
        ## Depedency

        This resource depends on CloudAMQP instance identifier, `cloudamqp_instance.instance.id`.

        Since `region` also is required, suggest to reuse the argument from CloudAMQP instance,
        `cloudamqp_instance.instance.region`.

        ## Import

        `cloudamqp_vpc_connect` can be imported using CloudAMQP internal identifier.

        ```sh
         $ pulumi import cloudamqp:index/vpcConnect:VpcConnect vpc_connect <id>`
        ```

        :param str resource_name: The name of the resource.
        :param VpcConnectArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(VpcConnectArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 allowed_principals: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 allowed_projects: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 approved_subscriptions: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 instance_id: Optional[pulumi.Input[int]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 sleep: Optional[pulumi.Input[int]] = None,
                 timeout: Optional[pulumi.Input[int]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = VpcConnectArgs.__new__(VpcConnectArgs)

            __props__.__dict__["allowed_principals"] = allowed_principals
            __props__.__dict__["allowed_projects"] = allowed_projects
            __props__.__dict__["approved_subscriptions"] = approved_subscriptions
            if instance_id is None and not opts.urn:
                raise TypeError("Missing required property 'instance_id'")
            __props__.__dict__["instance_id"] = instance_id
            if region is None and not opts.urn:
                raise TypeError("Missing required property 'region'")
            __props__.__dict__["region"] = region
            __props__.__dict__["sleep"] = sleep
            __props__.__dict__["timeout"] = timeout
            __props__.__dict__["active_zones"] = None
            __props__.__dict__["service_name"] = None
            __props__.__dict__["status"] = None
        super(VpcConnect, __self__).__init__(
            'cloudamqp:index/vpcConnect:VpcConnect',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            active_zones: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            allowed_principals: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            allowed_projects: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            approved_subscriptions: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            instance_id: Optional[pulumi.Input[int]] = None,
            region: Optional[pulumi.Input[str]] = None,
            service_name: Optional[pulumi.Input[str]] = None,
            sleep: Optional[pulumi.Input[int]] = None,
            status: Optional[pulumi.Input[str]] = None,
            timeout: Optional[pulumi.Input[int]] = None) -> 'VpcConnect':
        """
        Get an existing VpcConnect resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] active_zones: Covering availability zones used when creating an endpoint from other VPC. (AWS)
        :param pulumi.Input[Sequence[pulumi.Input[str]]] allowed_principals: List of allowed prinicpals used by AWS, see below table.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] allowed_projects: List of allowed projects used by GCP, see below table.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] approved_subscriptions: List of approved subscriptions used by Azure, see below table.
        :param pulumi.Input[int] instance_id: The CloudAMQP instance identifier.
        :param pulumi.Input[str] region: The region where the CloudAMQP instance is hosted.
        :param pulumi.Input[str] service_name: Service name (alias for Azure) of the PrivateLink.
        :param pulumi.Input[int] sleep: Configurable sleep time (seconds) when enable Private Service Connect.
               Default set to 10 seconds.
        :param pulumi.Input[str] status: Private Service Connect status [enable, pending, disable]
        :param pulumi.Input[int] timeout: Configurable timeout time (seconds) when enable Private Service Connect.
               Default set to 1800 seconds.
               
               ___
               
               The `allowed_principals`, `approved_subscriptions` or `allowed_projects` data depends on the provider platform:
               
               | Platform | Description         | Format                                                                                                                             |
               |----------|---------------------|------------------------------------------------------------------------------------------------------------------------------------|
               | AWS      | IAM ARN principals  | arn:aws:iam::aws-account-id:root<br /> arn:aws:iam::aws-account-id:user/user-name<br /> arn:aws:iam::aws-account-id:role/role-name |
               | Azure    | Subscription (GUID) | XXXXXXXX-XXXX-XXXX-XXXX-XXXXXXXXXXXX                                                                                               |
               | GCP      | Project IDs*        | 6 to 30 lowercase letters, digits, or hyphens                                                                                      |
               
               *https://cloud.google.com/resource-manager/reference/rest/v1/projects
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _VpcConnectState.__new__(_VpcConnectState)

        __props__.__dict__["active_zones"] = active_zones
        __props__.__dict__["allowed_principals"] = allowed_principals
        __props__.__dict__["allowed_projects"] = allowed_projects
        __props__.__dict__["approved_subscriptions"] = approved_subscriptions
        __props__.__dict__["instance_id"] = instance_id
        __props__.__dict__["region"] = region
        __props__.__dict__["service_name"] = service_name
        __props__.__dict__["sleep"] = sleep
        __props__.__dict__["status"] = status
        __props__.__dict__["timeout"] = timeout
        return VpcConnect(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="activeZones")
    def active_zones(self) -> pulumi.Output[Sequence[str]]:
        """
        Covering availability zones used when creating an endpoint from other VPC. (AWS)
        """
        return pulumi.get(self, "active_zones")

    @property
    @pulumi.getter(name="allowedPrincipals")
    def allowed_principals(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        List of allowed prinicpals used by AWS, see below table.
        """
        return pulumi.get(self, "allowed_principals")

    @property
    @pulumi.getter(name="allowedProjects")
    def allowed_projects(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        List of allowed projects used by GCP, see below table.
        """
        return pulumi.get(self, "allowed_projects")

    @property
    @pulumi.getter(name="approvedSubscriptions")
    def approved_subscriptions(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        List of approved subscriptions used by Azure, see below table.
        """
        return pulumi.get(self, "approved_subscriptions")

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> pulumi.Output[int]:
        """
        The CloudAMQP instance identifier.
        """
        return pulumi.get(self, "instance_id")

    @property
    @pulumi.getter
    def region(self) -> pulumi.Output[str]:
        """
        The region where the CloudAMQP instance is hosted.
        """
        return pulumi.get(self, "region")

    @property
    @pulumi.getter(name="serviceName")
    def service_name(self) -> pulumi.Output[str]:
        """
        Service name (alias for Azure) of the PrivateLink.
        """
        return pulumi.get(self, "service_name")

    @property
    @pulumi.getter
    def sleep(self) -> pulumi.Output[Optional[int]]:
        """
        Configurable sleep time (seconds) when enable Private Service Connect.
        Default set to 10 seconds.
        """
        return pulumi.get(self, "sleep")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[str]:
        """
        Private Service Connect status [enable, pending, disable]
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter
    def timeout(self) -> pulumi.Output[Optional[int]]:
        """
        Configurable timeout time (seconds) when enable Private Service Connect.
        Default set to 1800 seconds.

        ___

        The `allowed_principals`, `approved_subscriptions` or `allowed_projects` data depends on the provider platform:

        | Platform | Description         | Format                                                                                                                             |
        |----------|---------------------|------------------------------------------------------------------------------------------------------------------------------------|
        | AWS      | IAM ARN principals  | arn:aws:iam::aws-account-id:root<br /> arn:aws:iam::aws-account-id:user/user-name<br /> arn:aws:iam::aws-account-id:role/role-name |
        | Azure    | Subscription (GUID) | XXXXXXXX-XXXX-XXXX-XXXX-XXXXXXXXXXXX                                                                                               |
        | GCP      | Project IDs*        | 6 to 30 lowercase letters, digits, or hyphens                                                                                      |

        *https://cloud.google.com/resource-manager/reference/rest/v1/projects
        """
        return pulumi.get(self, "timeout")

