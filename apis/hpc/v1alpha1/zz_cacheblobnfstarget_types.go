/*
Copyright 2020 The Crossplane Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by terrajet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type CacheBlobNfsTargetObservation struct {
}

type CacheBlobNfsTargetParameters struct {

	// +kubebuilder:validation:Optional
	AccessPolicyName *string `json:"accessPolicyName,omitempty" tf:"access_policy_name,omitempty"`

	// +kubebuilder:validation:Required
	CacheName *string `json:"cacheName" tf:"cache_name,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Required
	NamespacePath *string `json:"namespacePath" tf:"namespace_path,omitempty"`

	// +kubebuilder:validation:Required
	ResourceGroupName *string `json:"resourceGroupName" tf:"resource_group_name,omitempty"`

	// +kubebuilder:validation:Required
	StorageContainerID *string `json:"storageContainerId" tf:"storage_container_id,omitempty"`

	// +kubebuilder:validation:Required
	UsageModel *string `json:"usageModel" tf:"usage_model,omitempty"`
}

// CacheBlobNfsTargetSpec defines the desired state of CacheBlobNfsTarget
type CacheBlobNfsTargetSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     CacheBlobNfsTargetParameters `json:"forProvider"`
}

// CacheBlobNfsTargetStatus defines the observed state of CacheBlobNfsTarget.
type CacheBlobNfsTargetStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        CacheBlobNfsTargetObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// CacheBlobNfsTarget is the Schema for the CacheBlobNfsTargets API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azurejet}
type CacheBlobNfsTarget struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CacheBlobNfsTargetSpec   `json:"spec"`
	Status            CacheBlobNfsTargetStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CacheBlobNfsTargetList contains a list of CacheBlobNfsTargets
type CacheBlobNfsTargetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CacheBlobNfsTarget `json:"items"`
}

// Repository type metadata.
var (
	CacheBlobNfsTarget_Kind             = "CacheBlobNfsTarget"
	CacheBlobNfsTarget_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: CacheBlobNfsTarget_Kind}.String()
	CacheBlobNfsTarget_KindAPIVersion   = CacheBlobNfsTarget_Kind + "." + CRDGroupVersion.String()
	CacheBlobNfsTarget_GroupVersionKind = CRDGroupVersion.WithKind(CacheBlobNfsTarget_Kind)
)

func init() {
	SchemeBuilder.Register(&CacheBlobNfsTarget{}, &CacheBlobNfsTargetList{})
}
