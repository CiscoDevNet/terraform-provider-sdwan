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
func TestAccDataSourceSdwanSystemBFDProfileParcel(t *testing.T) {
	if os.Getenv("SDWAN_2012") == "" {
		t.Skip("skipping test, set environment variable SDWAN_2012")
	}
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_system_bfd_profile_parcel.test", "multiplier", "3"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_system_bfd_profile_parcel.test", "poll_interval", "100"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_system_bfd_profile_parcel.test", "default_dscp", "8"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_system_bfd_profile_parcel.test", "colors.0.color", "3g"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_system_bfd_profile_parcel.test", "colors.0.hello_interval", "200"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_system_bfd_profile_parcel.test", "colors.0.multiplier", "3"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_system_bfd_profile_parcel.test", "colors.0.pmtu_discovery", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_system_bfd_profile_parcel.test", "colors.0.dscp", "16"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceSdwanSystemBFDPrerequisitesProfileParcelConfig + testAccDataSourceSdwanSystemBFDProfileParcelConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAccDataSource

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites
const testAccDataSourceSdwanSystemBFDPrerequisitesProfileParcelConfig = `
resource "sdwan_system_feature_profile" "test" {
  name = "TF_TEST"
  description = "Terraform test"
}
`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSourceConfig
func testAccDataSourceSdwanSystemBFDProfileParcelConfig() string {
	config := `resource "sdwan_system_bfd_profile_parcel" "test" {` + "\n"
	config += ` name = "TF_TEST"` + "\n"
	config += ` description = "Terraform integration test"` + "\n"
	config += `	feature_profile_id = sdwan_system_feature_profile.test.id` + "\n"
	config += `	multiplier = 3` + "\n"
	config += `	poll_interval = 100` + "\n"
	config += `	default_dscp = 8` + "\n"
	config += `	colors = [{` + "\n"
	config += `	  color = "3g"` + "\n"

	config += `	  hello_interval = 200` + "\n"

	config += `	  multiplier = 3` + "\n"

	config += `	  pmtu_discovery = true` + "\n"

	config += `	  dscp = 16` + "\n"

	config += `	}]` + "\n"
	config += `}` + "\n"

	config += `
		data "sdwan_system_bfd_profile_parcel" "test" {
			id = sdwan_system_bfd_profile_parcel.test.id
			feature_profile_id = sdwan_system_feature_profile.test.id
		}
	`
	return config
}

// End of section. //template:end testAccDataSourceConfig
