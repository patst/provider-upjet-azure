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

type FactoryInitParameters struct {

	// Specifies the Azure Key Vault Key ID to be used as the Customer Managed Key (CMK) for double encryption. Required with user assigned identity.
	CustomerManagedKeyID *string `json:"customerManagedKeyId,omitempty" tf:"customer_managed_key_id,omitempty"`

	// Specifies the ID of the user assigned identity associated with the Customer Managed Key. Must be supplied if customer_managed_key_id is set.
	CustomerManagedKeyIdentityID *string `json:"customerManagedKeyIdentityId,omitempty" tf:"customer_managed_key_identity_id,omitempty"`

	// A github_configuration block as defined below.
	GithubConfiguration []GithubConfigurationInitParameters `json:"githubConfiguration,omitempty" tf:"github_configuration,omitempty"`

	// A list of global_parameter blocks as defined above.
	GlobalParameter []GlobalParameterInitParameters `json:"globalParameter,omitempty" tf:"global_parameter,omitempty"`

	// An identity block as defined below.
	Identity []IdentityInitParameters `json:"identity,omitempty" tf:"identity,omitempty"`

	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// Is Managed Virtual Network enabled?
	ManagedVirtualNetworkEnabled *bool `json:"managedVirtualNetworkEnabled,omitempty" tf:"managed_virtual_network_enabled,omitempty"`

	// Is the Data Factory visible to the public network? Defaults to true.
	PublicNetworkEnabled *bool `json:"publicNetworkEnabled,omitempty" tf:"public_network_enabled,omitempty"`

	// Specifies the ID of the purview account resource associated with the Data Factory.
	PurviewID *string `json:"purviewId,omitempty" tf:"purview_id,omitempty"`

	// A mapping of tags to assign to the resource.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A vsts_configuration block as defined below.
	VstsConfiguration []VstsConfigurationInitParameters `json:"vstsConfiguration,omitempty" tf:"vsts_configuration,omitempty"`
}

type FactoryObservation struct {

	// Specifies the Azure Key Vault Key ID to be used as the Customer Managed Key (CMK) for double encryption. Required with user assigned identity.
	CustomerManagedKeyID *string `json:"customerManagedKeyId,omitempty" tf:"customer_managed_key_id,omitempty"`

	// Specifies the ID of the user assigned identity associated with the Customer Managed Key. Must be supplied if customer_managed_key_id is set.
	CustomerManagedKeyIdentityID *string `json:"customerManagedKeyIdentityId,omitempty" tf:"customer_managed_key_identity_id,omitempty"`

	// A github_configuration block as defined below.
	GithubConfiguration []GithubConfigurationObservation `json:"githubConfiguration,omitempty" tf:"github_configuration,omitempty"`

	// A list of global_parameter blocks as defined above.
	GlobalParameter []GlobalParameterObservation `json:"globalParameter,omitempty" tf:"global_parameter,omitempty"`

	// The ID of the Data Factory.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// An identity block as defined below.
	Identity []IdentityObservation `json:"identity,omitempty" tf:"identity,omitempty"`

	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// Is Managed Virtual Network enabled?
	ManagedVirtualNetworkEnabled *bool `json:"managedVirtualNetworkEnabled,omitempty" tf:"managed_virtual_network_enabled,omitempty"`

	// Is the Data Factory visible to the public network? Defaults to true.
	PublicNetworkEnabled *bool `json:"publicNetworkEnabled,omitempty" tf:"public_network_enabled,omitempty"`

	// Specifies the ID of the purview account resource associated with the Data Factory.
	PurviewID *string `json:"purviewId,omitempty" tf:"purview_id,omitempty"`

	// The name of the resource group in which to create the Data Factory. Changing this forces a new resource to be created.
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// A mapping of tags to assign to the resource.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A vsts_configuration block as defined below.
	VstsConfiguration []VstsConfigurationObservation `json:"vstsConfiguration,omitempty" tf:"vsts_configuration,omitempty"`
}

