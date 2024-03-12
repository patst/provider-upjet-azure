// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type TopicAuthorizationRuleInitParameters struct {

	// Grants listen access to this this Authorization Rule. Defaults to false.
	Listen *bool `json:"listen,omitempty" tf:"listen,omitempty"`

	// Grants manage access to this this Authorization Rule. When this property is true - both listen and send must be too. Defaults to false.
	Manage *bool `json:"manage,omitempty" tf:"manage,omitempty"`

	// Grants send access to this this Authorization Rule. Defaults to false.
	Send *bool `json:"send,omitempty" tf:"send,omitempty"`
}

type TopicAuthorizationRuleObservation struct {

	// The ServiceBus Topic ID.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Grants listen access to this this Authorization Rule. Defaults to false.
	Listen *bool `json:"listen,omitempty" tf:"listen,omitempty"`

	// Grants manage access to this this Authorization Rule. When this property is true - both listen and send must be too. Defaults to false.
	Manage *bool `json:"manage,omitempty" tf:"manage,omitempty"`

	// Grants send access to this this Authorization Rule. Defaults to false.
	Send *bool `json:"send,omitempty" tf:"send,omitempty"`

	// Specifies the ID of the ServiceBus Topic. Changing this forces a new resource to be created.
	TopicID *string `json:"topicId,omitempty" tf:"topic_id,omitempty"`
}

type TopicAuthorizationRuleParameters struct {

	// Grants listen access to this this Authorization Rule. Defaults to false.
	// +kubebuilder:validation:Optional
	Listen *bool `json:"listen,omitempty" tf:"listen,omitempty"`

	// Grants manage access to this this Authorization Rule. When this property is true - both listen and send must be too. Defaults to false.
	// +kubebuilder:validation:Optional
	Manage *bool `json:"manage,omitempty" tf:"manage,omitempty"`

	// Grants send access to this this Authorization Rule. Defaults to false.
	// +kubebuilder:validation:Optional
	Send *bool `json:"send,omitempty" tf:"send,omitempty"`

	// Specifies the ID of the ServiceBus Topic. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/servicebus/v1beta1.Topic
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	TopicID *string `json:"topicId,omitempty" tf:"topic_id,omitempty"`

	// Reference to a Topic in servicebus to populate topicId.
	// +kubebuilder:validation:Optional
	TopicIDRef *v1.Reference `json:"topicIdRef,omitempty" tf:"-"`

	// Selector for a Topic in servicebus to populate topicId.
	// +kubebuilder:validation:Optional
	TopicIDSelector *v1.Selector `json:"topicIdSelector,omitempty" tf:"-"`
}

// TopicAuthorizationRuleSpec defines the desired state of TopicAuthorizationRule
type TopicAuthorizationRuleSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     TopicAuthorizationRuleParameters `json:"forProvider"`
	// THIS IS A BETA FIELD. It will be honored
	// unless the Management Policies feature flag is disabled.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider TopicAuthorizationRuleInitParameters `json:"initProvider,omitempty"`
}

// TopicAuthorizationRuleStatus defines the observed state of TopicAuthorizationRule.
type TopicAuthorizationRuleStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        TopicAuthorizationRuleObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// TopicAuthorizationRule is the Schema for the TopicAuthorizationRules API. Manages a ServiceBus Topic authorization Rule within a ServiceBus Topic.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type TopicAuthorizationRule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              TopicAuthorizationRuleSpec   `json:"spec"`
	Status            TopicAuthorizationRuleStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// TopicAuthorizationRuleList contains a list of TopicAuthorizationRules
type TopicAuthorizationRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TopicAuthorizationRule `json:"items"`
}

// Repository type metadata.
var (
	TopicAuthorizationRule_Kind             = "TopicAuthorizationRule"
	TopicAuthorizationRule_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: TopicAuthorizationRule_Kind}.String()
	TopicAuthorizationRule_KindAPIVersion   = TopicAuthorizationRule_Kind + "." + CRDGroupVersion.String()
	TopicAuthorizationRule_GroupVersionKind = CRDGroupVersion.WithKind(TopicAuthorizationRule_Kind)
)

func init() {
	SchemeBuilder.Register(&TopicAuthorizationRule{}, &TopicAuthorizationRuleList{})
}
