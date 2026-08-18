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
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/tidwall/gjson"
)

// Issue #744: a static ipv4_subnet_mask must stay optionType=global when
// ipv4_address is a device variable. addressV4.ipAddress and
// addressV4.subnetMask are independent in the Manager schema.
func TestServiceLANVPNInterfaceSVIToBodyMixedAddressVariableStaticMask(t *testing.T) {
	data := ServiceLANVPNInterfaceSVI{
		Name:                types.StringValue("SVI-VLAN10"),
		InterfaceName:       types.StringValue("Vlan10"),
		Ipv4AddressVariable: types.StringValue("{{SVI-IP-Address}}"),
		Ipv4SubnetMask:      types.StringValue("255.255.255.0"),
		Shutdown:            types.BoolValue(false),
	}

	body := data.toBody(context.Background())
	res := gjson.Parse(body)

	if got := res.Get("data.ipv4.addressV4.ipAddress.optionType").String(); got != "variable" {
		t.Fatalf("ipAddress.optionType = %q, want variable", got)
	}
	if got := res.Get("data.ipv4.addressV4.ipAddress.value").String(); got != "{{SVI-IP-Address}}" {
		t.Fatalf("ipAddress.value = %q, want {{SVI-IP-Address}}", got)
	}
	if got := res.Get("data.ipv4.addressV4.subnetMask.optionType").String(); got != "global" {
		t.Fatalf("subnetMask.optionType = %q, want global (issue #744)", got)
	}
	if got := res.Get("data.ipv4.addressV4.subnetMask.value").String(); got != "255.255.255.0" {
		t.Fatalf("subnetMask.value = %q, want 255.255.255.0", got)
	}
	if got := res.Get("data.shutdown.optionType").String(); got != "global" {
		t.Fatalf("shutdown.optionType = %q, want global", got)
	}
}

func TestServiceLANVPNInterfaceSVIFromBodyMixedAddressVariableStaticMask(t *testing.T) {
	payload := `{
		"payload": {
			"name": "SVI-VLAN10",
			"data": {
				"ipv4": {
					"addressV4": {
						"ipAddress": {"optionType": "variable", "value": "{{SVI-IP-Address}}"},
						"subnetMask": {"optionType": "global", "value": "255.255.255.0"}
					}
				}
			}
		}
	}`

	data := ServiceLANVPNInterfaceSVI{}
	data.fromBody(context.Background(), gjson.Parse(payload), true)

	if data.Ipv4AddressVariable.ValueString() != "{{SVI-IP-Address}}" {
		t.Fatalf("Ipv4AddressVariable = %q", data.Ipv4AddressVariable.ValueString())
	}
	if !data.Ipv4Address.IsNull() {
		t.Fatalf("Ipv4Address should be null, got %q", data.Ipv4Address.ValueString())
	}
	if data.Ipv4SubnetMask.ValueString() != "255.255.255.0" {
		t.Fatalf("Ipv4SubnetMask = %q, want 255.255.255.0", data.Ipv4SubnetMask.ValueString())
	}
	if !data.Ipv4SubnetMaskVariable.IsNull() {
		t.Fatalf("Ipv4SubnetMaskVariable should be null, got %q", data.Ipv4SubnetMaskVariable.ValueString())
	}
}
