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
func TestAccDataSourceSdwanSystemRemoteAccessProfileParcel(t *testing.T) {
	if os.Getenv("SDWAN_2012") == "" {
		t.Skip("skipping test, set environment variable SDWAN_2012")
	}
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_system_remote_access_profile_parcel.test", "connection_type_ssl", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_system_remote_access_profile_parcel.test", "any_connect_eap_authentication_type", "user"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_system_remote_access_profile_parcel.test", "ipv4_pool_size", "50"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_system_remote_access_profile_parcel.test", "ipv6_pool_size", "1024"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_system_remote_access_profile_parcel.test", "enable_certificate_list_check", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_system_remote_access_profile_parcel.test", "psk_authentication_type", "aaa"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_system_remote_access_profile_parcel.test", "psk_authentication_pre_shared_key", "Cisco123"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_system_remote_access_profile_parcel.test", "radius_group_name", "radius-1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_system_remote_access_profile_parcel.test", "aaa_derive_name_from_peer_identity", "MyPassword"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_system_remote_access_profile_parcel.test", "aaa_enable_accounting", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_system_remote_access_profile_parcel.test", "ikev2_local_ike_identity_type", "EMAIL"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_system_remote_access_profile_parcel.test", "ikev2_local_ike_identity_value", "abc@xyz.com"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_system_remote_access_profile_parcel.test", "ikev2_security_association_lifetime", "86400"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_system_remote_access_profile_parcel.test", "ikev2_anti_dos_threshold", "99"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_system_remote_access_profile_parcel.test", "ipsec_enable_anti_replay", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_system_remote_access_profile_parcel.test", "ipsec_anti_replay_window_size", "64"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_system_remote_access_profile_parcel.test", "ipsec_security_association_lifetime", "3600"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_system_remote_access_profile_parcel.test", "ipsec_enable_perfect_foward_secrecy", "false"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceSdwanSystemRemoteAccessPrerequisitesProfileParcelConfig + testAccDataSourceSdwanSystemRemoteAccessProfileParcelConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAccDataSource

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites
const testAccDataSourceSdwanSystemRemoteAccessPrerequisitesProfileParcelConfig = `
resource "sdwan_system_feature_profile" "test" {
  name = "TF_TEST"
  description = "Terraform test"
}
`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSourceConfig
func testAccDataSourceSdwanSystemRemoteAccessProfileParcelConfig() string {
	config := `resource "sdwan_system_remote_access_profile_parcel" "test" {` + "\n"
	config += ` name = "TF_TEST"` + "\n"
	config += ` description = "Terraform integration test"` + "\n"
	config += `	feature_profile_id = sdwan_system_feature_profile.test.id` + "\n"
	config += `	connection_type_ssl = false` + "\n"
	config += `	any_connect_eap_authentication_type = "user"` + "\n"
	config += `	ipv4_pool_size = 50` + "\n"
	config += `	ipv6_pool_size = 1024` + "\n"
	config += `	enable_certificate_list_check = false` + "\n"
	config += `	psk_authentication_type = "aaa"` + "\n"
	config += `	psk_authentication_pre_shared_key = "Cisco123"` + "\n"
	config += `	radius_group_name = "radius-1"` + "\n"
	config += `	aaa_derive_name_from_peer_identity = "MyPassword"` + "\n"
	config += `	aaa_enable_accounting = false` + "\n"
	config += `	ikev2_local_ike_identity_type = "EMAIL"` + "\n"
	config += `	ikev2_local_ike_identity_value = "abc@xyz.com"` + "\n"
	config += `	ikev2_security_association_lifetime = 86400` + "\n"
	config += `	ikev2_anti_dos_threshold = 99` + "\n"
	config += `	ipsec_enable_anti_replay = false` + "\n"
	config += `	ipsec_anti_replay_window_size = 64` + "\n"
	config += `	ipsec_security_association_lifetime = 3600` + "\n"
	config += `	ipsec_enable_perfect_foward_secrecy = false` + "\n"
	config += `}` + "\n"

	config += `
		data "sdwan_system_remote_access_profile_parcel" "test" {
			id = sdwan_system_remote_access_profile_parcel.test.id
			feature_profile_id = sdwan_system_feature_profile.test.id
		}
	`
	return config
}

// End of section. //template:end testAccDataSourceConfig
