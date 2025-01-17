# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class Notification(pulumi.CustomResource):
    bucket: pulumi.Output[str]
    """
    The name of the bucket.
    """
    custom_attributes: pulumi.Output[dict]
    """
    A set of key/value attribute pairs to attach to each Cloud PubSub message published for this notification subscription
    """
    event_types: pulumi.Output[list]
    """
    List of event type filters for this notification config. If not specified, Cloud Storage will send notifications for all event types. The valid types are: `"OBJECT_FINALIZE"`, `"OBJECT_METADATA_UPDATE"`, `"OBJECT_DELETE"`, `"OBJECT_ARCHIVE"`
    """
    notification_id: pulumi.Output[str]
    """
    The ID of the created notification.
    """
    object_name_prefix: pulumi.Output[str]
    """
    Specifies a prefix path filter for this notification config. Cloud Storage will only send notifications for objects in this bucket whose names begin with the specified prefix.
    """
    payload_format: pulumi.Output[str]
    """
    The desired content of the Payload. One of `"JSON_API_V1"` or `"NONE"`.
    """
    self_link: pulumi.Output[str]
    """
    The URI of the created resource.
    """
    topic: pulumi.Output[str]
    """
    The Cloud PubSub topic to which this subscription publishes. Expects either the 
    topic name, assumed to belong to the default GCP provider project, or the project-level name,
    i.e. `projects/my-gcp-project/topics/my-topic` or `my-topic`.
    """
    def __init__(__self__, resource_name, opts=None, bucket=None, custom_attributes=None, event_types=None, object_name_prefix=None, payload_format=None, topic=None, __props__=None, __name__=None, __opts__=None):
        """
        Creates a new notification configuration on a specified bucket, establishing a flow of event notifications from GCS to a Cloud Pub/Sub topic.
         For more information see 
        [the official documentation](https://cloud.google.com/storage/docs/pubsub-notifications) 
        and 
        [API](https://cloud.google.com/storage/docs/json_api/v1/notifications).
        
        In order to enable notifications, a special Google Cloud Storage service account unique to the project
        must have the IAM permission "projects.topics.publish" for a Cloud Pub/Sub topic in the project. To get the service
        account's email address, use the `storage.getProjectServiceAccount` datasource's `email_address` value, and see below
        for an example of enabling notifications by granting the correct IAM permission. See
        [the notifications documentation](https://cloud.google.com/storage/docs/gsutil/commands/notification) for more details.
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] bucket: The name of the bucket.
        :param pulumi.Input[dict] custom_attributes: A set of key/value attribute pairs to attach to each Cloud PubSub message published for this notification subscription
        :param pulumi.Input[list] event_types: List of event type filters for this notification config. If not specified, Cloud Storage will send notifications for all event types. The valid types are: `"OBJECT_FINALIZE"`, `"OBJECT_METADATA_UPDATE"`, `"OBJECT_DELETE"`, `"OBJECT_ARCHIVE"`
        :param pulumi.Input[str] object_name_prefix: Specifies a prefix path filter for this notification config. Cloud Storage will only send notifications for objects in this bucket whose names begin with the specified prefix.
        :param pulumi.Input[str] payload_format: The desired content of the Payload. One of `"JSON_API_V1"` or `"NONE"`.
        :param pulumi.Input[str] topic: The Cloud PubSub topic to which this subscription publishes. Expects either the 
               topic name, assumed to belong to the default GCP provider project, or the project-level name,
               i.e. `projects/my-gcp-project/topics/my-topic` or `my-topic`.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/storage_notification.html.markdown.
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            if bucket is None:
                raise TypeError("Missing required property 'bucket'")
            __props__['bucket'] = bucket
            __props__['custom_attributes'] = custom_attributes
            __props__['event_types'] = event_types
            __props__['object_name_prefix'] = object_name_prefix
            if payload_format is None:
                raise TypeError("Missing required property 'payload_format'")
            __props__['payload_format'] = payload_format
            if topic is None:
                raise TypeError("Missing required property 'topic'")
            __props__['topic'] = topic
            __props__['notification_id'] = None
            __props__['self_link'] = None
        super(Notification, __self__).__init__(
            'gcp:storage/notification:Notification',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, bucket=None, custom_attributes=None, event_types=None, notification_id=None, object_name_prefix=None, payload_format=None, self_link=None, topic=None):
        """
        Get an existing Notification resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.
        
        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] bucket: The name of the bucket.
        :param pulumi.Input[dict] custom_attributes: A set of key/value attribute pairs to attach to each Cloud PubSub message published for this notification subscription
        :param pulumi.Input[list] event_types: List of event type filters for this notification config. If not specified, Cloud Storage will send notifications for all event types. The valid types are: `"OBJECT_FINALIZE"`, `"OBJECT_METADATA_UPDATE"`, `"OBJECT_DELETE"`, `"OBJECT_ARCHIVE"`
        :param pulumi.Input[str] notification_id: The ID of the created notification.
        :param pulumi.Input[str] object_name_prefix: Specifies a prefix path filter for this notification config. Cloud Storage will only send notifications for objects in this bucket whose names begin with the specified prefix.
        :param pulumi.Input[str] payload_format: The desired content of the Payload. One of `"JSON_API_V1"` or `"NONE"`.
        :param pulumi.Input[str] self_link: The URI of the created resource.
        :param pulumi.Input[str] topic: The Cloud PubSub topic to which this subscription publishes. Expects either the 
               topic name, assumed to belong to the default GCP provider project, or the project-level name,
               i.e. `projects/my-gcp-project/topics/my-topic` or `my-topic`.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/storage_notification.html.markdown.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()
        __props__["bucket"] = bucket
        __props__["custom_attributes"] = custom_attributes
        __props__["event_types"] = event_types
        __props__["notification_id"] = notification_id
        __props__["object_name_prefix"] = object_name_prefix
        __props__["payload_format"] = payload_format
        __props__["self_link"] = self_link
        __props__["topic"] = topic
        return Notification(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

