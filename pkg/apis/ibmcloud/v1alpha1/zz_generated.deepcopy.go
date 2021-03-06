// +build !ignore_autogenerated

/*
 * Copyright 2019 IBM Corporation
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
// Code generated by operator-sdk. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AccessGroup) DeepCopyInto(out *AccessGroup) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccessGroup.
func (in *AccessGroup) DeepCopy() *AccessGroup {
	if in == nil {
		return nil
	}
	out := new(AccessGroup)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AccessGroup) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AccessGroupDef) DeepCopyInto(out *AccessGroupDef) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccessGroupDef.
func (in *AccessGroupDef) DeepCopy() *AccessGroupDef {
	if in == nil {
		return nil
	}
	out := new(AccessGroupDef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AccessGroupList) DeepCopyInto(out *AccessGroupList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AccessGroup, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccessGroupList.
func (in *AccessGroupList) DeepCopy() *AccessGroupList {
	if in == nil {
		return nil
	}
	out := new(AccessGroupList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AccessGroupList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AccessGroupSpec) DeepCopyInto(out *AccessGroupSpec) {
	*out = *in
	if in.UserEmails != nil {
		in, out := &in.UserEmails, &out.UserEmails
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ServiceIDs != nil {
		in, out := &in.ServiceIDs, &out.ServiceIDs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccessGroupSpec.
func (in *AccessGroupSpec) DeepCopy() *AccessGroupSpec {
	if in == nil {
		return nil
	}
	out := new(AccessGroupSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AccessGroupStatus) DeepCopyInto(out *AccessGroupStatus) {
	*out = *in
	out.ResourceStatus = in.ResourceStatus
	if in.UserEmails != nil {
		in, out := &in.UserEmails, &out.UserEmails
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ServiceIDs != nil {
		in, out := &in.ServiceIDs, &out.ServiceIDs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccessGroupStatus.
func (in *AccessGroupStatus) DeepCopy() *AccessGroupStatus {
	if in == nil {
		return nil
	}
	out := new(AccessGroupStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AccessPolicy) DeepCopyInto(out *AccessPolicy) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccessPolicy.
func (in *AccessPolicy) DeepCopy() *AccessPolicy {
	if in == nil {
		return nil
	}
	out := new(AccessPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AccessPolicy) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AccessPolicyList) DeepCopyInto(out *AccessPolicyList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AccessPolicy, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccessPolicyList.
func (in *AccessPolicyList) DeepCopy() *AccessPolicyList {
	if in == nil {
		return nil
	}
	out := new(AccessPolicyList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AccessPolicyList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AccessPolicySpec) DeepCopyInto(out *AccessPolicySpec) {
	*out = *in
	out.Subject = in.Subject
	in.Roles.DeepCopyInto(&out.Roles)
	out.Target = in.Target
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccessPolicySpec.
func (in *AccessPolicySpec) DeepCopy() *AccessPolicySpec {
	if in == nil {
		return nil
	}
	out := new(AccessPolicySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AccessPolicyStatus) DeepCopyInto(out *AccessPolicyStatus) {
	*out = *in
	out.ResourceStatus = in.ResourceStatus
	out.Subject = in.Subject
	in.Roles.DeepCopyInto(&out.Roles)
	out.Target = in.Target
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccessPolicyStatus.
func (in *AccessPolicyStatus) DeepCopy() *AccessPolicyStatus {
	if in == nil {
		return nil
	}
	out := new(AccessPolicyStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CustomRole) DeepCopyInto(out *CustomRole) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CustomRole.
func (in *CustomRole) DeepCopy() *CustomRole {
	if in == nil {
		return nil
	}
	out := new(CustomRole)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CustomRole) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CustomRoleList) DeepCopyInto(out *CustomRoleList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]CustomRole, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CustomRoleList.
func (in *CustomRoleList) DeepCopy() *CustomRoleList {
	if in == nil {
		return nil
	}
	out := new(CustomRoleList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CustomRoleList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CustomRoleSpec) DeepCopyInto(out *CustomRoleSpec) {
	*out = *in
	if in.Actions != nil {
		in, out := &in.Actions, &out.Actions
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CustomRoleSpec.
func (in *CustomRoleSpec) DeepCopy() *CustomRoleSpec {
	if in == nil {
		return nil
	}
	out := new(CustomRoleSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CustomRoleStatus) DeepCopyInto(out *CustomRoleStatus) {
	*out = *in
	out.ResourceStatus = in.ResourceStatus
	if in.Actions != nil {
		in, out := &in.Actions, &out.Actions
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CustomRoleStatus.
func (in *CustomRoleStatus) DeepCopy() *CustomRoleStatus {
	if in == nil {
		return nil
	}
	out := new(CustomRoleStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CustomRolesDef) DeepCopyInto(out *CustomRolesDef) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CustomRolesDef.
func (in *CustomRolesDef) DeepCopy() *CustomRolesDef {
	if in == nil {
		return nil
	}
	out := new(CustomRolesDef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Roles) DeepCopyInto(out *Roles) {
	*out = *in
	if in.DefinedRoles != nil {
		in, out := &in.DefinedRoles, &out.DefinedRoles
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.CustomRolesDName != nil {
		in, out := &in.CustomRolesDName, &out.CustomRolesDName
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.CustomRolesDef != nil {
		in, out := &in.CustomRolesDef, &out.CustomRolesDef
		*out = make([]CustomRolesDef, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Roles.
func (in *Roles) DeepCopy() *Roles {
	if in == nil {
		return nil
	}
	out := new(Roles)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Subject) DeepCopyInto(out *Subject) {
	*out = *in
	out.AccessGroupDef = in.AccessGroupDef
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Subject.
func (in *Subject) DeepCopy() *Subject {
	if in == nil {
		return nil
	}
	out := new(Subject)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Target) DeepCopyInto(out *Target) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Target.
func (in *Target) DeepCopy() *Target {
	if in == nil {
		return nil
	}
	out := new(Target)
	in.DeepCopyInto(out)
	return out
}
