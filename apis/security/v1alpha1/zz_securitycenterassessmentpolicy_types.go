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

type SecurityCenterAssessmentPolicyObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type SecurityCenterAssessmentPolicyParameters struct {

	// +kubebuilder:validation:Optional
	Categories []*string `json:"categories,omitempty" tf:"categories,omitempty"`

	// +kubebuilder:validation:Required
	Description *string `json:"description" tf:"description,omitempty"`

	// +kubebuilder:validation:Required
	DisplayName *string `json:"displayName" tf:"display_name,omitempty"`

	// +kubebuilder:validation:Optional
	ImplementationEffort *string `json:"implementationEffort,omitempty" tf:"implementation_effort,omitempty"`

	// +kubebuilder:validation:Optional
	RemediationDescription *string `json:"remediationDescription,omitempty" tf:"remediation_description,omitempty"`

	// +kubebuilder:validation:Optional
	Severity *string `json:"severity,omitempty" tf:"severity,omitempty"`

	// +kubebuilder:validation:Optional
	Threats []*string `json:"threats,omitempty" tf:"threats,omitempty"`

	// +kubebuilder:validation:Optional
	UserImpact *string `json:"userImpact,omitempty" tf:"user_impact,omitempty"`
}

// SecurityCenterAssessmentPolicySpec defines the desired state of SecurityCenterAssessmentPolicy
type SecurityCenterAssessmentPolicySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SecurityCenterAssessmentPolicyParameters `json:"forProvider"`
}

// SecurityCenterAssessmentPolicyStatus defines the observed state of SecurityCenterAssessmentPolicy.
type SecurityCenterAssessmentPolicyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SecurityCenterAssessmentPolicyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// SecurityCenterAssessmentPolicy is the Schema for the SecurityCenterAssessmentPolicys API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azurejet}
type SecurityCenterAssessmentPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SecurityCenterAssessmentPolicySpec   `json:"spec"`
	Status            SecurityCenterAssessmentPolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SecurityCenterAssessmentPolicyList contains a list of SecurityCenterAssessmentPolicys
type SecurityCenterAssessmentPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SecurityCenterAssessmentPolicy `json:"items"`
}

// Repository type metadata.
var (
	SecurityCenterAssessmentPolicy_Kind             = "SecurityCenterAssessmentPolicy"
	SecurityCenterAssessmentPolicy_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: SecurityCenterAssessmentPolicy_Kind}.String()
	SecurityCenterAssessmentPolicy_KindAPIVersion   = SecurityCenterAssessmentPolicy_Kind + "." + CRDGroupVersion.String()
	SecurityCenterAssessmentPolicy_GroupVersionKind = CRDGroupVersion.WithKind(SecurityCenterAssessmentPolicy_Kind)
)

func init() {
	SchemeBuilder.Register(&SecurityCenterAssessmentPolicy{}, &SecurityCenterAssessmentPolicyList{})
}
