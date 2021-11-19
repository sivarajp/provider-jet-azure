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

type DomainTopicObservation struct {
}

type DomainTopicParameters struct {

	// +kubebuilder:validation:Required
	DomainName *string `json:"domainName" tf:"domain_name,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Required
	ResourceGroupName *string `json:"resourceGroupName" tf:"resource_group_name,omitempty"`
}

// DomainTopicSpec defines the desired state of DomainTopic
type DomainTopicSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DomainTopicParameters `json:"forProvider"`
}

// DomainTopicStatus defines the observed state of DomainTopic.
type DomainTopicStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DomainTopicObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// DomainTopic is the Schema for the DomainTopics API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azurejet}
type DomainTopic struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DomainTopicSpec   `json:"spec"`
	Status            DomainTopicStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DomainTopicList contains a list of DomainTopics
type DomainTopicList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DomainTopic `json:"items"`
}

// Repository type metadata.
var (
	DomainTopic_Kind             = "DomainTopic"
	DomainTopic_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: DomainTopic_Kind}.String()
	DomainTopic_KindAPIVersion   = DomainTopic_Kind + "." + CRDGroupVersion.String()
	DomainTopic_GroupVersionKind = CRDGroupVersion.WithKind(DomainTopic_Kind)
)

func init() {
	SchemeBuilder.Register(&DomainTopic{}, &DomainTopicList{})
}
