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

// Section below is generated&owned by "gen/generator.go". //template:begin testAcc
func TestAccSdwanServiceWirelessLANProfileParcel(t *testing.T) {
	if os.Getenv("SDWAN_2012") == "" {
		t.Skip("skipping test, set environment variable SDWAN_2012")
	}
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_service_wireless_lan_feature.test", "enable_24g", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_service_wireless_lan_feature.test", "enable_5g", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_service_wireless_lan_feature.test", "ssids.0.ssid_name", "SSID_1"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_service_wireless_lan_feature.test", "ssids.0.admin_state", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_service_wireless_lan_feature.test", "ssids.0.broadcast_ssid", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_service_wireless_lan_feature.test", "ssids.0.vlan_id", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_service_wireless_lan_feature.test", "ssids.0.radio_type", "all"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_service_wireless_lan_feature.test", "ssids.0.security_type", "personal"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_service_wireless_lan_feature.test", "ssids.0.qos_profile", "silver"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_service_wireless_lan_feature.test", "country", "GB"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_service_wireless_lan_feature.test", "username", "user1"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_service_wireless_lan_feature.test", "me_dynamic_ip_enabled", "true"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccSdwanServiceWirelessLANPrerequisitesProfileParcelConfig + testAccSdwanServiceWirelessLANProfileParcelConfig_minimum(),
			},
			{
				Config: testAccSdwanServiceWirelessLANPrerequisitesProfileParcelConfig + testAccSdwanServiceWirelessLANProfileParcelConfig_all(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAcc

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites
const testAccSdwanServiceWirelessLANPrerequisitesProfileParcelConfig = `
resource "sdwan_service_feature_profile" "test" {
  name = "TF_TEST"
  description = "Terraform test"
}
`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigMinimum
func testAccSdwanServiceWirelessLANProfileParcelConfig_minimum() string {
	config := `resource "sdwan_service_wireless_lan_feature" "test" {` + "\n"
	config += ` name = "TF_TEST_MIN"` + "\n"
	config += ` description = "Terraform integration test"` + "\n"
	config += `	feature_profile_id = sdwan_service_feature_profile.test.id` + "\n"
	config += `	ssids = [{` + "\n"
	config += `	  ssid_name = "SSID_1"` + "\n"
	config += `	  admin_state = true` + "\n"
	config += `	  broadcast_ssid = true` + "\n"
	config += `	  vlan_id = 1` + "\n"
	config += `	  security_type = "personal"` + "\n"
	config += `	  passphrase = "MyPassword123"` + "\n"
	config += `	}]` + "\n"
	config += `	country = "GB"` + "\n"
	config += `	username = "user1"` + "\n"
	config += `	password = "Test@316s13"` + "\n"
	config += `	me_dynamic_ip_enabled = true` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigMinimum

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAll
func testAccSdwanServiceWirelessLANProfileParcelConfig_all() string {
	config := `resource "sdwan_service_wireless_lan_feature" "test" {` + "\n"
	config += ` name = "TF_TEST_ALL"` + "\n"
	config += ` description = "Terraform integration test"` + "\n"
	config += `	feature_profile_id = sdwan_service_feature_profile.test.id` + "\n"
	config += `	enable_24g = true` + "\n"
	config += `	enable_5g = true` + "\n"
	config += `	ssids = [{` + "\n"
	config += `	  ssid_name = "SSID_1"` + "\n"
	config += `	  admin_state = true` + "\n"
	config += `	  broadcast_ssid = true` + "\n"
	config += `	  vlan_id = 1` + "\n"
	config += `	  radio_type = "all"` + "\n"
	config += `	  security_type = "personal"` + "\n"
	config += `	  passphrase = "MyPassword123"` + "\n"
	config += `	  qos_profile = "silver"` + "\n"
	config += `	}]` + "\n"
	config += `	country = "GB"` + "\n"
	config += `	username = "user1"` + "\n"
	config += `	password = "Test@316s13"` + "\n"
	config += `	me_dynamic_ip_enabled = true` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigAll
