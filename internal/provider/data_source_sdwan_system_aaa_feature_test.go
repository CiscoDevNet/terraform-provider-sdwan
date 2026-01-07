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
func TestAccDataSourceSdwanSystemAAAProfileParcel(t *testing.T) {
	if os.Getenv("SDWAN_2015_IN_PROGRESS") == "" {
		t.Skip("skipping test, set environment variable SDWAN_2015_IN_PROGRESS")
	}
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_system_aaa_feature.test", "authentication_group", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_system_aaa_feature.test", "accounting_group", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_system_aaa_feature.test", "users.0.name", "User1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_system_aaa_feature.test", "users.0.privilege", "15"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_system_aaa_feature.test", "users.0.public_keys.0.key_string", "AAAAB3NzaC1yc2"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_system_aaa_feature.test", "users.0.public_keys.0.key_type", "ssh-rsa"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_system_aaa_feature.test", "radius_groups.0.group_name", "RGROUP1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_system_aaa_feature.test", "radius_groups.0.vpn", "10"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_system_aaa_feature.test", "radius_groups.0.source_interface", "GigabitEthernet0"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_system_aaa_feature.test", "radius_groups.0.servers.0.address", "1.2.3.4"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_system_aaa_feature.test", "radius_groups.0.servers.0.auth_port", "1812"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_system_aaa_feature.test", "radius_groups.0.servers.0.acct_port", "1813"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_system_aaa_feature.test", "radius_groups.0.servers.0.timeout", "5"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_system_aaa_feature.test", "radius_groups.0.servers.0.retransmit", "3"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_system_aaa_feature.test", "radius_groups.0.servers.0.secret_key", "cisco123"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_system_aaa_feature.test", "radius_groups.0.servers.0.key_enum", "7"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_system_aaa_feature.test", "radius_groups.0.servers.0.key_type", "key"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_system_aaa_feature.test", "tacacs_groups.0.group_name", "TGROUP1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_system_aaa_feature.test", "tacacs_groups.0.vpn", "10"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_system_aaa_feature.test", "tacacs_groups.0.source_interface", "GigabitEthernet0"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_system_aaa_feature.test", "tacacs_groups.0.servers.0.address", "1.2.3.4"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_system_aaa_feature.test", "tacacs_groups.0.servers.0.port", "49"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_system_aaa_feature.test", "tacacs_groups.0.servers.0.timeout", "5"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_system_aaa_feature.test", "tacacs_groups.0.servers.0.secret_key", "cisco123"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_system_aaa_feature.test", "tacacs_groups.0.servers.0.key_enum", "7"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_system_aaa_feature.test", "accounting_rules.0.rule_id", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_system_aaa_feature.test", "accounting_rules.0.method", "commands"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_system_aaa_feature.test", "accounting_rules.0.level", "15"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_system_aaa_feature.test", "accounting_rules.0.start_stop", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_system_aaa_feature.test", "authorization_console", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_system_aaa_feature.test", "authorization_config_commands", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_system_aaa_feature.test", "authorization_rules.0.rule_id", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_system_aaa_feature.test", "authorization_rules.0.method", "commands"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_system_aaa_feature.test", "authorization_rules.0.level", "15"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_system_aaa_feature.test", "authorization_rules.0.if_authenticated", "true"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceSdwanSystemAAAPrerequisitesProfileParcelConfig + testAccDataSourceSdwanSystemAAAProfileParcelConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAccDataSource

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites
const testAccDataSourceSdwanSystemAAAPrerequisitesProfileParcelConfig = `
resource "sdwan_system_feature_profile" "test" {
  name = "TF_TEST"
  description = "Terraform test"
}
`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSourceConfig
func testAccDataSourceSdwanSystemAAAProfileParcelConfig() string {
	config := `resource "sdwan_system_aaa_feature" "test" {` + "\n"
	config += ` name = "TF_TEST"` + "\n"
	config += ` description = "Terraform integration test"` + "\n"
	config += `	feature_profile_id = sdwan_system_feature_profile.test.id` + "\n"
	config += `	authentication_group = true` + "\n"
	config += `	accounting_group = true` + "\n"
	config += `	server_auth_order = ["local"]` + "\n"
	config += `	users = [{` + "\n"
	config += `	  name = "User1"` + "\n"
	config += `	  password = "cisco123"` + "\n"
	config += `	  privilege = "15"` + "\n"
	config += `	  public_keys = [{` + "\n"
	config += `		key_string = "AAAAB3NzaC1yc2"` + "\n"
	config += `		key_type = "ssh-rsa"` + "\n"
	config += `	}]` + "\n"
	config += `	}]` + "\n"
	config += `	radius_groups = [{` + "\n"
	config += `	  group_name = "RGROUP1"` + "\n"
	config += `	  vpn = 10` + "\n"
	config += `	  source_interface = "GigabitEthernet0"` + "\n"
	config += `	  servers = [{` + "\n"
	config += `		address = "1.2.3.4"` + "\n"
	config += `		auth_port = 1812` + "\n"
	config += `		acct_port = 1813` + "\n"
	config += `		timeout = 5` + "\n"
	config += `		retransmit = 3` + "\n"
	config += `		key = "cisco123"` + "\n"
	config += `		secret_key = "cisco123"` + "\n"
	config += `		key_enum = "7"` + "\n"
	config += `		key_type = "key"` + "\n"
	config += `	}]` + "\n"
	config += `	}]` + "\n"
	config += `	tacacs_groups = [{` + "\n"
	config += `	  group_name = "TGROUP1"` + "\n"
	config += `	  vpn = 10` + "\n"
	config += `	  source_interface = "GigabitEthernet0"` + "\n"
	config += `	  servers = [{` + "\n"
	config += `		address = "1.2.3.4"` + "\n"
	config += `		port = 49` + "\n"
	config += `		timeout = 5` + "\n"
	config += `		key = "cisco123"` + "\n"
	config += `		secret_key = "cisco123"` + "\n"
	config += `		key_enum = "7"` + "\n"
	config += `	}]` + "\n"
	config += `	}]` + "\n"
	config += `	accounting_rules = [{` + "\n"
	config += `	  rule_id = "1"` + "\n"
	config += `	  method = "commands"` + "\n"
	config += `	  level = "15"` + "\n"
	config += `	  start_stop = true` + "\n"
	config += `	  group = ["RGROUP1"]` + "\n"
	config += `	}]` + "\n"
	config += `	authorization_console = true` + "\n"
	config += `	authorization_config_commands = true` + "\n"
	config += `	authorization_rules = [{` + "\n"
	config += `	  rule_id = "1"` + "\n"
	config += `	  method = "commands"` + "\n"
	config += `	  level = "15"` + "\n"
	config += `	  group = ["RGROUP1"]` + "\n"
	config += `	  if_authenticated = true` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"

	config += `
		data "sdwan_system_aaa_feature" "test" {
			id = sdwan_system_aaa_feature.test.id
			feature_profile_id = sdwan_system_feature_profile.test.id
		}
	`
	return config
}

// End of section. //template:end testAccDataSourceConfig
