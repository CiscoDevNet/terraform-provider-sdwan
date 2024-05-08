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
type SystemRemoteAccess struct {
	Id                                         types.String `tfsdk:"id"`
	Version                                    types.Int64  `tfsdk:"version"`
	Name                                       types.String `tfsdk:"name"`
	Description                                types.String `tfsdk:"description"`
	FeatureProfileId                           types.String `tfsdk:"feature_profile_id"`
	ConnectionTypeSsl                          types.Bool   `tfsdk:"connection_type_ssl"`
	AnyConnectEapAuthenticationType            types.String `tfsdk:"any_connect_eap_authentication_type"`
	AnyConnectEapProfileDownloadStatus         types.String `tfsdk:"any_connect_eap_profile_download_status"`
	AnyConnectEapProfileDownloadStatusVariable types.String `tfsdk:"any_connect_eap_profile_download_status_variable"`
	AnyConnectEapProfileFileName               types.String `tfsdk:"any_connect_eap_profile_file_name"`
	AnyConnectEapProfileFileNameVariable       types.String `tfsdk:"any_connect_eap_profile_file_name_variable"`
	Ipv4PoolSize                               types.Int64  `tfsdk:"ipv4_pool_size"`
	Ipv4PoolSizeVariable                       types.String `tfsdk:"ipv4_pool_size_variable"`
	Ipv6PoolSize                               types.Int64  `tfsdk:"ipv6_pool_size"`
	Ipv6PoolSizeVariable                       types.String `tfsdk:"ipv6_pool_size_variable"`
	EnableCrlCheck                             types.Bool   `tfsdk:"enable_crl_check"`
	EnableCrlCheckVariable                     types.String `tfsdk:"enable_crl_check_variable"`
	PskAuthenticationType                      types.String `tfsdk:"psk_authentication_type"`
	PskAuthenticationTypeVariable              types.String `tfsdk:"psk_authentication_type_variable"`
	PskAuthenticationPreSharedKey              types.String `tfsdk:"psk_authentication_pre_shared_key"`
	PskAuthenticationPreSharedKeyVariable      types.String `tfsdk:"psk_authentication_pre_shared_key_variable"`
	RadiusGroupName                            types.String `tfsdk:"radius_group_name"`
	RadiusGroupNameVariable                    types.String `tfsdk:"radius_group_name_variable"`
	AaaSpecifyNamePolicyName                   types.String `tfsdk:"aaa_specify_name_policy_name"`
	AaaSpecifyNamePolicyNameVariable           types.String `tfsdk:"aaa_specify_name_policy_name_variable"`
	AaaSpecifyNamePolicyPassword               types.String `tfsdk:"aaa_specify_name_policy_password"`
	AaaSpecifyNamePolicyPasswordVariable       types.String `tfsdk:"aaa_specify_name_policy_password_variable"`
	AaaDeriveNameIdentity                      types.String `tfsdk:"aaa_derive_name_identity"`
	AaaDeriveNameIdentityVariable              types.String `tfsdk:"aaa_derive_name_identity_variable"`
	AaaDeriveNameDomain                        types.String `tfsdk:"aaa_derive_name_domain"`
	AaaDeriveNameDomainVariable                types.String `tfsdk:"aaa_derive_name_domain_variable"`
	AaaEnableAccounting                        types.Bool   `tfsdk:"aaa_enable_accounting"`
	AaaEnableAccountingVariable                types.String `tfsdk:"aaa_enable_accounting_variable"`
	Ikev2LocalIkeIdentityType                  types.String `tfsdk:"ikev2_local_ike_identity_type"`
	Ikev2LocalIkeIdentityTypeVariable          types.String `tfsdk:"ikev2_local_ike_identity_type_variable"`
	Ikev2LocalIkeIdentityValue                 types.String `tfsdk:"ikev2_local_ike_identity_value"`
	Ikev2LocalIkeIdentityValueVariable         types.String `tfsdk:"ikev2_local_ike_identity_value_variable"`
	Ikev2SecurityAssociationLifetime           types.Int64  `tfsdk:"ikev2_security_association_lifetime"`
	Ikev2SecurityAssociationLifetimeVariable   types.String `tfsdk:"ikev2_security_association_lifetime_variable"`
	Ikev2AntiDosThreshold                      types.Int64  `tfsdk:"ikev2_anti_dos_threshold"`
	Ikev2AntiDosThresholdVariable              types.String `tfsdk:"ikev2_anti_dos_threshold_variable"`
	IpsecEnableAntiReplay                      types.Bool   `tfsdk:"ipsec_enable_anti_replay"`
	IpsecEnableAntiReplayVariable              types.String `tfsdk:"ipsec_enable_anti_replay_variable"`
	IpsecAntiReplayWindowSize                  types.Int64  `tfsdk:"ipsec_anti_replay_window_size"`
	IpsecAntiReplayWindowSizeVariable          types.String `tfsdk:"ipsec_anti_replay_window_size_variable"`
	IpsecSecurityAssociationLifetime           types.Int64  `tfsdk:"ipsec_security_association_lifetime"`
	IpsecSecurityAssociationLifetimeVariable   types.String `tfsdk:"ipsec_security_association_lifetime_variable"`
	IpsecEnablePerfectFowardSecrecy            types.Bool   `tfsdk:"ipsec_enable_perfect_foward_secrecy"`
	IpsecEnablePerfectFowardSecrecyVariable    types.String `tfsdk:"ipsec_enable_perfect_foward_secrecy_variable"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getModel
func (data SystemRemoteAccess) getModel() string {
	return "system_remote_access"
}

// End of section. //template:end getModel

// Section below is generated&owned by "gen/generator.go". //template:begin getPath
func (data SystemRemoteAccess) getPath() string {
	return fmt.Sprintf("/v1/feature-profile/sdwan/system/%v/remote-access", url.QueryEscape(data.FeatureProfileId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody
func (data SystemRemoteAccess) toBody(ctx context.Context) string {
	body := ""
	body, _ = sjson.Set(body, "name", data.Name.ValueString())
	body, _ = sjson.Set(body, "description", data.Description.ValueString())
	path := "data."
	if data.ConnectionTypeSsl.IsNull() {
		body, _ = sjson.Set(body, path+"enableSSLVPN.optionType", "default")
		body, _ = sjson.Set(body, path+"enableSSLVPN.value", false)
	} else {
		body, _ = sjson.Set(body, path+"enableSSLVPN.optionType", "global")
		body, _ = sjson.Set(body, path+"enableSSLVPN.value", data.ConnectionTypeSsl.ValueBool())
	}
	if !data.AnyConnectEapAuthenticationType.IsNull() {
		body, _ = sjson.Set(body, path+"anyConnectEapAuth.optionType", "global")
		body, _ = sjson.Set(body, path+"anyConnectEapAuth.value", data.AnyConnectEapAuthenticationType.ValueString())
	}

	if !data.AnyConnectEapProfileDownloadStatusVariable.IsNull() {
		body, _ = sjson.Set(body, path+"anyConnectProfileDownloadStatus.optionType", "variable")
		body, _ = sjson.Set(body, path+"anyConnectProfileDownloadStatus.value", data.AnyConnectEapProfileDownloadStatusVariable.ValueString())
	} else if data.AnyConnectEapProfileDownloadStatus.IsNull() {
		body, _ = sjson.Set(body, path+"anyConnectProfileDownloadStatus.optionType", "default")
		body, _ = sjson.Set(body, path+"anyConnectProfileDownloadStatus.value", "NONE")
	} else {
		body, _ = sjson.Set(body, path+"anyConnectProfileDownloadStatus.optionType", "global")
		body, _ = sjson.Set(body, path+"anyConnectProfileDownloadStatus.value", data.AnyConnectEapProfileDownloadStatus.ValueString())
	}

	if !data.AnyConnectEapProfileFileNameVariable.IsNull() {
		body, _ = sjson.Set(body, path+"anyConnectProfileFileName.optionType", "variable")
		body, _ = sjson.Set(body, path+"anyConnectProfileFileName.value", data.AnyConnectEapProfileFileNameVariable.ValueString())
	} else if data.AnyConnectEapProfileFileName.IsNull() {
		body, _ = sjson.Set(body, path+"anyConnectProfileFileName.optionType", "default")
		body, _ = sjson.Set(body, path+"anyConnectProfileFileName.value", "")
	} else {
		body, _ = sjson.Set(body, path+"anyConnectProfileFileName.optionType", "global")
		body, _ = sjson.Set(body, path+"anyConnectProfileFileName.value", data.AnyConnectEapProfileFileName.ValueString())
	}

	if !data.Ipv4PoolSizeVariable.IsNull() {
		body, _ = sjson.Set(body, path+"ipv4PoolSize.optionType", "variable")
		body, _ = sjson.Set(body, path+"ipv4PoolSize.value", data.Ipv4PoolSizeVariable.ValueString())
	} else if data.Ipv4PoolSize.IsNull() {
		body, _ = sjson.Set(body, path+"ipv4PoolSize.optionType", "default")
		body, _ = sjson.Set(body, path+"ipv4PoolSize.value", 1000)
	} else {
		body, _ = sjson.Set(body, path+"ipv4PoolSize.optionType", "global")
		body, _ = sjson.Set(body, path+"ipv4PoolSize.value", data.Ipv4PoolSize.ValueInt64())
	}

	if !data.Ipv6PoolSizeVariable.IsNull() {
		body, _ = sjson.Set(body, path+"ipv6PoolSize.optionType", "variable")
		body, _ = sjson.Set(body, path+"ipv6PoolSize.value", data.Ipv6PoolSizeVariable.ValueString())
	} else if data.Ipv6PoolSize.IsNull() {
		body, _ = sjson.Set(body, path+"ipv6PoolSize.optionType", "default")
		body, _ = sjson.Set(body, path+"ipv6PoolSize.value", 1024)
	} else {
		body, _ = sjson.Set(body, path+"ipv6PoolSize.optionType", "global")
		body, _ = sjson.Set(body, path+"ipv6PoolSize.value", data.Ipv6PoolSize.ValueInt64())
	}

	if !data.EnableCrlCheckVariable.IsNull() {
		body, _ = sjson.Set(body, path+"enableCrlCheck.optionType", "variable")
		body, _ = sjson.Set(body, path+"enableCrlCheck.value", data.EnableCrlCheckVariable.ValueString())
	} else if data.EnableCrlCheck.IsNull() {
		body, _ = sjson.Set(body, path+"enableCrlCheck.optionType", "default")
		body, _ = sjson.Set(body, path+"enableCrlCheck.value", false)
	} else {
		body, _ = sjson.Set(body, path+"enableCrlCheck.optionType", "global")
		body, _ = sjson.Set(body, path+"enableCrlCheck.value", data.EnableCrlCheck.ValueBool())
	}

	if !data.PskAuthenticationTypeVariable.IsNull() {
		body, _ = sjson.Set(body, path+"pskAuth.pskSelection.optionType", "variable")
		body, _ = sjson.Set(body, path+"pskAuth.pskSelection.value", data.PskAuthenticationTypeVariable.ValueString())
	} else if data.PskAuthenticationType.IsNull() {
		body, _ = sjson.Set(body, path+"pskAuth.pskSelection.optionType", "default")

	} else {
		body, _ = sjson.Set(body, path+"pskAuth.pskSelection.optionType", "global")
		body, _ = sjson.Set(body, path+"pskAuth.pskSelection.value", data.PskAuthenticationType.ValueString())
	}

	if !data.PskAuthenticationPreSharedKeyVariable.IsNull() {
		body, _ = sjson.Set(body, path+"pskAuth.preSharedKey.optionType", "variable")
		body, _ = sjson.Set(body, path+"pskAuth.preSharedKey.value", data.PskAuthenticationPreSharedKeyVariable.ValueString())
	} else if !data.PskAuthenticationPreSharedKey.IsNull() {
		body, _ = sjson.Set(body, path+"pskAuth.preSharedKey.optionType", "global")
		body, _ = sjson.Set(body, path+"pskAuth.preSharedKey.value", data.PskAuthenticationPreSharedKey.ValueString())
	}

	if !data.RadiusGroupNameVariable.IsNull() {
		body, _ = sjson.Set(body, path+"radiusGroupName.optionType", "variable")
		body, _ = sjson.Set(body, path+"radiusGroupName.value", data.RadiusGroupNameVariable.ValueString())
	} else if !data.RadiusGroupName.IsNull() {
		body, _ = sjson.Set(body, path+"radiusGroupName.optionType", "global")
		body, _ = sjson.Set(body, path+"radiusGroupName.value", data.RadiusGroupName.ValueString())
	}

	if !data.AaaSpecifyNamePolicyNameVariable.IsNull() {
		body, _ = sjson.Set(body, path+"aaaPolicy.specifyName.policyName.optionType", "variable")
		body, _ = sjson.Set(body, path+"aaaPolicy.specifyName.policyName.value", data.AaaSpecifyNamePolicyNameVariable.ValueString())
	} else if !data.AaaSpecifyNamePolicyName.IsNull() {
		body, _ = sjson.Set(body, path+"aaaPolicy.specifyName.policyName.optionType", "global")
		body, _ = sjson.Set(body, path+"aaaPolicy.specifyName.policyName.value", data.AaaSpecifyNamePolicyName.ValueString())
	}

	if !data.AaaSpecifyNamePolicyPasswordVariable.IsNull() {
		body, _ = sjson.Set(body, path+"aaaPolicy.specifyName.password.optionType", "variable")
		body, _ = sjson.Set(body, path+"aaaPolicy.specifyName.password.value", data.AaaSpecifyNamePolicyPasswordVariable.ValueString())
	} else if !data.AaaSpecifyNamePolicyPassword.IsNull() {
		body, _ = sjson.Set(body, path+"aaaPolicy.specifyName.password.optionType", "global")
		body, _ = sjson.Set(body, path+"aaaPolicy.specifyName.password.value", data.AaaSpecifyNamePolicyPassword.ValueString())
	}

	if !data.AaaDeriveNameIdentityVariable.IsNull() {
		body, _ = sjson.Set(body, path+"aaaPolicy.deriveNameIdentity.optionType", "variable")
		body, _ = sjson.Set(body, path+"aaaPolicy.deriveNameIdentity.value", data.AaaDeriveNameIdentityVariable.ValueString())
	} else if !data.AaaDeriveNameIdentity.IsNull() {
		body, _ = sjson.Set(body, path+"aaaPolicy.deriveNameIdentity.optionType", "global")
		body, _ = sjson.Set(body, path+"aaaPolicy.deriveNameIdentity.value", data.AaaDeriveNameIdentity.ValueString())
	}

	if !data.AaaDeriveNameDomainVariable.IsNull() {
		body, _ = sjson.Set(body, path+"aaaPolicy.deriveNameDomain.optionType", "variable")
		body, _ = sjson.Set(body, path+"aaaPolicy.deriveNameDomain.value", data.AaaDeriveNameDomainVariable.ValueString())
	} else if !data.AaaDeriveNameDomain.IsNull() {
		body, _ = sjson.Set(body, path+"aaaPolicy.deriveNameDomain.optionType", "global")
		body, _ = sjson.Set(body, path+"aaaPolicy.deriveNameDomain.value", data.AaaDeriveNameDomain.ValueString())
	}

	if !data.AaaEnableAccountingVariable.IsNull() {
		body, _ = sjson.Set(body, path+"enableAccounting.optionType", "variable")
		body, _ = sjson.Set(body, path+"enableAccounting.value", data.AaaEnableAccountingVariable.ValueString())
	} else if data.AaaEnableAccounting.IsNull() {
		body, _ = sjson.Set(body, path+"enableAccounting.optionType", "default")
		body, _ = sjson.Set(body, path+"enableAccounting.value", true)
	} else {
		body, _ = sjson.Set(body, path+"enableAccounting.optionType", "global")
		body, _ = sjson.Set(body, path+"enableAccounting.value", data.AaaEnableAccounting.ValueBool())
	}

	if !data.Ikev2LocalIkeIdentityTypeVariable.IsNull() {
		body, _ = sjson.Set(body, path+"localIkev2IdentityType.optionType", "variable")
		body, _ = sjson.Set(body, path+"localIkev2IdentityType.value", data.Ikev2LocalIkeIdentityTypeVariable.ValueString())
	} else if data.Ikev2LocalIkeIdentityType.IsNull() {
		body, _ = sjson.Set(body, path+"localIkev2IdentityType.optionType", "default")

	} else {
		body, _ = sjson.Set(body, path+"localIkev2IdentityType.optionType", "global")
		body, _ = sjson.Set(body, path+"localIkev2IdentityType.value", data.Ikev2LocalIkeIdentityType.ValueString())
	}

	if !data.Ikev2LocalIkeIdentityValueVariable.IsNull() {
		body, _ = sjson.Set(body, path+"localIkev2IdentityValue.optionType", "variable")
		body, _ = sjson.Set(body, path+"localIkev2IdentityValue.value", data.Ikev2LocalIkeIdentityValueVariable.ValueString())
	} else if data.Ikev2LocalIkeIdentityValue.IsNull() {
		body, _ = sjson.Set(body, path+"localIkev2IdentityValue.optionType", "default")

	} else {
		body, _ = sjson.Set(body, path+"localIkev2IdentityValue.optionType", "global")
		body, _ = sjson.Set(body, path+"localIkev2IdentityValue.value", data.Ikev2LocalIkeIdentityValue.ValueString())
	}

	if !data.Ikev2SecurityAssociationLifetimeVariable.IsNull() {
		body, _ = sjson.Set(body, path+"ikev2SaLifetime.optionType", "variable")
		body, _ = sjson.Set(body, path+"ikev2SaLifetime.value", data.Ikev2SecurityAssociationLifetimeVariable.ValueString())
	} else if data.Ikev2SecurityAssociationLifetime.IsNull() {
		body, _ = sjson.Set(body, path+"ikev2SaLifetime.optionType", "default")
		body, _ = sjson.Set(body, path+"ikev2SaLifetime.value", 86400)
	} else {
		body, _ = sjson.Set(body, path+"ikev2SaLifetime.optionType", "global")
		body, _ = sjson.Set(body, path+"ikev2SaLifetime.value", data.Ikev2SecurityAssociationLifetime.ValueInt64())
	}

	if !data.Ikev2AntiDosThresholdVariable.IsNull() {
		body, _ = sjson.Set(body, path+"antiDosThreshold.optionType", "variable")
		body, _ = sjson.Set(body, path+"antiDosThreshold.value", data.Ikev2AntiDosThresholdVariable.ValueString())
	} else if data.Ikev2AntiDosThreshold.IsNull() {
		body, _ = sjson.Set(body, path+"antiDosThreshold.optionType", "default")
		body, _ = sjson.Set(body, path+"antiDosThreshold.value", 100)
	} else {
		body, _ = sjson.Set(body, path+"antiDosThreshold.optionType", "global")
		body, _ = sjson.Set(body, path+"antiDosThreshold.value", data.Ikev2AntiDosThreshold.ValueInt64())
	}

	if !data.IpsecEnableAntiReplayVariable.IsNull() {
		body, _ = sjson.Set(body, path+"enableAntiReplay.optionType", "variable")
		body, _ = sjson.Set(body, path+"enableAntiReplay.value", data.IpsecEnableAntiReplayVariable.ValueString())
	} else if data.IpsecEnableAntiReplay.IsNull() {
		body, _ = sjson.Set(body, path+"enableAntiReplay.optionType", "default")
		body, _ = sjson.Set(body, path+"enableAntiReplay.value", true)
	} else {
		body, _ = sjson.Set(body, path+"enableAntiReplay.optionType", "global")
		body, _ = sjson.Set(body, path+"enableAntiReplay.value", data.IpsecEnableAntiReplay.ValueBool())
	}

	if !data.IpsecAntiReplayWindowSizeVariable.IsNull() {
		body, _ = sjson.Set(body, path+"antiReplayWindowSize.optionType", "variable")
		body, _ = sjson.Set(body, path+"antiReplayWindowSize.value", data.IpsecAntiReplayWindowSizeVariable.ValueString())
	} else if data.IpsecAntiReplayWindowSize.IsNull() {
		body, _ = sjson.Set(body, path+"antiReplayWindowSize.optionType", "default")
		body, _ = sjson.Set(body, path+"antiReplayWindowSize.value", 64)
	} else {
		body, _ = sjson.Set(body, path+"antiReplayWindowSize.optionType", "global")
		body, _ = sjson.Set(body, path+"antiReplayWindowSize.value", data.IpsecAntiReplayWindowSize.ValueInt64())
	}

	if !data.IpsecSecurityAssociationLifetimeVariable.IsNull() {
		body, _ = sjson.Set(body, path+"ipSecSaLifetime.optionType", "variable")
		body, _ = sjson.Set(body, path+"ipSecSaLifetime.value", data.IpsecSecurityAssociationLifetimeVariable.ValueString())
	} else if data.IpsecSecurityAssociationLifetime.IsNull() {
		body, _ = sjson.Set(body, path+"ipSecSaLifetime.optionType", "default")
		body, _ = sjson.Set(body, path+"ipSecSaLifetime.value", 3600)
	} else {
		body, _ = sjson.Set(body, path+"ipSecSaLifetime.optionType", "global")
		body, _ = sjson.Set(body, path+"ipSecSaLifetime.value", data.IpsecSecurityAssociationLifetime.ValueInt64())
	}

	if !data.IpsecEnablePerfectFowardSecrecyVariable.IsNull() {
		body, _ = sjson.Set(body, path+"enablePFS.optionType", "variable")
		body, _ = sjson.Set(body, path+"enablePFS.value", data.IpsecEnablePerfectFowardSecrecyVariable.ValueString())
	} else if data.IpsecEnablePerfectFowardSecrecy.IsNull() {
		body, _ = sjson.Set(body, path+"enablePFS.optionType", "default")
		body, _ = sjson.Set(body, path+"enablePFS.value", false)
	} else {
		body, _ = sjson.Set(body, path+"enablePFS.optionType", "global")
		body, _ = sjson.Set(body, path+"enablePFS.value", data.IpsecEnablePerfectFowardSecrecy.ValueBool())
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody
func (data *SystemRemoteAccess) fromBody(ctx context.Context, res gjson.Result) {
	data.Name = types.StringValue(res.Get("payload.name").String())
	if value := res.Get("payload.description"); value.Exists() && value.String() != "" {
		data.Description = types.StringValue(value.String())
	} else {
		data.Description = types.StringNull()
	}
	path := "payload.data."
	data.ConnectionTypeSsl = types.BoolNull()

	if t := res.Get(path + "enableSSLVPN.optionType"); t.Exists() {
		va := res.Get(path + "enableSSLVPN.value")
		if t.String() == "global" {
			data.ConnectionTypeSsl = types.BoolValue(va.Bool())
		}
	}
	data.AnyConnectEapAuthenticationType = types.StringNull()

	if t := res.Get(path + "anyConnectEapAuth.optionType"); t.Exists() {
		va := res.Get(path + "anyConnectEapAuth.value")
		if t.String() == "global" {
			data.AnyConnectEapAuthenticationType = types.StringValue(va.String())
		}
	}
	data.AnyConnectEapProfileDownloadStatus = types.StringNull()
	data.AnyConnectEapProfileDownloadStatusVariable = types.StringNull()
	if t := res.Get(path + "anyConnectProfileDownloadStatus.optionType"); t.Exists() {
		va := res.Get(path + "anyConnectProfileDownloadStatus.value")
		if t.String() == "variable" {
			data.AnyConnectEapProfileDownloadStatusVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.AnyConnectEapProfileDownloadStatus = types.StringValue(va.String())
		}
	}
	data.AnyConnectEapProfileFileName = types.StringNull()
	data.AnyConnectEapProfileFileNameVariable = types.StringNull()
	if t := res.Get(path + "anyConnectProfileFileName.optionType"); t.Exists() {
		va := res.Get(path + "anyConnectProfileFileName.value")
		if t.String() == "variable" {
			data.AnyConnectEapProfileFileNameVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.AnyConnectEapProfileFileName = types.StringValue(va.String())
		}
	}
	data.Ipv4PoolSize = types.Int64Null()
	data.Ipv4PoolSizeVariable = types.StringNull()
	if t := res.Get(path + "ipv4PoolSize.optionType"); t.Exists() {
		va := res.Get(path + "ipv4PoolSize.value")
		if t.String() == "variable" {
			data.Ipv4PoolSizeVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.Ipv4PoolSize = types.Int64Value(va.Int())
		}
	}
	data.Ipv6PoolSize = types.Int64Null()
	data.Ipv6PoolSizeVariable = types.StringNull()
	if t := res.Get(path + "ipv6PoolSize.optionType"); t.Exists() {
		va := res.Get(path + "ipv6PoolSize.value")
		if t.String() == "variable" {
			data.Ipv6PoolSizeVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.Ipv6PoolSize = types.Int64Value(va.Int())
		}
	}
	data.EnableCrlCheck = types.BoolNull()
	data.EnableCrlCheckVariable = types.StringNull()
	if t := res.Get(path + "enableCrlCheck.optionType"); t.Exists() {
		va := res.Get(path + "enableCrlCheck.value")
		if t.String() == "variable" {
			data.EnableCrlCheckVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.EnableCrlCheck = types.BoolValue(va.Bool())
		}
	}
	data.PskAuthenticationType = types.StringNull()
	data.PskAuthenticationTypeVariable = types.StringNull()
	if t := res.Get(path + "pskAuth.pskSelection.optionType"); t.Exists() {
		va := res.Get(path + "pskAuth.pskSelection.value")
		if t.String() == "variable" {
			data.PskAuthenticationTypeVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.PskAuthenticationType = types.StringValue(va.String())
		}
	}
	data.PskAuthenticationPreSharedKey = types.StringNull()
	data.PskAuthenticationPreSharedKeyVariable = types.StringNull()
	if t := res.Get(path + "pskAuth.preSharedKey.optionType"); t.Exists() {
		va := res.Get(path + "pskAuth.preSharedKey.value")
		if t.String() == "variable" {
			data.PskAuthenticationPreSharedKeyVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.PskAuthenticationPreSharedKey = types.StringValue(va.String())
		}
	}
	data.RadiusGroupName = types.StringNull()
	data.RadiusGroupNameVariable = types.StringNull()
	if t := res.Get(path + "radiusGroupName.optionType"); t.Exists() {
		va := res.Get(path + "radiusGroupName.value")
		if t.String() == "variable" {
			data.RadiusGroupNameVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.RadiusGroupName = types.StringValue(va.String())
		}
	}
	data.AaaSpecifyNamePolicyName = types.StringNull()
	data.AaaSpecifyNamePolicyNameVariable = types.StringNull()
	if t := res.Get(path + "aaaPolicy.specifyName.policyName.optionType"); t.Exists() {
		va := res.Get(path + "aaaPolicy.specifyName.policyName.value")
		if t.String() == "variable" {
			data.AaaSpecifyNamePolicyNameVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.AaaSpecifyNamePolicyName = types.StringValue(va.String())
		}
	}
	data.AaaSpecifyNamePolicyPassword = types.StringNull()
	data.AaaSpecifyNamePolicyPasswordVariable = types.StringNull()
	if t := res.Get(path + "aaaPolicy.specifyName.password.optionType"); t.Exists() {
		va := res.Get(path + "aaaPolicy.specifyName.password.value")
		if t.String() == "variable" {
			data.AaaSpecifyNamePolicyPasswordVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.AaaSpecifyNamePolicyPassword = types.StringValue(va.String())
		}
	}
	data.AaaDeriveNameIdentity = types.StringNull()
	data.AaaDeriveNameIdentityVariable = types.StringNull()
	if t := res.Get(path + "aaaPolicy.deriveNameIdentity.optionType"); t.Exists() {
		va := res.Get(path + "aaaPolicy.deriveNameIdentity.value")
		if t.String() == "variable" {
			data.AaaDeriveNameIdentityVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.AaaDeriveNameIdentity = types.StringValue(va.String())
		}
	}
	data.AaaDeriveNameDomain = types.StringNull()
	data.AaaDeriveNameDomainVariable = types.StringNull()
	if t := res.Get(path + "aaaPolicy.deriveNameDomain.optionType"); t.Exists() {
		va := res.Get(path + "aaaPolicy.deriveNameDomain.value")
		if t.String() == "variable" {
			data.AaaDeriveNameDomainVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.AaaDeriveNameDomain = types.StringValue(va.String())
		}
	}
	data.AaaEnableAccounting = types.BoolNull()
	data.AaaEnableAccountingVariable = types.StringNull()
	if t := res.Get(path + "enableAccounting.optionType"); t.Exists() {
		va := res.Get(path + "enableAccounting.value")
		if t.String() == "variable" {
			data.AaaEnableAccountingVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.AaaEnableAccounting = types.BoolValue(va.Bool())
		}
	}
	data.Ikev2LocalIkeIdentityType = types.StringNull()
	data.Ikev2LocalIkeIdentityTypeVariable = types.StringNull()
	if t := res.Get(path + "localIkev2IdentityType.optionType"); t.Exists() {
		va := res.Get(path + "localIkev2IdentityType.value")
		if t.String() == "variable" {
			data.Ikev2LocalIkeIdentityTypeVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.Ikev2LocalIkeIdentityType = types.StringValue(va.String())
		}
	}
	data.Ikev2LocalIkeIdentityValue = types.StringNull()
	data.Ikev2LocalIkeIdentityValueVariable = types.StringNull()
	if t := res.Get(path + "localIkev2IdentityValue.optionType"); t.Exists() {
		va := res.Get(path + "localIkev2IdentityValue.value")
		if t.String() == "variable" {
			data.Ikev2LocalIkeIdentityValueVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.Ikev2LocalIkeIdentityValue = types.StringValue(va.String())
		}
	}
	data.Ikev2SecurityAssociationLifetime = types.Int64Null()
	data.Ikev2SecurityAssociationLifetimeVariable = types.StringNull()
	if t := res.Get(path + "ikev2SaLifetime.optionType"); t.Exists() {
		va := res.Get(path + "ikev2SaLifetime.value")
		if t.String() == "variable" {
			data.Ikev2SecurityAssociationLifetimeVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.Ikev2SecurityAssociationLifetime = types.Int64Value(va.Int())
		}
	}
	data.Ikev2AntiDosThreshold = types.Int64Null()
	data.Ikev2AntiDosThresholdVariable = types.StringNull()
	if t := res.Get(path + "antiDosThreshold.optionType"); t.Exists() {
		va := res.Get(path + "antiDosThreshold.value")
		if t.String() == "variable" {
			data.Ikev2AntiDosThresholdVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.Ikev2AntiDosThreshold = types.Int64Value(va.Int())
		}
	}
	data.IpsecEnableAntiReplay = types.BoolNull()
	data.IpsecEnableAntiReplayVariable = types.StringNull()
	if t := res.Get(path + "enableAntiReplay.optionType"); t.Exists() {
		va := res.Get(path + "enableAntiReplay.value")
		if t.String() == "variable" {
			data.IpsecEnableAntiReplayVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.IpsecEnableAntiReplay = types.BoolValue(va.Bool())
		}
	}
	data.IpsecAntiReplayWindowSize = types.Int64Null()
	data.IpsecAntiReplayWindowSizeVariable = types.StringNull()
	if t := res.Get(path + "antiReplayWindowSize.optionType"); t.Exists() {
		va := res.Get(path + "antiReplayWindowSize.value")
		if t.String() == "variable" {
			data.IpsecAntiReplayWindowSizeVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.IpsecAntiReplayWindowSize = types.Int64Value(va.Int())
		}
	}
	data.IpsecSecurityAssociationLifetime = types.Int64Null()
	data.IpsecSecurityAssociationLifetimeVariable = types.StringNull()
	if t := res.Get(path + "ipSecSaLifetime.optionType"); t.Exists() {
		va := res.Get(path + "ipSecSaLifetime.value")
		if t.String() == "variable" {
			data.IpsecSecurityAssociationLifetimeVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.IpsecSecurityAssociationLifetime = types.Int64Value(va.Int())
		}
	}
	data.IpsecEnablePerfectFowardSecrecy = types.BoolNull()
	data.IpsecEnablePerfectFowardSecrecyVariable = types.StringNull()
	if t := res.Get(path + "enablePFS.optionType"); t.Exists() {
		va := res.Get(path + "enablePFS.value")
		if t.String() == "variable" {
			data.IpsecEnablePerfectFowardSecrecyVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.IpsecEnablePerfectFowardSecrecy = types.BoolValue(va.Bool())
		}
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin updateFromBody
func (data *SystemRemoteAccess) updateFromBody(ctx context.Context, res gjson.Result) {
	data.Name = types.StringValue(res.Get("payload.name").String())
	if value := res.Get("payload.description"); value.Exists() && value.String() != "" {
		data.Description = types.StringValue(value.String())
	} else {
		data.Description = types.StringNull()
	}
	path := "payload.data."
	data.ConnectionTypeSsl = types.BoolNull()

	if t := res.Get(path + "enableSSLVPN.optionType"); t.Exists() {
		va := res.Get(path + "enableSSLVPN.value")
		if t.String() == "global" {
			data.ConnectionTypeSsl = types.BoolValue(va.Bool())
		}
	}
	data.AnyConnectEapAuthenticationType = types.StringNull()

	if t := res.Get(path + "anyConnectEapAuth.optionType"); t.Exists() {
		va := res.Get(path + "anyConnectEapAuth.value")
		if t.String() == "global" {
			data.AnyConnectEapAuthenticationType = types.StringValue(va.String())
		}
	}
	data.AnyConnectEapProfileDownloadStatus = types.StringNull()
	data.AnyConnectEapProfileDownloadStatusVariable = types.StringNull()
	if t := res.Get(path + "anyConnectProfileDownloadStatus.optionType"); t.Exists() {
		va := res.Get(path + "anyConnectProfileDownloadStatus.value")
		if t.String() == "variable" {
			data.AnyConnectEapProfileDownloadStatusVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.AnyConnectEapProfileDownloadStatus = types.StringValue(va.String())
		}
	}
	data.AnyConnectEapProfileFileName = types.StringNull()
	data.AnyConnectEapProfileFileNameVariable = types.StringNull()
	if t := res.Get(path + "anyConnectProfileFileName.optionType"); t.Exists() {
		va := res.Get(path + "anyConnectProfileFileName.value")
		if t.String() == "variable" {
			data.AnyConnectEapProfileFileNameVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.AnyConnectEapProfileFileName = types.StringValue(va.String())
		}
	}
	data.Ipv4PoolSize = types.Int64Null()
	data.Ipv4PoolSizeVariable = types.StringNull()
	if t := res.Get(path + "ipv4PoolSize.optionType"); t.Exists() {
		va := res.Get(path + "ipv4PoolSize.value")
		if t.String() == "variable" {
			data.Ipv4PoolSizeVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.Ipv4PoolSize = types.Int64Value(va.Int())
		}
	}
	data.Ipv6PoolSize = types.Int64Null()
	data.Ipv6PoolSizeVariable = types.StringNull()
	if t := res.Get(path + "ipv6PoolSize.optionType"); t.Exists() {
		va := res.Get(path + "ipv6PoolSize.value")
		if t.String() == "variable" {
			data.Ipv6PoolSizeVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.Ipv6PoolSize = types.Int64Value(va.Int())
		}
	}
	data.EnableCrlCheck = types.BoolNull()
	data.EnableCrlCheckVariable = types.StringNull()
	if t := res.Get(path + "enableCrlCheck.optionType"); t.Exists() {
		va := res.Get(path + "enableCrlCheck.value")
		if t.String() == "variable" {
			data.EnableCrlCheckVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.EnableCrlCheck = types.BoolValue(va.Bool())
		}
	}
	data.PskAuthenticationType = types.StringNull()
	data.PskAuthenticationTypeVariable = types.StringNull()
	if t := res.Get(path + "pskAuth.pskSelection.optionType"); t.Exists() {
		va := res.Get(path + "pskAuth.pskSelection.value")
		if t.String() == "variable" {
			data.PskAuthenticationTypeVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.PskAuthenticationType = types.StringValue(va.String())
		}
	}
	data.PskAuthenticationPreSharedKey = types.StringNull()
	data.PskAuthenticationPreSharedKeyVariable = types.StringNull()
	if t := res.Get(path + "pskAuth.preSharedKey.optionType"); t.Exists() {
		va := res.Get(path + "pskAuth.preSharedKey.value")
		if t.String() == "variable" {
			data.PskAuthenticationPreSharedKeyVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.PskAuthenticationPreSharedKey = types.StringValue(va.String())
		}
	}
	data.RadiusGroupName = types.StringNull()
	data.RadiusGroupNameVariable = types.StringNull()
	if t := res.Get(path + "radiusGroupName.optionType"); t.Exists() {
		va := res.Get(path + "radiusGroupName.value")
		if t.String() == "variable" {
			data.RadiusGroupNameVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.RadiusGroupName = types.StringValue(va.String())
		}
	}
	data.AaaSpecifyNamePolicyName = types.StringNull()
	data.AaaSpecifyNamePolicyNameVariable = types.StringNull()
	if t := res.Get(path + "aaaPolicy.specifyName.policyName.optionType"); t.Exists() {
		va := res.Get(path + "aaaPolicy.specifyName.policyName.value")
		if t.String() == "variable" {
			data.AaaSpecifyNamePolicyNameVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.AaaSpecifyNamePolicyName = types.StringValue(va.String())
		}
	}
	data.AaaSpecifyNamePolicyPassword = types.StringNull()
	data.AaaSpecifyNamePolicyPasswordVariable = types.StringNull()
	if t := res.Get(path + "aaaPolicy.specifyName.password.optionType"); t.Exists() {
		va := res.Get(path + "aaaPolicy.specifyName.password.value")
		if t.String() == "variable" {
			data.AaaSpecifyNamePolicyPasswordVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.AaaSpecifyNamePolicyPassword = types.StringValue(va.String())
		}
	}
	data.AaaDeriveNameIdentity = types.StringNull()
	data.AaaDeriveNameIdentityVariable = types.StringNull()
	if t := res.Get(path + "aaaPolicy.deriveNameIdentity.optionType"); t.Exists() {
		va := res.Get(path + "aaaPolicy.deriveNameIdentity.value")
		if t.String() == "variable" {
			data.AaaDeriveNameIdentityVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.AaaDeriveNameIdentity = types.StringValue(va.String())
		}
	}
	data.AaaDeriveNameDomain = types.StringNull()
	data.AaaDeriveNameDomainVariable = types.StringNull()
	if t := res.Get(path + "aaaPolicy.deriveNameDomain.optionType"); t.Exists() {
		va := res.Get(path + "aaaPolicy.deriveNameDomain.value")
		if t.String() == "variable" {
			data.AaaDeriveNameDomainVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.AaaDeriveNameDomain = types.StringValue(va.String())
		}
	}
	data.AaaEnableAccounting = types.BoolNull()
	data.AaaEnableAccountingVariable = types.StringNull()
	if t := res.Get(path + "enableAccounting.optionType"); t.Exists() {
		va := res.Get(path + "enableAccounting.value")
		if t.String() == "variable" {
			data.AaaEnableAccountingVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.AaaEnableAccounting = types.BoolValue(va.Bool())
		}
	}
	data.Ikev2LocalIkeIdentityType = types.StringNull()
	data.Ikev2LocalIkeIdentityTypeVariable = types.StringNull()
	if t := res.Get(path + "localIkev2IdentityType.optionType"); t.Exists() {
		va := res.Get(path + "localIkev2IdentityType.value")
		if t.String() == "variable" {
			data.Ikev2LocalIkeIdentityTypeVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.Ikev2LocalIkeIdentityType = types.StringValue(va.String())
		}
	}
	data.Ikev2LocalIkeIdentityValue = types.StringNull()
	data.Ikev2LocalIkeIdentityValueVariable = types.StringNull()
	if t := res.Get(path + "localIkev2IdentityValue.optionType"); t.Exists() {
		va := res.Get(path + "localIkev2IdentityValue.value")
		if t.String() == "variable" {
			data.Ikev2LocalIkeIdentityValueVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.Ikev2LocalIkeIdentityValue = types.StringValue(va.String())
		}
	}
	data.Ikev2SecurityAssociationLifetime = types.Int64Null()
	data.Ikev2SecurityAssociationLifetimeVariable = types.StringNull()
	if t := res.Get(path + "ikev2SaLifetime.optionType"); t.Exists() {
		va := res.Get(path + "ikev2SaLifetime.value")
		if t.String() == "variable" {
			data.Ikev2SecurityAssociationLifetimeVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.Ikev2SecurityAssociationLifetime = types.Int64Value(va.Int())
		}
	}
	data.Ikev2AntiDosThreshold = types.Int64Null()
	data.Ikev2AntiDosThresholdVariable = types.StringNull()
	if t := res.Get(path + "antiDosThreshold.optionType"); t.Exists() {
		va := res.Get(path + "antiDosThreshold.value")
		if t.String() == "variable" {
			data.Ikev2AntiDosThresholdVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.Ikev2AntiDosThreshold = types.Int64Value(va.Int())
		}
	}
	data.IpsecEnableAntiReplay = types.BoolNull()
	data.IpsecEnableAntiReplayVariable = types.StringNull()
	if t := res.Get(path + "enableAntiReplay.optionType"); t.Exists() {
		va := res.Get(path + "enableAntiReplay.value")
		if t.String() == "variable" {
			data.IpsecEnableAntiReplayVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.IpsecEnableAntiReplay = types.BoolValue(va.Bool())
		}
	}
	data.IpsecAntiReplayWindowSize = types.Int64Null()
	data.IpsecAntiReplayWindowSizeVariable = types.StringNull()
	if t := res.Get(path + "antiReplayWindowSize.optionType"); t.Exists() {
		va := res.Get(path + "antiReplayWindowSize.value")
		if t.String() == "variable" {
			data.IpsecAntiReplayWindowSizeVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.IpsecAntiReplayWindowSize = types.Int64Value(va.Int())
		}
	}
	data.IpsecSecurityAssociationLifetime = types.Int64Null()
	data.IpsecSecurityAssociationLifetimeVariable = types.StringNull()
	if t := res.Get(path + "ipSecSaLifetime.optionType"); t.Exists() {
		va := res.Get(path + "ipSecSaLifetime.value")
		if t.String() == "variable" {
			data.IpsecSecurityAssociationLifetimeVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.IpsecSecurityAssociationLifetime = types.Int64Value(va.Int())
		}
	}
	data.IpsecEnablePerfectFowardSecrecy = types.BoolNull()
	data.IpsecEnablePerfectFowardSecrecyVariable = types.StringNull()
	if t := res.Get(path + "enablePFS.optionType"); t.Exists() {
		va := res.Get(path + "enablePFS.value")
		if t.String() == "variable" {
			data.IpsecEnablePerfectFowardSecrecyVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.IpsecEnablePerfectFowardSecrecy = types.BoolValue(va.Bool())
		}
	}
}

// End of section. //template:end updateFromBody

// Section below is generated&owned by "gen/generator.go". //template:begin isNull
func (data *SystemRemoteAccess) isNull(ctx context.Context, res gjson.Result) bool {
	if !data.FeatureProfileId.IsNull() {
		return false
	}
	if !data.ConnectionTypeSsl.IsNull() {
		return false
	}
	if !data.AnyConnectEapAuthenticationType.IsNull() {
		return false
	}
	if !data.AnyConnectEapProfileDownloadStatus.IsNull() {
		return false
	}
	if !data.AnyConnectEapProfileDownloadStatusVariable.IsNull() {
		return false
	}
	if !data.AnyConnectEapProfileFileName.IsNull() {
		return false
	}
	if !data.AnyConnectEapProfileFileNameVariable.IsNull() {
		return false
	}
	if !data.Ipv4PoolSize.IsNull() {
		return false
	}
	if !data.Ipv4PoolSizeVariable.IsNull() {
		return false
	}
	if !data.Ipv6PoolSize.IsNull() {
		return false
	}
	if !data.Ipv6PoolSizeVariable.IsNull() {
		return false
	}
	if !data.EnableCrlCheck.IsNull() {
		return false
	}
	if !data.EnableCrlCheckVariable.IsNull() {
		return false
	}
	if !data.PskAuthenticationType.IsNull() {
		return false
	}
	if !data.PskAuthenticationTypeVariable.IsNull() {
		return false
	}
	if !data.PskAuthenticationPreSharedKey.IsNull() {
		return false
	}
	if !data.PskAuthenticationPreSharedKeyVariable.IsNull() {
		return false
	}
	if !data.RadiusGroupName.IsNull() {
		return false
	}
	if !data.RadiusGroupNameVariable.IsNull() {
		return false
	}
	if !data.AaaSpecifyNamePolicyName.IsNull() {
		return false
	}
	if !data.AaaSpecifyNamePolicyNameVariable.IsNull() {
		return false
	}
	if !data.AaaSpecifyNamePolicyPassword.IsNull() {
		return false
	}
	if !data.AaaSpecifyNamePolicyPasswordVariable.IsNull() {
		return false
	}
	if !data.AaaDeriveNameIdentity.IsNull() {
		return false
	}
	if !data.AaaDeriveNameIdentityVariable.IsNull() {
		return false
	}
	if !data.AaaDeriveNameDomain.IsNull() {
		return false
	}
	if !data.AaaDeriveNameDomainVariable.IsNull() {
		return false
	}
	if !data.AaaEnableAccounting.IsNull() {
		return false
	}
	if !data.AaaEnableAccountingVariable.IsNull() {
		return false
	}
	if !data.Ikev2LocalIkeIdentityType.IsNull() {
		return false
	}
	if !data.Ikev2LocalIkeIdentityTypeVariable.IsNull() {
		return false
	}
	if !data.Ikev2LocalIkeIdentityValue.IsNull() {
		return false
	}
	if !data.Ikev2LocalIkeIdentityValueVariable.IsNull() {
		return false
	}
	if !data.Ikev2SecurityAssociationLifetime.IsNull() {
		return false
	}
	if !data.Ikev2SecurityAssociationLifetimeVariable.IsNull() {
		return false
	}
	if !data.Ikev2AntiDosThreshold.IsNull() {
		return false
	}
	if !data.Ikev2AntiDosThresholdVariable.IsNull() {
		return false
	}
	if !data.IpsecEnableAntiReplay.IsNull() {
		return false
	}
	if !data.IpsecEnableAntiReplayVariable.IsNull() {
		return false
	}
	if !data.IpsecAntiReplayWindowSize.IsNull() {
		return false
	}
	if !data.IpsecAntiReplayWindowSizeVariable.IsNull() {
		return false
	}
	if !data.IpsecSecurityAssociationLifetime.IsNull() {
		return false
	}
	if !data.IpsecSecurityAssociationLifetimeVariable.IsNull() {
		return false
	}
	if !data.IpsecEnablePerfectFowardSecrecy.IsNull() {
		return false
	}
	if !data.IpsecEnablePerfectFowardSecrecyVariable.IsNull() {
		return false
	}
	return true
}

// End of section. //template:end isNull
