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

package v1

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// IndexerClusterSpec defines the desired state of IndexerCluster
type IndexerClusterSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Foo is an example field of IndexerCluster. Edit indexercluster_types.go to remove/update
	Spec `json:",inline"`
	// ClusterManagerRef refers to a Splunk Enterprise indexer cluster managed by the operator within Kubernetes
	ClusterManagerRef corev1.ObjectReference `json:"clusterManagerRef"`
}

// IndexerClusterStatus defines the observed state of IndexerCluster
type IndexerClusterStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	Phase Phase  `json:"phase"`
	Image string `image:"image"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// IndexerCluster is the Schema for the indexerclusters API
type IndexerCluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   IndexerClusterSpec   `json:"spec,omitempty"`
	Status IndexerClusterStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// IndexerClusterList contains a list of IndexerCluster
type IndexerClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []IndexerCluster `json:"items"`
}

func init() {
	SchemeBuilder.Register(&IndexerCluster{}, &IndexerClusterList{})
}
