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

__all__ = ['WebhookArgs', 'Webhook']

@pulumi.input_type
class WebhookArgs:
    def __init__(__self__, *,
                 concurrency: pulumi.Input[builtins.int],
                 instance_id: pulumi.Input[builtins.int],
                 queue: pulumi.Input[builtins.str],
                 vhost: pulumi.Input[builtins.str],
                 webhook_uri: pulumi.Input[builtins.str],
                 sleep: Optional[pulumi.Input[builtins.int]] = None,
                 timeout: Optional[pulumi.Input[builtins.int]] = None):
        """
        The set of arguments for constructing a Webhook resource.
        :param pulumi.Input[builtins.int] concurrency: Max simultaneous requests to the endpoint.
        :param pulumi.Input[builtins.int] instance_id: The CloudAMQP instance ID.
        :param pulumi.Input[builtins.str] queue: A (durable) queue on your RabbitMQ instance.
        :param pulumi.Input[builtins.str] vhost: The vhost the queue resides in.
        :param pulumi.Input[builtins.str] webhook_uri: A POST request will be made for each message in the queue to this
               endpoint.
        :param pulumi.Input[builtins.int] sleep: Configurable sleep time in seconds between retries for webhook
        :param pulumi.Input[builtins.int] timeout: Configurable timeout time in seconds for webhook
        """
        pulumi.set(__self__, "concurrency", concurrency)
        pulumi.set(__self__, "instance_id", instance_id)
        pulumi.set(__self__, "queue", queue)
        pulumi.set(__self__, "vhost", vhost)
        pulumi.set(__self__, "webhook_uri", webhook_uri)
        if sleep is not None:
            pulumi.set(__self__, "sleep", sleep)
        if timeout is not None:
            pulumi.set(__self__, "timeout", timeout)

    @property
    @pulumi.getter
    def concurrency(self) -> pulumi.Input[builtins.int]:
        """
        Max simultaneous requests to the endpoint.
        """
        return pulumi.get(self, "concurrency")

    @concurrency.setter
    def concurrency(self, value: pulumi.Input[builtins.int]):
        pulumi.set(self, "concurrency", value)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> pulumi.Input[builtins.int]:
        """
        The CloudAMQP instance ID.
        """
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: pulumi.Input[builtins.int]):
        pulumi.set(self, "instance_id", value)

    @property
    @pulumi.getter
    def queue(self) -> pulumi.Input[builtins.str]:
        """
        A (durable) queue on your RabbitMQ instance.
        """
        return pulumi.get(self, "queue")

    @queue.setter
    def queue(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "queue", value)

    @property
    @pulumi.getter
    def vhost(self) -> pulumi.Input[builtins.str]:
        """
        The vhost the queue resides in.
        """
        return pulumi.get(self, "vhost")

    @vhost.setter
    def vhost(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "vhost", value)

    @property
    @pulumi.getter(name="webhookUri")
    def webhook_uri(self) -> pulumi.Input[builtins.str]:
        """
        A POST request will be made for each message in the queue to this
        endpoint.
        """
        return pulumi.get(self, "webhook_uri")

    @webhook_uri.setter
    def webhook_uri(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "webhook_uri", value)

    @property
    @pulumi.getter
    def sleep(self) -> Optional[pulumi.Input[builtins.int]]:
        """
        Configurable sleep time in seconds between retries for webhook
        """
        return pulumi.get(self, "sleep")

    @sleep.setter
    def sleep(self, value: Optional[pulumi.Input[builtins.int]]):
        pulumi.set(self, "sleep", value)

    @property
    @pulumi.getter
    def timeout(self) -> Optional[pulumi.Input[builtins.int]]:
        """
        Configurable timeout time in seconds for webhook
        """
        return pulumi.get(self, "timeout")

    @timeout.setter
    def timeout(self, value: Optional[pulumi.Input[builtins.int]]):
        pulumi.set(self, "timeout", value)


