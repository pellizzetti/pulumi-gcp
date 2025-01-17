// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/compute_region_backend_service.html.markdown.
type RegionBackendService struct {
	s *pulumi.ResourceState
}

// NewRegionBackendService registers a new resource with the given unique name, arguments, and options.
func NewRegionBackendService(ctx *pulumi.Context,
	name string, args *RegionBackendServiceArgs, opts ...pulumi.ResourceOpt) (*RegionBackendService, error) {
	if args == nil || args.HealthChecks == nil {
		return nil, errors.New("missing required argument 'HealthChecks'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["affinityCookieTtlSec"] = nil
		inputs["backends"] = nil
		inputs["circuitBreakers"] = nil
		inputs["connectionDrainingTimeoutSec"] = nil
		inputs["consistentHash"] = nil
		inputs["description"] = nil
		inputs["failoverPolicy"] = nil
		inputs["healthChecks"] = nil
		inputs["loadBalancingScheme"] = nil
		inputs["localityLbPolicy"] = nil
		inputs["logConfig"] = nil
		inputs["name"] = nil
		inputs["network"] = nil
		inputs["outlierDetection"] = nil
		inputs["project"] = nil
		inputs["protocol"] = nil
		inputs["region"] = nil
		inputs["sessionAffinity"] = nil
		inputs["timeoutSec"] = nil
	} else {
		inputs["affinityCookieTtlSec"] = args.AffinityCookieTtlSec
		inputs["backends"] = args.Backends
		inputs["circuitBreakers"] = args.CircuitBreakers
		inputs["connectionDrainingTimeoutSec"] = args.ConnectionDrainingTimeoutSec
		inputs["consistentHash"] = args.ConsistentHash
		inputs["description"] = args.Description
		inputs["failoverPolicy"] = args.FailoverPolicy
		inputs["healthChecks"] = args.HealthChecks
		inputs["loadBalancingScheme"] = args.LoadBalancingScheme
		inputs["localityLbPolicy"] = args.LocalityLbPolicy
		inputs["logConfig"] = args.LogConfig
		inputs["name"] = args.Name
		inputs["network"] = args.Network
		inputs["outlierDetection"] = args.OutlierDetection
		inputs["project"] = args.Project
		inputs["protocol"] = args.Protocol
		inputs["region"] = args.Region
		inputs["sessionAffinity"] = args.SessionAffinity
		inputs["timeoutSec"] = args.TimeoutSec
	}
	inputs["creationTimestamp"] = nil
	inputs["fingerprint"] = nil
	inputs["selfLink"] = nil
	s, err := ctx.RegisterResource("gcp:compute/regionBackendService:RegionBackendService", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &RegionBackendService{s: s}, nil
}

// GetRegionBackendService gets an existing RegionBackendService resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRegionBackendService(ctx *pulumi.Context,
	name string, id pulumi.ID, state *RegionBackendServiceState, opts ...pulumi.ResourceOpt) (*RegionBackendService, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["affinityCookieTtlSec"] = state.AffinityCookieTtlSec
		inputs["backends"] = state.Backends
		inputs["circuitBreakers"] = state.CircuitBreakers
		inputs["connectionDrainingTimeoutSec"] = state.ConnectionDrainingTimeoutSec
		inputs["consistentHash"] = state.ConsistentHash
		inputs["creationTimestamp"] = state.CreationTimestamp
		inputs["description"] = state.Description
		inputs["failoverPolicy"] = state.FailoverPolicy
		inputs["fingerprint"] = state.Fingerprint
		inputs["healthChecks"] = state.HealthChecks
		inputs["loadBalancingScheme"] = state.LoadBalancingScheme
		inputs["localityLbPolicy"] = state.LocalityLbPolicy
		inputs["logConfig"] = state.LogConfig
		inputs["name"] = state.Name
		inputs["network"] = state.Network
		inputs["outlierDetection"] = state.OutlierDetection
		inputs["project"] = state.Project
		inputs["protocol"] = state.Protocol
		inputs["region"] = state.Region
		inputs["selfLink"] = state.SelfLink
		inputs["sessionAffinity"] = state.SessionAffinity
		inputs["timeoutSec"] = state.TimeoutSec
	}
	s, err := ctx.ReadResource("gcp:compute/regionBackendService:RegionBackendService", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &RegionBackendService{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *RegionBackendService) URN() pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *RegionBackendService) ID() pulumi.IDOutput {
	return r.s.ID()
}

// Lifetime of cookies in seconds if session_affinity is GENERATED_COOKIE. If set to 0, the cookie is non-persistent and
// lasts only until the end of the browser session (or equivalent). The maximum allowed value for TTL is one day. When the
// load balancing scheme is INTERNAL, this field is not used.
func (r *RegionBackendService) AffinityCookieTtlSec() pulumi.IntOutput {
	return (pulumi.IntOutput)(r.s.State["affinityCookieTtlSec"])
}

// The set of backends that serve this RegionBackendService.
func (r *RegionBackendService) Backends() pulumi.ArrayOutput {
	return (pulumi.ArrayOutput)(r.s.State["backends"])
}

// Settings controlling the volume of connections to a backend service. This field is applicable only when the
// 'load_balancing_scheme' is set to INTERNAL_MANAGED and the 'protocol' is set to HTTP, HTTPS, or HTTP2.
func (r *RegionBackendService) CircuitBreakers() pulumi.Output {
	return r.s.State["circuitBreakers"]
}

// Time for which instance will be drained (not accept new connections, but still work to finish started).
func (r *RegionBackendService) ConnectionDrainingTimeoutSec() pulumi.IntOutput {
	return (pulumi.IntOutput)(r.s.State["connectionDrainingTimeoutSec"])
}

// Consistent Hash-based load balancing can be used to provide soft session affinity based on HTTP headers, cookies or
// other properties. This load balancing policy is applicable only for HTTP connections. The affinity to a particular
// destination host will be lost when one or more hosts are added/removed from the destination service. This field
// specifies parameters that control consistent hashing. This field only applies when all of the following are true - *
// 'load_balancing_scheme' is set to INTERNAL_MANAGED * 'protocol' is set to HTTP, HTTPS, or HTTP2 * 'locality_lb_policy'
// is set to MAGLEV or RING_HASH
func (r *RegionBackendService) ConsistentHash() pulumi.Output {
	return r.s.State["consistentHash"]
}

// Creation timestamp in RFC3339 text format.
func (r *RegionBackendService) CreationTimestamp() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["creationTimestamp"])
}

