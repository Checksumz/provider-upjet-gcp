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

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type SecretCiphertextObservation struct {

	// Contains the result of encrypting the provided plaintext, encoded in base64.
	// Contains the result of encrypting the provided plaintext, encoded in base64.
	Ciphertext *string `json:"ciphertext,omitempty" tf:"ciphertext,omitempty"`

	// an identifier for the resource with format {{crypto_key}}/{{ciphertext}}
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type SecretCiphertextParameters struct {

	// The additional authenticated data used for integrity checks during encryption and decryption.
	// Note: This property is sensitive and will not be displayed in the plan.
	// The additional authenticated data used for integrity checks during encryption and decryption.
	// +kubebuilder:validation:Optional
	AdditionalAuthenticatedDataSecretRef *v1.SecretKeySelector `json:"additionalAuthenticatedDataSecretRef,omitempty" tf:"-"`

	// The full name of the CryptoKey that will be used to encrypt the provided plaintext.
	// Format: 'projects/{{project}}/locations/{{location}}/keyRings/{{keyRing}}/cryptoKeys/{{cryptoKey}}'
	// The full name of the CryptoKey that will be used to encrypt the provided plaintext.
	// Format: ''projects/{{project}}/locations/{{location}}/keyRings/{{keyRing}}/cryptoKeys/{{cryptoKey}}''
	// +crossplane:generate:reference:type=github.com/upbound/official-providers/provider-gcp/apis/kms/v1beta1.CryptoKey
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	CryptoKey *string `json:"cryptoKey,omitempty" tf:"crypto_key,omitempty"`

	// Reference to a CryptoKey in kms to populate cryptoKey.
	// +kubebuilder:validation:Optional
	CryptoKeyRef *v1.Reference `json:"cryptoKeyRef,omitempty" tf:"-"`

	// Selector for a CryptoKey in kms to populate cryptoKey.
	// +kubebuilder:validation:Optional
	CryptoKeySelector *v1.Selector `json:"cryptoKeySelector,omitempty" tf:"-"`

	// The plaintext to be encrypted.
	// Note: This property is sensitive and will not be displayed in the plan.
	// The plaintext to be encrypted.
	// +kubebuilder:validation:Required
	Plaintext *string `json:"plaintext" tf:"plaintext,omitempty"`
}

// SecretCiphertextSpec defines the desired state of SecretCiphertext
type SecretCiphertextSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SecretCiphertextParameters `json:"forProvider"`
}

// SecretCiphertextStatus defines the observed state of SecretCiphertext.
type SecretCiphertextStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SecretCiphertextObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// SecretCiphertext is the Schema for the SecretCiphertexts API. Encrypts secret data with Google Cloud KMS and provides access to the ciphertext.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type SecretCiphertext struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SecretCiphertextSpec   `json:"spec"`
	Status            SecretCiphertextStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SecretCiphertextList contains a list of SecretCiphertexts
type SecretCiphertextList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SecretCiphertext `json:"items"`
}

// Repository type metadata.
var (
	SecretCiphertext_Kind             = "SecretCiphertext"
	SecretCiphertext_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: SecretCiphertext_Kind}.String()
	SecretCiphertext_KindAPIVersion   = SecretCiphertext_Kind + "." + CRDGroupVersion.String()
	SecretCiphertext_GroupVersionKind = CRDGroupVersion.WithKind(SecretCiphertext_Kind)
)

func init() {
	SchemeBuilder.Register(&SecretCiphertext{}, &SecretCiphertextList{})
}
