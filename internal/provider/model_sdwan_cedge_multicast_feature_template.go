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
	"strconv"

	"github.com/CiscoDevNet/terraform-provider-sdwan/internal/provider/helpers"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types
type CEdgeMulticast struct {
	Id                      types.String `tfsdk:"id"`
	Version                 types.Int64  `tfsdk:"version"`
	TemplateType            types.String `tfsdk:"template_type"`
	Name                    types.String `tfsdk:"name"`
	Description             types.String `tfsdk:"description"`
	DeviceTypes             types.Set    `tfsdk:"device_types"`
	SptOnly                 types.Bool   `tfsdk:"spt_only"`
	SptOnlyVariable         types.String `tfsdk:"spt_only_variable"`
	LocalReplicator         types.Bool   `tfsdk:"local_replicator"`
	LocalReplicatorVariable types.String `tfsdk:"local_replicator_variable"`
	Threshold               types.Int64  `tfsdk:"threshold"`
	ThresholdVariable       types.String `tfsdk:"threshold_variable"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getModel
func (data CEdgeMulticast) getModel() string {
	return "cedge_multicast"
}

// End of section. //template:end getModel

// Section below is generated&owned by "gen/generator.go". //template:begin toBody
func (data CEdgeMulticast) toBody(ctx context.Context) string {
	body := ""

	var device_types []string
	data.DeviceTypes.ElementsAs(ctx, &device_types, false)
	body, _ = sjson.Set(body, "deviceType", device_types)
	body, _ = sjson.Set(body, "factoryDefault", false)
	body, _ = sjson.Set(body, "templateDescription", data.Description.ValueString())
	body, _ = sjson.Set(body, "templateMinVersion", "15.0.0")
	body, _ = sjson.Set(body, "templateName", data.Name.ValueString())
	body, _ = sjson.Set(body, "templateType", "cedge_multicast")
	body, _ = sjson.Set(body, "templateDefinition", map[string]interface{}{})

	path := "templateDefinition."

	if !data.SptOnlyVariable.IsNull() {
		body, _ = sjson.Set(body, path+"multicast.spt-only."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"multicast.spt-only."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"multicast.spt-only."+"vipVariableName", data.SptOnlyVariable.ValueString())
	} else if data.SptOnly.IsNull() {
		if !gjson.Get(body, path+"multicast").Exists() {
			body, _ = sjson.Set(body, path+"multicast", map[string]interface{}{})
		}
	} else {
		body, _ = sjson.Set(body, path+"multicast.spt-only."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"multicast.spt-only."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"multicast.spt-only."+"vipValue", strconv.FormatBool(data.SptOnly.ValueBool()))
	}

	if !data.LocalReplicatorVariable.IsNull() {
		body, _ = sjson.Set(body, path+"multicast-replicator.local."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"multicast-replicator.local."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"multicast-replicator.local."+"vipVariableName", data.LocalReplicatorVariable.ValueString())
	} else if data.LocalReplicator.IsNull() {
		if !gjson.Get(body, path+"multicast-replicator").Exists() {
			body, _ = sjson.Set(body, path+"multicast-replicator", map[string]interface{}{})
		}
	} else {
		body, _ = sjson.Set(body, path+"multicast-replicator.local."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"multicast-replicator.local."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"multicast-replicator.local."+"vipValue", strconv.FormatBool(data.LocalReplicator.ValueBool()))
	}

	if !data.ThresholdVariable.IsNull() {
		body, _ = sjson.Set(body, path+"multicast-replicator.threshold."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"multicast-replicator.threshold."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"multicast-replicator.threshold."+"vipVariableName", data.ThresholdVariable.ValueString())
	} else if data.Threshold.IsNull() {
		body, _ = sjson.Set(body, path+"multicast-replicator.threshold."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"multicast-replicator.threshold."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"multicast-replicator.threshold."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"multicast-replicator.threshold."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"multicast-replicator.threshold."+"vipValue", data.Threshold.ValueInt64())
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody
func (data *CEdgeMulticast) fromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("deviceType"); value.Exists() {
		data.DeviceTypes = helpers.GetStringSet(value.Array())
	} else {
		data.DeviceTypes = types.SetNull(types.StringType)
	}
	if value := res.Get("templateDescription"); value.Exists() && value.String() != "" {
		data.Description = types.StringValue(value.String())
	} else {
		data.Description = types.StringNull()
	}
	if value := res.Get("templateName"); value.Exists() {
		data.Name = types.StringValue(value.String())
	} else {
		data.Name = types.StringNull()
	}
	if value := res.Get("templateType"); value.Exists() {
		data.TemplateType = types.StringValue(value.String())
	} else {
		data.TemplateType = types.StringNull()
	}

	path := "templateDefinition."
	if value := res.Get(path + "multicast.spt-only.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.SptOnly = types.BoolNull()

			v := res.Get(path + "multicast.spt-only.vipVariableName")
			data.SptOnlyVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.SptOnly = types.BoolNull()
			data.SptOnlyVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "multicast.spt-only.vipValue")
			data.SptOnly = types.BoolValue(v.Bool())
			data.SptOnlyVariable = types.StringNull()
		}
	} else {
		data.SptOnly = types.BoolNull()
		data.SptOnlyVariable = types.StringNull()
	}
	if value := res.Get(path + "multicast-replicator.local.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.LocalReplicator = types.BoolNull()

			v := res.Get(path + "multicast-replicator.local.vipVariableName")
			data.LocalReplicatorVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.LocalReplicator = types.BoolNull()
			data.LocalReplicatorVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "multicast-replicator.local.vipValue")
			data.LocalReplicator = types.BoolValue(v.Bool())
			data.LocalReplicatorVariable = types.StringNull()
		}
	} else {
		data.LocalReplicator = types.BoolNull()
		data.LocalReplicatorVariable = types.StringNull()
	}
	if value := res.Get(path + "multicast-replicator.threshold.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.Threshold = types.Int64Null()

			v := res.Get(path + "multicast-replicator.threshold.vipVariableName")
			data.ThresholdVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.Threshold = types.Int64Null()
			data.ThresholdVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "multicast-replicator.threshold.vipValue")
			data.Threshold = types.Int64Value(v.Int())
			data.ThresholdVariable = types.StringNull()
		}
	} else {
		data.Threshold = types.Int64Null()
		data.ThresholdVariable = types.StringNull()
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin hasChanges
func (data *CEdgeMulticast) hasChanges(ctx context.Context, state *CEdgeMulticast) bool {
	hasChanges := false
	if !data.SptOnly.Equal(state.SptOnly) {
		hasChanges = true
	}
	if !data.LocalReplicator.Equal(state.LocalReplicator) {
		hasChanges = true
	}
	if !data.Threshold.Equal(state.Threshold) {
		hasChanges = true
	}
	return hasChanges
}

// End of section. //template:end hasChanges
