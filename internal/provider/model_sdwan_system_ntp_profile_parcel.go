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
type SystemNTP struct {
	Id                             types.String                  `tfsdk:"id"`
	Version                        types.Int64                   `tfsdk:"version"`
	Name                           types.String                  `tfsdk:"name"`
	Description                    types.String                  `tfsdk:"description"`
	FeatureProfileId               types.String                  `tfsdk:"feature_profile_id"`
	Servers                        []SystemNTPServers            `tfsdk:"servers"`
	AuthenticationKeys             []SystemNTPAuthenticationKeys `tfsdk:"authentication_keys"`
	TrustedKeys                    types.Set                     `tfsdk:"trusted_keys"`
	TrustedKeysVariable            types.String                  `tfsdk:"trusted_keys_variable"`
	AuthoritativeNtpServer         types.Bool                    `tfsdk:"authoritative_ntp_server"`
	AuthoritativeNtpServerVariable types.String                  `tfsdk:"authoritative_ntp_server_variable"`
	Stratum                        types.Int64                   `tfsdk:"stratum"`
	StratumVariable                types.String                  `tfsdk:"stratum_variable"`
	SourceInterface                types.String                  `tfsdk:"source_interface"`
	SourceInterfaceVariable        types.String                  `tfsdk:"source_interface_variable"`
}

type SystemNTPServers struct {
	HostnameIpAddress           types.String `tfsdk:"hostname_ip_address"`
	HostnameIpAddressVariable   types.String `tfsdk:"hostname_ip_address_variable"`
	AuthenticationKey           types.Int64  `tfsdk:"authentication_key"`
	AuthenticationKeyVariable   types.String `tfsdk:"authentication_key_variable"`
	Vpn                         types.Int64  `tfsdk:"vpn"`
	VpnVariable                 types.String `tfsdk:"vpn_variable"`
	NtpVersion                  types.Int64  `tfsdk:"ntp_version"`
	NtpVersionVariable          types.String `tfsdk:"ntp_version_variable"`
	SourceInterface             types.String `tfsdk:"source_interface"`
	SourceInterfaceVariable     types.String `tfsdk:"source_interface_variable"`
	PreferThisNtpServer         types.Bool   `tfsdk:"prefer_this_ntp_server"`
	PreferThisNtpServerVariable types.String `tfsdk:"prefer_this_ntp_server_variable"`
}

