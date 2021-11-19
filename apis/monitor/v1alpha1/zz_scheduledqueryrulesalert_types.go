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

type ScheduledQueryRulesAlertActionObservation struct {
}

type ScheduledQueryRulesAlertActionParameters struct {

	// +kubebuilder:validation:Required
	ActionGroup []*string `json:"actionGroup" tf:"action_group,omitempty"`

	// +kubebuilder:validation:Optional
	CustomWebhookPayload *string `json:"customWebhookPayload,omitempty" tf:"custom_webhook_payload,omitempty"`

	// +kubebuilder:validation:Optional
	EmailSubject *string `json:"emailSubject,omitempty" tf:"email_subject,omitempty"`
}

type ScheduledQueryRulesAlertObservation struct {
}

type ScheduledQueryRulesAlertParameters struct {

	// +kubebuilder:validation:Required
	Action []ScheduledQueryRulesAlertActionParameters `json:"action" tf:"action,omitempty"`

	// +kubebuilder:validation:Optional
	AuthorizedResourceIds []*string `json:"authorizedResourceIds,omitempty" tf:"authorized_resource_ids,omitempty"`

	// +kubebuilder:validation:Optional
	AutoMitigationEnabled *bool `json:"autoMitigationEnabled,omitempty" tf:"auto_mitigation_enabled,omitempty"`

	// +kubebuilder:validation:Required
	DataSourceID *string `json:"dataSourceId" tf:"data_source_id,omitempty"`

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// +kubebuilder:validation:Required
	Frequency *int64 `json:"frequency" tf:"frequency,omitempty"`

	// +kubebuilder:validation:Required
	Location *string `json:"location" tf:"location,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Required
	Query *string `json:"query" tf:"query,omitempty"`

	// +kubebuilder:validation:Optional
	QueryType *string `json:"queryType,omitempty" tf:"query_type,omitempty"`

	// +kubebuilder:validation:Required
	ResourceGroupName *string `json:"resourceGroupName" tf:"resource_group_name,omitempty"`

	// +kubebuilder:validation:Optional
	Severity *int64 `json:"severity,omitempty" tf:"severity,omitempty"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// +kubebuilder:validation:Optional
	Throttling *int64 `json:"throttling,omitempty" tf:"throttling,omitempty"`

	// +kubebuilder:validation:Required
	TimeWindow *int64 `json:"timeWindow" tf:"time_window,omitempty"`

	// +kubebuilder:validation:Required
	Trigger []TriggerParameters `json:"trigger" tf:"trigger,omitempty"`
}

type TriggerMetricTriggerObservation struct {
}

type TriggerMetricTriggerParameters struct {

	// +kubebuilder:validation:Required
	MetricColumn *string `json:"metricColumn" tf:"metric_column,omitempty"`

	// +kubebuilder:validation:Required
	MetricTriggerType *string `json:"metricTriggerType" tf:"metric_trigger_type,omitempty"`

	// +kubebuilder:validation:Required
	Operator *string `json:"operator" tf:"operator,omitempty"`

	// +kubebuilder:validation:Required
	Threshold *float64 `json:"threshold" tf:"threshold,omitempty"`
}

type TriggerObservation struct {
}

type TriggerParameters struct {

	// +kubebuilder:validation:Optional
	MetricTrigger []TriggerMetricTriggerParameters `json:"metricTrigger,omitempty" tf:"metric_trigger,omitempty"`

	// +kubebuilder:validation:Required
	Operator *string `json:"operator" tf:"operator,omitempty"`

	// +kubebuilder:validation:Required
	Threshold *float64 `json:"threshold" tf:"threshold,omitempty"`
}

// ScheduledQueryRulesAlertSpec defines the desired state of ScheduledQueryRulesAlert
type ScheduledQueryRulesAlertSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ScheduledQueryRulesAlertParameters `json:"forProvider"`
}

// ScheduledQueryRulesAlertStatus defines the observed state of ScheduledQueryRulesAlert.
type ScheduledQueryRulesAlertStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ScheduledQueryRulesAlertObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ScheduledQueryRulesAlert is the Schema for the ScheduledQueryRulesAlerts API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azurejet}
type ScheduledQueryRulesAlert struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ScheduledQueryRulesAlertSpec   `json:"spec"`
	Status            ScheduledQueryRulesAlertStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ScheduledQueryRulesAlertList contains a list of ScheduledQueryRulesAlerts
type ScheduledQueryRulesAlertList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ScheduledQueryRulesAlert `json:"items"`
}

// Repository type metadata.
var (
	ScheduledQueryRulesAlert_Kind             = "ScheduledQueryRulesAlert"
	ScheduledQueryRulesAlert_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ScheduledQueryRulesAlert_Kind}.String()
	ScheduledQueryRulesAlert_KindAPIVersion   = ScheduledQueryRulesAlert_Kind + "." + CRDGroupVersion.String()
	ScheduledQueryRulesAlert_GroupVersionKind = CRDGroupVersion.WithKind(ScheduledQueryRulesAlert_Kind)
)

func init() {
	SchemeBuilder.Register(&ScheduledQueryRulesAlert{}, &ScheduledQueryRulesAlertList{})
}
