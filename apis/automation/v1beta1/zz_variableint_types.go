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

type VariableIntInitParameters struct {

	// The description of the Automation Variable.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Specifies if the Automation Variable is encrypted. Defaults to false.
	Encrypted *bool `json:"encrypted,omitempty" tf:"encrypted,omitempty"`

	// The value of the Automation Variable as a integer.
	Value *float64 `json:"value,omitempty" tf:"value,omitempty"`
}

type VariableIntObservation struct {

	// The name of the automation account in which the Variable is created. Changing this forces a new resource to be created.
	AutomationAccountName *string `json:"automationAccountName,omitempty" tf:"automation_account_name,omitempty"`

	// The description of the Automation Variable.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Specifies if the Automation Variable is encrypted. Defaults to false.
	Encrypted *bool `json:"encrypted,omitempty" tf:"encrypted,omitempty"`

	// The ID of the Automation Variable.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The name of the resource group in which to create the Automation Variable. Changing this forces a new resource to be created.
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// The value of the Automation Variable as a integer.
	Value *float64 `json:"value,omitempty" tf:"value,omitempty"`
}

type VariableIntParameters struct {

	// The name of the automation account in which the Variable is created. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/automation/v1beta1.Account
	// +kubebuilder:validation:Optional
	AutomationAccountName *string `json:"automationAccountName,omitempty" tf:"automation_account_name,omitempty"`

	// Reference to a Account in automation to populate automationAccountName.
	// +kubebuilder:validation:Optional
	AutomationAccountNameRef *v1.Reference `json:"automationAccountNameRef,omitempty" tf:"-"`

	// Selector for a Account in automation to populate automationAccountName.
	// +kubebuilder:validation:Optional
	AutomationAccountNameSelector *v1.Selector `json:"automationAccountNameSelector,omitempty" tf:"-"`

	// The description of the Automation Variable.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Specifies if the Automation Variable is encrypted. Defaults to false.
	// +kubebuilder:validation:Optional
	Encrypted *bool `json:"encrypted,omitempty" tf:"encrypted,omitempty"`

	// The name of the resource group in which to create the Automation Variable. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// The value of the Automation Variable as a integer.
	// +kubebuilder:validation:Optional
	Value *float64 `json:"value,omitempty" tf:"value,omitempty"`
}

// VariableIntSpec defines the desired state of VariableInt
type VariableIntSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     VariableIntParameters `json:"forProvider"`
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
	InitProvider VariableIntInitParameters `json:"initProvider,omitempty"`
}

// VariableIntStatus defines the observed state of VariableInt.
type VariableIntStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        VariableIntObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// VariableInt is the Schema for the VariableInts API. Manages a integer variable in Azure Automation.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type VariableInt struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              VariableIntSpec   `json:"spec"`
	Status            VariableIntStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// VariableIntList contains a list of VariableInts
type VariableIntList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VariableInt `json:"items"`
}

// Repository type metadata.
var (
	VariableInt_Kind             = "VariableInt"
	VariableInt_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: VariableInt_Kind}.String()
	VariableInt_KindAPIVersion   = VariableInt_Kind + "." + CRDGroupVersion.String()
	VariableInt_GroupVersionKind = CRDGroupVersion.WithKind(VariableInt_Kind)
)

func init() {
	SchemeBuilder.Register(&VariableInt{}, &VariableIntList{})
}
