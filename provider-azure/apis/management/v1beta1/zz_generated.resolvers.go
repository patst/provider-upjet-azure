/*
Copyright 2022 Upbound Inc.
*/
// Code generated by angryjet. DO NOT EDIT.

package v1beta1

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	errors "github.com/pkg/errors"
	resource "github.com/upbound/upjet/pkg/resource"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this ManagementGroup.
func (mg *ManagementGroup) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ParentManagementGroupID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.ParentManagementGroupIDRef,
		Selector:     mg.Spec.ForProvider.ParentManagementGroupIDSelector,
		To: reference.To{
			List:    &ManagementGroupList{},
			Managed: &ManagementGroup{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ParentManagementGroupID")
	}
	mg.Spec.ForProvider.ParentManagementGroupID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ParentManagementGroupIDRef = rsp.ResolvedReference

	return nil
}
