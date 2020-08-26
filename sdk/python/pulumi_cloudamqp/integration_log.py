# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from . import _utilities, _tables

__all__ = ['IntegrationLog']


class IntegrationLog(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 access_key_id: Optional[pulumi.Input[str]] = None,
                 api_key: Optional[pulumi.Input[str]] = None,
                 client_email: Optional[pulumi.Input[str]] = None,
                 host_port: Optional[pulumi.Input[str]] = None,
                 instance_id: Optional[pulumi.Input[float]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 private_key: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 secret_access_key: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[str]] = None,
                 token: Optional[pulumi.Input[str]] = None,
                 url: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        This resource allows you to create and manage third party log integrations for a CloudAMQP instance. Once configured, the logs produced will be forward to corresponding integration.

        Only available for dedicated subscription plans.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_cloudamqp as cloudamqp

        cloudwatch = cloudamqp.IntegrationLog("cloudwatch",
            instance_id=cloudamqp_instance["instance"]["id"],
            access_key_id=var["aws_access_key_id"],
            secret_access_key=var["aws_secret_access_key"],
            region=var["aws_region"])
        logentries = cloudamqp.IntegrationLog("logentries",
            instance_id=cloudamqp_instance["instance"]["id"],
            token=var["logentries_token"])
        loggly = cloudamqp.IntegrationLog("loggly",
            instance_id=cloudamqp_instance["instance"]["id"],
            token=var["loggly_token"])
        papertrail = cloudamqp.IntegrationLog("papertrail",
            instance_id=cloudamqp_instance["instance"]["id"],
            url=var["papertrail_url"])
        splunk = cloudamqp.IntegrationLog("splunk",
            instance_id=cloudamqp_instance["instance"]["id"],
            token=var["splunk_token"],
            host_port=var["splunk_host_port"])
        datadog = cloudamqp.IntegrationLog("datadog",
            instance_id=cloudamqp_instance["instance"]["id"],
            region=var["datadog_region"],
            api_key=var["datadog_api_key"],
            tags=var["datadog_tags"])
        stackdriver = cloudamqp.IntegrationLog("stackdriver",
            instance_id=cloudamqp_instance["instance"]["id"],
            project_id=var["stackdriver_project_id"],
            private_key=var["stackdriver_private_key"],
            client_email=var["stackdriver_client_email"])
        ```
        ## Argument Reference (cloudwatchlog)

        Cloudwatch argument reference and example. Create an IAM user with programmatic access and the following permissions:

        * CreateLogGroup
        * CreateLogStream
        * DescribeLogGroups
        * DescribeLogStreams
        * PutLogEvents

        ## Integration service reference

        Valid names for third party log integration.

        | Name       | Description |
        |------------|---------------------------------------------------------------|
        | cloudwatchlog | Create a IAM with programmatic access. |
        | logentries | Create a Logentries token at https://logentries.com/app#/add-log/manual  |
        | loggly     | Create a Loggly token at https://{your-company}.loggly.com/tokens |
        | papertrail | Create a Papertrail endpoint https://papertrailapp.com/systems/setup |
        | splunk     | Create a HTTP Event Collector token at https://.cloud.splunk.com/en-US/manager/search/http-eventcollector |
        | datadog       | Create a Datadog API key at app.datadoghq.com |
        | stackdriver   | Create a service account and add 'monitor metrics writer' role, then download credentials. |

        ## Integration Type reference

        Valid arguments for third party log integrations.

        Required arguments for all integrations: name

        | Name | Type | Required arguments |
        | ---- | ---- | ---- |
        | CloudWatch | cloudwatchlog | access_key_id, secret_access_key, region |
        | Log Entries | logentries | token |
        | Loggly | loggly | token |
        | Papertrail | papertrail | url |
        | Splunk | splunk | token, host_port |
        | Data Dog | datadog | region, api_keys, tags |
        | Stackdriver | stackdriver | project_id, private_key, client_email |

        ## Dependency

        This resource depends on CloudAMQP instance identifier, `cloudamqp_instance.instance.id`.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] access_key_id: AWS access key identifier.
        :param pulumi.Input[str] api_key: The API key.
        :param pulumi.Input[str] client_email: The client email registered for the integration service.
        :param pulumi.Input[str] host_port: Destination to send the logs.
        :param pulumi.Input[float] instance_id: Instance identifier used to make proxy calls
        :param pulumi.Input[str] name: The name of the third party log integration. See
        :param pulumi.Input[str] private_key: The private access key.
        :param pulumi.Input[str] project_id: The project identifier.
        :param pulumi.Input[str] region: Region hosting the integration service.
        :param pulumi.Input[str] secret_access_key: AWS secret access key.
        :param pulumi.Input[str] tags: Tag the integration, e.g. env=prod, region=europe.
        :param pulumi.Input[str] token: Token used for authentication.
        :param pulumi.Input[str] url: Endpoint to log integration.
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
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            __props__['access_key_id'] = access_key_id
            __props__['api_key'] = api_key
            __props__['client_email'] = client_email
            __props__['host_port'] = host_port
            if instance_id is None:
                raise TypeError("Missing required property 'instance_id'")
            __props__['instance_id'] = instance_id
            __props__['name'] = name
            __props__['private_key'] = private_key
            __props__['project_id'] = project_id
            __props__['region'] = region
            __props__['secret_access_key'] = secret_access_key
            __props__['tags'] = tags
            __props__['token'] = token
            __props__['url'] = url
        super(IntegrationLog, __self__).__init__(
            'cloudamqp:index/integrationLog:IntegrationLog',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            access_key_id: Optional[pulumi.Input[str]] = None,
            api_key: Optional[pulumi.Input[str]] = None,
            client_email: Optional[pulumi.Input[str]] = None,
            host_port: Optional[pulumi.Input[str]] = None,
            instance_id: Optional[pulumi.Input[float]] = None,
            name: Optional[pulumi.Input[str]] = None,
            private_key: Optional[pulumi.Input[str]] = None,
            project_id: Optional[pulumi.Input[str]] = None,
            region: Optional[pulumi.Input[str]] = None,
            secret_access_key: Optional[pulumi.Input[str]] = None,
            tags: Optional[pulumi.Input[str]] = None,
            token: Optional[pulumi.Input[str]] = None,
            url: Optional[pulumi.Input[str]] = None) -> 'IntegrationLog':
        """
        Get an existing IntegrationLog resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] access_key_id: AWS access key identifier.
        :param pulumi.Input[str] api_key: The API key.
        :param pulumi.Input[str] client_email: The client email registered for the integration service.
        :param pulumi.Input[str] host_port: Destination to send the logs.
        :param pulumi.Input[float] instance_id: Instance identifier used to make proxy calls
        :param pulumi.Input[str] name: The name of the third party log integration. See
        :param pulumi.Input[str] private_key: The private access key.
        :param pulumi.Input[str] project_id: The project identifier.
        :param pulumi.Input[str] region: Region hosting the integration service.
        :param pulumi.Input[str] secret_access_key: AWS secret access key.
        :param pulumi.Input[str] tags: Tag the integration, e.g. env=prod, region=europe.
        :param pulumi.Input[str] token: Token used for authentication.
        :param pulumi.Input[str] url: Endpoint to log integration.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["access_key_id"] = access_key_id
        __props__["api_key"] = api_key
        __props__["client_email"] = client_email
        __props__["host_port"] = host_port
        __props__["instance_id"] = instance_id
        __props__["name"] = name
        __props__["private_key"] = private_key
        __props__["project_id"] = project_id
        __props__["region"] = region
        __props__["secret_access_key"] = secret_access_key
        __props__["tags"] = tags
        __props__["token"] = token
        __props__["url"] = url
        return IntegrationLog(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="accessKeyId")
    def access_key_id(self) -> Optional[str]:
        """
        AWS access key identifier.
        """
        return pulumi.get(self, "access_key_id")

    @property
    @pulumi.getter(name="apiKey")
    def api_key(self) -> Optional[str]:
        """
        The API key.
        """
        return pulumi.get(self, "api_key")

    @property
    @pulumi.getter(name="clientEmail")
    def client_email(self) -> Optional[str]:
        """
        The client email registered for the integration service.
        """
        return pulumi.get(self, "client_email")

    @property
    @pulumi.getter(name="hostPort")
    def host_port(self) -> Optional[str]:
        """
        Destination to send the logs.
        """
        return pulumi.get(self, "host_port")

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> float:
        """
        Instance identifier used to make proxy calls
        """
        return pulumi.get(self, "instance_id")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The name of the third party log integration. See
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="privateKey")
    def private_key(self) -> Optional[str]:
        """
        The private access key.
        """
        return pulumi.get(self, "private_key")

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> Optional[str]:
        """
        The project identifier.
        """
        return pulumi.get(self, "project_id")

    @property
    @pulumi.getter
    def region(self) -> Optional[str]:
        """
        Region hosting the integration service.
        """
        return pulumi.get(self, "region")

    @property
    @pulumi.getter(name="secretAccessKey")
    def secret_access_key(self) -> Optional[str]:
        """
        AWS secret access key.
        """
        return pulumi.get(self, "secret_access_key")

    @property
    @pulumi.getter
    def tags(self) -> Optional[str]:
        """
        Tag the integration, e.g. env=prod, region=europe.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def token(self) -> Optional[str]:
        """
        Token used for authentication.
        """
        return pulumi.get(self, "token")

    @property
    @pulumi.getter
    def url(self) -> Optional[str]:
        """
        Endpoint to log integration.
        """
        return pulumi.get(self, "url")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

