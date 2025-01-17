// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Get the IP address from a static address reserved for a Global Forwarding Rule which are only used for HTTP load balancing. For more information see
// the official [API](https://cloud.google.com/compute/docs/reference/latest/globalAddresses) documentation.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/d/compute_global_address.html.markdown.
func LookupGlobalAddress(ctx *pulumi.Context, args *GetGlobalAddressArgs) (*GetGlobalAddressResult, error) {
	inputs := make(map[string]interface{})
	if args != nil {
		inputs["name"] = args.Name
		inputs["project"] = args.Project
	}
	outputs, err := ctx.Invoke("gcp:compute/getGlobalAddress:getGlobalAddress", inputs)
	if err != nil {
		return nil, err
	}
	return &GetGlobalAddressResult{
		Address: outputs["address"],
		Name: outputs["name"],
		Project: outputs["project"],
		SelfLink: outputs["selfLink"],
		Status: outputs["status"],
		Id: outputs["id"],
	}, nil
}

// A collection of arguments for invoking getGlobalAddress.
type GetGlobalAddressArgs struct {
	// A unique name for the resource, required by GCE.
	Name interface{}
	// The project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project interface{}
}

// A collection of values returned by getGlobalAddress.
type GetGlobalAddressResult struct {
	// The IP of the created resource.
	Address interface{}
	Name interface{}
	Project interface{}
	// The URI of the created resource.
	SelfLink interface{}
	// Indicates if the address is used. Possible values are: RESERVED or IN_USE.
	Status interface{}
	// id is the provider-assigned unique ID for this managed resource.
	Id interface{}
}
