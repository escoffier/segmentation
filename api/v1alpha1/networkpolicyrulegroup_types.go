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
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.
type EntityReference struct {
	Cluster   string `json:"cluster,omitempty"`
	Namespace string `json:"namespace"`
	Name      string `json:"name"`
}

type Address struct {
	IP           string           `json:"ip"`
	PodReference *EntityReference `json:"podReference,omitempty"`
}

type NodeRule struct {
	Priority    int                 `json:"priority,omitempty"`
	Protocol    string              `json:"prtocol,omitempty"`
	Direction   string              `json:"direction,omitempty"`
	Action      string              `json:"action"`
	Ports       []NetworkPolicyPort `json:"ports,omitempty"`
	ToAddresses []Address           `json:"toAddresses,omitempty"`
	ToIPBlock   *IPBlock            `json:"toIPBlock,omitempty"`
	FromAddress []Address           `json:"fromAddress,omitempty"`
	FromIPBlock *IPBlock            `json:"fromIPBlock,omitempty"`
}

// NetworkPolicyRuleGroupSpec defines the desired state of NetworkPolicyRuleGroup
type NetworkPolicyRuleGroupSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Foo is an example field of NetworkPolicyRuleGroup. Edit networkpolicyrulegroup_types.go to remove/update
	Policy   string     `json:"policy"`
	NodeName string     `json:"nodeName"`
	Rules    []NodeRule `json:"rules"`
}

// NetworkPolicyRuleGroupStatus defines the observed state of NetworkPolicyRuleGroup
type NetworkPolicyRuleGroupStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status
//+kubebuilder:resource:scope=Cluster

// NetworkPolicyRuleGroup is the Schema for the networkpolicyrulegroups API
type NetworkPolicyRuleGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   NetworkPolicyRuleGroupSpec   `json:"spec,omitempty"`
	Status NetworkPolicyRuleGroupStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// NetworkPolicyRuleGroupList contains a list of NetworkPolicyRuleGroup
type NetworkPolicyRuleGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NetworkPolicyRuleGroup `json:"items"`
}

func init() {
	SchemeBuilder.Register(&NetworkPolicyRuleGroup{}, &NetworkPolicyRuleGroupList{})
}
