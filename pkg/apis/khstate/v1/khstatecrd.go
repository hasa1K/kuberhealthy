// +build !ignore_autogenerated

/*
Copyright The Kubernetes Authors.

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

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1

import (
	"log"

	"k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KuberhealthyState) DeepCopyInto(out *KuberhealthyState) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KuberhealthyState.
func (in *KuberhealthyState) DeepCopy() *KuberhealthyState {
	if in == nil {
		return nil
	}
	out := new(KuberhealthyState)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KuberhealthyState) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KuberhealthyStateList) DeepCopyInto(out *KuberhealthyStateList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]KuberhealthyState, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KuberhealthyStateList.
func (in *KuberhealthyStateList) DeepCopy() *KuberhealthyStateList {
	if in == nil {
		return nil
	}
	out := new(KuberhealthyStateList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KuberhealthyStateList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkloadDetails) DeepCopyInto(out *WorkloadDetails) {
	*out = *in
	if in.Errors != nil {
		in, out := &in.Errors, &out.Errors
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	in.LastRun.DeepCopyInto(out.LastRun)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkloadDetails.
func (in *WorkloadDetails) DeepCopy() *WorkloadDetails {
	if in == nil {
		return nil
	}
	out := new(WorkloadDetails)
	in.DeepCopyInto(out)
	return out
}

// NewKuberhealthyState creates a KuberhealthyState struct which represents
// the data inside a KuberhealthyState resource
func NewKuberhealthyState(name string, spec WorkloadDetails) KuberhealthyState {
	state := KuberhealthyState{}
	state.SetName(name)
	state.Spec = spec
	return state
}

// NewWorkloadDetails creates a new WorkloadDetails struct
func NewWorkloadDetails(workloadType KHWorkload) WorkloadDetails {
	if workloadType == "" {
		log.Panic("Creating workload details with empty workload type.  This will probably cause errors when the struct is used.")
	}
	return WorkloadDetails{
		Errors:     []string{},
		khWorkload: &workloadType,
		Successes:  []string{},
	}
}

// GetKHWorkload returns the workload for the WorkloadDetails struct
func (wd *WorkloadDetails) GetKHWorkload() KHWorkload {
	// failsafe if the workload is empty
	if len(*wd.khWorkload) == 0 {
		log.Panicln("Fetched a workload type from a WorkloadDetails struct, but it was blank!")
	}
	return *wd.khWorkload
}
