// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/compute_backend_service.html.markdown.
type BackendService struct {
	s *pulumi.ResourceState
}

// NewBackendService registers a new resource with the given unique name, arguments, and options.
func NewBackendService(ctx *pulumi.Context,
	name string, args *BackendServiceArgs, opts ...pulumi.ResourceOpt) (*BackendService, error) {
	if args == nil || args.HealthChecks == nil {
		return nil, errors.New("missing required argument 'HealthChecks'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["affinityCookieTtlSec"] = nil
		inputs["backends"] = nil
		inputs["cdnPolicy"] = nil
		inputs["circuitBreakers"] = nil
		inputs["connectionDrainingTimeoutSec"] = nil
		inputs["consistentHash"] = nil
		inputs["customRequestHeaders"] = nil
		inputs["description"] = nil
		inputs["enableCdn"] = nil
		inputs["healthChecks"] = nil
		inputs["iap"] = nil
		inputs["loadBalancingScheme"] = nil
		inputs["localityLbPolicy"] = nil
		inputs["logConfig"] = nil
		inputs["name"] = nil
		inputs["outlierDetection"] = nil
		inputs["portName"] = nil
		inputs["project"] = nil
		inputs["protocol"] = nil
		inputs["securityPolicy"] = nil
		inputs["sessionAffinity"] = nil
		inputs["timeoutSec"] = nil
	} else {
		inputs["affinityCookieTtlSec"] = args.AffinityCookieTtlSec
		inputs["backends"] = args.Backends
		inputs["cdnPolicy"] = args.CdnPolicy
		inputs["circuitBreakers"] = args.CircuitBreakers
		inputs["connectionDrainingTimeoutSec"] = args.ConnectionDrainingTimeoutSec
		inputs["consistentHash"] = args.ConsistentHash
		inputs["customRequestHeaders"] = args.CustomRequestHeaders
		inputs["description"] = args.Description
		inputs["enableCdn"] = args.EnableCdn
		inputs["healthChecks"] = args.HealthChecks
		inputs["iap"] = args.Iap
		inputs["loadBalancingScheme"] = args.LoadBalancingScheme
		inputs["localityLbPolicy"] = args.LocalityLbPolicy
		inputs["logConfig"] = args.LogConfig
		inputs["name"] = args.Name
		inputs["outlierDetection"] = args.OutlierDetection
		inputs["portName"] = args.PortName
		inputs["project"] = args.Project
		inputs["protocol"] = args.Protocol
		inputs["securityPolicy"] = args.SecurityPolicy
		inputs["sessionAffinity"] = args.SessionAffinity
		inputs["timeoutSec"] = args.TimeoutSec
	}
	inputs["creationTimestamp"] = nil
	inputs["fingerprint"] = nil
	inputs["selfLink"] = nil
	s, err := ctx.RegisterResource("gcp:compute/backendService:BackendService", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &BackendService{s: s}, nil
}

// GetBackendService gets an existing BackendService resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBackendService(ctx *pulumi.Context,
	name string, id pulumi.ID, state *BackendServiceState, opts ...pulumi.ResourceOpt) (*BackendService, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["affinityCookieTtlSec"] = state.AffinityCookieTtlSec
		inputs["backends"] = state.Backends
		inputs["cdnPolicy"] = state.CdnPolicy
		inputs["circuitBreakers"] = state.CircuitBreakers
		inputs["connectionDrainingTimeoutSec"] = state.ConnectionDrainingTimeoutSec
		inputs["consistentHash"] = state.ConsistentHash
		inputs["creationTimestamp"] = state.CreationTimestamp
		inputs["customRequestHeaders"] = state.CustomRequestHeaders
		inputs["description"] = state.Description
		inputs["enableCdn"] = state.EnableCdn
		inputs["fingerprint"] = state.Fingerprint
		inputs["healthChecks"] = state.HealthChecks
		inputs["iap"] = state.Iap
		inputs["loadBalancingScheme"] = state.LoadBalancingScheme
		inputs["localityLbPolicy"] = state.LocalityLbPolicy
		inputs["logConfig"] = state.LogConfig
		inputs["name"] = state.Name
		inputs["outlierDetection"] = state.OutlierDetection
		inputs["portName"] = state.PortName
		inputs["project"] = state.Project
		inputs["protocol"] = state.Protocol
		inputs["securityPolicy"] = state.SecurityPolicy
		inputs["selfLink"] = state.SelfLink
		inputs["sessionAffinity"] = state.SessionAffinity
		inputs["timeoutSec"] = state.TimeoutSec
	}
	s, err := ctx.ReadResource("gcp:compute/backendService:BackendService", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &BackendService{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *BackendService) URN() pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *BackendService) ID() pulumi.IDOutput {
	return r.s.ID()
}

// Lifetime of cookies in seconds if session_affinity is GENERATED_COOKIE. If set to 0, the cookie is non-persistent and
// lasts only until the end of the browser session (or equivalent). The maximum allowed value for TTL is one day. When the
// load balancing scheme is INTERNAL, this field is not used.
func (r *BackendService) AffinityCookieTtlSec() pulumi.IntOutput {
	return (pulumi.IntOutput)(r.s.State["affinityCookieTtlSec"])
}

// The set of backends that serve this BackendService.
func (r *BackendService) Backends() pulumi.ArrayOutput {
	return (pulumi.ArrayOutput)(r.s.State["backends"])
}

// Cloud CDN configuration for this BackendService.
func (r *BackendService) CdnPolicy() pulumi.Output {
	return r.s.State["cdnPolicy"]
}

// Settings controlling the volume of connections to a backend service. This field is applicable only when the
// load_balancing_scheme is set to INTERNAL_SELF_MANAGED.
func (r *BackendService) CircuitBreakers() pulumi.Output {
	return r.s.State["circuitBreakers"]
}

// Time for which instance will be drained (not accept new connections, but still work to finish started).
func (r *BackendService) ConnectionDrainingTimeoutSec() pulumi.IntOutput {
	return (pulumi.IntOutput)(r.s.State["connectionDrainingTimeoutSec"])
}

// Consistent Hash-based load balancing can be used to provide soft session affinity based on HTTP headers, cookies or
// other properties. This load balancing policy is applicable only for HTTP connections. The affinity to a particular
// destination host will be lost when one or more hosts are added/removed from the destination service. This field
// specifies parameters that control consistent hashing. This field only applies if the load_balancing_scheme is set to
// INTERNAL_SELF_MANAGED. This field is only applicable when locality_lb_policy is set to MAGLEV or RING_HASH.
func (r *BackendService) ConsistentHash() pulumi.Output {
	return r.s.State["consistentHash"]
}

// Creation timestamp in RFC3339 text format.
func (r *BackendService) CreationTimestamp() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["creationTimestamp"])
}

