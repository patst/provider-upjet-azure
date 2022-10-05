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

type MSSQLServerDNSAliasObservation struct {

	// The fully qualified DNS record for alias.
	DNSRecord *string `json:"dnsRecord,omitempty" tf:"dns_record,omitempty"`

	// The ID of the MSSQL Server DNS Alias.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type MSSQLServerDNSAliasParameters struct {

	// The ID of the mssql server. Changing this forces a new MSSQL Server DNS Alias to be created.
	// +crossplane:generate:reference:type=MSSQLServer
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-azure/apis/rconfig.ExtractResourceID()
	// +kubebuilder:validation:Optional
	MSSQLServerID *string `json:"mssqlServerId,omitempty" tf:"mssql_server_id,omitempty"`

	// Reference to a MSSQLServer to populate mssqlServerId.
	// +kubebuilder:validation:Optional
	MSSQLServerIDRef *v1.Reference `json:"mssqlServerIdRef,omitempty" tf:"-"`

	// Selector for a MSSQLServer to populate mssqlServerId.
	// +kubebuilder:validation:Optional
	MSSQLServerIDSelector *v1.Selector `json:"mssqlServerIdSelector,omitempty" tf:"-"`
}

// MSSQLServerDNSAliasSpec defines the desired state of MSSQLServerDNSAlias
type MSSQLServerDNSAliasSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     MSSQLServerDNSAliasParameters `json:"forProvider"`
}

// MSSQLServerDNSAliasStatus defines the observed state of MSSQLServerDNSAlias.
type MSSQLServerDNSAliasStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        MSSQLServerDNSAliasObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// MSSQLServerDNSAlias is the Schema for the MSSQLServerDNSAliass API. Manages a MS SQL Server DNS Alias.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type MSSQLServerDNSAlias struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              MSSQLServerDNSAliasSpec   `json:"spec"`
	Status            MSSQLServerDNSAliasStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// MSSQLServerDNSAliasList contains a list of MSSQLServerDNSAliass
type MSSQLServerDNSAliasList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MSSQLServerDNSAlias `json:"items"`
}

// Repository type metadata.
var (
	MSSQLServerDNSAlias_Kind             = "MSSQLServerDNSAlias"
	MSSQLServerDNSAlias_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: MSSQLServerDNSAlias_Kind}.String()
	MSSQLServerDNSAlias_KindAPIVersion   = MSSQLServerDNSAlias_Kind + "." + CRDGroupVersion.String()
	MSSQLServerDNSAlias_GroupVersionKind = CRDGroupVersion.WithKind(MSSQLServerDNSAlias_Kind)
)

func init() {
	SchemeBuilder.Register(&MSSQLServerDNSAlias{}, &MSSQLServerDNSAliasList{})
}
