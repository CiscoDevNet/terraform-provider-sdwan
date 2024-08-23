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
func TestAccDataSourceSdwanSystemNTPProfileParcel(t *testing.T) {
	if os.Getenv("SDWAN_2012") == "" {
		t.Skip("skipping test, set environment variable SDWAN_2012")
	}
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_system_ntp_profile_parcel.test", "servers.0.hostname_ip_address", "1.1.1.1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_system_ntp_profile_parcel.test", "servers.0.authentication_key", "41673"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_system_ntp_profile_parcel.test", "servers.0.vpn", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_system_ntp_profile_parcel.test", "servers.0.ntp_version", "4"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_system_ntp_profile_parcel.test", "servers.0.source_interface", "Ethernet"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_system_ntp_profile_parcel.test", "servers.0.prefer_this_ntp_server", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_system_ntp_profile_parcel.test", "authentication_keys.0.key_id", "49737"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_system_ntp_profile_parcel.test", "authentication_keys.0.md5_value", "$CRYPT_CLUSTER"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_system_ntp_profile_parcel.test", "authoritative_ntp_server", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_system_ntp_profile_parcel.test", "stratum", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_system_ntp_profile_parcel.test", "source_interface", "ATM"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceSdwanSystemNTPPrerequisitesProfileParcelConfig + testAccDataSourceSdwanSystemNTPProfileParcelConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAccDataSource

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites
const testAccDataSourceSdwanSystemNTPPrerequisitesProfileParcelConfig = `
resource "sdwan_system_feature_profile" "test" {
  name = "TF_TEST"
  description = "Terraform test"
}
`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSourceConfig
func testAccDataSourceSdwanSystemNTPProfileParcelConfig() string {
	config := `resource "sdwan_system_ntp_profile_parcel" "test" {` + "\n"
	config += ` name = "TF_TEST"` + "\n"
	config += ` description = "Terraform integration test"` + "\n"
	config += `	feature_profile_id = sdwan_system_feature_profile.test.id` + "\n"
	config += `	servers = [{` + "\n"
	config += `	  hostname_ip_address = "1.1.1.1"` + "\n"
	config += `	  authentication_key = 41673` + "\n"
	config += `	  vpn = 1` + "\n"
	config += `	  ntp_version = 4` + "\n"
	config += `	  source_interface = "Ethernet"` + "\n"
	config += `	  prefer_this_ntp_server = false` + "\n"
	config += `	}]` + "\n"
	config += `	authentication_keys = [{` + "\n"
	config += `	  key_id = 49737` + "\n"
	config += `	  md5_value = "$CRYPT_CLUSTER"` + "\n"
	config += `	}]` + "\n"
	config += `	trusted_keys = [49737]` + "\n"
	config += `	authoritative_ntp_server = false` + "\n"
	config += `	stratum = 1` + "\n"
	config += `	source_interface = "ATM"` + "\n"
	config += `}` + "\n"

	config += `
		data "sdwan_system_ntp_profile_parcel" "test" {
			id = sdwan_system_ntp_profile_parcel.test.id
			feature_profile_id = sdwan_system_feature_profile.test.id
		}
	`
	return config
}

// End of section. //template:end testAccDataSourceConfig
