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
func TestAccSdwanServiceTrackerGroupProfileParcel(t *testing.T) {
	if os.Getenv("SDWAN_2012") == "" {
		t.Skip("skipping test, set environment variable SDWAN_2012")
	}
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_service_tracker_group_profile_parcel.test", "tracker_elements.0.tracker_id", "615d948f-34ee-4a2e-810e-a9bd8d3d48ec"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_service_tracker_group_profile_parcel.test", "tracker_boolean", "or"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccSdwanServiceTrackerGroupPrerequisitesProfileParcelConfig + testAccSdwanServiceTrackerGroupProfileParcelConfig_minimum(),
			},
			{
				Config: testAccSdwanServiceTrackerGroupPrerequisitesProfileParcelConfig + testAccSdwanServiceTrackerGroupProfileParcelConfig_all(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAcc

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites
const testAccSdwanServiceTrackerGroupPrerequisitesProfileParcelConfig = `
resource "sdwan_service_feature_profile" "test" {
  name = "TF_TEST"
  description = "Terraform test"
}

resource "sdwan_service_tracker_profile_parcel" "test" {
  name                  = "TF_TEST"
  description           = "Terraform test"
  feature_profile_id    = sdwan_service_feature_profile.test.id
  tracker_name          = "TRACKER_2"
  endpoint_api_url      = "google.com"
  endpoint_dns_name     = "google.com"
  endpoint_ip           = "1.2.3.4"
  protocol              = "tcp"
  port                  = 123
  interval              = 30
  multiplier            = 3
  threshold             = 300
  endpoint_tracker_type = "static-route"
  tracker_type          = "endpoint"
}
`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigMinimum
func testAccSdwanServiceTrackerGroupProfileParcelConfig_minimum() string {
	config := `resource "sdwan_service_tracker_group_profile_parcel" "test" {` + "\n"
	config += ` name = "TF_TEST_MIN"` + "\n"
	config += ` description = "Terraform integration test"` + "\n"
	config += `	feature_profile_id = sdwan_service_feature_profile.test.id` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigMinimum

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAll
func testAccSdwanServiceTrackerGroupProfileParcelConfig_all() string {
	config := `resource "sdwan_service_tracker_group_profile_parcel" "test" {` + "\n"
	config += ` name = "TF_TEST_ALL"` + "\n"
	config += ` description = "Terraform integration test"` + "\n"
	config += `	feature_profile_id = sdwan_service_feature_profile.test.id` + "\n"
	config += `	tracker_elements = [{` + "\n"
	config += `	  tracker_id = sdwan_service_tracker_profile_parcel.test.id` + "\n"
	config += `	}]` + "\n"
	config += `	tracker_boolean = "or"` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigAll
