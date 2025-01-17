// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package iap

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/iap_app_engine_version_iam_binding.html.markdown.
type AppEngineVersionIamBinding struct {
	s *pulumi.ResourceState
}

// NewAppEngineVersionIamBinding registers a new resource with the given unique name, arguments, and options.
func NewAppEngineVersionIamBinding(ctx *pulumi.Context,
	name string, args *AppEngineVersionIamBindingArgs, opts ...pulumi.ResourceOpt) (*AppEngineVersionIamBinding, error) {
	if args == nil || args.AppId == nil {
		return nil, errors.New("missing required argument 'AppId'")
	}
	if args == nil || args.Members == nil {
		return nil, errors.New("missing required argument 'Members'")
	}
	if args == nil || args.Role == nil {
		return nil, errors.New("missing required argument 'Role'")
	}
	if args == nil || args.Service == nil {
		return nil, errors.New("missing required argument 'Service'")
	}
	if args == nil || args.VersionId == nil {
		return nil, errors.New("missing required argument 'VersionId'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["appId"] = nil
		inputs["condition"] = nil
		inputs["members"] = nil
		inputs["project"] = nil
		inputs["role"] = nil
		inputs["service"] = nil
		inputs["versionId"] = nil
	} else {
		inputs["appId"] = args.AppId
		inputs["condition"] = args.Condition
		inputs["members"] = args.Members
		inputs["project"] = args.Project
		inputs["role"] = args.Role
		inputs["service"] = args.Service
		inputs["versionId"] = args.VersionId
	}
	inputs["etag"] = nil
	s, err := ctx.RegisterResource("gcp:iap/appEngineVersionIamBinding:AppEngineVersionIamBinding", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &AppEngineVersionIamBinding{s: s}, nil
}

// GetAppEngineVersionIamBinding gets an existing AppEngineVersionIamBinding resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAppEngineVersionIamBinding(ctx *pulumi.Context,
	name string, id pulumi.ID, state *AppEngineVersionIamBindingState, opts ...pulumi.ResourceOpt) (*AppEngineVersionIamBinding, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["appId"] = state.AppId
		inputs["condition"] = state.Condition
		inputs["etag"] = state.Etag
		inputs["members"] = state.Members
		inputs["project"] = state.Project
		inputs["role"] = state.Role
		inputs["service"] = state.Service
		inputs["versionId"] = state.VersionId
	}
	s, err := ctx.ReadResource("gcp:iap/appEngineVersionIamBinding:AppEngineVersionIamBinding", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &AppEngineVersionIamBinding{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *AppEngineVersionIamBinding) URN() pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *AppEngineVersionIamBinding) ID() pulumi.IDOutput {
	return r.s.ID()
}

// Id of the App Engine application. Used to find the parent resource to bind the IAM policy to
func (r *AppEngineVersionIamBinding) AppId() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["appId"])
}

func (r *AppEngineVersionIamBinding) Condition() pulumi.Output {
	return r.s.State["condition"]
}

// (Computed) The etag of the IAM policy.
func (r *AppEngineVersionIamBinding) Etag() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["etag"])
}

func (r *AppEngineVersionIamBinding) Members() pulumi.ArrayOutput {
	return (pulumi.ArrayOutput)(r.s.State["members"])
}

// The ID of the project in which the resource belongs.
// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
func (r *AppEngineVersionIamBinding) Project() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["project"])
}

// The role that should be applied. Only one
// `iap.AppEngineVersionIamBinding` can be used per role. Note that custom roles must be of the format
// `[projects|organizations]/{parent-name}/roles/{role-name}`.
func (r *AppEngineVersionIamBinding) Role() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["role"])
}

// Service id of the App Engine application Used to find the parent resource to bind the IAM policy to
func (r *AppEngineVersionIamBinding) Service() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["service"])
}

// Version id of the App Engine application Used to find the parent resource to bind the IAM policy to
func (r *AppEngineVersionIamBinding) VersionId() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["versionId"])
}

// Input properties used for looking up and filtering AppEngineVersionIamBinding resources.
type AppEngineVersionIamBindingState struct {
	// Id of the App Engine application. Used to find the parent resource to bind the IAM policy to
	AppId interface{}
	Condition interface{}
	// (Computed) The etag of the IAM policy.
	Etag interface{}
	Members interface{}
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project interface{}
	// The role that should be applied. Only one
	// `iap.AppEngineVersionIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role interface{}
	// Service id of the App Engine application Used to find the parent resource to bind the IAM policy to
	Service interface{}
	// Version id of the App Engine application Used to find the parent resource to bind the IAM policy to
	VersionId interface{}
}

// The set of arguments for constructing a AppEngineVersionIamBinding resource.
type AppEngineVersionIamBindingArgs struct {
	// Id of the App Engine application. Used to find the parent resource to bind the IAM policy to
	AppId interface{}
	Condition interface{}
	Members interface{}
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project interface{}
	// The role that should be applied. Only one
	// `iap.AppEngineVersionIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role interface{}
	// Service id of the App Engine application Used to find the parent resource to bind the IAM policy to
	Service interface{}
	// Version id of the App Engine application Used to find the parent resource to bind the IAM policy to
	VersionId interface{}
}
