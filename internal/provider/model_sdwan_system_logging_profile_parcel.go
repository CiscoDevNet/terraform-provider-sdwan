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

// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	"net/url"

	"github.com/CiscoDevNet/terraform-provider-sdwan/internal/provider/helpers"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type SystemLogging struct {
	Id                     types.String              `tfsdk:"id"`
	Version                types.Int64               `tfsdk:"version"`
	Name                   types.String              `tfsdk:"name"`
	Description            types.String              `tfsdk:"description"`
	FeatureProfileId       types.String              `tfsdk:"feature_profile_id"`
	DiskEnable             types.Bool                `tfsdk:"disk_enable"`
	DiskEnableVariable     types.String              `tfsdk:"disk_enable_variable"`
	DiskFileSize           types.Int64               `tfsdk:"disk_file_size"`
	DiskFileSizeVariable   types.String              `tfsdk:"disk_file_size_variable"`
	DiskFileRotate         types.Int64               `tfsdk:"disk_file_rotate"`
	DiskFileRotateVariable types.String              `tfsdk:"disk_file_rotate_variable"`
	TlsProfile             []SystemLoggingTlsProfile `tfsdk:"tls_profile"`
	Server                 []SystemLoggingServer     `tfsdk:"server"`
	Ipv6Server             []SystemLoggingIpv6Server `tfsdk:"ipv6_server"`
}

type SystemLoggingTlsProfile struct {
	Profile                 types.String `tfsdk:"profile"`
	ProfileVariable         types.String `tfsdk:"profile_variable"`
	TlsVersion              types.String `tfsdk:"tls_version"`
	TlsVersionVariable      types.String `tfsdk:"tls_version_variable"`
	AuthType                types.String `tfsdk:"auth_type"`
	CipherSuiteList         types.Set    `tfsdk:"cipher_suite_list"`
	CipherSuiteListVariable types.String `tfsdk:"cipher_suite_list_variable"`
}