type FactoryParameters struct {

	// Specifies the Azure Key Vault Key ID to be used as the Customer Managed Key (CMK) for double encryption. Required with user assigned identity.
	// +kubebuilder:validation:Optional
	CustomerManagedKeyID *string `json:"customerManagedKeyId,omitempty" tf:"customer_managed_key_id,omitempty"`

	// Specifies the ID of the user assigned identity associated with the Customer Managed Key. Must be supplied if customer_managed_key_id is set.
	// +kubebuilder:validation:Optional
	CustomerManagedKeyIdentityID *string `json:"customerManagedKeyIdentityId,omitempty" tf:"customer_managed_key_identity_id,omitempty"`

	// A github_configuration block as defined below.
	// +kubebuilder:validation:Optional
	GithubConfiguration []GithubConfigurationParameters `json:"githubConfiguration,omitempty" tf:"github_configuration,omitempty"`

	// A list of global_parameter blocks as defined above.
	// +kubebuilder:validation:Optional
	GlobalParameter []GlobalParameterParameters `json:"globalParameter,omitempty" tf:"global_parameter,omitempty"`

	// An identity block as defined below.
	// +kubebuilder:validation:Optional
	Identity []IdentityParameters `json:"identity,omitempty" tf:"identity,omitempty"`

	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// Is Managed Virtual Network enabled?
	// +kubebuilder:validation:Optional
	ManagedVirtualNetworkEnabled *bool `json:"managedVirtualNetworkEnabled,omitempty" tf:"managed_virtual_network_enabled,omitempty"`

	// Is the Data Factory visible to the public network? Defaults to true.
	// +kubebuilder:validation:Optional
	PublicNetworkEnabled *bool `json:"publicNetworkEnabled,omitempty" tf:"public_network_enabled,omitempty"`

	// Specifies the ID of the purview account resource associated with the Data Factory.
	// +kubebuilder:validation:Optional
	PurviewID *string `json:"purviewId,omitempty" tf:"purview_id,omitempty"`

	// The name of the resource group in which to create the Data Factory. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// A mapping of tags to assign to the resource.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A vsts_configuration block as defined below.
	// +kubebuilder:validation:Optional
	VstsConfiguration []VstsConfigurationParameters `json:"vstsConfiguration,omitempty" tf:"vsts_configuration,omitempty"`
}

type GithubConfigurationInitParameters struct {

	// Specifies the GitHub account name.
	AccountName *string `json:"accountName,omitempty" tf:"account_name,omitempty"`

	// Specifies the branch of the repository to get code from.
	BranchName *string `json:"branchName,omitempty" tf:"branch_name,omitempty"`

	// Specifies the GitHub Enterprise host name. For example: https://github.mydomain.com. Use https://github.com for open source repositories.
	GitURL *string `json:"gitUrl,omitempty" tf:"git_url,omitempty"`

	// Is automated publishing enabled? Defaults to true.
	PublishingEnabled *bool `json:"publishingEnabled,omitempty" tf:"publishing_enabled,omitempty"`

	// Specifies the name of the git repository.
	RepositoryName *string `json:"repositoryName,omitempty" tf:"repository_name,omitempty"`

	// Specifies the root folder within the repository. Set to / for the top level.
	RootFolder *string `json:"rootFolder,omitempty" tf:"root_folder,omitempty"`
}

type GithubConfigurationObservation struct {

	// Specifies the GitHub account name.
	AccountName *string `json:"accountName,omitempty" tf:"account_name,omitempty"`

	// Specifies the branch of the repository to get code from.
	BranchName *string `json:"branchName,omitempty" tf:"branch_name,omitempty"`

	// Specifies the GitHub Enterprise host name. For example: https://github.mydomain.com. Use https://github.com for open source repositories.
	GitURL *string `json:"gitUrl,omitempty" tf:"git_url,omitempty"`

	// Is automated publishing enabled? Defaults to true.
	PublishingEnabled *bool `json:"publishingEnabled,omitempty" tf:"publishing_enabled,omitempty"`

	// Specifies the name of the git repository.
	RepositoryName *string `json:"repositoryName,omitempty" tf:"repository_name,omitempty"`

	// Specifies the root folder within the repository. Set to / for the top level.
	RootFolder *string `json:"rootFolder,omitempty" tf:"root_folder,omitempty"`
}

