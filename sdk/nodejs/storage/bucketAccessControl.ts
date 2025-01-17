// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * The BucketAccessControls resource represents the Access Control Lists
 * (ACLs) for buckets within Google Cloud Storage. ACLs let you specify who
 * has access to your data and to what extent.
 * 
 * There are three roles that can be assigned to an entity:
 * 
 * READERs can get the bucket, though no acl property will be returned, and
 * list the bucket's objects.  WRITERs are READERs, and they can insert
 * objects into the bucket and delete the bucket's objects.  OWNERs are
 * WRITERs, and they can get the acl property of a bucket, update a bucket,
 * and call all BucketAccessControls methods on the bucket.  For more
 * information, see Access Control, with the caveat that this API uses
 * READER, WRITER, and OWNER instead of READ, WRITE, and FULL_CONTROL.
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/storage_bucket_access_control.html.markdown.
 */
export class BucketAccessControl extends pulumi.CustomResource {
    /**
     * Get an existing BucketAccessControl resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: BucketAccessControlState, opts?: pulumi.CustomResourceOptions): BucketAccessControl {
        return new BucketAccessControl(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:storage/bucketAccessControl:BucketAccessControl';

    /**
     * Returns true if the given object is an instance of BucketAccessControl.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is BucketAccessControl {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === BucketAccessControl.__pulumiType;
    }

    /**
     * The name of the bucket.
     */
    public readonly bucket!: pulumi.Output<string>;
    /**
     * The domain associated with the entity.
     */
    public /*out*/ readonly domain!: pulumi.Output<string>;
    /**
     * The email address associated with the entity.
     */
    public /*out*/ readonly email!: pulumi.Output<string>;
    /**
     * The entity holding the permission, in one of the following forms: user-userId user-email group-groupId group-email
     * domain-domain project-team-projectId allUsers allAuthenticatedUsers Examples: The user liz@example.com would be
     * user-liz@example.com. The group example@googlegroups.com would be group-example@googlegroups.com. To refer to all
     * members of the Google Apps for Business domain example.com, the entity would be domain-example.com.
     */
    public readonly entity!: pulumi.Output<string>;
    /**
     * The access permission for the entity.
     */
    public readonly role!: pulumi.Output<string | undefined>;

    /**
     * Create a BucketAccessControl resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: BucketAccessControlArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: BucketAccessControlArgs | BucketAccessControlState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as BucketAccessControlState | undefined;
            inputs["bucket"] = state ? state.bucket : undefined;
            inputs["domain"] = state ? state.domain : undefined;
            inputs["email"] = state ? state.email : undefined;
            inputs["entity"] = state ? state.entity : undefined;
            inputs["role"] = state ? state.role : undefined;
        } else {
            const args = argsOrState as BucketAccessControlArgs | undefined;
            if (!args || args.bucket === undefined) {
                throw new Error("Missing required property 'bucket'");
            }
            if (!args || args.entity === undefined) {
                throw new Error("Missing required property 'entity'");
            }
            inputs["bucket"] = args ? args.bucket : undefined;
            inputs["entity"] = args ? args.entity : undefined;
            inputs["role"] = args ? args.role : undefined;
            inputs["domain"] = undefined /*out*/;
            inputs["email"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(BucketAccessControl.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering BucketAccessControl resources.
 */
export interface BucketAccessControlState {
    /**
     * The name of the bucket.
     */
    readonly bucket?: pulumi.Input<string>;
    /**
     * The domain associated with the entity.
     */
    readonly domain?: pulumi.Input<string>;
    /**
     * The email address associated with the entity.
     */
    readonly email?: pulumi.Input<string>;
    /**
     * The entity holding the permission, in one of the following forms: user-userId user-email group-groupId group-email
     * domain-domain project-team-projectId allUsers allAuthenticatedUsers Examples: The user liz@example.com would be
     * user-liz@example.com. The group example@googlegroups.com would be group-example@googlegroups.com. To refer to all
     * members of the Google Apps for Business domain example.com, the entity would be domain-example.com.
     */
    readonly entity?: pulumi.Input<string>;
    /**
     * The access permission for the entity.
     */
    readonly role?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a BucketAccessControl resource.
 */
export interface BucketAccessControlArgs {
    /**
     * The name of the bucket.
     */
    readonly bucket: pulumi.Input<string>;
    /**
     * The entity holding the permission, in one of the following forms: user-userId user-email group-groupId group-email
     * domain-domain project-team-projectId allUsers allAuthenticatedUsers Examples: The user liz@example.com would be
     * user-liz@example.com. The group example@googlegroups.com would be group-example@googlegroups.com. To refer to all
     * members of the Google Apps for Business domain example.com, the entity would be domain-example.com.
     */
    readonly entity: pulumi.Input<string>;
    /**
     * The access permission for the entity.
     */
    readonly role?: pulumi.Input<string>;
}
