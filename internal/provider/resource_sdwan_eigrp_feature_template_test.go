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
func TestAccSdwanEigrpFeatureTemplate(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_eigrp_feature_template.test", "as_number", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_eigrp_feature_template.test", "address_families.0.type", "ipv4"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_eigrp_feature_template.test", "address_families.0.redistributes.0.protocol", "bgp"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_eigrp_feature_template.test", "address_families.0.redistributes.0.route_policy", "1.2.3.4"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_eigrp_feature_template.test", "address_families.0.networks.0.prefix", "1.2.3.4/24"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_eigrp_feature_template.test", "hello_interval", "5"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_eigrp_feature_template.test", "hold_time", "15"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_eigrp_feature_template.test", "route_policy_name", "RP1"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_eigrp_feature_template.test", "filter", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_eigrp_feature_template.test", "authentication_type", "hmac-sha-256"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_eigrp_feature_template.test", "hmac_authentication_key", "myAuthKey"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_eigrp_feature_template.test", "interfaces.0.interface_name", "Ethernet1"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_eigrp_feature_template.test", "interfaces.0.shutdown", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_eigrp_feature_template.test", "interfaces.0.summary_addresses.0.prefix", "1.2.3.4/24"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccSdwanEigrpFeatureTemplateConfig_minimum(),
			},
			{
				Config: testAccSdwanEigrpFeatureTemplateConfig_all(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAcc

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites
// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigMinimum
func testAccSdwanEigrpFeatureTemplateConfig_minimum() string {
	config := `resource "sdwan_eigrp_feature_template" "test" {` + "\n"
	config += ` name = "TF_TEST_MIN"` + "\n"
	config += ` description = "Terraform integration test"` + "\n"
	config += ` device_types = ["vedge-C8000V"]` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigMinimum

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAll
func testAccSdwanEigrpFeatureTemplateConfig_all() string {
	config := `resource "sdwan_eigrp_feature_template" "test" {` + "\n"
	config += ` name = "TF_TEST_ALL"` + "\n"
	config += ` description = "Terraform integration test"` + "\n"
	config += ` device_types = ["vedge-C8000V"]` + "\n"
	config += `	as_number = 1` + "\n"
	config += `	address_families = [{` + "\n"
	config += `	  type = "ipv4"` + "\n"
	config += `	  redistributes = [{` + "\n"
	config += `		protocol = "bgp"` + "\n"
	config += `		route_policy = "1.2.3.4"` + "\n"
	config += `	}]` + "\n"
	config += `	  networks = [{` + "\n"
	config += `		prefix = "1.2.3.4/24"` + "\n"
	config += `	}]` + "\n"
	config += `	}]` + "\n"
	config += `	hello_interval = 5` + "\n"
	config += `	hold_time = 15` + "\n"
	config += `	route_policy_name = "RP1"` + "\n"
	config += `	filter = false` + "\n"
	config += `	authentication_type = "hmac-sha-256"` + "\n"
	config += `	hmac_authentication_key = "myAuthKey"` + "\n"
	config += `	interfaces = [{` + "\n"
	config += `	  interface_name = "Ethernet1"` + "\n"
	config += `	  shutdown = false` + "\n"
	config += `	  summary_addresses = [{` + "\n"
	config += `		prefix = "1.2.3.4/24"` + "\n"
	config += `	}]` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigAll
