// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.cloudamqp;

import com.pulumi.core.internal.Codegen;
import java.lang.Boolean;
import java.lang.String;
import java.util.Optional;

public final class Config {

    private static final com.pulumi.Config config = com.pulumi.Config.of("cloudamqp");
/**
 * Key used to authentication to the CloudAMQP Customer API
 * 
 */
    public String apikey() {
        return Codegen.stringProp("apikey").config(config).require();
    }
/**
 * Base URL to CloudAMQP Customer website
 * 
 */
    public Optional<String> baseurl() {
        return Codegen.stringProp("baseurl").config(config).get();
    }
    public Optional<Boolean> enableFasterInstanceDestroy() {
        return Codegen.booleanProp("enableFasterInstanceDestroy").config(config).get();
    }
}
