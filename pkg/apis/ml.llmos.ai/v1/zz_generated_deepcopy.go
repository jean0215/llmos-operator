//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2024 llmos.ai.

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

package v1

import (
	common "github.com/llmos-ai/llmos-operator/pkg/apis/common"
	corev1 "k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ModelDetails) DeepCopyInto(out *ModelDetails) {
	*out = *in
	if in.Families != nil {
		in, out := &in.Families, &out.Families
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ModelDetails.
func (in *ModelDetails) DeepCopy() *ModelDetails {
	if in == nil {
		return nil
	}
	out := new(ModelDetails)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ModelFile) DeepCopyInto(out *ModelFile) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ModelFile.
func (in *ModelFile) DeepCopy() *ModelFile {
	if in == nil {
		return nil
	}
	out := new(ModelFile)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ModelFile) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ModelFileList) DeepCopyInto(out *ModelFileList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ModelFile, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ModelFileList.
func (in *ModelFileList) DeepCopy() *ModelFileList {
	if in == nil {
		return nil
	}
	out := new(ModelFileList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ModelFileList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ModelFileSpec) DeepCopyInto(out *ModelFileSpec) {
	*out = *in
	if in.PromptSuggestions != nil {
		in, out := &in.PromptSuggestions, &out.PromptSuggestions
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Categories != nil {
		in, out := &in.Categories, &out.Categories
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ModelFileSpec.
func (in *ModelFileSpec) DeepCopy() *ModelFileSpec {
	if in == nil {
		return nil
	}
	out := new(ModelFileSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ModelFileStatus) DeepCopyInto(out *ModelFileStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]common.Condition, len(*in))
		copy(*out, *in)
	}
	in.Details.DeepCopyInto(&out.Details)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ModelFileStatus.
func (in *ModelFileStatus) DeepCopy() *ModelFileStatus {
	if in == nil {
		return nil
	}
	out := new(ModelFileStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ModelService) DeepCopyInto(out *ModelService) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ModelService.
func (in *ModelService) DeepCopy() *ModelService {
	if in == nil {
		return nil
	}
	out := new(ModelService)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ModelService) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ModelServiceList) DeepCopyInto(out *ModelServiceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ModelService, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ModelServiceList.
func (in *ModelServiceList) DeepCopy() *ModelServiceList {
	if in == nil {
		return nil
	}
	out := new(ModelServiceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ModelServiceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ModelServiceSpec) DeepCopyInto(out *ModelServiceSpec) {
	*out = *in
	if in.VolumeClaimTemplates != nil {
		in, out := &in.VolumeClaimTemplates, &out.VolumeClaimTemplates
		*out = make([]corev1.PersistentVolumeClaim, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	in.UpdateStrategy.DeepCopyInto(&out.UpdateStrategy)
	in.Template.DeepCopyInto(&out.Template)
	if in.Accelerators != nil {
		in, out := &in.Accelerators, &out.Accelerators
		*out = make(map[string]byte, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ModelServiceSpec.
func (in *ModelServiceSpec) DeepCopy() *ModelServiceSpec {
	if in == nil {
		return nil
	}
	out := new(ModelServiceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ModelServiceStatus) DeepCopyInto(out *ModelServiceStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]common.Condition, len(*in))
		copy(*out, *in)
	}
	in.ContainerState.DeepCopyInto(&out.ContainerState)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ModelServiceStatus.
func (in *ModelServiceStatus) DeepCopy() *ModelServiceStatus {
	if in == nil {
		return nil
	}
	out := new(ModelServiceStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ModelServiceTemplateSpec) DeepCopyInto(out *ModelServiceTemplateSpec) {
	*out = *in
	in.Spec.DeepCopyInto(&out.Spec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ModelServiceTemplateSpec.
func (in *ModelServiceTemplateSpec) DeepCopy() *ModelServiceTemplateSpec {
	if in == nil {
		return nil
	}
	out := new(ModelServiceTemplateSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Notebook) DeepCopyInto(out *Notebook) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Notebook.
func (in *Notebook) DeepCopy() *Notebook {
	if in == nil {
		return nil
	}
	out := new(Notebook)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Notebook) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NotebookList) DeepCopyInto(out *NotebookList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Notebook, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NotebookList.
func (in *NotebookList) DeepCopy() *NotebookList {
	if in == nil {
		return nil
	}
	out := new(NotebookList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NotebookList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NotebookSpec) DeepCopyInto(out *NotebookSpec) {
	*out = *in
	in.Template.DeepCopyInto(&out.Template)
	if in.Volumes != nil {
		in, out := &in.Volumes, &out.Volumes
		*out = make([]Volume, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NotebookSpec.
func (in *NotebookSpec) DeepCopy() *NotebookSpec {
	if in == nil {
		return nil
	}
	out := new(NotebookSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NotebookStatus) DeepCopyInto(out *NotebookStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]common.Condition, len(*in))
		copy(*out, *in)
	}
	in.State.DeepCopyInto(&out.State)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NotebookStatus.
func (in *NotebookStatus) DeepCopy() *NotebookStatus {
	if in == nil {
		return nil
	}
	out := new(NotebookStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NotebookTemplateSpec) DeepCopyInto(out *NotebookTemplateSpec) {
	*out = *in
	in.Spec.DeepCopyInto(&out.Spec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NotebookTemplateSpec.
func (in *NotebookTemplateSpec) DeepCopy() *NotebookTemplateSpec {
	if in == nil {
		return nil
	}
	out := new(NotebookTemplateSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Volume) DeepCopyInto(out *Volume) {
	*out = *in
	in.Spec.DeepCopyInto(&out.Spec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Volume.
func (in *Volume) DeepCopy() *Volume {
	if in == nil {
		return nil
	}
	out := new(Volume)
	in.DeepCopyInto(out)
	return out
}
