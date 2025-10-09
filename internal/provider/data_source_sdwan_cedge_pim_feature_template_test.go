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
func TestAccDataSourceSdwanCEdgePIMFeatureTemplate(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cedge_pim_feature_template.test", "auto_rp", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cedge_pim_feature_template.test", "rp_announce_fields.0.interface_name", "Ethernet1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cedge_pim_feature_template.test", "rp_announce_fields.0.scope", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cedge_pim_feature_template.test", "interface_name", "Ethernet1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cedge_pim_feature_template.test", "rp_candidates.0.interface", "Ethernet1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cedge_pim_feature_template.test", "rp_candidates.0.access_list", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cedge_pim_feature_template.test", "rp_candidates.0.interval", "100"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cedge_pim_feature_template.test", "rp_candidates.0.priority", "2"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cedge_pim_feature_template.test", "bsr_candidate", "Ethernet1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cedge_pim_feature_template.test", "hash_mask_length", "24"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cedge_pim_feature_template.test", "priority", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cedge_pim_feature_template.test", "rp_candidate_access_list", "120"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cedge_pim_feature_template.test", "scope", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cedge_pim_feature_template.test", "range", "16"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cedge_pim_feature_template.test", "default", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cedge_pim_feature_template.test", "rp_addresses.0.access_list", "99"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cedge_pim_feature_template.test", "rp_addresses.0.ip_address", "1.2.3.4"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cedge_pim_feature_template.test", "rp_addresses.0.override", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cedge_pim_feature_template.test", "spt_threshold", "0"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cedge_pim_feature_template.test", "interfaces.0.interface_name", "Ethernet1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cedge_pim_feature_template.test", "interfaces.0.query_interval", "30"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cedge_pim_feature_template.test", "interfaces.0.join_prune_interval", "60"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceSdwanCEdgePIMFeatureTemplateConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAccDataSource

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSourceConfig
func testAccDataSourceSdwanCEdgePIMFeatureTemplateConfig() string {
	config := `resource "sdwan_cedge_pim_feature_template" "test" {` + "\n"
	config += ` name = "TF_TEST"` + "\n"
	config += ` description = "Terraform integration test"` + "\n"
	config += ` device_types = ["vedge-C8000V"]` + "\n"
	config += `	auto_rp = true` + "\n"
	config += `	rp_announce_fields = [{` + "\n"
	config += `	  interface_name = "Ethernet1"` + "\n"
	config += `	  scope = 1` + "\n"
	config += `	}]` + "\n"
	config += `	interface_name = "Ethernet1"` + "\n"
	config += `	rp_candidates = [{` + "\n"
	config += `	  interface = "Ethernet1"` + "\n"
	config += `	  access_list = "1"` + "\n"
	config += `	  interval = 100` + "\n"
	config += `	  priority = 2` + "\n"
	config += `	}]` + "\n"
	config += `	bsr_candidate = "Ethernet1"` + "\n"
	config += `	hash_mask_length = "24"` + "\n"
	config += `	priority = 1` + "\n"
	config += `	rp_candidate_access_list = "120"` + "\n"
	config += `	scope = 1` + "\n"
	config += `	range = "16"` + "\n"
	config += `	default = true` + "\n"
	config += `	rp_addresses = [{` + "\n"
	config += `	  access_list = "99"` + "\n"
	config += `	  ip_address = "1.2.3.4"` + "\n"
	config += `	  override = false` + "\n"
	config += `	}]` + "\n"
	config += `	spt_threshold = "0"` + "\n"
	config += `	interfaces = [{` + "\n"
	config += `	  interface_name = "Ethernet1"` + "\n"
	config += `	  query_interval = 30` + "\n"
	config += `	  join_prune_interval = 60` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"

	config += `
		data "sdwan_cedge_pim_feature_template" "test" {
			id = sdwan_cedge_pim_feature_template.test.id
		}
	`
	return config
}

// End of section. //template:end testAccDataSourceConfig
