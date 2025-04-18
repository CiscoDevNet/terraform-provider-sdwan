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

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types
type ServiceObjectTracker struct {
	Id                      types.String `tfsdk:"id"`
	Version                 types.Int64  `tfsdk:"version"`
	Name                    types.String `tfsdk:"name"`
	Description             types.String `tfsdk:"description"`
	FeatureProfileId        types.String `tfsdk:"feature_profile_id"`
	ObjectTrackerId         types.Int64  `tfsdk:"object_tracker_id"`
	ObjectTrackerIdVariable types.String `tfsdk:"object_tracker_id_variable"`
	ObjectTrackerType       types.String `tfsdk:"object_tracker_type"`
	Interface               types.String `tfsdk:"interface"`
	InterfaceVariable       types.String `tfsdk:"interface_variable"`
	RouteIp                 types.String `tfsdk:"route_ip"`
	RouteIpVariable         types.String `tfsdk:"route_ip_variable"`
	RouteMask               types.String `tfsdk:"route_mask"`
	RouteMaskVariable       types.String `tfsdk:"route_mask_variable"`
	Vpn                     types.Int64  `tfsdk:"vpn"`
	VpnVariable             types.String `tfsdk:"vpn_variable"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getModel
func (data ServiceObjectTracker) getModel() string {
	return "service_object_tracker"
}

// End of section. //template:end getModel

// Section below is generated&owned by "gen/generator.go". //template:begin getPath
func (data ServiceObjectTracker) getPath() string {
	return fmt.Sprintf("/v1/feature-profile/sdwan/service/%v/objecttracker", url.QueryEscape(data.FeatureProfileId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody
func (data ServiceObjectTracker) toBody(ctx context.Context) string {
	body := ""
	body, _ = sjson.Set(body, "name", data.Name.ValueString())
	body, _ = sjson.Set(body, "description", data.Description.ValueString())
	path := "data."

	if !data.ObjectTrackerIdVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"objectId.optionType", "variable")
			body, _ = sjson.Set(body, path+"objectId.value", data.ObjectTrackerIdVariable.ValueString())
		}
	} else if !data.ObjectTrackerId.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"objectId.optionType", "global")
			body, _ = sjson.Set(body, path+"objectId.value", data.ObjectTrackerId.ValueInt64())
		}
	}
	if !data.ObjectTrackerType.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"objectTrackerType.optionType", "global")
			body, _ = sjson.Set(body, path+"objectTrackerType.value", data.ObjectTrackerType.ValueString())
		}
	}

	if !data.InterfaceVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"interface.optionType", "variable")
			body, _ = sjson.Set(body, path+"interface.value", data.InterfaceVariable.ValueString())
		}
	} else if !data.Interface.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"interface.optionType", "global")
			body, _ = sjson.Set(body, path+"interface.value", data.Interface.ValueString())
		}
	}

	if !data.RouteIpVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"routeIp.optionType", "variable")
			body, _ = sjson.Set(body, path+"routeIp.value", data.RouteIpVariable.ValueString())
		}
	} else if !data.RouteIp.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"routeIp.optionType", "global")
			body, _ = sjson.Set(body, path+"routeIp.value", data.RouteIp.ValueString())
		}
	}

	if !data.RouteMaskVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"routeMask.optionType", "variable")
			body, _ = sjson.Set(body, path+"routeMask.value", data.RouteMaskVariable.ValueString())
		}
	} else if data.RouteMask.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"routeMask.optionType", "default")
			body, _ = sjson.Set(body, path+"routeMask.value", "0.0.0.0")
		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"routeMask.optionType", "global")
			body, _ = sjson.Set(body, path+"routeMask.value", data.RouteMask.ValueString())
		}
	}

	if !data.VpnVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"vpn.optionType", "variable")
			body, _ = sjson.Set(body, path+"vpn.value", data.VpnVariable.ValueString())
		}
	} else if data.Vpn.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"vpn.optionType", "default")

		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"vpn.optionType", "global")
			body, _ = sjson.Set(body, path+"vpn.value", data.Vpn.ValueInt64())
		}
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody
func (data *ServiceObjectTracker) fromBody(ctx context.Context, res gjson.Result) {
	data.Name = types.StringValue(res.Get("payload.name").String())
	if value := res.Get("payload.description"); value.Exists() && value.String() != "" {
		data.Description = types.StringValue(value.String())
	} else {
		data.Description = types.StringNull()
	}
	path := "payload.data."
	data.ObjectTrackerId = types.Int64Null()
	data.ObjectTrackerIdVariable = types.StringNull()
	if t := res.Get(path + "objectId.optionType"); t.Exists() {
		va := res.Get(path + "objectId.value")
		if t.String() == "variable" {
			data.ObjectTrackerIdVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.ObjectTrackerId = types.Int64Value(va.Int())
		}
	}
	data.ObjectTrackerType = types.StringNull()

	if t := res.Get(path + "objectTrackerType.optionType"); t.Exists() {
		va := res.Get(path + "objectTrackerType.value")
		if t.String() == "global" {
			data.ObjectTrackerType = types.StringValue(va.String())
		}
	}
	data.Interface = types.StringNull()
	data.InterfaceVariable = types.StringNull()
	if t := res.Get(path + "interface.optionType"); t.Exists() {
		va := res.Get(path + "interface.value")
		if t.String() == "variable" {
			data.InterfaceVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.Interface = types.StringValue(va.String())
		}
	}
	data.RouteIp = types.StringNull()
	data.RouteIpVariable = types.StringNull()
	if t := res.Get(path + "routeIp.optionType"); t.Exists() {
		va := res.Get(path + "routeIp.value")
		if t.String() == "variable" {
			data.RouteIpVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.RouteIp = types.StringValue(va.String())
		}
	}
	data.RouteMask = types.StringNull()
	data.RouteMaskVariable = types.StringNull()
	if t := res.Get(path + "routeMask.optionType"); t.Exists() {
		va := res.Get(path + "routeMask.value")
		if t.String() == "variable" {
			data.RouteMaskVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.RouteMask = types.StringValue(va.String())
		}
	}
	data.Vpn = types.Int64Null()
	data.VpnVariable = types.StringNull()
	if t := res.Get(path + "vpn.optionType"); t.Exists() {
		va := res.Get(path + "vpn.value")
		if t.String() == "variable" {
			data.VpnVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.Vpn = types.Int64Value(va.Int())
		}
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin updateFromBody
func (data *ServiceObjectTracker) updateFromBody(ctx context.Context, res gjson.Result) {
	data.Name = types.StringValue(res.Get("payload.name").String())
	if value := res.Get("payload.description"); value.Exists() && value.String() != "" {
		data.Description = types.StringValue(value.String())
	} else {
		data.Description = types.StringNull()
	}
	path := "payload.data."
	data.ObjectTrackerId = types.Int64Null()
	data.ObjectTrackerIdVariable = types.StringNull()
	if t := res.Get(path + "objectId.optionType"); t.Exists() {
		va := res.Get(path + "objectId.value")
		if t.String() == "variable" {
			data.ObjectTrackerIdVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.ObjectTrackerId = types.Int64Value(va.Int())
		}
	}
	data.ObjectTrackerType = types.StringNull()

	if t := res.Get(path + "objectTrackerType.optionType"); t.Exists() {
		va := res.Get(path + "objectTrackerType.value")
		if t.String() == "global" {
			data.ObjectTrackerType = types.StringValue(va.String())
		}
	}
	data.Interface = types.StringNull()
	data.InterfaceVariable = types.StringNull()
	if t := res.Get(path + "interface.optionType"); t.Exists() {
		va := res.Get(path + "interface.value")
		if t.String() == "variable" {
			data.InterfaceVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.Interface = types.StringValue(va.String())
		}
	}
	data.RouteIp = types.StringNull()
	data.RouteIpVariable = types.StringNull()
	if t := res.Get(path + "routeIp.optionType"); t.Exists() {
		va := res.Get(path + "routeIp.value")
		if t.String() == "variable" {
			data.RouteIpVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.RouteIp = types.StringValue(va.String())
		}
	}
	data.RouteMask = types.StringNull()
	data.RouteMaskVariable = types.StringNull()
	if t := res.Get(path + "routeMask.optionType"); t.Exists() {
		va := res.Get(path + "routeMask.value")
		if t.String() == "variable" {
			data.RouteMaskVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.RouteMask = types.StringValue(va.String())
		}
	}
	data.Vpn = types.Int64Null()
	data.VpnVariable = types.StringNull()
	if t := res.Get(path + "vpn.optionType"); t.Exists() {
		va := res.Get(path + "vpn.value")
		if t.String() == "variable" {
			data.VpnVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.Vpn = types.Int64Value(va.Int())
		}
	}
}

// End of section. //template:end updateFromBody
