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

type ContentLinkHashObservation struct {
}

type ContentLinkHashParameters struct {

	// Specifies the hash algorithm used to hash the content.
	// +kubebuilder:validation:Required
	Algorithm *string `json:"algorithm" tf:"algorithm,omitempty"`

	// Specifies the expected hash value of the content.
	// +kubebuilder:validation:Required
	Value *string `json:"value" tf:"value,omitempty"`
}

type ContentLinkObservation struct {
}

type ContentLinkParameters struct {

	// A hash block as defined below.
	// +kubebuilder:validation:Optional
	Hash []ContentLinkHashParameters `json:"hash,omitempty" tf:"hash,omitempty"`

	// The URI of the runbook content.
	// +kubebuilder:validation:Required
	URI *string `json:"uri" tf:"uri,omitempty"`

	// Specifies the version of the content
	// +kubebuilder:validation:Optional
	Version *string `json:"version,omitempty" tf:"version,omitempty"`
}

type DraftObservation struct {
	CreationTime *string `json:"creationTime,omitempty" tf:"creation_time,omitempty"`

	LastModifiedTime *string `json:"lastModifiedTime,omitempty" tf:"last_modified_time,omitempty"`
}

type DraftParameters struct {

	// A publish_content_link block as defined above.
	// +kubebuilder:validation:Optional
	ContentLink []ContentLinkParameters `json:"contentLink,omitempty" tf:"content_link,omitempty"`

	// Whether the draft in edit mode.
	// +kubebuilder:validation:Optional
	EditModeEnabled *bool `json:"editModeEnabled,omitempty" tf:"edit_mode_enabled,omitempty"`

	// Specifies the output types of the runbook.
	// +kubebuilder:validation:Optional
	OutputTypes []*string `json:"outputTypes,omitempty" tf:"output_types,omitempty"`

	// A list of parameters block as defined below.
	// +kubebuilder:validation:Optional
	Parameters []ParametersParameters `json:"parameters,omitempty" tf:"parameters,omitempty"`
}

type JobScheduleObservation struct {
}

type JobScheduleParameters struct {

	// The Automation Runbook ID.
	// +kubebuilder:validation:Optional
	JobScheduleID *string `json:"jobScheduleId,omitempty" tf:"job_schedule_id"`

	// A list of parameters block as defined below.
	// +kubebuilder:validation:Optional
	Parameters map[string]*string `json:"parameters,omitempty" tf:"parameters"`

	// +kubebuilder:validation:Optional
	RunOn *string `json:"runOn,omitempty" tf:"run_on"`

	// Specifies the name of the Runbook. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	ScheduleName *string `json:"scheduleName,omitempty" tf:"schedule_name"`
}

type ParametersObservation struct {
}

type ParametersParameters struct {

	// Specifies the default value of the parameter.
	// +kubebuilder:validation:Optional
	DefaultValue *string `json:"defaultValue,omitempty" tf:"default_value,omitempty"`

	// The name of the parameter.
	// +kubebuilder:validation:Required
	Key *string `json:"key" tf:"key,omitempty"`

	// Whether this parameter is mandatory.
	// +kubebuilder:validation:Optional
	Mandatory *bool `json:"mandatory,omitempty" tf:"mandatory,omitempty"`

	// Specifies the position of the parameter.
	// +kubebuilder:validation:Optional
	Position *float64 `json:"position,omitempty" tf:"position,omitempty"`

	// Specifies the type of this parameter.
	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`
}

type PublishContentLinkHashObservation struct {
}

type PublishContentLinkHashParameters struct {

	// Specifies the hash algorithm used to hash the content.
	// +kubebuilder:validation:Required
	Algorithm *string `json:"algorithm" tf:"algorithm,omitempty"`

	// Specifies the expected hash value of the content.
	// +kubebuilder:validation:Required
	Value *string `json:"value" tf:"value,omitempty"`
}

type PublishContentLinkObservation struct {
}

type PublishContentLinkParameters struct {

	// A hash block as defined below.
	// +kubebuilder:validation:Optional
	Hash []PublishContentLinkHashParameters `json:"hash,omitempty" tf:"hash,omitempty"`

	// The URI of the runbook content.
	// +kubebuilder:validation:Required
	URI *string `json:"uri" tf:"uri,omitempty"`

	// Specifies the version of the content
	// +kubebuilder:validation:Optional
	Version *string `json:"version,omitempty" tf:"version,omitempty"`
}

type RunBookObservation struct {

	// A draft block as defined below .
	// +kubebuilder:validation:Optional
	Draft []DraftObservation `json:"draft,omitempty" tf:"draft,omitempty"`

	// The Automation Runbook ID.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type RunBookParameters struct {

	// The name of the automation account in which the Runbook is created. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/automation/v1beta1.Account
	// +kubebuilder:validation:Optional
	AutomationAccountName *string `json:"automationAccountName,omitempty" tf:"automation_account_name,omitempty"`

	// Reference to a Account in automation to populate automationAccountName.
	// +kubebuilder:validation:Optional
	AutomationAccountNameRef *v1.Reference `json:"automationAccountNameRef,omitempty" tf:"-"`

	// Selector for a Account in automation to populate automationAccountName.
	// +kubebuilder:validation:Optional
	AutomationAccountNameSelector *v1.Selector `json:"automationAccountNameSelector,omitempty" tf:"-"`

	// The desired content of the runbook.
	// +kubebuilder:validation:Optional
	Content *string `json:"content,omitempty" tf:"content,omitempty"`

	// A description for this credential.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// A draft block as defined below .
	// +kubebuilder:validation:Optional
	Draft []DraftParameters `json:"draft,omitempty" tf:"draft,omitempty"`

	// +kubebuilder:validation:Optional
	JobSchedule []JobScheduleParameters `json:"jobSchedule,omitempty" tf:"job_schedule,omitempty"`

	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Required
	Location *string `json:"location" tf:"location,omitempty"`

	// Specifies the activity-level tracing options of the runbook, available only for Graphical runbooks. Possible values are 0 for None, 9 for Basic, and 15 for Detailed. Must turn on Verbose logging in order to see the tracing.
	// +kubebuilder:validation:Optional
	LogActivityTraceLevel *float64 `json:"logActivityTraceLevel,omitempty" tf:"log_activity_trace_level,omitempty"`

	// Progress log option.
	// +kubebuilder:validation:Required
	LogProgress *bool `json:"logProgress" tf:"log_progress,omitempty"`

	// Verbose log option.
	// +kubebuilder:validation:Required
	LogVerbose *bool `json:"logVerbose" tf:"log_verbose,omitempty"`

	// Specifies the name of the Runbook. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// The published runbook content link.
	// +kubebuilder:validation:Optional
	PublishContentLink []PublishContentLinkParameters `json:"publishContentLink,omitempty" tf:"publish_content_link,omitempty"`

	// The name of the resource group in which the Runbook is created. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// The type of the runbook - can be either Graph, GraphPowerShell, GraphPowerShellWorkflow, PowerShellWorkflow, PowerShell, Python3, Python2 or Script. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Required
	RunBookType *string `json:"runbookType" tf:"runbook_type,omitempty"`

	// A mapping of tags to assign to the resource.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// RunBookSpec defines the desired state of RunBook
type RunBookSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     RunBookParameters `json:"forProvider"`
}

// RunBookStatus defines the observed state of RunBook.
type RunBookStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        RunBookObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// RunBook is the Schema for the RunBooks API. Manages a Automation Runbook.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type RunBook struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              RunBookSpec   `json:"spec"`
	Status            RunBookStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RunBookList contains a list of RunBooks
type RunBookList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RunBook `json:"items"`
}

// Repository type metadata.
var (
	RunBook_Kind             = "RunBook"
	RunBook_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: RunBook_Kind}.String()
	RunBook_KindAPIVersion   = RunBook_Kind + "." + CRDGroupVersion.String()
	RunBook_GroupVersionKind = CRDGroupVersion.WithKind(RunBook_Kind)
)

func init() {
	SchemeBuilder.Register(&RunBook{}, &RunBookList{})
}