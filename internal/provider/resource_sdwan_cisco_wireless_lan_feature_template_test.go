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
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccSdwanCiscoWirelessLANFeatureTemplate(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccSdwanCiscoWirelessLANFeatureTemplateConfig_minimum(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("sdwan_cisco_wireless_lan_feature_template.test", "country", "AE"),
					resource.TestCheckResourceAttr("sdwan_cisco_wireless_lan_feature_template.test", "username", "example"),
					resource.TestCheckResourceAttr("sdwan_cisco_wireless_lan_feature_template.test", "password", "myPassword01"),
				),
			},
			{
				Config: testAccSdwanCiscoWirelessLANFeatureTemplateConfig_all(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("sdwan_cisco_wireless_lan_feature_template.test", "shutdown_2_4ghz", "false"),
					resource.TestCheckResourceAttr("sdwan_cisco_wireless_lan_feature_template.test", "shutdown_5ghz", "false"),
					resource.TestCheckResourceAttr("sdwan_cisco_wireless_lan_feature_template.test", "ssid.0.wireless_network_name", "example"),
					resource.TestCheckResourceAttr("sdwan_cisco_wireless_lan_feature_template.test", "ssid.0.admin_state", "false"),
					resource.TestCheckResourceAttr("sdwan_cisco_wireless_lan_feature_template.test", "ssid.0.broadcast_ssid", "true"),
					resource.TestCheckResourceAttr("sdwan_cisco_wireless_lan_feature_template.test", "ssid.0.vlan_id", "1"),
					resource.TestCheckResourceAttr("sdwan_cisco_wireless_lan_feature_template.test", "ssid.0.radio_type", "24ghz"),
					resource.TestCheckResourceAttr("sdwan_cisco_wireless_lan_feature_template.test", "ssid.0.security_type", "enterprise"),
					resource.TestCheckResourceAttr("sdwan_cisco_wireless_lan_feature_template.test", "ssid.0.radius_server_ip", "1.2.3.4"),
					resource.TestCheckResourceAttr("sdwan_cisco_wireless_lan_feature_template.test", "ssid.0.authentication_port", "1812"),
					resource.TestCheckResourceAttr("sdwan_cisco_wireless_lan_feature_template.test", "ssid.0.shared_secret", "example"),
					resource.TestCheckResourceAttr("sdwan_cisco_wireless_lan_feature_template.test", "ssid.0.passphrase", "passphrase"),
					resource.TestCheckResourceAttr("sdwan_cisco_wireless_lan_feature_template.test", "ssid.0.qos_profile", "silver"),
					resource.TestCheckResourceAttr("sdwan_cisco_wireless_lan_feature_template.test", "country", "AE"),
					resource.TestCheckResourceAttr("sdwan_cisco_wireless_lan_feature_template.test", "username", "example"),
					resource.TestCheckResourceAttr("sdwan_cisco_wireless_lan_feature_template.test", "password", "myPassword01"),
					resource.TestCheckResourceAttr("sdwan_cisco_wireless_lan_feature_template.test", "controller_ip_address", "0.0.0.0"),
					resource.TestCheckResourceAttr("sdwan_cisco_wireless_lan_feature_template.test", "subnet_mask", "0.0.0.0"),
					resource.TestCheckResourceAttr("sdwan_cisco_wireless_lan_feature_template.test", "default_gateway", "0.0.0.0"),
				),
			},
		},
	})
}

func testAccSdwanCiscoWirelessLANFeatureTemplateConfig_minimum() string {
	return `
	resource "sdwan_cisco_wireless_lan_feature_template" "test" {
		name = "TF_TEST_MIN"
		description = "Terraform integration test"
		device_types = ["vedge-C8000V"]
		country = "AE"
		username = "example"
		password = "myPassword01"
	}
	`
}

func testAccSdwanCiscoWirelessLANFeatureTemplateConfig_all() string {
	return `
	resource "sdwan_cisco_wireless_lan_feature_template" "test" {
		name = "TF_TEST_ALL"
		description = "Terraform integration test"
		device_types = ["vedge-C8000V"]
		shutdown_2_4ghz = false
		shutdown_5ghz = false
		ssid = [{
			wireless_network_name = "example"
			admin_state = false
			broadcast_ssid = true
			vlan_id = 1
			radio_type = "24ghz"
			security_type = "enterprise"
			radius_server_ip = "1.2.3.4"
			authentication_port = 1812
			shared_secret = "example"
			passphrase = "passphrase"
			qos_profile = "silver"
		}]
		country = "AE"
		username = "example"
		password = "myPassword01"
		controller_ip_address = "0.0.0.0"
		subnet_mask = "0.0.0.0"
		default_gateway = "0.0.0.0"
	}
	`
}
