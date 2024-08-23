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
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSource
func TestAccDataSourceSdwanServiceSwitchportProfileParcel(t *testing.T) {
	if os.Getenv("SDWAN_2012") == "" {
		t.Skip("skipping test, set environment variable SDWAN_2012")
	}
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_service_switchport_profile_parcel.test", "interface.0.interface_name", "GigabitEthernet"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_service_switchport_profile_parcel.test", "interface.0.mode", "access"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_service_switchport_profile_parcel.test", "interface.0.shutdown", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_service_switchport_profile_parcel.test", "interface.0.speed", "10"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_service_switchport_profile_parcel.test", "interface.0.duplex", "full"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_service_switchport_profile_parcel.test", "interface.0.switchport_access_vlan", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_service_switchport_profile_parcel.test", "interface.0.switchport_trunk_allowed_vlans", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_service_switchport_profile_parcel.test", "interface.0.switchport_trunk_native_vlan", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_service_switchport_profile_parcel.test", "interface.0.port_control", "auto"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_service_switchport_profile_parcel.test", "interface.0.voice_vlan", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_service_switchport_profile_parcel.test", "interface.0.pae_enable", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_service_switchport_profile_parcel.test", "interface.0.mac_authentication_bypass", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_service_switchport_profile_parcel.test", "interface.0.host_mode", "single-host"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_service_switchport_profile_parcel.test", "interface.0.enable_periodic_reauth", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_service_switchport_profile_parcel.test", "interface.0.inactivity", "60"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_service_switchport_profile_parcel.test", "interface.0.reauthentication", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_service_switchport_profile_parcel.test", "interface.0.control_direction", "both"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_service_switchport_profile_parcel.test", "interface.0.restricted_vlan", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_service_switchport_profile_parcel.test", "interface.0.guest_vlan", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_service_switchport_profile_parcel.test", "interface.0.critical_vlan", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_service_switchport_profile_parcel.test", "interface.0.enable_voice", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_service_switchport_profile_parcel.test", "age_out_time", "300"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_service_switchport_profile_parcel.test", "static_mac_addresses.0.mac_address", "01:02:03:04:05:06"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_service_switchport_profile_parcel.test", "static_mac_addresses.0.vlan_id", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_service_switchport_profile_parcel.test", "static_mac_addresses.0.interface_name", "GigabitEthernet0/0/0"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceSdwanServiceSwitchportPrerequisitesProfileParcelConfig + testAccDataSourceSdwanServiceSwitchportProfileParcelConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAccDataSource

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites
const testAccDataSourceSdwanServiceSwitchportPrerequisitesProfileParcelConfig = `
resource "sdwan_service_feature_profile" "test" {
  name = "TF_TEST"
  description = "Terraform test"
}
`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSourceConfig
func testAccDataSourceSdwanServiceSwitchportProfileParcelConfig() string {
	config := `resource "sdwan_service_switchport_profile_parcel" "test" {` + "\n"
	config += ` name = "TF_TEST"` + "\n"
	config += ` description = "Terraform integration test"` + "\n"
	config += `	feature_profile_id = sdwan_service_feature_profile.test.id` + "\n"
	config += `	interface = [{` + "\n"
	config += `	  interface_name = "GigabitEthernet"` + "\n"
	config += `	  mode = "access"` + "\n"
	config += `	  shutdown = true` + "\n"
	config += `	  speed = "10"` + "\n"
	config += `	  duplex = "full"` + "\n"
	config += `	  switchport_access_vlan = 1` + "\n"
	config += `	  switchport_trunk_allowed_vlans = "1"` + "\n"
	config += `	  switchport_trunk_native_vlan = 1` + "\n"
	config += `	  port_control = "auto"` + "\n"
	config += `	  voice_vlan = 1` + "\n"
	config += `	  pae_enable = true` + "\n"
	config += `	  mac_authentication_bypass = false` + "\n"
	config += `	  host_mode = "single-host"` + "\n"
	config += `	  enable_periodic_reauth = false` + "\n"
	config += `	  inactivity = 60` + "\n"
	config += `	  reauthentication = 1` + "\n"
	config += `	  control_direction = "both"` + "\n"
	config += `	  restricted_vlan = 1` + "\n"
	config += `	  guest_vlan = 1` + "\n"
	config += `	  critical_vlan = 1` + "\n"
	config += `	  enable_voice = false` + "\n"
	config += `	}]` + "\n"
	config += `	age_out_time = 300` + "\n"
	config += `	static_mac_addresses = [{` + "\n"
	config += `	  mac_address = "01:02:03:04:05:06"` + "\n"
	config += `	  vlan_id = 1` + "\n"
	config += `	  interface_name = "GigabitEthernet0/0/0"` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"

	config += `
		data "sdwan_service_switchport_profile_parcel" "test" {
			id = sdwan_service_switchport_profile_parcel.test.id
			feature_profile_id = sdwan_service_feature_profile.test.id
		}
	`
	return config
}

// End of section. //template:end testAccDataSourceConfig
