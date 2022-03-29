// Copyright 2021 VMware
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package v1alpha1

import (
	"fmt"

	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
)

// +kubebuilder:webhook:path=/validate-carto-run-v1alpha1-clusterimagetemplate,mutating=false,failurePolicy=fail,sideEffects=none,admissionReviewVersions=v1beta1;v1,groups=carto.run,resources=clusterimagetemplates,verbs=create;update,versions=v1alpha1,name=image-template-validator.cartographer.com

var _ webhook.Validator = &ClusterImageTemplate{}

func (c *ClusterImageTemplate) ValidateCreate() error {
	if err := c.validate(); err != nil {
		return fmt.Errorf("error validating clusterimagetemplate [%s]: %w", c.Name, err)
	}
	return nil
}

func (c *ClusterImageTemplate) ValidateUpdate(_ runtime.Object) error {
	if err := c.validate(); err != nil {
		return fmt.Errorf("error validating clusterimagetemplate [%s]: %w", c.Name, err)
	}
	return nil
}

func (c *ClusterImageTemplate) ValidateDelete() error {
	return nil
}

func (c *ClusterImageTemplate) validate() error {
	err := validateName(c.ObjectMeta)
	if err != nil {
		return err
	}

	return c.Spec.TemplateSpec.validate()
}

func (c *ClusterImageTemplate) SetupWebhookWithManager(mgr ctrl.Manager) error {
	return ctrl.NewWebhookManagedBy(mgr).
		For(c).
		Complete()
}
