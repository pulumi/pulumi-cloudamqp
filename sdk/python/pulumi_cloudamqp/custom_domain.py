# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['CustomDomainArgs', 'CustomDomain']

@pulumi.input_type
class CustomDomainArgs:
    def __init__(__self__, *,
                 hostname: pulumi.Input[str],
                 instance_id: pulumi.Input[int]):
        """
        The set of arguments for constructing a CustomDomain resource.
        :param pulumi.Input[str] hostname: Your custom domain name.
        :param pulumi.Input[int] instance_id: The CloudAMQP instance ID.
        """
        pulumi.set(__self__, "hostname", hostname)
        pulumi.set(__self__, "instance_id", instance_id)

    @property
    @pulumi.getter
    def hostname(self) -> pulumi.Input[str]:
        """
        Your custom domain name.
        """
        return pulumi.get(self, "hostname")

    @hostname.setter
    def hostname(self, value: pulumi.Input[str]):
        pulumi.set(self, "hostname", value)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> pulumi.Input[int]:
        """
        The CloudAMQP instance ID.
        """
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: pulumi.Input[int]):
        pulumi.set(self, "instance_id", value)


@pulumi.input_type
class _CustomDomainState:
    def __init__(__self__, *,
                 hostname: Optional[pulumi.Input[str]] = None,
                 instance_id: Optional[pulumi.Input[int]] = None):
        """
        Input properties used for looking up and filtering CustomDomain resources.
        :param pulumi.Input[str] hostname: Your custom domain name.
        :param pulumi.Input[int] instance_id: The CloudAMQP instance ID.
        """
        if hostname is not None:
            pulumi.set(__self__, "hostname", hostname)
        if instance_id is not None:
            pulumi.set(__self__, "instance_id", instance_id)

    @property
    @pulumi.getter
    def hostname(self) -> Optional[pulumi.Input[str]]:
        """
        Your custom domain name.
        """
        return pulumi.get(self, "hostname")

    @hostname.setter
    def hostname(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "hostname", value)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> Optional[pulumi.Input[int]]:
        """
        The CloudAMQP instance ID.
        """
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "instance_id", value)


class CustomDomain(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 hostname: Optional[pulumi.Input[str]] = None,
                 instance_id: Optional[pulumi.Input[int]] = None,
                 __props__=None):
        """
        This resource allows you to configure and manage your custom domain for the CloudAMQP instance.

        Adding a custom domain to your instance will generate a TLS certificate from [Let's Encrypt], for the given hostname, and install it on all servers in your cluster. The certificate will be automatically renewed going forward.

        > **WARNING:** Please note that when creating, changing or deleting the custom domain, the listeners on your servers will be restarted in order to apply the changes. This will close your current connections.

        Your custom domain name needs to point to your CloudAMQP hostname, preferably using a CNAME DNS record.

        Only available for dedicated subscription plans.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_cloudamqp as cloudamqp

        settings = cloudamqp.CustomDomain("settings",
            instance_id=instance["id"],
            hostname="myname.mydomain")
        ```

        ## Depedency

        This resource depends on CloudAMQP instance identifier, `cloudamqp_instance.instance.id`.

        ## Import

        `cloudamqp_custom_domain` can be imported using CloudAMQP instance identifier.

        ```sh
        $ pulumi import cloudamqp:index/customDomain:CustomDomain settings <instance_id>`
        ```

        [Let's Encrypt]: https://letsencrypt.org/

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] hostname: Your custom domain name.
        :param pulumi.Input[int] instance_id: The CloudAMQP instance ID.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: CustomDomainArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        This resource allows you to configure and manage your custom domain for the CloudAMQP instance.

        Adding a custom domain to your instance will generate a TLS certificate from [Let's Encrypt], for the given hostname, and install it on all servers in your cluster. The certificate will be automatically renewed going forward.

        > **WARNING:** Please note that when creating, changing or deleting the custom domain, the listeners on your servers will be restarted in order to apply the changes. This will close your current connections.

        Your custom domain name needs to point to your CloudAMQP hostname, preferably using a CNAME DNS record.

        Only available for dedicated subscription plans.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_cloudamqp as cloudamqp

        settings = cloudamqp.CustomDomain("settings",
            instance_id=instance["id"],
            hostname="myname.mydomain")
        ```

        ## Depedency

        This resource depends on CloudAMQP instance identifier, `cloudamqp_instance.instance.id`.

        ## Import

        `cloudamqp_custom_domain` can be imported using CloudAMQP instance identifier.

        ```sh
        $ pulumi import cloudamqp:index/customDomain:CustomDomain settings <instance_id>`
        ```

        [Let's Encrypt]: https://letsencrypt.org/

        :param str resource_name: The name of the resource.
        :param CustomDomainArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(CustomDomainArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 hostname: Optional[pulumi.Input[str]] = None,
                 instance_id: Optional[pulumi.Input[int]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = CustomDomainArgs.__new__(CustomDomainArgs)

            if hostname is None and not opts.urn:
                raise TypeError("Missing required property 'hostname'")
            __props__.__dict__["hostname"] = hostname
            if instance_id is None and not opts.urn:
                raise TypeError("Missing required property 'instance_id'")
            __props__.__dict__["instance_id"] = instance_id
        super(CustomDomain, __self__).__init__(
            'cloudamqp:index/customDomain:CustomDomain',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            hostname: Optional[pulumi.Input[str]] = None,
            instance_id: Optional[pulumi.Input[int]] = None) -> 'CustomDomain':
        """
        Get an existing CustomDomain resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] hostname: Your custom domain name.
        :param pulumi.Input[int] instance_id: The CloudAMQP instance ID.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _CustomDomainState.__new__(_CustomDomainState)

        __props__.__dict__["hostname"] = hostname
        __props__.__dict__["instance_id"] = instance_id
        return CustomDomain(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def hostname(self) -> pulumi.Output[str]:
        """
        Your custom domain name.
        """
        return pulumi.get(self, "hostname")

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> pulumi.Output[int]:
        """
        The CloudAMQP instance ID.
        """
        return pulumi.get(self, "instance_id")

