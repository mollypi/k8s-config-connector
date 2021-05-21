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
func (in *AccessContextManagerAccessLevel) DeepCopyInto(out *AccessContextManagerAccessLevel) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccessContextManagerAccessLevel.
func (in *AccessContextManagerAccessLevel) DeepCopy() *AccessContextManagerAccessLevel {
	if in == nil {
		return nil
	}
	out := new(AccessContextManagerAccessLevel)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AccessContextManagerAccessLevel) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AccessContextManagerAccessLevelList) DeepCopyInto(out *AccessContextManagerAccessLevelList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AccessContextManagerAccessLevel, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccessContextManagerAccessLevelList.
func (in *AccessContextManagerAccessLevelList) DeepCopy() *AccessContextManagerAccessLevelList {
	if in == nil {
		return nil
	}
	out := new(AccessContextManagerAccessLevelList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AccessContextManagerAccessLevelList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AccessContextManagerAccessLevelSpec) DeepCopyInto(out *AccessContextManagerAccessLevelSpec) {
	*out = *in
	out.AccessPolicyRef = in.AccessPolicyRef
	if in.Basic != nil {
		in, out := &in.Basic, &out.Basic
		*out = new(AccesslevelBasic)
		(*in).DeepCopyInto(*out)
	}
	if in.Custom != nil {
		in, out := &in.Custom, &out.Custom
		*out = new(AccesslevelCustom)
		(*in).DeepCopyInto(*out)
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.ResourceID != nil {
		in, out := &in.ResourceID, &out.ResourceID
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccessContextManagerAccessLevelSpec.
func (in *AccessContextManagerAccessLevelSpec) DeepCopy() *AccessContextManagerAccessLevelSpec {
	if in == nil {
		return nil
	}
	out := new(AccessContextManagerAccessLevelSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AccessContextManagerAccessLevelStatus) DeepCopyInto(out *AccessContextManagerAccessLevelStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1alpha1.Condition, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccessContextManagerAccessLevelStatus.
func (in *AccessContextManagerAccessLevelStatus) DeepCopy() *AccessContextManagerAccessLevelStatus {
	if in == nil {
		return nil
	}
	out := new(AccessContextManagerAccessLevelStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AccessContextManagerAccessPolicy) DeepCopyInto(out *AccessContextManagerAccessPolicy) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccessContextManagerAccessPolicy.
func (in *AccessContextManagerAccessPolicy) DeepCopy() *AccessContextManagerAccessPolicy {
	if in == nil {
		return nil
	}
	out := new(AccessContextManagerAccessPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AccessContextManagerAccessPolicy) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AccessContextManagerAccessPolicyList) DeepCopyInto(out *AccessContextManagerAccessPolicyList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AccessContextManagerAccessPolicy, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccessContextManagerAccessPolicyList.
func (in *AccessContextManagerAccessPolicyList) DeepCopy() *AccessContextManagerAccessPolicyList {
	if in == nil {
		return nil
	}
	out := new(AccessContextManagerAccessPolicyList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AccessContextManagerAccessPolicyList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AccessContextManagerAccessPolicySpec) DeepCopyInto(out *AccessContextManagerAccessPolicySpec) {
	*out = *in
	if in.ResourceID != nil {
		in, out := &in.ResourceID, &out.ResourceID
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccessContextManagerAccessPolicySpec.
func (in *AccessContextManagerAccessPolicySpec) DeepCopy() *AccessContextManagerAccessPolicySpec {
	if in == nil {
		return nil
	}
	out := new(AccessContextManagerAccessPolicySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AccessContextManagerAccessPolicyStatus) DeepCopyInto(out *AccessContextManagerAccessPolicyStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1alpha1.Condition, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccessContextManagerAccessPolicyStatus.
func (in *AccessContextManagerAccessPolicyStatus) DeepCopy() *AccessContextManagerAccessPolicyStatus {
	if in == nil {
		return nil
	}
	out := new(AccessContextManagerAccessPolicyStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AccessContextManagerServicePerimeter) DeepCopyInto(out *AccessContextManagerServicePerimeter) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccessContextManagerServicePerimeter.
func (in *AccessContextManagerServicePerimeter) DeepCopy() *AccessContextManagerServicePerimeter {
	if in == nil {
		return nil
	}
	out := new(AccessContextManagerServicePerimeter)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AccessContextManagerServicePerimeter) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AccessContextManagerServicePerimeterList) DeepCopyInto(out *AccessContextManagerServicePerimeterList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AccessContextManagerServicePerimeter, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccessContextManagerServicePerimeterList.
func (in *AccessContextManagerServicePerimeterList) DeepCopy() *AccessContextManagerServicePerimeterList {
	if in == nil {
		return nil
	}
	out := new(AccessContextManagerServicePerimeterList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AccessContextManagerServicePerimeterList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AccessContextManagerServicePerimeterSpec) DeepCopyInto(out *AccessContextManagerServicePerimeterSpec) {
	*out = *in
	out.AccessPolicyRef = in.AccessPolicyRef
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.PerimeterType != nil {
		in, out := &in.PerimeterType, &out.PerimeterType
		*out = new(string)
		**out = **in
	}
	if in.ResourceID != nil {
		in, out := &in.ResourceID, &out.ResourceID
		*out = new(string)
		**out = **in
	}
	if in.Spec != nil {
		in, out := &in.Spec, &out.Spec
		*out = new(ServiceperimeterSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(ServiceperimeterStatus)
		(*in).DeepCopyInto(*out)
	}
	if in.UseExplicitDryRunSpec != nil {
		in, out := &in.UseExplicitDryRunSpec, &out.UseExplicitDryRunSpec
		*out = new(bool)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccessContextManagerServicePerimeterSpec.
func (in *AccessContextManagerServicePerimeterSpec) DeepCopy() *AccessContextManagerServicePerimeterSpec {
	if in == nil {
		return nil
	}
	out := new(AccessContextManagerServicePerimeterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AccessContextManagerServicePerimeterStatus) DeepCopyInto(out *AccessContextManagerServicePerimeterStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1alpha1.Condition, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccessContextManagerServicePerimeterStatus.
func (in *AccessContextManagerServicePerimeterStatus) DeepCopy() *AccessContextManagerServicePerimeterStatus {
	if in == nil {
		return nil
	}
	out := new(AccessContextManagerServicePerimeterStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AccesslevelBasic) DeepCopyInto(out *AccesslevelBasic) {
	*out = *in
	if in.CombiningFunction != nil {
		in, out := &in.CombiningFunction, &out.CombiningFunction
		*out = new(string)
		**out = **in
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]AccesslevelConditions, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccesslevelBasic.
func (in *AccesslevelBasic) DeepCopy() *AccesslevelBasic {
	if in == nil {
		return nil
	}
	out := new(AccesslevelBasic)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AccesslevelConditions) DeepCopyInto(out *AccesslevelConditions) {
	*out = *in
	if in.DevicePolicy != nil {
		in, out := &in.DevicePolicy, &out.DevicePolicy
		*out = new(AccesslevelDevicePolicy)
		(*in).DeepCopyInto(*out)
	}
	if in.IpSubnetworks != nil {
		in, out := &in.IpSubnetworks, &out.IpSubnetworks
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Members != nil {
		in, out := &in.Members, &out.Members
		*out = make([]AccesslevelMembers, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Negate != nil {
		in, out := &in.Negate, &out.Negate
		*out = new(bool)
		**out = **in
	}
	if in.Regions != nil {
		in, out := &in.Regions, &out.Regions
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.RequiredAccessLevels != nil {
		in, out := &in.RequiredAccessLevels, &out.RequiredAccessLevels
		*out = make([]v1alpha1.ResourceRef, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccesslevelConditions.
func (in *AccesslevelConditions) DeepCopy() *AccesslevelConditions {
	if in == nil {
		return nil
	}
	out := new(AccesslevelConditions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AccesslevelCustom) DeepCopyInto(out *AccesslevelCustom) {
	*out = *in
	in.Expr.DeepCopyInto(&out.Expr)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccesslevelCustom.
func (in *AccesslevelCustom) DeepCopy() *AccesslevelCustom {
	if in == nil {
		return nil
	}
	out := new(AccesslevelCustom)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AccesslevelDevicePolicy) DeepCopyInto(out *AccesslevelDevicePolicy) {
	*out = *in
	if in.AllowedDeviceManagementLevels != nil {
		in, out := &in.AllowedDeviceManagementLevels, &out.AllowedDeviceManagementLevels
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.AllowedEncryptionStatuses != nil {
		in, out := &in.AllowedEncryptionStatuses, &out.AllowedEncryptionStatuses
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.OsConstraints != nil {
		in, out := &in.OsConstraints, &out.OsConstraints
		*out = make([]AccesslevelOsConstraints, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.RequireAdminApproval != nil {
		in, out := &in.RequireAdminApproval, &out.RequireAdminApproval
		*out = new(bool)
		**out = **in
	}
	if in.RequireCorpOwned != nil {
		in, out := &in.RequireCorpOwned, &out.RequireCorpOwned
		*out = new(bool)
		**out = **in
	}
	if in.RequireScreenLock != nil {
		in, out := &in.RequireScreenLock, &out.RequireScreenLock
		*out = new(bool)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccesslevelDevicePolicy.
func (in *AccesslevelDevicePolicy) DeepCopy() *AccesslevelDevicePolicy {
	if in == nil {
		return nil
	}
	out := new(AccesslevelDevicePolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AccesslevelExpr) DeepCopyInto(out *AccesslevelExpr) {
	*out = *in
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.Location != nil {
		in, out := &in.Location, &out.Location
		*out = new(string)
		**out = **in
	}
	if in.Title != nil {
		in, out := &in.Title, &out.Title
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccesslevelExpr.
func (in *AccesslevelExpr) DeepCopy() *AccesslevelExpr {
	if in == nil {
		return nil
	}
	out := new(AccesslevelExpr)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AccesslevelMembers) DeepCopyInto(out *AccesslevelMembers) {
	*out = *in
	if in.ServiceAccountRef != nil {
		in, out := &in.ServiceAccountRef, &out.ServiceAccountRef
		*out = new(v1alpha1.ResourceRef)
		**out = **in
	}
	if in.User != nil {
		in, out := &in.User, &out.User
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccesslevelMembers.
func (in *AccesslevelMembers) DeepCopy() *AccesslevelMembers {
	if in == nil {
		return nil
	}
	out := new(AccesslevelMembers)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AccesslevelOsConstraints) DeepCopyInto(out *AccesslevelOsConstraints) {
	*out = *in
	if in.MinimumVersion != nil {
		in, out := &in.MinimumVersion, &out.MinimumVersion
		*out = new(string)
		**out = **in
	}
	if in.RequireVerifiedChromeOs != nil {
		in, out := &in.RequireVerifiedChromeOs, &out.RequireVerifiedChromeOs
		*out = new(bool)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccesslevelOsConstraints.
func (in *AccesslevelOsConstraints) DeepCopy() *AccesslevelOsConstraints {
	if in == nil {
		return nil
	}
	out := new(AccesslevelOsConstraints)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceperimeterResources) DeepCopyInto(out *ServiceperimeterResources) {
	*out = *in
	if in.ProjectRef != nil {
		in, out := &in.ProjectRef, &out.ProjectRef
		*out = new(v1alpha1.ResourceRef)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceperimeterResources.
func (in *ServiceperimeterResources) DeepCopy() *ServiceperimeterResources {
	if in == nil {
		return nil
	}
	out := new(ServiceperimeterResources)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceperimeterSpec) DeepCopyInto(out *ServiceperimeterSpec) {
	*out = *in
	if in.AccessLevels != nil {
		in, out := &in.AccessLevels, &out.AccessLevels
		*out = make([]v1alpha1.ResourceRef, len(*in))
		copy(*out, *in)
	}
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = make([]ServiceperimeterResources, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.RestrictedServices != nil {
		in, out := &in.RestrictedServices, &out.RestrictedServices
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.VpcAccessibleServices != nil {
		in, out := &in.VpcAccessibleServices, &out.VpcAccessibleServices
		*out = new(ServiceperimeterVpcAccessibleServices)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceperimeterSpec.
func (in *ServiceperimeterSpec) DeepCopy() *ServiceperimeterSpec {
	if in == nil {
		return nil
	}
	out := new(ServiceperimeterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceperimeterStatus) DeepCopyInto(out *ServiceperimeterStatus) {
	*out = *in
	if in.AccessLevels != nil {
		in, out := &in.AccessLevels, &out.AccessLevels
		*out = make([]v1alpha1.ResourceRef, len(*in))
		copy(*out, *in)
	}
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = make([]ServiceperimeterResources, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.RestrictedServices != nil {
		in, out := &in.RestrictedServices, &out.RestrictedServices
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.VpcAccessibleServices != nil {
		in, out := &in.VpcAccessibleServices, &out.VpcAccessibleServices
		*out = new(ServiceperimeterVpcAccessibleServices)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceperimeterStatus.
func (in *ServiceperimeterStatus) DeepCopy() *ServiceperimeterStatus {
	if in == nil {
		return nil
	}
	out := new(ServiceperimeterStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceperimeterVpcAccessibleServices) DeepCopyInto(out *ServiceperimeterVpcAccessibleServices) {
	*out = *in
	if in.AllowedServices != nil {
		in, out := &in.AllowedServices, &out.AllowedServices
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.EnableRestriction != nil {
		in, out := &in.EnableRestriction, &out.EnableRestriction
		*out = new(bool)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceperimeterVpcAccessibleServices.
func (in *ServiceperimeterVpcAccessibleServices) DeepCopy() *ServiceperimeterVpcAccessibleServices {
	if in == nil {
		return nil
	}
	out := new(ServiceperimeterVpcAccessibleServices)
	in.DeepCopyInto(out)
	return out
}
