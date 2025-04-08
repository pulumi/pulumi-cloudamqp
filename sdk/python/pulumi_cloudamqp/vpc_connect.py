# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import builtins
import copy
import warnings
import sys
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
if sys.version_info >= (3, 11):
    from typing import NotRequired, TypedDict, TypeAlias
else:
    from typing_extensions import NotRequired, TypedDict, TypeAlias
from . import _utilities

__all__ = ['VpcConnectArgs', 'VpcConnect']

@pulumi.input_type
class VpcConnectArgs:
    def __init__(__self__, *,
                 instance_id: pulumi.Input[builtins.int],
                 region: pulumi.Input[builtins.str],
                 allowed_principals: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
                 allowed_projects: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
                 approved_subscriptions: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
                 sleep: Optional[pulumi.Input[builtins.int]] = None,
                 timeout: Optional[pulumi.Input[builtins.int]] = None):
        """
        The set of arguments for constructing a VpcConnect resource.
        :param pulumi.Input[builtins.int] instance_id: The CloudAMQP instance identifier.
        :param pulumi.Input[builtins.str] region: The region where the CloudAMQP instance is hosted.
        :param pulumi.Input[Sequence[pulumi.Input[builtins.str]]] allowed_principals: List of allowed prinicpals used by AWS, see below table.
        :param pulumi.Input[Sequence[pulumi.Input[builtins.str]]] allowed_projects: List of allowed projects used by GCP, see below table.
        :param pulumi.Input[Sequence[pulumi.Input[builtins.str]]] approved_subscriptions: List of approved subscriptions used by Azure, see below
               table.
        :param pulumi.Input[builtins.int] sleep: Configurable sleep time (seconds) when enable Private
               Service Connect. Default set to 10 seconds.
        :param pulumi.Input[builtins.int] timeout: Configurable timeout time (seconds) when enable Private
               Service Connect. Default set to 1800 seconds.
               
               ___
               
               The `allowed_principals`, `approved_subscriptions` or `allowed_projects` data depends on the
               provider platform:
               
               | Platform | Description | Format |
               |---|---|---|
               | AWS | IAM ARN principals | arn:aws:iam::aws-account-id:root<br>arn:aws:iam::aws-account-id:user/user-name<br> arn:aws:iam::aws-account-id:role/role-name |
               | Azure | Subscription (GUID) | XXXXXXXX-XXXX-XXXX-XXXX-XXXXXXXXXXXX |
               | GCP | Project IDs [Google docs] | 6 to 30 lowercase letters, digits, or hyphens |
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
    def instance_id(self) -> pulumi.Input[builtins.int]:
        """
        The CloudAMQP instance identifier.
        """
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: pulumi.Input[builtins.int]):
        pulumi.set(self, "instance_id", value)

    @property
    @pulumi.getter
    def region(self) -> pulumi.Input[builtins.str]:
        """
        The region where the CloudAMQP instance is hosted.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter(name="allowedPrincipals")
    def allowed_principals(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]]:
        """
        List of allowed prinicpals used by AWS, see below table.
        """
        return pulumi.get(self, "allowed_principals")

    @allowed_principals.setter
    def allowed_principals(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]]):
        pulumi.set(self, "allowed_principals", value)

    @property
    @pulumi.getter(name="allowedProjects")
    def allowed_projects(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]]:
        """
        List of allowed projects used by GCP, see below table.
        """
        return pulumi.get(self, "allowed_projects")

    @allowed_projects.setter
    def allowed_projects(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]]):
        pulumi.set(self, "allowed_projects", value)

    @property
    @pulumi.getter(name="approvedSubscriptions")
    def approved_subscriptions(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]]:
        """
        List of approved subscriptions used by Azure, see below
        table.
        """
        return pulumi.get(self, "approved_subscriptions")

    @approved_subscriptions.setter
    def approved_subscriptions(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]]):
        pulumi.set(self, "approved_subscriptions", value)

    @property
    @pulumi.getter
    def sleep(self) -> Optional[pulumi.Input[builtins.int]]:
        """
        Configurable sleep time (seconds) when enable Private
        Service Connect. Default set to 10 seconds.
        """
        return pulumi.get(self, "sleep")

    @sleep.setter
    def sleep(self, value: Optional[pulumi.Input[builtins.int]]):
        pulumi.set(self, "sleep", value)

    @property
    @pulumi.getter
    def timeout(self) -> Optional[pulumi.Input[builtins.int]]:
        """
        Configurable timeout time (seconds) when enable Private
        Service Connect. Default set to 1800 seconds.

        ___

        The `allowed_principals`, `approved_subscriptions` or `allowed_projects` data depends on the
        provider platform:

        | Platform | Description | Format |
        |---|---|---|
        | AWS | IAM ARN principals | arn:aws:iam::aws-account-id:root<br>arn:aws:iam::aws-account-id:user/user-name<br> arn:aws:iam::aws-account-id:role/role-name |
        | Azure | Subscription (GUID) | XXXXXXXX-XXXX-XXXX-XXXX-XXXXXXXXXXXX |
        | GCP | Project IDs [Google docs] | 6 to 30 lowercase letters, digits, or hyphens |
        """
        return pulumi.get(self, "timeout")

    @timeout.setter
    def timeout(self, value: Optional[pulumi.Input[builtins.int]]):
        pulumi.set(self, "timeout", value)


