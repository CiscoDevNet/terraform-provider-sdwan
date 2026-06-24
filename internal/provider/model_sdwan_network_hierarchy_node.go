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

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// NetworkHierarchyNode represents a node in the network hierarchy
// This can be a Group (Area), Region, or Site
type NetworkHierarchyNode struct {
	Id          types.String                 `tfsdk:"id"`
	Name        types.String                 `tfsdk:"name"`
	Description types.String                 `tfsdk:"description"`
	ParentGroup types.String                 `tfsdk:"parent_group"`
	Type        types.String                 `tfsdk:"type"`
	SiteId      types.Int64                  `tfsdk:"site_id"`
	IsSecondary types.Bool                   `tfsdk:"is_secondary"`
	Address     *NetworkHierarchyNodeAddress `tfsdk:"address"`
	Controllers types.Set                    `tfsdk:"controllers"`
}

// NetworkHierarchyNodeAddress represents the address for a site node
type NetworkHierarchyNodeAddress struct {
	Street  types.String `tfsdk:"street"`
	City    types.String `tfsdk:"city"`
	State   types.String `tfsdk:"state"`
	Country types.String `tfsdk:"country"`
	Zipcode types.String `tfsdk:"zipcode"`
}

func (data NetworkHierarchyNode) getPath() string {
	return "/v1/network-hierarchy/"
}

func (data NetworkHierarchyNode) getHierarchyListPath() string {
	return "/v1/network-hierarchy?offset=0&limit=1500"
}

// resolveParentGroupToId finds the UUID for a group by name from the hierarchy list.
// Returns the UUID and an error message if not found.
func (data NetworkHierarchyNode) resolveParentGroupToId(hierarchyRes gjson.Result) (string, string) {
	parentGroupName := data.ParentGroup.ValueString()

	var foundId string
	var matchCount int

	hierarchyRes.ForEach(func(key, value gjson.Result) bool {
		nodeName := value.Get("name").String()
		nodeLabel := value.Get("data.label").String()

		// Match by name - must be either Global (no label) or a group (no label means group)
		// Groups don't have data.label, regions have "REGION", sites have "SITE"
		if nodeName == parentGroupName && nodeLabel != "REGION" && nodeLabel != "SITE" {
			// Use "id" field which exists for all nodes including Global
			foundId = value.Get("id").String()
			matchCount++
		}
		return true
	})

	if matchCount == 0 {
		return "", fmt.Sprintf("Parent group '%s' not found in network hierarchy", parentGroupName)
	}
	if matchCount > 1 {
		return "", fmt.Sprintf("Multiple groups found with name '%s'. Group names must be unique.", parentGroupName)
	}

	return foundId, ""
}

// resolveParentIdToGroup finds the group name for a given parent UUID from the hierarchy list.
func (data NetworkHierarchyNode) resolveParentIdToGroup(hierarchyRes gjson.Result, parentId string) string {
	var parentName string

	hierarchyRes.ForEach(func(key, value gjson.Result) bool {
		nodeId := value.Get("id").String()
		if nodeId == parentId {
			parentName = value.Get("name").String()
			return false // stop iteration
		}
		return true
	})

	return parentName
}

