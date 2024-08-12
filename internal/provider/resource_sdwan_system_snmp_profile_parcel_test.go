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
func TestAccSdwanSystemSNMPProfileParcel(t *testing.T) {
	if os.Getenv("SDWAN_2012") == "" {
		t.Skip("skipping test, set environment variable SDWAN_2012")
	}
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_system_snmp_profile_parcel.test", "shutdown", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_system_snmp_profile_parcel.test", "contact_person", "wixie.cisco"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_system_snmp_profile_parcel.test", "location_of_device", "SHANGHAI"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_system_snmp_profile_parcel.test", "views.0.name", "VIEW1"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_system_snmp_profile_parcel.test", "views.0.oids.0.id", "1.3.6.1.4.1.9.9.394"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_system_snmp_profile_parcel.test", "views.0.oids.0.exclude", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_system_snmp_profile_parcel.test", "communities.0.user_label", "COMMUNITY1"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_system_snmp_profile_parcel.test", "communities.0.view", "VIEW1"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_system_snmp_profile_parcel.test", "communities.0.authorization", "read-only"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_system_snmp_profile_parcel.test", "groups.0.name", "GROUP1"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_system_snmp_profile_parcel.test", "groups.0.security_level", "auth-priv"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_system_snmp_profile_parcel.test", "groups.0.view", "VIEW1"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_system_snmp_profile_parcel.test", "users.0.name", "USER1"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_system_snmp_profile_parcel.test", "users.0.authentication_protocol", "sha"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_system_snmp_profile_parcel.test", "users.0.authentication_password", "$CRYPT_CLUSTER$su56l1Z0Tk4Qc9N7+T/uOg==$sD6b0HLqEdI+RNwsEOoLcQ=="))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_system_snmp_profile_parcel.test", "users.0.privacy_protocol", "aes-256-cfb-128"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_system_snmp_profile_parcel.test", "users.0.privacy_password", "$CRYPT_CLUSTER$su56l1Z0Tk4Qc9N7+T/uOg==$sD6b0HLqEdI+RNwsEOoLcQ=="))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_system_snmp_profile_parcel.test", "users.0.group", "GROUP1"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_system_snmp_profile_parcel.test", "trap_target_servers.0.vpn_id", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_system_snmp_profile_parcel.test", "trap_target_servers.0.ip", "10.75.221.156"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_system_snmp_profile_parcel.test", "trap_target_servers.0.port", "161"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_system_snmp_profile_parcel.test", "trap_target_servers.0.user_label", "TARGET1"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_system_snmp_profile_parcel.test", "trap_target_servers.0.user", "USER1"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_system_snmp_profile_parcel.test", "trap_target_servers.0.source_interface", "GigabitEthernet1"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccSdwanSystemSNMPPrerequisitesProfileParcelConfig + testAccSdwanSystemSNMPProfileParcelConfig_minimum(),
			},
			{
				Config: testAccSdwanSystemSNMPPrerequisitesProfileParcelConfig + testAccSdwanSystemSNMPProfileParcelConfig_all(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAcc

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites
const testAccSdwanSystemSNMPPrerequisitesProfileParcelConfig = `
resource "sdwan_system_feature_profile" "test" {
  name = "TF_TEST"
  description = "Terraform test"
}
`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigMinimum
func testAccSdwanSystemSNMPProfileParcelConfig_minimum() string {
	config := `resource "sdwan_system_snmp_profile_parcel" "test" {` + "\n"
	config += ` name = "TF_TEST_MIN"` + "\n"
	config += ` description = "Terraform integration test"` + "\n"
	config += `	feature_profile_id = sdwan_system_feature_profile.test.id` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigMinimum

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAll
func testAccSdwanSystemSNMPProfileParcelConfig_all() string {
	config := `resource "sdwan_system_snmp_profile_parcel" "test" {` + "\n"
	config += ` name = "TF_TEST_ALL"` + "\n"
	config += ` description = "Terraform integration test"` + "\n"
	config += `	feature_profile_id = sdwan_system_feature_profile.test.id` + "\n"
	config += `	shutdown = false` + "\n"
	config += `	contact_person = "wixie.cisco"` + "\n"
	config += `	location_of_device = "SHANGHAI"` + "\n"
	config += `	views = [{` + "\n"
	config += `	  name = "VIEW1"` + "\n"

	config += `	  oids = [{` + "\n"
	config += `		id = "1.3.6.1.4.1.9.9.394"` + "\n"

	config += `		exclude = false` + "\n"

	config += `	}]` + "\n"
	config += `	}]` + "\n"
	config += `	communities = [{` + "\n"
	config += `	  name = "example"` + "\n"

	config += `	  user_label = "COMMUNITY1"` + "\n"

	config += `	  view = "VIEW1"` + "\n"

	config += `	  authorization = "read-only"` + "\n"

	config += `	}]` + "\n"
	config += `	groups = [{` + "\n"
	config += `	  name = "GROUP1"` + "\n"

	config += `	  security_level = "auth-priv"` + "\n"

	config += `	  view = "VIEW1"` + "\n"

	config += `	}]` + "\n"
	config += `	users = [{` + "\n"
	config += `	  name = "USER1"` + "\n"

	config += `	  authentication_protocol = "sha"` + "\n"

	config += `	  authentication_password = "$CRYPT_CLUSTER$su56l1Z0Tk4Qc9N7+T/uOg==$sD6b0HLqEdI+RNwsEOoLcQ=="` + "\n"

	config += `	  privacy_protocol = "aes-256-cfb-128"` + "\n"

	config += `	  privacy_password = "$CRYPT_CLUSTER$su56l1Z0Tk4Qc9N7+T/uOg==$sD6b0HLqEdI+RNwsEOoLcQ=="` + "\n"

	config += `	  group = "GROUP1"` + "\n"

	config += `	}]` + "\n"
	config += `	trap_target_servers = [{` + "\n"
	config += `	  vpn_id = 1` + "\n"

	config += `	  ip = "10.75.221.156"` + "\n"

	config += `	  port = 161` + "\n"

	config += `	  user_label = "TARGET1"` + "\n"

	config += `	  user = "USER1"` + "\n"

	config += `	  source_interface = "GigabitEthernet1"` + "\n"

	config += `	}]` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigAll
