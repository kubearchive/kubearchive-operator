/*
Copyright 2026.

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
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// KubeArchiveInstallationSpec defines the desired state of KubeArchiveInstallation.
type KubeArchiveInstallationSpec struct {
	// Important: Run "make" to regenerate code after modifying this file

	// Version is the version of KubeArchive to install. See https://github.com/kubearchive/kubearchive/releases
	// for a list of available versions
	Version string `json:"version"`
}

// KubeArchiveInstallationStatus defines the observed state of KubeArchiveInstallation.
type KubeArchiveInstallationStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// KubeArchiveInstallation is the Schema for the kubearchiveinstallations API.
type KubeArchiveInstallation struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   KubeArchiveInstallationSpec   `json:"spec,omitempty"`
	Status KubeArchiveInstallationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// KubeArchiveInstallationList contains a list of KubeArchiveInstallation.
type KubeArchiveInstallationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []KubeArchiveInstallation `json:"items"`
}

func init() {
	SchemeBuilder.Register(&KubeArchiveInstallation{}, &KubeArchiveInstallationList{})
}
