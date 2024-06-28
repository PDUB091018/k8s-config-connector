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
// Code generated by deepcopy-gen. DO NOT EDIT.

package v1beta1

import (
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	v1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/k8s/v1alpha1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Aggregation) DeepCopyInto(out *Aggregation) {
	*out = *in
	if in.AlignmentPeriod != nil {
		in, out := &in.AlignmentPeriod, &out.AlignmentPeriod
		*out = new(string)
		**out = **in
	}
	if in.PerSeriesAligner != nil {
		in, out := &in.PerSeriesAligner, &out.PerSeriesAligner
		*out = new(string)
		**out = **in
	}
	if in.CrossSeriesReducer != nil {
		in, out := &in.CrossSeriesReducer, &out.CrossSeriesReducer
		*out = new(string)
		**out = **in
	}
	if in.GroupByFields != nil {
		in, out := &in.GroupByFields, &out.GroupByFields
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Aggregation.
func (in *Aggregation) DeepCopy() *Aggregation {
	if in == nil {
		return nil
	}
	out := new(Aggregation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AlertChart) DeepCopyInto(out *AlertChart) {
	*out = *in
	if in.AlertPolicyRef != nil {
		in, out := &in.AlertPolicyRef, &out.AlertPolicyRef
		*out = new(MonitoringAlertPolicyRef)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AlertChart.
func (in *AlertChart) DeepCopy() *AlertChart {
	if in == nil {
		return nil
	}
	out := new(AlertChart)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ChartOptions) DeepCopyInto(out *ChartOptions) {
	*out = *in
	if in.Mode != nil {
		in, out := &in.Mode, &out.Mode
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ChartOptions.
func (in *ChartOptions) DeepCopy() *ChartOptions {
	if in == nil {
		return nil
	}
	out := new(ChartOptions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CollapsibleGroup) DeepCopyInto(out *CollapsibleGroup) {
	*out = *in
	if in.Collapsed != nil {
		in, out := &in.Collapsed, &out.Collapsed
		*out = new(bool)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CollapsibleGroup.
func (in *CollapsibleGroup) DeepCopy() *CollapsibleGroup {
	if in == nil {
		return nil
	}
	out := new(CollapsibleGroup)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ColumnLayout) DeepCopyInto(out *ColumnLayout) {
	*out = *in
	if in.Columns != nil {
		in, out := &in.Columns, &out.Columns
		*out = make([]ColumnLayout_Column, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ColumnLayout.
func (in *ColumnLayout) DeepCopy() *ColumnLayout {
	if in == nil {
		return nil
	}
	out := new(ColumnLayout)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ColumnLayout_Column) DeepCopyInto(out *ColumnLayout_Column) {
	*out = *in
	if in.Weight != nil {
		in, out := &in.Weight, &out.Weight
		*out = new(int64)
		**out = **in
	}
	if in.Widgets != nil {
		in, out := &in.Widgets, &out.Widgets
		*out = make([]Widget, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ColumnLayout_Column.
func (in *ColumnLayout_Column) DeepCopy() *ColumnLayout_Column {
	if in == nil {
		return nil
	}
	out := new(ColumnLayout_Column)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DashboardFilter) DeepCopyInto(out *DashboardFilter) {
	*out = *in
	if in.LabelKey != nil {
		in, out := &in.LabelKey, &out.LabelKey
		*out = new(string)
		**out = **in
	}
	if in.TemplateVariable != nil {
		in, out := &in.TemplateVariable, &out.TemplateVariable
		*out = new(string)
		**out = **in
	}
	if in.StringValue != nil {
		in, out := &in.StringValue, &out.StringValue
		*out = new(string)
		**out = **in
	}
	if in.FilterType != nil {
		in, out := &in.FilterType, &out.FilterType
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DashboardFilter.
func (in *DashboardFilter) DeepCopy() *DashboardFilter {
	if in == nil {
		return nil
	}
	out := new(DashboardFilter)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Empty) DeepCopyInto(out *Empty) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Empty.
func (in *Empty) DeepCopy() *Empty {
	if in == nil {
		return nil
	}
	out := new(Empty)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ErrorReportingPanel) DeepCopyInto(out *ErrorReportingPanel) {
	*out = *in
	if in.ProjectRefs != nil {
		in, out := &in.ProjectRefs, &out.ProjectRefs
		*out = make([]refsv1beta1.ProjectRef, len(*in))
		copy(*out, *in)
	}
	if in.Services != nil {
		in, out := &in.Services, &out.Services
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Versions != nil {
		in, out := &in.Versions, &out.Versions
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ErrorReportingPanel.
func (in *ErrorReportingPanel) DeepCopy() *ErrorReportingPanel {
	if in == nil {
		return nil
	}
	out := new(ErrorReportingPanel)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GridLayout) DeepCopyInto(out *GridLayout) {
	*out = *in
	if in.Columns != nil {
		in, out := &in.Columns, &out.Columns
		*out = new(int64)
		**out = **in
	}
	if in.Widgets != nil {
		in, out := &in.Widgets, &out.Widgets
		*out = make([]Widget, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GridLayout.
func (in *GridLayout) DeepCopy() *GridLayout {
	if in == nil {
		return nil
	}
	out := new(GridLayout)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IncidentList) DeepCopyInto(out *IncidentList) {
	*out = *in
	if in.MonitoredResources != nil {
		in, out := &in.MonitoredResources, &out.MonitoredResources
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.PolicyNames != nil {
		in, out := &in.PolicyNames, &out.PolicyNames
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IncidentList.
func (in *IncidentList) DeepCopy() *IncidentList {
	if in == nil {
		return nil
	}
	out := new(IncidentList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LogsPanel) DeepCopyInto(out *LogsPanel) {
	*out = *in
	if in.Filter != nil {
		in, out := &in.Filter, &out.Filter
		*out = new(string)
		**out = **in
	}
	if in.ResourceNames != nil {
		in, out := &in.ResourceNames, &out.ResourceNames
		*out = make([]v1alpha1.ResourceRef, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LogsPanel.
func (in *LogsPanel) DeepCopy() *LogsPanel {
	if in == nil {
		return nil
	}
	out := new(LogsPanel)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MonitoringAlertPolicyRef) DeepCopyInto(out *MonitoringAlertPolicyRef) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MonitoringAlertPolicyRef.
func (in *MonitoringAlertPolicyRef) DeepCopy() *MonitoringAlertPolicyRef {
	if in == nil {
		return nil
	}
	out := new(MonitoringAlertPolicyRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MonitoringDashboard) DeepCopyInto(out *MonitoringDashboard) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MonitoringDashboard.
func (in *MonitoringDashboard) DeepCopy() *MonitoringDashboard {
	if in == nil {
		return nil
	}
	out := new(MonitoringDashboard)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MonitoringDashboard) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MonitoringDashboardList) DeepCopyInto(out *MonitoringDashboardList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]MonitoringDashboard, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MonitoringDashboardList.
func (in *MonitoringDashboardList) DeepCopy() *MonitoringDashboardList {
	if in == nil {
		return nil
	}
	out := new(MonitoringDashboardList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MonitoringDashboardList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MonitoringDashboardSpec) DeepCopyInto(out *MonitoringDashboardSpec) {
	*out = *in
	out.ProjectRef = in.ProjectRef
	if in.ResourceID != nil {
		in, out := &in.ResourceID, &out.ResourceID
		*out = new(string)
		**out = **in
	}
	if in.DisplayName != nil {
		in, out := &in.DisplayName, &out.DisplayName
		*out = new(string)
		**out = **in
	}
	if in.GridLayout != nil {
		in, out := &in.GridLayout, &out.GridLayout
		*out = new(GridLayout)
		(*in).DeepCopyInto(*out)
	}
	if in.MosaicLayout != nil {
		in, out := &in.MosaicLayout, &out.MosaicLayout
		*out = new(MosaicLayout)
		(*in).DeepCopyInto(*out)
	}
	if in.RowLayout != nil {
		in, out := &in.RowLayout, &out.RowLayout
		*out = new(RowLayout)
		(*in).DeepCopyInto(*out)
	}
	if in.ColumnLayout != nil {
		in, out := &in.ColumnLayout, &out.ColumnLayout
		*out = new(ColumnLayout)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MonitoringDashboardSpec.
func (in *MonitoringDashboardSpec) DeepCopy() *MonitoringDashboardSpec {
	if in == nil {
		return nil
	}
	out := new(MonitoringDashboardSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MonitoringDashboardStatus) DeepCopyInto(out *MonitoringDashboardStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1alpha1.Condition, len(*in))
		copy(*out, *in)
	}
	if in.Etag != nil {
		in, out := &in.Etag, &out.Etag
		*out = new(string)
		**out = **in
	}
	if in.ObservedGeneration != nil {
		in, out := &in.ObservedGeneration, &out.ObservedGeneration
		*out = new(int64)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MonitoringDashboardStatus.
func (in *MonitoringDashboardStatus) DeepCopy() *MonitoringDashboardStatus {
	if in == nil {
		return nil
	}
	out := new(MonitoringDashboardStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MosaicLayout) DeepCopyInto(out *MosaicLayout) {
	*out = *in
	if in.Columns != nil {
		in, out := &in.Columns, &out.Columns
		*out = new(int32)
		**out = **in
	}
	if in.Tiles != nil {
		in, out := &in.Tiles, &out.Tiles
		*out = make([]MosaicLayout_Tile, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MosaicLayout.
func (in *MosaicLayout) DeepCopy() *MosaicLayout {
	if in == nil {
		return nil
	}
	out := new(MosaicLayout)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MosaicLayout_Tile) DeepCopyInto(out *MosaicLayout_Tile) {
	*out = *in
	if in.XPos != nil {
		in, out := &in.XPos, &out.XPos
		*out = new(int32)
		**out = **in
	}
	if in.YPos != nil {
		in, out := &in.YPos, &out.YPos
		*out = new(int32)
		**out = **in
	}
	if in.Width != nil {
		in, out := &in.Width, &out.Width
		*out = new(int32)
		**out = **in
	}
	if in.Height != nil {
		in, out := &in.Height, &out.Height
		*out = new(int32)
		**out = **in
	}
	if in.Widget != nil {
		in, out := &in.Widget, &out.Widget
		*out = new(Widget)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MosaicLayout_Tile.
func (in *MosaicLayout_Tile) DeepCopy() *MosaicLayout_Tile {
	if in == nil {
		return nil
	}
	out := new(MosaicLayout_Tile)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PickTimeSeriesFilter) DeepCopyInto(out *PickTimeSeriesFilter) {
	*out = *in
	if in.RankingMethod != nil {
		in, out := &in.RankingMethod, &out.RankingMethod
		*out = new(string)
		**out = **in
	}
	if in.NumTimeSeries != nil {
		in, out := &in.NumTimeSeries, &out.NumTimeSeries
		*out = new(int32)
		**out = **in
	}
	if in.Direction != nil {
		in, out := &in.Direction, &out.Direction
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PickTimeSeriesFilter.
func (in *PickTimeSeriesFilter) DeepCopy() *PickTimeSeriesFilter {
	if in == nil {
		return nil
	}
	out := new(PickTimeSeriesFilter)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PieChart) DeepCopyInto(out *PieChart) {
	*out = *in
	if in.DataSets != nil {
		in, out := &in.DataSets, &out.DataSets
		*out = make([]PieChart_PieChartDataSet, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ChartType != nil {
		in, out := &in.ChartType, &out.ChartType
		*out = new(string)
		**out = **in
	}
	if in.ShowLabels != nil {
		in, out := &in.ShowLabels, &out.ShowLabels
		*out = new(bool)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PieChart.
func (in *PieChart) DeepCopy() *PieChart {
	if in == nil {
		return nil
	}
	out := new(PieChart)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PieChart_PieChartDataSet) DeepCopyInto(out *PieChart_PieChartDataSet) {
	*out = *in
	if in.TimeSeriesQuery != nil {
		in, out := &in.TimeSeriesQuery, &out.TimeSeriesQuery
		*out = new(TimeSeriesQuery)
		(*in).DeepCopyInto(*out)
	}
	if in.SliceNameTemplate != nil {
		in, out := &in.SliceNameTemplate, &out.SliceNameTemplate
		*out = new(string)
		**out = **in
	}
	if in.MinAlignmentPeriod != nil {
		in, out := &in.MinAlignmentPeriod, &out.MinAlignmentPeriod
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PieChart_PieChartDataSet.
func (in *PieChart_PieChartDataSet) DeepCopy() *PieChart_PieChartDataSet {
	if in == nil {
		return nil
	}
	out := new(PieChart_PieChartDataSet)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RowLayout) DeepCopyInto(out *RowLayout) {
	*out = *in
	if in.Rows != nil {
		in, out := &in.Rows, &out.Rows
		*out = make([]RowLayout_Row, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RowLayout.
func (in *RowLayout) DeepCopy() *RowLayout {
	if in == nil {
		return nil
	}
	out := new(RowLayout)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RowLayout_Row) DeepCopyInto(out *RowLayout_Row) {
	*out = *in
	if in.Weight != nil {
		in, out := &in.Weight, &out.Weight
		*out = new(int64)
		**out = **in
	}
	if in.Widgets != nil {
		in, out := &in.Widgets, &out.Widgets
		*out = make([]Widget, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RowLayout_Row.
func (in *RowLayout_Row) DeepCopy() *RowLayout_Row {
	if in == nil {
		return nil
	}
	out := new(RowLayout_Row)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Scorecard) DeepCopyInto(out *Scorecard) {
	*out = *in
	if in.TimeSeriesQuery != nil {
		in, out := &in.TimeSeriesQuery, &out.TimeSeriesQuery
		*out = new(TimeSeriesQuery)
		(*in).DeepCopyInto(*out)
	}
	if in.GaugeView != nil {
		in, out := &in.GaugeView, &out.GaugeView
		*out = new(Scorecard_GaugeView)
		(*in).DeepCopyInto(*out)
	}
	if in.SparkChartView != nil {
		in, out := &in.SparkChartView, &out.SparkChartView
		*out = new(Scorecard_SparkChartView)
		(*in).DeepCopyInto(*out)
	}
	if in.Thresholds != nil {
		in, out := &in.Thresholds, &out.Thresholds
		*out = make([]Threshold, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Scorecard.
func (in *Scorecard) DeepCopy() *Scorecard {
	if in == nil {
		return nil
	}
	out := new(Scorecard)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Scorecard_GaugeView) DeepCopyInto(out *Scorecard_GaugeView) {
	*out = *in
	if in.LowerBound != nil {
		in, out := &in.LowerBound, &out.LowerBound
		*out = new(float64)
		**out = **in
	}
	if in.UpperBound != nil {
		in, out := &in.UpperBound, &out.UpperBound
		*out = new(float64)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Scorecard_GaugeView.
func (in *Scorecard_GaugeView) DeepCopy() *Scorecard_GaugeView {
	if in == nil {
		return nil
	}
	out := new(Scorecard_GaugeView)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Scorecard_SparkChartView) DeepCopyInto(out *Scorecard_SparkChartView) {
	*out = *in
	if in.SparkChartType != nil {
		in, out := &in.SparkChartType, &out.SparkChartType
		*out = new(string)
		**out = **in
	}
	if in.MinAlignmentPeriod != nil {
		in, out := &in.MinAlignmentPeriod, &out.MinAlignmentPeriod
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Scorecard_SparkChartView.
func (in *Scorecard_SparkChartView) DeepCopy() *Scorecard_SparkChartView {
	if in == nil {
		return nil
	}
	out := new(Scorecard_SparkChartView)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SectionHeader) DeepCopyInto(out *SectionHeader) {
	*out = *in
	if in.Subtitle != nil {
		in, out := &in.Subtitle, &out.Subtitle
		*out = new(string)
		**out = **in
	}
	if in.DividerBelow != nil {
		in, out := &in.DividerBelow, &out.DividerBelow
		*out = new(bool)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SectionHeader.
func (in *SectionHeader) DeepCopy() *SectionHeader {
	if in == nil {
		return nil
	}
	out := new(SectionHeader)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SingleViewGroup) DeepCopyInto(out *SingleViewGroup) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SingleViewGroup.
func (in *SingleViewGroup) DeepCopy() *SingleViewGroup {
	if in == nil {
		return nil
	}
	out := new(SingleViewGroup)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StatisticalTimeSeriesFilter) DeepCopyInto(out *StatisticalTimeSeriesFilter) {
	*out = *in
	if in.RankingMethod != nil {
		in, out := &in.RankingMethod, &out.RankingMethod
		*out = new(string)
		**out = **in
	}
	if in.NumTimeSeries != nil {
		in, out := &in.NumTimeSeries, &out.NumTimeSeries
		*out = new(int32)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StatisticalTimeSeriesFilter.
func (in *StatisticalTimeSeriesFilter) DeepCopy() *StatisticalTimeSeriesFilter {
	if in == nil {
		return nil
	}
	out := new(StatisticalTimeSeriesFilter)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TableDisplayOptions) DeepCopyInto(out *TableDisplayOptions) {
	*out = *in
	if in.ShownColumns != nil {
		in, out := &in.ShownColumns, &out.ShownColumns
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TableDisplayOptions.
func (in *TableDisplayOptions) DeepCopy() *TableDisplayOptions {
	if in == nil {
		return nil
	}
	out := new(TableDisplayOptions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Text) DeepCopyInto(out *Text) {
	*out = *in
	if in.Content != nil {
		in, out := &in.Content, &out.Content
		*out = new(string)
		**out = **in
	}
	if in.Format != nil {
		in, out := &in.Format, &out.Format
		*out = new(string)
		**out = **in
	}
	if in.Style != nil {
		in, out := &in.Style, &out.Style
		*out = new(Text_TextStyle)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Text.
func (in *Text) DeepCopy() *Text {
	if in == nil {
		return nil
	}
	out := new(Text)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Text_TextStyle) DeepCopyInto(out *Text_TextStyle) {
	*out = *in
	if in.BackgroundColor != nil {
		in, out := &in.BackgroundColor, &out.BackgroundColor
		*out = new(string)
		**out = **in
	}
	if in.TextColor != nil {
		in, out := &in.TextColor, &out.TextColor
		*out = new(string)
		**out = **in
	}
	if in.HorizontalAlignment != nil {
		in, out := &in.HorizontalAlignment, &out.HorizontalAlignment
		*out = new(string)
		**out = **in
	}
	if in.VerticalAlignment != nil {
		in, out := &in.VerticalAlignment, &out.VerticalAlignment
		*out = new(string)
		**out = **in
	}
	if in.Padding != nil {
		in, out := &in.Padding, &out.Padding
		*out = new(string)
		**out = **in
	}
	if in.FontSize != nil {
		in, out := &in.FontSize, &out.FontSize
		*out = new(string)
		**out = **in
	}
	if in.PointerLocation != nil {
		in, out := &in.PointerLocation, &out.PointerLocation
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Text_TextStyle.
func (in *Text_TextStyle) DeepCopy() *Text_TextStyle {
	if in == nil {
		return nil
	}
	out := new(Text_TextStyle)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Threshold) DeepCopyInto(out *Threshold) {
	*out = *in
	if in.Label != nil {
		in, out := &in.Label, &out.Label
		*out = new(string)
		**out = **in
	}
	if in.Value != nil {
		in, out := &in.Value, &out.Value
		*out = new(float64)
		**out = **in
	}
	if in.Color != nil {
		in, out := &in.Color, &out.Color
		*out = new(string)
		**out = **in
	}
	if in.Direction != nil {
		in, out := &in.Direction, &out.Direction
		*out = new(string)
		**out = **in
	}
	if in.TargetAxis != nil {
		in, out := &in.TargetAxis, &out.TargetAxis
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Threshold.
func (in *Threshold) DeepCopy() *Threshold {
	if in == nil {
		return nil
	}
	out := new(Threshold)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TimeSeriesFilter) DeepCopyInto(out *TimeSeriesFilter) {
	*out = *in
	if in.Filter != nil {
		in, out := &in.Filter, &out.Filter
		*out = new(string)
		**out = **in
	}
	if in.Aggregation != nil {
		in, out := &in.Aggregation, &out.Aggregation
		*out = new(Aggregation)
		(*in).DeepCopyInto(*out)
	}
	if in.SecondaryAggregation != nil {
		in, out := &in.SecondaryAggregation, &out.SecondaryAggregation
		*out = new(Aggregation)
		(*in).DeepCopyInto(*out)
	}
	if in.PickTimeSeriesFilter != nil {
		in, out := &in.PickTimeSeriesFilter, &out.PickTimeSeriesFilter
		*out = new(PickTimeSeriesFilter)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TimeSeriesFilter.
func (in *TimeSeriesFilter) DeepCopy() *TimeSeriesFilter {
	if in == nil {
		return nil
	}
	out := new(TimeSeriesFilter)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TimeSeriesFilterRatio) DeepCopyInto(out *TimeSeriesFilterRatio) {
	*out = *in
	if in.Numerator != nil {
		in, out := &in.Numerator, &out.Numerator
		*out = new(TimeSeriesFilterRatio_RatioPart)
		(*in).DeepCopyInto(*out)
	}
	if in.Denominator != nil {
		in, out := &in.Denominator, &out.Denominator
		*out = new(TimeSeriesFilterRatio_RatioPart)
		(*in).DeepCopyInto(*out)
	}
	if in.SecondaryAggregation != nil {
		in, out := &in.SecondaryAggregation, &out.SecondaryAggregation
		*out = new(Aggregation)
		(*in).DeepCopyInto(*out)
	}
	if in.PickTimeSeriesFilter != nil {
		in, out := &in.PickTimeSeriesFilter, &out.PickTimeSeriesFilter
		*out = new(PickTimeSeriesFilter)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TimeSeriesFilterRatio.
func (in *TimeSeriesFilterRatio) DeepCopy() *TimeSeriesFilterRatio {
	if in == nil {
		return nil
	}
	out := new(TimeSeriesFilterRatio)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TimeSeriesFilterRatio_RatioPart) DeepCopyInto(out *TimeSeriesFilterRatio_RatioPart) {
	*out = *in
	if in.Filter != nil {
		in, out := &in.Filter, &out.Filter
		*out = new(string)
		**out = **in
	}
	if in.Aggregation != nil {
		in, out := &in.Aggregation, &out.Aggregation
		*out = new(Aggregation)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TimeSeriesFilterRatio_RatioPart.
func (in *TimeSeriesFilterRatio_RatioPart) DeepCopy() *TimeSeriesFilterRatio_RatioPart {
	if in == nil {
		return nil
	}
	out := new(TimeSeriesFilterRatio_RatioPart)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TimeSeriesQuery) DeepCopyInto(out *TimeSeriesQuery) {
	*out = *in
	if in.TimeSeriesFilter != nil {
		in, out := &in.TimeSeriesFilter, &out.TimeSeriesFilter
		*out = new(TimeSeriesFilter)
		(*in).DeepCopyInto(*out)
	}
	if in.TimeSeriesFilterRatio != nil {
		in, out := &in.TimeSeriesFilterRatio, &out.TimeSeriesFilterRatio
		*out = new(TimeSeriesFilterRatio)
		(*in).DeepCopyInto(*out)
	}
	if in.TimeSeriesQueryLanguage != nil {
		in, out := &in.TimeSeriesQueryLanguage, &out.TimeSeriesQueryLanguage
		*out = new(string)
		**out = **in
	}
	if in.UnitOverride != nil {
		in, out := &in.UnitOverride, &out.UnitOverride
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TimeSeriesQuery.
func (in *TimeSeriesQuery) DeepCopy() *TimeSeriesQuery {
	if in == nil {
		return nil
	}
	out := new(TimeSeriesQuery)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TimeSeriesTable) DeepCopyInto(out *TimeSeriesTable) {
	*out = *in
	if in.DataSets != nil {
		in, out := &in.DataSets, &out.DataSets
		*out = make([]TimeSeriesTable_TableDataSet, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ColumnSettings != nil {
		in, out := &in.ColumnSettings, &out.ColumnSettings
		*out = make([]TimeSeriesTable_ColumnSettings, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TimeSeriesTable.
func (in *TimeSeriesTable) DeepCopy() *TimeSeriesTable {
	if in == nil {
		return nil
	}
	out := new(TimeSeriesTable)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TimeSeriesTable_ColumnSettings) DeepCopyInto(out *TimeSeriesTable_ColumnSettings) {
	*out = *in
	if in.Column != nil {
		in, out := &in.Column, &out.Column
		*out = new(string)
		**out = **in
	}
	if in.Visible != nil {
		in, out := &in.Visible, &out.Visible
		*out = new(bool)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TimeSeriesTable_ColumnSettings.
func (in *TimeSeriesTable_ColumnSettings) DeepCopy() *TimeSeriesTable_ColumnSettings {
	if in == nil {
		return nil
	}
	out := new(TimeSeriesTable_ColumnSettings)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TimeSeriesTable_TableDataSet) DeepCopyInto(out *TimeSeriesTable_TableDataSet) {
	*out = *in
	if in.TimeSeriesQuery != nil {
		in, out := &in.TimeSeriesQuery, &out.TimeSeriesQuery
		*out = new(TimeSeriesQuery)
		(*in).DeepCopyInto(*out)
	}
	if in.TableTemplate != nil {
		in, out := &in.TableTemplate, &out.TableTemplate
		*out = new(string)
		**out = **in
	}
	if in.MinAlignmentPeriod != nil {
		in, out := &in.MinAlignmentPeriod, &out.MinAlignmentPeriod
		*out = new(string)
		**out = **in
	}
	if in.TableDisplayOptions != nil {
		in, out := &in.TableDisplayOptions, &out.TableDisplayOptions
		*out = new(TableDisplayOptions)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TimeSeriesTable_TableDataSet.
func (in *TimeSeriesTable_TableDataSet) DeepCopy() *TimeSeriesTable_TableDataSet {
	if in == nil {
		return nil
	}
	out := new(TimeSeriesTable_TableDataSet)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Widget) DeepCopyInto(out *Widget) {
	*out = *in
	if in.Title != nil {
		in, out := &in.Title, &out.Title
		*out = new(string)
		**out = **in
	}
	if in.XyChart != nil {
		in, out := &in.XyChart, &out.XyChart
		*out = new(XyChart)
		(*in).DeepCopyInto(*out)
	}
	if in.Scorecard != nil {
		in, out := &in.Scorecard, &out.Scorecard
		*out = new(Scorecard)
		(*in).DeepCopyInto(*out)
	}
	if in.Text != nil {
		in, out := &in.Text, &out.Text
		*out = new(Text)
		(*in).DeepCopyInto(*out)
	}
	if in.Blank != nil {
		in, out := &in.Blank, &out.Blank
		*out = new(Empty)
		**out = **in
	}
	if in.CollapsibleGroup != nil {
		in, out := &in.CollapsibleGroup, &out.CollapsibleGroup
		*out = new(CollapsibleGroup)
		(*in).DeepCopyInto(*out)
	}
	if in.LogsPanel != nil {
		in, out := &in.LogsPanel, &out.LogsPanel
		*out = new(LogsPanel)
		(*in).DeepCopyInto(*out)
	}
	if in.PieChart != nil {
		in, out := &in.PieChart, &out.PieChart
		*out = new(PieChart)
		(*in).DeepCopyInto(*out)
	}
	if in.ErrorReportingPanel != nil {
		in, out := &in.ErrorReportingPanel, &out.ErrorReportingPanel
		*out = new(ErrorReportingPanel)
		(*in).DeepCopyInto(*out)
	}
	if in.SectionHeader != nil {
		in, out := &in.SectionHeader, &out.SectionHeader
		*out = new(SectionHeader)
		(*in).DeepCopyInto(*out)
	}
	if in.SingleViewGroup != nil {
		in, out := &in.SingleViewGroup, &out.SingleViewGroup
		*out = new(SingleViewGroup)
		**out = **in
	}
	if in.Id != nil {
		in, out := &in.Id, &out.Id
		*out = new(string)
		**out = **in
	}
	if in.AlertChart != nil {
		in, out := &in.AlertChart, &out.AlertChart
		*out = new(AlertChart)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Widget.
func (in *Widget) DeepCopy() *Widget {
	if in == nil {
		return nil
	}
	out := new(Widget)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *XyChart) DeepCopyInto(out *XyChart) {
	*out = *in
	if in.DataSets != nil {
		in, out := &in.DataSets, &out.DataSets
		*out = make([]XyChart_DataSet, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.TimeshiftDuration != nil {
		in, out := &in.TimeshiftDuration, &out.TimeshiftDuration
		*out = new(string)
		**out = **in
	}
	if in.Thresholds != nil {
		in, out := &in.Thresholds, &out.Thresholds
		*out = make([]Threshold, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.XAxis != nil {
		in, out := &in.XAxis, &out.XAxis
		*out = new(XyChart_Axis)
		(*in).DeepCopyInto(*out)
	}
	if in.YAxis != nil {
		in, out := &in.YAxis, &out.YAxis
		*out = new(XyChart_Axis)
		(*in).DeepCopyInto(*out)
	}
	if in.ChartOptions != nil {
		in, out := &in.ChartOptions, &out.ChartOptions
		*out = new(ChartOptions)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new XyChart.
func (in *XyChart) DeepCopy() *XyChart {
	if in == nil {
		return nil
	}
	out := new(XyChart)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *XyChart_Axis) DeepCopyInto(out *XyChart_Axis) {
	*out = *in
	if in.Label != nil {
		in, out := &in.Label, &out.Label
		*out = new(string)
		**out = **in
	}
	if in.Scale != nil {
		in, out := &in.Scale, &out.Scale
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new XyChart_Axis.
func (in *XyChart_Axis) DeepCopy() *XyChart_Axis {
	if in == nil {
		return nil
	}
	out := new(XyChart_Axis)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *XyChart_DataSet) DeepCopyInto(out *XyChart_DataSet) {
	*out = *in
	if in.TimeSeriesQuery != nil {
		in, out := &in.TimeSeriesQuery, &out.TimeSeriesQuery
		*out = new(TimeSeriesQuery)
		(*in).DeepCopyInto(*out)
	}
	if in.PlotType != nil {
		in, out := &in.PlotType, &out.PlotType
		*out = new(string)
		**out = **in
	}
	if in.LegendTemplate != nil {
		in, out := &in.LegendTemplate, &out.LegendTemplate
		*out = new(string)
		**out = **in
	}
	if in.MinAlignmentPeriod != nil {
		in, out := &in.MinAlignmentPeriod, &out.MinAlignmentPeriod
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new XyChart_DataSet.
func (in *XyChart_DataSet) DeepCopy() *XyChart_DataSet {
	if in == nil {
		return nil
	}
	out := new(XyChart_DataSet)
	in.DeepCopyInto(out)
	return out
}
