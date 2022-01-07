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

type HealthBotObservation struct {
	BotManagementPortalURL *string `json:"botManagementPortalUrl,omitempty" tf:"bot_management_portal_url,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type HealthBotParameters struct {

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

	// +kubebuilder:validation:Required
	SkuName *string `json:"skuName" tf:"sku_name,omitempty"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// HealthBotSpec defines the desired state of HealthBot
type HealthBotSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     HealthBotParameters `json:"forProvider"`
}

// HealthBotStatus defines the observed state of HealthBot.
type HealthBotStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        HealthBotObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// HealthBot is the Schema for the HealthBots API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azurejet}
type HealthBot struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              HealthBotSpec   `json:"spec"`
	Status            HealthBotStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// HealthBotList contains a list of HealthBots
type HealthBotList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []HealthBot `json:"items"`
}

// Repository type metadata.
var (
	HealthBot_Kind             = "HealthBot"
	HealthBot_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: HealthBot_Kind}.String()
	HealthBot_KindAPIVersion   = HealthBot_Kind + "." + CRDGroupVersion.String()
	HealthBot_GroupVersionKind = CRDGroupVersion.WithKind(HealthBot_Kind)
)

func init() {
	SchemeBuilder.Register(&HealthBot{}, &HealthBotList{})
}
