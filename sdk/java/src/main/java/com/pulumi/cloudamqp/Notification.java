// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.cloudamqp;

import com.pulumi.cloudamqp.NotificationArgs;
import com.pulumi.cloudamqp.Utilities;
import com.pulumi.cloudamqp.inputs.NotificationState;
import com.pulumi.cloudamqp.outputs.NotificationResponder;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * This resource allows you to create and manage recipients to receive alarm notifications. There will
 * always be a default recipient created upon instance creation. This recipient will use team email and
 * receive notifications from default alarms.
 * 
 * Available for all subscription plans.
 * 
 * ## Example Usage
 * 
 * &lt;details&gt;
 *   &lt;summary&gt;
 *     &lt;b&gt;Email recipient&lt;/b&gt;
 *   &lt;/summary&gt;
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.cloudamqp.Notification;
 * import com.pulumi.cloudamqp.NotificationArgs;
 * import java.util.List;
 * import java.util.ArrayList;
 * import java.util.Map;
 * import java.io.File;
 * import java.nio.file.Files;
 * import java.nio.file.Paths;
 * 
 * public class App }{{@code
 *     public static void main(String[] args) }{{@code
 *         Pulumi.run(App::stack);
 *     }}{@code
 * 
 *     public static void stack(Context ctx) }{{@code
 *         var emailRecipient = new Notification("emailRecipient", NotificationArgs.builder()
 *             .instanceId(instance.id())
 *             .type("email")
 *             .value("alarm}{@literal @}{@code example.com")
 *             .name("alarm")
 *             .build());
 * 
 *     }}{@code
 * }}{@code
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * &lt;/details&gt;
 * 
 * &lt;details&gt;
 *   &lt;summary&gt;
 *     &lt;b&gt;OpsGenie recipient with optional responders&lt;/b&gt;
 *   &lt;/summary&gt;
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.cloudamqp.Notification;
 * import com.pulumi.cloudamqp.NotificationArgs;
 * import com.pulumi.cloudamqp.inputs.NotificationResponderArgs;
 * import java.util.List;
 * import java.util.ArrayList;
 * import java.util.Map;
 * import java.io.File;
 * import java.nio.file.Files;
 * import java.nio.file.Paths;
 * 
 * public class App {
 *     public static void main(String[] args) {
 *         Pulumi.run(App::stack);
 *     }
 * 
 *     public static void stack(Context ctx) {
 *         var opsgenieRecipient = new Notification("opsgenieRecipient", NotificationArgs.builder()
 *             .instanceId(instance.id())
 *             .type("opsgenie")
 *             .value("<api-key>")
 *             .name("OpsGenie")
 *             .responders(            
 *                 NotificationResponderArgs.builder()
 *                     .type("team")
 *                     .id("<team-uuid>")
 *                     .build(),
 *                 NotificationResponderArgs.builder()
 *                     .type("user")
 *                     .username("<username>")
 *                     .build())
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * &lt;/details&gt;
 * 
 * &lt;details&gt;
 *   &lt;summary&gt;
 *     &lt;b&gt;Pagerduty recipient with optional dedup key&lt;/b&gt;
 *   &lt;/summary&gt;
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.cloudamqp.Notification;
 * import com.pulumi.cloudamqp.NotificationArgs;
 * import java.util.List;
 * import java.util.ArrayList;
 * import java.util.Map;
 * import java.io.File;
 * import java.nio.file.Files;
 * import java.nio.file.Paths;
 * 
 * public class App {
 *     public static void main(String[] args) {
 *         Pulumi.run(App::stack);
 *     }
 * 
 *     public static void stack(Context ctx) {
 *         var pagerdutyRecipient = new Notification("pagerdutyRecipient", NotificationArgs.builder()
 *             .instanceId(instance.id())
 *             .type("pagerduty")
 *             .value("<integration-key>")
 *             .name("PagerDuty")
 *             .options(Map.of("dedupkey", "DEDUPKEY"))
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * &lt;/details&gt;
 * 
 * &lt;details&gt;
 *   &lt;summary&gt;
 *     &lt;b&gt;Signl4 recipient&lt;/b&gt;
 *   &lt;/summary&gt;
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.cloudamqp.Notification;
 * import com.pulumi.cloudamqp.NotificationArgs;
 * import java.util.List;
 * import java.util.ArrayList;
 * import java.util.Map;
 * import java.io.File;
 * import java.nio.file.Files;
 * import java.nio.file.Paths;
 * 
 * public class App {
 *     public static void main(String[] args) {
 *         Pulumi.run(App::stack);
 *     }
 * 
 *     public static void stack(Context ctx) {
 *         var signl4Recipient = new Notification("signl4Recipient", NotificationArgs.builder()
 *             .instanceId(instance.id())
 *             .type("signl4")
 *             .value("<team-secret>")
 *             .name("Signl4")
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * &lt;/details&gt;
 * 
 * &lt;details&gt;
 *   &lt;summary&gt;
 *     &lt;b&gt;Teams recipient&lt;/b&gt;
 *   &lt;/summary&gt;
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.cloudamqp.Notification;
 * import com.pulumi.cloudamqp.NotificationArgs;
 * import java.util.List;
 * import java.util.ArrayList;
 * import java.util.Map;
 * import java.io.File;
 * import java.nio.file.Files;
 * import java.nio.file.Paths;
 * 
 * public class App {
 *     public static void main(String[] args) {
 *         Pulumi.run(App::stack);
 *     }
 * 
 *     public static void stack(Context ctx) {
 *         var teamsRecipient = new Notification("teamsRecipient", NotificationArgs.builder()
 *             .instanceId(instance.id())
 *             .type("teams")
 *             .value("<teams-webhook-url>")
 *             .name("Teams")
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * &lt;/details&gt;
 * 
 * &lt;details&gt;
 *   &lt;summary&gt;
 *     &lt;b&gt;Victorops recipient with optional routing key (rk)&lt;/b&gt;
 *   &lt;/summary&gt;
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.cloudamqp.Notification;
 * import com.pulumi.cloudamqp.NotificationArgs;
 * import java.util.List;
 * import java.util.ArrayList;
 * import java.util.Map;
 * import java.io.File;
 * import java.nio.file.Files;
 * import java.nio.file.Paths;
 * 
 * public class App {
 *     public static void main(String[] args) {
 *         Pulumi.run(App::stack);
 *     }
 * 
 *     public static void stack(Context ctx) {
 *         var victoropsRecipient = new Notification("victoropsRecipient", NotificationArgs.builder()
 *             .instanceId(instance.id())
 *             .type("victorops")
 *             .value("<integration-key>")
 *             .name("Victorops")
 *             .options(Map.of("rk", "ROUTINGKEY"))
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * &lt;/details&gt;
 * 
 * &lt;details&gt;
 *   &lt;summary&gt;
 *     &lt;b&gt;Slack recipient&lt;/b&gt;
 *   &lt;/summary&gt;
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.cloudamqp.Notification;
 * import com.pulumi.cloudamqp.NotificationArgs;
 * import java.util.List;
 * import java.util.ArrayList;
 * import java.util.Map;
 * import java.io.File;
 * import java.nio.file.Files;
 * import java.nio.file.Paths;
 * 
 * public class App {
 *     public static void main(String[] args) {
 *         Pulumi.run(App::stack);
 *     }
 * 
 *     public static void stack(Context ctx) {
 *         var slackRecipient = new Notification("slackRecipient", NotificationArgs.builder()
 *             .instanceId(instance.id())
 *             .type("slack")
 *             .value("<slack-webhook-url>")
 *             .name("Slack webhook recipient")
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * &lt;/details&gt;
 * 
 * &lt;details&gt;
 *   &lt;summary&gt;
 *     &lt;b&gt;Webhook recipient&lt;/b&gt;
 *   &lt;/summary&gt;
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.cloudamqp.Notification;
 * import com.pulumi.cloudamqp.NotificationArgs;
 * import java.util.List;
 * import java.util.ArrayList;
 * import java.util.Map;
 * import java.io.File;
 * import java.nio.file.Files;
 * import java.nio.file.Paths;
 * 
 * public class App {
 *     public static void main(String[] args) {
 *         Pulumi.run(App::stack);
 *     }
 * 
 *     public static void stack(Context ctx) {
 *         var webhookRecipient = new Notification("webhookRecipient", NotificationArgs.builder()
 *             .instanceId(instance.id())
 *             .type("webhook")
 *             .value("<webhook-url>")
 *             .name("Webhook")
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * &lt;/details&gt;
 * 
 * ## Notification Type reference
 * 
 * Valid options for notification type.
 * 
 * * email
 * * opsgenie
 * * opsgenie-eu
 * * pagerduty
 * * signl4
 * * slack
 * * teams
 * * victorops
 * * webhook
 * 
 * ## Options parameter
 * 
 * |   Type    | Options  |                                                    Description                                                    |                                    Note                                     |
 * |-----------|----------|-------------------------------------------------------------------------------------------------------------------|-----------------------------------------------------------------------------|
 * | Victorops | rk       | Routing key to route alarm                                                                                        | -                                                                           |
 * | PagerDuty | dedupkey | Default the dedup key for PagerDuty is generated depending on what alarm has triggered, but here you can set what | If multiple alarms are triggered using this recipient, since they all share |
 * 
 */
@ResourceType(type="cloudamqp:index/notification:Notification")
public class Notification extends com.pulumi.resources.CustomResource {
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
     * Name of the responder
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return Name of the responder
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * Options argument (e.g. `rk` used for VictorOps routing key).
     * 
     */
    @Export(name="options", refs={Map.class,String.class}, tree="[0,1,1]")
    private Output</* @Nullable */ Map<String,String>> options;

    /**
     * @return Options argument (e.g. `rk` used for VictorOps routing key).
     * 
     */
    public Output<Optional<Map<String,String>>> options() {
        return Codegen.optional(this.options);
    }
    /**
     * An array of reponders (only for OpsGenie). Each `responders` block
     * consists of the field documented below.
     * 
     */
    @Export(name="responders", refs={List.class,NotificationResponder.class}, tree="[0,1]")
    private Output</* @Nullable */ List<NotificationResponder>> responders;

    /**
     * @return An array of reponders (only for OpsGenie). Each `responders` block
     * consists of the field documented below.
     * 
     */
    public Output<Optional<List<NotificationResponder>>> responders() {
        return Codegen.optional(this.responders);
    }
    /**
     * Type of responder. [`team`, `user`, `escalation`, `schedule`]
     * 
     */
    @Export(name="type", refs={String.class}, tree="[0]")
    private Output<String> type;

    /**
     * @return Type of responder. [`team`, `user`, `escalation`, `schedule`]
     * 
     */
    public Output<String> type() {
        return this.type;
    }
    /**
     * Integration/API key or endpoint to send the notification.
     * 
     */
    @Export(name="value", refs={String.class}, tree="[0]")
    private Output<String> value;

    /**
     * @return Integration/API key or endpoint to send the notification.
     * 
     */
    public Output<String> value() {
        return this.value;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Notification(java.lang.String name) {
        this(name, NotificationArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Notification(java.lang.String name, NotificationArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Notification(java.lang.String name, NotificationArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("cloudamqp:index/notification:Notification", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private Notification(java.lang.String name, Output<java.lang.String> id, @Nullable NotificationState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("cloudamqp:index/notification:Notification", name, state, makeResourceOptions(options, id), false);
    }

    private static NotificationArgs makeArgs(NotificationArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? NotificationArgs.Empty : args;
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<java.lang.String> id) {
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
    public static Notification get(java.lang.String name, Output<java.lang.String> id, @Nullable NotificationState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Notification(name, id, state, options);
    }
}
