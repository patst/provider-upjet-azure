/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type NetworkInterfaceSecurityGroupAssociationObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type NetworkInterfaceSecurityGroupAssociationParameters struct {

	// The ID of the Network Interface. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=NetworkInterface
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-azure/apis/rconfig.ExtractResourceID()
	// +kubebuilder:validation:Optional
	NetworkInterfaceID *string `json:"networkInterfaceId,omitempty" tf:"network_interface_id,omitempty"`

	// Reference to a NetworkInterface to populate networkInterfaceId.
	// +kubebuilder:validation:Optional
	NetworkInterfaceIDRef *v1.Reference `json:"networkInterfaceIdRef,omitempty" tf:"-"`

	// Selector for a NetworkInterface to populate networkInterfaceId.
	// +kubebuilder:validation:Optional
	NetworkInterfaceIDSelector *v1.Selector `json:"networkInterfaceIdSelector,omitempty" tf:"-"`

	// The ID of the Network Security Group which should be attached to the Network Interface. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=SecurityGroup
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-azure/apis/rconfig.ExtractResourceID()
	// +kubebuilder:validation:Optional
	NetworkSecurityGroupID *string `json:"networkSecurityGroupId,omitempty" tf:"network_security_group_id,omitempty"`

	// Reference to a SecurityGroup to populate networkSecurityGroupId.
	// +kubebuilder:validation:Optional
	NetworkSecurityGroupIDRef *v1.Reference `json:"networkSecurityGroupIdRef,omitempty" tf:"-"`

	// Selector for a SecurityGroup to populate networkSecurityGroupId.
	// +kubebuilder:validation:Optional
	NetworkSecurityGroupIDSelector *v1.Selector `json:"networkSecurityGroupIdSelector,omitempty" tf:"-"`
}

// NetworkInterfaceSecurityGroupAssociationSpec defines the desired state of NetworkInterfaceSecurityGroupAssociation
type NetworkInterfaceSecurityGroupAssociationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     NetworkInterfaceSecurityGroupAssociationParameters `json:"forProvider"`
}

// NetworkInterfaceSecurityGroupAssociationStatus defines the observed state of NetworkInterfaceSecurityGroupAssociation.
type NetworkInterfaceSecurityGroupAssociationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        NetworkInterfaceSecurityGroupAssociationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// NetworkInterfaceSecurityGroupAssociation is the Schema for the NetworkInterfaceSecurityGroupAssociations API. Manages the association between a Network Interface and a Network Security Group.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type NetworkInterfaceSecurityGroupAssociation struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              NetworkInterfaceSecurityGroupAssociationSpec   `json:"spec"`
	Status            NetworkInterfaceSecurityGroupAssociationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// NetworkInterfaceSecurityGroupAssociationList contains a list of NetworkInterfaceSecurityGroupAssociations
type NetworkInterfaceSecurityGroupAssociationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NetworkInterfaceSecurityGroupAssociation `json:"items"`
}

// Repository type metadata.
var (
	NetworkInterfaceSecurityGroupAssociation_Kind             = "NetworkInterfaceSecurityGroupAssociation"
	NetworkInterfaceSecurityGroupAssociation_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: NetworkInterfaceSecurityGroupAssociation_Kind}.String()
	NetworkInterfaceSecurityGroupAssociation_KindAPIVersion   = NetworkInterfaceSecurityGroupAssociation_Kind + "." + CRDGroupVersion.String()
	NetworkInterfaceSecurityGroupAssociation_GroupVersionKind = CRDGroupVersion.WithKind(NetworkInterfaceSecurityGroupAssociation_Kind)
)

func init() {
	SchemeBuilder.Register(&NetworkInterfaceSecurityGroupAssociation{}, &NetworkInterfaceSecurityGroupAssociationList{})
}
