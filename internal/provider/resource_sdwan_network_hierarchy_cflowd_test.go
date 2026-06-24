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
func TestAccSdwanNetworkHierarchyCflowd(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_network_hierarchy_cflowd.test", "flow_active_timeout", "600"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_network_hierarchy_cflowd.test", "flow_inactive_timeout", "60"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_network_hierarchy_cflowd.test", "flow_refresh_time", "600"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_network_hierarchy_cflowd.test", "flow_sampling_interval", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_network_hierarchy_cflowd.test", "collect_tloc_loopback", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_network_hierarchy_cflowd.test", "protocol", "ipv4"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_network_hierarchy_cflowd.test", "collect_tos", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_network_hierarchy_cflowd.test", "collect_dscp_output", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_network_hierarchy_cflowd.test", "collectors.0.vpn_id", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_network_hierarchy_cflowd.test", "collectors.0.address", "10.0.0.1"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_network_hierarchy_cflowd.test", "collectors.0.udp_port", "4739"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_network_hierarchy_cflowd.test", "collectors.0.export_spread", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_network_hierarchy_cflowd.test", "collectors.0.bfd_metrics_export", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_network_hierarchy_cflowd.test", "collectors.0.export_interval", "600"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccSdwanNetworkHierarchyCflowdConfig_all(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAcc

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites
// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAll
func testAccSdwanNetworkHierarchyCflowdConfig_all() string {
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
	config += `	}]` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigAll

func TestAccSdwanNetworkHierarchyCflowdMultipleCollectors(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccSdwanNetworkHierarchyCflowdConfig_multipleCollectors(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("sdwan_network_hierarchy_cflowd.test", "id"),
					resource.TestCheckResourceAttrSet("sdwan_network_hierarchy_cflowd.test", "node_id"),
					resource.TestCheckResourceAttr("sdwan_network_hierarchy_cflowd.test", "collectors.#", "3"),
					resource.TestCheckResourceAttr("sdwan_network_hierarchy_cflowd.test", "protocol", "both"),
					resource.TestCheckResourceAttr("sdwan_network_hierarchy_cflowd.test", "flow_active_timeout", "1800"),
					resource.TestCheckResourceAttr("sdwan_network_hierarchy_cflowd.test", "flow_inactive_timeout", "120"),
					resource.TestCheckResourceAttr("sdwan_network_hierarchy_cflowd.test", "flow_refresh_time", "3600"),
					resource.TestCheckResourceAttr("sdwan_network_hierarchy_cflowd.test", "flow_sampling_interval", "100"),
					resource.TestCheckResourceAttr("sdwan_network_hierarchy_cflowd.test", "collect_tloc_loopback", "false"),
					resource.TestCheckResourceAttr("sdwan_network_hierarchy_cflowd.test", "collect_tos", "false"),
					resource.TestCheckResourceAttr("sdwan_network_hierarchy_cflowd.test", "collect_dscp_output", "false"),
				),
			},
		},
	})
}

func testAccSdwanNetworkHierarchyCflowdConfig_multipleCollectors() string {
	config := `resource "sdwan_network_hierarchy_cflowd" "test" {` + "\n"
	config += `	flow_active_timeout = 1800` + "\n"
	config += `	flow_inactive_timeout = 120` + "\n"
	config += `	flow_refresh_time = 3600` + "\n"
	config += `	flow_sampling_interval = 100` + "\n"
	config += `	collect_tloc_loopback = false` + "\n"
	config += `	protocol = "both"` + "\n"
	config += `	collect_tos = false` + "\n"
	config += `	collect_dscp_output = false` + "\n"
	config += `	collectors = [{` + "\n"
	config += `	  vpn_id = 10` + "\n"
	config += `	  address = "192.168.1.1"` + "\n"
	config += `	  udp_port = 5000` + "\n"
	config += `	  export_spread = true` + "\n"
	config += `	  bfd_metrics_export = true` + "\n"
	config += `	  export_interval = 120` + "\n"
	config += `	}, {` + "\n"
	config += `	  vpn_id = 20` + "\n"
	config += `	  address = "192.168.1.2"` + "\n"
	config += `	  udp_port = 5001` + "\n"
	config += `	  export_spread = false` + "\n"
	config += `	  bfd_metrics_export = true` + "\n"
	config += `	  export_interval = 240` + "\n"
	config += `	}, {` + "\n"
	config += `	  vpn_id = 30` + "\n"
	config += `	  address = "2001:db8::1"` + "\n"
	config += `	  udp_port = 5002` + "\n"
	config += `	  export_spread = true` + "\n"
	config += `	  bfd_metrics_export = false` + "\n"
	config += `	  export_interval = 480` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"
	return config
}

func TestAccSdwanNetworkHierarchyCflowdProtocolIPv6(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccSdwanNetworkHierarchyCflowdConfig_ipv6(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("sdwan_network_hierarchy_cflowd.test", "id"),
					resource.TestCheckResourceAttr("sdwan_network_hierarchy_cflowd.test", "protocol", "ipv6"),
					resource.TestCheckResourceAttr("sdwan_network_hierarchy_cflowd.test", "collectors.#", "1"),
				),
			},
		},
	})
}

func testAccSdwanNetworkHierarchyCflowdConfig_ipv6() string {
	config := `resource "sdwan_network_hierarchy_cflowd" "test" {` + "\n"
	config += `	flow_active_timeout = 600` + "\n"
	config += `	protocol = "ipv6"` + "\n"
	config += `	collectors = [{` + "\n"
	config += `	  vpn_id = 1` + "\n"
	config += `	  address = "2001:db8::100"` + "\n"
	config += `	  udp_port = 4739` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"
	return config
}
