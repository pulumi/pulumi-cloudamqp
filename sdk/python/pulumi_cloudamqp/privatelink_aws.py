# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['PrivatelinkAwsArgs', 'PrivatelinkAws']

@pulumi.input_type
class PrivatelinkAwsArgs:
    def __init__(__self__, *,
                 allowed_principals: pulumi.Input[Sequence[pulumi.Input[str]]],
                 instance_id: pulumi.Input[int],
                 sleep: Optional[pulumi.Input[int]] = None,
                 timeout: Optional[pulumi.Input[int]] = None):
        """
        The set of arguments for constructing a PrivatelinkAws resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] allowed_principals: Allowed principals to access the endpoint service.
        :param pulumi.Input[int] instance_id: The CloudAMQP instance identifier.
        :param pulumi.Input[int] sleep: Configurable sleep time (seconds) when enable PrivateLink.
               Default set to 10 seconds. *Available from v1.29.0*
        :param pulumi.Input[int] timeout: Configurable timeout time (seconds) when enable PrivateLink.
               Default set to 1800 seconds. *Available from v1.29.0*
               
               Allowed principals format: <br>
               `arn:aws:iam::aws-account-id:root` <br>
               `arn:aws:iam::aws-account-id:user/user-name` <br>
               `arn:aws:iam::aws-account-id:role/role-name`
        """
        pulumi.set(__self__, "allowed_principals", allowed_principals)
        pulumi.set(__self__, "instance_id", instance_id)
        if sleep is not None:
            pulumi.set(__self__, "sleep", sleep)
        if timeout is not None:
            pulumi.set(__self__, "timeout", timeout)

    @property
    @pulumi.getter(name="allowedPrincipals")
    def allowed_principals(self) -> pulumi.Input[Sequence[pulumi.Input[str]]]:
        """
        Allowed principals to access the endpoint service.
        """
        return pulumi.get(self, "allowed_principals")

    @allowed_principals.setter
    def allowed_principals(self, value: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(self, "allowed_principals", value)

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
    def sleep(self) -> Optional[pulumi.Input[int]]:
        """
        Configurable sleep time (seconds) when enable PrivateLink.
        Default set to 10 seconds. *Available from v1.29.0*
        """
        return pulumi.get(self, "sleep")

    @sleep.setter
    def sleep(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "sleep", value)

    @property
    @pulumi.getter
    def timeout(self) -> Optional[pulumi.Input[int]]:
        """
        Configurable timeout time (seconds) when enable PrivateLink.
        Default set to 1800 seconds. *Available from v1.29.0*

        Allowed principals format: <br>
        `arn:aws:iam::aws-account-id:root` <br>
        `arn:aws:iam::aws-account-id:user/user-name` <br>
        `arn:aws:iam::aws-account-id:role/role-name`
        """
        return pulumi.get(self, "timeout")

    @timeout.setter
    def timeout(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "timeout", value)


@pulumi.input_type
class _PrivatelinkAwsState:
    def __init__(__self__, *,
                 active_zones: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 allowed_principals: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 instance_id: Optional[pulumi.Input[int]] = None,
                 service_name: Optional[pulumi.Input[str]] = None,
                 sleep: Optional[pulumi.Input[int]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 timeout: Optional[pulumi.Input[int]] = None):
        """
        Input properties used for looking up and filtering PrivatelinkAws resources.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] active_zones: Covering availability zones used when creating an Endpoint from other VPC.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] allowed_principals: Allowed principals to access the endpoint service.
        :param pulumi.Input[int] instance_id: The CloudAMQP instance identifier.
        :param pulumi.Input[str] service_name: Service name of the PrivateLink used when creating the endpoint from other VPC.
        :param pulumi.Input[int] sleep: Configurable sleep time (seconds) when enable PrivateLink.
               Default set to 10 seconds. *Available from v1.29.0*
        :param pulumi.Input[str] status: PrivateLink status [enable, pending, disable]
        :param pulumi.Input[int] timeout: Configurable timeout time (seconds) when enable PrivateLink.
               Default set to 1800 seconds. *Available from v1.29.0*
               
               Allowed principals format: <br>
               `arn:aws:iam::aws-account-id:root` <br>
               `arn:aws:iam::aws-account-id:user/user-name` <br>
               `arn:aws:iam::aws-account-id:role/role-name`
        """
        if active_zones is not None:
            pulumi.set(__self__, "active_zones", active_zones)
        if allowed_principals is not None:
            pulumi.set(__self__, "allowed_principals", allowed_principals)
        if instance_id is not None:
            pulumi.set(__self__, "instance_id", instance_id)
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
        Covering availability zones used when creating an Endpoint from other VPC.
        """
        return pulumi.get(self, "active_zones")

    @active_zones.setter
    def active_zones(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "active_zones", value)

    @property
    @pulumi.getter(name="allowedPrincipals")
    def allowed_principals(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Allowed principals to access the endpoint service.
        """
        return pulumi.get(self, "allowed_principals")

    @allowed_principals.setter
    def allowed_principals(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "allowed_principals", value)

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
    @pulumi.getter(name="serviceName")
    def service_name(self) -> Optional[pulumi.Input[str]]:
        """
        Service name of the PrivateLink used when creating the endpoint from other VPC.
        """
        return pulumi.get(self, "service_name")

    @service_name.setter
    def service_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "service_name", value)

    @property
    @pulumi.getter
    def sleep(self) -> Optional[pulumi.Input[int]]:
        """
        Configurable sleep time (seconds) when enable PrivateLink.
        Default set to 10 seconds. *Available from v1.29.0*
        """
        return pulumi.get(self, "sleep")

    @sleep.setter
    def sleep(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "sleep", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        """
        PrivateLink status [enable, pending, disable]
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)

    @property
    @pulumi.getter
    def timeout(self) -> Optional[pulumi.Input[int]]:
        """
        Configurable timeout time (seconds) when enable PrivateLink.
        Default set to 1800 seconds. *Available from v1.29.0*

        Allowed principals format: <br>
        `arn:aws:iam::aws-account-id:root` <br>
        `arn:aws:iam::aws-account-id:user/user-name` <br>
        `arn:aws:iam::aws-account-id:role/role-name`
        """
        return pulumi.get(self, "timeout")

    @timeout.setter
    def timeout(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "timeout", value)


class PrivatelinkAws(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 allowed_principals: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 instance_id: Optional[pulumi.Input[int]] = None,
                 sleep: Optional[pulumi.Input[int]] = None,
                 timeout: Optional[pulumi.Input[int]] = None,
                 __props__=None):
        """
        Enable PrivateLink for a CloudAMQP instance hosted in AWS. If no existing VPC available when enable
        PrivateLink, a new VPC will be created with subnet `10.52.72.0/24`.

        > **Note:** Enabling PrivateLink will automatically add firewall rules for the peered subnet.
        <details>
         <summary>
            <i>Default PrivateLink firewall rule</i>
          </summary>

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        ```
        <!--End PulumiCodeChooser -->

        </details>

        Pricing is available at [cloudamqp.com](https://www.cloudamqp.com/plans.html)
        where you can also find more information about
        [CloudAMQP PrivateLink](https://www.cloudamqp.com/docs/cloudamqp-privatelink.html#aws-privatelink).

        Only available for dedicated subscription plans.

        > **Warning:** This resource considered deprecated and will be removed in next major version (v2.0).
        Recommended to start using the new resource`VpcConnect`.

        ## Example Usage

        <details>
          <summary>
            <b>
              <i>CloudAMQP instance without existing VPC</i>
            </b>
          </summary>

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumi_cloudamqp as cloudamqp

        instance = cloudamqp.Instance("instance",
            plan="bunny-1",
            region="amazon-web-services::us-west-1",
            tags=[])
        privatelink = cloudamqp.PrivatelinkAws("privatelink",
            instance_id=instance.id,
            allowed_principals=["arn:aws:iam::aws-account-id:user/user-name"])
        ```
        <!--End PulumiCodeChooser -->

        </details>

        <details>
          <summary>
            <b>
              <i>CloudAMQP instance in an existing VPC</i>
            </b>
          </summary>

        <!--Start PulumiCodeChooser -->
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
        privatelink = cloudamqp.PrivatelinkAws("privatelink",
            instance_id=instance.id,
            allowed_principals=["arn:aws:iam::aws-account-id:user/user-name"])
        ```
        <!--End PulumiCodeChooser -->

        </details>

        ### With Additional Firewall Rules

        <details>
          <summary>
            <b>
              <i>CloudAMQP instance in an existing VPC with managed firewall rules</i>
            </b>
          </summary>

        <!--Start PulumiCodeChooser -->
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
        privatelink = cloudamqp.PrivatelinkAws("privatelink",
            instance_id=instance.id,
            allowed_principals=["arn:aws:iam::aws-account-id:user/user-name"])
        firewall_settings = cloudamqp.SecurityFirewall("firewallSettings",
            instance_id=instance.id,
            rules=[
                cloudamqp.SecurityFirewallRuleArgs(
                    description="Custom PrivateLink setup",
                    ip=vpc.subnet,
                    ports=[],
                    services=[
                        "AMQP",
                        "AMQPS",
                        "HTTPS",
                        "STREAM",
                        "STREAM_SSL",
                    ],
                ),
                cloudamqp.SecurityFirewallRuleArgs(
                    description="MGMT interface",
                    ip="0.0.0.0/0",
                    ports=[],
                    services=["HTTPS"],
                ),
            ],
            opts=pulumi.ResourceOptions(depends_on=[privatelink]))
        ```
        <!--End PulumiCodeChooser -->

        </details>

        ## Depedency

        This resource depends on CloudAMQP instance identifier, `cloudamqp_instance.instance.id`.

        ## Create PrivateLink with additional firewall rules

        To create a PrivateLink configuration with additional firewall rules, it's required to chain the SecurityFirewall
        resource to avoid parallel conflicting resource calls. You can do this by making the firewall
        resource depend on the PrivateLink resource, `cloudamqp_privatelink_aws.privatelink`.

        Furthermore, since all firewall rules are overwritten, the otherwise automatically added rules for
        the PrivateLink also needs to be added.

        ## Import

        `cloudamqp_privatelink_aws` can be imported using CloudAMQP internal identifier.

        ```sh
        $ pulumi import cloudamqp:index/privatelinkAws:PrivatelinkAws privatelink <id>`
        ```

        The resource uses the same identifier as the CloudAMQP instance. To retrieve the identifier for an instance, either use [CloudAMQP customer API](https://docs.cloudamqp.com/#list-instances) or use the data source [`cloudamqp_account`](./data-sources/account.md).

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] allowed_principals: Allowed principals to access the endpoint service.
        :param pulumi.Input[int] instance_id: The CloudAMQP instance identifier.
        :param pulumi.Input[int] sleep: Configurable sleep time (seconds) when enable PrivateLink.
               Default set to 10 seconds. *Available from v1.29.0*
        :param pulumi.Input[int] timeout: Configurable timeout time (seconds) when enable PrivateLink.
               Default set to 1800 seconds. *Available from v1.29.0*
               
               Allowed principals format: <br>
               `arn:aws:iam::aws-account-id:root` <br>
               `arn:aws:iam::aws-account-id:user/user-name` <br>
               `arn:aws:iam::aws-account-id:role/role-name`
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: PrivatelinkAwsArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Enable PrivateLink for a CloudAMQP instance hosted in AWS. If no existing VPC available when enable
        PrivateLink, a new VPC will be created with subnet `10.52.72.0/24`.

        > **Note:** Enabling PrivateLink will automatically add firewall rules for the peered subnet.
        <details>
         <summary>
            <i>Default PrivateLink firewall rule</i>
          </summary>

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        ```
        <!--End PulumiCodeChooser -->

        </details>

        Pricing is available at [cloudamqp.com](https://www.cloudamqp.com/plans.html)
        where you can also find more information about
        [CloudAMQP PrivateLink](https://www.cloudamqp.com/docs/cloudamqp-privatelink.html#aws-privatelink).

        Only available for dedicated subscription plans.

        > **Warning:** This resource considered deprecated and will be removed in next major version (v2.0).
        Recommended to start using the new resource`VpcConnect`.

        ## Example Usage

        <details>
          <summary>
            <b>
              <i>CloudAMQP instance without existing VPC</i>
            </b>
          </summary>

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumi_cloudamqp as cloudamqp

        instance = cloudamqp.Instance("instance",
            plan="bunny-1",
            region="amazon-web-services::us-west-1",
            tags=[])
        privatelink = cloudamqp.PrivatelinkAws("privatelink",
            instance_id=instance.id,
            allowed_principals=["arn:aws:iam::aws-account-id:user/user-name"])
        ```
        <!--End PulumiCodeChooser -->

        </details>

        <details>
          <summary>
            <b>
              <i>CloudAMQP instance in an existing VPC</i>
            </b>
          </summary>

        <!--Start PulumiCodeChooser -->
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
        privatelink = cloudamqp.PrivatelinkAws("privatelink",
            instance_id=instance.id,
            allowed_principals=["arn:aws:iam::aws-account-id:user/user-name"])
        ```
        <!--End PulumiCodeChooser -->

        </details>

        ### With Additional Firewall Rules

        <details>
          <summary>
            <b>
              <i>CloudAMQP instance in an existing VPC with managed firewall rules</i>
            </b>
          </summary>

        <!--Start PulumiCodeChooser -->
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
        privatelink = cloudamqp.PrivatelinkAws("privatelink",
            instance_id=instance.id,
            allowed_principals=["arn:aws:iam::aws-account-id:user/user-name"])
        firewall_settings = cloudamqp.SecurityFirewall("firewallSettings",
            instance_id=instance.id,
            rules=[
                cloudamqp.SecurityFirewallRuleArgs(
                    description="Custom PrivateLink setup",
                    ip=vpc.subnet,
                    ports=[],
                    services=[
                        "AMQP",
                        "AMQPS",
                        "HTTPS",
                        "STREAM",
                        "STREAM_SSL",
                    ],
                ),
                cloudamqp.SecurityFirewallRuleArgs(
                    description="MGMT interface",
                    ip="0.0.0.0/0",
                    ports=[],
                    services=["HTTPS"],
                ),
            ],
            opts=pulumi.ResourceOptions(depends_on=[privatelink]))
        ```
        <!--End PulumiCodeChooser -->

        </details>

        ## Depedency

        This resource depends on CloudAMQP instance identifier, `cloudamqp_instance.instance.id`.

        ## Create PrivateLink with additional firewall rules

        To create a PrivateLink configuration with additional firewall rules, it's required to chain the SecurityFirewall
        resource to avoid parallel conflicting resource calls. You can do this by making the firewall
        resource depend on the PrivateLink resource, `cloudamqp_privatelink_aws.privatelink`.

        Furthermore, since all firewall rules are overwritten, the otherwise automatically added rules for
        the PrivateLink also needs to be added.

        ## Import

        `cloudamqp_privatelink_aws` can be imported using CloudAMQP internal identifier.

        ```sh
        $ pulumi import cloudamqp:index/privatelinkAws:PrivatelinkAws privatelink <id>`
        ```

        The resource uses the same identifier as the CloudAMQP instance. To retrieve the identifier for an instance, either use [CloudAMQP customer API](https://docs.cloudamqp.com/#list-instances) or use the data source [`cloudamqp_account`](./data-sources/account.md).

        :param str resource_name: The name of the resource.
        :param PrivatelinkAwsArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(PrivatelinkAwsArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 allowed_principals: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 instance_id: Optional[pulumi.Input[int]] = None,
                 sleep: Optional[pulumi.Input[int]] = None,
                 timeout: Optional[pulumi.Input[int]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = PrivatelinkAwsArgs.__new__(PrivatelinkAwsArgs)

            if allowed_principals is None and not opts.urn:
                raise TypeError("Missing required property 'allowed_principals'")
            __props__.__dict__["allowed_principals"] = allowed_principals
            if instance_id is None and not opts.urn:
                raise TypeError("Missing required property 'instance_id'")
            __props__.__dict__["instance_id"] = instance_id
            __props__.__dict__["sleep"] = sleep
            __props__.__dict__["timeout"] = timeout
            __props__.__dict__["active_zones"] = None
            __props__.__dict__["service_name"] = None
            __props__.__dict__["status"] = None
        super(PrivatelinkAws, __self__).__init__(
            'cloudamqp:index/privatelinkAws:PrivatelinkAws',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            active_zones: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            allowed_principals: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            instance_id: Optional[pulumi.Input[int]] = None,
            service_name: Optional[pulumi.Input[str]] = None,
            sleep: Optional[pulumi.Input[int]] = None,
            status: Optional[pulumi.Input[str]] = None,
            timeout: Optional[pulumi.Input[int]] = None) -> 'PrivatelinkAws':
        """
        Get an existing PrivatelinkAws resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] active_zones: Covering availability zones used when creating an Endpoint from other VPC.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] allowed_principals: Allowed principals to access the endpoint service.
        :param pulumi.Input[int] instance_id: The CloudAMQP instance identifier.
        :param pulumi.Input[str] service_name: Service name of the PrivateLink used when creating the endpoint from other VPC.
        :param pulumi.Input[int] sleep: Configurable sleep time (seconds) when enable PrivateLink.
               Default set to 10 seconds. *Available from v1.29.0*
        :param pulumi.Input[str] status: PrivateLink status [enable, pending, disable]
        :param pulumi.Input[int] timeout: Configurable timeout time (seconds) when enable PrivateLink.
               Default set to 1800 seconds. *Available from v1.29.0*
               
               Allowed principals format: <br>
               `arn:aws:iam::aws-account-id:root` <br>
               `arn:aws:iam::aws-account-id:user/user-name` <br>
               `arn:aws:iam::aws-account-id:role/role-name`
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _PrivatelinkAwsState.__new__(_PrivatelinkAwsState)

        __props__.__dict__["active_zones"] = active_zones
        __props__.__dict__["allowed_principals"] = allowed_principals
        __props__.__dict__["instance_id"] = instance_id
        __props__.__dict__["service_name"] = service_name
        __props__.__dict__["sleep"] = sleep
        __props__.__dict__["status"] = status
        __props__.__dict__["timeout"] = timeout
        return PrivatelinkAws(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="activeZones")
    def active_zones(self) -> pulumi.Output[Sequence[str]]:
        """
        Covering availability zones used when creating an Endpoint from other VPC.
        """
        return pulumi.get(self, "active_zones")

    @property
    @pulumi.getter(name="allowedPrincipals")
    def allowed_principals(self) -> pulumi.Output[Sequence[str]]:
        """
        Allowed principals to access the endpoint service.
        """
        return pulumi.get(self, "allowed_principals")

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> pulumi.Output[int]:
        """
        The CloudAMQP instance identifier.
        """
        return pulumi.get(self, "instance_id")

    @property
    @pulumi.getter(name="serviceName")
    def service_name(self) -> pulumi.Output[str]:
        """
        Service name of the PrivateLink used when creating the endpoint from other VPC.
        """
        return pulumi.get(self, "service_name")

    @property
    @pulumi.getter
    def sleep(self) -> pulumi.Output[Optional[int]]:
        """
        Configurable sleep time (seconds) when enable PrivateLink.
        Default set to 10 seconds. *Available from v1.29.0*
        """
        return pulumi.get(self, "sleep")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[str]:
        """
        PrivateLink status [enable, pending, disable]
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter
    def timeout(self) -> pulumi.Output[Optional[int]]:
        """
        Configurable timeout time (seconds) when enable PrivateLink.
        Default set to 1800 seconds. *Available from v1.29.0*

        Allowed principals format: <br>
        `arn:aws:iam::aws-account-id:root` <br>
        `arn:aws:iam::aws-account-id:user/user-name` <br>
        `arn:aws:iam::aws-account-id:role/role-name`
        """
        return pulumi.get(self, "timeout")

