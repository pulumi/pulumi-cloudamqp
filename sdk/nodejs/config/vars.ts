// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

let __config = new pulumi.Config("cloudamqp");

/**
 * Key used to authentication to the CloudAMQP Customer API
 */
export let apikey: string | undefined = __config.get("apikey") || utilities.getEnv("CLOUDAMQP_APIKEY");
/**
 * Base URL to CloudAMQP Customer website
 */
export let baseurl: string | undefined = __config.get("baseurl");
