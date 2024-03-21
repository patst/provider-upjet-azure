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

type FlexibleServerActiveDirectoryAdministratorInitParameters struct {

	// The name of Azure Active Directory principal. Changing this forces a new resource to be created.
	PrincipalName *string `json:"principalName,omitempty" tf:"principal_name,omitempty"`

	// The type of Azure Active Directory principal. Possible values are Group, ServicePrincipal and User. Changing this forces a new resource to be created.
	PrincipalType *string `json:"principalType,omitempty" tf:"principal_type,omitempty"`

	// The Azure Tenant ID. Changing this forces a new resource to be created.
	TenantID *string `json:"tenantId,omitempty" tf:"tenant_id,omitempty"`
}

type FlexibleServerActiveDirectoryAdministratorObservation struct {

	// The ID of the PostgreSQL Flexible Server Active Directory Administrator.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The object ID of a user, service principal or security group in the Azure Active Directory tenant set as the Flexible Server Admin. Changing this forces a new resource to be created.
	ObjectID *string `json:"objectId,omitempty" tf:"object_id,omitempty"`

	// The name of Azure Active Directory principal. Changing this forces a new resource to be created.
	PrincipalName *string `json:"principalName,omitempty" tf:"principal_name,omitempty"`

	// The type of Azure Active Directory principal. Possible values are Group, ServicePrincipal and User. Changing this forces a new resource to be created.
	PrincipalType *string `json:"principalType,omitempty" tf:"principal_type,omitempty"`

	// The name of the resource group for the PostgreSQL Server. Changing this forces a new resource to be created.
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// The name of the PostgreSQL Flexible Server on which to set the administrator. Changing this forces a new resource to be created.
	ServerName *string `json:"serverName,omitempty" tf:"server_name,omitempty"`

	// The Azure Tenant ID. Changing this forces a new resource to be created.
	TenantID *string `json:"tenantId,omitempty" tf:"tenant_id,omitempty"`
}

type FlexibleServerActiveDirectoryAdministratorParameters struct {

	// The object ID of a user, service principal or security group in the Azure Active Directory tenant set as the Flexible Server Admin. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Required
	ObjectID *string `json:"objectId" tf:"object_id,omitempty"`

	// The name of Azure Active Directory principal. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	PrincipalName *string `json:"principalName,omitempty" tf:"principal_name,omitempty"`

	// The type of Azure Active Directory principal. Possible values are Group, ServicePrincipal and User. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	PrincipalType *string `json:"principalType,omitempty" tf:"principal_type,omitempty"`

	// The name of the resource group for the PostgreSQL Server. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// The name of the PostgreSQL Flexible Server on which to set the administrator. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=FlexibleServer
	// +kubebuilder:validation:Optional
	ServerName *string `json:"serverName,omitempty" tf:"server_name,omitempty"`

	// Reference to a FlexibleServer to populate serverName.
	// +kubebuilder:validation:Optional
	ServerNameRef *v1.Reference `json:"serverNameRef,omitempty" tf:"-"`

	// Selector for a FlexibleServer to populate serverName.
	// +kubebuilder:validation:Optional
	ServerNameSelector *v1.Selector `json:"serverNameSelector,omitempty" tf:"-"`

	// The Azure Tenant ID. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	TenantID *string `json:"tenantId,omitempty" tf:"tenant_id,omitempty"`
}

// FlexibleServerActiveDirectoryAdministratorSpec defines the desired state of FlexibleServerActiveDirectoryAdministrator
type FlexibleServerActiveDirectoryAdministratorSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     FlexibleServerActiveDirectoryAdministratorParameters `json:"forProvider"`
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
	InitProvider FlexibleServerActiveDirectoryAdministratorInitParameters `json:"initProvider,omitempty"`
}

// FlexibleServerActiveDirectoryAdministratorStatus defines the observed state of FlexibleServerActiveDirectoryAdministrator.
type FlexibleServerActiveDirectoryAdministratorStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        FlexibleServerActiveDirectoryAdministratorObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// FlexibleServerActiveDirectoryAdministrator is the Schema for the FlexibleServerActiveDirectoryAdministrators API. Manages an Active Directory administrator on a PostgreSQL Flexible server
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type FlexibleServerActiveDirectoryAdministrator struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.principalName) || (has(self.initProvider) && has(self.initProvider.principalName))",message="spec.forProvider.principalName is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.principalType) || (has(self.initProvider) && has(self.initProvider.principalType))",message="spec.forProvider.principalType is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.tenantId) || (has(self.initProvider) && has(self.initProvider.tenantId))",message="spec.forProvider.tenantId is a required parameter"
	Spec   FlexibleServerActiveDirectoryAdministratorSpec   `json:"spec"`
	Status FlexibleServerActiveDirectoryAdministratorStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// FlexibleServerActiveDirectoryAdministratorList contains a list of FlexibleServerActiveDirectoryAdministrators
type FlexibleServerActiveDirectoryAdministratorList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []FlexibleServerActiveDirectoryAdministrator `json:"items"`
}

// Repository type metadata.
var (
	FlexibleServerActiveDirectoryAdministrator_Kind             = "FlexibleServerActiveDirectoryAdministrator"
	FlexibleServerActiveDirectoryAdministrator_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: FlexibleServerActiveDirectoryAdministrator_Kind}.String()
	FlexibleServerActiveDirectoryAdministrator_KindAPIVersion   = FlexibleServerActiveDirectoryAdministrator_Kind + "." + CRDGroupVersion.String()
	FlexibleServerActiveDirectoryAdministrator_GroupVersionKind = CRDGroupVersion.WithKind(FlexibleServerActiveDirectoryAdministrator_Kind)
)

func init() {
	SchemeBuilder.Register(&FlexibleServerActiveDirectoryAdministrator{}, &FlexibleServerActiveDirectoryAdministratorList{})
}
