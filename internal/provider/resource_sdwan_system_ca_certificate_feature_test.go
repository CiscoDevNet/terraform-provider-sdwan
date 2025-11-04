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

// Section below is generated&owned by "gen/generator.go". //template:begin testAcc
func TestAccSdwanSystemCACertificateProfileParcel(t *testing.T) {
	if os.Getenv("SDWAN_2015") == "" && os.Getenv("CA_CERT") == "" {
		t.Skip("skipping test, set environment variable SDWAN_2015 or CA_CERT")
	}
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_system_ca_certificate_feature.test", "certificates.0.trust_point_name", "example"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_system_ca_certificate_feature.test", "certificates.0.ca_certificate_id", "22a2f715-fed8-4031-b693-b5d43451a05e"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccSdwanSystemCACertificatePrerequisitesProfileParcelConfig + testAccSdwanSystemCACertificateProfileParcelConfig_minimum(),
			},
			{
				Config: testAccSdwanSystemCACertificatePrerequisitesProfileParcelConfig + testAccSdwanSystemCACertificateProfileParcelConfig_all(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAcc

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites
const testAccSdwanSystemCACertificatePrerequisitesProfileParcelConfig = `
resource "sdwan_system_feature_profile" "test" {
  name = "TF_TEST"
  description = "Terraform test"
}
`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigMinimum
func testAccSdwanSystemCACertificateProfileParcelConfig_minimum() string {
	config := `resource "sdwan_system_ca_certificate_feature" "test" {` + "\n"
	config += ` name = "TF_TEST_MIN"` + "\n"
	config += ` description = "Terraform integration test"` + "\n"
	config += `	feature_profile_id = sdwan_system_feature_profile.test.id` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigMinimum

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAll
func testAccSdwanSystemCACertificateProfileParcelConfig_all() string {
	config := `resource "sdwan_system_ca_certificate_feature" "test" {` + "\n"
	config += ` name = "TF_TEST_ALL"` + "\n"
	config += ` description = "Terraform integration test"` + "\n"
	config += `	feature_profile_id = sdwan_system_feature_profile.test.id` + "\n"
	config += `	certificates = [{` + "\n"
	config += `	  trust_point_name = "example"` + "\n"
	config += `	  ca_certificate_id = "22a2f715-fed8-4031-b693-b5d43451a05e"` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigAll
