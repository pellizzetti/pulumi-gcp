// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package cloudrun

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/cloud_run_service.html.markdown.
type Service struct {
	s *pulumi.ResourceState
}

// NewService registers a new resource with the given unique name, arguments, and options.
func NewService(ctx *pulumi.Context,
	name string, args *ServiceArgs, opts ...pulumi.ResourceOpt) (*Service, error) {
	if args == nil || args.Location == nil {
		return nil, errors.New("missing required argument 'Location'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["location"] = nil
		inputs["metadata"] = nil
		inputs["name"] = nil
		inputs["project"] = nil
		inputs["template"] = nil
		inputs["traffics"] = nil
	} else {
		inputs["location"] = args.Location
		inputs["metadata"] = args.Metadata
		inputs["name"] = args.Name
		inputs["project"] = args.Project
		inputs["template"] = args.Template
		inputs["traffics"] = args.Traffics
	}
	inputs["status"] = nil
	s, err := ctx.RegisterResource("gcp:cloudrun/service:Service", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Service{s: s}, nil
}

// GetService gets an existing Service resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetService(ctx *pulumi.Context,
	name string, id pulumi.ID, state *ServiceState, opts ...pulumi.ResourceOpt) (*Service, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["location"] = state.Location
		inputs["metadata"] = state.Metadata
		inputs["name"] = state.Name
		inputs["project"] = state.Project
		inputs["status"] = state.Status
		inputs["template"] = state.Template
		inputs["traffics"] = state.Traffics
	}
	s, err := ctx.ReadResource("gcp:cloudrun/service:Service", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Service{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *Service) URN() pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *Service) ID() pulumi.IDOutput {
	return r.s.ID()
}

// The location of the cloud run instance. eg us-central1
func (r *Service) Location() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["location"])
}

// Metadata associated with this Service, including name, namespace, labels, and annotations.
func (r *Service) Metadata() pulumi.Output {
	return r.s.State["metadata"]
}

// Name must be unique within a namespace, within a Cloud Run region. Is required when creating resources. Name is
// primarily intended for creation idempotence and configuration definition. Cannot be updated. More info:
// http://kubernetes.io/docs/user-guide/identifiers#names
func (r *Service) Name() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["name"])
}

func (r *Service) Project() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["project"])
}

// The current status of the Service.
func (r *Service) Status() pulumi.Output {
	return r.s.State["status"]
}

// template holds the latest specification for the Revision to be stamped out. The template references the container image,
// and may also include labels and annotations that should be attached to the Revision. To correlate a Revision, and/or to
// force a Revision to be created when the spec doesn't otherwise change, a nonce label may be provided in the template
// metadata. For more details, see:
// https://github.com/knative/serving/blob/master/docs/client-conventions.md#associate-modifications-with-revisions Cloud
// Run does not currently support referencing a build that is responsible for materializing the container image from
// source.
func (r *Service) Template() pulumi.Output {
	return r.s.State["template"]
}

// Traffic specifies how to distribute traffic over a collection of Knative Revisions and Configurations
func (r *Service) Traffics() pulumi.ArrayOutput {
	return (pulumi.ArrayOutput)(r.s.State["traffics"])
}

// Input properties used for looking up and filtering Service resources.
type ServiceState struct {
	// The location of the cloud run instance. eg us-central1
	Location interface{}
	// Metadata associated with this Service, including name, namespace, labels, and annotations.
	Metadata interface{}
	// Name must be unique within a namespace, within a Cloud Run region. Is required when creating resources. Name is
	// primarily intended for creation idempotence and configuration definition. Cannot be updated. More info:
	// http://kubernetes.io/docs/user-guide/identifiers#names
	Name interface{}
	Project interface{}
	// The current status of the Service.
	Status interface{}
	// template holds the latest specification for the Revision to be stamped out. The template references the container
	// image, and may also include labels and annotations that should be attached to the Revision. To correlate a Revision,
	// and/or to force a Revision to be created when the spec doesn't otherwise change, a nonce label may be provided in the
	// template metadata. For more details, see:
	// https://github.com/knative/serving/blob/master/docs/client-conventions.md#associate-modifications-with-revisions Cloud
	// Run does not currently support referencing a build that is responsible for materializing the container image from
	// source.
	Template interface{}
	// Traffic specifies how to distribute traffic over a collection of Knative Revisions and Configurations
	Traffics interface{}
}

// The set of arguments for constructing a Service resource.
type ServiceArgs struct {
	// The location of the cloud run instance. eg us-central1
	Location interface{}
	// Metadata associated with this Service, including name, namespace, labels, and annotations.
	Metadata interface{}
	// Name must be unique within a namespace, within a Cloud Run region. Is required when creating resources. Name is
	// primarily intended for creation idempotence and configuration definition. Cannot be updated. More info:
	// http://kubernetes.io/docs/user-guide/identifiers#names
	Name interface{}
	Project interface{}
	// template holds the latest specification for the Revision to be stamped out. The template references the container
	// image, and may also include labels and annotations that should be attached to the Revision. To correlate a Revision,
	// and/or to force a Revision to be created when the spec doesn't otherwise change, a nonce label may be provided in the
	// template metadata. For more details, see:
	// https://github.com/knative/serving/blob/master/docs/client-conventions.md#associate-modifications-with-revisions Cloud
	// Run does not currently support referencing a build that is responsible for materializing the container image from
	// source.
	Template interface{}
	// Traffic specifies how to distribute traffic over a collection of Knative Revisions and Configurations
	Traffics interface{}
}
