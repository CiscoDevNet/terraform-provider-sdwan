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
func TestAccSdwanCiscoSecurityFeatureTemplate(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_security_feature_template.test", "rekey_interval", "86400"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_security_feature_template.test", "replay_window", "64"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_security_feature_template.test", "extended_ar_window", "256"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_security_feature_template.test", "pairwise_keying", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_security_feature_template.test", "keychains.0.name", "CHAIN1"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_security_feature_template.test", "keychains.0.key_id", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_security_feature_template.test", "keys.0.id", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_security_feature_template.test", "keys.0.chain_name", "CHAIN1"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_security_feature_template.test", "keys.0.send_id", "0"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_security_feature_template.test", "keys.0.receive_id", "0"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_security_feature_template.test", "keys.0.crypto_algorithm", "hmac-sha-256"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_security_feature_template.test", "keys.0.key_string", "abc123"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_security_feature_template.test", "keys.0.send_lifetime_local", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_security_feature_template.test", "keys.0.send_lifetime_start_time", "2022-12-31T23:59"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_security_feature_template.test", "keys.0.send_lifetime_end_time_format", "infinite"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_security_feature_template.test", "keys.0.send_lifetime_duration", "1000"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_security_feature_template.test", "keys.0.send_lifetime_end_time", "2032-12-31T23:59"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_security_feature_template.test", "keys.0.send_lifetime_infinite", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_security_feature_template.test", "keys.0.accept_lifetime_local", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_security_feature_template.test", "keys.0.accept_lifetime_start_time", "2022-12-31T23:59"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_security_feature_template.test", "keys.0.accept_lifetime_end_time_format", "infinite"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_security_feature_template.test", "keys.0.accept_lifetime_duration", "1000"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_security_feature_template.test", "keys.0.accept_lifetime_end_time", "2032-12-31T23:59"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_security_feature_template.test", "keys.0.accept_lifetime_infinite", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_security_feature_template.test", "keys.0.include_tcp_options", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_security_feature_template.test", "keys.0.accept_ao_mismatch", "true"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccSdwanCiscoSecurityFeatureTemplateConfig_minimum(),
			},
			{
				Config: testAccSdwanCiscoSecurityFeatureTemplateConfig_all(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAcc

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites
// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigMinimum
func testAccSdwanCiscoSecurityFeatureTemplateConfig_minimum() string {
	config := `resource "sdwan_cisco_security_feature_template" "test" {` + "\n"
	config += ` name = "TF_TEST_MIN"` + "\n"
	config += ` description = "Terraform integration test"` + "\n"
	config += ` device_types = ["vedge-C8000V"]` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigMinimum

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAll
func testAccSdwanCiscoSecurityFeatureTemplateConfig_all() string {
	config := `resource "sdwan_cisco_security_feature_template" "test" {` + "\n"
	config += ` name = "TF_TEST_ALL"` + "\n"
	config += ` description = "Terraform integration test"` + "\n"
	config += ` device_types = ["vedge-C8000V"]` + "\n"
	config += `	rekey_interval = 86400` + "\n"
	config += `	replay_window = "64"` + "\n"
	config += `	extended_ar_window = 256` + "\n"
	config += `	authentication_type = ["none"]` + "\n"
	config += `	integrity_type = ["none"]` + "\n"
	config += `	pairwise_keying = true` + "\n"
	config += `	keychains = [{` + "\n"
	config += `	  name = "CHAIN1"` + "\n"
	config += `	  key_id = 1` + "\n"
	config += `	}]` + "\n"
	config += `	keys = [{` + "\n"
	config += `	  id = "1"` + "\n"
	config += `	  chain_name = "CHAIN1"` + "\n"
	config += `	  send_id = 0` + "\n"
	config += `	  receive_id = 0` + "\n"
	config += `	  crypto_algorithm = "hmac-sha-256"` + "\n"
	config += `	  key_string = "abc123"` + "\n"
	config += `	  send_lifetime_local = true` + "\n"
	config += `	  send_lifetime_start_time = "2022-12-31T23:59"` + "\n"
	config += `	  send_lifetime_end_time_format = "infinite"` + "\n"
	config += `	  send_lifetime_duration = 1000` + "\n"
	config += `	  send_lifetime_end_time = "2032-12-31T23:59"` + "\n"
	config += `	  send_lifetime_infinite = true` + "\n"
	config += `	  accept_lifetime_local = true` + "\n"
	config += `	  accept_lifetime_start_time = "2022-12-31T23:59"` + "\n"
	config += `	  accept_lifetime_end_time_format = "infinite"` + "\n"
	config += `	  accept_lifetime_duration = 1000` + "\n"
	config += `	  accept_lifetime_end_time = "2032-12-31T23:59"` + "\n"
	config += `	  accept_lifetime_infinite = true` + "\n"
	config += `	  include_tcp_options = false` + "\n"
	config += `	  accept_ao_mismatch = true` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigAll
