// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Three different resources help you manage your IAM policy for GCE instance. Each of these resources serves a different use case:
 * 
 * * `google_compute_instance_iam_policy`: Authoritative. Sets the IAM policy for the instance and replaces any existing policy already attached.
 * * `google_compute_instance_iam_binding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the instance are preserved.
 * * `google_compute_instance_iam_member`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the instance are preserved.
 * 
 * > **Note:** `google_compute_instance_iam_policy` **cannot** be used in conjunction with `google_compute_instance_iam_binding` and `google_compute_instance_iam_member` or they will fight over what your policy should be.
 * 
 * > **Note:** `google_compute_instance_iam_binding` resources **can be** used in conjunction with `google_compute_instance_iam_member` resources **only if** they do not grant privilege to the same role.
 * 
 * ## google\_compute\_instance\_iam\_policy
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 * 
 * const admin = pulumi.output(gcp.organizations.getIAMPolicy({
 *     bindings: [{
 *         members: ["user:jane@example.com"],
 *         role: "roles/compute.osLogin",
 *     }],
 * }));
 * const instance = new gcp.compute.InstanceIAMPolicy("instance", {
 *     instanceName: "your-instance-name",
 *     policyData: admin.policyData,
 * });
 * ```
 * 
 * ## google\_compute\_instance\_iam\_binding
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 * 
 * const instance = new gcp.compute.InstanceIAMBinding("instance", {
 *     instanceName: "your-instance-name",
 *     members: ["user:jane@example.com"],
 *     role: "roles/compute.osLogin",
 * });
 * ```
 * 
 * ## google\compute\_instance\_iam\_member
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 * 
 * const instance = new gcp.compute.InstanceIAMMember("instance", {
 *     instanceName: "your-instance-name",
 *     member: "user:jane@example.com",
 *     role: "roles/compute.osLogin",
 * });
 * ```
 */
export class InstanceIAMMember extends pulumi.CustomResource {
    /**
     * Get an existing InstanceIAMMember resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: InstanceIAMMemberState, opts?: pulumi.CustomResourceOptions): InstanceIAMMember {
        return new InstanceIAMMember(name, <any>state, { ...opts, id: id });
    }

    /**
     * (Computed) The etag of the instance's IAM policy.
     */
    public /*out*/ readonly etag!: pulumi.Output<string>;
    /**
     * The name of the instance.
     */
    public readonly instanceName!: pulumi.Output<string>;
    public readonly member!: pulumi.Output<string>;
    /**
     * The ID of the project in which the resource belongs. If it
     * is not provided, the provider project is used.
     */
    public readonly project!: pulumi.Output<string>;
    /**
     * The role that should be applied. Only one
     * `google_compute_instance_iam_binding` can be used per role. Note that custom roles must be of the format
     * `[projects|organizations]/{parent-name}/roles/{role-name}`.
     */
    public readonly role!: pulumi.Output<string>;
    /**
     * The zone of the instance. If
     * unspecified, this defaults to the zone configured in the provider.
     */
    public readonly zone!: pulumi.Output<string>;

    /**
     * Create a InstanceIAMMember resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: InstanceIAMMemberArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: InstanceIAMMemberArgs | InstanceIAMMemberState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as InstanceIAMMemberState | undefined;
            inputs["etag"] = state ? state.etag : undefined;
            inputs["instanceName"] = state ? state.instanceName : undefined;
            inputs["member"] = state ? state.member : undefined;
            inputs["project"] = state ? state.project : undefined;
            inputs["role"] = state ? state.role : undefined;
            inputs["zone"] = state ? state.zone : undefined;
        } else {
            const args = argsOrState as InstanceIAMMemberArgs | undefined;
            if (!args || args.instanceName === undefined) {
                throw new Error("Missing required property 'instanceName'");
            }
            if (!args || args.member === undefined) {
                throw new Error("Missing required property 'member'");
            }
            if (!args || args.role === undefined) {
                throw new Error("Missing required property 'role'");
            }
            inputs["instanceName"] = args ? args.instanceName : undefined;
            inputs["member"] = args ? args.member : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["role"] = args ? args.role : undefined;
            inputs["zone"] = args ? args.zone : undefined;
            inputs["etag"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super("gcp:compute/instanceIAMMember:InstanceIAMMember", name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering InstanceIAMMember resources.
 */
export interface InstanceIAMMemberState {
    /**
     * (Computed) The etag of the instance's IAM policy.
     */
    readonly etag?: pulumi.Input<string>;
    /**
     * The name of the instance.
     */
    readonly instanceName?: pulumi.Input<string>;
    readonly member?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs. If it
     * is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
    /**
     * The role that should be applied. Only one
     * `google_compute_instance_iam_binding` can be used per role. Note that custom roles must be of the format
     * `[projects|organizations]/{parent-name}/roles/{role-name}`.
     */
    readonly role?: pulumi.Input<string>;
    /**
     * The zone of the instance. If
     * unspecified, this defaults to the zone configured in the provider.
     */
    readonly zone?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a InstanceIAMMember resource.
 */
export interface InstanceIAMMemberArgs {
    /**
     * The name of the instance.
     */
    readonly instanceName: pulumi.Input<string>;
    readonly member: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs. If it
     * is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
    /**
     * The role that should be applied. Only one
     * `google_compute_instance_iam_binding` can be used per role. Note that custom roles must be of the format
     * `[projects|organizations]/{parent-name}/roles/{role-name}`.
     */
    readonly role: pulumi.Input<string>;
    /**
     * The zone of the instance. If
     * unspecified, this defaults to the zone configured in the provider.
     */
    readonly zone?: pulumi.Input<string>;
}
