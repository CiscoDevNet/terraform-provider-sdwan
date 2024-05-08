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

	"github.com/CiscoDevNet/terraform-provider-sdwan/internal/provider/helpers"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types
type DNSSecurityPolicyDefinition struct {
	Id                                        types.String                            `tfsdk:"id"`
	Version                                   types.Int64                             `tfsdk:"version"`
	Name                                      types.String                            `tfsdk:"name"`
	Description                               types.String                            `tfsdk:"description"`
	DomainListId                              types.String                            `tfsdk:"domain_list_id"`
	DomainListVersion                         types.Int64                             `tfsdk:"domain_list_version"`
	LocalDomainBypassEnabled                  types.Bool                              `tfsdk:"local_domain_bypass_enabled"`
	MatchAllVpn                               types.Bool                              `tfsdk:"match_all_vpn"`
	TargetVpns                                []DNSSecurityPolicyDefinitionTargetVpns `tfsdk:"target_vpns"`
	Dnscrypt                                  types.Bool                              `tfsdk:"dnscrypt"`
	UmbrellaDnsDefault                        types.Bool                              `tfsdk:"umbrella_dns_default"`
	CustomDnsServerIp                         types.String                            `tfsdk:"custom_dns_server_ip"`
	CiscoSigCredentialsFeatureTemplateId      types.String                            `tfsdk:"cisco_sig_credentials_feature_template_id"`
	CiscoSigCredentialsFeatureTemplateVersion types.Int64                             `tfsdk:"cisco_sig_credentials_feature_template_version"`
}

