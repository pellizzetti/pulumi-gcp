// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.DataFusion
{
    /// <summary>
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/data_fusion_instance.html.markdown.
    /// </summary>
    public partial class Instance : Pulumi.CustomResource
    {
        /// <summary>
        /// The time the instance was created in RFC3339 UTC "Zulu" format, accurate to nanoseconds.
        /// </summary>
        [Output("createTime")]
        public Output<string> CreateTime { get; private set; } = null!;

        /// <summary>
        /// An optional description of the instance.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// Option to enable Stackdriver Logging.
        /// </summary>
        [Output("enableStackdriverLogging")]
        public Output<bool?> EnableStackdriverLogging { get; private set; } = null!;

        /// <summary>
        /// Option to enable Stackdriver Monitoring.
        /// </summary>
        [Output("enableStackdriverMonitoring")]
        public Output<bool?> EnableStackdriverMonitoring { get; private set; } = null!;

        /// <summary>
        /// The resource labels for instance to use to annotate any related underlying resources, such as Compute Engine
        /// VMs.
        /// </summary>
        [Output("labels")]
        public Output<ImmutableDictionary<string, string>?> Labels { get; private set; } = null!;

        /// <summary>
        /// The ID of the instance or a fully qualified identifier for the instance.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Network configuration options. These are required when a private Data Fusion instance is to be created.
        /// </summary>
        [Output("networkConfig")]
        public Output<Outputs.InstanceNetworkConfig?> NetworkConfig { get; private set; } = null!;

        /// <summary>
        /// Map of additional options used to configure the behavior of Data Fusion instance.
        /// </summary>
        [Output("options")]
        public Output<ImmutableDictionary<string, string>?> Options { get; private set; } = null!;

        /// <summary>
        /// Specifies whether the Data Fusion instance should be private. If set to true, all Data Fusion nodes will
        /// have private IP addresses and will not be able to access the public internet.
        /// </summary>
        [Output("privateInstance")]
        public Output<bool?> PrivateInstance { get; private set; } = null!;

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        /// <summary>
        /// The region of the Data Fusion instance.
        /// </summary>
        [Output("region")]
        public Output<string> Region { get; private set; } = null!;

        /// <summary>
        /// Endpoint on which the Data Fusion UI and REST APIs are accessible.
        /// </summary>
        [Output("serviceEndpoint")]
        public Output<string> ServiceEndpoint { get; private set; } = null!;

        /// <summary>
        /// The current state of this Data Fusion instance. - CREATING: Instance is being created - RUNNING: Instance is
        /// running and ready for requests - FAILED: Instance creation failed - DELETING: Instance is being deleted -
        /// UPGRADING: Instance is being upgraded - RESTARTING: Instance is being restarted
        /// </summary>
        [Output("state")]
        public Output<string> State { get; private set; } = null!;

        /// <summary>
        /// Additional information about the current state of this Data Fusion instance if available.
        /// </summary>
        [Output("stateMessage")]
        public Output<string> StateMessage { get; private set; } = null!;

        /// <summary>
        /// Represents the type of Data Fusion instance. Each type is configured with the default settings for
        /// processing and memory. - BASIC: Basic Data Fusion instance. In Basic type, the user will be able to create
        /// data pipelines using point and click UI. However, there are certain limitations, such as fewer number of
        /// concurrent pipelines, no support for streaming pipelines, etc. - ENTERPRISE: Enterprise Data Fusion
        /// instance. In Enterprise type, the user will have more features available, such as support for streaming
        /// pipelines, higher number of concurrent pipelines, etc.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;

        /// <summary>
        /// The time the instance was last updated in RFC3339 UTC "Zulu" format, accurate to nanoseconds.
        /// </summary>
        [Output("updateTime")]
        public Output<string> UpdateTime { get; private set; } = null!;

        /// <summary>
        /// Current version of the Data Fusion.
        /// </summary>
        [Output("version")]
        public Output<string> Version { get; private set; } = null!;


        /// <summary>
        /// Create a Instance resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Instance(string name, InstanceArgs args, CustomResourceOptions? options = null)
            : base("gcp:datafusion/instance:Instance", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private Instance(string name, Input<string> id, InstanceState? state = null, CustomResourceOptions? options = null)
            : base("gcp:datafusion/instance:Instance", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Instance resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Instance Get(string name, Input<string> id, InstanceState? state = null, CustomResourceOptions? options = null)
        {
            return new Instance(name, id, state, options);
        }
    }

    public sealed class InstanceArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// An optional description of the instance.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Option to enable Stackdriver Logging.
        /// </summary>
        [Input("enableStackdriverLogging")]
        public Input<bool>? EnableStackdriverLogging { get; set; }

        /// <summary>
        /// Option to enable Stackdriver Monitoring.
        /// </summary>
        [Input("enableStackdriverMonitoring")]
        public Input<bool>? EnableStackdriverMonitoring { get; set; }

        [Input("labels")]
        private InputMap<string>? _labels;

        /// <summary>
        /// The resource labels for instance to use to annotate any related underlying resources, such as Compute Engine
        /// VMs.
        /// </summary>
        public InputMap<string> Labels
        {
            get => _labels ?? (_labels = new InputMap<string>());
            set => _labels = value;
        }

        /// <summary>
        /// The ID of the instance or a fully qualified identifier for the instance.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Network configuration options. These are required when a private Data Fusion instance is to be created.
        /// </summary>
        [Input("networkConfig")]
        public Input<Inputs.InstanceNetworkConfigArgs>? NetworkConfig { get; set; }

        [Input("options")]
        private InputMap<string>? _options;

        /// <summary>
        /// Map of additional options used to configure the behavior of Data Fusion instance.
        /// </summary>
        public InputMap<string> Options
        {
            get => _options ?? (_options = new InputMap<string>());
            set => _options = value;
        }

        /// <summary>
        /// Specifies whether the Data Fusion instance should be private. If set to true, all Data Fusion nodes will
        /// have private IP addresses and will not be able to access the public internet.
        /// </summary>
        [Input("privateInstance")]
        public Input<bool>? PrivateInstance { get; set; }

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// The region of the Data Fusion instance.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// Represents the type of Data Fusion instance. Each type is configured with the default settings for
        /// processing and memory. - BASIC: Basic Data Fusion instance. In Basic type, the user will be able to create
        /// data pipelines using point and click UI. However, there are certain limitations, such as fewer number of
        /// concurrent pipelines, no support for streaming pipelines, etc. - ENTERPRISE: Enterprise Data Fusion
        /// instance. In Enterprise type, the user will have more features available, such as support for streaming
        /// pipelines, higher number of concurrent pipelines, etc.
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        public InstanceArgs()
        {
        }
    }

    public sealed class InstanceState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The time the instance was created in RFC3339 UTC "Zulu" format, accurate to nanoseconds.
        /// </summary>
        [Input("createTime")]
        public Input<string>? CreateTime { get; set; }

        /// <summary>
        /// An optional description of the instance.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Option to enable Stackdriver Logging.
        /// </summary>
        [Input("enableStackdriverLogging")]
        public Input<bool>? EnableStackdriverLogging { get; set; }

        /// <summary>
        /// Option to enable Stackdriver Monitoring.
        /// </summary>
        [Input("enableStackdriverMonitoring")]
        public Input<bool>? EnableStackdriverMonitoring { get; set; }

        [Input("labels")]
        private InputMap<string>? _labels;

        /// <summary>
        /// The resource labels for instance to use to annotate any related underlying resources, such as Compute Engine
        /// VMs.
        /// </summary>
        public InputMap<string> Labels
        {
            get => _labels ?? (_labels = new InputMap<string>());
            set => _labels = value;
        }

        /// <summary>
        /// The ID of the instance or a fully qualified identifier for the instance.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Network configuration options. These are required when a private Data Fusion instance is to be created.
        /// </summary>
        [Input("networkConfig")]
        public Input<Inputs.InstanceNetworkConfigGetArgs>? NetworkConfig { get; set; }

        [Input("options")]
        private InputMap<string>? _options;

        /// <summary>
        /// Map of additional options used to configure the behavior of Data Fusion instance.
        /// </summary>
        public InputMap<string> Options
        {
            get => _options ?? (_options = new InputMap<string>());
            set => _options = value;
        }

        /// <summary>
        /// Specifies whether the Data Fusion instance should be private. If set to true, all Data Fusion nodes will
        /// have private IP addresses and will not be able to access the public internet.
        /// </summary>
        [Input("privateInstance")]
        public Input<bool>? PrivateInstance { get; set; }

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// The region of the Data Fusion instance.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// Endpoint on which the Data Fusion UI and REST APIs are accessible.
        /// </summary>
        [Input("serviceEndpoint")]
        public Input<string>? ServiceEndpoint { get; set; }

        /// <summary>
        /// The current state of this Data Fusion instance. - CREATING: Instance is being created - RUNNING: Instance is
        /// running and ready for requests - FAILED: Instance creation failed - DELETING: Instance is being deleted -
        /// UPGRADING: Instance is being upgraded - RESTARTING: Instance is being restarted
        /// </summary>
        [Input("state")]
        public Input<string>? State { get; set; }

        /// <summary>
        /// Additional information about the current state of this Data Fusion instance if available.
        /// </summary>
        [Input("stateMessage")]
        public Input<string>? StateMessage { get; set; }

        /// <summary>
        /// Represents the type of Data Fusion instance. Each type is configured with the default settings for
        /// processing and memory. - BASIC: Basic Data Fusion instance. In Basic type, the user will be able to create
        /// data pipelines using point and click UI. However, there are certain limitations, such as fewer number of
        /// concurrent pipelines, no support for streaming pipelines, etc. - ENTERPRISE: Enterprise Data Fusion
        /// instance. In Enterprise type, the user will have more features available, such as support for streaming
        /// pipelines, higher number of concurrent pipelines, etc.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        /// <summary>
        /// The time the instance was last updated in RFC3339 UTC "Zulu" format, accurate to nanoseconds.
        /// </summary>
        [Input("updateTime")]
        public Input<string>? UpdateTime { get; set; }

        /// <summary>
        /// Current version of the Data Fusion.
        /// </summary>
        [Input("version")]
        public Input<string>? Version { get; set; }

        public InstanceState()
        {
        }
    }

    namespace Inputs
    {

    public sealed class InstanceNetworkConfigArgs : Pulumi.ResourceArgs
    {
        [Input("ipAllocation", required: true)]
        public Input<string> IpAllocation { get; set; } = null!;

        [Input("network", required: true)]
        public Input<string> Network { get; set; } = null!;

        public InstanceNetworkConfigArgs()
        {
        }
    }

    public sealed class InstanceNetworkConfigGetArgs : Pulumi.ResourceArgs
    {
        [Input("ipAllocation", required: true)]
        public Input<string> IpAllocation { get; set; } = null!;

        [Input("network", required: true)]
        public Input<string> Network { get; set; } = null!;

        public InstanceNetworkConfigGetArgs()
        {
        }
    }
    }

    namespace Outputs
    {

    [OutputType]
    public sealed class InstanceNetworkConfig
    {
        public readonly string IpAllocation;
        public readonly string Network;

        [OutputConstructor]
        private InstanceNetworkConfig(
            string ipAllocation,
            string network)
        {
            IpAllocation = ipAllocation;
            Network = network;
        }
    }
    }
}
