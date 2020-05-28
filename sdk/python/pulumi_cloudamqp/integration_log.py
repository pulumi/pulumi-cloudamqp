# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from . import utilities, tables

class IntegrationLog(pulumi.CustomResource):
    access_key_id: pulumi.Output[str]
    """
    AWS access key identifier. (Cloudwatch)
    """
    host_port: pulumi.Output[str]
    """
    Destination to send the logs. (Splunk)
    """
    instance_id: pulumi.Output[float]
    """
    Instance identifier used to make proxy calls
    """
    name: pulumi.Output[str]
    """
    The name of log integration
    """
    region: pulumi.Output[str]
    """
    The region hosting integration service. (Cloudwatch)
    """
    secret_access_key: pulumi.Output[str]
    """
    AWS secret access key. (Cloudwatch)
    """
    token: pulumi.Output[str]
    """
    The token used for authentication. (Loggly, Logentries, Splunk)
    """
    url: pulumi.Output[str]
    """
    The URL to push the logs to. (Papertrail)
    """
    def __init__(__self__, resource_name, opts=None, access_key_id=None, host_port=None, instance_id=None, name=None, region=None, secret_access_key=None, token=None, url=None, __props__=None, __name__=None, __opts__=None):
        """
        Create a IntegrationLog resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] access_key_id: AWS access key identifier. (Cloudwatch)
        :param pulumi.Input[str] host_port: Destination to send the logs. (Splunk)
        :param pulumi.Input[float] instance_id: Instance identifier used to make proxy calls
        :param pulumi.Input[str] name: The name of log integration
        :param pulumi.Input[str] region: The region hosting integration service. (Cloudwatch)
        :param pulumi.Input[str] secret_access_key: AWS secret access key. (Cloudwatch)
        :param pulumi.Input[str] token: The token used for authentication. (Loggly, Logentries, Splunk)
        :param pulumi.Input[str] url: The URL to push the logs to. (Papertrail)
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            __props__['access_key_id'] = access_key_id
            __props__['host_port'] = host_port
            if instance_id is None:
                raise TypeError("Missing required property 'instance_id'")
            __props__['instance_id'] = instance_id
            __props__['name'] = name
            __props__['region'] = region
            __props__['secret_access_key'] = secret_access_key
            __props__['token'] = token
            __props__['url'] = url
        super(IntegrationLog, __self__).__init__(
            'cloudamqp:index/integrationLog:IntegrationLog',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, access_key_id=None, host_port=None, instance_id=None, name=None, region=None, secret_access_key=None, token=None, url=None):
        """
        Get an existing IntegrationLog resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] access_key_id: AWS access key identifier. (Cloudwatch)
        :param pulumi.Input[str] host_port: Destination to send the logs. (Splunk)
        :param pulumi.Input[float] instance_id: Instance identifier used to make proxy calls
        :param pulumi.Input[str] name: The name of log integration
        :param pulumi.Input[str] region: The region hosting integration service. (Cloudwatch)
        :param pulumi.Input[str] secret_access_key: AWS secret access key. (Cloudwatch)
        :param pulumi.Input[str] token: The token used for authentication. (Loggly, Logentries, Splunk)
        :param pulumi.Input[str] url: The URL to push the logs to. (Papertrail)
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["access_key_id"] = access_key_id
        __props__["host_port"] = host_port
        __props__["instance_id"] = instance_id
        __props__["name"] = name
        __props__["region"] = region
        __props__["secret_access_key"] = secret_access_key
        __props__["token"] = token
        __props__["url"] = url
        return IntegrationLog(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
