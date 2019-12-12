// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Storage
{
    /// <summary>
    /// Three different resources help you manage your IAM policy for storage bucket. Each of these resources serves a different use case:
    /// 
    /// * `gcp.storage.BucketIAMBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the storage bucket are preserved.
    /// * `gcp.storage.BucketIAMMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the storage bucket are preserved.
    /// * `gcp.storage.BucketIAMPolicy`: Setting a policy removes all other permissions on the bucket, and if done incorrectly, there's a real chance you will lock yourself out of the bucket. If possible for your use case, using multiple gcp.storage.BucketIAMBinding resources will be much safer. See the usage example on how to work with policy correctly.
    /// 
    /// 
    /// &gt; **Note:** `gcp.storage.BucketIAMBinding` resources **can be** used in conjunction with `gcp.storage.BucketIAMMember` resources **only if** they do not grant privilege to the same role.
    /// 
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/storage_bucket_iam_member.html.markdown.
    /// </summary>
    public partial class BucketIAMMember : Pulumi.CustomResource
    {
        /// <summary>
        /// The name of the bucket it applies to.
        /// </summary>
        [Output("bucket")]
        public Output<string> Bucket { get; private set; } = null!;

        [Output("condition")]
        public Output<Outputs.BucketIAMMemberCondition?> Condition { get; private set; } = null!;

        /// <summary>
        /// (Computed) The etag of the storage bucket's IAM policy.
        /// </summary>
        [Output("etag")]
        public Output<string> Etag { get; private set; } = null!;

        [Output("member")]
        public Output<string> Member { get; private set; } = null!;

        /// <summary>
        /// The role that should be applied. Note that custom roles must be of the format
        /// `[projects|organizations]/{parent-name}/roles/{role-name}`.
        /// </summary>
        [Output("role")]
        public Output<string> Role { get; private set; } = null!;


        /// <summary>
        /// Create a BucketIAMMember resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public BucketIAMMember(string name, BucketIAMMemberArgs args, CustomResourceOptions? options = null)
            : base("gcp:storage/bucketIAMMember:BucketIAMMember", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private BucketIAMMember(string name, Input<string> id, BucketIAMMemberState? state = null, CustomResourceOptions? options = null)
            : base("gcp:storage/bucketIAMMember:BucketIAMMember", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing BucketIAMMember resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static BucketIAMMember Get(string name, Input<string> id, BucketIAMMemberState? state = null, CustomResourceOptions? options = null)
        {
            return new BucketIAMMember(name, id, state, options);
        }
    }

    public sealed class BucketIAMMemberArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the bucket it applies to.
        /// </summary>
        [Input("bucket", required: true)]
        public Input<string> Bucket { get; set; } = null!;

        [Input("condition")]
        public Input<Inputs.BucketIAMMemberConditionArgs>? Condition { get; set; }

        [Input("member", required: true)]
        public Input<string> Member { get; set; } = null!;

        /// <summary>
        /// The role that should be applied. Note that custom roles must be of the format
        /// `[projects|organizations]/{parent-name}/roles/{role-name}`.
        /// </summary>
        [Input("role", required: true)]
        public Input<string> Role { get; set; } = null!;

        public BucketIAMMemberArgs()
        {
        }
    }

    public sealed class BucketIAMMemberState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the bucket it applies to.
        /// </summary>
        [Input("bucket")]
        public Input<string>? Bucket { get; set; }

        [Input("condition")]
        public Input<Inputs.BucketIAMMemberConditionGetArgs>? Condition { get; set; }

        /// <summary>
        /// (Computed) The etag of the storage bucket's IAM policy.
        /// </summary>
        [Input("etag")]
        public Input<string>? Etag { get; set; }

        [Input("member")]
        public Input<string>? Member { get; set; }

        /// <summary>
        /// The role that should be applied. Note that custom roles must be of the format
        /// `[projects|organizations]/{parent-name}/roles/{role-name}`.
        /// </summary>
        [Input("role")]
        public Input<string>? Role { get; set; }

        public BucketIAMMemberState()
        {
        }
    }

    namespace Inputs
    {

    public sealed class BucketIAMMemberConditionArgs : Pulumi.ResourceArgs
    {
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("expression", required: true)]
        public Input<string> Expression { get; set; } = null!;

        [Input("title", required: true)]
        public Input<string> Title { get; set; } = null!;

        public BucketIAMMemberConditionArgs()
        {
        }
    }

    public sealed class BucketIAMMemberConditionGetArgs : Pulumi.ResourceArgs
    {
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("expression", required: true)]
        public Input<string> Expression { get; set; } = null!;

        [Input("title", required: true)]
        public Input<string> Title { get; set; } = null!;

        public BucketIAMMemberConditionGetArgs()
        {
        }
    }
    }

    namespace Outputs
    {

    [OutputType]
    public sealed class BucketIAMMemberCondition
    {
        public readonly string? Description;
        public readonly string Expression;
        public readonly string Title;

        [OutputConstructor]
        private BucketIAMMemberCondition(
            string? description,
            string expression,
            string title)
        {
            Description = description;
            Expression = expression;
            Title = title;
        }
    }
    }
}