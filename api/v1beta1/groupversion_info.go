// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2025 Datadog, Inc.

// Package v1beta1 contains API Schema definitions for the chaos v1beta1 API group
// +kubebuilder:object:generate=true
// +groupName=chaos.datadoghq.com
package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/scheme"

	chaostypes "github.com/DataDog/chaos-controller/types"
)

// GroupName is exported for client-go purposes
const GroupName = chaostypes.GroupName

// APIVersion is exported for client-go purposes
const APIVersion = "v1beta1"

// DisruptionKind is the disruption kind
const DisruptionKind = "Disruption"

// DisruptionCronKind is the disruption cron kind
const DisruptionCronKind = "DisruptionCron"

var (
	// GroupVersion is group version used to register these objects
	GroupVersion = schema.GroupVersion{Group: GroupName, Version: APIVersion}

	// SchemeBuilder is used to add go types to the GroupVersionKind scheme
	SchemeBuilder = &scheme.Builder{GroupVersion: GroupVersion}

	// AddToScheme adds the types in this group-version to the given scheme.
	AddToScheme = SchemeBuilder.AddToScheme

	// TypeMeta is the Disruption Type Meta to create a disruption
	TypeMeta = metav1.TypeMeta{
		Kind:       DisruptionKind,
		APIVersion: GroupVersion.Identifier(),
	}
)
