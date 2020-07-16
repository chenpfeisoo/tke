// +build !ignore_autogenerated

/*
 * Tencent is pleased to support the open source community by making TKEStack
 * available.
 *
 * Copyright (C) 2012-2020 Tencent. All Rights Reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not use
 * this file except in compliance with the License. You may obtain a copy of the
 * License at
 *
 * https://opensource.org/licenses/Apache-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
 * WARRANTIES OF ANY KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations under the License.
 */

// Code generated by conversion-gen. DO NOT EDIT.

package v1

import (
	unsafe "unsafe"

	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
	monitor "tkestack.io/tke/api/monitor"
)

func init() {
	localSchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(s *runtime.Scheme) error {
	if err := s.AddGeneratedConversionFunc((*ClusterOverview)(nil), (*monitor.ClusterOverview)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_ClusterOverview_To_monitor_ClusterOverview(a.(*ClusterOverview), b.(*monitor.ClusterOverview), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*monitor.ClusterOverview)(nil), (*ClusterOverview)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_monitor_ClusterOverview_To_v1_ClusterOverview(a.(*monitor.ClusterOverview), b.(*ClusterOverview), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ClusterOverviewResult)(nil), (*monitor.ClusterOverviewResult)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_ClusterOverviewResult_To_monitor_ClusterOverviewResult(a.(*ClusterOverviewResult), b.(*monitor.ClusterOverviewResult), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*monitor.ClusterOverviewResult)(nil), (*ClusterOverviewResult)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_monitor_ClusterOverviewResult_To_v1_ClusterOverviewResult(a.(*monitor.ClusterOverviewResult), b.(*ClusterOverviewResult), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ClusterStatistic)(nil), (*monitor.ClusterStatistic)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_ClusterStatistic_To_monitor_ClusterStatistic(a.(*ClusterStatistic), b.(*monitor.ClusterStatistic), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*monitor.ClusterStatistic)(nil), (*ClusterStatistic)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_monitor_ClusterStatistic_To_v1_ClusterStatistic(a.(*monitor.ClusterStatistic), b.(*ClusterStatistic), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ConfigMap)(nil), (*monitor.ConfigMap)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_ConfigMap_To_monitor_ConfigMap(a.(*ConfigMap), b.(*monitor.ConfigMap), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*monitor.ConfigMap)(nil), (*ConfigMap)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_monitor_ConfigMap_To_v1_ConfigMap(a.(*monitor.ConfigMap), b.(*ConfigMap), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ConfigMapList)(nil), (*monitor.ConfigMapList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_ConfigMapList_To_monitor_ConfigMapList(a.(*ConfigMapList), b.(*monitor.ConfigMapList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*monitor.ConfigMapList)(nil), (*ConfigMapList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_monitor_ConfigMapList_To_v1_ConfigMapList(a.(*monitor.ConfigMapList), b.(*ConfigMapList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*Metric)(nil), (*monitor.Metric)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_Metric_To_monitor_Metric(a.(*Metric), b.(*monitor.Metric), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*monitor.Metric)(nil), (*Metric)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_monitor_Metric_To_v1_Metric(a.(*monitor.Metric), b.(*Metric), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*MetricList)(nil), (*monitor.MetricList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_MetricList_To_monitor_MetricList(a.(*MetricList), b.(*monitor.MetricList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*monitor.MetricList)(nil), (*MetricList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_monitor_MetricList_To_v1_MetricList(a.(*monitor.MetricList), b.(*MetricList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*MetricQuery)(nil), (*monitor.MetricQuery)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_MetricQuery_To_monitor_MetricQuery(a.(*MetricQuery), b.(*monitor.MetricQuery), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*monitor.MetricQuery)(nil), (*MetricQuery)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_monitor_MetricQuery_To_v1_MetricQuery(a.(*monitor.MetricQuery), b.(*MetricQuery), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*MetricQueryCondition)(nil), (*monitor.MetricQueryCondition)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_MetricQueryCondition_To_monitor_MetricQueryCondition(a.(*MetricQueryCondition), b.(*monitor.MetricQueryCondition), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*monitor.MetricQueryCondition)(nil), (*MetricQueryCondition)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_monitor_MetricQueryCondition_To_v1_MetricQueryCondition(a.(*monitor.MetricQueryCondition), b.(*MetricQueryCondition), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*Prometheus)(nil), (*monitor.Prometheus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_Prometheus_To_monitor_Prometheus(a.(*Prometheus), b.(*monitor.Prometheus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*monitor.Prometheus)(nil), (*Prometheus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_monitor_Prometheus_To_v1_Prometheus(a.(*monitor.Prometheus), b.(*Prometheus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*PrometheusList)(nil), (*monitor.PrometheusList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_PrometheusList_To_monitor_PrometheusList(a.(*PrometheusList), b.(*monitor.PrometheusList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*monitor.PrometheusList)(nil), (*PrometheusList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_monitor_PrometheusList_To_v1_PrometheusList(a.(*monitor.PrometheusList), b.(*PrometheusList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*PrometheusRemoteAddr)(nil), (*monitor.PrometheusRemoteAddr)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_PrometheusRemoteAddr_To_monitor_PrometheusRemoteAddr(a.(*PrometheusRemoteAddr), b.(*monitor.PrometheusRemoteAddr), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*monitor.PrometheusRemoteAddr)(nil), (*PrometheusRemoteAddr)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_monitor_PrometheusRemoteAddr_To_v1_PrometheusRemoteAddr(a.(*monitor.PrometheusRemoteAddr), b.(*PrometheusRemoteAddr), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*PrometheusSpec)(nil), (*monitor.PrometheusSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_PrometheusSpec_To_monitor_PrometheusSpec(a.(*PrometheusSpec), b.(*monitor.PrometheusSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*monitor.PrometheusSpec)(nil), (*PrometheusSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_monitor_PrometheusSpec_To_v1_PrometheusSpec(a.(*monitor.PrometheusSpec), b.(*PrometheusSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*PrometheusStatus)(nil), (*monitor.PrometheusStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_PrometheusStatus_To_monitor_PrometheusStatus(a.(*PrometheusStatus), b.(*monitor.PrometheusStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*monitor.PrometheusStatus)(nil), (*PrometheusStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_monitor_PrometheusStatus_To_v1_PrometheusStatus(a.(*monitor.PrometheusStatus), b.(*PrometheusStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ResourceRequirements)(nil), (*monitor.ResourceRequirements)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_ResourceRequirements_To_monitor_ResourceRequirements(a.(*ResourceRequirements), b.(*monitor.ResourceRequirements), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*monitor.ResourceRequirements)(nil), (*ResourceRequirements)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_monitor_ResourceRequirements_To_v1_ResourceRequirements(a.(*monitor.ResourceRequirements), b.(*ResourceRequirements), scope)
	}); err != nil {
		return err
	}
	return nil
}

func autoConvert_v1_ClusterOverview_To_monitor_ClusterOverview(in *ClusterOverview, out *monitor.ClusterOverview, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	out.Result = (*monitor.ClusterOverviewResult)(unsafe.Pointer(in.Result))
	return nil
}

// Convert_v1_ClusterOverview_To_monitor_ClusterOverview is an autogenerated conversion function.
func Convert_v1_ClusterOverview_To_monitor_ClusterOverview(in *ClusterOverview, out *monitor.ClusterOverview, s conversion.Scope) error {
	return autoConvert_v1_ClusterOverview_To_monitor_ClusterOverview(in, out, s)
}

func autoConvert_monitor_ClusterOverview_To_v1_ClusterOverview(in *monitor.ClusterOverview, out *ClusterOverview, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	out.Result = (*ClusterOverviewResult)(unsafe.Pointer(in.Result))
	return nil
}

// Convert_monitor_ClusterOverview_To_v1_ClusterOverview is an autogenerated conversion function.
func Convert_monitor_ClusterOverview_To_v1_ClusterOverview(in *monitor.ClusterOverview, out *ClusterOverview, s conversion.Scope) error {
	return autoConvert_monitor_ClusterOverview_To_v1_ClusterOverview(in, out, s)
}

func autoConvert_v1_ClusterOverviewResult_To_monitor_ClusterOverviewResult(in *ClusterOverviewResult, out *monitor.ClusterOverviewResult, s conversion.Scope) error {
	out.ClusterCount = in.ClusterCount
	out.ClusterAbnormal = in.ClusterAbnormal
	out.ProjectCount = in.ProjectCount
	out.ProjectAbnormal = in.ProjectAbnormal
	out.NodeCount = in.NodeCount
	out.NodeAbnormal = in.NodeAbnormal
	out.WorkloadCount = in.WorkloadCount
	out.WorkloadAbnormal = in.WorkloadAbnormal
	out.Clusters = *(*[]*monitor.ClusterStatistic)(unsafe.Pointer(&in.Clusters))
	return nil
}

// Convert_v1_ClusterOverviewResult_To_monitor_ClusterOverviewResult is an autogenerated conversion function.
func Convert_v1_ClusterOverviewResult_To_monitor_ClusterOverviewResult(in *ClusterOverviewResult, out *monitor.ClusterOverviewResult, s conversion.Scope) error {
	return autoConvert_v1_ClusterOverviewResult_To_monitor_ClusterOverviewResult(in, out, s)
}

func autoConvert_monitor_ClusterOverviewResult_To_v1_ClusterOverviewResult(in *monitor.ClusterOverviewResult, out *ClusterOverviewResult, s conversion.Scope) error {
	out.ClusterCount = in.ClusterCount
	out.ClusterAbnormal = in.ClusterAbnormal
	out.ProjectCount = in.ProjectCount
	out.ProjectAbnormal = in.ProjectAbnormal
	out.NodeCount = in.NodeCount
	out.NodeAbnormal = in.NodeAbnormal
	out.WorkloadCount = in.WorkloadCount
	out.WorkloadAbnormal = in.WorkloadAbnormal
	out.Clusters = *(*[]*ClusterStatistic)(unsafe.Pointer(&in.Clusters))
	return nil
}

// Convert_monitor_ClusterOverviewResult_To_v1_ClusterOverviewResult is an autogenerated conversion function.
func Convert_monitor_ClusterOverviewResult_To_v1_ClusterOverviewResult(in *monitor.ClusterOverviewResult, out *ClusterOverviewResult, s conversion.Scope) error {
	return autoConvert_monitor_ClusterOverviewResult_To_v1_ClusterOverviewResult(in, out, s)
}

func autoConvert_v1_ClusterStatistic_To_monitor_ClusterStatistic(in *ClusterStatistic, out *monitor.ClusterStatistic, s conversion.Scope) error {
	out.ClusterID = in.ClusterID
	out.ClusterPhase = in.ClusterPhase
	out.NodeCount = in.NodeCount
	out.NodeAbnormal = in.NodeAbnormal
	out.WorkloadCount = in.WorkloadCount
	out.WorkloadAbnormal = in.WorkloadAbnormal
	out.CPURequest = in.CPURequest
	out.CPULimit = in.CPULimit
	out.CPUCapacity = in.CPUCapacity
	out.CPUAllocatable = in.CPUAllocatable
	out.CPURequestRate = in.CPURequestRate
	out.CPUAllocatableRate = in.CPUAllocatableRate
	out.MemRequest = in.MemRequest
	out.MemLimit = in.MemLimit
	out.MemCapacity = in.MemCapacity
	out.MemAllocatable = in.MemAllocatable
	out.MemRequestRate = in.MemRequestRate
	out.MemAllocatableRate = in.MemAllocatableRate
	out.SchedulerHealthy = in.SchedulerHealthy
	out.ControllerManagerHealthy = in.ControllerManagerHealthy
	out.EtcdHealthy = in.EtcdHealthy
	return nil
}

// Convert_v1_ClusterStatistic_To_monitor_ClusterStatistic is an autogenerated conversion function.
func Convert_v1_ClusterStatistic_To_monitor_ClusterStatistic(in *ClusterStatistic, out *monitor.ClusterStatistic, s conversion.Scope) error {
	return autoConvert_v1_ClusterStatistic_To_monitor_ClusterStatistic(in, out, s)
}

func autoConvert_monitor_ClusterStatistic_To_v1_ClusterStatistic(in *monitor.ClusterStatistic, out *ClusterStatistic, s conversion.Scope) error {
	out.ClusterID = in.ClusterID
	out.ClusterPhase = in.ClusterPhase
	out.NodeCount = in.NodeCount
	out.NodeAbnormal = in.NodeAbnormal
	out.WorkloadCount = in.WorkloadCount
	out.WorkloadAbnormal = in.WorkloadAbnormal
	out.CPURequest = in.CPURequest
	out.CPULimit = in.CPULimit
	out.CPUCapacity = in.CPUCapacity
	out.CPUAllocatable = in.CPUAllocatable
	out.CPURequestRate = in.CPURequestRate
	out.CPUAllocatableRate = in.CPUAllocatableRate
	out.MemRequest = in.MemRequest
	out.MemLimit = in.MemLimit
	out.MemCapacity = in.MemCapacity
	out.MemAllocatable = in.MemAllocatable
	out.MemRequestRate = in.MemRequestRate
	out.MemAllocatableRate = in.MemAllocatableRate
	out.SchedulerHealthy = in.SchedulerHealthy
	out.ControllerManagerHealthy = in.ControllerManagerHealthy
	out.EtcdHealthy = in.EtcdHealthy
	return nil
}

// Convert_monitor_ClusterStatistic_To_v1_ClusterStatistic is an autogenerated conversion function.
func Convert_monitor_ClusterStatistic_To_v1_ClusterStatistic(in *monitor.ClusterStatistic, out *ClusterStatistic, s conversion.Scope) error {
	return autoConvert_monitor_ClusterStatistic_To_v1_ClusterStatistic(in, out, s)
}

func autoConvert_v1_ConfigMap_To_monitor_ConfigMap(in *ConfigMap, out *monitor.ConfigMap, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	out.Data = *(*map[string]string)(unsafe.Pointer(&in.Data))
	out.BinaryData = *(*map[string][]byte)(unsafe.Pointer(&in.BinaryData))
	return nil
}

// Convert_v1_ConfigMap_To_monitor_ConfigMap is an autogenerated conversion function.
func Convert_v1_ConfigMap_To_monitor_ConfigMap(in *ConfigMap, out *monitor.ConfigMap, s conversion.Scope) error {
	return autoConvert_v1_ConfigMap_To_monitor_ConfigMap(in, out, s)
}

func autoConvert_monitor_ConfigMap_To_v1_ConfigMap(in *monitor.ConfigMap, out *ConfigMap, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	out.Data = *(*map[string]string)(unsafe.Pointer(&in.Data))
	out.BinaryData = *(*map[string][]byte)(unsafe.Pointer(&in.BinaryData))
	return nil
}

// Convert_monitor_ConfigMap_To_v1_ConfigMap is an autogenerated conversion function.
func Convert_monitor_ConfigMap_To_v1_ConfigMap(in *monitor.ConfigMap, out *ConfigMap, s conversion.Scope) error {
	return autoConvert_monitor_ConfigMap_To_v1_ConfigMap(in, out, s)
}

func autoConvert_v1_ConfigMapList_To_monitor_ConfigMapList(in *ConfigMapList, out *monitor.ConfigMapList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]monitor.ConfigMap)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_v1_ConfigMapList_To_monitor_ConfigMapList is an autogenerated conversion function.
func Convert_v1_ConfigMapList_To_monitor_ConfigMapList(in *ConfigMapList, out *monitor.ConfigMapList, s conversion.Scope) error {
	return autoConvert_v1_ConfigMapList_To_monitor_ConfigMapList(in, out, s)
}

func autoConvert_monitor_ConfigMapList_To_v1_ConfigMapList(in *monitor.ConfigMapList, out *ConfigMapList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]ConfigMap)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_monitor_ConfigMapList_To_v1_ConfigMapList is an autogenerated conversion function.
func Convert_monitor_ConfigMapList_To_v1_ConfigMapList(in *monitor.ConfigMapList, out *ConfigMapList, s conversion.Scope) error {
	return autoConvert_monitor_ConfigMapList_To_v1_ConfigMapList(in, out, s)
}

func autoConvert_v1_Metric_To_monitor_Metric(in *Metric, out *monitor.Metric, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1_MetricQuery_To_monitor_MetricQuery(&in.Query, &out.Query, s); err != nil {
		return err
	}
	out.JSONResult = in.JSONResult
	return nil
}

// Convert_v1_Metric_To_monitor_Metric is an autogenerated conversion function.
func Convert_v1_Metric_To_monitor_Metric(in *Metric, out *monitor.Metric, s conversion.Scope) error {
	return autoConvert_v1_Metric_To_monitor_Metric(in, out, s)
}

func autoConvert_monitor_Metric_To_v1_Metric(in *monitor.Metric, out *Metric, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_monitor_MetricQuery_To_v1_MetricQuery(&in.Query, &out.Query, s); err != nil {
		return err
	}
	out.JSONResult = in.JSONResult
	return nil
}

// Convert_monitor_Metric_To_v1_Metric is an autogenerated conversion function.
func Convert_monitor_Metric_To_v1_Metric(in *monitor.Metric, out *Metric, s conversion.Scope) error {
	return autoConvert_monitor_Metric_To_v1_Metric(in, out, s)
}

func autoConvert_v1_MetricList_To_monitor_MetricList(in *MetricList, out *monitor.MetricList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]monitor.Metric)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_v1_MetricList_To_monitor_MetricList is an autogenerated conversion function.
func Convert_v1_MetricList_To_monitor_MetricList(in *MetricList, out *monitor.MetricList, s conversion.Scope) error {
	return autoConvert_v1_MetricList_To_monitor_MetricList(in, out, s)
}

func autoConvert_monitor_MetricList_To_v1_MetricList(in *monitor.MetricList, out *MetricList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]Metric)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_monitor_MetricList_To_v1_MetricList is an autogenerated conversion function.
func Convert_monitor_MetricList_To_v1_MetricList(in *monitor.MetricList, out *MetricList, s conversion.Scope) error {
	return autoConvert_monitor_MetricList_To_v1_MetricList(in, out, s)
}

func autoConvert_v1_MetricQuery_To_monitor_MetricQuery(in *MetricQuery, out *monitor.MetricQuery, s conversion.Scope) error {
	out.Table = in.Table
	out.StartTime = (*int64)(unsafe.Pointer(in.StartTime))
	out.EndTime = (*int64)(unsafe.Pointer(in.EndTime))
	out.Fields = *(*[]string)(unsafe.Pointer(&in.Fields))
	out.Conditions = *(*[]monitor.MetricQueryCondition)(unsafe.Pointer(&in.Conditions))
	out.OrderBy = in.OrderBy
	out.Order = in.Order
	out.GroupBy = *(*[]string)(unsafe.Pointer(&in.GroupBy))
	out.Limit = in.Limit
	out.Offset = in.Offset
	return nil
}

// Convert_v1_MetricQuery_To_monitor_MetricQuery is an autogenerated conversion function.
func Convert_v1_MetricQuery_To_monitor_MetricQuery(in *MetricQuery, out *monitor.MetricQuery, s conversion.Scope) error {
	return autoConvert_v1_MetricQuery_To_monitor_MetricQuery(in, out, s)
}

func autoConvert_monitor_MetricQuery_To_v1_MetricQuery(in *monitor.MetricQuery, out *MetricQuery, s conversion.Scope) error {
	out.Table = in.Table
	out.StartTime = (*int64)(unsafe.Pointer(in.StartTime))
	out.EndTime = (*int64)(unsafe.Pointer(in.EndTime))
	out.Fields = *(*[]string)(unsafe.Pointer(&in.Fields))
	out.Conditions = *(*[]MetricQueryCondition)(unsafe.Pointer(&in.Conditions))
	out.OrderBy = in.OrderBy
	out.Order = in.Order
	out.GroupBy = *(*[]string)(unsafe.Pointer(&in.GroupBy))
	out.Limit = in.Limit
	out.Offset = in.Offset
	return nil
}

// Convert_monitor_MetricQuery_To_v1_MetricQuery is an autogenerated conversion function.
func Convert_monitor_MetricQuery_To_v1_MetricQuery(in *monitor.MetricQuery, out *MetricQuery, s conversion.Scope) error {
	return autoConvert_monitor_MetricQuery_To_v1_MetricQuery(in, out, s)
}

func autoConvert_v1_MetricQueryCondition_To_monitor_MetricQueryCondition(in *MetricQueryCondition, out *monitor.MetricQueryCondition, s conversion.Scope) error {
	out.Key = in.Key
	out.Expr = in.Expr
	out.Value = in.Value
	return nil
}

// Convert_v1_MetricQueryCondition_To_monitor_MetricQueryCondition is an autogenerated conversion function.
func Convert_v1_MetricQueryCondition_To_monitor_MetricQueryCondition(in *MetricQueryCondition, out *monitor.MetricQueryCondition, s conversion.Scope) error {
	return autoConvert_v1_MetricQueryCondition_To_monitor_MetricQueryCondition(in, out, s)
}

func autoConvert_monitor_MetricQueryCondition_To_v1_MetricQueryCondition(in *monitor.MetricQueryCondition, out *MetricQueryCondition, s conversion.Scope) error {
	out.Key = in.Key
	out.Expr = in.Expr
	out.Value = in.Value
	return nil
}

// Convert_monitor_MetricQueryCondition_To_v1_MetricQueryCondition is an autogenerated conversion function.
func Convert_monitor_MetricQueryCondition_To_v1_MetricQueryCondition(in *monitor.MetricQueryCondition, out *MetricQueryCondition, s conversion.Scope) error {
	return autoConvert_monitor_MetricQueryCondition_To_v1_MetricQueryCondition(in, out, s)
}

func autoConvert_v1_Prometheus_To_monitor_Prometheus(in *Prometheus, out *monitor.Prometheus, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1_PrometheusSpec_To_monitor_PrometheusSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_v1_PrometheusStatus_To_monitor_PrometheusStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1_Prometheus_To_monitor_Prometheus is an autogenerated conversion function.
func Convert_v1_Prometheus_To_monitor_Prometheus(in *Prometheus, out *monitor.Prometheus, s conversion.Scope) error {
	return autoConvert_v1_Prometheus_To_monitor_Prometheus(in, out, s)
}

func autoConvert_monitor_Prometheus_To_v1_Prometheus(in *monitor.Prometheus, out *Prometheus, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_monitor_PrometheusSpec_To_v1_PrometheusSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_monitor_PrometheusStatus_To_v1_PrometheusStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_monitor_Prometheus_To_v1_Prometheus is an autogenerated conversion function.
func Convert_monitor_Prometheus_To_v1_Prometheus(in *monitor.Prometheus, out *Prometheus, s conversion.Scope) error {
	return autoConvert_monitor_Prometheus_To_v1_Prometheus(in, out, s)
}

func autoConvert_v1_PrometheusList_To_monitor_PrometheusList(in *PrometheusList, out *monitor.PrometheusList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]monitor.Prometheus)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_v1_PrometheusList_To_monitor_PrometheusList is an autogenerated conversion function.
func Convert_v1_PrometheusList_To_monitor_PrometheusList(in *PrometheusList, out *monitor.PrometheusList, s conversion.Scope) error {
	return autoConvert_v1_PrometheusList_To_monitor_PrometheusList(in, out, s)
}

func autoConvert_monitor_PrometheusList_To_v1_PrometheusList(in *monitor.PrometheusList, out *PrometheusList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]Prometheus)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_monitor_PrometheusList_To_v1_PrometheusList is an autogenerated conversion function.
func Convert_monitor_PrometheusList_To_v1_PrometheusList(in *monitor.PrometheusList, out *PrometheusList, s conversion.Scope) error {
	return autoConvert_monitor_PrometheusList_To_v1_PrometheusList(in, out, s)
}

func autoConvert_v1_PrometheusRemoteAddr_To_monitor_PrometheusRemoteAddr(in *PrometheusRemoteAddr, out *monitor.PrometheusRemoteAddr, s conversion.Scope) error {
	out.WriteAddr = *(*[]string)(unsafe.Pointer(&in.WriteAddr))
	out.ReadAddr = *(*[]string)(unsafe.Pointer(&in.ReadAddr))
	return nil
}

// Convert_v1_PrometheusRemoteAddr_To_monitor_PrometheusRemoteAddr is an autogenerated conversion function.
func Convert_v1_PrometheusRemoteAddr_To_monitor_PrometheusRemoteAddr(in *PrometheusRemoteAddr, out *monitor.PrometheusRemoteAddr, s conversion.Scope) error {
	return autoConvert_v1_PrometheusRemoteAddr_To_monitor_PrometheusRemoteAddr(in, out, s)
}

func autoConvert_monitor_PrometheusRemoteAddr_To_v1_PrometheusRemoteAddr(in *monitor.PrometheusRemoteAddr, out *PrometheusRemoteAddr, s conversion.Scope) error {
	out.WriteAddr = *(*[]string)(unsafe.Pointer(&in.WriteAddr))
	out.ReadAddr = *(*[]string)(unsafe.Pointer(&in.ReadAddr))
	return nil
}

// Convert_monitor_PrometheusRemoteAddr_To_v1_PrometheusRemoteAddr is an autogenerated conversion function.
func Convert_monitor_PrometheusRemoteAddr_To_v1_PrometheusRemoteAddr(in *monitor.PrometheusRemoteAddr, out *PrometheusRemoteAddr, s conversion.Scope) error {
	return autoConvert_monitor_PrometheusRemoteAddr_To_v1_PrometheusRemoteAddr(in, out, s)
}

func autoConvert_v1_PrometheusSpec_To_monitor_PrometheusSpec(in *PrometheusSpec, out *monitor.PrometheusSpec, s conversion.Scope) error {
	out.TenantID = in.TenantID
	out.ClusterName = in.ClusterName
	out.Version = in.Version
	out.SubVersion = *(*map[string]string)(unsafe.Pointer(&in.SubVersion))
	if err := Convert_v1_PrometheusRemoteAddr_To_monitor_PrometheusRemoteAddr(&in.RemoteAddress, &out.RemoteAddress, s); err != nil {
		return err
	}
	out.NotifyWebhook = in.NotifyWebhook
	if err := Convert_v1_ResourceRequirements_To_monitor_ResourceRequirements(&in.Resources, &out.Resources, s); err != nil {
		return err
	}
	out.RunOnMaster = in.RunOnMaster
	out.AlertRepeatInterval = in.AlertRepeatInterval
	out.WithNPD = in.WithNPD
	return nil
}

// Convert_v1_PrometheusSpec_To_monitor_PrometheusSpec is an autogenerated conversion function.
func Convert_v1_PrometheusSpec_To_monitor_PrometheusSpec(in *PrometheusSpec, out *monitor.PrometheusSpec, s conversion.Scope) error {
	return autoConvert_v1_PrometheusSpec_To_monitor_PrometheusSpec(in, out, s)
}

func autoConvert_monitor_PrometheusSpec_To_v1_PrometheusSpec(in *monitor.PrometheusSpec, out *PrometheusSpec, s conversion.Scope) error {
	out.TenantID = in.TenantID
	out.ClusterName = in.ClusterName
	out.Version = in.Version
	out.SubVersion = *(*map[string]string)(unsafe.Pointer(&in.SubVersion))
	if err := Convert_monitor_PrometheusRemoteAddr_To_v1_PrometheusRemoteAddr(&in.RemoteAddress, &out.RemoteAddress, s); err != nil {
		return err
	}
	out.NotifyWebhook = in.NotifyWebhook
	if err := Convert_monitor_ResourceRequirements_To_v1_ResourceRequirements(&in.Resources, &out.Resources, s); err != nil {
		return err
	}
	out.RunOnMaster = in.RunOnMaster
	out.AlertRepeatInterval = in.AlertRepeatInterval
	out.WithNPD = in.WithNPD
	return nil
}

// Convert_monitor_PrometheusSpec_To_v1_PrometheusSpec is an autogenerated conversion function.
func Convert_monitor_PrometheusSpec_To_v1_PrometheusSpec(in *monitor.PrometheusSpec, out *PrometheusSpec, s conversion.Scope) error {
	return autoConvert_monitor_PrometheusSpec_To_v1_PrometheusSpec(in, out, s)
}

func autoConvert_v1_PrometheusStatus_To_monitor_PrometheusStatus(in *PrometheusStatus, out *monitor.PrometheusStatus, s conversion.Scope) error {
	out.Version = in.Version
	out.Phase = monitor.AddonPhase(in.Phase)
	out.Reason = in.Reason
	out.RetryCount = in.RetryCount
	out.LastReInitializingTimestamp = in.LastReInitializingTimestamp
	out.SubVersion = *(*map[string]string)(unsafe.Pointer(&in.SubVersion))
	return nil
}

// Convert_v1_PrometheusStatus_To_monitor_PrometheusStatus is an autogenerated conversion function.
func Convert_v1_PrometheusStatus_To_monitor_PrometheusStatus(in *PrometheusStatus, out *monitor.PrometheusStatus, s conversion.Scope) error {
	return autoConvert_v1_PrometheusStatus_To_monitor_PrometheusStatus(in, out, s)
}

func autoConvert_monitor_PrometheusStatus_To_v1_PrometheusStatus(in *monitor.PrometheusStatus, out *PrometheusStatus, s conversion.Scope) error {
	out.Version = in.Version
	out.Phase = AddonPhase(in.Phase)
	out.Reason = in.Reason
	out.RetryCount = in.RetryCount
	out.LastReInitializingTimestamp = in.LastReInitializingTimestamp
	out.SubVersion = *(*map[string]string)(unsafe.Pointer(&in.SubVersion))
	return nil
}

// Convert_monitor_PrometheusStatus_To_v1_PrometheusStatus is an autogenerated conversion function.
func Convert_monitor_PrometheusStatus_To_v1_PrometheusStatus(in *monitor.PrometheusStatus, out *PrometheusStatus, s conversion.Scope) error {
	return autoConvert_monitor_PrometheusStatus_To_v1_PrometheusStatus(in, out, s)
}

func autoConvert_v1_ResourceRequirements_To_monitor_ResourceRequirements(in *ResourceRequirements, out *monitor.ResourceRequirements, s conversion.Scope) error {
	out.Limits = *(*monitor.ResourceList)(unsafe.Pointer(&in.Limits))
	out.Requests = *(*monitor.ResourceList)(unsafe.Pointer(&in.Requests))
	return nil
}

// Convert_v1_ResourceRequirements_To_monitor_ResourceRequirements is an autogenerated conversion function.
func Convert_v1_ResourceRequirements_To_monitor_ResourceRequirements(in *ResourceRequirements, out *monitor.ResourceRequirements, s conversion.Scope) error {
	return autoConvert_v1_ResourceRequirements_To_monitor_ResourceRequirements(in, out, s)
}

func autoConvert_monitor_ResourceRequirements_To_v1_ResourceRequirements(in *monitor.ResourceRequirements, out *ResourceRequirements, s conversion.Scope) error {
	out.Limits = *(*ResourceList)(unsafe.Pointer(&in.Limits))
	out.Requests = *(*ResourceList)(unsafe.Pointer(&in.Requests))
	return nil
}

// Convert_monitor_ResourceRequirements_To_v1_ResourceRequirements is an autogenerated conversion function.
func Convert_monitor_ResourceRequirements_To_v1_ResourceRequirements(in *monitor.ResourceRequirements, out *ResourceRequirements, s conversion.Scope) error {
	return autoConvert_monitor_ResourceRequirements_To_v1_ResourceRequirements(in, out, s)
}
