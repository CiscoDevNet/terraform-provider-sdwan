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
func TestAccSdwanSystemLoggingProfileParcel(t *testing.T) {
	if os.Getenv("SDWAN_2012") == "" {
		t.Skip("skipping test, set environment variable SDWAN_2012")
	}
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_system_logging_profile_parcel.test", "disk_enable", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_system_logging_profile_parcel.test", "disk_file_size", "9"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_system_logging_profile_parcel.test", "disk_file_rotate", "10"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_system_logging_profile_parcel.test", "tls_profiles.0.profile", "test"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_system_logging_profile_parcel.test", "tls_profiles.0.tls_version", "TLSv1.1"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_system_logging_profile_parcel.test", "ipv4_servers.0.hostname_ip", "1.1.1.1"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_system_logging_profile_parcel.test", "ipv4_servers.0.vpn", "512"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_system_logging_profile_parcel.test", "ipv4_servers.0.source_interface", "GigabitEthernet1"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_system_logging_profile_parcel.test", "ipv4_servers.0.priority", "informational"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_system_logging_profile_parcel.test", "ipv4_servers.0.tls_enable", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_system_logging_profile_parcel.test", "ipv4_servers.0.tls_properties_custom_profile", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_system_logging_profile_parcel.test", "ipv4_servers.0.tls_properties_profile", "test"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_system_logging_profile_parcel.test", "ipv6_servers.0.hostname_ip", "1.1.1.1"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_system_logging_profile_parcel.test", "ipv6_servers.0.vpn", "512"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_system_logging_profile_parcel.test", "ipv6_servers.0.source_interface", "GigabitEthernet1"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_system_logging_profile_parcel.test", "ipv6_servers.0.priority", "informational"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_system_logging_profile_parcel.test", "ipv6_servers.0.tls_enable", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_system_logging_profile_parcel.test", "ipv6_servers.0.tls_properties_custom_profile", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_system_logging_profile_parcel.test", "ipv6_servers.0.tls_properties_profile", "test"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccSdwanSystemLoggingPrerequisitesProfileParcelConfig + testAccSdwanSystemLoggingProfileParcelConfig_minimum(),
			},
			{
				Config: testAccSdwanSystemLoggingPrerequisitesProfileParcelConfig + testAccSdwanSystemLoggingProfileParcelConfig_all(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAcc

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites
const testAccSdwanSystemLoggingPrerequisitesProfileParcelConfig = `
resource "sdwan_system_feature_profile" "test" {
  name = "TF_TEST"
  description = "Terraform test"
}

`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigMinimum
func testAccSdwanSystemLoggingProfileParcelConfig_minimum() string {
	config := `resource "sdwan_system_logging_profile_parcel" "test" {` + "\n"
	config += ` name = "TF_TEST_MIN"` + "\n"
	config += ` description = "Terraform integration test"` + "\n"
	config += `	feature_profile_id = sdwan_system_feature_profile.test.id` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigMinimum

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAll
func testAccSdwanSystemLoggingProfileParcelConfig_all() string {
	config := `resource "sdwan_system_logging_profile_parcel" "test" {` + "\n"
	config += ` name = "TF_TEST_ALL"` + "\n"
	config += ` description = "Terraform integration test"` + "\n"
	config += `	feature_profile_id = sdwan_system_feature_profile.test.id` + "\n"
	config += `	disk_enable = true` + "\n"
	config += `	disk_file_size = 9` + "\n"
	config += `	disk_file_rotate = 10` + "\n"
	config += `	tls_profiles = [{` + "\n"
	config += `	  profile = "test"` + "\n"
	config += `	  tls_version = "TLSv1.1"` + "\n"
	config += `	  cipher_suites = ["aes-128-cbc-sha"]` + "\n"
	config += `	}]` + "\n"
	config += `	ipv4_servers = [{` + "\n"
	config += `	  hostname_ip = "1.1.1.1"` + "\n"
	config += `	  vpn = 512` + "\n"
	config += `	  source_interface = "GigabitEthernet1"` + "\n"
	config += `	  priority = "informational"` + "\n"
	config += `	  tls_enable = true` + "\n"
	config += `	  tls_properties_custom_profile = true` + "\n"
	config += `	  tls_properties_profile = "test"` + "\n"
	config += `	}]` + "\n"
	config += `	ipv6_servers = [{` + "\n"
	config += `	  hostname_ip = "1.1.1.1"` + "\n"
	config += `	  vpn = 512` + "\n"
	config += `	  source_interface = "GigabitEthernet1"` + "\n"
	config += `	  priority = "informational"` + "\n"
	config += `	  tls_enable = true` + "\n"
	config += `	  tls_properties_custom_profile = true` + "\n"
	config += `	  tls_properties_profile = "test"` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigAll
