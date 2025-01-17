// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package bigtable

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Creates a Google Bigtable instance. For more information see
// [the official documentation](https://cloud.google.com/bigtable/) and
// [API](https://cloud.google.com/bigtable/docs/go/reference).
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/bigtable_instance.html.markdown.
type Instance struct {
	s *pulumi.ResourceState
}

// NewInstance registers a new resource with the given unique name, arguments, and options.
func NewInstance(ctx *pulumi.Context,
	name string, args *InstanceArgs, opts ...pulumi.ResourceOpt) (*Instance, error) {
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["clusters"] = nil
		inputs["displayName"] = nil
		inputs["instanceType"] = nil
		inputs["name"] = nil
		inputs["project"] = nil
	} else {
		inputs["clusters"] = args.Clusters
		inputs["displayName"] = args.DisplayName
		inputs["instanceType"] = args.InstanceType
		inputs["name"] = args.Name
		inputs["project"] = args.Project
	}
	s, err := ctx.RegisterResource("gcp:bigtable/instance:Instance", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Instance{s: s}, nil
}

// GetInstance gets an existing Instance resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetInstance(ctx *pulumi.Context,
	name string, id pulumi.ID, state *InstanceState, opts ...pulumi.ResourceOpt) (*Instance, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["clusters"] = state.Clusters
		inputs["displayName"] = state.DisplayName
		inputs["instanceType"] = state.InstanceType
		inputs["name"] = state.Name
		inputs["project"] = state.Project
	}
	s, err := ctx.ReadResource("gcp:bigtable/instance:Instance", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Instance{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *Instance) URN() pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *Instance) ID() pulumi.IDOutput {
	return r.s.ID()
}

// A block of cluster configuration options. This can be specified 1 or 2 times. See structure below.
func (r *Instance) Clusters() pulumi.ArrayOutput {
	return (pulumi.ArrayOutput)(r.s.State["clusters"])
}

// The human-readable display name of the Bigtable instance. Defaults to the instance `name`.
func (r *Instance) DisplayName() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["displayName"])
}

// The instance type to create. One of `"DEVELOPMENT"` or `"PRODUCTION"`. Defaults to `"PRODUCTION"`.
func (r *Instance) InstanceType() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["instanceType"])
}

// The name (also called Instance Id in the Cloud Console) of the Cloud Bigtable instance.
func (r *Instance) Name() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["name"])
}

// The ID of the project in which the resource belongs. If it
// is not provided, the provider project is used.
func (r *Instance) Project() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["project"])
}

// Input properties used for looking up and filtering Instance resources.
type InstanceState struct {
	// A block of cluster configuration options. This can be specified 1 or 2 times. See structure below.
	Clusters interface{}
	// The human-readable display name of the Bigtable instance. Defaults to the instance `name`.
	DisplayName interface{}
	// The instance type to create. One of `"DEVELOPMENT"` or `"PRODUCTION"`. Defaults to `"PRODUCTION"`.
	InstanceType interface{}
	// The name (also called Instance Id in the Cloud Console) of the Cloud Bigtable instance.
	Name interface{}
	// The ID of the project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project interface{}
}

// The set of arguments for constructing a Instance resource.
type InstanceArgs struct {
	// A block of cluster configuration options. This can be specified 1 or 2 times. See structure below.
	Clusters interface{}
	// The human-readable display name of the Bigtable instance. Defaults to the instance `name`.
	DisplayName interface{}
	// The instance type to create. One of `"DEVELOPMENT"` or `"PRODUCTION"`. Defaults to `"PRODUCTION"`.
	InstanceType interface{}
	// The name (also called Instance Id in the Cloud Console) of the Cloud Bigtable instance.
	Name interface{}
	// The ID of the project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project interface{}
}
