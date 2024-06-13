// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.cloudamqp;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class IntegrationLogArgs extends com.pulumi.resources.ResourceArgs {

    public static final IntegrationLogArgs Empty = new IntegrationLogArgs();

    /**
     * AWS access key identifier.
     * 
     */
    @Import(name="accessKeyId")
    private @Nullable Output<String> accessKeyId;

    /**
     * @return AWS access key identifier.
     * 
     */
    public Optional<Output<String>> accessKeyId() {
        return Optional.ofNullable(this.accessKeyId);
    }

    /**
     * The API key.
     * 
     */
    @Import(name="apiKey")
    private @Nullable Output<String> apiKey;

    /**
     * @return The API key.
     * 
     */
    public Optional<Output<String>> apiKey() {
        return Optional.ofNullable(this.apiKey);
    }

    /**
     * The application name for Coralogix.
     * 
     */
    @Import(name="application")
    private @Nullable Output<String> application;

    /**
     * @return The application name for Coralogix.
     * 
     */
    public Optional<Output<String>> application() {
        return Optional.ofNullable(this.application);
    }

    /**
     * The application identifier for Azure monitor.
     * 
     */
    @Import(name="applicationId")
    private @Nullable Output<String> applicationId;

    /**
     * @return The application identifier for Azure monitor.
     * 
     */
    public Optional<Output<String>> applicationId() {
        return Optional.ofNullable(this.applicationId);
    }

    /**
     * The application secret for Azure monitor.
     * 
     */
    @Import(name="applicationSecret")
    private @Nullable Output<String> applicationSecret;

    /**
     * @return The application secret for Azure monitor.
     * 
     */
    public Optional<Output<String>> applicationSecret() {
        return Optional.ofNullable(this.applicationSecret);
    }

    /**
     * The client email registered for the integration service.
     * 
     */
    @Import(name="clientEmail")
    private @Nullable Output<String> clientEmail;

    /**
     * @return The client email registered for the integration service.
     * 
     */
    public Optional<Output<String>> clientEmail() {
        return Optional.ofNullable(this.clientEmail);
    }

    /**
     * Google Service Account private key credentials.
     * 
     */
    @Import(name="credentials")
    private @Nullable Output<String> credentials;

    /**
     * @return Google Service Account private key credentials.
     * 
     */
    public Optional<Output<String>> credentials() {
        return Optional.ofNullable(this.credentials);
    }

    /**
     * The data collection endpoint for Azure monitor.
     * 
     */
    @Import(name="dceUri")
    private @Nullable Output<String> dceUri;

    /**
     * @return The data collection endpoint for Azure monitor.
     * 
     */
    public Optional<Output<String>> dceUri() {
        return Optional.ofNullable(this.dceUri);
    }

    /**
     * ID of data collection rule that your DCE is linked to for Azure Monitor.
     * 
     * This is the full list of all arguments. Only a subset of arguments are used based on which type of integration used. See Integration Type reference table below for more information.
     * 
     */
    @Import(name="dcrId")
    private @Nullable Output<String> dcrId;

    /**
     * @return ID of data collection rule that your DCE is linked to for Azure Monitor.
     * 
     * This is the full list of all arguments. Only a subset of arguments are used based on which type of integration used. See Integration Type reference table below for more information.
     * 
     */
    public Optional<Output<String>> dcrId() {
        return Optional.ofNullable(this.dcrId);
    }

    /**
     * The syslog destination to send the logs to for Coralogix.
     * 
     */
    @Import(name="endpoint")
    private @Nullable Output<String> endpoint;

    /**
     * @return The syslog destination to send the logs to for Coralogix.
     * 
     */
    public Optional<Output<String>> endpoint() {
        return Optional.ofNullable(this.endpoint);
    }

    /**
     * The host for Scalyr integration. (app.scalyr.com, app.eu.scalyr.com)
     * 
     */
    @Import(name="host")
    private @Nullable Output<String> host;

    /**
     * @return The host for Scalyr integration. (app.scalyr.com, app.eu.scalyr.com)
     * 
     */
    public Optional<Output<String>> host() {
        return Optional.ofNullable(this.host);
    }

    /**
     * Destination to send the logs.
     * 
     */
    @Import(name="hostPort")
    private @Nullable Output<String> hostPort;

    /**
     * @return Destination to send the logs.
     * 
     */
    public Optional<Output<String>> hostPort() {
        return Optional.ofNullable(this.hostPort);
    }

    /**
     * Instance identifier used to make proxy calls
     * 
     */
    @Import(name="instanceId", required=true)
    private Output<Integer> instanceId;

    /**
     * @return Instance identifier used to make proxy calls
     * 
     */
    public Output<Integer> instanceId() {
        return this.instanceId;
    }

    /**
     * The name of the third party log integration. See
     * Integration type reference
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return The name of the third party log integration. See
     * Integration type reference
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * The private access key.
     * 
     */
    @Import(name="privateKey")
    private @Nullable Output<String> privateKey;

    /**
     * @return The private access key.
     * 
     */
    public Optional<Output<String>> privateKey() {
        return Optional.ofNullable(this.privateKey);
    }

    /**
     * Private key identifier. (Stackdriver)
     * 
     */
    @Import(name="privateKeyId")
    private @Nullable Output<String> privateKeyId;

    /**
     * @return Private key identifier. (Stackdriver)
     * 
     */
    public Optional<Output<String>> privateKeyId() {
        return Optional.ofNullable(this.privateKeyId);
    }

    /**
     * The project identifier.
     * 
     */
    @Import(name="projectId")
    private @Nullable Output<String> projectId;

    /**
     * @return The project identifier.
     * 
     */
    public Optional<Output<String>> projectId() {
        return Optional.ofNullable(this.projectId);
    }

    /**
     * Region hosting the integration service.
     * 
     */
    @Import(name="region")
    private @Nullable Output<String> region;

    /**
     * @return Region hosting the integration service.
     * 
     */
    public Optional<Output<String>> region() {
        return Optional.ofNullable(this.region);
    }

    /**
     * AWS secret access key.
     * 
     */
    @Import(name="secretAccessKey")
    private @Nullable Output<String> secretAccessKey;

    /**
     * @return AWS secret access key.
     * 
     */
    public Optional<Output<String>> secretAccessKey() {
        return Optional.ofNullable(this.secretAccessKey);
    }

    /**
     * Assign source type to the data exported, eg. generic_single_line. (Splunk)
     * 
     */
    @Import(name="sourcetype")
    private @Nullable Output<String> sourcetype;

    /**
     * @return Assign source type to the data exported, eg. generic_single_line. (Splunk)
     * 
     */
    public Optional<Output<String>> sourcetype() {
        return Optional.ofNullable(this.sourcetype);
    }

    /**
     * The subsystem name for Coralogix.
     * 
     */
    @Import(name="subsystem")
    private @Nullable Output<String> subsystem;

    /**
     * @return The subsystem name for Coralogix.
     * 
     */
    public Optional<Output<String>> subsystem() {
        return Optional.ofNullable(this.subsystem);
    }

    /**
     * The table name for Azure monitor.
     * 
     */
    @Import(name="table")
    private @Nullable Output<String> table;

    /**
     * @return The table name for Azure monitor.
     * 
     */
    public Optional<Output<String>> table() {
        return Optional.ofNullable(this.table);
    }

    /**
     * Tags. e.g. `env=prod,region=europe`.
     * 
     * ***Note: If tags are used with Datadog. The value part (prod, europe, ...) must start with a letter, read more about tags format in the [Datadog documentation](https://docs.datadoghq.com/getting_started/tagging/#define-tags)***
     * 
     */
    @Import(name="tags")
    private @Nullable Output<String> tags;

    /**
     * @return Tags. e.g. `env=prod,region=europe`.
     * 
     * ***Note: If tags are used with Datadog. The value part (prod, europe, ...) must start with a letter, read more about tags format in the [Datadog documentation](https://docs.datadoghq.com/getting_started/tagging/#define-tags)***
     * 
     */
    public Optional<Output<String>> tags() {
        return Optional.ofNullable(this.tags);
    }

    /**
     * The tenant identifier for Azure monitor.
     * 
     */
    @Import(name="tenantId")
    private @Nullable Output<String> tenantId;

    /**
     * @return The tenant identifier for Azure monitor.
     * 
     */
    public Optional<Output<String>> tenantId() {
        return Optional.ofNullable(this.tenantId);
    }

    /**
     * Token used for authentication.
     * 
     */
    @Import(name="token")
    private @Nullable Output<String> token;

    /**
     * @return Token used for authentication.
     * 
     */
    public Optional<Output<String>> token() {
        return Optional.ofNullable(this.token);
    }

    /**
     * Endpoint to log integration.
     * 
     */
    @Import(name="url")
    private @Nullable Output<String> url;

    /**
     * @return Endpoint to log integration.
     * 
     */
    public Optional<Output<String>> url() {
        return Optional.ofNullable(this.url);
    }

    private IntegrationLogArgs() {}

    private IntegrationLogArgs(IntegrationLogArgs $) {
        this.accessKeyId = $.accessKeyId;
        this.apiKey = $.apiKey;
        this.application = $.application;
        this.applicationId = $.applicationId;
        this.applicationSecret = $.applicationSecret;
        this.clientEmail = $.clientEmail;
        this.credentials = $.credentials;
        this.dceUri = $.dceUri;
        this.dcrId = $.dcrId;
        this.endpoint = $.endpoint;
        this.host = $.host;
        this.hostPort = $.hostPort;
        this.instanceId = $.instanceId;
        this.name = $.name;
        this.privateKey = $.privateKey;
        this.privateKeyId = $.privateKeyId;
        this.projectId = $.projectId;
        this.region = $.region;
        this.secretAccessKey = $.secretAccessKey;
        this.sourcetype = $.sourcetype;
        this.subsystem = $.subsystem;
        this.table = $.table;
        this.tags = $.tags;
        this.tenantId = $.tenantId;
        this.token = $.token;
        this.url = $.url;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(IntegrationLogArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private IntegrationLogArgs $;

        public Builder() {
            $ = new IntegrationLogArgs();
        }

        public Builder(IntegrationLogArgs defaults) {
            $ = new IntegrationLogArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param accessKeyId AWS access key identifier.
         * 
         * @return builder
         * 
         */
        public Builder accessKeyId(@Nullable Output<String> accessKeyId) {
            $.accessKeyId = accessKeyId;
            return this;
        }

        /**
         * @param accessKeyId AWS access key identifier.
         * 
         * @return builder
         * 
         */
        public Builder accessKeyId(String accessKeyId) {
            return accessKeyId(Output.of(accessKeyId));
        }

        /**
         * @param apiKey The API key.
         * 
         * @return builder
         * 
         */
        public Builder apiKey(@Nullable Output<String> apiKey) {
            $.apiKey = apiKey;
            return this;
        }

        /**
         * @param apiKey The API key.
         * 
         * @return builder
         * 
         */
        public Builder apiKey(String apiKey) {
            return apiKey(Output.of(apiKey));
        }

        /**
         * @param application The application name for Coralogix.
         * 
         * @return builder
         * 
         */
        public Builder application(@Nullable Output<String> application) {
            $.application = application;
            return this;
        }

        /**
         * @param application The application name for Coralogix.
         * 
         * @return builder
         * 
         */
        public Builder application(String application) {
            return application(Output.of(application));
        }

        /**
         * @param applicationId The application identifier for Azure monitor.
         * 
         * @return builder
         * 
         */
        public Builder applicationId(@Nullable Output<String> applicationId) {
            $.applicationId = applicationId;
            return this;
        }

        /**
         * @param applicationId The application identifier for Azure monitor.
         * 
         * @return builder
         * 
         */
        public Builder applicationId(String applicationId) {
            return applicationId(Output.of(applicationId));
        }

        /**
         * @param applicationSecret The application secret for Azure monitor.
         * 
         * @return builder
         * 
         */
        public Builder applicationSecret(@Nullable Output<String> applicationSecret) {
            $.applicationSecret = applicationSecret;
            return this;
        }

        /**
         * @param applicationSecret The application secret for Azure monitor.
         * 
         * @return builder
         * 
         */
        public Builder applicationSecret(String applicationSecret) {
            return applicationSecret(Output.of(applicationSecret));
        }

        /**
         * @param clientEmail The client email registered for the integration service.
         * 
         * @return builder
         * 
         */
        public Builder clientEmail(@Nullable Output<String> clientEmail) {
            $.clientEmail = clientEmail;
            return this;
        }

        /**
         * @param clientEmail The client email registered for the integration service.
         * 
         * @return builder
         * 
         */
        public Builder clientEmail(String clientEmail) {
            return clientEmail(Output.of(clientEmail));
        }

        /**
         * @param credentials Google Service Account private key credentials.
         * 
         * @return builder
         * 
         */
        public Builder credentials(@Nullable Output<String> credentials) {
            $.credentials = credentials;
            return this;
        }

        /**
         * @param credentials Google Service Account private key credentials.
         * 
         * @return builder
         * 
         */
        public Builder credentials(String credentials) {
            return credentials(Output.of(credentials));
        }

        /**
         * @param dceUri The data collection endpoint for Azure monitor.
         * 
         * @return builder
         * 
         */
        public Builder dceUri(@Nullable Output<String> dceUri) {
            $.dceUri = dceUri;
            return this;
        }

        /**
         * @param dceUri The data collection endpoint for Azure monitor.
         * 
         * @return builder
         * 
         */
        public Builder dceUri(String dceUri) {
            return dceUri(Output.of(dceUri));
        }

        /**
         * @param dcrId ID of data collection rule that your DCE is linked to for Azure Monitor.
         * 
         * This is the full list of all arguments. Only a subset of arguments are used based on which type of integration used. See Integration Type reference table below for more information.
         * 
         * @return builder
         * 
         */
        public Builder dcrId(@Nullable Output<String> dcrId) {
            $.dcrId = dcrId;
            return this;
        }

        /**
         * @param dcrId ID of data collection rule that your DCE is linked to for Azure Monitor.
         * 
         * This is the full list of all arguments. Only a subset of arguments are used based on which type of integration used. See Integration Type reference table below for more information.
         * 
         * @return builder
         * 
         */
        public Builder dcrId(String dcrId) {
            return dcrId(Output.of(dcrId));
        }

        /**
         * @param endpoint The syslog destination to send the logs to for Coralogix.
         * 
         * @return builder
         * 
         */
        public Builder endpoint(@Nullable Output<String> endpoint) {
            $.endpoint = endpoint;
            return this;
        }

        /**
         * @param endpoint The syslog destination to send the logs to for Coralogix.
         * 
         * @return builder
         * 
         */
        public Builder endpoint(String endpoint) {
            return endpoint(Output.of(endpoint));
        }

        /**
         * @param host The host for Scalyr integration. (app.scalyr.com, app.eu.scalyr.com)
         * 
         * @return builder
         * 
         */
        public Builder host(@Nullable Output<String> host) {
            $.host = host;
            return this;
        }

        /**
         * @param host The host for Scalyr integration. (app.scalyr.com, app.eu.scalyr.com)
         * 
         * @return builder
         * 
         */
        public Builder host(String host) {
            return host(Output.of(host));
        }

        /**
         * @param hostPort Destination to send the logs.
         * 
         * @return builder
         * 
         */
        public Builder hostPort(@Nullable Output<String> hostPort) {
            $.hostPort = hostPort;
            return this;
        }

        /**
         * @param hostPort Destination to send the logs.
         * 
         * @return builder
         * 
         */
        public Builder hostPort(String hostPort) {
            return hostPort(Output.of(hostPort));
        }

        /**
         * @param instanceId Instance identifier used to make proxy calls
         * 
         * @return builder
         * 
         */
        public Builder instanceId(Output<Integer> instanceId) {
            $.instanceId = instanceId;
            return this;
        }

        /**
         * @param instanceId Instance identifier used to make proxy calls
         * 
         * @return builder
         * 
         */
        public Builder instanceId(Integer instanceId) {
            return instanceId(Output.of(instanceId));
        }

        /**
         * @param name The name of the third party log integration. See
         * Integration type reference
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name The name of the third party log integration. See
         * Integration type reference
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param privateKey The private access key.
         * 
         * @return builder
         * 
         */
        public Builder privateKey(@Nullable Output<String> privateKey) {
            $.privateKey = privateKey;
            return this;
        }

        /**
         * @param privateKey The private access key.
         * 
         * @return builder
         * 
         */
        public Builder privateKey(String privateKey) {
            return privateKey(Output.of(privateKey));
        }

        /**
         * @param privateKeyId Private key identifier. (Stackdriver)
         * 
         * @return builder
         * 
         */
        public Builder privateKeyId(@Nullable Output<String> privateKeyId) {
            $.privateKeyId = privateKeyId;
            return this;
        }

        /**
         * @param privateKeyId Private key identifier. (Stackdriver)
         * 
         * @return builder
         * 
         */
        public Builder privateKeyId(String privateKeyId) {
            return privateKeyId(Output.of(privateKeyId));
        }

        /**
         * @param projectId The project identifier.
         * 
         * @return builder
         * 
         */
        public Builder projectId(@Nullable Output<String> projectId) {
            $.projectId = projectId;
            return this;
        }

        /**
         * @param projectId The project identifier.
         * 
         * @return builder
         * 
         */
        public Builder projectId(String projectId) {
            return projectId(Output.of(projectId));
        }

        /**
         * @param region Region hosting the integration service.
         * 
         * @return builder
         * 
         */
        public Builder region(@Nullable Output<String> region) {
            $.region = region;
            return this;
        }

        /**
         * @param region Region hosting the integration service.
         * 
         * @return builder
         * 
         */
        public Builder region(String region) {
            return region(Output.of(region));
        }

        /**
         * @param secretAccessKey AWS secret access key.
         * 
         * @return builder
         * 
         */
        public Builder secretAccessKey(@Nullable Output<String> secretAccessKey) {
            $.secretAccessKey = secretAccessKey;
            return this;
        }

        /**
         * @param secretAccessKey AWS secret access key.
         * 
         * @return builder
         * 
         */
        public Builder secretAccessKey(String secretAccessKey) {
            return secretAccessKey(Output.of(secretAccessKey));
        }

        /**
         * @param sourcetype Assign source type to the data exported, eg. generic_single_line. (Splunk)
         * 
         * @return builder
         * 
         */
        public Builder sourcetype(@Nullable Output<String> sourcetype) {
            $.sourcetype = sourcetype;
            return this;
        }

        /**
         * @param sourcetype Assign source type to the data exported, eg. generic_single_line. (Splunk)
         * 
         * @return builder
         * 
         */
        public Builder sourcetype(String sourcetype) {
            return sourcetype(Output.of(sourcetype));
        }

        /**
         * @param subsystem The subsystem name for Coralogix.
         * 
         * @return builder
         * 
         */
        public Builder subsystem(@Nullable Output<String> subsystem) {
            $.subsystem = subsystem;
            return this;
        }

        /**
         * @param subsystem The subsystem name for Coralogix.
         * 
         * @return builder
         * 
         */
        public Builder subsystem(String subsystem) {
            return subsystem(Output.of(subsystem));
        }

        /**
         * @param table The table name for Azure monitor.
         * 
         * @return builder
         * 
         */
        public Builder table(@Nullable Output<String> table) {
            $.table = table;
            return this;
        }

        /**
         * @param table The table name for Azure monitor.
         * 
         * @return builder
         * 
         */
        public Builder table(String table) {
            return table(Output.of(table));
        }

        /**
         * @param tags Tags. e.g. `env=prod,region=europe`.
         * 
         * ***Note: If tags are used with Datadog. The value part (prod, europe, ...) must start with a letter, read more about tags format in the [Datadog documentation](https://docs.datadoghq.com/getting_started/tagging/#define-tags)***
         * 
         * @return builder
         * 
         */
        public Builder tags(@Nullable Output<String> tags) {
            $.tags = tags;
            return this;
        }

        /**
         * @param tags Tags. e.g. `env=prod,region=europe`.
         * 
         * ***Note: If tags are used with Datadog. The value part (prod, europe, ...) must start with a letter, read more about tags format in the [Datadog documentation](https://docs.datadoghq.com/getting_started/tagging/#define-tags)***
         * 
         * @return builder
         * 
         */
        public Builder tags(String tags) {
            return tags(Output.of(tags));
        }

        /**
         * @param tenantId The tenant identifier for Azure monitor.
         * 
         * @return builder
         * 
         */
        public Builder tenantId(@Nullable Output<String> tenantId) {
            $.tenantId = tenantId;
            return this;
        }

        /**
         * @param tenantId The tenant identifier for Azure monitor.
         * 
         * @return builder
         * 
         */
        public Builder tenantId(String tenantId) {
            return tenantId(Output.of(tenantId));
        }

        /**
         * @param token Token used for authentication.
         * 
         * @return builder
         * 
         */
        public Builder token(@Nullable Output<String> token) {
            $.token = token;
            return this;
        }

        /**
         * @param token Token used for authentication.
         * 
         * @return builder
         * 
         */
        public Builder token(String token) {
            return token(Output.of(token));
        }

        /**
         * @param url Endpoint to log integration.
         * 
         * @return builder
         * 
         */
        public Builder url(@Nullable Output<String> url) {
            $.url = url;
            return this;
        }

        /**
         * @param url Endpoint to log integration.
         * 
         * @return builder
         * 
         */
        public Builder url(String url) {
            return url(Output.of(url));
        }

        public IntegrationLogArgs build() {
            if ($.instanceId == null) {
                throw new MissingRequiredPropertyException("IntegrationLogArgs", "instanceId");
            }
            return $;
        }
    }

}
