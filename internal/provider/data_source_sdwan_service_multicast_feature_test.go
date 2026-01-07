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
func TestAccDataSourceSdwanServiceMulticastProfileParcel(t *testing.T) {
	if os.Getenv("SDWAN_2015_IN_PROGRESS") == "" {
		t.Skip("skipping test, set environment variable SDWAN_2015_IN_PROGRESS")
	}
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_service_multicast_feature.test", "spt_only", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_service_multicast_feature.test", "local_replicator", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_service_multicast_feature.test", "local_replicator_threshold", "10"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_service_multicast_feature.test", "igmp_interfaces.0.interface_name", "GigabitEthernet1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_service_multicast_feature.test", "igmp_interfaces.0.version", "2"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_service_multicast_feature.test", "igmp_interfaces.0.join_groups.0.group_address", "224.0.0.0"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_service_multicast_feature.test", "igmp_interfaces.0.join_groups.0.source_address", "1.2.3.4"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_service_multicast_feature.test", "pim_source_specific_multicast_enable", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_service_multicast_feature.test", "pim_source_specific_multicast_access_list", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_service_multicast_feature.test", "pim_spt_threshold", "0"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_service_multicast_feature.test", "pim_interfaces.0.interface_name", "GigabitEthernet1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_service_multicast_feature.test", "pim_interfaces.0.query_interval", "30"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_service_multicast_feature.test", "pim_interfaces.0.join_prune_interval", "60"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_service_multicast_feature.test", "static_rp_addresses.0.ip_address", "1.2.3.4"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_service_multicast_feature.test", "static_rp_addresses.0.access_list", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_service_multicast_feature.test", "static_rp_addresses.0.override", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_service_multicast_feature.test", "enable_auto_rp", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_service_multicast_feature.test", "pim_bsr_rp_candidates.0.interface_name", "GigabitEthernet1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_service_multicast_feature.test", "pim_bsr_rp_candidates.0.access_list_id", "2"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_service_multicast_feature.test", "pim_bsr_rp_candidates.0.interval", "30"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_service_multicast_feature.test", "pim_bsr_rp_candidates.0.priority", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_service_multicast_feature.test", "pim_bsr_candidates.0.interface_name", "GigabitEthernet1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_service_multicast_feature.test", "pim_bsr_candidates.0.hash_mask_length", "30"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_service_multicast_feature.test", "pim_bsr_candidates.0.priority", "120"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_service_multicast_feature.test", "pim_bsr_candidates.0.accept_candidate_access_list", "test"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_service_multicast_feature.test", "msdp_groups.0.mesh_group_name", "Example"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_service_multicast_feature.test", "msdp_groups.0.peers.0.peer_ip", "1.2.3.4"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_service_multicast_feature.test", "msdp_groups.0.peers.0.connection_source_interface", "GigabitEthernet1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_service_multicast_feature.test", "msdp_groups.0.peers.0.remote_as", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_service_multicast_feature.test", "msdp_groups.0.peers.0.keepalive_interval", "15"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_service_multicast_feature.test", "msdp_groups.0.peers.0.keepalive_hold_time", "30"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_service_multicast_feature.test", "msdp_groups.0.peers.0.sa_limit", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_service_multicast_feature.test", "msdp_groups.0.peers.0.default_peer", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_service_multicast_feature.test", "msdp_originator_id", "GigabitEthernet1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_service_multicast_feature.test", "msdp_connection_retry_interval", "30"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceSdwanServiceMulticastPrerequisitesProfileParcelConfig + testAccDataSourceSdwanServiceMulticastProfileParcelConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAccDataSource

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites
const testAccDataSourceSdwanServiceMulticastPrerequisitesProfileParcelConfig = `
resource "sdwan_service_feature_profile" "test" {
  name = "TF_TEST"
  description = "Terraform test"
}
`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSourceConfig
func testAccDataSourceSdwanServiceMulticastProfileParcelConfig() string {
	config := `resource "sdwan_service_multicast_feature" "test" {` + "\n"
	config += ` name = "TF_TEST"` + "\n"
	config += ` description = "Terraform integration test"` + "\n"
	config += `	feature_profile_id = sdwan_service_feature_profile.test.id` + "\n"
	config += `	spt_only = false` + "\n"
	config += `	local_replicator = false` + "\n"
	config += `	local_replicator_threshold = 10` + "\n"
	config += `	igmp_interfaces = [{` + "\n"
	config += `	  interface_name = "GigabitEthernet1"` + "\n"
	config += `	  version = 2` + "\n"
	config += `	  join_groups = [{` + "\n"
	config += `		group_address = "224.0.0.0"` + "\n"
	config += `		source_address = "1.2.3.4"` + "\n"
	config += `	}]` + "\n"
	config += `	}]` + "\n"
	config += `	pim_source_specific_multicast_enable = true` + "\n"
	config += `	pim_source_specific_multicast_access_list = "1"` + "\n"
	config += `	pim_spt_threshold = "0"` + "\n"
	config += `	pim_interfaces = [{` + "\n"
	config += `	  interface_name = "GigabitEthernet1"` + "\n"
	config += `	  query_interval = 30` + "\n"
	config += `	  join_prune_interval = 60` + "\n"
	config += `	}]` + "\n"
	config += `	static_rp_addresses = [{` + "\n"
	config += `	  ip_address = "1.2.3.4"` + "\n"
	config += `	  access_list = "1"` + "\n"
	config += `	  override = false` + "\n"
	config += `	}]` + "\n"
	config += `	enable_auto_rp = false` + "\n"
	config += `	pim_bsr_rp_candidates = [{` + "\n"
	config += `	  interface_name = "GigabitEthernet1"` + "\n"
	config += `	  access_list_id = "2"` + "\n"
	config += `	  interval = 30` + "\n"
	config += `	  priority = 1` + "\n"
	config += `	}]` + "\n"
	config += `	pim_bsr_candidates = [{` + "\n"
	config += `	  interface_name = "GigabitEthernet1"` + "\n"
	config += `	  hash_mask_length = 30` + "\n"
	config += `	  priority = 120` + "\n"
	config += `	  accept_candidate_access_list = "test"` + "\n"
	config += `	}]` + "\n"
	config += `	msdp_groups = [{` + "\n"
	config += `	  mesh_group_name = "Example"` + "\n"
	config += `	  peers = [{` + "\n"
	config += `		peer_ip = "1.2.3.4"` + "\n"
	config += `		connection_source_interface = "GigabitEthernet1"` + "\n"
	config += `		remote_as = 1` + "\n"
	config += `		peer_authentication_password = "Password123!"` + "\n"
	config += `		keepalive_interval = 15` + "\n"
	config += `		keepalive_hold_time = 30` + "\n"
	config += `		sa_limit = 1` + "\n"
	config += `		default_peer = false` + "\n"
	config += `	}]` + "\n"
	config += `	}]` + "\n"
	config += `	msdp_originator_id = "GigabitEthernet1"` + "\n"
	config += `	msdp_connection_retry_interval = 30` + "\n"
	config += `}` + "\n"

	config += `
		data "sdwan_service_multicast_feature" "test" {
			id = sdwan_service_multicast_feature.test.id
			feature_profile_id = sdwan_service_feature_profile.test.id
		}
	`
	return config
}

// End of section. //template:end testAccDataSourceConfig
