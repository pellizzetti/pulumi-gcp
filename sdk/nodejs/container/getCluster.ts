// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Get info about a GKE cluster from its name and location.
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/d/container_cluster.html.markdown.
 */
export function getCluster(args: GetClusterArgs, opts?: pulumi.InvokeOptions): Promise<GetClusterResult> & GetClusterResult {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    const promise: Promise<GetClusterResult> = pulumi.runtime.invoke("gcp:container/getCluster:getCluster", {
        "location": args.location,
        "name": args.name,
        "project": args.project,
        "region": args.region,
        "zone": args.zone,
    }, opts);

    return pulumi.utils.liftProperties(promise, opts);
}

/**
 * A collection of arguments for invoking getCluster.
 */
export interface GetClusterArgs {
    /**
     * The location (zone or region) this cluster has been
     * created in. One of `location`, `region`, `zone`, or a provider-level `zone` must
     * be specified.
     */
    readonly location?: string;
    /**
     * The name of the cluster.
     */
    readonly name: string;
    /**
     * The project in which the resource belongs. If it
     * is not provided, the provider project is used.
     */
    readonly project?: string;
    /**
     * The region this cluster has been created in. Deprecated
     * in favour of `location`.
     */
    readonly region?: string;
    /**
     * The zone this cluster has been created in. Deprecated in
     * favour of `location`.
     */
    readonly zone?: string;
}

/**
 * A collection of values returned by getCluster.
 */
export interface GetClusterResult {
    readonly additionalZones: string[];
    readonly addonsConfigs: outputs.container.GetClusterAddonsConfig[];
    readonly authenticatorGroupsConfigs: outputs.container.GetClusterAuthenticatorGroupsConfig[];
    readonly clusterAutoscalings: outputs.container.GetClusterClusterAutoscaling[];
    readonly clusterIpv4Cidr: string;
    readonly databaseEncryptions: outputs.container.GetClusterDatabaseEncryption[];
    readonly defaultMaxPodsPerNode: number;
    readonly description: string;
    readonly enableBinaryAuthorization: boolean;
    readonly enableIntranodeVisibility: boolean;
    readonly enableKubernetesAlpha: boolean;
    readonly enableLegacyAbac: boolean;
    readonly enableShieldedNodes: boolean;
    readonly enableTpu: boolean;
    readonly endpoint: string;
    readonly initialNodeCount: number;
    readonly instanceGroupUrls: string[];
    readonly ipAllocationPolicies: outputs.container.GetClusterIpAllocationPolicy[];
    readonly location?: string;
    readonly loggingService: string;
    readonly maintenancePolicies: outputs.container.GetClusterMaintenancePolicy[];
    readonly masterAuths: outputs.container.GetClusterMasterAuth[];
    readonly masterAuthorizedNetworksConfigs: outputs.container.GetClusterMasterAuthorizedNetworksConfig[];
    readonly masterVersion: string;
    readonly minMasterVersion: string;
    readonly monitoringService: string;
    readonly name: string;
    readonly network: string;
    readonly networkPolicies: outputs.container.GetClusterNetworkPolicy[];
    readonly nodeConfigs: outputs.container.GetClusterNodeConfig[];
    readonly nodeLocations: string[];
    readonly nodePools: outputs.container.GetClusterNodePool[];
    readonly nodeVersion: string;
    readonly podSecurityPolicyConfigs: outputs.container.GetClusterPodSecurityPolicyConfig[];
    readonly privateClusterConfigs: outputs.container.GetClusterPrivateClusterConfig[];
    readonly project?: string;
    readonly region?: string;
    readonly releaseChannels: outputs.container.GetClusterReleaseChannel[];
    readonly removeDefaultNodePool: boolean;
    readonly resourceLabels: {[key: string]: string};
    readonly resourceUsageExportConfigs: outputs.container.GetClusterResourceUsageExportConfig[];
    readonly servicesIpv4Cidr: string;
    readonly subnetwork: string;
    readonly tpuIpv4CidrBlock: string;
    readonly verticalPodAutoscalings: outputs.container.GetClusterVerticalPodAutoscaling[];
    readonly workloadIdentityConfigs: outputs.container.GetClusterWorkloadIdentityConfig[];
    readonly zone?: string;
    /**
     * id is the provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}
