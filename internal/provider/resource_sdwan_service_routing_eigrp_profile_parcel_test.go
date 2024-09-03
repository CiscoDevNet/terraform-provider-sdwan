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
func TestAccSdwanServiceRoutingEIGRPProfileParcel(t *testing.T) {
	if os.Getenv("SDWAN_2012") == "" {
		t.Skip("skipping test, set environment variable SDWAN_2012")
	}
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_service_routing_eigrp_profile_parcel.test", "autonomous_system_id", "111"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_service_routing_eigrp_profile_parcel.test", "networks.0.ip_address", "100.2.2.3"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_service_routing_eigrp_profile_parcel.test", "networks.0.mask", "255.255.255.0"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_service_routing_eigrp_profile_parcel.test", "hello_interval", "5"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_service_routing_eigrp_profile_parcel.test", "hold_time", "15"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_service_routing_eigrp_profile_parcel.test", "type", "md5"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_service_routing_eigrp_profile_parcel.test", "md5_keys.0.key_id", "2"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_service_routing_eigrp_profile_parcel.test", "af_interfaces.0.name", "GigabitEthernet3"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_service_routing_eigrp_profile_parcel.test", "af_interfaces.0.shutdown", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_service_routing_eigrp_profile_parcel.test", "af_interfaces.0.summary_addresses.0.address", "10.0.0.1"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_service_routing_eigrp_profile_parcel.test", "af_interfaces.0.summary_addresses.0.mask", "255.255.255.0"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_service_routing_eigrp_profile_parcel.test", "filter", "false"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccSdwanServiceRoutingEIGRPPrerequisitesProfileParcelConfig + testAccSdwanServiceRoutingEIGRPProfileParcelConfig_minimum(),
			},
			{
				Config: testAccSdwanServiceRoutingEIGRPPrerequisitesProfileParcelConfig + testAccSdwanServiceRoutingEIGRPProfileParcelConfig_all(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAcc

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites
const testAccSdwanServiceRoutingEIGRPPrerequisitesProfileParcelConfig = `
resource "sdwan_service_feature_profile" "test" {
  name = "TF_TEST"
  description = "Terraform test"
}

`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigMinimum
func testAccSdwanServiceRoutingEIGRPProfileParcelConfig_minimum() string {
	config := `resource "sdwan_service_routing_eigrp_profile_parcel" "test" {` + "\n"
	config += ` name = "TF_TEST_MIN"` + "\n"
	config += ` description = "Terraform integration test"` + "\n"
	config += `	feature_profile_id = sdwan_service_feature_profile.test.id` + "\n"
	config += `	autonomous_system_id = 111` + "\n"
	config += `	networks = [{` + "\n"
	config += `	  ip_address = "100.2.2.3"` + "\n"
	config += `	  mask = "255.255.255.0"` + "\n"
	config += `	}]` + "\n"
	config += `	hello_interval = 5` + "\n"
	config += `	hold_time = 15` + "\n"
	config += `	type = "md5"` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigMinimum

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAll
func testAccSdwanServiceRoutingEIGRPProfileParcelConfig_all() string {
	config := `resource "sdwan_service_routing_eigrp_profile_parcel" "test" {` + "\n"
	config += ` name = "TF_TEST_ALL"` + "\n"
	config += ` description = "Terraform integration test"` + "\n"
	config += `	feature_profile_id = sdwan_service_feature_profile.test.id` + "\n"
	config += `	autonomous_system_id = 111` + "\n"
	config += `	networks = [{` + "\n"
	config += `	  ip_address = "100.2.2.3"` + "\n"
	config += `	  mask = "255.255.255.0"` + "\n"
	config += `	}]` + "\n"
	config += `	hello_interval = 5` + "\n"
	config += `	hold_time = 15` + "\n"
	config += `	type = "md5"` + "\n"
	config += `	md5_keys = [{` + "\n"
	config += `	  key_id = 2` + "\n"
	config += `	  authentication_key = "password123"` + "\n"
	config += `	}]` + "\n"
	config += `	af_interfaces = [{` + "\n"
	config += `	  name = "GigabitEthernet3"` + "\n"
	config += `	  shutdown = false` + "\n"
	config += `	  summary_addresses = [{` + "\n"
	config += `		address = "10.0.0.1"` + "\n"
	config += `		mask = "255.255.255.0"` + "\n"
	config += `	}]` + "\n"
	config += `	}]` + "\n"
	config += `	filter = false` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigAll
