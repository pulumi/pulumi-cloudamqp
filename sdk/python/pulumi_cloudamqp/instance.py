# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['InstanceArgs', 'Instance']

@pulumi.input_type
class InstanceArgs:
    def __init__(__self__, *,
                 plan: pulumi.Input[str],
                 region: pulumi.Input[str],
                 keep_associated_vpc: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 no_default_alarms: Optional[pulumi.Input[bool]] = None,
                 nodes: Optional[pulumi.Input[int]] = None,
                 rmq_version: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 vpc_id: Optional[pulumi.Input[int]] = None,
                 vpc_subnet: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Instance resource.
        :param pulumi.Input[str] plan: The subscription plan. See available plans
        :param pulumi.Input[str] region: The region to host the instance in. See Instance regions
        :param pulumi.Input[bool] keep_associated_vpc: Keep associated VPC when deleting instance, default set to false.
        :param pulumi.Input[str] name: Name of the CloudAMQP instance.
        :param pulumi.Input[bool] no_default_alarms: Set to true to discard creating default alarms when the instance is created. Can be left out, will then use default value = false.
        :param pulumi.Input[int] nodes: Number of nodes, 1, 3 or 5 depending on plan used.
        :param pulumi.Input[str] rmq_version: The Rabbit MQ version. Can be left out, will then be set to default value used by CloudAMQP API.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] tags: One or more tags for the CloudAMQP instance, makes it possible to categories multiple instances in console view. Default there is no tags assigned.
        :param pulumi.Input[int] vpc_id: The VPC ID. Use this to create your instance in an existing VPC. See available example.
        :param pulumi.Input[str] vpc_subnet: Creates a dedicated VPC subnet, shouldn't overlap with other VPC subnet, default subnet used 10.56.72.0/24.
        """
        pulumi.set(__self__, "plan", plan)
        pulumi.set(__self__, "region", region)
        if keep_associated_vpc is not None:
            pulumi.set(__self__, "keep_associated_vpc", keep_associated_vpc)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if no_default_alarms is not None:
            pulumi.set(__self__, "no_default_alarms", no_default_alarms)
        if nodes is not None:
            pulumi.set(__self__, "nodes", nodes)
        if rmq_version is not None:
            pulumi.set(__self__, "rmq_version", rmq_version)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if vpc_id is not None:
            pulumi.set(__self__, "vpc_id", vpc_id)
        if vpc_subnet is not None:
            pulumi.set(__self__, "vpc_subnet", vpc_subnet)

    @property
    @pulumi.getter
    def plan(self) -> pulumi.Input[str]:
        """
        The subscription plan. See available plans
        """
        return pulumi.get(self, "plan")

    @plan.setter
    def plan(self, value: pulumi.Input[str]):
        pulumi.set(self, "plan", value)

    @property
    @pulumi.getter
    def region(self) -> pulumi.Input[str]:
        """
        The region to host the instance in. See Instance regions
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: pulumi.Input[str]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter(name="keepAssociatedVpc")
    def keep_associated_vpc(self) -> Optional[pulumi.Input[bool]]:
        """
        Keep associated VPC when deleting instance, default set to false.
        """
        return pulumi.get(self, "keep_associated_vpc")

    @keep_associated_vpc.setter
    def keep_associated_vpc(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "keep_associated_vpc", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the CloudAMQP instance.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="noDefaultAlarms")
    def no_default_alarms(self) -> Optional[pulumi.Input[bool]]:
        """
        Set to true to discard creating default alarms when the instance is created. Can be left out, will then use default value = false.
        """
        return pulumi.get(self, "no_default_alarms")

    @no_default_alarms.setter
    def no_default_alarms(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "no_default_alarms", value)

    @property
    @pulumi.getter
    def nodes(self) -> Optional[pulumi.Input[int]]:
        """
        Number of nodes, 1, 3 or 5 depending on plan used.
        """
        return pulumi.get(self, "nodes")

    @nodes.setter
    def nodes(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "nodes", value)

    @property
    @pulumi.getter(name="rmqVersion")
    def rmq_version(self) -> Optional[pulumi.Input[str]]:
        """
        The Rabbit MQ version. Can be left out, will then be set to default value used by CloudAMQP API.
        """
        return pulumi.get(self, "rmq_version")

    @rmq_version.setter
    def rmq_version(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "rmq_version", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        One or more tags for the CloudAMQP instance, makes it possible to categories multiple instances in console view. Default there is no tags assigned.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)

    @property
    @pulumi.getter(name="vpcId")
    def vpc_id(self) -> Optional[pulumi.Input[int]]:
        """
        The VPC ID. Use this to create your instance in an existing VPC. See available example.
        """
        return pulumi.get(self, "vpc_id")

    @vpc_id.setter
    def vpc_id(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "vpc_id", value)

    @property
    @pulumi.getter(name="vpcSubnet")
    def vpc_subnet(self) -> Optional[pulumi.Input[str]]:
        """
        Creates a dedicated VPC subnet, shouldn't overlap with other VPC subnet, default subnet used 10.56.72.0/24.
        """
        return pulumi.get(self, "vpc_subnet")

    @vpc_subnet.setter
    def vpc_subnet(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vpc_subnet", value)


@pulumi.input_type
class _InstanceState:
    def __init__(__self__, *,
                 apikey: Optional[pulumi.Input[str]] = None,
                 dedicated: Optional[pulumi.Input[bool]] = None,
                 host: Optional[pulumi.Input[str]] = None,
                 host_internal: Optional[pulumi.Input[str]] = None,
                 keep_associated_vpc: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 no_default_alarms: Optional[pulumi.Input[bool]] = None,
                 nodes: Optional[pulumi.Input[int]] = None,
                 plan: Optional[pulumi.Input[str]] = None,
                 ready: Optional[pulumi.Input[bool]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 rmq_version: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 url: Optional[pulumi.Input[str]] = None,
                 vhost: Optional[pulumi.Input[str]] = None,
                 vpc_id: Optional[pulumi.Input[int]] = None,
                 vpc_subnet: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Instance resources.
        :param pulumi.Input[str] apikey: API key needed to communicate to CloudAMQP's second API. The second API is used to manage alarms, integration and more, full description [CloudAMQP API](https://docs.cloudamqp.com/cloudamqp_api.html).
        :param pulumi.Input[bool] dedicated: Is the instance hosted on a dedicated server
        :param pulumi.Input[str] host: The external hostname for the CloudAMQP instance.
        :param pulumi.Input[str] host_internal: The internal hostname for the CloudAMQP instance.
        :param pulumi.Input[bool] keep_associated_vpc: Keep associated VPC when deleting instance, default set to false.
        :param pulumi.Input[str] name: Name of the CloudAMQP instance.
        :param pulumi.Input[bool] no_default_alarms: Set to true to discard creating default alarms when the instance is created. Can be left out, will then use default value = false.
        :param pulumi.Input[int] nodes: Number of nodes, 1, 3 or 5 depending on plan used.
        :param pulumi.Input[str] plan: The subscription plan. See available plans
        :param pulumi.Input[bool] ready: Flag describing if the resource is ready
        :param pulumi.Input[str] region: The region to host the instance in. See Instance regions
        :param pulumi.Input[str] rmq_version: The Rabbit MQ version. Can be left out, will then be set to default value used by CloudAMQP API.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] tags: One or more tags for the CloudAMQP instance, makes it possible to categories multiple instances in console view. Default there is no tags assigned.
        :param pulumi.Input[str] url: The AMQP URL (uses the internal hostname if the instance was created with VPC). Has the format: `amqps://{username}:{password}@{hostname}/{vhost}`
        :param pulumi.Input[str] vhost: The virtual host used by Rabbit MQ.
        :param pulumi.Input[int] vpc_id: The VPC ID. Use this to create your instance in an existing VPC. See available example.
        :param pulumi.Input[str] vpc_subnet: Creates a dedicated VPC subnet, shouldn't overlap with other VPC subnet, default subnet used 10.56.72.0/24.
        """
        if apikey is not None:
            pulumi.set(__self__, "apikey", apikey)
        if dedicated is not None:
            pulumi.set(__self__, "dedicated", dedicated)
        if host is not None:
            pulumi.set(__self__, "host", host)
        if host_internal is not None:
            pulumi.set(__self__, "host_internal", host_internal)
        if keep_associated_vpc is not None:
            pulumi.set(__self__, "keep_associated_vpc", keep_associated_vpc)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if no_default_alarms is not None:
            pulumi.set(__self__, "no_default_alarms", no_default_alarms)
        if nodes is not None:
            pulumi.set(__self__, "nodes", nodes)
        if plan is not None:
            pulumi.set(__self__, "plan", plan)
        if ready is not None:
            pulumi.set(__self__, "ready", ready)
        if region is not None:
            pulumi.set(__self__, "region", region)
        if rmq_version is not None:
            pulumi.set(__self__, "rmq_version", rmq_version)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if url is not None:
            pulumi.set(__self__, "url", url)
        if vhost is not None:
            pulumi.set(__self__, "vhost", vhost)
        if vpc_id is not None:
            pulumi.set(__self__, "vpc_id", vpc_id)
        if vpc_subnet is not None:
            pulumi.set(__self__, "vpc_subnet", vpc_subnet)

    @property
    @pulumi.getter
    def apikey(self) -> Optional[pulumi.Input[str]]:
        """
        API key needed to communicate to CloudAMQP's second API. The second API is used to manage alarms, integration and more, full description [CloudAMQP API](https://docs.cloudamqp.com/cloudamqp_api.html).
        """
        return pulumi.get(self, "apikey")

    @apikey.setter
    def apikey(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "apikey", value)

    @property
    @pulumi.getter
    def dedicated(self) -> Optional[pulumi.Input[bool]]:
        """
        Is the instance hosted on a dedicated server
        """
        return pulumi.get(self, "dedicated")

    @dedicated.setter
    def dedicated(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "dedicated", value)

    @property
    @pulumi.getter
    def host(self) -> Optional[pulumi.Input[str]]:
        """
        The external hostname for the CloudAMQP instance.
        """
        return pulumi.get(self, "host")

    @host.setter
    def host(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "host", value)

    @property
    @pulumi.getter(name="hostInternal")
    def host_internal(self) -> Optional[pulumi.Input[str]]:
        """
        The internal hostname for the CloudAMQP instance.
        """
        return pulumi.get(self, "host_internal")

    @host_internal.setter
    def host_internal(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "host_internal", value)

    @property
    @pulumi.getter(name="keepAssociatedVpc")
    def keep_associated_vpc(self) -> Optional[pulumi.Input[bool]]:
        """
        Keep associated VPC when deleting instance, default set to false.
        """
        return pulumi.get(self, "keep_associated_vpc")

    @keep_associated_vpc.setter
    def keep_associated_vpc(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "keep_associated_vpc", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the CloudAMQP instance.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="noDefaultAlarms")
    def no_default_alarms(self) -> Optional[pulumi.Input[bool]]:
        """
        Set to true to discard creating default alarms when the instance is created. Can be left out, will then use default value = false.
        """
        return pulumi.get(self, "no_default_alarms")

    @no_default_alarms.setter
    def no_default_alarms(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "no_default_alarms", value)

    @property
    @pulumi.getter
    def nodes(self) -> Optional[pulumi.Input[int]]:
        """
        Number of nodes, 1, 3 or 5 depending on plan used.
        """
        return pulumi.get(self, "nodes")

    @nodes.setter
    def nodes(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "nodes", value)

    @property
    @pulumi.getter
    def plan(self) -> Optional[pulumi.Input[str]]:
        """
        The subscription plan. See available plans
        """
        return pulumi.get(self, "plan")

    @plan.setter
    def plan(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "plan", value)

    @property
    @pulumi.getter
    def ready(self) -> Optional[pulumi.Input[bool]]:
        """
        Flag describing if the resource is ready
        """
        return pulumi.get(self, "ready")

    @ready.setter
    def ready(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "ready", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        """
        The region to host the instance in. See Instance regions
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter(name="rmqVersion")
    def rmq_version(self) -> Optional[pulumi.Input[str]]:
        """
        The Rabbit MQ version. Can be left out, will then be set to default value used by CloudAMQP API.
        """
        return pulumi.get(self, "rmq_version")

    @rmq_version.setter
    def rmq_version(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "rmq_version", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        One or more tags for the CloudAMQP instance, makes it possible to categories multiple instances in console view. Default there is no tags assigned.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)

    @property
    @pulumi.getter
    def url(self) -> Optional[pulumi.Input[str]]:
        """
        The AMQP URL (uses the internal hostname if the instance was created with VPC). Has the format: `amqps://{username}:{password}@{hostname}/{vhost}`
        """
        return pulumi.get(self, "url")

    @url.setter
    def url(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "url", value)

    @property
    @pulumi.getter
    def vhost(self) -> Optional[pulumi.Input[str]]:
        """
        The virtual host used by Rabbit MQ.
        """
        return pulumi.get(self, "vhost")

    @vhost.setter
    def vhost(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vhost", value)

    @property
    @pulumi.getter(name="vpcId")
    def vpc_id(self) -> Optional[pulumi.Input[int]]:
        """
        The VPC ID. Use this to create your instance in an existing VPC. See available example.
        """
        return pulumi.get(self, "vpc_id")

    @vpc_id.setter
    def vpc_id(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "vpc_id", value)

    @property
    @pulumi.getter(name="vpcSubnet")
    def vpc_subnet(self) -> Optional[pulumi.Input[str]]:
        """
        Creates a dedicated VPC subnet, shouldn't overlap with other VPC subnet, default subnet used 10.56.72.0/24.
        """
        return pulumi.get(self, "vpc_subnet")

    @vpc_subnet.setter
    def vpc_subnet(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vpc_subnet", value)


class Instance(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 keep_associated_vpc: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 no_default_alarms: Optional[pulumi.Input[bool]] = None,
                 nodes: Optional[pulumi.Input[int]] = None,
                 plan: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 rmq_version: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 vpc_id: Optional[pulumi.Input[int]] = None,
                 vpc_subnet: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        ## Import

        `cloudamqp_instance`can be imported using CloudAMQP internal identifier.

        ```sh
         $ pulumi import cloudamqp:index/instance:Instance instance <id>`
        ```

         To retrieve the identifier for a VPC, either use [CloudAMQP customer API](https://docs.cloudamqp.com/#list-instances). Or use the data source [`cloudamqp_account`](https://registry.terraform.io/providers/cloudamqp/cloudamqp/latest/docs/data-sources/account) to list all available instances for an account.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] keep_associated_vpc: Keep associated VPC when deleting instance, default set to false.
        :param pulumi.Input[str] name: Name of the CloudAMQP instance.
        :param pulumi.Input[bool] no_default_alarms: Set to true to discard creating default alarms when the instance is created. Can be left out, will then use default value = false.
        :param pulumi.Input[int] nodes: Number of nodes, 1, 3 or 5 depending on plan used.
        :param pulumi.Input[str] plan: The subscription plan. See available plans
        :param pulumi.Input[str] region: The region to host the instance in. See Instance regions
        :param pulumi.Input[str] rmq_version: The Rabbit MQ version. Can be left out, will then be set to default value used by CloudAMQP API.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] tags: One or more tags for the CloudAMQP instance, makes it possible to categories multiple instances in console view. Default there is no tags assigned.
        :param pulumi.Input[int] vpc_id: The VPC ID. Use this to create your instance in an existing VPC. See available example.
        :param pulumi.Input[str] vpc_subnet: Creates a dedicated VPC subnet, shouldn't overlap with other VPC subnet, default subnet used 10.56.72.0/24.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: InstanceArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        ## Import

        `cloudamqp_instance`can be imported using CloudAMQP internal identifier.

        ```sh
         $ pulumi import cloudamqp:index/instance:Instance instance <id>`
        ```

         To retrieve the identifier for a VPC, either use [CloudAMQP customer API](https://docs.cloudamqp.com/#list-instances). Or use the data source [`cloudamqp_account`](https://registry.terraform.io/providers/cloudamqp/cloudamqp/latest/docs/data-sources/account) to list all available instances for an account.

        :param str resource_name: The name of the resource.
        :param InstanceArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(InstanceArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 keep_associated_vpc: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 no_default_alarms: Optional[pulumi.Input[bool]] = None,
                 nodes: Optional[pulumi.Input[int]] = None,
                 plan: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 rmq_version: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 vpc_id: Optional[pulumi.Input[int]] = None,
                 vpc_subnet: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = InstanceArgs.__new__(InstanceArgs)

            __props__.__dict__["keep_associated_vpc"] = keep_associated_vpc
            __props__.__dict__["name"] = name
            __props__.__dict__["no_default_alarms"] = no_default_alarms
            __props__.__dict__["nodes"] = nodes
            if plan is None and not opts.urn:
                raise TypeError("Missing required property 'plan'")
            __props__.__dict__["plan"] = plan
            if region is None and not opts.urn:
                raise TypeError("Missing required property 'region'")
            __props__.__dict__["region"] = region
            __props__.__dict__["rmq_version"] = rmq_version
            __props__.__dict__["tags"] = tags
            __props__.__dict__["vpc_id"] = vpc_id
            __props__.__dict__["vpc_subnet"] = vpc_subnet
            __props__.__dict__["apikey"] = None
            __props__.__dict__["dedicated"] = None
            __props__.__dict__["host"] = None
            __props__.__dict__["host_internal"] = None
            __props__.__dict__["ready"] = None
            __props__.__dict__["url"] = None
            __props__.__dict__["vhost"] = None
        super(Instance, __self__).__init__(
            'cloudamqp:index/instance:Instance',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            apikey: Optional[pulumi.Input[str]] = None,
            dedicated: Optional[pulumi.Input[bool]] = None,
            host: Optional[pulumi.Input[str]] = None,
            host_internal: Optional[pulumi.Input[str]] = None,
            keep_associated_vpc: Optional[pulumi.Input[bool]] = None,
            name: Optional[pulumi.Input[str]] = None,
            no_default_alarms: Optional[pulumi.Input[bool]] = None,
            nodes: Optional[pulumi.Input[int]] = None,
            plan: Optional[pulumi.Input[str]] = None,
            ready: Optional[pulumi.Input[bool]] = None,
            region: Optional[pulumi.Input[str]] = None,
            rmq_version: Optional[pulumi.Input[str]] = None,
            tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            url: Optional[pulumi.Input[str]] = None,
            vhost: Optional[pulumi.Input[str]] = None,
            vpc_id: Optional[pulumi.Input[int]] = None,
            vpc_subnet: Optional[pulumi.Input[str]] = None) -> 'Instance':
        """
        Get an existing Instance resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] apikey: API key needed to communicate to CloudAMQP's second API. The second API is used to manage alarms, integration and more, full description [CloudAMQP API](https://docs.cloudamqp.com/cloudamqp_api.html).
        :param pulumi.Input[bool] dedicated: Is the instance hosted on a dedicated server
        :param pulumi.Input[str] host: The external hostname for the CloudAMQP instance.
        :param pulumi.Input[str] host_internal: The internal hostname for the CloudAMQP instance.
        :param pulumi.Input[bool] keep_associated_vpc: Keep associated VPC when deleting instance, default set to false.
        :param pulumi.Input[str] name: Name of the CloudAMQP instance.
        :param pulumi.Input[bool] no_default_alarms: Set to true to discard creating default alarms when the instance is created. Can be left out, will then use default value = false.
        :param pulumi.Input[int] nodes: Number of nodes, 1, 3 or 5 depending on plan used.
        :param pulumi.Input[str] plan: The subscription plan. See available plans
        :param pulumi.Input[bool] ready: Flag describing if the resource is ready
        :param pulumi.Input[str] region: The region to host the instance in. See Instance regions
        :param pulumi.Input[str] rmq_version: The Rabbit MQ version. Can be left out, will then be set to default value used by CloudAMQP API.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] tags: One or more tags for the CloudAMQP instance, makes it possible to categories multiple instances in console view. Default there is no tags assigned.
        :param pulumi.Input[str] url: The AMQP URL (uses the internal hostname if the instance was created with VPC). Has the format: `amqps://{username}:{password}@{hostname}/{vhost}`
        :param pulumi.Input[str] vhost: The virtual host used by Rabbit MQ.
        :param pulumi.Input[int] vpc_id: The VPC ID. Use this to create your instance in an existing VPC. See available example.
        :param pulumi.Input[str] vpc_subnet: Creates a dedicated VPC subnet, shouldn't overlap with other VPC subnet, default subnet used 10.56.72.0/24.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _InstanceState.__new__(_InstanceState)

        __props__.__dict__["apikey"] = apikey
        __props__.__dict__["dedicated"] = dedicated
        __props__.__dict__["host"] = host
        __props__.__dict__["host_internal"] = host_internal
        __props__.__dict__["keep_associated_vpc"] = keep_associated_vpc
        __props__.__dict__["name"] = name
        __props__.__dict__["no_default_alarms"] = no_default_alarms
        __props__.__dict__["nodes"] = nodes
        __props__.__dict__["plan"] = plan
        __props__.__dict__["ready"] = ready
        __props__.__dict__["region"] = region
        __props__.__dict__["rmq_version"] = rmq_version
        __props__.__dict__["tags"] = tags
        __props__.__dict__["url"] = url
        __props__.__dict__["vhost"] = vhost
        __props__.__dict__["vpc_id"] = vpc_id
        __props__.__dict__["vpc_subnet"] = vpc_subnet
        return Instance(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def apikey(self) -> pulumi.Output[str]:
        """
        API key needed to communicate to CloudAMQP's second API. The second API is used to manage alarms, integration and more, full description [CloudAMQP API](https://docs.cloudamqp.com/cloudamqp_api.html).
        """
        return pulumi.get(self, "apikey")

    @property
    @pulumi.getter
    def dedicated(self) -> pulumi.Output[bool]:
        """
        Is the instance hosted on a dedicated server
        """
        return pulumi.get(self, "dedicated")

    @property
    @pulumi.getter
    def host(self) -> pulumi.Output[str]:
        """
        The external hostname for the CloudAMQP instance.
        """
        return pulumi.get(self, "host")

    @property
    @pulumi.getter(name="hostInternal")
    def host_internal(self) -> pulumi.Output[str]:
        """
        The internal hostname for the CloudAMQP instance.
        """
        return pulumi.get(self, "host_internal")

    @property
    @pulumi.getter(name="keepAssociatedVpc")
    def keep_associated_vpc(self) -> pulumi.Output[Optional[bool]]:
        """
        Keep associated VPC when deleting instance, default set to false.
        """
        return pulumi.get(self, "keep_associated_vpc")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Name of the CloudAMQP instance.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="noDefaultAlarms")
    def no_default_alarms(self) -> pulumi.Output[bool]:
        """
        Set to true to discard creating default alarms when the instance is created. Can be left out, will then use default value = false.
        """
        return pulumi.get(self, "no_default_alarms")

    @property
    @pulumi.getter
    def nodes(self) -> pulumi.Output[int]:
        """
        Number of nodes, 1, 3 or 5 depending on plan used.
        """
        return pulumi.get(self, "nodes")

    @property
    @pulumi.getter
    def plan(self) -> pulumi.Output[str]:
        """
        The subscription plan. See available plans
        """
        return pulumi.get(self, "plan")

    @property
    @pulumi.getter
    def ready(self) -> pulumi.Output[bool]:
        """
        Flag describing if the resource is ready
        """
        return pulumi.get(self, "ready")

    @property
    @pulumi.getter
    def region(self) -> pulumi.Output[str]:
        """
        The region to host the instance in. See Instance regions
        """
        return pulumi.get(self, "region")

    @property
    @pulumi.getter(name="rmqVersion")
    def rmq_version(self) -> pulumi.Output[str]:
        """
        The Rabbit MQ version. Can be left out, will then be set to default value used by CloudAMQP API.
        """
        return pulumi.get(self, "rmq_version")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        One or more tags for the CloudAMQP instance, makes it possible to categories multiple instances in console view. Default there is no tags assigned.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def url(self) -> pulumi.Output[str]:
        """
        The AMQP URL (uses the internal hostname if the instance was created with VPC). Has the format: `amqps://{username}:{password}@{hostname}/{vhost}`
        """
        return pulumi.get(self, "url")

    @property
    @pulumi.getter
    def vhost(self) -> pulumi.Output[str]:
        """
        The virtual host used by Rabbit MQ.
        """
        return pulumi.get(self, "vhost")

    @property
    @pulumi.getter(name="vpcId")
    def vpc_id(self) -> pulumi.Output[int]:
        """
        The VPC ID. Use this to create your instance in an existing VPC. See available example.
        """
        return pulumi.get(self, "vpc_id")

    @property
    @pulumi.getter(name="vpcSubnet")
    def vpc_subnet(self) -> pulumi.Output[str]:
        """
        Creates a dedicated VPC subnet, shouldn't overlap with other VPC subnet, default subnet used 10.56.72.0/24.
        """
        return pulumi.get(self, "vpc_subnet")

