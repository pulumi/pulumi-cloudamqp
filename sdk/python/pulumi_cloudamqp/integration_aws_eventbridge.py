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

__all__ = ['IntegrationAwsEventbridgeArgs', 'IntegrationAwsEventbridge']

@pulumi.input_type
class IntegrationAwsEventbridgeArgs:
    def __init__(__self__, *,
                 aws_account_id: pulumi.Input[builtins.str],
                 aws_region: pulumi.Input[builtins.str],
                 instance_id: pulumi.Input[builtins.int],
                 queue: pulumi.Input[builtins.str],
                 vhost: pulumi.Input[builtins.str],
                 with_headers: pulumi.Input[builtins.bool]):
        """
        The set of arguments for constructing a IntegrationAwsEventbridge resource.
        :param pulumi.Input[builtins.str] aws_account_id: The 12 digit AWS Account ID where you want the events to be sent to.
        :param pulumi.Input[builtins.str] aws_region: The AWS region where you the events to be sent to. (e.g. us-west-1, us-west-2, ..., etc.)
        :param pulumi.Input[builtins.int] instance_id: Instance identifier
        :param pulumi.Input[builtins.str] queue: A (durable) queue on your RabbitMQ instance.
        :param pulumi.Input[builtins.str] vhost: The VHost the queue resides in.
        :param pulumi.Input[builtins.bool] with_headers: Include message headers in the event data.
        """
        pulumi.set(__self__, "aws_account_id", aws_account_id)
        pulumi.set(__self__, "aws_region", aws_region)
        pulumi.set(__self__, "instance_id", instance_id)
        pulumi.set(__self__, "queue", queue)
        pulumi.set(__self__, "vhost", vhost)
        pulumi.set(__self__, "with_headers", with_headers)

    @property
    @pulumi.getter(name="awsAccountId")
    def aws_account_id(self) -> pulumi.Input[builtins.str]:
        """
        The 12 digit AWS Account ID where you want the events to be sent to.
        """
        return pulumi.get(self, "aws_account_id")

    @aws_account_id.setter
    def aws_account_id(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "aws_account_id", value)

    @property
    @pulumi.getter(name="awsRegion")
    def aws_region(self) -> pulumi.Input[builtins.str]:
        """
        The AWS region where you the events to be sent to. (e.g. us-west-1, us-west-2, ..., etc.)
        """
        return pulumi.get(self, "aws_region")

    @aws_region.setter
    def aws_region(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "aws_region", value)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> pulumi.Input[builtins.int]:
        """
        Instance identifier
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
        The VHost the queue resides in.
        """
        return pulumi.get(self, "vhost")

    @vhost.setter
    def vhost(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "vhost", value)

    @property
    @pulumi.getter(name="withHeaders")
    def with_headers(self) -> pulumi.Input[builtins.bool]:
        """
        Include message headers in the event data.
        """
        return pulumi.get(self, "with_headers")

    @with_headers.setter
    def with_headers(self, value: pulumi.Input[builtins.bool]):
        pulumi.set(self, "with_headers", value)


