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
func TestAccSdwanServiceDHCPServerProfileParcel(t *testing.T) {
	if os.Getenv("SDWAN_2015_IN_PROGRESS") == "" {
		t.Skip("skipping test, set environment variable SDWAN_2015_IN_PROGRESS")
	}
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_service_dhcp_server_feature.test", "lease_time", "86400"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_service_dhcp_server_feature.test", "interface_mtu", "65535"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_service_dhcp_server_feature.test", "domain_name", "example.com"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_service_dhcp_server_feature.test", "default_gateway", "1.2.3.4"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_service_dhcp_server_feature.test", "static_leases.0.mac_address", "01:02:03:04:05:06"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_service_dhcp_server_feature.test", "static_leases.0.ip_address", "1.2.3.4"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_service_dhcp_server_feature.test", "option_codes.0.code", "250"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccSdwanServiceDHCPServerPrerequisitesProfileParcelConfig + testAccSdwanServiceDHCPServerProfileParcelConfig_minimum(),
			},
			{
				Config: testAccSdwanServiceDHCPServerPrerequisitesProfileParcelConfig + testAccSdwanServiceDHCPServerProfileParcelConfig_all(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAcc

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites
const testAccSdwanServiceDHCPServerPrerequisitesProfileParcelConfig = `
resource "sdwan_service_feature_profile" "test" {
  name = "TF_TEST"
  description = "Terraform test"
}

`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigMinimum
func testAccSdwanServiceDHCPServerProfileParcelConfig_minimum() string {
	config := `resource "sdwan_service_dhcp_server_feature" "test" {` + "\n"
	config += ` name = "TF_TEST_MIN"` + "\n"
	config += ` description = "Terraform integration test"` + "\n"
	config += `	feature_profile_id = sdwan_service_feature_profile.test.id` + "\n"
	config += `	network_address = "1.2.3.4"` + "\n"
	config += `	subnet_mask = "255.255.255.0"` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigMinimum

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAll
func testAccSdwanServiceDHCPServerProfileParcelConfig_all() string {
	config := `resource "sdwan_service_dhcp_server_feature" "test" {` + "\n"
	config += ` name = "TF_TEST_ALL"` + "\n"
	config += ` description = "Terraform integration test"` + "\n"
	config += `	feature_profile_id = sdwan_service_feature_profile.test.id` + "\n"
	config += `	network_address = "1.2.3.4"` + "\n"
	config += `	subnet_mask = "255.255.255.0"` + "\n"
	config += `	exclude = ["192.168.1.1"]` + "\n"
	config += `	lease_time = 86400` + "\n"
	config += `	interface_mtu = 65535` + "\n"
	config += `	domain_name = "example.com"` + "\n"
	config += `	default_gateway = "1.2.3.4"` + "\n"
	config += `	dns_servers = ["8.8.8.8"]` + "\n"
	config += `	tftp_servers = ["1.1.1.1"]` + "\n"
	config += `	static_leases = [{` + "\n"
	config += `	  mac_address = "01:02:03:04:05:06"` + "\n"
	config += `	  ip_address = "1.2.3.4"` + "\n"
	config += `	}]` + "\n"
	config += `	option_codes = [{` + "\n"
	config += `	  code = 250` + "\n"
	config += `	  ascii = "example"` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigAll
