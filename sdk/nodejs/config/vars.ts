// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

declare var exports: any;
const __config = new pulumi.Config("cloudamqp");

/**
 * Key used to authentication to the CloudAMQP Customer API
 */
export declare const apikey: string | undefined;
Object.defineProperty(exports, "apikey", {
    get() {
        return __config.get("apikey");
    },
    enumerable: true,
});

/**
 * Base URL to CloudAMQP Customer website
 */
export declare const baseurl: string | undefined;
Object.defineProperty(exports, "baseurl", {
    get() {
        return __config.get("baseurl");
    },
    enumerable: true,
});

export declare const enableFasterInstanceDestroy: boolean | undefined;
Object.defineProperty(exports, "enableFasterInstanceDestroy", {
    get() {
        return __config.getObject<boolean>("enableFasterInstanceDestroy");
    },
    enumerable: true,
});

