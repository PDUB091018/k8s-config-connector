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
func (in *InstanceActiveDirectoryConfig) DeepCopyInto(out *InstanceActiveDirectoryConfig) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InstanceActiveDirectoryConfig.
func (in *InstanceActiveDirectoryConfig) DeepCopy() *InstanceActiveDirectoryConfig {
	if in == nil {
		return nil
	}
	out := new(InstanceActiveDirectoryConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InstanceAuthorizedNetworks) DeepCopyInto(out *InstanceAuthorizedNetworks) {
	*out = *in
	if in.ExpirationTime != nil {
		in, out := &in.ExpirationTime, &out.ExpirationTime
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InstanceAuthorizedNetworks.
func (in *InstanceAuthorizedNetworks) DeepCopy() *InstanceAuthorizedNetworks {
	if in == nil {
		return nil
	}
	out := new(InstanceAuthorizedNetworks)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InstanceBackupConfiguration) DeepCopyInto(out *InstanceBackupConfiguration) {
	*out = *in
	if in.BackupRetentionSettings != nil {
		in, out := &in.BackupRetentionSettings, &out.BackupRetentionSettings
		*out = new(InstanceBackupRetentionSettings)
		(*in).DeepCopyInto(*out)
	}
	if in.BinaryLogEnabled != nil {
		in, out := &in.BinaryLogEnabled, &out.BinaryLogEnabled
		*out = new(bool)
		**out = **in
	}
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.Location != nil {
		in, out := &in.Location, &out.Location
		*out = new(string)
		**out = **in
	}
	if in.PointInTimeRecoveryEnabled != nil {
		in, out := &in.PointInTimeRecoveryEnabled, &out.PointInTimeRecoveryEnabled
		*out = new(bool)
		**out = **in
	}
	if in.StartTime != nil {
		in, out := &in.StartTime, &out.StartTime
		*out = new(string)
		**out = **in
	}
	if in.TransactionLogRetentionDays != nil {
		in, out := &in.TransactionLogRetentionDays, &out.TransactionLogRetentionDays
		*out = new(int)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InstanceBackupConfiguration.
func (in *InstanceBackupConfiguration) DeepCopy() *InstanceBackupConfiguration {
	if in == nil {
		return nil
	}
	out := new(InstanceBackupConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InstanceBackupRetentionSettings) DeepCopyInto(out *InstanceBackupRetentionSettings) {
	*out = *in
	if in.RetentionUnit != nil {
		in, out := &in.RetentionUnit, &out.RetentionUnit
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InstanceBackupRetentionSettings.
func (in *InstanceBackupRetentionSettings) DeepCopy() *InstanceBackupRetentionSettings {
	if in == nil {
		return nil
	}
	out := new(InstanceBackupRetentionSettings)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InstanceDatabaseFlags) DeepCopyInto(out *InstanceDatabaseFlags) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InstanceDatabaseFlags.
func (in *InstanceDatabaseFlags) DeepCopy() *InstanceDatabaseFlags {
	if in == nil {
		return nil
	}
	out := new(InstanceDatabaseFlags)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InstanceDenyMaintenancePeriod) DeepCopyInto(out *InstanceDenyMaintenancePeriod) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InstanceDenyMaintenancePeriod.
func (in *InstanceDenyMaintenancePeriod) DeepCopy() *InstanceDenyMaintenancePeriod {
	if in == nil {
		return nil
	}
	out := new(InstanceDenyMaintenancePeriod)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InstanceInsightsConfig) DeepCopyInto(out *InstanceInsightsConfig) {
	*out = *in
	if in.QueryInsightsEnabled != nil {
		in, out := &in.QueryInsightsEnabled, &out.QueryInsightsEnabled
		*out = new(bool)
		**out = **in
	}
	if in.QueryPlansPerMinute != nil {
		in, out := &in.QueryPlansPerMinute, &out.QueryPlansPerMinute
		*out = new(int)
		**out = **in
	}
	if in.QueryStringLength != nil {
		in, out := &in.QueryStringLength, &out.QueryStringLength
		*out = new(int)
		**out = **in
	}
	if in.RecordApplicationTags != nil {
		in, out := &in.RecordApplicationTags, &out.RecordApplicationTags
		*out = new(bool)
		**out = **in
	}
	if in.RecordClientAddress != nil {
		in, out := &in.RecordClientAddress, &out.RecordClientAddress
		*out = new(bool)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InstanceInsightsConfig.
func (in *InstanceInsightsConfig) DeepCopy() *InstanceInsightsConfig {
	if in == nil {
		return nil
	}
	out := new(InstanceInsightsConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InstanceIpAddressStatus) DeepCopyInto(out *InstanceIpAddressStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InstanceIpAddressStatus.
func (in *InstanceIpAddressStatus) DeepCopy() *InstanceIpAddressStatus {
	if in == nil {
		return nil
	}
	out := new(InstanceIpAddressStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InstanceIpConfiguration) DeepCopyInto(out *InstanceIpConfiguration) {
	*out = *in
	if in.AllocatedIpRange != nil {
		in, out := &in.AllocatedIpRange, &out.AllocatedIpRange
		*out = new(string)
		**out = **in
	}
	if in.AuthorizedNetworks != nil {
		in, out := &in.AuthorizedNetworks, &out.AuthorizedNetworks
		*out = make([]InstanceAuthorizedNetworks, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Ipv4Enabled != nil {
		in, out := &in.Ipv4Enabled, &out.Ipv4Enabled
		*out = new(bool)
		**out = **in
	}
	if in.PrivateNetworkRef != nil {
		in, out := &in.PrivateNetworkRef, &out.PrivateNetworkRef
		*out = new(v1alpha1.ResourceRef)
		**out = **in
	}
	if in.RequireSsl != nil {
		in, out := &in.RequireSsl, &out.RequireSsl
		*out = new(bool)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InstanceIpConfiguration.
func (in *InstanceIpConfiguration) DeepCopy() *InstanceIpConfiguration {
	if in == nil {
		return nil
	}
	out := new(InstanceIpConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InstanceLocationPreference) DeepCopyInto(out *InstanceLocationPreference) {
	*out = *in
	if in.FollowGaeApplication != nil {
		in, out := &in.FollowGaeApplication, &out.FollowGaeApplication
		*out = new(string)
		**out = **in
	}
	if in.SecondaryZone != nil {
		in, out := &in.SecondaryZone, &out.SecondaryZone
		*out = new(string)
		**out = **in
	}
	if in.Zone != nil {
		in, out := &in.Zone, &out.Zone
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InstanceLocationPreference.
func (in *InstanceLocationPreference) DeepCopy() *InstanceLocationPreference {
	if in == nil {
		return nil
	}
	out := new(InstanceLocationPreference)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InstanceMaintenanceWindow) DeepCopyInto(out *InstanceMaintenanceWindow) {
	*out = *in
	if in.Day != nil {
		in, out := &in.Day, &out.Day
		*out = new(int)
		**out = **in
	}
	if in.Hour != nil {
		in, out := &in.Hour, &out.Hour
		*out = new(int)
		**out = **in
	}
	if in.UpdateTrack != nil {
		in, out := &in.UpdateTrack, &out.UpdateTrack
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InstanceMaintenanceWindow.
func (in *InstanceMaintenanceWindow) DeepCopy() *InstanceMaintenanceWindow {
	if in == nil {
		return nil
	}
	out := new(InstanceMaintenanceWindow)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InstancePassword) DeepCopyInto(out *InstancePassword) {
	*out = *in
	if in.Value != nil {
		in, out := &in.Value, &out.Value
		*out = new(string)
		**out = **in
	}
	if in.ValueFrom != nil {
		in, out := &in.ValueFrom, &out.ValueFrom
		*out = new(InstanceValueFrom)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InstancePassword.
func (in *InstancePassword) DeepCopy() *InstancePassword {
	if in == nil {
		return nil
	}
	out := new(InstancePassword)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InstancePasswordValidationPolicy) DeepCopyInto(out *InstancePasswordValidationPolicy) {
	*out = *in
	if in.Complexity != nil {
		in, out := &in.Complexity, &out.Complexity
		*out = new(string)
		**out = **in
	}
	if in.DisallowUsernameSubstring != nil {
		in, out := &in.DisallowUsernameSubstring, &out.DisallowUsernameSubstring
		*out = new(bool)
		**out = **in
	}
	if in.MinLength != nil {
		in, out := &in.MinLength, &out.MinLength
		*out = new(int)
		**out = **in
	}
	if in.PasswordChangeInterval != nil {
		in, out := &in.PasswordChangeInterval, &out.PasswordChangeInterval
		*out = new(string)
		**out = **in
	}
	if in.ReuseInterval != nil {
		in, out := &in.ReuseInterval, &out.ReuseInterval
		*out = new(int)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InstancePasswordValidationPolicy.
func (in *InstancePasswordValidationPolicy) DeepCopy() *InstancePasswordValidationPolicy {
	if in == nil {
		return nil
	}
	out := new(InstancePasswordValidationPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InstanceReplicaConfiguration) DeepCopyInto(out *InstanceReplicaConfiguration) {
	*out = *in
	if in.CaCertificate != nil {
		in, out := &in.CaCertificate, &out.CaCertificate
		*out = new(string)
		**out = **in
	}
	if in.ClientCertificate != nil {
		in, out := &in.ClientCertificate, &out.ClientCertificate
		*out = new(string)
		**out = **in
	}
	if in.ClientKey != nil {
		in, out := &in.ClientKey, &out.ClientKey
		*out = new(string)
		**out = **in
	}
	if in.ConnectRetryInterval != nil {
		in, out := &in.ConnectRetryInterval, &out.ConnectRetryInterval
		*out = new(int)
		**out = **in
	}
	if in.DumpFilePath != nil {
		in, out := &in.DumpFilePath, &out.DumpFilePath
		*out = new(string)
		**out = **in
	}
	if in.FailoverTarget != nil {
		in, out := &in.FailoverTarget, &out.FailoverTarget
		*out = new(bool)
		**out = **in
	}
	if in.MasterHeartbeatPeriod != nil {
		in, out := &in.MasterHeartbeatPeriod, &out.MasterHeartbeatPeriod
		*out = new(int)
		**out = **in
	}
	if in.Password != nil {
		in, out := &in.Password, &out.Password
		*out = new(InstancePassword)
		(*in).DeepCopyInto(*out)
	}
	if in.SslCipher != nil {
		in, out := &in.SslCipher, &out.SslCipher
		*out = new(string)
		**out = **in
	}
	if in.Username != nil {
		in, out := &in.Username, &out.Username
		*out = new(string)
		**out = **in
	}
	if in.VerifyServerCertificate != nil {
		in, out := &in.VerifyServerCertificate, &out.VerifyServerCertificate
		*out = new(bool)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InstanceReplicaConfiguration.
func (in *InstanceReplicaConfiguration) DeepCopy() *InstanceReplicaConfiguration {
	if in == nil {
		return nil
	}
	out := new(InstanceReplicaConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InstanceRootPassword) DeepCopyInto(out *InstanceRootPassword) {
	*out = *in
	if in.Value != nil {
		in, out := &in.Value, &out.Value
		*out = new(string)
		**out = **in
	}
	if in.ValueFrom != nil {
		in, out := &in.ValueFrom, &out.ValueFrom
		*out = new(InstanceValueFrom)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InstanceRootPassword.
func (in *InstanceRootPassword) DeepCopy() *InstanceRootPassword {
	if in == nil {
		return nil
	}
	out := new(InstanceRootPassword)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InstanceServerCaCertStatus) DeepCopyInto(out *InstanceServerCaCertStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InstanceServerCaCertStatus.
func (in *InstanceServerCaCertStatus) DeepCopy() *InstanceServerCaCertStatus {
	if in == nil {
		return nil
	}
	out := new(InstanceServerCaCertStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InstanceSettings) DeepCopyInto(out *InstanceSettings) {
	*out = *in
	if in.ActivationPolicy != nil {
		in, out := &in.ActivationPolicy, &out.ActivationPolicy
		*out = new(string)
		**out = **in
	}
	if in.ActiveDirectoryConfig != nil {
		in, out := &in.ActiveDirectoryConfig, &out.ActiveDirectoryConfig
		*out = new(InstanceActiveDirectoryConfig)
		**out = **in
	}
	if in.AuthorizedGaeApplications != nil {
		in, out := &in.AuthorizedGaeApplications, &out.AuthorizedGaeApplications
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.AvailabilityType != nil {
		in, out := &in.AvailabilityType, &out.AvailabilityType
		*out = new(string)
		**out = **in
	}
	if in.BackupConfiguration != nil {
		in, out := &in.BackupConfiguration, &out.BackupConfiguration
		*out = new(InstanceBackupConfiguration)
		(*in).DeepCopyInto(*out)
	}
	if in.Collation != nil {
		in, out := &in.Collation, &out.Collation
		*out = new(string)
		**out = **in
	}
	if in.ConnectorEnforcement != nil {
		in, out := &in.ConnectorEnforcement, &out.ConnectorEnforcement
		*out = new(string)
		**out = **in
	}
	if in.CrashSafeReplication != nil {
		in, out := &in.CrashSafeReplication, &out.CrashSafeReplication
		*out = new(bool)
		**out = **in
	}
	if in.DatabaseFlags != nil {
		in, out := &in.DatabaseFlags, &out.DatabaseFlags
		*out = make([]InstanceDatabaseFlags, len(*in))
		copy(*out, *in)
	}
	if in.DeletionProtectionEnabled != nil {
		in, out := &in.DeletionProtectionEnabled, &out.DeletionProtectionEnabled
		*out = new(bool)
		**out = **in
	}
	if in.DenyMaintenancePeriod != nil {
		in, out := &in.DenyMaintenancePeriod, &out.DenyMaintenancePeriod
		*out = new(InstanceDenyMaintenancePeriod)
		**out = **in
	}
	if in.DiskAutoresize != nil {
		in, out := &in.DiskAutoresize, &out.DiskAutoresize
		*out = new(bool)
		**out = **in
	}
	if in.DiskAutoresizeLimit != nil {
		in, out := &in.DiskAutoresizeLimit, &out.DiskAutoresizeLimit
		*out = new(int)
		**out = **in
	}
	if in.DiskSize != nil {
		in, out := &in.DiskSize, &out.DiskSize
		*out = new(int)
		**out = **in
	}
	if in.DiskType != nil {
		in, out := &in.DiskType, &out.DiskType
		*out = new(string)
		**out = **in
	}
	if in.InsightsConfig != nil {
		in, out := &in.InsightsConfig, &out.InsightsConfig
		*out = new(InstanceInsightsConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.IpConfiguration != nil {
		in, out := &in.IpConfiguration, &out.IpConfiguration
		*out = new(InstanceIpConfiguration)
		(*in).DeepCopyInto(*out)
	}
	if in.LocationPreference != nil {
		in, out := &in.LocationPreference, &out.LocationPreference
		*out = new(InstanceLocationPreference)
		(*in).DeepCopyInto(*out)
	}
	if in.MaintenanceWindow != nil {
		in, out := &in.MaintenanceWindow, &out.MaintenanceWindow
		*out = new(InstanceMaintenanceWindow)
		(*in).DeepCopyInto(*out)
	}
	if in.PasswordValidationPolicy != nil {
		in, out := &in.PasswordValidationPolicy, &out.PasswordValidationPolicy
		*out = new(InstancePasswordValidationPolicy)
		(*in).DeepCopyInto(*out)
	}
	if in.PricingPlan != nil {
		in, out := &in.PricingPlan, &out.PricingPlan
		*out = new(string)
		**out = **in
	}
	if in.ReplicationType != nil {
		in, out := &in.ReplicationType, &out.ReplicationType
		*out = new(string)
		**out = **in
	}
	if in.SqlServerAuditConfig != nil {
		in, out := &in.SqlServerAuditConfig, &out.SqlServerAuditConfig
		*out = new(InstanceSqlServerAuditConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.TimeZone != nil {
		in, out := &in.TimeZone, &out.TimeZone
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InstanceSettings.
func (in *InstanceSettings) DeepCopy() *InstanceSettings {
	if in == nil {
		return nil
	}
	out := new(InstanceSettings)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InstanceSqlServerAuditConfig) DeepCopyInto(out *InstanceSqlServerAuditConfig) {
	*out = *in
	if in.BucketRef != nil {
		in, out := &in.BucketRef, &out.BucketRef
		*out = new(v1alpha1.ResourceRef)
		**out = **in
	}
	if in.RetentionInterval != nil {
		in, out := &in.RetentionInterval, &out.RetentionInterval
		*out = new(string)
		**out = **in
	}
	if in.UploadInterval != nil {
		in, out := &in.UploadInterval, &out.UploadInterval
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InstanceSqlServerAuditConfig.
func (in *InstanceSqlServerAuditConfig) DeepCopy() *InstanceSqlServerAuditConfig {
	if in == nil {
		return nil
	}
	out := new(InstanceSqlServerAuditConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InstanceValueFrom) DeepCopyInto(out *InstanceValueFrom) {
	*out = *in
	if in.SecretKeyRef != nil {
		in, out := &in.SecretKeyRef, &out.SecretKeyRef
		*out = new(v1alpha1.ResourceRef)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InstanceValueFrom.
func (in *InstanceValueFrom) DeepCopy() *InstanceValueFrom {
	if in == nil {
		return nil
	}
	out := new(InstanceValueFrom)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SQLDatabase) DeepCopyInto(out *SQLDatabase) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SQLDatabase.
func (in *SQLDatabase) DeepCopy() *SQLDatabase {
	if in == nil {
		return nil
	}
	out := new(SQLDatabase)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SQLDatabase) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SQLDatabaseList) DeepCopyInto(out *SQLDatabaseList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]SQLDatabase, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SQLDatabaseList.
func (in *SQLDatabaseList) DeepCopy() *SQLDatabaseList {
	if in == nil {
		return nil
	}
	out := new(SQLDatabaseList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SQLDatabaseList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SQLDatabaseSpec) DeepCopyInto(out *SQLDatabaseSpec) {
	*out = *in
	if in.Charset != nil {
		in, out := &in.Charset, &out.Charset
		*out = new(string)
		**out = **in
	}
	if in.Collation != nil {
		in, out := &in.Collation, &out.Collation
		*out = new(string)
		**out = **in
	}
	if in.DeletionPolicy != nil {
		in, out := &in.DeletionPolicy, &out.DeletionPolicy
		*out = new(string)
		**out = **in
	}
	out.InstanceRef = in.InstanceRef
	if in.ResourceID != nil {
		in, out := &in.ResourceID, &out.ResourceID
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SQLDatabaseSpec.
func (in *SQLDatabaseSpec) DeepCopy() *SQLDatabaseSpec {
	if in == nil {
		return nil
	}
	out := new(SQLDatabaseSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SQLDatabaseStatus) DeepCopyInto(out *SQLDatabaseStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1alpha1.Condition, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SQLDatabaseStatus.
func (in *SQLDatabaseStatus) DeepCopy() *SQLDatabaseStatus {
	if in == nil {
		return nil
	}
	out := new(SQLDatabaseStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SQLInstance) DeepCopyInto(out *SQLInstance) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SQLInstance.
func (in *SQLInstance) DeepCopy() *SQLInstance {
	if in == nil {
		return nil
	}
	out := new(SQLInstance)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SQLInstance) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SQLInstanceList) DeepCopyInto(out *SQLInstanceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]SQLInstance, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SQLInstanceList.
func (in *SQLInstanceList) DeepCopy() *SQLInstanceList {
	if in == nil {
		return nil
	}
	out := new(SQLInstanceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SQLInstanceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SQLInstanceSpec) DeepCopyInto(out *SQLInstanceSpec) {
	*out = *in
	if in.DatabaseVersion != nil {
		in, out := &in.DatabaseVersion, &out.DatabaseVersion
		*out = new(string)
		**out = **in
	}
	if in.EncryptionKMSCryptoKeyRef != nil {
		in, out := &in.EncryptionKMSCryptoKeyRef, &out.EncryptionKMSCryptoKeyRef
		*out = new(v1alpha1.ResourceRef)
		**out = **in
	}
	if in.MaintenanceVersion != nil {
		in, out := &in.MaintenanceVersion, &out.MaintenanceVersion
		*out = new(string)
		**out = **in
	}
	if in.MasterInstanceRef != nil {
		in, out := &in.MasterInstanceRef, &out.MasterInstanceRef
		*out = new(v1alpha1.ResourceRef)
		**out = **in
	}
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(string)
		**out = **in
	}
	if in.ReplicaConfiguration != nil {
		in, out := &in.ReplicaConfiguration, &out.ReplicaConfiguration
		*out = new(InstanceReplicaConfiguration)
		(*in).DeepCopyInto(*out)
	}
	if in.ResourceID != nil {
		in, out := &in.ResourceID, &out.ResourceID
		*out = new(string)
		**out = **in
	}
	if in.RootPassword != nil {
		in, out := &in.RootPassword, &out.RootPassword
		*out = new(InstanceRootPassword)
		(*in).DeepCopyInto(*out)
	}
	in.Settings.DeepCopyInto(&out.Settings)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SQLInstanceSpec.
func (in *SQLInstanceSpec) DeepCopy() *SQLInstanceSpec {
	if in == nil {
		return nil
	}
	out := new(SQLInstanceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SQLInstanceStatus) DeepCopyInto(out *SQLInstanceStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1alpha1.Condition, len(*in))
		copy(*out, *in)
	}
	if in.AvailableMaintenanceVersions != nil {
		in, out := &in.AvailableMaintenanceVersions, &out.AvailableMaintenanceVersions
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.IpAddress != nil {
		in, out := &in.IpAddress, &out.IpAddress
		*out = make([]InstanceIpAddressStatus, len(*in))
		copy(*out, *in)
	}
	out.ServerCaCert = in.ServerCaCert
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SQLInstanceStatus.
func (in *SQLInstanceStatus) DeepCopy() *SQLInstanceStatus {
	if in == nil {
		return nil
	}
	out := new(SQLInstanceStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SQLSSLCert) DeepCopyInto(out *SQLSSLCert) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SQLSSLCert.
func (in *SQLSSLCert) DeepCopy() *SQLSSLCert {
	if in == nil {
		return nil
	}
	out := new(SQLSSLCert)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SQLSSLCert) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SQLSSLCertList) DeepCopyInto(out *SQLSSLCertList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]SQLSSLCert, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SQLSSLCertList.
func (in *SQLSSLCertList) DeepCopy() *SQLSSLCertList {
	if in == nil {
		return nil
	}
	out := new(SQLSSLCertList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SQLSSLCertList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SQLSSLCertSpec) DeepCopyInto(out *SQLSSLCertSpec) {
	*out = *in
	out.InstanceRef = in.InstanceRef
	if in.ResourceID != nil {
		in, out := &in.ResourceID, &out.ResourceID
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SQLSSLCertSpec.
func (in *SQLSSLCertSpec) DeepCopy() *SQLSSLCertSpec {
	if in == nil {
		return nil
	}
	out := new(SQLSSLCertSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SQLSSLCertStatus) DeepCopyInto(out *SQLSSLCertStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1alpha1.Condition, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SQLSSLCertStatus.
func (in *SQLSSLCertStatus) DeepCopy() *SQLSSLCertStatus {
	if in == nil {
		return nil
	}
	out := new(SQLSSLCertStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SQLUser) DeepCopyInto(out *SQLUser) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SQLUser.
func (in *SQLUser) DeepCopy() *SQLUser {
	if in == nil {
		return nil
	}
	out := new(SQLUser)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SQLUser) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SQLUserList) DeepCopyInto(out *SQLUserList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]SQLUser, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SQLUserList.
func (in *SQLUserList) DeepCopy() *SQLUserList {
	if in == nil {
		return nil
	}
	out := new(SQLUserList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SQLUserList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SQLUserSpec) DeepCopyInto(out *SQLUserSpec) {
	*out = *in
	if in.Host != nil {
		in, out := &in.Host, &out.Host
		*out = new(string)
		**out = **in
	}
	out.InstanceRef = in.InstanceRef
	if in.Password != nil {
		in, out := &in.Password, &out.Password
		*out = new(UserPassword)
		(*in).DeepCopyInto(*out)
	}
	if in.PasswordPolicy != nil {
		in, out := &in.PasswordPolicy, &out.PasswordPolicy
		*out = new(UserPasswordPolicy)
		(*in).DeepCopyInto(*out)
	}
	if in.ResourceID != nil {
		in, out := &in.ResourceID, &out.ResourceID
		*out = new(string)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SQLUserSpec.
func (in *SQLUserSpec) DeepCopy() *SQLUserSpec {
	if in == nil {
		return nil
	}
	out := new(SQLUserSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SQLUserStatus) DeepCopyInto(out *SQLUserStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1alpha1.Condition, len(*in))
		copy(*out, *in)
	}
	if in.SqlServerUserDetails != nil {
		in, out := &in.SqlServerUserDetails, &out.SqlServerUserDetails
		*out = make([]UserSqlServerUserDetailsStatus, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SQLUserStatus.
func (in *SQLUserStatus) DeepCopy() *SQLUserStatus {
	if in == nil {
		return nil
	}
	out := new(SQLUserStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UserPassword) DeepCopyInto(out *UserPassword) {
	*out = *in
	if in.Value != nil {
		in, out := &in.Value, &out.Value
		*out = new(string)
		**out = **in
	}
	if in.ValueFrom != nil {
		in, out := &in.ValueFrom, &out.ValueFrom
		*out = new(UserValueFrom)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UserPassword.
func (in *UserPassword) DeepCopy() *UserPassword {
	if in == nil {
		return nil
	}
	out := new(UserPassword)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UserPasswordPolicy) DeepCopyInto(out *UserPasswordPolicy) {
	*out = *in
	if in.AllowedFailedAttempts != nil {
		in, out := &in.AllowedFailedAttempts, &out.AllowedFailedAttempts
		*out = new(int)
		**out = **in
	}
	if in.EnableFailedAttemptsCheck != nil {
		in, out := &in.EnableFailedAttemptsCheck, &out.EnableFailedAttemptsCheck
		*out = new(bool)
		**out = **in
	}
	if in.EnablePasswordVerification != nil {
		in, out := &in.EnablePasswordVerification, &out.EnablePasswordVerification
		*out = new(bool)
		**out = **in
	}
	if in.PasswordExpirationDuration != nil {
		in, out := &in.PasswordExpirationDuration, &out.PasswordExpirationDuration
		*out = new(string)
		**out = **in
	}
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = make([]UserStatus, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UserPasswordPolicy.
func (in *UserPasswordPolicy) DeepCopy() *UserPasswordPolicy {
	if in == nil {
		return nil
	}
	out := new(UserPasswordPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UserSqlServerUserDetailsStatus) DeepCopyInto(out *UserSqlServerUserDetailsStatus) {
	*out = *in
	if in.ServerRoles != nil {
		in, out := &in.ServerRoles, &out.ServerRoles
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UserSqlServerUserDetailsStatus.
func (in *UserSqlServerUserDetailsStatus) DeepCopy() *UserSqlServerUserDetailsStatus {
	if in == nil {
		return nil
	}
	out := new(UserSqlServerUserDetailsStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UserStatus) DeepCopyInto(out *UserStatus) {
	*out = *in
	if in.Locked != nil {
		in, out := &in.Locked, &out.Locked
		*out = new(bool)
		**out = **in
	}
	if in.PasswordExpirationTime != nil {
		in, out := &in.PasswordExpirationTime, &out.PasswordExpirationTime
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UserStatus.
func (in *UserStatus) DeepCopy() *UserStatus {
	if in == nil {
		return nil
	}
	out := new(UserStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UserValueFrom) DeepCopyInto(out *UserValueFrom) {
	*out = *in
	if in.SecretKeyRef != nil {
		in, out := &in.SecretKeyRef, &out.SecretKeyRef
		*out = new(v1alpha1.ResourceRef)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UserValueFrom.
func (in *UserValueFrom) DeepCopy() *UserValueFrom {
	if in == nil {
		return nil
	}
	out := new(UserValueFrom)
	in.DeepCopyInto(out)
	return out
}