@pulumi.input_type
class _VpcConnectState:
    def __init__(__self__, *,
                 active_zones: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
                 allowed_principals: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
                 allowed_projects: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
                 approved_subscriptions: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
                 instance_id: Optional[pulumi.Input[builtins.int]] = None,
                 region: Optional[pulumi.Input[builtins.str]] = None,
                 service_name: Optional[pulumi.Input[builtins.str]] = None,
                 sleep: Optional[pulumi.Input[builtins.int]] = None,
                 status: Optional[pulumi.Input[builtins.str]] = None,
                 timeout: Optional[pulumi.Input[builtins.int]] = None):
        """
        Input properties used for looking up and filtering VpcConnect resources.
        :param pulumi.Input[Sequence[pulumi.Input[builtins.str]]] active_zones: Covering availability zones used when creating an endpoint from other VPC. (AWS)
        :param pulumi.Input[Sequence[pulumi.Input[builtins.str]]] allowed_principals: List of allowed prinicpals used by AWS, see below table.
        :param pulumi.Input[Sequence[pulumi.Input[builtins.str]]] allowed_projects: List of allowed projects used by GCP, see below table.
        :param pulumi.Input[Sequence[pulumi.Input[builtins.str]]] approved_subscriptions: List of approved subscriptions used by Azure, see below
               table.
        :param pulumi.Input[builtins.int] instance_id: The CloudAMQP instance identifier.
        :param pulumi.Input[builtins.str] region: The region where the CloudAMQP instance is hosted.
        :param pulumi.Input[builtins.str] service_name: Service name (alias for Azure, see example above) of the PrivateLink.
        :param pulumi.Input[builtins.int] sleep: Configurable sleep time (seconds) when enable Private
               Service Connect. Default set to 10 seconds.
        :param pulumi.Input[builtins.str] status: Private Service Connect status [enable, pending, disable]
        :param pulumi.Input[builtins.int] timeout: Configurable timeout time (seconds) when enable Private
               Service Connect. Default set to 1800 seconds.
               
               ___
               
               The `allowed_principals`, `approved_subscriptions` or `allowed_projects` data depends on the
               provider platform:
               
               | Platform | Description | Format |
               |---|---|---|
               | AWS | IAM ARN principals | arn:aws:iam::aws-account-id:root<br>arn:aws:iam::aws-account-id:user/user-name<br> arn:aws:iam::aws-account-id:role/role-name |
               | Azure | Subscription (GUID) | XXXXXXXX-XXXX-XXXX-XXXX-XXXXXXXXXXXX |
               | GCP | Project IDs [Google docs] | 6 to 30 lowercase letters, digits, or hyphens |
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
    def active_zones(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]]:
        """
        Covering availability zones used when creating an endpoint from other VPC. (AWS)
        """
        return pulumi.get(self, "active_zones")

    @active_zones.setter
    def active_zones(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]]):
        pulumi.set(self, "active_zones", value)

    @property
    @pulumi.getter(name="allowedPrincipals")
    def allowed_principals(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]]:
        """
        List of allowed prinicpals used by AWS, see below table.
        """
        return pulumi.get(self, "allowed_principals")

    @allowed_principals.setter
    def allowed_principals(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]]):
        pulumi.set(self, "allowed_principals", value)

    @property
    @pulumi.getter(name="allowedProjects")
    def allowed_projects(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]]:
        """
        List of allowed projects used by GCP, see below table.
        """
        return pulumi.get(self, "allowed_projects")

    @allowed_projects.setter
    def allowed_projects(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]]):
        pulumi.set(self, "allowed_projects", value)

    @property
    @pulumi.getter(name="approvedSubscriptions")
    def approved_subscriptions(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]]:
        """
        List of approved subscriptions used by Azure, see below
        table.
        """
        return pulumi.get(self, "approved_subscriptions")

    @approved_subscriptions.setter
    def approved_subscriptions(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]]):
        pulumi.set(self, "approved_subscriptions", value)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> Optional[pulumi.Input[builtins.int]]:
        """
        The CloudAMQP instance identifier.
        """
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: Optional[pulumi.Input[builtins.int]]):
        pulumi.set(self, "instance_id", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The region where the CloudAMQP instance is hosted.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter(name="serviceName")
    def service_name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Service name (alias for Azure, see example above) of the PrivateLink.
        """
        return pulumi.get(self, "service_name")

    @service_name.setter
    def service_name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "service_name", value)

    @property
    @pulumi.getter
    def sleep(self) -> Optional[pulumi.Input[builtins.int]]:
        """
        Configurable sleep time (seconds) when enable Private
        Service Connect. Default set to 10 seconds.
        """
        return pulumi.get(self, "sleep")

    @sleep.setter
    def sleep(self, value: Optional[pulumi.Input[builtins.int]]):
        pulumi.set(self, "sleep", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Private Service Connect status [enable, pending, disable]
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "status", value)

    @property
    @pulumi.getter
    def timeout(self) -> Optional[pulumi.Input[builtins.int]]:
        """
        Configurable timeout time (seconds) when enable Private
        Service Connect. Default set to 1800 seconds.

        ___

        The `allowed_principals`, `approved_subscriptions` or `allowed_projects` data depends on the
        provider platform:

        | Platform | Description | Format |
        |---|---|---|
        | AWS | IAM ARN principals | arn:aws:iam::aws-account-id:root<br>arn:aws:iam::aws-account-id:user/user-name<br> arn:aws:iam::aws-account-id:role/role-name |
        | Azure | Subscription (GUID) | XXXXXXXX-XXXX-XXXX-XXXX-XXXXXXXXXXXX |
        | GCP | Project IDs [Google docs] | 6 to 30 lowercase letters, digits, or hyphens |
        """
        return pulumi.get(self, "timeout")

    @timeout.setter
    def timeout(self, value: Optional[pulumi.Input[builtins.int]]):
        pulumi.set(self, "timeout", value)


