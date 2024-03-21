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

type BotChannelLineInitParameters struct {

	// One or more line_channel blocks as defined below.
	LineChannel []LineChannelInitParameters `json:"lineChannel,omitempty" tf:"line_channel,omitempty"`

	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`
}

type BotChannelLineObservation struct {

	// The name of the Bot Resource this channel will be associated with. Changing this forces a new resource to be created.
	BotName *string `json:"botName,omitempty" tf:"bot_name,omitempty"`

	// The ID of the Line Integration for a Bot Channel.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// One or more line_channel blocks as defined below.
	LineChannel []LineChannelParameters `json:"lineChannel,omitempty" tf:"line_channel,omitempty"`

	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// The name of the resource group where the Line Channel should be created. Changing this forces a new resource to be created.
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`
}

type BotChannelLineParameters struct {

	// The name of the Bot Resource this channel will be associated with. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/botservice/v1beta1.BotChannelsRegistration
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("name",false)
	// +kubebuilder:validation:Optional
	BotName *string `json:"botName,omitempty" tf:"bot_name,omitempty"`

	// Reference to a BotChannelsRegistration in botservice to populate botName.
	// +kubebuilder:validation:Optional
	BotNameRef *v1.Reference `json:"botNameRef,omitempty" tf:"-"`

	// Selector for a BotChannelsRegistration in botservice to populate botName.
	// +kubebuilder:validation:Optional
	BotNameSelector *v1.Selector `json:"botNameSelector,omitempty" tf:"-"`

	// One or more line_channel blocks as defined below.
	// +kubebuilder:validation:Optional
	LineChannel []LineChannelParameters `json:"lineChannel,omitempty" tf:"line_channel,omitempty"`

	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// The name of the resource group where the Line Channel should be created. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`
}

type LineChannelInitParameters struct {
}

type LineChannelObservation struct {
}

type LineChannelParameters struct {

	// The access token which is used to call the Line Channel API.
	// +kubebuilder:validation:Required
	AccessTokenSecretRef v1.SecretKeySelector `json:"accessTokenSecretRef" tf:"-"`

	// The secret which is used to access the Line Channel.
	// +kubebuilder:validation:Required
	SecretSecretRef v1.SecretKeySelector `json:"secretSecretRef" tf:"-"`
}

// BotChannelLineSpec defines the desired state of BotChannelLine
type BotChannelLineSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     BotChannelLineParameters `json:"forProvider"`
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
	InitProvider BotChannelLineInitParameters `json:"initProvider,omitempty"`
}

// BotChannelLineStatus defines the observed state of BotChannelLine.
type BotChannelLineStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        BotChannelLineObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// BotChannelLine is the Schema for the BotChannelLines API. Manages a Line integration for a Bot Channel
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type BotChannelLine struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.lineChannel) || (has(self.initProvider) && has(self.initProvider.lineChannel))",message="spec.forProvider.lineChannel is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.location) || (has(self.initProvider) && has(self.initProvider.location))",message="spec.forProvider.location is a required parameter"
	Spec   BotChannelLineSpec   `json:"spec"`
	Status BotChannelLineStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// BotChannelLineList contains a list of BotChannelLines
type BotChannelLineList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BotChannelLine `json:"items"`
}

// Repository type metadata.
var (
	BotChannelLine_Kind             = "BotChannelLine"
	BotChannelLine_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: BotChannelLine_Kind}.String()
	BotChannelLine_KindAPIVersion   = BotChannelLine_Kind + "." + CRDGroupVersion.String()
	BotChannelLine_GroupVersionKind = CRDGroupVersion.WithKind(BotChannelLine_Kind)
)

func init() {
	SchemeBuilder.Register(&BotChannelLine{}, &BotChannelLineList{})
}
