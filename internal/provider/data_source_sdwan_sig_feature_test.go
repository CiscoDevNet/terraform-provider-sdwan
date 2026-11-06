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
func TestAccDataSourceSdwanSigProfileParcel(t *testing.T) {
	if os.Getenv("SDWAN_2015") == "" && os.Getenv("SDWAN_2018") == "" {
		t.Skip("skipping test, set environment variable SDWAN_2015 or SDWAN_2018")
	}
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_sig_feature.test", "sig_provider", "Generic"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_sig_feature.test", "interfaces.0.interface_name", "gre1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_sig_feature.test", "interfaces.0.auto", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_sig_feature.test", "interfaces.0.interface_description", "Tunnel to SIG data center A"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_sig_feature.test", "interfaces.0.unnumbered", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_sig_feature.test", "interfaces.0.tunnel_source_interface", "GigabitEthernet1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_sig_feature.test", "interfaces.0.tunnel_destination", "198.51.100.10"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_sig_feature.test", "interfaces.0.tunnel_set", "secure-internet-gateway-other"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_sig_feature.test", "interfaces.0.tunnel_dc_preference", "primary-dc"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_sig_feature.test", "interfaces.0.tcp_mss_adjust", "1300"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_sig_feature.test", "interfaces.0.mtu", "1400"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_sig_feature.test", "interfaces.0.tracker", "tracker1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_sig_feature.test", "interfaces.0.track_enable", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_sig_feature.test", "interface_pairs.0.active_interface", "gre1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_sig_feature.test", "interface_pairs.0.active_interface_weight", "255"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_sig_feature.test", "interface_pairs.0.backup_interface", "None"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_sig_feature.test", "interface_pairs.0.backup_interface_weight", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_sig_feature.test", "tracker_source_ip", "192.0.2.10"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_sig_feature.test", "trackers.0.name", "tracker1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_sig_feature.test", "trackers.0.endpoint_api_url", "http://gateway.example.com/vpntest"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_sig_feature.test", "trackers.0.threshold", "1000"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_sig_feature.test", "trackers.0.interval", "30"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_sig_feature.test", "trackers.0.multiplier", "2"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceSdwanSigPrerequisitesProfileParcelConfig + testAccDataSourceSdwanSigProfileParcelConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
			{
				Config: testAccDataSourceSdwanSigPrerequisitesProfileParcelConfig + testAccDataSourceSdwanSigProfileParcelByNameConfig(),
				Check: resource.ComposeTestCheckFunc(
					append(checks,
						resource.TestCheckResourceAttr("data.sdwan_sig_feature.test", "name", "TF_TEST"),
						resource.TestCheckResourceAttrSet("data.sdwan_sig_feature.test", "id"),
					)...),
			},
		},
	})
}

// End of section. //template:end testAccDataSource

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites
const testAccDataSourceSdwanSigPrerequisitesProfileParcelConfig = `
resource "sdwan_sig_security_feature_profile" "test" {
  name        = "TF_TEST_SIG_SECURITY"
  description = "Terraform test"
}

`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSourceConfig
func testAccDataSourceSdwanSigProfileParcelConfig() string {
	config := `resource "sdwan_sig_feature" "test" {` + "\n"
	config += ` name = "TF_TEST"` + "\n"
	config += ` description = "Terraform integration test"` + "\n"
	config += `	feature_profile_id = sdwan_sig_security_feature_profile.test.id` + "\n"
	config += `	sig_provider = "Generic"` + "\n"
	config += `	interfaces = [{` + "\n"
	config += `	  interface_name = "gre1"` + "\n"
	config += `	  auto = false` + "\n"
	config += `	  interface_description = "Tunnel to SIG data center A"` + "\n"
	config += `	  unnumbered = true` + "\n"
	config += `	  tunnel_source_interface = "GigabitEthernet1"` + "\n"
	config += `	  tunnel_destination = "198.51.100.10"` + "\n"
	config += `	  tunnel_set = "secure-internet-gateway-other"` + "\n"
	config += `	  tunnel_dc_preference = "primary-dc"` + "\n"
	config += `	  tcp_mss_adjust = 1300` + "\n"
	config += `	  mtu = 1400` + "\n"
	config += `	  tracker = "tracker1"` + "\n"
	config += `	  track_enable = true` + "\n"
	config += `	}]` + "\n"
	config += `	interface_pairs = [{` + "\n"
	config += `	  active_interface = "gre1"` + "\n"
	config += `	  active_interface_weight = 255` + "\n"
	config += `	  backup_interface = "None"` + "\n"
	config += `	  backup_interface_weight = 1` + "\n"
	config += `	}]` + "\n"
	config += `	tracker_source_ip = "192.0.2.10"` + "\n"
	config += `	trackers = [{` + "\n"
	config += `	  name = "tracker1"` + "\n"
	config += `	  endpoint_api_url = "http://gateway.example.com/vpntest"` + "\n"
	config += `	  threshold = 1000` + "\n"
	config += `	  interval = 30` + "\n"
	config += `	  multiplier = 2` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"

	config += `
		data "sdwan_sig_feature" "test" {
			id = sdwan_sig_feature.test.id
			feature_profile_id = sdwan_sig_security_feature_profile.test.id
		}
	`
	return config
}

// End of section. //template:end testAccDataSourceConfig

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSourceByNameConfig
func testAccDataSourceSdwanSigProfileParcelByNameConfig() string {
	config := `resource "sdwan_sig_feature" "test" {` + "\n"
	config += ` name = "TF_TEST"` + "\n"
	config += ` description = "Terraform integration test"` + "\n"
	config += `	feature_profile_id = sdwan_sig_security_feature_profile.test.id` + "\n"
	config += `	sig_provider = "Generic"` + "\n"
	config += `	interfaces = [{` + "\n"
	config += `	  interface_name = "gre1"` + "\n"
	config += `	  auto = false` + "\n"
	config += `	  interface_description = "Tunnel to SIG data center A"` + "\n"
	config += `	  unnumbered = true` + "\n"
	config += `	  tunnel_source_interface = "GigabitEthernet1"` + "\n"
	config += `	  tunnel_destination = "198.51.100.10"` + "\n"
	config += `	  tunnel_set = "secure-internet-gateway-other"` + "\n"
	config += `	  tunnel_dc_preference = "primary-dc"` + "\n"
	config += `	  tcp_mss_adjust = 1300` + "\n"
	config += `	  mtu = 1400` + "\n"
	config += `	  tracker = "tracker1"` + "\n"
	config += `	  track_enable = true` + "\n"
	config += `	}]` + "\n"
	config += `	interface_pairs = [{` + "\n"
	config += `	  active_interface = "gre1"` + "\n"
	config += `	  active_interface_weight = 255` + "\n"
	config += `	  backup_interface = "None"` + "\n"
	config += `	  backup_interface_weight = 1` + "\n"
	config += `	}]` + "\n"
	config += `	tracker_source_ip = "192.0.2.10"` + "\n"
	config += `	trackers = [{` + "\n"
	config += `	  name = "tracker1"` + "\n"
	config += `	  endpoint_api_url = "http://gateway.example.com/vpntest"` + "\n"
	config += `	  threshold = 1000` + "\n"
	config += `	  interval = 30` + "\n"
	config += `	  multiplier = 2` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"

	config += `
		data "sdwan_sig_feature" "test" {
			name = "TF_TEST"
			feature_profile_id = sdwan_sig_security_feature_profile.test.id
		}
	`
	return config
}

// End of section. //template:end testAccDataSourceByNameConfig
