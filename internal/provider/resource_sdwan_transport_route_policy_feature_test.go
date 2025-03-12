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
func TestAccSdwanTransportRoutePolicyProfileParcel(t *testing.T) {
	if os.Getenv("SDWAN_2012") == "" {
		t.Skip("skipping test, set environment variable SDWAN_2012")
	}
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_transport_route_policy_feature.test", "default_action", "accept"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_transport_route_policy_feature.test", "sequences.0.id", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_transport_route_policy_feature.test", "sequences.0.name", "SEQ_1"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_transport_route_policy_feature.test", "sequences.0.base_action", "reject"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_transport_route_policy_feature.test", "sequences.0.protocol", "IPV4"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_transport_route_policy_feature.test", "sequences.0.actions.0.community_additive", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_transport_route_policy_feature.test", "sequences.0.actions.0.local_preference", "100"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_transport_route_policy_feature.test", "sequences.0.actions.0.metric", "20"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_transport_route_policy_feature.test", "sequences.0.actions.0.metric_type", "type1"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_transport_route_policy_feature.test", "sequences.0.actions.0.omp_tag", "200"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_transport_route_policy_feature.test", "sequences.0.actions.0.origin", "EGP"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_transport_route_policy_feature.test", "sequences.0.actions.0.ospf_tag", "1200"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_transport_route_policy_feature.test", "sequences.0.actions.0.weight", "2200"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_transport_route_policy_feature.test", "sequences.0.actions.0.ipv4_next_hop", "10.0.0.1"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccSdwanTransportRoutePolicyPrerequisitesProfileParcelConfig + testAccSdwanTransportRoutePolicyProfileParcelConfig_minimum(),
			},
			{
				Config: testAccSdwanTransportRoutePolicyPrerequisitesProfileParcelConfig + testAccSdwanTransportRoutePolicyProfileParcelConfig_all(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAcc

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites
const testAccSdwanTransportRoutePolicyPrerequisitesProfileParcelConfig = `
resource "sdwan_transport_feature_profile" "test" {
  name = "TF_TEST"
  description = "Terraform test"
}
`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigMinimum
func testAccSdwanTransportRoutePolicyProfileParcelConfig_minimum() string {
	config := `resource "sdwan_transport_route_policy_feature" "test" {` + "\n"
	config += ` name = "TF_TEST_MIN"` + "\n"
	config += ` description = "Terraform integration test"` + "\n"
	config += `	feature_profile_id = sdwan_transport_feature_profile.test.id` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigMinimum

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAll
func testAccSdwanTransportRoutePolicyProfileParcelConfig_all() string {
	config := `resource "sdwan_transport_route_policy_feature" "test" {` + "\n"
	config += ` name = "TF_TEST_ALL"` + "\n"
	config += ` description = "Terraform integration test"` + "\n"
	config += `	feature_profile_id = sdwan_transport_feature_profile.test.id` + "\n"
	config += `	default_action = "accept"` + "\n"
	config += `	sequences = [{` + "\n"
	config += `	  id = 1` + "\n"
	config += `	  name = "SEQ_1"` + "\n"
	config += `	  base_action = "reject"` + "\n"
	config += `	  protocol = "IPV4"` + "\n"
	config += `	  actions = [{` + "\n"
	config += `		as_path_prepend = ["65521"]` + "\n"
	config += `		community_additive = false` + "\n"
	config += `		community = ["internet"]` + "\n"
	config += `		local_preference = 100` + "\n"
	config += `		metric = 20` + "\n"
	config += `		metric_type = "type1"` + "\n"
	config += `		omp_tag = 200` + "\n"
	config += `		origin = "EGP"` + "\n"
	config += `		ospf_tag = 1200` + "\n"
	config += `		weight = 2200` + "\n"
	config += `		ipv4_next_hop = "10.0.0.1"` + "\n"
	config += `	}]` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigAll
