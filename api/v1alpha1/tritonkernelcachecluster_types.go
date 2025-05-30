/*
Copyright 2025.

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

// TritonKernelCacheClusterSpec defines the desired state of TritonKernelCacheCluster
type TritonKernelCacheClusterSpec struct {
	CacheImage        string `json:"cacheImage"`
	ValidateSignature bool   `json:"validateSignature"`
}

// TritonKernelCacheClusterStatus defines the observed state of TritonKernelCacheCluster
type TritonKernelCacheClusterStatus struct {
	Digest     string             `json:"digest,omitempty"`
	LastSynced metav1.Time        `json:"lastSynced,omitempty"`
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// TritonKernelCacheCluster is the Schema for the tritonkernelcacheclusters API
type TritonKernelCacheCluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   TritonKernelCacheClusterSpec   `json:"spec,omitempty"`
	Status TritonKernelCacheClusterStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// TritonKernelCacheClusterList contains a list of TritonKernelCacheCluster
type TritonKernelCacheClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TritonKernelCacheCluster `json:"items"`
}

func init() {
	SchemeBuilder.Register(&TritonKernelCacheCluster{}, &TritonKernelCacheClusterList{})
}
