// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.CloudAmqp
{
    /// <summary>
    /// This resource allows you to invoke actions on a specific node.
    /// 
    /// Only available for dedicated subscription plans.
    /// 
    /// ## Example Usage
    /// 
    /// &lt;details&gt;
    ///   &lt;summary&gt;
    ///     &lt;b&gt;
    ///       &lt;i&gt;Already know the node identifier (e.g. from state file)&lt;/i&gt;
    ///     &lt;/b&gt;
    ///   &lt;/summary&gt;
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using CloudAmqp = Pulumi.CloudAmqp;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         // New recipient to receieve notifications
    ///         var nodeAction = new CloudAmqp.NodeActions("nodeAction", new CloudAmqp.NodeActionsArgs
    ///         {
    ///             InstanceId = cloudamqp_instance.Instance.Id,
    ///             NodeName = "&lt;node name&gt;",
    ///             Action = "restart",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// &lt;/details&gt;
    /// 
    /// Using data source `cloudamqp.getNodes` to restart RabbitMQ on all nodes.&lt;/br&gt;
    /// ***Note: RabbitMQ restart on multiple nodes need to be chained, so one node restart at the time.***
    /// 
    /// &lt;details&gt;
    ///   &lt;summary&gt;
    ///     &lt;b&gt;
    ///       &lt;i&gt;Multi node RabbitMQ restart&lt;/i&gt;
    ///     &lt;/b&gt;
    ///   &lt;/summary&gt;
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using CloudAmqp = Pulumi.CloudAmqp;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var listNodes = Output.Create(CloudAmqp.GetNodes.InvokeAsync(new CloudAmqp.GetNodesArgs
    ///         {
    ///             InstanceId = cloudamqp_instance.Instance.Id,
    ///         }));
    ///         var restart01 = new CloudAmqp.NodeActions("restart01", new CloudAmqp.NodeActionsArgs
    ///         {
    ///             InstanceId = cloudamqp_instance.Instance.Id,
    ///             Action = "restart",
    ///             NodeName = listNodes.Apply(listNodes =&gt; listNodes.Nodes?[0]?.Name),
    ///         });
    ///         var restart02 = new CloudAmqp.NodeActions("restart02", new CloudAmqp.NodeActionsArgs
    ///         {
    ///             InstanceId = cloudamqp_instance.Instance.Id,
    ///             Action = "restart",
    ///             NodeName = listNodes.Apply(listNodes =&gt; listNodes.Nodes?[1]?.Name),
    ///         }, new CustomResourceOptions
    ///         {
    ///             DependsOn = 
    ///             {
    ///                 restart01,
    ///             },
    ///         });
    ///         var restart03 = new CloudAmqp.NodeActions("restart03", new CloudAmqp.NodeActionsArgs
    ///         {
    ///             InstanceId = cloudamqp_instance.Instance.Id,
    ///             Action = "restart",
    ///             NodeName = listNodes.Apply(listNodes =&gt; listNodes.Nodes?[2]?.Name),
    ///         }, new CustomResourceOptions
    ///         {
    ///             DependsOn = 
    ///             {
    ///                 restart01,
    ///                 restart02,
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// &lt;/details&gt;
    /// 
    /// &lt;details&gt;
    ///   &lt;summary&gt;
    ///     &lt;b&gt;
    ///       &lt;i&gt;Combine log level configuration change with multi node RabbitMQ restart&lt;/i&gt;
    ///     &lt;/b&gt;
    ///   &lt;/summary&gt;
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using CloudAmqp = Pulumi.CloudAmqp;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var listNodes = Output.Create(CloudAmqp.GetNodes.InvokeAsync(new CloudAmqp.GetNodesArgs
    ///         {
    ///             InstanceId = cloudamqp_instance.Instance.Id,
    ///         }));
    ///         var rabbitmqConfig = new CloudAmqp.RabbitConfiguration("rabbitmqConfig", new CloudAmqp.RabbitConfigurationArgs
    ///         {
    ///             InstanceId = cloudamqp_instance.Instance.Id,
    ///             LogExchangeLevel = "info",
    ///         });
    ///         var restart01 = new CloudAmqp.NodeActions("restart01", new CloudAmqp.NodeActionsArgs
    ///         {
    ///             InstanceId = cloudamqp_instance.Instance.Id,
    ///             Action = "restart",
    ///             NodeName = listNodes.Apply(listNodes =&gt; listNodes.Nodes?[0]?.Name),
    ///         }, new CustomResourceOptions
    ///         {
    ///             DependsOn = 
    ///             {
    ///                 rabbitmqConfig,
    ///             },
    ///         });
    ///         var restart02 = new CloudAmqp.NodeActions("restart02", new CloudAmqp.NodeActionsArgs
    ///         {
    ///             InstanceId = cloudamqp_instance.Instance.Id,
    ///             Action = "restart",
    ///             NodeName = listNodes.Apply(listNodes =&gt; listNodes.Nodes?[1]?.Name),
    ///         }, new CustomResourceOptions
    ///         {
    ///             DependsOn = 
    ///             {
    ///                 rabbitmqConfig,
    ///                 restart01,
    ///             },
    ///         });
    ///         var restart03 = new CloudAmqp.NodeActions("restart03", new CloudAmqp.NodeActionsArgs
    ///         {
    ///             InstanceId = cloudamqp_instance.Instance.Id,
    ///             Action = "restart",
    ///             NodeName = listNodes.Apply(listNodes =&gt; listNodes.Nodes?[2]?.Name),
    ///         }, new CustomResourceOptions
    ///         {
    ///             DependsOn = 
    ///             {
    ///                 rabbitmqConfig,
    ///                 restart01,
    ///                 restart02,
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// &lt;/details&gt;
    /// ## Action reference
    /// 
    /// Valid options for action.
    /// 
    /// | Action       | Info                               |
    /// |--------------|------------------------------------|
    /// | start        | Start RabbitMQ                     |
    /// | stop         | Stop RabbitMQ                      |
    /// | restart      | Restart RabbitMQ                   |
    /// | reboot       | Reboot the node                    |
    /// | mgmt.restart | Restart the RabbitMQ mgmt interace |
    /// 
    /// ## Dependency
    /// 
    /// This resource depends on CloudAMQP instance identifier, `cloudamqp_instance.instance.id` and node name.
    /// 
    /// ## Import
    /// 
    /// This resource cannot be imported.
    /// </summary>
    [CloudAmqpResourceType("cloudamqp:index/nodeActions:NodeActions")]
    public partial class NodeActions : Pulumi.CustomResource
    {
        /// <summary>
        /// The action to invoke on the node.
        /// </summary>
        [Output("action")]
        public Output<string> Action { get; private set; } = null!;

        /// <summary>
        /// The CloudAMQP instance ID.
        /// </summary>
        [Output("instanceId")]
        public Output<int> InstanceId { get; private set; } = null!;

        /// <summary>
        /// The node name, e.g `green-guinea-pig-01`.
        /// </summary>
        [Output("nodeName")]
        public Output<string> NodeName { get; private set; } = null!;

        /// <summary>
        /// If the node is running.
        /// </summary>
        [Output("running")]
        public Output<bool> Running { get; private set; } = null!;


        /// <summary>
        /// Create a NodeActions resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public NodeActions(string name, NodeActionsArgs args, CustomResourceOptions? options = null)
            : base("cloudamqp:index/nodeActions:NodeActions", name, args ?? new NodeActionsArgs(), MakeResourceOptions(options, ""))
        {
        }

        private NodeActions(string name, Input<string> id, NodeActionsState? state = null, CustomResourceOptions? options = null)
            : base("cloudamqp:index/nodeActions:NodeActions", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing NodeActions resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static NodeActions Get(string name, Input<string> id, NodeActionsState? state = null, CustomResourceOptions? options = null)
        {
            return new NodeActions(name, id, state, options);
        }
    }

    public sealed class NodeActionsArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The action to invoke on the node.
        /// </summary>
        [Input("action", required: true)]
        public Input<string> Action { get; set; } = null!;

        /// <summary>
        /// The CloudAMQP instance ID.
        /// </summary>
        [Input("instanceId", required: true)]
        public Input<int> InstanceId { get; set; } = null!;

        /// <summary>
        /// The node name, e.g `green-guinea-pig-01`.
        /// </summary>
        [Input("nodeName", required: true)]
        public Input<string> NodeName { get; set; } = null!;

        public NodeActionsArgs()
        {
        }
    }

    public sealed class NodeActionsState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The action to invoke on the node.
        /// </summary>
        [Input("action")]
        public Input<string>? Action { get; set; }

        /// <summary>
        /// The CloudAMQP instance ID.
        /// </summary>
        [Input("instanceId")]
        public Input<int>? InstanceId { get; set; }

        /// <summary>
        /// The node name, e.g `green-guinea-pig-01`.
        /// </summary>
        [Input("nodeName")]
        public Input<string>? NodeName { get; set; }

        /// <summary>
        /// If the node is running.
        /// </summary>
        [Input("running")]
        public Input<bool>? Running { get; set; }

        public NodeActionsState()
        {
        }
    }
}