type SystemLoggingServer struct {
	Name                               types.String `tfsdk:"name"`
	NameVariable                       types.String `tfsdk:"name_variable"`
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

type SystemLoggingIpv6Server struct {
	Name                               types.String `tfsdk:"name"`
	NameVariable                       types.String `tfsdk:"name_variable"`
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

func (data SystemLogging) getModel() string {
	return "system_logging"
}

func (data SystemLogging) getPath() string {
	return fmt.Sprintf("/v1/feature-profile/sdwan/system/%v/logging", url.QueryEscape(data.FeatureProfileId.ValueString()))
}

func (data SystemLogging) toBody(ctx context.Context) string {
	body := ""
	body, _ = sjson.Set(body, "name", data.Name.ValueString())
	body, _ = sjson.Set(body, "description", data.Description.ValueString())
	path := "data."

	if !data.DiskEnableVariable.IsNull() {
		body, _ = sjson.Set(body, path+"disk.diskEnable.optionType", "variable")
		body, _ = sjson.Set(body, path+"disk.diskEnable.value", data.DiskEnableVariable.ValueString())
	} else if data.DiskEnable.IsNull() {
		body, _ = sjson.Set(body, path+"disk.diskEnable.optionType", "default")
		body, _ = sjson.Set(body, path+"disk.diskEnable.value", true)
	} else {
		body, _ = sjson.Set(body, path+"disk.diskEnable.optionType", "global")
		body, _ = sjson.Set(body, path+"disk.diskEnable.value", data.DiskEnable.ValueBool())
	}

	if !data.DiskFileSizeVariable.IsNull() {
		body, _ = sjson.Set(body, path+"disk.file.diskFileSize.optionType", "variable")
		body, _ = sjson.Set(body, path+"disk.file.diskFileSize.value", data.DiskFileSizeVariable.ValueString())
	} else if data.DiskFileSize.IsNull() {
		body, _ = sjson.Set(body, path+"disk.file.diskFileSize.optionType", "default")
		body, _ = sjson.Set(body, path+"disk.file.diskFileSize.value", 10)
	} else {
		body, _ = sjson.Set(body, path+"disk.file.diskFileSize.optionType", "global")
		body, _ = sjson.Set(body, path+"disk.file.diskFileSize.value", data.DiskFileSize.ValueInt64())
	}

	if !data.DiskFileRotateVariable.IsNull() {
		body, _ = sjson.Set(body, path+"disk.file.diskFileRotate.optionType", "variable")
		body, _ = sjson.Set(body, path+"disk.file.diskFileRotate.value", data.DiskFileRotateVariable.ValueString())
	} else if data.DiskFileRotate.IsNull() {
		body, _ = sjson.Set(body, path+"disk.file.diskFileRotate.optionType", "default")
		body, _ = sjson.Set(body, path+"disk.file.diskFileRotate.value", 10)
	} else {
		body, _ = sjson.Set(body, path+"disk.file.diskFileRotate.optionType", "global")
		body, _ = sjson.Set(body, path+"disk.file.diskFileRotate.value", data.DiskFileRotate.ValueInt64())
	}

	for _, item := range data.TlsProfile {
		itemBody := ""

		if !item.ProfileVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "profile.optionType", "variable")
			itemBody, _ = sjson.Set(itemBody, "profile.value", item.ProfileVariable.ValueString())
		} else if true {
			itemBody, _ = sjson.Set(itemBody, "profile.optionType", "global")
			itemBody, _ = sjson.Set(itemBody, "profile.value", item.Profile.ValueString())
		}

		if !item.TlsVersionVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "tlsVersion.optionType", "variable")
			itemBody, _ = sjson.Set(itemBody, "tlsVersion.value", item.TlsVersionVariable.ValueString())
		} else if item.TlsVersion.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "tlsVersion.optionType", "default")
			itemBody, _ = sjson.Set(itemBody, "tlsVersion.value", "TLSv1.1")
		} else {
			itemBody, _ = sjson.Set(itemBody, "tlsVersion.optionType", "global")
			itemBody, _ = sjson.Set(itemBody, "tlsVersion.value", item.TlsVersion.ValueString())
		}

		itemBody, _ = sjson.Set(itemBody, "authType.optionType", "default")
		itemBody, _ = sjson.Set(itemBody, "authType.value", "Server")

		if !item.CipherSuiteListVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "cipherSuiteList.optionType", "variable")
			itemBody, _ = sjson.Set(itemBody, "cipherSuiteList.value", item.CipherSuiteListVariable.ValueString())
		} else if true {
			itemBody, _ = sjson.Set(itemBody, "cipherSuiteList.optionType", "global")
			var values []string
			item.CipherSuiteList.ElementsAs(ctx, &values, false)
			itemBody, _ = sjson.Set(itemBody, "cipherSuiteList.value", values)
		}

		body, _ = sjson.SetRaw(body, path+"tlsProfile.-1", itemBody)
	}
	for _, item := range data.Server {
		itemBody := ""

		if !item.NameVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "name.optionType", "variable")
			itemBody, _ = sjson.Set(itemBody, "name.value", item.NameVariable.ValueString())
		} else if true {
			itemBody, _ = sjson.Set(itemBody, "name.optionType", "global")
			itemBody, _ = sjson.Set(itemBody, "name.value", item.Name.ValueString())
		}

		if !item.VpnVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "vpn.optionType", "variable")
			itemBody, _ = sjson.Set(itemBody, "vpn.value", item.VpnVariable.ValueString())
		} else if item.Vpn.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "vpn.optionType", "default")
			itemBody, _ = sjson.Set(itemBody, "vpn.value", 0)
		} else {
			itemBody, _ = sjson.Set(itemBody, "vpn.optionType", "global")
			itemBody, _ = sjson.Set(itemBody, "vpn.value", item.Vpn.ValueInt64())
		}

		if !item.SourceInterfaceVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "sourceInterface.optionType", "variable")
			itemBody, _ = sjson.Set(itemBody, "sourceInterface.value", item.SourceInterfaceVariable.ValueString())
		} else if true {
			itemBody, _ = sjson.Set(itemBody, "sourceInterface.optionType", "global")
			itemBody, _ = sjson.Set(itemBody, "sourceInterface.value", item.SourceInterface.ValueString())
		}

		if !item.PriorityVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "priority.optionType", "variable")
			itemBody, _ = sjson.Set(itemBody, "priority.value", item.PriorityVariable.ValueString())
		} else if item.Priority.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "priority.optionType", "default")
			itemBody, _ = sjson.Set(itemBody, "priority.value", "informational")
		} else {
			itemBody, _ = sjson.Set(itemBody, "priority.optionType", "global")
			itemBody, _ = sjson.Set(itemBody, "priority.value", item.Priority.ValueString())
		}

		if !item.TlsEnableVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "tlsEnable.optionType", "variable")
			itemBody, _ = sjson.Set(itemBody, "tlsEnable.value", item.TlsEnableVariable.ValueString())
		} else if item.TlsEnable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "tlsEnable.optionType", "default")
			itemBody, _ = sjson.Set(itemBody, "tlsEnable.value", false)
		} else {
			itemBody, _ = sjson.Set(itemBody, "tlsEnable.optionType", "global")
			itemBody, _ = sjson.Set(itemBody, "tlsEnable.value", item.TlsEnable.ValueBool())
		}

		if !item.TlsPropertiesCustomProfileVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "tlsPropertiesCustomProfile.optionType", "variable")
			itemBody, _ = sjson.Set(itemBody, "tlsPropertiesCustomProfile.value", item.TlsPropertiesCustomProfileVariable.ValueString())
		} else if item.TlsPropertiesCustomProfile.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "tlsPropertiesCustomProfile.optionType", "default")
			itemBody, _ = sjson.Set(itemBody, "tlsPropertiesCustomProfile.value", false)
		} else {
			itemBody, _ = sjson.Set(itemBody, "tlsPropertiesCustomProfile.optionType", "global")
			itemBody, _ = sjson.Set(itemBody, "tlsPropertiesCustomProfile.value", item.TlsPropertiesCustomProfile.ValueBool())
		}

		if !item.TlsPropertiesProfileVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "tlsPropertiesProfile.optionType", "variable")
			itemBody, _ = sjson.Set(itemBody, "tlsPropertiesProfile.value", item.TlsPropertiesProfileVariable.ValueString())
		} else if true {
			itemBody, _ = sjson.Set(itemBody, "tlsPropertiesProfile.optionType", "global")
			itemBody, _ = sjson.Set(itemBody, "tlsPropertiesProfile.value", item.TlsPropertiesProfile.ValueString())
		}

		body, _ = sjson.SetRaw(body, path+"server.-1", itemBody)
	}
	for _, item := range data.Ipv6Server {
		itemBody := ""

		if !item.NameVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "name.optionType", "variable")
			itemBody, _ = sjson.Set(itemBody, "name.value", item.NameVariable.ValueString())
		} else if true {
			itemBody, _ = sjson.Set(itemBody, "name.optionType", "global")
			itemBody, _ = sjson.Set(itemBody, "name.value", item.Name.ValueString())
		}

		if !item.VpnVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "vpn.optionType", "variable")
			itemBody, _ = sjson.Set(itemBody, "vpn.value", item.VpnVariable.ValueString())
		} else if item.Vpn.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "vpn.optionType", "default")
			itemBody, _ = sjson.Set(itemBody, "vpn.value", 0)
		} else {
			itemBody, _ = sjson.Set(itemBody, "vpn.optionType", "global")
			itemBody, _ = sjson.Set(itemBody, "vpn.value", item.Vpn.ValueInt64())
		}

		if !item.SourceInterfaceVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "sourceInterface.optionType", "variable")
			itemBody, _ = sjson.Set(itemBody, "sourceInterface.value", item.SourceInterfaceVariable.ValueString())
		} else if true {
			itemBody, _ = sjson.Set(itemBody, "sourceInterface.optionType", "global")
			itemBody, _ = sjson.Set(itemBody, "sourceInterface.value", item.SourceInterface.ValueString())
		}

		if !item.PriorityVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "priority.optionType", "variable")
			itemBody, _ = sjson.Set(itemBody, "priority.value", item.PriorityVariable.ValueString())
		} else if item.Priority.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "priority.optionType", "default")
			itemBody, _ = sjson.Set(itemBody, "priority.value", "informational")
		} else {
			itemBody, _ = sjson.Set(itemBody, "priority.optionType", "global")
			itemBody, _ = sjson.Set(itemBody, "priority.value", item.Priority.ValueString())
		}

		if !item.TlsEnableVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "tlsEnable.optionType", "variable")
			itemBody, _ = sjson.Set(itemBody, "tlsEnable.value", item.TlsEnableVariable.ValueString())
		} else if item.TlsEnable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "tlsEnable.optionType", "default")
			itemBody, _ = sjson.Set(itemBody, "tlsEnable.value", false)
		} else {
			itemBody, _ = sjson.Set(itemBody, "tlsEnable.optionType", "global")
			itemBody, _ = sjson.Set(itemBody, "tlsEnable.value", item.TlsEnable.ValueBool())
		}

		if !item.TlsPropertiesCustomProfileVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "tlsPropertiesCustomProfile.optionType", "variable")
			itemBody, _ = sjson.Set(itemBody, "tlsPropertiesCustomProfile.value", item.TlsPropertiesCustomProfileVariable.ValueString())
		} else if item.TlsPropertiesCustomProfile.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "tlsPropertiesCustomProfile.optionType", "default")
			itemBody, _ = sjson.Set(itemBody, "tlsPropertiesCustomProfile.value", false)
		} else {
			itemBody, _ = sjson.Set(itemBody, "tlsPropertiesCustomProfile.optionType", "global")
			itemBody, _ = sjson.Set(itemBody, "tlsPropertiesCustomProfile.value", item.TlsPropertiesCustomProfile.ValueBool())
		}

		if !item.TlsPropertiesProfileVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "tlsPropertiesProfile.optionType", "variable")
			itemBody, _ = sjson.Set(itemBody, "tlsPropertiesProfile.value", item.TlsPropertiesProfileVariable.ValueString())
		} else if true {
			itemBody, _ = sjson.Set(itemBody, "tlsPropertiesProfile.optionType", "global")
			itemBody, _ = sjson.Set(itemBody, "tlsPropertiesProfile.value", item.TlsPropertiesProfile.ValueString())
		}

		body, _ = sjson.SetRaw(body, path+"ipv6Server.-1", itemBody)
	}
	return body
}

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
		data.TlsProfile = make([]SystemLoggingTlsProfile, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := SystemLoggingTlsProfile{}
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
			item.AuthType = types.StringNull()

			if t := v.Get("authType.optionType"); t.Exists() {
				va := v.Get("authType.value")
				if t.String() == "global" {
					item.AuthType = types.StringValue(va.String())
				}
			}
			item.CipherSuiteList = types.SetNull(types.StringType)
			item.CipherSuiteListVariable = types.StringNull()
			if t := v.Get("cipherSuiteList.optionType"); t.Exists() {
				va := v.Get("cipherSuiteList.value")
				if t.String() == "variable" {
					item.CipherSuiteListVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.CipherSuiteList = helpers.GetStringSet(va.Array())
				}
			}
			data.TlsProfile = append(data.TlsProfile, item)
			return true
		})
	}
	if value := res.Get(path + "server"); value.Exists() {
		data.Server = make([]SystemLoggingServer, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := SystemLoggingServer{}
			item.Name = types.StringNull()
			item.NameVariable = types.StringNull()
			if t := v.Get("name.optionType"); t.Exists() {
				va := v.Get("name.value")
				if t.String() == "variable" {
					item.NameVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.Name = types.StringValue(va.String())
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
			data.Server = append(data.Server, item)
			return true
		})
	}
	if value := res.Get(path + "ipv6Server"); value.Exists() {
		data.Ipv6Server = make([]SystemLoggingIpv6Server, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := SystemLoggingIpv6Server{}
			item.Name = types.StringNull()
			item.NameVariable = types.StringNull()
			if t := v.Get("name.optionType"); t.Exists() {
				va := v.Get("name.value")
				if t.String() == "variable" {
					item.NameVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.Name = types.StringValue(va.String())
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
			data.Ipv6Server = append(data.Ipv6Server, item)
			return true
		})
	}
}

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
	for i := range data.TlsProfile {
		keys := [...]string{"profile"}
		keyValues := [...]string{data.TlsProfile[i].Profile.ValueString()}
		keyValuesVariables := [...]string{data.TlsProfile[i].ProfileVariable.ValueString()}

		var r gjson.Result
		res.Get(path + "tlsProfile").ForEach(
			func(_, v gjson.Result) bool {
				found := false
				for ik := range keys {
					tt := v.Get(keys[ik] + ".optionType").String()
					vv := v.Get(keys[ik] + ".value").String()
					if (tt == "variable" && vv == keyValuesVariables[ik]) || (tt == "global" && vv == keyValues[ik]) {
						found = true
						continue
					}
					found = false
					break
				}
				if found {
					r = v
					return false
				}
				return true
			},
		)
		data.TlsProfile[i].Profile = types.StringNull()
		data.TlsProfile[i].ProfileVariable = types.StringNull()
		if t := r.Get("profile.optionType"); t.Exists() {
			va := r.Get("profile.value")
			if t.String() == "variable" {
				data.TlsProfile[i].ProfileVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.TlsProfile[i].Profile = types.StringValue(va.String())
			}
		}
		data.TlsProfile[i].TlsVersion = types.StringNull()
		data.TlsProfile[i].TlsVersionVariable = types.StringNull()
		if t := r.Get("tlsVersion.optionType"); t.Exists() {
			va := r.Get("tlsVersion.value")
			if t.String() == "variable" {
				data.TlsProfile[i].TlsVersionVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.TlsProfile[i].TlsVersion = types.StringValue(va.String())
			}
		}
		data.TlsProfile[i].AuthType = types.StringNull()

		if t := r.Get("authType.optionType"); t.Exists() {
			va := r.Get("authType.value")
			if t.String() == "global" {
				data.TlsProfile[i].AuthType = types.StringValue(va.String())
			}
		}
		data.TlsProfile[i].CipherSuiteList = types.SetNull(types.StringType)
		data.TlsProfile[i].CipherSuiteListVariable = types.StringNull()
		if t := r.Get("cipherSuiteList.optionType"); t.Exists() {
			va := r.Get("cipherSuiteList.value")
			if t.String() == "variable" {
				data.TlsProfile[i].CipherSuiteListVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.TlsProfile[i].CipherSuiteList = helpers.GetStringSet(va.Array())
			}
		}
	}
	for i := range data.Server {
		keys := [...]string{"name"}
		keyValues := [...]string{data.Server[i].Name.ValueString()}
		keyValuesVariables := [...]string{data.Server[i].NameVariable.ValueString()}

		var r gjson.Result
		res.Get(path + "server").ForEach(
			func(_, v gjson.Result) bool {
				found := false
				for ik := range keys {
					tt := v.Get(keys[ik] + ".optionType").String()
					vv := v.Get(keys[ik] + ".value").String()
					if (tt == "variable" && vv == keyValuesVariables[ik]) || (tt == "global" && vv == keyValues[ik]) {
						found = true
						continue
					}
					found = false
					break
				}
				if found {
					r = v
					return false
				}
				return true
			},
		)
		data.Server[i].Name = types.StringNull()
		data.Server[i].NameVariable = types.StringNull()
		if t := r.Get("name.optionType"); t.Exists() {
			va := r.Get("name.value")
			if t.String() == "variable" {
				data.Server[i].NameVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.Server[i].Name = types.StringValue(va.String())
			}
		}
		data.Server[i].Vpn = types.Int64Null()
		data.Server[i].VpnVariable = types.StringNull()
		if t := r.Get("vpn.optionType"); t.Exists() {
			va := r.Get("vpn.value")
			if t.String() == "variable" {
				data.Server[i].VpnVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.Server[i].Vpn = types.Int64Value(va.Int())
			}
		}
		data.Server[i].SourceInterface = types.StringNull()
		data.Server[i].SourceInterfaceVariable = types.StringNull()
		if t := r.Get("sourceInterface.optionType"); t.Exists() {
			va := r.Get("sourceInterface.value")
			if t.String() == "variable" {
				data.Server[i].SourceInterfaceVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.Server[i].SourceInterface = types.StringValue(va.String())
			}
		}
		data.Server[i].Priority = types.StringNull()
		data.Server[i].PriorityVariable = types.StringNull()
		if t := r.Get("priority.optionType"); t.Exists() {
			va := r.Get("priority.value")
			if t.String() == "variable" {
				data.Server[i].PriorityVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.Server[i].Priority = types.StringValue(va.String())
			}
		}
		data.Server[i].TlsEnable = types.BoolNull()
		data.Server[i].TlsEnableVariable = types.StringNull()
		if t := r.Get("tlsEnable.optionType"); t.Exists() {
			va := r.Get("tlsEnable.value")
			if t.String() == "variable" {
				data.Server[i].TlsEnableVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.Server[i].TlsEnable = types.BoolValue(va.Bool())
			}
		}
		data.Server[i].TlsPropertiesCustomProfile = types.BoolNull()
		data.Server[i].TlsPropertiesCustomProfileVariable = types.StringNull()
		if t := r.Get("tlsPropertiesCustomProfile.optionType"); t.Exists() {
			va := r.Get("tlsPropertiesCustomProfile.value")
			if t.String() == "variable" {
				data.Server[i].TlsPropertiesCustomProfileVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.Server[i].TlsPropertiesCustomProfile = types.BoolValue(va.Bool())
			}
		}
		data.Server[i].TlsPropertiesProfile = types.StringNull()
		data.Server[i].TlsPropertiesProfileVariable = types.StringNull()
		if t := r.Get("tlsPropertiesProfile.optionType"); t.Exists() {
			va := r.Get("tlsPropertiesProfile.value")
			if t.String() == "variable" {
				data.Server[i].TlsPropertiesProfileVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.Server[i].TlsPropertiesProfile = types.StringValue(va.String())
			}
		}
	}
	for i := range data.Ipv6Server {
		keys := [...]string{"name"}
		keyValues := [...]string{data.Ipv6Server[i].Name.ValueString()}
		keyValuesVariables := [...]string{data.Ipv6Server[i].NameVariable.ValueString()}

		var r gjson.Result
		res.Get(path + "ipv6Server").ForEach(
			func(_, v gjson.Result) bool {
				found := false
				for ik := range keys {
					tt := v.Get(keys[ik] + ".optionType").String()
					vv := v.Get(keys[ik] + ".value").String()
					if (tt == "variable" && vv == keyValuesVariables[ik]) || (tt == "global" && vv == keyValues[ik]) {
						found = true
						continue
					}
					found = false
					break
				}
				if found {
					r = v
					return false
				}
				return true
			},
		)
		data.Ipv6Server[i].Name = types.StringNull()
		data.Ipv6Server[i].NameVariable = types.StringNull()
		if t := r.Get("name.optionType"); t.Exists() {
			va := r.Get("name.value")
			if t.String() == "variable" {
				data.Ipv6Server[i].NameVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.Ipv6Server[i].Name = types.StringValue(va.String())
			}
		}
		data.Ipv6Server[i].Vpn = types.Int64Null()
		data.Ipv6Server[i].VpnVariable = types.StringNull()
		if t := r.Get("vpn.optionType"); t.Exists() {
			va := r.Get("vpn.value")
			if t.String() == "variable" {
				data.Ipv6Server[i].VpnVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.Ipv6Server[i].Vpn = types.Int64Value(va.Int())
			}
		}
		data.Ipv6Server[i].SourceInterface = types.StringNull()
		data.Ipv6Server[i].SourceInterfaceVariable = types.StringNull()
		if t := r.Get("sourceInterface.optionType"); t.Exists() {
			va := r.Get("sourceInterface.value")
			if t.String() == "variable" {
				data.Ipv6Server[i].SourceInterfaceVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.Ipv6Server[i].SourceInterface = types.StringValue(va.String())
			}
		}
		data.Ipv6Server[i].Priority = types.StringNull()
		data.Ipv6Server[i].PriorityVariable = types.StringNull()
		if t := r.Get("priority.optionType"); t.Exists() {
			va := r.Get("priority.value")
			if t.String() == "variable" {
				data.Ipv6Server[i].PriorityVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.Ipv6Server[i].Priority = types.StringValue(va.String())
			}
		}
		data.Ipv6Server[i].TlsEnable = types.BoolNull()
		data.Ipv6Server[i].TlsEnableVariable = types.StringNull()
		if t := r.Get("tlsEnable.optionType"); t.Exists() {
			va := r.Get("tlsEnable.value")
			if t.String() == "variable" {
				data.Ipv6Server[i].TlsEnableVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.Ipv6Server[i].TlsEnable = types.BoolValue(va.Bool())
			}
		}
		data.Ipv6Server[i].TlsPropertiesCustomProfile = types.BoolNull()
		data.Ipv6Server[i].TlsPropertiesCustomProfileVariable = types.StringNull()
		if t := r.Get("tlsPropertiesCustomProfile.optionType"); t.Exists() {
			va := r.Get("tlsPropertiesCustomProfile.value")
			if t.String() == "variable" {
				data.Ipv6Server[i].TlsPropertiesCustomProfileVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.Ipv6Server[i].TlsPropertiesCustomProfile = types.BoolValue(va.Bool())
			}
		}
		data.Ipv6Server[i].TlsPropertiesProfile = types.StringNull()
		data.Ipv6Server[i].TlsPropertiesProfileVariable = types.StringNull()
		if t := r.Get("tlsPropertiesProfile.optionType"); t.Exists() {
			va := r.Get("tlsPropertiesProfile.value")
			if t.String() == "variable" {
				data.Ipv6Server[i].TlsPropertiesProfileVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.Ipv6Server[i].TlsPropertiesProfile = types.StringValue(va.String())
			}
		}
	}
}

func (data *SystemLogging) isNull(ctx context.Context, res gjson.Result) bool {
	if !data.FeatureProfileId.IsNull() {
		return false
	}
	if !data.DiskEnable.IsNull() {
		return false
	}
	if !data.DiskEnableVariable.IsNull() {
		return false
	}
	if !data.DiskFileSize.IsNull() {
		return false
	}
	if !data.DiskFileSizeVariable.IsNull() {
		return false
	}
	if !data.DiskFileRotate.IsNull() {
		return false
	}
	if !data.DiskFileRotateVariable.IsNull() {
		return false
	}
	if len(data.TlsProfile) > 0 {
		return false
	}
	if len(data.Server) > 0 {
		return false
	}
	if len(data.Ipv6Server) > 0 {
		return false
	}
	return true
}
