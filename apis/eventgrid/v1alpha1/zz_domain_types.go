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

type DomainObservation struct {
	Endpoint *string `json:"endpoint,omitempty" tf:"endpoint,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type DomainParameters struct {

	// +kubebuilder:validation:Optional
	Identity []IdentityParameters `json:"identity,omitempty" tf:"identity,omitempty"`

	// +kubebuilder:validation:Optional
	InboundIPRule []InboundIPRuleParameters `json:"inboundIpRule,omitempty" tf:"inbound_ip_rule,omitempty"`

	// +kubebuilder:validation:Optional
	InputMappingDefaultValues []InputMappingDefaultValuesParameters `json:"inputMappingDefaultValues,omitempty" tf:"input_mapping_default_values,omitempty"`

	// +kubebuilder:validation:Optional
	InputMappingFields []InputMappingFieldsParameters `json:"inputMappingFields,omitempty" tf:"input_mapping_fields,omitempty"`

	// +kubebuilder:validation:Optional
	InputSchema *string `json:"inputSchema,omitempty" tf:"input_schema,omitempty"`

	// +kubebuilder:validation:Required
	Location *string `json:"location" tf:"location,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	PublicNetworkAccessEnabled *bool `json:"publicNetworkAccessEnabled,omitempty" tf:"public_network_access_enabled,omitempty"`

	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-jet-azure/apis/azure/v1alpha2.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type IdentityObservation struct {
	PrincipalID *string `json:"principalId,omitempty" tf:"principal_id,omitempty"`

	TenantID *string `json:"tenantId,omitempty" tf:"tenant_id,omitempty"`
}

type IdentityParameters struct {

	// +kubebuilder:validation:Optional
	IdentityIds []*string `json:"identityIds,omitempty" tf:"identity_ids,omitempty"`

	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`
}

type InboundIPRuleObservation struct {
}

type InboundIPRuleParameters struct {

	// +kubebuilder:validation:Optional
	Action *string `json:"action,omitempty" tf:"action,omitempty"`

	// +kubebuilder:validation:Required
	IPMask *string `json:"ipMask" tf:"ip_mask,omitempty"`
}

type InputMappingDefaultValuesObservation struct {
}

type InputMappingDefaultValuesParameters struct {

	// +kubebuilder:validation:Optional
	DataVersion *string `json:"dataVersion,omitempty" tf:"data_version,omitempty"`

	// +kubebuilder:validation:Optional
	EventType *string `json:"eventType,omitempty" tf:"event_type,omitempty"`

	// +kubebuilder:validation:Optional
	Subject *string `json:"subject,omitempty" tf:"subject,omitempty"`
}

type InputMappingFieldsObservation struct {
}

type InputMappingFieldsParameters struct {

	// +kubebuilder:validation:Optional
	DataVersion *string `json:"dataVersion,omitempty" tf:"data_version,omitempty"`

	// +kubebuilder:validation:Optional
	EventTime *string `json:"eventTime,omitempty" tf:"event_time,omitempty"`

	// +kubebuilder:validation:Optional
	EventType *string `json:"eventType,omitempty" tf:"event_type,omitempty"`

	// +kubebuilder:validation:Optional
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// +kubebuilder:validation:Optional
	Subject *string `json:"subject,omitempty" tf:"subject,omitempty"`

	// +kubebuilder:validation:Optional
	Topic *string `json:"topic,omitempty" tf:"topic,omitempty"`
}

// DomainSpec defines the desired state of Domain
type DomainSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DomainParameters `json:"forProvider"`
}

// DomainStatus defines the observed state of Domain.
type DomainStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DomainObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Domain is the Schema for the Domains API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azurejet}
type Domain struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DomainSpec   `json:"spec"`
	Status            DomainStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DomainList contains a list of Domains
type DomainList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Domain `json:"items"`
}

// Repository type metadata.
var (
	Domain_Kind             = "Domain"
	Domain_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Domain_Kind}.String()
	Domain_KindAPIVersion   = Domain_Kind + "." + CRDGroupVersion.String()
	Domain_GroupVersionKind = CRDGroupVersion.WithKind(Domain_Kind)
)

func init() {
	SchemeBuilder.Register(&Domain{}, &DomainList{})
}
