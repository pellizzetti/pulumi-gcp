// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides access to available Google Compute regions for a given project.
// See more about [regions and regions](https://cloud.google.com/compute/docs/regions-zones/) in the upstream docs.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/d/compute_regions.html.markdown.
func LookupRegions(ctx *pulumi.Context, args *GetRegionsArgs) (*GetRegionsResult, error) {
	inputs := make(map[string]interface{})
	if args != nil {
		inputs["project"] = args.Project
		inputs["status"] = args.Status
	}
	outputs, err := ctx.Invoke("gcp:compute/getRegions:getRegions", inputs)
	if err != nil {
		return nil, err
	}
	return &GetRegionsResult{
		Names: outputs["names"],
		Project: outputs["project"],
		Status: outputs["status"],
		Id: outputs["id"],
	}, nil
}

// A collection of arguments for invoking getRegions.
type GetRegionsArgs struct {
	// Project from which to list available regions. Defaults to project declared in the provider.
	Project interface{}
	// Allows to filter list of regions based on their current status. Status can be either `UP` or `DOWN`.
	// Defaults to no filtering (all available regions - both `UP` and `DOWN`).
	Status interface{}
}

// A collection of values returned by getRegions.
type GetRegionsResult struct {
	// A list of regions available in the given project
	Names interface{}
	Project interface{}
	Status interface{}
	// id is the provider-assigned unique ID for this managed resource.
	Id interface{}
}
