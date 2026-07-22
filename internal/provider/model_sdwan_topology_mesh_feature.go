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
type TopologyMesh struct {
	Id               types.String `tfsdk:"id"`
	Version          types.Int64  `tfsdk:"version"`
	Name             types.String `tfsdk:"name"`
	Description      types.String `tfsdk:"description"`
	FeatureProfileId types.String `tfsdk:"feature_profile_id"`
	TargetVpns       types.Set    `tfsdk:"target_vpns"`
	Sites            types.Set    `tfsdk:"sites"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getModel
func (data TopologyMesh) getModel() string {
	return "topology_mesh"
}

// End of section. //template:end getModel

// Section below is generated&owned by "gen/generator.go". //template:begin getPath
func (data TopologyMesh) getPath() string {
	return fmt.Sprintf("/v1/feature-profile/sdwan/topology/%v/mesh", url.QueryEscape(data.FeatureProfileId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody
func (data TopologyMesh) toBody(ctx context.Context) string {
	body := ""
	body, _ = sjson.Set(body, "name", data.Name.ValueString())
	body, _ = sjson.Set(body, "description", data.Description.ValueString())
	path := "data."
	if !data.TargetVpns.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"target.vpn.optionType", "global")
			var values []string
			data.TargetVpns.ElementsAs(ctx, &values, false)
			body, _ = sjson.Set(body, path+"target.vpn.value", values)
		}
	}
	if !data.Sites.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"sites.optionType", "global")
			var values []string
			data.Sites.ElementsAs(ctx, &values, false)
			body, _ = sjson.Set(body, path+"sites.value", values)
		}
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody
func (data *TopologyMesh) fromBody(ctx context.Context, res gjson.Result, fullRead bool) {
	data.Name = types.StringValue(res.Get("payload.name").String())
	if value := res.Get("payload.description"); value.Exists() && value.String() != "" {
		data.Description = types.StringValue(value.String())
	} else {
		data.Description = types.StringNull()
	}
	path := "payload.data."
	data.TargetVpns = types.SetNull(types.StringType)

	if t := res.Get(path + "target.vpn.optionType"); t.Exists() {
		va := res.Get(path + "target.vpn.value")
		if t.String() == "global" {
			data.TargetVpns = helpers.GetStringSet(va.Array())
		}
	}
	data.Sites = types.SetNull(types.StringType)

	if t := res.Get(path + "sites.optionType"); t.Exists() {
		va := res.Get(path + "sites.value")
		if t.String() == "global" {
			data.Sites = helpers.GetStringSet(va.Array())
		}
	}
}

// End of section. //template:end fromBody
