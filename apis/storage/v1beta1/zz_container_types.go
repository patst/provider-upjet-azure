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

type ContainerInitParameters struct {

	// The Access Level configured for this Container. Possible values are blob, container or private. Defaults to private.
	ContainerAccessType *string `json:"containerAccessType,omitempty" tf:"container_access_type,omitempty"`

	// A mapping of MetaData for this Container. All metadata keys should be lowercase.
	// +mapType=granular
	Metadata map[string]*string `json:"metadata,omitempty" tf:"metadata,omitempty"`
}

type ContainerObservation struct {

	// The Access Level configured for this Container. Possible values are blob, container or private. Defaults to private.
	ContainerAccessType *string `json:"containerAccessType,omitempty" tf:"container_access_type,omitempty"`

	// Is there an Immutability Policy configured on this Storage Container?
	HasImmutabilityPolicy *bool `json:"hasImmutabilityPolicy,omitempty" tf:"has_immutability_policy,omitempty"`

	// Is there a Legal Hold configured on this Storage Container?
	HasLegalHold *bool `json:"hasLegalHold,omitempty" tf:"has_legal_hold,omitempty"`

	// The ID of the Storage Container.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// A mapping of MetaData for this Container. All metadata keys should be lowercase.
	// +mapType=granular
	Metadata map[string]*string `json:"metadata,omitempty" tf:"metadata,omitempty"`

	// The Resource Manager ID of this Storage Container.
	ResourceManagerID *string `json:"resourceManagerId,omitempty" tf:"resource_manager_id,omitempty"`

	// The name of the Storage Account where the Container should be created. Changing this forces a new resource to be created.
	StorageAccountName *string `json:"storageAccountName,omitempty" tf:"storage_account_name,omitempty"`
}

type ContainerParameters struct {

	// The Access Level configured for this Container. Possible values are blob, container or private. Defaults to private.
	// +kubebuilder:validation:Optional
	ContainerAccessType *string `json:"containerAccessType,omitempty" tf:"container_access_type,omitempty"`

	// A mapping of MetaData for this Container. All metadata keys should be lowercase.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Metadata map[string]*string `json:"metadata,omitempty" tf:"metadata,omitempty"`

	// The name of the Storage Account where the Container should be created. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=Account
	// +kubebuilder:validation:Optional
	StorageAccountName *string `json:"storageAccountName,omitempty" tf:"storage_account_name,omitempty"`

	// Reference to a Account to populate storageAccountName.
	// +kubebuilder:validation:Optional
	StorageAccountNameRef *v1.Reference `json:"storageAccountNameRef,omitempty" tf:"-"`

	// Selector for a Account to populate storageAccountName.
	// +kubebuilder:validation:Optional
	StorageAccountNameSelector *v1.Selector `json:"storageAccountNameSelector,omitempty" tf:"-"`
}

// ContainerSpec defines the desired state of Container
type ContainerSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ContainerParameters `json:"forProvider"`
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
	InitProvider ContainerInitParameters `json:"initProvider,omitempty"`
}

// ContainerStatus defines the observed state of Container.
type ContainerStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ContainerObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// Container is the Schema for the Containers API. Manages a Container within an Azure Storage Account.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type Container struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ContainerSpec   `json:"spec"`
	Status            ContainerStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ContainerList contains a list of Containers
type ContainerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Container `json:"items"`
}

// Repository type metadata.
var (
	Container_Kind             = "Container"
	Container_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Container_Kind}.String()
	Container_KindAPIVersion   = Container_Kind + "." + CRDGroupVersion.String()
	Container_GroupVersionKind = CRDGroupVersion.WithKind(Container_Kind)
)

func init() {
	SchemeBuilder.Register(&Container{}, &ContainerList{})
}
