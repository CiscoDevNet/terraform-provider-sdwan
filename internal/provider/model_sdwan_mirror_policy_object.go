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

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types
type MirrorPolicyObject struct {
	Id                  types.String `tfsdk:"id"`
	Version             types.Int64  `tfsdk:"version"`
	Name                types.String `tfsdk:"name"`
	RemoteDestinationIp types.String `tfsdk:"remote_destination_ip"`
	SourceIp            types.String `tfsdk:"source_ip"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath
func (data MirrorPolicyObject) getPath() string {
	return "/template/policy/list/mirror/"
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody
func (data MirrorPolicyObject) toBody(ctx context.Context) string {
	body := ""
	if true {
		body, _ = sjson.Set(body, "type", "mirror")
	}
	if !data.Name.IsNull() {
		body, _ = sjson.Set(body, "name", data.Name.ValueString())
	}
	if !data.RemoteDestinationIp.IsNull() {
		body, _ = sjson.Set(body, "entries.0.remoteDest", data.RemoteDestinationIp.ValueString())
	}
	if !data.SourceIp.IsNull() {
		body, _ = sjson.Set(body, "entries.0.source", data.SourceIp.ValueString())
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody
func (data *MirrorPolicyObject) fromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("name"); value.Exists() {
		data.Name = types.StringValue(value.String())
	} else {
		data.Name = types.StringNull()
	}
	if value := res.Get("entries.0.remoteDest"); value.Exists() {
		data.RemoteDestinationIp = types.StringValue(value.String())
	} else {
		data.RemoteDestinationIp = types.StringNull()
	}
	if value := res.Get("entries.0.source"); value.Exists() {
		data.SourceIp = types.StringValue(value.String())
	} else {
		data.SourceIp = types.StringNull()
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin hasChanges
func (data *MirrorPolicyObject) hasChanges(ctx context.Context, state *MirrorPolicyObject) bool {
	hasChanges := false
	if !data.Name.Equal(state.Name) {
		hasChanges = true
	}
	if !data.RemoteDestinationIp.Equal(state.RemoteDestinationIp) {
		hasChanges = true
	}
	if !data.SourceIp.Equal(state.SourceIp) {
		hasChanges = true
	}
	return hasChanges
}

// End of section. //template:end hasChanges

// Section below is generated&owned by "gen/generator.go". //template:begin updateVersions

// End of section. //template:end updateVersions
