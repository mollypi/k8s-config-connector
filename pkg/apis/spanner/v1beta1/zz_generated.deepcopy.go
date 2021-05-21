// +build !ignore_autogenerated

// Copyright 2020 Google LLC
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

// *** DISCLAIMER ***
// Config Connector's go-client for CRDs is currently in ALPHA, which means
// that future versions of the go-client may include breaking changes.
// Please try it out and give us feedback!

// Code generated by main. DO NOT EDIT.

package v1beta1

import (
	v1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DatabaseEncryptionConfig) DeepCopyInto(out *DatabaseEncryptionConfig) {
	*out = *in
	out.KmsKeyRef = in.KmsKeyRef
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DatabaseEncryptionConfig.
func (in *DatabaseEncryptionConfig) DeepCopy() *DatabaseEncryptionConfig {
	if in == nil {
		return nil
	}
	out := new(DatabaseEncryptionConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SpannerDatabase) DeepCopyInto(out *SpannerDatabase) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SpannerDatabase.
func (in *SpannerDatabase) DeepCopy() *SpannerDatabase {
	if in == nil {
		return nil
	}
	out := new(SpannerDatabase)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SpannerDatabase) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SpannerDatabaseList) DeepCopyInto(out *SpannerDatabaseList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]SpannerDatabase, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SpannerDatabaseList.
func (in *SpannerDatabaseList) DeepCopy() *SpannerDatabaseList {
	if in == nil {
		return nil
	}
	out := new(SpannerDatabaseList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SpannerDatabaseList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SpannerDatabaseSpec) DeepCopyInto(out *SpannerDatabaseSpec) {
	*out = *in
	if in.Ddl != nil {
		in, out := &in.Ddl, &out.Ddl
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.EncryptionConfig != nil {
		in, out := &in.EncryptionConfig, &out.EncryptionConfig
		*out = new(DatabaseEncryptionConfig)
		**out = **in
	}
	out.InstanceRef = in.InstanceRef
	if in.ResourceID != nil {
		in, out := &in.ResourceID, &out.ResourceID
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SpannerDatabaseSpec.
func (in *SpannerDatabaseSpec) DeepCopy() *SpannerDatabaseSpec {
	if in == nil {
		return nil
	}
	out := new(SpannerDatabaseSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SpannerDatabaseStatus) DeepCopyInto(out *SpannerDatabaseStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1alpha1.Condition, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SpannerDatabaseStatus.
func (in *SpannerDatabaseStatus) DeepCopy() *SpannerDatabaseStatus {
	if in == nil {
		return nil
	}
	out := new(SpannerDatabaseStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SpannerInstance) DeepCopyInto(out *SpannerInstance) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SpannerInstance.
func (in *SpannerInstance) DeepCopy() *SpannerInstance {
	if in == nil {
		return nil
	}
	out := new(SpannerInstance)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SpannerInstance) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SpannerInstanceList) DeepCopyInto(out *SpannerInstanceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]SpannerInstance, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SpannerInstanceList.
func (in *SpannerInstanceList) DeepCopy() *SpannerInstanceList {
	if in == nil {
		return nil
	}
	out := new(SpannerInstanceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SpannerInstanceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SpannerInstanceSpec) DeepCopyInto(out *SpannerInstanceSpec) {
	*out = *in
	if in.NumNodes != nil {
		in, out := &in.NumNodes, &out.NumNodes
		*out = new(int)
		**out = **in
	}
	if in.ResourceID != nil {
		in, out := &in.ResourceID, &out.ResourceID
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SpannerInstanceSpec.
func (in *SpannerInstanceSpec) DeepCopy() *SpannerInstanceSpec {
	if in == nil {
		return nil
	}
	out := new(SpannerInstanceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SpannerInstanceStatus) DeepCopyInto(out *SpannerInstanceStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1alpha1.Condition, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SpannerInstanceStatus.
func (in *SpannerInstanceStatus) DeepCopy() *SpannerInstanceStatus {
	if in == nil {
		return nil
	}
	out := new(SpannerInstanceStatus)
	in.DeepCopyInto(out)
	return out
}
