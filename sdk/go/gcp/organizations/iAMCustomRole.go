// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package organizations

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Allows management of a customized Cloud IAM organization role. For more information see
// [the official documentation](https://cloud.google.com/iam/docs/understanding-custom-roles)
// and
// [API](https://cloud.google.com/iam/reference/rest/v1/organizations.roles).
// 
// > **Warning:** Note that custom roles in GCP have the concept of a soft-delete. There are two issues that may arise
//  from this and how roles are propagated. 1) creating a role may involve undeleting and then updating a role with the
//  same name, possibly causing confusing behavior between undelete and update. 2) A deleted role is permanently deleted
//  after 7 days, but it can take up to 30 more days (i.e. between 7 and 37 days after deletion) before the role name is
//  made available again. This means a deleted role that has been deleted for more than 7 days cannot be changed at all
//  by this provider, and new roles cannot share that name.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/organization_iam_custom_role.html.markdown.
type IAMCustomRole struct {
	s *pulumi.ResourceState
}

// NewIAMCustomRole registers a new resource with the given unique name, arguments, and options.
func NewIAMCustomRole(ctx *pulumi.Context,
	name string, args *IAMCustomRoleArgs, opts ...pulumi.ResourceOpt) (*IAMCustomRole, error) {
	if args == nil || args.OrgId == nil {
		return nil, errors.New("missing required argument 'OrgId'")
	}
	if args == nil || args.Permissions == nil {
		return nil, errors.New("missing required argument 'Permissions'")
	}
	if args == nil || args.RoleId == nil {
		return nil, errors.New("missing required argument 'RoleId'")
	}
	if args == nil || args.Title == nil {
		return nil, errors.New("missing required argument 'Title'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["description"] = nil
		inputs["orgId"] = nil
		inputs["permissions"] = nil
		inputs["roleId"] = nil
		inputs["stage"] = nil
		inputs["title"] = nil
	} else {
		inputs["description"] = args.Description
		inputs["orgId"] = args.OrgId
		inputs["permissions"] = args.Permissions
		inputs["roleId"] = args.RoleId
		inputs["stage"] = args.Stage
		inputs["title"] = args.Title
	}
	inputs["deleted"] = nil
	s, err := ctx.RegisterResource("gcp:organizations/iAMCustomRole:IAMCustomRole", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &IAMCustomRole{s: s}, nil
}

// GetIAMCustomRole gets an existing IAMCustomRole resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIAMCustomRole(ctx *pulumi.Context,
	name string, id pulumi.ID, state *IAMCustomRoleState, opts ...pulumi.ResourceOpt) (*IAMCustomRole, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["deleted"] = state.Deleted
		inputs["description"] = state.Description
		inputs["orgId"] = state.OrgId
		inputs["permissions"] = state.Permissions
		inputs["roleId"] = state.RoleId
		inputs["stage"] = state.Stage
		inputs["title"] = state.Title
	}
	s, err := ctx.ReadResource("gcp:organizations/iAMCustomRole:IAMCustomRole", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &IAMCustomRole{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *IAMCustomRole) URN() pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *IAMCustomRole) ID() pulumi.IDOutput {
	return r.s.ID()
}

// (Optional) The current deleted state of the role.
func (r *IAMCustomRole) Deleted() pulumi.BoolOutput {
	return (pulumi.BoolOutput)(r.s.State["deleted"])
}

// A human-readable description for the role.
func (r *IAMCustomRole) Description() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["description"])
}

// The numeric ID of the organization in which you want to create a custom role.
func (r *IAMCustomRole) OrgId() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["orgId"])
}

// The names of the permissions this role grants when bound in an IAM policy. At least one permission must be specified.
func (r *IAMCustomRole) Permissions() pulumi.ArrayOutput {
	return (pulumi.ArrayOutput)(r.s.State["permissions"])
}

// The role id to use for this role.
func (r *IAMCustomRole) RoleId() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["roleId"])
}

// The current launch stage of the role.
// Defaults to `GA`.
// List of possible stages is [here](https://cloud.google.com/iam/reference/rest/v1/organizations.roles#Role.RoleLaunchStage).
func (r *IAMCustomRole) Stage() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["stage"])
}

// A human-readable title for the role.
func (r *IAMCustomRole) Title() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["title"])
}

// Input properties used for looking up and filtering IAMCustomRole resources.
type IAMCustomRoleState struct {
	// (Optional) The current deleted state of the role.
	Deleted interface{}
	// A human-readable description for the role.
	Description interface{}
	// The numeric ID of the organization in which you want to create a custom role.
	OrgId interface{}
	// The names of the permissions this role grants when bound in an IAM policy. At least one permission must be specified.
	Permissions interface{}
	// The role id to use for this role.
	RoleId interface{}
	// The current launch stage of the role.
	// Defaults to `GA`.
	// List of possible stages is [here](https://cloud.google.com/iam/reference/rest/v1/organizations.roles#Role.RoleLaunchStage).
	Stage interface{}
	// A human-readable title for the role.
	Title interface{}
}

// The set of arguments for constructing a IAMCustomRole resource.
type IAMCustomRoleArgs struct {
	// A human-readable description for the role.
	Description interface{}
	// The numeric ID of the organization in which you want to create a custom role.
	OrgId interface{}
	// The names of the permissions this role grants when bound in an IAM policy. At least one permission must be specified.
	Permissions interface{}
	// The role id to use for this role.
	RoleId interface{}
	// The current launch stage of the role.
	// Defaults to `GA`.
	// List of possible stages is [here](https://cloud.google.com/iam/reference/rest/v1/organizations.roles#Role.RoleLaunchStage).
	Stage interface{}
	// A human-readable title for the role.
	Title interface{}
}
