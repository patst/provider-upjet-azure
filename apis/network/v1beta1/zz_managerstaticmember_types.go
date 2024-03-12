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

type ManagerStaticMemberInitParameters struct {

	// Specifies the Resource ID of the Virtual Network using as the Static Member. Changing this forces a new Network Manager Static Member to be created.
	// +crossplane:generate:reference:type=VirtualNetwork
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-azure/apis/rconfig.ExtractResourceID()
	TargetVirtualNetworkID *string `json:"targetVirtualNetworkId,omitempty" tf:"target_virtual_network_id,omitempty"`

	// Reference to a VirtualNetwork to populate targetVirtualNetworkId.
	// +kubebuilder:validation:Optional
	TargetVirtualNetworkIDRef *v1.Reference `json:"targetVirtualNetworkIdRef,omitempty" tf:"-"`

	// Selector for a VirtualNetwork to populate targetVirtualNetworkId.
	// +kubebuilder:validation:Optional
	TargetVirtualNetworkIDSelector *v1.Selector `json:"targetVirtualNetworkIdSelector,omitempty" tf:"-"`
}

type ManagerStaticMemberObservation struct {

	// The ID of the Network Manager Static Member.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Specifies the ID of the Network Manager Group. Changing this forces a new Network Manager Static Member to be created.
	NetworkGroupID *string `json:"networkGroupId,omitempty" tf:"network_group_id,omitempty"`

	// The region of the Network Manager Static Member.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// Specifies the Resource ID of the Virtual Network using as the Static Member. Changing this forces a new Network Manager Static Member to be created.
	TargetVirtualNetworkID *string `json:"targetVirtualNetworkId,omitempty" tf:"target_virtual_network_id,omitempty"`
}

type ManagerStaticMemberParameters struct {

	// Specifies the ID of the Network Manager Group. Changing this forces a new Network Manager Static Member to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/network/v1beta1.ManagerNetworkGroup
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	NetworkGroupID *string `json:"networkGroupId,omitempty" tf:"network_group_id,omitempty"`

	// Reference to a ManagerNetworkGroup in network to populate networkGroupId.
	// +kubebuilder:validation:Optional
	NetworkGroupIDRef *v1.Reference `json:"networkGroupIdRef,omitempty" tf:"-"`

	// Selector for a ManagerNetworkGroup in network to populate networkGroupId.
	// +kubebuilder:validation:Optional
	NetworkGroupIDSelector *v1.Selector `json:"networkGroupIdSelector,omitempty" tf:"-"`

	// Specifies the Resource ID of the Virtual Network using as the Static Member. Changing this forces a new Network Manager Static Member to be created.
	// +crossplane:generate:reference:type=VirtualNetwork
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-azure/apis/rconfig.ExtractResourceID()
	// +kubebuilder:validation:Optional
	TargetVirtualNetworkID *string `json:"targetVirtualNetworkId,omitempty" tf:"target_virtual_network_id,omitempty"`

	// Reference to a VirtualNetwork to populate targetVirtualNetworkId.
	// +kubebuilder:validation:Optional
	TargetVirtualNetworkIDRef *v1.Reference `json:"targetVirtualNetworkIdRef,omitempty" tf:"-"`

	// Selector for a VirtualNetwork to populate targetVirtualNetworkId.
	// +kubebuilder:validation:Optional
	TargetVirtualNetworkIDSelector *v1.Selector `json:"targetVirtualNetworkIdSelector,omitempty" tf:"-"`
}

// ManagerStaticMemberSpec defines the desired state of ManagerStaticMember
type ManagerStaticMemberSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ManagerStaticMemberParameters `json:"forProvider"`
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
	InitProvider ManagerStaticMemberInitParameters `json:"initProvider,omitempty"`
}

// ManagerStaticMemberStatus defines the observed state of ManagerStaticMember.
type ManagerStaticMemberStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ManagerStaticMemberObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// ManagerStaticMember is the Schema for the ManagerStaticMembers API. Manages a Network Manager Static Member.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type ManagerStaticMember struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ManagerStaticMemberSpec   `json:"spec"`
	Status            ManagerStaticMemberStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ManagerStaticMemberList contains a list of ManagerStaticMembers
type ManagerStaticMemberList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ManagerStaticMember `json:"items"`
}

// Repository type metadata.
var (
	ManagerStaticMember_Kind             = "ManagerStaticMember"
	ManagerStaticMember_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ManagerStaticMember_Kind}.String()
	ManagerStaticMember_KindAPIVersion   = ManagerStaticMember_Kind + "." + CRDGroupVersion.String()
	ManagerStaticMember_GroupVersionKind = CRDGroupVersion.WithKind(ManagerStaticMember_Kind)
)

func init() {
	SchemeBuilder.Register(&ManagerStaticMember{}, &ManagerStaticMemberList{})
}
