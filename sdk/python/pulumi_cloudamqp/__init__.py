# coding=utf-8
# *** WARNING: this file was generated by pulumi-language-python. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import builtins
from . import _utilities
import typing
# Export this package's modules as members:
from .account_actions import *
from .alarm import *
from .custom_domain import *
from .extra_disk_size import *
from .get_account import *
from .get_account_vpcs import *
from .get_alarm import *
from .get_alarms import *
from .get_credentials import *
from .get_instance import *
from .get_nodes import *
from .get_notification import *
from .get_notifications import *
from .get_plugins import *
from .get_plugins_community import *
from .get_upgradable_versions import *
from .get_vpc_gcp_info import *
from .get_vpc_info import *
from .instance import *
from .integration_aws_eventbridge import *
from .integration_log import *
from .integration_metric import *
from .maintenance_window import *
from .node_actions import *
from .notification import *
from .plugin import *
from .plugin_community import *
from .privatelink_aws import *
from .privatelink_azure import *
from .provider import *
from .rabbit_configuration import *
from .security_firewall import *
from .upgrade_lavinmq import *
from .upgrade_rabbitmq import *
from .vpc import *
from .vpc_connect import *
from .vpc_gcp_peering import *
from .vpc_peering import *
from .webhook import *
from ._inputs import *
from . import outputs

# Make subpackages available:
if typing.TYPE_CHECKING:
    import pulumi_cloudamqp.config as __config
    config = __config
else:
    config = _utilities.lazy_import('pulumi_cloudamqp.config')