@pulumi.input_type
class _WebhookState:
    def __init__(__self__, *,
                 concurrency: Optional[pulumi.Input[builtins.int]] = None,
                 instance_id: Optional[pulumi.Input[builtins.int]] = None,
                 queue: Optional[pulumi.Input[builtins.str]] = None,
                 sleep: Optional[pulumi.Input[builtins.int]] = None,
                 timeout: Optional[pulumi.Input[builtins.int]] = None,
                 vhost: Optional[pulumi.Input[builtins.str]] = None,
                 webhook_uri: Optional[pulumi.Input[builtins.str]] = None):
        """
        Input properties used for looking up and filtering Webhook resources.
        :param pulumi.Input[builtins.int] concurrency: Max simultaneous requests to the endpoint.
        :param pulumi.Input[builtins.int] instance_id: The CloudAMQP instance ID.
        :param pulumi.Input[builtins.str] queue: A (durable) queue on your RabbitMQ instance.
        :param pulumi.Input[builtins.int] sleep: Configurable sleep time in seconds between retries for webhook
        :param pulumi.Input[builtins.int] timeout: Configurable timeout time in seconds for webhook
        :param pulumi.Input[builtins.str] vhost: The vhost the queue resides in.
        :param pulumi.Input[builtins.str] webhook_uri: A POST request will be made for each message in the queue to this
               endpoint.
        """
        if concurrency is not None:
            pulumi.set(__self__, "concurrency", concurrency)
        if instance_id is not None:
            pulumi.set(__self__, "instance_id", instance_id)
        if queue is not None:
            pulumi.set(__self__, "queue", queue)
        if sleep is not None:
            pulumi.set(__self__, "sleep", sleep)
        if timeout is not None:
            pulumi.set(__self__, "timeout", timeout)
        if vhost is not None:
            pulumi.set(__self__, "vhost", vhost)
        if webhook_uri is not None:
            pulumi.set(__self__, "webhook_uri", webhook_uri)

    @property
    @pulumi.getter
    def concurrency(self) -> Optional[pulumi.Input[builtins.int]]:
        """
        Max simultaneous requests to the endpoint.
        """
        return pulumi.get(self, "concurrency")

    @concurrency.setter
    def concurrency(self, value: Optional[pulumi.Input[builtins.int]]):
        pulumi.set(self, "concurrency", value)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> Optional[pulumi.Input[builtins.int]]:
        """
        The CloudAMQP instance ID.
        """
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: Optional[pulumi.Input[builtins.int]]):
        pulumi.set(self, "instance_id", value)

    @property
    @pulumi.getter
    def queue(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        A (durable) queue on your RabbitMQ instance.
        """
        return pulumi.get(self, "queue")

    @queue.setter
    def queue(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "queue", value)

    @property
    @pulumi.getter
    def sleep(self) -> Optional[pulumi.Input[builtins.int]]:
        """
        Configurable sleep time in seconds between retries for webhook
        """
        return pulumi.get(self, "sleep")

    @sleep.setter
    def sleep(self, value: Optional[pulumi.Input[builtins.int]]):
        pulumi.set(self, "sleep", value)

    @property
    @pulumi.getter
    def timeout(self) -> Optional[pulumi.Input[builtins.int]]:
        """
        Configurable timeout time in seconds for webhook
        """
        return pulumi.get(self, "timeout")

    @timeout.setter
    def timeout(self, value: Optional[pulumi.Input[builtins.int]]):
        pulumi.set(self, "timeout", value)

    @property
    @pulumi.getter
    def vhost(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The vhost the queue resides in.
        """
        return pulumi.get(self, "vhost")

    @vhost.setter
    def vhost(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "vhost", value)

    @property
    @pulumi.getter(name="webhookUri")
    def webhook_uri(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        A POST request will be made for each message in the queue to this
        endpoint.
        """
        return pulumi.get(self, "webhook_uri")

    @webhook_uri.setter
    def webhook_uri(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "webhook_uri", value)


class Webhook(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 concurrency: Optional[pulumi.Input[builtins.int]] = None,
                 instance_id: Optional[pulumi.Input[builtins.int]] = None,
                 queue: Optional[pulumi.Input[builtins.str]] = None,
                 sleep: Optional[pulumi.Input[builtins.int]] = None,
                 timeout: Optional[pulumi.Input[builtins.int]] = None,
                 vhost: Optional[pulumi.Input[builtins.str]] = None,
                 webhook_uri: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        """
        ## Import

        `cloudamqp_webhook` can be imported using the resource identifier together with CloudAMQP instance

        identifier (CSV separated). To retrieve the resource identifier, use [CloudAMQP API list webhooks].

        From Terraform v1.5.0, the `import` block can be used to import this resource:

        hcl

        import {

          to = cloudamqp_webhook.webhook_queue

          id = format("<id>,%s", cloudamqp_instance.instance.id)

        }

        Or use Terraform CLI:

        ```sh
        $ pulumi import cloudamqp:index/webhook:Webhook webhook_queue <id>,<instance_id>`
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.int] concurrency: Max simultaneous requests to the endpoint.
        :param pulumi.Input[builtins.int] instance_id: The CloudAMQP instance ID.
        :param pulumi.Input[builtins.str] queue: A (durable) queue on your RabbitMQ instance.
        :param pulumi.Input[builtins.int] sleep: Configurable sleep time in seconds between retries for webhook
        :param pulumi.Input[builtins.int] timeout: Configurable timeout time in seconds for webhook
        :param pulumi.Input[builtins.str] vhost: The vhost the queue resides in.
        :param pulumi.Input[builtins.str] webhook_uri: A POST request will be made for each message in the queue to this
               endpoint.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: WebhookArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        ## Import

        `cloudamqp_webhook` can be imported using the resource identifier together with CloudAMQP instance

        identifier (CSV separated). To retrieve the resource identifier, use [CloudAMQP API list webhooks].

        From Terraform v1.5.0, the `import` block can be used to import this resource:

        hcl

        import {

          to = cloudamqp_webhook.webhook_queue

          id = format("<id>,%s", cloudamqp_instance.instance.id)

        }

        Or use Terraform CLI:

        ```sh
        $ pulumi import cloudamqp:index/webhook:Webhook webhook_queue <id>,<instance_id>`
        ```

        :param str resource_name: The name of the resource.
        :param WebhookArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(WebhookArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 concurrency: Optional[pulumi.Input[builtins.int]] = None,
                 instance_id: Optional[pulumi.Input[builtins.int]] = None,
                 queue: Optional[pulumi.Input[builtins.str]] = None,
                 sleep: Optional[pulumi.Input[builtins.int]] = None,
                 timeout: Optional[pulumi.Input[builtins.int]] = None,
                 vhost: Optional[pulumi.Input[builtins.str]] = None,
                 webhook_uri: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = WebhookArgs.__new__(WebhookArgs)

            if concurrency is None and not opts.urn:
                raise TypeError("Missing required property 'concurrency'")
            __props__.__dict__["concurrency"] = concurrency
            if instance_id is None and not opts.urn:
                raise TypeError("Missing required property 'instance_id'")
            __props__.__dict__["instance_id"] = instance_id
            if queue is None and not opts.urn:
                raise TypeError("Missing required property 'queue'")
            __props__.__dict__["queue"] = queue
            __props__.__dict__["sleep"] = sleep
            __props__.__dict__["timeout"] = timeout
            if vhost is None and not opts.urn:
                raise TypeError("Missing required property 'vhost'")
            __props__.__dict__["vhost"] = vhost
            if webhook_uri is None and not opts.urn:
                raise TypeError("Missing required property 'webhook_uri'")
            __props__.__dict__["webhook_uri"] = webhook_uri
        super(Webhook, __self__).__init__(
            'cloudamqp:index/webhook:Webhook',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            concurrency: Optional[pulumi.Input[builtins.int]] = None,
            instance_id: Optional[pulumi.Input[builtins.int]] = None,
            queue: Optional[pulumi.Input[builtins.str]] = None,
            sleep: Optional[pulumi.Input[builtins.int]] = None,
            timeout: Optional[pulumi.Input[builtins.int]] = None,
            vhost: Optional[pulumi.Input[builtins.str]] = None,
            webhook_uri: Optional[pulumi.Input[builtins.str]] = None) -> 'Webhook':
        """
        Get an existing Webhook resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.int] concurrency: Max simultaneous requests to the endpoint.
        :param pulumi.Input[builtins.int] instance_id: The CloudAMQP instance ID.
        :param pulumi.Input[builtins.str] queue: A (durable) queue on your RabbitMQ instance.
        :param pulumi.Input[builtins.int] sleep: Configurable sleep time in seconds between retries for webhook
        :param pulumi.Input[builtins.int] timeout: Configurable timeout time in seconds for webhook
        :param pulumi.Input[builtins.str] vhost: The vhost the queue resides in.
        :param pulumi.Input[builtins.str] webhook_uri: A POST request will be made for each message in the queue to this
               endpoint.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _WebhookState.__new__(_WebhookState)

        __props__.__dict__["concurrency"] = concurrency
        __props__.__dict__["instance_id"] = instance_id
        __props__.__dict__["queue"] = queue
        __props__.__dict__["sleep"] = sleep
        __props__.__dict__["timeout"] = timeout
        __props__.__dict__["vhost"] = vhost
        __props__.__dict__["webhook_uri"] = webhook_uri
        return Webhook(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def concurrency(self) -> pulumi.Output[builtins.int]:
        """
        Max simultaneous requests to the endpoint.
        """
        return pulumi.get(self, "concurrency")

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> pulumi.Output[builtins.int]:
        """
        The CloudAMQP instance ID.
        """
        return pulumi.get(self, "instance_id")

    @property
    @pulumi.getter
    def queue(self) -> pulumi.Output[builtins.str]:
        """
        A (durable) queue on your RabbitMQ instance.
        """
        return pulumi.get(self, "queue")

    @property
    @pulumi.getter
    def sleep(self) -> pulumi.Output[Optional[builtins.int]]:
        """
        Configurable sleep time in seconds between retries for webhook
        """
        return pulumi.get(self, "sleep")

    @property
    @pulumi.getter
    def timeout(self) -> pulumi.Output[Optional[builtins.int]]:
        """
        Configurable timeout time in seconds for webhook
        """
        return pulumi.get(self, "timeout")

    @property
    @pulumi.getter
    def vhost(self) -> pulumi.Output[builtins.str]:
        """
        The vhost the queue resides in.
        """
        return pulumi.get(self, "vhost")

    @property
    @pulumi.getter(name="webhookUri")
    def webhook_uri(self) -> pulumi.Output[builtins.str]:
        """
        A POST request will be made for each message in the queue to this
        endpoint.
        """
        return pulumi.get(self, "webhook_uri")

