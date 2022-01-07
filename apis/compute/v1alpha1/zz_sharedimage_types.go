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

type IdentifierObservation struct {
}

type IdentifierParameters struct {

	// +kubebuilder:validation:Required
	Offer *string `json:"offer" tf:"offer,omitempty"`

	// +kubebuilder:validation:Required
	Publisher *string `json:"publisher" tf:"publisher,omitempty"`

	// +kubebuilder:validation:Required
	Sku *string `json:"sku" tf:"sku,omitempty"`
}

type PurchasePlanObservation struct {
}

type PurchasePlanParameters struct {

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	Product *string `json:"product,omitempty" tf:"product,omitempty"`

	// +kubebuilder:validation:Optional
	Publisher *string `json:"publisher,omitempty" tf:"publisher,omitempty"`
}

type SharedImageObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type SharedImageParameters struct {

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Optional
	Eula *string `json:"eula,omitempty" tf:"eula,omitempty"`

	// +kubebuilder:validation:Required
	GalleryName *string `json:"galleryName" tf:"gallery_name,omitempty"`

	// +kubebuilder:validation:Optional
	HyperVGeneration *string `json:"hyperVGeneration,omitempty" tf:"hyper_v_generation,omitempty"`

	// +kubebuilder:validation:Required
	Identifier []IdentifierParameters `json:"identifier" tf:"identifier,omitempty"`

	// +kubebuilder:validation:Required
	Location *string `json:"location" tf:"location,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Required
	OsType *string `json:"osType" tf:"os_type,omitempty"`

	// +kubebuilder:validation:Optional
	PrivacyStatementURI *string `json:"privacyStatementUri,omitempty" tf:"privacy_statement_uri,omitempty"`

	// +kubebuilder:validation:Optional
	PurchasePlan []PurchasePlanParameters `json:"purchasePlan,omitempty" tf:"purchase_plan,omitempty"`

	// +kubebuilder:validation:Optional
	ReleaseNoteURI *string `json:"releaseNoteUri,omitempty" tf:"release_note_uri,omitempty"`

	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-jet-azure/apis/azure/v1alpha2.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	Specialized *bool `json:"specialized,omitempty" tf:"specialized,omitempty"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// SharedImageSpec defines the desired state of SharedImage
type SharedImageSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SharedImageParameters `json:"forProvider"`
}

// SharedImageStatus defines the observed state of SharedImage.
type SharedImageStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SharedImageObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// SharedImage is the Schema for the SharedImages API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azurejet}
type SharedImage struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SharedImageSpec   `json:"spec"`
	Status            SharedImageStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SharedImageList contains a list of SharedImages
type SharedImageList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SharedImage `json:"items"`
}

// Repository type metadata.
var (
	SharedImage_Kind             = "SharedImage"
	SharedImage_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: SharedImage_Kind}.String()
	SharedImage_KindAPIVersion   = SharedImage_Kind + "." + CRDGroupVersion.String()
	SharedImage_GroupVersionKind = CRDGroupVersion.WithKind(SharedImage_Kind)
)

func init() {
	SchemeBuilder.Register(&SharedImage{}, &SharedImageList{})
}
