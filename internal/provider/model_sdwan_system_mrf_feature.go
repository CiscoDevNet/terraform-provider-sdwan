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
type SystemMRF struct {
	Id                             types.String `tfsdk:"id"`
	Version                        types.Int64  `tfsdk:"version"`
	Name                           types.String `tfsdk:"name"`
	Description                    types.String `tfsdk:"description"`
	FeatureProfileId               types.String `tfsdk:"feature_profile_id"`
	SecondaryRegionId              types.Int64  `tfsdk:"secondary_region_id"`
	SecondaryRegionIdVariable      types.String `tfsdk:"secondary_region_id_variable"`
	Role                           types.String `tfsdk:"role"`
	RoleVariable                   types.String `tfsdk:"role_variable"`
	EnableMigrationToMrf           types.String `tfsdk:"enable_migration_to_mrf"`
	MigrationBgpCommunity          types.Int64  `tfsdk:"migration_bgp_community"`
	EnableManagementRegion         types.Bool   `tfsdk:"enable_management_region"`
	EnableManagementRegionVariable types.String `tfsdk:"enable_management_region_variable"`
	VrfId                          types.Int64  `tfsdk:"vrf_id"`
	VrfIdVariable                  types.String `tfsdk:"vrf_id_variable"`
	GatewayPreference              types.Set    `tfsdk:"gateway_preference"`
	GatewayPreferenceVariable      types.String `tfsdk:"gateway_preference_variable"`
	ManagementGateway              types.Bool   `tfsdk:"management_gateway"`
	ManagementGatewayVariable      types.String `tfsdk:"management_gateway_variable"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getModel
func (data SystemMRF) getModel() string {
	return "system_mrf"
}

// End of section. //template:end getModel

// Section below is generated&owned by "gen/generator.go". //template:begin getPath
func (data SystemMRF) getPath() string {
	return fmt.Sprintf("/v1/feature-profile/sdwan/system/%v/mrf", url.QueryEscape(data.FeatureProfileId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody
func (data SystemMRF) toBody(ctx context.Context) string {
	body := ""
	body, _ = sjson.Set(body, "name", data.Name.ValueString())
	body, _ = sjson.Set(body, "description", data.Description.ValueString())
	path := "data."

	if !data.SecondaryRegionIdVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"secondaryRegion.optionType", "variable")
			body, _ = sjson.Set(body, path+"secondaryRegion.value", data.SecondaryRegionIdVariable.ValueString())
		}
	} else if data.SecondaryRegionId.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"secondaryRegion.optionType", "default")

		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"secondaryRegion.optionType", "global")
			body, _ = sjson.Set(body, path+"secondaryRegion.value", data.SecondaryRegionId.ValueInt64())
		}
	}

	if !data.RoleVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"role.optionType", "variable")
			body, _ = sjson.Set(body, path+"role.value", data.RoleVariable.ValueString())
		}
	} else if data.Role.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"role.optionType", "default")

		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"role.optionType", "global")
			body, _ = sjson.Set(body, path+"role.value", data.Role.ValueString())
		}
	}
	if data.EnableMigrationToMrf.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"enableMrfMigration.optionType", "default")

		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"enableMrfMigration.optionType", "global")
			body, _ = sjson.Set(body, path+"enableMrfMigration.value", data.EnableMigrationToMrf.ValueString())
		}
	}
	if data.MigrationBgpCommunity.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"migrationBgpCommunity.optionType", "default")

		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"migrationBgpCommunity.optionType", "global")
			body, _ = sjson.Set(body, path+"migrationBgpCommunity.value", data.MigrationBgpCommunity.ValueInt64())
		}
	}

	if !data.EnableManagementRegionVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"enableManagementRegion.optionType", "variable")
			body, _ = sjson.Set(body, path+"enableManagementRegion.value", data.EnableManagementRegionVariable.ValueString())
		}
	} else if data.EnableManagementRegion.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"enableManagementRegion.optionType", "default")
			body, _ = sjson.Set(body, path+"enableManagementRegion.value", false)
		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"enableManagementRegion.optionType", "global")
			body, _ = sjson.Set(body, path+"enableManagementRegion.value", data.EnableManagementRegion.ValueBool())
		}
	}

	if !data.VrfIdVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"managementRegion.vrfId.optionType", "variable")
			body, _ = sjson.Set(body, path+"managementRegion.vrfId.value", data.VrfIdVariable.ValueString())
		}
	} else if data.VrfId.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"managementRegion.vrfId.optionType", "default")

		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"managementRegion.vrfId.optionType", "global")
			body, _ = sjson.Set(body, path+"managementRegion.vrfId.value", data.VrfId.ValueInt64())
		}
	}

	if !data.GatewayPreferenceVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"managementRegion.gatewayPreference.optionType", "variable")
			body, _ = sjson.Set(body, path+"managementRegion.gatewayPreference.value", data.GatewayPreferenceVariable.ValueString())
		}
	} else if data.GatewayPreference.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"managementRegion.gatewayPreference.optionType", "default")

		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"managementRegion.gatewayPreference.optionType", "global")
			var values []int64
			data.GatewayPreference.ElementsAs(ctx, &values, false)
			body, _ = sjson.Set(body, path+"managementRegion.gatewayPreference.value", values)
		}
	}

	if !data.ManagementGatewayVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"managementRegion.managementGateway.optionType", "variable")
			body, _ = sjson.Set(body, path+"managementRegion.managementGateway.value", data.ManagementGatewayVariable.ValueString())
		}
	} else if data.ManagementGateway.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"managementRegion.managementGateway.optionType", "default")
			body, _ = sjson.Set(body, path+"managementRegion.managementGateway.value", false)
		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"managementRegion.managementGateway.optionType", "global")
			body, _ = sjson.Set(body, path+"managementRegion.managementGateway.value", data.ManagementGateway.ValueBool())
		}
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody
func (data *SystemMRF) fromBody(ctx context.Context, res gjson.Result, fullRead bool) {
	data.Name = types.StringValue(res.Get("payload.name").String())
	if value := res.Get("payload.description"); value.Exists() && value.String() != "" {
		data.Description = types.StringValue(value.String())
	} else {
		data.Description = types.StringNull()
	}
	path := "payload.data."
	data.SecondaryRegionId = types.Int64Null()
	data.SecondaryRegionIdVariable = types.StringNull()
	if t := res.Get(path + "secondaryRegion.optionType"); t.Exists() {
		va := res.Get(path + "secondaryRegion.value")
		if t.String() == "variable" {
			data.SecondaryRegionIdVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.SecondaryRegionId = types.Int64Value(va.Int())
		}
	}
	data.Role = types.StringNull()
	data.RoleVariable = types.StringNull()
	if t := res.Get(path + "role.optionType"); t.Exists() {
		va := res.Get(path + "role.value")
		if t.String() == "variable" {
			data.RoleVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.Role = types.StringValue(va.String())
		}
	}
	data.EnableMigrationToMrf = types.StringNull()

	if t := res.Get(path + "enableMrfMigration.optionType"); t.Exists() {
		va := res.Get(path + "enableMrfMigration.value")
		if t.String() == "global" {
			data.EnableMigrationToMrf = types.StringValue(va.String())
		}
	}
	data.MigrationBgpCommunity = types.Int64Null()

	if t := res.Get(path + "migrationBgpCommunity.optionType"); t.Exists() {
		va := res.Get(path + "migrationBgpCommunity.value")
		if t.String() == "global" {
			data.MigrationBgpCommunity = types.Int64Value(va.Int())
		}
	}
	data.EnableManagementRegion = types.BoolNull()
	data.EnableManagementRegionVariable = types.StringNull()
	if t := res.Get(path + "enableManagementRegion.optionType"); t.Exists() {
		va := res.Get(path + "enableManagementRegion.value")
		if t.String() == "variable" {
			data.EnableManagementRegionVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.EnableManagementRegion = types.BoolValue(va.Bool())
		}
	}
	data.VrfId = types.Int64Null()
	data.VrfIdVariable = types.StringNull()
	if t := res.Get(path + "managementRegion.vrfId.optionType"); t.Exists() {
		va := res.Get(path + "managementRegion.vrfId.value")
		if t.String() == "variable" {
			data.VrfIdVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.VrfId = types.Int64Value(va.Int())
		}
	}
	data.GatewayPreference = types.SetNull(types.Int64Type)
	data.GatewayPreferenceVariable = types.StringNull()
	if t := res.Get(path + "managementRegion.gatewayPreference.optionType"); t.Exists() {
		va := res.Get(path + "managementRegion.gatewayPreference.value")
		if t.String() == "variable" {
			data.GatewayPreferenceVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.GatewayPreference = helpers.GetInt64Set(va.Array())
		}
	}
	data.ManagementGateway = types.BoolNull()
	data.ManagementGatewayVariable = types.StringNull()
	if t := res.Get(path + "managementRegion.managementGateway.optionType"); t.Exists() {
		va := res.Get(path + "managementRegion.managementGateway.value")
		if t.String() == "variable" {
			data.ManagementGatewayVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.ManagementGateway = types.BoolValue(va.Bool())
		}
	}
}

// End of section. //template:end fromBody
