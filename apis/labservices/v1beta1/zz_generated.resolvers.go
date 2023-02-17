/*
Copyright 2022 Upbound Inc.
*/
// Code generated by angryjet. DO NOT EDIT.

package v1beta1

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	errors "github.com/pkg/errors"
	v1beta11 "github.com/upbound/provider-azure/apis/azure/v1beta1"
	v1beta1 "github.com/upbound/provider-azure/apis/network/v1beta1"
	rconfig "github.com/upbound/provider-azure/apis/rconfig"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this LabServiceLab.
func (mg *LabServiceLab) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	for i3 := 0; i3 < len(mg.Spec.ForProvider.Network); i3++ {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Network[i3].SubnetID),
			Extract:      rconfig.ExtractResourceID(),
			Reference:    mg.Spec.ForProvider.Network[i3].SubnetIDRef,
			Selector:     mg.Spec.ForProvider.Network[i3].SubnetIDSelector,
			To: reference.To{
				List:    &v1beta1.SubnetList{},
				Managed: &v1beta1.Subnet{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.Network[i3].SubnetID")
		}
		mg.Spec.ForProvider.Network[i3].SubnetID = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.Network[i3].SubnetIDRef = rsp.ResolvedReference

	}
	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ResourceGroupName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ResourceGroupNameRef,
		Selector:     mg.Spec.ForProvider.ResourceGroupNameSelector,
		To: reference.To{
			List:    &v1beta11.ResourceGroupList{},
			Managed: &v1beta11.ResourceGroup{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ResourceGroupName")
	}
	mg.Spec.ForProvider.ResourceGroupName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ResourceGroupNameRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this LabServicePlan.
func (mg *LabServicePlan) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.DefaultNetworkSubnetID),
		Extract:      rconfig.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.DefaultNetworkSubnetIDRef,
		Selector:     mg.Spec.ForProvider.DefaultNetworkSubnetIDSelector,
		To: reference.To{
			List:    &v1beta1.SubnetList{},
			Managed: &v1beta1.Subnet{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.DefaultNetworkSubnetID")
	}
	mg.Spec.ForProvider.DefaultNetworkSubnetID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.DefaultNetworkSubnetIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ResourceGroupName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ResourceGroupNameRef,
		Selector:     mg.Spec.ForProvider.ResourceGroupNameSelector,
		To: reference.To{
			List:    &v1beta11.ResourceGroupList{},
			Managed: &v1beta11.ResourceGroup{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ResourceGroupName")
	}
	mg.Spec.ForProvider.ResourceGroupName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ResourceGroupNameRef = rsp.ResolvedReference

	return nil
}
