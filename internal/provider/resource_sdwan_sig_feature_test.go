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
func TestAccSdwanSigProfileParcel(t *testing.T) {
	if os.Getenv("SDWAN_2015") == "" && os.Getenv("SDWAN_2018") == "" {
		t.Skip("skipping test, set environment variable SDWAN_2015 or SDWAN_2018")
	}
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_sig_feature.test", "sig_provider", "Generic"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_sig_feature.test", "interfaces.0.interface_name", "gre1"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_sig_feature.test", "interfaces.0.interface_description", "Tunnel to SIG data center A"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_sig_feature.test", "interfaces.0.tunnel_source_interface", "GigabitEthernet1"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_sig_feature.test", "interfaces.0.tunnel_destination", "198.51.100.10"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_sig_feature.test", "interfaces.0.tunnel_set", "secure-internet-gateway-other"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_sig_feature.test", "interfaces.0.tunnel_dc_preference", "primary-dc"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_sig_feature.test", "interfaces.0.tcp_mss_adjust", "1300"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_sig_feature.test", "interfaces.0.mtu", "1400"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_sig_feature.test", "interfaces.0.tracker", "tracker1"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_sig_feature.test", "interfaces.0.track_enable", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_sig_feature.test", "interface_pairs.0.active_interface", "gre1"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_sig_feature.test", "interface_pairs.0.active_interface_weight", "255"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_sig_feature.test", "interface_pairs.0.backup_interface", "None"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_sig_feature.test", "interface_pairs.0.backup_interface_weight", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_sig_feature.test", "tracker_source_ip", "192.0.2.10"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_sig_feature.test", "trackers.0.name", "tracker1"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_sig_feature.test", "trackers.0.endpoint_api_url", "http://gateway.example.com/vpntest"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_sig_feature.test", "trackers.0.threshold", "1000"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_sig_feature.test", "trackers.0.interval", "30"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_sig_feature.test", "trackers.0.multiplier", "2"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccSdwanSigPrerequisitesProfileParcelConfig + testAccSdwanSigProfileParcelConfig_minimum(),
			},
			{
				Config: testAccSdwanSigPrerequisitesProfileParcelConfig + testAccSdwanSigProfileParcelConfig_all(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAcc

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites
const testAccSdwanSigPrerequisitesProfileParcelConfig = `
resource "sdwan_sig_security_feature_profile" "test" {
  name        = "TF_TEST_SIG_SECURITY"
  description = "Terraform test"
}

`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigMinimum
func testAccSdwanSigProfileParcelConfig_minimum() string {
	config := `resource "sdwan_sig_feature" "test" {` + "\n"
	config += ` name = "TF_TEST_MIN"` + "\n"
	config += ` description = "Terraform integration test"` + "\n"
	config += `	feature_profile_id = sdwan_sig_security_feature_profile.test.id` + "\n"
	config += `	sig_provider = "Generic"` + "\n"
	config += `	interfaces = [{` + "\n"
	config += `	  interface_name = "gre1"` + "\n"
	config += `	  tunnel_source_interface = "GigabitEthernet1"` + "\n"
	config += `	  mtu = 1400` + "\n"
	config += `	}]` + "\n"
	config += `	tracker_source_ip = "192.0.2.10"` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigMinimum

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAll
func testAccSdwanSigProfileParcelConfig_all() string {
	config := `resource "sdwan_sig_feature" "test" {` + "\n"
	config += ` name = "TF_TEST_ALL"` + "\n"
	config += ` description = "Terraform integration test"` + "\n"
	config += `	feature_profile_id = sdwan_sig_security_feature_profile.test.id` + "\n"
	config += `	sig_provider = "Generic"` + "\n"
	config += `	interfaces = [{` + "\n"
	config += `	  interface_name = "gre1"` + "\n"
	config += `	  interface_description = "Tunnel to SIG data center A"` + "\n"
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
	return config
}

// End of section. //template:end testAccConfigAll

// Verifies a Generic GRE build with an active/backup interface pair, a shared tracker and
// device-specific values passed as variables (tracker source IP, tunnel source and destination).
func TestAccSdwanSigProfileParcelGenericGreInterfacePair(t *testing.T) {
	if os.Getenv("SDWAN_2015") == "" {
		t.Skip("skipping test, set environment variable SDWAN_2015")
	}
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccSdwanSigPrerequisitesProfileParcelConfig + testAccSdwanSigProfileParcelConfig_genericGreInterfacePair(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("sdwan_sig_feature.test", "id"),
					resource.TestCheckResourceAttr("sdwan_sig_feature.test", "sig_provider", "Generic"),
					resource.TestCheckResourceAttr("sdwan_sig_feature.test", "tracker_source_ip_variable", "{{sig_tracker_src_ip}}"),
					resource.TestCheckNoResourceAttr("sdwan_sig_feature.test", "tracker_source_ip"),
					resource.TestCheckResourceAttr("sdwan_sig_feature.test", "interfaces.#", "2"),
					resource.TestCheckResourceAttr("sdwan_sig_feature.test", "interfaces.0.tunnel_source_interface_variable", "{{sig_wan1_if}}"),
					resource.TestCheckResourceAttr("sdwan_sig_feature.test", "interfaces.0.tunnel_destination_variable", "{{sig_dc_a_ip}}"),
					resource.TestCheckResourceAttr("sdwan_sig_feature.test", "interfaces.1.tunnel_dc_preference", "secondary-dc"),
					resource.TestCheckResourceAttr("sdwan_sig_feature.test", "interface_pairs.0.active_interface", "gre1"),
					resource.TestCheckResourceAttr("sdwan_sig_feature.test", "interface_pairs.0.active_interface_weight", "255"),
					resource.TestCheckResourceAttr("sdwan_sig_feature.test", "interface_pairs.0.backup_interface", "gre2"),
					resource.TestCheckResourceAttr("sdwan_sig_feature.test", "interface_pairs.0.backup_interface_weight", "10"),
					resource.TestCheckResourceAttr("sdwan_sig_feature.test", "trackers.0.name", "tracker1"),
				),
			},
		},
	})
}

func testAccSdwanSigProfileParcelConfig_genericGreInterfacePair() string {
	return `resource "sdwan_sig_feature" "test" {
  name               = "TF_TEST_GENERIC_GRE"
  description        = "Terraform integration test"
  feature_profile_id = sdwan_sig_security_feature_profile.test.id
  sig_provider       = "Generic"
  interfaces = [
    {
      interface_name                   = "gre1"
      interface_description            = "Tunnel to SIG data center A"
      tunnel_source_interface_variable = "{{sig_wan1_if}}"
      tunnel_destination_variable      = "{{sig_dc_a_ip}}"
      tunnel_set                       = "secure-internet-gateway-other"
      tunnel_dc_preference             = "primary-dc"
      tcp_mss_adjust                   = 1300
      mtu                              = 1400
      tracker                          = "tracker1"
      track_enable                     = true
    },
    {
      interface_name                   = "gre2"
      interface_description            = "Tunnel to SIG data center B"
      tunnel_source_interface_variable = "{{sig_wan1_if}}"
      tunnel_destination_variable      = "{{sig_dc_b_ip}}"
      tunnel_set                       = "secure-internet-gateway-other"
      tunnel_dc_preference             = "secondary-dc"
      tcp_mss_adjust                   = 1300
      mtu                              = 1400
      tracker                          = "tracker1"
      track_enable                     = true
    }
  ]
  interface_pairs = [
    {
      active_interface        = "gre1"
      active_interface_weight = 255
      backup_interface        = "gre2"
      backup_interface_weight = 10
    }
  ]
  tracker_source_ip_variable = "{{sig_tracker_src_ip}}"
  trackers = [
    {
      name             = "tracker1"
      endpoint_api_url = "http://gateway.example.com/vpntest"
      threshold        = 1000
      interval         = 30
      multiplier       = 2
    }
  ]
}
`
}
