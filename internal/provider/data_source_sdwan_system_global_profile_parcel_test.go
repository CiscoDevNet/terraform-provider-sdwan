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

// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccDataSourceSdwanSystemGlobalProfileParcel(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceSdwanSystemGlobalProfileParcelConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.sdwan_system_global_profile_parcel.test", "http_server", "false"),
					resource.TestCheckResourceAttr("data.sdwan_system_global_profile_parcel.test", "https_server", "false"),
					resource.TestCheckResourceAttr("data.sdwan_system_global_profile_parcel.test", "ftp_passive", "false"),
					resource.TestCheckResourceAttr("data.sdwan_system_global_profile_parcel.test", "domain_lookup", "false"),
					resource.TestCheckResourceAttr("data.sdwan_system_global_profile_parcel.test", "arp_proxy", "false"),
					resource.TestCheckResourceAttr("data.sdwan_system_global_profile_parcel.test", "rsh_rcp", "false"),
					resource.TestCheckResourceAttr("data.sdwan_system_global_profile_parcel.test", "line_vty", "false"),
					resource.TestCheckResourceAttr("data.sdwan_system_global_profile_parcel.test", "cdp", "true"),
					resource.TestCheckResourceAttr("data.sdwan_system_global_profile_parcel.test", "lldp", "true"),
					resource.TestCheckResourceAttr("data.sdwan_system_global_profile_parcel.test", "source_interface", "GigabitEthernet0/0/1"),
					resource.TestCheckResourceAttr("data.sdwan_system_global_profile_parcel.test", "tcp_keepalives_in", "true"),
					resource.TestCheckResourceAttr("data.sdwan_system_global_profile_parcel.test", "tcp_keepalives_out", "true"),
					resource.TestCheckResourceAttr("data.sdwan_system_global_profile_parcel.test", "tcp_small_servers", "false"),
					resource.TestCheckResourceAttr("data.sdwan_system_global_profile_parcel.test", "udp_small_servers", "false"),
					resource.TestCheckResourceAttr("data.sdwan_system_global_profile_parcel.test", "console_logging", "true"),
					resource.TestCheckResourceAttr("data.sdwan_system_global_profile_parcel.test", "ip_source_routing", "false"),
					resource.TestCheckResourceAttr("data.sdwan_system_global_profile_parcel.test", "vty_line_logging", "false"),
					resource.TestCheckResourceAttr("data.sdwan_system_global_profile_parcel.test", "snmp_ifindex_persist", "true"),
					resource.TestCheckResourceAttr("data.sdwan_system_global_profile_parcel.test", "ignore_bootp", "true"),
					resource.TestCheckResourceAttr("data.sdwan_system_global_profile_parcel.test", "nat64_udp_timeout", "300"),
					resource.TestCheckResourceAttr("data.sdwan_system_global_profile_parcel.test", "nat64_tcp_timeout", "3600"),
					resource.TestCheckResourceAttr("data.sdwan_system_global_profile_parcel.test", "http_authentication", "aaa"),
					resource.TestCheckResourceAttr("data.sdwan_system_global_profile_parcel.test", "ssh_version", "2"),
				),
			},
		},
	})
}

const testAccDataSourceSdwanSystemGlobalProfileParcelConfig = `
resource "sdwan_system_feature_profile" "test" {
  name = "TF_TEST"
  description = "Terraform test"
}

resource "sdwan_system_global_profile_parcel" "test" {
  name = "TF_TEST"
  description = "Terraform integration test"
  feature_profile_id = sdwan_system_feature_profile.test.id
  http_server = false
  https_server = false
  ftp_passive = false
  domain_lookup = false
  arp_proxy = false
  rsh_rcp = false
  line_vty = false
  cdp = true
  lldp = true
  source_interface = "GigabitEthernet0/0/1"
  tcp_keepalives_in = true
  tcp_keepalives_out = true
  tcp_small_servers = false
  udp_small_servers = false
  console_logging = true
  ip_source_routing = false
  vty_line_logging = false
  snmp_ifindex_persist = true
  ignore_bootp = true
  nat64_udp_timeout = 300
  nat64_tcp_timeout = 3600
  http_authentication = "aaa"
  ssh_version = "2"
}

data "sdwan_system_global_profile_parcel" "test" {
  id = sdwan_system_global_profile_parcel.test.id
  feature_profile_id = sdwan_system_feature_profile.test.id
}
`
