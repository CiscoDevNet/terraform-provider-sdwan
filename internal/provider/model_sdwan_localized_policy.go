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

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type LocalizedPolicy struct {
	Id                         types.String                 `tfsdk:"id"`
	Version                    types.Int64                  `tfsdk:"version"`
	Name                       types.String                 `tfsdk:"name"`
	Description                types.String                 `tfsdk:"description"`
	FlowVisibilityIpv4         types.Bool                   `tfsdk:"flow_visibility_ipv4"`
	FlowVisibilityIpv6         types.Bool                   `tfsdk:"flow_visibility_ipv6"`
	ApplicationVisibilityIpv4  types.Bool                   `tfsdk:"application_visibility_ipv4"`
	ApplicationVisibilityIpv6  types.Bool                   `tfsdk:"application_visibility_ipv6"`
	CloudQos                   types.Bool                   `tfsdk:"cloud_qos"`
	CloudQosServerSide         types.Bool                   `tfsdk:"cloud_qos_service_side"`
	ImplicitAclLogging         types.Bool                   `tfsdk:"implicit_acl_logging"`
	LogFrequency               types.Int64                  `tfsdk:"log_frequency"`
	Ipv4VisibilityCacheEntries types.Int64                  `tfsdk:"ipv4_visibility_cache_entries"`
	Ipv6VisibilityCacheEntries types.Int64                  `tfsdk:"ipv6_visibility_cache_entries"`
	Definitions                []LocalizedPolicyDefinitions `tfsdk:"definitions"`
}

type LocalizedPolicyDefinitions struct {
	Id      types.String `tfsdk:"id"`
	Type    types.String `tfsdk:"type"`
	Version types.Int64  `tfsdk:"version"`
}

func (data *LocalizedPolicy) toBody(ctx context.Context) string {
	body := ""
	body, _ = sjson.Set(body, "policyName", data.Name.ValueString())
	body, _ = sjson.Set(body, "policyDescription", data.Description.ValueString())
	body, _ = sjson.Set(body, "policyType", "feature")
	path := "policyDefinition."
	body, _ = sjson.Set(body, path+"settings.flowVisibility", data.FlowVisibilityIpv4.ValueBool())
	body, _ = sjson.Set(body, path+"settings.flowVisibilityIPv6", data.FlowVisibilityIpv6.ValueBool())
	body, _ = sjson.Set(body, path+"settings.appVisibility", data.ApplicationVisibilityIpv4.ValueBool())
	body, _ = sjson.Set(body, path+"settings.appVisibilityIPv6", data.ApplicationVisibilityIpv6.ValueBool())
	body, _ = sjson.Set(body, path+"settings.cloudQos", data.CloudQos.ValueBool())
	body, _ = sjson.Set(body, path+"settings.cloudQosServiceSide", data.CloudQosServerSide.ValueBool())
	body, _ = sjson.Set(body, path+"settings.implicitAclLogging", data.ImplicitAclLogging.ValueBool())
	if !data.LogFrequency.IsNull() {
		body, _ = sjson.Set(body, path+"settings.logFrequency", data.LogFrequency.ValueInt64())
	}
	if !data.Ipv4VisibilityCacheEntries.IsNull() {
		body, _ = sjson.Set(body, path+"settings.ipVisibilityCacheEntries", data.Ipv4VisibilityCacheEntries.ValueInt64())
	}
	if !data.Ipv6VisibilityCacheEntries.IsNull() {
		body, _ = sjson.Set(body, path+"settings.ipV6VisibilityCacheEntries", data.Ipv6VisibilityCacheEntries.ValueInt64())
	}
	for _, item := range data.Definitions {
		itemBody := ""
		itemBody, _ = sjson.Set(itemBody, "definitionId", item.Id.ValueString())
		itemBody, _ = sjson.Set(itemBody, "type", item.Type.ValueString())
		body, _ = sjson.SetRaw(body, path+"assembly.-1", itemBody)
	}
	return body
}

