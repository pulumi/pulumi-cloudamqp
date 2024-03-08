# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['VpcGcpPeeringArgs', 'VpcGcpPeering']

@pulumi.input_type
class VpcGcpPeeringArgs:
    def __init__(__self__, *,
                 peer_network_uri: pulumi.Input[str],
                 instance_id: Optional[pulumi.Input[int]] = None,
                 sleep: Optional[pulumi.Input[int]] = None,
                 timeout: Optional[pulumi.Input[int]] = None,
                 vpc_id: Optional[pulumi.Input[str]] = None,
                 wait_on_peering_status: Optional[pulumi.Input[bool]] = None):
        """
        The set of arguments for constructing a VpcGcpPeering resource.
        :param pulumi.Input[str] peer_network_uri: Network uri of the VPC network to which you will peer with.
        :param pulumi.Input[int] instance_id: The CloudAMQP instance identifier. *Deprecated from v1.16.0*
        :param pulumi.Input[int] sleep: Configurable sleep time (seconds) between retries when requesting or reading
               peering. Default set to 10 seconds. *Available from v1.29.0*
        :param pulumi.Input[int] timeout: Configurable timeout time (seconds) before retries times out. Default set
               to 1800 seconds. *Available from v1.29.0*
        :param pulumi.Input[str] vpc_id: The managed VPC identifier. *Available from v1.16.0*
        :param pulumi.Input[bool] wait_on_peering_status: Makes the resource wait until the peering is connected.
               Default set to false. *Available from v1.28.0*
        """
        pulumi.set(__self__, "peer_network_uri", peer_network_uri)
        if instance_id is not None:
            pulumi.set(__self__, "instance_id", instance_id)
        if sleep is not None:
            pulumi.set(__self__, "sleep", sleep)
        if timeout is not None:
            pulumi.set(__self__, "timeout", timeout)
        if vpc_id is not None:
            pulumi.set(__self__, "vpc_id", vpc_id)
        if wait_on_peering_status is not None:
            pulumi.set(__self__, "wait_on_peering_status", wait_on_peering_status)

    @property
    @pulumi.getter(name="peerNetworkUri")
    def peer_network_uri(self) -> pulumi.Input[str]:
        """
        Network uri of the VPC network to which you will peer with.
        """
        return pulumi.get(self, "peer_network_uri")

    @peer_network_uri.setter
    def peer_network_uri(self, value: pulumi.Input[str]):
        pulumi.set(self, "peer_network_uri", value)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> Optional[pulumi.Input[int]]:
        """
        The CloudAMQP instance identifier. *Deprecated from v1.16.0*
        """
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "instance_id", value)

    @property
    @pulumi.getter
    def sleep(self) -> Optional[pulumi.Input[int]]:
        """
        Configurable sleep time (seconds) between retries when requesting or reading
        peering. Default set to 10 seconds. *Available from v1.29.0*
        """
        return pulumi.get(self, "sleep")

    @sleep.setter
    def sleep(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "sleep", value)

    @property
    @pulumi.getter
    def timeout(self) -> Optional[pulumi.Input[int]]:
        """
        Configurable timeout time (seconds) before retries times out. Default set
        to 1800 seconds. *Available from v1.29.0*
        """
        return pulumi.get(self, "timeout")

    @timeout.setter
    def timeout(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "timeout", value)

    @property
    @pulumi.getter(name="vpcId")
    def vpc_id(self) -> Optional[pulumi.Input[str]]:
        """
        The managed VPC identifier. *Available from v1.16.0*
        """
        return pulumi.get(self, "vpc_id")

    @vpc_id.setter
    def vpc_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vpc_id", value)

    @property
    @pulumi.getter(name="waitOnPeeringStatus")
    def wait_on_peering_status(self) -> Optional[pulumi.Input[bool]]:
        """
        Makes the resource wait until the peering is connected.
        Default set to false. *Available from v1.28.0*
        """
        return pulumi.get(self, "wait_on_peering_status")

    @wait_on_peering_status.setter
    def wait_on_peering_status(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "wait_on_peering_status", value)