@pulumi.input_type
class _IntegrationAwsEventbridgeState:
    def __init__(__self__, *,
                 aws_account_id: Optional[pulumi.Input[builtins.str]] = None,
                 aws_region: Optional[pulumi.Input[builtins.str]] = None,
                 instance_id: Optional[pulumi.Input[builtins.int]] = None,
                 queue: Optional[pulumi.Input[builtins.str]] = None,
                 status: Optional[pulumi.Input[builtins.str]] = None,
                 vhost: Optional[pulumi.Input[builtins.str]] = None,
                 with_headers: Optional[pulumi.Input[builtins.bool]] = None):
        """
        Input properties used for looking up and filtering IntegrationAwsEventbridge resources.
        :param pulumi.Input[builtins.str] aws_account_id: The 12 digit AWS Account ID where you want the events to be sent to.
        :param pulumi.Input[builtins.str] aws_region: The AWS region where you the events to be sent to. (e.g. us-west-1, us-west-2, ..., etc.)
        :param pulumi.Input[builtins.int] instance_id: Instance identifier
        :param pulumi.Input[builtins.str] queue: A (durable) queue on your RabbitMQ instance.
        :param pulumi.Input[builtins.str] status: Always set to null, unless there is an error starting the EventBridge.
        :param pulumi.Input[builtins.str] vhost: The VHost the queue resides in.
        :param pulumi.Input[builtins.bool] with_headers: Include message headers in the event data.
        """
        if aws_account_id is not None:
            pulumi.set(__self__, "aws_account_id", aws_account_id)
        if aws_region is not None:
            pulumi.set(__self__, "aws_region", aws_region)
        if instance_id is not None:
            pulumi.set(__self__, "instance_id", instance_id)
        if queue is not None:
            pulumi.set(__self__, "queue", queue)
        if status is not None:
            pulumi.set(__self__, "status", status)
        if vhost is not None:
            pulumi.set(__self__, "vhost", vhost)
        if with_headers is not None:
            pulumi.set(__self__, "with_headers", with_headers)

    @property
    @pulumi.getter(name="awsAccountId")
    def aws_account_id(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The 12 digit AWS Account ID where you want the events to be sent to.
        """
        return pulumi.get(self, "aws_account_id")

    @aws_account_id.setter
    def aws_account_id(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "aws_account_id", value)

    @property
    @pulumi.getter(name="awsRegion")
    def aws_region(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The AWS region where you the events to be sent to. (e.g. us-west-1, us-west-2, ..., etc.)
        """
        return pulumi.get(self, "aws_region")

    @aws_region.setter
    def aws_region(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "aws_region", value)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> Optional[pulumi.Input[builtins.int]]:
        """
        Instance identifier
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
    def status(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Always set to null, unless there is an error starting the EventBridge.
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "status", value)

    @property
    @pulumi.getter
    def vhost(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The VHost the queue resides in.
        """
        return pulumi.get(self, "vhost")

    @vhost.setter
    def vhost(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "vhost", value)

    @property
    @pulumi.getter(name="withHeaders")
    def with_headers(self) -> Optional[pulumi.Input[builtins.bool]]:
        """
        Include message headers in the event data.
        """
        return pulumi.get(self, "with_headers")

    @with_headers.setter
    def with_headers(self, value: Optional[pulumi.Input[builtins.bool]]):
        pulumi.set(self, "with_headers", value)


class IntegrationAwsEventbridge(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 aws_account_id: Optional[pulumi.Input[builtins.str]] = None,
                 aws_region: Optional[pulumi.Input[builtins.str]] = None,
                 instance_id: Optional[pulumi.Input[builtins.int]] = None,
                 queue: Optional[pulumi.Input[builtins.str]] = None,
                 vhost: Optional[pulumi.Input[builtins.str]] = None,
                 with_headers: Optional[pulumi.Input[builtins.bool]] = None,
                 __props__=None):
        """
        This resource allows you to create and manage, an [AWS EventBridge] for a CloudAMQP instance. Once
        created, continue to map the EventBridge in the [AWS Eventbridge console].

        >  Our consumer needs to have exclusive usage to the configured queue and the maximum body size
        allowed on msgs by AWS is 256kb. The message body has to be valid JSON for AWS Eventbridge to accept
        it. If messages are too large or are not valid JSON, they will be rejected (tip: setup a dead-letter
        queue to catch them).

        Not possible to update this resource. Any changes made to the argument will destroy and recreate the
        resource. Hence why all arguments use ForceNew.

        Only available for dedicated subscription plans.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_cloudamqp as cloudamqp

        instance = cloudamqp.Instance("instance",
            name="Test instance",
            plan="penguin-1",
            region="amazon-web-services::us-west-1",
            rmq_version="3.11.5",
            tags=["aws"])
        this = cloudamqp.IntegrationAwsEventbridge("this",
            instance_id=instance.id,
            vhost=instance.vhost,
            queue="<QUEUE-NAME>",
            aws_account_id="<AWS-ACCOUNT-ID>",
            aws_region="us-west-1",
            with_headers=True)
        ```

        ## Argument References

        The following arguments are supported:

        * `aws_account_id`  - (ForceNew/Required) The 12 digit AWS Account ID where you want the events to
                              be sent to.
        * `aws_region`      - (ForceNew/Required) The AWS region where you the events to be sent to.
                              (e.g. us-west-1, us-west-2, ..., etc.)
        * `vhost`           - (ForceNew/Required) The VHost the queue resides in.
        * `queue`           - (ForceNew/Required) A (durable) queue on your RabbitMQ instance.
        * `with_headers`    - (ForceNew/Required) Include message headers in the event data.
                              `({ "headers": { }, "body": { "your": "message" } })`

        ## Dependency

        This resource depends on CloudAMQP instance identifier, `cloudamqp_instance.instance.id`.

        ## Import

        `cloudamqp_integration_aws_eventbridge` can be imported using the resource identifier together with

        CloudAMQP instance identifier (CSV separated). To retrieve the resource identifier, use

        [CloudAMQP API list eventbridges].

        From Terraform v1.5.0, the `import` block can be used to import this resource:

        hcl

        import {

          to = cloudamqp_integration_aws_eventbridge.this

          id = format("<id>,%s", cloudamqp_instance.instance.id)

        }

        Or with Terraform CLI:

        ```sh
        $ pulumi import cloudamqp:index/integrationAwsEventbridge:IntegrationAwsEventbridge this <id>,<instance_id>`
        ```

        [AWS EventBridge]: https://aws.amazon.com/eventbridge

        [AWS Eventbridge console]: https://console.aws.amazon.com/events/home

        [CloudAMQP API list eventbridges]: https://docs.cloudamqp.com/cloudamqp_api.html#list-eventbridges

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.str] aws_account_id: The 12 digit AWS Account ID where you want the events to be sent to.
        :param pulumi.Input[builtins.str] aws_region: The AWS region where you the events to be sent to. (e.g. us-west-1, us-west-2, ..., etc.)
        :param pulumi.Input[builtins.int] instance_id: Instance identifier
        :param pulumi.Input[builtins.str] queue: A (durable) queue on your RabbitMQ instance.
        :param pulumi.Input[builtins.str] vhost: The VHost the queue resides in.
        :param pulumi.Input[builtins.bool] with_headers: Include message headers in the event data.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: IntegrationAwsEventbridgeArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        This resource allows you to create and manage, an [AWS EventBridge] for a CloudAMQP instance. Once
        created, continue to map the EventBridge in the [AWS Eventbridge console].

        >  Our consumer needs to have exclusive usage to the configured queue and the maximum body size
        allowed on msgs by AWS is 256kb. The message body has to be valid JSON for AWS Eventbridge to accept
        it. If messages are too large or are not valid JSON, they will be rejected (tip: setup a dead-letter
        queue to catch them).

        Not possible to update this resource. Any changes made to the argument will destroy and recreate the
        resource. Hence why all arguments use ForceNew.

        Only available for dedicated subscription plans.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_cloudamqp as cloudamqp

        instance = cloudamqp.Instance("instance",
            name="Test instance",
            plan="penguin-1",
            region="amazon-web-services::us-west-1",
            rmq_version="3.11.5",
            tags=["aws"])
        this = cloudamqp.IntegrationAwsEventbridge("this",
            instance_id=instance.id,
            vhost=instance.vhost,
            queue="<QUEUE-NAME>",
            aws_account_id="<AWS-ACCOUNT-ID>",
            aws_region="us-west-1",
            with_headers=True)
        ```

        ## Argument References

        The following arguments are supported:

        * `aws_account_id`  - (ForceNew/Required) The 12 digit AWS Account ID where you want the events to
                              be sent to.
        * `aws_region`      - (ForceNew/Required) The AWS region where you the events to be sent to.
                              (e.g. us-west-1, us-west-2, ..., etc.)
        * `vhost`           - (ForceNew/Required) The VHost the queue resides in.
        * `queue`           - (ForceNew/Required) A (durable) queue on your RabbitMQ instance.
        * `with_headers`    - (ForceNew/Required) Include message headers in the event data.
                              `({ "headers": { }, "body": { "your": "message" } })`

        ## Dependency

        This resource depends on CloudAMQP instance identifier, `cloudamqp_instance.instance.id`.

        ## Import

        `cloudamqp_integration_aws_eventbridge` can be imported using the resource identifier together with

        CloudAMQP instance identifier (CSV separated). To retrieve the resource identifier, use

        [CloudAMQP API list eventbridges].

        From Terraform v1.5.0, the `import` block can be used to import this resource:

        hcl

        import {

          to = cloudamqp_integration_aws_eventbridge.this

          id = format("<id>,%s", cloudamqp_instance.instance.id)

        }

        Or with Terraform CLI:

        ```sh
        $ pulumi import cloudamqp:index/integrationAwsEventbridge:IntegrationAwsEventbridge this <id>,<instance_id>`
        ```

        [AWS EventBridge]: https://aws.amazon.com/eventbridge

        [AWS Eventbridge console]: https://console.aws.amazon.com/events/home

        [CloudAMQP API list eventbridges]: https://docs.cloudamqp.com/cloudamqp_api.html#list-eventbridges

        :param str resource_name: The name of the resource.
        :param IntegrationAwsEventbridgeArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(IntegrationAwsEventbridgeArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 aws_account_id: Optional[pulumi.Input[builtins.str]] = None,
                 aws_region: Optional[pulumi.Input[builtins.str]] = None,
                 instance_id: Optional[pulumi.Input[builtins.int]] = None,
                 queue: Optional[pulumi.Input[builtins.str]] = None,
                 vhost: Optional[pulumi.Input[builtins.str]] = None,
                 with_headers: Optional[pulumi.Input[builtins.bool]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = IntegrationAwsEventbridgeArgs.__new__(IntegrationAwsEventbridgeArgs)

            if aws_account_id is None and not opts.urn:
                raise TypeError("Missing required property 'aws_account_id'")
            __props__.__dict__["aws_account_id"] = aws_account_id
            if aws_region is None and not opts.urn:
                raise TypeError("Missing required property 'aws_region'")
            __props__.__dict__["aws_region"] = aws_region
            if instance_id is None and not opts.urn:
                raise TypeError("Missing required property 'instance_id'")
            __props__.__dict__["instance_id"] = instance_id
            if queue is None and not opts.urn:
                raise TypeError("Missing required property 'queue'")
            __props__.__dict__["queue"] = queue
            if vhost is None and not opts.urn:
                raise TypeError("Missing required property 'vhost'")
            __props__.__dict__["vhost"] = vhost
            if with_headers is None and not opts.urn:
                raise TypeError("Missing required property 'with_headers'")
            __props__.__dict__["with_headers"] = with_headers
            __props__.__dict__["status"] = None
        super(IntegrationAwsEventbridge, __self__).__init__(
            'cloudamqp:index/integrationAwsEventbridge:IntegrationAwsEventbridge',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            aws_account_id: Optional[pulumi.Input[builtins.str]] = None,
            aws_region: Optional[pulumi.Input[builtins.str]] = None,
            instance_id: Optional[pulumi.Input[builtins.int]] = None,
            queue: Optional[pulumi.Input[builtins.str]] = None,
            status: Optional[pulumi.Input[builtins.str]] = None,
            vhost: Optional[pulumi.Input[builtins.str]] = None,
            with_headers: Optional[pulumi.Input[builtins.bool]] = None) -> 'IntegrationAwsEventbridge':
        """
        Get an existing IntegrationAwsEventbridge resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.str] aws_account_id: The 12 digit AWS Account ID where you want the events to be sent to.
        :param pulumi.Input[builtins.str] aws_region: The AWS region where you the events to be sent to. (e.g. us-west-1, us-west-2, ..., etc.)
        :param pulumi.Input[builtins.int] instance_id: Instance identifier
        :param pulumi.Input[builtins.str] queue: A (durable) queue on your RabbitMQ instance.
        :param pulumi.Input[builtins.str] status: Always set to null, unless there is an error starting the EventBridge.
        :param pulumi.Input[builtins.str] vhost: The VHost the queue resides in.
        :param pulumi.Input[builtins.bool] with_headers: Include message headers in the event data.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _IntegrationAwsEventbridgeState.__new__(_IntegrationAwsEventbridgeState)

        __props__.__dict__["aws_account_id"] = aws_account_id
        __props__.__dict__["aws_region"] = aws_region
        __props__.__dict__["instance_id"] = instance_id
        __props__.__dict__["queue"] = queue
        __props__.__dict__["status"] = status
        __props__.__dict__["vhost"] = vhost
        __props__.__dict__["with_headers"] = with_headers
        return IntegrationAwsEventbridge(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="awsAccountId")
    def aws_account_id(self) -> pulumi.Output[builtins.str]:
        """
        The 12 digit AWS Account ID where you want the events to be sent to.
        """
        return pulumi.get(self, "aws_account_id")

    @property
    @pulumi.getter(name="awsRegion")
    def aws_region(self) -> pulumi.Output[builtins.str]:
        """
        The AWS region where you the events to be sent to. (e.g. us-west-1, us-west-2, ..., etc.)
        """
        return pulumi.get(self, "aws_region")

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> pulumi.Output[builtins.int]:
        """
        Instance identifier
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
    def status(self) -> pulumi.Output[builtins.str]:
        """
        Always set to null, unless there is an error starting the EventBridge.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter
    def vhost(self) -> pulumi.Output[builtins.str]:
        """
        The VHost the queue resides in.
        """
        return pulumi.get(self, "vhost")

    @property
    @pulumi.getter(name="withHeaders")
    def with_headers(self) -> pulumi.Output[builtins.bool]:
        """
        Include message headers in the event data.
        """
        return pulumi.get(self, "with_headers")

