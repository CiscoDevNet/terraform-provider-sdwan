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
func TestAccDataSourceSdwanSwitchportFeatureTemplate(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_switchport_feature_template.test", "slot", "0"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_switchport_feature_template.test", "sub_slot", "0"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_switchport_feature_template.test", "module_type", "4"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_switchport_feature_template.test", "interfaces.0.name", "GigabitEthernet0/0/0"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_switchport_feature_template.test", "interfaces.0.switchport_mode", "access"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_switchport_feature_template.test", "interfaces.0.shutdown", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_switchport_feature_template.test", "interfaces.0.speed", "100"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_switchport_feature_template.test", "interfaces.0.duplex", "full"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_switchport_feature_template.test", "interfaces.0.switchport_access_vlan", "100"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_switchport_feature_template.test", "interfaces.0.switchport_trunk_allowed_vlans", "100,200"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_switchport_feature_template.test", "interfaces.0.switchport_trunk_native_vlan", "100"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_switchport_feature_template.test", "interfaces.0.dot1x_enable", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_switchport_feature_template.test", "interfaces.0.dot1x_port_control", "auto"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_switchport_feature_template.test", "interfaces.0.voice_vlan", "200"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_switchport_feature_template.test", "interfaces.0.dot1x_pae_enable", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_switchport_feature_template.test", "interfaces.0.dot1x_mac_authentication_bypass", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_switchport_feature_template.test", "interfaces.0.dot1x_host_mode", "multi-domain"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_switchport_feature_template.test", "interfaces.0.dot1x_enable_periodic_reauth", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_switchport_feature_template.test", "interfaces.0.dot1x_periodic_reauth_inactivity_timeout", "100"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_switchport_feature_template.test", "interfaces.0.dot1x_periodic_reauth_interval", "60"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_switchport_feature_template.test", "interfaces.0.dot1x_control_direction", "both"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_switchport_feature_template.test", "interfaces.0.dot1x_restricted_vlan", "100"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_switchport_feature_template.test", "interfaces.0.dot1x_guest_vlan", "101"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_switchport_feature_template.test", "interfaces.0.dot1x_critical_vlan", "102"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_switchport_feature_template.test", "interfaces.0.dot1x_enable_criticial_voice_vlan", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_switchport_feature_template.test", "age_out_time", "500"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_switchport_feature_template.test", "static_mac_addresses.0.mac_address", "0000.0000.0000"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_switchport_feature_template.test", "static_mac_addresses.0.if_name", "GigabitEthernet0/0/0"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_switchport_feature_template.test", "static_mac_addresses.0.vlan", "100"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceSdwanSwitchportFeatureTemplateConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAccDataSource

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSourceConfig
func testAccDataSourceSdwanSwitchportFeatureTemplateConfig() string {
	config := `resource "sdwan_switchport_feature_template" "test" {` + "\n"
	config += ` name = "TF_TEST"` + "\n"
	config += ` description = "Terraform integration test"` + "\n"
	config += ` device_types = ["vedge-C8000V"]` + "\n"
	config += `	slot = 0` + "\n"
	config += `	sub_slot = 0` + "\n"
	config += `	module_type = "4"` + "\n"
	config += `	interfaces = [{` + "\n"
	config += `	  name = "GigabitEthernet0/0/0"` + "\n"
	config += `	  switchport_mode = "access"` + "\n"
	config += `	  shutdown = true` + "\n"
	config += `	  speed = "100"` + "\n"
	config += `	  duplex = "full"` + "\n"
	config += `	  switchport_access_vlan = 100` + "\n"
	config += `	  switchport_trunk_allowed_vlans = "100,200"` + "\n"
	config += `	  switchport_trunk_native_vlan = 100` + "\n"
	config += `	  dot1x_enable = true` + "\n"
	config += `	  dot1x_port_control = "auto"` + "\n"
	config += `	  dot1x_authentication_order = ["dot1x"]` + "\n"
	config += `	  voice_vlan = 200` + "\n"
	config += `	  dot1x_pae_enable = true` + "\n"
	config += `	  dot1x_mac_authentication_bypass = true` + "\n"
	config += `	  dot1x_host_mode = "multi-domain"` + "\n"
	config += `	  dot1x_enable_periodic_reauth = true` + "\n"
	config += `	  dot1x_periodic_reauth_inactivity_timeout = 100` + "\n"
	config += `	  dot1x_periodic_reauth_interval = 60` + "\n"
	config += `	  dot1x_control_direction = "both"` + "\n"
	config += `	  dot1x_restricted_vlan = 100` + "\n"
	config += `	  dot1x_guest_vlan = 101` + "\n"
	config += `	  dot1x_critical_vlan = 102` + "\n"
	config += `	  dot1x_enable_criticial_voice_vlan = true` + "\n"
	config += `	}]` + "\n"
	config += `	age_out_time = 500` + "\n"
	config += `	static_mac_addresses = [{` + "\n"
	config += `	  mac_address = "0000.0000.0000"` + "\n"
	config += `	  if_name = "GigabitEthernet0/0/0"` + "\n"
	config += `	  vlan = 100` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"

	config += `
		data "sdwan_switchport_feature_template" "test" {
			id = sdwan_switchport_feature_template.test.id
		}
	`
	return config
}

// End of section. //template:end testAccDataSourceConfig