type DNSSecurityPolicyDefinitionTargetVpns struct {
	VpnIds                   types.Set    `tfsdk:"vpn_ids"`
	UmbrellaDnsDefault       types.Bool   `tfsdk:"umbrella_dns_default"`
	CustomDnsServerIp        types.String `tfsdk:"custom_dns_server_ip"`
	LocalDomainBypassEnabled types.Bool   `tfsdk:"local_domain_bypass_enabled"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath
func (data DNSSecurityPolicyDefinition) getPath() string {
	return "/template/policy/definition/dnssecurity/"
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody
func (data DNSSecurityPolicyDefinition) toBody(ctx context.Context) string {
	body := ""
	if true {
		body, _ = sjson.Set(body, "type", "DNSSecurity")
	}
	if !data.Name.IsNull() {
		body, _ = sjson.Set(body, "name", data.Name.ValueString())
	}
	if !data.Description.IsNull() {
		body, _ = sjson.Set(body, "description", data.Description.ValueString())
	}
	if !data.DomainListId.IsNull() {
		body, _ = sjson.Set(body, "definition.localDomainBypassList.ref", data.DomainListId.ValueString())
	}
	if !data.LocalDomainBypassEnabled.IsNull() {
		if false && data.LocalDomainBypassEnabled.ValueBool() {
			body, _ = sjson.Set(body, "definition.localDomainBypassEnabled", "")
		} else {
			body, _ = sjson.Set(body, "definition.localDomainBypassEnabled", data.LocalDomainBypassEnabled.ValueBool())
		}
	}
	if !data.MatchAllVpn.IsNull() {
		if false && data.MatchAllVpn.ValueBool() {
			body, _ = sjson.Set(body, "definition.matchAllVpn", "")
		} else {
			body, _ = sjson.Set(body, "definition.matchAllVpn", data.MatchAllVpn.ValueBool())
		}
	}
	if true {
		body, _ = sjson.Set(body, "definition.targetVpns", []interface{}{})
		for _, item := range data.TargetVpns {
			itemBody := ""
			if !item.VpnIds.IsNull() {
				var values []string
				item.VpnIds.ElementsAs(ctx, &values, false)
				itemBody, _ = sjson.Set(itemBody, "vpns", values)
			}
			if !item.UmbrellaDnsDefault.IsNull() {
				if false && item.UmbrellaDnsDefault.ValueBool() {
					itemBody, _ = sjson.Set(itemBody, "umbrellaDefault", "")
				} else {
					itemBody, _ = sjson.Set(itemBody, "umbrellaDefault", item.UmbrellaDnsDefault.ValueBool())
				}
			}
			if !item.CustomDnsServerIp.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "dnsServerIP", item.CustomDnsServerIp.ValueString())
			}
			if !item.LocalDomainBypassEnabled.IsNull() {
				if false && item.LocalDomainBypassEnabled.ValueBool() {
					itemBody, _ = sjson.Set(itemBody, "localDomainBypassEnabled", "")
				} else {
					itemBody, _ = sjson.Set(itemBody, "localDomainBypassEnabled", item.LocalDomainBypassEnabled.ValueBool())
				}
			}
			body, _ = sjson.SetRaw(body, "definition.targetVpns.-1", itemBody)
		}
	}
	if !data.Dnscrypt.IsNull() {
		if false && data.Dnscrypt.ValueBool() {
			body, _ = sjson.Set(body, "definition.dnsCrypt", "")
		} else {
			body, _ = sjson.Set(body, "definition.dnsCrypt", data.Dnscrypt.ValueBool())
		}
	}
	if !data.UmbrellaDnsDefault.IsNull() {
		if false && data.UmbrellaDnsDefault.ValueBool() {
			body, _ = sjson.Set(body, "definition.umbrellaDefault", "")
		} else {
			body, _ = sjson.Set(body, "definition.umbrellaDefault", data.UmbrellaDnsDefault.ValueBool())
		}
	}
	if !data.CustomDnsServerIp.IsNull() {
		body, _ = sjson.Set(body, "definition.dnsServerIP", data.CustomDnsServerIp.ValueString())
	}
	if !data.CiscoSigCredentialsFeatureTemplateId.IsNull() {
		body, _ = sjson.Set(body, "definition.umbrellaData.ref", data.CiscoSigCredentialsFeatureTemplateId.ValueString())
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody
func (data *DNSSecurityPolicyDefinition) fromBody(ctx context.Context, res gjson.Result) {
	state := *data
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
	if value := res.Get("definition.localDomainBypassList.ref"); value.Exists() {
		data.DomainListId = types.StringValue(value.String())
	} else {
		data.DomainListId = types.StringNull()
	}
	if value := res.Get("definition.localDomainBypassEnabled"); value.Exists() {
		if false && value.String() == "" {
			data.LocalDomainBypassEnabled = types.BoolValue(true)
		} else {
			data.LocalDomainBypassEnabled = types.BoolValue(value.Bool())
		}
	} else {
		data.LocalDomainBypassEnabled = types.BoolNull()
	}
	if value := res.Get("definition.matchAllVpn"); value.Exists() {
		if false && value.String() == "" {
			data.MatchAllVpn = types.BoolValue(true)
		} else {
			data.MatchAllVpn = types.BoolValue(value.Bool())
		}
	} else {
		data.MatchAllVpn = types.BoolNull()
	}
	if value := res.Get("definition.targetVpns"); value.Exists() && len(value.Array()) > 0 {
		data.TargetVpns = make([]DNSSecurityPolicyDefinitionTargetVpns, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := DNSSecurityPolicyDefinitionTargetVpns{}
			if cValue := v.Get("vpns"); cValue.Exists() {
				item.VpnIds = helpers.GetStringSet(cValue.Array())
			} else {
				item.VpnIds = types.SetNull(types.StringType)
			}
			if cValue := v.Get("umbrellaDefault"); cValue.Exists() {
				if false && cValue.String() == "" {
					item.UmbrellaDnsDefault = types.BoolValue(true)
				} else {
					item.UmbrellaDnsDefault = types.BoolValue(cValue.Bool())
				}
			} else {
				item.UmbrellaDnsDefault = types.BoolNull()
			}
			if cValue := v.Get("dnsServerIP"); cValue.Exists() {
				item.CustomDnsServerIp = types.StringValue(cValue.String())
			} else {
				item.CustomDnsServerIp = types.StringNull()
			}
			if cValue := v.Get("localDomainBypassEnabled"); cValue.Exists() {
				if false && cValue.String() == "" {
					item.LocalDomainBypassEnabled = types.BoolValue(true)
				} else {
					item.LocalDomainBypassEnabled = types.BoolValue(cValue.Bool())
				}
			} else {
				item.LocalDomainBypassEnabled = types.BoolNull()
			}
			data.TargetVpns = append(data.TargetVpns, item)
			return true
		})
	} else {
		if len(data.TargetVpns) > 0 {
			data.TargetVpns = []DNSSecurityPolicyDefinitionTargetVpns{}
		}
	}
	if value := res.Get("definition.dnsCrypt"); value.Exists() {
		if false && value.String() == "" {
			data.Dnscrypt = types.BoolValue(true)
		} else {
			data.Dnscrypt = types.BoolValue(value.Bool())
		}
	} else {
		data.Dnscrypt = types.BoolNull()
	}
	if value := res.Get("definition.umbrellaDefault"); value.Exists() {
		if false && value.String() == "" {
			data.UmbrellaDnsDefault = types.BoolValue(true)
		} else {
			data.UmbrellaDnsDefault = types.BoolValue(value.Bool())
		}
	} else {
		data.UmbrellaDnsDefault = types.BoolNull()
	}
	if value := res.Get("definition.dnsServerIP"); value.Exists() {
		data.CustomDnsServerIp = types.StringValue(value.String())
	} else {
		data.CustomDnsServerIp = types.StringNull()
	}
	if value := res.Get("definition.umbrellaData.ref"); value.Exists() {
		data.CiscoSigCredentialsFeatureTemplateId = types.StringValue(value.String())
	} else {
		data.CiscoSigCredentialsFeatureTemplateId = types.StringNull()
	}
	data.updateVersions(ctx, &state)
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin hasChanges
func (data *DNSSecurityPolicyDefinition) hasChanges(ctx context.Context, state *DNSSecurityPolicyDefinition) bool {
	hasChanges := false
	if !data.Name.Equal(state.Name) {
		hasChanges = true
	}
	if !data.Description.Equal(state.Description) {
		hasChanges = true
	}
	if !data.DomainListId.Equal(state.DomainListId) {
		hasChanges = true
	}
	if !data.LocalDomainBypassEnabled.Equal(state.LocalDomainBypassEnabled) {
		hasChanges = true
	}
	if !data.MatchAllVpn.Equal(state.MatchAllVpn) {
		hasChanges = true
	}
	if len(data.TargetVpns) != len(state.TargetVpns) {
		hasChanges = true
	} else {
		for i := range data.TargetVpns {
			if !data.TargetVpns[i].VpnIds.Equal(state.TargetVpns[i].VpnIds) {
				hasChanges = true
			}
			if !data.TargetVpns[i].UmbrellaDnsDefault.Equal(state.TargetVpns[i].UmbrellaDnsDefault) {
				hasChanges = true
			}
			if !data.TargetVpns[i].CustomDnsServerIp.Equal(state.TargetVpns[i].CustomDnsServerIp) {
				hasChanges = true
			}
			if !data.TargetVpns[i].LocalDomainBypassEnabled.Equal(state.TargetVpns[i].LocalDomainBypassEnabled) {
				hasChanges = true
			}
		}
	}
	if !data.Dnscrypt.Equal(state.Dnscrypt) {
		hasChanges = true
	}
	if !data.UmbrellaDnsDefault.Equal(state.UmbrellaDnsDefault) {
		hasChanges = true
	}
	if !data.CustomDnsServerIp.Equal(state.CustomDnsServerIp) {
		hasChanges = true
	}
	if !data.CiscoSigCredentialsFeatureTemplateId.Equal(state.CiscoSigCredentialsFeatureTemplateId) {
		hasChanges = true
	}
	return hasChanges
}

// End of section. //template:end hasChanges

// Section below is generated&owned by "gen/generator.go". //template:begin updateVersions

func (data *DNSSecurityPolicyDefinition) updateVersions(ctx context.Context, state *DNSSecurityPolicyDefinition) {
	data.DomainListVersion = state.DomainListVersion
	data.CiscoSigCredentialsFeatureTemplateVersion = state.CiscoSigCredentialsFeatureTemplateVersion
}

// End of section. //template:end updateVersions
