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

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type DomainListPolicyObject struct {
	Id      types.String                    `tfsdk:"id"`
	Version types.Int64                     `tfsdk:"version"`
	Name    types.String                    `tfsdk:"name"`
	Entries []DomainListPolicyObjectEntries `tfsdk:"entries"`
}

type DomainListPolicyObjectEntries struct {
	Domain types.String `tfsdk:"domain"`
}

func (data DomainListPolicyObject) toBody(ctx context.Context) string {
	body := ""
	if true {
		body, _ = sjson.Set(body, "type", "localdomain")
	}
	if !data.Name.IsNull() {
		body, _ = sjson.Set(body, "name", data.Name.ValueString())
	}
	if true {
		body, _ = sjson.Set(body, "entries", []interface{}{})
		for _, item := range data.Entries {
			itemBody := ""
			if !item.Domain.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "nameServer", item.Domain.ValueString())
			}
			body, _ = sjson.SetRaw(body, "entries.-1", itemBody)
		}
	}
	return body
}

func (data *DomainListPolicyObject) fromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("name"); value.Exists() {
		data.Name = types.StringValue(value.String())
	} else {
		data.Name = types.StringNull()
	}
	if value := res.Get("entries"); value.Exists() && len(value.Array()) > 0 {
		data.Entries = make([]DomainListPolicyObjectEntries, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := DomainListPolicyObjectEntries{}
			if cValue := v.Get("nameServer"); cValue.Exists() {
				item.Domain = types.StringValue(cValue.String())
			} else {
				item.Domain = types.StringNull()
			}
			data.Entries = append(data.Entries, item)
			return true
		})
	} else {
		if len(data.Entries) > 0 {
			data.Entries = []DomainListPolicyObjectEntries{}
		}
	}
}

func (data *DomainListPolicyObject) hasChanges(ctx context.Context, state *DomainListPolicyObject) bool {
	hasChanges := false
	if !data.Name.Equal(state.Name) {
		hasChanges = true
	}
	if len(data.Entries) != len(state.Entries) {
		hasChanges = true
	} else {
		for i := range data.Entries {
			if !data.Entries[i].Domain.Equal(state.Entries[i].Domain) {
				hasChanges = true
			}
		}
	}
	return hasChanges
}
