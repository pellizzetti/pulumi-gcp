// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package containeranalysis

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/container_analysis_note.html.markdown.
type Note struct {
	s *pulumi.ResourceState
}

// NewNote registers a new resource with the given unique name, arguments, and options.
func NewNote(ctx *pulumi.Context,
	name string, args *NoteArgs, opts ...pulumi.ResourceOpt) (*Note, error) {
	if args == nil || args.AttestationAuthority == nil {
		return nil, errors.New("missing required argument 'AttestationAuthority'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["attestationAuthority"] = nil
		inputs["name"] = nil
		inputs["project"] = nil
	} else {
		inputs["attestationAuthority"] = args.AttestationAuthority
		inputs["name"] = args.Name
		inputs["project"] = args.Project
	}
	s, err := ctx.RegisterResource("gcp:containeranalysis/note:Note", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Note{s: s}, nil
}

// GetNote gets an existing Note resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNote(ctx *pulumi.Context,
	name string, id pulumi.ID, state *NoteState, opts ...pulumi.ResourceOpt) (*Note, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["attestationAuthority"] = state.AttestationAuthority
		inputs["name"] = state.Name
		inputs["project"] = state.Project
	}
	s, err := ctx.ReadResource("gcp:containeranalysis/note:Note", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Note{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *Note) URN() pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *Note) ID() pulumi.IDOutput {
	return r.s.ID()
}

// Note kind that represents a logical attestation "role" or "authority". For example, an organization might have one
// AttestationAuthority for "QA" and one for "build". This Note is intended to act strictly as a grouping mechanism for the
// attached Occurrences (Attestations). This grouping mechanism also provides a security boundary, since IAM ACLs gate the
// ability for a principle to attach an Occurrence to a given Note. It also provides a single point of lookup to find all
// attached Attestation Occurrences, even if they don't all live in the same project.
func (r *Note) AttestationAuthority() pulumi.Output {
	return r.s.State["attestationAuthority"]
}

// The name of the note.
func (r *Note) Name() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["name"])
}

func (r *Note) Project() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["project"])
}

// Input properties used for looking up and filtering Note resources.
type NoteState struct {
	// Note kind that represents a logical attestation "role" or "authority". For example, an organization might have one
	// AttestationAuthority for "QA" and one for "build". This Note is intended to act strictly as a grouping mechanism for
	// the attached Occurrences (Attestations). This grouping mechanism also provides a security boundary, since IAM ACLs gate
	// the ability for a principle to attach an Occurrence to a given Note. It also provides a single point of lookup to find
	// all attached Attestation Occurrences, even if they don't all live in the same project.
	AttestationAuthority interface{}
	// The name of the note.
	Name interface{}
	Project interface{}
}

// The set of arguments for constructing a Note resource.
type NoteArgs struct {
	// Note kind that represents a logical attestation "role" or "authority". For example, an organization might have one
	// AttestationAuthority for "QA" and one for "build". This Note is intended to act strictly as a grouping mechanism for
	// the attached Occurrences (Attestations). This grouping mechanism also provides a security boundary, since IAM ACLs gate
	// the ability for a principle to attach an Occurrence to a given Note. It also provides a single point of lookup to find
	// all attached Attestation Occurrences, even if they don't all live in the same project.
	AttestationAuthority interface{}
	// The name of the note.
	Name interface{}
	Project interface{}
}