type SystemNTPAuthenticationKeys struct {
	KeyId            types.Int64  `tfsdk:"key_id"`
	KeyIdVariable    types.String `tfsdk:"key_id_variable"`
	Md5Value         types.String `tfsdk:"md5_value"`
	Md5ValueVariable types.String `tfsdk:"md5_value_variable"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getModel
func (data SystemNTP) getModel() string {
	return "system_ntp"
}

// End of section. //template:end getModel

// Section below is generated&owned by "gen/generator.go". //template:begin getPath
func (data SystemNTP) getPath() string {
	return fmt.Sprintf("/v1/feature-profile/sdwan/system/%v/ntp", url.QueryEscape(data.FeatureProfileId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody
func (data SystemNTP) toBody(ctx context.Context) string {
	body := ""
	body, _ = sjson.Set(body, "name", data.Name.ValueString())
	body, _ = sjson.Set(body, "description", data.Description.ValueString())
	path := "data."
	if true {
		body, _ = sjson.Set(body, path+"server", []interface{}{})
		for _, item := range data.Servers {
			itemBody := ""

			if !item.HostnameIpAddressVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "name.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "name.value", item.HostnameIpAddressVariable.ValueString())
				}
			} else if !item.HostnameIpAddress.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "name.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "name.value", item.HostnameIpAddress.ValueString())
				}
			}

			if !item.AuthenticationKeyVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "key.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "key.value", item.AuthenticationKeyVariable.ValueString())
				}
			} else if item.AuthenticationKey.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "key.optionType", "default")

				}
			} else {
				if true {
					itemBody, _ = sjson.Set(itemBody, "key.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "key.value", item.AuthenticationKey.ValueInt64())
				}
			}

			if !item.VpnVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "vpn.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "vpn.value", item.VpnVariable.ValueString())
				}
			} else if item.Vpn.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "vpn.optionType", "default")
					itemBody, _ = sjson.Set(itemBody, "vpn.value", 0)
				}
			} else {
				if true {
					itemBody, _ = sjson.Set(itemBody, "vpn.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "vpn.value", item.Vpn.ValueInt64())
				}
			}

			if !item.NtpVersionVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "version.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "version.value", item.NtpVersionVariable.ValueString())
				}
			} else if item.NtpVersion.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "version.optionType", "default")
					itemBody, _ = sjson.Set(itemBody, "version.value", 4)
				}
			} else {
				if true {
					itemBody, _ = sjson.Set(itemBody, "version.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "version.value", item.NtpVersion.ValueInt64())
				}
			}

			if !item.SourceInterfaceVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "sourceInterface.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "sourceInterface.value", item.SourceInterfaceVariable.ValueString())
				}
			} else if item.SourceInterface.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "sourceInterface.optionType", "default")

				}
			} else {
				if true {
					itemBody, _ = sjson.Set(itemBody, "sourceInterface.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "sourceInterface.value", item.SourceInterface.ValueString())
				}
			}

			if !item.PreferThisNtpServerVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "prefer.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "prefer.value", item.PreferThisNtpServerVariable.ValueString())
				}
			} else if item.PreferThisNtpServer.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "prefer.optionType", "default")
					itemBody, _ = sjson.Set(itemBody, "prefer.value", false)
				}
			} else {
				if true {
					itemBody, _ = sjson.Set(itemBody, "prefer.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "prefer.value", item.PreferThisNtpServer.ValueBool())
				}
			}
			body, _ = sjson.SetRaw(body, path+"server.-1", itemBody)
		}
	}
	if true {
		body, _ = sjson.Set(body, path+"authentication.authenticationKeys", []interface{}{})
		for _, item := range data.AuthenticationKeys {
			itemBody := ""

			if !item.KeyIdVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "keyId.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "keyId.value", item.KeyIdVariable.ValueString())
				}
			} else if !item.KeyId.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "keyId.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "keyId.value", item.KeyId.ValueInt64())
				}
			}

			if !item.Md5ValueVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "md5Value.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "md5Value.value", item.Md5ValueVariable.ValueString())
				}
			} else if !item.Md5Value.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "md5Value.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "md5Value.value", item.Md5Value.ValueString())
				}
			}
			body, _ = sjson.SetRaw(body, path+"authentication.authenticationKeys.-1", itemBody)
		}
	}

	if !data.TrustedKeysVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"authentication.trustedKeys.optionType", "variable")
			body, _ = sjson.Set(body, path+"authentication.trustedKeys.value", data.TrustedKeysVariable.ValueString())
		}
	} else if data.TrustedKeys.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"authentication.trustedKeys.optionType", "default")

		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"authentication.trustedKeys.optionType", "global")
			var values []int64
			data.TrustedKeys.ElementsAs(ctx, &values, false)
			body, _ = sjson.Set(body, path+"authentication.trustedKeys.value", values)
		}
	}

	if !data.AuthoritativeNtpServerVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"leader.enable.optionType", "variable")
			body, _ = sjson.Set(body, path+"leader.enable.value", data.AuthoritativeNtpServerVariable.ValueString())
		}
	} else if data.AuthoritativeNtpServer.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"leader.enable.optionType", "default")
			body, _ = sjson.Set(body, path+"leader.enable.value", false)
		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"leader.enable.optionType", "global")
			body, _ = sjson.Set(body, path+"leader.enable.value", data.AuthoritativeNtpServer.ValueBool())
		}
	}

	if !data.StratumVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"leader.stratum.optionType", "variable")
			body, _ = sjson.Set(body, path+"leader.stratum.value", data.StratumVariable.ValueString())
		}
	} else if data.Stratum.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"leader.stratum.optionType", "default")

		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"leader.stratum.optionType", "global")
			body, _ = sjson.Set(body, path+"leader.stratum.value", data.Stratum.ValueInt64())
		}
	}

	if !data.SourceInterfaceVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"leader.source.optionType", "variable")
			body, _ = sjson.Set(body, path+"leader.source.value", data.SourceInterfaceVariable.ValueString())
		}
	} else if data.SourceInterface.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"leader.source.optionType", "default")

		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"leader.source.optionType", "global")
			body, _ = sjson.Set(body, path+"leader.source.value", data.SourceInterface.ValueString())
		}
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody
func (data *SystemNTP) fromBody(ctx context.Context, res gjson.Result) {
	data.Name = types.StringValue(res.Get("payload.name").String())
	if value := res.Get("payload.description"); value.Exists() && value.String() != "" {
		data.Description = types.StringValue(value.String())
	} else {
		data.Description = types.StringNull()
	}
	path := "payload.data."
	if value := res.Get(path + "server"); value.Exists() {
		data.Servers = make([]SystemNTPServers, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := SystemNTPServers{}
			item.HostnameIpAddress = types.StringNull()
			item.HostnameIpAddressVariable = types.StringNull()
			if t := v.Get("name.optionType"); t.Exists() {
				va := v.Get("name.value")
				if t.String() == "variable" {
					item.HostnameIpAddressVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.HostnameIpAddress = types.StringValue(va.String())
				}
			}
			item.AuthenticationKey = types.Int64Null()
			item.AuthenticationKeyVariable = types.StringNull()
			if t := v.Get("key.optionType"); t.Exists() {
				va := v.Get("key.value")
				if t.String() == "variable" {
					item.AuthenticationKeyVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.AuthenticationKey = types.Int64Value(va.Int())
				}
			}
			item.Vpn = types.Int64Null()
			item.VpnVariable = types.StringNull()
			if t := v.Get("vpn.optionType"); t.Exists() {
				va := v.Get("vpn.value")
				if t.String() == "variable" {
					item.VpnVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.Vpn = types.Int64Value(va.Int())
				}
			}
			item.NtpVersion = types.Int64Null()
			item.NtpVersionVariable = types.StringNull()
			if t := v.Get("version.optionType"); t.Exists() {
				va := v.Get("version.value")
				if t.String() == "variable" {
					item.NtpVersionVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.NtpVersion = types.Int64Value(va.Int())
				}
			}
			item.SourceInterface = types.StringNull()
			item.SourceInterfaceVariable = types.StringNull()
			if t := v.Get("sourceInterface.optionType"); t.Exists() {
				va := v.Get("sourceInterface.value")
				if t.String() == "variable" {
					item.SourceInterfaceVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.SourceInterface = types.StringValue(va.String())
				}
			}
			item.PreferThisNtpServer = types.BoolNull()
			item.PreferThisNtpServerVariable = types.StringNull()
			if t := v.Get("prefer.optionType"); t.Exists() {
				va := v.Get("prefer.value")
				if t.String() == "variable" {
					item.PreferThisNtpServerVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.PreferThisNtpServer = types.BoolValue(va.Bool())
				}
			}
			data.Servers = append(data.Servers, item)
			return true
		})
	}
	if value := res.Get(path + "authentication.authenticationKeys"); value.Exists() {
		data.AuthenticationKeys = make([]SystemNTPAuthenticationKeys, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := SystemNTPAuthenticationKeys{}
			item.KeyId = types.Int64Null()
			item.KeyIdVariable = types.StringNull()
			if t := v.Get("keyId.optionType"); t.Exists() {
				va := v.Get("keyId.value")
				if t.String() == "variable" {
					item.KeyIdVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.KeyId = types.Int64Value(va.Int())
				}
			}
			item.Md5Value = types.StringNull()
			item.Md5ValueVariable = types.StringNull()
			if t := v.Get("md5Value.optionType"); t.Exists() {
				va := v.Get("md5Value.value")
				if t.String() == "variable" {
					item.Md5ValueVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.Md5Value = types.StringValue(va.String())
				}
			}
			data.AuthenticationKeys = append(data.AuthenticationKeys, item)
			return true
		})
	}
	data.TrustedKeys = types.SetNull(types.Int64Type)
	data.TrustedKeysVariable = types.StringNull()
	if t := res.Get(path + "authentication.trustedKeys.optionType"); t.Exists() {
		va := res.Get(path + "authentication.trustedKeys.value")
		if t.String() == "variable" {
			data.TrustedKeysVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.TrustedKeys = helpers.GetInt64Set(va.Array())
		}
	}
	data.AuthoritativeNtpServer = types.BoolNull()
	data.AuthoritativeNtpServerVariable = types.StringNull()
	if t := res.Get(path + "leader.enable.optionType"); t.Exists() {
		va := res.Get(path + "leader.enable.value")
		if t.String() == "variable" {
			data.AuthoritativeNtpServerVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.AuthoritativeNtpServer = types.BoolValue(va.Bool())
		}
	}
	data.Stratum = types.Int64Null()
	data.StratumVariable = types.StringNull()
	if t := res.Get(path + "leader.stratum.optionType"); t.Exists() {
		va := res.Get(path + "leader.stratum.value")
		if t.String() == "variable" {
			data.StratumVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.Stratum = types.Int64Value(va.Int())
		}
	}
	data.SourceInterface = types.StringNull()
	data.SourceInterfaceVariable = types.StringNull()
	if t := res.Get(path + "leader.source.optionType"); t.Exists() {
		va := res.Get(path + "leader.source.value")
		if t.String() == "variable" {
			data.SourceInterfaceVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.SourceInterface = types.StringValue(va.String())
		}
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin updateFromBody
func (data *SystemNTP) updateFromBody(ctx context.Context, res gjson.Result) {
	data.Name = types.StringValue(res.Get("payload.name").String())
	if value := res.Get("payload.description"); value.Exists() && value.String() != "" {
		data.Description = types.StringValue(value.String())
	} else {
		data.Description = types.StringNull()
	}
	path := "payload.data."
	for i := range data.Servers {
		keys := [...]string{"name"}
		keyValues := [...]string{data.Servers[i].HostnameIpAddress.ValueString()}
		keyValuesVariables := [...]string{data.Servers[i].HostnameIpAddressVariable.ValueString()}

		var r gjson.Result
		res.Get(path + "server").ForEach(
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
		data.Servers[i].HostnameIpAddress = types.StringNull()
		data.Servers[i].HostnameIpAddressVariable = types.StringNull()
		if t := r.Get("name.optionType"); t.Exists() {
			va := r.Get("name.value")
			if t.String() == "variable" {
				data.Servers[i].HostnameIpAddressVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.Servers[i].HostnameIpAddress = types.StringValue(va.String())
			}
		}
		data.Servers[i].AuthenticationKey = types.Int64Null()
		data.Servers[i].AuthenticationKeyVariable = types.StringNull()
		if t := r.Get("key.optionType"); t.Exists() {
			va := r.Get("key.value")
			if t.String() == "variable" {
				data.Servers[i].AuthenticationKeyVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.Servers[i].AuthenticationKey = types.Int64Value(va.Int())
			}
		}
		data.Servers[i].Vpn = types.Int64Null()
		data.Servers[i].VpnVariable = types.StringNull()
		if t := r.Get("vpn.optionType"); t.Exists() {
			va := r.Get("vpn.value")
			if t.String() == "variable" {
				data.Servers[i].VpnVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.Servers[i].Vpn = types.Int64Value(va.Int())
			}
		}
		data.Servers[i].NtpVersion = types.Int64Null()
		data.Servers[i].NtpVersionVariable = types.StringNull()
		if t := r.Get("version.optionType"); t.Exists() {
			va := r.Get("version.value")
			if t.String() == "variable" {
				data.Servers[i].NtpVersionVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.Servers[i].NtpVersion = types.Int64Value(va.Int())
			}
		}
		data.Servers[i].SourceInterface = types.StringNull()
		data.Servers[i].SourceInterfaceVariable = types.StringNull()
		if t := r.Get("sourceInterface.optionType"); t.Exists() {
			va := r.Get("sourceInterface.value")
			if t.String() == "variable" {
				data.Servers[i].SourceInterfaceVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.Servers[i].SourceInterface = types.StringValue(va.String())
			}
		}
		data.Servers[i].PreferThisNtpServer = types.BoolNull()
		data.Servers[i].PreferThisNtpServerVariable = types.StringNull()
		if t := r.Get("prefer.optionType"); t.Exists() {
			va := r.Get("prefer.value")
			if t.String() == "variable" {
				data.Servers[i].PreferThisNtpServerVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.Servers[i].PreferThisNtpServer = types.BoolValue(va.Bool())
			}
		}
	}
	for i := range data.AuthenticationKeys {
		keys := [...]string{"keyId"}
		keyValues := [...]string{strconv.FormatInt(data.AuthenticationKeys[i].KeyId.ValueInt64(), 10)}
		keyValuesVariables := [...]string{data.AuthenticationKeys[i].KeyIdVariable.ValueString()}

		var r gjson.Result
		res.Get(path + "authentication.authenticationKeys").ForEach(
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
		data.AuthenticationKeys[i].KeyId = types.Int64Null()
		data.AuthenticationKeys[i].KeyIdVariable = types.StringNull()
		if t := r.Get("keyId.optionType"); t.Exists() {
			va := r.Get("keyId.value")
			if t.String() == "variable" {
				data.AuthenticationKeys[i].KeyIdVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.AuthenticationKeys[i].KeyId = types.Int64Value(va.Int())
			}
		}
		data.AuthenticationKeys[i].Md5Value = types.StringNull()
		data.AuthenticationKeys[i].Md5ValueVariable = types.StringNull()
		if t := r.Get("md5Value.optionType"); t.Exists() {
			va := r.Get("md5Value.value")
			if t.String() == "variable" {
				data.AuthenticationKeys[i].Md5ValueVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.AuthenticationKeys[i].Md5Value = types.StringValue(va.String())
			}
		}
	}
	data.TrustedKeys = types.SetNull(types.Int64Type)
	data.TrustedKeysVariable = types.StringNull()
	if t := res.Get(path + "authentication.trustedKeys.optionType"); t.Exists() {
		va := res.Get(path + "authentication.trustedKeys.value")
		if t.String() == "variable" {
			data.TrustedKeysVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.TrustedKeys = helpers.GetInt64Set(va.Array())
		}
	}
	data.AuthoritativeNtpServer = types.BoolNull()
	data.AuthoritativeNtpServerVariable = types.StringNull()
	if t := res.Get(path + "leader.enable.optionType"); t.Exists() {
		va := res.Get(path + "leader.enable.value")
		if t.String() == "variable" {
			data.AuthoritativeNtpServerVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.AuthoritativeNtpServer = types.BoolValue(va.Bool())
		}
	}
	data.Stratum = types.Int64Null()
	data.StratumVariable = types.StringNull()
	if t := res.Get(path + "leader.stratum.optionType"); t.Exists() {
		va := res.Get(path + "leader.stratum.value")
		if t.String() == "variable" {
			data.StratumVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.Stratum = types.Int64Value(va.Int())
		}
	}
	data.SourceInterface = types.StringNull()
	data.SourceInterfaceVariable = types.StringNull()
	if t := res.Get(path + "leader.source.optionType"); t.Exists() {
		va := res.Get(path + "leader.source.value")
		if t.String() == "variable" {
			data.SourceInterfaceVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.SourceInterface = types.StringValue(va.String())
		}
	}
}

// End of section. //template:end updateFromBody

// Section below is generated&owned by "gen/generator.go". //template:begin isNull
func (data *SystemNTP) isNull(ctx context.Context, res gjson.Result) bool {
	if !data.FeatureProfileId.IsNull() {
		return false
	}
	if len(data.Servers) > 0 {
		return false
	}
	if len(data.AuthenticationKeys) > 0 {
		return false
	}
	if !data.TrustedKeys.IsNull() {
		return false
	}
	if !data.TrustedKeysVariable.IsNull() {
		return false
	}
	if !data.AuthoritativeNtpServer.IsNull() {
		return false
	}
	if !data.AuthoritativeNtpServerVariable.IsNull() {
		return false
	}
	if !data.Stratum.IsNull() {
		return false
	}
	if !data.StratumVariable.IsNull() {
		return false
	}
	if !data.SourceInterface.IsNull() {
		return false
	}
	if !data.SourceInterfaceVariable.IsNull() {
		return false
	}
	return true
}

// End of section. //template:end isNull
