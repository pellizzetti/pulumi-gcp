// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/d/compute_resource_policy.html.markdown.
func LookupResourcePolicy(ctx *pulumi.Context, args *GetResourcePolicyArgs) (*GetResourcePolicyResult, error) {
	inputs := make(map[string]interface{})
	if args != nil {
		inputs["name"] = args.Name
		inputs["project"] = args.Project
		inputs["region"] = args.Region
	}
	outputs, err := ctx.Invoke("gcp:compute/getResourcePolicy:getResourcePolicy", inputs)
	if err != nil {
		return nil, err
	}
	return &GetResourcePolicyResult{
		Description: outputs["description"],
		Name: outputs["name"],
		Project: outputs["project"],
		Region: outputs["region"],
		SelfLink: outputs["selfLink"],
		Id: outputs["id"],
	}, nil
}

// A collection of arguments for invoking getResourcePolicy.
type GetResourcePolicyArgs struct {
	// The name of the Resource Policy.
	Name interface{}
	// Project from which to list the Resource Policy. Defaults to project declared in the provider.
	Project interface{}
	// Region where the Resource Policy resides.
	Region interface{}
}

// A collection of values returned by getResourcePolicy.
type GetResourcePolicyResult struct {
	// Description of this Resource Policy.
	Description interface{}
	Name interface{}
	Project interface{}
	Region interface{}
	// The URI of the resource.
	SelfLink interface{}
	// id is the provider-assigned unique ID for this managed resource.
	Id interface{}
}
