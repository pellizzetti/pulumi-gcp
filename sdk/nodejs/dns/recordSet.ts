// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputApi from "../types/input";
import * as outputApi from "../types/output";
import * as utilities from "../utilities";

/**
 * Manages a set of DNS records within Google Cloud DNS. For more information see [the official documentation](https://cloud.google.com/dns/records/) and
 * [API](https://cloud.google.com/dns/api/v1/resourceRecordSets).
 * 
 * > **Note:** The provider treats this resource as an authoritative record set. This means existing records (including the default records) for the given type will be overwritten when you create this resource with this provider. In addition, the Google Cloud DNS API requires NS records to be present at all times, so this provider will not actually remove NS records during destroy but will report that it did.
 * 
 * ## Example Usage
 * 
 * ### Binding a DNS name to the ephemeral IP of a new instance:
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 * 
 * const frontendInstance = new gcp.compute.Instance("frontend", {
 *     bootDisk: {
 *         initializeParams: {
 *             image: "debian-cloud/debian-9",
 *         },
 *     },
 *     machineType: "g1-small",
 *     networkInterfaces: [{
 *         accessConfigs: [{}],
 *         network: "default",
 *     }],
 *     zone: "us-central1-b",
 * });
 * const prod = new gcp.dns.ManagedZone("prod", {
 *     dnsName: "prod.mydomain.com.",
 * });
 * const frontendRecordSet = new gcp.dns.RecordSet("frontend", {
 *     managedZone: prod.name,
 *     rrdatas: [frontendInstance.networkInterfaces[0].accessConfig.0.natIp],
 *     ttl: 300,
 *     type: "A",
 * });
 * ```
 * 
 * ### Adding an A record
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 * 
 * const prod = new gcp.dns.ManagedZone("prod", {
 *     dnsName: "prod.mydomain.com.",
 * });
 * const recordSet = new gcp.dns.RecordSet("a", {
 *     managedZone: prod.name,
 *     rrdatas: ["8.8.8.8"],
 *     ttl: 300,
 *     type: "A",
 * });
 * ```
 * 
 * ### Adding an MX record
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 * 
 * const prod = new gcp.dns.ManagedZone("prod", {
 *     dnsName: "prod.mydomain.com.",
 * });
 * const mx = new gcp.dns.RecordSet("mx", {
 *     managedZone: prod.name,
 *     rrdatas: [
 *         "1 aspmx.l.google.com.",
 *         "5 alt1.aspmx.l.google.com.",
 *         "5 alt2.aspmx.l.google.com.",
 *         "10 alt3.aspmx.l.google.com.",
 *         "10 alt4.aspmx.l.google.com.",
 *     ],
 *     ttl: 3600,
 *     type: "MX",
 * });
 * ```
 * 
 * ### Adding an SPF record
 * 
 * Quotes (`""`) must be added around your `rrdatas` for a SPF record. Otherwise `rrdatas` string gets split on spaces.
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 * 
 * const prod = new gcp.dns.ManagedZone("prod", {
 *     dnsName: "prod.mydomain.com.",
 * });
 * const spf = new gcp.dns.RecordSet("spf", {
 *     managedZone: prod.name,
 *     rrdatas: ["\"v=spf1 ip4:111.111.111.111 include:backoff.email-example.com -all\""],
 *     ttl: 300,
 *     type: "TXT",
 * });
 * ```
 * 
 * ### Adding a CNAME record
 * 
 *  The list of `rrdatas` should only contain a single string corresponding to the Canonical Name intended.
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 * 
 * const prod = new gcp.dns.ManagedZone("prod", {
 *     dnsName: "prod.mydomain.com.",
 * });
 * const cname = new gcp.dns.RecordSet("cname", {
 *     managedZone: prod.name,
 *     rrdatas: ["frontend.mydomain.com."],
 *     ttl: 300,
 *     type: "CNAME",
 * });
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/dns_record_set.html.markdown.
 */
export class RecordSet extends pulumi.CustomResource {
    /**
     * Get an existing RecordSet resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: RecordSetState, opts?: pulumi.CustomResourceOptions): RecordSet {
        return new RecordSet(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:dns/recordSet:RecordSet';

    /**
     * Returns true if the given object is an instance of RecordSet.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is RecordSet {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === RecordSet.__pulumiType;
    }

    /**
     * The name of the zone in which this record set will
     * reside.
     */
    public readonly managedZone!: pulumi.Output<string>;
    /**
     * The DNS name this record set will apply to.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The ID of the project in which the resource belongs. If it
     * is not provided, the provider project is used.
     */
    public readonly project!: pulumi.Output<string>;
    /**
     * The string data for the records in this record set
     * whose meaning depends on the DNS type. For TXT record, if the string data contains spaces, add surrounding `\"` if you don't want your string to get split on spaces. To specify a single record value longer than 255 characters such as a TXT record for DKIM, add `\"\"` inside this provider's configuration string (e.g. `"first255characters\"\"morecharacters"`).
     */
    public readonly rrdatas!: pulumi.Output<string[]>;
    /**
     * The time-to-live of this record set (seconds).
     */
    public readonly ttl!: pulumi.Output<number>;
    /**
     * The DNS record set type.
     */
    public readonly type!: pulumi.Output<string>;

    /**
     * Create a RecordSet resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: RecordSetArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: RecordSetArgs | RecordSetState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as RecordSetState | undefined;
            inputs["managedZone"] = state ? state.managedZone : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["project"] = state ? state.project : undefined;
            inputs["rrdatas"] = state ? state.rrdatas : undefined;
            inputs["ttl"] = state ? state.ttl : undefined;
            inputs["type"] = state ? state.type : undefined;
        } else {
            const args = argsOrState as RecordSetArgs | undefined;
            if (!args || args.managedZone === undefined) {
                throw new Error("Missing required property 'managedZone'");
            }
            if (!args || args.rrdatas === undefined) {
                throw new Error("Missing required property 'rrdatas'");
            }
            if (!args || args.ttl === undefined) {
                throw new Error("Missing required property 'ttl'");
            }
            if (!args || args.type === undefined) {
                throw new Error("Missing required property 'type'");
            }
            inputs["managedZone"] = args ? args.managedZone : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["rrdatas"] = args ? args.rrdatas : undefined;
            inputs["ttl"] = args ? args.ttl : undefined;
            inputs["type"] = args ? args.type : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(RecordSet.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering RecordSet resources.
 */
export interface RecordSetState {
    /**
     * The name of the zone in which this record set will
     * reside.
     */
    readonly managedZone?: pulumi.Input<string>;
    /**
     * The DNS name this record set will apply to.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs. If it
     * is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
    /**
     * The string data for the records in this record set
     * whose meaning depends on the DNS type. For TXT record, if the string data contains spaces, add surrounding `\"` if you don't want your string to get split on spaces. To specify a single record value longer than 255 characters such as a TXT record for DKIM, add `\"\"` inside this provider's configuration string (e.g. `"first255characters\"\"morecharacters"`).
     */
    readonly rrdatas?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The time-to-live of this record set (seconds).
     */
    readonly ttl?: pulumi.Input<number>;
    /**
     * The DNS record set type.
     */
    readonly type?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a RecordSet resource.
 */
export interface RecordSetArgs {
    /**
     * The name of the zone in which this record set will
     * reside.
     */
    readonly managedZone: pulumi.Input<string>;
    /**
     * The DNS name this record set will apply to.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs. If it
     * is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
    /**
     * The string data for the records in this record set
     * whose meaning depends on the DNS type. For TXT record, if the string data contains spaces, add surrounding `\"` if you don't want your string to get split on spaces. To specify a single record value longer than 255 characters such as a TXT record for DKIM, add `\"\"` inside this provider's configuration string (e.g. `"first255characters\"\"morecharacters"`).
     */
    readonly rrdatas: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The time-to-live of this record set (seconds).
     */
    readonly ttl: pulumi.Input<number>;
    /**
     * The DNS record set type.
     */
    readonly type: pulumi.Input<string>;
}
