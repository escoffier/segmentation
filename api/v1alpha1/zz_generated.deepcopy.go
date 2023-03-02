//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2023.

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
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/intstr"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IPBlock) DeepCopyInto(out *IPBlock) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IPBlock.
func (in *IPBlock) DeepCopy() *IPBlock {
	if in == nil {
		return nil
	}
	out := new(IPBlock)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MicrosegClusterGroup) DeepCopyInto(out *MicrosegClusterGroup) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MicrosegClusterGroup.
func (in *MicrosegClusterGroup) DeepCopy() *MicrosegClusterGroup {
	if in == nil {
		return nil
	}
	out := new(MicrosegClusterGroup)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MicrosegClusterGroup) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MicrosegClusterGroupList) DeepCopyInto(out *MicrosegClusterGroupList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]MicrosegClusterGroup, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MicrosegClusterGroupList.
func (in *MicrosegClusterGroupList) DeepCopy() *MicrosegClusterGroupList {
	if in == nil {
		return nil
	}
	out := new(MicrosegClusterGroupList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MicrosegClusterGroupList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MicrosegClusterGroupSpec) DeepCopyInto(out *MicrosegClusterGroupSpec) {
	*out = *in
	if in.PodSelector != nil {
		in, out := &in.PodSelector, &out.PodSelector
		*out = new(v1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
	if in.NamespaceSelector != nil {
		in, out := &in.NamespaceSelector, &out.NamespaceSelector
		*out = new(v1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
	if in.IPBlock != nil {
		in, out := &in.IPBlock, &out.IPBlock
		*out = new(IPBlock)
		**out = **in
	}
	if in.ServiceReference != nil {
		in, out := &in.ServiceReference, &out.ServiceReference
		*out = new(NamespacedName)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MicrosegClusterGroupSpec.
func (in *MicrosegClusterGroupSpec) DeepCopy() *MicrosegClusterGroupSpec {
	if in == nil {
		return nil
	}
	out := new(MicrosegClusterGroupSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MicrosegClusterGroupStatus) DeepCopyInto(out *MicrosegClusterGroupStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MicrosegClusterGroupStatus.
func (in *MicrosegClusterGroupStatus) DeepCopy() *MicrosegClusterGroupStatus {
	if in == nil {
		return nil
	}
	out := new(MicrosegClusterGroupStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MicrosegClusterNetworkPolicy) DeepCopyInto(out *MicrosegClusterNetworkPolicy) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MicrosegClusterNetworkPolicy.
func (in *MicrosegClusterNetworkPolicy) DeepCopy() *MicrosegClusterNetworkPolicy {
	if in == nil {
		return nil
	}
	out := new(MicrosegClusterNetworkPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MicrosegClusterNetworkPolicy) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MicrosegClusterNetworkPolicyList) DeepCopyInto(out *MicrosegClusterNetworkPolicyList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]MicrosegClusterNetworkPolicy, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MicrosegClusterNetworkPolicyList.
func (in *MicrosegClusterNetworkPolicyList) DeepCopy() *MicrosegClusterNetworkPolicyList {
	if in == nil {
		return nil
	}
	out := new(MicrosegClusterNetworkPolicyList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MicrosegClusterNetworkPolicyList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MicrosegClusterNetworkPolicySpec) DeepCopyInto(out *MicrosegClusterNetworkPolicySpec) {
	*out = *in
	if in.PodSelector != nil {
		in, out := &in.PodSelector, &out.PodSelector
		*out = new(v1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
	if in.NamespaceSelector != nil {
		in, out := &in.NamespaceSelector, &out.NamespaceSelector
		*out = new(v1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
	if in.ServiceAccount != nil {
		in, out := &in.ServiceAccount, &out.ServiceAccount
		*out = new(NamespacedName)
		**out = **in
	}
	if in.Ingress != nil {
		in, out := &in.Ingress, &out.Ingress
		*out = make([]Rule, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Egress != nil {
		in, out := &in.Egress, &out.Egress
		*out = make([]Rule, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MicrosegClusterNetworkPolicySpec.
func (in *MicrosegClusterNetworkPolicySpec) DeepCopy() *MicrosegClusterNetworkPolicySpec {
	if in == nil {
		return nil
	}
	out := new(MicrosegClusterNetworkPolicySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MicrosegClusterNetworkPolicyStatus) DeepCopyInto(out *MicrosegClusterNetworkPolicyStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MicrosegClusterNetworkPolicyStatus.
func (in *MicrosegClusterNetworkPolicyStatus) DeepCopy() *MicrosegClusterNetworkPolicyStatus {
	if in == nil {
		return nil
	}
	out := new(MicrosegClusterNetworkPolicyStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MicrosegGroup) DeepCopyInto(out *MicrosegGroup) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MicrosegGroup.
func (in *MicrosegGroup) DeepCopy() *MicrosegGroup {
	if in == nil {
		return nil
	}
	out := new(MicrosegGroup)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MicrosegGroup) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MicrosegGroupList) DeepCopyInto(out *MicrosegGroupList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]MicrosegGroup, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MicrosegGroupList.
func (in *MicrosegGroupList) DeepCopy() *MicrosegGroupList {
	if in == nil {
		return nil
	}
	out := new(MicrosegGroupList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MicrosegGroupList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MicrosegGroupSpec) DeepCopyInto(out *MicrosegGroupSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MicrosegGroupSpec.
func (in *MicrosegGroupSpec) DeepCopy() *MicrosegGroupSpec {
	if in == nil {
		return nil
	}
	out := new(MicrosegGroupSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MicrosegGroupStatus) DeepCopyInto(out *MicrosegGroupStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MicrosegGroupStatus.
func (in *MicrosegGroupStatus) DeepCopy() *MicrosegGroupStatus {
	if in == nil {
		return nil
	}
	out := new(MicrosegGroupStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MicrosegNetworkPolicy) DeepCopyInto(out *MicrosegNetworkPolicy) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MicrosegNetworkPolicy.
func (in *MicrosegNetworkPolicy) DeepCopy() *MicrosegNetworkPolicy {
	if in == nil {
		return nil
	}
	out := new(MicrosegNetworkPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MicrosegNetworkPolicy) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MicrosegNetworkPolicyList) DeepCopyInto(out *MicrosegNetworkPolicyList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]MicrosegNetworkPolicy, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MicrosegNetworkPolicyList.
func (in *MicrosegNetworkPolicyList) DeepCopy() *MicrosegNetworkPolicyList {
	if in == nil {
		return nil
	}
	out := new(MicrosegNetworkPolicyList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MicrosegNetworkPolicyList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MicrosegNetworkPolicySpec) DeepCopyInto(out *MicrosegNetworkPolicySpec) {
	*out = *in
	if in.PodSelector != nil {
		in, out := &in.PodSelector, &out.PodSelector
		*out = new(v1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
	if in.ServiceAccount != nil {
		in, out := &in.ServiceAccount, &out.ServiceAccount
		*out = new(NamespacedName)
		**out = **in
	}
	if in.Ingress != nil {
		in, out := &in.Ingress, &out.Ingress
		*out = make([]Rule, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Egress != nil {
		in, out := &in.Egress, &out.Egress
		*out = make([]Rule, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MicrosegNetworkPolicySpec.
func (in *MicrosegNetworkPolicySpec) DeepCopy() *MicrosegNetworkPolicySpec {
	if in == nil {
		return nil
	}
	out := new(MicrosegNetworkPolicySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MicrosegNetworkPolicyStatus) DeepCopyInto(out *MicrosegNetworkPolicyStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MicrosegNetworkPolicyStatus.
func (in *MicrosegNetworkPolicyStatus) DeepCopy() *MicrosegNetworkPolicyStatus {
	if in == nil {
		return nil
	}
	out := new(MicrosegNetworkPolicyStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NamespacedName) DeepCopyInto(out *NamespacedName) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NamespacedName.
func (in *NamespacedName) DeepCopy() *NamespacedName {
	if in == nil {
		return nil
	}
	out := new(NamespacedName)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetwokPolicyRule) DeepCopyInto(out *NetwokPolicyRule) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetwokPolicyRule.
func (in *NetwokPolicyRule) DeepCopy() *NetwokPolicyRule {
	if in == nil {
		return nil
	}
	out := new(NetwokPolicyRule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NetwokPolicyRule) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetwokPolicyRuleList) DeepCopyInto(out *NetwokPolicyRuleList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]NetwokPolicyRule, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetwokPolicyRuleList.
func (in *NetwokPolicyRuleList) DeepCopy() *NetwokPolicyRuleList {
	if in == nil {
		return nil
	}
	out := new(NetwokPolicyRuleList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NetwokPolicyRuleList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetwokPolicyRuleSpec) DeepCopyInto(out *NetwokPolicyRuleSpec) {
	*out = *in
	if in.Rules != nil {
		in, out := &in.Rules, &out.Rules
		*out = make([]NodeRule, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetwokPolicyRuleSpec.
func (in *NetwokPolicyRuleSpec) DeepCopy() *NetwokPolicyRuleSpec {
	if in == nil {
		return nil
	}
	out := new(NetwokPolicyRuleSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetwokPolicyRuleStatus) DeepCopyInto(out *NetwokPolicyRuleStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetwokPolicyRuleStatus.
func (in *NetwokPolicyRuleStatus) DeepCopy() *NetwokPolicyRuleStatus {
	if in == nil {
		return nil
	}
	out := new(NetwokPolicyRuleStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkPolicyPeer) DeepCopyInto(out *NetworkPolicyPeer) {
	*out = *in
	if in.IPBlock != nil {
		in, out := &in.IPBlock, &out.IPBlock
		*out = new(IPBlock)
		**out = **in
	}
	if in.PodSelector != nil {
		in, out := &in.PodSelector, &out.PodSelector
		*out = new(v1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
	if in.NamespaceSelector != nil {
		in, out := &in.NamespaceSelector, &out.NamespaceSelector
		*out = new(v1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
	if in.ServiceAccount != nil {
		in, out := &in.ServiceAccount, &out.ServiceAccount
		*out = new(NamespacedName)
		**out = **in
	}
	if in.NodeSelector != nil {
		in, out := &in.NodeSelector, &out.NodeSelector
		*out = new(v1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkPolicyPeer.
func (in *NetworkPolicyPeer) DeepCopy() *NetworkPolicyPeer {
	if in == nil {
		return nil
	}
	out := new(NetworkPolicyPeer)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkPolicyPort) DeepCopyInto(out *NetworkPolicyPort) {
	*out = *in
	if in.Protocol != nil {
		in, out := &in.Protocol, &out.Protocol
		*out = new(corev1.Protocol)
		**out = **in
	}
	if in.Port != nil {
		in, out := &in.Port, &out.Port
		*out = new(intstr.IntOrString)
		**out = **in
	}
	if in.EndPort != nil {
		in, out := &in.EndPort, &out.EndPort
		*out = new(int32)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkPolicyPort.
func (in *NetworkPolicyPort) DeepCopy() *NetworkPolicyPort {
	if in == nil {
		return nil
	}
	out := new(NetworkPolicyPort)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkPolicyRuleSlice) DeepCopyInto(out *NetworkPolicyRuleSlice) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	if in.Rules != nil {
		in, out := &in.Rules, &out.Rules
		*out = make([]NodeRule, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkPolicyRuleSlice.
func (in *NetworkPolicyRuleSlice) DeepCopy() *NetworkPolicyRuleSlice {
	if in == nil {
		return nil
	}
	out := new(NetworkPolicyRuleSlice)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NetworkPolicyRuleSlice) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkPolicyRuleSliceList) DeepCopyInto(out *NetworkPolicyRuleSliceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]NetworkPolicyRuleSlice, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkPolicyRuleSliceList.
func (in *NetworkPolicyRuleSliceList) DeepCopy() *NetworkPolicyRuleSliceList {
	if in == nil {
		return nil
	}
	out := new(NetworkPolicyRuleSliceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NetworkPolicyRuleSliceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkPolicyRuleSliceSpec) DeepCopyInto(out *NetworkPolicyRuleSliceSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkPolicyRuleSliceSpec.
func (in *NetworkPolicyRuleSliceSpec) DeepCopy() *NetworkPolicyRuleSliceSpec {
	if in == nil {
		return nil
	}
	out := new(NetworkPolicyRuleSliceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkPolicyRuleSliceStatus) DeepCopyInto(out *NetworkPolicyRuleSliceStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkPolicyRuleSliceStatus.
func (in *NetworkPolicyRuleSliceStatus) DeepCopy() *NetworkPolicyRuleSliceStatus {
	if in == nil {
		return nil
	}
	out := new(NetworkPolicyRuleSliceStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeRule) DeepCopyInto(out *NodeRule) {
	*out = *in
	if in.Ports != nil {
		in, out := &in.Ports, &out.Ports
		*out = make([]NetworkPolicyPort, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ToAddresses != nil {
		in, out := &in.ToAddresses, &out.ToAddresses
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ToIPBlock != nil {
		in, out := &in.ToIPBlock, &out.ToIPBlock
		*out = new(IPBlock)
		**out = **in
	}
	if in.FromAddress != nil {
		in, out := &in.FromAddress, &out.FromAddress
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.FromIPBlock != nil {
		in, out := &in.FromIPBlock, &out.FromIPBlock
		*out = new(IPBlock)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeRule.
func (in *NodeRule) DeepCopy() *NodeRule {
	if in == nil {
		return nil
	}
	out := new(NodeRule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Rule) DeepCopyInto(out *Rule) {
	*out = *in
	if in.Action != nil {
		in, out := &in.Action, &out.Action
		*out = new(RuleAction)
		**out = **in
	}
	if in.Ports != nil {
		in, out := &in.Ports, &out.Ports
		*out = make([]NetworkPolicyPort, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.From != nil {
		in, out := &in.From, &out.From
		*out = make([]NetworkPolicyPeer, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.To != nil {
		in, out := &in.To, &out.To
		*out = make([]NetworkPolicyPeer, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Rule.
func (in *Rule) DeepCopy() *Rule {
	if in == nil {
		return nil
	}
	out := new(Rule)
	in.DeepCopyInto(out)
	return out
}