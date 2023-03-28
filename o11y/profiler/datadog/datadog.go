// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2023 Datadog, Inc.

package datadog

import (
	"github.com/DataDog/chaos-controller/o11y/profiler/types"
)

// Sink describes a datadog profiler
type Sink struct{}

// New datadog sink
func New() *Sink {
	return &Sink{}
}

// Start returns nil
func (d *Sink) Start() {
	//FIXME
}

// Stop returns nil
func (d *Sink) Stop() {
	//FIXME
}

// GetSinkName returns the name of the sink
func (d *Sink) GetSinkName() string {
	return string(types.SinkDriverDatadog)
}
