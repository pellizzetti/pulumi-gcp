// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Configures the Google Compute Engine
// [Default Network Tier](https://cloud.google.com/network-tiers/docs/using-network-service-tiers#setting_the_tier_for_all_resources_in_a_project)
// for a project.
// 
// For more information, see,
// [the Project API documentation](https://cloud.google.com/compute/docs/reference/rest/v1/projects/setDefaultNetworkTier).
type ProjectDefaultNetworkTier struct {
	s *pulumi.ResourceState
}

// NewProjectDefaultNetworkTier registers a new resource with the given unique name, arguments, and options.
func NewProjectDefaultNetworkTier(ctx *pulumi.Context,
	name string, args *ProjectDefaultNetworkTierArgs, opts ...pulumi.ResourceOpt) (*ProjectDefaultNetworkTier, error) {
	if args == nil || args.NetworkTier == nil {
		return nil, errors.New("missing required argument 'NetworkTier'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["networkTier"] = nil
		inputs["project"] = nil
	} else {
		inputs["networkTier"] = args.NetworkTier
		inputs["project"] = args.Project
	}
	s, err := ctx.RegisterResource("gcp:compute/projectDefaultNetworkTier:ProjectDefaultNetworkTier", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &ProjectDefaultNetworkTier{s: s}, nil
}

// GetProjectDefaultNetworkTier gets an existing ProjectDefaultNetworkTier resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetProjectDefaultNetworkTier(ctx *pulumi.Context,
	name string, id pulumi.ID, state *ProjectDefaultNetworkTierState, opts ...pulumi.ResourceOpt) (*ProjectDefaultNetworkTier, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["networkTier"] = state.NetworkTier
		inputs["project"] = state.Project
	}
	s, err := ctx.ReadResource("gcp:compute/projectDefaultNetworkTier:ProjectDefaultNetworkTier", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &ProjectDefaultNetworkTier{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *ProjectDefaultNetworkTier) URN() *pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *ProjectDefaultNetworkTier) ID() *pulumi.IDOutput {
	return r.s.ID()
}

// The default network tier to be configured for the project.
// This field can take the following values: `PREMIUM` or `STANDARD`.
func (r *ProjectDefaultNetworkTier) NetworkTier() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["networkTier"])
}

// The ID of the project in which the resource belongs. If it
// is not provided, the provider project is used.
func (r *ProjectDefaultNetworkTier) Project() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["project"])
}

// Input properties used for looking up and filtering ProjectDefaultNetworkTier resources.
type ProjectDefaultNetworkTierState struct {
	// The default network tier to be configured for the project.
	// This field can take the following values: `PREMIUM` or `STANDARD`.
	NetworkTier interface{}
	// The ID of the project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project interface{}
}

// The set of arguments for constructing a ProjectDefaultNetworkTier resource.
type ProjectDefaultNetworkTierArgs struct {
	// The default network tier to be configured for the project.
	// This field can take the following values: `PREMIUM` or `STANDARD`.
	NetworkTier interface{}
	// The ID of the project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project interface{}
}