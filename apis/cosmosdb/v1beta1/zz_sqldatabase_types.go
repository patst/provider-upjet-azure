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

type SQLDatabaseAutoscaleSettingsObservation struct {
}

type SQLDatabaseAutoscaleSettingsParameters struct {

	// The maximum throughput of the SQL database (RU/s). Must be between 1,000 and 1,000,000. Must be set in increments of 1,000. Conflicts with throughput.
	// +kubebuilder:validation:Optional
	MaxThroughput *float64 `json:"maxThroughput,omitempty" tf:"max_throughput,omitempty"`
}

type SQLDatabaseObservation struct {

	// The ID of the CosmosDB SQL Database.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type SQLDatabaseParameters struct {

	// The name of the Cosmos DB SQL Database to create the table within. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=Account
	// +kubebuilder:validation:Optional
	AccountName *string `json:"accountName,omitempty" tf:"account_name,omitempty"`

	// Reference to a Account to populate accountName.
	// +kubebuilder:validation:Optional
	AccountNameRef *v1.Reference `json:"accountNameRef,omitempty" tf:"-"`

	// Selector for a Account to populate accountName.
	// +kubebuilder:validation:Optional
	AccountNameSelector *v1.Selector `json:"accountNameSelector,omitempty" tf:"-"`

	// An autoscale_settings block as defined below.
	// +kubebuilder:validation:Optional
	AutoscaleSettings []SQLDatabaseAutoscaleSettingsParameters `json:"autoscaleSettings,omitempty" tf:"autoscale_settings,omitempty"`

	// The name of the resource group in which the Cosmos DB SQL Database is created. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// The throughput of SQL database (RU/s). Must be set in increments of 100. The minimum value is 400.  Do not set when azurerm_cosmosdb_account is configured with EnableServerless capability.
	// +kubebuilder:validation:Optional
	Throughput *float64 `json:"throughput,omitempty" tf:"throughput,omitempty"`
}

// SQLDatabaseSpec defines the desired state of SQLDatabase
type SQLDatabaseSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SQLDatabaseParameters `json:"forProvider"`
}

// SQLDatabaseStatus defines the observed state of SQLDatabase.
type SQLDatabaseStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SQLDatabaseObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// SQLDatabase is the Schema for the SQLDatabases API. Manages a SQL Database within a Cosmos DB Account.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type SQLDatabase struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SQLDatabaseSpec   `json:"spec"`
	Status            SQLDatabaseStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SQLDatabaseList contains a list of SQLDatabases
type SQLDatabaseList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SQLDatabase `json:"items"`
}

// Repository type metadata.
var (
	SQLDatabase_Kind             = "SQLDatabase"
	SQLDatabase_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: SQLDatabase_Kind}.String()
	SQLDatabase_KindAPIVersion   = SQLDatabase_Kind + "." + CRDGroupVersion.String()
	SQLDatabase_GroupVersionKind = CRDGroupVersion.WithKind(SQLDatabase_Kind)
)

func init() {
	SchemeBuilder.Register(&SQLDatabase{}, &SQLDatabaseList{})
}
