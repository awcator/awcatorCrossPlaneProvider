/*
Copyright 2020 The Crossplane Authors.

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

// Package apis contains Kubernetes API for the AwcatorProvider provider.
package apis

import (
	"k8s.io/apimachinery/pkg/runtime"

	awcatorgroup1v1alpha1 "github.com/crossplane/provider-awcatorprovider/apis/awcatorgroup1/v1alpha1"
	awcatorgroup2passv1alpha1 "github.com/crossplane/provider-awcatorprovider/apis/awcatorgroup2pass/v1alpha1"
	awcatorproviderv1alpha1 "github.com/crossplane/provider-awcatorprovider/apis/v1alpha1"
)

/*
*
Register API in API directory
*/
func init() {
	// Register the types with the Scheme so the components can map objects to GroupVersionKinds and back
	AddToSchemes = append(AddToSchemes,
		awcatorproviderv1alpha1.SchemeBuilder.AddToScheme,
		awcatorgroup1v1alpha1.SchemeBuilder.AddToScheme,
		awcatorgroup2passv1alpha1.SchemeBuilder.AddToScheme,
	)
}

// AddToSchemes may be used to add all resources defined in the project to a Scheme
var AddToSchemes runtime.SchemeBuilder

// AddToScheme adds all Resources to the Scheme
func AddToScheme(s *runtime.Scheme) error {
	return AddToSchemes.AddToScheme(s)
}
