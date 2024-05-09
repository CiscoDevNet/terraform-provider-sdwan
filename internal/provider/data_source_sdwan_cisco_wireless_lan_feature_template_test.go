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
func TestAccDataSourceSdwanCiscoWirelessLANFeatureTemplate(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_wireless_lan_feature_template.test", "shutdown_2_4ghz", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_wireless_lan_feature_template.test", "shutdown_5ghz", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_wireless_lan_feature_template.test", "ssids.0.wireless_network_name", "WLAN1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_wireless_lan_feature_template.test", "ssids.0.admin_state", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_wireless_lan_feature_template.test", "ssids.0.broadcast_ssid", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_wireless_lan_feature_template.test", "ssids.0.vlan_id", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_wireless_lan_feature_template.test", "ssids.0.radio_type", "24ghz"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_wireless_lan_feature_template.test", "ssids.0.security_type", "enterprise"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_wireless_lan_feature_template.test", "ssids.0.radius_server_ip", "1.2.3.4"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_wireless_lan_feature_template.test", "ssids.0.radius_server_port", "1812"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_wireless_lan_feature_template.test", "ssids.0.radius_server_secret", "MySecret1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_wireless_lan_feature_template.test", "ssids.0.passphrase", "passphrase"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_wireless_lan_feature_template.test", "ssids.0.qos_profile", "silver"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_wireless_lan_feature_template.test", "country", "AE"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_wireless_lan_feature_template.test", "username", "user1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_wireless_lan_feature_template.test", "password", "myPassword01"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_wireless_lan_feature_template.test", "controller_ip_address", "0.0.0.0"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_wireless_lan_feature_template.test", "controller_subnet_mask", "0.0.0.0"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_wireless_lan_feature_template.test", "controller_default_gateway", "0.0.0.0"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceSdwanCiscoWirelessLANFeatureTemplateConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAccDataSource

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSourceConfig
func testAccDataSourceSdwanCiscoWirelessLANFeatureTemplateConfig() string {
	config := `resource "sdwan_cisco_wireless_lan_feature_template" "test" {` + "\n"
	config += ` name = "TF_TEST"` + "\n"
	config += ` description = "Terraform integration test"` + "\n"
	config += ` device_types = ["vedge-C8000V"]` + "\n"
	config += `	shutdown_2_4ghz = false` + "\n"
	config += `	shutdown_5ghz = false` + "\n"
	config += `	ssids = [{` + "\n"
	config += `	  wireless_network_name = "WLAN1"` + "\n"
	config += `	  admin_state = false` + "\n"
	config += `	  broadcast_ssid = true` + "\n"
	config += `	  vlan_id = 1` + "\n"
	config += `	  radio_type = "24ghz"` + "\n"
	config += `	  security_type = "enterprise"` + "\n"
	config += `	  radius_server_ip = "1.2.3.4"` + "\n"
	config += `	  radius_server_port = 1812` + "\n"
	config += `	  radius_server_secret = "MySecret1"` + "\n"
	config += `	  passphrase = "passphrase"` + "\n"
	config += `	  qos_profile = "silver"` + "\n"
	config += `	}]` + "\n"
	config += `	country = "AE"` + "\n"
	config += `	username = "user1"` + "\n"
	config += `	password = "myPassword01"` + "\n"
	config += `	controller_ip_address = "0.0.0.0"` + "\n"
	config += `	controller_subnet_mask = "0.0.0.0"` + "\n"
	config += `	controller_default_gateway = "0.0.0.0"` + "\n"
	config += `}` + "\n"

	config += `
		data "sdwan_cisco_wireless_lan_feature_template" "test" {
			id = sdwan_cisco_wireless_lan_feature_template.test.id
		}
	`
	return config
}

// End of section. //template:end testAccDataSourceConfig
