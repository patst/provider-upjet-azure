//go:build !ignore_autogenerated

// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by controller-gen. DO NOT EDIT.

package v1beta2

import (
	"github.com/crossplane/crossplane-runtime/apis/common/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

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
func (in *Workspace) DeepCopyInto(out *Workspace) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Workspace.
func (in *Workspace) DeepCopy() *Workspace {
	if in == nil {
		return nil
	}
	out := new(Workspace)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Workspace) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkspaceInitParameters) DeepCopyInto(out *WorkspaceInitParameters) {
	*out = *in
	if in.AllowResourceOnlyPermissions != nil {
		in, out := &in.AllowResourceOnlyPermissions, &out.AllowResourceOnlyPermissions
		*out = new(bool)
		**out = **in
	}
	if in.CmkForQueryForced != nil {
		in, out := &in.CmkForQueryForced, &out.CmkForQueryForced
		*out = new(bool)
		**out = **in
	}
	if in.DailyQuotaGb != nil {
		in, out := &in.DailyQuotaGb, &out.DailyQuotaGb
		*out = new(float64)
		**out = **in
	}
	if in.DataCollectionRuleID != nil {
		in, out := &in.DataCollectionRuleID, &out.DataCollectionRuleID
		*out = new(string)
		**out = **in
	}
	if in.Identity != nil {
		in, out := &in.Identity, &out.Identity
		*out = new(IdentityInitParameters)
		(*in).DeepCopyInto(*out)
	}
	if in.ImmediateDataPurgeOn30DaysEnabled != nil {
		in, out := &in.ImmediateDataPurgeOn30DaysEnabled, &out.ImmediateDataPurgeOn30DaysEnabled
		*out = new(bool)
		**out = **in
	}
	if in.InternetIngestionEnabled != nil {
		in, out := &in.InternetIngestionEnabled, &out.InternetIngestionEnabled
		*out = new(bool)
		**out = **in
	}
	if in.InternetQueryEnabled != nil {
		in, out := &in.InternetQueryEnabled, &out.InternetQueryEnabled
		*out = new(bool)
		**out = **in
	}
	if in.LocalAuthenticationDisabled != nil {
		in, out := &in.LocalAuthenticationDisabled, &out.LocalAuthenticationDisabled
		*out = new(bool)
		**out = **in
	}
	if in.Location != nil {
		in, out := &in.Location, &out.Location
		*out = new(string)
		**out = **in
	}
	if in.ReservationCapacityInGbPerDay != nil {
		in, out := &in.ReservationCapacityInGbPerDay, &out.ReservationCapacityInGbPerDay
		*out = new(float64)
		**out = **in
	}
	if in.RetentionInDays != nil {
		in, out := &in.RetentionInDays, &out.RetentionInDays
		*out = new(float64)
		**out = **in
	}
	if in.Sku != nil {
		in, out := &in.Sku, &out.Sku
		*out = new(string)
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkspaceInitParameters.
func (in *WorkspaceInitParameters) DeepCopy() *WorkspaceInitParameters {
	if in == nil {
		return nil
	}
	out := new(WorkspaceInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkspaceList) DeepCopyInto(out *WorkspaceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Workspace, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkspaceList.
func (in *WorkspaceList) DeepCopy() *WorkspaceList {
	if in == nil {
		return nil
	}
	out := new(WorkspaceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *WorkspaceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkspaceObservation) DeepCopyInto(out *WorkspaceObservation) {
	*out = *in
	if in.AllowResourceOnlyPermissions != nil {
		in, out := &in.AllowResourceOnlyPermissions, &out.AllowResourceOnlyPermissions
		*out = new(bool)
		**out = **in
	}
	if in.CmkForQueryForced != nil {
		in, out := &in.CmkForQueryForced, &out.CmkForQueryForced
		*out = new(bool)
		**out = **in
	}
	if in.DailyQuotaGb != nil {
		in, out := &in.DailyQuotaGb, &out.DailyQuotaGb
		*out = new(float64)
		**out = **in
	}
	if in.DataCollectionRuleID != nil {
		in, out := &in.DataCollectionRuleID, &out.DataCollectionRuleID
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
		*out = new(IdentityObservation)
		(*in).DeepCopyInto(*out)
	}
	if in.ImmediateDataPurgeOn30DaysEnabled != nil {
		in, out := &in.ImmediateDataPurgeOn30DaysEnabled, &out.ImmediateDataPurgeOn30DaysEnabled
		*out = new(bool)
		**out = **in
	}
	if in.InternetIngestionEnabled != nil {
		in, out := &in.InternetIngestionEnabled, &out.InternetIngestionEnabled
		*out = new(bool)
		**out = **in
	}
	if in.InternetQueryEnabled != nil {
		in, out := &in.InternetQueryEnabled, &out.InternetQueryEnabled
		*out = new(bool)
		**out = **in
	}
	if in.LocalAuthenticationDisabled != nil {
		in, out := &in.LocalAuthenticationDisabled, &out.LocalAuthenticationDisabled
		*out = new(bool)
		**out = **in
	}
	if in.Location != nil {
		in, out := &in.Location, &out.Location
		*out = new(string)
		**out = **in
	}
	if in.ReservationCapacityInGbPerDay != nil {
		in, out := &in.ReservationCapacityInGbPerDay, &out.ReservationCapacityInGbPerDay
		*out = new(float64)
		**out = **in
	}
	if in.ResourceGroupName != nil {
		in, out := &in.ResourceGroupName, &out.ResourceGroupName
		*out = new(string)
		**out = **in
	}
	if in.RetentionInDays != nil {
		in, out := &in.RetentionInDays, &out.RetentionInDays
		*out = new(float64)
		**out = **in
	}
	if in.Sku != nil {
		in, out := &in.Sku, &out.Sku
		*out = new(string)
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
	if in.WorkspaceID != nil {
		in, out := &in.WorkspaceID, &out.WorkspaceID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkspaceObservation.
func (in *WorkspaceObservation) DeepCopy() *WorkspaceObservation {
	if in == nil {
		return nil
	}
	out := new(WorkspaceObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkspaceParameters) DeepCopyInto(out *WorkspaceParameters) {
	*out = *in
	if in.AllowResourceOnlyPermissions != nil {
		in, out := &in.AllowResourceOnlyPermissions, &out.AllowResourceOnlyPermissions
		*out = new(bool)
		**out = **in
	}
	if in.CmkForQueryForced != nil {
		in, out := &in.CmkForQueryForced, &out.CmkForQueryForced
		*out = new(bool)
		**out = **in
	}
	if in.DailyQuotaGb != nil {
		in, out := &in.DailyQuotaGb, &out.DailyQuotaGb
		*out = new(float64)
		**out = **in
	}
	if in.DataCollectionRuleID != nil {
		in, out := &in.DataCollectionRuleID, &out.DataCollectionRuleID
		*out = new(string)
		**out = **in
	}
	if in.Identity != nil {
		in, out := &in.Identity, &out.Identity
		*out = new(IdentityParameters)
		(*in).DeepCopyInto(*out)
	}
	if in.ImmediateDataPurgeOn30DaysEnabled != nil {
		in, out := &in.ImmediateDataPurgeOn30DaysEnabled, &out.ImmediateDataPurgeOn30DaysEnabled
		*out = new(bool)
		**out = **in
	}
	if in.InternetIngestionEnabled != nil {
		in, out := &in.InternetIngestionEnabled, &out.InternetIngestionEnabled
		*out = new(bool)
		**out = **in
	}
	if in.InternetQueryEnabled != nil {
		in, out := &in.InternetQueryEnabled, &out.InternetQueryEnabled
		*out = new(bool)
		**out = **in
	}
	if in.LocalAuthenticationDisabled != nil {
		in, out := &in.LocalAuthenticationDisabled, &out.LocalAuthenticationDisabled
		*out = new(bool)
		**out = **in
	}
	if in.Location != nil {
		in, out := &in.Location, &out.Location
		*out = new(string)
		**out = **in
	}
	if in.ReservationCapacityInGbPerDay != nil {
		in, out := &in.ReservationCapacityInGbPerDay, &out.ReservationCapacityInGbPerDay
		*out = new(float64)
		**out = **in
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
	if in.RetentionInDays != nil {
		in, out := &in.RetentionInDays, &out.RetentionInDays
		*out = new(float64)
		**out = **in
	}
	if in.Sku != nil {
		in, out := &in.Sku, &out.Sku
		*out = new(string)
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkspaceParameters.
func (in *WorkspaceParameters) DeepCopy() *WorkspaceParameters {
	if in == nil {
		return nil
	}
	out := new(WorkspaceParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkspaceSpec) DeepCopyInto(out *WorkspaceSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
	in.InitProvider.DeepCopyInto(&out.InitProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkspaceSpec.
func (in *WorkspaceSpec) DeepCopy() *WorkspaceSpec {
	if in == nil {
		return nil
	}
	out := new(WorkspaceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkspaceStatus) DeepCopyInto(out *WorkspaceStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkspaceStatus.
func (in *WorkspaceStatus) DeepCopy() *WorkspaceStatus {
	if in == nil {
		return nil
	}
	out := new(WorkspaceStatus)
	in.DeepCopyInto(out)
	return out
}