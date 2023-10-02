//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2020 The Crossplane Authors.

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

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AwcatorKind1) DeepCopyInto(out *AwcatorKind1) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AwcatorKind1.
func (in *AwcatorKind1) DeepCopy() *AwcatorKind1 {
	if in == nil {
		return nil
	}
	out := new(AwcatorKind1)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AwcatorKind1) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AwcatorKind1List) DeepCopyInto(out *AwcatorKind1List) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AwcatorKind1, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AwcatorKind1List.
func (in *AwcatorKind1List) DeepCopy() *AwcatorKind1List {
	if in == nil {
		return nil
	}
	out := new(AwcatorKind1List)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AwcatorKind1List) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AwcatorKind1Observation) DeepCopyInto(out *AwcatorKind1Observation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AwcatorKind1Observation.
func (in *AwcatorKind1Observation) DeepCopy() *AwcatorKind1Observation {
	if in == nil {
		return nil
	}
	out := new(AwcatorKind1Observation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AwcatorKind1Parameters) DeepCopyInto(out *AwcatorKind1Parameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AwcatorKind1Parameters.
func (in *AwcatorKind1Parameters) DeepCopy() *AwcatorKind1Parameters {
	if in == nil {
		return nil
	}
	out := new(AwcatorKind1Parameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AwcatorKind1Spec) DeepCopyInto(out *AwcatorKind1Spec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	out.ForProvider = in.ForProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AwcatorKind1Spec.
func (in *AwcatorKind1Spec) DeepCopy() *AwcatorKind1Spec {
	if in == nil {
		return nil
	}
	out := new(AwcatorKind1Spec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AwcatorKind1Status) DeepCopyInto(out *AwcatorKind1Status) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	out.AtProvider = in.AtProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AwcatorKind1Status.
func (in *AwcatorKind1Status) DeepCopy() *AwcatorKind1Status {
	if in == nil {
		return nil
	}
	out := new(AwcatorKind1Status)
	in.DeepCopyInto(out)
	return out
}