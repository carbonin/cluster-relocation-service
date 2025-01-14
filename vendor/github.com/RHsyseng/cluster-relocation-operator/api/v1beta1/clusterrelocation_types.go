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

package v1beta1

import (
	configv1 "github.com/openshift/api/config/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ClusterRelocationSpec defines the desired state of ClusterRelocation
type ClusterRelocationSpec struct {
	// Important: Run "make" to regenerate code after modifying this file

	// APICertRef is a reference to a TLS secret that will be used for the API server.
	// If it is omitted, a self-signed certificate will be generated.
	APICertRef *corev1.SecretReference `json:"apiCertRef,omitempty"`

	// CatalogSources define new CatalogSources to install on the cluster.
	CatalogSources []CatalogSource `json:"catalogSources,omitempty"`

	// Domain defines the new base domain for the cluster.
	Domain string `json:"domain"`

	// ImageDigestMirrors is used to configured a mirror registry on the cluster.
	ImageDigestMirrors []configv1.ImageDigestMirrors `json:"imageDigestMirrors,omitempty"`

	// IngressCertRef is a reference to a TLS secret that will be used for the Ingress Controller.
	// If it is omitted, a self-signed certificate will be generated.
	IngressCertRef *corev1.SecretReference `json:"ingressCertRef,omitempty"`

	// PullSecretRef is a reference to new cluster-wide pull secret.
	// If defined, it will replace the secret located at openshift-config/pull-secret.
	PullSecretRef *corev1.SecretReference `json:"pullSecretRef,omitempty"`

	// RegistryCert is a new trusted CA certificate.
	// It will be added to image.config.openshift.io/cluster (additionalTrustedCA).
	RegistryCert *RegistryCert `json:"registryCert,omitempty"`

	// SSHKeys defines a list of authorized SSH keys for the 'core' user.
	// If defined, it will be appended to the existing authorized SSH key(s).
	SSHKeys []string `json:"sshKeys,omitempty"`
}

// ClusterRelocationStatus defines the observed state of ClusterRelocation
type ClusterRelocationStatus struct {
	// Important: Run "make" to regenerate code after modifying this file

	// Conditions represent the latest available observations of an object's state
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status
//+kubebuilder:resource:scope=Cluster

// ClusterRelocation is the Schema for the clusterrelocations API
type ClusterRelocation struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ClusterRelocationSpec   `json:"spec"`
	Status ClusterRelocationStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// ClusterRelocationList contains a list of ClusterRelocation
type ClusterRelocationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ClusterRelocation `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ClusterRelocation{}, &ClusterRelocationList{})
}

type CatalogSource struct {
	// Name is the name of the CatalogSource.
	Name string `json:"name"`

	// Image is an operator-registry container image to instantiate a registry-server with.
	Image string `json:"image"`
}

type RegistryCert struct {
	// RegistryHostname is the hostname of the new registry.
	RegistryHostname string `json:"registryHostname"`

	// RegistryPort is the port number that the registry is served on.
	RegistryPort *int `json:"registryPort,omitempty"`

	// Certificate is the certificate for the trusted certificate authority associated with the registry.
	Certificate string `json:"certificate"`
}

const (
	ConditionTypeReady      string = "Ready"
	ConditionTypeReconciled string = "Reconciled"
)

const (
	PullSecretName       string = "pull-secret"
	BackupPullSecretName string = "backup-pull-secret"
	ConfigNamespace      string = "openshift-config"
	IngressNamespace     string = "openshift-ingress"
)

const (
	// ValidationSucceededReason represents the fact that the validation of
	// the resource has succeeded.
	ValidationSucceededReason string = "ValidationSucceeded"

	// ValidationFailedReason represents the fact that the validation of
	// the resource has failed.
	ValidationFailedReason string = "ValidationFailed"

	// ReconciliationSucceededReason represents the fact that the validation of
	// the resource has succeeded.
	ReconciliationSucceededReason string = "ReconciliationSucceeded"

	APIReconciliationFailedReason        string = "APIReconciliationFailed"
	IngressReconciliationFailedReason    string = "IngressReconciliationFailed"
	PullSecretReconciliationFailedReason string = "PullSecretReconciliationFailed"
	SSHReconciliationFailedReason        string = "SSHReconciliationFailed"
	RegistryReconciliationFailedReason   string = "RegistryReconciliationFailed"
	MirrorReconciliationFailedReason     string = "MirrorReconciliationFailed"
	CatalogReconciliationFailedReason    string = "CatalogReconciliationFailed"
	DNSReconciliationFailedReason        string = "DNSReconciliationFailed"
	InProgressReconciliationFailedReason string = "ReconcileInProgress"
)
