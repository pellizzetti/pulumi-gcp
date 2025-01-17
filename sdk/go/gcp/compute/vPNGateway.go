// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/compute_vpn_gateway.html.markdown.
type VPNGateway struct {
	s *pulumi.ResourceState
}

// NewVPNGateway registers a new resource with the given unique name, arguments, and options.
func NewVPNGateway(ctx *pulumi.Context,
	name string, args *VPNGatewayArgs, opts ...pulumi.ResourceOpt) (*VPNGateway, error) {
	if args == nil || args.Network == nil {
		return nil, errors.New("missing required argument 'Network'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["description"] = nil
		inputs["name"] = nil
		inputs["network"] = nil
		inputs["project"] = nil
		inputs["region"] = nil
	} else {
		inputs["description"] = args.Description
		inputs["name"] = args.Name
		inputs["network"] = args.Network
		inputs["project"] = args.Project
		inputs["region"] = args.Region
	}
	inputs["creationTimestamp"] = nil
	inputs["gatewayId"] = nil
	inputs["selfLink"] = nil
	s, err := ctx.RegisterResource("gcp:compute/vPNGateway:VPNGateway", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &VPNGateway{s: s}, nil
}

// GetVPNGateway gets an existing VPNGateway resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVPNGateway(ctx *pulumi.Context,
	name string, id pulumi.ID, state *VPNGatewayState, opts ...pulumi.ResourceOpt) (*VPNGateway, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["creationTimestamp"] = state.CreationTimestamp
		inputs["description"] = state.Description
		inputs["gatewayId"] = state.GatewayId
		inputs["name"] = state.Name
		inputs["network"] = state.Network
		inputs["project"] = state.Project
		inputs["region"] = state.Region
		inputs["selfLink"] = state.SelfLink
	}
	s, err := ctx.ReadResource("gcp:compute/vPNGateway:VPNGateway", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &VPNGateway{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *VPNGateway) URN() pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *VPNGateway) ID() pulumi.IDOutput {
	return r.s.ID()
}

// Creation timestamp in RFC3339 text format.
func (r *VPNGateway) CreationTimestamp() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["creationTimestamp"])
}

// An optional description of this resource.
func (r *VPNGateway) Description() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["description"])
}

// The unique identifier for the resource.
func (r *VPNGateway) GatewayId() pulumi.IntOutput {
	return (pulumi.IntOutput)(r.s.State["gatewayId"])
}

// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and
// comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression
// '[a-z]([-a-z0-9]*[a-z0-9])?' which means the first character must be a lowercase letter, and all following characters
// must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
func (r *VPNGateway) Name() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["name"])
}

// The network this VPN gateway is accepting traffic for.
func (r *VPNGateway) Network() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["network"])
}

// The ID of the project in which the resource belongs.
// If it is not provided, the provider project is used.
func (r *VPNGateway) Project() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["project"])
}

// The region this gateway should sit in.
func (r *VPNGateway) Region() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["region"])
}

// The URI of the created resource.
func (r *VPNGateway) SelfLink() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["selfLink"])
}

// Input properties used for looking up and filtering VPNGateway resources.
type VPNGatewayState struct {
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp interface{}
	// An optional description of this resource.
	Description interface{}
	// The unique identifier for the resource.
	GatewayId interface{}
	// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and
	// comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression
	// '[a-z]([-a-z0-9]*[a-z0-9])?' which means the first character must be a lowercase letter, and all following characters
	// must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name interface{}
	// The network this VPN gateway is accepting traffic for.
	Network interface{}
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project interface{}
	// The region this gateway should sit in.
	Region interface{}
	// The URI of the created resource.
	SelfLink interface{}
}

// The set of arguments for constructing a VPNGateway resource.
type VPNGatewayArgs struct {
	// An optional description of this resource.
	Description interface{}
	// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and
	// comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression
	// '[a-z]([-a-z0-9]*[a-z0-9])?' which means the first character must be a lowercase letter, and all following characters
	// must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name interface{}
	// The network this VPN gateway is accepting traffic for.
	Network interface{}
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project interface{}
	// The region this gateway should sit in.
	Region interface{}
}
