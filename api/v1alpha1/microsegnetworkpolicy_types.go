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

// MicrosegNetworkPolicySpec defines the desired state of MicrosegNetworkPolicy
type MicrosegNetworkPolicySpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	Priority       int                   `json:"priority"`
	PodSelector    *metav1.LabelSelector `json:"podSelector,omitempty"`
	Group          string                `json:"group,omitempty"`
	ServiceAccount *NamespacedName       `json:"serviceAccount,omitempty"`
	// Ingress rules
	Ingress []Rule `json:"ingress,omitempty"`

	// Egress rules
	Egress []Rule `json:"egress,omitempty"`
}

// MicrosegNetworkPolicyStatus defines the observed state of MicrosegNetworkPolicy
type MicrosegNetworkPolicyStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// MicrosegNetworkPolicy is the Schema for the microsegnetworkpolicies API
type MicrosegNetworkPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   MicrosegNetworkPolicySpec   `json:"spec,omitempty"`
	Status MicrosegNetworkPolicyStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// MicrosegNetworkPolicyList contains a list of MicrosegNetworkPolicy
type MicrosegNetworkPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MicrosegNetworkPolicy `json:"items"`
}

func init() {
	SchemeBuilder.Register(&MicrosegNetworkPolicy{}, &MicrosegNetworkPolicyList{})
}
