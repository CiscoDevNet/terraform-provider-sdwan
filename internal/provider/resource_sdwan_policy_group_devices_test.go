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
func TestAccSdwanPolicyGroupDevices(t *testing.T) {
	if os.Getenv("SDWAN_2015") == "" && os.Getenv("SDWAN_2018") == "" {
		t.Skip("skipping test, set environment variable SDWAN_2015 or SDWAN_2018")
	}
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_policy_group_devices.test", "solution", "sdwan"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccSdwanPolicyGroupDevicesPrerequisitesConfig + testAccSdwanPolicyGroupDevicesConfig_all(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAcc

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites
const testAccSdwanPolicyGroupDevicesPrerequisitesConfig = `
resource "sdwan_system_feature_profile" "test" {
  name = "SYSTEM_PGD_TF"
  description = "Terraform test"
}

resource "sdwan_system_basic_feature" "test" {
  name = "BASIC_TF"
  feature_profile_id = sdwan_system_feature_profile.test.id
}

resource "sdwan_system_aaa_feature" "test" {
  name               = "AAA_TF"
  feature_profile_id = sdwan_system_feature_profile.test.id
  server_auth_order  = ["local"]
  users = [{
    name     = "admin"
    password = "admin"
  }]
}

resource "sdwan_system_bfd_feature" "test" {
  name               = "BFD_TF"
  feature_profile_id = sdwan_system_feature_profile.test.id
}

resource "sdwan_system_global_feature" "test" {
  name               = "GLOBAL_TF"
  feature_profile_id = sdwan_system_feature_profile.test.id
}

resource "sdwan_system_logging_feature" "test" {
  name               = "LOGGING_TF"
  feature_profile_id = sdwan_system_feature_profile.test.id
}

resource "sdwan_system_omp_feature" "test" {
  name               = "OMP_TF"
  feature_profile_id = sdwan_system_feature_profile.test.id
}

resource "sdwan_transport_feature_profile" "test" {
  name        = "TRANSPORT_PGD_TF"
  description = "My transport feature profile 1"
}

resource "sdwan_transport_wan_vpn_feature" "test" {
  name               = "WAN_VPN_TF"
  feature_profile_id = sdwan_transport_feature_profile.test.id
  vpn                = 0
}

resource "sdwan_transport_wan_vpn_interface_ethernet_feature" "test" {
  name                         = "WAN_VPN_INT_TF"
  feature_profile_id           = sdwan_transport_feature_profile.test.id
  transport_wan_vpn_feature_id = sdwan_transport_wan_vpn_feature.test.id
  interface_name               = "GigabitEthernet1"
  shutdown                     = false
  ipv4_address_type            = "dynamic"
  ipv4_dhcp_distance           = 1
  tunnel_interface             = true
  tunnel_interface_encapsulations = [
    {
      encapsulation = "ipsec"
    }
  ]
}

resource "sdwan_service_feature_profile" "test" {
  name        = "SERVICE_PGD_TF"
  description = "Terraform test"
}

resource "sdwan_service_lan_vpn_feature" "test" {
  name                = "SERVICE_VPN_TF"
  feature_profile_id  = sdwan_service_feature_profile.test.id
  vpn                 = 1
}

resource "sdwan_configuration_group" "test" {
  name        = "CG_PGD_TF"
  description = "My config group 1"
  solution    = "sdwan"
  feature_profile_ids = [
    sdwan_system_feature_profile.test.id,
    sdwan_transport_feature_profile.test.id,
    sdwan_service_feature_profile.test.id,
  ]
  feature_versions = [
    sdwan_system_basic_feature.test.version,
    sdwan_system_aaa_feature.test.version,
    sdwan_system_bfd_feature.test.version,
    sdwan_system_global_feature.test.version,
    sdwan_system_logging_feature.test.version,
    sdwan_system_omp_feature.test.version,
    sdwan_transport_wan_vpn_interface_ethernet_feature.test.version,
    sdwan_service_lan_vpn_feature.test.version,
  ]
}

resource "sdwan_configuration_group_devices" "test" {
  configuration_group_id      = sdwan_configuration_group.test.id
  solution                    = "sdwan"
  configuration_group_version = sdwan_configuration_group.test.version
  devices = [{
    id     = "C8K-40C0CCFD-9EA8-2B2E-E73B-32C5924EC79B"
    deploy = true
    variables = [
      { name = "host_name", value = "edge1" },
      { name = "pseudo_commit_timer", value = 0 },
      { name = "site_id", value = 1 },
      { name = "system_ip", value = "10.1.1.1" },
      { name = "ipv6_strict_control", value = "false" },
    ]
  }]
}

resource "sdwan_application_priority_feature_profile" "test" {
  name        = "APPLICATION_PRIORITY_TF"
  description = "Terraform test"
}

resource "sdwan_application_priority_qos_policy" "test" {
  name                       = "qos"
  description                = "QoS policy for application priority"
  feature_profile_id         = sdwan_application_priority_feature_profile.test.id
  target_interfaces_variable = "{{qos_interfaces}}"
}

resource "sdwan_policy_group" "test" {
  name                = "PG_PGD_TF"
  description         = "My policy group 1"
  solution            = "sdwan"
  feature_profile_ids = [sdwan_application_priority_feature_profile.test.id]
  policy_versions     = [sdwan_application_priority_qos_policy.test.version]
}

`

// End of section. //template:end testPrerequisites

func testAccSdwanPolicyGroupDevicesConfig_all() string {
	config := `resource "sdwan_policy_group_devices" "test" {` + "\n"
	config += `	policy_group_id = sdwan_policy_group.test.id` + "\n"
	config += `	solution = "sdwan"` + "\n"
	config += `	policy_group_version = sdwan_policy_group.test.version` + "\n"
	config += `	devices = [{` + "\n"
	config += `	  id = "C8K-40C0CCFD-9EA8-2B2E-E73B-32C5924EC79B"` + "\n"
	config += `	  deploy = true` + "\n"
	config += `	  variables = [` + "\n"
	config += `	    {` + "\n"
	config += `	      name = "qos_interfaces"` + "\n"
	config += `	      list_value = [` + "\n"
	config += `	        "GigabitEthernet1",` + "\n"
	config += `	        "GigabitEthernet2"` + "\n"
	config += `	      ]` + "\n"
	config += `	    },` + "\n"
	config += `	  ]` + "\n"
	config += `	}]` + "\n"
	config += `	depends_on = [sdwan_configuration_group_devices.test]` + "\n"
	config += `}` + "\n"
	return config
}
