//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2022 The Koordinator Authors.

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

package extension

import (
	"github.com/koordinator-sh/koordinator/apis/slo/v1alpha1"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CPUBurstCfg) DeepCopyInto(out *CPUBurstCfg) {
	*out = *in
	if in.ClusterStrategy != nil {
		in, out := &in.ClusterStrategy, &out.ClusterStrategy
		*out = new(v1alpha1.CPUBurstStrategy)
		(*in).DeepCopyInto(*out)
	}
	if in.NodeStrategies != nil {
		in, out := &in.NodeStrategies, &out.NodeStrategies
		*out = make([]NodeCPUBurstCfg, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CPUBurstCfg.
func (in *CPUBurstCfg) DeepCopy() *CPUBurstCfg {
	if in == nil {
		return nil
	}
	out := new(CPUBurstCfg)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ColocationCfg) DeepCopyInto(out *ColocationCfg) {
	*out = *in
	in.ColocationStrategy.DeepCopyInto(&out.ColocationStrategy)
	if in.NodeConfigs != nil {
		in, out := &in.NodeConfigs, &out.NodeConfigs
		*out = make([]NodeColocationCfg, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ColocationCfg.
func (in *ColocationCfg) DeepCopy() *ColocationCfg {
	if in == nil {
		return nil
	}
	out := new(ColocationCfg)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ColocationStrategy) DeepCopyInto(out *ColocationStrategy) {
	*out = *in
	if in.Enable != nil {
		in, out := &in.Enable, &out.Enable
		*out = new(bool)
		**out = **in
	}
	if in.MetricAggregateDurationSeconds != nil {
		in, out := &in.MetricAggregateDurationSeconds, &out.MetricAggregateDurationSeconds
		*out = new(int64)
		**out = **in
	}
	if in.MetricReportIntervalSeconds != nil {
		in, out := &in.MetricReportIntervalSeconds, &out.MetricReportIntervalSeconds
		*out = new(int64)
		**out = **in
	}
	if in.MetricAggregatePolicy != nil {
		in, out := &in.MetricAggregatePolicy, &out.MetricAggregatePolicy
		*out = new(v1alpha1.AggregatePolicy)
		(*in).DeepCopyInto(*out)
	}
	if in.CPUReclaimThresholdPercent != nil {
		in, out := &in.CPUReclaimThresholdPercent, &out.CPUReclaimThresholdPercent
		*out = new(int64)
		**out = **in
	}
	if in.MemoryReclaimThresholdPercent != nil {
		in, out := &in.MemoryReclaimThresholdPercent, &out.MemoryReclaimThresholdPercent
		*out = new(int64)
		**out = **in
	}
	if in.MemoryCalculatePolicy != nil {
		in, out := &in.MemoryCalculatePolicy, &out.MemoryCalculatePolicy
		*out = new(CalculatePolicy)
		**out = **in
	}
	if in.DegradeTimeMinutes != nil {
		in, out := &in.DegradeTimeMinutes, &out.DegradeTimeMinutes
		*out = new(int64)
		**out = **in
	}
	if in.UpdateTimeThresholdSeconds != nil {
		in, out := &in.UpdateTimeThresholdSeconds, &out.UpdateTimeThresholdSeconds
		*out = new(int64)
		**out = **in
	}
	if in.ResourceDiffThreshold != nil {
		in, out := &in.ResourceDiffThreshold, &out.ResourceDiffThreshold
		*out = new(float64)
		**out = **in
	}
	if in.MidCPUThresholdPercent != nil {
		in, out := &in.MidCPUThresholdPercent, &out.MidCPUThresholdPercent
		*out = new(int64)
		**out = **in
	}
	if in.MidMemoryThresholdPercent != nil {
		in, out := &in.MidMemoryThresholdPercent, &out.MidMemoryThresholdPercent
		*out = new(int64)
		**out = **in
	}
	in.ColocationStrategyExtender.DeepCopyInto(&out.ColocationStrategyExtender)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ColocationStrategy.
func (in *ColocationStrategy) DeepCopy() *ColocationStrategy {
	if in == nil {
		return nil
	}
	out := new(ColocationStrategy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ColocationStrategyExtender) DeepCopyInto(out *ColocationStrategyExtender) {
	*out = *in
	in.Extensions.DeepCopyInto(&out.Extensions)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ColocationStrategyExtender.
func (in *ColocationStrategyExtender) DeepCopy() *ColocationStrategyExtender {
	if in == nil {
		return nil
	}
	out := new(ColocationStrategyExtender)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExtensionCfgMap) DeepCopyInto(out *ExtensionCfgMap) {
	*out = *in
	if in.Object != nil {
		in, out := &in.Object, &out.Object
		*out = make(map[string]ExtensionCfg, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExtensionCfgMap.
func (in *ExtensionCfgMap) DeepCopy() *ExtensionCfgMap {
	if in == nil {
		return nil
	}
	out := new(ExtensionCfgMap)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeCPUBurstCfg) DeepCopyInto(out *NodeCPUBurstCfg) {
	*out = *in
	in.NodeCfgProfile.DeepCopyInto(&out.NodeCfgProfile)
	if in.CPUBurstStrategy != nil {
		in, out := &in.CPUBurstStrategy, &out.CPUBurstStrategy
		*out = new(v1alpha1.CPUBurstStrategy)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeCPUBurstCfg.
func (in *NodeCPUBurstCfg) DeepCopy() *NodeCPUBurstCfg {
	if in == nil {
		return nil
	}
	out := new(NodeCPUBurstCfg)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeCfgProfile) DeepCopyInto(out *NodeCfgProfile) {
	*out = *in
	if in.NodeSelector != nil {
		in, out := &in.NodeSelector, &out.NodeSelector
		*out = new(v1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeCfgProfile.
func (in *NodeCfgProfile) DeepCopy() *NodeCfgProfile {
	if in == nil {
		return nil
	}
	out := new(NodeCfgProfile)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeColocationCfg) DeepCopyInto(out *NodeColocationCfg) {
	*out = *in
	in.NodeCfgProfile.DeepCopyInto(&out.NodeCfgProfile)
	in.ColocationStrategy.DeepCopyInto(&out.ColocationStrategy)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeColocationCfg.
func (in *NodeColocationCfg) DeepCopy() *NodeColocationCfg {
	if in == nil {
		return nil
	}
	out := new(NodeColocationCfg)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeResourceQOSStrategy) DeepCopyInto(out *NodeResourceQOSStrategy) {
	*out = *in
	in.NodeCfgProfile.DeepCopyInto(&out.NodeCfgProfile)
	if in.ResourceQOSStrategy != nil {
		in, out := &in.ResourceQOSStrategy, &out.ResourceQOSStrategy
		*out = new(v1alpha1.ResourceQOSStrategy)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeResourceQOSStrategy.
func (in *NodeResourceQOSStrategy) DeepCopy() *NodeResourceQOSStrategy {
	if in == nil {
		return nil
	}
	out := new(NodeResourceQOSStrategy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeResourceThresholdStrategy) DeepCopyInto(out *NodeResourceThresholdStrategy) {
	*out = *in
	in.NodeCfgProfile.DeepCopyInto(&out.NodeCfgProfile)
	if in.ResourceThresholdStrategy != nil {
		in, out := &in.ResourceThresholdStrategy, &out.ResourceThresholdStrategy
		*out = new(v1alpha1.ResourceThresholdStrategy)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeResourceThresholdStrategy.
func (in *NodeResourceThresholdStrategy) DeepCopy() *NodeResourceThresholdStrategy {
	if in == nil {
		return nil
	}
	out := new(NodeResourceThresholdStrategy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeSystemStrategy) DeepCopyInto(out *NodeSystemStrategy) {
	*out = *in
	in.NodeCfgProfile.DeepCopyInto(&out.NodeCfgProfile)
	if in.SystemStrategy != nil {
		in, out := &in.SystemStrategy, &out.SystemStrategy
		*out = new(v1alpha1.SystemStrategy)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeSystemStrategy.
func (in *NodeSystemStrategy) DeepCopy() *NodeSystemStrategy {
	if in == nil {
		return nil
	}
	out := new(NodeSystemStrategy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceQOSCfg) DeepCopyInto(out *ResourceQOSCfg) {
	*out = *in
	if in.ClusterStrategy != nil {
		in, out := &in.ClusterStrategy, &out.ClusterStrategy
		*out = new(v1alpha1.ResourceQOSStrategy)
		(*in).DeepCopyInto(*out)
	}
	if in.NodeStrategies != nil {
		in, out := &in.NodeStrategies, &out.NodeStrategies
		*out = make([]NodeResourceQOSStrategy, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceQOSCfg.
func (in *ResourceQOSCfg) DeepCopy() *ResourceQOSCfg {
	if in == nil {
		return nil
	}
	out := new(ResourceQOSCfg)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceThresholdCfg) DeepCopyInto(out *ResourceThresholdCfg) {
	*out = *in
	if in.ClusterStrategy != nil {
		in, out := &in.ClusterStrategy, &out.ClusterStrategy
		*out = new(v1alpha1.ResourceThresholdStrategy)
		(*in).DeepCopyInto(*out)
	}
	if in.NodeStrategies != nil {
		in, out := &in.NodeStrategies, &out.NodeStrategies
		*out = make([]NodeResourceThresholdStrategy, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceThresholdCfg.
func (in *ResourceThresholdCfg) DeepCopy() *ResourceThresholdCfg {
	if in == nil {
		return nil
	}
	out := new(ResourceThresholdCfg)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SystemCfg) DeepCopyInto(out *SystemCfg) {
	*out = *in
	if in.ClusterStrategy != nil {
		in, out := &in.ClusterStrategy, &out.ClusterStrategy
		*out = new(v1alpha1.SystemStrategy)
		(*in).DeepCopyInto(*out)
	}
	if in.NodeStrategies != nil {
		in, out := &in.NodeStrategies, &out.NodeStrategies
		*out = make([]NodeSystemStrategy, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SystemCfg.
func (in *SystemCfg) DeepCopy() *SystemCfg {
	if in == nil {
		return nil
	}
	out := new(SystemCfg)
	in.DeepCopyInto(out)
	return out
}
