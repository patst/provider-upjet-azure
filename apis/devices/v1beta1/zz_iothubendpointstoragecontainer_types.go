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

type IOTHubEndpointStorageContainerObservation struct {

	// The ID of the IoTHub Storage Container Endpoint.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type IOTHubEndpointStorageContainerParameters struct {

	// Type used to authenticate against the storage endpoint. Possible values are keyBased and identityBased. Defaults to keyBased.
	// +kubebuilder:validation:Optional
	AuthenticationType *string `json:"authenticationType,omitempty" tf:"authentication_type,omitempty"`

	// Time interval at which blobs are written to storage. Value should be between 60 and 720 seconds. Default value is 300 seconds.
	// +kubebuilder:validation:Optional
	BatchFrequencyInSeconds *float64 `json:"batchFrequencyInSeconds,omitempty" tf:"batch_frequency_in_seconds,omitempty"`

	// The connection string for the endpoint. This attribute can only be specified and is mandatory when authentication_type is keyBased.
	// +kubebuilder:validation:Optional
	ConnectionStringSecretRef *v1.SecretKeySelector `json:"connectionStringSecretRef,omitempty" tf:"-"`

	// The name of storage container in the storage account.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/storage/v1beta1.Container
	// +kubebuilder:validation:Optional
	ContainerName *string `json:"containerName,omitempty" tf:"container_name,omitempty"`

	// Reference to a Container in storage to populate containerName.
	// +kubebuilder:validation:Optional
	ContainerNameRef *v1.Reference `json:"containerNameRef,omitempty" tf:"-"`

	// Selector for a Container in storage to populate containerName.
	// +kubebuilder:validation:Optional
	ContainerNameSelector *v1.Selector `json:"containerNameSelector,omitempty" tf:"-"`

	// Encoding that is used to serialize messages to blobs. Supported values are Avro, AvroDeflate and JSON. Default value is Avro. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	Encoding *string `json:"encoding,omitempty" tf:"encoding,omitempty"`

	// URI of the Storage Container endpoint. This corresponds to the primary_blob_endpoint of the parent storage account. This attribute can only be specified and is mandatory when authentication_type is identityBased.
	// +kubebuilder:validation:Optional
	EndpointURI *string `json:"endpointUri,omitempty" tf:"endpoint_uri,omitempty"`

	// File name format for the blob. Default format is {iothub}/{partition}/{YYYY}/{MM}/{DD}/{HH}/{mm}. All parameters are mandatory but can be reordered.
	// +kubebuilder:validation:Optional
	FileNameFormat *string `json:"fileNameFormat,omitempty" tf:"file_name_format,omitempty"`

	// The IoTHub ID for the endpoint.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/devices/v1beta1.IOTHub
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	IOTHubID *string `json:"iothubId,omitempty" tf:"iothub_id,omitempty"`

	// Reference to a IOTHub in devices to populate iothubId.
	// +kubebuilder:validation:Optional
	IOTHubIDRef *v1.Reference `json:"iothubIdRef,omitempty" tf:"-"`

	// Selector for a IOTHub in devices to populate iothubId.
	// +kubebuilder:validation:Optional
	IOTHubIDSelector *v1.Selector `json:"iothubIdSelector,omitempty" tf:"-"`

	// ID of the User Managed Identity used to authenticate against the storage endpoint.
	// +kubebuilder:validation:Optional
	IdentityID *string `json:"identityId,omitempty" tf:"identity_id,omitempty"`

	// Maximum number of bytes for each blob written to storage. Value should be between 10485760(10MB) and 524288000(500MB). Default value is 314572800(300MB).
	// +kubebuilder:validation:Optional
	MaxChunkSizeInBytes *float64 `json:"maxChunkSizeInBytes,omitempty" tf:"max_chunk_size_in_bytes,omitempty"`

	// The name of the resource group under which the Storage Container has been created. Changing this forces a new resource to be created.
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

// IOTHubEndpointStorageContainerSpec defines the desired state of IOTHubEndpointStorageContainer
type IOTHubEndpointStorageContainerSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     IOTHubEndpointStorageContainerParameters `json:"forProvider"`
}

// IOTHubEndpointStorageContainerStatus defines the observed state of IOTHubEndpointStorageContainer.
type IOTHubEndpointStorageContainerStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        IOTHubEndpointStorageContainerObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// IOTHubEndpointStorageContainer is the Schema for the IOTHubEndpointStorageContainers API. Manages an IotHub Storage Container Endpoint
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type IOTHubEndpointStorageContainer struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              IOTHubEndpointStorageContainerSpec   `json:"spec"`
	Status            IOTHubEndpointStorageContainerStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// IOTHubEndpointStorageContainerList contains a list of IOTHubEndpointStorageContainers
type IOTHubEndpointStorageContainerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []IOTHubEndpointStorageContainer `json:"items"`
}

// Repository type metadata.
var (
	IOTHubEndpointStorageContainer_Kind             = "IOTHubEndpointStorageContainer"
	IOTHubEndpointStorageContainer_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: IOTHubEndpointStorageContainer_Kind}.String()
	IOTHubEndpointStorageContainer_KindAPIVersion   = IOTHubEndpointStorageContainer_Kind + "." + CRDGroupVersion.String()
	IOTHubEndpointStorageContainer_GroupVersionKind = CRDGroupVersion.WithKind(IOTHubEndpointStorageContainer_Kind)
)

func init() {
	SchemeBuilder.Register(&IOTHubEndpointStorageContainer{}, &IOTHubEndpointStorageContainerList{})
}
