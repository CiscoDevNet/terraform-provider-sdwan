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
func TestAccSdwanSSECiscoProfileParcel(t *testing.T) {
	if os.Getenv("SDWAN_2015") == "" {
		t.Skip("skipping test, set environment variable SDWAN_2015")
	}
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_sse_cisco_feature.test", "context_sharing_for_vpn", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_sse_cisco_feature.test", "context_sharing_for_sgt", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_sse_cisco_feature.test", "interfaces.0.interface_name", "ipsec1"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_sse_cisco_feature.test", "interfaces.0.tunnel_dc_preference", "primary-dc"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_sse_cisco_feature.test", "interfaces.0.mtu", "1400"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_sse_cisco_feature.test", "interfaces.0.ike_version", "2"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_sse_cisco_feature.test", "interface_pairs.0.active_interface", "ipsec1"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_sse_cisco_feature.test", "interface_pairs.0.active_interface_weight", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_sse_cisco_feature.test", "interface_pairs.0.backup_interface", "None"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_sse_cisco_feature.test", "interface_pairs.0.backup_interface_weight", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_sse_cisco_feature.test", "region", "us-east-1"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_sse_cisco_feature.test", "tracker_source_ip", "1.2.3.4"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_sse_cisco_feature.test", "trackers.0.name", "tracker1"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_sse_cisco_feature.test", "trackers.0.endpoint_api_url", "http://cisco.com"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_sse_cisco_feature.test", "trackers.0.threshold", "300"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_sse_cisco_feature.test", "trackers.0.interval", "60"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_sse_cisco_feature.test", "trackers.0.multiplier", "3"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccSdwanSSECiscoPrerequisitesProfileParcelConfig + testAccSdwanSSECiscoProfileParcelConfig_minimum(),
			},
			{
				Config: testAccSdwanSSECiscoPrerequisitesProfileParcelConfig + testAccSdwanSSECiscoProfileParcelConfig_all(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAcc

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites
const testAccSdwanSSECiscoPrerequisitesProfileParcelConfig = `
resource "sdwan_sse_feature_profile" "test" {
  name        = "TF_TEST"
  description = "Terraform test"
}
`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigMinimum
func testAccSdwanSSECiscoProfileParcelConfig_minimum() string {
	config := `resource "sdwan_sse_cisco_feature" "test" {` + "\n"
	config += ` name = "TF_TEST_MIN"` + "\n"
	config += ` description = "Terraform integration test"` + "\n"
	config += `	feature_profile_id = sdwan_sse_feature_profile.test.id` + "\n"
	config += `	interfaces = [{` + "\n"
	config += `	  interface_name = "ipsec1"` + "\n"
	config += `	  tunnel_dc_preference = "primary-dc"` + "\n"
	config += `	  mtu = 1400` + "\n"
	config += `	}]` + "\n"
	config += `	interface_pairs = [{` + "\n"
	config += `	  active_interface = "ipsec1"` + "\n"
	config += `	  active_interface_weight = 1` + "\n"
	config += `	  backup_interface = "None"` + "\n"
	config += `	  backup_interface_weight = 1` + "\n"
	config += `	}]` + "\n"
	config += `	region = "us-east-1"` + "\n"
	config += `	tracker_source_ip = "1.2.3.4"` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigMinimum

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAll
func testAccSdwanSSECiscoProfileParcelConfig_all() string {
	config := `resource "sdwan_sse_cisco_feature" "test" {` + "\n"
	config += ` name = "TF_TEST_ALL"` + "\n"
	config += ` description = "Terraform integration test"` + "\n"
	config += `	feature_profile_id = sdwan_sse_feature_profile.test.id` + "\n"
	config += `	context_sharing_for_vpn = false` + "\n"
	config += `	context_sharing_for_sgt = false` + "\n"
	config += `	interfaces = [{` + "\n"
	config += `	  interface_name = "ipsec1"` + "\n"
	config += `	  tunnel_dc_preference = "primary-dc"` + "\n"
	config += `	  mtu = 1400` + "\n"
	config += `	  ike_version = 2` + "\n"
	config += `	}]` + "\n"
	config += `	interface_pairs = [{` + "\n"
	config += `	  active_interface = "ipsec1"` + "\n"
	config += `	  active_interface_weight = 1` + "\n"
	config += `	  backup_interface = "None"` + "\n"
	config += `	  backup_interface_weight = 1` + "\n"
	config += `	}]` + "\n"
	config += `	region = "us-east-1"` + "\n"
	config += `	tracker_source_ip = "1.2.3.4"` + "\n"
	config += `	trackers = [{` + "\n"
	config += `	  name = "tracker1"` + "\n"
	config += `	  endpoint_api_url = "http://cisco.com"` + "\n"
	config += `	  threshold = 300` + "\n"
	config += `	  interval = 60` + "\n"
	config += `	  multiplier = 3` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigAll
