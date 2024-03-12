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

type DedicatedHostInitParameters struct {

	// Should the Dedicated Host automatically be replaced in case of a Hardware Failure? Defaults to true.
	AutoReplaceOnFailure *bool `json:"autoReplaceOnFailure,omitempty" tf:"auto_replace_on_failure,omitempty"`

	// Specifies the software license type that will be applied to the VMs deployed on the Dedicated Host. Possible values are None, Windows_Server_Hybrid and Windows_Server_Perpetual. Defaults to None.
	LicenseType *string `json:"licenseType,omitempty" tf:"license_type,omitempty"`

	// Specify the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// Specify the fault domain of the Dedicated Host Group in which to create the Dedicated Host. Changing this forces a new resource to be created.
	PlatformFaultDomain *float64 `json:"platformFaultDomain,omitempty" tf:"platform_fault_domain,omitempty"`

	// Specify the SKU name of the Dedicated Host. Possible values are DADSv5-Type1, DASv4-Type1, DASv4-Type2, DASv5-Type1, DCSv2-Type1, DDSv4-Type1, DDSv4-Type2, DDSv5-Type1, DSv3-Type1, DSv3-Type2, DSv3-Type3, DSv3-Type4, DSv4-Type1, DSv4-Type2, DSv5-Type1, EADSv5-Type1, EASv4-Type1, EASv4-Type2, EASv5-Type1, EDSv4-Type1, EDSv4-Type2, EDSv5-Type1, ESv3-Type1, ESv3-Type2, ESv3-Type3, ESv3-Type4, ESv4-Type1, ESv4-Type2, ESv5-Type1, FSv2-Type2, FSv2-Type3, FSv2-Type4, FXmds-Type1, LSv2-Type1, LSv3-Type1, MDMSv2MedMem-Type1, MDSv2MedMem-Type1, MMSv2MedMem-Type1, MS-Type1, MSm-Type1, MSmv2-Type1, MSv2-Type1, MSv2MedMem-Type1, NVASv4-Type1 and NVSv3-Type1. Changing this forces a new resource to be created.
	SkuName *string `json:"skuName,omitempty" tf:"sku_name,omitempty"`

	// A mapping of tags to assign to the resource.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type DedicatedHostObservation struct {

	// Should the Dedicated Host automatically be replaced in case of a Hardware Failure? Defaults to true.
	AutoReplaceOnFailure *bool `json:"autoReplaceOnFailure,omitempty" tf:"auto_replace_on_failure,omitempty"`

	// Specifies the ID of the Dedicated Host Group where the Dedicated Host should exist. Changing this forces a new resource to be created.
	DedicatedHostGroupID *string `json:"dedicatedHostGroupId,omitempty" tf:"dedicated_host_group_id,omitempty"`

	// The ID of the Dedicated Host.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Specifies the software license type that will be applied to the VMs deployed on the Dedicated Host. Possible values are None, Windows_Server_Hybrid and Windows_Server_Perpetual. Defaults to None.
	LicenseType *string `json:"licenseType,omitempty" tf:"license_type,omitempty"`

	// Specify the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// Specify the fault domain of the Dedicated Host Group in which to create the Dedicated Host. Changing this forces a new resource to be created.
	PlatformFaultDomain *float64 `json:"platformFaultDomain,omitempty" tf:"platform_fault_domain,omitempty"`

	// Specify the SKU name of the Dedicated Host. Possible values are DADSv5-Type1, DASv4-Type1, DASv4-Type2, DASv5-Type1, DCSv2-Type1, DDSv4-Type1, DDSv4-Type2, DDSv5-Type1, DSv3-Type1, DSv3-Type2, DSv3-Type3, DSv3-Type4, DSv4-Type1, DSv4-Type2, DSv5-Type1, EADSv5-Type1, EASv4-Type1, EASv4-Type2, EASv5-Type1, EDSv4-Type1, EDSv4-Type2, EDSv5-Type1, ESv3-Type1, ESv3-Type2, ESv3-Type3, ESv3-Type4, ESv4-Type1, ESv4-Type2, ESv5-Type1, FSv2-Type2, FSv2-Type3, FSv2-Type4, FXmds-Type1, LSv2-Type1, LSv3-Type1, MDMSv2MedMem-Type1, MDSv2MedMem-Type1, MMSv2MedMem-Type1, MS-Type1, MSm-Type1, MSmv2-Type1, MSv2-Type1, MSv2MedMem-Type1, NVASv4-Type1 and NVSv3-Type1. Changing this forces a new resource to be created.
	SkuName *string `json:"skuName,omitempty" tf:"sku_name,omitempty"`

	// A mapping of tags to assign to the resource.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type DedicatedHostParameters struct {

	// Should the Dedicated Host automatically be replaced in case of a Hardware Failure? Defaults to true.
	// +kubebuilder:validation:Optional
	AutoReplaceOnFailure *bool `json:"autoReplaceOnFailure,omitempty" tf:"auto_replace_on_failure,omitempty"`

	// Specifies the ID of the Dedicated Host Group where the Dedicated Host should exist. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Required
	DedicatedHostGroupID *string `json:"dedicatedHostGroupId" tf:"dedicated_host_group_id,omitempty"`

	// Specifies the software license type that will be applied to the VMs deployed on the Dedicated Host. Possible values are None, Windows_Server_Hybrid and Windows_Server_Perpetual. Defaults to None.
	// +kubebuilder:validation:Optional
	LicenseType *string `json:"licenseType,omitempty" tf:"license_type,omitempty"`

	// Specify the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// Specify the fault domain of the Dedicated Host Group in which to create the Dedicated Host. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	PlatformFaultDomain *float64 `json:"platformFaultDomain,omitempty" tf:"platform_fault_domain,omitempty"`

	// Specify the SKU name of the Dedicated Host. Possible values are DADSv5-Type1, DASv4-Type1, DASv4-Type2, DASv5-Type1, DCSv2-Type1, DDSv4-Type1, DDSv4-Type2, DDSv5-Type1, DSv3-Type1, DSv3-Type2, DSv3-Type3, DSv3-Type4, DSv4-Type1, DSv4-Type2, DSv5-Type1, EADSv5-Type1, EASv4-Type1, EASv4-Type2, EASv5-Type1, EDSv4-Type1, EDSv4-Type2, EDSv5-Type1, ESv3-Type1, ESv3-Type2, ESv3-Type3, ESv3-Type4, ESv4-Type1, ESv4-Type2, ESv5-Type1, FSv2-Type2, FSv2-Type3, FSv2-Type4, FXmds-Type1, LSv2-Type1, LSv3-Type1, MDMSv2MedMem-Type1, MDSv2MedMem-Type1, MMSv2MedMem-Type1, MS-Type1, MSm-Type1, MSmv2-Type1, MSv2-Type1, MSv2MedMem-Type1, NVASv4-Type1 and NVSv3-Type1. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	SkuName *string `json:"skuName,omitempty" tf:"sku_name,omitempty"`

	// A mapping of tags to assign to the resource.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// DedicatedHostSpec defines the desired state of DedicatedHost
type DedicatedHostSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DedicatedHostParameters `json:"forProvider"`
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
	InitProvider DedicatedHostInitParameters `json:"initProvider,omitempty"`
}

// DedicatedHostStatus defines the observed state of DedicatedHost.
type DedicatedHostStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DedicatedHostObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// DedicatedHost is the Schema for the DedicatedHosts API. Manage a Dedicated Host within a Dedicated Host Group.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type DedicatedHost struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.location) || (has(self.initProvider) && has(self.initProvider.location))",message="spec.forProvider.location is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.platformFaultDomain) || (has(self.initProvider) && has(self.initProvider.platformFaultDomain))",message="spec.forProvider.platformFaultDomain is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.skuName) || (has(self.initProvider) && has(self.initProvider.skuName))",message="spec.forProvider.skuName is a required parameter"
	Spec   DedicatedHostSpec   `json:"spec"`
	Status DedicatedHostStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DedicatedHostList contains a list of DedicatedHosts
type DedicatedHostList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DedicatedHost `json:"items"`
}

// Repository type metadata.
var (
	DedicatedHost_Kind             = "DedicatedHost"
	DedicatedHost_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: DedicatedHost_Kind}.String()
	DedicatedHost_KindAPIVersion   = DedicatedHost_Kind + "." + CRDGroupVersion.String()
	DedicatedHost_GroupVersionKind = CRDGroupVersion.WithKind(DedicatedHost_Kind)
)

func init() {
	SchemeBuilder.Register(&DedicatedHost{}, &DedicatedHostList{})
}
