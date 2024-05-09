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
type SystemMRF struct {
	Id                        types.String `tfsdk:"id"`
	Version                   types.Int64  `tfsdk:"version"`
	Name                      types.String `tfsdk:"name"`
	Description               types.String `tfsdk:"description"`
	FeatureProfileId          types.String `tfsdk:"feature_profile_id"`
	RegionId                  types.Int64  `tfsdk:"region_id"`
	SecondaryRegionId         types.Int64  `tfsdk:"secondary_region_id"`
	SecondaryRegionIdVariable types.String `tfsdk:"secondary_region_id_variable"`
	Role                      types.String `tfsdk:"role"`
	RoleVariable              types.String `tfsdk:"role_variable"`
	EnableMigrationToMrf      types.String `tfsdk:"enable_migration_to_mrf"`
	MigrationBgpCommunity     types.Int64  `tfsdk:"migration_bgp_community"`
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
	if data.RegionId.IsNull() {
		body, _ = sjson.Set(body, path+"regionId.optionType", "default")

	} else {
		body, _ = sjson.Set(body, path+"regionId.optionType", "global")
		body, _ = sjson.Set(body, path+"regionId.value", data.RegionId.ValueInt64())
	}

	if !data.SecondaryRegionIdVariable.IsNull() {
		body, _ = sjson.Set(body, path+"secondaryRegion.optionType", "variable")
		body, _ = sjson.Set(body, path+"secondaryRegion.value", data.SecondaryRegionIdVariable.ValueString())
	} else if data.SecondaryRegionId.IsNull() {
		body, _ = sjson.Set(body, path+"secondaryRegion.optionType", "default")

	} else {
		body, _ = sjson.Set(body, path+"secondaryRegion.optionType", "global")
		body, _ = sjson.Set(body, path+"secondaryRegion.value", data.SecondaryRegionId.ValueInt64())
	}

	if !data.RoleVariable.IsNull() {
		body, _ = sjson.Set(body, path+"role.optionType", "variable")
		body, _ = sjson.Set(body, path+"role.value", data.RoleVariable.ValueString())
	} else if data.Role.IsNull() {
		body, _ = sjson.Set(body, path+"role.optionType", "default")

	} else {
		body, _ = sjson.Set(body, path+"role.optionType", "global")
		body, _ = sjson.Set(body, path+"role.value", data.Role.ValueString())
	}
	if data.EnableMigrationToMrf.IsNull() {
		body, _ = sjson.Set(body, path+"enableMrfMigration.optionType", "default")

	} else {
		body, _ = sjson.Set(body, path+"enableMrfMigration.optionType", "global")
		body, _ = sjson.Set(body, path+"enableMrfMigration.value", data.EnableMigrationToMrf.ValueString())
	}
	if data.MigrationBgpCommunity.IsNull() {
		body, _ = sjson.Set(body, path+"migrationBgpCommunity.optionType", "default")

	} else {
		body, _ = sjson.Set(body, path+"migrationBgpCommunity.optionType", "global")
		body, _ = sjson.Set(body, path+"migrationBgpCommunity.value", data.MigrationBgpCommunity.ValueInt64())
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody
func (data *SystemMRF) fromBody(ctx context.Context, res gjson.Result) {
	data.Name = types.StringValue(res.Get("payload.name").String())
	if value := res.Get("payload.description"); value.Exists() && value.String() != "" {
		data.Description = types.StringValue(value.String())
	} else {
		data.Description = types.StringNull()
	}
	path := "payload.data."
	data.RegionId = types.Int64Null()

	if t := res.Get(path + "regionId.optionType"); t.Exists() {
		va := res.Get(path + "regionId.value")
		if t.String() == "global" {
			data.RegionId = types.Int64Value(va.Int())
		}
	}
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
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin updateFromBody
func (data *SystemMRF) updateFromBody(ctx context.Context, res gjson.Result) {
	data.Name = types.StringValue(res.Get("payload.name").String())
	if value := res.Get("payload.description"); value.Exists() && value.String() != "" {
		data.Description = types.StringValue(value.String())
	} else {
		data.Description = types.StringNull()
	}
	path := "payload.data."
	data.RegionId = types.Int64Null()

	if t := res.Get(path + "regionId.optionType"); t.Exists() {
		va := res.Get(path + "regionId.value")
		if t.String() == "global" {
			data.RegionId = types.Int64Value(va.Int())
		}
	}
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
}

// End of section. //template:end updateFromBody

// Section below is generated&owned by "gen/generator.go". //template:begin isNull
func (data *SystemMRF) isNull(ctx context.Context, res gjson.Result) bool {
	if !data.FeatureProfileId.IsNull() {
		return false
	}
	if !data.RegionId.IsNull() {
		return false
	}
	if !data.SecondaryRegionId.IsNull() {
		return false
	}
	if !data.SecondaryRegionIdVariable.IsNull() {
		return false
	}
	if !data.Role.IsNull() {
		return false
	}
	if !data.RoleVariable.IsNull() {
		return false
	}
	if !data.EnableMigrationToMrf.IsNull() {
		return false
	}
	if !data.MigrationBgpCommunity.IsNull() {
		return false
	}
	return true
}

// End of section. //template:end isNull
