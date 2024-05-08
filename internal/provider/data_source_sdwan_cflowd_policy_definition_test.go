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
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSource
func TestAccDataSourceSdwanCflowdPolicyDefinition(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cflowd_policy_definition.test", "name", "Example"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cflowd_policy_definition.test", "description", "My description"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cflowd_policy_definition.test", "active_flow_timeout", "100"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cflowd_policy_definition.test", "inactive_flow_timeout", "10"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cflowd_policy_definition.test", "sampling_interval", "10"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cflowd_policy_definition.test", "flow_refresh", "120"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cflowd_policy_definition.test", "protocol", "ipv4"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cflowd_policy_definition.test", "tos", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cflowd_policy_definition.test", "remarked_dscp", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cflowd_policy_definition.test", "collectors.0.vpn_id", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cflowd_policy_definition.test", "collectors.0.ip_address", "10.0.0.1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cflowd_policy_definition.test", "collectors.0.port", "12345"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cflowd_policy_definition.test", "collectors.0.transport", "transport_tcp"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cflowd_policy_definition.test", "collectors.0.source_interface", "Ethernet1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cflowd_policy_definition.test", "collectors.0.export_spreading", "enable"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceSdwanCflowdPolicyDefinitionConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAccDataSource

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites
// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSourceConfig
func testAccDataSourceSdwanCflowdPolicyDefinitionConfig() string {
	config := ""
	config += `resource "sdwan_cflowd_policy_definition" "test" {` + "\n"
	config += `	name = "Example"` + "\n"
	config += `	description = "My description"` + "\n"
	config += `	active_flow_timeout = 100` + "\n"
	config += `	inactive_flow_timeout = 10` + "\n"
	config += `	sampling_interval = 10` + "\n"
	config += `	flow_refresh = 120` + "\n"
	config += `	protocol = "ipv4"` + "\n"
	config += `	tos = true` + "\n"
	config += `	remarked_dscp = true` + "\n"
	config += `	collectors = [{` + "\n"
	config += `	  vpn_id = 1` + "\n"
	config += `	  ip_address = "10.0.0.1"` + "\n"
	config += `	  port = 12345` + "\n"
	config += `	  transport = "transport_tcp"` + "\n"
	config += `	  source_interface = "Ethernet1"` + "\n"
	config += `	  export_spreading = "enable"` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"

	config += `
		data "sdwan_cflowd_policy_definition" "test" {
			id = sdwan_cflowd_policy_definition.test.id
		}
	`
	return config
}

// End of section. //template:end testAccDataSourceConfig
