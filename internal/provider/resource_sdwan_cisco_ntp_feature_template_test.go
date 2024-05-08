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
func TestAccSdwanCiscoNTPFeatureTemplate(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_ntp_feature_template.test", "master", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_ntp_feature_template.test", "master_stratum", "6"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_ntp_feature_template.test", "master_source_interface", "e1"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_ntp_feature_template.test", "authentication_keys.0.id", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_ntp_feature_template.test", "authentication_keys.0.value", "12345"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_ntp_feature_template.test", "servers.0.hostname_ip", "NTP_SERVER1"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_ntp_feature_template.test", "servers.0.authentication_key_id", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_ntp_feature_template.test", "servers.0.vpn_id", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_ntp_feature_template.test", "servers.0.version", "4"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_ntp_feature_template.test", "servers.0.source_interface", "e1"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cisco_ntp_feature_template.test", "servers.0.prefer", "true"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccSdwanCiscoNTPFeatureTemplateConfig_minimum(),
			},
			{
				Config: testAccSdwanCiscoNTPFeatureTemplateConfig_all(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAcc

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites
// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigMinimum
func testAccSdwanCiscoNTPFeatureTemplateConfig_minimum() string {
	config := `resource "sdwan_cisco_ntp_feature_template" "test" {` + "\n"
	config += ` name = "TF_TEST_MIN"` + "\n"
	config += ` description = "Terraform integration test"` + "\n"
	config += ` device_types = ["vedge-C8000V"]` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigMinimum

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAll
func testAccSdwanCiscoNTPFeatureTemplateConfig_all() string {
	config := `resource "sdwan_cisco_ntp_feature_template" "test" {` + "\n"
	config += ` name = "TF_TEST_ALL"` + "\n"
	config += ` description = "Terraform integration test"` + "\n"
	config += ` device_types = ["vedge-C8000V"]` + "\n"
	config += `	master = true` + "\n"
	config += `	master_stratum = 6` + "\n"
	config += `	master_source_interface = "e1"` + "\n"
	config += `	trusted_keys = [1]` + "\n"
	config += `	authentication_keys = [{` + "\n"
	config += `	  id = 1` + "\n"
	config += `	  value = "12345"` + "\n"
	config += `	}]` + "\n"
	config += `	servers = [{` + "\n"
	config += `	  hostname_ip = "NTP_SERVER1"` + "\n"
	config += `	  authentication_key_id = 1` + "\n"
	config += `	  vpn_id = 1` + "\n"
	config += `	  version = 4` + "\n"
	config += `	  source_interface = "e1"` + "\n"
	config += `	  prefer = true` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigAll
