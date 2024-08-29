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
func TestAccDataSourceSdwanPolicyObjectSecurityIPsSignatureProfileParcel(t *testing.T) {
	if os.Getenv("SDWAN_2012") == "" && os.Getenv("TF_VAR_policy_object_feature_template_id") == "" {
		t.Skip("skipping test, set environment variable SDWAN_2012 or TF_VAR_policy_object_feature_template_id")
	}
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_policy_object_security_ips_signature_profile_parcel.test", "entries.0.generator_id", "1234"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_policy_object_security_ips_signature_profile_parcel.test", "entries.0.signature_id", "5678"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceSdwanPolicyObjectSecurityIPsSignaturePrerequisitesProfileParcelConfig + testAccDataSourceSdwanPolicyObjectSecurityIPsSignatureProfileParcelConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAccDataSource

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites
const testAccDataSourceSdwanPolicyObjectSecurityIPsSignaturePrerequisitesProfileParcelConfig = `
variable "policy_object_feature_template_id" {} 
`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSourceConfig
func testAccDataSourceSdwanPolicyObjectSecurityIPsSignatureProfileParcelConfig() string {
	config := `resource "sdwan_policy_object_security_ips_signature_profile_parcel" "test" {` + "\n"
	config += ` name = "TF_TEST"` + "\n"
	config += ` description = "Terraform integration test"` + "\n"
	config += `	feature_profile_id = var.policy_object_feature_template_id` + "\n"
	config += `	entries = [{` + "\n"
	config += `	  generator_id = "1234"` + "\n"
	config += `	  signature_id = "5678"` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"

	config += `
		data "sdwan_policy_object_security_ips_signature_profile_parcel" "test" {
			id = sdwan_policy_object_security_ips_signature_profile_parcel.test.id
			feature_profile_id = var.policy_object_feature_template_id
		}
	`
	return config
}

// End of section. //template:end testAccDataSourceConfig
