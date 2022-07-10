//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Code generated by controller-gen. DO NOT EDIT.

package v2alpha1

import (
	apiextensionsv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BlueprintParam) DeepCopyInto(out *BlueprintParam) {
	*out = *in
	if in.Value != nil {
		in, out := &in.Value, &out.Value
		*out = new(apiextensionsv1.JSON)
		(*in).DeepCopyInto(*out)
	}
	if in.DefaultValue != nil {
		in, out := &in.DefaultValue, &out.DefaultValue
		*out = new(apiextensionsv1.JSON)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BlueprintParam.
func (in *BlueprintParam) DeepCopy() *BlueprintParam {
	if in == nil {
		return nil
	}
	out := new(BlueprintParam)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BlueprintRef) DeepCopyInto(out *BlueprintRef) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BlueprintRef.
func (in *BlueprintRef) DeepCopy() *BlueprintRef {
	if in == nil {
		return nil
	}
	out := new(BlueprintRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BlueprintSpec) DeepCopyInto(out *BlueprintSpec) {
	*out = *in
	out.OutputTypeRef = in.OutputTypeRef
	if in.Components != nil {
		in, out := &in.Components, &out.Components
		*out = make([]Component, len(*in))
		copy(*out, *in)
	}
	in.Template.DeepCopyInto(&out.Template)
	if in.Params != nil {
		in, out := &in.Params, &out.Params
		*out = make([]BlueprintParam, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	out.ServiceAccountRef = in.ServiceAccountRef
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BlueprintSpec.
func (in *BlueprintSpec) DeepCopy() *BlueprintSpec {
	if in == nil {
		return nil
	}
	out := new(BlueprintSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BlueprintStatus) DeepCopyInto(out *BlueprintStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Calculated != nil {
		in, out := &in.Calculated, &out.Calculated
		*out = make([]CalculatedParam, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BlueprintStatus.
func (in *BlueprintStatus) DeepCopy() *BlueprintStatus {
	if in == nil {
		return nil
	}
	out := new(BlueprintStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CalculatedParam) DeepCopyInto(out *CalculatedParam) {
	*out = *in
	if in.Value != nil {
		in, out := &in.Value, &out.Value
		*out = new(apiextensionsv1.JSON)
		(*in).DeepCopyInto(*out)
	}
	if in.Default != nil {
		in, out := &in.Default, &out.Default
		*out = new(apiextensionsv1.JSON)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CalculatedParam.
func (in *CalculatedParam) DeepCopy() *CalculatedParam {
	if in == nil {
		return nil
	}
	out := new(CalculatedParam)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterBlueprint) DeepCopyInto(out *ClusterBlueprint) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterBlueprint.
func (in *ClusterBlueprint) DeepCopy() *ClusterBlueprint {
	if in == nil {
		return nil
	}
	out := new(ClusterBlueprint)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterBlueprint) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterBlueprintList) DeepCopyInto(out *ClusterBlueprintList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ClusterBlueprint, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterBlueprintList.
func (in *ClusterBlueprintList) DeepCopy() *ClusterBlueprintList {
	if in == nil {
		return nil
	}
	out := new(ClusterBlueprintList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterBlueprintList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterOutputType) DeepCopyInto(out *ClusterOutputType) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Schema.DeepCopyInto(&out.Schema)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterOutputType.
func (in *ClusterOutputType) DeepCopy() *ClusterOutputType {
	if in == nil {
		return nil
	}
	out := new(ClusterOutputType)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterOutputType) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterOutputTypeList) DeepCopyInto(out *ClusterOutputTypeList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ClusterOutputType, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterOutputTypeList.
func (in *ClusterOutputTypeList) DeepCopy() *ClusterOutputTypeList {
	if in == nil {
		return nil
	}
	out := new(ClusterOutputTypeList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterOutputTypeList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterSelector) DeepCopyInto(out *ClusterSelector) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterSelector.
func (in *ClusterSelector) DeepCopy() *ClusterSelector {
	if in == nil {
		return nil
	}
	out := new(ClusterSelector)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterSelector) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterSelectorList) DeepCopyInto(out *ClusterSelectorList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ClusterSelector, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterSelectorList.
func (in *ClusterSelectorList) DeepCopy() *ClusterSelectorList {
	if in == nil {
		return nil
	}
	out := new(ClusterSelectorList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterSelectorList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Component) DeepCopyInto(out *Component) {
	*out = *in
	out.BlueprintRef = in.BlueprintRef
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Component.
func (in *Component) DeepCopy() *Component {
	if in == nil {
		return nil
	}
	out := new(Component)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConditionRequirement) DeepCopyInto(out *ConditionRequirement) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConditionRequirement.
func (in *ConditionRequirement) DeepCopy() *ConditionRequirement {
	if in == nil {
		return nil
	}
	out := new(ConditionRequirement)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FieldSelectorRequirement) DeepCopyInto(out *FieldSelectorRequirement) {
	*out = *in
	if in.Values != nil {
		in, out := &in.Values, &out.Values
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FieldSelectorRequirement.
func (in *FieldSelectorRequirement) DeepCopy() *FieldSelectorRequirement {
	if in == nil {
		return nil
	}
	out := new(FieldSelectorRequirement)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HealthMatchFieldSelectorRequirement) DeepCopyInto(out *HealthMatchFieldSelectorRequirement) {
	*out = *in
	in.FieldSelectorRequirement.DeepCopyInto(&out.FieldSelectorRequirement)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HealthMatchFieldSelectorRequirement.
func (in *HealthMatchFieldSelectorRequirement) DeepCopy() *HealthMatchFieldSelectorRequirement {
	if in == nil {
		return nil
	}
	out := new(HealthMatchFieldSelectorRequirement)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HealthMatchRule) DeepCopyInto(out *HealthMatchRule) {
	*out = *in
	if in.MatchConditions != nil {
		in, out := &in.MatchConditions, &out.MatchConditions
		*out = make([]ConditionRequirement, len(*in))
		copy(*out, *in)
	}
	if in.MatchFields != nil {
		in, out := &in.MatchFields, &out.MatchFields
		*out = make([]HealthMatchFieldSelectorRequirement, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HealthMatchRule.
func (in *HealthMatchRule) DeepCopy() *HealthMatchRule {
	if in == nil {
		return nil
	}
	out := new(HealthMatchRule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HealthRule) DeepCopyInto(out *HealthRule) {
	*out = *in
	if in.AlwaysHealthy != nil {
		in, out := &in.AlwaysHealthy, &out.AlwaysHealthy
		*out = new(runtime.RawExtension)
		(*in).DeepCopyInto(*out)
	}
	if in.MultiMatch != nil {
		in, out := &in.MultiMatch, &out.MultiMatch
		*out = new(MultiMatchHealthRule)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HealthRule.
func (in *HealthRule) DeepCopy() *HealthRule {
	if in == nil {
		return nil
	}
	out := new(HealthRule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImprintStatus) DeepCopyInto(out *ImprintStatus) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImprintStatus.
func (in *ImprintStatus) DeepCopy() *ImprintStatus {
	if in == nil {
		return nil
	}
	out := new(ImprintStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ImprintStatus) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImprintStatusList) DeepCopyInto(out *ImprintStatusList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ImprintStatus, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImprintStatusList.
func (in *ImprintStatusList) DeepCopy() *ImprintStatusList {
	if in == nil {
		return nil
	}
	out := new(ImprintStatusList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ImprintStatusList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MultiMatchHealthRule) DeepCopyInto(out *MultiMatchHealthRule) {
	*out = *in
	in.Healthy.DeepCopyInto(&out.Healthy)
	in.Unhealthy.DeepCopyInto(&out.Unhealthy)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MultiMatchHealthRule.
func (in *MultiMatchHealthRule) DeepCopy() *MultiMatchHealthRule {
	if in == nil {
		return nil
	}
	out := new(MultiMatchHealthRule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OutputTypeRef) DeepCopyInto(out *OutputTypeRef) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OutputTypeRef.
func (in *OutputTypeRef) DeepCopy() *OutputTypeRef {
	if in == nil {
		return nil
	}
	out := new(OutputTypeRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ParameterMapping) DeepCopyInto(out *ParameterMapping) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ParameterMapping.
func (in *ParameterMapping) DeepCopy() *ParameterMapping {
	if in == nil {
		return nil
	}
	out := new(ParameterMapping)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SelectorSpec) DeepCopyInto(out *SelectorSpec) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.BlueprintRef = in.BlueprintRef
	if in.ParamMap != nil {
		in, out := &in.ParamMap, &out.ParamMap
		*out = make([]ParameterMapping, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SelectorSpec.
func (in *SelectorSpec) DeepCopy() *SelectorSpec {
	if in == nil {
		return nil
	}
	out := new(SelectorSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceAccountRef) DeepCopyInto(out *ServiceAccountRef) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceAccountRef.
func (in *ServiceAccountRef) DeepCopy() *ServiceAccountRef {
	if in == nil {
		return nil
	}
	out := new(ServiceAccountRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TemplateSpec) DeepCopyInto(out *TemplateSpec) {
	*out = *in
	if in.Template != nil {
		in, out := &in.Template, &out.Template
		*out = new(runtime.RawExtension)
		(*in).DeepCopyInto(*out)
	}
	if in.HealthRule != nil {
		in, out := &in.HealthRule, &out.HealthRule
		*out = new(HealthRule)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TemplateSpec.
func (in *TemplateSpec) DeepCopy() *TemplateSpec {
	if in == nil {
		return nil
	}
	out := new(TemplateSpec)
	in.DeepCopyInto(out)
	return out
}
