// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package folder

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Allows creation and management of a single member for a single binding within
// the IAM policy for an existing Google Cloud Platform folder.
// 
// > **Note:** This resource _must not_ be used in conjunction with
//    `folder.IAMPolicy` or they will fight over what your policy
//    should be. Similarly, roles controlled by `folder.IAMBinding`
//    should not be assigned to using `folder.IAMMember`.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/folder_iam_member.html.markdown.
type IAMMember struct {
	s *pulumi.ResourceState
}

// NewIAMMember registers a new resource with the given unique name, arguments, and options.
func NewIAMMember(ctx *pulumi.Context,
	name string, args *IAMMemberArgs, opts ...pulumi.ResourceOpt) (*IAMMember, error) {
	if args == nil || args.Folder == nil {
		return nil, errors.New("missing required argument 'Folder'")
	}
	if args == nil || args.Member == nil {
		return nil, errors.New("missing required argument 'Member'")
	}
	if args == nil || args.Role == nil {
		return nil, errors.New("missing required argument 'Role'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["condition"] = nil
		inputs["folder"] = nil
		inputs["member"] = nil
		inputs["role"] = nil
	} else {
		inputs["condition"] = args.Condition
		inputs["folder"] = args.Folder
		inputs["member"] = args.Member
		inputs["role"] = args.Role
	}
	inputs["etag"] = nil
	s, err := ctx.RegisterResource("gcp:folder/iAMMember:IAMMember", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &IAMMember{s: s}, nil
}

// GetIAMMember gets an existing IAMMember resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIAMMember(ctx *pulumi.Context,
	name string, id pulumi.ID, state *IAMMemberState, opts ...pulumi.ResourceOpt) (*IAMMember, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["condition"] = state.Condition
		inputs["etag"] = state.Etag
		inputs["folder"] = state.Folder
		inputs["member"] = state.Member
		inputs["role"] = state.Role
	}
	s, err := ctx.ReadResource("gcp:folder/iAMMember:IAMMember", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &IAMMember{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *IAMMember) URN() pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *IAMMember) ID() pulumi.IDOutput {
	return r.s.ID()
}

func (r *IAMMember) Condition() pulumi.Output {
	return r.s.State["condition"]
}

// (Computed) The etag of the folder's IAM policy.
func (r *IAMMember) Etag() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["etag"])
}

// The resource name of the folder the policy is attached to. Its format is folders/{folder_id}.
func (r *IAMMember) Folder() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["folder"])
}

// The identity that will be granted the privilege in the `role`. For more details on format and restrictions see https://cloud.google.com/billing/reference/rest/v1/Policy#Binding
// This field can have one of the following values:
// * **user:{emailid}**: An email address that represents a specific Google account. For example, alice@gmail.com or joe@example.com.
// * **serviceAccount:{emailid}**: An email address that represents a service account. For example, my-other-app@appspot.gserviceaccount.com.
// * **group:{emailid}**: An email address that represents a Google group. For example, admins@example.com.
// * **domain:{domain}**: A G Suite domain (primary, instead of alias) name that represents all the users of that domain. For example, google.com or example.com.
func (r *IAMMember) Member() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["member"])
}

// The role that should be applied. Note that custom roles must be of the format
// `[projects|organizations]/{parent-name}/roles/{role-name}`.
func (r *IAMMember) Role() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["role"])
}

// Input properties used for looking up and filtering IAMMember resources.
type IAMMemberState struct {
	Condition interface{}
	// (Computed) The etag of the folder's IAM policy.
	Etag interface{}
	// The resource name of the folder the policy is attached to. Its format is folders/{folder_id}.
	Folder interface{}
	// The identity that will be granted the privilege in the `role`. For more details on format and restrictions see https://cloud.google.com/billing/reference/rest/v1/Policy#Binding
	// This field can have one of the following values:
	// * **user:{emailid}**: An email address that represents a specific Google account. For example, alice@gmail.com or joe@example.com.
	// * **serviceAccount:{emailid}**: An email address that represents a service account. For example, my-other-app@appspot.gserviceaccount.com.
	// * **group:{emailid}**: An email address that represents a Google group. For example, admins@example.com.
	// * **domain:{domain}**: A G Suite domain (primary, instead of alias) name that represents all the users of that domain. For example, google.com or example.com.
	Member interface{}
	// The role that should be applied. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role interface{}
}

// The set of arguments for constructing a IAMMember resource.
type IAMMemberArgs struct {
	Condition interface{}
	// The resource name of the folder the policy is attached to. Its format is folders/{folder_id}.
	Folder interface{}
	// The identity that will be granted the privilege in the `role`. For more details on format and restrictions see https://cloud.google.com/billing/reference/rest/v1/Policy#Binding
	// This field can have one of the following values:
	// * **user:{emailid}**: An email address that represents a specific Google account. For example, alice@gmail.com or joe@example.com.
	// * **serviceAccount:{emailid}**: An email address that represents a service account. For example, my-other-app@appspot.gserviceaccount.com.
	// * **group:{emailid}**: An email address that represents a Google group. For example, admins@example.com.
	// * **domain:{domain}**: A G Suite domain (primary, instead of alias) name that represents all the users of that domain. For example, google.com or example.com.
	Member interface{}
	// The role that should be applied. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role interface{}
}
