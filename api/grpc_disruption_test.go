// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2025 Datadog, Inc.

package api_test

import (
	"github.com/DataDog/chaos-controller/api/v1beta1"
	"github.com/hashicorp/go-multierror"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("GRPCDisruption Validation", func() {
	var spec v1beta1.GRPCDisruptionSpec

	BeforeEach(func() {
		spec = v1beta1.GRPCDisruptionSpec{
			Port:      50051,
			Endpoints: []v1beta1.EndpointAlteration{},
		}
	})

	Context("Error and override are both undefined", func() {
		It("errors because exactly one of error or override must be defined for an alteration", func() {
			spec.Endpoints = []v1beta1.EndpointAlteration{
				{
					TargetEndpoint:   "/chaosdogfood.ChaosDogfood/order",
					ErrorToReturn:    "",
					OverrideToReturn: "",
					QueryPercent:     100,
				},
			}
			err := spec.Validate().(*multierror.Error)
			Expect(err.Len()).To(Equal(1))
			Expect(err.Errors[0].Error()).To(Equal("GRPC: the gRPC disruption must have exactly one of ErrorToReturn or OverrideToReturn specified for endpoint /chaosdogfood.ChaosDogfood/order"))
		})
	})

	Describe("One target endpoint", func() {
		Context("with query percentage 100", func() {
			It("Passes validation", func() {
				spec.Endpoints = []v1beta1.EndpointAlteration{
					{
						TargetEndpoint:   "/chaosdogfood.ChaosDogfood/order",
						ErrorToReturn:    "CANCELED",
						OverrideToReturn: "",
						QueryPercent:     100,
					},
				}

				Expect(spec.Validate()).To(Succeed())
			})
		})

		Context("with query percentage 99", func() {
			It("Passes validation", func() {
				spec.Endpoints = []v1beta1.EndpointAlteration{
					{
						TargetEndpoint:   "/chaosdogfood.ChaosDogfood/order",
						ErrorToReturn:    "CANCELED",
						OverrideToReturn: "",
						QueryPercent:     99,
					},
				}

				Expect(spec.Validate()).To(Succeed())
			})
		})
	})

	Context("Two separate target endpoints each have query percentage 100", func() {
		It("Passes validation", func() {
			spec.Endpoints = []v1beta1.EndpointAlteration{
				{
					TargetEndpoint:   "/chaosdogfood.ChaosDogfood/order",
					ErrorToReturn:    "CANCELED",
					OverrideToReturn: "",
					QueryPercent:     100,
				},
				{
					TargetEndpoint:   "/chaosdogfood.ChaosDogfood/getCatalog",
					ErrorToReturn:    "CANCELED",
					OverrideToReturn: "",
					QueryPercent:     100,
				},
			}

			Expect(spec.Validate()).To(Succeed())
		})
	})

	Describe("One target endpoint with two alterations", func() {
		Context("which in total have query percentage 100", func() {
			It("Passes validation", func() {
				spec.Endpoints = []v1beta1.EndpointAlteration{
					{
						TargetEndpoint:   "/chaosdogfood.ChaosDogfood/order",
						ErrorToReturn:    "CANCELED",
						OverrideToReturn: "",
						QueryPercent:     60,
					},
					{
						TargetEndpoint:   "/chaosdogfood.ChaosDogfood/order",
						ErrorToReturn:    "ALREADY_EXISTS",
						OverrideToReturn: "",
						QueryPercent:     40,
					},
				}

				Expect(spec.Validate()).To(Succeed())
			})
		})

		Context("which in total have query percentage less than 100", func() {
			It("Passes validation", func() {
				spec.Endpoints = []v1beta1.EndpointAlteration{
					{
						TargetEndpoint:   "/chaosdogfood.ChaosDogfood/order",
						ErrorToReturn:    "CANCELED",
						OverrideToReturn: "",
						QueryPercent:     49,
					},
					{
						TargetEndpoint:   "/chaosdogfood.ChaosDogfood/order",
						ErrorToReturn:    "ALREADY_EXISTS",
						OverrideToReturn: "",
						QueryPercent:     49,
					},
				}

				Expect(spec.Validate()).To(Succeed())
			})
		})

		Context("which in total have query percentage greater than 100", func() {
			It("Fails validation", func() {
				spec.Endpoints = []v1beta1.EndpointAlteration{
					{
						TargetEndpoint:   "/chaosdogfood.ChaosDogfood/order",
						ErrorToReturn:    "CANCELED",
						OverrideToReturn: "",
						QueryPercent:     50,
					},
					{
						TargetEndpoint:   "/chaosdogfood.ChaosDogfood/order",
						ErrorToReturn:    "ALREADY_EXISTS",
						OverrideToReturn: "",
						QueryPercent:     51,
					},
				}

				Expect(spec.Validate()).ToNot(Succeed())
			})
		})
	})

	Describe("Alterations with ErrorToReturn", func() {
		Context("which are in the standard grpc errors", func() {
			It("Passes validation", func() {
				spec.Endpoints = []v1beta1.EndpointAlteration{
					{
						TargetEndpoint:   "/chaosdogfood.ChaosDogfood/order",
						ErrorToReturn:    "CANCELED",
						OverrideToReturn: "",
						QueryPercent:     0,
					},
					{
						TargetEndpoint:   "/chaosdogfood.ChaosDogfood/order",
						ErrorToReturn:    "ALREADY_EXISTS",
						OverrideToReturn: "",
						QueryPercent:     0,
					},
					{
						TargetEndpoint:   "/chaosdogfood.ChaosDogfood/order",
						ErrorToReturn:    "UNKNOWN",
						OverrideToReturn: "",
						QueryPercent:     0,
					},
					{
						TargetEndpoint:   "/chaosdogfood.ChaosDogfood/order",
						ErrorToReturn:    "INVALID_ARGUMENT",
						OverrideToReturn: "",
						QueryPercent:     0,
					},
					{
						TargetEndpoint:   "/chaosdogfood.ChaosDogfood/order",
						ErrorToReturn:    "DEADLINE_EXCEEDED",
						OverrideToReturn: "",
						QueryPercent:     0,
					},
					{
						TargetEndpoint:   "/chaosdogfood.ChaosDogfood/order",
						ErrorToReturn:    "NOT_FOUND",
						OverrideToReturn: "",
						QueryPercent:     0,
					},
					{
						TargetEndpoint:   "/chaosdogfood.ChaosDogfood/order",
						ErrorToReturn:    "PERMISSION_DENIED",
						OverrideToReturn: "",
						QueryPercent:     0,
					},
					{
						TargetEndpoint:   "/chaosdogfood.ChaosDogfood/order",
						ErrorToReturn:    "RESOURCE_EXHAUSTED",
						OverrideToReturn: "",
						QueryPercent:     0,
					},
					{
						TargetEndpoint:   "/chaosdogfood.ChaosDogfood/order",
						ErrorToReturn:    "FAILED_PRECONDITION",
						OverrideToReturn: "",
						QueryPercent:     0,
					},
					{
						TargetEndpoint:   "/chaosdogfood.ChaosDogfood/order",
						ErrorToReturn:    "ABORTED",
						OverrideToReturn: "",
						QueryPercent:     0,
					},
					{
						TargetEndpoint:   "/chaosdogfood.ChaosDogfood/order",
						ErrorToReturn:    "OUT_OF_RANGE",
						OverrideToReturn: "",
						QueryPercent:     0,
					},
					{
						TargetEndpoint:   "/chaosdogfood.ChaosDogfood/order",
						ErrorToReturn:    "UNIMPLEMENTED",
						OverrideToReturn: "",
						QueryPercent:     0,
					},
				}

				Expect(spec.Validate()).To(Succeed())
			})
		})

		Context("In the standard grpc errors but which in total exceed 100%", func() {
			It("Fails validation", func() {
				spec.Endpoints = []v1beta1.EndpointAlteration{
					{
						TargetEndpoint:   "/chaosdogfood.ChaosDogfood/order",
						ErrorToReturn:    "CANCELED",
						OverrideToReturn: "",
						QueryPercent:     90,
					},
					{
						TargetEndpoint:   "/chaosdogfood.ChaosDogfood/order",
						ErrorToReturn:    "ALREADY_EXISTS",
						OverrideToReturn: "",
						QueryPercent:     0,
					},
					{
						TargetEndpoint:   "/chaosdogfood.ChaosDogfood/order",
						ErrorToReturn:    "UNKNOWN",
						OverrideToReturn: "",
						QueryPercent:     0,
					},
					{
						TargetEndpoint:   "/chaosdogfood.ChaosDogfood/order",
						ErrorToReturn:    "INVALID_ARGUMENT",
						OverrideToReturn: "",
						QueryPercent:     0,
					},
					{
						TargetEndpoint:   "/chaosdogfood.ChaosDogfood/order",
						ErrorToReturn:    "DEADLINE_EXCEEDED",
						OverrideToReturn: "",
						QueryPercent:     0,
					},
					{
						TargetEndpoint:   "/chaosdogfood.ChaosDogfood/order",
						ErrorToReturn:    "NOT_FOUND",
						OverrideToReturn: "",
						QueryPercent:     0,
					},
					{
						TargetEndpoint:   "/chaosdogfood.ChaosDogfood/order",
						ErrorToReturn:    "PERMISSION_DENIED",
						OverrideToReturn: "",
						QueryPercent:     0,
					},
					{
						TargetEndpoint:   "/chaosdogfood.ChaosDogfood/order",
						ErrorToReturn:    "RESOURCE_EXHAUSTED",
						OverrideToReturn: "",
						QueryPercent:     0,
					},
					{
						TargetEndpoint:   "/chaosdogfood.ChaosDogfood/order",
						ErrorToReturn:    "FAILED_PRECONDITION",
						OverrideToReturn: "",
						QueryPercent:     0,
					},
					{
						TargetEndpoint:   "/chaosdogfood.ChaosDogfood/order",
						ErrorToReturn:    "ABORTED",
						OverrideToReturn: "",
						QueryPercent:     0,
					},
					{
						TargetEndpoint:   "/chaosdogfood.ChaosDogfood/order",
						ErrorToReturn:    "OUT_OF_RANGE",
						OverrideToReturn: "",
						QueryPercent:     0,
					},
					{
						TargetEndpoint:   "/chaosdogfood.ChaosDogfood/order",
						ErrorToReturn:    "UNIMPLEMENTED",
						OverrideToReturn: "",
						QueryPercent:     0,
					},
				}

				Expect(spec.Validate()).ToNot(Succeed())
			})
		})
	})
})
