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
func TestAccSdwanTransportGPSProfileParcel(t *testing.T) {
	if os.Getenv("SDWAN_2012") == "" {
		t.Skip("skipping test, set environment variable SDWAN_2012")
	}
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_transport_gps_profile_parcel.test", "gps_enable", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_transport_gps_profile_parcel.test", "gps_mode", "ms-based"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_transport_gps_profile_parcel.test", "nmea_enable", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_transport_gps_profile_parcel.test", "nmea_source_address", "1.2.3.4"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_transport_gps_profile_parcel.test", "nmea_destination_address", "2.3.4.5"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_transport_gps_profile_parcel.test", "nmea_destination_port", "22"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccSdwanTransportGPSPrerequisitesProfileParcelConfig + testAccSdwanTransportGPSProfileParcelConfig_minimum(),
			},
			{
				Config: testAccSdwanTransportGPSPrerequisitesProfileParcelConfig + testAccSdwanTransportGPSProfileParcelConfig_all(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAcc

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites
const testAccSdwanTransportGPSPrerequisitesProfileParcelConfig = `
resource "sdwan_transport_feature_profile" "test" {
  name = "TF_TEST"
  description = "Terraform test"
}
`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigMinimum
func testAccSdwanTransportGPSProfileParcelConfig_minimum() string {
	config := `resource "sdwan_transport_gps_profile_parcel" "test" {` + "\n"
	config += ` name = "TF_TEST_MIN"` + "\n"
	config += ` description = "Terraform integration test"` + "\n"
	config += `	feature_profile_id = sdwan_transport_feature_profile.test.id` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigMinimum

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAll
func testAccSdwanTransportGPSProfileParcelConfig_all() string {
	config := `resource "sdwan_transport_gps_profile_parcel" "test" {` + "\n"
	config += ` name = "TF_TEST_ALL"` + "\n"
	config += ` description = "Terraform integration test"` + "\n"
	config += `	feature_profile_id = sdwan_transport_feature_profile.test.id` + "\n"
	config += `	gps_enable = false` + "\n"
	config += `	gps_mode = "ms-based"` + "\n"
	config += `	nmea_enable = false` + "\n"
	config += `	nmea_source_address = "1.2.3.4"` + "\n"
	config += `	nmea_destination_address = "2.3.4.5"` + "\n"
	config += `	nmea_destination_port = 22` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigAll
