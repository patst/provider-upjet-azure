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

type FirewallIPConfigurationInitParameters struct {

	// Specifies the name of the IP Configuration.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The ID of the Public IP Address associated with the firewall.
	// +crossplane:generate:reference:type=PublicIP
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-azure/apis/rconfig.ExtractResourceID()
	PublicIPAddressID *string `json:"publicIpAddressId,omitempty" tf:"public_ip_address_id,omitempty"`

	// Reference to a PublicIP to populate publicIpAddressId.
	// +kubebuilder:validation:Optional
	PublicIPAddressIDRef *v1.Reference `json:"publicIpAddressIdRef,omitempty" tf:"-"`

	// Selector for a PublicIP to populate publicIpAddressId.
	// +kubebuilder:validation:Optional
	PublicIPAddressIDSelector *v1.Selector `json:"publicIpAddressIdSelector,omitempty" tf:"-"`

	// Reference to the subnet associated with the IP Configuration. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=Subnet
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-azure/apis/rconfig.ExtractResourceID()
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`

	// Reference to a Subnet to populate subnetId.
	// +kubebuilder:validation:Optional
	SubnetIDRef *v1.Reference `json:"subnetIdRef,omitempty" tf:"-"`

	// Selector for a Subnet to populate subnetId.
	// +kubebuilder:validation:Optional
	SubnetIDSelector *v1.Selector `json:"subnetIdSelector,omitempty" tf:"-"`
}

type FirewallIPConfigurationObservation struct {

	// Specifies the name of the IP Configuration.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The Private IP address of the Azure Firewall.
	PrivateIPAddress *string `json:"privateIpAddress,omitempty" tf:"private_ip_address,omitempty"`

	// The ID of the Public IP Address associated with the firewall.
	PublicIPAddressID *string `json:"publicIpAddressId,omitempty" tf:"public_ip_address_id,omitempty"`

	// Reference to the subnet associated with the IP Configuration. Changing this forces a new resource to be created.
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`
}

type FirewallIPConfigurationParameters struct {

	// Specifies the name of the IP Configuration.
	// +kubebuilder:validation:Optional
	Name *string `json:"name" tf:"name,omitempty"`

	// The ID of the Public IP Address associated with the firewall.
	// +crossplane:generate:reference:type=PublicIP
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-azure/apis/rconfig.ExtractResourceID()
	// +kubebuilder:validation:Optional
	PublicIPAddressID *string `json:"publicIpAddressId,omitempty" tf:"public_ip_address_id,omitempty"`

	// Reference to a PublicIP to populate publicIpAddressId.
	// +kubebuilder:validation:Optional
	PublicIPAddressIDRef *v1.Reference `json:"publicIpAddressIdRef,omitempty" tf:"-"`

	// Selector for a PublicIP to populate publicIpAddressId.
	// +kubebuilder:validation:Optional
	PublicIPAddressIDSelector *v1.Selector `json:"publicIpAddressIdSelector,omitempty" tf:"-"`

	// Reference to the subnet associated with the IP Configuration. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=Subnet
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-azure/apis/rconfig.ExtractResourceID()
	// +kubebuilder:validation:Optional
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`

	// Reference to a Subnet to populate subnetId.
	// +kubebuilder:validation:Optional
	SubnetIDRef *v1.Reference `json:"subnetIdRef,omitempty" tf:"-"`

	// Selector for a Subnet to populate subnetId.
	// +kubebuilder:validation:Optional
	SubnetIDSelector *v1.Selector `json:"subnetIdSelector,omitempty" tf:"-"`
}

type FirewallInitParameters struct {

	// Whether DNS proxy is enabled. It will forward DNS requests to the DNS servers when set to true. It will be set to true if dns_servers provided with a not empty list.
	DNSProxyEnabled *bool `json:"dnsProxyEnabled,omitempty" tf:"dns_proxy_enabled,omitempty"`

	// A list of DNS servers that the Azure Firewall will direct DNS traffic to the for name resolution.
	DNSServers []*string `json:"dnsServers,omitempty" tf:"dns_servers,omitempty"`

	// The ID of the Firewall Policy applied to this Firewall.
	FirewallPolicyID *string `json:"firewallPolicyId,omitempty" tf:"firewall_policy_id,omitempty"`

	// An ip_configuration block as documented below.
	IPConfiguration []FirewallIPConfigurationInitParameters `json:"ipConfiguration,omitempty" tf:"ip_configuration,omitempty"`

	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// A management_ip_configuration block as documented below, which allows force-tunnelling of traffic to be performed by the firewall. Adding or removing this block or changing the subnet_id in an existing block forces a new resource to be created. Changing this forces a new resource to be created.
	ManagementIPConfiguration []ManagementIPConfigurationInitParameters `json:"managementIpConfiguration,omitempty" tf:"management_ip_configuration,omitempty"`

	// A list of SNAT private CIDR IP ranges, or the special string IANAPrivateRanges, which indicates Azure Firewall does not SNAT when the destination IP address is a private range per IANA RFC 1918.
	// +listType=set
	PrivateIPRanges []*string `json:"privateIpRanges,omitempty" tf:"private_ip_ranges,omitempty"`

	// SKU name of the Firewall. Possible values are AZFW_Hub and AZFW_VNet. Changing this forces a new resource to be created.
	SkuName *string `json:"skuName,omitempty" tf:"sku_name,omitempty"`

	// SKU tier of the Firewall. Possible values are Premium, Standard and Basic.
	SkuTier *string `json:"skuTier,omitempty" tf:"sku_tier,omitempty"`

	// A mapping of tags to assign to the resource.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The operation mode for threat intelligence-based filtering. Possible values are: Off, Alert and Deny. Defaults to Alert.
	ThreatIntelMode *string `json:"threatIntelMode,omitempty" tf:"threat_intel_mode,omitempty"`

	// A virtual_hub block as documented below.
	VirtualHub []VirtualHubInitParameters `json:"virtualHub,omitempty" tf:"virtual_hub,omitempty"`

	// Specifies a list of Availability Zones in which this Azure Firewall should be located. Changing this forces a new Azure Firewall to be created.
	// +listType=set
	Zones []*string `json:"zones,omitempty" tf:"zones,omitempty"`
}

type FirewallObservation struct {

	// Whether DNS proxy is enabled. It will forward DNS requests to the DNS servers when set to true. It will be set to true if dns_servers provided with a not empty list.
	DNSProxyEnabled *bool `json:"dnsProxyEnabled,omitempty" tf:"dns_proxy_enabled,omitempty"`

	// A list of DNS servers that the Azure Firewall will direct DNS traffic to the for name resolution.
	DNSServers []*string `json:"dnsServers,omitempty" tf:"dns_servers,omitempty"`

	// The ID of the Firewall Policy applied to this Firewall.
	FirewallPolicyID *string `json:"firewallPolicyId,omitempty" tf:"firewall_policy_id,omitempty"`

	// The ID of the Azure Firewall.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// An ip_configuration block as documented below.
	IPConfiguration []FirewallIPConfigurationObservation `json:"ipConfiguration,omitempty" tf:"ip_configuration,omitempty"`

	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// A management_ip_configuration block as documented below, which allows force-tunnelling of traffic to be performed by the firewall. Adding or removing this block or changing the subnet_id in an existing block forces a new resource to be created. Changing this forces a new resource to be created.
	ManagementIPConfiguration []ManagementIPConfigurationObservation `json:"managementIpConfiguration,omitempty" tf:"management_ip_configuration,omitempty"`

	// A list of SNAT private CIDR IP ranges, or the special string IANAPrivateRanges, which indicates Azure Firewall does not SNAT when the destination IP address is a private range per IANA RFC 1918.
	// +listType=set
	PrivateIPRanges []*string `json:"privateIpRanges,omitempty" tf:"private_ip_ranges,omitempty"`

	// The name of the resource group in which to create the resource. Changing this forces a new resource to be created.
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// SKU name of the Firewall. Possible values are AZFW_Hub and AZFW_VNet. Changing this forces a new resource to be created.
	SkuName *string `json:"skuName,omitempty" tf:"sku_name,omitempty"`

	// SKU tier of the Firewall. Possible values are Premium, Standard and Basic.
	SkuTier *string `json:"skuTier,omitempty" tf:"sku_tier,omitempty"`

	// A mapping of tags to assign to the resource.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The operation mode for threat intelligence-based filtering. Possible values are: Off, Alert and Deny. Defaults to Alert.
	ThreatIntelMode *string `json:"threatIntelMode,omitempty" tf:"threat_intel_mode,omitempty"`

	// A virtual_hub block as documented below.
	VirtualHub []VirtualHubObservation `json:"virtualHub,omitempty" tf:"virtual_hub,omitempty"`

	// Specifies a list of Availability Zones in which this Azure Firewall should be located. Changing this forces a new Azure Firewall to be created.
	// +listType=set
	Zones []*string `json:"zones,omitempty" tf:"zones,omitempty"`
}

type FirewallParameters struct {

	// Whether DNS proxy is enabled. It will forward DNS requests to the DNS servers when set to true. It will be set to true if dns_servers provided with a not empty list.
	// +kubebuilder:validation:Optional
	DNSProxyEnabled *bool `json:"dnsProxyEnabled,omitempty" tf:"dns_proxy_enabled,omitempty"`

	// A list of DNS servers that the Azure Firewall will direct DNS traffic to the for name resolution.
	// +kubebuilder:validation:Optional
	DNSServers []*string `json:"dnsServers,omitempty" tf:"dns_servers,omitempty"`

	// The ID of the Firewall Policy applied to this Firewall.
	// +kubebuilder:validation:Optional
	FirewallPolicyID *string `json:"firewallPolicyId,omitempty" tf:"firewall_policy_id,omitempty"`

	// An ip_configuration block as documented below.
	// +kubebuilder:validation:Optional
	IPConfiguration []FirewallIPConfigurationParameters `json:"ipConfiguration,omitempty" tf:"ip_configuration,omitempty"`

	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// A management_ip_configuration block as documented below, which allows force-tunnelling of traffic to be performed by the firewall. Adding or removing this block or changing the subnet_id in an existing block forces a new resource to be created. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	ManagementIPConfiguration []ManagementIPConfigurationParameters `json:"managementIpConfiguration,omitempty" tf:"management_ip_configuration,omitempty"`

	// A list of SNAT private CIDR IP ranges, or the special string IANAPrivateRanges, which indicates Azure Firewall does not SNAT when the destination IP address is a private range per IANA RFC 1918.
	// +kubebuilder:validation:Optional
	// +listType=set
	PrivateIPRanges []*string `json:"privateIpRanges,omitempty" tf:"private_ip_ranges,omitempty"`

	// The name of the resource group in which to create the resource. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// SKU name of the Firewall. Possible values are AZFW_Hub and AZFW_VNet. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	SkuName *string `json:"skuName,omitempty" tf:"sku_name,omitempty"`

	// SKU tier of the Firewall. Possible values are Premium, Standard and Basic.
	// +kubebuilder:validation:Optional
	SkuTier *string `json:"skuTier,omitempty" tf:"sku_tier,omitempty"`

	// A mapping of tags to assign to the resource.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The operation mode for threat intelligence-based filtering. Possible values are: Off, Alert and Deny. Defaults to Alert.
	// +kubebuilder:validation:Optional
	ThreatIntelMode *string `json:"threatIntelMode,omitempty" tf:"threat_intel_mode,omitempty"`

	// A virtual_hub block as documented below.
	// +kubebuilder:validation:Optional
	VirtualHub []VirtualHubParameters `json:"virtualHub,omitempty" tf:"virtual_hub,omitempty"`

	// Specifies a list of Availability Zones in which this Azure Firewall should be located. Changing this forces a new Azure Firewall to be created.
	// +kubebuilder:validation:Optional
	// +listType=set
	Zones []*string `json:"zones,omitempty" tf:"zones,omitempty"`
}

type ManagementIPConfigurationInitParameters struct {

	// Specifies the name of the IP Configuration.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The ID of the Public IP Address associated with the firewall.
	PublicIPAddressID *string `json:"publicIpAddressId,omitempty" tf:"public_ip_address_id,omitempty"`

	// Reference to the subnet associated with the IP Configuration. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=Subnet
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-azure/apis/rconfig.ExtractResourceID()
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`

	// Reference to a Subnet to populate subnetId.
	// +kubebuilder:validation:Optional
	SubnetIDRef *v1.Reference `json:"subnetIdRef,omitempty" tf:"-"`

	// Selector for a Subnet to populate subnetId.
	// +kubebuilder:validation:Optional
	SubnetIDSelector *v1.Selector `json:"subnetIdSelector,omitempty" tf:"-"`
}

type ManagementIPConfigurationObservation struct {

	// Specifies the name of the IP Configuration.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The private IP address associated with the Firewall.
	PrivateIPAddress *string `json:"privateIpAddress,omitempty" tf:"private_ip_address,omitempty"`

	// The ID of the Public IP Address associated with the firewall.
	PublicIPAddressID *string `json:"publicIpAddressId,omitempty" tf:"public_ip_address_id,omitempty"`

	// Reference to the subnet associated with the IP Configuration. Changing this forces a new resource to be created.
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`
}

type ManagementIPConfigurationParameters struct {

	// Specifies the name of the IP Configuration.
	// +kubebuilder:validation:Optional
	Name *string `json:"name" tf:"name,omitempty"`

	// The ID of the Public IP Address associated with the firewall.
	// +kubebuilder:validation:Optional
	PublicIPAddressID *string `json:"publicIpAddressId" tf:"public_ip_address_id,omitempty"`

	// Reference to the subnet associated with the IP Configuration. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=Subnet
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-azure/apis/rconfig.ExtractResourceID()
	// +kubebuilder:validation:Optional
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`

	// Reference to a Subnet to populate subnetId.
	// +kubebuilder:validation:Optional
	SubnetIDRef *v1.Reference `json:"subnetIdRef,omitempty" tf:"-"`

	// Selector for a Subnet to populate subnetId.
	// +kubebuilder:validation:Optional
	SubnetIDSelector *v1.Selector `json:"subnetIdSelector,omitempty" tf:"-"`
}

