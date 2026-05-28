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
func TestAccSdwanSSEZscalerProfileParcel(t *testing.T) {
	if os.Getenv("SDWAN_2015") == "" {
		t.Skip("skipping test, set environment variable SDWAN_2015")
	}
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_sse_zscaler_feature.test", "interfaces.0.interface_name", "ipsec1"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_sse_zscaler_feature.test", "interfaces.0.auto", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_sse_zscaler_feature.test", "interfaces.0.tunnel_source_interface", "GigabitEthernet8"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_sse_zscaler_feature.test", "interfaces.0.tunnel_set", "secure-internet-gateway-zscaler"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_sse_zscaler_feature.test", "interfaces.0.tunnel_dc_preference", "primary-dc"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_sse_zscaler_feature.test", "interfaces.0.mtu", "1400"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_sse_zscaler_feature.test", "interfaces.0.ike_version", "2"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_sse_zscaler_feature.test", "interfaces.0.pre_shared_key_dynamic", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_sse_zscaler_feature.test", "interface_pairs.0.active_interface", "ipsec1"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_sse_zscaler_feature.test", "interface_pairs.0.active_interface_weight", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_sse_zscaler_feature.test", "interface_pairs.0.backup_interface", "None"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_sse_zscaler_feature.test", "interface_pairs.0.backup_interface_weight", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_sse_zscaler_feature.test", "refresh_time", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_sse_zscaler_feature.test", "refresh_time_unit", "MINUTE"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_sse_zscaler_feature.test", "sub_locations.0.name", "zscaler_sub1"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_sse_zscaler_feature.test", "sub_locations.0.ofw_enabled", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_sse_zscaler_feature.test", "sub_locations.0.caution_enabled", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_sse_zscaler_feature.test", "sub_locations.0.internal_ip.0.internal_ip_value", "172.16.2.1"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_sse_zscaler_feature.test", "sub_locations.0.aup_enabled", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_sse_zscaler_feature.test", "sub_locations.0.enforce_bandwidth_control", "disabled"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_sse_zscaler_feature.test", "tracker_source_ip", "1.2.3.4"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_sse_zscaler_feature.test", "trackers.0.name", "tracker1"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_sse_zscaler_feature.test", "trackers.0.endpoint_api_url", "http://cisco.com"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_sse_zscaler_feature.test", "trackers.0.threshold", "300"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_sse_zscaler_feature.test", "trackers.0.interval", "60"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_sse_zscaler_feature.test", "trackers.0.multiplier", "3"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccSdwanSSEZscalerPrerequisitesProfileParcelConfig + testAccSdwanSSEZscalerProfileParcelConfig_minimum(),
			},
			{
				Config: testAccSdwanSSEZscalerPrerequisitesProfileParcelConfig + testAccSdwanSSEZscalerProfileParcelConfig_all(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAcc

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites
const testAccSdwanSSEZscalerPrerequisitesProfileParcelConfig = `
resource "sdwan_sse_feature_profile" "test" {
  name        = "TF_TEST"
  description = "Terraform test"
}

`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigMinimum
func testAccSdwanSSEZscalerProfileParcelConfig_minimum() string {
	config := `resource "sdwan_sse_zscaler_feature" "test" {` + "\n"
	config += ` name = "TF_TEST_MIN"` + "\n"
	config += ` description = "Terraform integration test"` + "\n"
	config += `	feature_profile_id = sdwan_sse_feature_profile.test.id` + "\n"
	config += `	interfaces = [{` + "\n"
	config += `	  interface_name = "ipsec1"` + "\n"
	config += `	  tunnel_source_interface = "GigabitEthernet8"` + "\n"
	config += `	  tunnel_dc_preference = "primary-dc"` + "\n"
	config += `	  mtu = 1400` + "\n"
	config += `	}]` + "\n"
	config += `	interface_pairs = [{` + "\n"
	config += `	  active_interface = "ipsec1"` + "\n"
	config += `	  active_interface_weight = 1` + "\n"
	config += `	  backup_interface = "None"` + "\n"
	config += `	  backup_interface_weight = 1` + "\n"
	config += `	}]` + "\n"
	config += `	tracker_source_ip = "1.2.3.4"` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigMinimum

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAll
func testAccSdwanSSEZscalerProfileParcelConfig_all() string {
	config := `resource "sdwan_sse_zscaler_feature" "test" {` + "\n"
	config += ` name = "TF_TEST_ALL"` + "\n"
	config += ` description = "Terraform integration test"` + "\n"
	config += `	feature_profile_id = sdwan_sse_feature_profile.test.id` + "\n"
	config += `	interfaces = [{` + "\n"
	config += `	  interface_name = "ipsec1"` + "\n"
	config += `	  auto = true` + "\n"
	config += `	  tunnel_source_interface = "GigabitEthernet8"` + "\n"
	config += `	  tunnel_set = "secure-internet-gateway-zscaler"` + "\n"
	config += `	  tunnel_dc_preference = "primary-dc"` + "\n"
	config += `	  mtu = 1400` + "\n"
	config += `	  ike_version = 2` + "\n"
	config += `	  pre_shared_key_dynamic = true` + "\n"
	config += `	}]` + "\n"
	config += `	interface_pairs = [{` + "\n"
	config += `	  active_interface = "ipsec1"` + "\n"
	config += `	  active_interface_weight = 1` + "\n"
	config += `	  backup_interface = "None"` + "\n"
	config += `	  backup_interface_weight = 1` + "\n"
	config += `	}]` + "\n"
	config += `	refresh_time = 1` + "\n"
	config += `	refresh_time_unit = "MINUTE"` + "\n"
	config += `	sub_locations = [{` + "\n"
	config += `	  name = "zscaler_sub1"` + "\n"
	config += `	  ofw_enabled = false` + "\n"
	config += `	  caution_enabled = false` + "\n"
	config += `	  service_vpn = ["service_lan_vpn1,service_lan_vpn2"]` + "\n"
	config += `	  internal_ip = [{` + "\n"
	config += `		internal_ip_value = "172.16.2.1"` + "\n"
	config += `	}]` + "\n"
	config += `	  aup_enabled = false` + "\n"
	config += `	  enforce_bandwidth_control = "disabled"` + "\n"
	config += `	}]` + "\n"
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