// Headers that the HTTP/S load balancer should add to proxied requests.
func (r *BackendService) CustomRequestHeaders() pulumi.ArrayOutput {
	return (pulumi.ArrayOutput)(r.s.State["customRequestHeaders"])
}

// An optional description of this resource.
func (r *BackendService) Description() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["description"])
}

// If true, enable Cloud CDN for this BackendService.
func (r *BackendService) EnableCdn() pulumi.BoolOutput {
	return (pulumi.BoolOutput)(r.s.State["enableCdn"])
}

// Fingerprint of this resource. A hash of the contents stored in this object. This field is used in optimistic locking.
func (r *BackendService) Fingerprint() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["fingerprint"])
}

// The set of URLs to the HttpHealthCheck or HttpsHealthCheck resource for health checking this BackendService. Currently
// at most one health check can be specified, and a health check is required. For internal load balancing, a URL to a
// HealthCheck resource must be specified instead.
func (r *BackendService) HealthChecks() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["healthChecks"])
}

// Settings for enabling Cloud Identity Aware Proxy
func (r *BackendService) Iap() pulumi.Output {
	return r.s.State["iap"]
}

// Indicates whether the backend service will be used with internal or external load balancing. A backend service created
// for one type of load balancing cannot be used with the other. Must be 'EXTERNAL' or 'INTERNAL_SELF_MANAGED' for a global
// backend service. Defaults to 'EXTERNAL'.
func (r *BackendService) LoadBalancingScheme() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["loadBalancingScheme"])
}

// The load balancing algorithm used within the scope of the locality. The possible values are - ROUND_ROBIN - This is a
// simple policy in which each healthy backend is selected in round robin order. LEAST_REQUEST - An O(1) algorithm which
// selects two random healthy hosts and picks the host which has fewer active requests. RING_HASH - The ring/modulo hash
// load balancer implements consistent hashing to backends. The algorithm has the property that the addition/removal of a
// host from a set of N hosts only affects 1/N of the requests. RANDOM - The load balancer selects a random healthy host.
// ORIGINAL_DESTINATION - Backend host is selected based on the client connection metadata, i.e., connections are opened to
// the same address as the destination address of the incoming connection before the connection was redirected to the load
// balancer. MAGLEV - used as a drop in replacement for the ring hash load balancer. Maglev is not as stable as ring hash
// but has faster table lookup build times and host selection times. For more information about Maglev, refer to
// https://ai.google/research/pubs/pub44824 This field is applicable only when the load_balancing_scheme is set to
// INTERNAL_SELF_MANAGED.
func (r *BackendService) LocalityLbPolicy() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["localityLbPolicy"])
}

// This field denotes the logging options for the load balancer traffic served by this backend service. If logging is
// enabled, logs will be exported to Stackdriver.
func (r *BackendService) LogConfig() pulumi.Output {
	return r.s.State["logConfig"]
}

// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and
// comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression
// '[a-z]([-a-z0-9]*[a-z0-9])?' which means the first character must be a lowercase letter, and all following characters
// must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
func (r *BackendService) Name() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["name"])
}

// Settings controlling eviction of unhealthy hosts from the load balancing pool. This field is applicable only when the
// load_balancing_scheme is set to INTERNAL_SELF_MANAGED.
func (r *BackendService) OutlierDetection() pulumi.Output {
	return r.s.State["outlierDetection"]
}

// Name of backend port. The same name should appear in the instance groups referenced by this service. Required when the
// load balancing scheme is EXTERNAL.
func (r *BackendService) PortName() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["portName"])
}

// The ID of the project in which the resource belongs.
// If it is not provided, the provider project is used.
func (r *BackendService) Project() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["project"])
}

// The protocol this BackendService uses to communicate with backends. Possible values are HTTP, HTTPS, HTTP2, TCP, and
// SSL. The default is HTTP. **NOTE**: HTTP2 is only valid for beta HTTP/2 load balancer types and may result in errors if
// used with the GA API.
func (r *BackendService) Protocol() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["protocol"])
}

// The security policy associated with this backend service.
func (r *BackendService) SecurityPolicy() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["securityPolicy"])
}

// The URI of the created resource.
func (r *BackendService) SelfLink() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["selfLink"])
}

// Type of session affinity to use. The default is NONE. Session affinity is not applicable if the protocol is UDP.
func (r *BackendService) SessionAffinity() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["sessionAffinity"])
}

// How many seconds to wait for the backend before considering it a failed request. Default is 30 seconds. Valid range is
// [1, 86400].
func (r *BackendService) TimeoutSec() pulumi.IntOutput {
	return (pulumi.IntOutput)(r.s.State["timeoutSec"])
}

