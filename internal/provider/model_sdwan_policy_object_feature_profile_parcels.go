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
type PolicyObjectFeatureProfileParcels struct {
	Id               types.String                               `tfsdk:"id"`
	FeatureProfileId types.String                               `tfsdk:"feature_profile_id"`
	Parcels          []PolicyObjectFeatureProfileParcelsParcels `tfsdk:"parcels"`
	CreatedBy        types.String                               `tfsdk:"created_by"`
	ParcelType       types.String                               `tfsdk:"parcel_type"`
}

type PolicyObjectFeatureProfileParcelsParcels struct {
	ParcelId          types.String `tfsdk:"parcel_id"`
	ParcelType        types.String `tfsdk:"parcel_type"`
	CreatedBy         types.String `tfsdk:"created_by"`
	LastUpdatedBy     types.String `tfsdk:"last_updated_by"`
	ParcelName        types.String `tfsdk:"parcel_name"`
	ParcelDescription types.String `tfsdk:"parcel_description"`
	ReferenceCount    types.Int64  `tfsdk:"reference_count"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath
func (data PolicyObjectFeatureProfileParcels) getPath() string {
	return fmt.Sprintf("/v1/feature-profile/sdwan/policy-object/%v", url.QueryEscape(data.FeatureProfileId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody
func (data PolicyObjectFeatureProfileParcels) toBody(ctx context.Context) string {
	body := ""
	if true {
		body, _ = sjson.Set(body, "associatedProfileParcels", []interface{}{})
		for _, item := range data.Parcels {
			itemBody := ""
			if !item.ParcelId.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "parcelId", item.ParcelId.ValueString())
			}
			if !item.ParcelType.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "parcelType", item.ParcelType.ValueString())
			}
			if !item.CreatedBy.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "createdBy", item.CreatedBy.ValueString())
			}
			if !item.LastUpdatedBy.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "lastUpdatedBy", item.LastUpdatedBy.ValueString())
			}
			if !item.ParcelName.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "payload.name", item.ParcelName.ValueString())
			}
			if !item.ParcelDescription.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "payload.description", item.ParcelDescription.ValueString())
			}
			if !item.ReferenceCount.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "referenceCount", item.ReferenceCount.ValueInt64())
			}
			body, _ = sjson.SetRaw(body, "associatedProfileParcels.-1", itemBody)
		}
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody
func (data *PolicyObjectFeatureProfileParcels) fromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("associatedProfileParcels"); value.Exists() && len(value.Array()) > 0 {
		data.Parcels = make([]PolicyObjectFeatureProfileParcelsParcels, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := PolicyObjectFeatureProfileParcelsParcels{}
			if cValue := v.Get("parcelId"); cValue.Exists() {
				item.ParcelId = types.StringValue(cValue.String())
			} else {
				item.ParcelId = types.StringNull()
			}
			if cValue := v.Get("parcelType"); cValue.Exists() {
				item.ParcelType = types.StringValue(cValue.String())
			} else {
				item.ParcelType = types.StringNull()
			}
			if cValue := v.Get("createdBy"); cValue.Exists() {
				item.CreatedBy = types.StringValue(cValue.String())
			} else {
				item.CreatedBy = types.StringNull()
			}
			if cValue := v.Get("lastUpdatedBy"); cValue.Exists() {
				item.LastUpdatedBy = types.StringValue(cValue.String())
			} else {
				item.LastUpdatedBy = types.StringNull()
			}
			if cValue := v.Get("payload.name"); cValue.Exists() {
				item.ParcelName = types.StringValue(cValue.String())
			} else {
				item.ParcelName = types.StringNull()
			}
			if cValue := v.Get("payload.description"); cValue.Exists() {
				item.ParcelDescription = types.StringValue(cValue.String())
			} else {
				item.ParcelDescription = types.StringNull()
			}
			if cValue := v.Get("referenceCount"); cValue.Exists() {
				item.ReferenceCount = types.Int64Value(cValue.Int())
			} else {
				item.ReferenceCount = types.Int64Null()
			}
			data.Parcels = append(data.Parcels, item)
			return true
		})
	} else {
		if len(data.Parcels) > 0 {
			data.Parcels = []PolicyObjectFeatureProfileParcelsParcels{}
		}
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin hasChanges
func (data *PolicyObjectFeatureProfileParcels) hasChanges(ctx context.Context, state *PolicyObjectFeatureProfileParcels) bool {
	hasChanges := false
	if !data.FeatureProfileId.Equal(state.FeatureProfileId) {
		hasChanges = true
	}
	if len(data.Parcels) != len(state.Parcels) {
		hasChanges = true
	} else {
		for i := range data.Parcels {
			if !data.Parcels[i].ParcelId.Equal(state.Parcels[i].ParcelId) {
				hasChanges = true
			}
			if !data.Parcels[i].ParcelType.Equal(state.Parcels[i].ParcelType) {
				hasChanges = true
			}
			if !data.Parcels[i].CreatedBy.Equal(state.Parcels[i].CreatedBy) {
				hasChanges = true
			}
			if !data.Parcels[i].LastUpdatedBy.Equal(state.Parcels[i].LastUpdatedBy) {
				hasChanges = true
			}
			if !data.Parcels[i].ParcelName.Equal(state.Parcels[i].ParcelName) {
				hasChanges = true
			}
			if !data.Parcels[i].ParcelDescription.Equal(state.Parcels[i].ParcelDescription) {
				hasChanges = true
			}
			if !data.Parcels[i].ReferenceCount.Equal(state.Parcels[i].ReferenceCount) {
				hasChanges = true
			}
		}
	}
	return hasChanges
}

// End of section. //template:end hasChanges

// Section below is generated&owned by "gen/generator.go". //template:begin updateVersions

// End of section. //template:end updateVersions

// Section below is generated&owned by "gen/generator.go". //template:begin processImport
func (data *PolicyObjectFeatureProfileParcels) processImport(ctx context.Context) {
}

// End of section. //template:end processImport

// Section below is generated&owned by "gen/generator.go". //template:begin applyFilters
func (data *PolicyObjectFeatureProfileParcels) applyFilters(ctx context.Context) {
	filtered := make([]PolicyObjectFeatureProfileParcelsParcels, 0, len(data.Parcels))
	for _, item := range data.Parcels {
		match := true
		if !data.CreatedBy.IsNull() && !data.CreatedBy.IsUnknown() {
			if item.CreatedBy.ValueString() != data.CreatedBy.ValueString() {
				match = false
			}
		}
		if !data.ParcelType.IsNull() && !data.ParcelType.IsUnknown() {
			if item.ParcelType.ValueString() != data.ParcelType.ValueString() {
				match = false
			}
		}
		if match {
			filtered = append(filtered, item)
		}
	}
	data.Parcels = filtered
}

// End of section. //template:end applyFilters
