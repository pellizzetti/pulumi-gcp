// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package logging

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/logging_metric.html.markdown.
type Metric struct {
	s *pulumi.ResourceState
}

// NewMetric registers a new resource with the given unique name, arguments, and options.
func NewMetric(ctx *pulumi.Context,
	name string, args *MetricArgs, opts ...pulumi.ResourceOpt) (*Metric, error) {
	if args == nil || args.Filter == nil {
		return nil, errors.New("missing required argument 'Filter'")
	}
	if args == nil || args.MetricDescriptor == nil {
		return nil, errors.New("missing required argument 'MetricDescriptor'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["bucketOptions"] = nil
		inputs["description"] = nil
		inputs["filter"] = nil
		inputs["labelExtractors"] = nil
		inputs["metricDescriptor"] = nil
		inputs["name"] = nil
		inputs["project"] = nil
		inputs["valueExtractor"] = nil
	} else {
		inputs["bucketOptions"] = args.BucketOptions
		inputs["description"] = args.Description
		inputs["filter"] = args.Filter
		inputs["labelExtractors"] = args.LabelExtractors
		inputs["metricDescriptor"] = args.MetricDescriptor
		inputs["name"] = args.Name
		inputs["project"] = args.Project
		inputs["valueExtractor"] = args.ValueExtractor
	}
	s, err := ctx.RegisterResource("gcp:logging/metric:Metric", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Metric{s: s}, nil
}

// GetMetric gets an existing Metric resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMetric(ctx *pulumi.Context,
	name string, id pulumi.ID, state *MetricState, opts ...pulumi.ResourceOpt) (*Metric, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["bucketOptions"] = state.BucketOptions
		inputs["description"] = state.Description
		inputs["filter"] = state.Filter
		inputs["labelExtractors"] = state.LabelExtractors
		inputs["metricDescriptor"] = state.MetricDescriptor
		inputs["name"] = state.Name
		inputs["project"] = state.Project
		inputs["valueExtractor"] = state.ValueExtractor
	}
	s, err := ctx.ReadResource("gcp:logging/metric:Metric", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Metric{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *Metric) URN() pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *Metric) ID() pulumi.IDOutput {
	return r.s.ID()
}

// The bucketOptions are required when the logs-based metric is using a DISTRIBUTION value type and it describes the bucket
// boundaries used to create a histogram of the extracted values.
func (r *Metric) BucketOptions() pulumi.Output {
	return r.s.State["bucketOptions"]
}

// A description of this metric, which is used in documentation. The maximum length of the description is 8000 characters.
func (r *Metric) Description() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["description"])
}

// An advanced logs filter (https://cloud.google.com/logging/docs/view/advanced-filters) which is used to match log
// entries.
func (r *Metric) Filter() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["filter"])
}

// A map from a label key string to an extractor expression which is used to extract data from a log entry field and assign
// as the label value. Each label key specified in the LabelDescriptor must have an associated extractor expression in this
// map. The syntax of the extractor expression is the same as for the valueExtractor field.
func (r *Metric) LabelExtractors() pulumi.MapOutput {
	return (pulumi.MapOutput)(r.s.State["labelExtractors"])
}

// The metric descriptor associated with the logs-based metric.
func (r *Metric) MetricDescriptor() pulumi.Output {
	return r.s.State["metricDescriptor"]
}

// The client-assigned metric identifier. Examples - "error_count", "nginx/requests". Metric identifiers are limited to 100
// characters and can include only the following characters A-Z, a-z, 0-9, and the special characters _-.,+!*',()%!/(MISSING). The
// forward-slash character (/) denotes a hierarchy of name pieces, and it cannot be the first character of the name.
func (r *Metric) Name() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["name"])
}

func (r *Metric) Project() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["project"])
}

// A valueExtractor is required when using a distribution logs-based metric to extract the values to record from a log
// entry. Two functions are supported for value extraction - EXTRACT(field) or REGEXP_EXTRACT(field, regex). The argument
// are 1. field - The name of the log entry field from which the value is to be extracted. 2. regex - A regular expression
// using the Google RE2 syntax (https://github.com/google/re2/wiki/Syntax) with a single capture group to extract data from
// the specified log entry field. The value of the field is converted to a string before applying the regex. It is an error
// to specify a regex that does not include exactly one capture group.
func (r *Metric) ValueExtractor() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["valueExtractor"])
}

// Input properties used for looking up and filtering Metric resources.
type MetricState struct {
	// The bucketOptions are required when the logs-based metric is using a DISTRIBUTION value type and it describes the
	// bucket boundaries used to create a histogram of the extracted values.
	BucketOptions interface{}
	// A description of this metric, which is used in documentation. The maximum length of the description is 8000 characters.
	Description interface{}
	// An advanced logs filter (https://cloud.google.com/logging/docs/view/advanced-filters) which is used to match log
	// entries.
	Filter interface{}
	// A map from a label key string to an extractor expression which is used to extract data from a log entry field and
	// assign as the label value. Each label key specified in the LabelDescriptor must have an associated extractor expression
	// in this map. The syntax of the extractor expression is the same as for the valueExtractor field.
	LabelExtractors interface{}
	// The metric descriptor associated with the logs-based metric.
	MetricDescriptor interface{}
	// The client-assigned metric identifier. Examples - "error_count", "nginx/requests". Metric identifiers are limited to
	// 100 characters and can include only the following characters A-Z, a-z, 0-9, and the special characters _-.,+!*',()%!/(MISSING).
	// The forward-slash character (/) denotes a hierarchy of name pieces, and it cannot be the first character of the name.
	Name interface{}
	Project interface{}
	// A valueExtractor is required when using a distribution logs-based metric to extract the values to record from a log
	// entry. Two functions are supported for value extraction - EXTRACT(field) or REGEXP_EXTRACT(field, regex). The argument
	// are 1. field - The name of the log entry field from which the value is to be extracted. 2. regex - A regular expression
	// using the Google RE2 syntax (https://github.com/google/re2/wiki/Syntax) with a single capture group to extract data
	// from the specified log entry field. The value of the field is converted to a string before applying the regex. It is an
	// error to specify a regex that does not include exactly one capture group.
	ValueExtractor interface{}
}

// The set of arguments for constructing a Metric resource.
type MetricArgs struct {
	// The bucketOptions are required when the logs-based metric is using a DISTRIBUTION value type and it describes the
	// bucket boundaries used to create a histogram of the extracted values.
	BucketOptions interface{}
	// A description of this metric, which is used in documentation. The maximum length of the description is 8000 characters.
	Description interface{}
	// An advanced logs filter (https://cloud.google.com/logging/docs/view/advanced-filters) which is used to match log
	// entries.
	Filter interface{}
	// A map from a label key string to an extractor expression which is used to extract data from a log entry field and
	// assign as the label value. Each label key specified in the LabelDescriptor must have an associated extractor expression
	// in this map. The syntax of the extractor expression is the same as for the valueExtractor field.
	LabelExtractors interface{}
	// The metric descriptor associated with the logs-based metric.
	MetricDescriptor interface{}
	// The client-assigned metric identifier. Examples - "error_count", "nginx/requests". Metric identifiers are limited to
	// 100 characters and can include only the following characters A-Z, a-z, 0-9, and the special characters _-.,+!*',()%!/(MISSING).
	// The forward-slash character (/) denotes a hierarchy of name pieces, and it cannot be the first character of the name.
	Name interface{}
	Project interface{}
	// A valueExtractor is required when using a distribution logs-based metric to extract the values to record from a log
	// entry. Two functions are supported for value extraction - EXTRACT(field) or REGEXP_EXTRACT(field, regex). The argument
	// are 1. field - The name of the log entry field from which the value is to be extracted. 2. regex - A regular expression
	// using the Google RE2 syntax (https://github.com/google/re2/wiki/Syntax) with a single capture group to extract data
	// from the specified log entry field. The value of the field is converted to a string before applying the regex. It is an
	// error to specify a regex that does not include exactly one capture group.
	ValueExtractor interface{}
}
