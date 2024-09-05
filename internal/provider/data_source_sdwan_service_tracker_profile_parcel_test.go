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
func TestAccDataSourceSdwanServiceTrackerProfileParcel(t *testing.T) {
	if os.Getenv("SDWAN_2012") == "" {
		t.Skip("skipping test, set environment variable SDWAN_2012")
	}
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_service_tracker_feature.test", "tracker_name", "TRACKER_2"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_service_tracker_feature.test", "endpoint_api_url", "google.com"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_service_tracker_feature.test", "endpoint_dns_name", "google.com"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_service_tracker_feature.test", "endpoint_ip", "1.2.3.4"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_service_tracker_feature.test", "protocol", "tcp"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_service_tracker_feature.test", "port", "123"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_service_tracker_feature.test", "interval", "30"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_service_tracker_feature.test", "multiplier", "3"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_service_tracker_feature.test", "threshold", "300"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_service_tracker_feature.test", "endpoint_tracker_type", "static-route"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_service_tracker_feature.test", "tracker_type", "endpoint"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceSdwanServiceTrackerPrerequisitesProfileParcelConfig + testAccDataSourceSdwanServiceTrackerProfileParcelConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAccDataSource

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites
const testAccDataSourceSdwanServiceTrackerPrerequisitesProfileParcelConfig = `
resource "sdwan_service_feature_profile" "test" {
  name = "TF_TEST"
  description = "Terraform test"
}
`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSourceConfig
func testAccDataSourceSdwanServiceTrackerProfileParcelConfig() string {
	config := `resource "sdwan_service_tracker_feature" "test" {` + "\n"
	config += ` name = "TF_TEST"` + "\n"
	config += ` description = "Terraform integration test"` + "\n"
	config += `	feature_profile_id = sdwan_service_feature_profile.test.id` + "\n"
	config += `	tracker_name = "TRACKER_2"` + "\n"
	config += `	endpoint_api_url = "google.com"` + "\n"
	config += `	endpoint_dns_name = "google.com"` + "\n"
	config += `	endpoint_ip = "1.2.3.4"` + "\n"
	config += `	protocol = "tcp"` + "\n"
	config += `	port = 123` + "\n"
	config += `	interval = 30` + "\n"
	config += `	multiplier = 3` + "\n"
	config += `	threshold = 300` + "\n"
	config += `	endpoint_tracker_type = "static-route"` + "\n"
	config += `	tracker_type = "endpoint"` + "\n"
	config += `}` + "\n"

	config += `
		data "sdwan_service_tracker_feature" "test" {
			id = sdwan_service_tracker_feature.test.id
			feature_profile_id = sdwan_service_feature_profile.test.id
		}
	`
	return config
}

// End of section. //template:end testAccDataSourceConfig
