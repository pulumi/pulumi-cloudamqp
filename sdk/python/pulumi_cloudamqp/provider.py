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

__all__ = ['ProviderArgs', 'Provider']

@pulumi.input_type
class ProviderArgs:
    def __init__(__self__, *,
                 apikey: pulumi.Input[builtins.str],
                 baseurl: Optional[pulumi.Input[builtins.str]] = None,
                 enable_faster_instance_destroy: Optional[pulumi.Input[builtins.bool]] = None):
        """
        The set of arguments for constructing a Provider resource.
        :param pulumi.Input[builtins.str] apikey: Key used to authentication to the CloudAMQP Customer API
        :param pulumi.Input[builtins.str] baseurl: Base URL to CloudAMQP Customer website
        """
        pulumi.set(__self__, "apikey", apikey)
        if baseurl is not None:
            pulumi.set(__self__, "baseurl", baseurl)
        if enable_faster_instance_destroy is not None:
            pulumi.set(__self__, "enable_faster_instance_destroy", enable_faster_instance_destroy)

    @property
    @pulumi.getter
    def apikey(self) -> pulumi.Input[builtins.str]:
        """
        Key used to authentication to the CloudAMQP Customer API
        """
        return pulumi.get(self, "apikey")

    @apikey.setter
    def apikey(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "apikey", value)

    @property
    @pulumi.getter
    def baseurl(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Base URL to CloudAMQP Customer website
        """
        return pulumi.get(self, "baseurl")

    @baseurl.setter
    def baseurl(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "baseurl", value)

    @property
    @pulumi.getter(name="enableFasterInstanceDestroy")
    def enable_faster_instance_destroy(self) -> Optional[pulumi.Input[builtins.bool]]:
        return pulumi.get(self, "enable_faster_instance_destroy")

    @enable_faster_instance_destroy.setter
    def enable_faster_instance_destroy(self, value: Optional[pulumi.Input[builtins.bool]]):
        pulumi.set(self, "enable_faster_instance_destroy", value)


class Provider(pulumi.ProviderResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 apikey: Optional[pulumi.Input[builtins.str]] = None,
                 baseurl: Optional[pulumi.Input[builtins.str]] = None,
                 enable_faster_instance_destroy: Optional[pulumi.Input[builtins.bool]] = None,
                 __props__=None):
        """
        The provider type for the cloudamqp package. By default, resources use package-wide configuration
        settings, however an explicit `Provider` instance may be created and passed during resource
        construction to achieve fine-grained programmatic control over provider settings. See the
        [documentation](https://www.pulumi.com/docs/reference/programming-model/#providers) for more information.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.str] apikey: Key used to authentication to the CloudAMQP Customer API
        :param pulumi.Input[builtins.str] baseurl: Base URL to CloudAMQP Customer website
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ProviderArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        The provider type for the cloudamqp package. By default, resources use package-wide configuration
        settings, however an explicit `Provider` instance may be created and passed during resource
        construction to achieve fine-grained programmatic control over provider settings. See the
        [documentation](https://www.pulumi.com/docs/reference/programming-model/#providers) for more information.

        :param str resource_name: The name of the resource.
        :param ProviderArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ProviderArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 apikey: Optional[pulumi.Input[builtins.str]] = None,
                 baseurl: Optional[pulumi.Input[builtins.str]] = None,
                 enable_faster_instance_destroy: Optional[pulumi.Input[builtins.bool]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ProviderArgs.__new__(ProviderArgs)

            if apikey is None and not opts.urn:
                raise TypeError("Missing required property 'apikey'")
            __props__.__dict__["apikey"] = apikey
            __props__.__dict__["baseurl"] = baseurl
            __props__.__dict__["enable_faster_instance_destroy"] = pulumi.Output.from_input(enable_faster_instance_destroy).apply(pulumi.runtime.to_json) if enable_faster_instance_destroy is not None else None
        super(Provider, __self__).__init__(
            'cloudamqp',
            resource_name,
            __props__,
            opts)

    @property
    @pulumi.getter
    def apikey(self) -> pulumi.Output[builtins.str]:
        """
        Key used to authentication to the CloudAMQP Customer API
        """
        return pulumi.get(self, "apikey")

    @property
    @pulumi.getter
    def baseurl(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        Base URL to CloudAMQP Customer website
        """
        return pulumi.get(self, "baseurl")

