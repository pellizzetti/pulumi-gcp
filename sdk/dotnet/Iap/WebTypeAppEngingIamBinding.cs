// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Iap
{
    /// <summary>
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/iap_web_type_app_engine_iam_binding.html.markdown.
    /// </summary>
    public partial class WebTypeAppEngingIamBinding : Pulumi.CustomResource
    {
        /// <summary>
        /// Id of the App Engine application. Used to find the parent resource to bind the IAM policy to
        /// </summary>
        [Output("appId")]
        public Output<string> AppId { get; private set; } = null!;

        [Output("condition")]
        public Output<Outputs.WebTypeAppEngingIamBindingCondition?> Condition { get; private set; } = null!;

        /// <summary>
        /// (Computed) The etag of the IAM policy.
        /// </summary>
        [Output("etag")]
        public Output<string> Etag { get; private set; } = null!;

        [Output("members")]
        public Output<ImmutableArray<string>> Members { get; private set; } = null!;

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
        /// </summary>
        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        /// <summary>
        /// The role that should be applied. Only one
        /// `gcp.iap.WebTypeAppEngingIamBinding` can be used per role. Note that custom roles must be of the format
        /// `[projects|organizations]/{parent-name}/roles/{role-name}`.
        /// </summary>
        [Output("role")]
        public Output<string> Role { get; private set; } = null!;


        /// <summary>
        /// Create a WebTypeAppEngingIamBinding resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public WebTypeAppEngingIamBinding(string name, WebTypeAppEngingIamBindingArgs args, CustomResourceOptions? options = null)
            : base("gcp:iap/webTypeAppEngingIamBinding:WebTypeAppEngingIamBinding", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private WebTypeAppEngingIamBinding(string name, Input<string> id, WebTypeAppEngingIamBindingState? state = null, CustomResourceOptions? options = null)
            : base("gcp:iap/webTypeAppEngingIamBinding:WebTypeAppEngingIamBinding", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing WebTypeAppEngingIamBinding resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static WebTypeAppEngingIamBinding Get(string name, Input<string> id, WebTypeAppEngingIamBindingState? state = null, CustomResourceOptions? options = null)
        {
            return new WebTypeAppEngingIamBinding(name, id, state, options);
        }
    }

    public sealed class WebTypeAppEngingIamBindingArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Id of the App Engine application. Used to find the parent resource to bind the IAM policy to
        /// </summary>
        [Input("appId", required: true)]
        public Input<string> AppId { get; set; } = null!;

        [Input("condition")]
        public Input<Inputs.WebTypeAppEngingIamBindingConditionArgs>? Condition { get; set; }

        [Input("members", required: true)]
        private InputList<string>? _members;
        public InputList<string> Members
        {
            get => _members ?? (_members = new InputList<string>());
            set => _members = value;
        }

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// The role that should be applied. Only one
        /// `gcp.iap.WebTypeAppEngingIamBinding` can be used per role. Note that custom roles must be of the format
        /// `[projects|organizations]/{parent-name}/roles/{role-name}`.
        /// </summary>
        [Input("role", required: true)]
        public Input<string> Role { get; set; } = null!;

        public WebTypeAppEngingIamBindingArgs()
        {
        }
    }

    public sealed class WebTypeAppEngingIamBindingState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Id of the App Engine application. Used to find the parent resource to bind the IAM policy to
        /// </summary>
        [Input("appId")]
        public Input<string>? AppId { get; set; }

        [Input("condition")]
        public Input<Inputs.WebTypeAppEngingIamBindingConditionGetArgs>? Condition { get; set; }

        /// <summary>
        /// (Computed) The etag of the IAM policy.
        /// </summary>
        [Input("etag")]
        public Input<string>? Etag { get; set; }

        [Input("members")]
        private InputList<string>? _members;
        public InputList<string> Members
        {
            get => _members ?? (_members = new InputList<string>());
            set => _members = value;
        }

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// The role that should be applied. Only one
        /// `gcp.iap.WebTypeAppEngingIamBinding` can be used per role. Note that custom roles must be of the format
        /// `[projects|organizations]/{parent-name}/roles/{role-name}`.
        /// </summary>
        [Input("role")]
        public Input<string>? Role { get; set; }

        public WebTypeAppEngingIamBindingState()
        {
        }
    }

    namespace Inputs
    {

    public sealed class WebTypeAppEngingIamBindingConditionArgs : Pulumi.ResourceArgs
    {
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("expression", required: true)]
        public Input<string> Expression { get; set; } = null!;

        [Input("title", required: true)]
        public Input<string> Title { get; set; } = null!;

        public WebTypeAppEngingIamBindingConditionArgs()
        {
        }
    }

    public sealed class WebTypeAppEngingIamBindingConditionGetArgs : Pulumi.ResourceArgs
    {
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("expression", required: true)]
        public Input<string> Expression { get; set; } = null!;

        [Input("title", required: true)]
        public Input<string> Title { get; set; } = null!;

        public WebTypeAppEngingIamBindingConditionGetArgs()
        {
        }
    }
    }

    namespace Outputs
    {

    [OutputType]
    public sealed class WebTypeAppEngingIamBindingCondition
    {
        public readonly string? Description;
        public readonly string Expression;
        public readonly string Title;

        [OutputConstructor]
        private WebTypeAppEngingIamBindingCondition(
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
