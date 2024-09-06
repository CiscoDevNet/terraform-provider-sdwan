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
func TestAccDataSourceSdwanPolicyObjectIPv6PrefixListProfileParcel(t *testing.T) {
	if os.Getenv("SDWAN_2012") == "" {
		t.Skip("skipping test, set environment variable SDWAN_2012")
	}
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_policy_object_ipv6_prefix_list.test", "entries.0.ipv6_address", "2001:db8:85a3::8a2e:370:7334"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_policy_object_ipv6_prefix_list.test", "entries.0.ipv6_prefix_length", "64"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_policy_object_ipv6_prefix_list.test", "entries.0.le", "100"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_policy_object_ipv6_prefix_list.test", "entries.0.ge", "70"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceSdwanPolicyObjectIPv6PrefixListPrerequisitesProfileParcelConfig + testAccDataSourceSdwanPolicyObjectIPv6PrefixListProfileParcelConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAccDataSource

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites
const testAccDataSourceSdwanPolicyObjectIPv6PrefixListPrerequisitesProfileParcelConfig = `
data "sdwan_policy_object_feature_profile" "test" {}

`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSourceConfig
func testAccDataSourceSdwanPolicyObjectIPv6PrefixListProfileParcelConfig() string {
	config := `resource "sdwan_policy_object_ipv6_prefix_list" "test" {` + "\n"
	config += ` name = "TF_TEST"` + "\n"
	config += ` description = "Terraform integration test"` + "\n"
	config += `	feature_profile_id = data.sdwan_policy_object_feature_profile.test.id` + "\n"
	config += `	entries = [{` + "\n"
	config += `	  ipv6_address = "2001:db8:85a3::8a2e:370:7334"` + "\n"
	config += `	  ipv6_prefix_length = 64` + "\n"
	config += `	  le = 100` + "\n"
	config += `	  ge = 70` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"

	config += `
		data "sdwan_policy_object_ipv6_prefix_list" "test" {
			id = sdwan_policy_object_ipv6_prefix_list.test.id
			feature_profile_id = data.sdwan_policy_object_feature_profile.test.id
		}
	`
	return config
}

// End of section. //template:end testAccDataSourceConfig