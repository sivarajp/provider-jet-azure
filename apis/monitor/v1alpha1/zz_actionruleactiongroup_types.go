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

type ActionRuleActionGroupObservation struct {
}

type ActionRuleActionGroupParameters struct {

	// +kubebuilder:validation:Required
	ActionGroupID *string `json:"actionGroupId" tf:"action_group_id,omitempty"`

	// +kubebuilder:validation:Optional
	Condition []ConditionParameters `json:"condition,omitempty" tf:"condition,omitempty"`

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Required
	ResourceGroupName *string `json:"resourceGroupName" tf:"resource_group_name,omitempty"`

	// +kubebuilder:validation:Optional
	Scope []ScopeParameters `json:"scope,omitempty" tf:"scope,omitempty"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type AlertContextObservation struct {
}

type AlertContextParameters struct {

	// +kubebuilder:validation:Required
	Operator *string `json:"operator" tf:"operator,omitempty"`

	// +kubebuilder:validation:Required
	Values []*string `json:"values" tf:"values,omitempty"`
}

type AlertRuleIDObservation struct {
}

type AlertRuleIDParameters struct {

	// +kubebuilder:validation:Required
	Operator *string `json:"operator" tf:"operator,omitempty"`

	// +kubebuilder:validation:Required
	Values []*string `json:"values" tf:"values,omitempty"`
}

type ConditionObservation struct {
}

type ConditionParameters struct {

	// +kubebuilder:validation:Optional
	AlertContext []AlertContextParameters `json:"alertContext,omitempty" tf:"alert_context,omitempty"`

	// +kubebuilder:validation:Optional
	AlertRuleID []AlertRuleIDParameters `json:"alertRuleId,omitempty" tf:"alert_rule_id,omitempty"`

	// +kubebuilder:validation:Optional
	Description []DescriptionParameters `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Optional
	Monitor []MonitorParameters `json:"monitor,omitempty" tf:"monitor,omitempty"`

	// +kubebuilder:validation:Optional
	MonitorService []MonitorServiceParameters `json:"monitorService,omitempty" tf:"monitor_service,omitempty"`

	// +kubebuilder:validation:Optional
	Severity []SeverityParameters `json:"severity,omitempty" tf:"severity,omitempty"`

	// +kubebuilder:validation:Optional
	TargetResourceType []TargetResourceTypeParameters `json:"targetResourceType,omitempty" tf:"target_resource_type,omitempty"`
}

type DescriptionObservation struct {
}

type DescriptionParameters struct {

	// +kubebuilder:validation:Required
	Operator *string `json:"operator" tf:"operator,omitempty"`

	// +kubebuilder:validation:Required
	Values []*string `json:"values" tf:"values,omitempty"`
}

type MonitorObservation struct {
}

type MonitorParameters struct {

	// +kubebuilder:validation:Required
	Operator *string `json:"operator" tf:"operator,omitempty"`

	// +kubebuilder:validation:Required
	Values []*string `json:"values" tf:"values,omitempty"`
}

type MonitorServiceObservation struct {
}

type MonitorServiceParameters struct {

	// +kubebuilder:validation:Required
	Operator *string `json:"operator" tf:"operator,omitempty"`

	// +kubebuilder:validation:Required
	Values []*string `json:"values" tf:"values,omitempty"`
}

type ScopeObservation struct {
}

type ScopeParameters struct {

	// +kubebuilder:validation:Required
	ResourceIds []*string `json:"resourceIds" tf:"resource_ids,omitempty"`

	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`
}

type SeverityObservation struct {
}

type SeverityParameters struct {

	// +kubebuilder:validation:Required
	Operator *string `json:"operator" tf:"operator,omitempty"`

	// +kubebuilder:validation:Required
	Values []*string `json:"values" tf:"values,omitempty"`
}

type TargetResourceTypeObservation struct {
}

type TargetResourceTypeParameters struct {

	// +kubebuilder:validation:Required
	Operator *string `json:"operator" tf:"operator,omitempty"`

	// +kubebuilder:validation:Required
	Values []*string `json:"values" tf:"values,omitempty"`
}

// ActionRuleActionGroupSpec defines the desired state of ActionRuleActionGroup
type ActionRuleActionGroupSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ActionRuleActionGroupParameters `json:"forProvider"`
}

// ActionRuleActionGroupStatus defines the observed state of ActionRuleActionGroup.
type ActionRuleActionGroupStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ActionRuleActionGroupObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ActionRuleActionGroup is the Schema for the ActionRuleActionGroups API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azurejet}
type ActionRuleActionGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ActionRuleActionGroupSpec   `json:"spec"`
	Status            ActionRuleActionGroupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ActionRuleActionGroupList contains a list of ActionRuleActionGroups
type ActionRuleActionGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ActionRuleActionGroup `json:"items"`
}

// Repository type metadata.
var (
	ActionRuleActionGroup_Kind             = "ActionRuleActionGroup"
	ActionRuleActionGroup_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ActionRuleActionGroup_Kind}.String()
	ActionRuleActionGroup_KindAPIVersion   = ActionRuleActionGroup_Kind + "." + CRDGroupVersion.String()
	ActionRuleActionGroup_GroupVersionKind = CRDGroupVersion.WithKind(ActionRuleActionGroup_Kind)
)

func init() {
	SchemeBuilder.Register(&ActionRuleActionGroup{}, &ActionRuleActionGroupList{})
}