class VpcConnect(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 allowed_principals: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
                 allowed_projects: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
                 approved_subscriptions: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
                 instance_id: Optional[pulumi.Input[builtins.int]] = None,
                 region: Optional[pulumi.Input[builtins.str]] = None,
                 sleep: Optional[pulumi.Input[builtins.int]] = None,
                 timeout: Optional[pulumi.Input[builtins.int]] = None,
                 __props__=None):
        """
        ## Import

        `cloudamqp_vpc_connect` can be imported using CloudAMQP instance identifier. To

        retrieve the identifier, use [CloudAMQP API list intances].

        From Terraform v1.5.0, the `import` block can be used to import this resource:

        hcl

        import {

          to = cloudamqp_vpc_connect.this

          id = cloudamqp_instance.instance.id

        }

        Or use Terraform CLI:

        ```sh
        $ pulumi import cloudamqp:index/vpcConnect:VpcConnect vpc_connect <id>`
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[builtins.str]]] allowed_principals: List of allowed prinicpals used by AWS, see below table.
        :param pulumi.Input[Sequence[pulumi.Input[builtins.str]]] allowed_projects: List of allowed projects used by GCP, see below table.
        :param pulumi.Input[Sequence[pulumi.Input[builtins.str]]] approved_subscriptions: List of approved subscriptions used by Azure, see below
               table.
        :param pulumi.Input[builtins.int] instance_id: The CloudAMQP instance identifier.
        :param pulumi.Input[builtins.str] region: The region where the CloudAMQP instance is hosted.
        :param pulumi.Input[builtins.int] sleep: Configurable sleep time (seconds) when enable Private
               Service Connect. Default set to 10 seconds.
        :param pulumi.Input[builtins.int] timeout: Configurable timeout time (seconds) when enable Private
               Service Connect. Default set to 1800 seconds.
               
               ___
               
               The `allowed_principals`, `approved_subscriptions` or `allowed_projects` data depends on the
               provider platform:
               
               | Platform | Description | Format |
               |---|---|---|
               | AWS | IAM ARN principals | arn:aws:iam::aws-account-id:root<br>arn:aws:iam::aws-account-id:user/user-name<br> arn:aws:iam::aws-account-id:role/role-name |
               | Azure | Subscription (GUID) | XXXXXXXX-XXXX-XXXX-XXXX-XXXXXXXXXXXX |
               | GCP | Project IDs [Google docs] | 6 to 30 lowercase letters, digits, or hyphens |
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: VpcConnectArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        ## Import

        `cloudamqp_vpc_connect` can be imported using CloudAMQP instance identifier. To

        retrieve the identifier, use [CloudAMQP API list intances].

        From Terraform v1.5.0, the `import` block can be used to import this resource:

        hcl

        import {

          to = cloudamqp_vpc_connect.this

          id = cloudamqp_instance.instance.id

        }

        Or use Terraform CLI:

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
                 allowed_principals: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
                 allowed_projects: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
                 approved_subscriptions: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
                 instance_id: Optional[pulumi.Input[builtins.int]] = None,
                 region: Optional[pulumi.Input[builtins.str]] = None,
                 sleep: Optional[pulumi.Input[builtins.int]] = None,
                 timeout: Optional[pulumi.Input[builtins.int]] = None,
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
            active_zones: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
            allowed_principals: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
            allowed_projects: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
            approved_subscriptions: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
            instance_id: Optional[pulumi.Input[builtins.int]] = None,
            region: Optional[pulumi.Input[builtins.str]] = None,
            service_name: Optional[pulumi.Input[builtins.str]] = None,
            sleep: Optional[pulumi.Input[builtins.int]] = None,
            status: Optional[pulumi.Input[builtins.str]] = None,
            timeout: Optional[pulumi.Input[builtins.int]] = None) -> 'VpcConnect':
        """
        Get an existing VpcConnect resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[builtins.str]]] active_zones: Covering availability zones used when creating an endpoint from other VPC. (AWS)
        :param pulumi.Input[Sequence[pulumi.Input[builtins.str]]] allowed_principals: List of allowed prinicpals used by AWS, see below table.
        :param pulumi.Input[Sequence[pulumi.Input[builtins.str]]] allowed_projects: List of allowed projects used by GCP, see below table.
        :param pulumi.Input[Sequence[pulumi.Input[builtins.str]]] approved_subscriptions: List of approved subscriptions used by Azure, see below
               table.
        :param pulumi.Input[builtins.int] instance_id: The CloudAMQP instance identifier.
        :param pulumi.Input[builtins.str] region: The region where the CloudAMQP instance is hosted.
        :param pulumi.Input[builtins.str] service_name: Service name (alias for Azure, see example above) of the PrivateLink.
        :param pulumi.Input[builtins.int] sleep: Configurable sleep time (seconds) when enable Private
               Service Connect. Default set to 10 seconds.
        :param pulumi.Input[builtins.str] status: Private Service Connect status [enable, pending, disable]
        :param pulumi.Input[builtins.int] timeout: Configurable timeout time (seconds) when enable Private
               Service Connect. Default set to 1800 seconds.
               
               ___
               
               The `allowed_principals`, `approved_subscriptions` or `allowed_projects` data depends on the
               provider platform:
               
               | Platform | Description | Format |
               |---|---|---|
               | AWS | IAM ARN principals | arn:aws:iam::aws-account-id:root<br>arn:aws:iam::aws-account-id:user/user-name<br> arn:aws:iam::aws-account-id:role/role-name |
               | Azure | Subscription (GUID) | XXXXXXXX-XXXX-XXXX-XXXX-XXXXXXXXXXXX |
               | GCP | Project IDs [Google docs] | 6 to 30 lowercase letters, digits, or hyphens |
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
    def active_zones(self) -> pulumi.Output[Sequence[builtins.str]]:
        """
        Covering availability zones used when creating an endpoint from other VPC. (AWS)
        """
        return pulumi.get(self, "active_zones")

    @property
    @pulumi.getter(name="allowedPrincipals")
    def allowed_principals(self) -> pulumi.Output[Optional[Sequence[builtins.str]]]:
        """
        List of allowed prinicpals used by AWS, see below table.
        """
        return pulumi.get(self, "allowed_principals")

    @property
    @pulumi.getter(name="allowedProjects")
    def allowed_projects(self) -> pulumi.Output[Optional[Sequence[builtins.str]]]:
        """
        List of allowed projects used by GCP, see below table.
        """
        return pulumi.get(self, "allowed_projects")

    @property
    @pulumi.getter(name="approvedSubscriptions")
    def approved_subscriptions(self) -> pulumi.Output[Optional[Sequence[builtins.str]]]:
        """
        List of approved subscriptions used by Azure, see below
        table.
        """
        return pulumi.get(self, "approved_subscriptions")

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> pulumi.Output[builtins.int]:
        """
        The CloudAMQP instance identifier.
        """
        return pulumi.get(self, "instance_id")

    @property
    @pulumi.getter
    def region(self) -> pulumi.Output[builtins.str]:
        """
        The region where the CloudAMQP instance is hosted.
        """
        return pulumi.get(self, "region")

    @property
    @pulumi.getter(name="serviceName")
    def service_name(self) -> pulumi.Output[builtins.str]:
        """
        Service name (alias for Azure, see example above) of the PrivateLink.
        """
        return pulumi.get(self, "service_name")

    @property
    @pulumi.getter
    def sleep(self) -> pulumi.Output[Optional[builtins.int]]:
        """
        Configurable sleep time (seconds) when enable Private
        Service Connect. Default set to 10 seconds.
        """
        return pulumi.get(self, "sleep")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[builtins.str]:
        """
        Private Service Connect status [enable, pending, disable]
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter
    def timeout(self) -> pulumi.Output[Optional[builtins.int]]:
        """
        Configurable timeout time (seconds) when enable Private
        Service Connect. Default set to 1800 seconds.

        ___

        The `allowed_principals`, `approved_subscriptions` or `allowed_projects` data depends on the
        provider platform:

        | Platform | Description | Format |
        |---|---|---|
        | AWS | IAM ARN principals | arn:aws:iam::aws-account-id:root<br>arn:aws:iam::aws-account-id:user/user-name<br> arn:aws:iam::aws-account-id:role/role-name |
        | Azure | Subscription (GUID) | XXXXXXXX-XXXX-XXXX-XXXX-XXXXXXXXXXXX |
        | GCP | Project IDs [Google docs] | 6 to 30 lowercase letters, digits, or hyphens |
        """
        return pulumi.get(self, "timeout")

