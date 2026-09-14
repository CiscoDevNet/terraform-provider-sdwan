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

	"github.com/hashicorp/go-version"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// minVersionNetworkHierarchyNodeGeo is the minimum SD-WAN Manager version that
// supports the location/latitude/longitude geolocation model on site nodes.
var minVersionNetworkHierarchyNodeGeo = version.Must(version.NewVersion("20.18.1"))

// networkHierarchyNodeAddressAttrTypes describes `address`'s object shape for
// types.Object construction/decoding. `address` is types.Object rather than a
// plain *NetworkHierarchyNodeAddress struct pointer specifically because a
// plain struct pointer cannot represent Unknown (only nil=null or
// populated=known) - confirmed live, referencing this attribute from a
// for_each-derived expression (e.g. `try(each.value.address, null)`, standard
// Terraform module authoring) crashes with "Received unknown value, however
// the target type cannot handle unknown values" even when the real resolved
// value would be null. types.Object has a proper null/unknown/known state and
// does not have this problem.
var networkHierarchyNodeAddressAttrTypes = map[string]attr.Type{
	"street":  types.StringType,
	"city":    types.StringType,
	"state":   types.StringType,
	"country": types.StringType,
	"zipcode": types.StringType,
}

type NetworkHierarchyNode struct {
	Id          types.String  `tfsdk:"id"`
	Name        types.String  `tfsdk:"name"`
	Description types.String  `tfsdk:"description"`
	ParentGroup types.String  `tfsdk:"parent_group"`
	Type        types.String  `tfsdk:"type"`
	SiteId      types.Int64   `tfsdk:"site_id"`
	IsSecondary types.Bool    `tfsdk:"is_secondary"`
	Address     types.Object  `tfsdk:"address"`
	Location    types.String  `tfsdk:"location"`
	Latitude    types.Float64 `tfsdk:"latitude"`
	Longitude   types.Float64 `tfsdk:"longitude"`
}

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

func (data NetworkHierarchyNode) resolveParentGroupToId(hierarchyRes gjson.Result) (string, string) {
	parentGroupName := data.ParentGroup.ValueString()
	childType := data.Type.ValueString()

	var foundId string
	var matchCount int

	hierarchyRes.ForEach(func(key, value gjson.Result) bool {
		nodeName := value.Get("name").String()
		nodeLabel := value.Get("data.label").String()

		// A SITE can never be a parent of anything. A REGION can parent a group
		// or a site, but not another region - a region cannot itself be nested
		// under a region, only under a group or Global (matches the Manager GUI).
		isValidParent := nodeLabel != "SITE"
		if childType == "region" {
			isValidParent = isValidParent && nodeLabel != "REGION"
		}

		if nodeName == parentGroupName && isValidParent {
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

func (data NetworkHierarchyNode) resolveParentIdToGroup(hierarchyRes gjson.Result, parentId string) string {
	var parentName string

	hierarchyRes.ForEach(func(key, value gjson.Result) bool {
		nodeId := value.Get("id").String()
		if nodeId == parentId {
			parentName = value.Get("name").String()
			return false
		}
		return true
	})

	return parentName
}

func (data NetworkHierarchyNode) toBody(ctx context.Context, parentId string, ver *version.Version) string {
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
		isGeoVersion := ver != nil && ver.GreaterThanOrEqual(minVersionNetworkHierarchyNodeGeo)
		// On SD-WAN Manager < 20.18.1 sites use a structured `address` block.
		// On 20.18.1+ that block is invalid; sites use a place-name `location`
		// plus a `gpsLocation` object instead.
		if !isGeoVersion && !data.Address.IsNull() {
			var addr NetworkHierarchyNodeAddress
			data.Address.As(ctx, &addr, basetypes.ObjectAsOptions{})
			if !addr.Street.IsNull() {
				body, _ = sjson.Set(body, "data.address.street", addr.Street.ValueString())
			}
			if !addr.City.IsNull() {
				body, _ = sjson.Set(body, "data.address.city", addr.City.ValueString())
			}
			if !addr.State.IsNull() {
				body, _ = sjson.Set(body, "data.address.state", addr.State.ValueString())
			}
			if !addr.Country.IsNull() {
				body, _ = sjson.Set(body, "data.address.country", addr.Country.ValueString())
			}
			if !addr.Zipcode.IsNull() {
				body, _ = sjson.Set(body, "data.address.zipcode", addr.Zipcode.ValueString())
			}
		}
		if isGeoVersion {
			// On 20.18.1+ `location` is always part of the payload. When the
			// user does not set a real place name, send "Undisclosed" (the
			// Manager's own default). `location` is Optional+Computed on the
			// resource so the read-back value does not drift against a null
			// config. `latitude`/`longitude` remain optional and are only sent
			// when the user provides them (a real location is required for GPS).
			loc := "Undisclosed"
			if !data.Location.IsNull() && data.Location.ValueString() != "" {
				loc = data.Location.ValueString()
			}
			body, _ = sjson.Set(body, "data.location", loc)
			if !data.Latitude.IsNull() {
				body, _ = sjson.Set(body, "data.gpsLocation.latitude", data.Latitude.ValueFloat64())
			}
			if !data.Longitude.IsNull() {
				body, _ = sjson.Set(body, "data.gpsLocation.longitude", data.Longitude.ValueFloat64())
			}
		}
	}

	return body
}

func (data NetworkHierarchyNode) toUpdateBody(ctx context.Context, currentNodeRes gjson.Result, parentId string, ver *version.Version) string {
	body := data.toBody(ctx, parentId, ver)

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

	// `location` is Optional+Computed, so it must always be resolved to a known
	// value. `address` is types.Object, which (unlike a plain struct pointer)
	// must always be resolved to a known state too - an object left at its Go
	// zero value is not a valid null/unknown/known types.Object. Both only
	// exist on site nodes; for region/group nodes they stay null. The SITE
	// case below overrides them with the Manager-assigned values.
	data.Location = types.StringNull()
	data.Address = types.ObjectNull(networkHierarchyNodeAddressAttrTypes)

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
				addr := NetworkHierarchyNodeAddress{
					Street:  types.StringNull(),
					City:    types.StringNull(),
					State:   types.StringNull(),
					Country: types.StringNull(),
					Zipcode: types.StringNull(),
				}
				if value := res.Get("data.address.street"); value.Exists() {
					addr.Street = types.StringValue(value.String())
				}
				if value := res.Get("data.address.city"); value.Exists() {
					addr.City = types.StringValue(value.String())
				}
				if value := res.Get("data.address.state"); value.Exists() {
					addr.State = types.StringValue(value.String())
				}
				if value := res.Get("data.address.country"); value.Exists() {
					addr.Country = types.StringValue(value.String())
				}
				if value := res.Get("data.address.zipcode"); value.Exists() {
					addr.Zipcode = types.StringValue(value.String())
				}
				objVal, _ := types.ObjectValueFrom(ctx, networkHierarchyNodeAddressAttrTypes, addr)
				data.Address = objVal
			}
			if value := res.Get("data.location"); value.Exists() {
				data.Location = types.StringValue(value.String())
			} else {
				data.Location = types.StringNull()
			}
			if value := res.Get("data.gpsLocation.latitude"); value.Exists() {
				data.Latitude = types.Float64Value(value.Float())
			} else {
				data.Latitude = types.Float64Null()
			}
			if value := res.Get("data.gpsLocation.longitude"); value.Exists() {
				data.Longitude = types.Float64Value(value.Float())
			} else {
				data.Longitude = types.Float64Null()
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
	// types.Object.Equal handles null-vs-null, null-vs-populated, and
	// populated-vs-populated-with-different-fields in one call - the explicit
	// per-field/nil-mismatch handling the *NetworkHierarchyNodeAddress pointer
	// version needed is no longer necessary.
	if !data.Address.Equal(state.Address) {
		hasChanges = true
	}
	if !data.Location.Equal(state.Location) {
		hasChanges = true
	}
	if !data.Latitude.Equal(state.Latitude) {
		hasChanges = true
	}
	if !data.Longitude.Equal(state.Longitude) {
		hasChanges = true
	}
	return hasChanges
}
