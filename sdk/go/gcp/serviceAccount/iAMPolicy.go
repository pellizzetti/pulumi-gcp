// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package serviceAccount

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// When managing IAM roles, you can treat a service account either as a resource or as an identity. This resource is to add iam policy bindings to a service account resource **to configure permissions for who can edit the service account**. To configure permissions for a service account to act as an identity that can manage other GCP resources, use the googleProjectIam set of resources.
// 
// Three different resources help you manage your IAM policy for a service account. Each of these resources serves a different use case:
// 
// * `serviceAccount.IAMPolicy`: Authoritative. Sets the IAM policy for the service account and replaces any existing policy already attached.
// * `serviceAccount.IAMBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the service account are preserved.
// * `serviceAccount.IAMMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the service account are preserved.
// 
// > **Note:** `serviceAccount.IAMPolicy` **cannot** be used in conjunction with `serviceAccount.IAMBinding` and `serviceAccount.IAMMember` or they will fight over what your policy should be.
// 
// > **Note:** `serviceAccount.IAMBinding` resources **can be** used in conjunction with `serviceAccount.IAMMember` resources **only if** they do not grant privilege to the same role.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/service_account_iam_policy.html.markdown.
type IAMPolicy struct {
	s *pulumi.ResourceState
}

// NewIAMPolicy registers a new resource with the given unique name, arguments, and options.
func NewIAMPolicy(ctx *pulumi.Context,
	name string, args *IAMPolicyArgs, opts ...pulumi.ResourceOpt) (*IAMPolicy, error) {
	if args == nil || args.PolicyData == nil {
		return nil, errors.New("missing required argument 'PolicyData'")
	}
	if args == nil || args.ServiceAccountId == nil {
		return nil, errors.New("missing required argument 'ServiceAccountId'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["policyData"] = nil
		inputs["serviceAccountId"] = nil
	} else {
		inputs["policyData"] = args.PolicyData
		inputs["serviceAccountId"] = args.ServiceAccountId
	}
	inputs["etag"] = nil
	s, err := ctx.RegisterResource("gcp:serviceAccount/iAMPolicy:IAMPolicy", name, true, inputs, opts...)
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
		inputs["serviceAccountId"] = state.ServiceAccountId
	}
	s, err := ctx.ReadResource("gcp:serviceAccount/iAMPolicy:IAMPolicy", name, id, inputs, opts...)
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

// (Computed) The etag of the service account IAM policy.
func (r *IAMPolicy) Etag() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["etag"])
}

// The policy data generated by
// a `organizations.getIAMPolicy` data source.
func (r *IAMPolicy) PolicyData() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["policyData"])
}

// The fully-qualified name of the service account to apply policy to.
func (r *IAMPolicy) ServiceAccountId() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["serviceAccountId"])
}

// Input properties used for looking up and filtering IAMPolicy resources.
type IAMPolicyState struct {
	// (Computed) The etag of the service account IAM policy.
	Etag interface{}
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData interface{}
	// The fully-qualified name of the service account to apply policy to.
	ServiceAccountId interface{}
}

// The set of arguments for constructing a IAMPolicy resource.
type IAMPolicyArgs struct {
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData interface{}
	// The fully-qualified name of the service account to apply policy to.
	ServiceAccountId interface{}
}