// An optional description of this resource.
func (r *RegionBackendService) Description() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["description"])
}

// Policy for failovers.
func (r *RegionBackendService) FailoverPolicy() pulumi.Output {
	return r.s.State["failoverPolicy"]
}

// Fingerprint of this resource. A hash of the contents stored in this object. This field is used in optimistic locking.
func (r *RegionBackendService) Fingerprint() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["fingerprint"])
}

// The set of URLs to HealthCheck resources for health checking this RegionBackendService. Currently at most one health
// check can be specified, and a health check is required.
func (r *RegionBackendService) HealthChecks() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["healthChecks"])
}

// Indicates what kind of load balancing this regional backend service will be used for. A backend service created for one
// type of load balancing cannot be used with the other(s). Must be 'INTERNAL' or 'INTERNAL_MANAGED'. Defaults to
// 'INTERNAL'.
func (r *RegionBackendService) LoadBalancingScheme() pulumi.StringOutput {
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
// https://ai.google/research/pubs/pub44824 This field is applicable only when the 'load_balancing_scheme' is set to
// INTERNAL_MANAGED and the 'protocol' is set to HTTP, HTTPS, or HTTP2.
func (r *RegionBackendService) LocalityLbPolicy() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["localityLbPolicy"])
}

// This field denotes the logging options for the load balancer traffic served by this backend service. If logging is
// enabled, logs will be exported to Stackdriver.
func (r *RegionBackendService) LogConfig() pulumi.Output {
	return r.s.State["logConfig"]
}

// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and
// comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression
// '[a-z]([-a-z0-9]*[a-z0-9])?' which means the first character must be a lowercase letter, and all following characters
// must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
func (r *RegionBackendService) Name() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["name"])
}

// The URL of the network to which this backend service belongs. This field can only be specified when the load balancing
// scheme is set to INTERNAL.
func (r *RegionBackendService) Network() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["network"])
}

// Settings controlling eviction of unhealthy hosts from the load balancing pool. This field is applicable only when the
// 'load_balancing_scheme' is set to INTERNAL_MANAGED and the 'protocol' is set to HTTP, HTTPS, or HTTP2.
func (r *RegionBackendService) OutlierDetection() pulumi.Output {
	return r.s.State["outlierDetection"]
}

// The ID of the project in which the resource belongs.
// If it is not provided, the provider project is used.
func (r *RegionBackendService) Project() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["project"])
}

