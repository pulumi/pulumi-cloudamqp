# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from . import utilities, tables


class IntegrationMetric(pulumi.CustomResource):
    access_key_id: pulumi.Output[str]
    """
    AWS access key identifier. (Cloudwatch)
    """
    api_key: pulumi.Output[str]
    """
    The API key for the integration service. (Librato)
    """
    email: pulumi.Output[str]
    """
    The email address registred for the integration service. (Librato)
    """
    instance_id: pulumi.Output[float]
    """
    Instance identifier
    """
    license_key: pulumi.Output[str]
    """
    The license key registred for the integration service. (New Relic)
    """
    name: pulumi.Output[str]
    """
    The name of metrics integration
    """
    queue_whitelist: pulumi.Output[str]
    """
    (optional) whitelist using regular expression
    """
    region: pulumi.Output[str]
    """
    AWS region for Cloudwatch and [US/EU] for Data dog/New relic. (Cloudwatch, Data Dog, New Relic)
    """
    secret_access_key: pulumi.Output[str]
    """
    AWS secret key. (Cloudwatch)
    """
    tags: pulumi.Output[str]
    """
    (optional) tags. E.g. env=prod,region=europe
    """
    vhost_whitelist: pulumi.Output[str]
    """
    (optional) whitelist using regular expression
    """
    def __init__(__self__, resource_name, opts=None, access_key_id=None, api_key=None, email=None, instance_id=None, license_key=None, name=None, queue_whitelist=None, region=None, secret_access_key=None, tags=None, vhost_whitelist=None, __props__=None, __name__=None, __opts__=None):
        """
        This resource allows you to create and manage, forwarding metrics to third party integrations for a CloudAMQP instance. Once configured, the metrics produced will be forward to corresponding integration.

        Only available for dedicated subscription plans.

        ## Argument references

        The following arguments are supported:

        * `name`              - (Required) The name of the third party log integration. See `Integration service reference`
        * `region`            - (Optional) Region hosting the integration service.
        * `access_key_id`     - (Optional) AWS access key identifier.
        * `secret_access_key` - (Optional) AWS secret access key.
        * `api_key`           - (Optional) The API key for the integration service.
        * `email`             - (Optional) The email address registred for the integration service.
        * `project_id`        - (Optional) The project identifier.
        * `private_key`       - (Optional) The private access key.
        * `client_email`      - (Optional) The client email registered for the integration service.
        * `tags`              - (Optional) Tags. e.g. env=prod, region=europe.
        * `queue_whitelist`   - (Optional) Whitelist queues using regular expression. Leave empty to include all queues.
        * `vhost_whitelist`   - (Optional) Whitelist vhost using regular expression. Leave empty to include all vhosts.

        This is the full list of all arguments. Only a subset of arguments are used based on which type of integration used. See Integration type reference below for more information.

        ## Integration service references

        Valid names for third party log integration.

        | Name          | Description |
        |---------------|---------------------------------------------------------------|
        | cloudwatch    | Create an IAM with programmatic access. |
        | cloudwatch_v2 | Create an IAM with programmatic access. |
        | datadog       | Create a Datadog API key at app.datadoghq.com |
        | datadog_v2    | Create a Datadog API key at app.datadoghq.com
        | librato       | Create a new API token (with record only permissions) here: https://metrics.librato.com/tokens |
        | newrelic      | Deprecated! |
        | newrelic_v2   | Find or register an Insert API key for your account: Go to insights.newrelic.com > Manage data > API keys. |
        | stackdriver   | Create a service account and add 'monitor metrics writer' role, then download credentials. |

        ## Integration type reference

        Valid arguments for third party log integrations.

        Required arguments for all integrations: *name*<br>
        Optional arguments for all integrations: *tags*, *queue_whitelist*, *vhost_whitelist*

        | Name | Type | Required arguments |
        | ---- | ---- | ---- |
        | Cloudwatch             | cloudwatch     | region, access_key_id, secret_access_key |
        | Cloudwatch v2          | cloudwatch_v2  | region, access_key_id, secret_access_key |
        | Datadog                | datadog        | api_key, region |
        | Datadog v2             | datadog_v2     | api_key, region |
        | Librato                | librato        | email, api_key |
        | New relic (deprecated) | newrelic       | - |
        | New relic v2           | newrelic_v2    | api_key, region |
        | Stackdriver            | stackdriver    | project_id, private_key, client_email |

        ## Dependency

        This resource depends on CloudAMQP instance identifier, `cloudamqp_instance.instance.id`.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] access_key_id: AWS access key identifier. (Cloudwatch)
        :param pulumi.Input[str] api_key: The API key for the integration service. (Librato)
        :param pulumi.Input[str] email: The email address registred for the integration service. (Librato)
        :param pulumi.Input[float] instance_id: Instance identifier
        :param pulumi.Input[str] license_key: The license key registred for the integration service. (New Relic)
        :param pulumi.Input[str] name: The name of metrics integration
        :param pulumi.Input[str] queue_whitelist: (optional) whitelist using regular expression
        :param pulumi.Input[str] region: AWS region for Cloudwatch and [US/EU] for Data dog/New relic. (Cloudwatch, Data Dog, New Relic)
        :param pulumi.Input[str] secret_access_key: AWS secret key. (Cloudwatch)
        :param pulumi.Input[str] tags: (optional) tags. E.g. env=prod,region=europe
        :param pulumi.Input[str] vhost_whitelist: (optional) whitelist using regular expression
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
            __props__['api_key'] = api_key
            __props__['email'] = email
            if instance_id is None:
                raise TypeError("Missing required property 'instance_id'")
            __props__['instance_id'] = instance_id
            __props__['license_key'] = license_key
            __props__['name'] = name
            __props__['queue_whitelist'] = queue_whitelist
            __props__['region'] = region
            __props__['secret_access_key'] = secret_access_key
            __props__['tags'] = tags
            __props__['vhost_whitelist'] = vhost_whitelist
        super(IntegrationMetric, __self__).__init__(
            'cloudamqp:index/integrationMetric:IntegrationMetric',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, access_key_id=None, api_key=None, email=None, instance_id=None, license_key=None, name=None, queue_whitelist=None, region=None, secret_access_key=None, tags=None, vhost_whitelist=None):
        """
        Get an existing IntegrationMetric resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] access_key_id: AWS access key identifier. (Cloudwatch)
        :param pulumi.Input[str] api_key: The API key for the integration service. (Librato)
        :param pulumi.Input[str] email: The email address registred for the integration service. (Librato)
        :param pulumi.Input[float] instance_id: Instance identifier
        :param pulumi.Input[str] license_key: The license key registred for the integration service. (New Relic)
        :param pulumi.Input[str] name: The name of metrics integration
        :param pulumi.Input[str] queue_whitelist: (optional) whitelist using regular expression
        :param pulumi.Input[str] region: AWS region for Cloudwatch and [US/EU] for Data dog/New relic. (Cloudwatch, Data Dog, New Relic)
        :param pulumi.Input[str] secret_access_key: AWS secret key. (Cloudwatch)
        :param pulumi.Input[str] tags: (optional) tags. E.g. env=prod,region=europe
        :param pulumi.Input[str] vhost_whitelist: (optional) whitelist using regular expression
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["access_key_id"] = access_key_id
        __props__["api_key"] = api_key
        __props__["email"] = email
        __props__["instance_id"] = instance_id
        __props__["license_key"] = license_key
        __props__["name"] = name
        __props__["queue_whitelist"] = queue_whitelist
        __props__["region"] = region
        __props__["secret_access_key"] = secret_access_key
        __props__["tags"] = tags
        __props__["vhost_whitelist"] = vhost_whitelist
        return IntegrationMetric(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
