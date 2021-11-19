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

type DatabaseObservation struct {
	CreationDate *string `json:"creationDate,omitempty" tf:"creation_date,omitempty"`

	DefaultSecondaryLocation *string `json:"defaultSecondaryLocation,omitempty" tf:"default_secondary_location,omitempty"`

	Encryption *string `json:"encryption,omitempty" tf:"encryption,omitempty"`
}

type DatabaseParameters struct {

	// +kubebuilder:validation:Optional
	Collation *string `json:"collation,omitempty" tf:"collation,omitempty"`

	// +kubebuilder:validation:Optional
	CreateMode *string `json:"createMode,omitempty" tf:"create_mode,omitempty"`

	// +kubebuilder:validation:Optional
	Edition *string `json:"edition,omitempty" tf:"edition,omitempty"`

	// +kubebuilder:validation:Optional
	ElasticPoolName *string `json:"elasticPoolName,omitempty" tf:"elastic_pool_name,omitempty"`

	// +kubebuilder:validation:Optional
	ExtendedAuditingPolicy []ExtendedAuditingPolicyParameters `json:"extendedAuditingPolicy,omitempty" tf:"extended_auditing_policy,omitempty"`

	// +kubebuilder:validation:Optional
	Import []ImportParameters `json:"import,omitempty" tf:"import,omitempty"`

	// +kubebuilder:validation:Required
	Location *string `json:"location" tf:"location,omitempty"`

	// +kubebuilder:validation:Optional
	MaxSizeBytes *string `json:"maxSizeBytes,omitempty" tf:"max_size_bytes,omitempty"`

	// +kubebuilder:validation:Optional
	MaxSizeGb *string `json:"maxSizeGb,omitempty" tf:"max_size_gb,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	ReadScale *bool `json:"readScale,omitempty" tf:"read_scale,omitempty"`

	// +kubebuilder:validation:Optional
	RequestedServiceObjectiveID *string `json:"requestedServiceObjectiveId,omitempty" tf:"requested_service_objective_id,omitempty"`

	// +kubebuilder:validation:Optional
	RequestedServiceObjectiveName *string `json:"requestedServiceObjectiveName,omitempty" tf:"requested_service_objective_name,omitempty"`

	// +kubebuilder:validation:Required
	ResourceGroupName *string `json:"resourceGroupName" tf:"resource_group_name,omitempty"`

	// +kubebuilder:validation:Optional
	RestorePointInTime *string `json:"restorePointInTime,omitempty" tf:"restore_point_in_time,omitempty"`

	// +kubebuilder:validation:Required
	ServerName *string `json:"serverName" tf:"server_name,omitempty"`

	// +kubebuilder:validation:Optional
	SourceDatabaseDeletionDate *string `json:"sourceDatabaseDeletionDate,omitempty" tf:"source_database_deletion_date,omitempty"`

	// +kubebuilder:validation:Optional
	SourceDatabaseID *string `json:"sourceDatabaseId,omitempty" tf:"source_database_id,omitempty"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// +kubebuilder:validation:Optional
	ThreatDetectionPolicy []ThreatDetectionPolicyParameters `json:"threatDetectionPolicy,omitempty" tf:"threat_detection_policy,omitempty"`

	// +kubebuilder:validation:Optional
	ZoneRedundant *bool `json:"zoneRedundant,omitempty" tf:"zone_redundant,omitempty"`
}

type ExtendedAuditingPolicyObservation struct {
}

type ExtendedAuditingPolicyParameters struct {

	// +kubebuilder:validation:Optional
	LogMonitoringEnabled *bool `json:"logMonitoringEnabled,omitempty" tf:"log_monitoring_enabled,omitempty"`

	// +kubebuilder:validation:Optional
	RetentionInDays *int64 `json:"retentionInDays,omitempty" tf:"retention_in_days,omitempty"`

	// +kubebuilder:validation:Optional
	StorageAccountAccessKeyIsSecondary *bool `json:"storageAccountAccessKeyIsSecondary,omitempty" tf:"storage_account_access_key_is_secondary,omitempty"`

	// +kubebuilder:validation:Optional
	StorageAccountAccessKeySecretRef *v1.SecretKeySelector `json:"storageAccountAccessKeySecretRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	StorageEndpoint *string `json:"storageEndpoint,omitempty" tf:"storage_endpoint,omitempty"`
}

type ImportObservation struct {
}

type ImportParameters struct {

	// +kubebuilder:validation:Required
	AdministratorLogin *string `json:"administratorLogin" tf:"administrator_login,omitempty"`

	// +kubebuilder:validation:Required
	AdministratorLoginPasswordSecretRef v1.SecretKeySelector `json:"administratorLoginPasswordSecretRef" tf:"-"`

	// +kubebuilder:validation:Required
	AuthenticationType *string `json:"authenticationType" tf:"authentication_type,omitempty"`

	// +kubebuilder:validation:Optional
	OperationMode *string `json:"operationMode,omitempty" tf:"operation_mode,omitempty"`

	// +kubebuilder:validation:Required
	StorageKeySecretRef v1.SecretKeySelector `json:"storageKeySecretRef" tf:"-"`

	// +kubebuilder:validation:Required
	StorageKeyType *string `json:"storageKeyType" tf:"storage_key_type,omitempty"`

	// +kubebuilder:validation:Required
	StorageURI *string `json:"storageUri" tf:"storage_uri,omitempty"`
}

type ThreatDetectionPolicyObservation struct {
}

type ThreatDetectionPolicyParameters struct {

	// +kubebuilder:validation:Optional
	DisabledAlerts []*string `json:"disabledAlerts,omitempty" tf:"disabled_alerts,omitempty"`

	// +kubebuilder:validation:Optional
	EmailAccountAdmins *string `json:"emailAccountAdmins,omitempty" tf:"email_account_admins,omitempty"`

	// +kubebuilder:validation:Optional
	EmailAddresses []*string `json:"emailAddresses,omitempty" tf:"email_addresses,omitempty"`

	// +kubebuilder:validation:Optional
	RetentionDays *int64 `json:"retentionDays,omitempty" tf:"retention_days,omitempty"`

	// +kubebuilder:validation:Optional
	State *string `json:"state,omitempty" tf:"state,omitempty"`

	// +kubebuilder:validation:Optional
	StorageAccountAccessKeySecretRef *v1.SecretKeySelector `json:"storageAccountAccessKeySecretRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	StorageEndpoint *string `json:"storageEndpoint,omitempty" tf:"storage_endpoint,omitempty"`

	// +kubebuilder:validation:Optional
	UseServerDefault *string `json:"useServerDefault,omitempty" tf:"use_server_default,omitempty"`
}

// DatabaseSpec defines the desired state of Database
type DatabaseSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DatabaseParameters `json:"forProvider"`
}

// DatabaseStatus defines the observed state of Database.
type DatabaseStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DatabaseObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Database is the Schema for the Databases API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azurejet}
type Database struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DatabaseSpec   `json:"spec"`
	Status            DatabaseStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DatabaseList contains a list of Databases
type DatabaseList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Database `json:"items"`
}

// Repository type metadata.
var (
	Database_Kind             = "Database"
	Database_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Database_Kind}.String()
	Database_KindAPIVersion   = Database_Kind + "." + CRDGroupVersion.String()
	Database_GroupVersionKind = CRDGroupVersion.WithKind(Database_Kind)
)

func init() {
	SchemeBuilder.Register(&Database{}, &DatabaseList{})
}
