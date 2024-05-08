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
type CiscoSIGCredentials struct {
	Id                             types.String `tfsdk:"id"`
	Version                        types.Int64  `tfsdk:"version"`
	TemplateType                   types.String `tfsdk:"template_type"`
	Name                           types.String `tfsdk:"name"`
	Description                    types.String `tfsdk:"description"`
	DeviceTypes                    types.Set    `tfsdk:"device_types"`
	ZscalerOrganization            types.String `tfsdk:"zscaler_organization"`
	ZscalerOrganizationVariable    types.String `tfsdk:"zscaler_organization_variable"`
	ZscalerPartnerBaseUri          types.String `tfsdk:"zscaler_partner_base_uri"`
	ZscalerPartnerBaseUriVariable  types.String `tfsdk:"zscaler_partner_base_uri_variable"`
	ZscalerUsername                types.String `tfsdk:"zscaler_username"`
	ZscalerUsernameVariable        types.String `tfsdk:"zscaler_username_variable"`
	ZscalerPassword                types.String `tfsdk:"zscaler_password"`
	ZscalerPasswordVariable        types.String `tfsdk:"zscaler_password_variable"`
	ZscalerCloudName               types.Int64  `tfsdk:"zscaler_cloud_name"`
	ZscalerCloudNameVariable       types.String `tfsdk:"zscaler_cloud_name_variable"`
	ZscalerPartnerUsername         types.String `tfsdk:"zscaler_partner_username"`
	ZscalerPartnerUsernameVariable types.String `tfsdk:"zscaler_partner_username_variable"`
	ZscalerPartnerPassword         types.String `tfsdk:"zscaler_partner_password"`
	ZscalerPartnerPasswordVariable types.String `tfsdk:"zscaler_partner_password_variable"`
	ZscalerPartnerApiKey           types.String `tfsdk:"zscaler_partner_api_key"`
	ZscalerPartnerApiKeyVariable   types.String `tfsdk:"zscaler_partner_api_key_variable"`
	UmbrellaApiKey                 types.String `tfsdk:"umbrella_api_key"`
	UmbrellaApiKeyVariable         types.String `tfsdk:"umbrella_api_key_variable"`
	UmbrellaApiSecret              types.String `tfsdk:"umbrella_api_secret"`
	UmbrellaApiSecretVariable      types.String `tfsdk:"umbrella_api_secret_variable"`
	UmbrellaOrganizationId         types.String `tfsdk:"umbrella_organization_id"`
	UmbrellaOrganizationIdVariable types.String `tfsdk:"umbrella_organization_id_variable"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getModel
func (data CiscoSIGCredentials) getModel() string {
	return "cisco_sig_credentials"
}

// End of section. //template:end getModel

// Section below is generated&owned by "gen/generator.go". //template:begin toBody
func (data CiscoSIGCredentials) toBody(ctx context.Context) string {
	body := ""

	var device_types []string
	data.DeviceTypes.ElementsAs(ctx, &device_types, false)
	body, _ = sjson.Set(body, "deviceType", device_types)
	body, _ = sjson.Set(body, "factoryDefault", false)
	body, _ = sjson.Set(body, "templateDescription", data.Description.ValueString())
	body, _ = sjson.Set(body, "templateMinVersion", "15.0.0")
	body, _ = sjson.Set(body, "templateName", data.Name.ValueString())
	body, _ = sjson.Set(body, "templateType", "cisco_sig_credentials")
	body, _ = sjson.Set(body, "templateDefinition", map[string]interface{}{})
	body, _ = sjson.Set(body, "isGlobal", true)

	path := "templateDefinition."

	if !data.ZscalerOrganizationVariable.IsNull() {
		body, _ = sjson.Set(body, path+"zscaler.organization."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"zscaler.organization."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"zscaler.organization."+"vipVariableName", data.ZscalerOrganizationVariable.ValueString())
	} else if data.ZscalerOrganization.IsNull() {
	} else {
		body, _ = sjson.Set(body, path+"zscaler.organization."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"zscaler.organization."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"zscaler.organization."+"vipValue", data.ZscalerOrganization.ValueString())
	}

	if !data.ZscalerPartnerBaseUriVariable.IsNull() {
		body, _ = sjson.Set(body, path+"zscaler.partner-base-uri."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"zscaler.partner-base-uri."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"zscaler.partner-base-uri."+"vipVariableName", data.ZscalerPartnerBaseUriVariable.ValueString())
	} else if data.ZscalerPartnerBaseUri.IsNull() {
	} else {
		body, _ = sjson.Set(body, path+"zscaler.partner-base-uri."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"zscaler.partner-base-uri."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"zscaler.partner-base-uri."+"vipValue", data.ZscalerPartnerBaseUri.ValueString())
	}

	if !data.ZscalerUsernameVariable.IsNull() {
		body, _ = sjson.Set(body, path+"zscaler.username."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"zscaler.username."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"zscaler.username."+"vipVariableName", data.ZscalerUsernameVariable.ValueString())
	} else if data.ZscalerUsername.IsNull() {
	} else {
		body, _ = sjson.Set(body, path+"zscaler.username."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"zscaler.username."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"zscaler.username."+"vipValue", data.ZscalerUsername.ValueString())
	}

	if !data.ZscalerPasswordVariable.IsNull() {
		body, _ = sjson.Set(body, path+"zscaler.password."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"zscaler.password."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"zscaler.password."+"vipVariableName", data.ZscalerPasswordVariable.ValueString())
	} else if data.ZscalerPassword.IsNull() {
	} else {
		body, _ = sjson.Set(body, path+"zscaler.password."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"zscaler.password."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"zscaler.password."+"vipValue", data.ZscalerPassword.ValueString())
	}

	if !data.ZscalerCloudNameVariable.IsNull() {
		body, _ = sjson.Set(body, path+"zscaler.cloud-gateway."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"zscaler.cloud-gateway."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"zscaler.cloud-gateway."+"vipVariableName", data.ZscalerCloudNameVariable.ValueString())
	} else if data.ZscalerCloudName.IsNull() {
	} else {
		body, _ = sjson.Set(body, path+"zscaler.cloud-gateway."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"zscaler.cloud-gateway."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"zscaler.cloud-gateway."+"vipValue", data.ZscalerCloudName.ValueInt64())
	}

	if !data.ZscalerPartnerUsernameVariable.IsNull() {
		body, _ = sjson.Set(body, path+"zscaler.partner-id."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"zscaler.partner-id."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"zscaler.partner-id."+"vipVariableName", data.ZscalerPartnerUsernameVariable.ValueString())
	} else if data.ZscalerPartnerUsername.IsNull() {
	} else {
		body, _ = sjson.Set(body, path+"zscaler.partner-id."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"zscaler.partner-id."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"zscaler.partner-id."+"vipValue", data.ZscalerPartnerUsername.ValueString())
	}

	if !data.ZscalerPartnerPasswordVariable.IsNull() {
		body, _ = sjson.Set(body, path+"zscaler.partner-secret."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"zscaler.partner-secret."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"zscaler.partner-secret."+"vipVariableName", data.ZscalerPartnerPasswordVariable.ValueString())
	} else if data.ZscalerPartnerPassword.IsNull() {
	} else {
		body, _ = sjson.Set(body, path+"zscaler.partner-secret."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"zscaler.partner-secret."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"zscaler.partner-secret."+"vipValue", data.ZscalerPartnerPassword.ValueString())
	}

	if !data.ZscalerPartnerApiKeyVariable.IsNull() {
		body, _ = sjson.Set(body, path+"zscaler.partner-key."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"zscaler.partner-key."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"zscaler.partner-key."+"vipVariableName", data.ZscalerPartnerApiKeyVariable.ValueString())
	} else if data.ZscalerPartnerApiKey.IsNull() {
	} else {
		body, _ = sjson.Set(body, path+"zscaler.partner-key."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"zscaler.partner-key."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"zscaler.partner-key."+"vipValue", data.ZscalerPartnerApiKey.ValueString())
	}

	if !data.UmbrellaApiKeyVariable.IsNull() {
		body, _ = sjson.Set(body, path+"umbrella.api-key."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"umbrella.api-key."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"umbrella.api-key."+"vipVariableName", data.UmbrellaApiKeyVariable.ValueString())
	} else if data.UmbrellaApiKey.IsNull() {
	} else {
		body, _ = sjson.Set(body, path+"umbrella.api-key."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"umbrella.api-key."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"umbrella.api-key."+"vipValue", data.UmbrellaApiKey.ValueString())
	}

	if !data.UmbrellaApiSecretVariable.IsNull() {
		body, _ = sjson.Set(body, path+"umbrella.api-secret."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"umbrella.api-secret."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"umbrella.api-secret."+"vipVariableName", data.UmbrellaApiSecretVariable.ValueString())
	} else if data.UmbrellaApiSecret.IsNull() {
	} else {
		body, _ = sjson.Set(body, path+"umbrella.api-secret."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"umbrella.api-secret."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"umbrella.api-secret."+"vipValue", data.UmbrellaApiSecret.ValueString())
	}

	if !data.UmbrellaOrganizationIdVariable.IsNull() {
		body, _ = sjson.Set(body, path+"umbrella.org-id."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"umbrella.org-id."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"umbrella.org-id."+"vipVariableName", data.UmbrellaOrganizationIdVariable.ValueString())
	} else if data.UmbrellaOrganizationId.IsNull() {
	} else {
		body, _ = sjson.Set(body, path+"umbrella.org-id."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"umbrella.org-id."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"umbrella.org-id."+"vipValue", data.UmbrellaOrganizationId.ValueString())
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody
func (data *CiscoSIGCredentials) fromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("deviceType"); value.Exists() {
		data.DeviceTypes = helpers.GetStringSet(value.Array())
	} else {
		data.DeviceTypes = types.SetNull(types.StringType)
	}
	if value := res.Get("templateDescription"); value.Exists() && value.String() != "" {
		data.Description = types.StringValue(value.String())
	} else {
		data.Description = types.StringNull()
	}
	if value := res.Get("templateName"); value.Exists() {
		data.Name = types.StringValue(value.String())
	} else {
		data.Name = types.StringNull()
	}
	if value := res.Get("templateType"); value.Exists() {
		data.TemplateType = types.StringValue(value.String())
	} else {
		data.TemplateType = types.StringNull()
	}

	path := "templateDefinition."
	if value := res.Get(path + "zscaler.organization.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.ZscalerOrganization = types.StringNull()

			v := res.Get(path + "zscaler.organization.vipVariableName")
			data.ZscalerOrganizationVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.ZscalerOrganization = types.StringNull()
			data.ZscalerOrganizationVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "zscaler.organization.vipValue")
			data.ZscalerOrganization = types.StringValue(v.String())
			data.ZscalerOrganizationVariable = types.StringNull()
		}
	} else {
		data.ZscalerOrganization = types.StringNull()
		data.ZscalerOrganizationVariable = types.StringNull()
	}
	if value := res.Get(path + "zscaler.partner-base-uri.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.ZscalerPartnerBaseUri = types.StringNull()

			v := res.Get(path + "zscaler.partner-base-uri.vipVariableName")
			data.ZscalerPartnerBaseUriVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.ZscalerPartnerBaseUri = types.StringNull()
			data.ZscalerPartnerBaseUriVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "zscaler.partner-base-uri.vipValue")
			data.ZscalerPartnerBaseUri = types.StringValue(v.String())
			data.ZscalerPartnerBaseUriVariable = types.StringNull()
		}
	} else {
		data.ZscalerPartnerBaseUri = types.StringNull()
		data.ZscalerPartnerBaseUriVariable = types.StringNull()
	}
	if value := res.Get(path + "zscaler.username.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.ZscalerUsername = types.StringNull()

			v := res.Get(path + "zscaler.username.vipVariableName")
			data.ZscalerUsernameVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.ZscalerUsername = types.StringNull()
			data.ZscalerUsernameVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "zscaler.username.vipValue")
			data.ZscalerUsername = types.StringValue(v.String())
			data.ZscalerUsernameVariable = types.StringNull()
		}
	} else {
		data.ZscalerUsername = types.StringNull()
		data.ZscalerUsernameVariable = types.StringNull()
	}
	if value := res.Get(path + "zscaler.password.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.ZscalerPassword = types.StringNull()

			v := res.Get(path + "zscaler.password.vipVariableName")
			data.ZscalerPasswordVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.ZscalerPassword = types.StringNull()
			data.ZscalerPasswordVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "zscaler.password.vipValue")
			data.ZscalerPassword = types.StringValue(v.String())
			data.ZscalerPasswordVariable = types.StringNull()
		}
	} else {
		data.ZscalerPassword = types.StringNull()
		data.ZscalerPasswordVariable = types.StringNull()
	}
	if value := res.Get(path + "zscaler.cloud-gateway.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.ZscalerCloudName = types.Int64Null()

			v := res.Get(path + "zscaler.cloud-gateway.vipVariableName")
			data.ZscalerCloudNameVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.ZscalerCloudName = types.Int64Null()
			data.ZscalerCloudNameVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "zscaler.cloud-gateway.vipValue")
			data.ZscalerCloudName = types.Int64Value(v.Int())
			data.ZscalerCloudNameVariable = types.StringNull()
		}
	} else {
		data.ZscalerCloudName = types.Int64Null()
		data.ZscalerCloudNameVariable = types.StringNull()
	}
	if value := res.Get(path + "zscaler.partner-id.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.ZscalerPartnerUsername = types.StringNull()

			v := res.Get(path + "zscaler.partner-id.vipVariableName")
			data.ZscalerPartnerUsernameVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.ZscalerPartnerUsername = types.StringNull()
			data.ZscalerPartnerUsernameVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "zscaler.partner-id.vipValue")
			data.ZscalerPartnerUsername = types.StringValue(v.String())
			data.ZscalerPartnerUsernameVariable = types.StringNull()
		}
	} else {
		data.ZscalerPartnerUsername = types.StringNull()
		data.ZscalerPartnerUsernameVariable = types.StringNull()
	}
	if value := res.Get(path + "zscaler.partner-secret.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.ZscalerPartnerPassword = types.StringNull()

			v := res.Get(path + "zscaler.partner-secret.vipVariableName")
			data.ZscalerPartnerPasswordVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.ZscalerPartnerPassword = types.StringNull()
			data.ZscalerPartnerPasswordVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "zscaler.partner-secret.vipValue")
			data.ZscalerPartnerPassword = types.StringValue(v.String())
			data.ZscalerPartnerPasswordVariable = types.StringNull()
		}
	} else {
		data.ZscalerPartnerPassword = types.StringNull()
		data.ZscalerPartnerPasswordVariable = types.StringNull()
	}
	if value := res.Get(path + "zscaler.partner-key.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.ZscalerPartnerApiKey = types.StringNull()

			v := res.Get(path + "zscaler.partner-key.vipVariableName")
			data.ZscalerPartnerApiKeyVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.ZscalerPartnerApiKey = types.StringNull()
			data.ZscalerPartnerApiKeyVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "zscaler.partner-key.vipValue")
			data.ZscalerPartnerApiKey = types.StringValue(v.String())
			data.ZscalerPartnerApiKeyVariable = types.StringNull()
		}
	} else {
		data.ZscalerPartnerApiKey = types.StringNull()
		data.ZscalerPartnerApiKeyVariable = types.StringNull()
	}
	if value := res.Get(path + "umbrella.api-key.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.UmbrellaApiKey = types.StringNull()

			v := res.Get(path + "umbrella.api-key.vipVariableName")
			data.UmbrellaApiKeyVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.UmbrellaApiKey = types.StringNull()
			data.UmbrellaApiKeyVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "umbrella.api-key.vipValue")
			data.UmbrellaApiKey = types.StringValue(v.String())
			data.UmbrellaApiKeyVariable = types.StringNull()
		}
	} else {
		data.UmbrellaApiKey = types.StringNull()
		data.UmbrellaApiKeyVariable = types.StringNull()
	}
	if value := res.Get(path + "umbrella.api-secret.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.UmbrellaApiSecret = types.StringNull()

			v := res.Get(path + "umbrella.api-secret.vipVariableName")
			data.UmbrellaApiSecretVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.UmbrellaApiSecret = types.StringNull()
			data.UmbrellaApiSecretVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "umbrella.api-secret.vipValue")
			data.UmbrellaApiSecret = types.StringValue(v.String())
			data.UmbrellaApiSecretVariable = types.StringNull()
		}
	} else {
		data.UmbrellaApiSecret = types.StringNull()
		data.UmbrellaApiSecretVariable = types.StringNull()
	}
	if value := res.Get(path + "umbrella.org-id.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.UmbrellaOrganizationId = types.StringNull()

			v := res.Get(path + "umbrella.org-id.vipVariableName")
			data.UmbrellaOrganizationIdVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.UmbrellaOrganizationId = types.StringNull()
			data.UmbrellaOrganizationIdVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "umbrella.org-id.vipValue")
			data.UmbrellaOrganizationId = types.StringValue(v.String())
			data.UmbrellaOrganizationIdVariable = types.StringNull()
		}
	} else {
		data.UmbrellaOrganizationId = types.StringNull()
		data.UmbrellaOrganizationIdVariable = types.StringNull()
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin hasChanges
func (data *CiscoSIGCredentials) hasChanges(ctx context.Context, state *CiscoSIGCredentials) bool {
	hasChanges := false
	if !data.ZscalerOrganization.Equal(state.ZscalerOrganization) {
		hasChanges = true
	}
	if !data.ZscalerPartnerBaseUri.Equal(state.ZscalerPartnerBaseUri) {
		hasChanges = true
	}
	if !data.ZscalerUsername.Equal(state.ZscalerUsername) {
		hasChanges = true
	}
	if !data.ZscalerPassword.Equal(state.ZscalerPassword) {
		hasChanges = true
	}
	if !data.ZscalerCloudName.Equal(state.ZscalerCloudName) {
		hasChanges = true
	}
	if !data.ZscalerPartnerUsername.Equal(state.ZscalerPartnerUsername) {
		hasChanges = true
	}
	if !data.ZscalerPartnerPassword.Equal(state.ZscalerPartnerPassword) {
		hasChanges = true
	}
	if !data.ZscalerPartnerApiKey.Equal(state.ZscalerPartnerApiKey) {
		hasChanges = true
	}
	if !data.UmbrellaApiKey.Equal(state.UmbrellaApiKey) {
		hasChanges = true
	}
	if !data.UmbrellaApiSecret.Equal(state.UmbrellaApiSecret) {
		hasChanges = true
	}
	if !data.UmbrellaOrganizationId.Equal(state.UmbrellaOrganizationId) {
		hasChanges = true
	}
	return hasChanges
}

// End of section. //template:end hasChanges
