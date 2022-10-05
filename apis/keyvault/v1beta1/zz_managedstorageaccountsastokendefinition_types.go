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

type ManagedStorageAccountSASTokenDefinitionObservation struct {

	// The ID of the Managed Storage Account SAS Definition.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The ID of the Secret that is created by Managed Storage Account SAS Definition.
	SecretID *string `json:"secretId,omitempty" tf:"secret_id,omitempty"`
}

type ManagedStorageAccountSASTokenDefinitionParameters struct {

	// The ID of the Managed Storage Account.
	// +crossplane:generate:reference:type=ManagedStorageAccount
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-azure/apis/rconfig.ExtractResourceID()
	// +kubebuilder:validation:Optional
	ManagedStorageAccountID *string `json:"managedStorageAccountId,omitempty" tf:"managed_storage_account_id,omitempty"`

	// Reference to a ManagedStorageAccount to populate managedStorageAccountId.
	// +kubebuilder:validation:Optional
	ManagedStorageAccountIDRef *v1.Reference `json:"managedStorageAccountIdRef,omitempty" tf:"-"`

	// Selector for a ManagedStorageAccount to populate managedStorageAccountId.
	// +kubebuilder:validation:Optional
	ManagedStorageAccountIDSelector *v1.Selector `json:"managedStorageAccountIdSelector,omitempty" tf:"-"`

	// The SAS definition token template signed with an arbitrary key. Tokens created according to the SAS definition will have the same properties as the template, but regenerated with a new validity period.
	// +kubebuilder:validation:Required
	SASTemplateURI *string `json:"sasTemplateUri" tf:"sas_template_uri,omitempty"`

	// The type of SAS token the SAS definition will create. Possible values are account and service.
	// +kubebuilder:validation:Required
	SASType *string `json:"sasType" tf:"sas_type,omitempty"`

	// A mapping of tags which should be assigned to the SAS Definition.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Validity period of SAS token. Value needs to be in ISO 8601 duration format.
	// +kubebuilder:validation:Required
	ValidityPeriod *string `json:"validityPeriod" tf:"validity_period,omitempty"`
}

// ManagedStorageAccountSASTokenDefinitionSpec defines the desired state of ManagedStorageAccountSASTokenDefinition
type ManagedStorageAccountSASTokenDefinitionSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ManagedStorageAccountSASTokenDefinitionParameters `json:"forProvider"`
}

// ManagedStorageAccountSASTokenDefinitionStatus defines the observed state of ManagedStorageAccountSASTokenDefinition.
type ManagedStorageAccountSASTokenDefinitionStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ManagedStorageAccountSASTokenDefinitionObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ManagedStorageAccountSASTokenDefinition is the Schema for the ManagedStorageAccountSASTokenDefinitions API. Manages a Key Vault Managed Storage Account SAS Definition.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type ManagedStorageAccountSASTokenDefinition struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ManagedStorageAccountSASTokenDefinitionSpec   `json:"spec"`
	Status            ManagedStorageAccountSASTokenDefinitionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ManagedStorageAccountSASTokenDefinitionList contains a list of ManagedStorageAccountSASTokenDefinitions
type ManagedStorageAccountSASTokenDefinitionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ManagedStorageAccountSASTokenDefinition `json:"items"`
}

// Repository type metadata.
var (
	ManagedStorageAccountSASTokenDefinition_Kind             = "ManagedStorageAccountSASTokenDefinition"
	ManagedStorageAccountSASTokenDefinition_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ManagedStorageAccountSASTokenDefinition_Kind}.String()
	ManagedStorageAccountSASTokenDefinition_KindAPIVersion   = ManagedStorageAccountSASTokenDefinition_Kind + "." + CRDGroupVersion.String()
	ManagedStorageAccountSASTokenDefinition_GroupVersionKind = CRDGroupVersion.WithKind(ManagedStorageAccountSASTokenDefinition_Kind)
)

func init() {
	SchemeBuilder.Register(&ManagedStorageAccountSASTokenDefinition{}, &ManagedStorageAccountSASTokenDefinitionList{})
}
