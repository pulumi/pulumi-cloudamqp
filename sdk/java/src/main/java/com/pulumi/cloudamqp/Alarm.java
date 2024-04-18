// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.cloudamqp;

import com.pulumi.cloudamqp.AlarmArgs;
import com.pulumi.cloudamqp.Utilities;
import com.pulumi.cloudamqp.inputs.AlarmState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * This resource allows you to create and manage alarms to trigger based on a set of conditions. Once triggerd a notification will be sent to the assigned recipients. When creating a new instance, there will also be a set of default alarms (cpu, memory and disk) created. All default alarms uses the default recipient for notifications.
 * 
 * By setting `no_default_alarms` to *true* in `cloudamqp.Instance`. This will create the instance without default alarms and avoid the need to import them to get full control.
 * 
 * Available for all subscription plans, but `lemur`and `tiger`are limited to fewer alarm types. The limited types supported can be seen in the table below in Alarm Type Reference.
 * 
 * ## Example Usage
 * 
 * &lt;details&gt;
 *   &lt;summary&gt;
 *     &lt;b&gt;
 *       &lt;i&gt;Basic example of CPU and memory alarm&lt;/i&gt;
 *     &lt;/b&gt;
 *   &lt;/summary&gt;
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.cloudamqp.Notification;
 * import com.pulumi.cloudamqp.NotificationArgs;
 * import com.pulumi.cloudamqp.Alarm;
 * import com.pulumi.cloudamqp.AlarmArgs;
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
 *         // New recipient
 *         var recipient01 = new Notification(&#34;recipient01&#34;, NotificationArgs.builder()        
 *             .instanceId(instance.id())
 *             .type(&#34;email&#34;)
 *             .value(&#34;alarm@example.com&#34;)
 *             .name(&#34;alarm&#34;)
 *             .build());
 * 
 *         // New cpu alarm
 *         var cpuAlarm = new Alarm(&#34;cpuAlarm&#34;, AlarmArgs.builder()        
 *             .instanceId(instance.id())
 *             .type(&#34;cpu&#34;)
 *             .enabled(true)
 *             .reminderInterval(600)
 *             .valueThreshold(95)
 *             .timeThreshold(600)
 *             .recipients(recipient01.id())
 *             .build());
 * 
 *         // New memory alarm
 *         var memoryAlarm = new Alarm(&#34;memoryAlarm&#34;, AlarmArgs.builder()        
 *             .instanceId(instance.id())
 *             .type(&#34;memory&#34;)
 *             .enabled(true)
 *             .reminderInterval(600)
 *             .valueThreshold(95)
 *             .timeThreshold(600)
 *             .recipients(recipient01.id())
 *             .build());
 * 
 *     }
 * }
 * ```
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * &lt;/details&gt;
 * 
 * &lt;details&gt;
 *   &lt;summary&gt;
 *     &lt;b&gt;
 *       &lt;i&gt;Manage notice alarm, available from v1.29.5&lt;/i&gt;
 *     &lt;/b&gt;
 *   &lt;/summary&gt;
 * 
 * Only one notice alarm can exists and cannot be created, instead the alarm resource will be updated.
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.cloudamqp.Notification;
 * import com.pulumi.cloudamqp.NotificationArgs;
 * import com.pulumi.cloudamqp.Alarm;
 * import com.pulumi.cloudamqp.AlarmArgs;
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
 *         // New recipient
 *         var recipient01 = new Notification(&#34;recipient01&#34;, NotificationArgs.builder()        
 *             .instanceId(instance.id())
 *             .type(&#34;email&#34;)
 *             .value(&#34;alarm@example.com&#34;)
 *             .name(&#34;alarm&#34;)
 *             .build());
 * 
 *         // Update existing notice alarm
 *         var notice = new Alarm(&#34;notice&#34;, AlarmArgs.builder()        
 *             .instanceId(instance.id())
 *             .type(&#34;notice&#34;)
 *             .enabled(true)
 *             .recipients(recipient01.id())
 *             .build());
 * 
 *     }
 * }
 * ```
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * &lt;/details&gt;
 * 
 * ## Alarm Type reference
 * 
 * Supported alarm types: `cpu, memory, disk, queue, connection, flow, consumer, netsplit, server_unreachable, notice`
 * 
 * Required arguments for all alarms: `instance_id, type, enabled`&lt;br&gt;
 * Optional argument for all alarms: `tags, queue_regex, vhost_regex`
 * 
 * | Name | Type | Shared | Dedicated | Required arguments |
 * | ---- | ---- | ---- | ---- | ---- |
 * | CPU | cpu | - | &amp;#10004; | time_threshold, value_threshold |
 * | Memory | memory | - | &amp;#10004;  | time_threshold, value_threshold |
 * | Disk space | disk | - | &amp;#10004;  | time_threshold, value_threshold |
 * | Queue | queue | &amp;#10004;  | &amp;#10004; | time_threshold, value_threshold, queue_regex, vhost_regex, message_type |
 * | Connection | connection | &amp;#10004; | &amp;#10004; | time_threshold, value_threshold |
 * | Connection flow | flow | &amp;#10004; | &amp;#10004; | time_threshold, value_threshold |
 * | Consumer | consumer | &amp;#10004; | &amp;#10004; | time_threshold, value_threshold, queue, vhost |
 * | Netsplit | netsplit | - | &amp;#10004; | time_threshold |
 * | Server unreachable | server_unreachable  | - | &amp;#10004;  | time_threshold |
 * | Notice | notice | &amp;#10004; | &amp;#10004; | |
 * 
 * &gt; Notice alarm is manadatory! Only one can exists and cannot be deleted. Setting `no_default_alarm` to true, will still create this alarm. See updated changes to notice alarm below.
 * 
 * ## Dependency
 * 
 * This resource depends on CloudAMQP instance identifier, `cloudamqp_instance.instance.id`.
 * 
 * ## Notice alarm
 * 
 * There is a limitation for notice alarm in the API backend. This alarm is mandatory, multiple
 * alarms cannot exists or be deleted.
 * 
 * From provider version v1.29.5
 * it&#39;s possible to manage the notice alarm and no longer needs to be imported. Just create the
 * alarm resource as usually and it will be updated with given recipients. If the alarm is deleted
 * it will only be removed from the state file, but will still be enabled in the backend.
 * 
 * ## Import
 * 
 * `cloudamqp_alarm` can be imported using CloudAMQP internal identifier of the alarm together (CSV separated) with the instance identifier. To retrieve the alarm identifier, use [CloudAMQP API](https://docs.cloudamqp.com/cloudamqp_api.html#list-alarms)
 * 
 * ```sh
 * $ pulumi import cloudamqp:index/alarm:Alarm alarm &lt;id&gt;,&lt;instance_id&gt;`
 * ```
 * 
 */
@ResourceType(type="cloudamqp:index/alarm:Alarm")
public class Alarm extends com.pulumi.resources.CustomResource {
    /**
     * Enable or disable the alarm to trigger.
     * 
     */
    @Export(name="enabled", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> enabled;

    /**
     * @return Enable or disable the alarm to trigger.
     * 
     */
    public Output<Boolean> enabled() {
        return this.enabled;
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
     * Message type `(total, unacked, ready)` used by queue alarm type.
     * 
     * Specific argument for `disk` alarm
     * 
     */
    @Export(name="messageType", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> messageType;

    /**
     * @return Message type `(total, unacked, ready)` used by queue alarm type.
     * 
     * Specific argument for `disk` alarm
     * 
     */
    public Output<Optional<String>> messageType() {
        return Codegen.optional(this.messageType);
    }
    /**
     * Regex for which queue to check.
     * 
     */
    @Export(name="queueRegex", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> queueRegex;

    /**
     * @return Regex for which queue to check.
     * 
     */
    public Output<Optional<String>> queueRegex() {
        return Codegen.optional(this.queueRegex);
    }
    /**
     * Identifier for recipient to be notified. Leave empty to notify all recipients.
     * 
     */
    @Export(name="recipients", refs={List.class,Integer.class}, tree="[0,1]")
    private Output<List<Integer>> recipients;

    /**
     * @return Identifier for recipient to be notified. Leave empty to notify all recipients.
     * 
     */
    public Output<List<Integer>> recipients() {
        return this.recipients;
    }
    /**
     * The reminder interval (in seconds) to resend the alarm if not resolved. Set to 0 for no reminders. The Default is 0.
     * 
     */
    @Export(name="reminderInterval", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> reminderInterval;

    /**
     * @return The reminder interval (in seconds) to resend the alarm if not resolved. Set to 0 for no reminders. The Default is 0.
     * 
     */
    public Output<Optional<Integer>> reminderInterval() {
        return Codegen.optional(this.reminderInterval);
    }
    /**
     * The time interval (in seconds) the `value_threshold` should be active before triggering an alarm.
     * 
     */
    @Export(name="timeThreshold", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> timeThreshold;

    /**
     * @return The time interval (in seconds) the `value_threshold` should be active before triggering an alarm.
     * 
     */
    public Output<Optional<Integer>> timeThreshold() {
        return Codegen.optional(this.timeThreshold);
    }
    /**
     * The alarm type, see valid options below.
     * 
     */
    @Export(name="type", refs={String.class}, tree="[0]")
    private Output<String> type;

    /**
     * @return The alarm type, see valid options below.
     * 
     */
    public Output<String> type() {
        return this.type;
    }
    /**
     * Disk value threshold calculation, `fixed, percentage` of disk space remaining.
     * 
     * Based on alarm type, different arguments are flagged as required or optional.
     * 
     */
    @Export(name="valueCalculation", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> valueCalculation;

    /**
     * @return Disk value threshold calculation, `fixed, percentage` of disk space remaining.
     * 
     * Based on alarm type, different arguments are flagged as required or optional.
     * 
     */
    public Output<Optional<String>> valueCalculation() {
        return Codegen.optional(this.valueCalculation);
    }
    /**
     * The value to trigger the alarm for.
     * 
     */
    @Export(name="valueThreshold", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> valueThreshold;

    /**
     * @return The value to trigger the alarm for.
     * 
     */
    public Output<Optional<Integer>> valueThreshold() {
        return Codegen.optional(this.valueThreshold);
    }
    /**
     * Regex for which vhost to check
     * 
     */
    @Export(name="vhostRegex", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> vhostRegex;

    /**
     * @return Regex for which vhost to check
     * 
     */
    public Output<Optional<String>> vhostRegex() {
        return Codegen.optional(this.vhostRegex);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Alarm(String name) {
        this(name, AlarmArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Alarm(String name, AlarmArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Alarm(String name, AlarmArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("cloudamqp:index/alarm:Alarm", name, args == null ? AlarmArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Alarm(String name, Output<String> id, @Nullable AlarmState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("cloudamqp:index/alarm:Alarm", name, state, makeResourceOptions(options, id));
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
    public static Alarm get(String name, Output<String> id, @Nullable AlarmState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Alarm(name, id, state, options);
    }
}