@pulumi.input_type
class _VpcGcpPeeringState:
    def __init__(__self__, *,
                 auto_create_routes: Optional[pulumi.Input[bool]] = None,
                 instance_id: Optional[pulumi.Input[int]] = None,
                 peer_network_uri: Optional[pulumi.Input[str]] = None,
                 sleep: Optional[pulumi.Input[int]] = None,
                 state: Optional[pulumi.Input[str]] = None,
                 state_details: Optional[pulumi.Input[str]] = None,
                 timeout: Optional[pulumi.Input[int]] = None,
                 vpc_id: Optional[pulumi.Input[str]] = None,
                 wait_on_peering_status: Optional[pulumi.Input[bool]] = None):
        """
        Input properties used for looking up and filtering VpcGcpPeering resources.
        :param pulumi.Input[bool] auto_create_routes: VPC peering auto created routes
        :param pulumi.Input[int] instance_id: The CloudAMQP instance identifier. *Deprecated from v1.16.0*
        :param pulumi.Input[str] peer_network_uri: Network uri of the VPC network to which you will peer with.
        :param pulumi.Input[int] sleep: Configurable sleep time (seconds) between retries when requesting or reading
               peering. Default set to 10 seconds. *Available from v1.29.0*
        :param pulumi.Input[str] state: VPC peering state
        :param pulumi.Input[str] state_details: VPC peering state details
        :param pulumi.Input[int] timeout: Configurable timeout time (seconds) before retries times out. Default set
               to 1800 seconds. *Available from v1.29.0*
        :param pulumi.Input[str] vpc_id: The managed VPC identifier. *Available from v1.16.0*
        :param pulumi.Input[bool] wait_on_peering_status: Makes the resource wait until the peering is connected.
               Default set to false. *Available from v1.28.0*
        """
        if auto_create_routes is not None:
            pulumi.set(__self__, "auto_create_routes", auto_create_routes)
        if instance_id is not None:
            pulumi.set(__self__, "instance_id", instance_id)
        if peer_network_uri is not None:
            pulumi.set(__self__, "peer_network_uri", peer_network_uri)
        if sleep is not None:
            pulumi.set(__self__, "sleep", sleep)
        if state is not None:
            pulumi.set(__self__, "state", state)
        if state_details is not None:
            pulumi.set(__self__, "state_details", state_details)
        if timeout is not None:
            pulumi.set(__self__, "timeout", timeout)
        if vpc_id is not None:
            pulumi.set(__self__, "vpc_id", vpc_id)
        if wait_on_peering_status is not None:
            pulumi.set(__self__, "wait_on_peering_status", wait_on_peering_status)

    @property
    @pulumi.getter(name="autoCreateRoutes")
    def auto_create_routes(self) -> Optional[pulumi.Input[bool]]:
        """
        VPC peering auto created routes
        """
        return pulumi.get(self, "auto_create_routes")

    @auto_create_routes.setter
    def auto_create_routes(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "auto_create_routes", value)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> Optional[pulumi.Input[int]]:
        """
        The CloudAMQP instance identifier. *Deprecated from v1.16.0*
        """
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "instance_id", value)

    @property
    @pulumi.getter(name="peerNetworkUri")
    def peer_network_uri(self) -> Optional[pulumi.Input[str]]:
        """
        Network uri of the VPC network to which you will peer with.
        """
        return pulumi.get(self, "peer_network_uri")

    @peer_network_uri.setter
    def peer_network_uri(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "peer_network_uri", value)

    @property
    @pulumi.getter
    def sleep(self) -> Optional[pulumi.Input[int]]:
        """
        Configurable sleep time (seconds) between retries when requesting or reading
        peering. Default set to 10 seconds. *Available from v1.29.0*
        """
        return pulumi.get(self, "sleep")

    @sleep.setter
    def sleep(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "sleep", value)

    @property
    @pulumi.getter
    def state(self) -> Optional[pulumi.Input[str]]:
        """
        VPC peering state
        """
        return pulumi.get(self, "state")

    @state.setter
    def state(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "state", value)

    @property
    @pulumi.getter(name="stateDetails")
    def state_details(self) -> Optional[pulumi.Input[str]]:
        """
        VPC peering state details
        """
        return pulumi.get(self, "state_details")

    @state_details.setter
    def state_details(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "state_details", value)

    @property
    @pulumi.getter
    def timeout(self) -> Optional[pulumi.Input[int]]:
        """
        Configurable timeout time (seconds) before retries times out. Default set
        to 1800 seconds. *Available from v1.29.0*
        """
        return pulumi.get(self, "timeout")

    @timeout.setter
    def timeout(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "timeout", value)

    @property
    @pulumi.getter(name="vpcId")
    def vpc_id(self) -> Optional[pulumi.Input[str]]:
        """
        The managed VPC identifier. *Available from v1.16.0*
        """
        return pulumi.get(self, "vpc_id")

    @vpc_id.setter
    def vpc_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vpc_id", value)

    @property
    @pulumi.getter(name="waitOnPeeringStatus")
    def wait_on_peering_status(self) -> Optional[pulumi.Input[bool]]:
        """
        Makes the resource wait until the peering is connected.
        Default set to false. *Available from v1.28.0*
        """
        return pulumi.get(self, "wait_on_peering_status")

    @wait_on_peering_status.setter
    def wait_on_peering_status(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "wait_on_peering_status", value)