type VirtualHubInitParameters struct {

	// Specifies the number of public IPs to assign to the Firewall. Defaults to 1.
	PublicIPCount *float64 `json:"publicIpCount,omitempty" tf:"public_ip_count,omitempty"`

	// Specifies the ID of the Virtual Hub where the Firewall resides in.
	VirtualHubID *string `json:"virtualHubId,omitempty" tf:"virtual_hub_id,omitempty"`
}

type VirtualHubObservation struct {

	// The private IP address associated with the Firewall.
	PrivateIPAddress *string `json:"privateIpAddress,omitempty" tf:"private_ip_address,omitempty"`

	// The list of public IP addresses associated with the Firewall.
	PublicIPAddresses []*string `json:"publicIpAddresses,omitempty" tf:"public_ip_addresses,omitempty"`

	// Specifies the number of public IPs to assign to the Firewall. Defaults to 1.
	PublicIPCount *float64 `json:"publicIpCount,omitempty" tf:"public_ip_count,omitempty"`

	// Specifies the ID of the Virtual Hub where the Firewall resides in.
	VirtualHubID *string `json:"virtualHubId,omitempty" tf:"virtual_hub_id,omitempty"`
}

type VirtualHubParameters struct {

	// Specifies the number of public IPs to assign to the Firewall. Defaults to 1.
	// +kubebuilder:validation:Optional
	PublicIPCount *float64 `json:"publicIpCount,omitempty" tf:"public_ip_count,omitempty"`

	// Specifies the ID of the Virtual Hub where the Firewall resides in.
	// +kubebuilder:validation:Optional
	VirtualHubID *string `json:"virtualHubId" tf:"virtual_hub_id,omitempty"`
}

// FirewallSpec defines the desired state of Firewall
type FirewallSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     FirewallParameters `json:"forProvider"`
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
	InitProvider FirewallInitParameters `json:"initProvider,omitempty"`
}

// FirewallStatus defines the observed state of Firewall.
type FirewallStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        FirewallObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// Firewall is the Schema for the Firewalls API. Manages an Azure Firewall.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type Firewall struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.location) || (has(self.initProvider) && has(self.initProvider.location))",message="spec.forProvider.location is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.skuName) || (has(self.initProvider) && has(self.initProvider.skuName))",message="spec.forProvider.skuName is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.skuTier) || (has(self.initProvider) && has(self.initProvider.skuTier))",message="spec.forProvider.skuTier is a required parameter"
	Spec   FirewallSpec   `json:"spec"`
	Status FirewallStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// FirewallList contains a list of Firewalls
type FirewallList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Firewall `json:"items"`
}

// Repository type metadata.
var (
	Firewall_Kind             = "Firewall"
	Firewall_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Firewall_Kind}.String()
	Firewall_KindAPIVersion   = Firewall_Kind + "." + CRDGroupVersion.String()
	Firewall_GroupVersionKind = CRDGroupVersion.WithKind(Firewall_Kind)
)

func init() {
	SchemeBuilder.Register(&Firewall{}, &FirewallList{})
}