_utilities.register(
    resource_modules="""
[
 {
  "pkg": "cloudamqp",
  "mod": "index/accountActions",
  "fqn": "pulumi_cloudamqp",
  "classes": {
   "cloudamqp:index/accountActions:AccountActions": "AccountActions"
  }
 },
 {
  "pkg": "cloudamqp",
  "mod": "index/alarm",
  "fqn": "pulumi_cloudamqp",
  "classes": {
   "cloudamqp:index/alarm:Alarm": "Alarm"
  }
 },
 {
  "pkg": "cloudamqp",
  "mod": "index/customDomain",
  "fqn": "pulumi_cloudamqp",
  "classes": {
   "cloudamqp:index/customDomain:CustomDomain": "CustomDomain"
  }
 },
 {
  "pkg": "cloudamqp",
  "mod": "index/extraDiskSize",
  "fqn": "pulumi_cloudamqp",
  "classes": {
   "cloudamqp:index/extraDiskSize:ExtraDiskSize": "ExtraDiskSize"
  }
 },
 {
  "pkg": "cloudamqp",
  "mod": "index/instance",
  "fqn": "pulumi_cloudamqp",
  "classes": {
   "cloudamqp:index/instance:Instance": "Instance"
  }
 },
 {
  "pkg": "cloudamqp",
  "mod": "index/integrationAwsEventbridge",
  "fqn": "pulumi_cloudamqp",
  "classes": {
   "cloudamqp:index/integrationAwsEventbridge:IntegrationAwsEventbridge": "IntegrationAwsEventbridge"
  }
 },
 {
  "pkg": "cloudamqp",
  "mod": "index/integrationLog",
  "fqn": "pulumi_cloudamqp",
  "classes": {
   "cloudamqp:index/integrationLog:IntegrationLog": "IntegrationLog"
  }
 },
 {
  "pkg": "cloudamqp",
  "mod": "index/integrationMetric",
  "fqn": "pulumi_cloudamqp",
  "classes": {
   "cloudamqp:index/integrationMetric:IntegrationMetric": "IntegrationMetric"
  }
 },
 {
  "pkg": "cloudamqp",
  "mod": "index/maintenanceWindow",
  "fqn": "pulumi_cloudamqp",
  "classes": {
   "cloudamqp:index/maintenanceWindow:MaintenanceWindow": "MaintenanceWindow"
  }
 },
 {
  "pkg": "cloudamqp",
  "mod": "index/nodeActions",
  "fqn": "pulumi_cloudamqp",
  "classes": {
   "cloudamqp:index/nodeActions:NodeActions": "NodeActions"
  }
 },
 {
  "pkg": "cloudamqp",
  "mod": "index/notification",
  "fqn": "pulumi_cloudamqp",
  "classes": {
   "cloudamqp:index/notification:Notification": "Notification"
  }
 },
 {
  "pkg": "cloudamqp",
  "mod": "index/plugin",
  "fqn": "pulumi_cloudamqp",
  "classes": {
   "cloudamqp:index/plugin:Plugin": "Plugin"
  }
 },
 {
  "pkg": "cloudamqp",
  "mod": "index/pluginCommunity",
  "fqn": "pulumi_cloudamqp",
  "classes": {
   "cloudamqp:index/pluginCommunity:PluginCommunity": "PluginCommunity"
  }
 },
 {
  "pkg": "cloudamqp",
  "mod": "index/privatelinkAws",
  "fqn": "pulumi_cloudamqp",
  "classes": {
   "cloudamqp:index/privatelinkAws:PrivatelinkAws": "PrivatelinkAws"
  }
 },
 {
  "pkg": "cloudamqp",
  "mod": "index/privatelinkAzure",
  "fqn": "pulumi_cloudamqp",
  "classes": {
   "cloudamqp:index/privatelinkAzure:PrivatelinkAzure": "PrivatelinkAzure"
  }
 },
 {
  "pkg": "cloudamqp",
  "mod": "index/rabbitConfiguration",
  "fqn": "pulumi_cloudamqp",
  "classes": {
   "cloudamqp:index/rabbitConfiguration:RabbitConfiguration": "RabbitConfiguration"
  }
 },
 {
  "pkg": "cloudamqp",
  "mod": "index/securityFirewall",
  "fqn": "pulumi_cloudamqp",
  "classes": {
   "cloudamqp:index/securityFirewall:SecurityFirewall": "SecurityFirewall"
  }
 },
 {
  "pkg": "cloudamqp",
  "mod": "index/upgradeLavinmq",
  "fqn": "pulumi_cloudamqp",
  "classes": {
   "cloudamqp:index/upgradeLavinmq:UpgradeLavinmq": "UpgradeLavinmq"
  }
 },
 {
  "pkg": "cloudamqp",
  "mod": "index/upgradeRabbitmq",
  "fqn": "pulumi_cloudamqp",
  "classes": {
   "cloudamqp:index/upgradeRabbitmq:UpgradeRabbitmq": "UpgradeRabbitmq"
  }
 },
 {
  "pkg": "cloudamqp",
  "mod": "index/vpc",
  "fqn": "pulumi_cloudamqp",
  "classes": {
   "cloudamqp:index/vpc:Vpc": "Vpc"
  }
 },
 {
  "pkg": "cloudamqp",
  "mod": "index/vpcConnect",
  "fqn": "pulumi_cloudamqp",
  "classes": {
   "cloudamqp:index/vpcConnect:VpcConnect": "VpcConnect"
  }
 },
 {
  "pkg": "cloudamqp",
  "mod": "index/vpcGcpPeering",
  "fqn": "pulumi_cloudamqp",
  "classes": {
   "cloudamqp:index/vpcGcpPeering:VpcGcpPeering": "VpcGcpPeering"
  }
 },
 {
  "pkg": "cloudamqp",
  "mod": "index/vpcPeering",
  "fqn": "pulumi_cloudamqp",
  "classes": {
   "cloudamqp:index/vpcPeering:VpcPeering": "VpcPeering"
  }
 },
 {
  "pkg": "cloudamqp",
  "mod": "index/webhook",
  "fqn": "pulumi_cloudamqp",
  "classes": {
   "cloudamqp:index/webhook:Webhook": "Webhook"
  }
 }
]
""",
    resource_packages="""
[
 {
  "pkg": "cloudamqp",
  "token": "pulumi:providers:cloudamqp",
  "fqn": "pulumi_cloudamqp",
  "class": "Provider"
 }
]
"""
)
