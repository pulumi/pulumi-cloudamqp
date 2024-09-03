// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.cloudamqp.inputs;

import com.pulumi.cloudamqp.inputs.ExtraDiskSizeNodeArgs;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.Integer;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ExtraDiskSizeState extends com.pulumi.resources.ResourceArgs {

    public static final ExtraDiskSizeState Empty = new ExtraDiskSizeState();

    /**
     * When resizing disk, allow cluster downtime to do so
     * 
     */
    @Import(name="allowDowntime")
    private @Nullable Output<Boolean> allowDowntime;

    /**
     * @return When resizing disk, allow cluster downtime to do so
     * 
     */
    public Optional<Output<Boolean>> allowDowntime() {
        return Optional.ofNullable(this.allowDowntime);
    }

    /**
     * Extra disk size in GB
     * 
     */
    @Import(name="extraDiskSize")
    private @Nullable Output<Integer> extraDiskSize;

    /**
     * @return Extra disk size in GB
     * 
     */
    public Optional<Output<Integer>> extraDiskSize() {
        return Optional.ofNullable(this.extraDiskSize);
    }

    /**
     * Instance identifier
     * 
     */
    @Import(name="instanceId")
    private @Nullable Output<Integer> instanceId;

    /**
     * @return Instance identifier
     * 
     */
    public Optional<Output<Integer>> instanceId() {
        return Optional.ofNullable(this.instanceId);
    }

    @Import(name="nodes")
    private @Nullable Output<List<ExtraDiskSizeNodeArgs>> nodes;

    public Optional<Output<List<ExtraDiskSizeNodeArgs>>> nodes() {
        return Optional.ofNullable(this.nodes);
    }

    /**
     * Configurable sleep time in seconds between retries for resizing the disk
     * 
     */
    @Import(name="sleep")
    private @Nullable Output<Integer> sleep;

    /**
     * @return Configurable sleep time in seconds between retries for resizing the disk
     * 
     */
    public Optional<Output<Integer>> sleep() {
        return Optional.ofNullable(this.sleep);
    }

    /**
     * Configurable timeout time in seconds for resizing the disk
     * 
     */
    @Import(name="timeout")
    private @Nullable Output<Integer> timeout;

    /**
     * @return Configurable timeout time in seconds for resizing the disk
     * 
     */
    public Optional<Output<Integer>> timeout() {
        return Optional.ofNullable(this.timeout);
    }

    private ExtraDiskSizeState() {}

    private ExtraDiskSizeState(ExtraDiskSizeState $) {
        this.allowDowntime = $.allowDowntime;
        this.extraDiskSize = $.extraDiskSize;
        this.instanceId = $.instanceId;
        this.nodes = $.nodes;
        this.sleep = $.sleep;
        this.timeout = $.timeout;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ExtraDiskSizeState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ExtraDiskSizeState $;

        public Builder() {
            $ = new ExtraDiskSizeState();
        }

        public Builder(ExtraDiskSizeState defaults) {
            $ = new ExtraDiskSizeState(Objects.requireNonNull(defaults));
        }

        /**
         * @param allowDowntime When resizing disk, allow cluster downtime to do so
         * 
         * @return builder
         * 
         */
        public Builder allowDowntime(@Nullable Output<Boolean> allowDowntime) {
            $.allowDowntime = allowDowntime;
            return this;
        }

        /**
         * @param allowDowntime When resizing disk, allow cluster downtime to do so
         * 
         * @return builder
         * 
         */
        public Builder allowDowntime(Boolean allowDowntime) {
            return allowDowntime(Output.of(allowDowntime));
        }

        /**
         * @param extraDiskSize Extra disk size in GB
         * 
         * @return builder
         * 
         */
        public Builder extraDiskSize(@Nullable Output<Integer> extraDiskSize) {
            $.extraDiskSize = extraDiskSize;
            return this;
        }

        /**
         * @param extraDiskSize Extra disk size in GB
         * 
         * @return builder
         * 
         */
        public Builder extraDiskSize(Integer extraDiskSize) {
            return extraDiskSize(Output.of(extraDiskSize));
        }

        /**
         * @param instanceId Instance identifier
         * 
         * @return builder
         * 
         */
        public Builder instanceId(@Nullable Output<Integer> instanceId) {
            $.instanceId = instanceId;
            return this;
        }

        /**
         * @param instanceId Instance identifier
         * 
         * @return builder
         * 
         */
        public Builder instanceId(Integer instanceId) {
            return instanceId(Output.of(instanceId));
        }

        public Builder nodes(@Nullable Output<List<ExtraDiskSizeNodeArgs>> nodes) {
            $.nodes = nodes;
            return this;
        }

        public Builder nodes(List<ExtraDiskSizeNodeArgs> nodes) {
            return nodes(Output.of(nodes));
        }

        public Builder nodes(ExtraDiskSizeNodeArgs... nodes) {
            return nodes(List.of(nodes));
        }

        /**
         * @param sleep Configurable sleep time in seconds between retries for resizing the disk
         * 
         * @return builder
         * 
         */
        public Builder sleep(@Nullable Output<Integer> sleep) {
            $.sleep = sleep;
            return this;
        }

        /**
         * @param sleep Configurable sleep time in seconds between retries for resizing the disk
         * 
         * @return builder
         * 
         */
        public Builder sleep(Integer sleep) {
            return sleep(Output.of(sleep));
        }

        /**
         * @param timeout Configurable timeout time in seconds for resizing the disk
         * 
         * @return builder
         * 
         */
        public Builder timeout(@Nullable Output<Integer> timeout) {
            $.timeout = timeout;
            return this;
        }

        /**
         * @param timeout Configurable timeout time in seconds for resizing the disk
         * 
         * @return builder
         * 
         */
        public Builder timeout(Integer timeout) {
            return timeout(Output.of(timeout));
        }

        public ExtraDiskSizeState build() {
            return $;
        }
    }

}
