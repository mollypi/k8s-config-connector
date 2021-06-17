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

type AutoscalingpolicyBasicAlgorithm struct {
	/* Optional. Duration between scaling events. A scaling period starts after the update operation from the previous event has completed. Bounds: . Default: 2m. */
	// +optional
	CooldownPeriod *string `json:"cooldownPeriod,omitempty"`

	/* Required. YARN autoscaling configuration. */
	YarnConfig AutoscalingpolicyYarnConfig `json:"yarnConfig"`
}

type AutoscalingpolicySecondaryWorkerConfig struct {
	/* Optional. Maximum number of instances for this group. Note that by default, clusters will not use secondary workers. Required for secondary workers if the minimum secondary instances is set. Primary workers - Bounds: [min_instances, ). Secondary workers - Bounds: [min_instances, ). Default: 0. */
	// +optional
	MaxInstances *int `json:"maxInstances,omitempty"`

	/* Optional. Minimum number of instances for this group. Primary workers - Bounds: . Default: 0. */
	// +optional
	MinInstances *int `json:"minInstances,omitempty"`

	/* Optional. Weight for the instance group, which is used to determine the fraction of total workers in the cluster from this instance group. For example, if primary workers have weight 2, and secondary workers have weight 1, the cluster will have approximately 2 primary workers for each secondary worker. The cluster may not reach the specified balance if constrained by min/max bounds or other autoscaling settings. For example, if `max_instances` for secondary workers is 0, then only primary workers will be added. The cluster can also be out of balance when created. If weight is not set on any instance group, the cluster will default to equal weight for all groups: the cluster will attempt to maintain an equal number of workers in each group within the configured size bounds for each group. If weight is set for one group only, the cluster will default to zero weight on the unset group. For example if weight is set only on primary workers, the cluster will use primary workers only and no secondary workers. */
	// +optional
	Weight *int `json:"weight,omitempty"`
}

type AutoscalingpolicyWorkerConfig struct {
	/* Required. Maximum number of instances for this group. Required for primary workers. Note that by default, clusters will not use secondary workers. Required for secondary workers if the minimum secondary instances is set. Primary workers - Bounds: [min_instances, ). Secondary workers - Bounds: [min_instances, ). Default: 0. */
	MaxInstances int `json:"maxInstances"`

	/* Optional. Minimum number of instances for this group. Primary workers - Bounds: . Default: 0. */
	// +optional
	MinInstances *int `json:"minInstances,omitempty"`

	/* Optional. Weight for the instance group, which is used to determine the fraction of total workers in the cluster from this instance group. For example, if primary workers have weight 2, and secondary workers have weight 1, the cluster will have approximately 2 primary workers for each secondary worker. The cluster may not reach the specified balance if constrained by min/max bounds or other autoscaling settings. For example, if `max_instances` for secondary workers is 0, then only primary workers will be added. The cluster can also be out of balance when created. If weight is not set on any instance group, the cluster will default to equal weight for all groups: the cluster will attempt to maintain an equal number of workers in each group within the configured size bounds for each group. If weight is set for one group only, the cluster will default to zero weight on the unset group. For example if weight is set only on primary workers, the cluster will use primary workers only and no secondary workers. */
	// +optional
	Weight *int `json:"weight,omitempty"`
}

type AutoscalingpolicyYarnConfig struct {
	/* Required. Timeout for YARN graceful decommissioning of Node Managers. Specifies the duration to wait for jobs to complete before forcefully removing workers (and potentially interrupting jobs). Only applicable to downscaling operations. */
	GracefulDecommissionTimeout string `json:"gracefulDecommissionTimeout"`

	/* Required. Fraction of average YARN pending memory in the last cooldown period for which to remove workers. A scale-down factor of 1 will result in scaling down so that there is no available memory remaining after the update (more aggressive scaling). A scale-down factor of 0 disables removing workers, which can be beneficial for autoscaling a single job. See . */
	ScaleDownFactor float64 `json:"scaleDownFactor"`

	/* Optional. Minimum scale-down threshold as a fraction of total cluster size before scaling occurs. For example, in a 20-worker cluster, a threshold of 0.1 means the autoscaler must recommend at least a 2 worker scale-down for the cluster to scale. A threshold of 0 means the autoscaler will scale down on any recommended change. Bounds: . Default: 0.0. */
	// +optional
	ScaleDownMinWorkerFraction *float64 `json:"scaleDownMinWorkerFraction,omitempty"`

	/* Required. Fraction of average YARN pending memory in the last cooldown period for which to add workers. A scale-up factor of 1.0 will result in scaling up so that there is no pending memory remaining after the update (more aggressive scaling). A scale-up factor closer to 0 will result in a smaller magnitude of scaling up (less aggressive scaling). See . */
	ScaleUpFactor float64 `json:"scaleUpFactor"`

	/* Optional. Minimum scale-up threshold as a fraction of total cluster size before scaling occurs. For example, in a 20-worker cluster, a threshold of 0.1 means the autoscaler must recommend at least a 2-worker scale-up for the cluster to scale. A threshold of 0 means the autoscaler will scale up on any recommended change. Bounds: . Default: 0.0. */
	// +optional
	ScaleUpMinWorkerFraction *float64 `json:"scaleUpMinWorkerFraction,omitempty"`
}

type DataprocAutoscalingPolicySpec struct {
	/*  */
	BasicAlgorithm AutoscalingpolicyBasicAlgorithm `json:"basicAlgorithm"`

	/* The location for the resource */
	Location string `json:"location"`

	/* The Project that this resource belongs to. */
	// +optional
	ProjectRef *v1alpha1.ResourceRef `json:"projectRef,omitempty"`

	/* Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default. */
	// +optional
	ResourceID *string `json:"resourceID,omitempty"`

	/* Optional. Describes how the autoscaler will operate for secondary workers. */
	// +optional
	SecondaryWorkerConfig *AutoscalingpolicySecondaryWorkerConfig `json:"secondaryWorkerConfig,omitempty"`

	/* Required. Describes how the autoscaler will operate for primary workers. */
	WorkerConfig AutoscalingpolicyWorkerConfig `json:"workerConfig"`
}

type DataprocAutoscalingPolicyStatus struct {
	/* Conditions represent the latest available observations of the
	   DataprocAutoscalingPolicy's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`
	/* ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource. */
	ObservedGeneration int `json:"observedGeneration,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// DataprocAutoscalingPolicy is the Schema for the dataproc API
// +k8s:openapi-gen=true
type DataprocAutoscalingPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DataprocAutoscalingPolicySpec   `json:"spec,omitempty"`
	Status DataprocAutoscalingPolicyStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// DataprocAutoscalingPolicyList contains a list of DataprocAutoscalingPolicy
type DataprocAutoscalingPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DataprocAutoscalingPolicy `json:"items"`
}

func init() {
	SchemeBuilder.Register(&DataprocAutoscalingPolicy{}, &DataprocAutoscalingPolicyList{})
}