// The protocol this RegionBackendService uses to communicate with backends. Possible values are HTTP, HTTPS, HTTP2, SSL,
// TCP, and UDP. The default is HTTP. **NOTE**: HTTP2 is only valid for beta HTTP/2 load balancer types and may result in
// errors if used with the GA API.
func (r *RegionBackendService) Protocol() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["protocol"])
}

// The Region in which the created backend service should reside. If it is not provided, the provider region is used.
func (r *RegionBackendService) Region() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["region"])
}

// The URI of the created resource.
func (r *RegionBackendService) SelfLink() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["selfLink"])
}

// Type of session affinity to use. The default is NONE. Session affinity is not applicable if the protocol is UDP.
func (r *RegionBackendService) SessionAffinity() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["sessionAffinity"])
}

// How many seconds to wait for the backend before considering it a failed request. Default is 30 seconds. Valid range is
// [1, 86400].
func (r *RegionBackendService) TimeoutSec() pulumi.IntOutput {
	return (pulumi.IntOutput)(r.s.State["timeoutSec"])
}

// Input properties used for looking up and filtering RegionBackendService resources.
type RegionBackendServiceState struct {
	// Lifetime of cookies in seconds if session_affinity is GENERATED_COOKIE. If set to 0, the cookie is non-persistent and
	// lasts only until the end of the browser session (or equivalent). The maximum allowed value for TTL is one day. When the
	// load balancing scheme is INTERNAL, this field is not used.
	AffinityCookieTtlSec interface{}
	// The set of backends that serve this RegionBackendService.
	Backends interface{}
	// Settings controlling the volume of connections to a backend service. This field is applicable only when the
	// 'load_balancing_scheme' is set to INTERNAL_MANAGED and the 'protocol' is set to HTTP, HTTPS, or HTTP2.
	CircuitBreakers interface{}
	// Time for which instance will be drained (not accept new connections, but still work to finish started).
	ConnectionDrainingTimeoutSec interface{}
	// Consistent Hash-based load balancing can be used to provide soft session affinity based on HTTP headers, cookies or
	// other properties. This load balancing policy is applicable only for HTTP connections. The affinity to a particular
	// destination host will be lost when one or more hosts are added/removed from the destination service. This field
	// specifies parameters that control consistent hashing. This field only applies when all of the following are true - *
	// 'load_balancing_scheme' is set to INTERNAL_MANAGED * 'protocol' is set to HTTP, HTTPS, or HTTP2 * 'locality_lb_policy'
	// is set to MAGLEV or RING_HASH
	ConsistentHash interface{}
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp interface{}
	// An optional description of this resource.
	Description interface{}
	// Policy for failovers.
	FailoverPolicy interface{}
	// Fingerprint of this resource. A hash of the contents stored in this object. This field is used in optimistic locking.
	Fingerprint interface{}
	// The set of URLs to HealthCheck resources for health checking this RegionBackendService. Currently at most one health
	// check can be specified, and a health check is required.
	HealthChecks interface{}
	// Indicates what kind of load balancing this regional backend service will be used for. A backend service created for one
	// type of load balancing cannot be used with the other(s). Must be 'INTERNAL' or 'INTERNAL_MANAGED'. Defaults to
	// 'INTERNAL'.
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
	// https://ai.google/research/pubs/pub44824 This field is applicable only when the 'load_balancing_scheme' is set to
	// INTERNAL_MANAGED and the 'protocol' is set to HTTP, HTTPS, or HTTP2.
	LocalityLbPolicy interface{}
	// This field denotes the logging options for the load balancer traffic served by this backend service. If logging is
	// enabled, logs will be exported to Stackdriver.
	LogConfig interface{}
	// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and
	// comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression
	// '[a-z]([-a-z0-9]*[a-z0-9])?' which means the first character must be a lowercase letter, and all following characters
	// must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name interface{}
	// The URL of the network to which this backend service belongs. This field can only be specified when the load balancing
	// scheme is set to INTERNAL.
	Network interface{}
	// Settings controlling eviction of unhealthy hosts from the load balancing pool. This field is applicable only when the
	// 'load_balancing_scheme' is set to INTERNAL_MANAGED and the 'protocol' is set to HTTP, HTTPS, or HTTP2.
	OutlierDetection interface{}
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project interface{}
	// The protocol this RegionBackendService uses to communicate with backends. Possible values are HTTP, HTTPS, HTTP2, SSL,
	// TCP, and UDP. The default is HTTP. **NOTE**: HTTP2 is only valid for beta HTTP/2 load balancer types and may result in
	// errors if used with the GA API.
	Protocol interface{}
	// The Region in which the created backend service should reside. If it is not provided, the provider region is used.
	Region interface{}
	// The URI of the created resource.
	SelfLink interface{}
	// Type of session affinity to use. The default is NONE. Session affinity is not applicable if the protocol is UDP.
	SessionAffinity interface{}
	// How many seconds to wait for the backend before considering it a failed request. Default is 30 seconds. Valid range is
	// [1, 86400].
	TimeoutSec interface{}
}

// The set of arguments for constructing a RegionBackendService resource.
type RegionBackendServiceArgs struct {
	// Lifetime of cookies in seconds if session_affinity is GENERATED_COOKIE. If set to 0, the cookie is non-persistent and
	// lasts only until the end of the browser session (or equivalent). The maximum allowed value for TTL is one day. When the
	// load balancing scheme is INTERNAL, this field is not used.
	AffinityCookieTtlSec interface{}
	// The set of backends that serve this RegionBackendService.
	Backends interface{}
	// Settings controlling the volume of connections to a backend service. This field is applicable only when the
	// 'load_balancing_scheme' is set to INTERNAL_MANAGED and the 'protocol' is set to HTTP, HTTPS, or HTTP2.
	CircuitBreakers interface{}
	// Time for which instance will be drained (not accept new connections, but still work to finish started).
	ConnectionDrainingTimeoutSec interface{}
	// Consistent Hash-based load balancing can be used to provide soft session affinity based on HTTP headers, cookies or
	// other properties. This load balancing policy is applicable only for HTTP connections. The affinity to a particular
	// destination host will be lost when one or more hosts are added/removed from the destination service. This field
	// specifies parameters that control consistent hashing. This field only applies when all of the following are true - *
	// 'load_balancing_scheme' is set to INTERNAL_MANAGED * 'protocol' is set to HTTP, HTTPS, or HTTP2 * 'locality_lb_policy'
	// is set to MAGLEV or RING_HASH
	ConsistentHash interface{}
	// An optional description of this resource.
	Description interface{}
	// Policy for failovers.
	FailoverPolicy interface{}
	// The set of URLs to HealthCheck resources for health checking this RegionBackendService. Currently at most one health
	// check can be specified, and a health check is required.
	HealthChecks interface{}
	// Indicates what kind of load balancing this regional backend service will be used for. A backend service created for one
	// type of load balancing cannot be used with the other(s). Must be 'INTERNAL' or 'INTERNAL_MANAGED'. Defaults to
	// 'INTERNAL'.
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
	// https://ai.google/research/pubs/pub44824 This field is applicable only when the 'load_balancing_scheme' is set to
	// INTERNAL_MANAGED and the 'protocol' is set to HTTP, HTTPS, or HTTP2.
	LocalityLbPolicy interface{}
	// This field denotes the logging options for the load balancer traffic served by this backend service. If logging is
	// enabled, logs will be exported to Stackdriver.
	LogConfig interface{}
	// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and
	// comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression
	// '[a-z]([-a-z0-9]*[a-z0-9])?' which means the first character must be a lowercase letter, and all following characters
	// must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name interface{}
	// The URL of the network to which this backend service belongs. This field can only be specified when the load balancing
	// scheme is set to INTERNAL.
	Network interface{}
	// Settings controlling eviction of unhealthy hosts from the load balancing pool. This field is applicable only when the
	// 'load_balancing_scheme' is set to INTERNAL_MANAGED and the 'protocol' is set to HTTP, HTTPS, or HTTP2.
	OutlierDetection interface{}
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project interface{}
	// The protocol this RegionBackendService uses to communicate with backends. Possible values are HTTP, HTTPS, HTTP2, SSL,
	// TCP, and UDP. The default is HTTP. **NOTE**: HTTP2 is only valid for beta HTTP/2 load balancer types and may result in
	// errors if used with the GA API.
	Protocol interface{}
	// The Region in which the created backend service should reside. If it is not provided, the provider region is used.
	Region interface{}
	// Type of session affinity to use. The default is NONE. Session affinity is not applicable if the protocol is UDP.
	SessionAffinity interface{}
	// How many seconds to wait for the backend before considering it a failed request. Default is 30 seconds. Valid range is
	// [1, 86400].
	TimeoutSec interface{}
}
