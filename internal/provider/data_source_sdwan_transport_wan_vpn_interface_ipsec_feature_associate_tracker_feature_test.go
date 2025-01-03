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
func TestAccDataSourceSdwanTransportWANVPNInterfaceIPSECFeatureAssociateTrackerFeature(t *testing.T) {
	if os.Getenv("SDWAN_2012") == "" {
		t.Skip("skipping test, set environment variable SDWAN_2012")
	}
	var checks []resource.TestCheckFunc
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceSdwanTransportWANVPNInterfaceIPSECFeatureAssociateTrackerFeaturePrerequisitesConfig + testAccDataSourceSdwanTransportWANVPNInterfaceIPSECFeatureAssociateTrackerFeatureConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAccDataSource

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites
const testAccDataSourceSdwanTransportWANVPNInterfaceIPSECFeatureAssociateTrackerFeaturePrerequisitesConfig = `
resource "sdwan_transport_feature_profile" "test" {
  name        = "TF_TEST"
  description = "Terraform test"
}

resource "sdwan_transport_wan_vpn_feature" "test" {
  name               = "TF_TEST_WAN_VPN"
  description        = "Terraform test"
  feature_profile_id = sdwan_transport_feature_profile.test.id
  vpn                = 0
}

resource "sdwan_transport_tracker_feature" "test" {
  name                  = "TF_TEST_TRACER"
  description           = "My Example"
  feature_profile_id    = sdwan_transport_feature_profile.test.id
  tracker_name          = "TRACKER_1"
  endpoint_api_url      = "google.com"
  endpoint_dns_name     = "google.com"
  endpoint_ip           = "1.2.3.4"
  interval              = 30
  multiplier            = 3
  threshold             = 300
  endpoint_tracker_type = "interface"
  tracker_type          = "endpoint"
}

resource "sdwan_transport_wan_vpn_interface_ipsec_feature" "test" {
  name                                = "TF_TEST_INTERFACE_IPSEC"
  description                         = "My Example"
  feature_profile_id                  = sdwan_transport_feature_profile.test.id
  transport_wan_vpn_feature_id        = sdwan_transport_wan_vpn_feature.test.id
  interface_name                      = "ipsec987"
  shutdown                            = true
  interface_description               = "ipsec987"
  ipv4_address                        = "9.7.5.4"
  ipv4_subnet_mask                    = "255.255.255.0"
  tunnel_source_ipv4_address          = "1.3.5.88"
  tunnel_source_ipv4_subnet_mask      = "255.255.255.0"
  tunnel_source_interface             = "GigabitEthernet8"
  tunnel_destination_ipv4_address     = "2.55.67.99"
  tunnel_destination_ipv4_subnet_mask = "255.255.255.0"
  application_tunnel_type             = "none"
  tcp_mss                             = 1460
  clear_dont_fragment                 = false
  ip_mtu                              = 1500
  dpd_interval                        = 10
  dpd_retries                         = 3
  ike_preshared_key                   = "123"
  ike_version                         = 1
  ike_integrity_protocol              = "main"
  ike_rekey_interval                  = 14400
  ike_ciphersuite                     = "aes256-cbc-sha1"
  ike_diffie_hellman_group            = "16"
  ike_id_local_end_point              = "xxx"
  ike_id_remote_end_point             = "xxx"
  ipsec_rekey_interval                = 3600
  ipsec_replay_window                 = 512
  ipsec_ciphersuite                   = "aes256-gcm"
  perfect_forward_secrecy             = "group-16"
  tunnel_route_via                    = "2222"
}
`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSourceConfig
func testAccDataSourceSdwanTransportWANVPNInterfaceIPSECFeatureAssociateTrackerFeatureConfig() string {
	config := ""
	config += `resource "sdwan_transport_wan_vpn_interface_ipsec_feature_associate_tracker_feature" "test" {` + "\n"
	config += `	feature_profile_id = sdwan_transport_feature_profile.test.id` + "\n"
	config += `	transport_wan_vpn_feature_id = sdwan_transport_wan_vpn_feature.test.id` + "\n"
	config += `	transport_wan_vpn_interface_ipsec_feature_id = sdwan_transport_wan_vpn_interface_ipsec_feature.test.id` + "\n"
	config += `	transport_tracker_feature_id = sdwan_transport_tracker_feature.test.id` + "\n"
	config += `}` + "\n"

	config += `
		data "sdwan_transport_wan_vpn_interface_ipsec_feature_associate_tracker_feature" "test" {
			feature_profile_id = sdwan_transport_feature_profile.test.id
			transport_wan_vpn_feature_id = sdwan_transport_wan_vpn_feature.test.id
			transport_wan_vpn_interface_ipsec_feature_id = sdwan_transport_wan_vpn_interface_ipsec_feature.test.id
			id = sdwan_transport_wan_vpn_interface_ipsec_feature_associate_tracker_feature.test.id
		}
	`
	return config
}

// End of section. //template:end testAccDataSourceConfig
