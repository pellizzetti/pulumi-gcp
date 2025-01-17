// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Organizations
{
    public static partial class Invokes
    {
        /// <summary>
        /// Generates an IAM policy document that may be referenced by and applied to
        /// other Google Cloud Platform resources, such as the `gcp.organizations.Project` resource.
        /// 
        /// 
        /// This data source is used to define IAM policies to apply to other resources.
        /// Currently, defining a policy through a datasource and referencing that policy
        /// from another resource is the only way to apply an IAM policy to a resource.
        /// 
        /// **Note:** Several restrictions apply when setting IAM policies through this API.
        /// See the [setIamPolicy docs](https://cloud.google.com/resource-manager/reference/rest/v1/projects/setIamPolicy)
        /// for a list of these restrictions.
        /// 
        /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/d/iam_policy.html.markdown.
        /// </summary>
        public static Task<GetIAMPolicyResult> GetIAMPolicy(GetIAMPolicyArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetIAMPolicyResult>("gcp:organizations/getIAMPolicy:getIAMPolicy", args ?? InvokeArgs.Empty, options.WithVersion());
    }

    public sealed class GetIAMPolicyArgs : Pulumi.InvokeArgs
    {
        [Input("auditConfigs")]
        private List<Inputs.GetIAMPolicyAuditConfigsArgs>? _auditConfigs;

        /// <summary>
        /// A nested configuration block that defines logging additional configuration for your project.
        /// </summary>
        public List<Inputs.GetIAMPolicyAuditConfigsArgs> AuditConfigs
        {
            get => _auditConfigs ?? (_auditConfigs = new List<Inputs.GetIAMPolicyAuditConfigsArgs>());
            set => _auditConfigs = value;
        }

        [Input("bindings", required: true)]
        private List<Inputs.GetIAMPolicyBindingsArgs>? _bindings;

        /// <summary>
        /// A nested configuration block (described below)
        /// defining a binding to be included in the policy document. Multiple
        /// `binding` arguments are supported.
        /// </summary>
        public List<Inputs.GetIAMPolicyBindingsArgs> Bindings
        {
            get => _bindings ?? (_bindings = new List<Inputs.GetIAMPolicyBindingsArgs>());
            set => _bindings = value;
        }

        public GetIAMPolicyArgs()
        {
        }
    }

    [OutputType]
    public sealed class GetIAMPolicyResult
    {
        public readonly ImmutableArray<Outputs.GetIAMPolicyAuditConfigsResult> AuditConfigs;
        public readonly ImmutableArray<Outputs.GetIAMPolicyBindingsResult> Bindings;
        /// <summary>
        /// The above bindings serialized in a format suitable for
        /// referencing from a resource that supports IAM.
        /// </summary>
        public readonly string PolicyData;
        /// <summary>
        /// id is the provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;

        [OutputConstructor]
        private GetIAMPolicyResult(
            ImmutableArray<Outputs.GetIAMPolicyAuditConfigsResult> auditConfigs,
            ImmutableArray<Outputs.GetIAMPolicyBindingsResult> bindings,
            string policyData,
            string id)
        {
            AuditConfigs = auditConfigs;
            Bindings = bindings;
            PolicyData = policyData;
            Id = id;
        }
    }

    namespace Inputs
    {

    public sealed class GetIAMPolicyAuditConfigsArgs : Pulumi.InvokeArgs
    {
        [Input("auditLogConfigs", required: true)]
        private List<GetIAMPolicyAuditConfigsAuditLogConfigsArgs>? _auditLogConfigs;

        /// <summary>
        /// A nested block that defines the operations you'd like to log.
        /// </summary>
        public List<GetIAMPolicyAuditConfigsAuditLogConfigsArgs> AuditLogConfigs
        {
            get => _auditLogConfigs ?? (_auditLogConfigs = new List<GetIAMPolicyAuditConfigsAuditLogConfigsArgs>());
            set => _auditLogConfigs = value;
        }

        /// <summary>
        /// Defines a service that will be enabled for audit logging. For example, `storage.googleapis.com`, `cloudsql.googleapis.com`. `allServices` is a special value that covers all services.
        /// </summary>
        [Input("service", required: true)]
        public string Service { get; set; } = null!;

        public GetIAMPolicyAuditConfigsArgs()
        {
        }
    }

    public sealed class GetIAMPolicyAuditConfigsAuditLogConfigsArgs : Pulumi.InvokeArgs
    {
        [Input("exemptedMembers")]
        private List<string>? _exemptedMembers;

        /// <summary>
        /// Specifies the identities that are exempt from these types of logging operations. Follows the same format of the `members` array for `binding`.
        /// </summary>
        public List<string> ExemptedMembers
        {
            get => _exemptedMembers ?? (_exemptedMembers = new List<string>());
            set => _exemptedMembers = value;
        }

        /// <summary>
        /// Defines the logging level. `DATA_READ`, `DATA_WRITE` and `ADMIN_READ` capture different types of events. See [the audit configuration documentation](https://cloud.google.com/resource-manager/reference/rest/Shared.Types/AuditConfig) for more details.
        /// </summary>
        [Input("logType", required: true)]
        public string LogType { get; set; } = null!;

        public GetIAMPolicyAuditConfigsAuditLogConfigsArgs()
        {
        }
    }

    public sealed class GetIAMPolicyBindingsArgs : Pulumi.InvokeArgs
    {
        [Input("condition")]
        public GetIAMPolicyBindingsConditionArgs? Condition { get; set; }

        [Input("members", required: true)]
        private List<string>? _members;

        /// <summary>
        /// An array of identities that will be granted the privilege in the `role`. For more details on format and restrictions see https://cloud.google.com/billing/reference/rest/v1/Policy#Binding
        /// Each entry can have one of the following values:
        /// * **allUsers**: A special identifier that represents anyone who is on the internet; with or without a Google account. It **can't** be used with the `gcp.organizations.Project` resource.
        /// * **allAuthenticatedUsers**: A special identifier that represents anyone who is authenticated with a Google account or a service account. It **can't** be used with the `gcp.organizations.Project` resource.
        /// * **user:{emailid}**: An email address that represents a specific Google account. For example, alice@gmail.com.
        /// * **serviceAccount:{emailid}**: An email address that represents a service account. For example, my-other-app@appspot.gserviceaccount.com.
        /// * **group:{emailid}**: An email address that represents a Google group. For example, admins@example.com.
        /// * **domain:{domain}**: A G Suite domain (primary, instead of alias) name that represents all the users of that domain. For example, google.com or example.com.
        /// </summary>
        public List<string> Members
        {
            get => _members ?? (_members = new List<string>());
            set => _members = value;
        }

        /// <summary>
        /// The role/permission that will be granted to the members.
        /// See the [IAM Roles](https://cloud.google.com/compute/docs/access/iam) documentation for a complete list of roles.
        /// Note that custom roles must be of the format `[projects|organizations]/{parent-name}/roles/{role-name}`.
        /// </summary>
        [Input("role", required: true)]
        public string Role { get; set; } = null!;

        public GetIAMPolicyBindingsArgs()
        {
        }
    }

    public sealed class GetIAMPolicyBindingsConditionArgs : Pulumi.InvokeArgs
    {
        [Input("description")]
        public string? Description { get; set; }

        [Input("expression", required: true)]
        public string Expression { get; set; } = null!;

        [Input("title", required: true)]
        public string Title { get; set; } = null!;

        public GetIAMPolicyBindingsConditionArgs()
        {
        }
    }
    }

    namespace Outputs
    {

    [OutputType]
    public sealed class GetIAMPolicyAuditConfigsAuditLogConfigsResult
    {
        /// <summary>
        /// Specifies the identities that are exempt from these types of logging operations. Follows the same format of the `members` array for `binding`.
        /// </summary>
        public readonly ImmutableArray<string> ExemptedMembers;
        /// <summary>
        /// Defines the logging level. `DATA_READ`, `DATA_WRITE` and `ADMIN_READ` capture different types of events. See [the audit configuration documentation](https://cloud.google.com/resource-manager/reference/rest/Shared.Types/AuditConfig) for more details.
        /// </summary>
        public readonly string LogType;

        [OutputConstructor]
        private GetIAMPolicyAuditConfigsAuditLogConfigsResult(
            ImmutableArray<string> exemptedMembers,
            string logType)
        {
            ExemptedMembers = exemptedMembers;
            LogType = logType;
        }
    }

    [OutputType]
    public sealed class GetIAMPolicyAuditConfigsResult
    {
        /// <summary>
        /// A nested block that defines the operations you'd like to log.
        /// </summary>
        public readonly ImmutableArray<GetIAMPolicyAuditConfigsAuditLogConfigsResult> AuditLogConfigs;
        /// <summary>
        /// Defines a service that will be enabled for audit logging. For example, `storage.googleapis.com`, `cloudsql.googleapis.com`. `allServices` is a special value that covers all services.
        /// </summary>
        public readonly string Service;

        [OutputConstructor]
        private GetIAMPolicyAuditConfigsResult(
            ImmutableArray<GetIAMPolicyAuditConfigsAuditLogConfigsResult> auditLogConfigs,
            string service)
        {
            AuditLogConfigs = auditLogConfigs;
            Service = service;
        }
    }

    [OutputType]
    public sealed class GetIAMPolicyBindingsConditionResult
    {
        public readonly string? Description;
        public readonly string Expression;
        public readonly string Title;

        [OutputConstructor]
        private GetIAMPolicyBindingsConditionResult(
            string? description,
            string expression,
            string title)
        {
            Description = description;
            Expression = expression;
            Title = title;
        }
    }

    [OutputType]
    public sealed class GetIAMPolicyBindingsResult
    {
        public readonly GetIAMPolicyBindingsConditionResult? Condition;
        /// <summary>
        /// An array of identities that will be granted the privilege in the `role`. For more details on format and restrictions see https://cloud.google.com/billing/reference/rest/v1/Policy#Binding
        /// Each entry can have one of the following values:
        /// * **allUsers**: A special identifier that represents anyone who is on the internet; with or without a Google account. It **can't** be used with the `gcp.organizations.Project` resource.
        /// * **allAuthenticatedUsers**: A special identifier that represents anyone who is authenticated with a Google account or a service account. It **can't** be used with the `gcp.organizations.Project` resource.
        /// * **user:{emailid}**: An email address that represents a specific Google account. For example, alice@gmail.com.
        /// * **serviceAccount:{emailid}**: An email address that represents a service account. For example, my-other-app@appspot.gserviceaccount.com.
        /// * **group:{emailid}**: An email address that represents a Google group. For example, admins@example.com.
        /// * **domain:{domain}**: A G Suite domain (primary, instead of alias) name that represents all the users of that domain. For example, google.com or example.com.
        /// </summary>
        public readonly ImmutableArray<string> Members;
        /// <summary>
        /// The role/permission that will be granted to the members.
        /// See the [IAM Roles](https://cloud.google.com/compute/docs/access/iam) documentation for a complete list of roles.
        /// Note that custom roles must be of the format `[projects|organizations]/{parent-name}/roles/{role-name}`.
        /// </summary>
        public readonly string Role;

        [OutputConstructor]
        private GetIAMPolicyBindingsResult(
            GetIAMPolicyBindingsConditionResult? condition,
            ImmutableArray<string> members,
            string role)
        {
            Condition = condition;
            Members = members;
            Role = role;
        }
    }
    }
}
