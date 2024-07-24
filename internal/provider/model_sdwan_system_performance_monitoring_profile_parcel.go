// Copyright © 2023 Cisco Systems, Inc. and its affiliates.
// All rights reserved.
//
// Licensed under the Mozilla Public License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://mozilla.org/MPL/2.0/
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// SPDX-License-Identifier: MPL-2.0

package provider

// Section below is generated&owned by "gen/generator.go". //template:begin imports
import (
	"context"
	"fmt"
	"net/url"

	"github.com/CiscoDevNet/terraform-provider-sdwan/internal/provider/helpers"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types
type SystemPerformanceMonitoring struct {
	Id                       types.String `tfsdk:"id"`
	Version                  types.Int64  `tfsdk:"version"`
	Name                     types.String `tfsdk:"name"`
	Description              types.String `tfsdk:"description"`
	FeatureProfileId         types.String `tfsdk:"feature_profile_id"`
	AppPerfMonitorEnabled    types.Bool   `tfsdk:"app_perf_monitor_enabled"`
	AppPerfMonitorAppGroup   types.Set    `tfsdk:"app_perf_monitor_app_group"`
	MonitoringConfigEnabled  types.Bool   `tfsdk:"monitoring_config_enabled"`
	MonitoringConfigInterval types.String `tfsdk:"monitoring_config_interval"`
	EventDrivenConfigEnabled types.Bool   `tfsdk:"event_driven_config_enabled"`
	EventDrivenEvents        types.Set    `tfsdk:"event_driven_events"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getModel
func (data SystemPerformanceMonitoring) getModel() string {
	return "system_performance_monitoring"
}

// End of section. //template:end getModel

// Section below is generated&owned by "gen/generator.go". //template:begin getPath
func (data SystemPerformanceMonitoring) getPath() string {
	return fmt.Sprintf("/v1/feature-profile/sdwan/system/%v/perfmonitor", url.QueryEscape(data.FeatureProfileId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody
func (data SystemPerformanceMonitoring) toBody(ctx context.Context) string {
	body := ""
	body, _ = sjson.Set(body, "name", data.Name.ValueString())
	body, _ = sjson.Set(body, "description", data.Description.ValueString())
	path := "data."
	if data.AppPerfMonitorEnabled.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"appPerfMonitorConfig.enabled.optionType", "default")
			body, _ = sjson.Set(body, path+"appPerfMonitorConfig.enabled.value", false)
		}
	} else {
		body, _ = sjson.Set(body, path+"appPerfMonitorConfig.enabled.optionType", "global")
		body, _ = sjson.Set(body, path+"appPerfMonitorConfig.enabled.value", data.AppPerfMonitorEnabled.ValueBool())
	}
	if !data.AppPerfMonitorAppGroup.IsNull() {
		body, _ = sjson.Set(body, path+"appPerfMonitorConfig.policyFilters.appGroups.optionType", "global")
		var values []string
		data.AppPerfMonitorAppGroup.ElementsAs(ctx, &values, false)
		body, _ = sjson.Set(body, path+"appPerfMonitorConfig.policyFilters.appGroups.value", values)
	}
	if data.MonitoringConfigEnabled.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"umtsConfig.monitoringConfig.enabled.optionType", "default")
			body, _ = sjson.Set(body, path+"umtsConfig.monitoringConfig.enabled.value", false)
		}
	} else {
		body, _ = sjson.Set(body, path+"umtsConfig.monitoringConfig.enabled.optionType", "global")
		body, _ = sjson.Set(body, path+"umtsConfig.monitoringConfig.enabled.value", data.MonitoringConfigEnabled.ValueBool())
	}
	if !data.MonitoringConfigInterval.IsNull() {
		body, _ = sjson.Set(body, path+"umtsConfig.monitoringConfig.interval.optionType", "global")
		body, _ = sjson.Set(body, path+"umtsConfig.monitoringConfig.interval.value", data.MonitoringConfigInterval.ValueString())
	}
	if data.EventDrivenConfigEnabled.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"umtsConfig.eventDrivenConfig.enabled.optionType", "default")
			body, _ = sjson.Set(body, path+"umtsConfig.eventDrivenConfig.enabled.value", false)
		}
	} else {
		body, _ = sjson.Set(body, path+"umtsConfig.eventDrivenConfig.enabled.optionType", "global")
		body, _ = sjson.Set(body, path+"umtsConfig.eventDrivenConfig.enabled.value", data.EventDrivenConfigEnabled.ValueBool())
	}
	if !data.EventDrivenEvents.IsNull() {
		body, _ = sjson.Set(body, path+"umtsConfig.eventDrivenConfig.events.optionType", "global")
		var values []string
		data.EventDrivenEvents.ElementsAs(ctx, &values, false)
		body, _ = sjson.Set(body, path+"umtsConfig.eventDrivenConfig.events.value", values)
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody
func (data *SystemPerformanceMonitoring) fromBody(ctx context.Context, res gjson.Result) {
	data.Name = types.StringValue(res.Get("payload.name").String())
	if value := res.Get("payload.description"); value.Exists() && value.String() != "" {
		data.Description = types.StringValue(value.String())
	} else {
		data.Description = types.StringNull()
	}
	path := "payload.data."
	data.AppPerfMonitorEnabled = types.BoolNull()

	if t := res.Get(path + "appPerfMonitorConfig.enabled.optionType"); t.Exists() {
		va := res.Get(path + "appPerfMonitorConfig.enabled.value")
		if t.String() == "global" {
			data.AppPerfMonitorEnabled = types.BoolValue(va.Bool())
		}
	}
	data.AppPerfMonitorAppGroup = types.SetNull(types.StringType)

	if t := res.Get(path + "appPerfMonitorConfig.policyFilters.appGroups.optionType"); t.Exists() {
		va := res.Get(path + "appPerfMonitorConfig.policyFilters.appGroups.value")
		if t.String() == "global" {
			data.AppPerfMonitorAppGroup = helpers.GetStringSet(va.Array())
		}
	}
	data.MonitoringConfigEnabled = types.BoolNull()

	if t := res.Get(path + "umtsConfig.monitoringConfig.enabled.optionType"); t.Exists() {
		va := res.Get(path + "umtsConfig.monitoringConfig.enabled.value")
		if t.String() == "global" {
			data.MonitoringConfigEnabled = types.BoolValue(va.Bool())
		}
	}
	data.MonitoringConfigInterval = types.StringNull()

	if t := res.Get(path + "umtsConfig.monitoringConfig.interval.optionType"); t.Exists() {
		va := res.Get(path + "umtsConfig.monitoringConfig.interval.value")
		if t.String() == "global" {
			data.MonitoringConfigInterval = types.StringValue(va.String())
		}
	}
	data.EventDrivenConfigEnabled = types.BoolNull()

	if t := res.Get(path + "umtsConfig.eventDrivenConfig.enabled.optionType"); t.Exists() {
		va := res.Get(path + "umtsConfig.eventDrivenConfig.enabled.value")
		if t.String() == "global" {
			data.EventDrivenConfigEnabled = types.BoolValue(va.Bool())
		}
	}
	data.EventDrivenEvents = types.SetNull(types.StringType)

	if t := res.Get(path + "umtsConfig.eventDrivenConfig.events.optionType"); t.Exists() {
		va := res.Get(path + "umtsConfig.eventDrivenConfig.events.value")
		if t.String() == "global" {
			data.EventDrivenEvents = helpers.GetStringSet(va.Array())
		}
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin updateFromBody
func (data *SystemPerformanceMonitoring) updateFromBody(ctx context.Context, res gjson.Result) {
	data.Name = types.StringValue(res.Get("payload.name").String())
	if value := res.Get("payload.description"); value.Exists() && value.String() != "" {
		data.Description = types.StringValue(value.String())
	} else {
		data.Description = types.StringNull()
	}
	path := "payload.data."
	data.AppPerfMonitorEnabled = types.BoolNull()

	if t := res.Get(path + "appPerfMonitorConfig.enabled.optionType"); t.Exists() {
		va := res.Get(path + "appPerfMonitorConfig.enabled.value")
		if t.String() == "global" {
			data.AppPerfMonitorEnabled = types.BoolValue(va.Bool())
		}
	}
	data.AppPerfMonitorAppGroup = types.SetNull(types.StringType)

	if t := res.Get(path + "appPerfMonitorConfig.policyFilters.appGroups.optionType"); t.Exists() {
		va := res.Get(path + "appPerfMonitorConfig.policyFilters.appGroups.value")
		if t.String() == "global" {
			data.AppPerfMonitorAppGroup = helpers.GetStringSet(va.Array())
		}
	}
	data.MonitoringConfigEnabled = types.BoolNull()

	if t := res.Get(path + "umtsConfig.monitoringConfig.enabled.optionType"); t.Exists() {
		va := res.Get(path + "umtsConfig.monitoringConfig.enabled.value")
		if t.String() == "global" {
			data.MonitoringConfigEnabled = types.BoolValue(va.Bool())
		}
	}
	data.MonitoringConfigInterval = types.StringNull()

	if t := res.Get(path + "umtsConfig.monitoringConfig.interval.optionType"); t.Exists() {
		va := res.Get(path + "umtsConfig.monitoringConfig.interval.value")
		if t.String() == "global" {
			data.MonitoringConfigInterval = types.StringValue(va.String())
		}
	}
	data.EventDrivenConfigEnabled = types.BoolNull()

	if t := res.Get(path + "umtsConfig.eventDrivenConfig.enabled.optionType"); t.Exists() {
		va := res.Get(path + "umtsConfig.eventDrivenConfig.enabled.value")
		if t.String() == "global" {
			data.EventDrivenConfigEnabled = types.BoolValue(va.Bool())
		}
	}
	data.EventDrivenEvents = types.SetNull(types.StringType)

	if t := res.Get(path + "umtsConfig.eventDrivenConfig.events.optionType"); t.Exists() {
		va := res.Get(path + "umtsConfig.eventDrivenConfig.events.value")
		if t.String() == "global" {
			data.EventDrivenEvents = helpers.GetStringSet(va.Array())
		}
	}
}

// End of section. //template:end updateFromBody

// Section below is generated&owned by "gen/generator.go". //template:begin isNull
func (data *SystemPerformanceMonitoring) isNull(ctx context.Context, res gjson.Result) bool {
	if !data.FeatureProfileId.IsNull() {
		return false
	}
	if !data.AppPerfMonitorEnabled.IsNull() {
		return false
	}
	if !data.AppPerfMonitorAppGroup.IsNull() {
		return false
	}
	if !data.MonitoringConfigEnabled.IsNull() {
		return false
	}
	if !data.MonitoringConfigInterval.IsNull() {
		return false
	}
	if !data.EventDrivenConfigEnabled.IsNull() {
		return false
	}
	if !data.EventDrivenEvents.IsNull() {
		return false
	}
	return true
}

// End of section. //template:end isNull
