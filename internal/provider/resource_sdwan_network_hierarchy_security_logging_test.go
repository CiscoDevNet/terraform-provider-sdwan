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

// Section below is generated&owned by "gen/generator.go". //template:begin testAcc
func TestAccSdwanNetworkHierarchySecurityLogging(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_network_hierarchy_security_logging.test", "high_speed_logging.0.vrf", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_network_hierarchy_security_logging.test", "high_speed_logging.0.server_ip", "10.1.2.1"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_network_hierarchy_security_logging.test", "high_speed_logging.0.port", "2055"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_network_hierarchy_security_logging.test", "utd_syslog.0.vpn", "service_lan_vpn1"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_network_hierarchy_security_logging.test", "utd_syslog.0.server_ip", "10.1.1.2"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccSdwanNetworkHierarchySecurityLoggingConfig_all(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAcc

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites
// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAll
func testAccSdwanNetworkHierarchySecurityLoggingConfig_all() string {
	config := `resource "sdwan_network_hierarchy_security_logging" "test" {` + "\n"
	config += `	high_speed_logging = [{` + "\n"
	config += `	  vrf = "1"` + "\n"
	config += `	  server_ip = "10.1.2.1"` + "\n"
	config += `	  port = 2055` + "\n"
	config += `	}]` + "\n"
	config += `	utd_syslog = [{` + "\n"
	config += `	  vpn = "service_lan_vpn1"` + "\n"
	config += `	  server_ip = "10.1.1.2"` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigAll

// Custom tests - persist through code generation

func TestAccSdwanNetworkHierarchySecurityLoggingMultipleHSL(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccSdwanNetworkHierarchySecurityLoggingConfig_multipleHSL(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("sdwan_network_hierarchy_security_logging.test", "id"),
					resource.TestCheckResourceAttr("sdwan_network_hierarchy_security_logging.test", "high_speed_logging.#", "4"),
				),
			},
		},
	})
}

func testAccSdwanNetworkHierarchySecurityLoggingConfig_multipleHSL() string {
	config := `resource "sdwan_network_hierarchy_security_logging" "test" {` + "\n"
	config += `	high_speed_logging = [{` + "\n"
	config += `	  vrf = "1"` + "\n"
	config += `	  server_ip = "10.1.1.1"` + "\n"
	config += `	  port = 2055` + "\n"
	config += `	}, {` + "\n"
	config += `	  vrf = "2"` + "\n"
	config += `	  server_ip = "10.1.1.2"` + "\n"
	config += `	  port = 2056` + "\n"
	config += `	}, {` + "\n"
	config += `	  vrf = "3"` + "\n"
	config += `	  server_ip = "10.1.1.3"` + "\n"
	config += `	  port = 2057` + "\n"
	config += `	}, {` + "\n"
	config += `	  vrf = "4"` + "\n"
	config += `	  server_ip = "10.1.1.4"` + "\n"
	config += `	  port = 2058` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"
	return config
}

func TestAccSdwanNetworkHierarchySecurityLoggingHSLOnly(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccSdwanNetworkHierarchySecurityLoggingConfig_hslOnly(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("sdwan_network_hierarchy_security_logging.test", "id"),
					resource.TestCheckResourceAttr("sdwan_network_hierarchy_security_logging.test", "high_speed_logging.#", "2"),
					resource.TestCheckResourceAttr("sdwan_network_hierarchy_security_logging.test", "high_speed_logging.0.vrf", "10"),
					resource.TestCheckResourceAttr("sdwan_network_hierarchy_security_logging.test", "high_speed_logging.0.server_ip", "192.168.1.100"),
					resource.TestCheckResourceAttr("sdwan_network_hierarchy_security_logging.test", "high_speed_logging.0.port", "3000"),
					resource.TestCheckResourceAttr("sdwan_network_hierarchy_security_logging.test", "utd_syslog.#", "0"),
				),
			},
		},
	})
}

func testAccSdwanNetworkHierarchySecurityLoggingConfig_hslOnly() string {
	config := `resource "sdwan_network_hierarchy_security_logging" "test" {` + "\n"
	config += `	high_speed_logging = [{` + "\n"
	config += `	  vrf = "10"` + "\n"
	config += `	  server_ip = "192.168.1.100"` + "\n"
	config += `	  port = 3000` + "\n"
	config += `	}, {` + "\n"
	config += `	  vrf = "20"` + "\n"
	config += `	  server_ip = "192.168.1.101"` + "\n"
	config += `	  port = 3001` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"
	return config
}

func TestAccSdwanNetworkHierarchySecurityLoggingUtdOnly(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccSdwanNetworkHierarchySecurityLoggingConfig_utdOnly(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("sdwan_network_hierarchy_security_logging.test", "id"),
					resource.TestCheckResourceAttr("sdwan_network_hierarchy_security_logging.test", "utd_syslog.#", "1"),
					resource.TestCheckResourceAttr("sdwan_network_hierarchy_security_logging.test", "utd_syslog.0.vpn", "service_vpn_100"),
					resource.TestCheckResourceAttr("sdwan_network_hierarchy_security_logging.test", "utd_syslog.0.server_ip", "172.16.0.50"),
					resource.TestCheckResourceAttr("sdwan_network_hierarchy_security_logging.test", "high_speed_logging.#", "0"),
				),
			},
		},
	})
}

func testAccSdwanNetworkHierarchySecurityLoggingConfig_utdOnly() string {
	config := `resource "sdwan_network_hierarchy_security_logging" "test" {` + "\n"
	config += `	utd_syslog = [{` + "\n"
	config += `	  vpn = "service_vpn_100"` + "\n"
	config += `	  server_ip = "172.16.0.50"` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"
	return config
}

func TestAccSdwanNetworkHierarchySecurityLoggingBothTypes(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccSdwanNetworkHierarchySecurityLoggingConfig_bothTypes(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("sdwan_network_hierarchy_security_logging.test", "id"),
					resource.TestCheckResourceAttr("sdwan_network_hierarchy_security_logging.test", "high_speed_logging.#", "2"),
					resource.TestCheckResourceAttr("sdwan_network_hierarchy_security_logging.test", "utd_syslog.#", "1"),
					resource.TestCheckResourceAttr("sdwan_network_hierarchy_security_logging.test", "utd_syslog.0.vpn", "corporate_vpn"),
					resource.TestCheckResourceAttr("sdwan_network_hierarchy_security_logging.test", "utd_syslog.0.server_ip", "10.100.0.1"),
				),
			},
		},
	})
}

func testAccSdwanNetworkHierarchySecurityLoggingConfig_bothTypes() string {
	config := `resource "sdwan_network_hierarchy_security_logging" "test" {` + "\n"
	config += `	high_speed_logging = [{` + "\n"
	config += `	  vrf = "100"` + "\n"
	config += `	  server_ip = "10.200.1.1"` + "\n"
	config += `	  port = 4000` + "\n"
	config += `	}, {` + "\n"
	config += `	  vrf = "200"` + "\n"
	config += `	  server_ip = "10.200.1.2"` + "\n"
	config += `	  port = 4001` + "\n"
	config += `	}]` + "\n"
	config += `	utd_syslog = [{` + "\n"
	config += `	  vpn = "corporate_vpn"` + "\n"
	config += `	  server_ip = "10.100.0.1"` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"
	return config
}
