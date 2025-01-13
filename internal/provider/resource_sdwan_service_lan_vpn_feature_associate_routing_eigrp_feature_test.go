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
func TestAccSdwanServiceLANVPNFeatureAssociateRoutingEIGRPFeature(t *testing.T) {
	if os.Getenv("SDWAN_2012") == "" {
		t.Skip("skipping test, set environment variable SDWAN_2012")
	}
	var checks []resource.TestCheckFunc
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccSdwanServiceLANVPNFeatureAssociateRoutingEIGRPFeaturePrerequisitesConfig + testAccSdwanServiceLANVPNFeatureAssociateRoutingEIGRPFeatureConfig_all(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAcc

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites
const testAccSdwanServiceLANVPNFeatureAssociateRoutingEIGRPFeaturePrerequisitesConfig = `
resource "sdwan_service_feature_profile" "test" {
  name        = "TF_TEST"
  description = "Terraform test"
}

resource "sdwan_service_lan_vpn_feature" "test" {
  name                       = "TF_TEST_SERVICE_LAN"
  description                = "Terraform test"
  feature_profile_id         = sdwan_service_feature_profile.test.id
}

resource "sdwan_service_routing_eigrp_feature" "test" {
  name                     = "TF_TEST_ROUTING_EIGRP"
  description              = "Terraform test"
  feature_profile_id         = sdwan_service_feature_profile.test.id
  autonomous_system_id = 111
  networks = [
    {
      ip_address = "100.2.2.3"
      mask       = "255.255.255.0"
    }
  ]
  hello_interval      = 5
  hold_time           = 15
  authentication_type = "md5"
}
`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAll
func testAccSdwanServiceLANVPNFeatureAssociateRoutingEIGRPFeatureConfig_all() string {
	config := `resource "sdwan_service_lan_vpn_feature_associate_routing_eigrp_feature" "test" {` + "\n"
	config += `	feature_profile_id = sdwan_service_feature_profile.test.id` + "\n"
	config += `	service_lan_vpn_feature_id = sdwan_service_lan_vpn_feature.test.id` + "\n"
	config += `	service_routing_eigrp_feature_id = sdwan_service_routing_eigrp_feature.test.id` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigAll
