# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['PrivatelinkAzureArgs', 'PrivatelinkAzure']

@pulumi.input_type
class PrivatelinkAzureArgs:
    def __init__(__self__, *,
                 approved_subscriptions: pulumi.Input[Sequence[pulumi.Input[str]]],
                 instance_id: pulumi.Input[int],
                 sleep: Optional[pulumi.Input[int]] = None,
                 timeout: Optional[pulumi.Input[int]] = None):
        """
        The set of arguments for constructing a PrivatelinkAzure resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] approved_subscriptions: Approved subscriptions to access the endpoint service.
               See format below.
        :param pulumi.Input[int] instance_id: The CloudAMQP instance identifier.
        :param pulumi.Input[int] sleep: Configurable sleep time (seconds) when enable PrivateLink.
               Default set to 10 seconds. *Available from v1.29.0*
        :param pulumi.Input[int] timeout: Configurable timeout time (seconds) when enable PrivateLink.
               Default set to 1800 seconds. *Available from v1.29.0*
               
               Approved subscriptions format (GUID): <br>
               `XXXXXXXX-XXXX-XXXX-XXXX-XXXXXXXXXXXX`
        """
        pulumi.set(__self__, "approved_subscriptions", approved_subscriptions)
        pulumi.set(__self__, "instance_id", instance_id)
        if sleep is not None:
            pulumi.set(__self__, "sleep", sleep)
        if timeout is not None:
            pulumi.set(__self__, "timeout", timeout)

    @property
    @pulumi.getter(name="approvedSubscriptions")
    def approved_subscriptions(self) -> pulumi.Input[Sequence[pulumi.Input[str]]]:
        """
        Approved subscriptions to access the endpoint service.
        See format below.
        """
        return pulumi.get(self, "approved_subscriptions")

    @approved_subscriptions.setter
    def approved_subscriptions(self, value: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(self, "approved_subscriptions", value)

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

        Approved subscriptions format (GUID): <br>
        `XXXXXXXX-XXXX-XXXX-XXXX-XXXXXXXXXXXX`
        """
        return pulumi.get(self, "timeout")

    @timeout.setter
    def timeout(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "timeout", value)


@pulumi.input_type
class _PrivatelinkAzureState:
    def __init__(__self__, *,
                 approved_subscriptions: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 instance_id: Optional[pulumi.Input[int]] = None,
                 server_name: Optional[pulumi.Input[str]] = None,
                 service_name: Optional[pulumi.Input[str]] = None,
                 sleep: Optional[pulumi.Input[int]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 timeout: Optional[pulumi.Input[int]] = None):
        """
        Input properties used for looking up and filtering PrivatelinkAzure resources.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] approved_subscriptions: Approved subscriptions to access the endpoint service.
               See format below.
        :param pulumi.Input[int] instance_id: The CloudAMQP instance identifier.
        :param pulumi.Input[str] server_name: Name of the server having the PrivateLink enabled.
        :param pulumi.Input[str] service_name: Service name (alias) of the PrivateLink, needed when creating the endpoint.
        :param pulumi.Input[int] sleep: Configurable sleep time (seconds) when enable PrivateLink.
               Default set to 10 seconds. *Available from v1.29.0*
        :param pulumi.Input[str] status: PrivateLink status [enable, pending, disable]
        :param pulumi.Input[int] timeout: Configurable timeout time (seconds) when enable PrivateLink.
               Default set to 1800 seconds. *Available from v1.29.0*
               
               Approved subscriptions format (GUID): <br>
               `XXXXXXXX-XXXX-XXXX-XXXX-XXXXXXXXXXXX`
        """
        if approved_subscriptions is not None:
            pulumi.set(__self__, "approved_subscriptions", approved_subscriptions)
        if instance_id is not None:
            pulumi.set(__self__, "instance_id", instance_id)
        if server_name is not None:
            pulumi.set(__self__, "server_name", server_name)
        if service_name is not None:
            pulumi.set(__self__, "service_name", service_name)
        if sleep is not None:
            pulumi.set(__self__, "sleep", sleep)
        if status is not None:
            pulumi.set(__self__, "status", status)
        if timeout is not None:
            pulumi.set(__self__, "timeout", timeout)

    @property
    @pulumi.getter(name="approvedSubscriptions")
    def approved_subscriptions(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Approved subscriptions to access the endpoint service.
        See format below.
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
    @pulumi.getter(name="serverName")
    def server_name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the server having the PrivateLink enabled.
        """
        return pulumi.get(self, "server_name")

    @server_name.setter
    def server_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "server_name", value)

    @property
    @pulumi.getter(name="serviceName")
    def service_name(self) -> Optional[pulumi.Input[str]]:
        """
        Service name (alias) of the PrivateLink, needed when creating the endpoint.
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

        Approved subscriptions format (GUID): <br>
        `XXXXXXXX-XXXX-XXXX-XXXX-XXXXXXXXXXXX`
        """
        return pulumi.get(self, "timeout")

    @timeout.setter
    def timeout(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "timeout", value)


class PrivatelinkAzure(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 approved_subscriptions: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 instance_id: Optional[pulumi.Input[int]] = None,
                 sleep: Optional[pulumi.Input[int]] = None,
                 timeout: Optional[pulumi.Input[int]] = None,
                 __props__=None):
        """
        Enable PrivateLink for a CloudAMQP instance hosted in Azure. If no existing VPC available when
        enable PrivateLink, a new VPC will be created with subnet `10.52.72.0/24`.

        > **Note:** Enabling PrivateLink will automatically add firewall rules for the peered subnet.

        <details>
         <summary>
            <i>Default PrivateLink firewall rule</i>
          </summary>

        ## Example Usage

        <details>
          <summary>
            <b>
              <i>CloudAMQP instance without existing VPC</i>
            </b>
          </summary>

        ```python
        import pulumi
        import pulumi_cloudamqp as cloudamqp

        instance = cloudamqp.Instance("instance",
            name="Instance 01",
            plan="bunny-1",
            region="azure-arm::westus",
            tags=[])
        privatelink = cloudamqp.PrivatelinkAzure("privatelink",
            instance_id=instance.id,
            approved_subscriptions=["XXXXXXXX-XXXX-XXXX-XXXX-XXXXXXXXXXXX"])
        ```
        </details>

        <details>
          <summary>
            <b>
              <i>CloudAMQP instance in an existing VPC</i>
            </b>
          </summary>

        ```python
        import pulumi
        import pulumi_cloudamqp as cloudamqp

        vpc = cloudamqp.Vpc("vpc",
            name="Standalone VPC",
            region="azure-arm::westus",
            subnet="10.56.72.0/24",
            tags=[])
        instance = cloudamqp.Instance("instance",
            name="Instance 01",
            plan="bunny-1",
            region="azure-arm::westus",
            tags=[],
            vpc_id=vpc.id,
            keep_associated_vpc=True)
        privatelink = cloudamqp.PrivatelinkAzure("privatelink",
            instance_id=instance.id,
            approved_subscriptions=["XXXXXXXX-XXXX-XXXX-XXXX-XXXXXXXXXXXX"])
        ```

        </details>

        ### With Additional Firewall Rules

        <details>
          <summary>
            <b>
              <i>CloudAMQP instance in an existing VPC with managed firewall rules</i>
            </b>
          </summary>

        ```python
        import pulumi
        import pulumi_cloudamqp as cloudamqp

        vpc = cloudamqp.Vpc("vpc",
            name="Standalone VPC",
            region="azure-arm::westus",
            subnet="10.56.72.0/24",
            tags=[])
        instance = cloudamqp.Instance("instance",
            name="Instance 01",
            plan="bunny-1",
            region="azure-arm::westus",
            tags=[],
            vpc_id=vpc.id,
            keep_associated_vpc=True)
        privatelink = cloudamqp.PrivatelinkAzure("privatelink",
            instance_id=instance.id,
            approved_subscriptions=["XXXXXXXX-XXXX-XXXX-XXXX-XXXXXXXXXXXX"])
        firewall_settings = cloudamqp.SecurityFirewall("firewall_settings",
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

        </details>

        ## Depedency

        This resource depends on CloudAMQP instance identifier, `cloudamqp_instance.instance.id`.

        ## Create PrivateLink with additional firewall rules

        To create a PrivateLink configuration with additional firewall rules, it's required to chain the SecurityFirewall
        resource to avoid parallel conflicting resource calls. You can do this by making the firewall
        resource depend on the PrivateLink resource, `cloudamqp_privatelink_azure.privatelink`.

        Furthermore, since all firewall rules are overwritten, the otherwise automatically added rules for
        the PrivateLink also needs to be added.

        ## Import

        `cloudamqp_privatelink_aws` can be imported using CloudAMQP internal identifier.

        ```sh
        $ pulumi import cloudamqp:index/privatelinkAzure:PrivatelinkAzure privatelink <id>`
        ```

        The resource uses the same identifier as the CloudAMQP instance. To retrieve the identifier for an instance, either use [CloudAMQP customer API](https://docs.cloudamqp.com/#list-instances) or use the data source [`cloudamqp_account`](./data-sources/account.md).

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] approved_subscriptions: Approved subscriptions to access the endpoint service.
               See format below.
        :param pulumi.Input[int] instance_id: The CloudAMQP instance identifier.
        :param pulumi.Input[int] sleep: Configurable sleep time (seconds) when enable PrivateLink.
               Default set to 10 seconds. *Available from v1.29.0*
        :param pulumi.Input[int] timeout: Configurable timeout time (seconds) when enable PrivateLink.
               Default set to 1800 seconds. *Available from v1.29.0*
               
               Approved subscriptions format (GUID): <br>
               `XXXXXXXX-XXXX-XXXX-XXXX-XXXXXXXXXXXX`
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: PrivatelinkAzureArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Enable PrivateLink for a CloudAMQP instance hosted in Azure. If no existing VPC available when
        enable PrivateLink, a new VPC will be created with subnet `10.52.72.0/24`.

        > **Note:** Enabling PrivateLink will automatically add firewall rules for the peered subnet.

        <details>
         <summary>
            <i>Default PrivateLink firewall rule</i>
          </summary>

        ## Example Usage

        <details>
          <summary>
            <b>
              <i>CloudAMQP instance without existing VPC</i>
            </b>
          </summary>

        ```python
        import pulumi
        import pulumi_cloudamqp as cloudamqp

        instance = cloudamqp.Instance("instance",
            name="Instance 01",
            plan="bunny-1",
            region="azure-arm::westus",
            tags=[])
        privatelink = cloudamqp.PrivatelinkAzure("privatelink",
            instance_id=instance.id,
            approved_subscriptions=["XXXXXXXX-XXXX-XXXX-XXXX-XXXXXXXXXXXX"])
        ```
        </details>

        <details>
          <summary>
            <b>
              <i>CloudAMQP instance in an existing VPC</i>
            </b>
          </summary>

        ```python
        import pulumi
        import pulumi_cloudamqp as cloudamqp

        vpc = cloudamqp.Vpc("vpc",
            name="Standalone VPC",
            region="azure-arm::westus",
            subnet="10.56.72.0/24",
            tags=[])
        instance = cloudamqp.Instance("instance",
            name="Instance 01",
            plan="bunny-1",
            region="azure-arm::westus",
            tags=[],
            vpc_id=vpc.id,
            keep_associated_vpc=True)
        privatelink = cloudamqp.PrivatelinkAzure("privatelink",
            instance_id=instance.id,
            approved_subscriptions=["XXXXXXXX-XXXX-XXXX-XXXX-XXXXXXXXXXXX"])
        ```

        </details>

        ### With Additional Firewall Rules

        <details>
          <summary>
            <b>
              <i>CloudAMQP instance in an existing VPC with managed firewall rules</i>
            </b>
          </summary>

        ```python
        import pulumi
        import pulumi_cloudamqp as cloudamqp

        vpc = cloudamqp.Vpc("vpc",
            name="Standalone VPC",
            region="azure-arm::westus",
            subnet="10.56.72.0/24",
            tags=[])
        instance = cloudamqp.Instance("instance",
            name="Instance 01",
            plan="bunny-1",
            region="azure-arm::westus",
            tags=[],
            vpc_id=vpc.id,
            keep_associated_vpc=True)
        privatelink = cloudamqp.PrivatelinkAzure("privatelink",
            instance_id=instance.id,
            approved_subscriptions=["XXXXXXXX-XXXX-XXXX-XXXX-XXXXXXXXXXXX"])
        firewall_settings = cloudamqp.SecurityFirewall("firewall_settings",
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

        </details>

        ## Depedency

        This resource depends on CloudAMQP instance identifier, `cloudamqp_instance.instance.id`.

        ## Create PrivateLink with additional firewall rules

        To create a PrivateLink configuration with additional firewall rules, it's required to chain the SecurityFirewall
        resource to avoid parallel conflicting resource calls. You can do this by making the firewall
        resource depend on the PrivateLink resource, `cloudamqp_privatelink_azure.privatelink`.

        Furthermore, since all firewall rules are overwritten, the otherwise automatically added rules for
        the PrivateLink also needs to be added.

        ## Import

        `cloudamqp_privatelink_aws` can be imported using CloudAMQP internal identifier.

        ```sh
        $ pulumi import cloudamqp:index/privatelinkAzure:PrivatelinkAzure privatelink <id>`
        ```

        The resource uses the same identifier as the CloudAMQP instance. To retrieve the identifier for an instance, either use [CloudAMQP customer API](https://docs.cloudamqp.com/#list-instances) or use the data source [`cloudamqp_account`](./data-sources/account.md).

        :param str resource_name: The name of the resource.
        :param PrivatelinkAzureArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(PrivatelinkAzureArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 approved_subscriptions: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
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
            __props__ = PrivatelinkAzureArgs.__new__(PrivatelinkAzureArgs)

            if approved_subscriptions is None and not opts.urn:
                raise TypeError("Missing required property 'approved_subscriptions'")
            __props__.__dict__["approved_subscriptions"] = approved_subscriptions
            if instance_id is None and not opts.urn:
                raise TypeError("Missing required property 'instance_id'")
            __props__.__dict__["instance_id"] = instance_id
            __props__.__dict__["sleep"] = sleep
            __props__.__dict__["timeout"] = timeout
            __props__.__dict__["server_name"] = None
            __props__.__dict__["service_name"] = None
            __props__.__dict__["status"] = None
        super(PrivatelinkAzure, __self__).__init__(
            'cloudamqp:index/privatelinkAzure:PrivatelinkAzure',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            approved_subscriptions: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            instance_id: Optional[pulumi.Input[int]] = None,
            server_name: Optional[pulumi.Input[str]] = None,
            service_name: Optional[pulumi.Input[str]] = None,
            sleep: Optional[pulumi.Input[int]] = None,
            status: Optional[pulumi.Input[str]] = None,
            timeout: Optional[pulumi.Input[int]] = None) -> 'PrivatelinkAzure':
        """
        Get an existing PrivatelinkAzure resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] approved_subscriptions: Approved subscriptions to access the endpoint service.
               See format below.
        :param pulumi.Input[int] instance_id: The CloudAMQP instance identifier.
        :param pulumi.Input[str] server_name: Name of the server having the PrivateLink enabled.
        :param pulumi.Input[str] service_name: Service name (alias) of the PrivateLink, needed when creating the endpoint.
        :param pulumi.Input[int] sleep: Configurable sleep time (seconds) when enable PrivateLink.
               Default set to 10 seconds. *Available from v1.29.0*
        :param pulumi.Input[str] status: PrivateLink status [enable, pending, disable]
        :param pulumi.Input[int] timeout: Configurable timeout time (seconds) when enable PrivateLink.
               Default set to 1800 seconds. *Available from v1.29.0*
               
               Approved subscriptions format (GUID): <br>
               `XXXXXXXX-XXXX-XXXX-XXXX-XXXXXXXXXXXX`
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _PrivatelinkAzureState.__new__(_PrivatelinkAzureState)

        __props__.__dict__["approved_subscriptions"] = approved_subscriptions
        __props__.__dict__["instance_id"] = instance_id
        __props__.__dict__["server_name"] = server_name
        __props__.__dict__["service_name"] = service_name
        __props__.__dict__["sleep"] = sleep
        __props__.__dict__["status"] = status
        __props__.__dict__["timeout"] = timeout
        return PrivatelinkAzure(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="approvedSubscriptions")
    def approved_subscriptions(self) -> pulumi.Output[Sequence[str]]:
        """
        Approved subscriptions to access the endpoint service.
        See format below.
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
    @pulumi.getter(name="serverName")
    def server_name(self) -> pulumi.Output[str]:
        """
        Name of the server having the PrivateLink enabled.
        """
        return pulumi.get(self, "server_name")

    @property
    @pulumi.getter(name="serviceName")
    def service_name(self) -> pulumi.Output[str]:
        """
        Service name (alias) of the PrivateLink, needed when creating the endpoint.
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

        Approved subscriptions format (GUID): <br>
        `XXXXXXXX-XXXX-XXXX-XXXX-XXXXXXXXXXXX`
        """
        return pulumi.get(self, "timeout")

