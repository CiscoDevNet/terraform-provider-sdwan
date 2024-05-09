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

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types
type IPSSignatureListPolicyObject struct {
	Id      types.String                          `tfsdk:"id"`
	Version types.Int64                           `tfsdk:"version"`
	Name    types.String                          `tfsdk:"name"`
	Entries []IPSSignatureListPolicyObjectEntries `tfsdk:"entries"`
}

type IPSSignatureListPolicyObjectEntries struct {
	GeneratorId types.Int64 `tfsdk:"generator_id"`
	SignatureId types.Int64 `tfsdk:"signature_id"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath
func (data IPSSignatureListPolicyObject) getPath() string {
	return "/template/policy/list/ipssignature/"
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody
func (data IPSSignatureListPolicyObject) toBody(ctx context.Context) string {
	body := ""
	if true {
		body, _ = sjson.Set(body, "type", "ipssignature")
	}
	if !data.Name.IsNull() {
		body, _ = sjson.Set(body, "name", data.Name.ValueString())
	}
	if true {
		body, _ = sjson.Set(body, "entries", []interface{}{})
		for _, item := range data.Entries {
			itemBody := ""
			if !item.GeneratorId.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "generatorId", fmt.Sprint(item.GeneratorId.ValueInt64()))
			}
			if !item.SignatureId.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "signatureId", fmt.Sprint(item.SignatureId.ValueInt64()))
			}
			body, _ = sjson.SetRaw(body, "entries.-1", itemBody)
		}
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody
func (data *IPSSignatureListPolicyObject) fromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("name"); value.Exists() {
		data.Name = types.StringValue(value.String())
	} else {
		data.Name = types.StringNull()
	}
	if value := res.Get("entries"); value.Exists() && len(value.Array()) > 0 {
		data.Entries = make([]IPSSignatureListPolicyObjectEntries, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := IPSSignatureListPolicyObjectEntries{}
			if cValue := v.Get("generatorId"); cValue.Exists() {
				item.GeneratorId = types.Int64Value(cValue.Int())
			} else {
				item.GeneratorId = types.Int64Null()
			}
			if cValue := v.Get("signatureId"); cValue.Exists() {
				item.SignatureId = types.Int64Value(cValue.Int())
			} else {
				item.SignatureId = types.Int64Null()
			}
			data.Entries = append(data.Entries, item)
			return true
		})
	} else {
		if len(data.Entries) > 0 {
			data.Entries = []IPSSignatureListPolicyObjectEntries{}
		}
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin hasChanges
func (data *IPSSignatureListPolicyObject) hasChanges(ctx context.Context, state *IPSSignatureListPolicyObject) bool {
	hasChanges := false
	if !data.Name.Equal(state.Name) {
		hasChanges = true
	}
	if len(data.Entries) != len(state.Entries) {
		hasChanges = true
	} else {
		for i := range data.Entries {
			if !data.Entries[i].GeneratorId.Equal(state.Entries[i].GeneratorId) {
				hasChanges = true
			}
			if !data.Entries[i].SignatureId.Equal(state.Entries[i].SignatureId) {
				hasChanges = true
			}
		}
	}
	return hasChanges
}

// End of section. //template:end hasChanges

// Section below is generated&owned by "gen/generator.go". //template:begin updateVersions

// End of section. //template:end updateVersions
