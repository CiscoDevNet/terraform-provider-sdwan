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
func TestAccSdwanCiscoOMPFeatureTemplate(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_omp_feature_template.test", "graceful_restart", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_omp_feature_template.test", "overlay_as", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_omp_feature_template.test", "send_path_limit", "4"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_omp_feature_template.test", "ecmp_limit", "4"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_omp_feature_template.test", "shutdown", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_omp_feature_template.test", "omp_admin_distance_ipv4", "10"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_omp_feature_template.test", "omp_admin_distance_ipv6", "10"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_omp_feature_template.test", "advertisement_interval", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_omp_feature_template.test", "graceful_restart_timer", "43200"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_omp_feature_template.test", "eor_timer", "300"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_omp_feature_template.test", "holdtime", "60"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_omp_feature_template.test", "ignore_region_path_length", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_omp_feature_template.test", "transport_gateway", "prefer"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_omp_feature_template.test", "advertise_ipv4_routes.0.protocol", "ospf"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_omp_feature_template.test", "advertise_ipv4_routes.0.advertise_external_ospf", "external"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_omp_feature_template.test", "advertise_ipv6_routes.0.protocol", "ospf"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccSdwanCiscoOMPFeatureTemplateConfig_minimum(),
			},
			{
				Config: testAccSdwanCiscoOMPFeatureTemplateConfig_all(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAcc

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites
// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigMinimum
func testAccSdwanCiscoOMPFeatureTemplateConfig_minimum() string {
	config := `resource "sdwan_cisco_omp_feature_template" "test" {` + "\n"
	config += ` name = "TF_TEST_MIN"` + "\n"
	config += ` description = "Terraform integration test"` + "\n"
	config += ` device_types = ["vedge-C8000V"]` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigMinimum

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAll
func testAccSdwanCiscoOMPFeatureTemplateConfig_all() string {
	config := `resource "sdwan_cisco_omp_feature_template" "test" {` + "\n"
	config += ` name = "TF_TEST_ALL"` + "\n"
	config += ` description = "Terraform integration test"` + "\n"
	config += ` device_types = ["vedge-C8000V"]` + "\n"
	config += `	graceful_restart = true` + "\n"
	config += `	overlay_as = 1` + "\n"
	config += `	send_path_limit = 4` + "\n"
	config += `	ecmp_limit = 4` + "\n"
	config += `	shutdown = false` + "\n"
	config += `	omp_admin_distance_ipv4 = 10` + "\n"
	config += `	omp_admin_distance_ipv6 = 10` + "\n"
	config += `	advertisement_interval = 1` + "\n"
	config += `	graceful_restart_timer = 43200` + "\n"
	config += `	eor_timer = 300` + "\n"
	config += `	holdtime = 60` + "\n"
	config += `	ignore_region_path_length = false` + "\n"
	config += `	transport_gateway = "prefer"` + "\n"
	config += `	advertise_ipv4_routes = [{` + "\n"
	config += `	  protocol = "ospf"` + "\n"
	config += `	  advertise_external_ospf = "external"` + "\n"
	config += `	}]` + "\n"
	config += `	advertise_ipv6_routes = [{` + "\n"
	config += `	  protocol = "ospf"` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigAll
