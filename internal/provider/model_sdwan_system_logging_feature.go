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

	"github.com/CiscoDevNet/terraform-provider-sdwan/internal/provider/helpers"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types
type SystemLogging struct {
	Id                     types.String               `tfsdk:"id"`
	Version                types.Int64                `tfsdk:"version"`
	Name                   types.String               `tfsdk:"name"`
	Description            types.String               `tfsdk:"description"`
	FeatureProfileId       types.String               `tfsdk:"feature_profile_id"`
	DiskEnable             types.Bool                 `tfsdk:"disk_enable"`
	DiskEnableVariable     types.String               `tfsdk:"disk_enable_variable"`
	DiskFileSize           types.Int64                `tfsdk:"disk_file_size"`
	DiskFileSizeVariable   types.String               `tfsdk:"disk_file_size_variable"`
	DiskFileRotate         types.Int64                `tfsdk:"disk_file_rotate"`
	DiskFileRotateVariable types.String               `tfsdk:"disk_file_rotate_variable"`
	TlsProfiles            []SystemLoggingTlsProfiles `tfsdk:"tls_profiles"`
	Ipv4Servers            []SystemLoggingIpv4Servers `tfsdk:"ipv4_servers"`
	Ipv6Servers            []SystemLoggingIpv6Servers `tfsdk:"ipv6_servers"`
}

type SystemLoggingTlsProfiles struct {
	Profile              types.String `tfsdk:"profile"`
	ProfileVariable      types.String `tfsdk:"profile_variable"`
	TlsVersion           types.String `tfsdk:"tls_version"`
	TlsVersionVariable   types.String `tfsdk:"tls_version_variable"`
	CipherSuites         types.Set    `tfsdk:"cipher_suites"`
	CipherSuitesVariable types.String `tfsdk:"cipher_suites_variable"`
}

type SystemLoggingIpv4Servers struct {
	HostnameIp                         types.String `tfsdk:"hostname_ip"`
	HostnameIpVariable                 types.String `tfsdk:"hostname_ip_variable"`
	Vpn                                types.Int64  `tfsdk:"vpn"`
	VpnVariable                        types.String `tfsdk:"vpn_variable"`
	SourceInterface                    types.String `tfsdk:"source_interface"`
	SourceInterfaceVariable            types.String `tfsdk:"source_interface_variable"`
	Priority                           types.String `tfsdk:"priority"`
	PriorityVariable                   types.String `tfsdk:"priority_variable"`
	TlsEnable                          types.Bool   `tfsdk:"tls_enable"`
	TlsEnableVariable                  types.String `tfsdk:"tls_enable_variable"`
	TlsPropertiesCustomProfile         types.Bool   `tfsdk:"tls_properties_custom_profile"`
	TlsPropertiesCustomProfileVariable types.String `tfsdk:"tls_properties_custom_profile_variable"`
	TlsPropertiesProfile               types.String `tfsdk:"tls_properties_profile"`
	TlsPropertiesProfileVariable       types.String `tfsdk:"tls_properties_profile_variable"`
}

type SystemLoggingIpv6Servers struct {
	HostnameIp                         types.String `tfsdk:"hostname_ip"`
	HostnameIpVariable                 types.String `tfsdk:"hostname_ip_variable"`
	Vpn                                types.Int64  `tfsdk:"vpn"`
	VpnVariable                        types.String `tfsdk:"vpn_variable"`
	SourceInterface                    types.String `tfsdk:"source_interface"`
	SourceInterfaceVariable            types.String `tfsdk:"source_interface_variable"`
	Priority                           types.String `tfsdk:"priority"`
	PriorityVariable                   types.String `tfsdk:"priority_variable"`
	TlsEnable                          types.Bool   `tfsdk:"tls_enable"`
	TlsEnableVariable                  types.String `tfsdk:"tls_enable_variable"`
	TlsPropertiesCustomProfile         types.Bool   `tfsdk:"tls_properties_custom_profile"`
	TlsPropertiesCustomProfileVariable types.String `tfsdk:"tls_properties_custom_profile_variable"`
	TlsPropertiesProfile               types.String `tfsdk:"tls_properties_profile"`
	TlsPropertiesProfileVariable       types.String `tfsdk:"tls_properties_profile_variable"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getModel
func (data SystemLogging) getModel() string {
	return "system_logging"
}

// End of section. //template:end getModel

// Section below is generated&owned by "gen/generator.go". //template:begin getPath
func (data SystemLogging) getPath() string {
	return fmt.Sprintf("/v1/feature-profile/sdwan/system/%v/logging", url.QueryEscape(data.FeatureProfileId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody
func (data SystemLogging) toBody(ctx context.Context) string {
	body := ""
	body, _ = sjson.Set(body, "name", data.Name.ValueString())
	body, _ = sjson.Set(body, "description", data.Description.ValueString())
	path := "data."

	if !data.DiskEnableVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"disk.diskEnable.optionType", "variable")
			body, _ = sjson.Set(body, path+"disk.diskEnable.value", data.DiskEnableVariable.ValueString())
		}
	} else if data.DiskEnable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"disk.diskEnable.optionType", "default")
			body, _ = sjson.Set(body, path+"disk.diskEnable.value", true)
		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"disk.diskEnable.optionType", "global")
			body, _ = sjson.Set(body, path+"disk.diskEnable.value", data.DiskEnable.ValueBool())
		}
	}

	if !data.DiskFileSizeVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"disk.file.diskFileSize.optionType", "variable")
			body, _ = sjson.Set(body, path+"disk.file.diskFileSize.value", data.DiskFileSizeVariable.ValueString())
		}
	} else if data.DiskFileSize.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"disk.file.diskFileSize.optionType", "default")
			body, _ = sjson.Set(body, path+"disk.file.diskFileSize.value", 10)
		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"disk.file.diskFileSize.optionType", "global")
			body, _ = sjson.Set(body, path+"disk.file.diskFileSize.value", data.DiskFileSize.ValueInt64())
		}
	}

	if !data.DiskFileRotateVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"disk.file.diskFileRotate.optionType", "variable")
			body, _ = sjson.Set(body, path+"disk.file.diskFileRotate.value", data.DiskFileRotateVariable.ValueString())
		}
	} else if data.DiskFileRotate.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"disk.file.diskFileRotate.optionType", "default")
			body, _ = sjson.Set(body, path+"disk.file.diskFileRotate.value", 10)
		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"disk.file.diskFileRotate.optionType", "global")
			body, _ = sjson.Set(body, path+"disk.file.diskFileRotate.value", data.DiskFileRotate.ValueInt64())
		}
	}
	if true {
		body, _ = sjson.Set(body, path+"tlsProfile", []interface{}{})
		for _, item := range data.TlsProfiles {
			itemBody := ""

			if !item.ProfileVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "profile.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "profile.value", item.ProfileVariable.ValueString())
				}
			} else if !item.Profile.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "profile.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "profile.value", item.Profile.ValueString())
				}
			}

			if !item.TlsVersionVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "tlsVersion.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "tlsVersion.value", item.TlsVersionVariable.ValueString())
				}
			} else if item.TlsVersion.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "tlsVersion.optionType", "default")
					itemBody, _ = sjson.Set(itemBody, "tlsVersion.value", "TLSv1.1")
				}
			} else {
				if true {
					itemBody, _ = sjson.Set(itemBody, "tlsVersion.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "tlsVersion.value", item.TlsVersion.ValueString())
				}
			}
			if true {
				itemBody, _ = sjson.Set(itemBody, "authType.optionType", "default")
				itemBody, _ = sjson.Set(itemBody, "authType.value", "Server")
			}

			if !item.CipherSuitesVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "cipherSuiteList.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "cipherSuiteList.value", item.CipherSuitesVariable.ValueString())
				}
			} else if item.CipherSuites.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "cipherSuiteList.optionType", "default")

				}
			} else {
				if true {
					itemBody, _ = sjson.Set(itemBody, "cipherSuiteList.optionType", "global")
					var values []string
					item.CipherSuites.ElementsAs(ctx, &values, false)
					itemBody, _ = sjson.Set(itemBody, "cipherSuiteList.value", values)
				}
			}
			body, _ = sjson.SetRaw(body, path+"tlsProfile.-1", itemBody)
		}
	}
	if true {
		body, _ = sjson.Set(body, path+"server", []interface{}{})
		for _, item := range data.Ipv4Servers {
			itemBody := ""

			if !item.HostnameIpVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "name.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "name.value", item.HostnameIpVariable.ValueString())
				}
			} else if !item.HostnameIp.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "name.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "name.value", item.HostnameIp.ValueString())
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

			if !item.PriorityVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "priority.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "priority.value", item.PriorityVariable.ValueString())
				}
			} else if item.Priority.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "priority.optionType", "default")
					itemBody, _ = sjson.Set(itemBody, "priority.value", "informational")
				}
			} else {
				if true {
					itemBody, _ = sjson.Set(itemBody, "priority.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "priority.value", item.Priority.ValueString())
				}
			}

			if !item.TlsEnableVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "tlsEnable.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "tlsEnable.value", item.TlsEnableVariable.ValueString())
				}
			} else if item.TlsEnable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "tlsEnable.optionType", "default")
					itemBody, _ = sjson.Set(itemBody, "tlsEnable.value", false)
				}
			} else {
				if true {
					itemBody, _ = sjson.Set(itemBody, "tlsEnable.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "tlsEnable.value", item.TlsEnable.ValueBool())
				}
			}

			if !item.TlsPropertiesCustomProfileVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "tlsPropertiesCustomProfile.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "tlsPropertiesCustomProfile.value", item.TlsPropertiesCustomProfileVariable.ValueString())
				}
			} else if item.TlsPropertiesCustomProfile.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "tlsPropertiesCustomProfile.optionType", "default")
					itemBody, _ = sjson.Set(itemBody, "tlsPropertiesCustomProfile.value", false)
				}
			} else {
				if true {
					itemBody, _ = sjson.Set(itemBody, "tlsPropertiesCustomProfile.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "tlsPropertiesCustomProfile.value", item.TlsPropertiesCustomProfile.ValueBool())
				}
			}

			if !item.TlsPropertiesProfileVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "tlsPropertiesProfile.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "tlsPropertiesProfile.value", item.TlsPropertiesProfileVariable.ValueString())
				}
			} else if item.TlsPropertiesProfile.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "tlsPropertiesProfile.optionType", "default")

				}
			} else {
				if true {
					itemBody, _ = sjson.Set(itemBody, "tlsPropertiesProfile.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "tlsPropertiesProfile.value", item.TlsPropertiesProfile.ValueString())
				}
			}
			body, _ = sjson.SetRaw(body, path+"server.-1", itemBody)
		}
	}
	if true {
		body, _ = sjson.Set(body, path+"ipv6Server", []interface{}{})
		for _, item := range data.Ipv6Servers {
			itemBody := ""

			if !item.HostnameIpVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "name.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "name.value", item.HostnameIpVariable.ValueString())
				}
			} else if !item.HostnameIp.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "name.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "name.value", item.HostnameIp.ValueString())
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

			if !item.PriorityVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "priority.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "priority.value", item.PriorityVariable.ValueString())
				}
			} else if item.Priority.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "priority.optionType", "default")
					itemBody, _ = sjson.Set(itemBody, "priority.value", "informational")
				}
			} else {
				if true {
					itemBody, _ = sjson.Set(itemBody, "priority.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "priority.value", item.Priority.ValueString())
				}
			}

			if !item.TlsEnableVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "tlsEnable.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "tlsEnable.value", item.TlsEnableVariable.ValueString())
				}
			} else if item.TlsEnable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "tlsEnable.optionType", "default")
					itemBody, _ = sjson.Set(itemBody, "tlsEnable.value", false)
				}
			} else {
				if true {
					itemBody, _ = sjson.Set(itemBody, "tlsEnable.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "tlsEnable.value", item.TlsEnable.ValueBool())
				}
			}

			if !item.TlsPropertiesCustomProfileVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "tlsPropertiesCustomProfile.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "tlsPropertiesCustomProfile.value", item.TlsPropertiesCustomProfileVariable.ValueString())
				}
			} else if item.TlsPropertiesCustomProfile.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "tlsPropertiesCustomProfile.optionType", "default")
					itemBody, _ = sjson.Set(itemBody, "tlsPropertiesCustomProfile.value", false)
				}
			} else {
				if true {
					itemBody, _ = sjson.Set(itemBody, "tlsPropertiesCustomProfile.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "tlsPropertiesCustomProfile.value", item.TlsPropertiesCustomProfile.ValueBool())
				}
			}

			if !item.TlsPropertiesProfileVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "tlsPropertiesProfile.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "tlsPropertiesProfile.value", item.TlsPropertiesProfileVariable.ValueString())
				}
			} else if item.TlsPropertiesProfile.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "tlsPropertiesProfile.optionType", "default")

				}
			} else {
				if true {
					itemBody, _ = sjson.Set(itemBody, "tlsPropertiesProfile.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "tlsPropertiesProfile.value", item.TlsPropertiesProfile.ValueString())
				}
			}
			body, _ = sjson.SetRaw(body, path+"ipv6Server.-1", itemBody)
		}
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody
func (data *SystemLogging) fromBody(ctx context.Context, res gjson.Result, fullRead bool) {
	data.Name = types.StringValue(res.Get("payload.name").String())
	if value := res.Get("payload.description"); value.Exists() && value.String() != "" {
		data.Description = types.StringValue(value.String())
	} else {
		data.Description = types.StringNull()
	}
	path := "payload.data."
	data.DiskEnable = types.BoolNull()
	data.DiskEnableVariable = types.StringNull()
	if t := res.Get(path + "disk.diskEnable.optionType"); t.Exists() {
		va := res.Get(path + "disk.diskEnable.value")
		if t.String() == "variable" {
			data.DiskEnableVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.DiskEnable = types.BoolValue(va.Bool())
		}
	}
	data.DiskFileSize = types.Int64Null()
	data.DiskFileSizeVariable = types.StringNull()
	if t := res.Get(path + "disk.file.diskFileSize.optionType"); t.Exists() {
		va := res.Get(path + "disk.file.diskFileSize.value")
		if t.String() == "variable" {
			data.DiskFileSizeVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.DiskFileSize = types.Int64Value(va.Int())
		}
	}
	data.DiskFileRotate = types.Int64Null()
	data.DiskFileRotateVariable = types.StringNull()
	if t := res.Get(path + "disk.file.diskFileRotate.optionType"); t.Exists() {
		va := res.Get(path + "disk.file.diskFileRotate.value")
		if t.String() == "variable" {
			data.DiskFileRotateVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.DiskFileRotate = types.Int64Value(va.Int())
		}
	}
	oldTlsProfiles := data.TlsProfiles
	if value := res.Get(path + "tlsProfile"); value.Exists() && len(value.Array()) > 0 {
		data.TlsProfiles = make([]SystemLoggingTlsProfiles, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := SystemLoggingTlsProfiles{}
			item.Profile = types.StringNull()
			item.ProfileVariable = types.StringNull()
			if t := v.Get("profile.optionType"); t.Exists() {
				va := v.Get("profile.value")
				if t.String() == "variable" {
					item.ProfileVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.Profile = types.StringValue(va.String())
				}
			}
			item.TlsVersion = types.StringNull()
			item.TlsVersionVariable = types.StringNull()
			if t := v.Get("tlsVersion.optionType"); t.Exists() {
				va := v.Get("tlsVersion.value")
				if t.String() == "variable" {
					item.TlsVersionVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.TlsVersion = types.StringValue(va.String())
				}
			}
			item.CipherSuites = types.SetNull(types.StringType)
			item.CipherSuitesVariable = types.StringNull()
			if t := v.Get("cipherSuiteList.optionType"); t.Exists() {
				va := v.Get("cipherSuiteList.value")
				if t.String() == "variable" {
					item.CipherSuitesVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.CipherSuites = helpers.GetStringSet(va.Array())
				}
			}
			data.TlsProfiles = append(data.TlsProfiles, item)
			return true
		})
	} else {
		data.TlsProfiles = nil
	}
	if !fullRead {
		resultTlsProfiles := make([]SystemLoggingTlsProfiles, 0, len(data.TlsProfiles))
		matchedTlsProfiles := make([]bool, len(data.TlsProfiles))
		for _, oldItem := range oldTlsProfiles {
			for ni := range data.TlsProfiles {
				if matchedTlsProfiles[ni] {
					continue
				}
				keyMatch := true
				if keyMatch && (oldItem.ProfileVariable.ValueString() != "" || data.TlsProfiles[ni].ProfileVariable.ValueString() != "") {
					if oldItem.ProfileVariable.ValueString() != data.TlsProfiles[ni].ProfileVariable.ValueString() {
						keyMatch = false
					}
				} else if keyMatch {
					if oldItem.Profile.ValueString() != data.TlsProfiles[ni].Profile.ValueString() {
						keyMatch = false
					}
				}
				if keyMatch {
					matchedTlsProfiles[ni] = true
					resultTlsProfiles = append(resultTlsProfiles, data.TlsProfiles[ni])
					break
				}
			}
		}
		for ni := range data.TlsProfiles {
			if !matchedTlsProfiles[ni] {
				resultTlsProfiles = append(resultTlsProfiles, data.TlsProfiles[ni])
			}
		}
		data.TlsProfiles = resultTlsProfiles
	}
	oldIpv4Servers := data.Ipv4Servers
	if value := res.Get(path + "server"); value.Exists() && len(value.Array()) > 0 {
		data.Ipv4Servers = make([]SystemLoggingIpv4Servers, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := SystemLoggingIpv4Servers{}
			item.HostnameIp = types.StringNull()
			item.HostnameIpVariable = types.StringNull()
			if t := v.Get("name.optionType"); t.Exists() {
				va := v.Get("name.value")
				if t.String() == "variable" {
					item.HostnameIpVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.HostnameIp = types.StringValue(va.String())
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
			item.Priority = types.StringNull()
			item.PriorityVariable = types.StringNull()
			if t := v.Get("priority.optionType"); t.Exists() {
				va := v.Get("priority.value")
				if t.String() == "variable" {
					item.PriorityVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.Priority = types.StringValue(va.String())
				}
			}
			item.TlsEnable = types.BoolNull()
			item.TlsEnableVariable = types.StringNull()
			if t := v.Get("tlsEnable.optionType"); t.Exists() {
				va := v.Get("tlsEnable.value")
				if t.String() == "variable" {
					item.TlsEnableVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.TlsEnable = types.BoolValue(va.Bool())
				}
			}
			item.TlsPropertiesCustomProfile = types.BoolNull()
			item.TlsPropertiesCustomProfileVariable = types.StringNull()
			if t := v.Get("tlsPropertiesCustomProfile.optionType"); t.Exists() {
				va := v.Get("tlsPropertiesCustomProfile.value")
				if t.String() == "variable" {
					item.TlsPropertiesCustomProfileVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.TlsPropertiesCustomProfile = types.BoolValue(va.Bool())
				}
			}
			item.TlsPropertiesProfile = types.StringNull()
			item.TlsPropertiesProfileVariable = types.StringNull()
			if t := v.Get("tlsPropertiesProfile.optionType"); t.Exists() {
				va := v.Get("tlsPropertiesProfile.value")
				if t.String() == "variable" {
					item.TlsPropertiesProfileVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.TlsPropertiesProfile = types.StringValue(va.String())
				}
			}
			data.Ipv4Servers = append(data.Ipv4Servers, item)
			return true
		})
	} else {
		data.Ipv4Servers = nil
	}
	if !fullRead {
		resultIpv4Servers := make([]SystemLoggingIpv4Servers, 0, len(data.Ipv4Servers))
		matchedIpv4Servers := make([]bool, len(data.Ipv4Servers))
		for _, oldItem := range oldIpv4Servers {
			for ni := range data.Ipv4Servers {
				if matchedIpv4Servers[ni] {
					continue
				}
				keyMatch := true
				if keyMatch && (oldItem.HostnameIpVariable.ValueString() != "" || data.Ipv4Servers[ni].HostnameIpVariable.ValueString() != "") {
					if oldItem.HostnameIpVariable.ValueString() != data.Ipv4Servers[ni].HostnameIpVariable.ValueString() {
						keyMatch = false
					}
				} else if keyMatch {
					if oldItem.HostnameIp.ValueString() != data.Ipv4Servers[ni].HostnameIp.ValueString() {
						keyMatch = false
					}
				}
				if keyMatch {
					matchedIpv4Servers[ni] = true
					resultIpv4Servers = append(resultIpv4Servers, data.Ipv4Servers[ni])
					break
				}
			}
		}
		for ni := range data.Ipv4Servers {
			if !matchedIpv4Servers[ni] {
				resultIpv4Servers = append(resultIpv4Servers, data.Ipv4Servers[ni])
			}
		}
		data.Ipv4Servers = resultIpv4Servers
	}
	oldIpv6Servers := data.Ipv6Servers
	if value := res.Get(path + "ipv6Server"); value.Exists() && len(value.Array()) > 0 {
		data.Ipv6Servers = make([]SystemLoggingIpv6Servers, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := SystemLoggingIpv6Servers{}
			item.HostnameIp = types.StringNull()
			item.HostnameIpVariable = types.StringNull()
			if t := v.Get("name.optionType"); t.Exists() {
				va := v.Get("name.value")
				if t.String() == "variable" {
					item.HostnameIpVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.HostnameIp = types.StringValue(va.String())
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
			item.Priority = types.StringNull()
			item.PriorityVariable = types.StringNull()
			if t := v.Get("priority.optionType"); t.Exists() {
				va := v.Get("priority.value")
				if t.String() == "variable" {
					item.PriorityVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.Priority = types.StringValue(va.String())
				}
			}
			item.TlsEnable = types.BoolNull()
			item.TlsEnableVariable = types.StringNull()
			if t := v.Get("tlsEnable.optionType"); t.Exists() {
				va := v.Get("tlsEnable.value")
				if t.String() == "variable" {
					item.TlsEnableVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.TlsEnable = types.BoolValue(va.Bool())
				}
			}
			item.TlsPropertiesCustomProfile = types.BoolNull()
			item.TlsPropertiesCustomProfileVariable = types.StringNull()
			if t := v.Get("tlsPropertiesCustomProfile.optionType"); t.Exists() {
				va := v.Get("tlsPropertiesCustomProfile.value")
				if t.String() == "variable" {
					item.TlsPropertiesCustomProfileVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.TlsPropertiesCustomProfile = types.BoolValue(va.Bool())
				}
			}
			item.TlsPropertiesProfile = types.StringNull()
			item.TlsPropertiesProfileVariable = types.StringNull()
			if t := v.Get("tlsPropertiesProfile.optionType"); t.Exists() {
				va := v.Get("tlsPropertiesProfile.value")
				if t.String() == "variable" {
					item.TlsPropertiesProfileVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.TlsPropertiesProfile = types.StringValue(va.String())
				}
			}
			data.Ipv6Servers = append(data.Ipv6Servers, item)
			return true
		})
	} else {
		data.Ipv6Servers = nil
	}
	if !fullRead {
		resultIpv6Servers := make([]SystemLoggingIpv6Servers, 0, len(data.Ipv6Servers))
		matchedIpv6Servers := make([]bool, len(data.Ipv6Servers))
		for _, oldItem := range oldIpv6Servers {
			for ni := range data.Ipv6Servers {
				if matchedIpv6Servers[ni] {
					continue
				}
				keyMatch := true
				if keyMatch && (oldItem.HostnameIpVariable.ValueString() != "" || data.Ipv6Servers[ni].HostnameIpVariable.ValueString() != "") {
					if oldItem.HostnameIpVariable.ValueString() != data.Ipv6Servers[ni].HostnameIpVariable.ValueString() {
						keyMatch = false
					}
				} else if keyMatch {
					if oldItem.HostnameIp.ValueString() != data.Ipv6Servers[ni].HostnameIp.ValueString() {
						keyMatch = false
					}
				}
				if keyMatch {
					matchedIpv6Servers[ni] = true
					resultIpv6Servers = append(resultIpv6Servers, data.Ipv6Servers[ni])
					break
				}
			}
		}
		for ni := range data.Ipv6Servers {
			if !matchedIpv6Servers[ni] {
				resultIpv6Servers = append(resultIpv6Servers, data.Ipv6Servers[ni])
			}
		}
		data.Ipv6Servers = resultIpv6Servers
	}
}

// End of section. //template:end fromBody
