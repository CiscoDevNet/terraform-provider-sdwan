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
	"reflect"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func parcelsForFilterTest() []PolicyObjectFeatureProfileParcelsParcels {
	parcel := func(name, parcelType, createdBy string) PolicyObjectFeatureProfileParcelsParcels {
		return PolicyObjectFeatureProfileParcelsParcels{
			ParcelName: types.StringValue(name),
			ParcelType: types.StringValue(parcelType),
			CreatedBy:  types.StringValue(createdBy),
		}
	}
	return []PolicyObjectFeatureProfileParcelsParcels{
		parcel("dpl_dc", "data-prefix", "admin"),
		parcel("dpl_branch", "data-prefix", "admin"),
		parcel("apps_branch", "app-list", "admin"),
		parcel("box_net_apps", "app-list", "system"),
	}
}

func filteredNames(data PolicyObjectFeatureProfileParcels) []string {
	names := []string{}
	for _, p := range data.Parcels {
		names = append(names, p.ParcelName.ValueString())
	}
	return names
}

func stringSet(values ...string) types.Set {
	elems := make([]attr.Value, 0, len(values))
	for _, v := range values {
		elems = append(elems, types.StringValue(v))
	}
	return types.SetValueMust(types.StringType, elems)
}

func TestPolicyObjectFeatureProfileParcelsApplyFilters(t *testing.T) {
	tests := []struct {
		name       string
		createdBy  types.String
		parcelType types.String
		names      types.Set
		want       []string
	}{
		{"no filter", types.StringNull(), types.StringNull(), types.SetNull(types.StringType), []string{"dpl_dc", "dpl_branch", "apps_branch", "box_net_apps"}},
		{"names", types.StringNull(), types.StringNull(), stringSet("dpl_branch", "box_net_apps", "missing"), []string{"dpl_branch", "box_net_apps"}},
		{"empty names", types.StringNull(), types.StringNull(), stringSet(), []string{}},
		{"names and type", types.StringNull(), types.StringValue("app-list"), stringSet("dpl_branch", "apps_branch"), []string{"apps_branch"}},
		{"type and creator", types.StringValue("system"), types.StringValue("app-list"), types.SetNull(types.StringType), []string{"box_net_apps"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data := PolicyObjectFeatureProfileParcels{
				Parcels:     parcelsForFilterTest(),
				CreatedBy:   tt.createdBy,
				ParcelType:  tt.parcelType,
				ParcelNames: tt.names,
			}
			data.applyFilters(context.Background())
			if got := filteredNames(data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("applyFilters() kept %v, want %v", got, tt.want)
			}
		})
	}
}
