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

type LoadBalancerRuleObservation struct {

	// The ID of the Load Balancer Rule.
	FrontendIPConfigurationID *string `json:"frontendIpConfigurationId,omitempty" tf:"frontend_ip_configuration_id,omitempty"`

	// The ID of the Load Balancer Rule.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type LoadBalancerRuleParameters struct {

	// A list of reference to a Backend Address Pool over which this Load Balancing Rule operates.
	// +kubebuilder:validation:Optional
	BackendAddressPoolIds []*string `json:"backendAddressPoolIds,omitempty" tf:"backend_address_pool_ids,omitempty"`

	// The port used for internal connections on the endpoint. Possible values range between 0 and 65535, inclusive.
	// +kubebuilder:validation:Required
	BackendPort *float64 `json:"backendPort" tf:"backend_port,omitempty"`

	// Is snat enabled for this Load Balancer Rule? Default false.
	// +kubebuilder:validation:Optional
	DisableOutboundSnat *bool `json:"disableOutboundSnat,omitempty" tf:"disable_outbound_snat,omitempty"`

	// Are the Floating IPs enabled for this Load Balncer Rule? A "floating” IP is reassigned to a secondary server in case the primary server fails. Required to configure a SQL AlwaysOn Availability Group. Defaults to false.
	// +kubebuilder:validation:Optional
	EnableFloatingIP *bool `json:"enableFloatingIp,omitempty" tf:"enable_floating_ip,omitempty"`

	// Is TCP Reset enabled for this Load Balancer Rule? Defaults to false.
	// +kubebuilder:validation:Optional
	EnableTCPReset *bool `json:"enableTcpReset,omitempty" tf:"enable_tcp_reset,omitempty"`

	// The name of the frontend IP configuration to which the rule is associated.
	// +kubebuilder:validation:Required
	FrontendIPConfigurationName *string `json:"frontendIpConfigurationName" tf:"frontend_ip_configuration_name,omitempty"`

	// The port for the external endpoint. Port numbers for each Rule must be unique within the Load Balancer. Possible values range between 0 and 65534, inclusive.
	// +kubebuilder:validation:Required
	FrontendPort *float64 `json:"frontendPort" tf:"frontend_port,omitempty"`

	// Specifies the idle timeout in minutes for TCP connections. Valid values are between 4 and 30 minutes. Defaults to 4 minutes.
	// +kubebuilder:validation:Optional
	IdleTimeoutInMinutes *float64 `json:"idleTimeoutInMinutes,omitempty" tf:"idle_timeout_in_minutes,omitempty"`

	// Specifies the load balancing distribution type to be used by the Load Balancer. Possible values are: Default – The load balancer is configured to use a 5 tuple hash to map traffic to available servers. SourceIP – The load balancer is configured to use a 2 tuple hash to map traffic to available servers. SourceIPProtocol – The load balancer is configured to use a 3 tuple hash to map traffic to available servers. Also known as Session Persistence, where  the options are called None, Client IP and Client IP and Protocol respectively.
	// +kubebuilder:validation:Optional
	LoadDistribution *string `json:"loadDistribution,omitempty" tf:"load_distribution,omitempty"`

	// The ID of the Load Balancer in which to create the Rule.
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

	// A reference to a Probe used by this Load Balancing Rule.
	// +kubebuilder:validation:Optional
	ProbeID *string `json:"probeId,omitempty" tf:"probe_id,omitempty"`

	// The transport protocol for the external endpoint. Possible values are Tcp, Udp or All.
	// +kubebuilder:validation:Required
	Protocol *string `json:"protocol" tf:"protocol,omitempty"`
}

// LoadBalancerRuleSpec defines the desired state of LoadBalancerRule
type LoadBalancerRuleSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     LoadBalancerRuleParameters `json:"forProvider"`
}

// LoadBalancerRuleStatus defines the observed state of LoadBalancerRule.
type LoadBalancerRuleStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        LoadBalancerRuleObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// LoadBalancerRule is the Schema for the LoadBalancerRules API. Manages a Load Balancer Rule.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type LoadBalancerRule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              LoadBalancerRuleSpec   `json:"spec"`
	Status            LoadBalancerRuleStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// LoadBalancerRuleList contains a list of LoadBalancerRules
type LoadBalancerRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LoadBalancerRule `json:"items"`
}

// Repository type metadata.
var (
	LoadBalancerRule_Kind             = "LoadBalancerRule"
	LoadBalancerRule_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: LoadBalancerRule_Kind}.String()
	LoadBalancerRule_KindAPIVersion   = LoadBalancerRule_Kind + "." + CRDGroupVersion.String()
	LoadBalancerRule_GroupVersionKind = CRDGroupVersion.WithKind(LoadBalancerRule_Kind)
)

func init() {
	SchemeBuilder.Register(&LoadBalancerRule{}, &LoadBalancerRuleList{})
}