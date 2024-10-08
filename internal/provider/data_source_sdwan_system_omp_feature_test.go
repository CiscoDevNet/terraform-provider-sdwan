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
func TestAccDataSourceSdwanSystemOMPProfileParcel(t *testing.T) {
	if os.Getenv("SDWAN_2012") == "" {
		t.Skip("skipping test, set environment variable SDWAN_2012")
	}
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_system_omp_feature.test", "graceful_restart", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_system_omp_feature.test", "overlay_as", "10"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_system_omp_feature.test", "paths_advertised_per_prefix", "4"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_system_omp_feature.test", "ecmp_limit", "4"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_system_omp_feature.test", "shutdown", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_system_omp_feature.test", "omp_admin_distance_ipv4", "10"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_system_omp_feature.test", "omp_admin_distance_ipv6", "20"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_system_omp_feature.test", "advertisement_interval", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_system_omp_feature.test", "graceful_restart_timer", "43200"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_system_omp_feature.test", "eor_timer", "300"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_system_omp_feature.test", "holdtime", "60"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_system_omp_feature.test", "advertise_ipv4_bgp", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_system_omp_feature.test", "advertise_ipv4_ospf", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_system_omp_feature.test", "advertise_ipv4_ospf_v3", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_system_omp_feature.test", "advertise_ipv4_connected", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_system_omp_feature.test", "advertise_ipv4_static", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_system_omp_feature.test", "advertise_ipv4_eigrp", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_system_omp_feature.test", "advertise_ipv4_lisp", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_system_omp_feature.test", "advertise_ipv4_isis", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_system_omp_feature.test", "advertise_ipv6_bgp", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_system_omp_feature.test", "advertise_ipv6_ospf", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_system_omp_feature.test", "advertise_ipv6_connected", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_system_omp_feature.test", "advertise_ipv6_static", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_system_omp_feature.test", "advertise_ipv6_eigrp", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_system_omp_feature.test", "advertise_ipv6_lisp", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_system_omp_feature.test", "advertise_ipv6_isis", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_system_omp_feature.test", "ignore_region_path_length", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_system_omp_feature.test", "transport_gateway", "prefer"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceSdwanSystemOMPPrerequisitesProfileParcelConfig + testAccDataSourceSdwanSystemOMPProfileParcelConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAccDataSource

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites
const testAccDataSourceSdwanSystemOMPPrerequisitesProfileParcelConfig = `
resource "sdwan_system_feature_profile" "test" {
  name = "TF_TEST"
  description = "Terraform test"
}
`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSourceConfig
func testAccDataSourceSdwanSystemOMPProfileParcelConfig() string {
	config := `resource "sdwan_system_omp_feature" "test" {` + "\n"
	config += ` name = "TF_TEST"` + "\n"
	config += ` description = "Terraform integration test"` + "\n"
	config += `	feature_profile_id = sdwan_system_feature_profile.test.id` + "\n"
	config += `	graceful_restart = true` + "\n"
	config += `	overlay_as = 10` + "\n"
	config += `	paths_advertised_per_prefix = 4` + "\n"
	config += `	ecmp_limit = 4` + "\n"
	config += `	shutdown = false` + "\n"
	config += `	omp_admin_distance_ipv4 = 10` + "\n"
	config += `	omp_admin_distance_ipv6 = 20` + "\n"
	config += `	advertisement_interval = 1` + "\n"
	config += `	graceful_restart_timer = 43200` + "\n"
	config += `	eor_timer = 300` + "\n"
	config += `	holdtime = 60` + "\n"
	config += `	advertise_ipv4_bgp = false` + "\n"
	config += `	advertise_ipv4_ospf = false` + "\n"
	config += `	advertise_ipv4_ospf_v3 = false` + "\n"
	config += `	advertise_ipv4_connected = false` + "\n"
	config += `	advertise_ipv4_static = false` + "\n"
	config += `	advertise_ipv4_eigrp = false` + "\n"
	config += `	advertise_ipv4_lisp = false` + "\n"
	config += `	advertise_ipv4_isis = false` + "\n"
	config += `	advertise_ipv6_bgp = true` + "\n"
	config += `	advertise_ipv6_ospf = true` + "\n"
	config += `	advertise_ipv6_connected = true` + "\n"
	config += `	advertise_ipv6_static = true` + "\n"
	config += `	advertise_ipv6_eigrp = true` + "\n"
	config += `	advertise_ipv6_lisp = true` + "\n"
	config += `	advertise_ipv6_isis = true` + "\n"
	config += `	ignore_region_path_length = false` + "\n"
	config += `	transport_gateway = "prefer"` + "\n"
	config += `	site_types = ["type-1"]` + "\n"
	config += `}` + "\n"

	config += `
		data "sdwan_system_omp_feature" "test" {
			id = sdwan_system_omp_feature.test.id
			feature_profile_id = sdwan_system_feature_profile.test.id
		}
	`
	return config
}

// End of section. //template:end testAccDataSourceConfig
