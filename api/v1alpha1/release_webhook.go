/*
Copyright 2022.

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

package v1alpha1

import (
	"context"
	"fmt"
	"github.com/go-logr/logr"
	"k8s.io/apimachinery/pkg/runtime"
	"reflect"
	ctrl "sigs.k8s.io/controller-runtime"
)

// ReleaseWebhook describes the data structure for the release webhook
type ReleaseWebhook struct{}

func (w *ReleaseWebhook) Register(mgr ctrl.Manager, log *logr.Logger) error {
	return ctrl.NewWebhookManagedBy(mgr).
		For(&Release{}).
		WithValidator(w).
		Complete()
}

//+kubebuilder:webhook:path=/validate-appstudio-redhat-com-v1alpha1-release,mutating=false,failurePolicy=fail,sideEffects=None,groups=appstudio.redhat.com,resources=releases,verbs=create;update,versions=v1alpha1,name=vrelease.kb.io,admissionReviewVersions=v1

// ValidateCreate implements webhook.Validator so a webhook will be registered for the type
func (w *ReleaseWebhook) ValidateCreate(ctx context.Context, obj runtime.Object) error {
	return nil
}

// ValidateUpdate implements webhook.Validator so a webhook will be registered for the type
func (w *ReleaseWebhook) ValidateUpdate(ctx context.Context, oldObj, newObj runtime.Object) error {
	oldRelease := oldObj.(*Release)
	newRelease := newObj.(*Release)

	if !reflect.DeepEqual(newRelease.Spec, oldRelease.Spec) {
		return fmt.Errorf("release resources spec cannot be updated")
	}

	return nil
}

// ValidateDelete implements webhook.Validator so a webhook will be registered for the type
func (w *ReleaseWebhook) ValidateDelete(ctx context.Context, obj runtime.Object) error {
	return nil
}
