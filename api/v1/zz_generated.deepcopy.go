//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2023.

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

package v1

import (
	corev1 "k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterManager) DeepCopyInto(out *ClusterManager) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterManager.
func (in *ClusterManager) DeepCopy() *ClusterManager {
	if in == nil {
		return nil
	}
	out := new(ClusterManager)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterManager) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterManagerList) DeepCopyInto(out *ClusterManagerList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ClusterManager, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterManagerList.
func (in *ClusterManagerList) DeepCopy() *ClusterManagerList {
	if in == nil {
		return nil
	}
	out := new(ClusterManagerList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterManagerList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterManagerSpec) DeepCopyInto(out *ClusterManagerSpec) {
	*out = *in
	in.Spec.DeepCopyInto(&out.Spec)
	out.LicenseManagerRef = in.LicenseManagerRef
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterManagerSpec.
func (in *ClusterManagerSpec) DeepCopy() *ClusterManagerSpec {
	if in == nil {
		return nil
	}
	out := new(ClusterManagerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterManagerStatus) DeepCopyInto(out *ClusterManagerStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterManagerStatus.
func (in *ClusterManagerStatus) DeepCopy() *ClusterManagerStatus {
	if in == nil {
		return nil
	}
	out := new(ClusterManagerStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IndexerCluster) DeepCopyInto(out *IndexerCluster) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IndexerCluster.
func (in *IndexerCluster) DeepCopy() *IndexerCluster {
	if in == nil {
		return nil
	}
	out := new(IndexerCluster)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *IndexerCluster) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IndexerClusterList) DeepCopyInto(out *IndexerClusterList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]IndexerCluster, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IndexerClusterList.
func (in *IndexerClusterList) DeepCopy() *IndexerClusterList {
	if in == nil {
		return nil
	}
	out := new(IndexerClusterList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *IndexerClusterList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IndexerClusterSpec) DeepCopyInto(out *IndexerClusterSpec) {
	*out = *in
	in.Spec.DeepCopyInto(&out.Spec)
	out.ClusterManagerRef = in.ClusterManagerRef
	out.LicenseManagerRef = in.LicenseManagerRef
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IndexerClusterSpec.
func (in *IndexerClusterSpec) DeepCopy() *IndexerClusterSpec {
	if in == nil {
		return nil
	}
	out := new(IndexerClusterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IndexerClusterStatus) DeepCopyInto(out *IndexerClusterStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IndexerClusterStatus.
func (in *IndexerClusterStatus) DeepCopy() *IndexerClusterStatus {
	if in == nil {
		return nil
	}
	out := new(IndexerClusterStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LicenseManager) DeepCopyInto(out *LicenseManager) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LicenseManager.
func (in *LicenseManager) DeepCopy() *LicenseManager {
	if in == nil {
		return nil
	}
	out := new(LicenseManager)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *LicenseManager) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LicenseManagerList) DeepCopyInto(out *LicenseManagerList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]LicenseManager, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LicenseManagerList.
func (in *LicenseManagerList) DeepCopy() *LicenseManagerList {
	if in == nil {
		return nil
	}
	out := new(LicenseManagerList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *LicenseManagerList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LicenseManagerSpec) DeepCopyInto(out *LicenseManagerSpec) {
	*out = *in
	in.Spec.DeepCopyInto(&out.Spec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LicenseManagerSpec.
func (in *LicenseManagerSpec) DeepCopy() *LicenseManagerSpec {
	if in == nil {
		return nil
	}
	out := new(LicenseManagerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LicenseManagerStatus) DeepCopyInto(out *LicenseManagerStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LicenseManagerStatus.
func (in *LicenseManagerStatus) DeepCopy() *LicenseManagerStatus {
	if in == nil {
		return nil
	}
	out := new(LicenseManagerStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MonitoringConsole) DeepCopyInto(out *MonitoringConsole) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MonitoringConsole.
func (in *MonitoringConsole) DeepCopy() *MonitoringConsole {
	if in == nil {
		return nil
	}
	out := new(MonitoringConsole)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MonitoringConsole) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MonitoringConsoleList) DeepCopyInto(out *MonitoringConsoleList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]MonitoringConsole, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MonitoringConsoleList.
func (in *MonitoringConsoleList) DeepCopy() *MonitoringConsoleList {
	if in == nil {
		return nil
	}
	out := new(MonitoringConsoleList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MonitoringConsoleList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MonitoringConsoleSpec) DeepCopyInto(out *MonitoringConsoleSpec) {
	*out = *in
	in.Spec.DeepCopyInto(&out.Spec)
	out.ClusterManagerRef = in.ClusterManagerRef
	out.LicenseManagerRef = in.LicenseManagerRef
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MonitoringConsoleSpec.
func (in *MonitoringConsoleSpec) DeepCopy() *MonitoringConsoleSpec {
	if in == nil {
		return nil
	}
	out := new(MonitoringConsoleSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MonitoringConsoleStatus) DeepCopyInto(out *MonitoringConsoleStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MonitoringConsoleStatus.
func (in *MonitoringConsoleStatus) DeepCopy() *MonitoringConsoleStatus {
	if in == nil {
		return nil
	}
	out := new(MonitoringConsoleStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SearchHeadCluster) DeepCopyInto(out *SearchHeadCluster) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SearchHeadCluster.
func (in *SearchHeadCluster) DeepCopy() *SearchHeadCluster {
	if in == nil {
		return nil
	}
	out := new(SearchHeadCluster)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SearchHeadCluster) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SearchHeadClusterList) DeepCopyInto(out *SearchHeadClusterList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]SearchHeadCluster, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SearchHeadClusterList.
func (in *SearchHeadClusterList) DeepCopy() *SearchHeadClusterList {
	if in == nil {
		return nil
	}
	out := new(SearchHeadClusterList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SearchHeadClusterList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SearchHeadClusterSpec) DeepCopyInto(out *SearchHeadClusterSpec) {
	*out = *in
	in.Spec.DeepCopyInto(&out.Spec)
	out.ClusterManagerRef = in.ClusterManagerRef
	out.LicenseManagerRef = in.LicenseManagerRef
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SearchHeadClusterSpec.
func (in *SearchHeadClusterSpec) DeepCopy() *SearchHeadClusterSpec {
	if in == nil {
		return nil
	}
	out := new(SearchHeadClusterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SearchHeadClusterStatus) DeepCopyInto(out *SearchHeadClusterStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SearchHeadClusterStatus.
func (in *SearchHeadClusterStatus) DeepCopy() *SearchHeadClusterStatus {
	if in == nil {
		return nil
	}
	out := new(SearchHeadClusterStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Spec) DeepCopyInto(out *Spec) {
	*out = *in
	in.Affinity.DeepCopyInto(&out.Affinity)
	if in.Tolerations != nil {
		in, out := &in.Tolerations, &out.Tolerations
		*out = make([]corev1.Toleration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	in.Resources.DeepCopyInto(&out.Resources)
	in.ServiceTemplate.DeepCopyInto(&out.ServiceTemplate)
	if in.TopologySpreadConstraints != nil {
		in, out := &in.TopologySpreadConstraints, &out.TopologySpreadConstraints
		*out = make([]corev1.TopologySpreadConstraint, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Spec.
func (in *Spec) DeepCopy() *Spec {
	if in == nil {
		return nil
	}
	out := new(Spec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Standalone) DeepCopyInto(out *Standalone) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Standalone.
func (in *Standalone) DeepCopy() *Standalone {
	if in == nil {
		return nil
	}
	out := new(Standalone)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Standalone) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StandaloneList) DeepCopyInto(out *StandaloneList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Standalone, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StandaloneList.
func (in *StandaloneList) DeepCopy() *StandaloneList {
	if in == nil {
		return nil
	}
	out := new(StandaloneList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *StandaloneList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StandaloneSpec) DeepCopyInto(out *StandaloneSpec) {
	*out = *in
	in.Spec.DeepCopyInto(&out.Spec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StandaloneSpec.
func (in *StandaloneSpec) DeepCopy() *StandaloneSpec {
	if in == nil {
		return nil
	}
	out := new(StandaloneSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StandaloneStatus) DeepCopyInto(out *StandaloneStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StandaloneStatus.
func (in *StandaloneStatus) DeepCopy() *StandaloneStatus {
	if in == nil {
		return nil
	}
	out := new(StandaloneStatus)
	in.DeepCopyInto(out)
	return out
}
