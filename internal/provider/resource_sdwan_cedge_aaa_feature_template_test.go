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
func TestAccSdwanCEdgeAAAFeatureTemplate(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cedge_aaa_feature_template.test", "dot1x_authentication", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cedge_aaa_feature_template.test", "dot1x_accounting", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cedge_aaa_feature_template.test", "server_groups_priority_order", "100"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cedge_aaa_feature_template.test", "users.0.name", "user1"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cedge_aaa_feature_template.test", "users.0.password", "password123"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cedge_aaa_feature_template.test", "users.0.secret", "secret123"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cedge_aaa_feature_template.test", "users.0.privilege_level", "15"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cedge_aaa_feature_template.test", "users.0.ssh_pubkeys.0.key_string", "abc123"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cedge_aaa_feature_template.test", "users.0.ssh_pubkeys.0.key_type", "rsa"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cedge_aaa_feature_template.test", "radius_server_groups.0.group_name", "GROUP1"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cedge_aaa_feature_template.test", "radius_server_groups.0.vpn_id", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cedge_aaa_feature_template.test", "radius_server_groups.0.source_interface", "e1"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cedge_aaa_feature_template.test", "radius_server_groups.0.servers.0.address", "1.1.1.1"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cedge_aaa_feature_template.test", "radius_server_groups.0.servers.0.authentication_port", "1812"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cedge_aaa_feature_template.test", "radius_server_groups.0.servers.0.accounting_port", "1813"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cedge_aaa_feature_template.test", "radius_server_groups.0.servers.0.timeout", "5"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cedge_aaa_feature_template.test", "radius_server_groups.0.servers.0.retransmit", "3"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cedge_aaa_feature_template.test", "radius_server_groups.0.servers.0.key", "key123"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cedge_aaa_feature_template.test", "radius_server_groups.0.servers.0.secret_key", "1234567"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cedge_aaa_feature_template.test", "radius_server_groups.0.servers.0.encryption_type", "7"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cedge_aaa_feature_template.test", "radius_server_groups.0.servers.0.key_type", "pac"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cedge_aaa_feature_template.test", "radius_clients.0.client_ip", "2.2.2.2"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cedge_aaa_feature_template.test", "radius_clients.0.vpn_configurations.0.vpn_id", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cedge_aaa_feature_template.test", "radius_clients.0.vpn_configurations.0.server_key", "key123"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cedge_aaa_feature_template.test", "radius_dynamic_author_server_key", "key123"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cedge_aaa_feature_template.test", "radius_dynamic_author_domain_stripping", "yes"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cedge_aaa_feature_template.test", "radius_dynamic_author_authentication_type", "all"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cedge_aaa_feature_template.test", "radius_dynamic_author_port", "1700"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cedge_aaa_feature_template.test", "radius_trustsec_cts_authorization_list", "ALIST1"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cedge_aaa_feature_template.test", "radius_trustsec_group", "GROUP1"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cedge_aaa_feature_template.test", "tacacs_server_groups.0.group_name", "GROUP1"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cedge_aaa_feature_template.test", "tacacs_server_groups.0.vpn_id", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cedge_aaa_feature_template.test", "tacacs_server_groups.0.source_interface", "e1"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cedge_aaa_feature_template.test", "tacacs_server_groups.0.servers.0.address", "1.1.1.1"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cedge_aaa_feature_template.test", "tacacs_server_groups.0.servers.0.port", "49"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cedge_aaa_feature_template.test", "tacacs_server_groups.0.servers.0.timeout", "5"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cedge_aaa_feature_template.test", "tacacs_server_groups.0.servers.0.key", "key123"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cedge_aaa_feature_template.test", "tacacs_server_groups.0.servers.0.secret_key", "1234567"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cedge_aaa_feature_template.test", "tacacs_server_groups.0.servers.0.encryption_type", "7"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cedge_aaa_feature_template.test", "accounting_rules.0.name", "RULE1"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cedge_aaa_feature_template.test", "accounting_rules.0.method", "exec"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cedge_aaa_feature_template.test", "accounting_rules.0.privilege_level", "15"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cedge_aaa_feature_template.test", "accounting_rules.0.start_stop", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cedge_aaa_feature_template.test", "accounting_rules.0.groups", "GROUP1"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cedge_aaa_feature_template.test", "authorization_console", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cedge_aaa_feature_template.test", "authorization_config_commands", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cedge_aaa_feature_template.test", "authorization_rules.0.name", "RULE1"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cedge_aaa_feature_template.test", "authorization_rules.0.method", "commands"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cedge_aaa_feature_template.test", "authorization_rules.0.privilege_level", "15"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cedge_aaa_feature_template.test", "authorization_rules.0.groups", "GROUP1"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cedge_aaa_feature_template.test", "authorization_rules.0.authenticated", "true"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccSdwanCEdgeAAAFeatureTemplateConfig_minimum(),
			},
			{
				Config: testAccSdwanCEdgeAAAFeatureTemplateConfig_all(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAcc

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites
// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigMinimum
func testAccSdwanCEdgeAAAFeatureTemplateConfig_minimum() string {
	config := `resource "sdwan_cedge_aaa_feature_template" "test" {` + "\n"
	config += ` name = "TF_TEST_MIN"` + "\n"
	config += ` description = "Terraform integration test"` + "\n"
	config += ` device_types = ["vedge-C8000V"]` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigMinimum

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAll
func testAccSdwanCEdgeAAAFeatureTemplateConfig_all() string {
	config := `resource "sdwan_cedge_aaa_feature_template" "test" {` + "\n"
	config += ` name = "TF_TEST_ALL"` + "\n"
	config += ` description = "Terraform integration test"` + "\n"
	config += ` device_types = ["vedge-C8000V"]` + "\n"
	config += `	dot1x_authentication = true` + "\n"
	config += `	dot1x_accounting = true` + "\n"
	config += `	server_groups_priority_order = "100"` + "\n"
	config += `	users = [{` + "\n"
	config += `	  name = "user1"` + "\n"
	config += `	  password = "password123"` + "\n"
	config += `	  secret = "secret123"` + "\n"
	config += `	  privilege_level = "15"` + "\n"
	config += `	  ssh_pubkeys = [{` + "\n"
	config += `		key_string = "abc123"` + "\n"
	config += `		key_type = "rsa"` + "\n"
	config += `	}]` + "\n"
	config += `	}]` + "\n"
	config += `	radius_server_groups = [{` + "\n"
	config += `	  group_name = "GROUP1"` + "\n"
	config += `	  vpn_id = 1` + "\n"
	config += `	  source_interface = "e1"` + "\n"
	config += `	  servers = [{` + "\n"
	config += `		address = "1.1.1.1"` + "\n"
	config += `		authentication_port = 1812` + "\n"
	config += `		accounting_port = 1813` + "\n"
	config += `		timeout = 5` + "\n"
	config += `		retransmit = 3` + "\n"
	config += `		key = "key123"` + "\n"
	config += `		secret_key = "1234567"` + "\n"
	config += `		encryption_type = "7"` + "\n"
	config += `		key_type = "pac"` + "\n"
	config += `	}]` + "\n"
	config += `	}]` + "\n"
	config += `	radius_clients = [{` + "\n"
	config += `	  client_ip = "2.2.2.2"` + "\n"
	config += `	  vpn_configurations = [{` + "\n"
	config += `		vpn_id = 1` + "\n"
	config += `		server_key = "key123"` + "\n"
	config += `	}]` + "\n"
	config += `	}]` + "\n"
	config += `	radius_dynamic_author_server_key = "key123"` + "\n"
	config += `	radius_dynamic_author_domain_stripping = "yes"` + "\n"
	config += `	radius_dynamic_author_authentication_type = "all"` + "\n"
	config += `	radius_dynamic_author_port = 1700` + "\n"
	config += `	radius_trustsec_cts_authorization_list = "ALIST1"` + "\n"
	config += `	radius_trustsec_group = "GROUP1"` + "\n"
	config += `	tacacs_server_groups = [{` + "\n"
	config += `	  group_name = "GROUP1"` + "\n"
	config += `	  vpn_id = 1` + "\n"
	config += `	  source_interface = "e1"` + "\n"
	config += `	  servers = [{` + "\n"
	config += `		address = "1.1.1.1"` + "\n"
	config += `		port = 49` + "\n"
	config += `		timeout = 5` + "\n"
	config += `		key = "key123"` + "\n"
	config += `		secret_key = "1234567"` + "\n"
	config += `		encryption_type = "7"` + "\n"
	config += `	}]` + "\n"
	config += `	}]` + "\n"
	config += `	accounting_rules = [{` + "\n"
	config += `	  name = "RULE1"` + "\n"
	config += `	  method = "exec"` + "\n"
	config += `	  privilege_level = "15"` + "\n"
	config += `	  start_stop = true` + "\n"
	config += `	  groups = "GROUP1"` + "\n"
	config += `	}]` + "\n"
	config += `	authorization_console = true` + "\n"
	config += `	authorization_config_commands = true` + "\n"
	config += `	authorization_rules = [{` + "\n"
	config += `	  name = "RULE1"` + "\n"
	config += `	  method = "commands"` + "\n"
	config += `	  privilege_level = "15"` + "\n"
	config += `	  groups = "GROUP1"` + "\n"
	config += `	  authenticated = true` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigAll
