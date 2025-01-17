// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package projects

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Four different resources help you manage your IAM policy for a project. Each of these resources serves a different use case:
// 
// * `projects.IAMPolicy`: Authoritative. Sets the IAM policy for the project and replaces any existing policy already attached.
// * `projects.IAMBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the project are preserved.
// * `projects.IAMMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the project are preserved.
// * `projects.IAMAuditConfig`: Authoritative for a given service. Updates the IAM policy to enable audit logging for the given service.
// 
// 
// > **Note:** `projects.IAMPolicy` **cannot** be used in conjunction with `projects.IAMBinding`, `projects.IAMMember`, or `projects.IAMAuditConfig` or they will fight over what your policy should be.
// 
// > **Note:** `projects.IAMBinding` resources **can be** used in conjunction with `projects.IAMMember` resources **only if** they do not grant privilege to the same role.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/project_iam_policy.html.markdown.
type IAMPolicy struct {
	s *pulumi.ResourceState
}

// NewIAMPolicy registers a new resource with the given unique name, arguments, and options.
func NewIAMPolicy(ctx *pulumi.Context,
	name string, args *IAMPolicyArgs, opts ...pulumi.ResourceOpt) (*IAMPolicy, error) {
	if args == nil || args.PolicyData == nil {
		return nil, errors.New("missing required argument 'PolicyData'")
	}
	if args == nil || args.Project == nil {
		return nil, errors.New("missing required argument 'Project'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["policyData"] = nil
		inputs["project"] = nil
	} else {
		inputs["policyData"] = args.PolicyData
		inputs["project"] = args.Project
	}
	inputs["etag"] = nil
	s, err := ctx.RegisterResource("gcp:projects/iAMPolicy:IAMPolicy", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &IAMPolicy{s: s}, nil
}

// GetIAMPolicy gets an existing IAMPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIAMPolicy(ctx *pulumi.Context,
	name string, id pulumi.ID, state *IAMPolicyState, opts ...pulumi.ResourceOpt) (*IAMPolicy, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["etag"] = state.Etag
		inputs["policyData"] = state.PolicyData
		inputs["project"] = state.Project
	}
	s, err := ctx.ReadResource("gcp:projects/iAMPolicy:IAMPolicy", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &IAMPolicy{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *IAMPolicy) URN() pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *IAMPolicy) ID() pulumi.IDOutput {
	return r.s.ID()
}

// (Computed) The etag of the project's IAM policy.
func (r *IAMPolicy) Etag() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["etag"])
}

// The `organizations.getIAMPolicy` data source that represents
// the IAM policy that will be applied to the project. The policy will be
// merged with any existing policy applied to the project.
func (r *IAMPolicy) PolicyData() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["policyData"])
}

// The project ID. If not specified for `projects.IAMBinding`, `projects.IAMMember`, or `projects.IAMAuditConfig`, uses the ID of the project configured with the provider.
// Required for `projects.IAMPolicy` - you must explicitly set the project, and it
// will not be inferred from the provider.
func (r *IAMPolicy) Project() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["project"])
}

// Input properties used for looking up and filtering IAMPolicy resources.
type IAMPolicyState struct {
	// (Computed) The etag of the project's IAM policy.
	Etag interface{}
	// The `organizations.getIAMPolicy` data source that represents
	// the IAM policy that will be applied to the project. The policy will be
	// merged with any existing policy applied to the project.
	PolicyData interface{}
	// The project ID. If not specified for `projects.IAMBinding`, `projects.IAMMember`, or `projects.IAMAuditConfig`, uses the ID of the project configured with the provider.
	// Required for `projects.IAMPolicy` - you must explicitly set the project, and it
	// will not be inferred from the provider.
	Project interface{}
}

// The set of arguments for constructing a IAMPolicy resource.
type IAMPolicyArgs struct {
	// The `organizations.getIAMPolicy` data source that represents
	// the IAM policy that will be applied to the project. The policy will be
	// merged with any existing policy applied to the project.
	PolicyData interface{}
	// The project ID. If not specified for `projects.IAMBinding`, `projects.IAMMember`, or `projects.IAMAuditConfig`, uses the ID of the project configured with the provider.
	// Required for `projects.IAMPolicy` - you must explicitly set the project, and it
	// will not be inferred from the provider.
	Project interface{}
}
