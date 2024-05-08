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
func TestAccSdwanTLOCListPolicyObject(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_tloc_list_policy_object.test", "name", "Example"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_tloc_list_policy_object.test", "entries.0.tloc_ip", "1.1.1.2"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_tloc_list_policy_object.test", "entries.0.color", "blue"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_tloc_list_policy_object.test", "entries.0.encapsulation", "gre"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_tloc_list_policy_object.test", "entries.0.preference", "10"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccSdwanTLOCListPolicyObjectConfig_all(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAcc

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites
// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAll
func testAccSdwanTLOCListPolicyObjectConfig_all() string {
	config := `resource "sdwan_tloc_list_policy_object" "test" {` + "\n"
	config += `	name = "Example"` + "\n"
	config += `	entries = [{` + "\n"
	config += `	  tloc_ip = "1.1.1.2"` + "\n"
	config += `	  color = "blue"` + "\n"
	config += `	  encapsulation = "gre"` + "\n"
	config += `	  preference = 10` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigAll
