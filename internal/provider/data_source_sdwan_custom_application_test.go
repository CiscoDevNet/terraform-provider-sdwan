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
func TestAccDataSourceSdwanCustomApplication(t *testing.T) {
	if os.Getenv("SDWAN_2015") == "" {
		t.Skip("skipping test, set environment variable SDWAN_2015")
	}
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_custom_application.test", "app_name", "Example-Custom"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_custom_application.test", "l3l4.0.l4_protocol", "TCP"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_custom_application.test", "l3l4.0.ports", "1 10-20"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_custom_application.test", "application_family", "routing"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_custom_application.test", "application_group", "ipsec-group"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_custom_application.test", "traffic_class", "signaling"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_custom_application.test", "business_relevance", "business-relevant"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceSdwanCustomApplicationConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAccDataSource

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites
// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSourceConfig
func testAccDataSourceSdwanCustomApplicationConfig() string {
	config := ""
	config += `resource "sdwan_custom_application" "test" {` + "\n"
	config += `	app_name = "Example-Custom"` + "\n"
	config += `	server_names = ["*customapp.com", "customapptest.com", "*appcustom"]` + "\n"
	config += `	l3l4 = [{` + "\n"
	config += `	  ip_addresses = ["10.2.2.25", "192.168.1.0/24"]` + "\n"
	config += `	  l4_protocol = "TCP"` + "\n"
	config += `	  ports = "1 10-20"` + "\n"
	config += `	}]` + "\n"
	config += `	application_family = "routing"` + "\n"
	config += `	application_group = "ipsec-group"` + "\n"
	config += `	traffic_class = "signaling"` + "\n"
	config += `	business_relevance = "business-relevant"` + "\n"
	config += `}` + "\n"

	config += `
		data "sdwan_custom_application" "test" {
			id = sdwan_custom_application.test.id
		}
	`
	return config
}

// End of section. //template:end testAccDataSourceConfig
