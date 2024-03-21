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

type PrivateDNSPTRRecordInitParameters struct {

	// List of Fully Qualified Domain Names.
	// +listType=set
	Records []*string `json:"records,omitempty" tf:"records,omitempty"`

	// The Time To Live (TTL) of the DNS record in seconds.
	TTL *float64 `json:"ttl,omitempty" tf:"ttl,omitempty"`

	// A mapping of tags to assign to the resource.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type PrivateDNSPTRRecordObservation struct {

	// The FQDN of the DNS PTR Record.
	Fqdn *string `json:"fqdn,omitempty" tf:"fqdn,omitempty"`

	// The Private DNS PTR Record ID.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// List of Fully Qualified Domain Names.
	// +listType=set
	Records []*string `json:"records,omitempty" tf:"records,omitempty"`

	// Specifies the resource group where the resource exists. Changing this forces a new resource to be created.
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// The Time To Live (TTL) of the DNS record in seconds.
	TTL *float64 `json:"ttl,omitempty" tf:"ttl,omitempty"`

	// A mapping of tags to assign to the resource.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Specifies the Private DNS Zone where the resource exists. Changing this forces a new resource to be created.
	ZoneName *string `json:"zoneName,omitempty" tf:"zone_name,omitempty"`
}

type PrivateDNSPTRRecordParameters struct {

	// List of Fully Qualified Domain Names.
	// +kubebuilder:validation:Optional
	// +listType=set
	Records []*string `json:"records,omitempty" tf:"records,omitempty"`

	// Specifies the resource group where the resource exists. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// The Time To Live (TTL) of the DNS record in seconds.
	// +kubebuilder:validation:Optional
	TTL *float64 `json:"ttl,omitempty" tf:"ttl,omitempty"`

	// A mapping of tags to assign to the resource.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Specifies the Private DNS Zone where the resource exists. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=PrivateDNSZone
	// +kubebuilder:validation:Optional
	ZoneName *string `json:"zoneName,omitempty" tf:"zone_name,omitempty"`

	// Reference to a PrivateDNSZone to populate zoneName.
	// +kubebuilder:validation:Optional
	ZoneNameRef *v1.Reference `json:"zoneNameRef,omitempty" tf:"-"`

	// Selector for a PrivateDNSZone to populate zoneName.
	// +kubebuilder:validation:Optional
	ZoneNameSelector *v1.Selector `json:"zoneNameSelector,omitempty" tf:"-"`
}

// PrivateDNSPTRRecordSpec defines the desired state of PrivateDNSPTRRecord
type PrivateDNSPTRRecordSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     PrivateDNSPTRRecordParameters `json:"forProvider"`
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
	InitProvider PrivateDNSPTRRecordInitParameters `json:"initProvider,omitempty"`
}

// PrivateDNSPTRRecordStatus defines the observed state of PrivateDNSPTRRecord.
type PrivateDNSPTRRecordStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        PrivateDNSPTRRecordObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// PrivateDNSPTRRecord is the Schema for the PrivateDNSPTRRecords API. Manages a Private DNS PTR Record.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type PrivateDNSPTRRecord struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.records) || (has(self.initProvider) && has(self.initProvider.records))",message="spec.forProvider.records is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.ttl) || (has(self.initProvider) && has(self.initProvider.ttl))",message="spec.forProvider.ttl is a required parameter"
	Spec   PrivateDNSPTRRecordSpec   `json:"spec"`
	Status PrivateDNSPTRRecordStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// PrivateDNSPTRRecordList contains a list of PrivateDNSPTRRecords
type PrivateDNSPTRRecordList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PrivateDNSPTRRecord `json:"items"`
}

// Repository type metadata.
var (
	PrivateDNSPTRRecord_Kind             = "PrivateDNSPTRRecord"
	PrivateDNSPTRRecord_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: PrivateDNSPTRRecord_Kind}.String()
	PrivateDNSPTRRecord_KindAPIVersion   = PrivateDNSPTRRecord_Kind + "." + CRDGroupVersion.String()
	PrivateDNSPTRRecord_GroupVersionKind = CRDGroupVersion.WithKind(PrivateDNSPTRRecord_Kind)
)

func init() {
	SchemeBuilder.Register(&PrivateDNSPTRRecord{}, &PrivateDNSPTRRecordList{})
}
