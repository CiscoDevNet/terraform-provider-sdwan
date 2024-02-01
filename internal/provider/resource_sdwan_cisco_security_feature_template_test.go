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

func TestAccSdwanCiscoSecurityFeatureTemplate(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccSdwanCiscoSecurityFeatureTemplateConfig_minimum(),
				Check:  resource.ComposeTestCheckFunc(),
			},
			{
				Config: testAccSdwanCiscoSecurityFeatureTemplateConfig_all(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("sdwan_cisco_security_feature_template.test", "rekey_interval", "86400"),
					resource.TestCheckResourceAttr("sdwan_cisco_security_feature_template.test", "replay_window", "64"),
					resource.TestCheckResourceAttr("sdwan_cisco_security_feature_template.test", "extended_ar_window", "256"),
					resource.TestCheckResourceAttr("sdwan_cisco_security_feature_template.test", "pairwise_keying", "true"),
					resource.TestCheckResourceAttr("sdwan_cisco_security_feature_template.test", "keychains.0.name", "CHAIN1"),
					resource.TestCheckResourceAttr("sdwan_cisco_security_feature_template.test", "keychains.0.key_id", "1"),
					resource.TestCheckResourceAttr("sdwan_cisco_security_feature_template.test", "keys.0.id", "1"),
					resource.TestCheckResourceAttr("sdwan_cisco_security_feature_template.test", "keys.0.chain_name", "CHAIN1"),
					resource.TestCheckResourceAttr("sdwan_cisco_security_feature_template.test", "keys.0.send_id", "0"),
					resource.TestCheckResourceAttr("sdwan_cisco_security_feature_template.test", "keys.0.receive_id", "0"),
					resource.TestCheckResourceAttr("sdwan_cisco_security_feature_template.test", "keys.0.crypto_algorithm", "hmac-sha-256"),
					resource.TestCheckResourceAttr("sdwan_cisco_security_feature_template.test", "keys.0.key_string", "abc123"),
					resource.TestCheckResourceAttr("sdwan_cisco_security_feature_template.test", "keys.0.send_lifetime_local", "true"),
					resource.TestCheckResourceAttr("sdwan_cisco_security_feature_template.test", "keys.0.send_lifetime_start_time", "2022-12-31T23:59"),
					resource.TestCheckResourceAttr("sdwan_cisco_security_feature_template.test", "keys.0.send_lifetime_end_time_format", "infinite"),
					resource.TestCheckResourceAttr("sdwan_cisco_security_feature_template.test", "keys.0.send_lifetime_duration", "1000"),
					resource.TestCheckResourceAttr("sdwan_cisco_security_feature_template.test", "keys.0.send_lifetime_end_time", "2032-12-31T23:59"),
					resource.TestCheckResourceAttr("sdwan_cisco_security_feature_template.test", "keys.0.send_lifetime_infinite", "true"),
					resource.TestCheckResourceAttr("sdwan_cisco_security_feature_template.test", "keys.0.accept_lifetime_local", "true"),
					resource.TestCheckResourceAttr("sdwan_cisco_security_feature_template.test", "keys.0.accept_lifetime_start_time", "2022-12-31T23:59"),
					resource.TestCheckResourceAttr("sdwan_cisco_security_feature_template.test", "keys.0.accept_lifetime_end_time_format", "infinite"),
					resource.TestCheckResourceAttr("sdwan_cisco_security_feature_template.test", "keys.0.accept_lifetime_duration", "1000"),
					resource.TestCheckResourceAttr("sdwan_cisco_security_feature_template.test", "keys.0.accept_lifetime_end_time", "2032-12-31T23:59"),
					resource.TestCheckResourceAttr("sdwan_cisco_security_feature_template.test", "keys.0.accept_lifetime_infinite", "true"),
					resource.TestCheckResourceAttr("sdwan_cisco_security_feature_template.test", "keys.0.include_tcp_options", "false"),
					resource.TestCheckResourceAttr("sdwan_cisco_security_feature_template.test", "keys.0.accept_ao_mismatch", "true"),
				),
			},
		},
	})
}

func testAccSdwanCiscoSecurityFeatureTemplateConfig_minimum() string {
	return `
	resource "sdwan_cisco_security_feature_template" "test" {
		name = "TF_TEST_MIN"
		description = "Terraform integration test"
		device_types = ["vedge-C8000V"]
	}
	`
}

func testAccSdwanCiscoSecurityFeatureTemplateConfig_all() string {
	return `
	resource "sdwan_cisco_security_feature_template" "test" {
		name = "TF_TEST_ALL"
		description = "Terraform integration test"
		device_types = ["vedge-C8000V"]
		rekey_interval = 86400
		replay_window = "64"
		extended_ar_window = 256
		authentication_type = ["none"]
		integrity_type = ["none"]
		pairwise_keying = true
		keychains = [{
			name = "CHAIN1"
			key_id = 1
		}]
		keys = [{
			id = "1"
			chain_name = "CHAIN1"
			send_id = 0
			receive_id = 0
			crypto_algorithm = "hmac-sha-256"
			key_string = "abc123"
			send_lifetime_local = true
			send_lifetime_start_time = "2022-12-31T23:59"
			send_lifetime_end_time_format = "infinite"
			send_lifetime_duration = 1000
			send_lifetime_end_time = "2032-12-31T23:59"
			send_lifetime_infinite = true
			accept_lifetime_local = true
			accept_lifetime_start_time = "2022-12-31T23:59"
			accept_lifetime_end_time_format = "infinite"
			accept_lifetime_duration = 1000
			accept_lifetime_end_time = "2032-12-31T23:59"
			accept_lifetime_infinite = true
			include_tcp_options = false
			accept_ao_mismatch = true
		}]
	}
	`
}
