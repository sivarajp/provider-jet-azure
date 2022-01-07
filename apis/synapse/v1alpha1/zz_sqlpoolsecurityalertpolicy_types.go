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

type SQLPoolSecurityAlertPolicyObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type SQLPoolSecurityAlertPolicyParameters struct {

	// +kubebuilder:validation:Optional
	DisabledAlerts []*string `json:"disabledAlerts,omitempty" tf:"disabled_alerts,omitempty"`

	// +kubebuilder:validation:Optional
	EmailAccountAdminsEnabled *bool `json:"emailAccountAdminsEnabled,omitempty" tf:"email_account_admins_enabled,omitempty"`

	// +kubebuilder:validation:Optional
	EmailAddresses []*string `json:"emailAddresses,omitempty" tf:"email_addresses,omitempty"`

	// +kubebuilder:validation:Required
	PolicyState *string `json:"policyState" tf:"policy_state,omitempty"`

	// +kubebuilder:validation:Optional
	RetentionDays *int64 `json:"retentionDays,omitempty" tf:"retention_days,omitempty"`

	// +kubebuilder:validation:Required
	SQLPoolID *string `json:"sqlPoolId" tf:"sql_pool_id,omitempty"`

	// +kubebuilder:validation:Optional
	StorageAccountAccessKeySecretRef *v1.SecretKeySelector `json:"storageAccountAccessKeySecretRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	StorageEndpoint *string `json:"storageEndpoint,omitempty" tf:"storage_endpoint,omitempty"`
}

// SQLPoolSecurityAlertPolicySpec defines the desired state of SQLPoolSecurityAlertPolicy
type SQLPoolSecurityAlertPolicySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SQLPoolSecurityAlertPolicyParameters `json:"forProvider"`
}

// SQLPoolSecurityAlertPolicyStatus defines the observed state of SQLPoolSecurityAlertPolicy.
type SQLPoolSecurityAlertPolicyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SQLPoolSecurityAlertPolicyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// SQLPoolSecurityAlertPolicy is the Schema for the SQLPoolSecurityAlertPolicys API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azurejet}
type SQLPoolSecurityAlertPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SQLPoolSecurityAlertPolicySpec   `json:"spec"`
	Status            SQLPoolSecurityAlertPolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SQLPoolSecurityAlertPolicyList contains a list of SQLPoolSecurityAlertPolicys
type SQLPoolSecurityAlertPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SQLPoolSecurityAlertPolicy `json:"items"`
}

// Repository type metadata.
var (
	SQLPoolSecurityAlertPolicy_Kind             = "SQLPoolSecurityAlertPolicy"
	SQLPoolSecurityAlertPolicy_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: SQLPoolSecurityAlertPolicy_Kind}.String()
	SQLPoolSecurityAlertPolicy_KindAPIVersion   = SQLPoolSecurityAlertPolicy_Kind + "." + CRDGroupVersion.String()
	SQLPoolSecurityAlertPolicy_GroupVersionKind = CRDGroupVersion.WithKind(SQLPoolSecurityAlertPolicy_Kind)
)

func init() {
	SchemeBuilder.Register(&SQLPoolSecurityAlertPolicy{}, &SQLPoolSecurityAlertPolicyList{})
}
