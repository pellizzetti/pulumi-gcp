// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Get a Compute Instance Group within GCE.
// For more information, see [the official documentation](https://cloud.google.com/compute/docs/instance-groups/#unmanaged_instance_groups)
// and [API](https://cloud.google.com/compute/docs/reference/latest/instanceGroups)
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/d/compute_instance_group.html.markdown.
func LookupInstanceGroup(ctx *pulumi.Context, args *GetInstanceGroupArgs) (*GetInstanceGroupResult, error) {
	inputs := make(map[string]interface{})
	if args != nil {
		inputs["name"] = args.Name
		inputs["project"] = args.Project
		inputs["selfLink"] = args.SelfLink
		inputs["zone"] = args.Zone
	}
	outputs, err := ctx.Invoke("gcp:compute/getInstanceGroup:getInstanceGroup", inputs)
	if err != nil {
		return nil, err
	}
	return &GetInstanceGroupResult{
		Description: outputs["description"],
		Instances: outputs["instances"],
		Name: outputs["name"],
		NamedPorts: outputs["namedPorts"],
		Network: outputs["network"],
		Project: outputs["project"],
		SelfLink: outputs["selfLink"],
		Size: outputs["size"],
		Zone: outputs["zone"],
		Id: outputs["id"],
	}, nil
}

// A collection of arguments for invoking getInstanceGroup.
type GetInstanceGroupArgs struct {
	// The name of the instance group. Either `name` or `selfLink` must be provided.
	Name interface{}
	// The ID of the project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project interface{}
	// The self link of the instance group. Either `name` or `selfLink` must be provided.
	SelfLink interface{}
	// The zone of the instance group. If referencing the instance group by name
	// and `zone` is not provided, the provider zone is used.
	Zone interface{}
}

// A collection of values returned by getInstanceGroup.
type GetInstanceGroupResult struct {
	// Textual description of the instance group.
	Description interface{}
	// List of instances in the group.
	Instances interface{}
	Name interface{}
	// List of named ports in the group.
	NamedPorts interface{}
	// The URL of the network the instance group is in.
	Network interface{}
	Project interface{}
	// The URI of the resource.
	SelfLink interface{}
	// The number of instances in the group.
	Size interface{}
	Zone interface{}
	// id is the provider-assigned unique ID for this managed resource.
	Id interface{}
}
