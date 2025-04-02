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
func (data *SystemLogging) fromBody(ctx context.Context, res gjson.Result) {
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
	if value := res.Get(path + "tlsProfile"); value.Exists() {
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
	}
	if value := res.Get(path + "server"); value.Exists() {
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
	}
	if value := res.Get(path + "ipv6Server"); value.Exists() {
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
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin updateFromBody
func (data *SystemLogging) updateFromBody(ctx context.Context, res gjson.Result) {
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
	for i := range data.TlsProfiles {
		keys := [...]string{"profile"}
		keyValues := [...]string{data.TlsProfiles[i].Profile.ValueString()}
		keyValuesVariables := [...]string{data.TlsProfiles[i].ProfileVariable.ValueString()}

		var r gjson.Result
		res.Get(path + "tlsProfile").ForEach(
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
		data.TlsProfiles[i].Profile = types.StringNull()
		data.TlsProfiles[i].ProfileVariable = types.StringNull()
		if t := r.Get("profile.optionType"); t.Exists() {
			va := r.Get("profile.value")
			if t.String() == "variable" {
				data.TlsProfiles[i].ProfileVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.TlsProfiles[i].Profile = types.StringValue(va.String())
			}
		}
		data.TlsProfiles[i].TlsVersion = types.StringNull()
		data.TlsProfiles[i].TlsVersionVariable = types.StringNull()
		if t := r.Get("tlsVersion.optionType"); t.Exists() {
			va := r.Get("tlsVersion.value")
			if t.String() == "variable" {
				data.TlsProfiles[i].TlsVersionVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.TlsProfiles[i].TlsVersion = types.StringValue(va.String())
			}
		}
		data.TlsProfiles[i].CipherSuites = types.SetNull(types.StringType)
		data.TlsProfiles[i].CipherSuitesVariable = types.StringNull()
		if t := r.Get("cipherSuiteList.optionType"); t.Exists() {
			va := r.Get("cipherSuiteList.value")
			if t.String() == "variable" {
				data.TlsProfiles[i].CipherSuitesVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.TlsProfiles[i].CipherSuites = helpers.GetStringSet(va.Array())
			}
		}
	}
	for i := range data.Ipv4Servers {
		keys := [...]string{"name"}
		keyValues := [...]string{data.Ipv4Servers[i].HostnameIp.ValueString()}
		keyValuesVariables := [...]string{data.Ipv4Servers[i].HostnameIpVariable.ValueString()}

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
		data.Ipv4Servers[i].HostnameIp = types.StringNull()
		data.Ipv4Servers[i].HostnameIpVariable = types.StringNull()
		if t := r.Get("name.optionType"); t.Exists() {
			va := r.Get("name.value")
			if t.String() == "variable" {
				data.Ipv4Servers[i].HostnameIpVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.Ipv4Servers[i].HostnameIp = types.StringValue(va.String())
			}
		}
		data.Ipv4Servers[i].Vpn = types.Int64Null()
		data.Ipv4Servers[i].VpnVariable = types.StringNull()
		if t := r.Get("vpn.optionType"); t.Exists() {
			va := r.Get("vpn.value")
			if t.String() == "variable" {
				data.Ipv4Servers[i].VpnVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.Ipv4Servers[i].Vpn = types.Int64Value(va.Int())
			}
		}
		data.Ipv4Servers[i].SourceInterface = types.StringNull()
		data.Ipv4Servers[i].SourceInterfaceVariable = types.StringNull()
		if t := r.Get("sourceInterface.optionType"); t.Exists() {
			va := r.Get("sourceInterface.value")
			if t.String() == "variable" {
				data.Ipv4Servers[i].SourceInterfaceVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.Ipv4Servers[i].SourceInterface = types.StringValue(va.String())
			}
		}
		data.Ipv4Servers[i].Priority = types.StringNull()
		data.Ipv4Servers[i].PriorityVariable = types.StringNull()
		if t := r.Get("priority.optionType"); t.Exists() {
			va := r.Get("priority.value")
			if t.String() == "variable" {
				data.Ipv4Servers[i].PriorityVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.Ipv4Servers[i].Priority = types.StringValue(va.String())
			}
		}
		data.Ipv4Servers[i].TlsEnable = types.BoolNull()
		data.Ipv4Servers[i].TlsEnableVariable = types.StringNull()
		if t := r.Get("tlsEnable.optionType"); t.Exists() {
			va := r.Get("tlsEnable.value")
			if t.String() == "variable" {
				data.Ipv4Servers[i].TlsEnableVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.Ipv4Servers[i].TlsEnable = types.BoolValue(va.Bool())
			}
		}
		data.Ipv4Servers[i].TlsPropertiesCustomProfile = types.BoolNull()
		data.Ipv4Servers[i].TlsPropertiesCustomProfileVariable = types.StringNull()
		if t := r.Get("tlsPropertiesCustomProfile.optionType"); t.Exists() {
			va := r.Get("tlsPropertiesCustomProfile.value")
			if t.String() == "variable" {
				data.Ipv4Servers[i].TlsPropertiesCustomProfileVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.Ipv4Servers[i].TlsPropertiesCustomProfile = types.BoolValue(va.Bool())
			}
		}
		data.Ipv4Servers[i].TlsPropertiesProfile = types.StringNull()
		data.Ipv4Servers[i].TlsPropertiesProfileVariable = types.StringNull()
		if t := r.Get("tlsPropertiesProfile.optionType"); t.Exists() {
			va := r.Get("tlsPropertiesProfile.value")
			if t.String() == "variable" {
				data.Ipv4Servers[i].TlsPropertiesProfileVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.Ipv4Servers[i].TlsPropertiesProfile = types.StringValue(va.String())
			}
		}
	}
	for i := range data.Ipv6Servers {
		keys := [...]string{"name"}
		keyValues := [...]string{data.Ipv6Servers[i].HostnameIp.ValueString()}
		keyValuesVariables := [...]string{data.Ipv6Servers[i].HostnameIpVariable.ValueString()}

		var r gjson.Result
		res.Get(path + "ipv6Server").ForEach(
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
		data.Ipv6Servers[i].HostnameIp = types.StringNull()
		data.Ipv6Servers[i].HostnameIpVariable = types.StringNull()
		if t := r.Get("name.optionType"); t.Exists() {
			va := r.Get("name.value")
			if t.String() == "variable" {
				data.Ipv6Servers[i].HostnameIpVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.Ipv6Servers[i].HostnameIp = types.StringValue(va.String())
			}
		}
		data.Ipv6Servers[i].Vpn = types.Int64Null()
		data.Ipv6Servers[i].VpnVariable = types.StringNull()
		if t := r.Get("vpn.optionType"); t.Exists() {
			va := r.Get("vpn.value")
			if t.String() == "variable" {
				data.Ipv6Servers[i].VpnVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.Ipv6Servers[i].Vpn = types.Int64Value(va.Int())
			}
		}
		data.Ipv6Servers[i].SourceInterface = types.StringNull()
		data.Ipv6Servers[i].SourceInterfaceVariable = types.StringNull()
		if t := r.Get("sourceInterface.optionType"); t.Exists() {
			va := r.Get("sourceInterface.value")
			if t.String() == "variable" {
				data.Ipv6Servers[i].SourceInterfaceVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.Ipv6Servers[i].SourceInterface = types.StringValue(va.String())
			}
		}
		data.Ipv6Servers[i].Priority = types.StringNull()
		data.Ipv6Servers[i].PriorityVariable = types.StringNull()
		if t := r.Get("priority.optionType"); t.Exists() {
			va := r.Get("priority.value")
			if t.String() == "variable" {
				data.Ipv6Servers[i].PriorityVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.Ipv6Servers[i].Priority = types.StringValue(va.String())
			}
		}
		data.Ipv6Servers[i].TlsEnable = types.BoolNull()
		data.Ipv6Servers[i].TlsEnableVariable = types.StringNull()
		if t := r.Get("tlsEnable.optionType"); t.Exists() {
			va := r.Get("tlsEnable.value")
			if t.String() == "variable" {
				data.Ipv6Servers[i].TlsEnableVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.Ipv6Servers[i].TlsEnable = types.BoolValue(va.Bool())
			}
		}
		data.Ipv6Servers[i].TlsPropertiesCustomProfile = types.BoolNull()
		data.Ipv6Servers[i].TlsPropertiesCustomProfileVariable = types.StringNull()
		if t := r.Get("tlsPropertiesCustomProfile.optionType"); t.Exists() {
			va := r.Get("tlsPropertiesCustomProfile.value")
			if t.String() == "variable" {
				data.Ipv6Servers[i].TlsPropertiesCustomProfileVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.Ipv6Servers[i].TlsPropertiesCustomProfile = types.BoolValue(va.Bool())
			}
		}
		data.Ipv6Servers[i].TlsPropertiesProfile = types.StringNull()
		data.Ipv6Servers[i].TlsPropertiesProfileVariable = types.StringNull()
		if t := r.Get("tlsPropertiesProfile.optionType"); t.Exists() {
			va := r.Get("tlsPropertiesProfile.value")
			if t.String() == "variable" {
				data.Ipv6Servers[i].TlsPropertiesProfileVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.Ipv6Servers[i].TlsPropertiesProfile = types.StringValue(va.String())
			}
		}
	}
}

// End of section. //template:end updateFromBody
