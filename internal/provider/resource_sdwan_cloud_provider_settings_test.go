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
func TestAccSdwanCloudProviderSettings(t *testing.T) {
	if os.Getenv("SDWAN_2015") == "" {
		t.Skip("skipping test, set environment variable SDWAN_2015")
	}
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cloud_provider_settings.test", "umbrella_org_id", "123456"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cloud_provider_settings.test", "umbrella_auth_key_v2", "aaaaaa"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cloud_provider_settings.test", "umbrella_sig_auth_key", "bbbbbb"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cloud_provider_settings.test", "umbrella_dns_auth_key", "cccccc"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cloud_provider_settings.test", "zscaler_organization", "z223456"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cloud_provider_settings.test", "zscaler_partner_base_uri", "zscaler.net"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cloud_provider_settings.test", "zscaler_partner_key", "secret012"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cloud_provider_settings.test", "zscaler_username", "user1"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cloud_provider_settings.test", "cisco_sse_org_id", "323456"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cloud_provider_settings.test", "cisco_sse_auth_key", "33333444"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cloud_provider_settings.test", "cisco_sse_context_sharing", "true"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccSdwanCloudProviderSettingsConfig_all(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAcc

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites
// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAll
func testAccSdwanCloudProviderSettingsConfig_all() string {
	config := `resource "sdwan_cloud_provider_settings" "test" {` + "\n"
	config += `	umbrella_org_id = "123456"` + "\n"
	config += `	umbrella_auth_key_v2 = "aaaaaa"` + "\n"
	config += `	umbrella_auth_secret_v2 = "$CRYPT_CLUSTER$a8nXr4uYaCN66qTQ1737+A==$gjmYVxnKMqvUEjo3BfIGbg=="` + "\n"
	config += `	umbrella_sig_auth_key = "bbbbbb"` + "\n"
	config += `	umbrella_sig_auth_secret = "secret456"` + "\n"
	config += `	umbrella_dns_auth_key = "cccccc"` + "\n"
	config += `	umbrella_dns_auth_secret = "secret789"` + "\n"
	config += `	zscaler_organization = "z223456"` + "\n"
	config += `	zscaler_partner_base_uri = "zscaler.net"` + "\n"
	config += `	zscaler_partner_key = "secret012"` + "\n"
	config += `	zscaler_username = "user1"` + "\n"
	config += `	zscaler_password = "pass123"` + "\n"
	config += `	cisco_sse_org_id = "323456"` + "\n"
	config += `	cisco_sse_auth_key = "33333444"` + "\n"
	config += `	cisco_sse_auth_secret = "$CRYPT_CLUSTER$Gg4nVpFdldXga1hLKhdJrA==$hiFPirWJnqNxMq3l/m1ekw=="` + "\n"
	config += `	cisco_sse_context_sharing = true` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigAll
