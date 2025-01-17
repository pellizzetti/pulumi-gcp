// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Compute
{
    /// <summary>
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/compute_forwarding_rule.html.markdown.
    /// </summary>
    public partial class ForwardingRule : Pulumi.CustomResource
    {
        /// <summary>
        /// For internal TCP/UDP load balancing (i.e. load balancing scheme is INTERNAL and protocol is TCP/UDP), set
        /// this to true to allow packets addressed to any ports to be forwarded to the backends configured with this
        /// forwarding rule. Used with backend service. Cannot be set if port or portRange are set.
        /// </summary>
        [Output("allPorts")]
        public Output<bool?> AllPorts { get; private set; } = null!;

        /// <summary>
        /// If true, clients can access ILB from all regions. Otherwise only allows from the local region the ILB is
        /// located at.
        /// </summary>
        [Output("allowGlobalAccess")]
        public Output<bool?> AllowGlobalAccess { get; private set; } = null!;

        /// <summary>
        /// A BackendService to receive the matched traffic. This is used only for INTERNAL load balancing.
        /// </summary>
        [Output("backendService")]
        public Output<string?> BackendService { get; private set; } = null!;

        /// <summary>
        /// Creation timestamp in RFC3339 text format.
        /// </summary>
        [Output("creationTimestamp")]
        public Output<string> CreationTimestamp { get; private set; } = null!;

        /// <summary>
        /// An optional description of this resource. Provide this property when you create the resource.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The IP address that this forwarding rule is serving on behalf of. Addresses are restricted based on the
        /// forwarding rule's load balancing scheme (EXTERNAL or INTERNAL) and scope (global or regional). When the load
        /// balancing scheme is EXTERNAL, for global forwarding rules, the address must be a global IP, and for regional
        /// forwarding rules, the address must live in the same region as the forwarding rule. If this field is empty,
        /// an ephemeral IPv4 address from the same scope (global or regional) will be assigned. A regional forwarding
        /// rule supports IPv4 only. A global forwarding rule supports either IPv4 or IPv6. When the load balancing
        /// scheme is INTERNAL, this can only be an RFC 1918 IP address belonging to the network/subnet configured for
        /// the forwarding rule. By default, if this field is empty, an ephemeral internal IP address will be
        /// automatically allocated from the IP range of the subnet or network configured for this forwarding rule. An
        /// address must be specified by a literal IP address. ~&gt; **NOTE**: While the API allows you to specify
        /// various resource paths for an address resource instead, Terraform requires this to specifically be an IP
        /// address to avoid needing to fetching the IP address from resource paths on refresh or unnecessary diffs.
        /// </summary>
        [Output("ipAddress")]
        public Output<string> IpAddress { get; private set; } = null!;

        /// <summary>
        /// The IP protocol to which this rule applies. Valid options are TCP, UDP, ESP, AH, SCTP or ICMP. When the load
        /// balancing scheme is INTERNAL, only TCP and UDP are valid.
        /// </summary>
        [Output("ipProtocol")]
        public Output<string> IpProtocol { get; private set; } = null!;

        /// <summary>
        /// The fingerprint used for optimistic locking of this resource. Used internally during updates.
        /// </summary>
        [Output("labelFingerprint")]
        public Output<string> LabelFingerprint { get; private set; } = null!;

        /// <summary>
        /// Labels to apply to this forwarding rule. A list of key-&gt;value pairs.
        /// </summary>
        [Output("labels")]
        public Output<ImmutableDictionary<string, string>?> Labels { get; private set; } = null!;

        /// <summary>
        /// This signifies what the ForwardingRule will be used for and can be EXTERNAL, INTERNAL, or INTERNAL_MANAGED.
        /// EXTERNAL is used for Classic Cloud VPN gateways, protocol forwarding to VMs from an external IP address, and
        /// HTTP(S), SSL Proxy, TCP Proxy, and Network TCP/UDP load balancers. INTERNAL is used for protocol forwarding
        /// to VMs from an internal IP address, and internal TCP/UDP load balancers. INTERNAL_MANAGED is used for
        /// internal HTTP(S) load balancers.
        /// </summary>
        [Output("loadBalancingScheme")]
        public Output<string?> LoadBalancingScheme { get; private set; } = null!;

        /// <summary>
        /// Name of the resource; provided by the client when the resource is created. The name must be 1-63 characters
        /// long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular
        /// expression '[a-z]([-a-z0-9]*[a-z0-9])?' which means the first character must be a lowercase letter, and all
        /// following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be
        /// a dash.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// For internal load balancing, this field identifies the network that the load balanced IP should belong to
        /// for this Forwarding Rule. If this field is not specified, the default network will be used. This field is
        /// only used for INTERNAL load balancing.
        /// </summary>
        [Output("network")]
        public Output<string> Network { get; private set; } = null!;

        /// <summary>
        /// The networking tier used for configuring this address. This field can take the following values: PREMIUM or
        /// STANDARD. If this field is not specified, it is assumed to be PREMIUM.
        /// </summary>
        [Output("networkTier")]
        public Output<string> NetworkTier { get; private set; } = null!;

        /// <summary>
        /// This field is used along with the target field for TargetHttpProxy, TargetHttpsProxy, TargetSslProxy,
        /// TargetTcpProxy, TargetVpnGateway, TargetPool, TargetInstance. Applicable only when IPProtocol is TCP, UDP,
        /// or SCTP, only packets addressed to ports in the specified range will be forwarded to target. Forwarding
        /// rules with the same [IPAddress, IPProtocol] pair must have disjoint port ranges. Some types of forwarding
        /// target have constraints on the acceptable ports: * TargetHttpProxy: 80, 8080 * TargetHttpsProxy: 443 *
        /// TargetTcpProxy: 25, 43, 110, 143, 195, 443, 465, 587, 700, 993, 995, 1883, 5222 * TargetSslProxy: 25, 43,
        /// 110, 143, 195, 443, 465, 587, 700, 993, 995, 1883, 5222 * TargetVpnGateway: 500, 4500
        /// </summary>
        [Output("portRange")]
        public Output<string?> PortRange { get; private set; } = null!;

        /// <summary>
        /// This field is used along with the backend_service field for internal load balancing. When the load balancing
        /// scheme is INTERNAL, a single port or a comma separated list of ports can be configured. Only packets
        /// addressed to these ports will be forwarded to the backends configured with this forwarding rule. You may
        /// specify a maximum of up to 5 ports.
        /// </summary>
        [Output("ports")]
        public Output<ImmutableArray<string>> Ports { get; private set; } = null!;

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        /// <summary>
        /// A reference to the region where the regional forwarding rule resides. This field is not applicable to global
        /// forwarding rules.
        /// </summary>
        [Output("region")]
        public Output<string> Region { get; private set; } = null!;

        /// <summary>
        /// The URI of the created resource.
        /// </summary>
        [Output("selfLink")]
        public Output<string> SelfLink { get; private set; } = null!;

        /// <summary>
        /// An optional prefix to the service name for this Forwarding Rule. If specified, will be the first label of
        /// the fully qualified service name. The label must be 1-63 characters long, and comply with RFC1035.
        /// Specifically, the label must be 1-63 characters long and match the regular expression
        /// '[a-z]([-a-z0-9]*[a-z0-9])?' which means the first character must be a lowercase letter, and all following
        /// characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
        /// This field is only used for INTERNAL load balancing.
        /// </summary>
        [Output("serviceLabel")]
        public Output<string?> ServiceLabel { get; private set; } = null!;

        /// <summary>
        /// The internal fully qualified service name for this Forwarding Rule. This field is only used for INTERNAL
        /// load balancing.
        /// </summary>
        [Output("serviceName")]
        public Output<string> ServiceName { get; private set; } = null!;

        /// <summary>
        /// The subnetwork that the load balanced IP should belong to for this Forwarding Rule. This field is only used
        /// for INTERNAL load balancing. If the network specified is in auto subnet mode, this field is optional.
        /// However, if the network is in custom subnet mode, a subnetwork must be specified.
        /// </summary>
        [Output("subnetwork")]
        public Output<string> Subnetwork { get; private set; } = null!;

        /// <summary>
        /// This field is only used for EXTERNAL load balancing. A reference to a TargetPool resource to receive the
        /// matched traffic. This target must live in the same region as the forwarding rule. The forwarded traffic must
        /// be of a type appropriate to the target object.
        /// </summary>
        [Output("target")]
        public Output<string?> Target { get; private set; } = null!;


        /// <summary>
        /// Create a ForwardingRule resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ForwardingRule(string name, ForwardingRuleArgs? args = null, CustomResourceOptions? options = null)
            : base("gcp:compute/forwardingRule:ForwardingRule", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private ForwardingRule(string name, Input<string> id, ForwardingRuleState? state = null, CustomResourceOptions? options = null)
            : base("gcp:compute/forwardingRule:ForwardingRule", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ForwardingRule resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ForwardingRule Get(string name, Input<string> id, ForwardingRuleState? state = null, CustomResourceOptions? options = null)
        {
            return new ForwardingRule(name, id, state, options);
        }
    }

    public sealed class ForwardingRuleArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// For internal TCP/UDP load balancing (i.e. load balancing scheme is INTERNAL and protocol is TCP/UDP), set
        /// this to true to allow packets addressed to any ports to be forwarded to the backends configured with this
        /// forwarding rule. Used with backend service. Cannot be set if port or portRange are set.
        /// </summary>
        [Input("allPorts")]
        public Input<bool>? AllPorts { get; set; }

        /// <summary>
        /// If true, clients can access ILB from all regions. Otherwise only allows from the local region the ILB is
        /// located at.
        /// </summary>
        [Input("allowGlobalAccess")]
        public Input<bool>? AllowGlobalAccess { get; set; }

        /// <summary>
        /// A BackendService to receive the matched traffic. This is used only for INTERNAL load balancing.
        /// </summary>
        [Input("backendService")]
        public Input<string>? BackendService { get; set; }

        /// <summary>
        /// An optional description of this resource. Provide this property when you create the resource.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The IP address that this forwarding rule is serving on behalf of. Addresses are restricted based on the
        /// forwarding rule's load balancing scheme (EXTERNAL or INTERNAL) and scope (global or regional). When the load
        /// balancing scheme is EXTERNAL, for global forwarding rules, the address must be a global IP, and for regional
        /// forwarding rules, the address must live in the same region as the forwarding rule. If this field is empty,
        /// an ephemeral IPv4 address from the same scope (global or regional) will be assigned. A regional forwarding
        /// rule supports IPv4 only. A global forwarding rule supports either IPv4 or IPv6. When the load balancing
        /// scheme is INTERNAL, this can only be an RFC 1918 IP address belonging to the network/subnet configured for
        /// the forwarding rule. By default, if this field is empty, an ephemeral internal IP address will be
        /// automatically allocated from the IP range of the subnet or network configured for this forwarding rule. An
        /// address must be specified by a literal IP address. ~&gt; **NOTE**: While the API allows you to specify
        /// various resource paths for an address resource instead, Terraform requires this to specifically be an IP
        /// address to avoid needing to fetching the IP address from resource paths on refresh or unnecessary diffs.
        /// </summary>
        [Input("ipAddress")]
        public Input<string>? IpAddress { get; set; }

        /// <summary>
        /// The IP protocol to which this rule applies. Valid options are TCP, UDP, ESP, AH, SCTP or ICMP. When the load
        /// balancing scheme is INTERNAL, only TCP and UDP are valid.
        /// </summary>
        [Input("ipProtocol")]
        public Input<string>? IpProtocol { get; set; }

        [Input("labels")]
        private InputMap<string>? _labels;

        /// <summary>
        /// Labels to apply to this forwarding rule. A list of key-&gt;value pairs.
        /// </summary>
        public InputMap<string> Labels
        {
            get => _labels ?? (_labels = new InputMap<string>());
            set => _labels = value;
        }

        /// <summary>
        /// This signifies what the ForwardingRule will be used for and can be EXTERNAL, INTERNAL, or INTERNAL_MANAGED.
        /// EXTERNAL is used for Classic Cloud VPN gateways, protocol forwarding to VMs from an external IP address, and
        /// HTTP(S), SSL Proxy, TCP Proxy, and Network TCP/UDP load balancers. INTERNAL is used for protocol forwarding
        /// to VMs from an internal IP address, and internal TCP/UDP load balancers. INTERNAL_MANAGED is used for
        /// internal HTTP(S) load balancers.
        /// </summary>
        [Input("loadBalancingScheme")]
        public Input<string>? LoadBalancingScheme { get; set; }

        /// <summary>
        /// Name of the resource; provided by the client when the resource is created. The name must be 1-63 characters
        /// long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular
        /// expression '[a-z]([-a-z0-9]*[a-z0-9])?' which means the first character must be a lowercase letter, and all
        /// following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be
        /// a dash.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// For internal load balancing, this field identifies the network that the load balanced IP should belong to
        /// for this Forwarding Rule. If this field is not specified, the default network will be used. This field is
        /// only used for INTERNAL load balancing.
        /// </summary>
        [Input("network")]
        public Input<string>? Network { get; set; }

        /// <summary>
        /// The networking tier used for configuring this address. This field can take the following values: PREMIUM or
        /// STANDARD. If this field is not specified, it is assumed to be PREMIUM.
        /// </summary>
        [Input("networkTier")]
        public Input<string>? NetworkTier { get; set; }

        /// <summary>
        /// This field is used along with the target field for TargetHttpProxy, TargetHttpsProxy, TargetSslProxy,
        /// TargetTcpProxy, TargetVpnGateway, TargetPool, TargetInstance. Applicable only when IPProtocol is TCP, UDP,
        /// or SCTP, only packets addressed to ports in the specified range will be forwarded to target. Forwarding
        /// rules with the same [IPAddress, IPProtocol] pair must have disjoint port ranges. Some types of forwarding
        /// target have constraints on the acceptable ports: * TargetHttpProxy: 80, 8080 * TargetHttpsProxy: 443 *
        /// TargetTcpProxy: 25, 43, 110, 143, 195, 443, 465, 587, 700, 993, 995, 1883, 5222 * TargetSslProxy: 25, 43,
        /// 110, 143, 195, 443, 465, 587, 700, 993, 995, 1883, 5222 * TargetVpnGateway: 500, 4500
        /// </summary>
        [Input("portRange")]
        public Input<string>? PortRange { get; set; }

        [Input("ports")]
        private InputList<string>? _ports;

        /// <summary>
        /// This field is used along with the backend_service field for internal load balancing. When the load balancing
        /// scheme is INTERNAL, a single port or a comma separated list of ports can be configured. Only packets
        /// addressed to these ports will be forwarded to the backends configured with this forwarding rule. You may
        /// specify a maximum of up to 5 ports.
        /// </summary>
        public InputList<string> Ports
        {
            get => _ports ?? (_ports = new InputList<string>());
            set => _ports = value;
        }

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// A reference to the region where the regional forwarding rule resides. This field is not applicable to global
        /// forwarding rules.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// An optional prefix to the service name for this Forwarding Rule. If specified, will be the first label of
        /// the fully qualified service name. The label must be 1-63 characters long, and comply with RFC1035.
        /// Specifically, the label must be 1-63 characters long and match the regular expression
        /// '[a-z]([-a-z0-9]*[a-z0-9])?' which means the first character must be a lowercase letter, and all following
        /// characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
        /// This field is only used for INTERNAL load balancing.
        /// </summary>
        [Input("serviceLabel")]
        public Input<string>? ServiceLabel { get; set; }

        /// <summary>
        /// The subnetwork that the load balanced IP should belong to for this Forwarding Rule. This field is only used
        /// for INTERNAL load balancing. If the network specified is in auto subnet mode, this field is optional.
        /// However, if the network is in custom subnet mode, a subnetwork must be specified.
        /// </summary>
        [Input("subnetwork")]
        public Input<string>? Subnetwork { get; set; }

        /// <summary>
        /// This field is only used for EXTERNAL load balancing. A reference to a TargetPool resource to receive the
        /// matched traffic. This target must live in the same region as the forwarding rule. The forwarded traffic must
        /// be of a type appropriate to the target object.
        /// </summary>
        [Input("target")]
        public Input<string>? Target { get; set; }

        public ForwardingRuleArgs()
        {
        }
    }

    public sealed class ForwardingRuleState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// For internal TCP/UDP load balancing (i.e. load balancing scheme is INTERNAL and protocol is TCP/UDP), set
        /// this to true to allow packets addressed to any ports to be forwarded to the backends configured with this
        /// forwarding rule. Used with backend service. Cannot be set if port or portRange are set.
        /// </summary>
        [Input("allPorts")]
        public Input<bool>? AllPorts { get; set; }

        /// <summary>
        /// If true, clients can access ILB from all regions. Otherwise only allows from the local region the ILB is
        /// located at.
        /// </summary>
        [Input("allowGlobalAccess")]
        public Input<bool>? AllowGlobalAccess { get; set; }

        /// <summary>
        /// A BackendService to receive the matched traffic. This is used only for INTERNAL load balancing.
        /// </summary>
        [Input("backendService")]
        public Input<string>? BackendService { get; set; }

        /// <summary>
        /// Creation timestamp in RFC3339 text format.
        /// </summary>
        [Input("creationTimestamp")]
        public Input<string>? CreationTimestamp { get; set; }

        /// <summary>
        /// An optional description of this resource. Provide this property when you create the resource.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The IP address that this forwarding rule is serving on behalf of. Addresses are restricted based on the
        /// forwarding rule's load balancing scheme (EXTERNAL or INTERNAL) and scope (global or regional). When the load
        /// balancing scheme is EXTERNAL, for global forwarding rules, the address must be a global IP, and for regional
        /// forwarding rules, the address must live in the same region as the forwarding rule. If this field is empty,
        /// an ephemeral IPv4 address from the same scope (global or regional) will be assigned. A regional forwarding
        /// rule supports IPv4 only. A global forwarding rule supports either IPv4 or IPv6. When the load balancing
        /// scheme is INTERNAL, this can only be an RFC 1918 IP address belonging to the network/subnet configured for
        /// the forwarding rule. By default, if this field is empty, an ephemeral internal IP address will be
        /// automatically allocated from the IP range of the subnet or network configured for this forwarding rule. An
        /// address must be specified by a literal IP address. ~&gt; **NOTE**: While the API allows you to specify
        /// various resource paths for an address resource instead, Terraform requires this to specifically be an IP
        /// address to avoid needing to fetching the IP address from resource paths on refresh or unnecessary diffs.
        /// </summary>
        [Input("ipAddress")]
        public Input<string>? IpAddress { get; set; }

        /// <summary>
        /// The IP protocol to which this rule applies. Valid options are TCP, UDP, ESP, AH, SCTP or ICMP. When the load
        /// balancing scheme is INTERNAL, only TCP and UDP are valid.
        /// </summary>
        [Input("ipProtocol")]
        public Input<string>? IpProtocol { get; set; }

        /// <summary>
        /// The fingerprint used for optimistic locking of this resource. Used internally during updates.
        /// </summary>
        [Input("labelFingerprint")]
        public Input<string>? LabelFingerprint { get; set; }

        [Input("labels")]
        private InputMap<string>? _labels;

        /// <summary>
        /// Labels to apply to this forwarding rule. A list of key-&gt;value pairs.
        /// </summary>
        public InputMap<string> Labels
        {
            get => _labels ?? (_labels = new InputMap<string>());
            set => _labels = value;
        }

        /// <summary>
        /// This signifies what the ForwardingRule will be used for and can be EXTERNAL, INTERNAL, or INTERNAL_MANAGED.
        /// EXTERNAL is used for Classic Cloud VPN gateways, protocol forwarding to VMs from an external IP address, and
        /// HTTP(S), SSL Proxy, TCP Proxy, and Network TCP/UDP load balancers. INTERNAL is used for protocol forwarding
        /// to VMs from an internal IP address, and internal TCP/UDP load balancers. INTERNAL_MANAGED is used for
        /// internal HTTP(S) load balancers.
        /// </summary>
        [Input("loadBalancingScheme")]
        public Input<string>? LoadBalancingScheme { get; set; }

        /// <summary>
        /// Name of the resource; provided by the client when the resource is created. The name must be 1-63 characters
        /// long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular
        /// expression '[a-z]([-a-z0-9]*[a-z0-9])?' which means the first character must be a lowercase letter, and all
        /// following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be
        /// a dash.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// For internal load balancing, this field identifies the network that the load balanced IP should belong to
        /// for this Forwarding Rule. If this field is not specified, the default network will be used. This field is
        /// only used for INTERNAL load balancing.
        /// </summary>
        [Input("network")]
        public Input<string>? Network { get; set; }

        /// <summary>
        /// The networking tier used for configuring this address. This field can take the following values: PREMIUM or
        /// STANDARD. If this field is not specified, it is assumed to be PREMIUM.
        /// </summary>
        [Input("networkTier")]
        public Input<string>? NetworkTier { get; set; }

        /// <summary>
        /// This field is used along with the target field for TargetHttpProxy, TargetHttpsProxy, TargetSslProxy,
        /// TargetTcpProxy, TargetVpnGateway, TargetPool, TargetInstance. Applicable only when IPProtocol is TCP, UDP,
        /// or SCTP, only packets addressed to ports in the specified range will be forwarded to target. Forwarding
        /// rules with the same [IPAddress, IPProtocol] pair must have disjoint port ranges. Some types of forwarding
        /// target have constraints on the acceptable ports: * TargetHttpProxy: 80, 8080 * TargetHttpsProxy: 443 *
        /// TargetTcpProxy: 25, 43, 110, 143, 195, 443, 465, 587, 700, 993, 995, 1883, 5222 * TargetSslProxy: 25, 43,
        /// 110, 143, 195, 443, 465, 587, 700, 993, 995, 1883, 5222 * TargetVpnGateway: 500, 4500
        /// </summary>
        [Input("portRange")]
        public Input<string>? PortRange { get; set; }

        [Input("ports")]
        private InputList<string>? _ports;

        /// <summary>
        /// This field is used along with the backend_service field for internal load balancing. When the load balancing
        /// scheme is INTERNAL, a single port or a comma separated list of ports can be configured. Only packets
        /// addressed to these ports will be forwarded to the backends configured with this forwarding rule. You may
        /// specify a maximum of up to 5 ports.
        /// </summary>
        public InputList<string> Ports
        {
            get => _ports ?? (_ports = new InputList<string>());
            set => _ports = value;
        }

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// A reference to the region where the regional forwarding rule resides. This field is not applicable to global
        /// forwarding rules.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// The URI of the created resource.
        /// </summary>
        [Input("selfLink")]
        public Input<string>? SelfLink { get; set; }

        /// <summary>
        /// An optional prefix to the service name for this Forwarding Rule. If specified, will be the first label of
        /// the fully qualified service name. The label must be 1-63 characters long, and comply with RFC1035.
        /// Specifically, the label must be 1-63 characters long and match the regular expression
        /// '[a-z]([-a-z0-9]*[a-z0-9])?' which means the first character must be a lowercase letter, and all following
        /// characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
        /// This field is only used for INTERNAL load balancing.
        /// </summary>
        [Input("serviceLabel")]
        public Input<string>? ServiceLabel { get; set; }

        /// <summary>
        /// The internal fully qualified service name for this Forwarding Rule. This field is only used for INTERNAL
        /// load balancing.
        /// </summary>
        [Input("serviceName")]
        public Input<string>? ServiceName { get; set; }

        /// <summary>
        /// The subnetwork that the load balanced IP should belong to for this Forwarding Rule. This field is only used
        /// for INTERNAL load balancing. If the network specified is in auto subnet mode, this field is optional.
        /// However, if the network is in custom subnet mode, a subnetwork must be specified.
        /// </summary>
        [Input("subnetwork")]
        public Input<string>? Subnetwork { get; set; }

        /// <summary>
        /// This field is only used for EXTERNAL load balancing. A reference to a TargetPool resource to receive the
        /// matched traffic. This target must live in the same region as the forwarding rule. The forwarded traffic must
        /// be of a type appropriate to the target object.
        /// </summary>
        [Input("target")]
        public Input<string>? Target { get; set; }

        public ForwardingRuleState()
        {
        }
    }
}
