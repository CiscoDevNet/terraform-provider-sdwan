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

func TestAccDataSourceSdwanCiscoVPNInterfaceGREFeatureTemplate(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceSdwanCiscoVPNInterfaceGREFeatureTemplateConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.sdwan_cisco_vpn_interface_gre_feature_template.test", "interface_name", "gre0/0"),
					resource.TestCheckResourceAttr("data.sdwan_cisco_vpn_interface_gre_feature_template.test", "interface_description", "My Description"),
					resource.TestCheckResourceAttr("data.sdwan_cisco_vpn_interface_gre_feature_template.test", "ip_address", "1.1.1.1/24"),
					resource.TestCheckResourceAttr("data.sdwan_cisco_vpn_interface_gre_feature_template.test", "tunnel_source", "1.2.3.4"),
					resource.TestCheckResourceAttr("data.sdwan_cisco_vpn_interface_gre_feature_template.test", "shutdown", "true"),
					resource.TestCheckResourceAttr("data.sdwan_cisco_vpn_interface_gre_feature_template.test", "tunnel_source_interface", "e1"),
					resource.TestCheckResourceAttr("data.sdwan_cisco_vpn_interface_gre_feature_template.test", "tunnel_destination", "3.4.5.6"),
					resource.TestCheckResourceAttr("data.sdwan_cisco_vpn_interface_gre_feature_template.test", "application", "sig"),
					resource.TestCheckResourceAttr("data.sdwan_cisco_vpn_interface_gre_feature_template.test", "ip_mtu", "1500"),
					resource.TestCheckResourceAttr("data.sdwan_cisco_vpn_interface_gre_feature_template.test", "tcp_mss_adjust", "1400"),
					resource.TestCheckResourceAttr("data.sdwan_cisco_vpn_interface_gre_feature_template.test", "clear_dont_fragment", "true"),
					resource.TestCheckResourceAttr("data.sdwan_cisco_vpn_interface_gre_feature_template.test", "rewrite_rule", "ACL1"),
					resource.TestCheckResourceAttr("data.sdwan_cisco_vpn_interface_gre_feature_template.test", "access_lists.0.direction", "in"),
					resource.TestCheckResourceAttr("data.sdwan_cisco_vpn_interface_gre_feature_template.test", "access_lists.0.acl_name", "ACL2"),
					resource.TestCheckResourceAttr("data.sdwan_cisco_vpn_interface_gre_feature_template.test", "tunnel_route_via", "g0/0"),
				),
			},
		},
	})
}

const testAccDataSourceSdwanCiscoVPNInterfaceGREFeatureTemplateConfig = `

resource "sdwan_cisco_vpn_interface_gre_feature_template" "test" {
  name = "TF_TEST_MIN"
  description = "Terraform integration test"
  device_types = ["vedge-C8000V"]
  interface_name = "gre0/0"
  interface_description = "My Description"
  ip_address = "1.1.1.1/24"
  tunnel_source = "1.2.3.4"
  shutdown = true
  tunnel_source_interface = "e1"
  tunnel_destination = "3.4.5.6"
  application = "sig"
  ip_mtu = 1500
  tcp_mss_adjust = 1400
  clear_dont_fragment = true
  rewrite_rule = "ACL1"
  access_lists = [{
    direction = "in"
    acl_name = "ACL2"
  }]
  tracker = ["TRACKER1"]
  tunnel_route_via = "g0/0"
}

data "sdwan_cisco_vpn_interface_gre_feature_template" "test" {
  id = sdwan_cisco_vpn_interface_gre_feature_template.test.id
}
`
