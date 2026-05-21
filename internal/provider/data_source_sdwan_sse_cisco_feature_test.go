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
func TestAccDataSourceSdwanSSECiscoProfileParcel(t *testing.T) {
	if os.Getenv("SDWAN_2015") == "" {
		t.Skip("skipping test, set environment variable SDWAN_2015")
	}
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_sse_cisco_feature.test", "context_sharing_for_vpn", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_sse_cisco_feature.test", "context_sharing_for_sgt", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_sse_cisco_feature.test", "interfaces.0.interface_name", "ipsec1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_sse_cisco_feature.test", "interfaces.0.shutdown", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_sse_cisco_feature.test", "interfaces.0.tunnel_dc_preference", "primary-dc"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_sse_cisco_feature.test", "interfaces.0.mtu", "1400"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_sse_cisco_feature.test", "interfaces.0.dpd_interval", "10"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_sse_cisco_feature.test", "interfaces.0.dpd_retries", "3"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_sse_cisco_feature.test", "interfaces.0.ike_version", "2"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_sse_cisco_feature.test", "interfaces.0.ike_rekey_interval", "14400"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_sse_cisco_feature.test", "interfaces.0.ike_ciphersuite", "aes256-cbc-sha1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_sse_cisco_feature.test", "interfaces.0.ike_group", "16"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_sse_cisco_feature.test", "interfaces.0.ipsec_rekey_interval", "3600"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_sse_cisco_feature.test", "interfaces.0.ipsec_replay_window", "512"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_sse_cisco_feature.test", "interfaces.0.ipsec_ciphersuite", "aes256-cbc-sha512"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_sse_cisco_feature.test", "interfaces.0.perfect_forward_secrecy", "group-16"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_sse_cisco_feature.test", "interfaces.0.track_enable", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_sse_cisco_feature.test", "interface_pairs.0.active_interface", "ipsec1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_sse_cisco_feature.test", "interface_pairs.0.active_interface_weight", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_sse_cisco_feature.test", "interface_pairs.0.backup_interface", "None"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_sse_cisco_feature.test", "interface_pairs.0.backup_interface_weight", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_sse_cisco_feature.test", "region", "us-east-1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_sse_cisco_feature.test", "tracker_source_ip", "1.2.3.4"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_sse_cisco_feature.test", "trackers.0.name", "tracker1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_sse_cisco_feature.test", "trackers.0.endpoint_api_url", "http://cisco.com"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_sse_cisco_feature.test", "trackers.0.threshold", "300"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_sse_cisco_feature.test", "trackers.0.interval", "60"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_sse_cisco_feature.test", "trackers.0.multiplier", "3"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceSdwanSSECiscoPrerequisitesProfileParcelConfig + testAccDataSourceSdwanSSECiscoProfileParcelConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
			{
				Config: testAccDataSourceSdwanSSECiscoPrerequisitesProfileParcelConfig + testAccDataSourceSdwanSSECiscoProfileParcelByNameConfig(),
				Check: resource.ComposeTestCheckFunc(
					append(checks,
						resource.TestCheckResourceAttr("data.sdwan_sse_cisco_feature.test", "name", "TF_TEST"),
						resource.TestCheckResourceAttrSet("data.sdwan_sse_cisco_feature.test", "id"),
					)...),
			},
		},
	})
}

// End of section. //template:end testAccDataSource

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites
const testAccDataSourceSdwanSSECiscoPrerequisitesProfileParcelConfig = `
resource "sdwan_sse_feature_profile" "test" {
  name        = "TF_TEST"
  description = "Terraform test"
}
`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSourceConfig
func testAccDataSourceSdwanSSECiscoProfileParcelConfig() string {
	config := `resource "sdwan_sse_cisco_feature" "test" {` + "\n"
	config += ` name = "TF_TEST"` + "\n"
	config += ` description = "Terraform integration test"` + "\n"
	config += `	feature_profile_id = sdwan_sse_feature_profile.test.id` + "\n"
	config += `	context_sharing_for_vpn = false` + "\n"
	config += `	context_sharing_for_sgt = false` + "\n"
	config += `	interfaces = [{` + "\n"
	config += `	  interface_name = "ipsec1"` + "\n"
	config += `	  shutdown = false` + "\n"
	config += `	  tunnel_dc_preference = "primary-dc"` + "\n"
	config += `	  mtu = 1400` + "\n"
	config += `	  dpd_interval = 10` + "\n"
	config += `	  dpd_retries = 3` + "\n"
	config += `	  ike_version = 2` + "\n"
	config += `	  ike_rekey_interval = 14400` + "\n"
	config += `	  ike_ciphersuite = "aes256-cbc-sha1"` + "\n"
	config += `	  ike_group = "16"` + "\n"
	config += `	  ipsec_rekey_interval = 3600` + "\n"
	config += `	  ipsec_replay_window = 512` + "\n"
	config += `	  ipsec_ciphersuite = "aes256-cbc-sha512"` + "\n"
	config += `	  perfect_forward_secrecy = "group-16"` + "\n"
	config += `	  track_enable = true` + "\n"
	config += `	}]` + "\n"
	config += `	interface_pairs = [{` + "\n"
	config += `	  active_interface = "ipsec1"` + "\n"
	config += `	  active_interface_weight = 1` + "\n"
	config += `	  backup_interface = "None"` + "\n"
	config += `	  backup_interface_weight = 1` + "\n"
	config += `	}]` + "\n"
	config += `	region = "us-east-1"` + "\n"
	config += `	tracker_source_ip = "1.2.3.4"` + "\n"
	config += `	trackers = [{` + "\n"
	config += `	  name = "tracker1"` + "\n"
	config += `	  endpoint_api_url = "http://cisco.com"` + "\n"
	config += `	  threshold = 300` + "\n"
	config += `	  interval = 60` + "\n"
	config += `	  multiplier = 3` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"

	config += `
		data "sdwan_sse_cisco_feature" "test" {
			id = sdwan_sse_cisco_feature.test.id
			feature_profile_id = sdwan_sse_feature_profile.test.id
		}
	`
	return config
}

// End of section. //template:end testAccDataSourceConfig

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSourceByNameConfig
func testAccDataSourceSdwanSSECiscoProfileParcelByNameConfig() string {
	config := `resource "sdwan_sse_cisco_feature" "test" {` + "\n"
	config += ` name = "TF_TEST"` + "\n"
	config += ` description = "Terraform integration test"` + "\n"
	config += `	feature_profile_id = sdwan_sse_feature_profile.test.id` + "\n"
	config += `	context_sharing_for_vpn = false` + "\n"
	config += `	context_sharing_for_sgt = false` + "\n"
	config += `	interfaces = [{` + "\n"
	config += `	  interface_name = "ipsec1"` + "\n"
	config += `	  shutdown = false` + "\n"
	config += `	  tunnel_dc_preference = "primary-dc"` + "\n"
	config += `	  mtu = 1400` + "\n"
	config += `	  dpd_interval = 10` + "\n"
	config += `	  dpd_retries = 3` + "\n"
	config += `	  ike_version = 2` + "\n"
	config += `	  ike_rekey_interval = 14400` + "\n"
	config += `	  ike_ciphersuite = "aes256-cbc-sha1"` + "\n"
	config += `	  ike_group = "16"` + "\n"
	config += `	  ipsec_rekey_interval = 3600` + "\n"
	config += `	  ipsec_replay_window = 512` + "\n"
	config += `	  ipsec_ciphersuite = "aes256-cbc-sha512"` + "\n"
	config += `	  perfect_forward_secrecy = "group-16"` + "\n"
	config += `	  track_enable = true` + "\n"
	config += `	}]` + "\n"
	config += `	interface_pairs = [{` + "\n"
	config += `	  active_interface = "ipsec1"` + "\n"
	config += `	  active_interface_weight = 1` + "\n"
	config += `	  backup_interface = "None"` + "\n"
	config += `	  backup_interface_weight = 1` + "\n"
	config += `	}]` + "\n"
	config += `	region = "us-east-1"` + "\n"
	config += `	tracker_source_ip = "1.2.3.4"` + "\n"
	config += `	trackers = [{` + "\n"
	config += `	  name = "tracker1"` + "\n"
	config += `	  endpoint_api_url = "http://cisco.com"` + "\n"
	config += `	  threshold = 300` + "\n"
	config += `	  interval = 60` + "\n"
	config += `	  multiplier = 3` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"

	config += `
		data "sdwan_sse_cisco_feature" "test" {
			name = "TF_TEST"
			feature_profile_id = sdwan_sse_feature_profile.test.id
		}
	`
	return config
}

// End of section. //template:end testAccDataSourceByNameConfig
