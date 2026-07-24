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
type Scope struct {
	Id          types.String   `tfsdk:"id"`
	Name        types.String   `tfsdk:"name"`
	Description types.String   `tfsdk:"description"`
	Objects     []ScopeObjects `tfsdk:"objects"`
	Users       types.Set      `tfsdk:"users"`
}

type ScopeObjects struct {
	ObjectType types.String `tfsdk:"object_type"`
	ObjectIds  types.Set    `tfsdk:"object_ids"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath
func (data Scope) getPath() string {
	return "/v1/resource-domain/"
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody
func (data Scope) toBody(ctx context.Context) string {
	body := ""
	if !data.Name.IsNull() {
		body, _ = sjson.Set(body, "name", data.Name.ValueString())
	}
	if !data.Description.IsNull() {
		body, _ = sjson.Set(body, "description", data.Description.ValueString())
	}
	if true {
		body, _ = sjson.Set(body, "objects", []interface{}{})
		for _, item := range data.Objects {
			itemBody := ""
			if !item.ObjectType.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "objectType", item.ObjectType.ValueString())
			}
			if !item.ObjectIds.IsNull() {
				var values []string
				item.ObjectIds.ElementsAs(ctx, &values, false)
				itemBody, _ = sjson.Set(itemBody, "objectIds", values)
			}
			body, _ = sjson.SetRaw(body, "objects.-1", itemBody)
		}
	}
	if !data.Users.IsNull() {
		var values []string
		data.Users.ElementsAs(ctx, &values, false)
		body, _ = sjson.Set(body, "users", values)
	}
	return body
}

// End of section. //template:end toBody

// fromBody is hand-maintained (not generated) because the resource-domain API's GET response
// shape does not match what POST/PUT accept:
//   - "objects" always lists every enum objectType (even unused ones), each wrapped as
//     {"objectType":..,"object":[{...full detail...}]} instead of the flat {"objectType":..,
//     "objectIds":[...]} sent on write, and the id key inside "object" differs per type
//     (e.g. "id" for config-group, "profileId" for feature-profile).
//   - "users" returns full user records instead of the flat usernames sent on write.
//
// Note: the controller silently attaches a default "feature-profile" association (the built-in
// "Default_Policy_Object_Profile") to every scope that was never explicitly requested. This is
// intentionally NOT filtered out here (matching by name would be fragile) - callers that want a
// diff-free plan should declare that association explicitly in their scope's "objects" config.
func (data *Scope) fromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("name"); value.Exists() {
		data.Name = types.StringValue(value.String())
	} else {
		data.Name = types.StringNull()
	}
	if value := res.Get("description"); value.Exists() {
		data.Description = types.StringValue(value.String())
	} else {
		data.Description = types.StringNull()
	}
	if value := res.Get("objects"); value.Exists() {
		objects := make([]ScopeObjects, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			objectType := v.Get("objectType").String()
			var ids []string
			v.Get("object").ForEach(func(k, item gjson.Result) bool {
				if id := item.Get("id"); id.Exists() {
					ids = append(ids, id.String())
				} else if id := item.Get("profileId"); id.Exists() {
					ids = append(ids, id.String())
				}
				return true
			})
			if len(ids) == 0 {
				return true
			}
			objectIds, _ := types.SetValueFrom(ctx, types.StringType, ids)
			objects = append(objects, ScopeObjects{
				ObjectType: types.StringValue(objectType),
				ObjectIds:  objectIds,
			})
			return true
		})
		if len(objects) > 0 {
			data.Objects = objects
		} else if len(data.Objects) > 0 {
			data.Objects = []ScopeObjects{}
		}
	} else if len(data.Objects) > 0 {
		data.Objects = []ScopeObjects{}
	}
	if value := res.Get("users"); value.Exists() {
		names := make([]string, 0, len(value.Array()))
		value.ForEach(func(k, v gjson.Result) bool {
			if un := v.Get("userName"); un.Exists() {
				names = append(names, un.String())
			} else {
				names = append(names, v.String())
			}
			return true
		})
		if len(names) > 0 {
			data.Users, _ = types.SetValueFrom(ctx, types.StringType, names)
		} else {
			data.Users = types.SetNull(types.StringType)
		}
	} else {
		data.Users = types.SetNull(types.StringType)
	}
}

// Section below is generated&owned by "gen/generator.go". //template:begin hasChanges
func (data *Scope) hasChanges(ctx context.Context, state *Scope) bool {
	hasChanges := false
	if !data.Name.Equal(state.Name) {
		hasChanges = true
	}
	if !data.Description.Equal(state.Description) {
		hasChanges = true
	}
	if len(data.Objects) != len(state.Objects) {
		hasChanges = true
	} else {
		for i := range data.Objects {
			if !data.Objects[i].ObjectType.Equal(state.Objects[i].ObjectType) {
				hasChanges = true
			}
			if !data.Objects[i].ObjectIds.Equal(state.Objects[i].ObjectIds) {
				hasChanges = true
			}
		}
	}
	if !data.Users.Equal(state.Users) {
		hasChanges = true
	}
	return hasChanges
}

// End of section. //template:end hasChanges

// Section below is generated&owned by "gen/generator.go". //template:begin updateVersions

// End of section. //template:end updateVersions

// Section below is generated&owned by "gen/generator.go". //template:begin processImport
func (data *Scope) processImport(ctx context.Context) {
}

// End of section. //template:end processImport

// Section below is generated&owned by "gen/generator.go". //template:begin applyFilters
// End of section. //template:end applyFilters
