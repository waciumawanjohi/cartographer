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

package supplychain_test

import (
	"context"
	"encoding/json"
	"time"

	rbacv1 "k8s.io/api/rbac/v1"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gbytes"
	. "github.com/onsi/gomega/gstruct"
	corev1 "k8s.io/api/core/v1"
	apiextensionsv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/yaml"

	"github.com/vmware-tanzu/cartographer/pkg/apis/v1alpha1"
	"github.com/vmware-tanzu/cartographer/pkg/utils"
	"github.com/vmware-tanzu/cartographer/tests/resources"
)

var _ = Describe("WorkloadReconciler", func() {
	var templateBytes = func() []byte {
		configMap := &corev1.ConfigMap{
			TypeMeta: metav1.TypeMeta{
				Kind:       "ConfigMap",
				APIVersion: "v1",
			},
			ObjectMeta: metav1.ObjectMeta{
				Name: "example-config-map",
			},
			Data: map[string]string{},
		}

		templateBytes, err := json.Marshal(configMap)
		Expect(err).ToNot(HaveOccurred())
		return templateBytes
	}

	var newClusterSupplyChain = func(name string, selector map[string]string) *v1alpha1.ClusterSupplyChain {
		return &v1alpha1.ClusterSupplyChain{
			TypeMeta: metav1.TypeMeta{},
			ObjectMeta: metav1.ObjectMeta{
				Name: name,
			},
			Spec: v1alpha1.SupplyChainSpec{
				Resources: []v1alpha1.SupplyChainResource{},
				Selector:  selector,
			},
		}
	}

	var reconcileAgain = func() {
		time.Sleep(1 * time.Second) //metav1.Time unmarshals with 1 second accuracy so this sleep avoids a race condition

		workload := &v1alpha1.Workload{}
		err := c.Get(context.Background(), client.ObjectKey{Name: "workload-bob", Namespace: testNS}, workload)
		Expect(err).NotTo(HaveOccurred())

		workload.Spec.ServiceAccountName = "my-service-account"
		workload.Spec.Params = []v1alpha1.Param{{Name: "foo", Value: apiextensionsv1.JSON{
			Raw: []byte(`"definitelybar"`),
		}}}
		err = c.Update(context.Background(), workload)
		Expect(err).NotTo(HaveOccurred())

		Eventually(func() bool {
			workload := &v1alpha1.Workload{}
			err := c.Get(context.Background(), client.ObjectKey{Name: "workload-bob", Namespace: testNS}, workload)
			Expect(err).NotTo(HaveOccurred())
			return workload.Status.ObservedGeneration == workload.Generation
		}).Should(BeTrue())
	}

	var (
		ctx      context.Context
		cleanups []client.Object
	)

	BeforeEach(func() {
		ctx = context.Background()

		myServiceAccount, err := serviceAccountHelper.CreateServiceAccount("my-service-account", testNS)
		Expect(err).NotTo(HaveOccurred())
		myServiceAccountSecret, err := serviceAccountHelper.CreateAuthableSecret(myServiceAccount)
		Expect(err).NotTo(HaveOccurred())

		cleanups = append(cleanups, myServiceAccount, myServiceAccountSecret)
	})

	AfterEach(func() {
		for _, obj := range cleanups {
			_ = c.Delete(ctx, obj, &client.DeleteOptions{})
		}
	})

	Context("Has the source template and workload installed", func() {
		BeforeEach(func() {
			workload := &v1alpha1.Workload{
				TypeMeta: metav1.TypeMeta{},
				ObjectMeta: metav1.ObjectMeta{
					Name:      "workload-bob",
					Namespace: testNS,
					Labels: map[string]string{
						"name": "webapp",
					},
				},
				Spec: v1alpha1.WorkloadSpec{ServiceAccountName: "my-service-account"},
			}

			cleanups = append(cleanups, workload)
			err := c.Create(ctx, workload, &client.CreateOptions{})
			Expect(err).NotTo(HaveOccurred())

		})

		It("does not update the lastTransitionTime on subsequent reconciliation if the status does not change", func() {
			var lastConditions []metav1.Condition

			Eventually(func() bool {
				workload := &v1alpha1.Workload{}
				err := c.Get(context.Background(), client.ObjectKey{Name: "workload-bob", Namespace: testNS}, workload)
				Expect(err).NotTo(HaveOccurred())
				lastConditions = workload.Status.Conditions
				return workload.Status.ObservedGeneration == workload.Generation
			}, 5*time.Second).Should(BeTrue())

			reconcileAgain()

			workload := &v1alpha1.Workload{}
			err := c.Get(ctx, client.ObjectKey{Name: "workload-bob", Namespace: testNS}, workload)
			Expect(err).NotTo(HaveOccurred())

			Expect(workload.Status.Conditions).To(Equal(lastConditions))
		})

		Context("when reconciliation will end in an unknown status", func() {
			BeforeEach(func() {
				template := &v1alpha1.ClusterSourceTemplate{
					TypeMeta: metav1.TypeMeta{},
					ObjectMeta: metav1.ObjectMeta{
						Name: "proper-template-bob",
					},
					Spec: v1alpha1.SourceTemplateSpec{
						TemplateSpec: v1alpha1.TemplateSpec{
							Template: &runtime.RawExtension{Raw: templateBytes()},
						},
						URLPath: "nonexistant.path",
					},
				}

				cleanups = append(cleanups, template)
				err := c.Create(ctx, template, &client.CreateOptions{})
				Expect(err).NotTo(HaveOccurred())

				supplyChain := newClusterSupplyChain("supplychain-bob", map[string]string{"name": "webapp"})
				supplyChain.Spec.Resources = []v1alpha1.SupplyChainResource{
					{
						Name: "fred-resource",
						TemplateRef: v1alpha1.ClusterTemplateReference{
							Kind: "ClusterSourceTemplate",
							Name: "proper-template-bob",
						},
					},
				}
				cleanups = append(cleanups, supplyChain)

				err = c.Create(ctx, supplyChain, &client.CreateOptions{})
				Expect(err).NotTo(HaveOccurred())
			})

			It("does not error if the reconciliation ends in an unknown status", func() {
				Eventually(func() []metav1.Condition {
					obj := &v1alpha1.Workload{}
					err := c.Get(ctx, client.ObjectKey{Name: "workload-bob", Namespace: testNS}, obj)
					Expect(err).NotTo(HaveOccurred())

					return obj.Status.Conditions
				}, 5*time.Second).Should(ContainElements(
					MatchFields(IgnoreExtras, Fields{
						"Type":   Equal("ResourcesSubmitted"),
						"Reason": Equal("MissingValueAtPath"),
						"Status": Equal(metav1.ConditionStatus("Unknown")),
					}),
					MatchFields(IgnoreExtras, Fields{
						"Type":   Equal("Ready"),
						"Reason": Equal("MissingValueAtPath"),
						"Status": Equal(metav1.ConditionStatus("Unknown")),
					}),
				))
				Expect(controllerBuffer).NotTo(gbytes.Say("Reconciler error.*unable to retrieve outputs from stamped object for resource"))
			})
		})
	})

	Context("insufficient service account permissions", func() {
		var (
			initiallyInsufficientRole *rbacv1.Role
			initiallyInsufficientClusterBoundClusterRole *rbacv1.ClusterRole
			initiallyInsufficientNonClusterBoundClusterRole *rbacv1.ClusterRole
		)
		BeforeEach(func() {
			serviceAccount, err := serviceAccountHelper.CreateServiceAccount("initially-insufficient-service-account", testNS)
			Expect(err).NotTo(HaveOccurred())

			secret, err := serviceAccountHelper.CreateAuthableSecret(serviceAccount)
			Expect(err).NotTo(HaveOccurred())

			cleanups = append(cleanups, serviceAccount, secret)

			roleBindingForRole := &rbacv1.RoleBinding{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "initially-insufficient-role.role-binding",
					Namespace: testNS,
				},
				Subjects: []rbacv1.Subject{
					{
						Kind:      "ServiceAccount",
						Name:      "initially-insufficient-service-account",
						Namespace: testNS,
					},
				},
				RoleRef: rbacv1.RoleRef{
					Kind: "Role",
					Name: "initially-insufficient-role",
				},
			}
			initiallyInsufficientRole = &rbacv1.Role{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "initially-insufficient-role",
					Namespace: testNS,
				},
			}
			roleBindingForClusterRole := &rbacv1.RoleBinding{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "initially-insufficient-non-cluster-bound-cluster-role.role-binding",
					Namespace: testNS,
				},
				Subjects: []rbacv1.Subject{
					{
						Kind:      "ServiceAccount",
						Name:      "initially-insufficient-service-account",
						Namespace: testNS,
					},
				},
				RoleRef: rbacv1.RoleRef{
					Kind: "ClusterRole",
					Name: "initially-insufficient-non-cluster-bound-cluster-role",
				},
			}
			initiallyInsufficientNonClusterBoundClusterRole = &rbacv1.ClusterRole{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "initially-insufficient-non-cluster-bound-cluster-role",
					Namespace: testNS,
				},
			}
			clusterRoleBinding := &rbacv1.ClusterRoleBinding{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "initially-insufficient-cluster-bound-cluster-role.cluster-role-binding",
					Namespace: testNS,
				},
				Subjects: []rbacv1.Subject{
					{
						Kind:      "ServiceAccount",
						Name:      "initially-insufficient-service-account",
						Namespace: testNS,
					},
				},
				RoleRef: rbacv1.RoleRef{
					Kind: "ClusterRole",
					Name: "initially-insufficient-cluster-bound-cluster-role",
				},
			}
			initiallyInsufficientClusterBoundClusterRole = &rbacv1.ClusterRole{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "initially-insufficient-cluster-bound-cluster-role",
					Namespace: testNS,
				},
			}

			err = c.Create(ctx, roleBindingForRole, &client.CreateOptions{})
			Expect(err).NotTo(HaveOccurred())
			cleanups = append(cleanups, roleBindingForRole)
			err = c.Create(ctx, initiallyInsufficientRole, &client.CreateOptions{})
			Expect(err).NotTo(HaveOccurred())
			cleanups = append(cleanups, initiallyInsufficientRole)
			err = c.Create(ctx, roleBindingForClusterRole, &client.CreateOptions{})
			Expect(err).NotTo(HaveOccurred())
			cleanups = append(cleanups, roleBindingForClusterRole)
			err = c.Create(ctx, initiallyInsufficientNonClusterBoundClusterRole, &client.CreateOptions{})
			Expect(err).NotTo(HaveOccurred())
			cleanups = append(cleanups, initiallyInsufficientNonClusterBoundClusterRole)
			err = c.Create(ctx, clusterRoleBinding, &client.CreateOptions{})
			Expect(err).NotTo(HaveOccurred())
			cleanups = append(cleanups, clusterRoleBinding)
			err = c.Create(ctx, initiallyInsufficientClusterBoundClusterRole, &client.CreateOptions{})
			Expect(err).NotTo(HaveOccurred())
			cleanups = append(cleanups, initiallyInsufficientClusterBoundClusterRole)

			template := &v1alpha1.ClusterSourceTemplate{
				TypeMeta: metav1.TypeMeta{},
				ObjectMeta: metav1.ObjectMeta{
					Name: "proper-template-ada",
				},
				Spec: v1alpha1.SourceTemplateSpec{
					TemplateSpec: v1alpha1.TemplateSpec{
						Template: &runtime.RawExtension{Raw: templateBytes()},
					},
				},
			}

			cleanups = append(cleanups, template)
			err = c.Create(ctx, template, &client.CreateOptions{})
			Expect(err).NotTo(HaveOccurred())

			supplyChain := newClusterSupplyChain("supplychain-ada", map[string]string{"name": "webapp"})
			supplyChain.Spec.Resources = []v1alpha1.SupplyChainResource{
				{
					Name: "ada-resource",
					TemplateRef: v1alpha1.ClusterTemplateReference{
						Kind: "ClusterSourceTemplate",
						Name: "proper-template-ada",
					},
				},
			}
			cleanups = append(cleanups, supplyChain)

			err = c.Create(ctx, supplyChain, &client.CreateOptions{})
			Expect(err).NotTo(HaveOccurred())

			workload := &v1alpha1.Workload{
				TypeMeta: metav1.TypeMeta{},
				ObjectMeta: metav1.ObjectMeta{
					Name:      "workload-ada",
					Namespace: testNS,
					Labels: map[string]string{
						"name": "webapp",
					},
				},
				Spec: v1alpha1.WorkloadSpec{ServiceAccountName: "initially-insufficient-service-account"},
			}

			cleanups = append(cleanups, workload)
			err = c.Create(ctx, workload, &client.CreateOptions{})
			Expect(err).NotTo(HaveOccurred())
		})

		It("throws a forbidden error", func() {
			Eventually(func() []metav1.Condition {
				obj := &v1alpha1.Workload{}
				err := c.Get(ctx, client.ObjectKey{Name: "workload-ada", Namespace: testNS}, obj)
				Expect(err).NotTo(HaveOccurred())

				return obj.Status.Conditions
			}, 10*time.Second).Should(ContainElements(
				MatchFields(IgnoreExtras, Fields{
					"Type":    Equal("ResourcesSubmitted"),
					"Reason":  Equal("TemplateRejectedByAPIServer"),
					"Status":  Equal(metav1.ConditionFalse),
					"Message": ContainSubstring("is forbidden"),
				}),
			))
		})

		FIt("reconciles on update to role", func() {
			Eventually(func() []metav1.Condition {
				obj := &v1alpha1.Workload{}
				err := c.Get(ctx, client.ObjectKey{Name: "workload-ada", Namespace: testNS}, obj)
				Expect(err).NotTo(HaveOccurred())

				return obj.Status.Conditions
			}, 10*time.Second).Should(ContainElements(
				MatchFields(IgnoreExtras, Fields{
					"Type":    Equal("ResourcesSubmitted"),
					"Reason":  Equal("TemplateRejectedByAPIServer"),
					"Status":  Equal(metav1.ConditionFalse),
					"Message": ContainSubstring("is forbidden"),
				}),
			))

			//initiallyInsufficientRole.Rules = []rbacv1.PolicyRule{
			//	{
			//		Verbs:     []string{"*"},
			//		APIGroups: []string{"*"},
			//		Resources: []string{"*"},
			//	},
			//}
			//err := c.Update(ctx, initiallyInsufficientRole)
			//Expect(err).NotTo(HaveOccurred())

			Eventually(func() []metav1.Condition {
				obj := &v1alpha1.Workload{}
				err := c.Get(ctx, client.ObjectKey{Name: "workload-ada", Namespace: testNS}, obj)
				Expect(err).NotTo(HaveOccurred())

				return obj.Status.Conditions
			}, 10*time.Second).Should(ContainElements(
				MatchFields(IgnoreExtras, Fields{
					"Type":   Equal("ResourcesSubmitted"),
					"Status": Equal(metav1.ConditionUnknown),
					"Reason": Equal("MissingValueAtPath"),
				}),
			))

			Fail("Stop here so we see debug output")
		})

		It("reconciles on update to clusterrole related by cluster role binding", func() {
			initiallyInsufficientClusterBoundClusterRole.Rules = []rbacv1.PolicyRule{
				{
					Verbs:     []string{"*"},
					APIGroups: []string{"*"},
					Resources: []string{"*"},
				},
			}
			err := c.Update(ctx, initiallyInsufficientClusterBoundClusterRole)
			Expect(err).NotTo(HaveOccurred())

			Eventually(func() []metav1.Condition {
				obj := &v1alpha1.Workload{}
				err := c.Get(ctx, client.ObjectKey{Name: "workload-ada", Namespace: testNS}, obj)
				Expect(err).NotTo(HaveOccurred())

				return obj.Status.Conditions
			}, 10*time.Second).Should(ContainElements(
				MatchFields(IgnoreExtras, Fields{
					"Type":   Equal("ResourcesSubmitted"),
					"Status": Equal(metav1.ConditionUnknown),
					"Reason": Equal("MissingValueAtPath"),
				}),
			))
		})

		It("reconciles on update to clusterrole related by role binding", func() {
			initiallyInsufficientNonClusterBoundClusterRole.Rules = []rbacv1.PolicyRule{
				{
					Verbs:     []string{"*"},
					APIGroups: []string{"*"},
					Resources: []string{"*"},
				},
			}
			err := c.Update(ctx, initiallyInsufficientNonClusterBoundClusterRole)
			Expect(err).NotTo(HaveOccurred())

			Eventually(func() []metav1.Condition {
				obj := &v1alpha1.Workload{}
				err := c.Get(ctx, client.ObjectKey{Name: "workload-ada", Namespace: testNS}, obj)
				Expect(err).NotTo(HaveOccurred())

				return obj.Status.Conditions
			}, 10*time.Second).Should(ContainElements(
				MatchFields(IgnoreExtras, Fields{
					"Type":   Equal("ResourcesSubmitted"),
					"Status": Equal(metav1.ConditionUnknown),
					"Reason": Equal("MissingValueAtPath"),
				}),
			))
		})

		It("reconciles on update to clusterrolebinding", func() {

		})

		It("reconciles on update to rolebinding", func() {

		})
	})

	Context("a supply chain with a template that has stamped a test crd", func() {
		var (
			test *resources.TestObj
		)

		BeforeEach(func() {
			templateYaml := utils.HereYaml(`
				---
				apiVersion: carto.run/v1alpha1
				kind: ClusterConfigTemplate
				metadata:
				  name: my-config-template
				spec:
				  configPath: status.conditions[?(@.type=="Ready")]
			      template:
					apiVersion: test.run/v1alpha1
					kind: TestObj
					metadata:
					  name: test-resource
					spec:
					  foo: "bar"
			`)

			template := &unstructured.Unstructured{}
			err := yaml.Unmarshal([]byte(templateYaml), template)
			Expect(err).NotTo(HaveOccurred())

			err = c.Create(ctx, template, &client.CreateOptions{})
			cleanups = append(cleanups, template)
			Expect(err).NotTo(HaveOccurred())

			supplyChainYaml := utils.HereYaml(`
				---
				apiVersion: carto.run/v1alpha1
				kind: ClusterSupplyChain
				metadata:
				  name: my-supply-chain
				spec:
				  selector:
					"some-key": "some-value"
			      resources:
			        - name: my-first-resource
					  templateRef:
				        kind: ClusterConfigTemplate
				        name: my-config-template
			`)

			supplyChain := &unstructured.Unstructured{}
			err = yaml.Unmarshal([]byte(supplyChainYaml), supplyChain)
			Expect(err).NotTo(HaveOccurred())

			err = c.Create(ctx, supplyChain, &client.CreateOptions{})
			cleanups = append(cleanups, supplyChain)
			Expect(err).NotTo(HaveOccurred())

			workload := &v1alpha1.Workload{
				TypeMeta: metav1.TypeMeta{},
				ObjectMeta: metav1.ObjectMeta{
					Name:      "workload-joe",
					Namespace: testNS,
					Labels: map[string]string{
						"some-key": "some-value",
					},
				},
				Spec: v1alpha1.WorkloadSpec{ServiceAccountName: "my-service-account"},
			}

			cleanups = append(cleanups, workload)
			err = c.Create(ctx, workload, &client.CreateOptions{})
			Expect(err).NotTo(HaveOccurred())

			test = &resources.TestObj{}

			// FIXME: make this more obvious
			Eventually(func() ([]metav1.Condition, error) {
				err := c.Get(ctx, client.ObjectKey{Name: "test-resource", Namespace: testNS}, test)
				return test.Status.Conditions, err
			}).Should(BeNil())

			Eventually(func() []metav1.Condition {
				obj := &v1alpha1.Workload{}
				err := c.Get(ctx, client.ObjectKey{Name: "workload-joe", Namespace: testNS}, obj)
				Expect(err).NotTo(HaveOccurred())

				return obj.Status.Conditions
			}, 5*time.Second).Should(ContainElements(
				MatchFields(IgnoreExtras, Fields{
					"Type":   Equal("SupplyChainReady"),
					"Reason": Equal("Ready"),
					"Status": Equal(metav1.ConditionTrue),
				}),
				MatchFields(IgnoreExtras, Fields{
					"Type":   Equal("ResourcesSubmitted"),
					"Reason": Equal("MissingValueAtPath"),
					"Status": Equal(metav1.ConditionStatus("Unknown")),
				}),
				MatchFields(IgnoreExtras, Fields{
					"Type":   Equal("Ready"),
					"Reason": Equal("MissingValueAtPath"),
					"Status": Equal(metav1.ConditionStatus("Unknown")),
				}),
			))
		})

		Context("a stamped object has changed", func() {

			BeforeEach(func() {
				test.Status.Conditions = []metav1.Condition{
					{
						Type:               "Ready",
						Status:             "True",
						Reason:             "LifeIsGood",
						LastTransitionTime: metav1.Now(),
					},
					{
						Type:               "Succeeded",
						Status:             "True",
						Reason:             "Success",
						LastTransitionTime: metav1.Now(),
					},
				}
				err := c.Status().Update(ctx, test)
				Expect(err).NotTo(HaveOccurred())

				Eventually(func() ([]metav1.Condition, error) {
					err := c.Get(ctx, client.ObjectKey{Name: "test-resource", Namespace: testNS}, test)
					return test.Status.Conditions, err
				}).Should(Not(BeNil()))
			})

			It("immediately reconciles", func() {
				Eventually(func() []metav1.Condition {
					obj := &v1alpha1.Workload{}
					err := c.Get(ctx, client.ObjectKey{Name: "workload-joe", Namespace: testNS}, obj)
					Expect(err).NotTo(HaveOccurred())

					return obj.Status.Conditions
				}, 5*time.Second).Should(ContainElements(
					MatchFields(IgnoreExtras, Fields{
						"Type":   Equal("SupplyChainReady"),
						"Reason": Equal("Ready"),
						"Status": Equal(metav1.ConditionStatus("True")),
					}),
					MatchFields(IgnoreExtras, Fields{
						"Type":   Equal("ResourcesSubmitted"),
						"Reason": Equal("ResourceSubmissionComplete"),
						"Status": Equal(metav1.ConditionStatus("True")),
					}),
					MatchFields(IgnoreExtras, Fields{
						"Type":   Equal("Ready"),
						"Reason": Equal("Ready"),
						"Status": Equal(metav1.ConditionStatus("True")),
					}),
				))
			})
		})
	})
})
