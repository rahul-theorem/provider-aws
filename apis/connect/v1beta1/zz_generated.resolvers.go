/*
Copyright 2022 Upbound Inc.
*/
// Code generated by angryjet. DO NOT EDIT.

package v1beta1

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	errors "github.com/pkg/errors"
	v1beta1 "github.com/upbound/provider-aws/apis/lexmodels/v1beta1"
	resource "github.com/upbound/upjet/pkg/resource"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this BotAssociation.
func (mg *BotAssociation) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.InstanceID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.InstanceIDRef,
		Selector:     mg.Spec.ForProvider.InstanceIDSelector,
		To: reference.To{
			List:    &InstanceList{},
			Managed: &Instance{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.InstanceID")
	}
	mg.Spec.ForProvider.InstanceID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.InstanceIDRef = rsp.ResolvedReference

	for i3 := 0; i3 < len(mg.Spec.ForProvider.LexBot); i3++ {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.LexBot[i3].Name),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.LexBot[i3].NameRef,
			Selector:     mg.Spec.ForProvider.LexBot[i3].NameSelector,
			To: reference.To{
				List:    &v1beta1.BotList{},
				Managed: &v1beta1.Bot{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.LexBot[i3].Name")
		}
		mg.Spec.ForProvider.LexBot[i3].Name = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.LexBot[i3].NameRef = rsp.ResolvedReference

	}

	return nil
}

// ResolveReferences of this ContactFlow.
func (mg *ContactFlow) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.InstanceID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.InstanceIDRef,
		Selector:     mg.Spec.ForProvider.InstanceIDSelector,
		To: reference.To{
			List:    &InstanceList{},
			Managed: &Instance{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.InstanceID")
	}
	mg.Spec.ForProvider.InstanceID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.InstanceIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this ContactFlowModule.
func (mg *ContactFlowModule) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.InstanceID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.InstanceIDRef,
		Selector:     mg.Spec.ForProvider.InstanceIDSelector,
		To: reference.To{
			List:    &InstanceList{},
			Managed: &Instance{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.InstanceID")
	}
	mg.Spec.ForProvider.InstanceID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.InstanceIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this HoursOfOperation.
func (mg *HoursOfOperation) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.InstanceID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.InstanceIDRef,
		Selector:     mg.Spec.ForProvider.InstanceIDSelector,
		To: reference.To{
			List:    &InstanceList{},
			Managed: &Instance{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.InstanceID")
	}
	mg.Spec.ForProvider.InstanceID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.InstanceIDRef = rsp.ResolvedReference

	return nil
}
