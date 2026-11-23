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
func TestAccDataSourceSdwanOtherTrustSecProfileParcel(t *testing.T) {
	if os.Getenv("SDWAN_2018") == "" {
		t.Skip("skipping test, set environment variable SDWAN_2018")
	}
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_other_trustsec_feature.test", "device_id", "trustsecDevice"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_other_trustsec_feature.test", "device_sgt", "100"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_other_trustsec_feature.test", "enable_enforcement", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_other_trustsec_feature.test", "enable_sxp", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_other_trustsec_feature.test", "listener_hold_time_min", "90"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_other_trustsec_feature.test", "listener_hold_time_max", "180"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_other_trustsec_feature.test", "speaker_hold_time", "120"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_other_trustsec_feature.test", "sxp_key_chain", "keychain1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_other_trustsec_feature.test", "sxp_log_binding_changes", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_other_trustsec_feature.test", "sxp_reconciliation_period", "120"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_other_trustsec_feature.test", "sxp_retry_period", "120"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_other_trustsec_feature.test", "sxp_source_ip", "10.0.0.1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_other_trustsec_feature.test", "sxp_connections.0.peer_ip", "10.0.0.2"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_other_trustsec_feature.test", "sxp_connections.0.source_ip", "10.0.0.1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_other_trustsec_feature.test", "sxp_connections.0.preshared_key", "none"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_other_trustsec_feature.test", "sxp_connections.0.mode", "peer"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_other_trustsec_feature.test", "sxp_connections.0.mode_type", "speaker"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_other_trustsec_feature.test", "sxp_connections.0.vpn_id", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_other_trustsec_feature.test", "sxp_connections.0.min_hold_time", "90"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_other_trustsec_feature.test", "sxp_connections.0.max_hold_time", "120"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceSdwanOtherTrustSecPrerequisitesProfileParcelConfig + testAccDataSourceSdwanOtherTrustSecProfileParcelConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
			{
				Config: testAccDataSourceSdwanOtherTrustSecPrerequisitesProfileParcelConfig + testAccDataSourceSdwanOtherTrustSecProfileParcelByNameConfig(),
				Check: resource.ComposeTestCheckFunc(
					append(checks,
						resource.TestCheckResourceAttr("data.sdwan_other_trustsec_feature.test", "name", "TF_TEST"),
						resource.TestCheckResourceAttrSet("data.sdwan_other_trustsec_feature.test", "id"),
					)...),
			},
		},
	})
}

// End of section. //template:end testAccDataSource

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites
const testAccDataSourceSdwanOtherTrustSecPrerequisitesProfileParcelConfig = `
resource "sdwan_other_feature_profile" "test" {
  name        = "TF_TEST"
  description = "Terraform test"
}

`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSourceConfig
func testAccDataSourceSdwanOtherTrustSecProfileParcelConfig() string {
	config := `resource "sdwan_other_trustsec_feature" "test" {` + "\n"
	config += ` name = "TF_TEST"` + "\n"
	config += ` description = "Terraform integration test"` + "\n"
	config += `	feature_profile_id = sdwan_other_feature_profile.test.id` + "\n"
	config += `	device_id = "trustsecDevice"` + "\n"
	config += `	device_password = "Cisco123"` + "\n"
	config += `	device_sgt = 100` + "\n"
	config += `	enable_enforcement = true` + "\n"
	config += `	enable_sxp = true` + "\n"
	config += `	listener_hold_time_min = 90` + "\n"
	config += `	listener_hold_time_max = 180` + "\n"
	config += `	speaker_hold_time = 120` + "\n"
	config += `	sxp_key_chain = "keychain1"` + "\n"
	config += `	sxp_log_binding_changes = true` + "\n"
	config += `	sxp_reconciliation_period = 120` + "\n"
	config += `	sxp_retry_period = 120` + "\n"
	config += `	sxp_source_ip = "10.0.0.1"` + "\n"
	config += `	sxp_connections = [{` + "\n"
	config += `	  peer_ip = "10.0.0.2"` + "\n"
	config += `	  source_ip = "10.0.0.1"` + "\n"
	config += `	  preshared_key = "none"` + "\n"
	config += `	  mode = "peer"` + "\n"
	config += `	  mode_type = "speaker"` + "\n"
	config += `	  vpn_id = 1` + "\n"
	config += `	  min_hold_time = 90` + "\n"
	config += `	  max_hold_time = 120` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"

	config += `
		data "sdwan_other_trustsec_feature" "test" {
			id = sdwan_other_trustsec_feature.test.id
			feature_profile_id = sdwan_other_feature_profile.test.id
		}
	`
	return config
}

// End of section. //template:end testAccDataSourceConfig

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSourceByNameConfig
func testAccDataSourceSdwanOtherTrustSecProfileParcelByNameConfig() string {
	config := `resource "sdwan_other_trustsec_feature" "test" {` + "\n"
	config += ` name = "TF_TEST"` + "\n"
	config += ` description = "Terraform integration test"` + "\n"
	config += `	feature_profile_id = sdwan_other_feature_profile.test.id` + "\n"
	config += `	device_id = "trustsecDevice"` + "\n"
	config += `	device_password = "Cisco123"` + "\n"
	config += `	device_sgt = 100` + "\n"
	config += `	enable_enforcement = true` + "\n"
	config += `	enable_sxp = true` + "\n"
	config += `	listener_hold_time_min = 90` + "\n"
	config += `	listener_hold_time_max = 180` + "\n"
	config += `	speaker_hold_time = 120` + "\n"
	config += `	sxp_key_chain = "keychain1"` + "\n"
	config += `	sxp_log_binding_changes = true` + "\n"
	config += `	sxp_reconciliation_period = 120` + "\n"
	config += `	sxp_retry_period = 120` + "\n"
	config += `	sxp_source_ip = "10.0.0.1"` + "\n"
	config += `	sxp_connections = [{` + "\n"
	config += `	  peer_ip = "10.0.0.2"` + "\n"
	config += `	  source_ip = "10.0.0.1"` + "\n"
	config += `	  preshared_key = "none"` + "\n"
	config += `	  mode = "peer"` + "\n"
	config += `	  mode_type = "speaker"` + "\n"
	config += `	  vpn_id = 1` + "\n"
	config += `	  min_hold_time = 90` + "\n"
	config += `	  max_hold_time = 120` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"

	config += `
		data "sdwan_other_trustsec_feature" "test" {
			name = "TF_TEST"
			feature_profile_id = sdwan_other_feature_profile.test.id
		}
	`
	return config
}

// End of section. //template:end testAccDataSourceByNameConfig
