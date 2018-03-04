/*
Copyright 2017 The Kubernetes Authors.

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

// This file was autogenerated by apiregister-gen. Do not edit it manually!

package v1alpha1

import (
	"github.com/kubernetes-incubator/apiserver-builder/pkg/builders"
	"github.com/marun/fnord/pkg/apis/federation"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

var (
	federationFederatedClusterStorage = builders.NewApiResource( // Resource status endpoint
		federation.InternalFederatedCluster,
		FederatedClusterSchemeFns{},
		func() runtime.Object { return &FederatedCluster{} },     // Register versioned resource
		func() runtime.Object { return &FederatedClusterList{} }, // Register versioned resource list
		&FederatedClusterStrategy{builders.StorageStrategySingleton},
	)
	federationFederatedReplicaSetStorage = builders.NewApiResource( // Resource status endpoint
		federation.InternalFederatedReplicaSet,
		FederatedReplicaSetSchemeFns{},
		func() runtime.Object { return &FederatedReplicaSet{} },     // Register versioned resource
		func() runtime.Object { return &FederatedReplicaSetList{} }, // Register versioned resource list
		&FederatedReplicaSetStrategy{builders.StorageStrategySingleton},
	)
	federationFederatedReplicaSetOverrideStorage = builders.NewApiResource( // Resource status endpoint
		federation.InternalFederatedReplicaSetOverride,
		FederatedReplicaSetOverrideSchemeFns{},
		func() runtime.Object { return &FederatedReplicaSetOverride{} },     // Register versioned resource
		func() runtime.Object { return &FederatedReplicaSetOverrideList{} }, // Register versioned resource list
		&FederatedReplicaSetOverrideStrategy{builders.StorageStrategySingleton},
	)
	federationFederatedSecretStorage = builders.NewApiResource( // Resource status endpoint
		federation.InternalFederatedSecret,
		FederatedSecretSchemeFns{},
		func() runtime.Object { return &FederatedSecret{} },     // Register versioned resource
		func() runtime.Object { return &FederatedSecretList{} }, // Register versioned resource list
		&FederatedSecretStrategy{builders.StorageStrategySingleton},
	)
	federationFederatedSecretOverrideStorage = builders.NewApiResource( // Resource status endpoint
		federation.InternalFederatedSecretOverride,
		FederatedSecretOverrideSchemeFns{},
		func() runtime.Object { return &FederatedSecretOverride{} },     // Register versioned resource
		func() runtime.Object { return &FederatedSecretOverrideList{} }, // Register versioned resource list
		&FederatedSecretOverrideStrategy{builders.StorageStrategySingleton},
	)
	federationFederatedSecretPlacementStorage = builders.NewApiResource( // Resource status endpoint
		federation.InternalFederatedSecretPlacement,
		FederatedSecretPlacementSchemeFns{},
		func() runtime.Object { return &FederatedSecretPlacement{} },     // Register versioned resource
		func() runtime.Object { return &FederatedSecretPlacementList{} }, // Register versioned resource list
		&FederatedSecretPlacementStrategy{builders.StorageStrategySingleton},
	)
	federationFederationPlacementStorage = builders.NewApiResource( // Resource status endpoint
		federation.InternalFederationPlacement,
		FederationPlacementSchemeFns{},
		func() runtime.Object { return &FederationPlacement{} },     // Register versioned resource
		func() runtime.Object { return &FederationPlacementList{} }, // Register versioned resource list
		&FederationPlacementStrategy{builders.StorageStrategySingleton},
	)
	ApiVersion = builders.NewApiVersion("federation.k8s.io", "v1alpha1").WithResources(
		federationFederatedClusterStorage,
		builders.NewApiResource( // Resource status endpoint
			federation.InternalFederatedClusterStatus,
			FederatedClusterSchemeFns{},
			func() runtime.Object { return &FederatedCluster{} },     // Register versioned resource
			func() runtime.Object { return &FederatedClusterList{} }, // Register versioned resource list
			&FederatedClusterStatusStrategy{builders.StatusStorageStrategySingleton},
		), federationFederatedReplicaSetStorage,
		builders.NewApiResource( // Resource status endpoint
			federation.InternalFederatedReplicaSetStatus,
			FederatedReplicaSetSchemeFns{},
			func() runtime.Object { return &FederatedReplicaSet{} },     // Register versioned resource
			func() runtime.Object { return &FederatedReplicaSetList{} }, // Register versioned resource list
			&FederatedReplicaSetStatusStrategy{builders.StatusStorageStrategySingleton},
		), federationFederatedReplicaSetOverrideStorage,
		builders.NewApiResource( // Resource status endpoint
			federation.InternalFederatedReplicaSetOverrideStatus,
			FederatedReplicaSetOverrideSchemeFns{},
			func() runtime.Object { return &FederatedReplicaSetOverride{} },     // Register versioned resource
			func() runtime.Object { return &FederatedReplicaSetOverrideList{} }, // Register versioned resource list
			&FederatedReplicaSetOverrideStatusStrategy{builders.StatusStorageStrategySingleton},
		), federationFederatedSecretStorage,
		builders.NewApiResource( // Resource status endpoint
			federation.InternalFederatedSecretStatus,
			FederatedSecretSchemeFns{},
			func() runtime.Object { return &FederatedSecret{} },     // Register versioned resource
			func() runtime.Object { return &FederatedSecretList{} }, // Register versioned resource list
			&FederatedSecretStatusStrategy{builders.StatusStorageStrategySingleton},
		), federationFederatedSecretOverrideStorage,
		builders.NewApiResource( // Resource status endpoint
			federation.InternalFederatedSecretOverrideStatus,
			FederatedSecretOverrideSchemeFns{},
			func() runtime.Object { return &FederatedSecretOverride{} },     // Register versioned resource
			func() runtime.Object { return &FederatedSecretOverrideList{} }, // Register versioned resource list
			&FederatedSecretOverrideStatusStrategy{builders.StatusStorageStrategySingleton},
		), federationFederatedSecretPlacementStorage,
		builders.NewApiResource( // Resource status endpoint
			federation.InternalFederatedSecretPlacementStatus,
			FederatedSecretPlacementSchemeFns{},
			func() runtime.Object { return &FederatedSecretPlacement{} },     // Register versioned resource
			func() runtime.Object { return &FederatedSecretPlacementList{} }, // Register versioned resource list
			&FederatedSecretPlacementStatusStrategy{builders.StatusStorageStrategySingleton},
		), federationFederationPlacementStorage,
		builders.NewApiResource( // Resource status endpoint
			federation.InternalFederationPlacementStatus,
			FederationPlacementSchemeFns{},
			func() runtime.Object { return &FederationPlacement{} },     // Register versioned resource
			func() runtime.Object { return &FederationPlacementList{} }, // Register versioned resource list
			&FederationPlacementStatusStrategy{builders.StatusStorageStrategySingleton},
		))

	// Required by code generated by go2idl
	AddToScheme        = ApiVersion.SchemaBuilder.AddToScheme
	SchemeBuilder      = ApiVersion.SchemaBuilder
	localSchemeBuilder = &SchemeBuilder
	SchemeGroupVersion = ApiVersion.GroupVersion
)

// Required by code generated by go2idl
// Kind takes an unqualified kind and returns a Group qualified GroupKind
func Kind(kind string) schema.GroupKind {
	return SchemeGroupVersion.WithKind(kind).GroupKind()
}

// Required by code generated by go2idl
// Resource takes an unqualified resource and returns a Group qualified GroupResource
func Resource(resource string) schema.GroupResource {
	return SchemeGroupVersion.WithResource(resource).GroupResource()
}

//
// FederatedCluster Functions and Structs
//
// +k8s:deepcopy-gen=false
type FederatedClusterSchemeFns struct {
	builders.DefaultSchemeFns
}

// +k8s:deepcopy-gen=false
type FederatedClusterStrategy struct {
	builders.DefaultStorageStrategy
}

// +k8s:deepcopy-gen=false
type FederatedClusterStatusStrategy struct {
	builders.DefaultStatusStorageStrategy
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type FederatedClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []FederatedCluster `json:"items"`
}

//
// FederatedReplicaSet Functions and Structs
//
// +k8s:deepcopy-gen=false
type FederatedReplicaSetSchemeFns struct {
	builders.DefaultSchemeFns
}

// +k8s:deepcopy-gen=false
type FederatedReplicaSetStrategy struct {
	builders.DefaultStorageStrategy
}

// +k8s:deepcopy-gen=false
type FederatedReplicaSetStatusStrategy struct {
	builders.DefaultStatusStorageStrategy
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type FederatedReplicaSetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []FederatedReplicaSet `json:"items"`
}

//
// FederatedReplicaSetOverride Functions and Structs
//
// +k8s:deepcopy-gen=false
type FederatedReplicaSetOverrideSchemeFns struct {
	builders.DefaultSchemeFns
}

// +k8s:deepcopy-gen=false
type FederatedReplicaSetOverrideStrategy struct {
	builders.DefaultStorageStrategy
}

// +k8s:deepcopy-gen=false
type FederatedReplicaSetOverrideStatusStrategy struct {
	builders.DefaultStatusStorageStrategy
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type FederatedReplicaSetOverrideList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []FederatedReplicaSetOverride `json:"items"`
}

//
// FederatedSecret Functions and Structs
//
// +k8s:deepcopy-gen=false
type FederatedSecretSchemeFns struct {
	builders.DefaultSchemeFns
}

// +k8s:deepcopy-gen=false
type FederatedSecretStrategy struct {
	builders.DefaultStorageStrategy
}

// +k8s:deepcopy-gen=false
type FederatedSecretStatusStrategy struct {
	builders.DefaultStatusStorageStrategy
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type FederatedSecretList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []FederatedSecret `json:"items"`
}

//
// FederatedSecretOverride Functions and Structs
//
// +k8s:deepcopy-gen=false
type FederatedSecretOverrideSchemeFns struct {
	builders.DefaultSchemeFns
}

// +k8s:deepcopy-gen=false
type FederatedSecretOverrideStrategy struct {
	builders.DefaultStorageStrategy
}

// +k8s:deepcopy-gen=false
type FederatedSecretOverrideStatusStrategy struct {
	builders.DefaultStatusStorageStrategy
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type FederatedSecretOverrideList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []FederatedSecretOverride `json:"items"`
}

//
// FederatedSecretPlacement Functions and Structs
//
// +k8s:deepcopy-gen=false
type FederatedSecretPlacementSchemeFns struct {
	builders.DefaultSchemeFns
}

// +k8s:deepcopy-gen=false
type FederatedSecretPlacementStrategy struct {
	builders.DefaultStorageStrategy
}

// +k8s:deepcopy-gen=false
type FederatedSecretPlacementStatusStrategy struct {
	builders.DefaultStatusStorageStrategy
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type FederatedSecretPlacementList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []FederatedSecretPlacement `json:"items"`
}

//
// FederationPlacement Functions and Structs
//
// +k8s:deepcopy-gen=false
type FederationPlacementSchemeFns struct {
	builders.DefaultSchemeFns
}

// +k8s:deepcopy-gen=false
type FederationPlacementStrategy struct {
	builders.DefaultStorageStrategy
}

// +k8s:deepcopy-gen=false
type FederationPlacementStatusStrategy struct {
	builders.DefaultStatusStorageStrategy
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type FederationPlacementList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []FederationPlacement `json:"items"`
}