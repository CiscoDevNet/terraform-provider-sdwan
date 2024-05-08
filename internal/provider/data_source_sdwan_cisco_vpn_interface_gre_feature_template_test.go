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
func TestAccDataSourceSdwanCiscoVPNInterfaceGREFeatureTemplate(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_vpn_interface_gre_feature_template.test", "interface_name", "gre0/0"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_vpn_interface_gre_feature_template.test", "interface_description", "My Description"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_vpn_interface_gre_feature_template.test", "ip_address", "1.1.1.1/24"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_vpn_interface_gre_feature_template.test", "tunnel_source", "1.2.3.4"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_vpn_interface_gre_feature_template.test", "shutdown", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_vpn_interface_gre_feature_template.test", "tunnel_source_interface", "e1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_vpn_interface_gre_feature_template.test", "tunnel_destination", "3.4.5.6"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_vpn_interface_gre_feature_template.test", "application", "sig"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_vpn_interface_gre_feature_template.test", "ip_mtu", "1500"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_vpn_interface_gre_feature_template.test", "tcp_mss_adjust", "1400"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_vpn_interface_gre_feature_template.test", "clear_dont_fragment", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_vpn_interface_gre_feature_template.test", "rewrite_rule", "ACL1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_vpn_interface_gre_feature_template.test", "access_lists.0.direction", "in"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_vpn_interface_gre_feature_template.test", "access_lists.0.acl_name", "ACL2"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_vpn_interface_gre_feature_template.test", "tunnel_route_via", "g0/0"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceSdwanCiscoVPNInterfaceGREFeatureTemplateConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAccDataSource

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSourceConfig
func testAccDataSourceSdwanCiscoVPNInterfaceGREFeatureTemplateConfig() string {
	config := `resource "sdwan_cisco_vpn_interface_gre_feature_template" "test" {` + "\n"
	config += ` name = "TF_TEST"` + "\n"
	config += ` description = "Terraform integration test"` + "\n"
	config += ` device_types = ["vedge-C8000V"]` + "\n"
	config += `	interface_name = "gre0/0"` + "\n"
	config += `	interface_description = "My Description"` + "\n"
	config += `	ip_address = "1.1.1.1/24"` + "\n"
	config += `	tunnel_source = "1.2.3.4"` + "\n"
	config += `	shutdown = true` + "\n"
	config += `	tunnel_source_interface = "e1"` + "\n"
	config += `	tunnel_destination = "3.4.5.6"` + "\n"
	config += `	application = "sig"` + "\n"
	config += `	ip_mtu = 1500` + "\n"
	config += `	tcp_mss_adjust = 1400` + "\n"
	config += `	clear_dont_fragment = true` + "\n"
	config += `	rewrite_rule = "ACL1"` + "\n"
	config += `	access_lists = [{` + "\n"
	config += `	  direction = "in"` + "\n"
	config += `	  acl_name = "ACL2"` + "\n"
	config += `	}]` + "\n"
	config += `	tracker = ["TRACKER1"]` + "\n"
	config += `	tunnel_route_via = "g0/0"` + "\n"
	config += `}` + "\n"

	config += `
		data "sdwan_cisco_vpn_interface_gre_feature_template" "test" {
			id = sdwan_cisco_vpn_interface_gre_feature_template.test.id
		}
	`
	return config
}

// End of section. //template:end testAccDataSourceConfig
