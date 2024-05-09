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
func TestAccDataSourceSdwanServiceLANVPNProfileParcel(t *testing.T) {
	if os.Getenv("SDWAN_2012") == "" {
		t.Skip("skipping test, set environment variable SDWAN_2012")
	}
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_service_lan_vpn_profile_parcel.test", "vpn", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_service_lan_vpn_profile_parcel.test", "config_description", "VPN1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_service_lan_vpn_profile_parcel.test", "omp_admin_distance_ipv4", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_service_lan_vpn_profile_parcel.test", "omp_admin_distance_ipv6", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_service_lan_vpn_profile_parcel.test", "enable_sdwan_remote_access", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_service_lan_vpn_profile_parcel.test", "primary_dns_address_ipv4", "1.2.3.4"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_service_lan_vpn_profile_parcel.test", "secondary_dns_address_ipv4", "2.3.4.5"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_service_lan_vpn_profile_parcel.test", "primary_dns_address_ipv6", "2001:0:0:1::0"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_service_lan_vpn_profile_parcel.test", "secondary_dns_address_ipv6", "2001:0:0:2::0"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_service_lan_vpn_profile_parcel.test", "host_mappings.0.host_name", "HOSTMAPPING1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_service_lan_vpn_profile_parcel.test", "ipv4_static_routes.0.network_address", "1.2.3.4"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_service_lan_vpn_profile_parcel.test", "ipv4_static_routes.0.subnet_mask", "0.0.0.0"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_service_lan_vpn_profile_parcel.test", "ipv4_static_routes.0.next_hops.0.address", "1.2.3.4"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_service_lan_vpn_profile_parcel.test", "ipv4_static_routes.0.next_hops.0.administrative_distance", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_service_lan_vpn_profile_parcel.test", "ipv6_static_routes.0.prefix", "2001:0:0:1::0/12"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_service_lan_vpn_profile_parcel.test", "ipv6_static_routes.0.next_hops.0.address", "2001:0:0:1::0"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_service_lan_vpn_profile_parcel.test", "ipv6_static_routes.0.next_hops.0.administrative_distance", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_service_lan_vpn_profile_parcel.test", "services.0.service_type", "FW"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_service_lan_vpn_profile_parcel.test", "services.0.tracking", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_service_lan_vpn_profile_parcel.test", "service_routes.0.network_address", "1.2.3.4"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_service_lan_vpn_profile_parcel.test", "service_routes.0.subnet_mask", "0.0.0.0"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_service_lan_vpn_profile_parcel.test", "service_routes.0.service", "SIG"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_service_lan_vpn_profile_parcel.test", "service_routes.0.vpn", "0"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_service_lan_vpn_profile_parcel.test", "gre_routes.0.network_address", "1.2.3.4"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_service_lan_vpn_profile_parcel.test", "gre_routes.0.subnet_mask", "0.0.0.0"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_service_lan_vpn_profile_parcel.test", "gre_routes.0.vpn", "0"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_service_lan_vpn_profile_parcel.test", "ipsec_routes.0.network_address", "1.2.3.4"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_service_lan_vpn_profile_parcel.test", "ipsec_routes.0.subnet_mask", "0.0.0.0"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_service_lan_vpn_profile_parcel.test", "nat_pools.0.nat_pool_name", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_service_lan_vpn_profile_parcel.test", "nat_pools.0.prefix_length", "3"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_service_lan_vpn_profile_parcel.test", "nat_pools.0.range_start", "1.2.3.4"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_service_lan_vpn_profile_parcel.test", "nat_pools.0.range_end", "2.3.4.5"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_service_lan_vpn_profile_parcel.test", "nat_pools.0.overload", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_service_lan_vpn_profile_parcel.test", "nat_pools.0.direction", "inside"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_service_lan_vpn_profile_parcel.test", "nat_port_forwards.0.nat_pool_name", "2"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_service_lan_vpn_profile_parcel.test", "nat_port_forwards.0.source_port", "122"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_service_lan_vpn_profile_parcel.test", "nat_port_forwards.0.translate_port", "330"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_service_lan_vpn_profile_parcel.test", "nat_port_forwards.0.source_ip", "1.2.3.4"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_service_lan_vpn_profile_parcel.test", "nat_port_forwards.0.translated_source_ip", "2.3.4.5"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_service_lan_vpn_profile_parcel.test", "nat_port_forwards.0.protocol", "TCP"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_service_lan_vpn_profile_parcel.test", "static_nats.0.nat_pool_name", "3"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_service_lan_vpn_profile_parcel.test", "static_nats.0.source_ip", "1.2.3.4"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_service_lan_vpn_profile_parcel.test", "static_nats.0.translated_source_ip", "2.3.4.5"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_service_lan_vpn_profile_parcel.test", "static_nats.0.static_nat_direction", "inside"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_service_lan_vpn_profile_parcel.test", "nat_64_v4_pools.0.name", "NATPOOL1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_service_lan_vpn_profile_parcel.test", "nat_64_v4_pools.0.range_start", "1.2.3.4"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_service_lan_vpn_profile_parcel.test", "nat_64_v4_pools.0.range_end", "2.3.4.5"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_service_lan_vpn_profile_parcel.test", "nat_64_v4_pools.0.overload", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_service_lan_vpn_profile_parcel.test", "ipv4_import_route_targets.0.route_target", "1.1.1.3:200"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_service_lan_vpn_profile_parcel.test", "ipv4_export_route_targets.0.route_target", "1.1.1.3:200"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_service_lan_vpn_profile_parcel.test", "ipv6_import_route_targets.0.route_target", "1.1.1.3:200"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_service_lan_vpn_profile_parcel.test", "ipv6_export_route_targets.0.route_target", "1.1.1.3:200"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceSdwanServiceLANVPNPrerequisitesProfileParcelConfig + testAccDataSourceSdwanServiceLANVPNProfileParcelConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAccDataSource

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites
const testAccDataSourceSdwanServiceLANVPNPrerequisitesProfileParcelConfig = `
resource "sdwan_service_feature_profile" "test" {
  name = "TF_TEST"
  description = "Terraform test"
}
`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSourceConfig
func testAccDataSourceSdwanServiceLANVPNProfileParcelConfig() string {
	config := `resource "sdwan_service_lan_vpn_profile_parcel" "test" {` + "\n"
	config += ` name = "TF_TEST"` + "\n"
	config += ` description = "Terraform integration test"` + "\n"
	config += `	feature_profile_id = sdwan_service_feature_profile.test.id` + "\n"
	config += `	vpn = 1` + "\n"
	config += `	config_description = "VPN1"` + "\n"
	config += `	omp_admin_distance_ipv4 = 1` + "\n"
	config += `	omp_admin_distance_ipv6 = 1` + "\n"
	config += `	enable_sdwan_remote_access = false` + "\n"
	config += `	primary_dns_address_ipv4 = "1.2.3.4"` + "\n"
	config += `	secondary_dns_address_ipv4 = "2.3.4.5"` + "\n"
	config += `	primary_dns_address_ipv6 = "2001:0:0:1::0"` + "\n"
	config += `	secondary_dns_address_ipv6 = "2001:0:0:2::0"` + "\n"
	config += `	host_mappings = [{` + "\n"
	config += `	  host_name = "HOSTMAPPING1"` + "\n"
	config += `	  list_of_ips = ["1.2.3.4"]` + "\n"
	config += `	}]` + "\n"
	config += `	ipv4_static_routes = [{` + "\n"
	config += `	  network_address = "1.2.3.4"` + "\n"
	config += `	  subnet_mask = "0.0.0.0"` + "\n"
	config += `	  next_hops = [{` + "\n"
	config += `		address = "1.2.3.4"` + "\n"
	config += `		administrative_distance = 1` + "\n"
	config += `	}]` + "\n"
	config += `	}]` + "\n"
	config += `	ipv6_static_routes = [{` + "\n"
	config += `	  prefix = "2001:0:0:1::0/12"` + "\n"
	config += `	  next_hops = [{` + "\n"
	config += `		address = "2001:0:0:1::0"` + "\n"
	config += `		administrative_distance = 1` + "\n"
	config += `	}]` + "\n"
	config += `	}]` + "\n"
	config += `	services = [{` + "\n"
	config += `	  service_type = "FW"` + "\n"
	config += `	  ipv4_addresses = ["1.2.3.4"]` + "\n"
	config += `	  tracking = true` + "\n"
	config += `	}]` + "\n"
	config += `	service_routes = [{` + "\n"
	config += `	  network_address = "1.2.3.4"` + "\n"
	config += `	  subnet_mask = "0.0.0.0"` + "\n"
	config += `	  service = "SIG"` + "\n"
	config += `	  vpn = 0` + "\n"
	config += `	}]` + "\n"
	config += `	gre_routes = [{` + "\n"
	config += `	  network_address = "1.2.3.4"` + "\n"
	config += `	  subnet_mask = "0.0.0.0"` + "\n"
	config += `	  interface = ["gre01"]` + "\n"
	config += `	  vpn = 0` + "\n"
	config += `	}]` + "\n"
	config += `	ipsec_routes = [{` + "\n"
	config += `	  network_address = "1.2.3.4"` + "\n"
	config += `	  subnet_mask = "0.0.0.0"` + "\n"
	config += `	  interface = ["ipsec01"]` + "\n"
	config += `	}]` + "\n"
	config += `	nat_pools = [{` + "\n"
	config += `	  nat_pool_name = 1` + "\n"
	config += `	  prefix_length = 3` + "\n"
	config += `	  range_start = "1.2.3.4"` + "\n"
	config += `	  range_end = "2.3.4.5"` + "\n"
	config += `	  overload = true` + "\n"
	config += `	  direction = "inside"` + "\n"
	config += `	}]` + "\n"
	config += `	nat_port_forwards = [{` + "\n"
	config += `	  nat_pool_name = 2` + "\n"
	config += `	  source_port = 122` + "\n"
	config += `	  translate_port = 330` + "\n"
	config += `	  source_ip = "1.2.3.4"` + "\n"
	config += `	  translated_source_ip = "2.3.4.5"` + "\n"
	config += `	  protocol = "TCP"` + "\n"
	config += `	}]` + "\n"
	config += `	static_nats = [{` + "\n"
	config += `	  nat_pool_name = 3` + "\n"
	config += `	  source_ip = "1.2.3.4"` + "\n"
	config += `	  translated_source_ip = "2.3.4.5"` + "\n"
	config += `	  static_nat_direction = "inside"` + "\n"
	config += `	}]` + "\n"
	config += `	nat_64_v4_pools = [{` + "\n"
	config += `	  name = "NATPOOL1"` + "\n"
	config += `	  range_start = "1.2.3.4"` + "\n"
	config += `	  range_end = "2.3.4.5"` + "\n"
	config += `	  overload = false` + "\n"
	config += `	}]` + "\n"
	config += `	ipv4_import_route_targets = [{` + "\n"
	config += `	  route_target = "1.1.1.3:200"` + "\n"
	config += `	}]` + "\n"
	config += `	ipv4_export_route_targets = [{` + "\n"
	config += `	  route_target = "1.1.1.3:200"` + "\n"
	config += `	}]` + "\n"
	config += `	ipv6_import_route_targets = [{` + "\n"
	config += `	  route_target = "1.1.1.3:200"` + "\n"
	config += `	}]` + "\n"
	config += `	ipv6_export_route_targets = [{` + "\n"
	config += `	  route_target = "1.1.1.3:200"` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"

	config += `
		data "sdwan_service_lan_vpn_profile_parcel" "test" {
			id = sdwan_service_lan_vpn_profile_parcel.test.id
			feature_profile_id = sdwan_service_feature_profile.test.id
		}
	`
	return config
}

// End of section. //template:end testAccDataSourceConfig
