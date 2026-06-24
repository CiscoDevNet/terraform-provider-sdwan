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
			id = sdwan_network_hierarchy_cflowd.test.id
		}
	`
	return config
}

// End of section. //template:end testAccDataSourceConfig

// Custom data source test - creates resource first, then reads via data source

func TestAccDataSourceSdwanNetworkHierarchyCflowdWithResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceSdwanNetworkHierarchyCflowdConfig_withResource(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.sdwan_network_hierarchy_cflowd.test", "id"),
					resource.TestCheckResourceAttr("data.sdwan_network_hierarchy_cflowd.test", "flow_active_timeout", "600"),
					resource.TestCheckResourceAttr("data.sdwan_network_hierarchy_cflowd.test", "flow_inactive_timeout", "60"),
					resource.TestCheckResourceAttr("data.sdwan_network_hierarchy_cflowd.test", "flow_refresh_time", "600"),
					resource.TestCheckResourceAttr("data.sdwan_network_hierarchy_cflowd.test", "flow_sampling_interval", "1"),
					resource.TestCheckResourceAttr("data.sdwan_network_hierarchy_cflowd.test", "collect_tloc_loopback", "true"),
					resource.TestCheckResourceAttr("data.sdwan_network_hierarchy_cflowd.test", "protocol", "ipv4"),
					resource.TestCheckResourceAttr("data.sdwan_network_hierarchy_cflowd.test", "collect_tos", "true"),
					resource.TestCheckResourceAttr("data.sdwan_network_hierarchy_cflowd.test", "collect_dscp_output", "true"),
					resource.TestCheckResourceAttr("data.sdwan_network_hierarchy_cflowd.test", "collectors.#", "2"),
				),
			},
		},
	})
}

func testAccDataSourceSdwanNetworkHierarchyCflowdConfig_withResource() string {
	config := `resource "sdwan_network_hierarchy_cflowd" "test" {` + "\n"
	config += `	flow_active_timeout = 600` + "\n"
	config += `	flow_inactive_timeout = 60` + "\n"
	config += `	flow_refresh_time = 600` + "\n"
	config += `	flow_sampling_interval = 1` + "\n"
	config += `	collect_tloc_loopback = true` + "\n"
	config += `	protocol = "ipv4"` + "\n"
	config += `	collect_tos = true` + "\n"
	config += `	collect_dscp_output = true` + "\n"
	config += `	collectors = [{` + "\n"
	config += `	  vpn_id = 1` + "\n"
	config += `	  address = "10.0.0.1"` + "\n"
	config += `	  udp_port = 4739` + "\n"
	config += `	  export_spread = true` + "\n"
	config += `	  bfd_metrics_export = true` + "\n"
	config += `	  export_interval = 600` + "\n"
	config += `	}, {` + "\n"
	config += `	  vpn_id = 2` + "\n"
	config += `	  address = "10.0.0.2"` + "\n"
	config += `	  udp_port = 4740` + "\n"
	config += `	  export_spread = false` + "\n"
	config += `	  bfd_metrics_export = false` + "\n"
	config += `	  export_interval = 300` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"
	config += "\n"
	config += `data "sdwan_network_hierarchy_cflowd" "test" {` + "\n"
	config += `	id = sdwan_network_hierarchy_cflowd.test.id` + "\n"
	config += `}` + "\n"
	return config
}
