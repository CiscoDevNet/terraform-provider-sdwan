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

func TestAccDataSourceSdwanMirrorPolicyObject(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_mirror_policy_object.test", "name", "Example"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_mirror_policy_object.test", "remote_destination_ip", "10.1.1.1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_mirror_policy_object.test", "source_ip", "10.2.1.1"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceSdwanMirrorPolicyObjectConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

func testAccDataSourceSdwanMirrorPolicyObjectConfig() string {
	config := `resource "sdwan_mirror_policy_object" "test" {` + "\n"
	config += `	name = "Example"` + "\n"
	config += `	remote_destination_ip = "10.1.1.1"` + "\n"
	config += `	source_ip = "10.2.1.1"` + "\n"
	config += `}` + "\n"

	config += `
		data "sdwan_mirror_policy_object" "test" {
			id = sdwan_mirror_policy_object.test.id
		}
	`
	return config
}
