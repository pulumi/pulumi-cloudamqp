# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Callable, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['ProviderArgs', 'Provider']

@pulumi.input_type
class ProviderArgs:
    def __init__(__self__, *,
                 apikey: pulumi.Input[str],
                 baseurl: Optional[pulumi.Input[str]] = None,
                 enable_faster_instance_destroy: Optional[pulumi.Input[bool]] = None):
        """
        The set of arguments for constructing a Provider resource.
        :param pulumi.Input[str] apikey: Key used to authentication to the CloudAMQP Customer API
        :param pulumi.Input[str] baseurl: Base URL to CloudAMQP Customer website
        :param pulumi.Input[bool] enable_faster_instance_destroy: Skips destroying backend resources on 'terraform destroy'
        """
        ProviderArgs._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            apikey=apikey,
            baseurl=baseurl,
            enable_faster_instance_destroy=enable_faster_instance_destroy,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             apikey: pulumi.Input[str],
             baseurl: Optional[pulumi.Input[str]] = None,
             enable_faster_instance_destroy: Optional[pulumi.Input[bool]] = None,
             opts: Optional[pulumi.ResourceOptions]=None):
        _setter("apikey", apikey)
        if baseurl is not None:
            _setter("baseurl", baseurl)
        if enable_faster_instance_destroy is not None:
            _setter("enable_faster_instance_destroy", enable_faster_instance_destroy)

    @property
    @pulumi.getter
    def apikey(self) -> pulumi.Input[str]:
        """
        Key used to authentication to the CloudAMQP Customer API
        """
        return pulumi.get(self, "apikey")

    @apikey.setter
    def apikey(self, value: pulumi.Input[str]):
        pulumi.set(self, "apikey", value)

    @property
    @pulumi.getter
    def baseurl(self) -> Optional[pulumi.Input[str]]:
        """
        Base URL to CloudAMQP Customer website
        """
        return pulumi.get(self, "baseurl")

    @baseurl.setter
    def baseurl(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "baseurl", value)

    @property
    @pulumi.getter(name="enableFasterInstanceDestroy")
    def enable_faster_instance_destroy(self) -> Optional[pulumi.Input[bool]]:
        """
        Skips destroying backend resources on 'terraform destroy'
        """
        return pulumi.get(self, "enable_faster_instance_destroy")

    @enable_faster_instance_destroy.setter
    def enable_faster_instance_destroy(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enable_faster_instance_destroy", value)


class Provider(pulumi.ProviderResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 apikey: Optional[pulumi.Input[str]] = None,
                 baseurl: Optional[pulumi.Input[str]] = None,
                 enable_faster_instance_destroy: Optional[pulumi.Input[bool]] = None,
                 __props__=None):
        """
        The provider type for the cloudamqp package. By default, resources use package-wide configuration
        settings, however an explicit `Provider` instance may be created and passed during resource
        construction to achieve fine-grained programmatic control over provider settings. See the
        [documentation](https://www.pulumi.com/docs/reference/programming-model/#providers) for more information.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] apikey: Key used to authentication to the CloudAMQP Customer API
        :param pulumi.Input[str] baseurl: Base URL to CloudAMQP Customer website
        :param pulumi.Input[bool] enable_faster_instance_destroy: Skips destroying backend resources on 'terraform destroy'
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
            kwargs = kwargs or {}
            def _setter(key, value):
                kwargs[key] = value
            ProviderArgs._configure(_setter, **kwargs)
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 apikey: Optional[pulumi.Input[str]] = None,
                 baseurl: Optional[pulumi.Input[str]] = None,
                 enable_faster_instance_destroy: Optional[pulumi.Input[bool]] = None,
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
    def apikey(self) -> pulumi.Output[str]:
        """
        Key used to authentication to the CloudAMQP Customer API
        """
        return pulumi.get(self, "apikey")

    @property
    @pulumi.getter
    def baseurl(self) -> pulumi.Output[Optional[str]]:
        """
        Base URL to CloudAMQP Customer website
        """
        return pulumi.get(self, "baseurl")