// Input properties used for looking up and filtering BackendService resources.
type BackendServiceState struct {
	// Lifetime of cookies in seconds if session_affinity is GENERATED_COOKIE. If set to 0, the cookie is non-persistent and
	// lasts only until the end of the browser session (or equivalent). The maximum allowed value for TTL is one day. When the
	// load balancing scheme is INTERNAL, this field is not used.
	AffinityCookieTtlSec interface{}
	// The set of backends that serve this BackendService.
	Backends interface{}
	// Cloud CDN configuration for this BackendService.
	CdnPolicy interface{}
	// Settings controlling the volume of connections to a backend service. This field is applicable only when the
	// load_balancing_scheme is set to INTERNAL_SELF_MANAGED.
	CircuitBreakers interface{}
	// Time for which instance will be drained (not accept new connections, but still work to finish started).
	ConnectionDrainingTimeoutSec interface{}
	// Consistent Hash-based load balancing can be used to provide soft session affinity based on HTTP headers, cookies or
	// other properties. This load balancing policy is applicable only for HTTP connections. The affinity to a particular
	// destination host will be lost when one or more hosts are added/removed from the destination service. This field
	// specifies parameters that control consistent hashing. This field only applies if the load_balancing_scheme is set to
	// INTERNAL_SELF_MANAGED. This field is only applicable when locality_lb_policy is set to MAGLEV or RING_HASH.
	ConsistentHash interface{}
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp interface{}
	// Headers that the HTTP/S load balancer should add to proxied requests.
	CustomRequestHeaders interface{}
	// An optional description of this resource.
	Description interface{}
	// If true, enable Cloud CDN for this BackendService.
	EnableCdn interface{}
	// Fingerprint of this resource. A hash of the contents stored in this object. This field is used in optimistic locking.
	Fingerprint interface{}
	// The set of URLs to the HttpHealthCheck or HttpsHealthCheck resource for health checking this BackendService. Currently
	// at most one health check can be specified, and a health check is required. For internal load balancing, a URL to a
	// HealthCheck resource must be specified instead.
	HealthChecks interface{}
	// Settings for enabling Cloud Identity Aware Proxy
	Iap interface{}
	// Indicates whether the backend service will be used with internal or external load balancing. A backend service created
	// for one type of load balancing cannot be used with the other. Must be 'EXTERNAL' or 'INTERNAL_SELF_MANAGED' for a
	// global backend service. Defaults to 'EXTERNAL'.
	LoadBalancingScheme interface{}
	// The load balancing algorithm used within the scope of the locality. The possible values are - ROUND_ROBIN - This is a
	// simple policy in which each healthy backend is selected in round robin order. LEAST_REQUEST - An O(1) algorithm which
	// selects two random healthy hosts and picks the host which has fewer active requests. RING_HASH - The ring/modulo hash
	// load balancer implements consistent hashing to backends. The algorithm has the property that the addition/removal of a
	// host from a set of N hosts only affects 1/N of the requests. RANDOM - The load balancer selects a random healthy host.
	// ORIGINAL_DESTINATION - Backend host is selected based on the client connection metadata, i.e., connections are opened
	// to the same address as the destination address of the incoming connection before the connection was redirected to the
	// load balancer. MAGLEV - used as a drop in replacement for the ring hash load balancer. Maglev is not as stable as ring
	// hash but has faster table lookup build times and host selection times. For more information about Maglev, refer to
	// https://ai.google/research/pubs/pub44824 This field is applicable only when the load_balancing_scheme is set to
	// INTERNAL_SELF_MANAGED.
	LocalityLbPolicy interface{}
	// This field denotes the logging options for the load balancer traffic served by this backend service. If logging is
	// enabled, logs will be exported to Stackdriver.
	LogConfig interface{}
	// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and
	// comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression
	// '[a-z]([-a-z0-9]*[a-z0-9])?' which means the first character must be a lowercase letter, and all following characters
	// must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name interface{}
	// Settings controlling eviction of unhealthy hosts from the load balancing pool. This field is applicable only when the
	// load_balancing_scheme is set to INTERNAL_SELF_MANAGED.
	OutlierDetection interface{}
	// Name of backend port. The same name should appear in the instance groups referenced by this service. Required when the
	// load balancing scheme is EXTERNAL.
	PortName interface{}
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project interface{}
	// The protocol this BackendService uses to communicate with backends. Possible values are HTTP, HTTPS, HTTP2, TCP, and
	// SSL. The default is HTTP. **NOTE**: HTTP2 is only valid for beta HTTP/2 load balancer types and may result in errors if
	// used with the GA API.
	Protocol interface{}
	// The security policy associated with this backend service.
	SecurityPolicy interface{}
	// The URI of the created resource.
	SelfLink interface{}
	// Type of session affinity to use. The default is NONE. Session affinity is not applicable if the protocol is UDP.
	SessionAffinity interface{}
	// How many seconds to wait for the backend before considering it a failed request. Default is 30 seconds. Valid range is
	// [1, 86400].
	TimeoutSec interface{}
}

// The set of arguments for constructing a BackendService resource.
type BackendServiceArgs struct {
	// Lifetime of cookies in seconds if session_affinity is GENERATED_COOKIE. If set to 0, the cookie is non-persistent and
	// lasts only until the end of the browser session (or equivalent). The maximum allowed value for TTL is one day. When the
	// load balancing scheme is INTERNAL, this field is not used.
	AffinityCookieTtlSec interface{}
	// The set of backends that serve this BackendService.
	Backends interface{}
	// Cloud CDN configuration for this BackendService.
	CdnPolicy interface{}
	// Settings controlling the volume of connections to a backend service. This field is applicable only when the
	// load_balancing_scheme is set to INTERNAL_SELF_MANAGED.
	CircuitBreakers interface{}
	// Time for which instance will be drained (not accept new connections, but still work to finish started).
	ConnectionDrainingTimeoutSec interface{}
	// Consistent Hash-based load balancing can be used to provide soft session affinity based on HTTP headers, cookies or
	// other properties. This load balancing policy is applicable only for HTTP connections. The affinity to a particular
	// destination host will be lost when one or more hosts are added/removed from the destination service. This field
	// specifies parameters that control consistent hashing. This field only applies if the load_balancing_scheme is set to
	// INTERNAL_SELF_MANAGED. This field is only applicable when locality_lb_policy is set to MAGLEV or RING_HASH.
	ConsistentHash interface{}
	// Headers that the HTTP/S load balancer should add to proxied requests.
	CustomRequestHeaders interface{}
	// An optional description of this resource.
	Description interface{}
	// If true, enable Cloud CDN for this BackendService.
	EnableCdn interface{}
	// The set of URLs to the HttpHealthCheck or HttpsHealthCheck resource for health checking this BackendService. Currently
	// at most one health check can be specified, and a health check is required. For internal load balancing, a URL to a
	// HealthCheck resource must be specified instead.
	HealthChecks interface{}
	// Settings for enabling Cloud Identity Aware Proxy
	Iap interface{}
	// Indicates whether the backend service will be used with internal or external load balancing. A backend service created
	// for one type of load balancing cannot be used with the other. Must be 'EXTERNAL' or 'INTERNAL_SELF_MANAGED' for a
	// global backend service. Defaults to 'EXTERNAL'.
	LoadBalancingScheme interface{}
	// The load balancing algorithm used within the scope of the locality. The possible values are - ROUND_ROBIN - This is a
	// simple policy in which each healthy backend is selected in round robin order. LEAST_REQUEST - An O(1) algorithm which
	// selects two random healthy hosts and picks the host which has fewer active requests. RING_HASH - The ring/modulo hash
	// load balancer implements consistent hashing to backends. The algorithm has the property that the addition/removal of a
	// host from a set of N hosts only affects 1/N of the requests. RANDOM - The load balancer selects a random healthy host.
	// ORIGINAL_DESTINATION - Backend host is selected based on the client connection metadata, i.e., connections are opened
	// to the same address as the destination address of the incoming connection before the connection was redirected to the
	// load balancer. MAGLEV - used as a drop in replacement for the ring hash load balancer. Maglev is not as stable as ring
	// hash but has faster table lookup build times and host selection times. For more information about Maglev, refer to
	// https://ai.google/research/pubs/pub44824 This field is applicable only when the load_balancing_scheme is set to
	// INTERNAL_SELF_MANAGED.
	LocalityLbPolicy interface{}
	// This field denotes the logging options for the load balancer traffic served by this backend service. If logging is
	// enabled, logs will be exported to Stackdriver.
	LogConfig interface{}
	// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and
	// comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression
	// '[a-z]([-a-z0-9]*[a-z0-9])?' which means the first character must be a lowercase letter, and all following characters
	// must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name interface{}
	// Settings controlling eviction of unhealthy hosts from the load balancing pool. This field is applicable only when the
	// load_balancing_scheme is set to INTERNAL_SELF_MANAGED.
	OutlierDetection interface{}
	// Name of backend port. The same name should appear in the instance groups referenced by this service. Required when the
	// load balancing scheme is EXTERNAL.
	PortName interface{}
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project interface{}
	// The protocol this BackendService uses to communicate with backends. Possible values are HTTP, HTTPS, HTTP2, TCP, and
	// SSL. The default is HTTP. **NOTE**: HTTP2 is only valid for beta HTTP/2 load balancer types and may result in errors if
	// used with the GA API.
	Protocol interface{}
	// The security policy associated with this backend service.
	SecurityPolicy interface{}
	// Type of session affinity to use. The default is NONE. Session affinity is not applicable if the protocol is UDP.
	SessionAffinity interface{}
	// How many seconds to wait for the backend before considering it a failed request. Default is 30 seconds. Valid range is
	// [1, 86400].
	TimeoutSec interface{}
}
