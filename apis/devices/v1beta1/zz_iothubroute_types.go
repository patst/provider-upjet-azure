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

type IOTHubRouteObservation struct {

	// The ID of the IoTHub Route.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type IOTHubRouteParameters struct {

	// The condition that is evaluated to apply the routing rule. If no condition is provided, it evaluates to true by default. For grammar, see: https://docs.microsoft.com/azure/iot-hub/iot-hub-devguide-query-language.
	// +kubebuilder:validation:Optional
	Condition *string `json:"condition,omitempty" tf:"condition,omitempty"`

	// Specifies whether a route is enabled.
	// +kubebuilder:validation:Required
	Enabled *bool `json:"enabled" tf:"enabled,omitempty"`

	// The list of endpoints to which messages that satisfy the condition are routed. Currently only one endpoint is allowed.
	// +kubebuilder:validation:Required
	EndpointNames []*string `json:"endpointNames" tf:"endpoint_names,omitempty"`

	// The name of the IoTHub to which this Route belongs. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/devices/v1beta1.IOTHub
	// +kubebuilder:validation:Optional
	IOTHubName *string `json:"iothubName,omitempty" tf:"iothub_name,omitempty"`

	// Reference to a IOTHub in devices to populate iothubName.
	// +kubebuilder:validation:Optional
	IOTHubNameRef *v1.Reference `json:"iothubNameRef,omitempty" tf:"-"`

	// Selector for a IOTHub in devices to populate iothubName.
	// +kubebuilder:validation:Optional
	IOTHubNameSelector *v1.Selector `json:"iothubNameSelector,omitempty" tf:"-"`

	// The name of the resource group under which the IotHub Route resource has to be created. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// The source that the routing rule is to be applied to. Possible values include: DeviceConnectionStateEvents, DeviceJobLifecycleEvents, DeviceLifecycleEvents, DeviceMessages, Invalid, TwinChangeEvents. Defaults to DeviceMessages.
	// +kubebuilder:validation:Required
	Source *string `json:"source" tf:"source,omitempty"`
}

// IOTHubRouteSpec defines the desired state of IOTHubRoute
type IOTHubRouteSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     IOTHubRouteParameters `json:"forProvider"`
}

// IOTHubRouteStatus defines the observed state of IOTHubRoute.
type IOTHubRouteStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        IOTHubRouteObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// IOTHubRoute is the Schema for the IOTHubRoutes API. Manages an IotHub Route
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type IOTHubRoute struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              IOTHubRouteSpec   `json:"spec"`
	Status            IOTHubRouteStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// IOTHubRouteList contains a list of IOTHubRoutes
type IOTHubRouteList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []IOTHubRoute `json:"items"`
}

// Repository type metadata.
var (
	IOTHubRoute_Kind             = "IOTHubRoute"
	IOTHubRoute_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: IOTHubRoute_Kind}.String()
	IOTHubRoute_KindAPIVersion   = IOTHubRoute_Kind + "." + CRDGroupVersion.String()
	IOTHubRoute_GroupVersionKind = CRDGroupVersion.WithKind(IOTHubRoute_Kind)
)

func init() {
	SchemeBuilder.Register(&IOTHubRoute{}, &IOTHubRouteList{})
}
