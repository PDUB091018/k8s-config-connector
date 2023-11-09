//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Copyright 2020 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// *** DISCLAIMER ***
// Config Connector's go-client for CRDs is currently in ALPHA, which means
// that future versions of the go-client may include breaking changes.
// Please try it out and give us feedback!

// Code generated by main. DO NOT EDIT.

package v1beta1

import (
	v1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/k8s/v1alpha1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterAdminUsers) DeepCopyInto(out *ClusterAdminUsers) {
	*out = *in
	out.UsernameRef = in.UsernameRef
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterAdminUsers.
func (in *ClusterAdminUsers) DeepCopy() *ClusterAdminUsers {
	if in == nil {
		return nil
	}
	out := new(ClusterAdminUsers)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterAuthorization) DeepCopyInto(out *ClusterAuthorization) {
	*out = *in
	out.AdminUsers = in.AdminUsers
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterAuthorization.
func (in *ClusterAuthorization) DeepCopy() *ClusterAuthorization {
	if in == nil {
		return nil
	}
	out := new(ClusterAuthorization)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterControlPlane) DeepCopyInto(out *ClusterControlPlane) {
	*out = *in
	if in.Local != nil {
		in, out := &in.Local, &out.Local
		*out = new(ClusterLocal)
		(*in).DeepCopyInto(*out)
	}
	if in.Remote != nil {
		in, out := &in.Remote, &out.Remote
		*out = new(ClusterRemote)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterControlPlane.
func (in *ClusterControlPlane) DeepCopy() *ClusterControlPlane {
	if in == nil {
		return nil
	}
	out := new(ClusterControlPlane)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterControlPlaneEncryption) DeepCopyInto(out *ClusterControlPlaneEncryption) {
	*out = *in
	if in.KmsKeyActiveVersion != nil {
		in, out := &in.KmsKeyActiveVersion, &out.KmsKeyActiveVersion
		*out = new(string)
		**out = **in
	}
	if in.KmsKeyRef != nil {
		in, out := &in.KmsKeyRef, &out.KmsKeyRef
		*out = new(v1alpha1.ResourceRef)
		**out = **in
	}
	if in.KmsKeyState != nil {
		in, out := &in.KmsKeyState, &out.KmsKeyState
		*out = new(string)
		**out = **in
	}
	if in.KmsStatus != nil {
		in, out := &in.KmsStatus, &out.KmsStatus
		*out = make([]ClusterKmsStatus, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterControlPlaneEncryption.
func (in *ClusterControlPlaneEncryption) DeepCopy() *ClusterControlPlaneEncryption {
	if in == nil {
		return nil
	}
	out := new(ClusterControlPlaneEncryption)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterFleet) DeepCopyInto(out *ClusterFleet) {
	*out = *in
	if in.Membership != nil {
		in, out := &in.Membership, &out.Membership
		*out = new(string)
		**out = **in
	}
	out.ProjectRef = in.ProjectRef
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterFleet.
func (in *ClusterFleet) DeepCopy() *ClusterFleet {
	if in == nil {
		return nil
	}
	out := new(ClusterFleet)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterIngress) DeepCopyInto(out *ClusterIngress) {
	*out = *in
	if in.Disabled != nil {
		in, out := &in.Disabled, &out.Disabled
		*out = new(bool)
		**out = **in
	}
	if in.Ipv4Vip != nil {
		in, out := &in.Ipv4Vip, &out.Ipv4Vip
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterIngress.
func (in *ClusterIngress) DeepCopy() *ClusterIngress {
	if in == nil {
		return nil
	}
	out := new(ClusterIngress)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterKmsStatus) DeepCopyInto(out *ClusterKmsStatus) {
	*out = *in
	if in.Code != nil {
		in, out := &in.Code, &out.Code
		*out = new(int)
		**out = **in
	}
	if in.Message != nil {
		in, out := &in.Message, &out.Message
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterKmsStatus.
func (in *ClusterKmsStatus) DeepCopy() *ClusterKmsStatus {
	if in == nil {
		return nil
	}
	out := new(ClusterKmsStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterLocal) DeepCopyInto(out *ClusterLocal) {
	*out = *in
	if in.MachineFilter != nil {
		in, out := &in.MachineFilter, &out.MachineFilter
		*out = new(string)
		**out = **in
	}
	if in.NodeCount != nil {
		in, out := &in.NodeCount, &out.NodeCount
		*out = new(int)
		**out = **in
	}
	if in.NodeLocation != nil {
		in, out := &in.NodeLocation, &out.NodeLocation
		*out = new(string)
		**out = **in
	}
	if in.SharedDeploymentPolicy != nil {
		in, out := &in.SharedDeploymentPolicy, &out.SharedDeploymentPolicy
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterLocal.
func (in *ClusterLocal) DeepCopy() *ClusterLocal {
	if in == nil {
		return nil
	}
	out := new(ClusterLocal)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterMaintenanceEventsStatus) DeepCopyInto(out *ClusterMaintenanceEventsStatus) {
	*out = *in
	if in.CreateTime != nil {
		in, out := &in.CreateTime, &out.CreateTime
		*out = new(string)
		**out = **in
	}
	if in.EndTime != nil {
		in, out := &in.EndTime, &out.EndTime
		*out = new(string)
		**out = **in
	}
	if in.Operation != nil {
		in, out := &in.Operation, &out.Operation
		*out = new(string)
		**out = **in
	}
	if in.Schedule != nil {
		in, out := &in.Schedule, &out.Schedule
		*out = new(string)
		**out = **in
	}
	if in.StartTime != nil {
		in, out := &in.StartTime, &out.StartTime
		*out = new(string)
		**out = **in
	}
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(string)
		**out = **in
	}
	if in.TargetVersion != nil {
		in, out := &in.TargetVersion, &out.TargetVersion
		*out = new(string)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
	if in.UpdateTime != nil {
		in, out := &in.UpdateTime, &out.UpdateTime
		*out = new(string)
		**out = **in
	}
	if in.Uuid != nil {
		in, out := &in.Uuid, &out.Uuid
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterMaintenanceEventsStatus.
func (in *ClusterMaintenanceEventsStatus) DeepCopy() *ClusterMaintenanceEventsStatus {
	if in == nil {
		return nil
	}
	out := new(ClusterMaintenanceEventsStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterMaintenancePolicy) DeepCopyInto(out *ClusterMaintenancePolicy) {
	*out = *in
	in.Window.DeepCopyInto(&out.Window)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterMaintenancePolicy.
func (in *ClusterMaintenancePolicy) DeepCopy() *ClusterMaintenancePolicy {
	if in == nil {
		return nil
	}
	out := new(ClusterMaintenancePolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterNetworking) DeepCopyInto(out *ClusterNetworking) {
	*out = *in
	if in.ClusterIpv4CidrBlocks != nil {
		in, out := &in.ClusterIpv4CidrBlocks, &out.ClusterIpv4CidrBlocks
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ClusterIpv6CidrBlocks != nil {
		in, out := &in.ClusterIpv6CidrBlocks, &out.ClusterIpv6CidrBlocks
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.NetworkType != nil {
		in, out := &in.NetworkType, &out.NetworkType
		*out = new(string)
		**out = **in
	}
	if in.ServicesIpv4CidrBlocks != nil {
		in, out := &in.ServicesIpv4CidrBlocks, &out.ServicesIpv4CidrBlocks
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ServicesIpv6CidrBlocks != nil {
		in, out := &in.ServicesIpv6CidrBlocks, &out.ServicesIpv6CidrBlocks
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterNetworking.
func (in *ClusterNetworking) DeepCopy() *ClusterNetworking {
	if in == nil {
		return nil
	}
	out := new(ClusterNetworking)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterRecurringWindow) DeepCopyInto(out *ClusterRecurringWindow) {
	*out = *in
	if in.Recurrence != nil {
		in, out := &in.Recurrence, &out.Recurrence
		*out = new(string)
		**out = **in
	}
	if in.Window != nil {
		in, out := &in.Window, &out.Window
		*out = new(ClusterWindow)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterRecurringWindow.
func (in *ClusterRecurringWindow) DeepCopy() *ClusterRecurringWindow {
	if in == nil {
		return nil
	}
	out := new(ClusterRecurringWindow)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterRemote) DeepCopyInto(out *ClusterRemote) {
	*out = *in
	if in.NodeLocation != nil {
		in, out := &in.NodeLocation, &out.NodeLocation
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterRemote.
func (in *ClusterRemote) DeepCopy() *ClusterRemote {
	if in == nil {
		return nil
	}
	out := new(ClusterRemote)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterSystemAddonsConfig) DeepCopyInto(out *ClusterSystemAddonsConfig) {
	*out = *in
	if in.Ingress != nil {
		in, out := &in.Ingress, &out.Ingress
		*out = new(ClusterIngress)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterSystemAddonsConfig.
func (in *ClusterSystemAddonsConfig) DeepCopy() *ClusterSystemAddonsConfig {
	if in == nil {
		return nil
	}
	out := new(ClusterSystemAddonsConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterWindow) DeepCopyInto(out *ClusterWindow) {
	*out = *in
	if in.EndTime != nil {
		in, out := &in.EndTime, &out.EndTime
		*out = new(string)
		**out = **in
	}
	if in.StartTime != nil {
		in, out := &in.StartTime, &out.StartTime
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterWindow.
func (in *ClusterWindow) DeepCopy() *ClusterWindow {
	if in == nil {
		return nil
	}
	out := new(ClusterWindow)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EdgeContainerCluster) DeepCopyInto(out *EdgeContainerCluster) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EdgeContainerCluster.
func (in *EdgeContainerCluster) DeepCopy() *EdgeContainerCluster {
	if in == nil {
		return nil
	}
	out := new(EdgeContainerCluster)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *EdgeContainerCluster) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EdgeContainerClusterList) DeepCopyInto(out *EdgeContainerClusterList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]EdgeContainerCluster, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EdgeContainerClusterList.
func (in *EdgeContainerClusterList) DeepCopy() *EdgeContainerClusterList {
	if in == nil {
		return nil
	}
	out := new(EdgeContainerClusterList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *EdgeContainerClusterList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EdgeContainerClusterSpec) DeepCopyInto(out *EdgeContainerClusterSpec) {
	*out = *in
	out.Authorization = in.Authorization
	if in.ControlPlane != nil {
		in, out := &in.ControlPlane, &out.ControlPlane
		*out = new(ClusterControlPlane)
		(*in).DeepCopyInto(*out)
	}
	if in.ControlPlaneEncryption != nil {
		in, out := &in.ControlPlaneEncryption, &out.ControlPlaneEncryption
		*out = new(ClusterControlPlaneEncryption)
		(*in).DeepCopyInto(*out)
	}
	if in.DefaultMaxPodsPerNode != nil {
		in, out := &in.DefaultMaxPodsPerNode, &out.DefaultMaxPodsPerNode
		*out = new(int)
		**out = **in
	}
	if in.ExternalLoadBalancerIpv4AddressPools != nil {
		in, out := &in.ExternalLoadBalancerIpv4AddressPools, &out.ExternalLoadBalancerIpv4AddressPools
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	in.Fleet.DeepCopyInto(&out.Fleet)
	if in.MaintenancePolicy != nil {
		in, out := &in.MaintenancePolicy, &out.MaintenancePolicy
		*out = new(ClusterMaintenancePolicy)
		(*in).DeepCopyInto(*out)
	}
	in.Networking.DeepCopyInto(&out.Networking)
	out.ProjectRef = in.ProjectRef
	if in.ReleaseChannel != nil {
		in, out := &in.ReleaseChannel, &out.ReleaseChannel
		*out = new(string)
		**out = **in
	}
	if in.ResourceID != nil {
		in, out := &in.ResourceID, &out.ResourceID
		*out = new(string)
		**out = **in
	}
	if in.SystemAddonsConfig != nil {
		in, out := &in.SystemAddonsConfig, &out.SystemAddonsConfig
		*out = new(ClusterSystemAddonsConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.TargetVersion != nil {
		in, out := &in.TargetVersion, &out.TargetVersion
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EdgeContainerClusterSpec.
func (in *EdgeContainerClusterSpec) DeepCopy() *EdgeContainerClusterSpec {
	if in == nil {
		return nil
	}
	out := new(EdgeContainerClusterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EdgeContainerClusterStatus) DeepCopyInto(out *EdgeContainerClusterStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1alpha1.Condition, len(*in))
		copy(*out, *in)
	}
	if in.ClusterCaCertificate != nil {
		in, out := &in.ClusterCaCertificate, &out.ClusterCaCertificate
		*out = new(string)
		**out = **in
	}
	if in.ControlPlaneVersion != nil {
		in, out := &in.ControlPlaneVersion, &out.ControlPlaneVersion
		*out = new(string)
		**out = **in
	}
	if in.CreateTime != nil {
		in, out := &in.CreateTime, &out.CreateTime
		*out = new(string)
		**out = **in
	}
	if in.Endpoint != nil {
		in, out := &in.Endpoint, &out.Endpoint
		*out = new(string)
		**out = **in
	}
	if in.MaintenanceEvents != nil {
		in, out := &in.MaintenanceEvents, &out.MaintenanceEvents
		*out = make([]ClusterMaintenanceEventsStatus, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.NodeVersion != nil {
		in, out := &in.NodeVersion, &out.NodeVersion
		*out = new(string)
		**out = **in
	}
	if in.ObservedGeneration != nil {
		in, out := &in.ObservedGeneration, &out.ObservedGeneration
		*out = new(int)
		**out = **in
	}
	if in.Port != nil {
		in, out := &in.Port, &out.Port
		*out = new(int)
		**out = **in
	}
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(string)
		**out = **in
	}
	if in.UpdateTime != nil {
		in, out := &in.UpdateTime, &out.UpdateTime
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EdgeContainerClusterStatus.
func (in *EdgeContainerClusterStatus) DeepCopy() *EdgeContainerClusterStatus {
	if in == nil {
		return nil
	}
	out := new(EdgeContainerClusterStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EdgeContainerNodePool) DeepCopyInto(out *EdgeContainerNodePool) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EdgeContainerNodePool.
func (in *EdgeContainerNodePool) DeepCopy() *EdgeContainerNodePool {
	if in == nil {
		return nil
	}
	out := new(EdgeContainerNodePool)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *EdgeContainerNodePool) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EdgeContainerNodePoolList) DeepCopyInto(out *EdgeContainerNodePoolList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]EdgeContainerNodePool, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EdgeContainerNodePoolList.
func (in *EdgeContainerNodePoolList) DeepCopy() *EdgeContainerNodePoolList {
	if in == nil {
		return nil
	}
	out := new(EdgeContainerNodePoolList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *EdgeContainerNodePoolList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EdgeContainerNodePoolSpec) DeepCopyInto(out *EdgeContainerNodePoolSpec) {
	*out = *in
	out.ClusterRef = in.ClusterRef
	if in.LocalDiskEncryption != nil {
		in, out := &in.LocalDiskEncryption, &out.LocalDiskEncryption
		*out = new(NodepoolLocalDiskEncryption)
		(*in).DeepCopyInto(*out)
	}
	if in.MachineFilter != nil {
		in, out := &in.MachineFilter, &out.MachineFilter
		*out = new(string)
		**out = **in
	}
	if in.NodeConfig != nil {
		in, out := &in.NodeConfig, &out.NodeConfig
		*out = new(NodepoolNodeConfig)
		(*in).DeepCopyInto(*out)
	}
	out.ProjectRef = in.ProjectRef
	if in.ResourceID != nil {
		in, out := &in.ResourceID, &out.ResourceID
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EdgeContainerNodePoolSpec.
func (in *EdgeContainerNodePoolSpec) DeepCopy() *EdgeContainerNodePoolSpec {
	if in == nil {
		return nil
	}
	out := new(EdgeContainerNodePoolSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EdgeContainerNodePoolStatus) DeepCopyInto(out *EdgeContainerNodePoolStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1alpha1.Condition, len(*in))
		copy(*out, *in)
	}
	if in.CreateTime != nil {
		in, out := &in.CreateTime, &out.CreateTime
		*out = new(string)
		**out = **in
	}
	if in.NodeVersion != nil {
		in, out := &in.NodeVersion, &out.NodeVersion
		*out = new(string)
		**out = **in
	}
	if in.ObservedGeneration != nil {
		in, out := &in.ObservedGeneration, &out.ObservedGeneration
		*out = new(int)
		**out = **in
	}
	if in.UpdateTime != nil {
		in, out := &in.UpdateTime, &out.UpdateTime
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EdgeContainerNodePoolStatus.
func (in *EdgeContainerNodePoolStatus) DeepCopy() *EdgeContainerNodePoolStatus {
	if in == nil {
		return nil
	}
	out := new(EdgeContainerNodePoolStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EdgeContainerVpnConnection) DeepCopyInto(out *EdgeContainerVpnConnection) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EdgeContainerVpnConnection.
func (in *EdgeContainerVpnConnection) DeepCopy() *EdgeContainerVpnConnection {
	if in == nil {
		return nil
	}
	out := new(EdgeContainerVpnConnection)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *EdgeContainerVpnConnection) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EdgeContainerVpnConnectionList) DeepCopyInto(out *EdgeContainerVpnConnectionList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]EdgeContainerVpnConnection, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EdgeContainerVpnConnectionList.
func (in *EdgeContainerVpnConnectionList) DeepCopy() *EdgeContainerVpnConnectionList {
	if in == nil {
		return nil
	}
	out := new(EdgeContainerVpnConnectionList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *EdgeContainerVpnConnectionList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EdgeContainerVpnConnectionSpec) DeepCopyInto(out *EdgeContainerVpnConnectionSpec) {
	*out = *in
	out.ClusterRef = in.ClusterRef
	if in.EnableHighAvailability != nil {
		in, out := &in.EnableHighAvailability, &out.EnableHighAvailability
		*out = new(bool)
		**out = **in
	}
	if in.NatGatewayIp != nil {
		in, out := &in.NatGatewayIp, &out.NatGatewayIp
		*out = new(string)
		**out = **in
	}
	out.ProjectRef = in.ProjectRef
	if in.ResourceID != nil {
		in, out := &in.ResourceID, &out.ResourceID
		*out = new(string)
		**out = **in
	}
	if in.Router != nil {
		in, out := &in.Router, &out.Router
		*out = new(string)
		**out = **in
	}
	if in.Vpc != nil {
		in, out := &in.Vpc, &out.Vpc
		*out = new(string)
		**out = **in
	}
	if in.VpcProject != nil {
		in, out := &in.VpcProject, &out.VpcProject
		*out = new(VpnconnectionVpcProject)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EdgeContainerVpnConnectionSpec.
func (in *EdgeContainerVpnConnectionSpec) DeepCopy() *EdgeContainerVpnConnectionSpec {
	if in == nil {
		return nil
	}
	out := new(EdgeContainerVpnConnectionSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EdgeContainerVpnConnectionStatus) DeepCopyInto(out *EdgeContainerVpnConnectionStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1alpha1.Condition, len(*in))
		copy(*out, *in)
	}
	if in.CreateTime != nil {
		in, out := &in.CreateTime, &out.CreateTime
		*out = new(string)
		**out = **in
	}
	if in.Details != nil {
		in, out := &in.Details, &out.Details
		*out = make([]VpnconnectionDetailsStatus, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ObservedGeneration != nil {
		in, out := &in.ObservedGeneration, &out.ObservedGeneration
		*out = new(int)
		**out = **in
	}
	if in.UpdateTime != nil {
		in, out := &in.UpdateTime, &out.UpdateTime
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EdgeContainerVpnConnectionStatus.
func (in *EdgeContainerVpnConnectionStatus) DeepCopy() *EdgeContainerVpnConnectionStatus {
	if in == nil {
		return nil
	}
	out := new(EdgeContainerVpnConnectionStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodepoolLocalDiskEncryption) DeepCopyInto(out *NodepoolLocalDiskEncryption) {
	*out = *in
	if in.KmsKeyActiveVersion != nil {
		in, out := &in.KmsKeyActiveVersion, &out.KmsKeyActiveVersion
		*out = new(string)
		**out = **in
	}
	if in.KmsKeyRef != nil {
		in, out := &in.KmsKeyRef, &out.KmsKeyRef
		*out = new(v1alpha1.ResourceRef)
		**out = **in
	}
	if in.KmsKeyState != nil {
		in, out := &in.KmsKeyState, &out.KmsKeyState
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodepoolLocalDiskEncryption.
func (in *NodepoolLocalDiskEncryption) DeepCopy() *NodepoolLocalDiskEncryption {
	if in == nil {
		return nil
	}
	out := new(NodepoolLocalDiskEncryption)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodepoolNodeConfig) DeepCopyInto(out *NodepoolNodeConfig) {
	*out = *in
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodepoolNodeConfig.
func (in *NodepoolNodeConfig) DeepCopy() *NodepoolNodeConfig {
	if in == nil {
		return nil
	}
	out := new(NodepoolNodeConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VpnconnectionCloudRouterStatus) DeepCopyInto(out *VpnconnectionCloudRouterStatus) {
	*out = *in
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VpnconnectionCloudRouterStatus.
func (in *VpnconnectionCloudRouterStatus) DeepCopy() *VpnconnectionCloudRouterStatus {
	if in == nil {
		return nil
	}
	out := new(VpnconnectionCloudRouterStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VpnconnectionCloudVpnsStatus) DeepCopyInto(out *VpnconnectionCloudVpnsStatus) {
	*out = *in
	if in.Gateway != nil {
		in, out := &in.Gateway, &out.Gateway
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VpnconnectionCloudVpnsStatus.
func (in *VpnconnectionCloudVpnsStatus) DeepCopy() *VpnconnectionCloudVpnsStatus {
	if in == nil {
		return nil
	}
	out := new(VpnconnectionCloudVpnsStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VpnconnectionDetailsStatus) DeepCopyInto(out *VpnconnectionDetailsStatus) {
	*out = *in
	if in.CloudRouter != nil {
		in, out := &in.CloudRouter, &out.CloudRouter
		*out = make([]VpnconnectionCloudRouterStatus, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.CloudVpns != nil {
		in, out := &in.CloudVpns, &out.CloudVpns
		*out = make([]VpnconnectionCloudVpnsStatus, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Error != nil {
		in, out := &in.Error, &out.Error
		*out = new(string)
		**out = **in
	}
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VpnconnectionDetailsStatus.
func (in *VpnconnectionDetailsStatus) DeepCopy() *VpnconnectionDetailsStatus {
	if in == nil {
		return nil
	}
	out := new(VpnconnectionDetailsStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VpnconnectionVpcProject) DeepCopyInto(out *VpnconnectionVpcProject) {
	*out = *in
	if in.ProjectId != nil {
		in, out := &in.ProjectId, &out.ProjectId
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VpnconnectionVpcProject.
func (in *VpnconnectionVpcProject) DeepCopy() *VpnconnectionVpcProject {
	if in == nil {
		return nil
	}
	out := new(VpnconnectionVpcProject)
	in.DeepCopyInto(out)
	return out
}
