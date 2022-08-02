/*
Copyright 2022 Upbound Inc.
*/
// Code generated by angryjet. DO NOT EDIT.

package v1beta1

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	errors "github.com/pkg/errors"
	v1beta1 "github.com/upbound/official-providers/provider-azure/apis/azure/v1beta1"
	v1beta11 "github.com/upbound/official-providers/provider-azure/apis/operationalinsights/v1beta1"
	rconfig "github.com/upbound/official-providers/provider-azure/apis/rconfig"
	v1beta12 "github.com/upbound/official-providers/provider-azure/apis/storage/v1beta1"
	resource "github.com/upbound/upjet/pkg/resource"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this ApplicationInsights.
func (mg *ApplicationInsights) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ResourceGroupName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ResourceGroupNameRef,
		Selector:     mg.Spec.ForProvider.ResourceGroupNameSelector,
		To: reference.To{
			List:    &v1beta1.ResourceGroupList{},
			Managed: &v1beta1.ResourceGroup{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ResourceGroupName")
	}
	mg.Spec.ForProvider.ResourceGroupName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ResourceGroupNameRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.WorkspaceID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.WorkspaceIDRef,
		Selector:     mg.Spec.ForProvider.WorkspaceIDSelector,
		To: reference.To{
			List:    &v1beta11.WorkspaceList{},
			Managed: &v1beta11.Workspace{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.WorkspaceID")
	}
	mg.Spec.ForProvider.WorkspaceID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.WorkspaceIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this MonitorActionGroup.
func (mg *MonitorActionGroup) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ResourceGroupName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ResourceGroupNameRef,
		Selector:     mg.Spec.ForProvider.ResourceGroupNameSelector,
		To: reference.To{
			List:    &v1beta1.ResourceGroupList{},
			Managed: &v1beta1.ResourceGroup{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ResourceGroupName")
	}
	mg.Spec.ForProvider.ResourceGroupName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ResourceGroupNameRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this MonitorMetricAlert.
func (mg *MonitorMetricAlert) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var mrsp reference.MultiResolutionResponse
	var err error

	for i3 := 0; i3 < len(mg.Spec.ForProvider.Action); i3++ {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Action[i3].ActionGroupID),
			Extract:      rconfig.ExtractResourceID(),
			Reference:    mg.Spec.ForProvider.Action[i3].ActionGroupIDRef,
			Selector:     mg.Spec.ForProvider.Action[i3].ActionGroupIDSelector,
			To: reference.To{
				List:    &MonitorActionGroupList{},
				Managed: &MonitorActionGroup{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.Action[i3].ActionGroupID")
		}
		mg.Spec.ForProvider.Action[i3].ActionGroupID = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.Action[i3].ActionGroupIDRef = rsp.ResolvedReference

	}
	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ResourceGroupName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ResourceGroupNameRef,
		Selector:     mg.Spec.ForProvider.ResourceGroupNameSelector,
		To: reference.To{
			List:    &v1beta1.ResourceGroupList{},
			Managed: &v1beta1.ResourceGroup{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ResourceGroupName")
	}
	mg.Spec.ForProvider.ResourceGroupName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ResourceGroupNameRef = rsp.ResolvedReference

	mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
		CurrentValues: reference.FromPtrValues(mg.Spec.ForProvider.Scopes),
		Extract:       rconfig.ExtractResourceID(),
		References:    mg.Spec.ForProvider.ScopesRefs,
		Selector:      mg.Spec.ForProvider.ScopesSelector,
		To: reference.To{
			List:    &v1beta12.AccountList{},
			Managed: &v1beta12.Account{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Scopes")
	}
	mg.Spec.ForProvider.Scopes = reference.ToPtrValues(mrsp.ResolvedValues)
	mg.Spec.ForProvider.ScopesRefs = mrsp.ResolvedReferences

	return nil
}

// ResolveReferences of this MonitorPrivateLinkScope.
func (mg *MonitorPrivateLinkScope) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ResourceGroupName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ResourceGroupNameRef,
		Selector:     mg.Spec.ForProvider.ResourceGroupNameSelector,
		To: reference.To{
			List:    &v1beta1.ResourceGroupList{},
			Managed: &v1beta1.ResourceGroup{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ResourceGroupName")
	}
	mg.Spec.ForProvider.ResourceGroupName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ResourceGroupNameRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this MonitorPrivateLinkScopedService.
func (mg *MonitorPrivateLinkScopedService) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.LinkedResourceID),
		Extract:      rconfig.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.LinkedResourceIDRef,
		Selector:     mg.Spec.ForProvider.LinkedResourceIDSelector,
		To: reference.To{
			List:    &ApplicationInsightsList{},
			Managed: &ApplicationInsights{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.LinkedResourceID")
	}
	mg.Spec.ForProvider.LinkedResourceID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.LinkedResourceIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ResourceGroupName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ResourceGroupNameRef,
		Selector:     mg.Spec.ForProvider.ResourceGroupNameSelector,
		To: reference.To{
			List:    &v1beta1.ResourceGroupList{},
			Managed: &v1beta1.ResourceGroup{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ResourceGroupName")
	}
	mg.Spec.ForProvider.ResourceGroupName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ResourceGroupNameRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ScopeName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ScopeNameRef,
		Selector:     mg.Spec.ForProvider.ScopeNameSelector,
		To: reference.To{
			List:    &MonitorPrivateLinkScopeList{},
			Managed: &MonitorPrivateLinkScope{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ScopeName")
	}
	mg.Spec.ForProvider.ScopeName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ScopeNameRef = rsp.ResolvedReference

	return nil
}
