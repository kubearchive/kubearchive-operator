/*
Copyright 2026.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1

import (
	"context"
	"fmt"

	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"

	kubearchiveorgv1 "github.com/kubearchive/kubearchive-operator/api/v1"
)

// nolint:unused
// log is for logging in this package.
var kubearchiveinstallationlog = logf.Log.WithName("kubearchiveinstallation-resource")

// SetupKubeArchiveInstallationWebhookWithManager registers the webhook for KubeArchiveInstallation in the manager.
func SetupKubeArchiveInstallationWebhookWithManager(mgr ctrl.Manager) error {
	return ctrl.NewWebhookManagedBy(mgr).For(&kubearchiveorgv1.KubeArchiveInstallation{}).
		WithValidator(&KubeArchiveInstallationCustomValidator{}).
		WithDefaulter(&KubeArchiveInstallationCustomDefaulter{}).
		Complete()
}

// +kubebuilder:webhook:path=/mutate-kubearchive-org-v1-kubearchiveinstallation,mutating=true,failurePolicy=fail,sideEffects=None,groups=kubearchive.org,resources=kubearchiveinstallations,verbs=create;update,versions=v1,name=mkubearchiveinstallation-v1.kb.io,admissionReviewVersions=v1

// KubeArchiveInstallationCustomDefaulter struct is responsible for setting default values on the custom resource of the
// Kind KubeArchiveInstallation when those are created or updated.
//
// NOTE: The +kubebuilder:object:generate=false marker prevents controller-gen from generating DeepCopy methods,
// as it is used only for temporary operations and does not need to be deeply copied.
type KubeArchiveInstallationCustomDefaulter struct {
	// TODO(user): Add more fields as needed for defaulting
}

var _ webhook.CustomDefaulter = &KubeArchiveInstallationCustomDefaulter{}

// Default implements webhook.CustomDefaulter so a webhook will be registered for the Kind KubeArchiveInstallation.
func (d *KubeArchiveInstallationCustomDefaulter) Default(_ context.Context, obj runtime.Object) error {
	kubearchiveinstallation, ok := obj.(*kubearchiveorgv1.KubeArchiveInstallation)

	if !ok {
		return fmt.Errorf("expected an KubeArchiveInstallation object but got %T", obj)
	}
	kubearchiveinstallationlog.Info("Defaulting for KubeArchiveInstallation", "name", kubearchiveinstallation.GetName())

	// TODO(user): fill in your defaulting logic.

	return nil
}

// TODO(user): change verbs to "verbs=create;update;delete" if you want to enable deletion validation.
// NOTE: The 'path' attribute must follow a specific pattern and should not be modified directly here.
// Modifying the path for an invalid path can cause API server errors; failing to locate the webhook.
// +kubebuilder:webhook:path=/validate-kubearchive-org-v1-kubearchiveinstallation,mutating=false,failurePolicy=fail,sideEffects=None,groups=kubearchive.org,resources=kubearchiveinstallations,verbs=create;update,versions=v1,name=vkubearchiveinstallation-v1.kb.io,admissionReviewVersions=v1

// KubeArchiveInstallationCustomValidator struct is responsible for validating the KubeArchiveInstallation resource
// when it is created, updated, or deleted.
//
// NOTE: The +kubebuilder:object:generate=false marker prevents controller-gen from generating DeepCopy methods,
// as this struct is used only for temporary operations and does not need to be deeply copied.
type KubeArchiveInstallationCustomValidator struct {
	// TODO(user): Add more fields as needed for validation
}

var _ webhook.CustomValidator = &KubeArchiveInstallationCustomValidator{}

// ValidateCreate implements webhook.CustomValidator so a webhook will be registered for the type KubeArchiveInstallation.
func (v *KubeArchiveInstallationCustomValidator) ValidateCreate(_ context.Context, obj runtime.Object) (admission.Warnings, error) {
	kubearchiveinstallation, ok := obj.(*kubearchiveorgv1.KubeArchiveInstallation)
	if !ok {
		return nil, fmt.Errorf("expected a KubeArchiveInstallation object but got %T", obj)
	}
	kubearchiveinstallationlog.Info("Validation for KubeArchiveInstallation upon creation", "name", kubearchiveinstallation.GetName())

	return nil, nil
}

// ValidateUpdate implements webhook.CustomValidator so a webhook will be registered for the type KubeArchiveInstallation.
func (v *KubeArchiveInstallationCustomValidator) ValidateUpdate(_ context.Context, oldObj, newObj runtime.Object) (admission.Warnings, error) {
	kubearchiveinstallation, ok := newObj.(*kubearchiveorgv1.KubeArchiveInstallation)
	if !ok {
		return nil, fmt.Errorf("expected a KubeArchiveInstallation object for the newObj but got %T", newObj)
	}
	kubearchiveinstallationlog.Info("Validation for KubeArchiveInstallation upon update", "name", kubearchiveinstallation.GetName())

	// TODO(user): fill in your validation logic upon object update.

	return nil, nil
}

// ValidateDelete implements webhook.CustomValidator so a webhook will be registered for the type KubeArchiveInstallation.
func (v *KubeArchiveInstallationCustomValidator) ValidateDelete(ctx context.Context, obj runtime.Object) (admission.Warnings, error) {
	kubearchiveinstallation, ok := obj.(*kubearchiveorgv1.KubeArchiveInstallation)
	if !ok {
		return nil, fmt.Errorf("expected a KubeArchiveInstallation object but got %T", obj)
	}
	kubearchiveinstallationlog.Info("Validation for KubeArchiveInstallation upon deletion", "name", kubearchiveinstallation.GetName())

	// TODO(user): fill in your validation logic upon object deletion.

	return nil, nil
}
