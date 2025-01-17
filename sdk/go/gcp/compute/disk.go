// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/compute_disk.html.markdown.
type Disk struct {
	s *pulumi.ResourceState
}

// NewDisk registers a new resource with the given unique name, arguments, and options.
func NewDisk(ctx *pulumi.Context,
	name string, args *DiskArgs, opts ...pulumi.ResourceOpt) (*Disk, error) {
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["description"] = nil
		inputs["diskEncryptionKey"] = nil
		inputs["image"] = nil
		inputs["labels"] = nil
		inputs["name"] = nil
		inputs["physicalBlockSizeBytes"] = nil
		inputs["project"] = nil
		inputs["resourcePolicies"] = nil
		inputs["size"] = nil
		inputs["snapshot"] = nil
		inputs["sourceImageEncryptionKey"] = nil
		inputs["sourceSnapshotEncryptionKey"] = nil
		inputs["type"] = nil
		inputs["zone"] = nil
	} else {
		inputs["description"] = args.Description
		inputs["diskEncryptionKey"] = args.DiskEncryptionKey
		inputs["image"] = args.Image
		inputs["labels"] = args.Labels
		inputs["name"] = args.Name
		inputs["physicalBlockSizeBytes"] = args.PhysicalBlockSizeBytes
		inputs["project"] = args.Project
		inputs["resourcePolicies"] = args.ResourcePolicies
		inputs["size"] = args.Size
		inputs["snapshot"] = args.Snapshot
		inputs["sourceImageEncryptionKey"] = args.SourceImageEncryptionKey
		inputs["sourceSnapshotEncryptionKey"] = args.SourceSnapshotEncryptionKey
		inputs["type"] = args.Type
		inputs["zone"] = args.Zone
	}
	inputs["creationTimestamp"] = nil
	inputs["labelFingerprint"] = nil
	inputs["lastAttachTimestamp"] = nil
	inputs["lastDetachTimestamp"] = nil
	inputs["selfLink"] = nil
	inputs["sourceImageId"] = nil
	inputs["sourceSnapshotId"] = nil
	inputs["users"] = nil
	s, err := ctx.RegisterResource("gcp:compute/disk:Disk", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Disk{s: s}, nil
}

// GetDisk gets an existing Disk resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDisk(ctx *pulumi.Context,
	name string, id pulumi.ID, state *DiskState, opts ...pulumi.ResourceOpt) (*Disk, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["creationTimestamp"] = state.CreationTimestamp
		inputs["description"] = state.Description
		inputs["diskEncryptionKey"] = state.DiskEncryptionKey
		inputs["image"] = state.Image
		inputs["labelFingerprint"] = state.LabelFingerprint
		inputs["labels"] = state.Labels
		inputs["lastAttachTimestamp"] = state.LastAttachTimestamp
		inputs["lastDetachTimestamp"] = state.LastDetachTimestamp
		inputs["name"] = state.Name
		inputs["physicalBlockSizeBytes"] = state.PhysicalBlockSizeBytes
		inputs["project"] = state.Project
		inputs["resourcePolicies"] = state.ResourcePolicies
		inputs["selfLink"] = state.SelfLink
		inputs["size"] = state.Size
		inputs["snapshot"] = state.Snapshot
		inputs["sourceImageEncryptionKey"] = state.SourceImageEncryptionKey
		inputs["sourceImageId"] = state.SourceImageId
		inputs["sourceSnapshotEncryptionKey"] = state.SourceSnapshotEncryptionKey
		inputs["sourceSnapshotId"] = state.SourceSnapshotId
		inputs["type"] = state.Type
		inputs["users"] = state.Users
		inputs["zone"] = state.Zone
	}
	s, err := ctx.ReadResource("gcp:compute/disk:Disk", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Disk{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *Disk) URN() pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *Disk) ID() pulumi.IDOutput {
	return r.s.ID()
}

// Creation timestamp in RFC3339 text format.
func (r *Disk) CreationTimestamp() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["creationTimestamp"])
}

// An optional description of this resource. Provide this property when you create the resource.
func (r *Disk) Description() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["description"])
}

