// Copyright 2020 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    AUTO GENERATED CODE     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Config Connector and manual
//     changes will be clobbered when the file is regenerated.
//
// ----------------------------------------------------------------------------

// *** DISCLAIMER ***
// Config Connector's go-client for CRDs is currently in ALPHA, which means
// that future versions of the go-client may include breaking changes.
// Please try it out and give us feedback!

package v1beta1

import (
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type ClusterAddonsConfig struct {
	/* The status of the CloudRun addon. It is disabled by default. Set disabled = false to enable. */
	// +optional
	CloudrunConfig *ClusterCloudrunConfig `json:"cloudrunConfig,omitempty"`

	/* The of the Config Connector addon. */
	// +optional
	ConfigConnectorConfig *ClusterConfigConnectorConfig `json:"configConnectorConfig,omitempty"`

	/* The status of the NodeLocal DNSCache addon. It is disabled by default. Set enabled = true to enable. */
	// +optional
	DnsCacheConfig *ClusterDnsCacheConfig `json:"dnsCacheConfig,omitempty"`

	/* Whether this cluster should enable the Google Compute Engine Persistent Disk Container Storage Interface (CSI) Driver. Defaults to disabled; set enabled = true to enable. */
	// +optional
	GcePersistentDiskCsiDriverConfig *ClusterGcePersistentDiskCsiDriverConfig `json:"gcePersistentDiskCsiDriverConfig,omitempty"`

	/* The status of the Horizontal Pod Autoscaling addon, which increases or decreases the number of replica pods a replication controller has based on the resource usage of the existing pods. It ensures that a Heapster pod is running in the cluster, which is also used by the Cloud Monitoring service. It is enabled by default; set disabled = true to disable. */
	// +optional
	HorizontalPodAutoscaling *ClusterHorizontalPodAutoscaling `json:"horizontalPodAutoscaling,omitempty"`

	/* The status of the HTTP (L7) load balancing controller addon, which makes it easy to set up HTTP load balancers for services in a cluster. It is enabled by default; set disabled = true to disable. */
	// +optional
	HttpLoadBalancing *ClusterHttpLoadBalancing `json:"httpLoadBalancing,omitempty"`

	/* The status of the Istio addon. */
	// +optional
	IstioConfig *ClusterIstioConfig `json:"istioConfig,omitempty"`

	/* Configuration for the KALM addon, which manages the lifecycle of k8s. It is disabled by default; Set enabled = true to enable. */
	// +optional
	KalmConfig *ClusterKalmConfig `json:"kalmConfig,omitempty"`

	/* Whether we should enable the network policy addon for the master. This must be enabled in order to enable network policy for the nodes. To enable this, you must also define a network_policy block, otherwise nothing will happen. It can only be disabled if the nodes already do not have network policies enabled. Defaults to disabled; set disabled = false to enable. */
	// +optional
	NetworkPolicyConfig *ClusterNetworkPolicyConfig `json:"networkPolicyConfig,omitempty"`
}

type ClusterAuthenticatorGroupsConfig struct {
	/* Immutable. The name of the RBAC security group for use with Google security groups in Kubernetes RBAC. Group name must be in format gke-security-groups@yourdomain.com. */
	SecurityGroup string `json:"securityGroup"`
}

type ClusterAutoProvisioningDefaults struct {
	/* Minimum CPU platform to be used by this instance. The instance may be scheduled on the specified or newer CPU platform. Applicable values are the friendly names of CPU platforms, such as Intel Haswell. */
	// +optional
	MinCpuPlatform *string `json:"minCpuPlatform,omitempty"`

	/* Scopes that are used by NAP when creating node pools. */
	// +optional
	OauthScopes []string `json:"oauthScopes,omitempty"`

	/*  */
	// +optional
	ServiceAccountRef *v1alpha1.ResourceRef `json:"serviceAccountRef,omitempty"`
}

type ClusterBigqueryDestination struct {
	/* The ID of a BigQuery Dataset. */
	DatasetId string `json:"datasetId"`
}

type ClusterCidrBlocks struct {
	/* External network that can access Kubernetes master through HTTPS. Must be specified in CIDR notation. */
	CidrBlock string `json:"cidrBlock"`

	/* Field for users to identify CIDR blocks. */
	// +optional
	DisplayName *string `json:"displayName,omitempty"`
}

type ClusterClientCertificateConfig struct {
	/* Immutable. Whether client certificate authorization is enabled for this cluster. */
	IssueClientCertificate bool `json:"issueClientCertificate"`
}

type ClusterCloudrunConfig struct {
	/*  */
	Disabled bool `json:"disabled"`

	/*  */
	// +optional
	LoadBalancerType *string `json:"loadBalancerType,omitempty"`
}

type ClusterClusterAutoscaling struct {
	/* Contains defaults for a node pool created by NAP. */
	// +optional
	AutoProvisioningDefaults *ClusterAutoProvisioningDefaults `json:"autoProvisioningDefaults,omitempty"`

	/* Configuration options for the Autoscaling profile feature, which lets you choose whether the cluster autoscaler should optimize for resource utilization or resource availability when deciding to remove nodes from a cluster. Can be BALANCED or OPTIMIZE_UTILIZATION. Defaults to BALANCED. */
	// +optional
	AutoscalingProfile *string `json:"autoscalingProfile,omitempty"`

	/* Whether node auto-provisioning is enabled. Resource limits for cpu and memory must be defined to enable node auto-provisioning. */
	Enabled bool `json:"enabled"`

	/* Global constraints for machine resources in the cluster. Configuring the cpu and memory types is required if node auto-provisioning is enabled. These limits will apply to node pool autoscaling in addition to node auto-provisioning. */
	// +optional
	ResourceLimits []ClusterResourceLimits `json:"resourceLimits,omitempty"`
}

type ClusterClusterTelemetry struct {
	/* Type of the integration. */
	Type string `json:"type"`
}

type ClusterConfidentialNodes struct {
	/* Immutable. Whether Confidential Nodes feature is enabled for all nodes in this cluster. */
	Enabled bool `json:"enabled"`
}

type ClusterConfigConnectorConfig struct {
	/*  */
	Enabled bool `json:"enabled"`
}

type ClusterDailyMaintenanceWindow struct {
	/*  */
	// +optional
	Duration *string `json:"duration,omitempty"`

	/*  */
	StartTime string `json:"startTime"`
}

type ClusterDatabaseEncryption struct {
	/* The key to use to encrypt/decrypt secrets. */
	// +optional
	KeyName *string `json:"keyName,omitempty"`

	/* ENCRYPTED or DECRYPTED. */
	State string `json:"state"`
}

type ClusterDefaultSnatStatus struct {
	/* When disabled is set to false, default IP masquerade rules will be applied to the nodes to prevent sNAT on cluster internal traffic. */
	Disabled bool `json:"disabled"`
}

type ClusterDnsCacheConfig struct {
	/*  */
	Enabled bool `json:"enabled"`
}

type ClusterEphemeralStorageConfig struct {
	/* Immutable. Number of local SSDs to use to back ephemeral storage. Uses NVMe interfaces. Each local SSD is 375 GB in size. */
	LocalSsdCount int `json:"localSsdCount"`
}

type ClusterGcePersistentDiskCsiDriverConfig struct {
	/*  */
	Enabled bool `json:"enabled"`
}

type ClusterGuestAccelerator struct {
	/* Immutable. The number of the accelerator cards exposed to an instance. */
	Count int `json:"count"`

	/* Immutable. The accelerator type resource name. */
	Type string `json:"type"`
}

type ClusterHorizontalPodAutoscaling struct {
	/*  */
	Disabled bool `json:"disabled"`
}

type ClusterHttpLoadBalancing struct {
	/*  */
	Disabled bool `json:"disabled"`
}

type ClusterIpAllocationPolicy struct {
	/* Immutable. The IP address range for the cluster pod IPs. Set to blank to have a range chosen with the default size. Set to /netmask (e.g. /14) to have a range chosen with a specific netmask. Set to a CIDR notation (e.g. 10.96.0.0/14) from the RFC-1918 private networks (e.g. 10.0.0.0/8, 172.16.0.0/12, 192.168.0.0/16) to pick a specific range to use. */
	// +optional
	ClusterIpv4CidrBlock *string `json:"clusterIpv4CidrBlock,omitempty"`

	/* Immutable. The name of the existing secondary range in the cluster's subnetwork to use for pod IP addresses. Alternatively, cluster_ipv4_cidr_block can be used to automatically create a GKE-managed one. */
	// +optional
	ClusterSecondaryRangeName *string `json:"clusterSecondaryRangeName,omitempty"`

	/* Immutable. The IP address range of the services IPs in this cluster. Set to blank to have a range chosen with the default size. Set to /netmask (e.g. /14) to have a range chosen with a specific netmask. Set to a CIDR notation (e.g. 10.96.0.0/14) from the RFC-1918 private networks (e.g. 10.0.0.0/8, 172.16.0.0/12, 192.168.0.0/16) to pick a specific range to use. */
	// +optional
	ServicesIpv4CidrBlock *string `json:"servicesIpv4CidrBlock,omitempty"`

	/* Immutable. The name of the existing secondary range in the cluster's subnetwork to use for service ClusterIPs. Alternatively, services_ipv4_cidr_block can be used to automatically create a GKE-managed one. */
	// +optional
	ServicesSecondaryRangeName *string `json:"servicesSecondaryRangeName,omitempty"`
}

type ClusterIstioConfig struct {
	/* The authentication type between services in Istio. Available options include AUTH_MUTUAL_TLS. */
	// +optional
	Auth *string `json:"auth,omitempty"`

	/* The status of the Istio addon, which makes it easy to set up Istio for services in a cluster. It is disabled by default. Set disabled = false to enable. */
	Disabled bool `json:"disabled"`
}

type ClusterKalmConfig struct {
	/*  */
	Enabled bool `json:"enabled"`
}

type ClusterKubeletConfig struct {
	/* Enable CPU CFS quota enforcement for containers that specify CPU limits. */
	// +optional
	CpuCfsQuota *bool `json:"cpuCfsQuota,omitempty"`

	/* Set the CPU CFS quota period value 'cpu.cfs_period_us'. */
	// +optional
	CpuCfsQuotaPeriod *string `json:"cpuCfsQuotaPeriod,omitempty"`

	/* Control the CPU management policy on the node. */
	CpuManagerPolicy string `json:"cpuManagerPolicy"`
}

type ClusterLinuxNodeConfig struct {
	/* The Linux kernel parameters to be applied to the nodes and all pods running on the nodes. */
	Sysctls map[string]string `json:"sysctls"`
}

type ClusterMaintenanceExclusion struct {
	/*  */
	EndTime string `json:"endTime"`

	/*  */
	ExclusionName string `json:"exclusionName"`

	/*  */
	StartTime string `json:"startTime"`
}

type ClusterMaintenancePolicy struct {
	/* Time window specified for daily maintenance operations. Specify start_time in RFC3339 format "HH:MM”, where HH : [00-23] and MM : [00-59] GMT. */
	// +optional
	DailyMaintenanceWindow *ClusterDailyMaintenanceWindow `json:"dailyMaintenanceWindow,omitempty"`

	/* Exceptions to maintenance window. Non-emergency maintenance should not occur in these windows. */
	// +optional
	MaintenanceExclusion []ClusterMaintenanceExclusion `json:"maintenanceExclusion,omitempty"`

	/* Time window for recurring maintenance operations. */
	// +optional
	RecurringWindow *ClusterRecurringWindow `json:"recurringWindow,omitempty"`
}

type ClusterMasterAuth struct {
	/* Base64 encoded public certificate used by clients to authenticate to the cluster endpoint. */
	// +optional
	ClientCertificate *string `json:"clientCertificate,omitempty"`

	/* Immutable. Whether client certificate authorization is enabled for this cluster. */
	// +optional
	ClientCertificateConfig *ClusterClientCertificateConfig `json:"clientCertificateConfig,omitempty"`

	/* Base64 encoded private key used by clients to authenticate to the cluster endpoint. */
	// +optional
	ClientKey *string `json:"clientKey,omitempty"`

	/* Base64 encoded public certificate that is the root of trust for the cluster. */
	// +optional
	ClusterCaCertificate *string `json:"clusterCaCertificate,omitempty"`

	/* The password to use for HTTP basic authentication when accessing the Kubernetes master endpoint. */
	// +optional
	Password *ClusterPassword `json:"password,omitempty"`

	/* The username to use for HTTP basic authentication when accessing the Kubernetes master endpoint. If not present basic auth will be disabled. */
	// +optional
	Username *string `json:"username,omitempty"`
}

type ClusterMasterAuthorizedNetworksConfig struct {
	/* External networks that can access the Kubernetes cluster master through HTTPS. */
	// +optional
	CidrBlocks []ClusterCidrBlocks `json:"cidrBlocks,omitempty"`
}

type ClusterMasterGlobalAccessConfig struct {
	/* Whether the cluster master is accessible globally or not. */
	Enabled bool `json:"enabled"`
}

type ClusterNetworkPolicy struct {
	/* Whether network policy is enabled on the cluster. */
	Enabled bool `json:"enabled"`

	/* The selected network policy provider. Defaults to PROVIDER_UNSPECIFIED. */
	// +optional
	Provider *string `json:"provider,omitempty"`
}

type ClusterNetworkPolicyConfig struct {
	/*  */
	Disabled bool `json:"disabled"`
}

type ClusterNodeConfig struct {
	/*  */
	// +optional
	BootDiskKMSCryptoKeyRef *v1alpha1.ResourceRef `json:"bootDiskKMSCryptoKeyRef,omitempty"`

	/* Immutable. Size of the disk attached to each node, specified in GB. The smallest allowed disk size is 10GB. */
	// +optional
	DiskSizeGb *int `json:"diskSizeGb,omitempty"`

	/* Immutable. Type of the disk attached to each node. */
	// +optional
	DiskType *string `json:"diskType,omitempty"`

	/* Immutable. Parameters for the ephemeral storage filesystem. */
	// +optional
	EphemeralStorageConfig *ClusterEphemeralStorageConfig `json:"ephemeralStorageConfig,omitempty"`

	/* Immutable. List of the type and count of accelerator cards attached to the instance. */
	// +optional
	GuestAccelerator []ClusterGuestAccelerator `json:"guestAccelerator,omitempty"`

	/* The image type to use for this node. Note that for a given image type, the latest version of it will be used. */
	// +optional
	ImageType *string `json:"imageType,omitempty"`

	/* Node kubelet configs. */
	// +optional
	KubeletConfig *ClusterKubeletConfig `json:"kubeletConfig,omitempty"`

	/* Immutable. The map of Kubernetes labels (key/value pairs) to be applied to each node. These will added in addition to any default label(s) that Kubernetes may apply to the node. */
	// +optional
	Labels map[string]string `json:"labels,omitempty"`

	/* Parameters that can be configured on Linux nodes. */
	// +optional
	LinuxNodeConfig *ClusterLinuxNodeConfig `json:"linuxNodeConfig,omitempty"`

	/* Immutable. The number of local SSD disks to be attached to the node. */
	// +optional
	LocalSsdCount *int `json:"localSsdCount,omitempty"`

	/* Immutable. The name of a Google Compute Engine machine type. */
	// +optional
	MachineType *string `json:"machineType,omitempty"`

	/* Immutable. The metadata key/value pairs assigned to instances in the cluster. */
	// +optional
	Metadata map[string]string `json:"metadata,omitempty"`

	/* Immutable. Minimum CPU platform to be used by this instance. The instance may be scheduled on the specified or newer CPU platform. */
	// +optional
	MinCpuPlatform *string `json:"minCpuPlatform,omitempty"`

	/* Immutable. The set of Google API scopes to be made available on all of the node VMs. */
	// +optional
	OauthScopes []string `json:"oauthScopes,omitempty"`

	/* Immutable. Whether the nodes are created as preemptible VM instances. */
	// +optional
	Preemptible *bool `json:"preemptible,omitempty"`

	/* Immutable. Sandbox configuration for this node. */
	// +optional
	SandboxConfig *ClusterSandboxConfig `json:"sandboxConfig,omitempty"`

	/*  */
	// +optional
	ServiceAccountRef *v1alpha1.ResourceRef `json:"serviceAccountRef,omitempty"`

	/* Immutable. Shielded Instance options. */
	// +optional
	ShieldedInstanceConfig *ClusterShieldedInstanceConfig `json:"shieldedInstanceConfig,omitempty"`

	/* Immutable. The list of instance tags applied to all nodes. */
	// +optional
	Tags []string `json:"tags,omitempty"`

	/* Immutable. List of Kubernetes taints to be applied to each node. */
	// +optional
	Taint []ClusterTaint `json:"taint,omitempty"`

	/* Immutable. The workload metadata configuration for this node. */
	// +optional
	WorkloadMetadataConfig *ClusterWorkloadMetadataConfig `json:"workloadMetadataConfig,omitempty"`
}

type ClusterNotificationConfig struct {
	/* Notification config for Cloud Pub/Sub */
	Pubsub ClusterPubsub `json:"pubsub"`
}

type ClusterPassword struct {
	/* Value of the field. Cannot be used if 'valueFrom' is specified. */
	// +optional
	Value *string `json:"value,omitempty"`

	/* Source for the field's value. Cannot be used if 'value' is specified. */
	// +optional
	ValueFrom *ClusterValueFrom `json:"valueFrom,omitempty"`
}

type ClusterPodSecurityPolicyConfig struct {
	/* Enable the PodSecurityPolicy controller for this cluster. If enabled, pods must be valid under a PodSecurityPolicy to be created. */
	Enabled bool `json:"enabled"`
}

type ClusterPrivateClusterConfig struct {
	/* Immutable. Enables the private cluster feature, creating a private endpoint on the cluster. In a private cluster, nodes only have RFC 1918 private addresses and communicate with the master's private endpoint via private networking. */
	EnablePrivateEndpoint bool `json:"enablePrivateEndpoint"`

	/* Immutable. When true, the cluster's private endpoint is used as the cluster endpoint and access through the public endpoint is disabled. When false, either endpoint can be used. This field only applies to private clusters, when enable_private_nodes is true. */
	// +optional
	EnablePrivateNodes *bool `json:"enablePrivateNodes,omitempty"`

	/* Controls cluster master global access settings. */
	// +optional
	MasterGlobalAccessConfig *ClusterMasterGlobalAccessConfig `json:"masterGlobalAccessConfig,omitempty"`

	/* Immutable. The IP range in CIDR notation to use for the hosted master network. This range will be used for assigning private IP addresses to the cluster master(s) and the ILB VIP. This range must not overlap with any other ranges in use within the cluster's network, and it must be a /28 subnet. See Private Cluster Limitations for more details. This field only applies to private clusters, when enable_private_nodes is true. */
	// +optional
	MasterIpv4CidrBlock *string `json:"masterIpv4CidrBlock,omitempty"`

	/* The name of the peering between this cluster and the Google owned VPC. */
	// +optional
	PeeringName *string `json:"peeringName,omitempty"`

	/* The internal IP address of this cluster's master endpoint. */
	// +optional
	PrivateEndpoint *string `json:"privateEndpoint,omitempty"`

	/* The external IP address of this cluster's master endpoint. */
	// +optional
	PublicEndpoint *string `json:"publicEndpoint,omitempty"`
}

type ClusterPubsub struct {
	/* Whether or not the notification config is enabled */
	Enabled bool `json:"enabled"`

	/* The PubSubTopic to send the notification to. */
	// +optional
	TopicRef *v1alpha1.ResourceRef `json:"topicRef,omitempty"`
}

type ClusterRecurringWindow struct {
	/*  */
	EndTime string `json:"endTime"`

	/*  */
	Recurrence string `json:"recurrence"`

	/*  */
	StartTime string `json:"startTime"`
}

type ClusterReleaseChannel struct {
	/* The selected release channel. Accepted values are:
	* UNSPECIFIED: Not set.
	* RAPID: Weekly upgrade cadence; Early testers and developers who requires new features.
	* REGULAR: Multiple per month upgrade cadence; Production users who need features not yet offered in the Stable channel.
	* STABLE: Every few months upgrade cadence; Production users who need stability above all else, and for whom frequent upgrades are too risky. */
	Channel string `json:"channel"`
}

type ClusterResourceLimits struct {
	/* Maximum amount of the resource in the cluster. */
	// +optional
	Maximum *int `json:"maximum,omitempty"`

	/* Minimum amount of the resource in the cluster. */
	// +optional
	Minimum *int `json:"minimum,omitempty"`

	/* The type of the resource. For example, cpu and memory. See the guide to using Node Auto-Provisioning for a list of types. */
	ResourceType string `json:"resourceType"`
}

type ClusterResourceUsageExportConfig struct {
	/* Parameters for using BigQuery as the destination of resource usage export. */
	BigqueryDestination ClusterBigqueryDestination `json:"bigqueryDestination"`

	/* Whether to enable network egress metering for this cluster. If enabled, a daemonset will be created in the cluster to meter network egress traffic. */
	// +optional
	EnableNetworkEgressMetering *bool `json:"enableNetworkEgressMetering,omitempty"`

	/* Whether to enable resource consumption metering on this cluster. When enabled, a table will be created in the resource export BigQuery dataset to store resource consumption data. The resulting table can be joined with the resource usage table or with BigQuery billing export. Defaults to true. */
	// +optional
	EnableResourceConsumptionMetering *bool `json:"enableResourceConsumptionMetering,omitempty"`
}

type ClusterSandboxConfig struct {
	/* Type of the sandbox to use for the node (e.g. 'gvisor') */
	SandboxType string `json:"sandboxType"`
}

type ClusterShieldedInstanceConfig struct {
	/* Immutable. Defines whether the instance has integrity monitoring enabled. */
	// +optional
	EnableIntegrityMonitoring *bool `json:"enableIntegrityMonitoring,omitempty"`

	/* Immutable. Defines whether the instance has Secure Boot enabled. */
	// +optional
	EnableSecureBoot *bool `json:"enableSecureBoot,omitempty"`
}

type ClusterTaint struct {
	/* Immutable. Effect for taint. */
	Effect string `json:"effect"`

	/* Immutable. Key for taint. */
	Key string `json:"key"`

	/* Immutable. Value for taint. */
	Value string `json:"value"`
}

type ClusterValueFrom struct {
	/* Reference to a value with the given key in the given Secret in the resource's namespace. */
	// +optional
	SecretKeyRef *v1alpha1.ResourceRef `json:"secretKeyRef,omitempty"`
}

type ClusterVerticalPodAutoscaling struct {
	/* Enables vertical pod autoscaling. */
	Enabled bool `json:"enabled"`
}

type ClusterWorkloadIdentityConfig struct {
	/* Enables workload identity. */
	IdentityNamespace string `json:"identityNamespace"`
}

type ClusterWorkloadMetadataConfig struct {
	/* NodeMetadata is the configuration for how to expose metadata to the workloads running on the node. */
	NodeMetadata string `json:"nodeMetadata"`
}

type ContainerClusterSpec struct {
	/* The configuration for addons supported by GKE. */
	// +optional
	AddonsConfig *ClusterAddonsConfig `json:"addonsConfig,omitempty"`

	/* Immutable. Configuration for the Google Groups for GKE feature. */
	// +optional
	AuthenticatorGroupsConfig *ClusterAuthenticatorGroupsConfig `json:"authenticatorGroupsConfig,omitempty"`

	/* Per-cluster configuration of Node Auto-Provisioning with Cluster Autoscaler to automatically adjust the size of the cluster and create/delete node pools based on the current needs of the cluster's workload. See the guide to using Node Auto-Provisioning for more details. */
	// +optional
	ClusterAutoscaling *ClusterClusterAutoscaling `json:"clusterAutoscaling,omitempty"`

	/* Immutable. The IP address range of the Kubernetes pods in this cluster in CIDR notation (e.g. 10.96.0.0/14). Leave blank to have one automatically chosen or specify a /14 block in 10.0.0.0/8. This field will only work for routes-based clusters, where ip_allocation_policy is not defined. */
	// +optional
	ClusterIpv4Cidr *string `json:"clusterIpv4Cidr,omitempty"`

	/* Telemetry integration for the cluster. */
	// +optional
	ClusterTelemetry *ClusterClusterTelemetry `json:"clusterTelemetry,omitempty"`

	/* Immutable. Configuration for the confidential nodes feature, which makes nodes run on confidential VMs. Warning: This configuration can't be changed (or added/removed) after cluster creation without deleting and recreating the entire cluster. */
	// +optional
	ConfidentialNodes *ClusterConfidentialNodes `json:"confidentialNodes,omitempty"`

	/* Application-layer Secrets Encryption settings. The object format is {state = string, key_name = string}. Valid values of state are: "ENCRYPTED"; "DECRYPTED". key_name is the name of a CloudKMS key. */
	// +optional
	DatabaseEncryption *ClusterDatabaseEncryption `json:"databaseEncryption,omitempty"`

	/* The desired datapath provider for this cluster. By default, uses the IPTables-based kube-proxy implementation. */
	// +optional
	DatapathProvider *string `json:"datapathProvider,omitempty"`

	/* Immutable. The default maximum number of pods per node in this cluster. This doesn't work on "routes-based" clusters, clusters that don't have IP Aliasing enabled. */
	// +optional
	DefaultMaxPodsPerNode *int `json:"defaultMaxPodsPerNode,omitempty"`

	/* Whether the cluster disables default in-node sNAT rules. In-node sNAT rules will be disabled when defaultSnatStatus is disabled. */
	// +optional
	DefaultSnatStatus *ClusterDefaultSnatStatus `json:"defaultSnatStatus,omitempty"`

	/* Immutable.  Description of the cluster. */
	// +optional
	Description *string `json:"description,omitempty"`

	/* Immutable. Enable Autopilot for this cluster. */
	// +optional
	EnableAutopilot *bool `json:"enableAutopilot,omitempty"`

	/* Enable Binary Authorization for this cluster. If enabled, all container images will be validated by Google Binary Authorization. */
	// +optional
	EnableBinaryAuthorization *bool `json:"enableBinaryAuthorization,omitempty"`

	/* Whether Intra-node visibility is enabled for this cluster. This makes same node pod to pod traffic visible for VPC network. */
	// +optional
	EnableIntranodeVisibility *bool `json:"enableIntranodeVisibility,omitempty"`

	/* Immutable. Whether to enable Kubernetes Alpha features for this cluster. Note that when this option is enabled, the cluster cannot be upgraded and will be automatically deleted after 30 days. */
	// +optional
	EnableKubernetesAlpha *bool `json:"enableKubernetesAlpha,omitempty"`

	/* Whether L4ILB Subsetting is enabled for this cluster. */
	// +optional
	EnableL4IlbSubsetting *bool `json:"enableL4IlbSubsetting,omitempty"`

	/* Whether the ABAC authorizer is enabled for this cluster. When enabled, identities in the system, including service accounts, nodes, and controllers, will have statically granted permissions beyond those provided by the RBAC configuration or IAM. Defaults to false. */
	// +optional
	EnableLegacyAbac *bool `json:"enableLegacyAbac,omitempty"`

	/* Enable Shielded Nodes features on all nodes in this cluster. */
	// +optional
	EnableShieldedNodes *bool `json:"enableShieldedNodes,omitempty"`

	/* Immutable. Whether to enable Cloud TPU resources in this cluster. */
	// +optional
	EnableTpu *bool `json:"enableTpu,omitempty"`

	/* Immutable. The number of nodes to create in this cluster's default node pool. In regional or multi-zonal clusters, this is the number of nodes per zone. Must be set if node_pool is not set. If you're using google_container_node_pool objects with no default node pool, you'll need to set this to a value of at least 1, alongside setting remove_default_node_pool to true. */
	// +optional
	InitialNodeCount *int `json:"initialNodeCount,omitempty"`

	/* Immutable. Configuration of cluster IP allocation for VPC-native clusters. Adding this block enables IP aliasing, making the cluster VPC-native instead of routes-based. */
	// +optional
	IpAllocationPolicy *ClusterIpAllocationPolicy `json:"ipAllocationPolicy,omitempty"`

	/* Immutable. The location (region or zone) in which the cluster master will be created, as well as the default node location. If you specify a zone (such as us-central1-a), the cluster will be a zonal cluster with a single cluster master. If you specify a region (such as us-west1), the cluster will be a regional cluster with multiple masters spread across zones in the region, and with default node locations in those zones as well. */
	Location string `json:"location"`

	/* The logging service that the cluster should write logs to. Available options include logging.googleapis.com(Legacy Stackdriver), logging.googleapis.com/kubernetes(Stackdriver Kubernetes Engine Logging), and none. Defaults to logging.googleapis.com/kubernetes. */
	// +optional
	LoggingService *string `json:"loggingService,omitempty"`

	/* The maintenance policy to use for the cluster. */
	// +optional
	MaintenancePolicy *ClusterMaintenancePolicy `json:"maintenancePolicy,omitempty"`

	/* The authentication information for accessing the Kubernetes master. Some values in this block are only returned by the API if your service account has permission to get credentials for your GKE cluster. If you see an unexpected diff removing a username/password or unsetting your client cert, ensure you have the container.clusters.getCredentials permission. */
	// +optional
	MasterAuth *ClusterMasterAuth `json:"masterAuth,omitempty"`

	/* The desired configuration options for master authorized networks. Omit the nested cidr_blocks attribute to disallow external access (except the cluster node IPs, which GKE automatically whitelists). */
	// +optional
	MasterAuthorizedNetworksConfig *ClusterMasterAuthorizedNetworksConfig `json:"masterAuthorizedNetworksConfig,omitempty"`

	/* The minimum version of the master. GKE will auto-update the master to new versions, so this does not guarantee the current master version--use the read-only master_version field to obtain that. If unset, the cluster's version will be set by GKE to the version of the most recent official release (which is not necessarily the latest version). */
	// +optional
	MinMasterVersion *string `json:"minMasterVersion,omitempty"`

	/* The monitoring service that the cluster should write metrics to. Automatically send metrics from pods in the cluster to the Google Cloud Monitoring API. VM metrics will be collected by Google Compute Engine regardless of this setting Available options include monitoring.googleapis.com(Legacy Stackdriver), monitoring.googleapis.com/kubernetes(Stackdriver Kubernetes Engine Monitoring), and none. Defaults to monitoring.googleapis.com/kubernetes. */
	// +optional
	MonitoringService *string `json:"monitoringService,omitempty"`

	/* Configuration options for the NetworkPolicy feature. */
	// +optional
	NetworkPolicy *ClusterNetworkPolicy `json:"networkPolicy,omitempty"`

	/*  */
	// +optional
	NetworkRef *v1alpha1.ResourceRef `json:"networkRef,omitempty"`

	/* Immutable. Determines whether alias IPs or routes will be used for pod IPs in the cluster. */
	// +optional
	NetworkingMode *string `json:"networkingMode,omitempty"`

	/* Immutable. The configuration of the nodepool */
	// +optional
	NodeConfig *ClusterNodeConfig `json:"nodeConfig,omitempty"`

	/* The list of zones in which the cluster's nodes are located. Nodes must be in the region of their regional cluster or in the same region as their cluster's zone for zonal clusters. If this is specified for a zonal cluster, omit the cluster's zone. */
	// +optional
	NodeLocations []string `json:"nodeLocations,omitempty"`

	/*  */
	// +optional
	NodeVersion *string `json:"nodeVersion,omitempty"`

	/* The notification config for sending cluster upgrade notifications */
	// +optional
	NotificationConfig *ClusterNotificationConfig `json:"notificationConfig,omitempty"`

	/* Configuration for the PodSecurityPolicy feature. */
	// +optional
	PodSecurityPolicyConfig *ClusterPodSecurityPolicyConfig `json:"podSecurityPolicyConfig,omitempty"`

	/* Configuration for private clusters, clusters with private nodes. */
	// +optional
	PrivateClusterConfig *ClusterPrivateClusterConfig `json:"privateClusterConfig,omitempty"`

	/* The desired state of IPv6 connectivity to Google Services. By default, no private IPv6 access to or from Google Services (all access will be via IPv4). */
	// +optional
	PrivateIpv6GoogleAccess *string `json:"privateIpv6GoogleAccess,omitempty"`

	/* Configuration options for the Release channel feature, which provide more control over automatic upgrades of your GKE clusters. Note that removing this field from your config will not unenroll it. Instead, use the "UNSPECIFIED" channel. */
	// +optional
	ReleaseChannel *ClusterReleaseChannel `json:"releaseChannel,omitempty"`

	/* Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default. */
	// +optional
	ResourceID *string `json:"resourceID,omitempty"`

	/* Configuration for the ResourceUsageExportConfig feature. */
	// +optional
	ResourceUsageExportConfig *ClusterResourceUsageExportConfig `json:"resourceUsageExportConfig,omitempty"`

	/*  */
	// +optional
	SubnetworkRef *v1alpha1.ResourceRef `json:"subnetworkRef,omitempty"`

	/* Vertical Pod Autoscaling automatically adjusts the resources of pods controlled by it. */
	// +optional
	VerticalPodAutoscaling *ClusterVerticalPodAutoscaling `json:"verticalPodAutoscaling,omitempty"`

	/* Configuration for the use of Kubernetes Service Accounts in GCP IAM policies. */
	// +optional
	WorkloadIdentityConfig *ClusterWorkloadIdentityConfig `json:"workloadIdentityConfig,omitempty"`
}

type ContainerClusterStatus struct {
	/* Conditions represent the latest available observations of the
	   ContainerCluster's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`
	/* The IP address of this cluster's Kubernetes master. */
	Endpoint string `json:"endpoint,omitempty"`
	/* List of instance group URLs which have been assigned to the cluster. */
	InstanceGroupUrls []string `json:"instanceGroupUrls,omitempty"`
	/* The fingerprint of the set of labels for this cluster. */
	LabelFingerprint string `json:"labelFingerprint,omitempty"`
	/* The current version of the master in the cluster. This may be different than the min_master_version set in the config if the master has been updated by GKE. */
	MasterVersion string `json:"masterVersion,omitempty"`
	/* ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource. */
	ObservedGeneration int `json:"observedGeneration,omitempty"`
	/*  */
	Operation string `json:"operation,omitempty"`
	/* Server-defined URL for the resource. */
	SelfLink string `json:"selfLink,omitempty"`
	/* The IP address range of the Kubernetes services in this cluster, in CIDR notation (e.g. 1.2.3.4/29). Service addresses are typically put in the last /16 from the container CIDR. */
	ServicesIpv4Cidr string `json:"servicesIpv4Cidr,omitempty"`
	/* The IP address range of the Cloud TPUs in this cluster, in CIDR notation (e.g. 1.2.3.4/29). */
	TpuIpv4CidrBlock string `json:"tpuIpv4CidrBlock,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ContainerCluster is the Schema for the container API
// +k8s:openapi-gen=true
type ContainerCluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ContainerClusterSpec   `json:"spec,omitempty"`
	Status ContainerClusterStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ContainerClusterList contains a list of ContainerCluster
type ContainerClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ContainerCluster `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ContainerCluster{}, &ContainerClusterList{})
}
