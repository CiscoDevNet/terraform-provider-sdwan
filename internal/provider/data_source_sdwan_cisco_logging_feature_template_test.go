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
func TestAccDataSourceSdwanCiscoLoggingFeatureTemplate(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_logging_feature_template.test", "disk_logging", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_logging_feature_template.test", "max_size", "10"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_logging_feature_template.test", "log_rotations", "10"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_logging_feature_template.test", "tls_profiles.0.name", "PROF1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_logging_feature_template.test", "tls_profiles.0.version", "TLSv1.2"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_logging_feature_template.test", "tls_profiles.0.authentication_type", "Server"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_logging_feature_template.test", "ipv4_servers.0.hostname_ip", "2.2.2.2"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_logging_feature_template.test", "ipv4_servers.0.vpn_id", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_logging_feature_template.test", "ipv4_servers.0.source_interface", "e1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_logging_feature_template.test", "ipv4_servers.0.logging_level", "information"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_logging_feature_template.test", "ipv4_servers.0.enable_tls", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_logging_feature_template.test", "ipv4_servers.0.custom_profile", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_logging_feature_template.test", "ipv4_servers.0.profile", "PROF1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_logging_feature_template.test", "ipv6_servers.0.hostname_ip", "2001::1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_logging_feature_template.test", "ipv6_servers.0.vpn_id", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_logging_feature_template.test", "ipv6_servers.0.source_interface", "e1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_logging_feature_template.test", "ipv6_servers.0.logging_level", "information"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_logging_feature_template.test", "ipv6_servers.0.enable_tls", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_logging_feature_template.test", "ipv6_servers.0.custom_profile", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_logging_feature_template.test", "ipv6_servers.0.profile", "PROF1"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceSdwanCiscoLoggingFeatureTemplateConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAccDataSource

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSourceConfig
func testAccDataSourceSdwanCiscoLoggingFeatureTemplateConfig() string {
	config := `resource "sdwan_cisco_logging_feature_template" "test" {` + "\n"
	config += ` name = "TF_TEST"` + "\n"
	config += ` description = "Terraform integration test"` + "\n"
	config += ` device_types = ["vedge-C8000V"]` + "\n"
	config += `	disk_logging = true` + "\n"
	config += `	max_size = 10` + "\n"
	config += `	log_rotations = 10` + "\n"
	config += `	tls_profiles = [{` + "\n"
	config += `	  name = "PROF1"` + "\n"
	config += `	  version = "TLSv1.2"` + "\n"
	config += `	  authentication_type = "Server"` + "\n"
	config += `	  ciphersuite_list = ["aes-128-cbc-sha"]` + "\n"
	config += `	}]` + "\n"
	config += `	ipv4_servers = [{` + "\n"
	config += `	  hostname_ip = "2.2.2.2"` + "\n"
	config += `	  vpn_id = 1` + "\n"
	config += `	  source_interface = "e1"` + "\n"
	config += `	  logging_level = "information"` + "\n"
	config += `	  enable_tls = true` + "\n"
	config += `	  custom_profile = true` + "\n"
	config += `	  profile = "PROF1"` + "\n"
	config += `	}]` + "\n"
	config += `	ipv6_servers = [{` + "\n"
	config += `	  hostname_ip = "2001::1"` + "\n"
	config += `	  vpn_id = 1` + "\n"
	config += `	  source_interface = "e1"` + "\n"
	config += `	  logging_level = "information"` + "\n"
	config += `	  enable_tls = true` + "\n"
	config += `	  custom_profile = true` + "\n"
	config += `	  profile = "PROF1"` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"

	config += `
		data "sdwan_cisco_logging_feature_template" "test" {
			id = sdwan_cisco_logging_feature_template.test.id
		}
	`
	return config
}

// End of section. //template:end testAccDataSourceConfig
