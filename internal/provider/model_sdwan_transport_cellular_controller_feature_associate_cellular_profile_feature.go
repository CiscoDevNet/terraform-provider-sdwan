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
type TransportCellularControllerFeatureAssociateCellularProfileFeature struct {
	Id                                   types.String `tfsdk:"id"`
	Version                              types.Int64  `tfsdk:"version"`
	FeatureProfileId                     types.String `tfsdk:"feature_profile_id"`
	TransportCellularControllerFeatureId types.String `tfsdk:"transport_cellular_controller_feature_id"`
	TransportCellularProfileFeatureId    types.String `tfsdk:"transport_cellular_profile_feature_id"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath
func (data TransportCellularControllerFeatureAssociateCellularProfileFeature) getPath() string {
	return fmt.Sprintf("/v1/feature-profile/sdwan/transport/%v/cellular-controller/%s/cellular-profile/", url.QueryEscape(data.FeatureProfileId.ValueString()), url.QueryEscape(data.TransportCellularControllerFeatureId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody
func (data TransportCellularControllerFeatureAssociateCellularProfileFeature) toBody(ctx context.Context) string {
	body := ""
	if !data.TransportCellularProfileFeatureId.IsNull() {
		body, _ = sjson.Set(body, "parcelId", data.TransportCellularProfileFeatureId.ValueString())
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody
func (data *TransportCellularControllerFeatureAssociateCellularProfileFeature) fromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("parcelId"); value.Exists() {
		data.TransportCellularProfileFeatureId = types.StringValue(value.String())
	} else {
		data.TransportCellularProfileFeatureId = types.StringNull()
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin hasChanges
func (data *TransportCellularControllerFeatureAssociateCellularProfileFeature) hasChanges(ctx context.Context, state *TransportCellularControllerFeatureAssociateCellularProfileFeature) bool {
	hasChanges := false
	if !data.FeatureProfileId.Equal(state.FeatureProfileId) {
		hasChanges = true
	}
	if !data.TransportCellularControllerFeatureId.Equal(state.TransportCellularControllerFeatureId) {
		hasChanges = true
	}
	if !data.TransportCellularProfileFeatureId.Equal(state.TransportCellularProfileFeatureId) {
		hasChanges = true
	}
	return hasChanges
}

// End of section. //template:end hasChanges

// Section below is generated&owned by "gen/generator.go". //template:begin updateVersions

// End of section. //template:end updateVersions

// Section below is generated&owned by "gen/generator.go". //template:begin processImport
func (data *TransportCellularControllerFeatureAssociateCellularProfileFeature) processImport(ctx context.Context) {
	data.Version = types.Int64Value(0)
}

// End of section. //template:end processImport

// Section below is generated&owned by "gen/generator.go". //template:begin applyFilters
// End of section. //template:end applyFilters
