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

type FirewallNATRuleCollectionObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type FirewallNATRuleCollectionParameters struct {

	// Specifies the action the rule will apply to matching traffic. Possible values are Dnat and Snat.
	// +kubebuilder:validation:Required
	Action *string `json:"action" tf:"action,omitempty"`

	// Specifies the name of the Firewall in which the NAT Rule Collection should be created. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=Firewall
	// +kubebuilder:validation:Optional
	AzureFirewallName *string `json:"azureFirewallName,omitempty" tf:"azure_firewall_name,omitempty"`

	// Reference to a Firewall to populate azureFirewallName.
	// +kubebuilder:validation:Optional
	AzureFirewallNameRef *v1.Reference `json:"azureFirewallNameRef,omitempty" tf:"-"`

	// Selector for a Firewall to populate azureFirewallName.
	// +kubebuilder:validation:Optional
	AzureFirewallNameSelector *v1.Selector `json:"azureFirewallNameSelector,omitempty" tf:"-"`

	// Specifies the priority of the rule collection. Possible values are between 100 - 65000.
	// +kubebuilder:validation:Required
	Priority *float64 `json:"priority" tf:"priority,omitempty"`

	// Specifies the name of the Resource Group in which the Firewall exists. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// One or more rule blocks as defined below.
	// +kubebuilder:validation:Required
	Rule []FirewallNATRuleCollectionRuleParameters `json:"rule" tf:"rule,omitempty"`
}

type FirewallNATRuleCollectionRuleObservation struct {
}

type FirewallNATRuleCollectionRuleParameters struct {

	// Specifies a description for the rule.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// A list of destination IP addresses and/or IP ranges.
	// +kubebuilder:validation:Required
	DestinationAddresses []*string `json:"destinationAddresses" tf:"destination_addresses,omitempty"`

	// A list of destination ports.
	// +kubebuilder:validation:Required
	DestinationPorts []*string `json:"destinationPorts" tf:"destination_ports,omitempty"`

	// Specifies the name of the rule.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// A list of protocols. Possible values are Any, ICMP, TCP and UDP.  If action is Dnat, protocols can only be TCP and UDP.
	// +kubebuilder:validation:Required
	Protocols []*string `json:"protocols" tf:"protocols,omitempty"`

	// A list of source IP addresses and/or IP ranges.
	// +kubebuilder:validation:Optional
	SourceAddresses []*string `json:"sourceAddresses,omitempty" tf:"source_addresses,omitempty"`

	// A list of source IP Group IDs for the rule.
	// +kubebuilder:validation:Optional
	SourceIPGroups []*string `json:"sourceIpGroups,omitempty" tf:"source_ip_groups,omitempty"`

	// The address of the service behind the Firewall.
	// +kubebuilder:validation:Required
	TranslatedAddress *string `json:"translatedAddress" tf:"translated_address,omitempty"`

	// The port of the service behind the Firewall.
	// +kubebuilder:validation:Required
	TranslatedPort *string `json:"translatedPort" tf:"translated_port,omitempty"`
}

// FirewallNATRuleCollectionSpec defines the desired state of FirewallNATRuleCollection
type FirewallNATRuleCollectionSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     FirewallNATRuleCollectionParameters `json:"forProvider"`
}

// FirewallNATRuleCollectionStatus defines the observed state of FirewallNATRuleCollection.
type FirewallNATRuleCollectionStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        FirewallNATRuleCollectionObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// FirewallNATRuleCollection is the Schema for the FirewallNATRuleCollections API. Manages a NAT Rule Collection within an Azure Firewall.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type FirewallNATRuleCollection struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              FirewallNATRuleCollectionSpec   `json:"spec"`
	Status            FirewallNATRuleCollectionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// FirewallNATRuleCollectionList contains a list of FirewallNATRuleCollections
type FirewallNATRuleCollectionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []FirewallNATRuleCollection `json:"items"`
}

// Repository type metadata.
var (
	FirewallNATRuleCollection_Kind             = "FirewallNATRuleCollection"
	FirewallNATRuleCollection_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: FirewallNATRuleCollection_Kind}.String()
	FirewallNATRuleCollection_KindAPIVersion   = FirewallNATRuleCollection_Kind + "." + CRDGroupVersion.String()
	FirewallNATRuleCollection_GroupVersionKind = CRDGroupVersion.WithKind(FirewallNATRuleCollection_Kind)
)

func init() {
	SchemeBuilder.Register(&FirewallNATRuleCollection{}, &FirewallNATRuleCollectionList{})
}
