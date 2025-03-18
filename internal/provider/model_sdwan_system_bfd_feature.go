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
type SystemBFD struct {
	Id                   types.String      `tfsdk:"id"`
	Version              types.Int64       `tfsdk:"version"`
	Name                 types.String      `tfsdk:"name"`
	Description          types.String      `tfsdk:"description"`
	FeatureProfileId     types.String      `tfsdk:"feature_profile_id"`
	Multiplier           types.Int64       `tfsdk:"multiplier"`
	MultiplierVariable   types.String      `tfsdk:"multiplier_variable"`
	PollInterval         types.Int64       `tfsdk:"poll_interval"`
	PollIntervalVariable types.String      `tfsdk:"poll_interval_variable"`
	DefaultDscp          types.Int64       `tfsdk:"default_dscp"`
	DefaultDscpVariable  types.String      `tfsdk:"default_dscp_variable"`
	Colors               []SystemBFDColors `tfsdk:"colors"`
}

type SystemBFDColors struct {
	Color                 types.String `tfsdk:"color"`
	ColorVariable         types.String `tfsdk:"color_variable"`
	HelloInterval         types.Int64  `tfsdk:"hello_interval"`
	HelloIntervalVariable types.String `tfsdk:"hello_interval_variable"`
	Multiplier            types.Int64  `tfsdk:"multiplier"`
	MultiplierVariable    types.String `tfsdk:"multiplier_variable"`
	PmtuDiscovery         types.Bool   `tfsdk:"pmtu_discovery"`
	PmtuDiscoveryVariable types.String `tfsdk:"pmtu_discovery_variable"`
	Dscp                  types.Int64  `tfsdk:"dscp"`
	DscpVariable          types.String `tfsdk:"dscp_variable"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getModel
func (data SystemBFD) getModel() string {
	return "system_bfd"
}

// End of section. //template:end getModel

// Section below is generated&owned by "gen/generator.go". //template:begin getPath
func (data SystemBFD) getPath() string {
	return fmt.Sprintf("/v1/feature-profile/sdwan/system/%v/bfd", url.QueryEscape(data.FeatureProfileId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody
func (data SystemBFD) toBody(ctx context.Context) string {
	body := ""
	body, _ = sjson.Set(body, "name", data.Name.ValueString())
	body, _ = sjson.Set(body, "description", data.Description.ValueString())
	path := "data."

	if !data.MultiplierVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"multiplier.optionType", "variable")
			body, _ = sjson.Set(body, path+"multiplier.value", data.MultiplierVariable.ValueString())
		}
	} else if data.Multiplier.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"multiplier.optionType", "default")
			body, _ = sjson.Set(body, path+"multiplier.value", 6)
		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"multiplier.optionType", "global")
			body, _ = sjson.Set(body, path+"multiplier.value", data.Multiplier.ValueInt64())
		}
	}

	if !data.PollIntervalVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"pollInterval.optionType", "variable")
			body, _ = sjson.Set(body, path+"pollInterval.value", data.PollIntervalVariable.ValueString())
		}
	} else if data.PollInterval.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"pollInterval.optionType", "default")
			body, _ = sjson.Set(body, path+"pollInterval.value", 600000)
		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"pollInterval.optionType", "global")
			body, _ = sjson.Set(body, path+"pollInterval.value", data.PollInterval.ValueInt64())
		}
	}

	if !data.DefaultDscpVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"defaultDscp.optionType", "variable")
			body, _ = sjson.Set(body, path+"defaultDscp.value", data.DefaultDscpVariable.ValueString())
		}
	} else if data.DefaultDscp.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"defaultDscp.optionType", "default")
			body, _ = sjson.Set(body, path+"defaultDscp.value", 48)
		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"defaultDscp.optionType", "global")
			body, _ = sjson.Set(body, path+"defaultDscp.value", data.DefaultDscp.ValueInt64())
		}
	}
	if true {
		body, _ = sjson.Set(body, path+"colors", []interface{}{})
		for _, item := range data.Colors {
			itemBody := ""

			if !item.ColorVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "color.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "color.value", item.ColorVariable.ValueString())
				}
			} else if !item.Color.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "color.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "color.value", item.Color.ValueString())
				}
			}

			if !item.HelloIntervalVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "helloInterval.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "helloInterval.value", item.HelloIntervalVariable.ValueString())
				}
			} else if item.HelloInterval.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "helloInterval.optionType", "default")
					itemBody, _ = sjson.Set(itemBody, "helloInterval.value", 1000)
				}
			} else {
				if true {
					itemBody, _ = sjson.Set(itemBody, "helloInterval.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "helloInterval.value", item.HelloInterval.ValueInt64())
				}
			}

			if !item.MultiplierVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "multiplier.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "multiplier.value", item.MultiplierVariable.ValueString())
				}
			} else if item.Multiplier.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "multiplier.optionType", "default")
					itemBody, _ = sjson.Set(itemBody, "multiplier.value", 7)
				}
			} else {
				if true {
					itemBody, _ = sjson.Set(itemBody, "multiplier.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "multiplier.value", item.Multiplier.ValueInt64())
				}
			}

			if !item.PmtuDiscoveryVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "pmtuDiscovery.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "pmtuDiscovery.value", item.PmtuDiscoveryVariable.ValueString())
				}
			} else if item.PmtuDiscovery.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "pmtuDiscovery.optionType", "default")
					itemBody, _ = sjson.Set(itemBody, "pmtuDiscovery.value", true)
				}
			} else {
				if true {
					itemBody, _ = sjson.Set(itemBody, "pmtuDiscovery.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "pmtuDiscovery.value", item.PmtuDiscovery.ValueBool())
				}
			}

			if !item.DscpVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "dscp.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "dscp.value", item.DscpVariable.ValueString())
				}
			} else if item.Dscp.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "dscp.optionType", "default")
					itemBody, _ = sjson.Set(itemBody, "dscp.value", 48)
				}
			} else {
				if true {
					itemBody, _ = sjson.Set(itemBody, "dscp.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "dscp.value", item.Dscp.ValueInt64())
				}
			}
			body, _ = sjson.SetRaw(body, path+"colors.-1", itemBody)
		}
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody
func (data *SystemBFD) fromBody(ctx context.Context, res gjson.Result) {
	data.Name = types.StringValue(res.Get("payload.name").String())
	if value := res.Get("payload.description"); value.Exists() && value.String() != "" {
		data.Description = types.StringValue(value.String())
	} else {
		data.Description = types.StringNull()
	}
	path := "payload.data."
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
	data.PollInterval = types.Int64Null()
	data.PollIntervalVariable = types.StringNull()
	if t := res.Get(path + "pollInterval.optionType"); t.Exists() {
		va := res.Get(path + "pollInterval.value")
		if t.String() == "variable" {
			data.PollIntervalVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.PollInterval = types.Int64Value(va.Int())
		}
	}
	data.DefaultDscp = types.Int64Null()
	data.DefaultDscpVariable = types.StringNull()
	if t := res.Get(path + "defaultDscp.optionType"); t.Exists() {
		va := res.Get(path + "defaultDscp.value")
		if t.String() == "variable" {
			data.DefaultDscpVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.DefaultDscp = types.Int64Value(va.Int())
		}
	}
	if value := res.Get(path + "colors"); value.Exists() {
		data.Colors = make([]SystemBFDColors, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := SystemBFDColors{}
			item.Color = types.StringNull()
			item.ColorVariable = types.StringNull()
			if t := v.Get("color.optionType"); t.Exists() {
				va := v.Get("color.value")
				if t.String() == "variable" {
					item.ColorVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.Color = types.StringValue(va.String())
				}
			}
			item.HelloInterval = types.Int64Null()
			item.HelloIntervalVariable = types.StringNull()
			if t := v.Get("helloInterval.optionType"); t.Exists() {
				va := v.Get("helloInterval.value")
				if t.String() == "variable" {
					item.HelloIntervalVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.HelloInterval = types.Int64Value(va.Int())
				}
			}
			item.Multiplier = types.Int64Null()
			item.MultiplierVariable = types.StringNull()
			if t := v.Get("multiplier.optionType"); t.Exists() {
				va := v.Get("multiplier.value")
				if t.String() == "variable" {
					item.MultiplierVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.Multiplier = types.Int64Value(va.Int())
				}
			}
			item.PmtuDiscovery = types.BoolNull()
			item.PmtuDiscoveryVariable = types.StringNull()
			if t := v.Get("pmtuDiscovery.optionType"); t.Exists() {
				va := v.Get("pmtuDiscovery.value")
				if t.String() == "variable" {
					item.PmtuDiscoveryVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.PmtuDiscovery = types.BoolValue(va.Bool())
				}
			}
			item.Dscp = types.Int64Null()
			item.DscpVariable = types.StringNull()
			if t := v.Get("dscp.optionType"); t.Exists() {
				va := v.Get("dscp.value")
				if t.String() == "variable" {
					item.DscpVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.Dscp = types.Int64Value(va.Int())
				}
			}
			data.Colors = append(data.Colors, item)
			return true
		})
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin updateFromBody
func (data *SystemBFD) updateFromBody(ctx context.Context, res gjson.Result) {
	data.Name = types.StringValue(res.Get("payload.name").String())
	if value := res.Get("payload.description"); value.Exists() && value.String() != "" {
		data.Description = types.StringValue(value.String())
	} else {
		data.Description = types.StringNull()
	}
	path := "payload.data."
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
	data.PollInterval = types.Int64Null()
	data.PollIntervalVariable = types.StringNull()
	if t := res.Get(path + "pollInterval.optionType"); t.Exists() {
		va := res.Get(path + "pollInterval.value")
		if t.String() == "variable" {
			data.PollIntervalVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.PollInterval = types.Int64Value(va.Int())
		}
	}
	data.DefaultDscp = types.Int64Null()
	data.DefaultDscpVariable = types.StringNull()
	if t := res.Get(path + "defaultDscp.optionType"); t.Exists() {
		va := res.Get(path + "defaultDscp.value")
		if t.String() == "variable" {
			data.DefaultDscpVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.DefaultDscp = types.Int64Value(va.Int())
		}
	}
	for i := range data.Colors {
		keys := [...]string{"color"}
		keyValues := [...]string{data.Colors[i].Color.ValueString()}
		keyValuesVariables := [...]string{data.Colors[i].ColorVariable.ValueString()}

		var r gjson.Result
		res.Get(path + "colors").ForEach(
			func(_, v gjson.Result) bool {
				found := false
				for ik := range keys {
					tt := v.Get(keys[ik] + ".optionType")
					vv := v.Get(keys[ik] + ".value")
					if tt.Exists() && vv.Exists() {
						if (tt.String() == "variable" && vv.String() == keyValuesVariables[ik]) || (tt.String() == "global" && vv.String() == keyValues[ik]) {
							found = true
							continue
						} else if tt.String() == "default" {
							continue
						}
						found = false
						break
					}
					continue
				}
				if found {
					r = v
					return false
				}
				return true
			},
		)
		data.Colors[i].Color = types.StringNull()
		data.Colors[i].ColorVariable = types.StringNull()
		if t := r.Get("color.optionType"); t.Exists() {
			va := r.Get("color.value")
			if t.String() == "variable" {
				data.Colors[i].ColorVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.Colors[i].Color = types.StringValue(va.String())
			}
		}
		data.Colors[i].HelloInterval = types.Int64Null()
		data.Colors[i].HelloIntervalVariable = types.StringNull()
		if t := r.Get("helloInterval.optionType"); t.Exists() {
			va := r.Get("helloInterval.value")
			if t.String() == "variable" {
				data.Colors[i].HelloIntervalVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.Colors[i].HelloInterval = types.Int64Value(va.Int())
			}
		}
		data.Colors[i].Multiplier = types.Int64Null()
		data.Colors[i].MultiplierVariable = types.StringNull()
		if t := r.Get("multiplier.optionType"); t.Exists() {
			va := r.Get("multiplier.value")
			if t.String() == "variable" {
				data.Colors[i].MultiplierVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.Colors[i].Multiplier = types.Int64Value(va.Int())
			}
		}
		data.Colors[i].PmtuDiscovery = types.BoolNull()
		data.Colors[i].PmtuDiscoveryVariable = types.StringNull()
		if t := r.Get("pmtuDiscovery.optionType"); t.Exists() {
			va := r.Get("pmtuDiscovery.value")
			if t.String() == "variable" {
				data.Colors[i].PmtuDiscoveryVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.Colors[i].PmtuDiscovery = types.BoolValue(va.Bool())
			}
		}
		data.Colors[i].Dscp = types.Int64Null()
		data.Colors[i].DscpVariable = types.StringNull()
		if t := r.Get("dscp.optionType"); t.Exists() {
			va := r.Get("dscp.value")
			if t.String() == "variable" {
				data.Colors[i].DscpVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.Colors[i].Dscp = types.Int64Value(va.Int())
			}
		}
	}
}

// End of section. //template:end updateFromBody

// Section below is generated&owned by "gen/generator.go". //template:begin isNull
func (data *SystemBFD) isNull(ctx context.Context, res gjson.Result) bool {
	if !data.FeatureProfileId.IsNull() {
		return false
	}
	if !data.Multiplier.IsNull() {
		return false
	}
	if !data.MultiplierVariable.IsNull() {
		return false
	}
	if !data.PollInterval.IsNull() {
		return false
	}
	if !data.PollIntervalVariable.IsNull() {
		return false
	}
	if !data.DefaultDscp.IsNull() {
		return false
	}
	if !data.DefaultDscpVariable.IsNull() {
		return false
	}
	if len(data.Colors) > 0 {
		return false
	}
	return true
}

// End of section. //template:end isNull
