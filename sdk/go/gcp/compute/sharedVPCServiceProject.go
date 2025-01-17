// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Enables the Google Compute Engine
// [Shared VPC](https://cloud.google.com/compute/docs/shared-vpc)
// feature for a project, assigning it as a Shared VPC service project associated
// with a given host project.
// 
// For more information, see,
// [the Project API documentation](https://cloud.google.com/compute/docs/reference/latest/projects),
// where the Shared VPC feature is referred to by its former name "XPN".
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/compute_shared_vpc_service_project.html.markdown.
type SharedVPCServiceProject struct {
	s *pulumi.ResourceState
}

// NewSharedVPCServiceProject registers a new resource with the given unique name, arguments, and options.
func NewSharedVPCServiceProject(ctx *pulumi.Context,
	name string, args *SharedVPCServiceProjectArgs, opts ...pulumi.ResourceOpt) (*SharedVPCServiceProject, error) {
	if args == nil || args.HostProject == nil {
		return nil, errors.New("missing required argument 'HostProject'")
	}
	if args == nil || args.ServiceProject == nil {
		return nil, errors.New("missing required argument 'ServiceProject'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["hostProject"] = nil
		inputs["serviceProject"] = nil
	} else {
		inputs["hostProject"] = args.HostProject
		inputs["serviceProject"] = args.ServiceProject
	}
	s, err := ctx.RegisterResource("gcp:compute/sharedVPCServiceProject:SharedVPCServiceProject", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &SharedVPCServiceProject{s: s}, nil
}

// GetSharedVPCServiceProject gets an existing SharedVPCServiceProject resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSharedVPCServiceProject(ctx *pulumi.Context,
	name string, id pulumi.ID, state *SharedVPCServiceProjectState, opts ...pulumi.ResourceOpt) (*SharedVPCServiceProject, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["hostProject"] = state.HostProject
		inputs["serviceProject"] = state.ServiceProject
	}
	s, err := ctx.ReadResource("gcp:compute/sharedVPCServiceProject:SharedVPCServiceProject", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &SharedVPCServiceProject{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *SharedVPCServiceProject) URN() pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *SharedVPCServiceProject) ID() pulumi.IDOutput {
	return r.s.ID()
}

// The ID of a host project to associate.
func (r *SharedVPCServiceProject) HostProject() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["hostProject"])
}

// The ID of the project that will serve as a Shared VPC service project.
func (r *SharedVPCServiceProject) ServiceProject() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["serviceProject"])
}

// Input properties used for looking up and filtering SharedVPCServiceProject resources.
type SharedVPCServiceProjectState struct {
	// The ID of a host project to associate.
	HostProject interface{}
	// The ID of the project that will serve as a Shared VPC service project.
	ServiceProject interface{}
}

// The set of arguments for constructing a SharedVPCServiceProject resource.
type SharedVPCServiceProjectArgs struct {
	// The ID of a host project to associate.
	HostProject interface{}
	// The ID of the project that will serve as a Shared VPC service project.
	ServiceProject interface{}
}
