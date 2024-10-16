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
func TestAccDataSourceSdwanOtherThousandEyesProfileParcel(t *testing.T) {
	if os.Getenv("SDWAN_2012") == "" {
		t.Skip("skipping test, set environment variable SDWAN_2012")
	}
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_other_thousandeyes_feature.test", "virtual_application.0.account_group_token", "qwer"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_other_thousandeyes_feature.test", "virtual_application.0.vpn", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_other_thousandeyes_feature.test", "virtual_application.0.management_ip", "10.0.0.2"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_other_thousandeyes_feature.test", "virtual_application.0.management_subnet_mask", "255.255.255.0"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_other_thousandeyes_feature.test", "virtual_application.0.agent_default_gateway", "10.0.0.1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_other_thousandeyes_feature.test", "virtual_application.0.name_server_ip", "77.77.77.71"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_other_thousandeyes_feature.test", "virtual_application.0.hostname", "thousandeyesHost"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_other_thousandeyes_feature.test", "virtual_application.0.proxy_type", "static"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_other_thousandeyes_feature.test", "virtual_application.0.proxy_host", "proxy.thousandeyes.com"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_other_thousandeyes_feature.test", "virtual_application.0.proxy_port", "3128"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceSdwanOtherThousandEyesPrerequisitesProfileParcelConfig + testAccDataSourceSdwanOtherThousandEyesProfileParcelConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAccDataSource

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites
const testAccDataSourceSdwanOtherThousandEyesPrerequisitesProfileParcelConfig = `
resource "sdwan_other_feature_profile" "test" {
  name        = "TF_TEST"
  description = "Terraform test"
}
`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSourceConfig
func testAccDataSourceSdwanOtherThousandEyesProfileParcelConfig() string {
	config := `resource "sdwan_other_thousandeyes_feature" "test" {` + "\n"
	config += ` name = "TF_TEST"` + "\n"
	config += ` description = "Terraform integration test"` + "\n"
	config += `	feature_profile_id = sdwan_other_feature_profile.test.id` + "\n"
	config += `	virtual_application = [{` + "\n"
	config += `	  account_group_token = "qwer"` + "\n"
	config += `	  vpn = 1` + "\n"
	config += `	  management_ip = "10.0.0.2"` + "\n"
	config += `	  management_subnet_mask = "255.255.255.0"` + "\n"
	config += `	  agent_default_gateway = "10.0.0.1"` + "\n"
	config += `	  name_server_ip = "77.77.77.71"` + "\n"
	config += `	  hostname = "thousandeyesHost"` + "\n"
	config += `	  proxy_type = "static"` + "\n"
	config += `	  proxy_host = "proxy.thousandeyes.com"` + "\n"
	config += `	  proxy_port = 3128` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"

	config += `
		data "sdwan_other_thousandeyes_feature" "test" {
			id = sdwan_other_thousandeyes_feature.test.id
			feature_profile_id = sdwan_other_feature_profile.test.id
		}
	`
	return config
}

// End of section. //template:end testAccDataSourceConfig
