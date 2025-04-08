// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.CloudAmqp.Outputs
{

    [OutputType]
    public sealed class GetNodesNodeResult
    {
        /// <summary>
        /// Additional added disk size
        /// </summary>
        public readonly int AdditionalDiskSize;
        /// <summary>
        /// Availability zone the node is hosted in.
        /// </summary>
        public readonly string AvailabilityZone;
        /// <summary>
        /// Is the node configured?
        /// </summary>
        public readonly bool Configured;
        /// <summary>
        /// Subscription plan disk size
        /// </summary>
        public readonly int DiskSize;
        /// <summary>
        /// Currently used Erlang version on the node.
        /// </summary>
        public readonly string ErlangVersion;
        /// <summary>
        /// Enable or disable High-performance Erlang.
        /// </summary>
        public readonly bool Hipe;
        /// <summary>
        /// External hostname assigned to the node.
        /// </summary>
        public readonly string Hostname;
        /// <summary>
        /// Internal hostname assigned to the node.
        /// </summary>
        public readonly string HostnameInternal;
        /// <summary>
        /// Name of the node.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Currently configured Rabbit MQ version on the node.
        /// </summary>
        public readonly string RabbitmqVersion;
        /// <summary>
        /// Is the node running?
        /// </summary>
        public readonly bool Running;

        [OutputConstructor]
        private GetNodesNodeResult(
            int additionalDiskSize,

            string availabilityZone,

            bool configured,

            int diskSize,

            string erlangVersion,

            bool hipe,

            string hostname,

            string hostnameInternal,

            string name,

            string rabbitmqVersion,

            bool running)
        {
            AdditionalDiskSize = additionalDiskSize;
            AvailabilityZone = availabilityZone;
            Configured = configured;
            DiskSize = diskSize;
            ErlangVersion = erlangVersion;
            Hipe = hipe;
            Hostname = hostname;
            HostnameInternal = hostnameInternal;
            Name = name;
            RabbitmqVersion = rabbitmqVersion;
            Running = running;
        }
    }
}
