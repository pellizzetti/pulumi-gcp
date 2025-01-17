// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package composer

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// An environment for running orchestration tasks.
// 
// Environments run Apache Airflow software on Google infrastructure.
// 
// To get more information about Environments, see:
// 
// * [API documentation](https://cloud.google.com/composer/docs/reference/rest/)
// * How-to Guides
//     * [Official Documentation](https://cloud.google.com/composer/docs)
//     * [Configuring Shared VPC for Composer Environments](https://cloud.google.com/composer/docs/how-to/managing/configuring-shared-vpc)
// * [Apache Airflow Documentation](http://airflow.apache.org/)
// 
// > **Warning:** We **STRONGLY** recommend  you read the [GCP guides](https://cloud.google.com/composer/docs/how-to)
//   as the Environment resource requires a long deployment process and involves several layers of GCP infrastructure, 
//   including a Kubernetes Engine cluster, Cloud Storage, and Compute networking resources. Due to limitations of the API,
//   this provider will not be able to automatically find or manage many of these underlying resources. In particular:
//   * It can take up to one hour to create or update an environment resource. In addition, GCP may only detect some 
//     errors in configuration when they are used (e.g. ~40-50 minutes into the creation process), and is prone to limited
//     error reporting. If you encounter confusing or uninformative errors, please verify your configuration is valid 
//     against GCP Cloud Composer before filing bugs against this provider. 
//   * **Environments create Google Cloud Storage buckets that do not get cleaned up automatically** on environment 
//     deletion. [More about Composer's use of Cloud Storage](https://cloud.google.com/composer/docs/concepts/cloud-storage).
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/composer_environment.html.markdown.
type Environment struct {
	s *pulumi.ResourceState
}

// NewEnvironment registers a new resource with the given unique name, arguments, and options.
func NewEnvironment(ctx *pulumi.Context,
	name string, args *EnvironmentArgs, opts ...pulumi.ResourceOpt) (*Environment, error) {
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["config"] = nil
		inputs["labels"] = nil
		inputs["name"] = nil
		inputs["project"] = nil
		inputs["region"] = nil
	} else {
		inputs["config"] = args.Config
		inputs["labels"] = args.Labels
		inputs["name"] = args.Name
		inputs["project"] = args.Project
		inputs["region"] = args.Region
	}
	s, err := ctx.RegisterResource("gcp:composer/environment:Environment", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Environment{s: s}, nil
}

// GetEnvironment gets an existing Environment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEnvironment(ctx *pulumi.Context,
	name string, id pulumi.ID, state *EnvironmentState, opts ...pulumi.ResourceOpt) (*Environment, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["config"] = state.Config
		inputs["labels"] = state.Labels
		inputs["name"] = state.Name
		inputs["project"] = state.Project
		inputs["region"] = state.Region
	}
	s, err := ctx.ReadResource("gcp:composer/environment:Environment", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Environment{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *Environment) URN() pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *Environment) ID() pulumi.IDOutput {
	return r.s.ID()
}

func (r *Environment) Config() pulumi.Output {
	return r.s.State["config"]
}

func (r *Environment) Labels() pulumi.MapOutput {
	return (pulumi.MapOutput)(r.s.State["labels"])
}

func (r *Environment) Name() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["name"])
}

// The ID of the project in which the resource belongs.
// If it is not provided, the provider project is used.
func (r *Environment) Project() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["project"])
}

func (r *Environment) Region() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["region"])
}

// Input properties used for looking up and filtering Environment resources.
type EnvironmentState struct {
	Config interface{}
	Labels interface{}
	Name interface{}
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project interface{}
	Region interface{}
}

// The set of arguments for constructing a Environment resource.
type EnvironmentArgs struct {
	Config interface{}
	Labels interface{}
	Name interface{}
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project interface{}
	Region interface{}
}
