# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

from . import _utilities
import typing
# Export this package's modules as members:
from .alarm import *
from .custom_domain import *
from .get_account import *
from .get_alarm import *
from .get_credentials import *
from .get_instance import *
from .get_nodes import *
from .get_notification import *
from .get_plugins import *
from .get_plugins_community import *
from .get_vpc_info import *
from .instance import *
from .integration_log import *
from .integration_metric import *
from .notification import *
from .plugin import *
from .plugin_community import *
from .provider import *
from .security_firewall import *
from .vpc_peering import *
from .webhook import *
from ._inputs import *
from . import outputs

# Make subpackages available:
if typing.TYPE_CHECKING:
    import pulumi_cloudamqp.config as config
else:
    config = _utilities.lazy_import('pulumi_cloudamqp.config')

_utilities.register(
    resource_modules="""
[
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
  "mod": "index/instance",
  "fqn": "pulumi_cloudamqp",
  "classes": {
   "cloudamqp:index/instance:Instance": "Instance"
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
  "mod": "index/securityFirewall",
  "fqn": "pulumi_cloudamqp",
  "classes": {
   "cloudamqp:index/securityFirewall:SecurityFirewall": "SecurityFirewall"
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
