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

type OutputSynapseObservation struct {

	// The ID of the Stream Analytics Output to an Azure Synapse Analytics Workspace.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type OutputSynapseParameters struct {

	// The name of the Azure SQL database. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Required
	Database *string `json:"database" tf:"database,omitempty"`

	// The password that will be used to connect to the Azure SQL database. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Required
	PasswordSecretRef v1.SecretKeySelector `json:"passwordSecretRef" tf:"-"`

	// The name of the Resource Group where the Stream Analytics Job exists. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// The name of the SQL server containing the Azure SQL database. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Required
	Server *string `json:"server" tf:"server,omitempty"`

	// The name of the Stream Analytics Job. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/streamanalytics/v1beta1.Job
	// +kubebuilder:validation:Optional
	StreamAnalyticsJobName *string `json:"streamAnalyticsJobName,omitempty" tf:"stream_analytics_job_name,omitempty"`

	// Reference to a Job in streamanalytics to populate streamAnalyticsJobName.
	// +kubebuilder:validation:Optional
	StreamAnalyticsJobNameRef *v1.Reference `json:"streamAnalyticsJobNameRef,omitempty" tf:"-"`

	// Selector for a Job in streamanalytics to populate streamAnalyticsJobName.
	// +kubebuilder:validation:Optional
	StreamAnalyticsJobNameSelector *v1.Selector `json:"streamAnalyticsJobNameSelector,omitempty" tf:"-"`

	// The name of the table in the Azure SQL database. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Required
	Table *string `json:"table" tf:"table,omitempty"`

	// The user name that will be used to connect to the Azure SQL database. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Required
	User *string `json:"user" tf:"user,omitempty"`
}

// OutputSynapseSpec defines the desired state of OutputSynapse
type OutputSynapseSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     OutputSynapseParameters `json:"forProvider"`
}

// OutputSynapseStatus defines the observed state of OutputSynapse.
type OutputSynapseStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        OutputSynapseObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// OutputSynapse is the Schema for the OutputSynapses API. Manages a Stream Analytics Output to an Azure Synapse Analytics Workspace.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type OutputSynapse struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              OutputSynapseSpec   `json:"spec"`
	Status            OutputSynapseStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// OutputSynapseList contains a list of OutputSynapses
type OutputSynapseList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []OutputSynapse `json:"items"`
}

// Repository type metadata.
var (
	OutputSynapse_Kind             = "OutputSynapse"
	OutputSynapse_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: OutputSynapse_Kind}.String()
	OutputSynapse_KindAPIVersion   = OutputSynapse_Kind + "." + CRDGroupVersion.String()
	OutputSynapse_GroupVersionKind = CRDGroupVersion.WithKind(OutputSynapse_Kind)
)

func init() {
	SchemeBuilder.Register(&OutputSynapse{}, &OutputSynapseList{})
}