// Encrypts the disk using a customer-supplied encryption key. After you encrypt a disk with a customer-supplied key, you
// must provide the same key if you use the disk later (e.g. to create a disk snapshot or an image, or to attach the disk
// to a virtual machine). Customer-supplied encryption keys do not protect access to metadata of the disk. If you do not
// provide an encryption key when creating the disk, then the disk will be encrypted using an automatically generated key
// and you do not need to provide a key to use the disk later.
func (r *Disk) DiskEncryptionKey() pulumi.Output {
	return r.s.State["diskEncryptionKey"]
}

// The image from which to initialize this disk. This can be one of: the image's 'self_link',
// 'projects/{project}/global/images/{image}', 'projects/{project}/global/images/family/{family}', 'global/images/{image}',
// 'global/images/family/{family}', 'family/{family}', '{project}/{family}', '{project}/{image}', '{family}', or '{image}'.
// If referred by family, the images names must include the family name. If they don't, use the [google_compute_image data
// source](/docs/providers/google/d/datasource_compute_image.html). For instance, the image 'centos-6-v20180104' includes
// its family name 'centos-6'. These images can be referred by family name here.
func (r *Disk) Image() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["image"])
}

// The fingerprint used for optimistic locking of this resource. Used internally during updates.
func (r *Disk) LabelFingerprint() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["labelFingerprint"])
}

// Labels to apply to this disk. A list of key->value pairs.
func (r *Disk) Labels() pulumi.MapOutput {
	return (pulumi.MapOutput)(r.s.State["labels"])
}

// Last attach timestamp in RFC3339 text format.
func (r *Disk) LastAttachTimestamp() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["lastAttachTimestamp"])
}

// Last detach timestamp in RFC3339 text format.
func (r *Disk) LastDetachTimestamp() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["lastDetachTimestamp"])
}

// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and
// comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression
// '[a-z]([-a-z0-9]*[a-z0-9])?' which means the first character must be a lowercase letter, and all following characters
// must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
func (r *Disk) Name() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["name"])
}

// Physical block size of the persistent disk, in bytes. If not present in a request, a default value is used. Currently
// supported sizes are 4096 and 16384, other sizes may be added in the future. If an unsupported value is requested, the
// error message will list the supported values for the caller's project.
func (r *Disk) PhysicalBlockSizeBytes() pulumi.IntOutput {
	return (pulumi.IntOutput)(r.s.State["physicalBlockSizeBytes"])
}

// The ID of the project in which the resource belongs.
// If it is not provided, the provider project is used.
func (r *Disk) Project() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["project"])
}

// Resource policies applied to this disk for automatic snapshot creations.
func (r *Disk) ResourcePolicies() pulumi.ArrayOutput {
	return (pulumi.ArrayOutput)(r.s.State["resourcePolicies"])
}

// The URI of the created resource.
func (r *Disk) SelfLink() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["selfLink"])
}

// Size of the persistent disk, specified in GB. You can specify this field when creating a persistent disk using the
// 'image' or 'snapshot' parameter, or specify it alone to create an empty persistent disk. If you specify this field along
// with 'image' or 'snapshot', the value must not be less than the size of the image or the size of the snapshot.
func (r *Disk) Size() pulumi.IntOutput {
	return (pulumi.IntOutput)(r.s.State["size"])
}

// The source snapshot used to create this disk. You can provide this as a partial or full URL to the resource. If the
// snapshot is in another project than this disk, you must supply a full URL. For example, the following are valid values:
// * 'https://www.googleapis.com/compute/v1/projects/project/global/snapshots/snapshot' *
// 'projects/project/global/snapshots/snapshot' * 'global/snapshots/snapshot' * 'snapshot'
func (r *Disk) Snapshot() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["snapshot"])
}

// The customer-supplied encryption key of the source image. Required if the source image is protected by a
// customer-supplied encryption key.
func (r *Disk) SourceImageEncryptionKey() pulumi.Output {
	return r.s.State["sourceImageEncryptionKey"]
}

// The ID value of the image used to create this disk. This value identifies the exact image that was used to create this
// persistent disk. For example, if you created the persistent disk from an image that was later deleted and recreated
// under the same name, the source image ID would identify the exact version of the image that was used.
func (r *Disk) SourceImageId() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["sourceImageId"])
}

