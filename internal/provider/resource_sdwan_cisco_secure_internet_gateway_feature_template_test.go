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
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin testAcc
func TestAccSdwanCiscoSecureInternetGatewayFeatureTemplate(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_secure_internet_gateway_feature_template.test", "vpn_id", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_secure_internet_gateway_feature_template.test", "interfaces.0.name", "ipsec1"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_secure_internet_gateway_feature_template.test", "interfaces.0.auto_tunnel_mode", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_secure_internet_gateway_feature_template.test", "interfaces.0.shutdown", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_secure_internet_gateway_feature_template.test", "interfaces.0.description", "My Description"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_secure_internet_gateway_feature_template.test", "interfaces.0.ip_unnumbered", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_secure_internet_gateway_feature_template.test", "interfaces.0.ipv4_address", "1.2.3.4/24"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_secure_internet_gateway_feature_template.test", "interfaces.0.tunnel_source", "3.3.3.3"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_secure_internet_gateway_feature_template.test", "interfaces.0.tunnel_source_interface", "ge0/1"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_secure_internet_gateway_feature_template.test", "interfaces.0.tunnel_route_via", "ge0/2"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_secure_internet_gateway_feature_template.test", "interfaces.0.tunnel_destination", "3.4.5.6"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_secure_internet_gateway_feature_template.test", "interfaces.0.application", "sig"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_secure_internet_gateway_feature_template.test", "interfaces.0.sig_provider", "secure-internet-gateway-umbrella"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_secure_internet_gateway_feature_template.test", "interfaces.0.tunnel_dc_preference", "primary-dc"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_secure_internet_gateway_feature_template.test", "interfaces.0.tcp_mss", "1400"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_secure_internet_gateway_feature_template.test", "interfaces.0.mtu", "1500"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_secure_internet_gateway_feature_template.test", "interfaces.0.dead_peer_detection_interval", "30"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_secure_internet_gateway_feature_template.test", "interfaces.0.dead_peer_detection_retries", "5"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_secure_internet_gateway_feature_template.test", "interfaces.0.ike_version", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_secure_internet_gateway_feature_template.test", "interfaces.0.ike_pre_shared_key", "A1234567"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_secure_internet_gateway_feature_template.test", "interfaces.0.ike_rekey_interval", "600"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_secure_internet_gateway_feature_template.test", "interfaces.0.ike_ciphersuite", "aes256-cbc-sha2"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_secure_internet_gateway_feature_template.test", "interfaces.0.ike_group", "14"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_secure_internet_gateway_feature_template.test", "interfaces.0.ike_pre_shared_key_dynamic", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_secure_internet_gateway_feature_template.test", "interfaces.0.ike_pre_shared_key_local_id", "1.2.3.4"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_secure_internet_gateway_feature_template.test", "interfaces.0.ike_pre_shared_key_remote_id", "2.3.4.5"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_secure_internet_gateway_feature_template.test", "interfaces.0.ipsec_rekey_interval", "7200"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_secure_internet_gateway_feature_template.test", "interfaces.0.ipsec_replay_window", "1024"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_secure_internet_gateway_feature_template.test", "interfaces.0.ipsec_ciphersuite", "aes256-cbc-sha1"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_secure_internet_gateway_feature_template.test", "interfaces.0.ipsec_perfect_forward_secrecy", "group-14"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_secure_internet_gateway_feature_template.test", "interfaces.0.tracker", "test"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_secure_internet_gateway_feature_template.test", "interfaces.0.track_enable", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_secure_internet_gateway_feature_template.test", "interfaces.0.tunnel_public_ip", "5.5.5.5"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_secure_internet_gateway_feature_template.test", "services.0.service_type", "sig"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_secure_internet_gateway_feature_template.test", "services.0.interface_pairs.0.active_interface", "e1"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_secure_internet_gateway_feature_template.test", "services.0.interface_pairs.0.active_interface_weight", "10"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_secure_internet_gateway_feature_template.test", "services.0.interface_pairs.0.backup_interface", "e2"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_secure_internet_gateway_feature_template.test", "services.0.interface_pairs.0.backup_interface_weight", "20"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_secure_internet_gateway_feature_template.test", "services.0.zscaler_authentication_required", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_secure_internet_gateway_feature_template.test", "services.0.zscaler_xff_forward", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_secure_internet_gateway_feature_template.test", "services.0.zscaler_firewall_enabled", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_secure_internet_gateway_feature_template.test", "services.0.zscaler_ips_control_enabled", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_secure_internet_gateway_feature_template.test", "services.0.zscaler_caution_enabled", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_secure_internet_gateway_feature_template.test", "services.0.zscaler_primary_data_center", "Auto"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_secure_internet_gateway_feature_template.test", "services.0.zscaler_secondary_data_center", "Auto"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_secure_internet_gateway_feature_template.test", "services.0.zscaler_surrogate_ip", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_secure_internet_gateway_feature_template.test", "services.0.zscaler_surrogate_idle_time", "100"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_secure_internet_gateway_feature_template.test", "services.0.zscaler_surrogate_display_time_unit", "MINUTE"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_secure_internet_gateway_feature_template.test", "services.0.zscaler_surrogate_ip_enforce_for_known_browsers", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_secure_internet_gateway_feature_template.test", "services.0.zscaler_surrogate_refresh_time", "12345"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_secure_internet_gateway_feature_template.test", "services.0.zscaler_surrogate_refresh_time_unit", "MINUTE"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_secure_internet_gateway_feature_template.test", "services.0.zscaler_aup_enabled", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_secure_internet_gateway_feature_template.test", "services.0.zscaler_aup_block_internet_until_accepted", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_secure_internet_gateway_feature_template.test", "services.0.zscaler_aup_force_ssl_inspection", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_secure_internet_gateway_feature_template.test", "services.0.zscaler_aup_timeout", "60"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_secure_internet_gateway_feature_template.test", "services.0.zscaler_location_name", "LOC1"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_secure_internet_gateway_feature_template.test", "services.0.umbrella_primary_data_center", "Auto"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_secure_internet_gateway_feature_template.test", "services.0.umbrella_secondary_data_center", "Auto"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_secure_internet_gateway_feature_template.test", "tracker_source_ip", "2.3.4.5"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_secure_internet_gateway_feature_template.test", "trackers.0.name", "TRACKER1"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_secure_internet_gateway_feature_template.test", "trackers.0.endpoint_api_url", "https://1.1.1.1"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_secure_internet_gateway_feature_template.test", "trackers.0.threshold", "500"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_secure_internet_gateway_feature_template.test", "trackers.0.interval", "60"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_secure_internet_gateway_feature_template.test", "trackers.0.multiplier", "4"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_secure_internet_gateway_feature_template.test", "trackers.0.tracker_type", "SIG"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccSdwanCiscoSecureInternetGatewayFeatureTemplateConfig_minimum(),
			},
			{
				Config: testAccSdwanCiscoSecureInternetGatewayFeatureTemplateConfig_all(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAcc

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites
// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigMinimum
func testAccSdwanCiscoSecureInternetGatewayFeatureTemplateConfig_minimum() string {
	config := `resource "sdwan_cisco_secure_internet_gateway_feature_template" "test" {` + "\n"
	config += ` name = "TF_TEST_MIN"` + "\n"
	config += ` description = "Terraform integration test"` + "\n"
	config += ` device_types = ["vedge-C8000V"]` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigMinimum

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAll
func testAccSdwanCiscoSecureInternetGatewayFeatureTemplateConfig_all() string {
	config := `resource "sdwan_cisco_secure_internet_gateway_feature_template" "test" {` + "\n"
	config += ` name = "TF_TEST_ALL"` + "\n"
	config += ` description = "Terraform integration test"` + "\n"
	config += ` device_types = ["vedge-C8000V"]` + "\n"
	config += `	vpn_id = 1` + "\n"
	config += `	interfaces = [{` + "\n"
	config += `	  name = "ipsec1"` + "\n"
	config += `	  auto_tunnel_mode = true` + "\n"
	config += `	  shutdown = true` + "\n"
	config += `	  description = "My Description"` + "\n"
	config += `	  ip_unnumbered = false` + "\n"
	config += `	  ipv4_address = "1.2.3.4/24"` + "\n"
	config += `	  tunnel_source = "3.3.3.3"` + "\n"
	config += `	  tunnel_source_interface = "ge0/1"` + "\n"
	config += `	  tunnel_route_via = "ge0/2"` + "\n"
	config += `	  tunnel_destination = "3.4.5.6"` + "\n"
	config += `	  application = "sig"` + "\n"
	config += `	  sig_provider = "secure-internet-gateway-umbrella"` + "\n"
	config += `	  tunnel_dc_preference = "primary-dc"` + "\n"
	config += `	  tcp_mss = 1400` + "\n"
	config += `	  mtu = 1500` + "\n"
	config += `	  dead_peer_detection_interval = 30` + "\n"
	config += `	  dead_peer_detection_retries = 5` + "\n"
	config += `	  ike_version = 1` + "\n"
	config += `	  ike_pre_shared_key = "A1234567"` + "\n"
	config += `	  ike_rekey_interval = 600` + "\n"
	config += `	  ike_ciphersuite = "aes256-cbc-sha2"` + "\n"
	config += `	  ike_group = "14"` + "\n"
	config += `	  ike_pre_shared_key_dynamic = false` + "\n"
	config += `	  ike_pre_shared_key_local_id = "1.2.3.4"` + "\n"
	config += `	  ike_pre_shared_key_remote_id = "2.3.4.5"` + "\n"
	config += `	  ipsec_rekey_interval = 7200` + "\n"
	config += `	  ipsec_replay_window = 1024` + "\n"
	config += `	  ipsec_ciphersuite = "aes256-cbc-sha1"` + "\n"
	config += `	  ipsec_perfect_forward_secrecy = "group-14"` + "\n"
	config += `	  tracker = "test"` + "\n"
	config += `	  track_enable = false` + "\n"
	config += `	  tunnel_public_ip = "5.5.5.5"` + "\n"
	config += `	}]` + "\n"
	config += `	services = [{` + "\n"
	config += `	  service_type = "sig"` + "\n"
	config += `	  interface_pairs = [{` + "\n"
	config += `		active_interface = "e1"` + "\n"
	config += `		active_interface_weight = 10` + "\n"
	config += `		backup_interface = "e2"` + "\n"
	config += `		backup_interface_weight = 20` + "\n"
	config += `	}]` + "\n"
	config += `	  zscaler_authentication_required = true` + "\n"
	config += `	  zscaler_xff_forward = true` + "\n"
	config += `	  zscaler_firewall_enabled = true` + "\n"
	config += `	  zscaler_ips_control_enabled = true` + "\n"
	config += `	  zscaler_caution_enabled = true` + "\n"
	config += `	  zscaler_primary_data_center = "Auto"` + "\n"
	config += `	  zscaler_secondary_data_center = "Auto"` + "\n"
	config += `	  zscaler_surrogate_ip = true` + "\n"
	config += `	  zscaler_surrogate_idle_time = 100` + "\n"
	config += `	  zscaler_surrogate_display_time_unit = "MINUTE"` + "\n"
	config += `	  zscaler_surrogate_ip_enforce_for_known_browsers = true` + "\n"
	config += `	  zscaler_surrogate_refresh_time = 12345` + "\n"
	config += `	  zscaler_surrogate_refresh_time_unit = "MINUTE"` + "\n"
	config += `	  zscaler_aup_enabled = true` + "\n"
	config += `	  zscaler_aup_block_internet_until_accepted = true` + "\n"
	config += `	  zscaler_aup_force_ssl_inspection = true` + "\n"
	config += `	  zscaler_aup_timeout = 60` + "\n"
	config += `	  zscaler_location_name = "LOC1"` + "\n"
	config += `	  umbrella_primary_data_center = "Auto"` + "\n"
	config += `	  umbrella_secondary_data_center = "Auto"` + "\n"
	config += `	}]` + "\n"
	config += `	tracker_source_ip = "2.3.4.5"` + "\n"
	config += `	trackers = [{` + "\n"
	config += `	  name = "TRACKER1"` + "\n"
	config += `	  endpoint_api_url = "https://1.1.1.1"` + "\n"
	config += `	  threshold = 500` + "\n"
	config += `	  interval = 60` + "\n"
	config += `	  multiplier = 4` + "\n"
	config += `	  tracker_type = "SIG"` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigAll
