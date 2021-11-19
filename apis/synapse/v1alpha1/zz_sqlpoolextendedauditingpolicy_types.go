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

type SqlPoolExtendedAuditingPolicyObservation struct {
}

type SqlPoolExtendedAuditingPolicyParameters struct {

	// +kubebuilder:validation:Optional
	LogMonitoringEnabled *bool `json:"logMonitoringEnabled,omitempty" tf:"log_monitoring_enabled,omitempty"`

	// +kubebuilder:validation:Optional
	RetentionInDays *int64 `json:"retentionInDays,omitempty" tf:"retention_in_days,omitempty"`

	// +kubebuilder:validation:Required
	SQLPoolID *string `json:"sqlPoolId" tf:"sql_pool_id,omitempty"`

	// +kubebuilder:validation:Optional
	StorageAccountAccessKeyIsSecondary *bool `json:"storageAccountAccessKeyIsSecondary,omitempty" tf:"storage_account_access_key_is_secondary,omitempty"`

	// +kubebuilder:validation:Optional
	StorageAccountAccessKeySecretRef *v1.SecretKeySelector `json:"storageAccountAccessKeySecretRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	StorageEndpoint *string `json:"storageEndpoint,omitempty" tf:"storage_endpoint,omitempty"`
}

// SqlPoolExtendedAuditingPolicySpec defines the desired state of SqlPoolExtendedAuditingPolicy
type SqlPoolExtendedAuditingPolicySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SqlPoolExtendedAuditingPolicyParameters `json:"forProvider"`
}

// SqlPoolExtendedAuditingPolicyStatus defines the observed state of SqlPoolExtendedAuditingPolicy.
type SqlPoolExtendedAuditingPolicyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SqlPoolExtendedAuditingPolicyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// SqlPoolExtendedAuditingPolicy is the Schema for the SqlPoolExtendedAuditingPolicys API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azurejet}
type SqlPoolExtendedAuditingPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SqlPoolExtendedAuditingPolicySpec   `json:"spec"`
	Status            SqlPoolExtendedAuditingPolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SqlPoolExtendedAuditingPolicyList contains a list of SqlPoolExtendedAuditingPolicys
type SqlPoolExtendedAuditingPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SqlPoolExtendedAuditingPolicy `json:"items"`
}

// Repository type metadata.
var (
	SqlPoolExtendedAuditingPolicy_Kind             = "SqlPoolExtendedAuditingPolicy"
	SqlPoolExtendedAuditingPolicy_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: SqlPoolExtendedAuditingPolicy_Kind}.String()
	SqlPoolExtendedAuditingPolicy_KindAPIVersion   = SqlPoolExtendedAuditingPolicy_Kind + "." + CRDGroupVersion.String()
	SqlPoolExtendedAuditingPolicy_GroupVersionKind = CRDGroupVersion.WithKind(SqlPoolExtendedAuditingPolicy_Kind)
)

func init() {
	SchemeBuilder.Register(&SqlPoolExtendedAuditingPolicy{}, &SqlPoolExtendedAuditingPolicyList{})
}
