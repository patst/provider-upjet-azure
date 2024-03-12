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

type BlobInitParameters struct {

	// The access tier of the storage blob. Possible values are Archive, Cool and Hot.
	AccessTier *string `json:"accessTier,omitempty" tf:"access_tier,omitempty"`

	// Controls the cache control header content of the response when blob is requested .
	CacheControl *string `json:"cacheControl,omitempty" tf:"cache_control,omitempty"`

	// The MD5 sum of the blob contents. Cannot be defined if source_uri is defined, or if blob type is Append or Page. Changing this forces a new resource to be created.
	ContentMd5 *string `json:"contentMd5,omitempty" tf:"content_md5,omitempty"`

	// The content type of the storage blob. Cannot be defined if source_uri is defined. Defaults to application/octet-stream.
	ContentType *string `json:"contentType,omitempty" tf:"content_type,omitempty"`

	// A map of custom blob metadata.
	// +mapType=granular
	Metadata map[string]*string `json:"metadata,omitempty" tf:"metadata,omitempty"`

	// The number of workers per CPU core to run for concurrent uploads. Defaults to 8. Changing this forces a new resource to be created.
	Parallelism *float64 `json:"parallelism,omitempty" tf:"parallelism,omitempty"`

	// Used only for page blobs to specify the size in bytes of the blob to be created. Must be a multiple of 512. Defaults to 0. Changing this forces a new resource to be created.
	Size *float64 `json:"size,omitempty" tf:"size,omitempty"`

	// An absolute path to a file on the local system. This field cannot be specified for Append blobs and cannot be specified if source_content or source_uri is specified. Changing this forces a new resource to be created.
	Source *string `json:"source,omitempty" tf:"source,omitempty"`

	// The content for this blob which should be defined inline. This field can only be specified for Block blobs and cannot be specified if source or source_uri is specified. Changing this forces a new resource to be created.
	SourceContent *string `json:"sourceContent,omitempty" tf:"source_content,omitempty"`

	// The URI of an existing blob, or a file in the Azure File service, to use as the source contents for the blob to be created. Changing this forces a new resource to be created. This field cannot be specified for Append blobs and cannot be specified if source or source_content is specified.
	SourceURI *string `json:"sourceUri,omitempty" tf:"source_uri,omitempty"`

	// The type of the storage blob to be created. Possible values are Append, Block or Page. Changing this forces a new resource to be created.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type BlobObservation struct {

	// The access tier of the storage blob. Possible values are Archive, Cool and Hot.
	AccessTier *string `json:"accessTier,omitempty" tf:"access_tier,omitempty"`

	// Controls the cache control header content of the response when blob is requested .
	CacheControl *string `json:"cacheControl,omitempty" tf:"cache_control,omitempty"`

	// The MD5 sum of the blob contents. Cannot be defined if source_uri is defined, or if blob type is Append or Page. Changing this forces a new resource to be created.
	ContentMd5 *string `json:"contentMd5,omitempty" tf:"content_md5,omitempty"`

	// The content type of the storage blob. Cannot be defined if source_uri is defined. Defaults to application/octet-stream.
	ContentType *string `json:"contentType,omitempty" tf:"content_type,omitempty"`

	// The ID of the Storage Blob.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// A map of custom blob metadata.
	// +mapType=granular
	Metadata map[string]*string `json:"metadata,omitempty" tf:"metadata,omitempty"`

	// The number of workers per CPU core to run for concurrent uploads. Defaults to 8. Changing this forces a new resource to be created.
	Parallelism *float64 `json:"parallelism,omitempty" tf:"parallelism,omitempty"`

	// Used only for page blobs to specify the size in bytes of the blob to be created. Must be a multiple of 512. Defaults to 0. Changing this forces a new resource to be created.
	Size *float64 `json:"size,omitempty" tf:"size,omitempty"`

	// An absolute path to a file on the local system. This field cannot be specified for Append blobs and cannot be specified if source_content or source_uri is specified. Changing this forces a new resource to be created.
	Source *string `json:"source,omitempty" tf:"source,omitempty"`

	// The content for this blob which should be defined inline. This field can only be specified for Block blobs and cannot be specified if source or source_uri is specified. Changing this forces a new resource to be created.
	SourceContent *string `json:"sourceContent,omitempty" tf:"source_content,omitempty"`

	// The URI of an existing blob, or a file in the Azure File service, to use as the source contents for the blob to be created. Changing this forces a new resource to be created. This field cannot be specified for Append blobs and cannot be specified if source or source_content is specified.
	SourceURI *string `json:"sourceUri,omitempty" tf:"source_uri,omitempty"`

	// Specifies the storage account in which to create the storage container. Changing this forces a new resource to be created.
	StorageAccountName *string `json:"storageAccountName,omitempty" tf:"storage_account_name,omitempty"`

	// The name of the storage container in which this blob should be created. Changing this forces a new resource to be created.
	StorageContainerName *string `json:"storageContainerName,omitempty" tf:"storage_container_name,omitempty"`

	// The type of the storage blob to be created. Possible values are Append, Block or Page. Changing this forces a new resource to be created.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// The URL of the blob
	URL *string `json:"url,omitempty" tf:"url,omitempty"`
}

type BlobParameters struct {

	// The access tier of the storage blob. Possible values are Archive, Cool and Hot.
	// +kubebuilder:validation:Optional
	AccessTier *string `json:"accessTier,omitempty" tf:"access_tier,omitempty"`

	// Controls the cache control header content of the response when blob is requested .
	// +kubebuilder:validation:Optional
	CacheControl *string `json:"cacheControl,omitempty" tf:"cache_control,omitempty"`

	// The MD5 sum of the blob contents. Cannot be defined if source_uri is defined, or if blob type is Append or Page. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	ContentMd5 *string `json:"contentMd5,omitempty" tf:"content_md5,omitempty"`

	// The content type of the storage blob. Cannot be defined if source_uri is defined. Defaults to application/octet-stream.
	// +kubebuilder:validation:Optional
	ContentType *string `json:"contentType,omitempty" tf:"content_type,omitempty"`

	// A map of custom blob metadata.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Metadata map[string]*string `json:"metadata,omitempty" tf:"metadata,omitempty"`

	// The number of workers per CPU core to run for concurrent uploads. Defaults to 8. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	Parallelism *float64 `json:"parallelism,omitempty" tf:"parallelism,omitempty"`

	// Used only for page blobs to specify the size in bytes of the blob to be created. Must be a multiple of 512. Defaults to 0. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	Size *float64 `json:"size,omitempty" tf:"size,omitempty"`

	// An absolute path to a file on the local system. This field cannot be specified for Append blobs and cannot be specified if source_content or source_uri is specified. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	Source *string `json:"source,omitempty" tf:"source,omitempty"`

	// The content for this blob which should be defined inline. This field can only be specified for Block blobs and cannot be specified if source or source_uri is specified. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	SourceContent *string `json:"sourceContent,omitempty" tf:"source_content,omitempty"`

	// The URI of an existing blob, or a file in the Azure File service, to use as the source contents for the blob to be created. Changing this forces a new resource to be created. This field cannot be specified for Append blobs and cannot be specified if source or source_content is specified.
	// +kubebuilder:validation:Optional
	SourceURI *string `json:"sourceUri,omitempty" tf:"source_uri,omitempty"`

	// Specifies the storage account in which to create the storage container. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=Account
	// +kubebuilder:validation:Optional
	StorageAccountName *string `json:"storageAccountName,omitempty" tf:"storage_account_name,omitempty"`

	// Reference to a Account to populate storageAccountName.
	// +kubebuilder:validation:Optional
	StorageAccountNameRef *v1.Reference `json:"storageAccountNameRef,omitempty" tf:"-"`

	// Selector for a Account to populate storageAccountName.
	// +kubebuilder:validation:Optional
	StorageAccountNameSelector *v1.Selector `json:"storageAccountNameSelector,omitempty" tf:"-"`

	// The name of the storage container in which this blob should be created. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=Container
	// +kubebuilder:validation:Optional
	StorageContainerName *string `json:"storageContainerName,omitempty" tf:"storage_container_name,omitempty"`

	// Reference to a Container to populate storageContainerName.
	// +kubebuilder:validation:Optional
	StorageContainerNameRef *v1.Reference `json:"storageContainerNameRef,omitempty" tf:"-"`

	// Selector for a Container to populate storageContainerName.
	// +kubebuilder:validation:Optional
	StorageContainerNameSelector *v1.Selector `json:"storageContainerNameSelector,omitempty" tf:"-"`

	// The type of the storage blob to be created. Possible values are Append, Block or Page. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

// BlobSpec defines the desired state of Blob
type BlobSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     BlobParameters `json:"forProvider"`
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
	InitProvider BlobInitParameters `json:"initProvider,omitempty"`
}

// BlobStatus defines the observed state of Blob.
type BlobStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        BlobObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// Blob is the Schema for the Blobs API. Manages a Blob within a Storage Container.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type Blob struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.type) || (has(self.initProvider) && has(self.initProvider.type))",message="spec.forProvider.type is a required parameter"
	Spec   BlobSpec   `json:"spec"`
	Status BlobStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// BlobList contains a list of Blobs
type BlobList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Blob `json:"items"`
}

// Repository type metadata.
var (
	Blob_Kind             = "Blob"
	Blob_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Blob_Kind}.String()
	Blob_KindAPIVersion   = Blob_Kind + "." + CRDGroupVersion.String()
	Blob_GroupVersionKind = CRDGroupVersion.WithKind(Blob_Kind)
)

func init() {
	SchemeBuilder.Register(&Blob{}, &BlobList{})
}