class VpcGcpPeering(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 instance_id: Optional[pulumi.Input[int]] = None,
                 peer_network_uri: Optional[pulumi.Input[str]] = None,
                 sleep: Optional[pulumi.Input[int]] = None,
                 timeout: Optional[pulumi.Input[int]] = None,
                 vpc_id: Optional[pulumi.Input[str]] = None,
                 wait_on_peering_status: Optional[pulumi.Input[bool]] = None,
                 __props__=None):
        """
        This resouce creates a VPC peering configuration for the CloudAMQP instance. The configuration will connect to another VPC network hosted on Google Cloud Platform (GCP). See the [GCP documentation](https://cloud.google.com/vpc/docs/using-vpc-peering) for more information on how to create the VPC peering configuration.

        > **Note:** Creating a VPC peering will automatically add firewall rules for the peered subnet.

        <details>
         <summary>
            <i>Default VPC peering firewall rule</i>
          </summary>

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        ```
        <!--End PulumiCodeChooser -->

        </details>

        Pricing is available at [cloudamqp.com](https://www.cloudamqp.com/plans.html).

        Only available for dedicated subscription plans.

        ## Example Usage

        <details>
          <summary>
            <b>
              <i>VPC peering pre v1.16.0</i>
            </b>
          </summary>

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumi_cloudamqp as cloudamqp

        # CloudAMQP instance
        instance = cloudamqp.Instance("instance",
            plan="bunny-1",
            region="google-compute-engine::europe-north1",
            tags=["terraform"],
            vpc_subnet="10.40.72.0/24")
        vpc_info = instance.id.apply(lambda id: cloudamqp.get_vpc_gcp_info_output(instance_id=id))
        # VPC peering configuration
        vpc_peering_request = cloudamqp.VpcGcpPeering("vpcPeeringRequest",
            instance_id=instance.id,
            peer_network_uri="https://www.googleapis.com/compute/v1/projects/<PROJECT-NAME>/global/networks/<NETWORK-NAME>")
        ```
        <!--End PulumiCodeChooser -->

        </details>

        <details>
          <summary>
            <b>
              <i>VPC peering post v1.16.0 (Managed VPC)</i>
            </b>
          </summary>

        ### With Additional Firewall Rules

        <details>
          <summary>
            <b>
              <i>VPC peering pre v1.16.0</i>
            </b>
          </summary>

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumi_cloudamqp as cloudamqp

        # VPC peering configuration
        vpc_peering_request = cloudamqp.VpcGcpPeering("vpcPeeringRequest",
            instance_id=cloudamqp_instance["instance"]["id"],
            peer_network_uri=var["peer_network_uri"])
        # Firewall rules
        firewall_settings = cloudamqp.SecurityFirewall("firewallSettings",
            instance_id=cloudamqp_instance["instance"]["id"],
            rules=[
                cloudamqp.SecurityFirewallRuleArgs(
                    ip=var["peer_subnet"],
                    ports=[15672],
                    services=[
                        "AMQP",
                        "AMQPS",
                        "STREAM",
                        "STREAM_SSL",
                    ],
                    description="VPC peering for <NETWORK>",
                ),
                cloudamqp.SecurityFirewallRuleArgs(
                    ip="192.168.0.0/24",
                    ports=[
                        4567,
                        4568,
                    ],
                    services=[
                        "AMQP",
                        "AMQPS",
                        "HTTPS",
                    ],
                ),
            ],
            opts=pulumi.ResourceOptions(depends_on=[vpc_peering_request]))
        ```
        <!--End PulumiCodeChooser -->

        </details>

        <details>
          <summary>
            <b>
              <i>VPC peering post v1.16.0 (Managed VPC)</i>
            </b>
          </summary>

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumi_cloudamqp as cloudamqp

        # VPC peering configuration
        vpc_peering_request = cloudamqp.VpcGcpPeering("vpcPeeringRequest",
            vpc_id=cloudamqp_vpc["vpc"]["id"],
            peer_network_uri=var["peer_network_uri"])
        # Firewall rules
        firewall_settings = cloudamqp.SecurityFirewall("firewallSettings",
            instance_id=cloudamqp_instance["instance"]["id"],
            rules=[
                cloudamqp.SecurityFirewallRuleArgs(
                    ip=var["peer_subnet"],
                    ports=[15672],
                    services=[
                        "AMQP",
                        "AMQPS",
                        "STREAM",
                        "STREAM_SSL",
                    ],
                    description="VPC peering for <NETWORK>",
                ),
                cloudamqp.SecurityFirewallRuleArgs(
                    ip="0.0.0.0/0",
                    ports=[],
                    services=["HTTPS"],
                    description="MGMT interface",
                ),
            ],
            opts=pulumi.ResourceOptions(depends_on=[vpc_peering_request]))
        ```
        <!--End PulumiCodeChooser -->

        </details>

        ## Depedency

        *Pre v1.16.0*
        This resource depends on CloudAMQP instance identifier, `cloudamqp_instance.instance.id`.

        *Post v1.16.0*
        This resource depends on CloudAMQP managed VPC identifier, `cloudamqp_vpc.vpc.id` or instance identifier, `cloudamqp_instance.instance.id`.

        ## Create VPC Peering with additional firewall rules

        To create a VPC peering configuration with additional firewall rules, it's required to chain the SecurityFirewall
        resource to avoid parallel conflicting resource calls. This is done by adding dependency from the firewall resource to the VPC peering resource.

        Furthermore, since all firewall rules are overwritten, the otherwise automatically added rules for the VPC peering also needs to be added.

        See example below.

        ## Import

        Not possible to import this resource.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] instance_id: The CloudAMQP instance identifier. *Deprecated from v1.16.0*
        :param pulumi.Input[str] peer_network_uri: Network uri of the VPC network to which you will peer with.
        :param pulumi.Input[int] sleep: Configurable sleep time (seconds) between retries when requesting or reading
               peering. Default set to 10 seconds. *Available from v1.29.0*
        :param pulumi.Input[int] timeout: Configurable timeout time (seconds) before retries times out. Default set
               to 1800 seconds. *Available from v1.29.0*
        :param pulumi.Input[str] vpc_id: The managed VPC identifier. *Available from v1.16.0*
        :param pulumi.Input[bool] wait_on_peering_status: Makes the resource wait until the peering is connected.
               Default set to false. *Available from v1.28.0*
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: VpcGcpPeeringArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        This resouce creates a VPC peering configuration for the CloudAMQP instance. The configuration will connect to another VPC network hosted on Google Cloud Platform (GCP). See the [GCP documentation](https://cloud.google.com/vpc/docs/using-vpc-peering) for more information on how to create the VPC peering configuration.

        > **Note:** Creating a VPC peering will automatically add firewall rules for the peered subnet.

        <details>
         <summary>
            <i>Default VPC peering firewall rule</i>
          </summary>

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        ```
        <!--End PulumiCodeChooser -->

        </details>

        Pricing is available at [cloudamqp.com](https://www.cloudamqp.com/plans.html).

        Only available for dedicated subscription plans.

        ## Example Usage

        <details>
          <summary>
            <b>
              <i>VPC peering pre v1.16.0</i>
            </b>
          </summary>

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumi_cloudamqp as cloudamqp

        # CloudAMQP instance
        instance = cloudamqp.Instance("instance",
            plan="bunny-1",
            region="google-compute-engine::europe-north1",
            tags=["terraform"],
            vpc_subnet="10.40.72.0/24")
        vpc_info = instance.id.apply(lambda id: cloudamqp.get_vpc_gcp_info_output(instance_id=id))
        # VPC peering configuration
        vpc_peering_request = cloudamqp.VpcGcpPeering("vpcPeeringRequest",
            instance_id=instance.id,
            peer_network_uri="https://www.googleapis.com/compute/v1/projects/<PROJECT-NAME>/global/networks/<NETWORK-NAME>")
        ```
        <!--End PulumiCodeChooser -->

        </details>

        <details>
          <summary>
            <b>
              <i>VPC peering post v1.16.0 (Managed VPC)</i>
            </b>
          </summary>

        ### With Additional Firewall Rules

        <details>
          <summary>
            <b>
              <i>VPC peering pre v1.16.0</i>
            </b>
          </summary>

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumi_cloudamqp as cloudamqp

        # VPC peering configuration
        vpc_peering_request = cloudamqp.VpcGcpPeering("vpcPeeringRequest",
            instance_id=cloudamqp_instance["instance"]["id"],
            peer_network_uri=var["peer_network_uri"])
        # Firewall rules
        firewall_settings = cloudamqp.SecurityFirewall("firewallSettings",
            instance_id=cloudamqp_instance["instance"]["id"],
            rules=[
                cloudamqp.SecurityFirewallRuleArgs(
                    ip=var["peer_subnet"],
                    ports=[15672],
                    services=[
                        "AMQP",
                        "AMQPS",
                        "STREAM",
                        "STREAM_SSL",
                    ],
                    description="VPC peering for <NETWORK>",
                ),
                cloudamqp.SecurityFirewallRuleArgs(
                    ip="192.168.0.0/24",
                    ports=[
                        4567,
                        4568,
                    ],
                    services=[
                        "AMQP",
                        "AMQPS",
                        "HTTPS",
                    ],
                ),
            ],
            opts=pulumi.ResourceOptions(depends_on=[vpc_peering_request]))
        ```
        <!--End PulumiCodeChooser -->

        </details>

        <details>
          <summary>
            <b>
              <i>VPC peering post v1.16.0 (Managed VPC)</i>
            </b>
          </summary>

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumi_cloudamqp as cloudamqp

        # VPC peering configuration
        vpc_peering_request = cloudamqp.VpcGcpPeering("vpcPeeringRequest",
            vpc_id=cloudamqp_vpc["vpc"]["id"],
            peer_network_uri=var["peer_network_uri"])
        # Firewall rules
        firewall_settings = cloudamqp.SecurityFirewall("firewallSettings",
            instance_id=cloudamqp_instance["instance"]["id"],
            rules=[
                cloudamqp.SecurityFirewallRuleArgs(
                    ip=var["peer_subnet"],
                    ports=[15672],
                    services=[
                        "AMQP",
                        "AMQPS",
                        "STREAM",
                        "STREAM_SSL",
                    ],
                    description="VPC peering for <NETWORK>",
                ),
                cloudamqp.SecurityFirewallRuleArgs(
                    ip="0.0.0.0/0",
                    ports=[],
                    services=["HTTPS"],
                    description="MGMT interface",
                ),
            ],
            opts=pulumi.ResourceOptions(depends_on=[vpc_peering_request]))
        ```
        <!--End PulumiCodeChooser -->

        </details>

        ## Depedency

        *Pre v1.16.0*
        This resource depends on CloudAMQP instance identifier, `cloudamqp_instance.instance.id`.

        *Post v1.16.0*
        This resource depends on CloudAMQP managed VPC identifier, `cloudamqp_vpc.vpc.id` or instance identifier, `cloudamqp_instance.instance.id`.

        ## Create VPC Peering with additional firewall rules

        To create a VPC peering configuration with additional firewall rules, it's required to chain the SecurityFirewall
        resource to avoid parallel conflicting resource calls. This is done by adding dependency from the firewall resource to the VPC peering resource.

        Furthermore, since all firewall rules are overwritten, the otherwise automatically added rules for the VPC peering also needs to be added.

        See example below.

        ## Import

        Not possible to import this resource.

        :param str resource_name: The name of the resource.
        :param VpcGcpPeeringArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(VpcGcpPeeringArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 instance_id: Optional[pulumi.Input[int]] = None,
                 peer_network_uri: Optional[pulumi.Input[str]] = None,
                 sleep: Optional[pulumi.Input[int]] = None,
                 timeout: Optional[pulumi.Input[int]] = None,
                 vpc_id: Optional[pulumi.Input[str]] = None,
                 wait_on_peering_status: Optional[pulumi.Input[bool]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = VpcGcpPeeringArgs.__new__(VpcGcpPeeringArgs)

            __props__.__dict__["instance_id"] = instance_id
            if peer_network_uri is None and not opts.urn:
                raise TypeError("Missing required property 'peer_network_uri'")
            __props__.__dict__["peer_network_uri"] = peer_network_uri
            __props__.__dict__["sleep"] = sleep
            __props__.__dict__["timeout"] = timeout
            __props__.__dict__["vpc_id"] = vpc_id
            __props__.__dict__["wait_on_peering_status"] = wait_on_peering_status
            __props__.__dict__["auto_create_routes"] = None
            __props__.__dict__["state"] = None
            __props__.__dict__["state_details"] = None
        super(VpcGcpPeering, __self__).__init__(
            'cloudamqp:index/vpcGcpPeering:VpcGcpPeering',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            auto_create_routes: Optional[pulumi.Input[bool]] = None,
            instance_id: Optional[pulumi.Input[int]] = None,
            peer_network_uri: Optional[pulumi.Input[str]] = None,
            sleep: Optional[pulumi.Input[int]] = None,
            state: Optional[pulumi.Input[str]] = None,
            state_details: Optional[pulumi.Input[str]] = None,
            timeout: Optional[pulumi.Input[int]] = None,
            vpc_id: Optional[pulumi.Input[str]] = None,
            wait_on_peering_status: Optional[pulumi.Input[bool]] = None) -> 'VpcGcpPeering':
        """
        Get an existing VpcGcpPeering resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] auto_create_routes: VPC peering auto created routes
        :param pulumi.Input[int] instance_id: The CloudAMQP instance identifier. *Deprecated from v1.16.0*
        :param pulumi.Input[str] peer_network_uri: Network uri of the VPC network to which you will peer with.
        :param pulumi.Input[int] sleep: Configurable sleep time (seconds) between retries when requesting or reading
               peering. Default set to 10 seconds. *Available from v1.29.0*
        :param pulumi.Input[str] state: VPC peering state
        :param pulumi.Input[str] state_details: VPC peering state details
        :param pulumi.Input[int] timeout: Configurable timeout time (seconds) before retries times out. Default set
               to 1800 seconds. *Available from v1.29.0*
        :param pulumi.Input[str] vpc_id: The managed VPC identifier. *Available from v1.16.0*
        :param pulumi.Input[bool] wait_on_peering_status: Makes the resource wait until the peering is connected.
               Default set to false. *Available from v1.28.0*
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _VpcGcpPeeringState.__new__(_VpcGcpPeeringState)

        __props__.__dict__["auto_create_routes"] = auto_create_routes
        __props__.__dict__["instance_id"] = instance_id
        __props__.__dict__["peer_network_uri"] = peer_network_uri
        __props__.__dict__["sleep"] = sleep
        __props__.__dict__["state"] = state
        __props__.__dict__["state_details"] = state_details
        __props__.__dict__["timeout"] = timeout
        __props__.__dict__["vpc_id"] = vpc_id
        __props__.__dict__["wait_on_peering_status"] = wait_on_peering_status
        return VpcGcpPeering(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="autoCreateRoutes")
    def auto_create_routes(self) -> pulumi.Output[bool]:
        """
        VPC peering auto created routes
        """
        return pulumi.get(self, "auto_create_routes")

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> pulumi.Output[Optional[int]]:
        """
        The CloudAMQP instance identifier. *Deprecated from v1.16.0*
        """
        return pulumi.get(self, "instance_id")

    @property
    @pulumi.getter(name="peerNetworkUri")
    def peer_network_uri(self) -> pulumi.Output[str]:
        """
        Network uri of the VPC network to which you will peer with.
        """
        return pulumi.get(self, "peer_network_uri")

    @property
    @pulumi.getter
    def sleep(self) -> pulumi.Output[Optional[int]]:
        """
        Configurable sleep time (seconds) between retries when requesting or reading
        peering. Default set to 10 seconds. *Available from v1.29.0*
        """
        return pulumi.get(self, "sleep")

    @property
    @pulumi.getter
    def state(self) -> pulumi.Output[str]:
        """
        VPC peering state
        """
        return pulumi.get(self, "state")

    @property
    @pulumi.getter(name="stateDetails")
    def state_details(self) -> pulumi.Output[str]:
        """
        VPC peering state details
        """
        return pulumi.get(self, "state_details")

    @property
    @pulumi.getter
    def timeout(self) -> pulumi.Output[Optional[int]]:
        """
        Configurable timeout time (seconds) before retries times out. Default set
        to 1800 seconds. *Available from v1.29.0*
        """
        return pulumi.get(self, "timeout")

    @property
    @pulumi.getter(name="vpcId")
    def vpc_id(self) -> pulumi.Output[Optional[str]]:
        """
        The managed VPC identifier. *Available from v1.16.0*
        """
        return pulumi.get(self, "vpc_id")

    @property
    @pulumi.getter(name="waitOnPeeringStatus")
    def wait_on_peering_status(self) -> pulumi.Output[Optional[bool]]:
        """
        Makes the resource wait until the peering is connected.
        Default set to false. *Available from v1.28.0*
        """
        return pulumi.get(self, "wait_on_peering_status")

