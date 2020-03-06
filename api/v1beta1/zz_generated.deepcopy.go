// +build !ignore_autogenerated

// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2020 Datadog, Inc.
/*

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

// Code generated by controller-gen. DO NOT EDIT.

package v1beta1

import (
	"k8s.io/apimachinery/pkg/labels"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Disruption) DeepCopyInto(out *Disruption) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Disruption.
func (in *Disruption) DeepCopy() *Disruption {
	if in == nil {
		return nil
	}
	out := new(Disruption)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Disruption) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DisruptionList) DeepCopyInto(out *DisruptionList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Disruption, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DisruptionList.
func (in *DisruptionList) DeepCopy() *DisruptionList {
	if in == nil {
		return nil
	}
	out := new(DisruptionList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DisruptionList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DisruptionSpec) DeepCopyInto(out *DisruptionSpec) {
	*out = *in
	if in.Count != nil {
		in, out := &in.Count, &out.Count
		*out = new(int)
		**out = **in
	}
	if in.Selector != nil {
		in, out := &in.Selector, &out.Selector
		*out = make(labels.Set, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.NetworkFailure != nil {
		in, out := &in.NetworkFailure, &out.NetworkFailure
		*out = new(NetworkFailureSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.NetworkLatency != nil {
		in, out := &in.NetworkLatency, &out.NetworkLatency
		*out = new(NetworkLatencySpec)
		(*in).DeepCopyInto(*out)
	}
	if in.NodeFailure != nil {
		in, out := &in.NodeFailure, &out.NodeFailure
		*out = new(NodeFailureSpec)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DisruptionSpec.
func (in *DisruptionSpec) DeepCopy() *DisruptionSpec {
	if in == nil {
		return nil
	}
	out := new(DisruptionSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DisruptionStatus) DeepCopyInto(out *DisruptionStatus) {
	*out = *in
	if in.TargetPods != nil {
		in, out := &in.TargetPods, &out.TargetPods
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DisruptionStatus.
func (in *DisruptionStatus) DeepCopy() *DisruptionStatus {
	if in == nil {
		return nil
	}
	out := new(DisruptionStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkFailureSpec) DeepCopyInto(out *NetworkFailureSpec) {
	*out = *in
	if in.Hosts != nil {
		in, out := &in.Hosts, &out.Hosts
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkFailureSpec.
func (in *NetworkFailureSpec) DeepCopy() *NetworkFailureSpec {
	if in == nil {
		return nil
	}
	out := new(NetworkFailureSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkLatencySpec) DeepCopyInto(out *NetworkLatencySpec) {
	*out = *in
	if in.Hosts != nil {
		in, out := &in.Hosts, &out.Hosts
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkLatencySpec.
func (in *NetworkLatencySpec) DeepCopy() *NetworkLatencySpec {
	if in == nil {
		return nil
	}
	out := new(NetworkLatencySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeFailureSpec) DeepCopyInto(out *NodeFailureSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeFailureSpec.
func (in *NodeFailureSpec) DeepCopy() *NodeFailureSpec {
	if in == nil {
		return nil
	}
	out := new(NodeFailureSpec)
	in.DeepCopyInto(out)
	return out
}
