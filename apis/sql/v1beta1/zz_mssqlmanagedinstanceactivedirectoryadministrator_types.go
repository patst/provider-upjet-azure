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

type MSSQLManagedInstanceActiveDirectoryAdministratorInitParameters struct {

	// When true, only permit logins from AAD users and administrators. When false, also allow local database users.
	AzureadAuthenticationOnly *bool `json:"azureadAuthenticationOnly,omitempty" tf:"azuread_authentication_only,omitempty"`

	// The login name of the principal to set as the Managed Instance Administrator.
	LoginUsername *string `json:"loginUsername,omitempty" tf:"login_username,omitempty"`

	// The Object ID of the principal to set as the Managed Instance Administrator.
	ObjectID *string `json:"objectId,omitempty" tf:"object_id,omitempty"`

	// The Azure Active Directory Tenant ID.
	TenantID *string `json:"tenantId,omitempty" tf:"tenant_id,omitempty"`
}

type MSSQLManagedInstanceActiveDirectoryAdministratorObservation struct {

	// When true, only permit logins from AAD users and administrators. When false, also allow local database users.
	AzureadAuthenticationOnly *bool `json:"azureadAuthenticationOnly,omitempty" tf:"azuread_authentication_only,omitempty"`

	// The ID of the SQL Managed Instance Active Directory Administrator.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The login name of the principal to set as the Managed Instance Administrator.
	LoginUsername *string `json:"loginUsername,omitempty" tf:"login_username,omitempty"`

	// The ID of the Azure SQL Managed Instance for which to set the administrator. Changing this forces a new resource to be created.
	ManagedInstanceID *string `json:"managedInstanceId,omitempty" tf:"managed_instance_id,omitempty"`

	// The Object ID of the principal to set as the Managed Instance Administrator.
	ObjectID *string `json:"objectId,omitempty" tf:"object_id,omitempty"`

	// The Azure Active Directory Tenant ID.
	TenantID *string `json:"tenantId,omitempty" tf:"tenant_id,omitempty"`
}

type MSSQLManagedInstanceActiveDirectoryAdministratorParameters struct {

	// When true, only permit logins from AAD users and administrators. When false, also allow local database users.
	// +kubebuilder:validation:Optional
	AzureadAuthenticationOnly *bool `json:"azureadAuthenticationOnly,omitempty" tf:"azuread_authentication_only,omitempty"`

	// The login name of the principal to set as the Managed Instance Administrator.
	// +kubebuilder:validation:Optional
	LoginUsername *string `json:"loginUsername,omitempty" tf:"login_username,omitempty"`

	// The ID of the Azure SQL Managed Instance for which to set the administrator. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=MSSQLManagedInstance
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-azure/apis/rconfig.ExtractResourceID()
	// +kubebuilder:validation:Optional
	ManagedInstanceID *string `json:"managedInstanceId,omitempty" tf:"managed_instance_id,omitempty"`

	// Reference to a MSSQLManagedInstance to populate managedInstanceId.
	// +kubebuilder:validation:Optional
	ManagedInstanceIDRef *v1.Reference `json:"managedInstanceIdRef,omitempty" tf:"-"`

	// Selector for a MSSQLManagedInstance to populate managedInstanceId.
	// +kubebuilder:validation:Optional
	ManagedInstanceIDSelector *v1.Selector `json:"managedInstanceIdSelector,omitempty" tf:"-"`

	// The Object ID of the principal to set as the Managed Instance Administrator.
	// +kubebuilder:validation:Optional
	ObjectID *string `json:"objectId,omitempty" tf:"object_id,omitempty"`

	// The Azure Active Directory Tenant ID.
	// +kubebuilder:validation:Optional
	TenantID *string `json:"tenantId,omitempty" tf:"tenant_id,omitempty"`
}

// MSSQLManagedInstanceActiveDirectoryAdministratorSpec defines the desired state of MSSQLManagedInstanceActiveDirectoryAdministrator
type MSSQLManagedInstanceActiveDirectoryAdministratorSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     MSSQLManagedInstanceActiveDirectoryAdministratorParameters `json:"forProvider"`
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
	InitProvider MSSQLManagedInstanceActiveDirectoryAdministratorInitParameters `json:"initProvider,omitempty"`
}

// MSSQLManagedInstanceActiveDirectoryAdministratorStatus defines the observed state of MSSQLManagedInstanceActiveDirectoryAdministrator.
type MSSQLManagedInstanceActiveDirectoryAdministratorStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        MSSQLManagedInstanceActiveDirectoryAdministratorObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// MSSQLManagedInstanceActiveDirectoryAdministrator is the Schema for the MSSQLManagedInstanceActiveDirectoryAdministrators API. Manages an Active Directory Administrator on a Microsoft Azure SQL Managed Instance
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type MSSQLManagedInstanceActiveDirectoryAdministrator struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.loginUsername) || (has(self.initProvider) && has(self.initProvider.loginUsername))",message="spec.forProvider.loginUsername is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.objectId) || (has(self.initProvider) && has(self.initProvider.objectId))",message="spec.forProvider.objectId is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.tenantId) || (has(self.initProvider) && has(self.initProvider.tenantId))",message="spec.forProvider.tenantId is a required parameter"
	Spec   MSSQLManagedInstanceActiveDirectoryAdministratorSpec   `json:"spec"`
	Status MSSQLManagedInstanceActiveDirectoryAdministratorStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// MSSQLManagedInstanceActiveDirectoryAdministratorList contains a list of MSSQLManagedInstanceActiveDirectoryAdministrators
type MSSQLManagedInstanceActiveDirectoryAdministratorList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MSSQLManagedInstanceActiveDirectoryAdministrator `json:"items"`
}

// Repository type metadata.
var (
	MSSQLManagedInstanceActiveDirectoryAdministrator_Kind             = "MSSQLManagedInstanceActiveDirectoryAdministrator"
	MSSQLManagedInstanceActiveDirectoryAdministrator_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: MSSQLManagedInstanceActiveDirectoryAdministrator_Kind}.String()
	MSSQLManagedInstanceActiveDirectoryAdministrator_KindAPIVersion   = MSSQLManagedInstanceActiveDirectoryAdministrator_Kind + "." + CRDGroupVersion.String()
	MSSQLManagedInstanceActiveDirectoryAdministrator_GroupVersionKind = CRDGroupVersion.WithKind(MSSQLManagedInstanceActiveDirectoryAdministrator_Kind)
)

func init() {
	SchemeBuilder.Register(&MSSQLManagedInstanceActiveDirectoryAdministrator{}, &MSSQLManagedInstanceActiveDirectoryAdministratorList{})
}
