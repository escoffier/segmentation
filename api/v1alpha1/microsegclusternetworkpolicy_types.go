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

package v1alpha1

import (
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

type RuleAction string

const (
	// RuleActionAllow describes that the traffic matching the rule must be allowed.
	RuleActionAllow RuleAction = "Allow"
	// RuleActionDrop describes that the traffic matching the rule must be dropped.
	RuleActionDrop RuleAction = "Drop"
)

type NamespacedName struct {
	Name      string `json:"name,omitempty"`
	Namespace string `json:"namespace,omitempty"`
}

type IPBlock struct {
	// CIDR is a string representing the IP Block
	// Valid examples are "192.168.1.1/24".
	CIDR string `json:"cidr"`
}

type NetworkPolicyPort struct {
	// The protocol (TCP, UDP, or SCTP) which traffic must match. If not specified, this
	// field defaults to TCP.
	// +optional
	Protocol *v1.Protocol `json:"protocol,omitempty" protobuf:"bytes,1,opt,name=protocol,casttype=k8s.io/api/core/v1.Protocol"`

	// The port on the given protocol. This can either be a numerical or named
	// port on a pod. If this field is not provided, this matches all port names and
	// numbers.
	// If present, only traffic on the specified protocol AND port will be matched.
	// +optional
	Port *intstr.IntOrString `json:"port,omitempty" protobuf:"bytes,2,opt,name=port"`

	// If set, indicates that the range of ports from port to endPort, inclusive,
	// should be allowed by the policy. This field cannot be defined if the port field
	// is not defined or if the port field is defined as a named (string) port.
	// The endPort must be equal or greater than port.
	// This feature is in Beta state and is enabled by default.
	// It can be disabled using the Feature Gate "NetworkPolicyEndPort".
	// +optional
	EndPort *int32 `json:"endPort,omitempty" protobuf:"bytes,3,opt,name=endPort"`
}

type Rule struct {
	// Action enforced on traffic which matched this rule
	Action *RuleAction `json:"action"`

	//
	Protocol string `json:"protocol,omitempty"`

	// List of destination ports for traffic.
	Ports []NetworkPolicyPort `json:"ports,omitempty"`

	// List of sources which should be able to access the pods selected for this rule.
	From []NetworkPolicyPeer `json:"from,omitempty"`

	// List of destinations for outgoing traffic of pods selected for this rule.
	To            []NetworkPolicyPeer `json:"to,omitempty"`
	Name          string              `json:"name,omitempty"`
	EnableLogging bool                `json:"enableLogging"`
}

type NetworkPolicyPeer struct {
	IPBlock           *IPBlock              `json:"ipBlock,omitempty"`
	PodSelector       *metav1.LabelSelector `json:"podSelector,omitempty"`
	NamespaceSelector *metav1.LabelSelector `json:"namespaceSelector,omitempty"`
	Group             string                `json:"group,omitempty"`
	ServiceAccount    *NamespacedName       `json:"serviceAccount,omitempty"`
	FQDN              string                `json:"fqdn,omitempty"`
	NodeSelector      *metav1.LabelSelector `json:"nodeSelector,omitempty"`
}

// MicrosegClusterNetworkPolicySpec defines the desired state of MicrosegClusterNetworkPolicy
type MicrosegClusterNetworkPolicySpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	Priority          int                   `json:"priority"`
	PodSelector       *metav1.LabelSelector `json:"podSelector,omitempty"`
	NamespaceSelector *metav1.LabelSelector `json:"namespaceSelector,omitempty"`
	Group             string                `json:"group,omitempty"`
	ServiceAccount    *NamespacedName       `json:"serviceAccount,omitempty"`
	// Ingress rules
	Ingress []Rule `json:"ingress,omitempty"`

	// Egress rules
	Egress []Rule `json:"egress,omitempty"`
}

// MicrosegClusterNetworkPolicyStatus defines the observed state of MicrosegClusterNetworkPolicy
type MicrosegClusterNetworkPolicyStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

// +genclient:nonNamespaced
// +genclient
//+kubebuilder:object:root=true
//+kubebuilder:subresource:status
//+kubebuilder:resource:scope=Cluster

// MicrosegClusterNetworkPolicy is the Schema for the microsegclusternetworkpolicies API
type MicrosegClusterNetworkPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   MicrosegClusterNetworkPolicySpec   `json:"spec,omitempty"`
	Status MicrosegClusterNetworkPolicyStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// MicrosegClusterNetworkPolicyList contains a list of MicrosegClusterNetworkPolicy
type MicrosegClusterNetworkPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MicrosegClusterNetworkPolicy `json:"items"`
}

func init() {
	SchemeBuilder.Register(&MicrosegClusterNetworkPolicy{}, &MicrosegClusterNetworkPolicyList{})
}
