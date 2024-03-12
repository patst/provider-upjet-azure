//go:build !ignore_autogenerated

// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by controller-gen. DO NOT EDIT.

package v1beta1

import (
	"github.com/crossplane/crossplane-runtime/apis/common/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Configuration) DeepCopyInto(out *Configuration) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Configuration.
func (in *Configuration) DeepCopy() *Configuration {
	if in == nil {
		return nil
	}
	out := new(Configuration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Configuration) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConfigurationInitParameters) DeepCopyInto(out *ConfigurationInitParameters) {
	*out = *in
	if in.Encryption != nil {
		in, out := &in.Encryption, &out.Encryption
		*out = make([]EncryptionInitParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Identity != nil {
		in, out := &in.Identity, &out.Identity
		*out = make([]IdentityInitParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.LocalAuthEnabled != nil {
		in, out := &in.LocalAuthEnabled, &out.LocalAuthEnabled
		*out = new(bool)
		**out = **in
	}
	if in.Location != nil {
		in, out := &in.Location, &out.Location
		*out = new(string)
		**out = **in
	}
	if in.PublicNetworkAccess != nil {
		in, out := &in.PublicNetworkAccess, &out.PublicNetworkAccess
		*out = new(string)
		**out = **in
	}
	if in.PurgeProtectionEnabled != nil {
		in, out := &in.PurgeProtectionEnabled, &out.PurgeProtectionEnabled
		*out = new(bool)
		**out = **in
	}
	if in.Replica != nil {
		in, out := &in.Replica, &out.Replica
		*out = make([]ReplicaInitParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Sku != nil {
		in, out := &in.Sku, &out.Sku
		*out = new(string)
		**out = **in
	}
	if in.SoftDeleteRetentionDays != nil {
		in, out := &in.SoftDeleteRetentionDays, &out.SoftDeleteRetentionDays
		*out = new(float64)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConfigurationInitParameters.
func (in *ConfigurationInitParameters) DeepCopy() *ConfigurationInitParameters {
	if in == nil {
		return nil
	}
	out := new(ConfigurationInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConfigurationList) DeepCopyInto(out *ConfigurationList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Configuration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConfigurationList.
func (in *ConfigurationList) DeepCopy() *ConfigurationList {
	if in == nil {
		return nil
	}
	out := new(ConfigurationList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ConfigurationList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConfigurationObservation) DeepCopyInto(out *ConfigurationObservation) {
	*out = *in
	if in.Encryption != nil {
		in, out := &in.Encryption, &out.Encryption
		*out = make([]EncryptionObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Endpoint != nil {
		in, out := &in.Endpoint, &out.Endpoint
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.Identity != nil {
		in, out := &in.Identity, &out.Identity
		*out = make([]IdentityObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.LocalAuthEnabled != nil {
		in, out := &in.LocalAuthEnabled, &out.LocalAuthEnabled
		*out = new(bool)
		**out = **in
	}
	if in.Location != nil {
		in, out := &in.Location, &out.Location
		*out = new(string)
		**out = **in
	}
	if in.PrimaryReadKey != nil {
		in, out := &in.PrimaryReadKey, &out.PrimaryReadKey
		*out = make([]PrimaryReadKeyObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.PrimaryWriteKey != nil {
		in, out := &in.PrimaryWriteKey, &out.PrimaryWriteKey
		*out = make([]PrimaryWriteKeyObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.PublicNetworkAccess != nil {
		in, out := &in.PublicNetworkAccess, &out.PublicNetworkAccess
		*out = new(string)
		**out = **in
	}
	if in.PurgeProtectionEnabled != nil {
		in, out := &in.PurgeProtectionEnabled, &out.PurgeProtectionEnabled
		*out = new(bool)
		**out = **in
	}
	if in.Replica != nil {
		in, out := &in.Replica, &out.Replica
		*out = make([]ReplicaObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ResourceGroupName != nil {
		in, out := &in.ResourceGroupName, &out.ResourceGroupName
		*out = new(string)
		**out = **in
	}
	if in.SecondaryReadKey != nil {
		in, out := &in.SecondaryReadKey, &out.SecondaryReadKey
		*out = make([]SecondaryReadKeyObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.SecondaryWriteKey != nil {
		in, out := &in.SecondaryWriteKey, &out.SecondaryWriteKey
		*out = make([]SecondaryWriteKeyObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Sku != nil {
		in, out := &in.Sku, &out.Sku
		*out = new(string)
		**out = **in
	}
	if in.SoftDeleteRetentionDays != nil {
		in, out := &in.SoftDeleteRetentionDays, &out.SoftDeleteRetentionDays
		*out = new(float64)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConfigurationObservation.
func (in *ConfigurationObservation) DeepCopy() *ConfigurationObservation {
	if in == nil {
		return nil
	}
	out := new(ConfigurationObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConfigurationParameters) DeepCopyInto(out *ConfigurationParameters) {
	*out = *in
	if in.Encryption != nil {
		in, out := &in.Encryption, &out.Encryption
		*out = make([]EncryptionParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Identity != nil {
		in, out := &in.Identity, &out.Identity
		*out = make([]IdentityParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.LocalAuthEnabled != nil {
		in, out := &in.LocalAuthEnabled, &out.LocalAuthEnabled
		*out = new(bool)
		**out = **in
	}
	if in.Location != nil {
		in, out := &in.Location, &out.Location
		*out = new(string)
		**out = **in
	}
	if in.PublicNetworkAccess != nil {
		in, out := &in.PublicNetworkAccess, &out.PublicNetworkAccess
		*out = new(string)
		**out = **in
	}
	if in.PurgeProtectionEnabled != nil {
		in, out := &in.PurgeProtectionEnabled, &out.PurgeProtectionEnabled
		*out = new(bool)
		**out = **in
	}
	if in.Replica != nil {
		in, out := &in.Replica, &out.Replica
		*out = make([]ReplicaParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ResourceGroupName != nil {
		in, out := &in.ResourceGroupName, &out.ResourceGroupName
		*out = new(string)
		**out = **in
	}
	if in.ResourceGroupNameRef != nil {
		in, out := &in.ResourceGroupNameRef, &out.ResourceGroupNameRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.ResourceGroupNameSelector != nil {
		in, out := &in.ResourceGroupNameSelector, &out.ResourceGroupNameSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.Sku != nil {
		in, out := &in.Sku, &out.Sku
		*out = new(string)
		**out = **in
	}
	if in.SoftDeleteRetentionDays != nil {
		in, out := &in.SoftDeleteRetentionDays, &out.SoftDeleteRetentionDays
		*out = new(float64)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConfigurationParameters.
func (in *ConfigurationParameters) DeepCopy() *ConfigurationParameters {
	if in == nil {
		return nil
	}
	out := new(ConfigurationParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConfigurationSpec) DeepCopyInto(out *ConfigurationSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
	in.InitProvider.DeepCopyInto(&out.InitProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConfigurationSpec.
func (in *ConfigurationSpec) DeepCopy() *ConfigurationSpec {
	if in == nil {
		return nil
	}
	out := new(ConfigurationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConfigurationStatus) DeepCopyInto(out *ConfigurationStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConfigurationStatus.
func (in *ConfigurationStatus) DeepCopy() *ConfigurationStatus {
	if in == nil {
		return nil
	}
	out := new(ConfigurationStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EncryptionInitParameters) DeepCopyInto(out *EncryptionInitParameters) {
	*out = *in
	if in.IdentityClientID != nil {
		in, out := &in.IdentityClientID, &out.IdentityClientID
		*out = new(string)
		**out = **in
	}
	if in.IdentityClientIDRef != nil {
		in, out := &in.IdentityClientIDRef, &out.IdentityClientIDRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.IdentityClientIDSelector != nil {
		in, out := &in.IdentityClientIDSelector, &out.IdentityClientIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.KeyVaultKeyIdentifier != nil {
		in, out := &in.KeyVaultKeyIdentifier, &out.KeyVaultKeyIdentifier
		*out = new(string)
		**out = **in
	}
	if in.KeyVaultKeyIdentifierRef != nil {
		in, out := &in.KeyVaultKeyIdentifierRef, &out.KeyVaultKeyIdentifierRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.KeyVaultKeyIdentifierSelector != nil {
		in, out := &in.KeyVaultKeyIdentifierSelector, &out.KeyVaultKeyIdentifierSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EncryptionInitParameters.
func (in *EncryptionInitParameters) DeepCopy() *EncryptionInitParameters {
	if in == nil {
		return nil
	}
	out := new(EncryptionInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EncryptionObservation) DeepCopyInto(out *EncryptionObservation) {
	*out = *in
	if in.IdentityClientID != nil {
		in, out := &in.IdentityClientID, &out.IdentityClientID
		*out = new(string)
		**out = **in
	}
	if in.KeyVaultKeyIdentifier != nil {
		in, out := &in.KeyVaultKeyIdentifier, &out.KeyVaultKeyIdentifier
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EncryptionObservation.
func (in *EncryptionObservation) DeepCopy() *EncryptionObservation {
	if in == nil {
		return nil
	}
	out := new(EncryptionObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EncryptionParameters) DeepCopyInto(out *EncryptionParameters) {
	*out = *in
	if in.IdentityClientID != nil {
		in, out := &in.IdentityClientID, &out.IdentityClientID
		*out = new(string)
		**out = **in
	}
	if in.IdentityClientIDRef != nil {
		in, out := &in.IdentityClientIDRef, &out.IdentityClientIDRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.IdentityClientIDSelector != nil {
		in, out := &in.IdentityClientIDSelector, &out.IdentityClientIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.KeyVaultKeyIdentifier != nil {
		in, out := &in.KeyVaultKeyIdentifier, &out.KeyVaultKeyIdentifier
		*out = new(string)
		**out = **in
	}
	if in.KeyVaultKeyIdentifierRef != nil {
		in, out := &in.KeyVaultKeyIdentifierRef, &out.KeyVaultKeyIdentifierRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.KeyVaultKeyIdentifierSelector != nil {
		in, out := &in.KeyVaultKeyIdentifierSelector, &out.KeyVaultKeyIdentifierSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EncryptionParameters.
func (in *EncryptionParameters) DeepCopy() *EncryptionParameters {
	if in == nil {
		return nil
	}
	out := new(EncryptionParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IdentityInitParameters) DeepCopyInto(out *IdentityInitParameters) {
	*out = *in
	if in.IdentityIds != nil {
		in, out := &in.IdentityIds, &out.IdentityIds
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IdentityInitParameters.
func (in *IdentityInitParameters) DeepCopy() *IdentityInitParameters {
	if in == nil {
		return nil
	}
	out := new(IdentityInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IdentityObservation) DeepCopyInto(out *IdentityObservation) {
	*out = *in
	if in.IdentityIds != nil {
		in, out := &in.IdentityIds, &out.IdentityIds
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.PrincipalID != nil {
		in, out := &in.PrincipalID, &out.PrincipalID
		*out = new(string)
		**out = **in
	}
	if in.TenantID != nil {
		in, out := &in.TenantID, &out.TenantID
		*out = new(string)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IdentityObservation.
func (in *IdentityObservation) DeepCopy() *IdentityObservation {
	if in == nil {
		return nil
	}
	out := new(IdentityObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IdentityParameters) DeepCopyInto(out *IdentityParameters) {
	*out = *in
	if in.IdentityIds != nil {
		in, out := &in.IdentityIds, &out.IdentityIds
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IdentityParameters.
func (in *IdentityParameters) DeepCopy() *IdentityParameters {
	if in == nil {
		return nil
	}
	out := new(IdentityParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PrimaryReadKeyInitParameters) DeepCopyInto(out *PrimaryReadKeyInitParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PrimaryReadKeyInitParameters.
func (in *PrimaryReadKeyInitParameters) DeepCopy() *PrimaryReadKeyInitParameters {
	if in == nil {
		return nil
	}
	out := new(PrimaryReadKeyInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PrimaryReadKeyObservation) DeepCopyInto(out *PrimaryReadKeyObservation) {
	*out = *in
	if in.ConnectionString != nil {
		in, out := &in.ConnectionString, &out.ConnectionString
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.Secret != nil {
		in, out := &in.Secret, &out.Secret
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PrimaryReadKeyObservation.
func (in *PrimaryReadKeyObservation) DeepCopy() *PrimaryReadKeyObservation {
	if in == nil {
		return nil
	}
	out := new(PrimaryReadKeyObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PrimaryReadKeyParameters) DeepCopyInto(out *PrimaryReadKeyParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PrimaryReadKeyParameters.
func (in *PrimaryReadKeyParameters) DeepCopy() *PrimaryReadKeyParameters {
	if in == nil {
		return nil
	}
	out := new(PrimaryReadKeyParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PrimaryWriteKeyInitParameters) DeepCopyInto(out *PrimaryWriteKeyInitParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PrimaryWriteKeyInitParameters.
func (in *PrimaryWriteKeyInitParameters) DeepCopy() *PrimaryWriteKeyInitParameters {
	if in == nil {
		return nil
	}
	out := new(PrimaryWriteKeyInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PrimaryWriteKeyObservation) DeepCopyInto(out *PrimaryWriteKeyObservation) {
	*out = *in
	if in.ConnectionString != nil {
		in, out := &in.ConnectionString, &out.ConnectionString
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.Secret != nil {
		in, out := &in.Secret, &out.Secret
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PrimaryWriteKeyObservation.
func (in *PrimaryWriteKeyObservation) DeepCopy() *PrimaryWriteKeyObservation {
	if in == nil {
		return nil
	}
	out := new(PrimaryWriteKeyObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PrimaryWriteKeyParameters) DeepCopyInto(out *PrimaryWriteKeyParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PrimaryWriteKeyParameters.
func (in *PrimaryWriteKeyParameters) DeepCopy() *PrimaryWriteKeyParameters {
	if in == nil {
		return nil
	}
	out := new(PrimaryWriteKeyParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReplicaInitParameters) DeepCopyInto(out *ReplicaInitParameters) {
	*out = *in
	if in.Location != nil {
		in, out := &in.Location, &out.Location
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReplicaInitParameters.
func (in *ReplicaInitParameters) DeepCopy() *ReplicaInitParameters {
	if in == nil {
		return nil
	}
	out := new(ReplicaInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReplicaObservation) DeepCopyInto(out *ReplicaObservation) {
	*out = *in
	if in.Endpoint != nil {
		in, out := &in.Endpoint, &out.Endpoint
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.Location != nil {
		in, out := &in.Location, &out.Location
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReplicaObservation.
func (in *ReplicaObservation) DeepCopy() *ReplicaObservation {
	if in == nil {
		return nil
	}
	out := new(ReplicaObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReplicaParameters) DeepCopyInto(out *ReplicaParameters) {
	*out = *in
	if in.Location != nil {
		in, out := &in.Location, &out.Location
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReplicaParameters.
func (in *ReplicaParameters) DeepCopy() *ReplicaParameters {
	if in == nil {
		return nil
	}
	out := new(ReplicaParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecondaryReadKeyInitParameters) DeepCopyInto(out *SecondaryReadKeyInitParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecondaryReadKeyInitParameters.
func (in *SecondaryReadKeyInitParameters) DeepCopy() *SecondaryReadKeyInitParameters {
	if in == nil {
		return nil
	}
	out := new(SecondaryReadKeyInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecondaryReadKeyObservation) DeepCopyInto(out *SecondaryReadKeyObservation) {
	*out = *in
	if in.ConnectionString != nil {
		in, out := &in.ConnectionString, &out.ConnectionString
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.Secret != nil {
		in, out := &in.Secret, &out.Secret
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecondaryReadKeyObservation.
func (in *SecondaryReadKeyObservation) DeepCopy() *SecondaryReadKeyObservation {
	if in == nil {
		return nil
	}
	out := new(SecondaryReadKeyObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecondaryReadKeyParameters) DeepCopyInto(out *SecondaryReadKeyParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecondaryReadKeyParameters.
func (in *SecondaryReadKeyParameters) DeepCopy() *SecondaryReadKeyParameters {
	if in == nil {
		return nil
	}
	out := new(SecondaryReadKeyParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecondaryWriteKeyInitParameters) DeepCopyInto(out *SecondaryWriteKeyInitParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecondaryWriteKeyInitParameters.
func (in *SecondaryWriteKeyInitParameters) DeepCopy() *SecondaryWriteKeyInitParameters {
	if in == nil {
		return nil
	}
	out := new(SecondaryWriteKeyInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecondaryWriteKeyObservation) DeepCopyInto(out *SecondaryWriteKeyObservation) {
	*out = *in
	if in.ConnectionString != nil {
		in, out := &in.ConnectionString, &out.ConnectionString
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.Secret != nil {
		in, out := &in.Secret, &out.Secret
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecondaryWriteKeyObservation.
func (in *SecondaryWriteKeyObservation) DeepCopy() *SecondaryWriteKeyObservation {
	if in == nil {
		return nil
	}
	out := new(SecondaryWriteKeyObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecondaryWriteKeyParameters) DeepCopyInto(out *SecondaryWriteKeyParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecondaryWriteKeyParameters.
func (in *SecondaryWriteKeyParameters) DeepCopy() *SecondaryWriteKeyParameters {
	if in == nil {
		return nil
	}
	out := new(SecondaryWriteKeyParameters)
	in.DeepCopyInto(out)
	return out
}