// The customer-supplied encryption key of the source snapshot. Required if the source snapshot is protected by a
// customer-supplied encryption key.
func (r *Disk) SourceSnapshotEncryptionKey() pulumi.Output {
	return r.s.State["sourceSnapshotEncryptionKey"]
}

// The unique ID of the snapshot used to create this disk. This value identifies the exact snapshot that was used to create
// this persistent disk. For example, if you created the persistent disk from a snapshot that was later deleted and
// recreated under the same name, the source snapshot ID would identify the exact version of the snapshot that was used.
func (r *Disk) SourceSnapshotId() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["sourceSnapshotId"])
}

// URL of the disk type resource describing which disk type to use to create the disk. Provide this when creating the disk.
func (r *Disk) Type() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["type"])
}

// Links to the users of the disk (attached instances) in form: project/zones/zone/instances/instance
func (r *Disk) Users() pulumi.ArrayOutput {
	return (pulumi.ArrayOutput)(r.s.State["users"])
}

// A reference to the zone where the disk resides.
func (r *Disk) Zone() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["zone"])
}

// Input properties used for looking up and filtering Disk resources.
type DiskState struct {
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp interface{}
	// An optional description of this resource. Provide this property when you create the resource.
	Description interface{}
	// Encrypts the disk using a customer-supplied encryption key. After you encrypt a disk with a customer-supplied key, you
	// must provide the same key if you use the disk later (e.g. to create a disk snapshot or an image, or to attach the disk
	// to a virtual machine). Customer-supplied encryption keys do not protect access to metadata of the disk. If you do not
	// provide an encryption key when creating the disk, then the disk will be encrypted using an automatically generated key
	// and you do not need to provide a key to use the disk later.
	DiskEncryptionKey interface{}
	// The image from which to initialize this disk. This can be one of: the image's 'self_link',
	// 'projects/{project}/global/images/{image}', 'projects/{project}/global/images/family/{family}',
	// 'global/images/{image}', 'global/images/family/{family}', 'family/{family}', '{project}/{family}', '{project}/{image}',
	// '{family}', or '{image}'. If referred by family, the images names must include the family name. If they don't, use the
	// [google_compute_image data source](/docs/providers/google/d/datasource_compute_image.html). For instance, the image
	// 'centos-6-v20180104' includes its family name 'centos-6'. These images can be referred by family name here.
	Image interface{}
	// The fingerprint used for optimistic locking of this resource. Used internally during updates.
	LabelFingerprint interface{}
	// Labels to apply to this disk. A list of key->value pairs.
	Labels interface{}
	// Last attach timestamp in RFC3339 text format.
	LastAttachTimestamp interface{}
	// Last detach timestamp in RFC3339 text format.
	LastDetachTimestamp interface{}
	// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and
	// comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression
	// '[a-z]([-a-z0-9]*[a-z0-9])?' which means the first character must be a lowercase letter, and all following characters
	// must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name interface{}
	// Physical block size of the persistent disk, in bytes. If not present in a request, a default value is used. Currently
	// supported sizes are 4096 and 16384, other sizes may be added in the future. If an unsupported value is requested, the
	// error message will list the supported values for the caller's project.
	PhysicalBlockSizeBytes interface{}
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project interface{}
	// Resource policies applied to this disk for automatic snapshot creations.
	ResourcePolicies interface{}
	// The URI of the created resource.
	SelfLink interface{}
	// Size of the persistent disk, specified in GB. You can specify this field when creating a persistent disk using the
	// 'image' or 'snapshot' parameter, or specify it alone to create an empty persistent disk. If you specify this field
	// along with 'image' or 'snapshot', the value must not be less than the size of the image or the size of the snapshot.
	Size interface{}
	// The source snapshot used to create this disk. You can provide this as a partial or full URL to the resource. If the
	// snapshot is in another project than this disk, you must supply a full URL. For example, the following are valid values:
	// * 'https://www.googleapis.com/compute/v1/projects/project/global/snapshots/snapshot' *
	// 'projects/project/global/snapshots/snapshot' * 'global/snapshots/snapshot' * 'snapshot'
	Snapshot interface{}
	// The customer-supplied encryption key of the source image. Required if the source image is protected by a
	// customer-supplied encryption key.
	SourceImageEncryptionKey interface{}
	// The ID value of the image used to create this disk. This value identifies the exact image that was used to create this
	// persistent disk. For example, if you created the persistent disk from an image that was later deleted and recreated
	// under the same name, the source image ID would identify the exact version of the image that was used.
	SourceImageId interface{}
	// The customer-supplied encryption key of the source snapshot. Required if the source snapshot is protected by a
	// customer-supplied encryption key.
	SourceSnapshotEncryptionKey interface{}
	// The unique ID of the snapshot used to create this disk. This value identifies the exact snapshot that was used to
	// create this persistent disk. For example, if you created the persistent disk from a snapshot that was later deleted and
	// recreated under the same name, the source snapshot ID would identify the exact version of the snapshot that was used.
	SourceSnapshotId interface{}
	// URL of the disk type resource describing which disk type to use to create the disk. Provide this when creating the
	// disk.
	Type interface{}
	// Links to the users of the disk (attached instances) in form: project/zones/zone/instances/instance
	Users interface{}
	// A reference to the zone where the disk resides.
	Zone interface{}
}

// The set of arguments for constructing a Disk resource.
type DiskArgs struct {
	// An optional description of this resource. Provide this property when you create the resource.
	Description interface{}
	// Encrypts the disk using a customer-supplied encryption key. After you encrypt a disk with a customer-supplied key, you
	// must provide the same key if you use the disk later (e.g. to create a disk snapshot or an image, or to attach the disk
	// to a virtual machine). Customer-supplied encryption keys do not protect access to metadata of the disk. If you do not
	// provide an encryption key when creating the disk, then the disk will be encrypted using an automatically generated key
	// and you do not need to provide a key to use the disk later.
	DiskEncryptionKey interface{}
	// The image from which to initialize this disk. This can be one of: the image's 'self_link',
	// 'projects/{project}/global/images/{image}', 'projects/{project}/global/images/family/{family}',
	// 'global/images/{image}', 'global/images/family/{family}', 'family/{family}', '{project}/{family}', '{project}/{image}',
	// '{family}', or '{image}'. If referred by family, the images names must include the family name. If they don't, use the
	// [google_compute_image data source](/docs/providers/google/d/datasource_compute_image.html). For instance, the image
	// 'centos-6-v20180104' includes its family name 'centos-6'. These images can be referred by family name here.
	Image interface{}
	// Labels to apply to this disk. A list of key->value pairs.
	Labels interface{}
	// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and
	// comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression
	// '[a-z]([-a-z0-9]*[a-z0-9])?' which means the first character must be a lowercase letter, and all following characters
	// must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name interface{}
	// Physical block size of the persistent disk, in bytes. If not present in a request, a default value is used. Currently
	// supported sizes are 4096 and 16384, other sizes may be added in the future. If an unsupported value is requested, the
	// error message will list the supported values for the caller's project.
	PhysicalBlockSizeBytes interface{}
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project interface{}
	// Resource policies applied to this disk for automatic snapshot creations.
	ResourcePolicies interface{}
	// Size of the persistent disk, specified in GB. You can specify this field when creating a persistent disk using the
	// 'image' or 'snapshot' parameter, or specify it alone to create an empty persistent disk. If you specify this field
	// along with 'image' or 'snapshot', the value must not be less than the size of the image or the size of the snapshot.
	Size interface{}
	// The source snapshot used to create this disk. You can provide this as a partial or full URL to the resource. If the
	// snapshot is in another project than this disk, you must supply a full URL. For example, the following are valid values:
	// * 'https://www.googleapis.com/compute/v1/projects/project/global/snapshots/snapshot' *
	// 'projects/project/global/snapshots/snapshot' * 'global/snapshots/snapshot' * 'snapshot'
	Snapshot interface{}
	// The customer-supplied encryption key of the source image. Required if the source image is protected by a
	// customer-supplied encryption key.
	SourceImageEncryptionKey interface{}
	// The customer-supplied encryption key of the source snapshot. Required if the source snapshot is protected by a
	// customer-supplied encryption key.
	SourceSnapshotEncryptionKey interface{}
	// URL of the disk type resource describing which disk type to use to create the disk. Provide this when creating the
	// disk.
	Type interface{}
	// A reference to the zone where the disk resides.
	Zone interface{}
}
