//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Code generated by diegen. DO NOT EDIT.

package dies

import (
	v1 "dies.dev/apis/meta/v1"
	json "encoding/json"
	fmtx "fmt"
	"github.com/vmware-tanzu/cartographer/pkg/apis/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	unstructured "k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	runtime "k8s.io/apimachinery/pkg/runtime"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
)

var RunnableBlank = (&RunnableDie{}).DieFeed(v1alpha1.Runnable{})

type RunnableDie struct {
	v1.FrozenObjectMeta
	mutable bool
	r       v1alpha1.Runnable
}

// DieImmutable returns a new die for the current die's state that is either mutable (`false`) or immutable (`true`).
func (d *RunnableDie) DieImmutable(immutable bool) *RunnableDie {
	if d.mutable == !immutable {
		return d
	}
	d = d.DeepCopy()
	d.mutable = !immutable
	return d
}

// DieFeed returns a new die with the provided resource.
func (d *RunnableDie) DieFeed(r v1alpha1.Runnable) *RunnableDie {
	if d.mutable {
		d.FrozenObjectMeta = v1.FreezeObjectMeta(r.ObjectMeta)
		d.r = r
		return d
	}
	return &RunnableDie{
		FrozenObjectMeta: v1.FreezeObjectMeta(r.ObjectMeta),
		mutable:          d.mutable,
		r:                r,
	}
}

// DieFeedPtr returns a new die with the provided resource pointer. If the resource is nil, the empty value is used instead.
func (d *RunnableDie) DieFeedPtr(r *v1alpha1.Runnable) *RunnableDie {
	if r == nil {
		r = &v1alpha1.Runnable{}
	}
	return d.DieFeed(*r)
}

// DieFeedRawExtension returns the resource managed by the die as an raw extension.
func (d *RunnableDie) DieFeedRawExtension(raw runtime.RawExtension) *RunnableDie {
	b, _ := json.Marshal(raw)
	r := v1alpha1.Runnable{}
	_ = json.Unmarshal(b, &r)
	return d.DieFeed(r)
}

// DieRelease returns the resource managed by the die.
func (d *RunnableDie) DieRelease() v1alpha1.Runnable {
	if d.mutable {
		return d.r
	}
	return *d.r.DeepCopy()
}

// DieReleasePtr returns a pointer to the resource managed by the die.
func (d *RunnableDie) DieReleasePtr() *v1alpha1.Runnable {
	r := d.DieRelease()
	return &r
}

// DieReleaseUnstructured returns the resource managed by the die as an unstructured object.
func (d *RunnableDie) DieReleaseUnstructured() runtime.Unstructured {
	r := d.DieReleasePtr()
	u, _ := runtime.DefaultUnstructuredConverter.ToUnstructured(r)
	return &unstructured.Unstructured{
		Object: u,
	}
}

// DieReleaseRawExtension returns the resource managed by the die as an raw extension.
func (d *RunnableDie) DieReleaseRawExtension() runtime.RawExtension {
	r := d.DieReleasePtr()
	b, _ := json.Marshal(r)
	raw := runtime.RawExtension{}
	_ = json.Unmarshal(b, &raw)
	return raw
}

// DieStamp returns a new die with the resource passed to the callback function. The resource is mutable.
func (d *RunnableDie) DieStamp(fn func(r *v1alpha1.Runnable)) *RunnableDie {
	r := d.DieRelease()
	fn(&r)
	return d.DieFeed(r)
}

// DeepCopy returns a new die with equivalent state. Useful for snapshotting a mutable die.
func (d *RunnableDie) DeepCopy() *RunnableDie {
	r := *d.r.DeepCopy()
	return &RunnableDie{
		FrozenObjectMeta: v1.FreezeObjectMeta(r.ObjectMeta),
		mutable:          d.mutable,
		r:                r,
	}
}

var _ runtime.Object = (*RunnableDie)(nil)

func (d *RunnableDie) DeepCopyObject() runtime.Object {
	return d.r.DeepCopy()
}

func (d *RunnableDie) GetObjectKind() schema.ObjectKind {
	r := d.DieRelease()
	return r.GetObjectKind()
}

func (d *RunnableDie) MarshalJSON() ([]byte, error) {
	return json.Marshal(d.r)
}

func (d *RunnableDie) UnmarshalJSON(b []byte) error {
	if d == RunnableBlank {
		return fmtx.Errorf("cannot unmarshal into the blank die, create a copy first")
	}
	if !d.mutable {
		return fmtx.Errorf("cannot unmarshal into immutable dies, create a mutable version first")
	}
	r := &v1alpha1.Runnable{}
	err := json.Unmarshal(b, r)
	*d = *d.DieFeed(*r)
	return err
}

// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
func (d *RunnableDie) APIVersion(v string) *RunnableDie {
	return d.DieStamp(func(r *v1alpha1.Runnable) {
		r.APIVersion = v
	})
}

// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
func (d *RunnableDie) Kind(v string) *RunnableDie {
	return d.DieStamp(func(r *v1alpha1.Runnable) {
		r.Kind = v
	})
}

// MetadataDie stamps the resource's ObjectMeta field with a mutable die.
func (d *RunnableDie) MetadataDie(fn func(d *v1.ObjectMetaDie)) *RunnableDie {
	return d.DieStamp(func(r *v1alpha1.Runnable) {
		d := v1.ObjectMetaBlank.DieImmutable(false).DieFeed(r.ObjectMeta)
		fn(d)
		r.ObjectMeta = d.DieRelease()
	})
}

// SpecDie stamps the resource's spec field with a mutable die.
func (d *RunnableDie) SpecDie(fn func(d *RunnableSpecDie)) *RunnableDie {
	return d.DieStamp(func(r *v1alpha1.Runnable) {
		d := RunnableSpecBlank.DieImmutable(false).DieFeed(r.Spec)
		fn(d)
		r.Spec = d.DieRelease()
	})
}

// StatusDie stamps the resource's status field with a mutable die.
func (d *RunnableDie) StatusDie(fn func(d *RunnableStatusDie)) *RunnableDie {
	return d.DieStamp(func(r *v1alpha1.Runnable) {
		d := RunnableStatusBlank.DieImmutable(false).DieFeed(r.Status)
		fn(d)
		r.Status = d.DieRelease()
	})
}

// Spec describes the runnable. More info: https://cartographer.sh/docs/latest/reference/runnable/#runnable
func (d *RunnableDie) Spec(v v1alpha1.RunnableSpec) *RunnableDie {
	return d.DieStamp(func(r *v1alpha1.Runnable) {
		r.Spec = v
	})
}

// Status conforms to the Kubernetes conventions: https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#typical-status-properties
func (d *RunnableDie) Status(v v1alpha1.RunnableStatus) *RunnableDie {
	return d.DieStamp(func(r *v1alpha1.Runnable) {
		r.Status = v
	})
}

var RunnableSpecBlank = (&RunnableSpecDie{}).DieFeed(v1alpha1.RunnableSpec{})

type RunnableSpecDie struct {
	mutable bool
	r       v1alpha1.RunnableSpec
}

// DieImmutable returns a new die for the current die's state that is either mutable (`false`) or immutable (`true`).
func (d *RunnableSpecDie) DieImmutable(immutable bool) *RunnableSpecDie {
	if d.mutable == !immutable {
		return d
	}
	d = d.DeepCopy()
	d.mutable = !immutable
	return d
}

// DieFeed returns a new die with the provided resource.
func (d *RunnableSpecDie) DieFeed(r v1alpha1.RunnableSpec) *RunnableSpecDie {
	if d.mutable {
		d.r = r
		return d
	}
	return &RunnableSpecDie{
		mutable: d.mutable,
		r:       r,
	}
}

// DieFeedPtr returns a new die with the provided resource pointer. If the resource is nil, the empty value is used instead.
func (d *RunnableSpecDie) DieFeedPtr(r *v1alpha1.RunnableSpec) *RunnableSpecDie {
	if r == nil {
		r = &v1alpha1.RunnableSpec{}
	}
	return d.DieFeed(*r)
}

// DieFeedRawExtension returns the resource managed by the die as an raw extension.
func (d *RunnableSpecDie) DieFeedRawExtension(raw runtime.RawExtension) *RunnableSpecDie {
	b, _ := json.Marshal(raw)
	r := v1alpha1.RunnableSpec{}
	_ = json.Unmarshal(b, &r)
	return d.DieFeed(r)
}

// DieRelease returns the resource managed by the die.
func (d *RunnableSpecDie) DieRelease() v1alpha1.RunnableSpec {
	if d.mutable {
		return d.r
	}
	return *d.r.DeepCopy()
}

// DieReleasePtr returns a pointer to the resource managed by the die.
func (d *RunnableSpecDie) DieReleasePtr() *v1alpha1.RunnableSpec {
	r := d.DieRelease()
	return &r
}

// DieReleaseRawExtension returns the resource managed by the die as an raw extension.
func (d *RunnableSpecDie) DieReleaseRawExtension() runtime.RawExtension {
	r := d.DieReleasePtr()
	b, _ := json.Marshal(r)
	raw := runtime.RawExtension{}
	_ = json.Unmarshal(b, &raw)
	return raw
}

// DieStamp returns a new die with the resource passed to the callback function. The resource is mutable.
func (d *RunnableSpecDie) DieStamp(fn func(r *v1alpha1.RunnableSpec)) *RunnableSpecDie {
	r := d.DieRelease()
	fn(&r)
	return d.DieFeed(r)
}

// DeepCopy returns a new die with equivalent state. Useful for snapshotting a mutable die.
func (d *RunnableSpecDie) DeepCopy() *RunnableSpecDie {
	r := *d.r.DeepCopy()
	return &RunnableSpecDie{
		mutable: d.mutable,
		r:       r,
	}
}

// RunTemplateRef identifies the run template used to produce resources for this runnable.
func (d *RunnableSpecDie) RunTemplateRef(v v1alpha1.TemplateReference) *RunnableSpecDie {
	return d.DieStamp(func(r *v1alpha1.RunnableSpec) {
		r.RunTemplateRef = v
	})
}

// Selector refers to an additional object that the template can refer to using: $(selected)$.
func (d *RunnableSpecDie) Selector(v *v1alpha1.ResourceSelector) *RunnableSpecDie {
	return d.DieStamp(func(r *v1alpha1.RunnableSpec) {
		r.Selector = v
	})
}

// ServiceAccountName refers to the Service account with permissions to create resources submitted by the ClusterRunTemplate.
//
// If not set, Cartographer will use the default service account in the runnable's namespace.
func (d *RunnableSpecDie) ServiceAccountName(v string) *RunnableSpecDie {
	return d.DieStamp(func(r *v1alpha1.RunnableSpec) {
		r.ServiceAccountName = v
	})
}

// RetentionPolicy specifies how many successful and failed runs should be retained. Runs older than this (ordered by creation time) will be deleted. Setting higher values will increase memory footprint.
func (d *RunnableSpecDie) RetentionPolicy(v v1alpha1.RetentionPolicy) *RunnableSpecDie {
	return d.DieStamp(func(r *v1alpha1.RunnableSpec) {
		r.RetentionPolicy = v
	})
}

var RunnableStatusBlank = (&RunnableStatusDie{}).DieFeed(v1alpha1.RunnableStatus{})

type RunnableStatusDie struct {
	mutable bool
	r       v1alpha1.RunnableStatus
}

// DieImmutable returns a new die for the current die's state that is either mutable (`false`) or immutable (`true`).
func (d *RunnableStatusDie) DieImmutable(immutable bool) *RunnableStatusDie {
	if d.mutable == !immutable {
		return d
	}
	d = d.DeepCopy()
	d.mutable = !immutable
	return d
}

// DieFeed returns a new die with the provided resource.
func (d *RunnableStatusDie) DieFeed(r v1alpha1.RunnableStatus) *RunnableStatusDie {
	if d.mutable {
		d.r = r
		return d
	}
	return &RunnableStatusDie{
		mutable: d.mutable,
		r:       r,
	}
}

// DieFeedPtr returns a new die with the provided resource pointer. If the resource is nil, the empty value is used instead.
func (d *RunnableStatusDie) DieFeedPtr(r *v1alpha1.RunnableStatus) *RunnableStatusDie {
	if r == nil {
		r = &v1alpha1.RunnableStatus{}
	}
	return d.DieFeed(*r)
}

// DieFeedRawExtension returns the resource managed by the die as an raw extension.
func (d *RunnableStatusDie) DieFeedRawExtension(raw runtime.RawExtension) *RunnableStatusDie {
	b, _ := json.Marshal(raw)
	r := v1alpha1.RunnableStatus{}
	_ = json.Unmarshal(b, &r)
	return d.DieFeed(r)
}

// DieRelease returns the resource managed by the die.
func (d *RunnableStatusDie) DieRelease() v1alpha1.RunnableStatus {
	if d.mutable {
		return d.r
	}
	return *d.r.DeepCopy()
}

// DieReleasePtr returns a pointer to the resource managed by the die.
func (d *RunnableStatusDie) DieReleasePtr() *v1alpha1.RunnableStatus {
	r := d.DieRelease()
	return &r
}

// DieReleaseRawExtension returns the resource managed by the die as an raw extension.
func (d *RunnableStatusDie) DieReleaseRawExtension() runtime.RawExtension {
	r := d.DieReleasePtr()
	b, _ := json.Marshal(r)
	raw := runtime.RawExtension{}
	_ = json.Unmarshal(b, &raw)
	return raw
}

// DieStamp returns a new die with the resource passed to the callback function. The resource is mutable.
func (d *RunnableStatusDie) DieStamp(fn func(r *v1alpha1.RunnableStatus)) *RunnableStatusDie {
	r := d.DieRelease()
	fn(&r)
	return d.DieFeed(r)
}

// DeepCopy returns a new die with equivalent state. Useful for snapshotting a mutable die.
func (d *RunnableStatusDie) DeepCopy() *RunnableStatusDie {
	r := *d.r.DeepCopy()
	return &RunnableStatusDie{
		mutable: d.mutable,
		r:       r,
	}
}

func (d *RunnableStatusDie) ObservedGeneration(v int64) *RunnableStatusDie {
	return d.DieStamp(func(r *v1alpha1.RunnableStatus) {
		r.ObservedGeneration = v
	})
}

func (d *RunnableStatusDie) Conditions(v ...metav1.Condition) *RunnableStatusDie {
	return d.DieStamp(func(r *v1alpha1.RunnableStatus) {
		r.Conditions = v
	})
}