type GithubConfigurationParameters struct {

	// Specifies the GitHub account name.
	// +kubebuilder:validation:Optional
	AccountName *string `json:"accountName" tf:"account_name,omitempty"`

	// Specifies the branch of the repository to get code from.
	// +kubebuilder:validation:Optional
	BranchName *string `json:"branchName" tf:"branch_name,omitempty"`

	// Specifies the GitHub Enterprise host name. For example: https://github.mydomain.com. Use https://github.com for open source repositories.
	// +kubebuilder:validation:Optional
	GitURL *string `json:"gitUrl,omitempty" tf:"git_url,omitempty"`

	// Is automated publishing enabled? Defaults to true.
	// +kubebuilder:validation:Optional
	PublishingEnabled *bool `json:"publishingEnabled,omitempty" tf:"publishing_enabled,omitempty"`

	// Specifies the name of the git repository.
	// +kubebuilder:validation:Optional
	RepositoryName *string `json:"repositoryName" tf:"repository_name,omitempty"`

	// Specifies the root folder within the repository. Set to / for the top level.
	// +kubebuilder:validation:Optional
	RootFolder *string `json:"rootFolder" tf:"root_folder,omitempty"`
}

type GlobalParameterInitParameters struct {

	// Specifies the global parameter name.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Specifies the global parameter type. Possible Values are Array, Bool, Float, Int, Object or String.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// Specifies the global parameter value.
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type GlobalParameterObservation struct {

	// Specifies the global parameter name.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Specifies the global parameter type. Possible Values are Array, Bool, Float, Int, Object or String.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// Specifies the global parameter value.
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type GlobalParameterParameters struct {

	// Specifies the global parameter name.
	// +kubebuilder:validation:Optional
	Name *string `json:"name" tf:"name,omitempty"`

	// Specifies the global parameter type. Possible Values are Array, Bool, Float, Int, Object or String.
	// +kubebuilder:validation:Optional
	Type *string `json:"type" tf:"type,omitempty"`

	// Specifies the global parameter value.
	// +kubebuilder:validation:Optional
	Value *string `json:"value" tf:"value,omitempty"`
}

type IdentityInitParameters struct {

	// Specifies a list of User Assigned Managed Identity IDs to be assigned to this Data Factory.
	// +listType=set
	IdentityIds []*string `json:"identityIds,omitempty" tf:"identity_ids,omitempty"`

	// Specifies the type of Managed Service Identity that should be configured on this Data Factory. Possible values are SystemAssigned, UserAssigned, SystemAssigned, UserAssigned (to enable both).
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type IdentityObservation struct {

	// Specifies a list of User Assigned Managed Identity IDs to be assigned to this Data Factory.
	// +listType=set
	IdentityIds []*string `json:"identityIds,omitempty" tf:"identity_ids,omitempty"`

	// The Principal ID associated with this Managed Service Identity.
	PrincipalID *string `json:"principalId,omitempty" tf:"principal_id,omitempty"`

	// The Tenant ID associated with this Managed Service Identity.
	TenantID *string `json:"tenantId,omitempty" tf:"tenant_id,omitempty"`

	// Specifies the type of Managed Service Identity that should be configured on this Data Factory. Possible values are SystemAssigned, UserAssigned, SystemAssigned, UserAssigned (to enable both).
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type IdentityParameters struct {

	// Specifies a list of User Assigned Managed Identity IDs to be assigned to this Data Factory.
	// +kubebuilder:validation:Optional
	// +listType=set
	IdentityIds []*string `json:"identityIds,omitempty" tf:"identity_ids,omitempty"`

	// Specifies the type of Managed Service Identity that should be configured on this Data Factory. Possible values are SystemAssigned, UserAssigned, SystemAssigned, UserAssigned (to enable both).
	// +kubebuilder:validation:Optional
	Type *string `json:"type" tf:"type,omitempty"`
}

type VstsConfigurationInitParameters struct {

	// Specifies the VSTS account name.
	AccountName *string `json:"accountName,omitempty" tf:"account_name,omitempty"`

	// Specifies the branch of the repository to get code from.
	BranchName *string `json:"branchName,omitempty" tf:"branch_name,omitempty"`

	// Specifies the name of the VSTS project.
	ProjectName *string `json:"projectName,omitempty" tf:"project_name,omitempty"`

	// Is automated publishing enabled? Defaults to true.
	PublishingEnabled *bool `json:"publishingEnabled,omitempty" tf:"publishing_enabled,omitempty"`

	// Specifies the name of the git repository.
	RepositoryName *string `json:"repositoryName,omitempty" tf:"repository_name,omitempty"`

	// Specifies the root folder within the repository. Set to / for the top level.
	RootFolder *string `json:"rootFolder,omitempty" tf:"root_folder,omitempty"`

	// Specifies the Tenant ID associated with the VSTS account.
	TenantID *string `json:"tenantId,omitempty" tf:"tenant_id,omitempty"`
}

type VstsConfigurationObservation struct {

	// Specifies the VSTS account name.
	AccountName *string `json:"accountName,omitempty" tf:"account_name,omitempty"`

	// Specifies the branch of the repository to get code from.
	BranchName *string `json:"branchName,omitempty" tf:"branch_name,omitempty"`

	// Specifies the name of the VSTS project.
	ProjectName *string `json:"projectName,omitempty" tf:"project_name,omitempty"`

	// Is automated publishing enabled? Defaults to true.
	PublishingEnabled *bool `json:"publishingEnabled,omitempty" tf:"publishing_enabled,omitempty"`

	// Specifies the name of the git repository.
	RepositoryName *string `json:"repositoryName,omitempty" tf:"repository_name,omitempty"`

	// Specifies the root folder within the repository. Set to / for the top level.
	RootFolder *string `json:"rootFolder,omitempty" tf:"root_folder,omitempty"`

	// Specifies the Tenant ID associated with the VSTS account.
	TenantID *string `json:"tenantId,omitempty" tf:"tenant_id,omitempty"`
}

type VstsConfigurationParameters struct {

	// Specifies the VSTS account name.
	// +kubebuilder:validation:Optional
	AccountName *string `json:"accountName" tf:"account_name,omitempty"`

	// Specifies the branch of the repository to get code from.
	// +kubebuilder:validation:Optional
	BranchName *string `json:"branchName" tf:"branch_name,omitempty"`

	// Specifies the name of the VSTS project.
	// +kubebuilder:validation:Optional
	ProjectName *string `json:"projectName" tf:"project_name,omitempty"`

	// Is automated publishing enabled? Defaults to true.
	// +kubebuilder:validation:Optional
	PublishingEnabled *bool `json:"publishingEnabled,omitempty" tf:"publishing_enabled,omitempty"`

	// Specifies the name of the git repository.
	// +kubebuilder:validation:Optional
	RepositoryName *string `json:"repositoryName" tf:"repository_name,omitempty"`

	// Specifies the root folder within the repository. Set to / for the top level.
	// +kubebuilder:validation:Optional
	RootFolder *string `json:"rootFolder" tf:"root_folder,omitempty"`

	// Specifies the Tenant ID associated with the VSTS account.
	// +kubebuilder:validation:Optional
	TenantID *string `json:"tenantId" tf:"tenant_id,omitempty"`
}

// FactorySpec defines the desired state of Factory
type FactorySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     FactoryParameters `json:"forProvider"`
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
	InitProvider FactoryInitParameters `json:"initProvider,omitempty"`
}

// FactoryStatus defines the observed state of Factory.
type FactoryStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        FactoryObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// Factory is the Schema for the Factorys API. Manages an Azure Data Factory (Version 2).
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type Factory struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.location) || (has(self.initProvider) && has(self.initProvider.location))",message="spec.forProvider.location is a required parameter"
	Spec   FactorySpec   `json:"spec"`
	Status FactoryStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// FactoryList contains a list of Factorys
type FactoryList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Factory `json:"items"`
}

// Repository type metadata.
var (
	Factory_Kind             = "Factory"
	Factory_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Factory_Kind}.String()
	Factory_KindAPIVersion   = Factory_Kind + "." + CRDGroupVersion.String()
	Factory_GroupVersionKind = CRDGroupVersion.WithKind(Factory_Kind)
)

func init() {
	SchemeBuilder.Register(&Factory{}, &FactoryList{})
}
