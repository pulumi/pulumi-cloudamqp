// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.cloudamqp;

import com.pulumi.cloudamqp.AccountActionArgs;
import com.pulumi.cloudamqp.Utilities;
import com.pulumi.cloudamqp.inputs.AccountActionState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Integer;
import java.lang.String;
import javax.annotation.Nullable;

/**
 * ## Import
 * 
 * Not possible to import this resource.
 * 
 */
@ResourceType(type="cloudamqp:index/accountAction:AccountAction")
public class AccountAction extends com.pulumi.resources.CustomResource {
    /**
     * The action to be invoked. Allowed actions `rotate-password`, `rotate-apikey`.
     * 
     */
    @Export(name="action", refs={String.class}, tree="[0]")
    private Output<String> action;

    /**
     * @return The action to be invoked. Allowed actions `rotate-password`, `rotate-apikey`.
     * 
     */
    public Output<String> action() {
        return this.action;
    }
    /**
     * The CloudAMQP instance ID.
     * 
     */
    @Export(name="instanceId", refs={Integer.class}, tree="[0]")
    private Output<Integer> instanceId;

    /**
     * @return The CloudAMQP instance ID.
     * 
     */
    public Output<Integer> instanceId() {
        return this.instanceId;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public AccountAction(String name) {
        this(name, AccountActionArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public AccountAction(String name, AccountActionArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public AccountAction(String name, AccountActionArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("cloudamqp:index/accountAction:AccountAction", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()));
    }

    private AccountAction(String name, Output<String> id, @Nullable AccountActionState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("cloudamqp:index/accountAction:AccountAction", name, state, makeResourceOptions(options, id));
    }

    private static AccountActionArgs makeArgs(AccountActionArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? AccountActionArgs.Empty : args;
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .build();
        return com.pulumi.resources.CustomResourceOptions.merge(defaultOptions, options, id);
    }

    /**
     * Get an existing Host resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state
     * @param options Optional settings to control the behavior of the CustomResource.
     */
    public static AccountAction get(String name, Output<String> id, @Nullable AccountActionState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new AccountAction(name, id, state, options);
    }
}
