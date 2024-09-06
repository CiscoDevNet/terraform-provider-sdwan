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
func TestAccSdwanServiceObjectTrackerGroupProfileParcel(t *testing.T) {
	if os.Getenv("SDWAN_2012") == "" {
		t.Skip("skipping test, set environment variable SDWAN_2012")
	}
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_service_object_tracker_group_feature.test", "reachable", "or"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{

			{
				Config: testAccSdwanServiceObjectTrackerGroupPrerequisitesProfileParcelConfig + testAccSdwanServiceObjectTrackerGroupProfileParcelConfig_all(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAcc

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites
const testAccSdwanServiceObjectTrackerGroupPrerequisitesProfileParcelConfig = `
resource "sdwan_service_feature_profile" "test" {
  name = "TF_TEST"
  description = "Terraform test"
}

resource "sdwan_service_object_tracker_feature" "test-1" {
  name               = "TF_TEST_1"
  description        = "My Example"
  feature_profile_id = sdwan_service_feature_profile.test.id
  object_tracker_id  = 11
  tracker_type       = "Interface"
  interface          = "GigabitEthernet1"
}

resource "sdwan_service_object_tracker_feature" "test-2" {
  name               = "TF_TEST_2"
  description        = "My Example"
  feature_profile_id = sdwan_service_feature_profile.test.id
  object_tracker_id  = 12
  tracker_type       = "Interface"
  interface          = "GigabitEthernet1"
}
`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigMinimum

// End of section. //template:end testAccConfigMinimum

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAll
func testAccSdwanServiceObjectTrackerGroupProfileParcelConfig_all() string {
	config := `resource "sdwan_service_object_tracker_group_feature" "test" {` + "\n"
	config += ` name = "TF_TEST_ALL"` + "\n"
	config += ` description = "Terraform integration test"` + "\n"
	config += `	feature_profile_id = sdwan_service_feature_profile.test.id` + "\n"
	config += `	object_tracker_id = 10` + "\n"
	config += `	tracker_elements = [{` + "\n"
	config += `	  object_tracker_id = sdwan_service_object_tracker_profile_parcel.test-1.id` + "\n"
	config += `	}, {` + "\n"
	config += `	  object_tracker_id = sdwan_service_object_tracker_profile_parcel.test-2.id` + "\n"
	config += `	}]` + "\n"
	config += `	reachable = "or"` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigAll
