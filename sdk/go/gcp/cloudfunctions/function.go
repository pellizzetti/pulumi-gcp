// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package cloudfunctions

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Creates a new Cloud Function. For more information see
// [the official documentation](https://cloud.google.com/functions/docs/)
// and
// [API](https://cloud.google.com/functions/docs/apis).
// 
// > **Warning:** As of November 1, 2019, newly created Functions are
// private-by-default and will require [appropriate IAM permissions](https://cloud.google.com/functions/docs/reference/iam/roles)
// to be invoked. See below examples for how to set up the appropriate permissions,
// or view the [Cloud Functions IAM resources](https://www.terraform.io/docs/providers/google/r/cloudfunctions_cloud_function_iam.html)
// for Cloud Functions.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/cloudfunctions_function.html.markdown.
type Function struct {
	s *pulumi.ResourceState
}

// NewFunction registers a new resource with the given unique name, arguments, and options.
func NewFunction(ctx *pulumi.Context,
	name string, args *FunctionArgs, opts ...pulumi.ResourceOpt) (*Function, error) {
	if args == nil || args.Runtime == nil {
		return nil, errors.New("missing required argument 'Runtime'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["availableMemoryMb"] = nil
		inputs["description"] = nil
		inputs["entryPoint"] = nil
		inputs["environmentVariables"] = nil
		inputs["eventTrigger"] = nil
		inputs["httpsTriggerUrl"] = nil
		inputs["labels"] = nil
		inputs["maxInstances"] = nil
		inputs["name"] = nil
		inputs["project"] = nil
		inputs["region"] = nil
		inputs["runtime"] = nil
		inputs["serviceAccountEmail"] = nil
		inputs["sourceArchiveBucket"] = nil
		inputs["sourceArchiveObject"] = nil
		inputs["sourceRepository"] = nil
		inputs["timeout"] = nil
		inputs["triggerHttp"] = nil
		inputs["vpcConnector"] = nil
	} else {
		inputs["availableMemoryMb"] = args.AvailableMemoryMb
		inputs["description"] = args.Description
		inputs["entryPoint"] = args.EntryPoint
		inputs["environmentVariables"] = args.EnvironmentVariables
		inputs["eventTrigger"] = args.EventTrigger
		inputs["httpsTriggerUrl"] = args.HttpsTriggerUrl
		inputs["labels"] = args.Labels
		inputs["maxInstances"] = args.MaxInstances
		inputs["name"] = args.Name
		inputs["project"] = args.Project
		inputs["region"] = args.Region
		inputs["runtime"] = args.Runtime
		inputs["serviceAccountEmail"] = args.ServiceAccountEmail
		inputs["sourceArchiveBucket"] = args.SourceArchiveBucket
		inputs["sourceArchiveObject"] = args.SourceArchiveObject
		inputs["sourceRepository"] = args.SourceRepository
		inputs["timeout"] = args.Timeout
		inputs["triggerHttp"] = args.TriggerHttp
		inputs["vpcConnector"] = args.VpcConnector
	}
	s, err := ctx.RegisterResource("gcp:cloudfunctions/function:Function", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Function{s: s}, nil
}

// GetFunction gets an existing Function resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFunction(ctx *pulumi.Context,
	name string, id pulumi.ID, state *FunctionState, opts ...pulumi.ResourceOpt) (*Function, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["availableMemoryMb"] = state.AvailableMemoryMb
		inputs["description"] = state.Description
		inputs["entryPoint"] = state.EntryPoint
		inputs["environmentVariables"] = state.EnvironmentVariables
		inputs["eventTrigger"] = state.EventTrigger
		inputs["httpsTriggerUrl"] = state.HttpsTriggerUrl
		inputs["labels"] = state.Labels
		inputs["maxInstances"] = state.MaxInstances
		inputs["name"] = state.Name
		inputs["project"] = state.Project
		inputs["region"] = state.Region
		inputs["runtime"] = state.Runtime
		inputs["serviceAccountEmail"] = state.ServiceAccountEmail
		inputs["sourceArchiveBucket"] = state.SourceArchiveBucket
		inputs["sourceArchiveObject"] = state.SourceArchiveObject
		inputs["sourceRepository"] = state.SourceRepository
		inputs["timeout"] = state.Timeout
		inputs["triggerHttp"] = state.TriggerHttp
		inputs["vpcConnector"] = state.VpcConnector
	}
	s, err := ctx.ReadResource("gcp:cloudfunctions/function:Function", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Function{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *Function) URN() pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *Function) ID() pulumi.IDOutput {
	return r.s.ID()
}

// Memory (in MB), available to the function. Default value is 256MB. Allowed values are: 128MB, 256MB, 512MB, 1024MB, and 2048MB.
func (r *Function) AvailableMemoryMb() pulumi.IntOutput {
	return (pulumi.IntOutput)(r.s.State["availableMemoryMb"])
}

// Description of the function.
func (r *Function) Description() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["description"])
}

// Name of the function that will be executed when the Google Cloud Function is triggered.
func (r *Function) EntryPoint() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["entryPoint"])
}

// A set of key/value environment variable pairs to assign to the function.
func (r *Function) EnvironmentVariables() pulumi.MapOutput {
	return (pulumi.MapOutput)(r.s.State["environmentVariables"])
}

// A source that fires events in response to a condition in another service. Structure is documented below. Cannot be used with `triggerHttp`.
func (r *Function) EventTrigger() pulumi.Output {
	return r.s.State["eventTrigger"]
}

// URL which triggers function execution. Returned only if `triggerHttp` is used.
func (r *Function) HttpsTriggerUrl() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["httpsTriggerUrl"])
}

// A set of key/value label pairs to assign to the function.
func (r *Function) Labels() pulumi.MapOutput {
	return (pulumi.MapOutput)(r.s.State["labels"])
}

// The limit on the maximum number of function instances that may coexist at a given time.
func (r *Function) MaxInstances() pulumi.IntOutput {
	return (pulumi.IntOutput)(r.s.State["maxInstances"])
}

// A user-defined name of the function. Function names must be unique globally.
func (r *Function) Name() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["name"])
}

// Project of the function. If it is not provided, the provider project is used.
func (r *Function) Project() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["project"])
}

// Region of function. Currently can be only "us-central1". If it is not provided, the provider region is used.
func (r *Function) Region() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["region"])
}

// The runtime in which the function is going to run.
// Eg. `"nodejs8"`, `"nodejs10"`, `"python37"`, `"go111"`.
func (r *Function) Runtime() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["runtime"])
}

// If provided, the self-provided service account to run the function with.
func (r *Function) ServiceAccountEmail() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["serviceAccountEmail"])
}

// The GCS bucket containing the zip archive which contains the function.
func (r *Function) SourceArchiveBucket() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["sourceArchiveBucket"])
}

// The source archive object (file) in archive bucket.
func (r *Function) SourceArchiveObject() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["sourceArchiveObject"])
}

// Represents parameters related to source repository where a function is hosted.
// Cannot be set alongside `sourceArchiveBucket` or `sourceArchiveObject`. Structure is documented below.
func (r *Function) SourceRepository() pulumi.Output {
	return r.s.State["sourceRepository"]
}

// Timeout (in seconds) for the function. Default value is 60 seconds. Cannot be more than 540 seconds.
func (r *Function) Timeout() pulumi.IntOutput {
	return (pulumi.IntOutput)(r.s.State["timeout"])
}

// Boolean variable. Any HTTP request (of a supported type) to the endpoint will trigger function execution. Supported HTTP request types are: POST, PUT, GET, DELETE, and OPTIONS. Endpoint is returned as `httpsTriggerUrl`. Cannot be used with `triggerBucket` and `triggerTopic`.
func (r *Function) TriggerHttp() pulumi.BoolOutput {
	return (pulumi.BoolOutput)(r.s.State["triggerHttp"])
}

// The VPC Network Connector that this cloud function can connect to. It can be either the fully-qualified URI, or the short name of the network connector resource. The format of this field is `projects/*/locations/*/connectors/*`.
func (r *Function) VpcConnector() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["vpcConnector"])
}

// Input properties used for looking up and filtering Function resources.
type FunctionState struct {
	// Memory (in MB), available to the function. Default value is 256MB. Allowed values are: 128MB, 256MB, 512MB, 1024MB, and 2048MB.
	AvailableMemoryMb interface{}
	// Description of the function.
	Description interface{}
	// Name of the function that will be executed when the Google Cloud Function is triggered.
	EntryPoint interface{}
	// A set of key/value environment variable pairs to assign to the function.
	EnvironmentVariables interface{}
	// A source that fires events in response to a condition in another service. Structure is documented below. Cannot be used with `triggerHttp`.
	EventTrigger interface{}
	// URL which triggers function execution. Returned only if `triggerHttp` is used.
	HttpsTriggerUrl interface{}
	// A set of key/value label pairs to assign to the function.
	Labels interface{}
	// The limit on the maximum number of function instances that may coexist at a given time.
	MaxInstances interface{}
	// A user-defined name of the function. Function names must be unique globally.
	Name interface{}
	// Project of the function. If it is not provided, the provider project is used.
	Project interface{}
	// Region of function. Currently can be only "us-central1". If it is not provided, the provider region is used.
	Region interface{}
	// The runtime in which the function is going to run.
	// Eg. `"nodejs8"`, `"nodejs10"`, `"python37"`, `"go111"`.
	Runtime interface{}
	// If provided, the self-provided service account to run the function with.
	ServiceAccountEmail interface{}
	// The GCS bucket containing the zip archive which contains the function.
	SourceArchiveBucket interface{}
	// The source archive object (file) in archive bucket.
	SourceArchiveObject interface{}
	// Represents parameters related to source repository where a function is hosted.
	// Cannot be set alongside `sourceArchiveBucket` or `sourceArchiveObject`. Structure is documented below.
	SourceRepository interface{}
	// Timeout (in seconds) for the function. Default value is 60 seconds. Cannot be more than 540 seconds.
	Timeout interface{}
	// Boolean variable. Any HTTP request (of a supported type) to the endpoint will trigger function execution. Supported HTTP request types are: POST, PUT, GET, DELETE, and OPTIONS. Endpoint is returned as `httpsTriggerUrl`. Cannot be used with `triggerBucket` and `triggerTopic`.
	TriggerHttp interface{}
	// The VPC Network Connector that this cloud function can connect to. It can be either the fully-qualified URI, or the short name of the network connector resource. The format of this field is `projects/*/locations/*/connectors/*`.
	VpcConnector interface{}
}

// The set of arguments for constructing a Function resource.
type FunctionArgs struct {
	// Memory (in MB), available to the function. Default value is 256MB. Allowed values are: 128MB, 256MB, 512MB, 1024MB, and 2048MB.
	AvailableMemoryMb interface{}
	// Description of the function.
	Description interface{}
	// Name of the function that will be executed when the Google Cloud Function is triggered.
	EntryPoint interface{}
	// A set of key/value environment variable pairs to assign to the function.
	EnvironmentVariables interface{}
	// A source that fires events in response to a condition in another service. Structure is documented below. Cannot be used with `triggerHttp`.
	EventTrigger interface{}
	// URL which triggers function execution. Returned only if `triggerHttp` is used.
	HttpsTriggerUrl interface{}
	// A set of key/value label pairs to assign to the function.
	Labels interface{}
	// The limit on the maximum number of function instances that may coexist at a given time.
	MaxInstances interface{}
	// A user-defined name of the function. Function names must be unique globally.
	Name interface{}
	// Project of the function. If it is not provided, the provider project is used.
	Project interface{}
	// Region of function. Currently can be only "us-central1". If it is not provided, the provider region is used.
	Region interface{}
	// The runtime in which the function is going to run.
	// Eg. `"nodejs8"`, `"nodejs10"`, `"python37"`, `"go111"`.
	Runtime interface{}
	// If provided, the self-provided service account to run the function with.
	ServiceAccountEmail interface{}
	// The GCS bucket containing the zip archive which contains the function.
	SourceArchiveBucket interface{}
	// The source archive object (file) in archive bucket.
	SourceArchiveObject interface{}
	// Represents parameters related to source repository where a function is hosted.
	// Cannot be set alongside `sourceArchiveBucket` or `sourceArchiveObject`. Structure is documented below.
	SourceRepository interface{}
	// Timeout (in seconds) for the function. Default value is 60 seconds. Cannot be more than 540 seconds.
	Timeout interface{}
	// Boolean variable. Any HTTP request (of a supported type) to the endpoint will trigger function execution. Supported HTTP request types are: POST, PUT, GET, DELETE, and OPTIONS. Endpoint is returned as `httpsTriggerUrl`. Cannot be used with `triggerBucket` and `triggerTopic`.
	TriggerHttp interface{}
	// The VPC Network Connector that this cloud function can connect to. It can be either the fully-qualified URI, or the short name of the network connector resource. The format of this field is `projects/*/locations/*/connectors/*`.
	VpcConnector interface{}
}
