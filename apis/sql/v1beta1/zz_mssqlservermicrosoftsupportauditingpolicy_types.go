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

type MSSQLServerMicrosoftSupportAuditingPolicyInitParameters struct {

	// Whether to enable the extended auditing policy. Possible values are true and false. Defaults to true.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// Enable audit events to Azure Monitor? To enable server audit events to Azure Monitor, please enable its main database audit events to Azure Monitor. Defaults to true.
	LogMonitoringEnabled *bool `json:"logMonitoringEnabled,omitempty" tf:"log_monitoring_enabled,omitempty"`
}

type MSSQLServerMicrosoftSupportAuditingPolicyObservation struct {

	// The blob storage endpoint (e.g. https://example.blob.core.windows.net). This blob storage will hold all Microsoft support auditing logs.
	BlobStorageEndpoint *string `json:"blobStorageEndpoint,omitempty" tf:"blob_storage_endpoint,omitempty"`

	// Whether to enable the extended auditing policy. Possible values are true and false. Defaults to true.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// The ID of the MS SQL Server Microsoft Support Auditing Policy.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Enable audit events to Azure Monitor? To enable server audit events to Azure Monitor, please enable its main database audit events to Azure Monitor. Defaults to true.
	LogMonitoringEnabled *bool `json:"logMonitoringEnabled,omitempty" tf:"log_monitoring_enabled,omitempty"`

	// The ID of the SQL Server to set the extended auditing policy. Changing this forces a new resource to be created.
	ServerID *string `json:"serverId,omitempty" tf:"server_id,omitempty"`
}

type MSSQLServerMicrosoftSupportAuditingPolicyParameters struct {

	// The blob storage endpoint (e.g. https://example.blob.core.windows.net). This blob storage will hold all Microsoft support auditing logs.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/storage/v1beta1.Account
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractParamPath("primary_blob_endpoint",true)
	// +kubebuilder:validation:Optional
	BlobStorageEndpoint *string `json:"blobStorageEndpoint,omitempty" tf:"blob_storage_endpoint,omitempty"`

	// Reference to a Account in storage to populate blobStorageEndpoint.
	// +kubebuilder:validation:Optional
	BlobStorageEndpointRef *v1.Reference `json:"blobStorageEndpointRef,omitempty" tf:"-"`

	// Selector for a Account in storage to populate blobStorageEndpoint.
	// +kubebuilder:validation:Optional
	BlobStorageEndpointSelector *v1.Selector `json:"blobStorageEndpointSelector,omitempty" tf:"-"`

	// Whether to enable the extended auditing policy. Possible values are true and false. Defaults to true.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// Enable audit events to Azure Monitor? To enable server audit events to Azure Monitor, please enable its main database audit events to Azure Monitor. Defaults to true.
	LogMonitoringEnabled *bool `json:"logMonitoringEnabled,omitempty" tf:"log_monitoring_enabled,omitempty"`

	// The ID of the SQL Server to set the extended auditing policy. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/sql/v1beta1.MSSQLServer
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	ServerID *string `json:"serverId,omitempty" tf:"server_id,omitempty"`

	// Reference to a MSSQLServer in sql to populate serverId.
	// +kubebuilder:validation:Optional
	ServerIDRef *v1.Reference `json:"serverIdRef,omitempty" tf:"-"`

	// Selector for a MSSQLServer in sql to populate serverId.
	// +kubebuilder:validation:Optional
	ServerIDSelector *v1.Selector `json:"serverIdSelector,omitempty" tf:"-"`

	// The access key to use for the auditing storage account.
	// +kubebuilder:validation:Optional
	StorageAccountAccessKeySecretRef *v1.SecretKeySelector `json:"storageAccountAccessKeySecretRef,omitempty" tf:"-"`

	// The ID of the Subscription containing the Storage Account.
	// +kubebuilder:validation:Optional
	StorageAccountSubscriptionIDSecretRef *v1.SecretKeySelector `json:"storageAccountSubscriptionIdSecretRef,omitempty" tf:"-"`
}

// MSSQLServerMicrosoftSupportAuditingPolicySpec defines the desired state of MSSQLServerMicrosoftSupportAuditingPolicy
type MSSQLServerMicrosoftSupportAuditingPolicySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     MSSQLServerMicrosoftSupportAuditingPolicyParameters `json:"forProvider"`
	// THIS IS AN ALPHA FIELD. Do not use it in production. It is not honored
	// unless the relevant Crossplane feature flag is enabled, and may be
	// changed or removed without notice.
	InitProvider MSSQLServerMicrosoftSupportAuditingPolicyInitParameters `json:"initProvider,omitempty"`
}

// MSSQLServerMicrosoftSupportAuditingPolicyStatus defines the observed state of MSSQLServerMicrosoftSupportAuditingPolicy.
type MSSQLServerMicrosoftSupportAuditingPolicyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        MSSQLServerMicrosoftSupportAuditingPolicyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// MSSQLServerMicrosoftSupportAuditingPolicy is the Schema for the MSSQLServerMicrosoftSupportAuditingPolicys API. Manages a MS SQL Server Microsoft Support Auditing Policy.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type MSSQLServerMicrosoftSupportAuditingPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              MSSQLServerMicrosoftSupportAuditingPolicySpec   `json:"spec"`
	Status            MSSQLServerMicrosoftSupportAuditingPolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// MSSQLServerMicrosoftSupportAuditingPolicyList contains a list of MSSQLServerMicrosoftSupportAuditingPolicys
type MSSQLServerMicrosoftSupportAuditingPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MSSQLServerMicrosoftSupportAuditingPolicy `json:"items"`
}

// Repository type metadata.
var (
	MSSQLServerMicrosoftSupportAuditingPolicy_Kind             = "MSSQLServerMicrosoftSupportAuditingPolicy"
	MSSQLServerMicrosoftSupportAuditingPolicy_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: MSSQLServerMicrosoftSupportAuditingPolicy_Kind}.String()
	MSSQLServerMicrosoftSupportAuditingPolicy_KindAPIVersion   = MSSQLServerMicrosoftSupportAuditingPolicy_Kind + "." + CRDGroupVersion.String()
	MSSQLServerMicrosoftSupportAuditingPolicy_GroupVersionKind = CRDGroupVersion.WithKind(MSSQLServerMicrosoftSupportAuditingPolicy_Kind)
)

func init() {
	SchemeBuilder.Register(&MSSQLServerMicrosoftSupportAuditingPolicy{}, &MSSQLServerMicrosoftSupportAuditingPolicyList{})
}
