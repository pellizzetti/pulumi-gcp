// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package storage

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// The Google Cloud storage signed URL data source generates a signed URL for a given storage object. Signed URLs provide a way to give time-limited read or write access to anyone in possession of the URL, regardless of whether they have a Google account.
// 
// For more info about signed URL's is available [here](https://cloud.google.com/storage/docs/access-control/signed-urls).
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/d/storage_object_signed_url.html.markdown.
func LookupObjectSignedUrl(ctx *pulumi.Context, args *GetObjectSignedUrlArgs) (*GetObjectSignedUrlResult, error) {
	inputs := make(map[string]interface{})
	if args != nil {
		inputs["bucket"] = args.Bucket
		inputs["contentMd5"] = args.ContentMd5
		inputs["contentType"] = args.ContentType
		inputs["credentials"] = args.Credentials
		inputs["duration"] = args.Duration
		inputs["extensionHeaders"] = args.ExtensionHeaders
		inputs["httpMethod"] = args.HttpMethod
		inputs["path"] = args.Path
	}
	outputs, err := ctx.Invoke("gcp:storage/getObjectSignedUrl:getObjectSignedUrl", inputs)
	if err != nil {
		return nil, err
	}
	return &GetObjectSignedUrlResult{
		Bucket: outputs["bucket"],
		ContentMd5: outputs["contentMd5"],
		ContentType: outputs["contentType"],
		Credentials: outputs["credentials"],
		Duration: outputs["duration"],
		ExtensionHeaders: outputs["extensionHeaders"],
		HttpMethod: outputs["httpMethod"],
		Path: outputs["path"],
		SignedUrl: outputs["signedUrl"],
		Id: outputs["id"],
	}, nil
}

// A collection of arguments for invoking getObjectSignedUrl.
type GetObjectSignedUrlArgs struct {
	// The name of the bucket to read the object from
	Bucket interface{}
	// The [MD5 digest](https://cloud.google.com/storage/docs/hashes-etags#_MD5) value in Base64.
	// Typically retrieved from `google_storage_bucket_object.object.md5hash` attribute.
	// If you provide this in the datasource, the client (e.g. browser, curl) must provide the `Content-MD5` HTTP header with this same value in its request.
	ContentMd5 interface{}
	// If you specify this in the datasource, the client must provide the `Content-Type` HTTP header with the same value in its request.
	ContentType interface{}
	// What Google service account credentials json should be used to sign the URL.
	// This data source checks the following locations for credentials, in order of preference: data source `credentials` attribute, provider `credentials` attribute and finally the GOOGLE_APPLICATION_CREDENTIALS environment variable.
	Credentials interface{}
	// For how long shall the signed URL be valid (defaults to 1 hour - i.e. `1h`).
	// See [here](https://golang.org/pkg/time/#ParseDuration) for info on valid duration formats.
	Duration interface{}
	// As needed. The server checks to make sure that the client provides matching values in requests using the signed URL.
	// Any header starting with `x-goog-` is accepted but see the [Google Docs](https://cloud.google.com/storage/docs/xml-api/reference-headers) for list of headers that are supported by Google.
	ExtensionHeaders interface{}
	// What HTTP Method will the signed URL allow (defaults to `GET`)
	HttpMethod interface{}
	// The full path to the object inside the bucket
	Path interface{}
}

// A collection of values returned by getObjectSignedUrl.
type GetObjectSignedUrlResult struct {
	Bucket interface{}
	ContentMd5 interface{}
	ContentType interface{}
	Credentials interface{}
	Duration interface{}
	ExtensionHeaders interface{}
	HttpMethod interface{}
	Path interface{}
	// The signed URL that can be used to access the storage object without authentication.
	SignedUrl interface{}
	// id is the provider-assigned unique ID for this managed resource.
	Id interface{}
}
