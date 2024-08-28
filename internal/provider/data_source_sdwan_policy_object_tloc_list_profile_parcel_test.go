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
func TestAccDataSourceSdwanPolicyObjectTLOCListProfileParcel(t *testing.T) {
	if os.Getenv("SDWAN_2012") == "" {
		t.Skip("skipping test, set environment variable SDWAN_2012")
	}
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_policy_object_tloc_list_profile_parcel.test", "entries.0.tloc_ip", "10.0.0.0"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_policy_object_tloc_list_profile_parcel.test", "entries.0.color", "3g"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_policy_object_tloc_list_profile_parcel.test", "entries.0.encapsulation", "gre"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_policy_object_tloc_list_profile_parcel.test", "entries.0.preference", "33"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceSdwanPolicyObjectTLOCListProfileParcelConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAccDataSource

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites
// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSourceConfig
func testAccDataSourceSdwanPolicyObjectTLOCListProfileParcelConfig() string {
	config := `resource "sdwan_policy_object_tloc_list_profile_parcel" "test" {` + "\n"
	config += ` name = "TF_TEST"` + "\n"
	config += ` description = "Terraform integration test"` + "\n"
	config += `	feature_profile_id = "f6dd22c8-0b4f-496c-9a0b-6813d1f8b8ac"` + "\n"
	config += `	entries = [{` + "\n"
	config += `	  tloc_ip = "10.0.0.0"` + "\n"
	config += `	  color = "3g"` + "\n"
	config += `	  encapsulation = "gre"` + "\n"
	config += `	  preference = "33"` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"

	config += `
		data "sdwan_policy_object_tloc_list_profile_parcel" "test" {
			id = sdwan_policy_object_tloc_list_profile_parcel.test.id
			feature_profile_id = "f6dd22c8-0b4f-496c-9a0b-6813d1f8b8ac"
		}
	`
	return config
}

// End of section. //template:end testAccDataSourceConfig
