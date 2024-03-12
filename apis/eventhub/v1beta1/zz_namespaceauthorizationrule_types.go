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

type NamespaceAuthorizationRuleInitParameters struct {

	// Grants listen access to this this Authorization Rule. Defaults to false.
	Listen *bool `json:"listen,omitempty" tf:"listen,omitempty"`

	// Grants manage access to this this Authorization Rule. When this property is true - both listen and send must be too. Defaults to false.
	Manage *bool `json:"manage,omitempty" tf:"manage,omitempty"`

	// Grants send access to this this Authorization Rule. Defaults to false.
	Send *bool `json:"send,omitempty" tf:"send,omitempty"`
}

type NamespaceAuthorizationRuleObservation struct {

	// The EventHub Namespace Authorization Rule ID.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Grants listen access to this this Authorization Rule. Defaults to false.
	Listen *bool `json:"listen,omitempty" tf:"listen,omitempty"`

	// Grants manage access to this this Authorization Rule. When this property is true - both listen and send must be too. Defaults to false.
	Manage *bool `json:"manage,omitempty" tf:"manage,omitempty"`

	// Specifies the name of the EventHub Namespace. Changing this forces a new resource to be created.
	NamespaceName *string `json:"namespaceName,omitempty" tf:"namespace_name,omitempty"`

	// The name of the resource group in which the EventHub Namespace exists. Changing this forces a new resource to be created.
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Grants send access to this this Authorization Rule. Defaults to false.
	Send *bool `json:"send,omitempty" tf:"send,omitempty"`
}

type NamespaceAuthorizationRuleParameters struct {

	// Grants listen access to this this Authorization Rule. Defaults to false.
	// +kubebuilder:validation:Optional
	Listen *bool `json:"listen,omitempty" tf:"listen,omitempty"`

	// Grants manage access to this this Authorization Rule. When this property is true - both listen and send must be too. Defaults to false.
	// +kubebuilder:validation:Optional
	Manage *bool `json:"manage,omitempty" tf:"manage,omitempty"`

	// Specifies the name of the EventHub Namespace. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/eventhub/v1beta1.EventHubNamespace
	// +kubebuilder:validation:Optional
	NamespaceName *string `json:"namespaceName,omitempty" tf:"namespace_name,omitempty"`

	// Reference to a EventHubNamespace in eventhub to populate namespaceName.
	// +kubebuilder:validation:Optional
	NamespaceNameRef *v1.Reference `json:"namespaceNameRef,omitempty" tf:"-"`

	// Selector for a EventHubNamespace in eventhub to populate namespaceName.
	// +kubebuilder:validation:Optional
	NamespaceNameSelector *v1.Selector `json:"namespaceNameSelector,omitempty" tf:"-"`

	// The name of the resource group in which the EventHub Namespace exists. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// Grants send access to this this Authorization Rule. Defaults to false.
	// +kubebuilder:validation:Optional
	Send *bool `json:"send,omitempty" tf:"send,omitempty"`
}

// NamespaceAuthorizationRuleSpec defines the desired state of NamespaceAuthorizationRule
type NamespaceAuthorizationRuleSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     NamespaceAuthorizationRuleParameters `json:"forProvider"`
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
	InitProvider NamespaceAuthorizationRuleInitParameters `json:"initProvider,omitempty"`
}

// NamespaceAuthorizationRuleStatus defines the observed state of NamespaceAuthorizationRule.
type NamespaceAuthorizationRuleStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        NamespaceAuthorizationRuleObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// NamespaceAuthorizationRule is the Schema for the NamespaceAuthorizationRules API. Manages an Authorization Rule for an Event Hub Namespace.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type NamespaceAuthorizationRule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              NamespaceAuthorizationRuleSpec   `json:"spec"`
	Status            NamespaceAuthorizationRuleStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// NamespaceAuthorizationRuleList contains a list of NamespaceAuthorizationRules
type NamespaceAuthorizationRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NamespaceAuthorizationRule `json:"items"`
}

// Repository type metadata.
var (
	NamespaceAuthorizationRule_Kind             = "NamespaceAuthorizationRule"
	NamespaceAuthorizationRule_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: NamespaceAuthorizationRule_Kind}.String()
	NamespaceAuthorizationRule_KindAPIVersion   = NamespaceAuthorizationRule_Kind + "." + CRDGroupVersion.String()
	NamespaceAuthorizationRule_GroupVersionKind = CRDGroupVersion.WithKind(NamespaceAuthorizationRule_Kind)
)

func init() {
	SchemeBuilder.Register(&NamespaceAuthorizationRule{}, &NamespaceAuthorizationRuleList{})
}
