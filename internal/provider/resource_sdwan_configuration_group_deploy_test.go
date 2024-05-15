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
func TestAccSdwanConfigurationGroupDeploy(t *testing.T) {
	if os.Getenv("CONFIG_GROUP_DEPLOY") == "" {
		t.Skip("skipping test, set environment variable CONFIG_GROUP_DEPLOY")
	}
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_configuration_group_deploy.test", "devices.0.id", "C8K-15411CCC-D476-0B3B-21F2-5D6AC387EE7B"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccSdwanConfigurationGroupDeployPrerequisitesConfig + testAccSdwanConfigurationGroupDeployConfig_all(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAcc

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites
const testAccSdwanConfigurationGroupDeployPrerequisitesConfig = `
resource "sdwan_system_feature_profile" "test" {
  name        = "SYSTEM_TF"
  description = "My system feature profile 1"
}

resource "sdwan_system_basic_profile_parcel" "test" {
  name = "BASIC_TF"
  feature_profile_id = sdwan_system_feature_profile.test.id
}

resource "sdwan_system_aaa_profile_parcel" "test" {
  name               = "AAA_TF"
  feature_profile_id = sdwan_system_feature_profile.test.id
  server_auth_order  = ["local"]
  users = [{
    name     = "admin"
    password = "admin"
  }]
}

resource "sdwan_system_bfd_profile_parcel" "test" {
  name               = "BFD_TF"
  feature_profile_id = sdwan_system_feature_profile.test.id
}

resource "sdwan_system_global_profile_parcel" "test" {
  name               = "GLOBAL_TF"
  feature_profile_id = sdwan_system_feature_profile.test.id
}

resource "sdwan_system_logging_profile_parcel" "test" {
  name               = "LOGGING_TF"
  feature_profile_id = sdwan_system_feature_profile.test.id
}

resource "sdwan_transport_feature_profile" "test" {
  name        = "TRANSPORT_TF"
  description = "My transport feature profile 1"
}

resource "sdwan_transport_wan_vpn_profile_parcel" "test" {
  name               = "WAN_VPN_TF"
  feature_profile_id = sdwan_transport_feature_profile.test.id
  vpn                = 0
}

resource "sdwan_transport_wan_vpn_interface_ethernet_profile_parcel" "test" {
  name                                = "WAN_VPN_INT_TF"
  feature_profile_id                  = sdwan_transport_feature_profile.test.id
  transport_wan_vpn_profile_parcel_id = sdwan_transport_wan_vpn_profile_parcel.test.id
  interface_name                      = "GigabitEthernet1"
  ipv4_dhcp_distance                  = 1
  tunnel_interface                    = true
  tunnel_interface_encapsulations = [
    {
      encapsulation = "ipsec"
    }
  ]
}

resource "sdwan_configuration_group" "test" {
  name        = "TF_TEST"
  description = "Terraform test"
  solution    = "sdwan"
  feature_profiles = [
    {
      id = sdwan_system_feature_profile.test.id
    },
    {
      id = sdwan_transport_feature_profile.test.id
    }
  ]

  depends_on = [
    sdwan_system_basic_profile_parcel.test,
    sdwan_system_aaa_profile_parcel.test,
    sdwan_system_bfd_profile_parcel.test,
    sdwan_system_global_profile_parcel.test,
    sdwan_system_logging_profile_parcel.test,
    sdwan_transport_wan_vpn_interface_ethernet_profile_parcel.test
  ]
}

resource "sdwan_configuration_group_devices" "test" {
  configuration_group_id = sdwan_configuration_group.test.id
  devices = [{
    id = "C8K-15411CCC-D476-0B3B-21F2-5D6AC387EE7B"
  }]
}

resource "sdwan_configuration_group_device_variables" "test" {
  configuration_group_id = sdwan_configuration_group_devices.test.configuration_group_id
  solution = "sdwan"
  devices = [{
    device_id = "C8K-15411CCC-D476-0B3B-21F2-5D6AC387EE7B"
    variables = [
      {
        name  = "host_name"
        value = "edge1"
      },
      {
        name  = "pseudo_commit_timer"
        value = "0"
      },
      {
        name  = "site_id"
        value = "1"
      },
      {
        name  = "system_ip"
        value = "10.1.1.1"
      },
    ]
  }]
}
`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAll
func testAccSdwanConfigurationGroupDeployConfig_all() string {
	config := `resource "sdwan_configuration_group_deploy" "test" {` + "\n"
	config += `	configuration_group_id = sdwan_configuration_group_device_variables.test.configuration_group_id` + "\n"
	config += `	devices = [{` + "\n"
	config += `	  id = "C8K-15411CCC-D476-0B3B-21F2-5D6AC387EE7B"` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigAll
