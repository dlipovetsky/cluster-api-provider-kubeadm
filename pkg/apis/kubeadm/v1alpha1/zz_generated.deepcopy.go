// +build !ignore_autogenerated

/*
Copyright 2018 The Kubernetes Authors.

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
// Code generated by main. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubeadmClusterProviderSpec) DeepCopyInto(out *KubeadmClusterProviderSpec) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubeadmClusterProviderSpec.
func (in *KubeadmClusterProviderSpec) DeepCopy() *KubeadmClusterProviderSpec {
	if in == nil {
		return nil
	}
	out := new(KubeadmClusterProviderSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KubeadmClusterProviderSpec) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubeadmClusterProviderSpecList) DeepCopyInto(out *KubeadmClusterProviderSpecList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]KubeadmClusterProviderSpec, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubeadmClusterProviderSpecList.
func (in *KubeadmClusterProviderSpecList) DeepCopy() *KubeadmClusterProviderSpecList {
	if in == nil {
		return nil
	}
	out := new(KubeadmClusterProviderSpecList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KubeadmClusterProviderSpecList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubeadmClusterProviderSpecSpec) DeepCopyInto(out *KubeadmClusterProviderSpecSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubeadmClusterProviderSpecSpec.
func (in *KubeadmClusterProviderSpecSpec) DeepCopy() *KubeadmClusterProviderSpecSpec {
	if in == nil {
		return nil
	}
	out := new(KubeadmClusterProviderSpecSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubeadmClusterProviderSpecStatus) DeepCopyInto(out *KubeadmClusterProviderSpecStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubeadmClusterProviderSpecStatus.
func (in *KubeadmClusterProviderSpecStatus) DeepCopy() *KubeadmClusterProviderSpecStatus {
	if in == nil {
		return nil
	}
	out := new(KubeadmClusterProviderSpecStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubeadmClusterProviderStatus) DeepCopyInto(out *KubeadmClusterProviderStatus) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubeadmClusterProviderStatus.
func (in *KubeadmClusterProviderStatus) DeepCopy() *KubeadmClusterProviderStatus {
	if in == nil {
		return nil
	}
	out := new(KubeadmClusterProviderStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KubeadmClusterProviderStatus) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubeadmClusterProviderStatusList) DeepCopyInto(out *KubeadmClusterProviderStatusList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]KubeadmClusterProviderStatus, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubeadmClusterProviderStatusList.
func (in *KubeadmClusterProviderStatusList) DeepCopy() *KubeadmClusterProviderStatusList {
	if in == nil {
		return nil
	}
	out := new(KubeadmClusterProviderStatusList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KubeadmClusterProviderStatusList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubeadmClusterProviderStatusSpec) DeepCopyInto(out *KubeadmClusterProviderStatusSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubeadmClusterProviderStatusSpec.
func (in *KubeadmClusterProviderStatusSpec) DeepCopy() *KubeadmClusterProviderStatusSpec {
	if in == nil {
		return nil
	}
	out := new(KubeadmClusterProviderStatusSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubeadmClusterProviderStatusStatus) DeepCopyInto(out *KubeadmClusterProviderStatusStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubeadmClusterProviderStatusStatus.
func (in *KubeadmClusterProviderStatusStatus) DeepCopy() *KubeadmClusterProviderStatusStatus {
	if in == nil {
		return nil
	}
	out := new(KubeadmClusterProviderStatusStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubeadmMachineProviderSpec) DeepCopyInto(out *KubeadmMachineProviderSpec) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubeadmMachineProviderSpec.
func (in *KubeadmMachineProviderSpec) DeepCopy() *KubeadmMachineProviderSpec {
	if in == nil {
		return nil
	}
	out := new(KubeadmMachineProviderSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KubeadmMachineProviderSpec) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubeadmMachineProviderSpecList) DeepCopyInto(out *KubeadmMachineProviderSpecList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]KubeadmMachineProviderSpec, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubeadmMachineProviderSpecList.
func (in *KubeadmMachineProviderSpecList) DeepCopy() *KubeadmMachineProviderSpecList {
	if in == nil {
		return nil
	}
	out := new(KubeadmMachineProviderSpecList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KubeadmMachineProviderSpecList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubeadmMachineProviderSpecSpec) DeepCopyInto(out *KubeadmMachineProviderSpecSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubeadmMachineProviderSpecSpec.
func (in *KubeadmMachineProviderSpecSpec) DeepCopy() *KubeadmMachineProviderSpecSpec {
	if in == nil {
		return nil
	}
	out := new(KubeadmMachineProviderSpecSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubeadmMachineProviderSpecStatus) DeepCopyInto(out *KubeadmMachineProviderSpecStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubeadmMachineProviderSpecStatus.
func (in *KubeadmMachineProviderSpecStatus) DeepCopy() *KubeadmMachineProviderSpecStatus {
	if in == nil {
		return nil
	}
	out := new(KubeadmMachineProviderSpecStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubeadmMachineProviderStatus) DeepCopyInto(out *KubeadmMachineProviderStatus) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubeadmMachineProviderStatus.
func (in *KubeadmMachineProviderStatus) DeepCopy() *KubeadmMachineProviderStatus {
	if in == nil {
		return nil
	}
	out := new(KubeadmMachineProviderStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KubeadmMachineProviderStatus) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubeadmMachineProviderStatusList) DeepCopyInto(out *KubeadmMachineProviderStatusList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]KubeadmMachineProviderStatus, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubeadmMachineProviderStatusList.
func (in *KubeadmMachineProviderStatusList) DeepCopy() *KubeadmMachineProviderStatusList {
	if in == nil {
		return nil
	}
	out := new(KubeadmMachineProviderStatusList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KubeadmMachineProviderStatusList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubeadmMachineProviderStatusSpec) DeepCopyInto(out *KubeadmMachineProviderStatusSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubeadmMachineProviderStatusSpec.
func (in *KubeadmMachineProviderStatusSpec) DeepCopy() *KubeadmMachineProviderStatusSpec {
	if in == nil {
		return nil
	}
	out := new(KubeadmMachineProviderStatusSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubeadmMachineProviderStatusStatus) DeepCopyInto(out *KubeadmMachineProviderStatusStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubeadmMachineProviderStatusStatus.
func (in *KubeadmMachineProviderStatusStatus) DeepCopy() *KubeadmMachineProviderStatusStatus {
	if in == nil {
		return nil
	}
	out := new(KubeadmMachineProviderStatusStatus)
	in.DeepCopyInto(out)
	return out
}
