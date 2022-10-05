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

type LoadBalancerOutboundRuleFrontendIPConfigurationObservation struct {

	// The ID of the Load Balancer Outbound Rule.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type LoadBalancerOutboundRuleFrontendIPConfigurationParameters struct {

	// The name of the Frontend IP Configuration.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`
}

type LoadBalancerOutboundRuleObservation struct {

	// One or more frontend_ip_configuration blocks as defined below.
	// +kubebuilder:validation:Optional
	FrontendIPConfiguration []LoadBalancerOutboundRuleFrontendIPConfigurationObservation `json:"frontendIpConfiguration,omitempty" tf:"frontend_ip_configuration,omitempty"`

	// The ID of the Load Balancer Outbound Rule.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type LoadBalancerOutboundRuleParameters struct {

	// The number of outbound ports to be used for NAT. Defaults to 1024.
	// +kubebuilder:validation:Optional
	AllocatedOutboundPorts *float64 `json:"allocatedOutboundPorts,omitempty" tf:"allocated_outbound_ports,omitempty"`

	// The ID of the Backend Address Pool. Outbound traffic is randomly load balanced across IPs in the backend IPs.
	// +crossplane:generate:reference:type=LoadBalancerBackendAddressPool
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-azure/apis/rconfig.ExtractResourceID()
	// +kubebuilder:validation:Optional
	BackendAddressPoolID *string `json:"backendAddressPoolId,omitempty" tf:"backend_address_pool_id,omitempty"`

	// Reference to a LoadBalancerBackendAddressPool to populate backendAddressPoolId.
	// +kubebuilder:validation:Optional
	BackendAddressPoolIDRef *v1.Reference `json:"backendAddressPoolIdRef,omitempty" tf:"-"`

	// Selector for a LoadBalancerBackendAddressPool to populate backendAddressPoolId.
	// +kubebuilder:validation:Optional
	BackendAddressPoolIDSelector *v1.Selector `json:"backendAddressPoolIdSelector,omitempty" tf:"-"`

	// Receive bidirectional TCP Reset on TCP flow idle timeout or unexpected connection termination. This element is only used when the protocol is set to TCP.
	// +kubebuilder:validation:Optional
	EnableTCPReset *bool `json:"enableTcpReset,omitempty" tf:"enable_tcp_reset,omitempty"`

	// One or more frontend_ip_configuration blocks as defined below.
	// +kubebuilder:validation:Optional
	FrontendIPConfiguration []LoadBalancerOutboundRuleFrontendIPConfigurationParameters `json:"frontendIpConfiguration,omitempty" tf:"frontend_ip_configuration,omitempty"`

	// The timeout for the TCP idle connection
	// +kubebuilder:validation:Optional
	IdleTimeoutInMinutes *float64 `json:"idleTimeoutInMinutes,omitempty" tf:"idle_timeout_in_minutes,omitempty"`

	// The ID of the Load Balancer in which to create the Outbound Rule. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=LoadBalancer
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-azure/apis/rconfig.ExtractResourceID()
	// +kubebuilder:validation:Optional
	LoadbalancerID *string `json:"loadbalancerId,omitempty" tf:"loadbalancer_id,omitempty"`

	// Reference to a LoadBalancer to populate loadbalancerId.
	// +kubebuilder:validation:Optional
	LoadbalancerIDRef *v1.Reference `json:"loadbalancerIdRef,omitempty" tf:"-"`

	// Selector for a LoadBalancer to populate loadbalancerId.
	// +kubebuilder:validation:Optional
	LoadbalancerIDSelector *v1.Selector `json:"loadbalancerIdSelector,omitempty" tf:"-"`

	// The transport protocol for the external endpoint. Possible values are Udp, Tcp or All.
	// +kubebuilder:validation:Required
	Protocol *string `json:"protocol" tf:"protocol,omitempty"`
}

// LoadBalancerOutboundRuleSpec defines the desired state of LoadBalancerOutboundRule
type LoadBalancerOutboundRuleSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     LoadBalancerOutboundRuleParameters `json:"forProvider"`
}

// LoadBalancerOutboundRuleStatus defines the observed state of LoadBalancerOutboundRule.
type LoadBalancerOutboundRuleStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        LoadBalancerOutboundRuleObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// LoadBalancerOutboundRule is the Schema for the LoadBalancerOutboundRules API. Manages a Load Balancer Outbound Rule.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type LoadBalancerOutboundRule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              LoadBalancerOutboundRuleSpec   `json:"spec"`
	Status            LoadBalancerOutboundRuleStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// LoadBalancerOutboundRuleList contains a list of LoadBalancerOutboundRules
type LoadBalancerOutboundRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LoadBalancerOutboundRule `json:"items"`
}

// Repository type metadata.
var (
	LoadBalancerOutboundRule_Kind             = "LoadBalancerOutboundRule"
	LoadBalancerOutboundRule_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: LoadBalancerOutboundRule_Kind}.String()
	LoadBalancerOutboundRule_KindAPIVersion   = LoadBalancerOutboundRule_Kind + "." + CRDGroupVersion.String()
	LoadBalancerOutboundRule_GroupVersionKind = CRDGroupVersion.WithKind(LoadBalancerOutboundRule_Kind)
)

func init() {
	SchemeBuilder.Register(&LoadBalancerOutboundRule{}, &LoadBalancerOutboundRuleList{})
}
