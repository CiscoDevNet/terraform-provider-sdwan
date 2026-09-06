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

func TestAccSdwanTag(t *testing.T) {
	if os.Getenv("SDWAN_2015") == "" && os.Getenv("SDWAN_2018") == "" {
		t.Skip("skipping test, set environment variable SDWAN_2015 or SDWAN_2018")
	}
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_tag.test", "name", "TAG_1"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_tag.test", "description", "My tag"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_tag.test", "devices.#", "2"))
	checks = append(checks, resource.TestCheckTypeSetElemAttr("sdwan_tag.test", "devices.*", "C8K-40C0CCFD-9EA8-2B2E-E73B-32C5924EC79B"))
	checks = append(checks, resource.TestCheckTypeSetElemAttr("sdwan_tag.test", "devices.*", "C8K-E94D7B88-4B9E-3323-C6C3-F29079FAAC3B"))

	// TAG_1 must keep both its devices after TAG_2 is created on one of them.
	var overlapChecks []resource.TestCheckFunc
	overlapChecks = append(overlapChecks, resource.TestCheckResourceAttr("sdwan_tag.test", "devices.#", "2"))
	overlapChecks = append(overlapChecks, resource.TestCheckTypeSetElemAttr("sdwan_tag.test", "devices.*", "C8K-40C0CCFD-9EA8-2B2E-E73B-32C5924EC79B"))
	overlapChecks = append(overlapChecks, resource.TestCheckTypeSetElemAttr("sdwan_tag.test", "devices.*", "C8K-E94D7B88-4B9E-3323-C6C3-F29079FAAC3B"))
	overlapChecks = append(overlapChecks, resource.TestCheckResourceAttr("sdwan_tag.test2", "name", "TAG_2"))
	overlapChecks = append(overlapChecks, resource.TestCheckResourceAttr("sdwan_tag.test2", "devices.#", "1"))
	overlapChecks = append(overlapChecks, resource.TestCheckTypeSetElemAttr("sdwan_tag.test2", "devices.*", "C8K-40C0CCFD-9EA8-2B2E-E73B-32C5924EC79B"))

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccSdwanTagConfig_all(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
			{
				// Adds TAG_2, sharing one device with TAG_1.
				Config: testAccSdwanTagConfig_all() + testAccSdwanTagConfig_overlap(),
				Check:  resource.ComposeTestCheckFunc(overlapChecks...),
			},
			{
				// Re-apply the identical config: no drift expected once both
				// tags coexist on the shared device.
				Config:   testAccSdwanTagConfig_all() + testAccSdwanTagConfig_overlap(),
				PlanOnly: true,
			},
		},
	})
}

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites
// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAll
func testAccSdwanTagConfig_all() string {
	config := `resource "sdwan_tag" "test" {` + "\n"
	config += `	name = "TAG_1"` + "\n"
	config += `	description = "My tag"` + "\n"
	config += `	devices = ["C8K-40C0CCFD-9EA8-2B2E-E73B-32C5924EC79B", "C8K-E94D7B88-4B9E-3323-C6C3-F29079FAAC3B"]` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigAll

// Second tag, sharing one device with sdwan_tag.test.
func testAccSdwanTagConfig_overlap() string {
	config := `resource "sdwan_tag" "test2" {` + "\n"
	config += `	name = "TAG_2"` + "\n"
	config += `	description = "My second tag"` + "\n"
	config += `	devices = ["C8K-40C0CCFD-9EA8-2B2E-E73B-32C5924EC79B"]` + "\n"
	config += `}` + "\n"
	return config
}
