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

func TestAccDataSourceSdwanCellularCEdgeProfileFeatureTemplate(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceSdwanCellularCEdgeProfileFeatureTemplateConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.sdwan_cellular_cedge_profile_feature_template.test", "profile_id", "1"),
					resource.TestCheckResourceAttr("data.sdwan_cellular_cedge_profile_feature_template.test", "access_point_name", "APN1"),
					resource.TestCheckResourceAttr("data.sdwan_cellular_cedge_profile_feature_template.test", "authentication_type", "chap"),
					resource.TestCheckResourceAttr("data.sdwan_cellular_cedge_profile_feature_template.test", "packet_data_network_type", "ipv4"),
					resource.TestCheckResourceAttr("data.sdwan_cellular_cedge_profile_feature_template.test", "profile_username", "MyUsername"),
					resource.TestCheckResourceAttr("data.sdwan_cellular_cedge_profile_feature_template.test", "profile_password", "MyPassword"),
					resource.TestCheckResourceAttr("data.sdwan_cellular_cedge_profile_feature_template.test", "no_overwrite", "false"),
				),
			},
		},
	})
}

const testAccDataSourceSdwanCellularCEdgeProfileFeatureTemplateConfig = `

resource "sdwan_cellular_cedge_profile_feature_template" "test" {
  name = "TF_TEST_MIN"
  description = "Terraform integration test"
  device_types = ["vedge-C8000V"]
  profile_id = 1
  access_point_name = "APN1"
  authentication_type = "chap"
  packet_data_network_type = "ipv4"
  profile_username = "MyUsername"
  profile_password = "MyPassword"
  no_overwrite = false
}

data "sdwan_cellular_cedge_profile_feature_template" "test" {
  id = sdwan_cellular_cedge_profile_feature_template.test.id
}
`
