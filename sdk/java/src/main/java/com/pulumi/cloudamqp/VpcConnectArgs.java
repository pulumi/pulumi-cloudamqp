// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.cloudamqp;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class VpcConnectArgs extends com.pulumi.resources.ResourceArgs {

    public static final VpcConnectArgs Empty = new VpcConnectArgs();

    /**
     * List of allowed prinicpals used by AWS, see below table.
     * 
     */
    @Import(name="allowedPrincipals")
    private @Nullable Output<List<String>> allowedPrincipals;

    /**
     * @return List of allowed prinicpals used by AWS, see below table.
     * 
     */
    public Optional<Output<List<String>>> allowedPrincipals() {
        return Optional.ofNullable(this.allowedPrincipals);
    }

    /**
     * List of allowed projects used by GCP, see below table.
     * 
     */
    @Import(name="allowedProjects")
    private @Nullable Output<List<String>> allowedProjects;

    /**
     * @return List of allowed projects used by GCP, see below table.
     * 
     */
    public Optional<Output<List<String>>> allowedProjects() {
        return Optional.ofNullable(this.allowedProjects);
    }

    /**
     * List of approved subscriptions used by Azure, see below
     * table.
     * 
     */
    @Import(name="approvedSubscriptions")
    private @Nullable Output<List<String>> approvedSubscriptions;

    /**
     * @return List of approved subscriptions used by Azure, see below
     * table.
     * 
     */
    public Optional<Output<List<String>>> approvedSubscriptions() {
        return Optional.ofNullable(this.approvedSubscriptions);
    }

    /**
     * The CloudAMQP instance identifier.
     * 
     */
    @Import(name="instanceId", required=true)
    private Output<Integer> instanceId;

    /**
     * @return The CloudAMQP instance identifier.
     * 
     */
    public Output<Integer> instanceId() {
        return this.instanceId;
    }

    /**
     * The region where the CloudAMQP instance is hosted.
     * 
     */
    @Import(name="region", required=true)
    private Output<String> region;

    /**
     * @return The region where the CloudAMQP instance is hosted.
     * 
     */
    public Output<String> region() {
        return this.region;
    }

    /**
     * Configurable sleep time (seconds) when enable Private
     * Service Connect. Default set to 10 seconds.
     * 
     */
    @Import(name="sleep")
    private @Nullable Output<Integer> sleep;

    /**
     * @return Configurable sleep time (seconds) when enable Private
     * Service Connect. Default set to 10 seconds.
     * 
     */
    public Optional<Output<Integer>> sleep() {
        return Optional.ofNullable(this.sleep);
    }

    /**
     * Configurable timeout time (seconds) when enable Private
     * Service Connect. Default set to 1800 seconds.
     * 
     * ***
     * 
     * The `allowed_principals`, `approved_subscriptions` or `allowed_projects` data depends on the
     * provider platform:
     * 
     * | Platform | Description | Format |
     * |---|---|---|
     * | AWS | IAM ARN principals | arn:aws:iam::aws-account-id:root&lt;br&gt;arn:aws:iam::aws-account-id:user/user-name&lt;br&gt; arn:aws:iam::aws-account-id:role/role-name |
     * | Azure | Subscription (GUID) | XXXXXXXX-XXXX-XXXX-XXXX-XXXXXXXXXXXX |
     * | GCP | Project IDs [Google docs] | 6 to 30 lowercase letters, digits, or hyphens |
     * 
     */
    @Import(name="timeout")
    private @Nullable Output<Integer> timeout;

    /**
     * @return Configurable timeout time (seconds) when enable Private
     * Service Connect. Default set to 1800 seconds.
     * 
     * ***
     * 
     * The `allowed_principals`, `approved_subscriptions` or `allowed_projects` data depends on the
     * provider platform:
     * 
     * | Platform | Description | Format |
     * |---|---|---|
     * | AWS | IAM ARN principals | arn:aws:iam::aws-account-id:root&lt;br&gt;arn:aws:iam::aws-account-id:user/user-name&lt;br&gt; arn:aws:iam::aws-account-id:role/role-name |
     * | Azure | Subscription (GUID) | XXXXXXXX-XXXX-XXXX-XXXX-XXXXXXXXXXXX |
     * | GCP | Project IDs [Google docs] | 6 to 30 lowercase letters, digits, or hyphens |
     * 
     */
    public Optional<Output<Integer>> timeout() {
        return Optional.ofNullable(this.timeout);
    }

    private VpcConnectArgs() {}

    private VpcConnectArgs(VpcConnectArgs $) {
        this.allowedPrincipals = $.allowedPrincipals;
        this.allowedProjects = $.allowedProjects;
        this.approvedSubscriptions = $.approvedSubscriptions;
        this.instanceId = $.instanceId;
        this.region = $.region;
        this.sleep = $.sleep;
        this.timeout = $.timeout;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(VpcConnectArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private VpcConnectArgs $;

        public Builder() {
            $ = new VpcConnectArgs();
        }

        public Builder(VpcConnectArgs defaults) {
            $ = new VpcConnectArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param allowedPrincipals List of allowed prinicpals used by AWS, see below table.
         * 
         * @return builder
         * 
         */
        public Builder allowedPrincipals(@Nullable Output<List<String>> allowedPrincipals) {
            $.allowedPrincipals = allowedPrincipals;
            return this;
        }

        /**
         * @param allowedPrincipals List of allowed prinicpals used by AWS, see below table.
         * 
         * @return builder
         * 
         */
        public Builder allowedPrincipals(List<String> allowedPrincipals) {
            return allowedPrincipals(Output.of(allowedPrincipals));
        }

        /**
         * @param allowedPrincipals List of allowed prinicpals used by AWS, see below table.
         * 
         * @return builder
         * 
         */
        public Builder allowedPrincipals(String... allowedPrincipals) {
            return allowedPrincipals(List.of(allowedPrincipals));
        }

        /**
         * @param allowedProjects List of allowed projects used by GCP, see below table.
         * 
         * @return builder
         * 
         */
        public Builder allowedProjects(@Nullable Output<List<String>> allowedProjects) {
            $.allowedProjects = allowedProjects;
            return this;
        }

        /**
         * @param allowedProjects List of allowed projects used by GCP, see below table.
         * 
         * @return builder
         * 
         */
        public Builder allowedProjects(List<String> allowedProjects) {
            return allowedProjects(Output.of(allowedProjects));
        }

        /**
         * @param allowedProjects List of allowed projects used by GCP, see below table.
         * 
         * @return builder
         * 
         */
        public Builder allowedProjects(String... allowedProjects) {
            return allowedProjects(List.of(allowedProjects));
        }

        /**
         * @param approvedSubscriptions List of approved subscriptions used by Azure, see below
         * table.
         * 
         * @return builder
         * 
         */
        public Builder approvedSubscriptions(@Nullable Output<List<String>> approvedSubscriptions) {
            $.approvedSubscriptions = approvedSubscriptions;
            return this;
        }

        /**
         * @param approvedSubscriptions List of approved subscriptions used by Azure, see below
         * table.
         * 
         * @return builder
         * 
         */
        public Builder approvedSubscriptions(List<String> approvedSubscriptions) {
            return approvedSubscriptions(Output.of(approvedSubscriptions));
        }

        /**
         * @param approvedSubscriptions List of approved subscriptions used by Azure, see below
         * table.
         * 
         * @return builder
         * 
         */
        public Builder approvedSubscriptions(String... approvedSubscriptions) {
            return approvedSubscriptions(List.of(approvedSubscriptions));
        }

        /**
         * @param instanceId The CloudAMQP instance identifier.
         * 
         * @return builder
         * 
         */
        public Builder instanceId(Output<Integer> instanceId) {
            $.instanceId = instanceId;
            return this;
        }

        /**
         * @param instanceId The CloudAMQP instance identifier.
         * 
         * @return builder
         * 
         */
        public Builder instanceId(Integer instanceId) {
            return instanceId(Output.of(instanceId));
        }

        /**
         * @param region The region where the CloudAMQP instance is hosted.
         * 
         * @return builder
         * 
         */
        public Builder region(Output<String> region) {
            $.region = region;
            return this;
        }

        /**
         * @param region The region where the CloudAMQP instance is hosted.
         * 
         * @return builder
         * 
         */
        public Builder region(String region) {
            return region(Output.of(region));
        }

        /**
         * @param sleep Configurable sleep time (seconds) when enable Private
         * Service Connect. Default set to 10 seconds.
         * 
         * @return builder
         * 
         */
        public Builder sleep(@Nullable Output<Integer> sleep) {
            $.sleep = sleep;
            return this;
        }

        /**
         * @param sleep Configurable sleep time (seconds) when enable Private
         * Service Connect. Default set to 10 seconds.
         * 
         * @return builder
         * 
         */
        public Builder sleep(Integer sleep) {
            return sleep(Output.of(sleep));
        }

        /**
         * @param timeout Configurable timeout time (seconds) when enable Private
         * Service Connect. Default set to 1800 seconds.
         * 
         * ***
         * 
         * The `allowed_principals`, `approved_subscriptions` or `allowed_projects` data depends on the
         * provider platform:
         * 
         * | Platform | Description | Format |
         * |---|---|---|
         * | AWS | IAM ARN principals | arn:aws:iam::aws-account-id:root&lt;br&gt;arn:aws:iam::aws-account-id:user/user-name&lt;br&gt; arn:aws:iam::aws-account-id:role/role-name |
         * | Azure | Subscription (GUID) | XXXXXXXX-XXXX-XXXX-XXXX-XXXXXXXXXXXX |
         * | GCP | Project IDs [Google docs] | 6 to 30 lowercase letters, digits, or hyphens |
         * 
         * @return builder
         * 
         */
        public Builder timeout(@Nullable Output<Integer> timeout) {
            $.timeout = timeout;
            return this;
        }

        /**
         * @param timeout Configurable timeout time (seconds) when enable Private
         * Service Connect. Default set to 1800 seconds.
         * 
         * ***
         * 
         * The `allowed_principals`, `approved_subscriptions` or `allowed_projects` data depends on the
         * provider platform:
         * 
         * | Platform | Description | Format |
         * |---|---|---|
         * | AWS | IAM ARN principals | arn:aws:iam::aws-account-id:root&lt;br&gt;arn:aws:iam::aws-account-id:user/user-name&lt;br&gt; arn:aws:iam::aws-account-id:role/role-name |
         * | Azure | Subscription (GUID) | XXXXXXXX-XXXX-XXXX-XXXX-XXXXXXXXXXXX |
         * | GCP | Project IDs [Google docs] | 6 to 30 lowercase letters, digits, or hyphens |
         * 
         * @return builder
         * 
         */
        public Builder timeout(Integer timeout) {
            return timeout(Output.of(timeout));
        }

        public VpcConnectArgs build() {
            if ($.instanceId == null) {
                throw new MissingRequiredPropertyException("VpcConnectArgs", "instanceId");
            }
            if ($.region == null) {
                throw new MissingRequiredPropertyException("VpcConnectArgs", "region");
            }
            return $;
        }
    }

}