func (data *LocalizedPolicy) fromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("policyName"); value.Exists() {
		data.Name = types.StringValue(value.String())
	} else {
		data.Name = types.StringNull()
	}
	if value := res.Get("policyDescription"); value.Exists() {
		data.Description = types.StringValue(value.String())
	} else {
		data.Description = types.StringNull()
	}
	path := "policyDefinition."
	if value := res.Get(path + "settings.flowVisibility"); value.Exists() {
		data.FlowVisibilityIpv4 = types.BoolValue(value.Bool())
	} else {
		data.FlowVisibilityIpv4 = types.BoolNull()
	}
	if value := res.Get(path + "settings.flowVisibilityIPv6"); value.Exists() {
		data.FlowVisibilityIpv6 = types.BoolValue(value.Bool())
	} else {
		data.FlowVisibilityIpv6 = types.BoolNull()
	}
	if value := res.Get(path + "settings.appVisibility"); value.Exists() {
		data.ApplicationVisibilityIpv4 = types.BoolValue(value.Bool())
	} else {
		data.ApplicationVisibilityIpv4 = types.BoolNull()
	}
	if value := res.Get(path + "settings.appVisibilityIPv6"); value.Exists() {
		data.ApplicationVisibilityIpv6 = types.BoolValue(value.Bool())
	} else {
		data.ApplicationVisibilityIpv6 = types.BoolNull()
	}
	if value := res.Get(path + "settings.cloudQos"); value.Exists() {
		data.CloudQos = types.BoolValue(value.Bool())
	} else {
		data.CloudQos = types.BoolNull()
	}
	if value := res.Get(path + "settings.cloudQosServiceSide"); value.Exists() {
		data.CloudQosServerSide = types.BoolValue(value.Bool())
	} else {
		data.CloudQosServerSide = types.BoolNull()
	}
	if value := res.Get(path + "settings.implicitAclLogging"); value.Exists() {
		data.ImplicitAclLogging = types.BoolValue(value.Bool())
	} else {
		data.ImplicitAclLogging = types.BoolNull()
	}
	if value := res.Get(path + "settings.logFrequency"); value.Exists() {
		data.LogFrequency = types.Int64Value(value.Int())
	} else {
		data.LogFrequency = types.Int64Null()
	}
	if value := res.Get(path + "settings.ipVisibilityCacheEntries"); value.Exists() {
		data.Ipv4VisibilityCacheEntries = types.Int64Value(value.Int())
	} else {
		data.Ipv4VisibilityCacheEntries = types.Int64Null()
	}
	if value := res.Get(path + "settings.ipV6VisibilityCacheEntries"); value.Exists() {
		data.Ipv6VisibilityCacheEntries = types.Int64Value(value.Int())
	} else {
		data.Ipv6VisibilityCacheEntries = types.Int64Null()
	}
	if value := res.Get(path + "assembly"); value.Exists() {
		data.Definitions = make([]LocalizedPolicyDefinitions, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := LocalizedPolicyDefinitions{}
			if cValue := v.Get("definitionId"); cValue.Exists() {
				item.Id = types.StringValue(cValue.String())
			} else {
				item.Id = types.StringNull()
			}
			if cValue := v.Get("type"); cValue.Exists() {
				item.Type = types.StringValue(cValue.String())
			} else {
				item.Type = types.StringNull()
			}
			data.Definitions = append(data.Definitions, item)
			return true
		})
	}
}

func (data *LocalizedPolicy) hasChanges(ctx context.Context, state *LocalizedPolicy) bool {
	hasChanges := false
	if !data.Name.Equal(state.Name) {
		hasChanges = true
	}
	if !data.Description.Equal(state.Description) {
		hasChanges = true
	}
	if !data.FlowVisibilityIpv4.Equal(state.FlowVisibilityIpv4) {
		hasChanges = true
	}
	if !data.FlowVisibilityIpv6.Equal(state.FlowVisibilityIpv6) {
		hasChanges = true
	}
	if !data.ApplicationVisibilityIpv4.Equal(state.ApplicationVisibilityIpv4) {
		hasChanges = true
	}
	if !data.ApplicationVisibilityIpv6.Equal(state.ApplicationVisibilityIpv6) {
		hasChanges = true
	}
	if !data.CloudQos.Equal(state.CloudQos) {
		hasChanges = true
	}
	if !data.CloudQosServerSide.Equal(state.CloudQosServerSide) {
		hasChanges = true
	}
	if !data.ImplicitAclLogging.Equal(state.ImplicitAclLogging) {
		hasChanges = true
	}
	if !data.LogFrequency.Equal(state.LogFrequency) {
		hasChanges = true
	}
	if !data.Ipv4VisibilityCacheEntries.Equal(state.Ipv4VisibilityCacheEntries) {
		hasChanges = true
	}
	if !data.Ipv6VisibilityCacheEntries.Equal(state.Ipv6VisibilityCacheEntries) {
		hasChanges = true
	}
	if len(data.Definitions) != len(state.Definitions) {
		hasChanges = true
	} else {
		for d := range data.Definitions {
			if !data.Definitions[d].Id.Equal(state.Definitions[d].Id) {
				hasChanges = true
			}
			if !data.Definitions[d].Type.Equal(state.Definitions[d].Type) {
				hasChanges = true
			}
		}
	}
	return hasChanges
}

func (data *LocalizedPolicy) getDefinitionVersion(ctx context.Context, id string) types.Int64 {
	for _, item := range data.Definitions {
		if item.Id.ValueString() == id {
			return item.Version
		}
	}
	return types.Int64Null()
}

func (data *LocalizedPolicy) updateDefinitionVersions(ctx context.Context, state LocalizedPolicy) {
	for d := range data.Definitions {
		id := data.Definitions[d].Id.ValueString()
		data.Definitions[d].Version = state.getDefinitionVersion(ctx, id)
	}
}
