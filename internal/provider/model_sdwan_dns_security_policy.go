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
	"strconv"

	"github.com/CiscoDevNet/terraform-provider-sdwan/internal/provider/helpers"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types
type DNSSecurity struct {
	Id                       types.String            `tfsdk:"id"`
	Version                  types.Int64             `tfsdk:"version"`
	Name                     types.String            `tfsdk:"name"`
	Description              types.String            `tfsdk:"description"`
	FeatureProfileId         types.String            `tfsdk:"feature_profile_id"`
	LocalDomainBypassListId  types.String            `tfsdk:"local_domain_bypass_list_id"`
	MatchAllVpn              types.Bool              `tfsdk:"match_all_vpn"`
	UmbrellaDefault          types.Bool              `tfsdk:"umbrella_default"`
	DnsServerIp              types.String            `tfsdk:"dns_server_ip"`
	LocalDomainBypassEnabled types.Bool              `tfsdk:"local_domain_bypass_enabled"`
	DnsCrypt                 types.Bool              `tfsdk:"dns_crypt"`
	ChildOrgId               types.String            `tfsdk:"child_org_id"`
	TargetVpns               []DNSSecurityTargetVpns `tfsdk:"target_vpns"`
}

type DNSSecurityTargetVpns struct {
	Vpns                     types.Set    `tfsdk:"vpns"`
	UmbrellaDefault          types.Bool   `tfsdk:"umbrella_default"`
	DnsServerIp              types.String `tfsdk:"dns_server_ip"`
	LocalDomainBypassEnabled types.Bool   `tfsdk:"local_domain_bypass_enabled"`
	Uid                      types.String `tfsdk:"uid"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getModel
func (data DNSSecurity) getModel() string {
	return "dns_security"
}

// End of section. //template:end getModel

// Section below is generated&owned by "gen/generator.go". //template:begin getPath
func (data DNSSecurity) getPath() string {
	return fmt.Sprintf("/v1/feature-profile/sdwan/dns-security/%v/dns", url.QueryEscape(data.FeatureProfileId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody
func (data DNSSecurity) toBody(ctx context.Context) string {
	body := ""
	body, _ = sjson.Set(body, "name", data.Name.ValueString())
	body, _ = sjson.Set(body, "description", data.Description.ValueString())
	path := "data."
	if !data.LocalDomainBypassListId.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"localDomainBypassList.refId.optionType", "global")
			body, _ = sjson.Set(body, path+"localDomainBypassList.refId.value", data.LocalDomainBypassListId.ValueString())
		}
	}
	if !data.MatchAllVpn.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"matchAllVpn.optionType", "global")
			body, _ = sjson.Set(body, path+"matchAllVpn.value", data.MatchAllVpn.ValueBool())
		}
	}
	if !data.UmbrellaDefault.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"umbrellaDefault.optionType", "global")
			body, _ = sjson.Set(body, path+"umbrellaDefault.value", data.UmbrellaDefault.ValueBool())
		}
	}
	if !data.DnsServerIp.IsNull() {
		if true && data.MatchAllVpn.ValueBool() == true {
			body, _ = sjson.Set(body, path+"dnsServerIP.optionType", "global")
			body, _ = sjson.Set(body, path+"dnsServerIP.value", data.DnsServerIp.ValueString())
		}
	}
	if !data.LocalDomainBypassEnabled.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"localDomainBypassEnabled.optionType", "global")
			body, _ = sjson.Set(body, path+"localDomainBypassEnabled.value", data.LocalDomainBypassEnabled.ValueBool())
		}
	}
	if !data.DnsCrypt.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"dnsCrypt.optionType", "global")
			body, _ = sjson.Set(body, path+"dnsCrypt.value", data.DnsCrypt.ValueBool())
		}
	}
	if !data.ChildOrgId.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"childOrgId.optionType", "global")
			body, _ = sjson.Set(body, path+"childOrgId.value", data.ChildOrgId.ValueString())
		}
	}
	if true && data.MatchAllVpn.ValueBool() == false {

		for _, item := range data.TargetVpns {
			itemBody := ""
			if !item.Vpns.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "vpns.optionType", "global")
					var values []string
					item.Vpns.ElementsAs(ctx, &values, false)
					itemBody, _ = sjson.Set(itemBody, "vpns.value", values)
				}
			}
			if !item.UmbrellaDefault.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "umbrellaDefault.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "umbrellaDefault.value", item.UmbrellaDefault.ValueBool())
				}
			}
			if !item.DnsServerIp.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "dnsServerIP.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "dnsServerIP.value", item.DnsServerIp.ValueString())
				}
			}
			if !item.LocalDomainBypassEnabled.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "localDomainBypassEnabled.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "localDomainBypassEnabled.value", item.LocalDomainBypassEnabled.ValueBool())
				}
			}
			if !item.Uid.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "uid.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "uid.value", item.Uid.ValueString())
				}
			}
			body, _ = sjson.SetRaw(body, path+"targetVpns.-1", itemBody)
		}
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody
func (data *DNSSecurity) fromBody(ctx context.Context, res gjson.Result) {
	data.Name = types.StringValue(res.Get("payload.name").String())
	if value := res.Get("payload.description"); value.Exists() && value.String() != "" {
		data.Description = types.StringValue(value.String())
	} else {
		data.Description = types.StringNull()
	}
	path := "payload.data."
	data.LocalDomainBypassListId = types.StringNull()

	if t := res.Get(path + "localDomainBypassList.refId.optionType"); t.Exists() {
		va := res.Get(path + "localDomainBypassList.refId.value")
		if t.String() == "global" {
			data.LocalDomainBypassListId = types.StringValue(va.String())
		}
	}
	data.MatchAllVpn = types.BoolNull()

	if t := res.Get(path + "matchAllVpn.optionType"); t.Exists() {
		va := res.Get(path + "matchAllVpn.value")
		if t.String() == "global" {
			data.MatchAllVpn = types.BoolValue(va.Bool())
		}
	}
	data.UmbrellaDefault = types.BoolNull()

	if t := res.Get(path + "umbrellaDefault.optionType"); t.Exists() {
		va := res.Get(path + "umbrellaDefault.value")
		if t.String() == "global" {
			data.UmbrellaDefault = types.BoolValue(va.Bool())
		}
	}
	data.DnsServerIp = types.StringNull()

	if t := res.Get(path + "dnsServerIP.optionType"); t.Exists() {
		va := res.Get(path + "dnsServerIP.value")
		if t.String() == "global" {
			data.DnsServerIp = types.StringValue(va.String())
		}
	}
	data.LocalDomainBypassEnabled = types.BoolNull()

	if t := res.Get(path + "localDomainBypassEnabled.optionType"); t.Exists() {
		va := res.Get(path + "localDomainBypassEnabled.value")
		if t.String() == "global" {
			data.LocalDomainBypassEnabled = types.BoolValue(va.Bool())
		}
	}
	data.DnsCrypt = types.BoolNull()

	if t := res.Get(path + "dnsCrypt.optionType"); t.Exists() {
		va := res.Get(path + "dnsCrypt.value")
		if t.String() == "global" {
			data.DnsCrypt = types.BoolValue(va.Bool())
		}
	}
	data.ChildOrgId = types.StringNull()

	if t := res.Get(path + "childOrgId.optionType"); t.Exists() {
		va := res.Get(path + "childOrgId.value")
		if t.String() == "global" {
			data.ChildOrgId = types.StringValue(va.String())
		}
	}
	if value := res.Get(path + "targetVpns"); value.Exists() {
		data.TargetVpns = make([]DNSSecurityTargetVpns, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := DNSSecurityTargetVpns{}
			item.Vpns = types.SetNull(types.StringType)

			if t := v.Get("vpns.optionType"); t.Exists() {
				va := v.Get("vpns.value")
				if t.String() == "global" {
					item.Vpns = helpers.GetStringSet(va.Array())
				}
			}
			item.UmbrellaDefault = types.BoolNull()

			if t := v.Get("umbrellaDefault.optionType"); t.Exists() {
				va := v.Get("umbrellaDefault.value")
				if t.String() == "global" {
					item.UmbrellaDefault = types.BoolValue(va.Bool())
				}
			}
			item.DnsServerIp = types.StringNull()

			if t := v.Get("dnsServerIP.optionType"); t.Exists() {
				va := v.Get("dnsServerIP.value")
				if t.String() == "global" {
					item.DnsServerIp = types.StringValue(va.String())
				}
			}
			item.LocalDomainBypassEnabled = types.BoolNull()

			if t := v.Get("localDomainBypassEnabled.optionType"); t.Exists() {
				va := v.Get("localDomainBypassEnabled.value")
				if t.String() == "global" {
					item.LocalDomainBypassEnabled = types.BoolValue(va.Bool())
				}
			}
			item.Uid = types.StringNull()

			if t := v.Get("uid.optionType"); t.Exists() {
				va := v.Get("uid.value")
				if t.String() == "global" {
					item.Uid = types.StringValue(va.String())
				}
			}
			data.TargetVpns = append(data.TargetVpns, item)
			return true
		})
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin updateFromBody
func (data *DNSSecurity) updateFromBody(ctx context.Context, res gjson.Result) {
	data.Name = types.StringValue(res.Get("payload.name").String())
	if value := res.Get("payload.description"); value.Exists() && value.String() != "" {
		data.Description = types.StringValue(value.String())
	} else {
		data.Description = types.StringNull()
	}
	path := "payload.data."
	data.LocalDomainBypassListId = types.StringNull()

	if t := res.Get(path + "localDomainBypassList.refId.optionType"); t.Exists() {
		va := res.Get(path + "localDomainBypassList.refId.value")
		if t.String() == "global" {
			data.LocalDomainBypassListId = types.StringValue(va.String())
		}
	}
	data.MatchAllVpn = types.BoolNull()

	if t := res.Get(path + "matchAllVpn.optionType"); t.Exists() {
		va := res.Get(path + "matchAllVpn.value")
		if t.String() == "global" {
			data.MatchAllVpn = types.BoolValue(va.Bool())
		}
	}
	data.UmbrellaDefault = types.BoolNull()

	if t := res.Get(path + "umbrellaDefault.optionType"); t.Exists() {
		va := res.Get(path + "umbrellaDefault.value")
		if t.String() == "global" {
			data.UmbrellaDefault = types.BoolValue(va.Bool())
		}
	}
	data.DnsServerIp = types.StringNull()

	if t := res.Get(path + "dnsServerIP.optionType"); t.Exists() {
		va := res.Get(path + "dnsServerIP.value")
		if t.String() == "global" {
			data.DnsServerIp = types.StringValue(va.String())
		}
	}
	data.LocalDomainBypassEnabled = types.BoolNull()

	if t := res.Get(path + "localDomainBypassEnabled.optionType"); t.Exists() {
		va := res.Get(path + "localDomainBypassEnabled.value")
		if t.String() == "global" {
			data.LocalDomainBypassEnabled = types.BoolValue(va.Bool())
		}
	}
	data.DnsCrypt = types.BoolNull()

	if t := res.Get(path + "dnsCrypt.optionType"); t.Exists() {
		va := res.Get(path + "dnsCrypt.value")
		if t.String() == "global" {
			data.DnsCrypt = types.BoolValue(va.Bool())
		}
	}
	data.ChildOrgId = types.StringNull()

	if t := res.Get(path + "childOrgId.optionType"); t.Exists() {
		va := res.Get(path + "childOrgId.value")
		if t.String() == "global" {
			data.ChildOrgId = types.StringValue(va.String())
		}
	}
	for i := range data.TargetVpns {
		keys := [...]string{"vpns", "umbrellaDefault", "dnsServerIP", "localDomainBypassEnabled", "uid"}
		keyValues := [...]string{helpers.GetStringFromSet(data.TargetVpns[i].Vpns).ValueString(), strconv.FormatBool(data.TargetVpns[i].UmbrellaDefault.ValueBool()), data.TargetVpns[i].DnsServerIp.ValueString(), strconv.FormatBool(data.TargetVpns[i].LocalDomainBypassEnabled.ValueBool()), data.TargetVpns[i].Uid.ValueString()}
		keyValuesVariables := [...]string{"", "", "", "", ""}

		var r gjson.Result
		res.Get(path + "targetVpns").ForEach(
			func(_, v gjson.Result) bool {
				found := false
				for ik := range keys {
					tt := v.Get(keys[ik] + ".optionType")
					vv := v.Get(keys[ik] + ".value")
					if tt.Exists() && vv.Exists() {
						if (tt.String() == "variable" && vv.String() == keyValuesVariables[ik]) || (tt.String() == "global" && vv.String() == keyValues[ik]) {
							found = true
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
		data.TargetVpns[i].Vpns = types.SetNull(types.StringType)

		if t := r.Get("vpns.optionType"); t.Exists() {
			va := r.Get("vpns.value")
			if t.String() == "global" {
				data.TargetVpns[i].Vpns = helpers.GetStringSet(va.Array())
			}
		}
		data.TargetVpns[i].UmbrellaDefault = types.BoolNull()

		if t := r.Get("umbrellaDefault.optionType"); t.Exists() {
			va := r.Get("umbrellaDefault.value")
			if t.String() == "global" {
				data.TargetVpns[i].UmbrellaDefault = types.BoolValue(va.Bool())
			}
		}
		data.TargetVpns[i].DnsServerIp = types.StringNull()

		if t := r.Get("dnsServerIP.optionType"); t.Exists() {
			va := r.Get("dnsServerIP.value")
			if t.String() == "global" {
				data.TargetVpns[i].DnsServerIp = types.StringValue(va.String())
			}
		}
		data.TargetVpns[i].LocalDomainBypassEnabled = types.BoolNull()

		if t := r.Get("localDomainBypassEnabled.optionType"); t.Exists() {
			va := r.Get("localDomainBypassEnabled.value")
			if t.String() == "global" {
				data.TargetVpns[i].LocalDomainBypassEnabled = types.BoolValue(va.Bool())
			}
		}
		data.TargetVpns[i].Uid = types.StringNull()

		if t := r.Get("uid.optionType"); t.Exists() {
			va := r.Get("uid.value")
			if t.String() == "global" {
				data.TargetVpns[i].Uid = types.StringValue(va.String())
			}
		}
	}
}

// End of section. //template:end updateFromBody
