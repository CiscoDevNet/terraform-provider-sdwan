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

	"github.com/CiscoDevNet/terraform-provider-sdwan/internal/provider/helpers"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types
type CustomApplication struct {
	Id                types.String            `tfsdk:"id"`
	Version           types.Int64             `tfsdk:"version"`
	AppName           types.String            `tfsdk:"app_name"`
	ServerNames       types.Set               `tfsdk:"server_names"`
	L3l4              []CustomApplicationL3l4 `tfsdk:"l3l4"`
	ApplicationFamily types.String            `tfsdk:"application_family"`
	ApplicationGroup  types.String            `tfsdk:"application_group"`
	TrafficClass      types.String            `tfsdk:"traffic_class"`
	BusinessRelevance types.String            `tfsdk:"business_relevance"`
}

type CustomApplicationL3l4 struct {
	IpAddresses types.Set    `tfsdk:"ip_addresses"`
	L4Protocol  types.String `tfsdk:"l4_protocol"`
	Ports       types.String `tfsdk:"ports"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath
func (data CustomApplication) getPath() string {
	return "/template/policy/customapp/"
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody
func (data CustomApplication) toBody(ctx context.Context) string {
	body := ""
	if !data.AppName.IsNull() {
		body, _ = sjson.Set(body, "appName", data.AppName.ValueString())
	}
	if !data.ServerNames.IsNull() {
		var values []string
		data.ServerNames.ElementsAs(ctx, &values, false)
		body, _ = sjson.Set(body, "serverNames", values)
	}
	if true {
		body, _ = sjson.Set(body, "L3L4", []interface{}{})
		for _, item := range data.L3l4 {
			itemBody := ""
			if !item.IpAddresses.IsNull() {
				var values []string
				item.IpAddresses.ElementsAs(ctx, &values, false)
				itemBody, _ = sjson.Set(itemBody, "ipAddresses", values)
			}
			if !item.L4Protocol.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "l4Protocol", item.L4Protocol.ValueString())
			}
			if !item.Ports.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "ports", item.Ports.ValueString())
			}
			body, _ = sjson.SetRaw(body, "L3L4.-1", itemBody)
		}
	}
	if !data.ApplicationFamily.IsNull() {
		body, _ = sjson.Set(body, "attributes.application-family", data.ApplicationFamily.ValueString())
	}
	if !data.ApplicationGroup.IsNull() {
		body, _ = sjson.Set(body, "attributes.application-group", data.ApplicationGroup.ValueString())
	}
	if !data.TrafficClass.IsNull() {
		body, _ = sjson.Set(body, "attributes.traffic-class", data.TrafficClass.ValueString())
	}
	if !data.BusinessRelevance.IsNull() {
		body, _ = sjson.Set(body, "attributes.business-relevance", data.BusinessRelevance.ValueString())
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody
func (data *CustomApplication) fromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("data.appName"); value.Exists() {
		data.AppName = types.StringValue(value.String())
	} else {
		data.AppName = types.StringNull()
	}
	if value := res.Get("data.serverNames"); value.Exists() {
		data.ServerNames = helpers.GetStringSet(value.Array())
	} else {
		data.ServerNames = types.SetNull(types.StringType)
	}
	if value := res.Get("data.L3L4"); value.Exists() && len(value.Array()) > 0 {
		data.L3l4 = make([]CustomApplicationL3l4, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := CustomApplicationL3l4{}
			if cValue := v.Get("ipAddresses"); cValue.Exists() {
				item.IpAddresses = helpers.GetStringSet(cValue.Array())
			} else {
				item.IpAddresses = types.SetNull(types.StringType)
			}
			if cValue := v.Get("l4Protocol"); cValue.Exists() {
				item.L4Protocol = types.StringValue(cValue.String())
			} else {
				item.L4Protocol = types.StringNull()
			}
			if cValue := v.Get("ports"); cValue.Exists() {
				item.Ports = types.StringValue(cValue.String())
			} else {
				item.Ports = types.StringNull()
			}
			data.L3l4 = append(data.L3l4, item)
			return true
		})
	} else {
		if len(data.L3l4) > 0 {
			data.L3l4 = []CustomApplicationL3l4{}
		}
	}
	if value := res.Get("data.attributes.application-family"); value.Exists() {
		data.ApplicationFamily = types.StringValue(value.String())
	} else {
		data.ApplicationFamily = types.StringNull()
	}
	if value := res.Get("data.attributes.application-group"); value.Exists() {
		data.ApplicationGroup = types.StringValue(value.String())
	} else {
		data.ApplicationGroup = types.StringNull()
	}
	if value := res.Get("data.attributes.traffic-class"); value.Exists() {
		data.TrafficClass = types.StringValue(value.String())
	} else {
		data.TrafficClass = types.StringNull()
	}
	if value := res.Get("data.attributes.business-relevance"); value.Exists() {
		data.BusinessRelevance = types.StringValue(value.String())
	} else {
		data.BusinessRelevance = types.StringNull()
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin hasChanges
func (data *CustomApplication) hasChanges(ctx context.Context, state *CustomApplication) bool {
	hasChanges := false
	if !data.AppName.Equal(state.AppName) {
		hasChanges = true
	}
	if !data.ServerNames.Equal(state.ServerNames) {
		hasChanges = true
	}
	if len(data.L3l4) != len(state.L3l4) {
		hasChanges = true
	} else {
		for i := range data.L3l4 {
			if !data.L3l4[i].IpAddresses.Equal(state.L3l4[i].IpAddresses) {
				hasChanges = true
			}
			if !data.L3l4[i].L4Protocol.Equal(state.L3l4[i].L4Protocol) {
				hasChanges = true
			}
			if !data.L3l4[i].Ports.Equal(state.L3l4[i].Ports) {
				hasChanges = true
			}
		}
	}
	if !data.ApplicationFamily.Equal(state.ApplicationFamily) {
		hasChanges = true
	}
	if !data.ApplicationGroup.Equal(state.ApplicationGroup) {
		hasChanges = true
	}
	if !data.TrafficClass.Equal(state.TrafficClass) {
		hasChanges = true
	}
	if !data.BusinessRelevance.Equal(state.BusinessRelevance) {
		hasChanges = true
	}
	return hasChanges
}

// End of section. //template:end hasChanges

// Section below is generated&owned by "gen/generator.go". //template:begin updateVersions

// End of section. //template:end updateVersions

// Section below is generated&owned by "gen/generator.go". //template:begin processImport
func (data *CustomApplication) processImport(ctx context.Context) {
	data.Version = types.Int64Value(0)
}

// End of section. //template:end processImport

// Section below is generated&owned by "gen/generator.go". //template:begin applyFilters
// End of section. //template:end applyFilters
