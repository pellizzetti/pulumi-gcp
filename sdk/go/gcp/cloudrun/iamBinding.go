// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package cloudrun

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/cloud_run_service_iam_binding.html.markdown.
type IamBinding struct {
	s *pulumi.ResourceState
}

// NewIamBinding registers a new resource with the given unique name, arguments, and options.
func NewIamBinding(ctx *pulumi.Context,
	name string, args *IamBindingArgs, opts ...pulumi.ResourceOpt) (*IamBinding, error) {
	if args == nil || args.Members == nil {
		return nil, errors.New("missing required argument 'Members'")
	}
	if args == nil || args.Role == nil {
		return nil, errors.New("missing required argument 'Role'")
	}
	if args == nil || args.Service == nil {
		return nil, errors.New("missing required argument 'Service'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["condition"] = nil
		inputs["location"] = nil
		inputs["members"] = nil
		inputs["project"] = nil
		inputs["role"] = nil
		inputs["service"] = nil
	} else {
		inputs["condition"] = args.Condition
		inputs["location"] = args.Location
		inputs["members"] = args.Members
		inputs["project"] = args.Project
		inputs["role"] = args.Role
		inputs["service"] = args.Service
	}
	inputs["etag"] = nil
	s, err := ctx.RegisterResource("gcp:cloudrun/iamBinding:IamBinding", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &IamBinding{s: s}, nil
}

// GetIamBinding gets an existing IamBinding resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIamBinding(ctx *pulumi.Context,
	name string, id pulumi.ID, state *IamBindingState, opts ...pulumi.ResourceOpt) (*IamBinding, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["condition"] = state.Condition
		inputs["etag"] = state.Etag
		inputs["location"] = state.Location
		inputs["members"] = state.Members
		inputs["project"] = state.Project
		inputs["role"] = state.Role
		inputs["service"] = state.Service
	}
	s, err := ctx.ReadResource("gcp:cloudrun/iamBinding:IamBinding", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &IamBinding{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *IamBinding) URN() pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *IamBinding) ID() pulumi.IDOutput {
	return r.s.ID()
}

func (r *IamBinding) Condition() pulumi.Output {
	return r.s.State["condition"]
}

// (Computed) The etag of the IAM policy.
func (r *IamBinding) Etag() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["etag"])
}

// The location of the cloud run instance. eg us-central1 Used to find the parent resource to bind the IAM policy to
func (r *IamBinding) Location() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["location"])
}

func (r *IamBinding) Members() pulumi.ArrayOutput {
	return (pulumi.ArrayOutput)(r.s.State["members"])
}

// The ID of the project in which the resource belongs.
// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
func (r *IamBinding) Project() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["project"])
}

// The role that should be applied. Only one
// `cloudrun.IamBinding` can be used per role. Note that custom roles must be of the format
// `[projects|organizations]/{parent-name}/roles/{role-name}`.
func (r *IamBinding) Role() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["role"])
}

// Used to find the parent resource to bind the IAM policy to
func (r *IamBinding) Service() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["service"])
}

// Input properties used for looking up and filtering IamBinding resources.
type IamBindingState struct {
	Condition interface{}
	// (Computed) The etag of the IAM policy.
	Etag interface{}
	// The location of the cloud run instance. eg us-central1 Used to find the parent resource to bind the IAM policy to
	Location interface{}
	Members interface{}
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project interface{}
	// The role that should be applied. Only one
	// `cloudrun.IamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role interface{}
	// Used to find the parent resource to bind the IAM policy to
	Service interface{}
}

// The set of arguments for constructing a IamBinding resource.
type IamBindingArgs struct {
	Condition interface{}
	// The location of the cloud run instance. eg us-central1 Used to find the parent resource to bind the IAM policy to
	Location interface{}
	Members interface{}
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project interface{}
	// The role that should be applied. Only one
	// `cloudrun.IamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role interface{}
	// Used to find the parent resource to bind the IAM policy to
	Service interface{}
}
