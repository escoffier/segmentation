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

// MicrosegClusterGroupSpec defines the desired state of MicrosegClusterGroup
type MicrosegClusterGroupSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	PodSelector       *metav1.LabelSelector `json:"podSelector,omitempty"`
	NamespaceSelector *metav1.LabelSelector `json:"namespaceSelector,omitempty"`
	IPBlock           *IPBlock              `json:"ipBlock,omitempty"`
	ServiceReference  *NamespacedName       `json:"serviceReference,omitempty"`
}

// MicrosegClusterGroupStatus defines the observed state of MicrosegClusterGroup
type MicrosegClusterGroupStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

//+genclient:nonNamespaced
//+genclient
//+kubebuilder:object:root=true
//+kubebuilder:subresource:status
//+kubebuilder:resource:scope=Cluster

// MicrosegClusterGroup is the Schema for the microsegclustergroups API
type MicrosegClusterGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   MicrosegClusterGroupSpec   `json:"spec,omitempty"`
	Status MicrosegClusterGroupStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// MicrosegClusterGroupList contains a list of MicrosegClusterGroup
type MicrosegClusterGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MicrosegClusterGroup `json:"items"`
}

func init() {
	SchemeBuilder.Register(&MicrosegClusterGroup{}, &MicrosegClusterGroupList{})
}
