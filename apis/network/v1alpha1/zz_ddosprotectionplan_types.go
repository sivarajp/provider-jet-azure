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

type DDOSProtectionPlanObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	VirtualNetworkIds []*string `json:"virtualNetworkIds,omitempty" tf:"virtual_network_ids,omitempty"`
}

type DDOSProtectionPlanParameters struct {

	// +kubebuilder:validation:Required
	Location *string `json:"location" tf:"location,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-jet-azure/apis/azure/v1alpha2.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// DDOSProtectionPlanSpec defines the desired state of DDOSProtectionPlan
type DDOSProtectionPlanSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DDOSProtectionPlanParameters `json:"forProvider"`
}

// DDOSProtectionPlanStatus defines the observed state of DDOSProtectionPlan.
type DDOSProtectionPlanStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DDOSProtectionPlanObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// DDOSProtectionPlan is the Schema for the DDOSProtectionPlans API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azurejet}
type DDOSProtectionPlan struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DDOSProtectionPlanSpec   `json:"spec"`
	Status            DDOSProtectionPlanStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DDOSProtectionPlanList contains a list of DDOSProtectionPlans
type DDOSProtectionPlanList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DDOSProtectionPlan `json:"items"`
}

// Repository type metadata.
var (
	DDOSProtectionPlan_Kind             = "DDOSProtectionPlan"
	DDOSProtectionPlan_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: DDOSProtectionPlan_Kind}.String()
	DDOSProtectionPlan_KindAPIVersion   = DDOSProtectionPlan_Kind + "." + CRDGroupVersion.String()
	DDOSProtectionPlan_GroupVersionKind = CRDGroupVersion.WithKind(DDOSProtectionPlan_Kind)
)

func init() {
	SchemeBuilder.Register(&DDOSProtectionPlan{}, &DDOSProtectionPlanList{})
}
