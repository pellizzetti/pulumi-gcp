// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package dataproc

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Three different resources help you manage IAM policies on dataproc clusters. Each of these resources serves a different use case:
// 
// * `dataproc.ClusterIAMPolicy`: Authoritative. Sets the IAM policy for the cluster and replaces any existing policy already attached.
// * `dataproc.ClusterIAMBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the cluster are preserved.
// * `dataproc.ClusterIAMMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the cluster are preserved.
// 
// > **Note:** `dataproc.ClusterIAMPolicy` **cannot** be used in conjunction with `dataproc.ClusterIAMBinding` and `dataproc.ClusterIAMMember` or they will fight over what your policy should be. In addition, be careful not to accidentally unset ownership of the cluster as `dataproc.ClusterIAMPolicy` replaces the entire policy.
// 
// > **Note:** `dataproc.ClusterIAMBinding` resources **can be** used in conjunction with `dataproc.ClusterIAMMember` resources **only if** they do not grant privilege to the same role.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/dataproc_cluster_iam_member.html.markdown.
type ClusterIAMMember struct {
	s *pulumi.ResourceState
}

// NewClusterIAMMember registers a new resource with the given unique name, arguments, and options.
func NewClusterIAMMember(ctx *pulumi.Context,
	name string, args *ClusterIAMMemberArgs, opts ...pulumi.ResourceOpt) (*ClusterIAMMember, error) {
	if args == nil || args.Cluster == nil {
		return nil, errors.New("missing required argument 'Cluster'")
	}
	if args == nil || args.Member == nil {
		return nil, errors.New("missing required argument 'Member'")
	}
	if args == nil || args.Role == nil {
		return nil, errors.New("missing required argument 'Role'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["cluster"] = nil
		inputs["condition"] = nil
		inputs["member"] = nil
		inputs["project"] = nil
		inputs["region"] = nil
		inputs["role"] = nil
	} else {
		inputs["cluster"] = args.Cluster
		inputs["condition"] = args.Condition
		inputs["member"] = args.Member
		inputs["project"] = args.Project
		inputs["region"] = args.Region
		inputs["role"] = args.Role
	}
	inputs["etag"] = nil
	s, err := ctx.RegisterResource("gcp:dataproc/clusterIAMMember:ClusterIAMMember", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &ClusterIAMMember{s: s}, nil
}

// GetClusterIAMMember gets an existing ClusterIAMMember resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetClusterIAMMember(ctx *pulumi.Context,
	name string, id pulumi.ID, state *ClusterIAMMemberState, opts ...pulumi.ResourceOpt) (*ClusterIAMMember, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["cluster"] = state.Cluster
		inputs["condition"] = state.Condition
		inputs["etag"] = state.Etag
		inputs["member"] = state.Member
		inputs["project"] = state.Project
		inputs["region"] = state.Region
		inputs["role"] = state.Role
	}
	s, err := ctx.ReadResource("gcp:dataproc/clusterIAMMember:ClusterIAMMember", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &ClusterIAMMember{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *ClusterIAMMember) URN() pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *ClusterIAMMember) ID() pulumi.IDOutput {
	return r.s.ID()
}

// The name or relative resource id of the cluster to manage IAM policies for.
func (r *ClusterIAMMember) Cluster() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["cluster"])
}

func (r *ClusterIAMMember) Condition() pulumi.Output {
	return r.s.State["condition"]
}

// (Computed) The etag of the clusters's IAM policy.
func (r *ClusterIAMMember) Etag() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["etag"])
}

func (r *ClusterIAMMember) Member() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["member"])
}

// The project in which the cluster belongs. If it
// is not provided, this provider will use the provider default.
func (r *ClusterIAMMember) Project() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["project"])
}

// The region in which the cluster belongs. If it
// is not provided, this provider will use the provider default.
func (r *ClusterIAMMember) Region() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["region"])
}

// The role that should be applied. Only one
// `dataproc.ClusterIAMBinding` can be used per role. Note that custom roles must be of the format
// `[projects|organizations]/{parent-name}/roles/{role-name}`.
func (r *ClusterIAMMember) Role() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["role"])
}

// Input properties used for looking up and filtering ClusterIAMMember resources.
type ClusterIAMMemberState struct {
	// The name or relative resource id of the cluster to manage IAM policies for.
	Cluster interface{}
	Condition interface{}
	// (Computed) The etag of the clusters's IAM policy.
	Etag interface{}
	Member interface{}
	// The project in which the cluster belongs. If it
	// is not provided, this provider will use the provider default.
	Project interface{}
	// The region in which the cluster belongs. If it
	// is not provided, this provider will use the provider default.
	Region interface{}
	// The role that should be applied. Only one
	// `dataproc.ClusterIAMBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role interface{}
}

// The set of arguments for constructing a ClusterIAMMember resource.
type ClusterIAMMemberArgs struct {
	// The name or relative resource id of the cluster to manage IAM policies for.
	Cluster interface{}
	Condition interface{}
	Member interface{}
	// The project in which the cluster belongs. If it
	// is not provided, this provider will use the provider default.
	Project interface{}
	// The region in which the cluster belongs. If it
	// is not provided, this provider will use the provider default.
	Region interface{}
	// The role that should be applied. Only one
	// `dataproc.ClusterIAMBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role interface{}
}
