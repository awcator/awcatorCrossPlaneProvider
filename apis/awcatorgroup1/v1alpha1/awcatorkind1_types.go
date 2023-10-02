/*
Copyright 2022 The Crossplane Authors.

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
	"reflect"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

// AwcatorKind1Parameters are the configurable fields of a AwcatorKind1.
type AwcatorKind1Parameters struct {

	// +optional
	City              *string `json:"City,omitempty"`
	UserName          string  `json:"UserName"`
	Message           string  `json:"Message"`
	ConfigurableField string  `json:"configurableField"`
}

// AwcatorKind1Observation are the observable fields of a AwcatorKind1.
type AwcatorKind1Observation struct {
	State           string `json:"state"`
	responseMsg     string `json:"state"`
	ObservableField string `json:"observableField,omitempty"`
}

// A AwcatorKind1Spec defines the desired state of a AwcatorKind1.
type AwcatorKind1Spec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       AwcatorKind1Parameters `json:"forProvider"`
}

// A AwcatorKind1Status represents the observed state of a AwcatorKind1.
type AwcatorKind1Status struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          AwcatorKind1Observation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// A AwcatorKind1 is an example API type.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,awcatorprovider}
type AwcatorKind1 struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   AwcatorKind1Spec   `json:"spec"`
	Status AwcatorKind1Status `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AwcatorKind1List contains a list of AwcatorKind1
type AwcatorKind1List struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AwcatorKind1 `json:"items"`
}

// AwcatorKind1 type metadata.
var (
	AwcatorKind1Kind             = reflect.TypeOf(AwcatorKind1{}).Name()
	AwcatorKind1GroupKind        = schema.GroupKind{Group: Group, Kind: AwcatorKind1Kind}.String()
	AwcatorKind1KindAPIVersion   = AwcatorKind1Kind + "." + SchemeGroupVersion.String()
	AwcatorKind1GroupVersionKind = SchemeGroupVersion.WithKind(AwcatorKind1Kind)
)

func init() {
	SchemeBuilder.Register(&AwcatorKind1{}, &AwcatorKind1List{})
}
