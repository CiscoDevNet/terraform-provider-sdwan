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
type SystemCACertificate struct {
	Id               types.String                      `tfsdk:"id"`
	Version          types.Int64                       `tfsdk:"version"`
	Name             types.String                      `tfsdk:"name"`
	Description      types.String                      `tfsdk:"description"`
	FeatureProfileId types.String                      `tfsdk:"feature_profile_id"`
	Certificates     []SystemCACertificateCertificates `tfsdk:"certificates"`
}

type SystemCACertificateCertificates struct {
	TrustPointName  types.String `tfsdk:"trust_point_name"`
	CaCertificateId types.String `tfsdk:"ca_certificate_id"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getModel
func (data SystemCACertificate) getModel() string {
	return "system_ca_certificate"
}

// End of section. //template:end getModel

// Section below is generated&owned by "gen/generator.go". //template:begin getPath
func (data SystemCACertificate) getPath() string {
	return fmt.Sprintf("/v1/feature-profile/sdwan/system/%v/ca-cert", url.QueryEscape(data.FeatureProfileId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody
func (data SystemCACertificate) toBody(ctx context.Context) string {
	body := ""
	body, _ = sjson.Set(body, "name", data.Name.ValueString())
	body, _ = sjson.Set(body, "description", data.Description.ValueString())
	path := "data."
	if true {
		body, _ = sjson.Set(body, path+"certificates", []interface{}{})
		for _, item := range data.Certificates {
			itemBody := ""
			if !item.TrustPointName.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "trustPointName.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "trustPointName.value", item.TrustPointName.ValueString())
				}
			}
			if !item.CaCertificateId.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "certificateUUID.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "certificateUUID.value", item.CaCertificateId.ValueString())
				}
			}
			body, _ = sjson.SetRaw(body, path+"certificates.-1", itemBody)
		}
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody
func (data *SystemCACertificate) fromBody(ctx context.Context, res gjson.Result) {
	data.Name = types.StringValue(res.Get("payload.name").String())
	if value := res.Get("payload.description"); value.Exists() && value.String() != "" {
		data.Description = types.StringValue(value.String())
	} else {
		data.Description = types.StringNull()
	}
	path := "payload.data."
	if value := res.Get(path + "certificates"); value.Exists() && len(value.Array()) > 0 {
		data.Certificates = make([]SystemCACertificateCertificates, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := SystemCACertificateCertificates{}
			item.TrustPointName = types.StringNull()

			if t := v.Get("trustPointName.optionType"); t.Exists() {
				va := v.Get("trustPointName.value")
				if t.String() == "global" {
					item.TrustPointName = types.StringValue(va.String())
				}
			}
			item.CaCertificateId = types.StringNull()

			if t := v.Get("certificateUUID.optionType"); t.Exists() {
				va := v.Get("certificateUUID.value")
				if t.String() == "global" {
					item.CaCertificateId = types.StringValue(va.String())
				}
			}
			data.Certificates = append(data.Certificates, item)
			return true
		})
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin updateFromBody
func (data *SystemCACertificate) updateFromBody(ctx context.Context, res gjson.Result) {
	data.Name = types.StringValue(res.Get("payload.name").String())
	if value := res.Get("payload.description"); value.Exists() && value.String() != "" {
		data.Description = types.StringValue(value.String())
	} else {
		data.Description = types.StringNull()
	}
	path := "payload.data."
	for i := range data.Certificates {
		keys := [...]string{"trustPointName", "certificateUUID"}
		keyValues := [...]string{data.Certificates[i].TrustPointName.ValueString(), data.Certificates[i].CaCertificateId.ValueString()}
		keyValuesVariables := [...]string{"", ""}

		var r gjson.Result
		res.Get(path + "certificates").ForEach(
			func(_, v gjson.Result) bool {
				found := false
				for ik := range keys {
					tt := v.Get(keys[ik] + ".optionType")
					vv := v.Get(keys[ik] + ".value")
					if tt.Exists() && vv.Exists() {
						if (tt.String() == "variable" && vv.String() == keyValuesVariables[ik]) || (tt.String() == "global" && vv.String() == keyValues[ik]) {
							found = true
							continue
						} else if tt.String() == "default" {
							continue
						}
						found = false
						break
					}
					continue
				}
				if found {
					r = v
					return false
				}
				return true
			},
		)
		data.Certificates[i].TrustPointName = types.StringNull()

		if t := r.Get("trustPointName.optionType"); t.Exists() {
			va := r.Get("trustPointName.value")
			if t.String() == "global" {
				data.Certificates[i].TrustPointName = types.StringValue(va.String())
			}
		}
		data.Certificates[i].CaCertificateId = types.StringNull()

		if t := r.Get("certificateUUID.optionType"); t.Exists() {
			va := r.Get("certificateUUID.value")
			if t.String() == "global" {
				data.Certificates[i].CaCertificateId = types.StringValue(va.String())
			}
		}
	}
}

// End of section. //template:end updateFromBody
