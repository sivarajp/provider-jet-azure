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

type ServiceEnvironmentObservation struct {
	ConnectorEndpointIPAddresses []*string `json:"connectorEndpointIpAddresses,omitempty" tf:"connector_endpoint_ip_addresses,omitempty"`

	ConnectorOutboundIPAddresses []*string `json:"connectorOutboundIpAddresses,omitempty" tf:"connector_outbound_ip_addresses,omitempty"`

	WorkflowEndpointIPAddresses []*string `json:"workflowEndpointIpAddresses,omitempty" tf:"workflow_endpoint_ip_addresses,omitempty"`

	WorkflowOutboundIPAddresses []*string `json:"workflowOutboundIpAddresses,omitempty" tf:"workflow_outbound_ip_addresses,omitempty"`
}

type ServiceEnvironmentParameters struct {

	// +kubebuilder:validation:Required
	AccessEndpointType *string `json:"accessEndpointType" tf:"access_endpoint_type,omitempty"`

	// +kubebuilder:validation:Required
	Location *string `json:"location" tf:"location,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Required
	ResourceGroupName *string `json:"resourceGroupName" tf:"resource_group_name,omitempty"`

	// +kubebuilder:validation:Optional
	SkuName *string `json:"skuName,omitempty" tf:"sku_name,omitempty"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// +kubebuilder:validation:Required
	VirtualNetworkSubnetIds []*string `json:"virtualNetworkSubnetIds" tf:"virtual_network_subnet_ids,omitempty"`
}

// ServiceEnvironmentSpec defines the desired state of ServiceEnvironment
type ServiceEnvironmentSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ServiceEnvironmentParameters `json:"forProvider"`
}

// ServiceEnvironmentStatus defines the observed state of ServiceEnvironment.
type ServiceEnvironmentStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ServiceEnvironmentObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ServiceEnvironment is the Schema for the ServiceEnvironments API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azurejet}
type ServiceEnvironment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ServiceEnvironmentSpec   `json:"spec"`
	Status            ServiceEnvironmentStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ServiceEnvironmentList contains a list of ServiceEnvironments
type ServiceEnvironmentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ServiceEnvironment `json:"items"`
}

// Repository type metadata.
var (
	ServiceEnvironment_Kind             = "ServiceEnvironment"
	ServiceEnvironment_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ServiceEnvironment_Kind}.String()
	ServiceEnvironment_KindAPIVersion   = ServiceEnvironment_Kind + "." + CRDGroupVersion.String()
	ServiceEnvironment_GroupVersionKind = CRDGroupVersion.WithKind(ServiceEnvironment_Kind)
)

func init() {
	SchemeBuilder.Register(&ServiceEnvironment{}, &ServiceEnvironmentList{})
}
