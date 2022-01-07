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

type LinkedServiceWebObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type LinkedServiceWebParameters struct {

	// +kubebuilder:validation:Optional
	AdditionalProperties map[string]*string `json:"additionalProperties,omitempty" tf:"additional_properties,omitempty"`

	// +kubebuilder:validation:Optional
	Annotations []*string `json:"annotations,omitempty" tf:"annotations,omitempty"`

	// +kubebuilder:validation:Required
	AuthenticationType *string `json:"authenticationType" tf:"authentication_type,omitempty"`

	// +kubebuilder:validation:Required
	DataFactoryName *string `json:"dataFactoryName" tf:"data_factory_name,omitempty"`

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Optional
	IntegrationRuntimeName *string `json:"integrationRuntimeName,omitempty" tf:"integration_runtime_name,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	Parameters map[string]*string `json:"parameters,omitempty" tf:"parameters,omitempty"`

	// +kubebuilder:validation:Optional
	PasswordSecretRef *v1.SecretKeySelector `json:"passwordSecretRef,omitempty" tf:"-"`

	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-jet-azure/apis/azure/v1alpha2.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Required
	URL *string `json:"url" tf:"url,omitempty"`

	// +kubebuilder:validation:Optional
	Username *string `json:"username,omitempty" tf:"username,omitempty"`
}

// LinkedServiceWebSpec defines the desired state of LinkedServiceWeb
type LinkedServiceWebSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     LinkedServiceWebParameters `json:"forProvider"`
}

// LinkedServiceWebStatus defines the observed state of LinkedServiceWeb.
type LinkedServiceWebStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        LinkedServiceWebObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// LinkedServiceWeb is the Schema for the LinkedServiceWebs API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azurejet}
type LinkedServiceWeb struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              LinkedServiceWebSpec   `json:"spec"`
	Status            LinkedServiceWebStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// LinkedServiceWebList contains a list of LinkedServiceWebs
type LinkedServiceWebList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LinkedServiceWeb `json:"items"`
}

// Repository type metadata.
var (
	LinkedServiceWeb_Kind             = "LinkedServiceWeb"
	LinkedServiceWeb_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: LinkedServiceWeb_Kind}.String()
	LinkedServiceWeb_KindAPIVersion   = LinkedServiceWeb_Kind + "." + CRDGroupVersion.String()
	LinkedServiceWeb_GroupVersionKind = CRDGroupVersion.WithKind(LinkedServiceWeb_Kind)
)

func init() {
	SchemeBuilder.Register(&LinkedServiceWeb{}, &LinkedServiceWebList{})
}
