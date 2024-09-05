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
func TestAccSdwanTransportManagementVPNProfileParcel(t *testing.T) {
	if os.Getenv("SDWAN_2012") == "" {
		t.Skip("skipping test, set environment variable SDWAN_2012")
	}
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_transport_management_vpn_feature.test", "vpn_description", "example"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_transport_management_vpn_feature.test", "primary_dns_address_ipv4", "1.2.3.4"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_transport_management_vpn_feature.test", "secondary_dns_address_ipv4", "2.3.4.5"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_transport_management_vpn_feature.test", "primary_dns_address_ipv6", "2001:0:0:1::0"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_transport_management_vpn_feature.test", "secondary_dns_address_ipv6", "2001:0:0:2::0"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_transport_management_vpn_feature.test", "new_host_mappings.0.host_name", "example"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_transport_management_vpn_feature.test", "ipv4_static_routes.0.network_address", "1.2.3.4"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_transport_management_vpn_feature.test", "ipv4_static_routes.0.subnet_mask", "0.0.0.0"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_transport_management_vpn_feature.test", "ipv4_static_routes.0.gateway", "nextHop"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_transport_management_vpn_feature.test", "ipv4_static_routes.0.next_hops.0.address", "1.2.3.4"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_transport_management_vpn_feature.test", "ipv4_static_routes.0.next_hops.0.administrative_distance", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_transport_management_vpn_feature.test", "ipv4_static_routes.0.administrative_distance", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_transport_management_vpn_feature.test", "ipv6_static_routes.0.prefix", "2002::/16"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_transport_management_vpn_feature.test", "ipv6_static_routes.0.next_hops.0.address", "2001:0:0:1::1"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_transport_management_vpn_feature.test", "ipv6_static_routes.0.next_hops.0.administrative_distance", "1"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccSdwanTransportManagementVPNPrerequisitesProfileParcelConfig + testAccSdwanTransportManagementVPNProfileParcelConfig_minimum(),
			},
			{
				Config: testAccSdwanTransportManagementVPNPrerequisitesProfileParcelConfig + testAccSdwanTransportManagementVPNProfileParcelConfig_all(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAcc

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites
const testAccSdwanTransportManagementVPNPrerequisitesProfileParcelConfig = `
resource "sdwan_transport_feature_profile" "test" {
  name = "TF_TEST"
  description = "Terraform test"
}

`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigMinimum
func testAccSdwanTransportManagementVPNProfileParcelConfig_minimum() string {
	config := `resource "sdwan_transport_management_vpn_feature" "test" {` + "\n"
	config += ` name = "TF_TEST_MIN"` + "\n"
	config += ` description = "Terraform integration test"` + "\n"
	config += `	feature_profile_id = sdwan_transport_feature_profile.test.id` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigMinimum

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAll
func testAccSdwanTransportManagementVPNProfileParcelConfig_all() string {
	config := `resource "sdwan_transport_management_vpn_feature" "test" {` + "\n"
	config += ` name = "TF_TEST_ALL"` + "\n"
	config += ` description = "Terraform integration test"` + "\n"
	config += `	feature_profile_id = sdwan_transport_feature_profile.test.id` + "\n"
	config += `	vpn_description = "example"` + "\n"
	config += `	primary_dns_address_ipv4 = "1.2.3.4"` + "\n"
	config += `	secondary_dns_address_ipv4 = "2.3.4.5"` + "\n"
	config += `	primary_dns_address_ipv6 = "2001:0:0:1::0"` + "\n"
	config += `	secondary_dns_address_ipv6 = "2001:0:0:2::0"` + "\n"
	config += `	new_host_mappings = [{` + "\n"
	config += `	  host_name = "example"` + "\n"
	config += `	  list_of_ip_addresses = ["1.2.3.4"]` + "\n"
	config += `	}]` + "\n"
	config += `	ipv4_static_routes = [{` + "\n"
	config += `	  network_address = "1.2.3.4"` + "\n"
	config += `	  subnet_mask = "0.0.0.0"` + "\n"
	config += `	  gateway = "nextHop"` + "\n"
	config += `	  next_hops = [{` + "\n"
	config += `		address = "1.2.3.4"` + "\n"
	config += `		administrative_distance = 1` + "\n"
	config += `	}]` + "\n"
	config += `	  administrative_distance = 1` + "\n"
	config += `	}]` + "\n"
	config += `	ipv6_static_routes = [{` + "\n"
	config += `	  prefix = "2002::/16"` + "\n"
	config += `	  next_hops = [{` + "\n"
	config += `		address = "2001:0:0:1::1"` + "\n"
	config += `		administrative_distance = 1` + "\n"
	config += `	}]` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigAll
