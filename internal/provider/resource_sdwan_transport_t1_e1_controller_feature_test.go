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
func TestAccSdwanTransportT1E1ControllerProfileParcel(t *testing.T) {
	if os.Getenv("SDWAN_2012") == "" {
		t.Skip("skipping test, set environment variable SDWAN_2012")
	}
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_transport_t1_e1_controller_feature.test", "type", "t1"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_transport_t1_e1_controller_feature.test", "entries.0.t1_description", "T1"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_transport_t1_e1_controller_feature.test", "entries.0.t1_framing", "esf"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_transport_t1_e1_controller_feature.test", "entries.0.t1_linecode", "ami"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_transport_t1_e1_controller_feature.test", "entries.0.cable_length", "long"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_transport_t1_e1_controller_feature.test", "entries.0.length_long", "-7.5db"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_transport_t1_e1_controller_feature.test", "entries.0.clock_source", "line"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_transport_t1_e1_controller_feature.test", "entries.0.line_mode", "primary"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_transport_t1_e1_controller_feature.test", "entries.0.description", "desc"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_transport_t1_e1_controller_feature.test", "entries.0.channel_groups.0.channel_group", "12"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_transport_t1_e1_controller_feature.test", "entries.0.channel_groups.0.time_slot", "timeslots 15"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccSdwanTransportT1E1ControllerPrerequisitesProfileParcelConfig + testAccSdwanTransportT1E1ControllerProfileParcelConfig_minimum(),
			},
			{
				Config: testAccSdwanTransportT1E1ControllerPrerequisitesProfileParcelConfig + testAccSdwanTransportT1E1ControllerProfileParcelConfig_all(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAcc

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites
const testAccSdwanTransportT1E1ControllerPrerequisitesProfileParcelConfig = `
resource "sdwan_transport_feature_profile" "test" {
  name = "TF_TEST"
  description = "Terraform test"
}
`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigMinimum
func testAccSdwanTransportT1E1ControllerProfileParcelConfig_minimum() string {
	config := `resource "sdwan_transport_t1_e1_controller_feature" "test" {` + "\n"
	config += ` name = "TF_TEST_MIN"` + "\n"
	config += ` description = "Terraform integration test"` + "\n"
	config += `	feature_profile_id = sdwan_transport_feature_profile.test.id` + "\n"
	config += `	type = "t1"` + "\n"
	config += `	slot = 11` + "\n"
	config += `	entries = [{` + "\n"
	config += `	  t1_description = "T1"` + "\n"
	config += `	  t1_framing = "esf"` + "\n"
	config += `	  t1_linecode = "ami"` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigMinimum

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAll
func testAccSdwanTransportT1E1ControllerProfileParcelConfig_all() string {
	config := `resource "sdwan_transport_t1_e1_controller_feature" "test" {` + "\n"
	config += ` name = "TF_TEST_ALL"` + "\n"
	config += ` description = "Terraform integration test"` + "\n"
	config += `	feature_profile_id = sdwan_transport_feature_profile.test.id` + "\n"
	config += `	type = "t1"` + "\n"
	config += `	slot = 11` + "\n"
	config += `	entries = [{` + "\n"
	config += `	  t1_description = "T1"` + "\n"
	config += `	  t1_framing = "esf"` + "\n"
	config += `	  t1_linecode = "ami"` + "\n"
	config += `	  cable_length = "long"` + "\n"
	config += `	  length_long = "-7.5db"` + "\n"
	config += `	  clock_source = "line"` + "\n"
	config += `	  line_mode = "primary"` + "\n"
	config += `	  description = "desc"` + "\n"
	config += `	  channel_groups = [{` + "\n"
	config += `		channel_group = 12` + "\n"
	config += `		time_slot = "timeslots 15"` + "\n"
	config += `	}]` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigAll
