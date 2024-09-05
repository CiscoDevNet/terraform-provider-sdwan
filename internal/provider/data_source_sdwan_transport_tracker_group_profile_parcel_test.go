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

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSource
func TestAccDataSourceSdwanTransportTrackerGroupProfileParcel(t *testing.T) {
	if os.Getenv("SDWAN_2012") == "" {
		t.Skip("skipping test, set environment variable SDWAN_2012")
	}
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_transport_tracker_group_feature.test", "tracker_boolean", "or"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceSdwanTransportTrackerGroupPrerequisitesProfileParcelConfig + testAccDataSourceSdwanTransportTrackerGroupProfileParcelConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAccDataSource

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites
const testAccDataSourceSdwanTransportTrackerGroupPrerequisitesProfileParcelConfig = `
resource "sdwan_transport_feature_profile" "test" {
  name = "TF_TEST"
  description = "Terraform test"
}

resource "sdwan_transport_tracker_profile_parcel" "test-1" {
  name                  = "TF_TEST_1"
  description           = "Terraform Test"
  feature_profile_id    = sdwan_transport_feature_profile.test.id
  tracker_name          = "TRACKER_1"
  endpoint_api_url      = "google.com"
  endpoint_dns_name     = "google.com"
  endpoint_ip           = "1.2.3.4"
  interval              = 30
  multiplier            = 3
  threshold             = 300
  endpoint_tracker_type = "interface"
  tracker_type          = "endpoint"
}

resource "sdwan_transport_tracker_profile_parcel" "test-2" {
  name                  = "TF_TEST_2"
  description           = "Terraform Test"
  feature_profile_id    = sdwan_transport_feature_profile.test.id
  tracker_name          = "TRACKER_1"
  endpoint_api_url      = "google.com"
  endpoint_dns_name     = "google.com"
  endpoint_ip           = "1.2.3.4"
  interval              = 30
  multiplier            = 3
  threshold             = 300
  endpoint_tracker_type = "interface"
  tracker_type          = "endpoint"
}
`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSourceConfig
func testAccDataSourceSdwanTransportTrackerGroupProfileParcelConfig() string {
	config := `resource "sdwan_transport_tracker_group_feature" "test" {` + "\n"
	config += ` name = "TF_TEST"` + "\n"
	config += ` description = "Terraform integration test"` + "\n"
	config += `	feature_profile_id = sdwan_transport_feature_profile.test.id` + "\n"
	config += `	tracker_elements = [{` + "\n"
	config += `	  tracker_id = sdwan_transport_tracker_profile_parcel.test-1.id` + "\n"
	config += `	}, {` + "\n"
	config += `	  tracker_id = sdwan_transport_tracker_profile_parcel.test-2.id` + "\n"
	config += `	}]` + "\n"
	config += `	tracker_boolean = "or"` + "\n"
	config += `}` + "\n"

	config += `
		data "sdwan_transport_tracker_group_feature" "test" {
			id = sdwan_transport_tracker_group_feature.test.id
			feature_profile_id = sdwan_transport_feature_profile.test.id
		}
	`
	return config
}

// End of section. //template:end testAccDataSourceConfig
