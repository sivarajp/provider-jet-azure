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

type EnrichmentObservation_2 struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type EnrichmentParameters_2 struct {

	// +kubebuilder:validation:Required
	EndpointNames []*string `json:"endpointNames" tf:"endpoint_names,omitempty"`

	// +kubebuilder:validation:Required
	IothubName *string `json:"iothubName" tf:"iothub_name,omitempty"`

	// +kubebuilder:validation:Required
	Key *string `json:"key" tf:"key,omitempty"`

	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-jet-azure/apis/azure/v1alpha1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Required
	Value *string `json:"value" tf:"value,omitempty"`
}

// EnrichmentSpec defines the desired state of Enrichment
type EnrichmentSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     EnrichmentParameters_2 `json:"forProvider"`
}

// EnrichmentStatus defines the observed state of Enrichment.
type EnrichmentStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        EnrichmentObservation_2 `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Enrichment is the Schema for the Enrichments API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azurejet}
type Enrichment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              EnrichmentSpec   `json:"spec"`
	Status            EnrichmentStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// EnrichmentList contains a list of Enrichments
type EnrichmentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Enrichment `json:"items"`
}

// Repository type metadata.
var (
	Enrichment_Kind             = "Enrichment"
	Enrichment_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Enrichment_Kind}.String()
	Enrichment_KindAPIVersion   = Enrichment_Kind + "." + CRDGroupVersion.String()
	Enrichment_GroupVersionKind = CRDGroupVersion.WithKind(Enrichment_Kind)
)

func init() {
	SchemeBuilder.Register(&Enrichment{}, &EnrichmentList{})
}
