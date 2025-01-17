// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package container

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// This data source fetches the project name, and provides the appropriate URLs to use for container registry for this project.
// 
// The URLs are computed entirely offline - as long as the project exists, they will be valid, but this data source does not contact Google Container Registry (GCR) at any point.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/d/container_registry_image.html.markdown.
func LookupRegistryImage(ctx *pulumi.Context, args *GetRegistryImageArgs) (*GetRegistryImageResult, error) {
	inputs := make(map[string]interface{})
	if args != nil {
		inputs["digest"] = args.Digest
		inputs["name"] = args.Name
		inputs["project"] = args.Project
		inputs["region"] = args.Region
		inputs["tag"] = args.Tag
	}
	outputs, err := ctx.Invoke("gcp:container/getRegistryImage:getRegistryImage", inputs)
	if err != nil {
		return nil, err
	}
	return &GetRegistryImageResult{
		Digest: outputs["digest"],
		ImageUrl: outputs["imageUrl"],
		Name: outputs["name"],
		Project: outputs["project"],
		Region: outputs["region"],
		Tag: outputs["tag"],
		Id: outputs["id"],
	}, nil
}

// A collection of arguments for invoking getRegistryImage.
type GetRegistryImageArgs struct {
	Digest interface{}
	Name interface{}
	Project interface{}
	Region interface{}
	Tag interface{}
}

// A collection of values returned by getRegistryImage.
type GetRegistryImageResult struct {
	Digest interface{}
	ImageUrl interface{}
	Name interface{}
	Project interface{}
	Region interface{}
	Tag interface{}
	// id is the provider-assigned unique ID for this managed resource.
	Id interface{}
}
