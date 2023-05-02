/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

	spatialanchorsaccount "github.com/upbound/provider-azure/internal/controller/mixedreality/spatialanchorsaccount"
)

// Setup_mixedreality creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_mixedreality(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		spatialanchorsaccount.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}