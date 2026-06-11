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

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSource
func TestAccDataSourceSdwanNetworkHierarchyCflowd(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_network_hierarchy_cflowd.test", "node_id", "5333c142-394b-47a7-afa6-760c44ca3cb5"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_network_hierarchy_cflowd.test", "flow_active_timeout", "600"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_network_hierarchy_cflowd.test", "flow_inactive_timeout", "60"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_network_hierarchy_cflowd.test", "flow_refresh_time", "600"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_network_hierarchy_cflowd.test", "flow_sampling_interval", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_network_hierarchy_cflowd.test", "collect_tloc_loopback", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_network_hierarchy_cflowd.test", "protocol", "ipv4"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_network_hierarchy_cflowd.test", "collect_tos", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_network_hierarchy_cflowd.test", "collect_dscp_output", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_network_hierarchy_cflowd.test", "collectors.0.vpn_id", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_network_hierarchy_cflowd.test", "collectors.0.address", "10.0.0.1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_network_hierarchy_cflowd.test", "collectors.0.udp_port", "4739"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_network_hierarchy_cflowd.test", "collectors.0.export_spread", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_network_hierarchy_cflowd.test", "collectors.0.bfd_metrics_export", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_network_hierarchy_cflowd.test", "collectors.0.export_interval", "600"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceSdwanNetworkHierarchyCflowdConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAccDataSource

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites
// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSourceConfig
func testAccDataSourceSdwanNetworkHierarchyCflowdConfig() string {
	config := ""

	config += `
		data "sdwan_network_hierarchy_cflowd" "test" {
			node_id = "5333c142-394b-47a7-afa6-760c44ca3cb5"
			id = sdwan_network_hierarchy_cflowd.test.id
		}
	`
	return config
}

// End of section. //template:end testAccDataSourceConfig
