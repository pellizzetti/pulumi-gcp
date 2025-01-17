// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/iap_tunnel_instance_iam_policy.html.markdown.
 */
export class TunnelInstanceIAMPolicy extends pulumi.CustomResource {
    /**
     * Get an existing TunnelInstanceIAMPolicy resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: TunnelInstanceIAMPolicyState, opts?: pulumi.CustomResourceOptions): TunnelInstanceIAMPolicy {
        return new TunnelInstanceIAMPolicy(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:iap/tunnelInstanceIAMPolicy:TunnelInstanceIAMPolicy';

    /**
     * Returns true if the given object is an instance of TunnelInstanceIAMPolicy.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is TunnelInstanceIAMPolicy {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === TunnelInstanceIAMPolicy.__pulumiType;
    }

    /**
     * (Computed) The etag of the instance's IAM policy.
     */
    public /*out*/ readonly etag!: pulumi.Output<string>;
    /**
     * The name of the instance.
     */
    public readonly instance!: pulumi.Output<string>;
    /**
     * The policy data generated by
     * a `gcp.organizations.getIAMPolicy` data source.
     */
    public readonly policyData!: pulumi.Output<string>;
    /**
     * The ID of the project in which the resource belongs. If it
     * is not provided, the provider project is used.
     */
    public readonly project!: pulumi.Output<string>;
    /**
     * The zone of the instance. If
     * unspecified, this defaults to the zone configured in the provider.
     */
    public readonly zone!: pulumi.Output<string>;

    /**
     * Create a TunnelInstanceIAMPolicy resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: TunnelInstanceIAMPolicyArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: TunnelInstanceIAMPolicyArgs | TunnelInstanceIAMPolicyState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as TunnelInstanceIAMPolicyState | undefined;
            inputs["etag"] = state ? state.etag : undefined;
            inputs["instance"] = state ? state.instance : undefined;
            inputs["policyData"] = state ? state.policyData : undefined;
            inputs["project"] = state ? state.project : undefined;
            inputs["zone"] = state ? state.zone : undefined;
        } else {
            const args = argsOrState as TunnelInstanceIAMPolicyArgs | undefined;
            if (!args || args.instance === undefined) {
                throw new Error("Missing required property 'instance'");
            }
            if (!args || args.policyData === undefined) {
                throw new Error("Missing required property 'policyData'");
            }
            inputs["instance"] = args ? args.instance : undefined;
            inputs["policyData"] = args ? args.policyData : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["zone"] = args ? args.zone : undefined;
            inputs["etag"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(TunnelInstanceIAMPolicy.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering TunnelInstanceIAMPolicy resources.
 */
export interface TunnelInstanceIAMPolicyState {
    /**
     * (Computed) The etag of the instance's IAM policy.
     */
    readonly etag?: pulumi.Input<string>;
    /**
     * The name of the instance.
     */
    readonly instance?: pulumi.Input<string>;
    /**
     * The policy data generated by
     * a `gcp.organizations.getIAMPolicy` data source.
     */
    readonly policyData?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs. If it
     * is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
    /**
     * The zone of the instance. If
     * unspecified, this defaults to the zone configured in the provider.
     */
    readonly zone?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a TunnelInstanceIAMPolicy resource.
 */
export interface TunnelInstanceIAMPolicyArgs {
    /**
     * The name of the instance.
     */
    readonly instance: pulumi.Input<string>;
    /**
     * The policy data generated by
     * a `gcp.organizations.getIAMPolicy` data source.
     */
    readonly policyData: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs. If it
     * is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
    /**
     * The zone of the instance. If
     * unspecified, this defaults to the zone configured in the provider.
     */
    readonly zone?: pulumi.Input<string>;
}