func (data NetworkHierarchyNode) toBody(ctx context.Context, parentId string) string {
	body := ""

	if !data.Name.IsNull() {
		body, _ = sjson.Set(body, "name", data.Name.ValueString())
	}
	if !data.Description.IsNull() {
		body, _ = sjson.Set(body, "description", data.Description.ValueString())
	}
	if parentId != "" {
		body, _ = sjson.Set(body, "data.parentUuid", parentId)
	}

	nodeType := data.Type.ValueString()
	switch nodeType {
	case "region":
		body, _ = sjson.Set(body, "data.label", "REGION")
		if !data.IsSecondary.IsNull() {
			body, _ = sjson.Set(body, "data.isSecondary", data.IsSecondary.ValueBool())
		} else {
			body, _ = sjson.Set(body, "data.isSecondary", false)
		}
	case "site":
		body, _ = sjson.Set(body, "data.label", "SITE")
		if !data.SiteId.IsNull() {
			body, _ = sjson.Set(body, "data.hierarchyId.siteId", data.SiteId.ValueInt64())
		}
		if data.Address != nil {
			if !data.Address.Street.IsNull() {
				body, _ = sjson.Set(body, "data.address.street", data.Address.Street.ValueString())
			}
			if !data.Address.City.IsNull() {
				body, _ = sjson.Set(body, "data.address.city", data.Address.City.ValueString())
			}
			if !data.Address.State.IsNull() {
				body, _ = sjson.Set(body, "data.address.state", data.Address.State.ValueString())
			}
			if !data.Address.Country.IsNull() {
				body, _ = sjson.Set(body, "data.address.country", data.Address.Country.ValueString())
			}
			if !data.Address.Zipcode.IsNull() {
				body, _ = sjson.Set(body, "data.address.zipcode", data.Address.Zipcode.ValueString())
			}
		}
	}

	return body
}

func (data NetworkHierarchyNode) toUpdateBody(ctx context.Context, currentNodeRes gjson.Result, parentId string) string {
	body := data.toBody(ctx, parentId)

	nodeType := data.Type.ValueString()
	switch nodeType {
	case "region":
		if regionId := currentNodeRes.Get("data.hierarchyId.regionId"); regionId.Exists() {
			body, _ = sjson.Set(body, "data.hierarchyId.regionId", regionId.Int())
		}
	case "site":
		if siteId := currentNodeRes.Get("data.hierarchyId.siteId"); siteId.Exists() {
			body, _ = sjson.Set(body, "data.hierarchyId.siteId", siteId.Int())
		}
	}

	return body
}

func (data NetworkHierarchyNode) getControllersPath() string {
	return "/system/device/controllers"
}

func (data NetworkHierarchyNode) toWorkflowBody(ctx context.Context, regionId int64) string {
	body := ""
	body, _ = sjson.Set(body, "handler", "MRF")

	nodeBody := ""
	nodeBody, _ = sjson.Set(nodeBody, "uuid", data.Id.ValueString())
	nodeBody, _ = sjson.Set(nodeBody, "name", data.Name.ValueString())
	nodeBody, _ = sjson.Set(nodeBody, "type", "Region")

	if !data.Description.IsNull() {
		nodeBody, _ = sjson.Set(nodeBody, "description", data.Description.ValueString())
	} else {
		nodeBody, _ = sjson.Set(nodeBody, "description", "")
	}

	nodeBody, _ = sjson.Set(nodeBody, "data.label", "REGION")
	nodeBody, _ = sjson.Set(nodeBody, "data.regionId", regionId)

	if !data.IsSecondary.IsNull() {
		nodeBody, _ = sjson.Set(nodeBody, "data.isSecondary", data.IsSecondary.ValueBool())
	} else {
		nodeBody, _ = sjson.Set(nodeBody, "data.isSecondary", false)
	}

	// Set controllers array
	if !data.Controllers.IsNull() && !data.Controllers.IsUnknown() {
		var controllers []string
		data.Controllers.ElementsAs(ctx, &controllers, false)
		nodeBody, _ = sjson.Set(nodeBody, "data.controllers", controllers)
	} else {
		nodeBody, _ = sjson.Set(nodeBody, "data.controllers", []string{})
	}

	body, _ = sjson.SetRaw(body, "nodes.0", nodeBody)
	return body
}

func (data *NetworkHierarchyNode) readControllersFromResponse(ctx context.Context, controllersRes gjson.Result, regionId int64) {
	regionIdStr := strconv.FormatInt(regionId, 10)
	var controllerUUIDs []attr.Value

	controllersRes.Get("data").ForEach(func(key, value gjson.Result) bool {
		// Only vsmart controllers have region-id
		if value.Get("deviceType").String() != "vsmart" {
			return true
		}

		regionIds := value.Get("region-id").String()
		if regionIds == "" {
			return true
		}

		// region-id is comma-separated like "1,4" or "2"
		for _, rid := range strings.Split(regionIds, ",") {
			if strings.TrimSpace(rid) == regionIdStr {
				uuid := value.Get("uuid").String()
				if uuid != "" {
					controllerUUIDs = append(controllerUUIDs, types.StringValue(uuid))
				}
				break
			}
		}
		return true
	})

	if len(controllerUUIDs) > 0 {
		data.Controllers, _ = types.SetValue(types.StringType, controllerUUIDs)
	} else {
		data.Controllers = types.SetNull(types.StringType)
	}
}

func (data *NetworkHierarchyNode) fromBody(ctx context.Context, res gjson.Result) string {
	var parentUuid string

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
	if value := res.Get("data.parentUuid"); value.Exists() {
		parentUuid = value.String()
	}

	if value := res.Get("data.label"); value.Exists() {
		label := value.String()
		switch label {
		case "REGION":
			data.Type = types.StringValue("region")
			if value := res.Get("data.isSecondary"); value.Exists() {
				data.IsSecondary = types.BoolValue(value.Bool())
			} else {
				data.IsSecondary = types.BoolNull()
			}
		case "SITE":
			data.Type = types.StringValue("site")
			if value := res.Get("data.hierarchyId.siteId"); value.Exists() {
				data.SiteId = types.Int64Value(value.Int())
			} else {
				data.SiteId = types.Int64Null()
			}
			if res.Get("data.address").Exists() {
				data.Address = &NetworkHierarchyNodeAddress{}
				if value := res.Get("data.address.street"); value.Exists() {
					data.Address.Street = types.StringValue(value.String())
				} else {
					data.Address.Street = types.StringNull()
				}
				if value := res.Get("data.address.city"); value.Exists() {
					data.Address.City = types.StringValue(value.String())
				} else {
					data.Address.City = types.StringNull()
				}
				if value := res.Get("data.address.state"); value.Exists() {
					data.Address.State = types.StringValue(value.String())
				} else {
					data.Address.State = types.StringNull()
				}
				if value := res.Get("data.address.country"); value.Exists() {
					data.Address.Country = types.StringValue(value.String())
				} else {
					data.Address.Country = types.StringNull()
				}
				if value := res.Get("data.address.zipcode"); value.Exists() {
					data.Address.Zipcode = types.StringValue(value.String())
				} else {
					data.Address.Zipcode = types.StringNull()
				}
			}
		default:
			data.Type = types.StringValue("group")
		}
	} else {
		data.Type = types.StringValue("group")
	}

	return parentUuid
}

func (data *NetworkHierarchyNode) hasChanges(ctx context.Context, state *NetworkHierarchyNode) bool {
	hasChanges := false
	if !data.Name.Equal(state.Name) {
		hasChanges = true
	}
	if !data.Description.Equal(state.Description) {
		hasChanges = true
	}
	if !data.ParentGroup.Equal(state.ParentGroup) {
		hasChanges = true
	}
	if !data.Type.Equal(state.Type) {
		hasChanges = true
	}
	if !data.SiteId.Equal(state.SiteId) {
		hasChanges = true
	}
	if !data.IsSecondary.Equal(state.IsSecondary) {
		hasChanges = true
	}
	if data.Address != nil && state.Address != nil {
		if !data.Address.Street.Equal(state.Address.Street) {
			hasChanges = true
		}
		if !data.Address.City.Equal(state.Address.City) {
			hasChanges = true
		}
		if !data.Address.State.Equal(state.Address.State) {
			hasChanges = true
		}
		if !data.Address.Country.Equal(state.Address.Country) {
			hasChanges = true
		}
		if !data.Address.Zipcode.Equal(state.Address.Zipcode) {
			hasChanges = true
		}
	} else if (data.Address != nil && state.Address == nil) || (data.Address == nil && state.Address != nil) {
		hasChanges = true
	}
	if !data.Controllers.Equal(state.Controllers) {
		hasChanges = true
	}
	return hasChanges
}
