# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from .. import _utilities, _tables

__all__ = [
    'apikey',
    'baseurl',
]

__config__ = pulumi.Config('cloudamqp')

apikey = __config__.get('apikey') or _utilities.get_env('CLOUDAMQP_APIKEY')
"""
Key used to authentication to the CloudAMQP Customer API
"""

baseurl = __config__.get('baseurl')
"""
Base URL to CloudAMQP Customer website
"""

