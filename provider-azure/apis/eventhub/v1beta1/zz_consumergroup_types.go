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

type ConsumerGroupObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type ConsumerGroupParameters struct {

	// +crossplane:generate:reference:type=EventHub
	// +kubebuilder:validation:Optional
	EventHubName *string `json:"eventhubName,omitempty" tf:"eventhub_name,omitempty"`

	// +kubebuilder:validation:Optional
	EventHubNameRef *v1.Reference `json:"eventhubNameRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	EventHubNameSelector *v1.Selector `json:"eventhubNameSelector,omitempty" tf:"-"`

	// +crossplane:generate:reference:type=EventHubNamespace
	// +kubebuilder:validation:Optional
	NamespaceName *string `json:"namespaceName,omitempty" tf:"namespace_name,omitempty"`

	// +kubebuilder:validation:Optional
	NamespaceNameRef *v1.Reference `json:"namespaceNameRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	NamespaceNameSelector *v1.Selector `json:"namespaceNameSelector,omitempty" tf:"-"`

	// +crossplane:generate:reference:type=github.com/upbound/official-providers/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	UserMetadata *string `json:"userMetadata,omitempty" tf:"user_metadata,omitempty"`
}

// ConsumerGroupSpec defines the desired state of ConsumerGroup
type ConsumerGroupSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ConsumerGroupParameters `json:"forProvider"`
}

// ConsumerGroupStatus defines the observed state of ConsumerGroup.
type ConsumerGroupStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ConsumerGroupObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ConsumerGroup is the Schema for the ConsumerGroups API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type ConsumerGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ConsumerGroupSpec   `json:"spec"`
	Status            ConsumerGroupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ConsumerGroupList contains a list of ConsumerGroups
type ConsumerGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ConsumerGroup `json:"items"`
}

// Repository type metadata.
var (
	ConsumerGroup_Kind             = "ConsumerGroup"
	ConsumerGroup_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ConsumerGroup_Kind}.String()
	ConsumerGroup_KindAPIVersion   = ConsumerGroup_Kind + "." + CRDGroupVersion.String()
	ConsumerGroup_GroupVersionKind = CRDGroupVersion.WithKind(ConsumerGroup_Kind)
)

func init() {
	SchemeBuilder.Register(&ConsumerGroup{}, &ConsumerGroupList{})
}
