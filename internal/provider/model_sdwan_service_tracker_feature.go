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
type ServiceTracker struct {
	Id                          types.String `tfsdk:"id"`
	Version                     types.Int64  `tfsdk:"version"`
	Name                        types.String `tfsdk:"name"`
	Description                 types.String `tfsdk:"description"`
	FeatureProfileId            types.String `tfsdk:"feature_profile_id"`
	TrackerName                 types.String `tfsdk:"tracker_name"`
	TrackerNameVariable         types.String `tfsdk:"tracker_name_variable"`
	EndpointApiUrl              types.String `tfsdk:"endpoint_api_url"`
	EndpointApiUrlVariable      types.String `tfsdk:"endpoint_api_url_variable"`
	EndpointDnsName             types.String `tfsdk:"endpoint_dns_name"`
	EndpointDnsNameVariable     types.String `tfsdk:"endpoint_dns_name_variable"`
	EndpointIp                  types.String `tfsdk:"endpoint_ip"`
	EndpointIpVariable          types.String `tfsdk:"endpoint_ip_variable"`
	Protocol                    types.String `tfsdk:"protocol"`
	ProtocolVariable            types.String `tfsdk:"protocol_variable"`
	Port                        types.Int64  `tfsdk:"port"`
	PortVariable                types.String `tfsdk:"port_variable"`
	Interval                    types.Int64  `tfsdk:"interval"`
	IntervalVariable            types.String `tfsdk:"interval_variable"`
	Multiplier                  types.Int64  `tfsdk:"multiplier"`
	MultiplierVariable          types.String `tfsdk:"multiplier_variable"`
	Threshold                   types.Int64  `tfsdk:"threshold"`
	ThresholdVariable           types.String `tfsdk:"threshold_variable"`
	EndpointTrackerType         types.String `tfsdk:"endpoint_tracker_type"`
	EndpointTrackerTypeVariable types.String `tfsdk:"endpoint_tracker_type_variable"`
	TrackerType                 types.String `tfsdk:"tracker_type"`
	TrackerTypeVariable         types.String `tfsdk:"tracker_type_variable"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getModel
func (data ServiceTracker) getModel() string {
	return "service_tracker"
}

// End of section. //template:end getModel

// Section below is generated&owned by "gen/generator.go". //template:begin getPath
func (data ServiceTracker) getPath() string {
	return fmt.Sprintf("/v1/feature-profile/sdwan/service/%v/tracker", url.QueryEscape(data.FeatureProfileId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody
func (data ServiceTracker) toBody(ctx context.Context) string {
	body := ""
	body, _ = sjson.Set(body, "name", data.Name.ValueString())
	body, _ = sjson.Set(body, "description", data.Description.ValueString())
	path := "data."

	if !data.TrackerNameVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"trackerName.optionType", "variable")
			body, _ = sjson.Set(body, path+"trackerName.value", data.TrackerNameVariable.ValueString())
		}
	} else if !data.TrackerName.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"trackerName.optionType", "global")
			body, _ = sjson.Set(body, path+"trackerName.value", data.TrackerName.ValueString())
		}
	}

	if !data.EndpointApiUrlVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"endpointApiUrl.optionType", "variable")
			body, _ = sjson.Set(body, path+"endpointApiUrl.value", data.EndpointApiUrlVariable.ValueString())
		}
	} else if !data.EndpointApiUrl.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"endpointApiUrl.optionType", "global")
			body, _ = sjson.Set(body, path+"endpointApiUrl.value", data.EndpointApiUrl.ValueString())
		}
	}

	if !data.EndpointDnsNameVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"endpointDnsName.optionType", "variable")
			body, _ = sjson.Set(body, path+"endpointDnsName.value", data.EndpointDnsNameVariable.ValueString())
		}
	} else if !data.EndpointDnsName.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"endpointDnsName.optionType", "global")
			body, _ = sjson.Set(body, path+"endpointDnsName.value", data.EndpointDnsName.ValueString())
		}
	}

	if !data.EndpointIpVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"endpointIp.optionType", "variable")
			body, _ = sjson.Set(body, path+"endpointIp.value", data.EndpointIpVariable.ValueString())
		}
	} else if !data.EndpointIp.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"endpointIp.optionType", "global")
			body, _ = sjson.Set(body, path+"endpointIp.value", data.EndpointIp.ValueString())
		}
	}

	if !data.ProtocolVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"endpointTcpUdp.protocol.optionType", "variable")
			body, _ = sjson.Set(body, path+"endpointTcpUdp.protocol.value", data.ProtocolVariable.ValueString())
		}
	} else if !data.Protocol.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"endpointTcpUdp.protocol.optionType", "global")
			body, _ = sjson.Set(body, path+"endpointTcpUdp.protocol.value", data.Protocol.ValueString())
		}
	}

	if !data.PortVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"endpointTcpUdp.port.optionType", "variable")
			body, _ = sjson.Set(body, path+"endpointTcpUdp.port.value", data.PortVariable.ValueString())
		}
	} else if !data.Port.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"endpointTcpUdp.port.optionType", "global")
			body, _ = sjson.Set(body, path+"endpointTcpUdp.port.value", data.Port.ValueInt64())
		}
	}

	if !data.IntervalVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"interval.optionType", "variable")
			body, _ = sjson.Set(body, path+"interval.value", data.IntervalVariable.ValueString())
		}
	} else if data.Interval.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"interval.optionType", "default")
			body, _ = sjson.Set(body, path+"interval.value", 60)
		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"interval.optionType", "global")
			body, _ = sjson.Set(body, path+"interval.value", data.Interval.ValueInt64())
		}
	}

	if !data.MultiplierVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"multiplier.optionType", "variable")
			body, _ = sjson.Set(body, path+"multiplier.value", data.MultiplierVariable.ValueString())
		}
	} else if data.Multiplier.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"multiplier.optionType", "default")
			body, _ = sjson.Set(body, path+"multiplier.value", 3)
		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"multiplier.optionType", "global")
			body, _ = sjson.Set(body, path+"multiplier.value", data.Multiplier.ValueInt64())
		}
	}

	if !data.ThresholdVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"threshold.optionType", "variable")
			body, _ = sjson.Set(body, path+"threshold.value", data.ThresholdVariable.ValueString())
		}
	} else if data.Threshold.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"threshold.optionType", "default")
			body, _ = sjson.Set(body, path+"threshold.value", 300)
		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"threshold.optionType", "global")
			body, _ = sjson.Set(body, path+"threshold.value", data.Threshold.ValueInt64())
		}
	}

	if !data.EndpointTrackerTypeVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"endpointTrackerType.optionType", "variable")
			body, _ = sjson.Set(body, path+"endpointTrackerType.value", data.EndpointTrackerTypeVariable.ValueString())
		}
	} else if data.EndpointTrackerType.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"endpointTrackerType.optionType", "default")
			body, _ = sjson.Set(body, path+"endpointTrackerType.value", "static-route")
		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"endpointTrackerType.optionType", "global")
			body, _ = sjson.Set(body, path+"endpointTrackerType.value", data.EndpointTrackerType.ValueString())
		}
	}

	if !data.TrackerTypeVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"trackerType.optionType", "variable")
			body, _ = sjson.Set(body, path+"trackerType.value", data.TrackerTypeVariable.ValueString())
		}
	} else if data.TrackerType.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"trackerType.optionType", "default")
			body, _ = sjson.Set(body, path+"trackerType.value", "endpoint")
		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"trackerType.optionType", "global")
			body, _ = sjson.Set(body, path+"trackerType.value", data.TrackerType.ValueString())
		}
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody
func (data *ServiceTracker) fromBody(ctx context.Context, res gjson.Result) {
	data.Name = types.StringValue(res.Get("payload.name").String())
	if value := res.Get("payload.description"); value.Exists() && value.String() != "" {
		data.Description = types.StringValue(value.String())
	} else {
		data.Description = types.StringNull()
	}
	path := "payload.data."
	data.TrackerName = types.StringNull()
	data.TrackerNameVariable = types.StringNull()
	if t := res.Get(path + "trackerName.optionType"); t.Exists() {
		va := res.Get(path + "trackerName.value")
		if t.String() == "variable" {
			data.TrackerNameVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.TrackerName = types.StringValue(va.String())
		}
	}
	data.EndpointApiUrl = types.StringNull()
	data.EndpointApiUrlVariable = types.StringNull()
	if t := res.Get(path + "endpointApiUrl.optionType"); t.Exists() {
		va := res.Get(path + "endpointApiUrl.value")
		if t.String() == "variable" {
			data.EndpointApiUrlVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.EndpointApiUrl = types.StringValue(va.String())
		}
	}
	data.EndpointDnsName = types.StringNull()
	data.EndpointDnsNameVariable = types.StringNull()
	if t := res.Get(path + "endpointDnsName.optionType"); t.Exists() {
		va := res.Get(path + "endpointDnsName.value")
		if t.String() == "variable" {
			data.EndpointDnsNameVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.EndpointDnsName = types.StringValue(va.String())
		}
	}
	data.EndpointIp = types.StringNull()
	data.EndpointIpVariable = types.StringNull()
	if t := res.Get(path + "endpointIp.optionType"); t.Exists() {
		va := res.Get(path + "endpointIp.value")
		if t.String() == "variable" {
			data.EndpointIpVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.EndpointIp = types.StringValue(va.String())
		}
	}
	data.Protocol = types.StringNull()
	data.ProtocolVariable = types.StringNull()
	if t := res.Get(path + "endpointTcpUdp.protocol.optionType"); t.Exists() {
		va := res.Get(path + "endpointTcpUdp.protocol.value")
		if t.String() == "variable" {
			data.ProtocolVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.Protocol = types.StringValue(va.String())
		}
	}
	data.Port = types.Int64Null()
	data.PortVariable = types.StringNull()
	if t := res.Get(path + "endpointTcpUdp.port.optionType"); t.Exists() {
		va := res.Get(path + "endpointTcpUdp.port.value")
		if t.String() == "variable" {
			data.PortVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.Port = types.Int64Value(va.Int())
		}
	}
	data.Interval = types.Int64Null()
	data.IntervalVariable = types.StringNull()
	if t := res.Get(path + "interval.optionType"); t.Exists() {
		va := res.Get(path + "interval.value")
		if t.String() == "variable" {
			data.IntervalVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.Interval = types.Int64Value(va.Int())
		}
	}
	data.Multiplier = types.Int64Null()
	data.MultiplierVariable = types.StringNull()
	if t := res.Get(path + "multiplier.optionType"); t.Exists() {
		va := res.Get(path + "multiplier.value")
		if t.String() == "variable" {
			data.MultiplierVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.Multiplier = types.Int64Value(va.Int())
		}
	}
	data.Threshold = types.Int64Null()
	data.ThresholdVariable = types.StringNull()
	if t := res.Get(path + "threshold.optionType"); t.Exists() {
		va := res.Get(path + "threshold.value")
		if t.String() == "variable" {
			data.ThresholdVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.Threshold = types.Int64Value(va.Int())
		}
	}
	data.EndpointTrackerType = types.StringNull()
	data.EndpointTrackerTypeVariable = types.StringNull()
	if t := res.Get(path + "endpointTrackerType.optionType"); t.Exists() {
		va := res.Get(path + "endpointTrackerType.value")
		if t.String() == "variable" {
			data.EndpointTrackerTypeVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.EndpointTrackerType = types.StringValue(va.String())
		}
	}
	data.TrackerType = types.StringNull()
	data.TrackerTypeVariable = types.StringNull()
	if t := res.Get(path + "trackerType.optionType"); t.Exists() {
		va := res.Get(path + "trackerType.value")
		if t.String() == "variable" {
			data.TrackerTypeVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.TrackerType = types.StringValue(va.String())
		}
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin updateFromBody
func (data *ServiceTracker) updateFromBody(ctx context.Context, res gjson.Result) {
	data.Name = types.StringValue(res.Get("payload.name").String())
	if value := res.Get("payload.description"); value.Exists() && value.String() != "" {
		data.Description = types.StringValue(value.String())
	} else {
		data.Description = types.StringNull()
	}
	path := "payload.data."
	data.TrackerName = types.StringNull()
	data.TrackerNameVariable = types.StringNull()
	if t := res.Get(path + "trackerName.optionType"); t.Exists() {
		va := res.Get(path + "trackerName.value")
		if t.String() == "variable" {
			data.TrackerNameVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.TrackerName = types.StringValue(va.String())
		}
	}
	data.EndpointApiUrl = types.StringNull()
	data.EndpointApiUrlVariable = types.StringNull()
	if t := res.Get(path + "endpointApiUrl.optionType"); t.Exists() {
		va := res.Get(path + "endpointApiUrl.value")
		if t.String() == "variable" {
			data.EndpointApiUrlVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.EndpointApiUrl = types.StringValue(va.String())
		}
	}
	data.EndpointDnsName = types.StringNull()
	data.EndpointDnsNameVariable = types.StringNull()
	if t := res.Get(path + "endpointDnsName.optionType"); t.Exists() {
		va := res.Get(path + "endpointDnsName.value")
		if t.String() == "variable" {
			data.EndpointDnsNameVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.EndpointDnsName = types.StringValue(va.String())
		}
	}
	data.EndpointIp = types.StringNull()
	data.EndpointIpVariable = types.StringNull()
	if t := res.Get(path + "endpointIp.optionType"); t.Exists() {
		va := res.Get(path + "endpointIp.value")
		if t.String() == "variable" {
			data.EndpointIpVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.EndpointIp = types.StringValue(va.String())
		}
	}
	data.Protocol = types.StringNull()
	data.ProtocolVariable = types.StringNull()
	if t := res.Get(path + "endpointTcpUdp.protocol.optionType"); t.Exists() {
		va := res.Get(path + "endpointTcpUdp.protocol.value")
		if t.String() == "variable" {
			data.ProtocolVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.Protocol = types.StringValue(va.String())
		}
	}
	data.Port = types.Int64Null()
	data.PortVariable = types.StringNull()
	if t := res.Get(path + "endpointTcpUdp.port.optionType"); t.Exists() {
		va := res.Get(path + "endpointTcpUdp.port.value")
		if t.String() == "variable" {
			data.PortVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.Port = types.Int64Value(va.Int())
		}
	}
	data.Interval = types.Int64Null()
	data.IntervalVariable = types.StringNull()
	if t := res.Get(path + "interval.optionType"); t.Exists() {
		va := res.Get(path + "interval.value")
		if t.String() == "variable" {
			data.IntervalVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.Interval = types.Int64Value(va.Int())
		}
	}
	data.Multiplier = types.Int64Null()
	data.MultiplierVariable = types.StringNull()
	if t := res.Get(path + "multiplier.optionType"); t.Exists() {
		va := res.Get(path + "multiplier.value")
		if t.String() == "variable" {
			data.MultiplierVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.Multiplier = types.Int64Value(va.Int())
		}
	}
	data.Threshold = types.Int64Null()
	data.ThresholdVariable = types.StringNull()
	if t := res.Get(path + "threshold.optionType"); t.Exists() {
		va := res.Get(path + "threshold.value")
		if t.String() == "variable" {
			data.ThresholdVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.Threshold = types.Int64Value(va.Int())
		}
	}
	data.EndpointTrackerType = types.StringNull()
	data.EndpointTrackerTypeVariable = types.StringNull()
	if t := res.Get(path + "endpointTrackerType.optionType"); t.Exists() {
		va := res.Get(path + "endpointTrackerType.value")
		if t.String() == "variable" {
			data.EndpointTrackerTypeVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.EndpointTrackerType = types.StringValue(va.String())
		}
	}
	data.TrackerType = types.StringNull()
	data.TrackerTypeVariable = types.StringNull()
	if t := res.Get(path + "trackerType.optionType"); t.Exists() {
		va := res.Get(path + "trackerType.value")
		if t.String() == "variable" {
			data.TrackerTypeVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.TrackerType = types.StringValue(va.String())
		}
	}
}

// End of section. //template:end updateFromBody
