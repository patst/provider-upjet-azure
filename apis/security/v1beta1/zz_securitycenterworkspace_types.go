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

type SecurityCenterWorkspaceInitParameters struct {

	// The scope of VMs to send their security data to the desired workspace, unless overridden by a setting with more specific scope.
	Scope *string `json:"scope,omitempty" tf:"scope,omitempty"`

	// The ID of the Log Analytics Workspace to save the data in.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/operationalinsights/v1beta1.Workspace
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	WorkspaceID *string `json:"workspaceId,omitempty" tf:"workspace_id,omitempty"`

	// Reference to a Workspace in operationalinsights to populate workspaceId.
	// +kubebuilder:validation:Optional
	WorkspaceIDRef *v1.Reference `json:"workspaceIdRef,omitempty" tf:"-"`

	// Selector for a Workspace in operationalinsights to populate workspaceId.
	// +kubebuilder:validation:Optional
	WorkspaceIDSelector *v1.Selector `json:"workspaceIdSelector,omitempty" tf:"-"`
}

type SecurityCenterWorkspaceObservation struct {

	// The Security Center Workspace ID.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The scope of VMs to send their security data to the desired workspace, unless overridden by a setting with more specific scope.
	Scope *string `json:"scope,omitempty" tf:"scope,omitempty"`

	// The ID of the Log Analytics Workspace to save the data in.
	WorkspaceID *string `json:"workspaceId,omitempty" tf:"workspace_id,omitempty"`
}

type SecurityCenterWorkspaceParameters struct {

	// The scope of VMs to send their security data to the desired workspace, unless overridden by a setting with more specific scope.
	// +kubebuilder:validation:Optional
	Scope *string `json:"scope,omitempty" tf:"scope,omitempty"`

	// The ID of the Log Analytics Workspace to save the data in.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/operationalinsights/v1beta1.Workspace
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	WorkspaceID *string `json:"workspaceId,omitempty" tf:"workspace_id,omitempty"`

	// Reference to a Workspace in operationalinsights to populate workspaceId.
	// +kubebuilder:validation:Optional
	WorkspaceIDRef *v1.Reference `json:"workspaceIdRef,omitempty" tf:"-"`

	// Selector for a Workspace in operationalinsights to populate workspaceId.
	// +kubebuilder:validation:Optional
	WorkspaceIDSelector *v1.Selector `json:"workspaceIdSelector,omitempty" tf:"-"`
}

// SecurityCenterWorkspaceSpec defines the desired state of SecurityCenterWorkspace
type SecurityCenterWorkspaceSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SecurityCenterWorkspaceParameters `json:"forProvider"`
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
	InitProvider SecurityCenterWorkspaceInitParameters `json:"initProvider,omitempty"`
}

// SecurityCenterWorkspaceStatus defines the observed state of SecurityCenterWorkspace.
type SecurityCenterWorkspaceStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SecurityCenterWorkspaceObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// SecurityCenterWorkspace is the Schema for the SecurityCenterWorkspaces API. Manages the subscription's Security Center Workspace.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type SecurityCenterWorkspace struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.scope) || (has(self.initProvider) && has(self.initProvider.scope))",message="spec.forProvider.scope is a required parameter"
	Spec   SecurityCenterWorkspaceSpec   `json:"spec"`
	Status SecurityCenterWorkspaceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SecurityCenterWorkspaceList contains a list of SecurityCenterWorkspaces
type SecurityCenterWorkspaceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SecurityCenterWorkspace `json:"items"`
}

// Repository type metadata.
var (
	SecurityCenterWorkspace_Kind             = "SecurityCenterWorkspace"
	SecurityCenterWorkspace_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: SecurityCenterWorkspace_Kind}.String()
	SecurityCenterWorkspace_KindAPIVersion   = SecurityCenterWorkspace_Kind + "." + CRDGroupVersion.String()
	SecurityCenterWorkspace_GroupVersionKind = CRDGroupVersion.WithKind(SecurityCenterWorkspace_Kind)
)

func init() {
	SchemeBuilder.Register(&SecurityCenterWorkspace{}, &SecurityCenterWorkspaceList{})
}
