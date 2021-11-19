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

type CloudAppRedisAssociationObservation struct {
}

type CloudAppRedisAssociationParameters struct {

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Required
	RedisAccessKey *string `json:"redisAccessKey" tf:"redis_access_key,omitempty"`

	// +kubebuilder:validation:Required
	RedisCacheID *string `json:"redisCacheId" tf:"redis_cache_id,omitempty"`

	// +kubebuilder:validation:Required
	SpringCloudAppID *string `json:"springCloudAppId" tf:"spring_cloud_app_id,omitempty"`

	// +kubebuilder:validation:Optional
	SslEnabled *bool `json:"sslEnabled,omitempty" tf:"ssl_enabled,omitempty"`
}

// CloudAppRedisAssociationSpec defines the desired state of CloudAppRedisAssociation
type CloudAppRedisAssociationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     CloudAppRedisAssociationParameters `json:"forProvider"`
}

// CloudAppRedisAssociationStatus defines the observed state of CloudAppRedisAssociation.
type CloudAppRedisAssociationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        CloudAppRedisAssociationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// CloudAppRedisAssociation is the Schema for the CloudAppRedisAssociations API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azurejet}
type CloudAppRedisAssociation struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CloudAppRedisAssociationSpec   `json:"spec"`
	Status            CloudAppRedisAssociationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CloudAppRedisAssociationList contains a list of CloudAppRedisAssociations
type CloudAppRedisAssociationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CloudAppRedisAssociation `json:"items"`
}

// Repository type metadata.
var (
	CloudAppRedisAssociation_Kind             = "CloudAppRedisAssociation"
	CloudAppRedisAssociation_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: CloudAppRedisAssociation_Kind}.String()
	CloudAppRedisAssociation_KindAPIVersion   = CloudAppRedisAssociation_Kind + "." + CRDGroupVersion.String()
	CloudAppRedisAssociation_GroupVersionKind = CRDGroupVersion.WithKind(CloudAppRedisAssociation_Kind)
)

func init() {
	SchemeBuilder.Register(&CloudAppRedisAssociation{}, &CloudAppRedisAssociationList{})
}
