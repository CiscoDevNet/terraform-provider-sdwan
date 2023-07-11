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

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccDataSourceSdwanLocalizedPolicy(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceSdwanLocalizedPolicyConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.sdwan_localized_policy.test", "name", "TF_TEST_ALL"),
					resource.TestCheckResourceAttr("data.sdwan_localized_policy.test", "description", "Terraform integration test"),
					resource.TestCheckResourceAttr("data.sdwan_localized_policy.test", "flow_visibility_ipv4", "true"),
				),
			},
		},
	})
}

const testAccDataSourceSdwanLocalizedPolicyConfig = `

resource "sdwan_localized_policy" "test" {
	name = "TF_TEST_ALL"
	description = "Terraform integration test"
	flow_visibility_ipv4 = true
}

data "sdwan_localized_policy" "test" {
  id = sdwan_localized_policy.test.id
}
`
