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

type EndpointServicebusQueueObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type EndpointServicebusQueueParameters struct {

	// +kubebuilder:validation:Required
	ConnectionStringSecretRef v1.SecretKeySelector `json:"connectionStringSecretRef" tf:"-"`

	// +kubebuilder:validation:Required
	IothubName *string `json:"iothubName" tf:"iothub_name,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-jet-azure/apis/azure/v1alpha1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`
}

// EndpointServicebusQueueSpec defines the desired state of EndpointServicebusQueue
type EndpointServicebusQueueSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     EndpointServicebusQueueParameters `json:"forProvider"`
}

// EndpointServicebusQueueStatus defines the observed state of EndpointServicebusQueue.
type EndpointServicebusQueueStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        EndpointServicebusQueueObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// EndpointServicebusQueue is the Schema for the EndpointServicebusQueues API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azurejet}
type EndpointServicebusQueue struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              EndpointServicebusQueueSpec   `json:"spec"`
	Status            EndpointServicebusQueueStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// EndpointServicebusQueueList contains a list of EndpointServicebusQueues
type EndpointServicebusQueueList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []EndpointServicebusQueue `json:"items"`
}

// Repository type metadata.
var (
	EndpointServicebusQueue_Kind             = "EndpointServicebusQueue"
	EndpointServicebusQueue_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: EndpointServicebusQueue_Kind}.String()
	EndpointServicebusQueue_KindAPIVersion   = EndpointServicebusQueue_Kind + "." + CRDGroupVersion.String()
	EndpointServicebusQueue_GroupVersionKind = CRDGroupVersion.WithKind(EndpointServicebusQueue_Kind)
)

func init() {
	SchemeBuilder.Register(&EndpointServicebusQueue{}, &EndpointServicebusQueueList{})
}
