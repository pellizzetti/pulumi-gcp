// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Three different resources help you manage IAM policies on dataproc jobs. Each of these resources serves a different use case:
 * 
 * * `gcp.dataproc.JobIAMPolicy`: Authoritative. Sets the IAM policy for the job and replaces any existing policy already attached.
 * * `gcp.dataproc.JobIAMBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the job are preserved.
 * * `gcp.dataproc.JobIAMMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the job are preserved.
 * 
 * > **Note:** `gcp.dataproc.JobIAMPolicy` **cannot** be used in conjunction with `gcp.dataproc.JobIAMBinding` and `gcp.dataproc.JobIAMMember` or they will fight over what your policy should be. In addition, be careful not to accidentaly unset ownership of the job as `gcp.dataproc.JobIAMPolicy` replaces the entire policy.
 * 
 * > **Note:** `gcp.dataproc.JobIAMBinding` resources **can be** used in conjunction with `gcp.dataproc.JobIAMMember` resources **only if** they do not grant privilege to the same role.
 * 
 * ## google\_pubsub\_subscription\_iam\_policy
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 * 
 * const admin = pulumi.output(gcp.organizations.getIAMPolicy({
 *     bindings: [{
 *         members: ["user:jane@example.com"],
 *         role: "roles/editor",
 *     }],
 * }));
 * const editor = new gcp.dataproc.JobIAMPolicy("editor", {
 *     jobId: "your-dataproc-job",
 *     policyData: admin.policyData,
 *     project: "your-project",
 *     region: "your-region",
 * });
 * ```
 * 
 * ## google\_pubsub\_subscription\_iam\_binding
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 * 
 * const editor = new gcp.dataproc.JobIAMBinding("editor", {
 *     jobId: "your-dataproc-job",
 *     members: ["user:jane@example.com"],
 *     role: "roles/editor",
 * });
 * ```
 * 
 * ## google\_pubsub\_subscription\_iam\_member
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 * 
 * const editor = new gcp.dataproc.JobIAMMember("editor", {
 *     jobId: "your-dataproc-job",
 *     member: "user:jane@example.com",
 *     role: "roles/editor",
 * });
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/dataproc_job_iam_policy.html.markdown.
 */
export class JobIAMPolicy extends pulumi.CustomResource {
    /**
     * Get an existing JobIAMPolicy resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: JobIAMPolicyState, opts?: pulumi.CustomResourceOptions): JobIAMPolicy {
        return new JobIAMPolicy(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:dataproc/jobIAMPolicy:JobIAMPolicy';

    /**
     * Returns true if the given object is an instance of JobIAMPolicy.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is JobIAMPolicy {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === JobIAMPolicy.__pulumiType;
    }

    /**
     * (Computed) The etag of the jobs's IAM policy.
     */
    public /*out*/ readonly etag!: pulumi.Output<string>;
    public readonly jobId!: pulumi.Output<string>;
    /**
     * The policy data generated by a `gcp.organizations.getIAMPolicy` data source.
     */
    public readonly policyData!: pulumi.Output<string>;
    /**
     * The project in which the job belongs. If it
     * is not provided, this provider will use the provider default.
     */
    public readonly project!: pulumi.Output<string>;
    /**
     * The region in which the job belongs. If it
     * is not provided, this provider will use the provider default.
     */
    public readonly region!: pulumi.Output<string>;

    /**
     * Create a JobIAMPolicy resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: JobIAMPolicyArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: JobIAMPolicyArgs | JobIAMPolicyState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as JobIAMPolicyState | undefined;
            inputs["etag"] = state ? state.etag : undefined;
            inputs["jobId"] = state ? state.jobId : undefined;
            inputs["policyData"] = state ? state.policyData : undefined;
            inputs["project"] = state ? state.project : undefined;
            inputs["region"] = state ? state.region : undefined;
        } else {
            const args = argsOrState as JobIAMPolicyArgs | undefined;
            if (!args || args.jobId === undefined) {
                throw new Error("Missing required property 'jobId'");
            }
            if (!args || args.policyData === undefined) {
                throw new Error("Missing required property 'policyData'");
            }
            inputs["jobId"] = args ? args.jobId : undefined;
            inputs["policyData"] = args ? args.policyData : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["region"] = args ? args.region : undefined;
            inputs["etag"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(JobIAMPolicy.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering JobIAMPolicy resources.
 */
export interface JobIAMPolicyState {
    /**
     * (Computed) The etag of the jobs's IAM policy.
     */
    readonly etag?: pulumi.Input<string>;
    readonly jobId?: pulumi.Input<string>;
    /**
     * The policy data generated by a `gcp.organizations.getIAMPolicy` data source.
     */
    readonly policyData?: pulumi.Input<string>;
    /**
     * The project in which the job belongs. If it
     * is not provided, this provider will use the provider default.
     */
    readonly project?: pulumi.Input<string>;
    /**
     * The region in which the job belongs. If it
     * is not provided, this provider will use the provider default.
     */
    readonly region?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a JobIAMPolicy resource.
 */
export interface JobIAMPolicyArgs {
    readonly jobId: pulumi.Input<string>;
    /**
     * The policy data generated by a `gcp.organizations.getIAMPolicy` data source.
     */
    readonly policyData: pulumi.Input<string>;
    /**
     * The project in which the job belongs. If it
     * is not provided, this provider will use the provider default.
     */
    readonly project?: pulumi.Input<string>;
    /**
     * The region in which the job belongs. If it
     * is not provided, this provider will use the provider default.
     */
    readonly region?: pulumi.Input<string>;
}
