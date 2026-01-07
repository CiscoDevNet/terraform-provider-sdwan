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
func TestAccSdwanTransportCellularProfileProfileParcel(t *testing.T) {
	if os.Getenv("SDWAN_2015_IN_PROGRESS") == "" {
		t.Skip("skipping test, set environment variable SDWAN_2015_IN_PROGRESS")
	}
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_transport_cellular_profile_feature.test", "profile_id", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_transport_cellular_profile_feature.test", "access_point_name", "apn1"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_transport_cellular_profile_feature.test", "authentication_type", "pap"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_transport_cellular_profile_feature.test", "profile_username", "example"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_transport_cellular_profile_feature.test", "packet_data_network_type", "ipv4"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_transport_cellular_profile_feature.test", "no_overwrite", "false"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{

			{
				Config: testAccSdwanTransportCellularProfilePrerequisitesProfileParcelConfig + testAccSdwanTransportCellularProfileProfileParcelConfig_all(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAcc

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites
const testAccSdwanTransportCellularProfilePrerequisitesProfileParcelConfig = `
resource "sdwan_transport_feature_profile" "test" {
  name = "TF_TEST"
  description = "Terraform test"
}
`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigMinimum

// End of section. //template:end testAccConfigMinimum

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAll
func testAccSdwanTransportCellularProfileProfileParcelConfig_all() string {
	config := `resource "sdwan_transport_cellular_profile_feature" "test" {` + "\n"
	config += ` name = "TF_TEST_ALL"` + "\n"
	config += ` description = "Terraform integration test"` + "\n"
	config += `	feature_profile_id = sdwan_transport_feature_profile.test.id` + "\n"
	config += `	profile_id = 1` + "\n"
	config += `	access_point_name = "apn1"` + "\n"
	config += `	requires_authentication = true` + "\n"
	config += `	authentication_type = "pap"` + "\n"
	config += `	profile_username = "example"` + "\n"
	config += `	profile_password = "example123!"` + "\n"
	config += `	packet_data_network_type = "ipv4"` + "\n"
	config += `	no_overwrite = false` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigAll
