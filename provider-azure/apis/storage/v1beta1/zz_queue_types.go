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

type QueueObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type QueueParameters struct {

	// +kubebuilder:validation:Optional
	Metadata map[string]*string `json:"metadata,omitempty" tf:"metadata,omitempty"`

	// +crossplane:generate:reference:type=github.com/upbound/official-providers/provider-azure/apis/storage/v1beta1.Account
	// +kubebuilder:validation:Optional
	StorageAccountName *string `json:"storageAccountName,omitempty" tf:"storage_account_name,omitempty"`

	// +kubebuilder:validation:Optional
	StorageAccountNameRef *v1.Reference `json:"storageAccountNameRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	StorageAccountNameSelector *v1.Selector `json:"storageAccountNameSelector,omitempty" tf:"-"`
}

// QueueSpec defines the desired state of Queue
type QueueSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     QueueParameters `json:"forProvider"`
}

// QueueStatus defines the observed state of Queue.
type QueueStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        QueueObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Queue is the Schema for the Queues API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type Queue struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              QueueSpec   `json:"spec"`
	Status            QueueStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// QueueList contains a list of Queues
type QueueList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Queue `json:"items"`
}

// Repository type metadata.
var (
	Queue_Kind             = "Queue"
	Queue_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Queue_Kind}.String()
	Queue_KindAPIVersion   = Queue_Kind + "." + CRDGroupVersion.String()
	Queue_GroupVersionKind = CRDGroupVersion.WithKind(Queue_Kind)
)

func init() {
	SchemeBuilder.Register(&Queue{}, &QueueList{})
}
