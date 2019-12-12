// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Compute
{
    /// <summary>
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/compute_router.html.markdown.
    /// </summary>
    public partial class Router : Pulumi.CustomResource
    {
        /// <summary>
        /// BGP information specific to this router.
        /// </summary>
        [Output("bgp")]
        public Output<Outputs.RouterBgp?> Bgp { get; private set; } = null!;

        /// <summary>
        /// Creation timestamp in RFC3339 text format.
        /// </summary>
        [Output("creationTimestamp")]
        public Output<string> CreationTimestamp { get; private set; } = null!;

        /// <summary>
        /// An optional description of this resource.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// Name of the resource. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name
        /// must be 1-63 characters long and match the regular expression '[a-z]([-a-z0-9]*[a-z0-9])?' which means the
        /// first character must be a lowercase letter, and all following characters must be a dash, lowercase letter,
        /// or digit, except the last character, which cannot be a dash.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// A reference to the network to which this router belongs.
        /// </summary>
        [Output("network")]
        public Output<string> Network { get; private set; } = null!;

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        /// <summary>
        /// Region where the router resides.
        /// </summary>
        [Output("region")]
        public Output<string> Region { get; private set; } = null!;

        /// <summary>
        /// The URI of the created resource.
        /// </summary>
        [Output("selfLink")]
        public Output<string> SelfLink { get; private set; } = null!;


        /// <summary>
        /// Create a Router resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Router(string name, RouterArgs args, CustomResourceOptions? options = null)
            : base("gcp:compute/router:Router", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private Router(string name, Input<string> id, RouterState? state = null, CustomResourceOptions? options = null)
            : base("gcp:compute/router:Router", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Router resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Router Get(string name, Input<string> id, RouterState? state = null, CustomResourceOptions? options = null)
        {
            return new Router(name, id, state, options);
        }
    }

    public sealed class RouterArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// BGP information specific to this router.
        /// </summary>
        [Input("bgp")]
        public Input<Inputs.RouterBgpArgs>? Bgp { get; set; }

        /// <summary>
        /// An optional description of this resource.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Name of the resource. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name
        /// must be 1-63 characters long and match the regular expression '[a-z]([-a-z0-9]*[a-z0-9])?' which means the
        /// first character must be a lowercase letter, and all following characters must be a dash, lowercase letter,
        /// or digit, except the last character, which cannot be a dash.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// A reference to the network to which this router belongs.
        /// </summary>
        [Input("network", required: true)]
        public Input<string> Network { get; set; } = null!;

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// Region where the router resides.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        public RouterArgs()
        {
        }
    }

    public sealed class RouterState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// BGP information specific to this router.
        /// </summary>
        [Input("bgp")]
        public Input<Inputs.RouterBgpGetArgs>? Bgp { get; set; }

        /// <summary>
        /// Creation timestamp in RFC3339 text format.
        /// </summary>
        [Input("creationTimestamp")]
        public Input<string>? CreationTimestamp { get; set; }

        /// <summary>
        /// An optional description of this resource.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Name of the resource. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name
        /// must be 1-63 characters long and match the regular expression '[a-z]([-a-z0-9]*[a-z0-9])?' which means the
        /// first character must be a lowercase letter, and all following characters must be a dash, lowercase letter,
        /// or digit, except the last character, which cannot be a dash.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// A reference to the network to which this router belongs.
        /// </summary>
        [Input("network")]
        public Input<string>? Network { get; set; }

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// Region where the router resides.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// The URI of the created resource.
        /// </summary>
        [Input("selfLink")]
        public Input<string>? SelfLink { get; set; }

        public RouterState()
        {
        }
    }

    namespace Inputs
    {

    public sealed class RouterBgpAdvertisedIpRangesArgs : Pulumi.ResourceArgs
    {
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("range")]
        public Input<string>? Range { get; set; }

        public RouterBgpAdvertisedIpRangesArgs()
        {
        }
    }

    public sealed class RouterBgpAdvertisedIpRangesGetArgs : Pulumi.ResourceArgs
    {
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("range")]
        public Input<string>? Range { get; set; }

        public RouterBgpAdvertisedIpRangesGetArgs()
        {
        }
    }

    public sealed class RouterBgpArgs : Pulumi.ResourceArgs
    {
        [Input("advertiseMode")]
        public Input<string>? AdvertiseMode { get; set; }

        [Input("advertisedGroups")]
        private InputList<string>? _advertisedGroups;
        public InputList<string> AdvertisedGroups
        {
            get => _advertisedGroups ?? (_advertisedGroups = new InputList<string>());
            set => _advertisedGroups = value;
        }

        [Input("advertisedIpRanges")]
        private InputList<RouterBgpAdvertisedIpRangesArgs>? _advertisedIpRanges;
        public InputList<RouterBgpAdvertisedIpRangesArgs> AdvertisedIpRanges
        {
            get => _advertisedIpRanges ?? (_advertisedIpRanges = new InputList<RouterBgpAdvertisedIpRangesArgs>());
            set => _advertisedIpRanges = value;
        }

        [Input("asn", required: true)]
        public Input<int> Asn { get; set; } = null!;

        public RouterBgpArgs()
        {
        }
    }

    public sealed class RouterBgpGetArgs : Pulumi.ResourceArgs
    {
        [Input("advertiseMode")]
        public Input<string>? AdvertiseMode { get; set; }

        [Input("advertisedGroups")]
        private InputList<string>? _advertisedGroups;
        public InputList<string> AdvertisedGroups
        {
            get => _advertisedGroups ?? (_advertisedGroups = new InputList<string>());
            set => _advertisedGroups = value;
        }

        [Input("advertisedIpRanges")]
        private InputList<RouterBgpAdvertisedIpRangesGetArgs>? _advertisedIpRanges;
        public InputList<RouterBgpAdvertisedIpRangesGetArgs> AdvertisedIpRanges
        {
            get => _advertisedIpRanges ?? (_advertisedIpRanges = new InputList<RouterBgpAdvertisedIpRangesGetArgs>());
            set => _advertisedIpRanges = value;
        }

        [Input("asn", required: true)]
        public Input<int> Asn { get; set; } = null!;

        public RouterBgpGetArgs()
        {
        }
    }
    }

    namespace Outputs
    {

    [OutputType]
    public sealed class RouterBgp
    {
        public readonly string? AdvertiseMode;
        public readonly ImmutableArray<string> AdvertisedGroups;
        public readonly ImmutableArray<RouterBgpAdvertisedIpRanges> AdvertisedIpRanges;
        public readonly int Asn;

        [OutputConstructor]
        private RouterBgp(
            string? advertiseMode,
            ImmutableArray<string> advertisedGroups,
            ImmutableArray<RouterBgpAdvertisedIpRanges> advertisedIpRanges,
            int asn)
        {
            AdvertiseMode = advertiseMode;
            AdvertisedGroups = advertisedGroups;
            AdvertisedIpRanges = advertisedIpRanges;
            Asn = asn;
        }
    }

    [OutputType]
    public sealed class RouterBgpAdvertisedIpRanges
    {
        public readonly string? Description;
        public readonly string? Range;

        [OutputConstructor]
        private RouterBgpAdvertisedIpRanges(
            string? description,
            string? range)
        {
            Description = description;
            Range = range;
        }
    }
    }
}