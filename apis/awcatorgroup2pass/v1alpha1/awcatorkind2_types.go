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

// AwcatorKind2Parameters are the configurable fields of a AwcatorKind2.
type AwcatorKind2Parameters struct {
	ConfigurableField string `json:"configurableField"`
}

// AwcatorKind2Observation are the observable fields of a AwcatorKind2.
type AwcatorKind2Observation struct {
	ObservableField string `json:"observableField,omitempty"`
}

// A AwcatorKind2Spec defines the desired state of a AwcatorKind2.
type AwcatorKind2Spec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       AwcatorKind2Parameters `json:"forProvider"`
}

// A AwcatorKind2Status represents the observed state of a AwcatorKind2.
type AwcatorKind2Status struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          AwcatorKind2Observation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// A AwcatorKind2 is an example API type.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,awcatorprovider}
type AwcatorKind2 struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   AwcatorKind2Spec   `json:"spec"`
	Status AwcatorKind2Status `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AwcatorKind2List contains a list of AwcatorKind2
type AwcatorKind2List struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AwcatorKind2 `json:"items"`
}

// AwcatorKind2 type metadata.
var (
	AwcatorKind2Kind             = reflect.TypeOf(AwcatorKind2{}).Name()
	AwcatorKind2GroupKind        = schema.GroupKind{Group: Group, Kind: AwcatorKind2Kind}.String()
	AwcatorKind2KindAPIVersion   = AwcatorKind2Kind + "." + SchemeGroupVersion.String()
	AwcatorKind2GroupVersionKind = SchemeGroupVersion.WithKind(AwcatorKind2Kind)
)

func init() {
	SchemeBuilder.Register(&AwcatorKind2{}, &AwcatorKind2List{})
}
