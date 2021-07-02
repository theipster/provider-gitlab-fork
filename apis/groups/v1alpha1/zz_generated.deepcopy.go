// +build !ignore_autogenerated

/*
Copyright 2021 The Crossplane Authors.

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
	"github.com/crossplane/crossplane-runtime/apis/common/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CustomAttribute) DeepCopyInto(out *CustomAttribute) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CustomAttribute.
func (in *CustomAttribute) DeepCopy() *CustomAttribute {
	if in == nil {
		return nil
	}
	out := new(CustomAttribute)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Group) DeepCopyInto(out *Group) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Group.
func (in *Group) DeepCopy() *Group {
	if in == nil {
		return nil
	}
	out := new(Group)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Group) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GroupList) DeepCopyInto(out *GroupList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Group, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GroupList.
func (in *GroupList) DeepCopy() *GroupList {
	if in == nil {
		return nil
	}
	out := new(GroupList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *GroupList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GroupMember) DeepCopyInto(out *GroupMember) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GroupMember.
func (in *GroupMember) DeepCopy() *GroupMember {
	if in == nil {
		return nil
	}
	out := new(GroupMember)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *GroupMember) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GroupMemberList) DeepCopyInto(out *GroupMemberList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]GroupMember, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GroupMemberList.
func (in *GroupMemberList) DeepCopy() *GroupMemberList {
	if in == nil {
		return nil
	}
	out := new(GroupMemberList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *GroupMemberList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GroupMemberObservation) DeepCopyInto(out *GroupMemberObservation) {
	*out = *in
	if in.GroupSAMLIdentity != nil {
		in, out := &in.GroupSAMLIdentity, &out.GroupSAMLIdentity
		*out = new(GroupMemberSAMLIdentity)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GroupMemberObservation.
func (in *GroupMemberObservation) DeepCopy() *GroupMemberObservation {
	if in == nil {
		return nil
	}
	out := new(GroupMemberObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GroupMemberParameters) DeepCopyInto(out *GroupMemberParameters) {
	*out = *in
	if in.ExpiresAt != nil {
		in, out := &in.ExpiresAt, &out.ExpiresAt
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GroupMemberParameters.
func (in *GroupMemberParameters) DeepCopy() *GroupMemberParameters {
	if in == nil {
		return nil
	}
	out := new(GroupMemberParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GroupMemberSAMLIdentity) DeepCopyInto(out *GroupMemberSAMLIdentity) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GroupMemberSAMLIdentity.
func (in *GroupMemberSAMLIdentity) DeepCopy() *GroupMemberSAMLIdentity {
	if in == nil {
		return nil
	}
	out := new(GroupMemberSAMLIdentity)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GroupMemberSpec) DeepCopyInto(out *GroupMemberSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GroupMemberSpec.
func (in *GroupMemberSpec) DeepCopy() *GroupMemberSpec {
	if in == nil {
		return nil
	}
	out := new(GroupMemberSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GroupMemberStatus) DeepCopyInto(out *GroupMemberStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GroupMemberStatus.
func (in *GroupMemberStatus) DeepCopy() *GroupMemberStatus {
	if in == nil {
		return nil
	}
	out := new(GroupMemberStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GroupObservation) DeepCopyInto(out *GroupObservation) {
	*out = *in
	if in.Statistics != nil {
		in, out := &in.Statistics, &out.Statistics
		*out = new(StorageStatistics)
		**out = **in
	}
	if in.CustomAttributes != nil {
		in, out := &in.CustomAttributes, &out.CustomAttributes
		*out = make([]CustomAttribute, len(*in))
		copy(*out, *in)
	}
	if in.SharedWithGroups != nil {
		in, out := &in.SharedWithGroups, &out.SharedWithGroups
		*out = make([]SharedWithGroups, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.LDAPGroupLinks != nil {
		in, out := &in.LDAPGroupLinks, &out.LDAPGroupLinks
		*out = make([]LDAPGroupLink, len(*in))
		copy(*out, *in)
	}
	if in.MarkedForDeletionOn != nil {
		in, out := &in.MarkedForDeletionOn, &out.MarkedForDeletionOn
		*out = (*in).DeepCopy()
	}
	if in.CreatedAt != nil {
		in, out := &in.CreatedAt, &out.CreatedAt
		*out = (*in).DeepCopy()
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GroupObservation.
func (in *GroupObservation) DeepCopy() *GroupObservation {
	if in == nil {
		return nil
	}
	out := new(GroupObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GroupParameters) DeepCopyInto(out *GroupParameters) {
	*out = *in
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.MembershipLock != nil {
		in, out := &in.MembershipLock, &out.MembershipLock
		*out = new(bool)
		**out = **in
	}
	if in.Visibility != nil {
		in, out := &in.Visibility, &out.Visibility
		*out = new(VisibilityValue)
		**out = **in
	}
	if in.ShareWithGroupLock != nil {
		in, out := &in.ShareWithGroupLock, &out.ShareWithGroupLock
		*out = new(bool)
		**out = **in
	}
	if in.RequireTwoFactorAuth != nil {
		in, out := &in.RequireTwoFactorAuth, &out.RequireTwoFactorAuth
		*out = new(bool)
		**out = **in
	}
	if in.TwoFactorGracePeriod != nil {
		in, out := &in.TwoFactorGracePeriod, &out.TwoFactorGracePeriod
		*out = new(int)
		**out = **in
	}
	if in.ProjectCreationLevel != nil {
		in, out := &in.ProjectCreationLevel, &out.ProjectCreationLevel
		*out = new(ProjectCreationLevelValue)
		**out = **in
	}
	if in.AutoDevopsEnabled != nil {
		in, out := &in.AutoDevopsEnabled, &out.AutoDevopsEnabled
		*out = new(bool)
		**out = **in
	}
	if in.SubGroupCreationLevel != nil {
		in, out := &in.SubGroupCreationLevel, &out.SubGroupCreationLevel
		*out = new(SubGroupCreationLevelValue)
		**out = **in
	}
	if in.EmailsDisabled != nil {
		in, out := &in.EmailsDisabled, &out.EmailsDisabled
		*out = new(bool)
		**out = **in
	}
	if in.MentionsDisabled != nil {
		in, out := &in.MentionsDisabled, &out.MentionsDisabled
		*out = new(bool)
		**out = **in
	}
	if in.LFSEnabled != nil {
		in, out := &in.LFSEnabled, &out.LFSEnabled
		*out = new(bool)
		**out = **in
	}
	if in.RequestAccessEnabled != nil {
		in, out := &in.RequestAccessEnabled, &out.RequestAccessEnabled
		*out = new(bool)
		**out = **in
	}
	if in.ParentID != nil {
		in, out := &in.ParentID, &out.ParentID
		*out = new(int)
		**out = **in
	}
	if in.ParentIDRef != nil {
		in, out := &in.ParentIDRef, &out.ParentIDRef
		*out = new(v1.Reference)
		**out = **in
	}
	if in.ParentIDSelector != nil {
		in, out := &in.ParentIDSelector, &out.ParentIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.SharedRunnersMinutesLimit != nil {
		in, out := &in.SharedRunnersMinutesLimit, &out.SharedRunnersMinutesLimit
		*out = new(int)
		**out = **in
	}
	if in.ExtraSharedRunnersMinutesLimit != nil {
		in, out := &in.ExtraSharedRunnersMinutesLimit, &out.ExtraSharedRunnersMinutesLimit
		*out = new(int)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GroupParameters.
func (in *GroupParameters) DeepCopy() *GroupParameters {
	if in == nil {
		return nil
	}
	out := new(GroupParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GroupSpec) DeepCopyInto(out *GroupSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GroupSpec.
func (in *GroupSpec) DeepCopy() *GroupSpec {
	if in == nil {
		return nil
	}
	out := new(GroupSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GroupStatus) DeepCopyInto(out *GroupStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GroupStatus.
func (in *GroupStatus) DeepCopy() *GroupStatus {
	if in == nil {
		return nil
	}
	out := new(GroupStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LDAPGroupLink) DeepCopyInto(out *LDAPGroupLink) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LDAPGroupLink.
func (in *LDAPGroupLink) DeepCopy() *LDAPGroupLink {
	if in == nil {
		return nil
	}
	out := new(LDAPGroupLink)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SharedWithGroups) DeepCopyInto(out *SharedWithGroups) {
	*out = *in
	if in.ExpiresAt != nil {
		in, out := &in.ExpiresAt, &out.ExpiresAt
		*out = (*in).DeepCopy()
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SharedWithGroups.
func (in *SharedWithGroups) DeepCopy() *SharedWithGroups {
	if in == nil {
		return nil
	}
	out := new(SharedWithGroups)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StorageStatistics) DeepCopyInto(out *StorageStatistics) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StorageStatistics.
func (in *StorageStatistics) DeepCopy() *StorageStatistics {
	if in == nil {
		return nil
	}
	out := new(StorageStatistics)
	in.DeepCopyInto(out)
	return out
}
