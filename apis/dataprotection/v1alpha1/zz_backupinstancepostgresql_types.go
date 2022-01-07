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

type BackupInstancePostgreSQLObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type BackupInstancePostgreSQLParameters struct {

	// +kubebuilder:validation:Required
	BackupPolicyID *string `json:"backupPolicyId" tf:"backup_policy_id,omitempty"`

	// +kubebuilder:validation:Optional
	DatabaseCredentialKeyVaultSecretID *string `json:"databaseCredentialKeyVaultSecretId,omitempty" tf:"database_credential_key_vault_secret_id,omitempty"`

	// +kubebuilder:validation:Required
	DatabaseID *string `json:"databaseId" tf:"database_id,omitempty"`

	// +kubebuilder:validation:Required
	Location *string `json:"location" tf:"location,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Required
	VaultID *string `json:"vaultId" tf:"vault_id,omitempty"`
}

// BackupInstancePostgreSQLSpec defines the desired state of BackupInstancePostgreSQL
type BackupInstancePostgreSQLSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     BackupInstancePostgreSQLParameters `json:"forProvider"`
}

// BackupInstancePostgreSQLStatus defines the observed state of BackupInstancePostgreSQL.
type BackupInstancePostgreSQLStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        BackupInstancePostgreSQLObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// BackupInstancePostgreSQL is the Schema for the BackupInstancePostgreSQLs API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azurejet}
type BackupInstancePostgreSQL struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              BackupInstancePostgreSQLSpec   `json:"spec"`
	Status            BackupInstancePostgreSQLStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// BackupInstancePostgreSQLList contains a list of BackupInstancePostgreSQLs
type BackupInstancePostgreSQLList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BackupInstancePostgreSQL `json:"items"`
}

// Repository type metadata.
var (
	BackupInstancePostgreSQL_Kind             = "BackupInstancePostgreSQL"
	BackupInstancePostgreSQL_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: BackupInstancePostgreSQL_Kind}.String()
	BackupInstancePostgreSQL_KindAPIVersion   = BackupInstancePostgreSQL_Kind + "." + CRDGroupVersion.String()
	BackupInstancePostgreSQL_GroupVersionKind = CRDGroupVersion.WithKind(BackupInstancePostgreSQL_Kind)
)

func init() {
	SchemeBuilder.Register(&BackupInstancePostgreSQL{}, &BackupInstancePostgreSQLList{})
}
