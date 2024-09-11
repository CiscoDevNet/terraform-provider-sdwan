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
func TestAccDataSourceSdwanServiceLANVPNInterfaceSVIProfileParcel(t *testing.T) {
	if os.Getenv("SDWAN_2012") == "" {
		t.Skip("skipping test, set environment variable SDWAN_2012")
	}
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_service_lan_vpn_interface_svi_feature.test", "shutdown", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_service_lan_vpn_interface_svi_feature.test", "interface_description", "SVI"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_service_lan_vpn_interface_svi_feature.test", "interface_mtu", "1500"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_service_lan_vpn_interface_svi_feature.test", "ip_mtu", "1500"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_service_lan_vpn_interface_svi_feature.test", "ipv4_secondary_addresses.0.address", "2.3.4.5"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_service_lan_vpn_interface_svi_feature.test", "ipv4_secondary_addresses.0.ipv4_subnet_mask", "0.0.0.0"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_service_lan_vpn_interface_svi_feature.test", "ipv6_address", "2001:0:0:1::0/32"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_service_lan_vpn_interface_svi_feature.test", "ipv6_secondary_addresses.0.address", "::2/32"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_service_lan_vpn_interface_svi_feature.test", "ipv6_dhcp_helpers.0.address", "2001:0:0:1::0"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_service_lan_vpn_interface_svi_feature.test", "ipv6_dhcp_helpers.0.vpn", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_service_lan_vpn_interface_svi_feature.test", "arps.0.ip_address", "1.2.3.4"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_service_lan_vpn_interface_svi_feature.test", "arps.0.mac_address", "00-B0-D0-63-C2-26"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_service_lan_vpn_interface_svi_feature.test", "ipv4_vrrps.0.group_id", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_service_lan_vpn_interface_svi_feature.test", "ipv4_vrrps.0.priority", "100"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_service_lan_vpn_interface_svi_feature.test", "ipv4_vrrps.0.timer", "1000"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_service_lan_vpn_interface_svi_feature.test", "ipv4_vrrps.0.track_omp", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_service_lan_vpn_interface_svi_feature.test", "ipv4_vrrps.0.prefix_list", "prefix"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_service_lan_vpn_interface_svi_feature.test", "ipv4_vrrps.0.address", "1.2.3.4"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_service_lan_vpn_interface_svi_feature.test", "ipv4_vrrps.0.secondary_addresses.0.address", "2.3.4.5"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_service_lan_vpn_interface_svi_feature.test", "ipv4_vrrps.0.tloc_prefix_change", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_service_lan_vpn_interface_svi_feature.test", "ipv4_vrrps.0.tloc_prefix_change_value", "100"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_service_lan_vpn_interface_svi_feature.test", "ipv6_vrrps.0.group_id", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_service_lan_vpn_interface_svi_feature.test", "ipv6_vrrps.0.priority", "100"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_service_lan_vpn_interface_svi_feature.test", "ipv6_vrrps.0.timer", "1000"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_service_lan_vpn_interface_svi_feature.test", "ipv6_vrrps.0.track_omp", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_service_lan_vpn_interface_svi_feature.test", "ipv6_vrrps.0.track_prefix_list", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_service_lan_vpn_interface_svi_feature.test", "ipv6_vrrps.0.addresses.0.link_local_address", "1::1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_service_lan_vpn_interface_svi_feature.test", "ipv6_vrrps.0.addresses.0.global_address", "1::1/24"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_service_lan_vpn_interface_svi_feature.test", "ipv6_vrrps.0.secondary_addresses.0.prefix", "::20/32"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_service_lan_vpn_interface_svi_feature.test", "enable_dhcpv6", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_service_lan_vpn_interface_svi_feature.test", "tcp_mss", "1024"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_service_lan_vpn_interface_svi_feature.test", "arp_timeout", "1200"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_service_lan_vpn_interface_svi_feature.test", "ip_directed_broadcast", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_service_lan_vpn_interface_svi_feature.test", "icmp_redirect_disable", "true"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceSdwanServiceLANVPNInterfaceSVIPrerequisitesProfileParcelConfig + testAccDataSourceSdwanServiceLANVPNInterfaceSVIProfileParcelConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAccDataSource

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites
const testAccDataSourceSdwanServiceLANVPNInterfaceSVIPrerequisitesProfileParcelConfig = `
resource "sdwan_service_feature_profile" "test" {
  name = "TF_TEST"
  description = "Terraform test"
}

resource "sdwan_service_lan_vpn_feature" "test" {
  name = "TF_TEST_SLAN"
  feature_profile_id = sdwan_service_feature_profile.test.id
}
`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSourceConfig
func testAccDataSourceSdwanServiceLANVPNInterfaceSVIProfileParcelConfig() string {
	config := `resource "sdwan_service_lan_vpn_interface_svi_feature" "test" {` + "\n"
	config += ` name = "TF_TEST"` + "\n"
	config += ` description = "Terraform integration test"` + "\n"
	config += `	feature_profile_id = sdwan_service_feature_profile.test.id` + "\n"
	config += `	service_lan_vpn_feature_id = sdwan_service_lan_vpn_feature.test.id` + "\n"
	config += `	shutdown = false` + "\n"
	config += `	interface_name = "Vlan1"` + "\n"
	config += `	interface_description = "SVI"` + "\n"
	config += `	interface_mtu = 1500` + "\n"
	config += `	ip_mtu = 1500` + "\n"
	config += `	ipv4_address = "1.2.3.4"` + "\n"
	config += `	ipv4_subnet_mask = "0.0.0.0"` + "\n"
	config += `	ipv4_secondary_addresses = [{` + "\n"
	config += `	  address = "2.3.4.5"` + "\n"
	config += `	  ipv4_subnet_mask = "0.0.0.0"` + "\n"
	config += `	}]` + "\n"
	config += `	ipv4_dhcp_helpers = ["4.5.6.7"]` + "\n"
	config += `	ipv6_address = "2001:0:0:1::0/32"` + "\n"
	config += `	ipv6_secondary_addresses = [{` + "\n"
	config += `	  address = "::2/32"` + "\n"
	config += `	}]` + "\n"
	config += `	ipv6_dhcp_helpers = [{` + "\n"
	config += `	  address = "2001:0:0:1::0"` + "\n"
	config += `	  vpn = 1` + "\n"
	config += `	}]` + "\n"
	config += `	arps = [{` + "\n"
	config += `	  ip_address = "1.2.3.4"` + "\n"
	config += `	  mac_address = "00-B0-D0-63-C2-26"` + "\n"
	config += `	}]` + "\n"
	config += `	ipv4_vrrps = [{` + "\n"
	config += `	  group_id = 1` + "\n"
	config += `	  priority = 100` + "\n"
	config += `	  timer = 1000` + "\n"
	config += `	  track_omp = false` + "\n"
	config += `	  prefix_list = "prefix"` + "\n"
	config += `	  address = "1.2.3.4"` + "\n"
	config += `	  secondary_addresses = [{` + "\n"
	config += `		address = "2.3.4.5"` + "\n"
	config += `	}]` + "\n"
	config += `	  tloc_prefix_change = true` + "\n"
	config += `	  tloc_prefix_change_value = 100` + "\n"
	config += `	}]` + "\n"
	config += `	ipv6_vrrps = [{` + "\n"
	config += `	  group_id = 1` + "\n"
	config += `	  priority = 100` + "\n"
	config += `	  timer = 1000` + "\n"
	config += `	  track_omp = false` + "\n"
	config += `	  track_prefix_list = "1"` + "\n"
	config += `	  addresses = [{` + "\n"
	config += `		link_local_address = "1::1"` + "\n"
	config += `		global_address = "1::1/24"` + "\n"
	config += `	}]` + "\n"
	config += `	  secondary_addresses = [{` + "\n"
	config += `		prefix = "::20/32"` + "\n"
	config += `	}]` + "\n"
	config += `	}]` + "\n"
	config += `	enable_dhcpv6 = false` + "\n"
	config += `	tcp_mss = 1024` + "\n"
	config += `	arp_timeout = 1200` + "\n"
	config += `	ip_directed_broadcast = false` + "\n"
	config += `	icmp_redirect_disable = true` + "\n"
	config += `}` + "\n"

	config += `
		data "sdwan_service_lan_vpn_interface_svi_feature" "test" {
			id = sdwan_service_lan_vpn_interface_svi_feature.test.id
			feature_profile_id = sdwan_service_feature_profile.test.id
			service_lan_vpn_feature_id = sdwan_service_lan_vpn_feature.test.id
		}
	`
	return config
}

// End of section. //template:end testAccDataSourceConfig
