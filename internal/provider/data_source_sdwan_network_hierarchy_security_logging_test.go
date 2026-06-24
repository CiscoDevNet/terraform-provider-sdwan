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
func TestAccDataSourceSdwanNetworkHierarchySecurityLogging(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_network_hierarchy_security_logging.test", "node_id", ""))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_network_hierarchy_security_logging.test", "high_speed_logging.0.vrf", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_network_hierarchy_security_logging.test", "high_speed_logging.0.server_ip", "10.1.2.1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_network_hierarchy_security_logging.test", "high_speed_logging.0.port", "2055"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_network_hierarchy_security_logging.test", "utd_syslog.0.vpn", "service_lan_vpn1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_network_hierarchy_security_logging.test", "utd_syslog.0.server_ip", "10.1.1.2"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceSdwanNetworkHierarchySecurityLoggingConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAccDataSource

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites
// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSourceConfig
func testAccDataSourceSdwanNetworkHierarchySecurityLoggingConfig() string {
	config := ""

	config += `
		data "sdwan_network_hierarchy_security_logging" "test" {
			id = sdwan_network_hierarchy_security_logging.test.id
		}
	`
	return config
}

// End of section. //template:end testAccDataSourceConfig

// Custom data source test - creates resource first, then reads via data source

func TestAccDataSourceSdwanNetworkHierarchySecurityLoggingWithResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceSdwanNetworkHierarchySecurityLoggingConfig_withResource(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.sdwan_network_hierarchy_security_logging.test", "id"),
					resource.TestCheckResourceAttr("data.sdwan_network_hierarchy_security_logging.test", "high_speed_logging.#", "2"),
					resource.TestCheckResourceAttr("data.sdwan_network_hierarchy_security_logging.test", "high_speed_logging.0.vrf", "1"),
					resource.TestCheckResourceAttr("data.sdwan_network_hierarchy_security_logging.test", "high_speed_logging.0.server_ip", "10.1.2.1"),
					resource.TestCheckResourceAttr("data.sdwan_network_hierarchy_security_logging.test", "high_speed_logging.0.port", "2055"),
					resource.TestCheckResourceAttr("data.sdwan_network_hierarchy_security_logging.test", "utd_syslog.#", "1"),
					resource.TestCheckResourceAttr("data.sdwan_network_hierarchy_security_logging.test", "utd_syslog.0.vpn", "service_lan_vpn1"),
					resource.TestCheckResourceAttr("data.sdwan_network_hierarchy_security_logging.test", "utd_syslog.0.server_ip", "10.1.1.2"),
				),
			},
		},
	})
}

func testAccDataSourceSdwanNetworkHierarchySecurityLoggingConfig_withResource() string {
	config := `resource "sdwan_network_hierarchy_security_logging" "test" {` + "\n"
	config += `	high_speed_logging = [{` + "\n"
	config += `	  vrf = "1"` + "\n"
	config += `	  server_ip = "10.1.2.1"` + "\n"
	config += `	  port = 2055` + "\n"
	config += `	}, {` + "\n"
	config += `	  vrf = "2"` + "\n"
	config += `	  server_ip = "10.1.2.2"` + "\n"
	config += `	  port = 2056` + "\n"
	config += `	}]` + "\n"
	config += `	utd_syslog = [{` + "\n"
	config += `	  vpn = "service_lan_vpn1"` + "\n"
	config += `	  server_ip = "10.1.1.2"` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"
	config += "\n"
	config += `data "sdwan_network_hierarchy_security_logging" "test" {` + "\n"
	config += `	id = sdwan_network_hierarchy_security_logging.test.id` + "\n"
	config += `}` + "\n"
	return config
}
