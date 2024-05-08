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
func TestAccDataSourceSdwanCEdgeGlobalFeatureTemplate(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cedge_global_feature_template.test", "nat64_udp_timeout", "300"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cedge_global_feature_template.test", "nat64_tcp_timeout", "3600"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cedge_global_feature_template.test", "http_authentication", "local"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cedge_global_feature_template.test", "ssh_version", "2"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cedge_global_feature_template.test", "http_server", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cedge_global_feature_template.test", "https_server", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cedge_global_feature_template.test", "source_interface", "e1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cedge_global_feature_template.test", "ip_source_routing", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cedge_global_feature_template.test", "arp_proxy", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cedge_global_feature_template.test", "ftp_passive", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cedge_global_feature_template.test", "rsh_rcp", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cedge_global_feature_template.test", "bootp", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cedge_global_feature_template.test", "domain_lookup", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cedge_global_feature_template.test", "tcp_keepalives_out", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cedge_global_feature_template.test", "tcp_keepalives_in", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cedge_global_feature_template.test", "tcp_small_servers", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cedge_global_feature_template.test", "udp_small_servers", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cedge_global_feature_template.test", "lldp", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cedge_global_feature_template.test", "cdp", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cedge_global_feature_template.test", "snmp_ifindex_persist", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cedge_global_feature_template.test", "console_logging", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cedge_global_feature_template.test", "vty_logging", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cedge_global_feature_template.test", "line_vty", "true"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceSdwanCEdgeGlobalFeatureTemplateConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAccDataSource

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSourceConfig
func testAccDataSourceSdwanCEdgeGlobalFeatureTemplateConfig() string {
	config := `resource "sdwan_cedge_global_feature_template" "test" {` + "\n"
	config += ` name = "TF_TEST"` + "\n"
	config += ` description = "Terraform integration test"` + "\n"
	config += ` device_types = ["vedge-C8000V"]` + "\n"
	config += `	nat64_udp_timeout = 300` + "\n"
	config += `	nat64_tcp_timeout = 3600` + "\n"
	config += `	http_authentication = "local"` + "\n"
	config += `	ssh_version = 2` + "\n"
	config += `	http_server = true` + "\n"
	config += `	https_server = true` + "\n"
	config += `	source_interface = "e1"` + "\n"
	config += `	ip_source_routing = true` + "\n"
	config += `	arp_proxy = true` + "\n"
	config += `	ftp_passive = true` + "\n"
	config += `	rsh_rcp = true` + "\n"
	config += `	bootp = true` + "\n"
	config += `	domain_lookup = true` + "\n"
	config += `	tcp_keepalives_out = true` + "\n"
	config += `	tcp_keepalives_in = true` + "\n"
	config += `	tcp_small_servers = true` + "\n"
	config += `	udp_small_servers = true` + "\n"
	config += `	lldp = true` + "\n"
	config += `	cdp = true` + "\n"
	config += `	snmp_ifindex_persist = true` + "\n"
	config += `	console_logging = true` + "\n"
	config += `	vty_logging = true` + "\n"
	config += `	line_vty = true` + "\n"
	config += `}` + "\n"

	config += `
		data "sdwan_cedge_global_feature_template" "test" {
			id = sdwan_cedge_global_feature_template.test.id
		}
	`
	return config
}

// End of section. //template:end testAccDataSourceConfig
