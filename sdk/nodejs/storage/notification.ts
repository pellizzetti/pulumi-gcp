// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Creates a new notification configuration on a specified bucket, establishing a flow of event notifications from GCS to a Cloud Pub/Sub topic.
 *  For more information see 
 * [the official documentation](https://cloud.google.com/storage/docs/pubsub-notifications) 
 * and 
 * [API](https://cloud.google.com/storage/docs/json_api/v1/notifications).
 * 
 * In order to enable notifications, a special Google Cloud Storage service account unique to the project
 * must have the IAM permission "projects.topics.publish" for a Cloud Pub/Sub topic in the project. To get the service
 * account's email address, use the `gcp.storage.getProjectServiceAccount` datasource's `emailAddress` value, and see below
 * for an example of enabling notifications by granting the correct IAM permission. See
 * [the notifications documentation](https://cloud.google.com/storage/docs/gsutil/commands/notification) for more details.
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/storage_notification.html.markdown.
 */
export class Notification extends pulumi.CustomResource {
    /**
     * Get an existing Notification resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: NotificationState, opts?: pulumi.CustomResourceOptions): Notification {
        return new Notification(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:storage/notification:Notification';

    /**
     * Returns true if the given object is an instance of Notification.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Notification {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Notification.__pulumiType;
    }

    /**
     * The name of the bucket.
     */
    public readonly bucket!: pulumi.Output<string>;
    /**
     * A set of key/value attribute pairs to attach to each Cloud PubSub message published for this notification subscription
     */
    public readonly customAttributes!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * List of event type filters for this notification config. If not specified, Cloud Storage will send notifications for all event types. The valid types are: `"OBJECT_FINALIZE"`, `"OBJECT_METADATA_UPDATE"`, `"OBJECT_DELETE"`, `"OBJECT_ARCHIVE"`
     */
    public readonly eventTypes!: pulumi.Output<string[] | undefined>;
    /**
     * The ID of the created notification.
     */
    public /*out*/ readonly notificationId!: pulumi.Output<string>;
    /**
     * Specifies a prefix path filter for this notification config. Cloud Storage will only send notifications for objects in this bucket whose names begin with the specified prefix.
     */
    public readonly objectNamePrefix!: pulumi.Output<string | undefined>;
    /**
     * The desired content of the Payload. One of `"JSON_API_V1"` or `"NONE"`.
     */
    public readonly payloadFormat!: pulumi.Output<string>;
    /**
     * The URI of the created resource.
     */
    public /*out*/ readonly selfLink!: pulumi.Output<string>;
    /**
     * The Cloud PubSub topic to which this subscription publishes. Expects either the 
     * topic name, assumed to belong to the default GCP provider project, or the project-level name,
     * i.e. `projects/my-gcp-project/topics/my-topic` or `my-topic`.
     */
    public readonly topic!: pulumi.Output<string>;

    /**
     * Create a Notification resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: NotificationArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: NotificationArgs | NotificationState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as NotificationState | undefined;
            inputs["bucket"] = state ? state.bucket : undefined;
            inputs["customAttributes"] = state ? state.customAttributes : undefined;
            inputs["eventTypes"] = state ? state.eventTypes : undefined;
            inputs["notificationId"] = state ? state.notificationId : undefined;
            inputs["objectNamePrefix"] = state ? state.objectNamePrefix : undefined;
            inputs["payloadFormat"] = state ? state.payloadFormat : undefined;
            inputs["selfLink"] = state ? state.selfLink : undefined;
            inputs["topic"] = state ? state.topic : undefined;
        } else {
            const args = argsOrState as NotificationArgs | undefined;
            if (!args || args.bucket === undefined) {
                throw new Error("Missing required property 'bucket'");
            }
            if (!args || args.payloadFormat === undefined) {
                throw new Error("Missing required property 'payloadFormat'");
            }
            if (!args || args.topic === undefined) {
                throw new Error("Missing required property 'topic'");
            }
            inputs["bucket"] = args ? args.bucket : undefined;
            inputs["customAttributes"] = args ? args.customAttributes : undefined;
            inputs["eventTypes"] = args ? args.eventTypes : undefined;
            inputs["objectNamePrefix"] = args ? args.objectNamePrefix : undefined;
            inputs["payloadFormat"] = args ? args.payloadFormat : undefined;
            inputs["topic"] = args ? args.topic : undefined;
            inputs["notificationId"] = undefined /*out*/;
            inputs["selfLink"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(Notification.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Notification resources.
 */
export interface NotificationState {
    /**
     * The name of the bucket.
     */
    readonly bucket?: pulumi.Input<string>;
    /**
     * A set of key/value attribute pairs to attach to each Cloud PubSub message published for this notification subscription
     */
    readonly customAttributes?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * List of event type filters for this notification config. If not specified, Cloud Storage will send notifications for all event types. The valid types are: `"OBJECT_FINALIZE"`, `"OBJECT_METADATA_UPDATE"`, `"OBJECT_DELETE"`, `"OBJECT_ARCHIVE"`
     */
    readonly eventTypes?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The ID of the created notification.
     */
    readonly notificationId?: pulumi.Input<string>;
    /**
     * Specifies a prefix path filter for this notification config. Cloud Storage will only send notifications for objects in this bucket whose names begin with the specified prefix.
     */
    readonly objectNamePrefix?: pulumi.Input<string>;
    /**
     * The desired content of the Payload. One of `"JSON_API_V1"` or `"NONE"`.
     */
    readonly payloadFormat?: pulumi.Input<string>;
    /**
     * The URI of the created resource.
     */
    readonly selfLink?: pulumi.Input<string>;
    /**
     * The Cloud PubSub topic to which this subscription publishes. Expects either the 
     * topic name, assumed to belong to the default GCP provider project, or the project-level name,
     * i.e. `projects/my-gcp-project/topics/my-topic` or `my-topic`.
     */
    readonly topic?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Notification resource.
 */
export interface NotificationArgs {
    /**
     * The name of the bucket.
     */
    readonly bucket: pulumi.Input<string>;
    /**
     * A set of key/value attribute pairs to attach to each Cloud PubSub message published for this notification subscription
     */
    readonly customAttributes?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * List of event type filters for this notification config. If not specified, Cloud Storage will send notifications for all event types. The valid types are: `"OBJECT_FINALIZE"`, `"OBJECT_METADATA_UPDATE"`, `"OBJECT_DELETE"`, `"OBJECT_ARCHIVE"`
     */
    readonly eventTypes?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Specifies a prefix path filter for this notification config. Cloud Storage will only send notifications for objects in this bucket whose names begin with the specified prefix.
     */
    readonly objectNamePrefix?: pulumi.Input<string>;
    /**
     * The desired content of the Payload. One of `"JSON_API_V1"` or `"NONE"`.
     */
    readonly payloadFormat: pulumi.Input<string>;
    /**
     * The Cloud PubSub topic to which this subscription publishes. Expects either the 
     * topic name, assumed to belong to the default GCP provider project, or the project-level name,
     * i.e. `projects/my-gcp-project/topics/my-topic` or `my-topic`.
     */
    readonly topic: pulumi.Input<string>;
}
